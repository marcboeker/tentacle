package client

import (
	"context"
	"net"
	"time"

	"github.com/marcboeker/tentacle/protocol"
	pb "github.com/marcboeker/tentacle/protocol"
	"google.golang.org/grpc"
)

// RPC represents an RPC client.
type RPC struct {
	token   string
	srvAddr *net.UDPAddr
	client  protocol.TentacleClient
}

type tokenCreds struct{ token string }

func (c *tokenCreds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"token": c.token}, nil
}

func (c *tokenCreds) RequireTransportSecurity() bool {
	return false
}

// NewRPC returns an instance of a GRPC client connected to the remote server.
func NewRPC(token, srvAddr string) (*RPC, error) {
	c := &RPC{token: token}

	conn, err := grpc.Dial(
		srvAddr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(&tokenCreds{token: token}),
	)
	if err != nil {
		return nil, err
	}
	c.client = pb.NewTentacleClient(conn)

	return c, nil
}

// Register registers a peer with its local port.
func (r *RPC) Register(port string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := r.client.Register(ctx, &pb.RegisterRequest{LocalPort: port})
	return err
}

// Subnets gets a list of all subnets for the current peer.
func (r *RPC) Subnets() ([]*pb.GetSubnetsResponse_Subnet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := r.client.GetSubnets(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}
	return res.Subnets, nil
}

// Connect resolves the local peer address into its public address.
func (r *RPC) Connect(peerIP string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := r.client.Connect(ctx, &pb.ConnectRequest{PeerIp: peerIP})
	if err != nil {
		return "", err
	}
	return res.PeerAddress, nil
}

// WaitForConnection waits for am incoming connection requests and starts
// the NAT hole punching process.
func (r *RPC) WaitForConnection(connectCB func(peerAddr string) error) error {
	stream, err := r.client.WaitForConnection(context.Background(), &pb.Empty{})
	if err != nil {
		return err
	}
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		if err := connectCB(msg.GetPeerAddress()); err != nil {
			return err
		}
	}
}
