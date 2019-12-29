package server

import (
	"context"
	"errors"
	"net"

	pb "github.com/marcboeker/tentacle/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type connDetails struct {
	nodeID   int
	peerAddr string
}

type apiServer struct {
	db     *DB
	connCh chan connDetails
}

type serverStreamWrapper struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *serverStreamWrapper) Context() context.Context { return w.ctx }

type mdKey string

var (
	tokenKey = mdKey("token")
	streams  = map[int]pb.Tentacle_WaitForConnectionServer{}
	peers    = map[int]string{}
)

// Auth interceptor for unary calls.
func authInterceptor(db *DB) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errUnauthorized
		}
		token, ok := md["token"]
		if !ok || len(token) == 0 {
			return nil, errUnauthorized
		}

		node, err := authenticate(db, token[0])
		if err != nil {
			return nil, errUnauthorized
		}

		ctx = context.WithValue(ctx, tokenKey, node)
		return handler(ctx, req)
	}
}

// Auth interceptor for stream calls.
func streamAuthInterceptor(db *DB) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := stream.Context()
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return errUnauthorized
		}
		token, ok := md["token"]
		if !ok || len(token) == 0 {
			return errUnauthorized
		}

		node, err := authenticate(db, token[0])
		if err != nil {
			return errUnauthorized
		}

		ctx = context.WithValue(ctx, tokenKey, node)
		ss := &serverStreamWrapper{ServerStream: stream, ctx: ctx}

		return handler(srv, ss)
	}
}

func authenticate(db *DB, token string) (*Node, error) {
	var node Node
	if err := db.Where("token = ?", token).First(&node).Error; err != nil {
		return nil, errUnauthorized
	}
	return &node, nil
}

// Register takes the remote address and port from the peer and writes it
// to the peer register.
func (s apiServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.Empty, error) {
	node := ctx.Value(tokenKey).(*Node)

	// Build remote address from remote IP and local port.
	peer, _ := peer.FromContext(ctx)
	host, _, _ := net.SplitHostPort(peer.Addr.String())
	peers[node.ID] = net.JoinHostPort(host, req.LocalPort)

	return &pb.Empty{}, nil
}

// GetSubnets returns all subnets for the peer.
func (s apiServer) GetSubnets(ctx context.Context, empty *pb.Empty) (*pb.GetSubnetsResponse, error) {
	node := ctx.Value(tokenKey).(*Node)

	// Find all subnet IDs a node belongs to.
	var subnetIDs []int
	s.db.Model(&SubnetNode{}).Where("node_id = ?", node.ID).Pluck("subnet_id", &subnetIDs)
	if len(subnetIDs) == 0 {
		return nil, errors.New("could not find active subnets for node")
	}

	// Retrieve all subnets for a node.
	var subnets []Subnet
	if err := s.db.Where("id IN (?)", subnetIDs).Find(&subnets).Error; err != nil {
		return nil, errors.New("could not find node's subnets")
	}

	//Find all nodes that are in the same subnet as the given node.
	var res []*pb.GetSubnetsResponse_Subnet
	for _, sn := range subnets {
		// Get all other peers in a subnet.
		var localIP string
		peers := []string{}
		var nodes []SubnetNode
		if err := s.db.Model(&SubnetNode{}).Where("subnet_id = ?", sn.ID).Find(&nodes).Error; err != nil {
			return nil, errors.New("could not find subnets")
		}

		for _, n := range nodes {
			if n.NodeID == node.ID {
				localIP = n.IP
			} else {
				peers = append(peers, n.IP)
			}
		}

		subnet := pb.GetSubnetsResponse_Subnet{
			Name:    sn.Name,
			Cidr:    sn.CIDR,
			Ip:      localIP,
			PeerIps: peers,
		}
		res = append(res, &subnet)
	}

	return &pb.GetSubnetsResponse{Subnets: res}, nil
}

// Connect looks up the peer's remote address by its local address.
func (s apiServer) Connect(ctx context.Context, req *pb.ConnectRequest) (*pb.ConnectResponse, error) {
	node := ctx.Value(tokenKey).(*Node)

	var subnetIDs []int
	s.db.Model(&SubnetNode{}).
		Where("node_id = ?", node.ID).
		Pluck("subnet_id", &subnetIDs)
	if len(subnetIDs) == 0 {
		return nil, errors.New("could not find active subnets for node")
	}

	var subnetNode SubnetNode
	if err := s.db.Model(&SubnetNode{}).
		Where("subnet_id IN (?) AND ip = ?", subnetIDs, req.GetPeerIp()).
		First(&subnetNode).
		Error; err != nil {
		return nil, err
	}

	remoteAddr, ok := peers[subnetNode.NodeID]
	if !ok {
		return nil, errors.New("peer is offline")
	}

	peerAddr, ok := peers[node.ID]
	if !ok {
		return nil, errors.New("peer is offline")
	}

	s.connCh <- connDetails{nodeID: subnetNode.NodeID, peerAddr: peerAddr}

	return &pb.ConnectResponse{PeerAddress: remoteAddr}, nil
}

// WaitForConnection waits until a peer wants another peer to establish a
// connection and performs a peer address exchange.
func (s apiServer) WaitForConnection(req *pb.Empty, stream pb.Tentacle_WaitForConnectionServer) error {
	node := stream.Context().Value(tokenKey).(*Node)
	streams[node.ID] = stream

	for {
		select {
		case cd := <-s.connCh:
			s, ok := streams[cd.nodeID]
			if !ok {
				break
			}
			cr := &pb.ConnectResponse{PeerAddress: cd.peerAddr}
			if err := s.Send(cr); err != nil {
				return err
			}
		}
	}
}

// StartAPIServer starts an UDP server that listens on the given address.
func StartAPIServer(listenAddr string, db *DB) error {
	as := apiServer{db: db, connCh: make(chan connDetails, 1)}
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor(db)),
		grpc.StreamInterceptor(streamAuthInterceptor(db)),
	)
	pb.RegisterTentacleServer(srv, &as)
	return srv.Serve(l)
}

var (
	errUnauthorized = errors.New("unauthorized")
)
