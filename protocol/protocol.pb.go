// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

package protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type RegisterRequest struct {
	LocalPort            string   `protobuf:"bytes,1,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{1}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetLocalPort() string {
	if m != nil {
		return m.LocalPort
	}
	return ""
}

type ConnectRequest struct {
	PeerIp               string   `protobuf:"bytes,1,opt,name=peer_ip,json=peerIp,proto3" json:"peer_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{2}
}

func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetPeerIp() string {
	if m != nil {
		return m.PeerIp
	}
	return ""
}

type ConnectResponse struct {
	PeerAddress          string   `protobuf:"bytes,1,opt,name=peer_address,json=peerAddress,proto3" json:"peer_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectResponse) Reset()         { *m = ConnectResponse{} }
func (m *ConnectResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectResponse) ProtoMessage()    {}
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{3}
}

func (m *ConnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectResponse.Unmarshal(m, b)
}
func (m *ConnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectResponse.Marshal(b, m, deterministic)
}
func (m *ConnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectResponse.Merge(m, src)
}
func (m *ConnectResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectResponse.Size(m)
}
func (m *ConnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectResponse proto.InternalMessageInfo

func (m *ConnectResponse) GetPeerAddress() string {
	if m != nil {
		return m.PeerAddress
	}
	return ""
}

type GetSubnetsResponse struct {
	Subnets              []*GetSubnetsResponse_Subnet `protobuf:"bytes,1,rep,name=subnets,proto3" json:"subnets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetSubnetsResponse) Reset()         { *m = GetSubnetsResponse{} }
func (m *GetSubnetsResponse) String() string { return proto.CompactTextString(m) }
func (*GetSubnetsResponse) ProtoMessage()    {}
func (*GetSubnetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{4}
}

func (m *GetSubnetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSubnetsResponse.Unmarshal(m, b)
}
func (m *GetSubnetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSubnetsResponse.Marshal(b, m, deterministic)
}
func (m *GetSubnetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSubnetsResponse.Merge(m, src)
}
func (m *GetSubnetsResponse) XXX_Size() int {
	return xxx_messageInfo_GetSubnetsResponse.Size(m)
}
func (m *GetSubnetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSubnetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSubnetsResponse proto.InternalMessageInfo

func (m *GetSubnetsResponse) GetSubnets() []*GetSubnetsResponse_Subnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

type GetSubnetsResponse_Subnet struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cidr                 string   `protobuf:"bytes,2,opt,name=cidr,proto3" json:"cidr,omitempty"`
	Ip                   string   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	PeerIps              []string `protobuf:"bytes,4,rep,name=peer_ips,json=peerIps,proto3" json:"peer_ips,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSubnetsResponse_Subnet) Reset()         { *m = GetSubnetsResponse_Subnet{} }
func (m *GetSubnetsResponse_Subnet) String() string { return proto.CompactTextString(m) }
func (*GetSubnetsResponse_Subnet) ProtoMessage()    {}
func (*GetSubnetsResponse_Subnet) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{4, 0}
}

func (m *GetSubnetsResponse_Subnet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSubnetsResponse_Subnet.Unmarshal(m, b)
}
func (m *GetSubnetsResponse_Subnet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSubnetsResponse_Subnet.Marshal(b, m, deterministic)
}
func (m *GetSubnetsResponse_Subnet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSubnetsResponse_Subnet.Merge(m, src)
}
func (m *GetSubnetsResponse_Subnet) XXX_Size() int {
	return xxx_messageInfo_GetSubnetsResponse_Subnet.Size(m)
}
func (m *GetSubnetsResponse_Subnet) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSubnetsResponse_Subnet.DiscardUnknown(m)
}

var xxx_messageInfo_GetSubnetsResponse_Subnet proto.InternalMessageInfo

func (m *GetSubnetsResponse_Subnet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetSubnetsResponse_Subnet) GetCidr() string {
	if m != nil {
		return m.Cidr
	}
	return ""
}

func (m *GetSubnetsResponse_Subnet) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *GetSubnetsResponse_Subnet) GetPeerIps() []string {
	if m != nil {
		return m.PeerIps
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "protocol.Empty")
	proto.RegisterType((*RegisterRequest)(nil), "protocol.RegisterRequest")
	proto.RegisterType((*ConnectRequest)(nil), "protocol.ConnectRequest")
	proto.RegisterType((*ConnectResponse)(nil), "protocol.ConnectResponse")
	proto.RegisterType((*GetSubnetsResponse)(nil), "protocol.GetSubnetsResponse")
	proto.RegisterType((*GetSubnetsResponse_Subnet)(nil), "protocol.GetSubnetsResponse.Subnet")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_2bc2336598a3f7e0) }

var fileDescriptor_2bc2336598a3f7e0 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4e, 0xe3, 0x30,
	0x10, 0xc6, 0x9b, 0xb4, 0xdb, 0xa4, 0xd3, 0x55, 0xab, 0x9d, 0xcb, 0xa6, 0xd1, 0xae, 0x54, 0xc2,
	0xa5, 0x5c, 0xaa, 0xaa, 0x70, 0xe0, 0x52, 0x09, 0x54, 0x01, 0xe2, 0x86, 0x02, 0x12, 0x07, 0x0e,
	0x55, 0x9a, 0x8e, 0x90, 0xa5, 0xd4, 0x36, 0xb6, 0x7b, 0xe0, 0x15, 0x78, 0x1a, 0x1e, 0x11, 0x25,
	0xce, 0x1f, 0xd1, 0xa2, 0xde, 0x66, 0x7e, 0xf3, 0x7d, 0xb6, 0xbf, 0x31, 0x0c, 0xa4, 0x12, 0x46,
	0xa4, 0x22, 0x9b, 0x16, 0x05, 0xfa, 0x55, 0x1f, 0x79, 0xf0, 0xeb, 0x66, 0x2b, 0xcd, 0x7b, 0x34,
	0x83, 0x61, 0x4c, 0xaf, 0x4c, 0x1b, 0x52, 0x31, 0xbd, 0xed, 0x48, 0x1b, 0xfc, 0x0f, 0x90, 0x89,
	0x34, 0xc9, 0x56, 0x52, 0x28, 0x13, 0x38, 0x63, 0x67, 0xd2, 0x8b, 0x7b, 0x05, 0x79, 0x10, 0xca,
	0x44, 0x67, 0x30, 0x58, 0x0a, 0xce, 0x29, 0x35, 0x95, 0xe1, 0x2f, 0x78, 0x92, 0x48, 0xad, 0x98,
	0x2c, 0xd5, 0xdd, 0xbc, 0xbd, 0x97, 0xd1, 0x05, 0x0c, 0x6b, 0xa9, 0x96, 0x82, 0x6b, 0xc2, 0x13,
	0xf8, 0x5d, 0x68, 0x93, 0xcd, 0x46, 0x91, 0xd6, 0xa5, 0xa1, 0x9f, 0xb3, 0x6b, 0x8b, 0xa2, 0x4f,
	0x07, 0xf0, 0x8e, 0xcc, 0xe3, 0x6e, 0xcd, 0xc9, 0xe8, 0xda, 0xb9, 0x00, 0x4f, 0x5b, 0x14, 0x38,
	0xe3, 0xf6, 0xa4, 0x3f, 0x3f, 0x9d, 0xd6, 0xf1, 0x0e, 0xe5, 0x53, 0xdb, 0xc7, 0x95, 0x27, 0x7c,
	0x81, 0xae, 0x45, 0x88, 0xd0, 0xe1, 0xc9, 0x96, 0xca, 0xab, 0x8b, 0x3a, 0x67, 0x29, 0xdb, 0xa8,
	0xc0, 0xb5, 0x2c, 0xaf, 0x71, 0x00, 0x2e, 0x93, 0x41, 0xbb, 0x20, 0x2e, 0x93, 0x38, 0x02, 0xbf,
	0x8c, 0xa9, 0x83, 0xce, 0xb8, 0x3d, 0xe9, 0xc5, 0x9e, 0xcd, 0xa9, 0xe7, 0x1f, 0x2e, 0xf8, 0x4f,
	0xc4, 0x4d, 0x92, 0x66, 0x84, 0x97, 0xe0, 0x57, 0x2b, 0xc5, 0x51, 0xf3, 0xc6, 0xbd, 0x35, 0x87,
	0xc3, 0x66, 0x64, 0xbf, 0xa2, 0x85, 0x0b, 0x80, 0x26, 0x09, 0xee, 0x0b, 0xc2, 0x7f, 0xc7, 0x02,
	0x47, 0x2d, 0xbc, 0x02, 0xaf, 0x5c, 0x37, 0x06, 0x8d, 0xf4, 0xfb, 0x67, 0x85, 0xa3, 0x1f, 0x26,
	0xf5, 0x09, 0x4b, 0xf8, 0xf3, 0x9c, 0x30, 0x73, 0x2b, 0x54, 0x39, 0x63, 0x82, 0x1f, 0xbe, 0xe3,
	0xd8, 0x11, 0x33, 0x67, 0xdd, 0x2d, 0xa6, 0xe7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x97,
	0x7e, 0xa5, 0x7e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TentacleClient is the client API for Tentacle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TentacleClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Empty, error)
	GetSubnets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetSubnetsResponse, error)
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	WaitForConnection(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Tentacle_WaitForConnectionClient, error)
}

type tentacleClient struct {
	cc *grpc.ClientConn
}

func NewTentacleClient(cc *grpc.ClientConn) TentacleClient {
	return &tentacleClient{cc}
}

func (c *tentacleClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protocol.Tentacle/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tentacleClient) GetSubnets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetSubnetsResponse, error) {
	out := new(GetSubnetsResponse)
	err := c.cc.Invoke(ctx, "/protocol.Tentacle/GetSubnets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tentacleClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, "/protocol.Tentacle/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tentacleClient) WaitForConnection(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Tentacle_WaitForConnectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tentacle_serviceDesc.Streams[0], "/protocol.Tentacle/WaitForConnection", opts...)
	if err != nil {
		return nil, err
	}
	x := &tentacleWaitForConnectionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tentacle_WaitForConnectionClient interface {
	Recv() (*ConnectResponse, error)
	grpc.ClientStream
}

type tentacleWaitForConnectionClient struct {
	grpc.ClientStream
}

func (x *tentacleWaitForConnectionClient) Recv() (*ConnectResponse, error) {
	m := new(ConnectResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TentacleServer is the server API for Tentacle service.
type TentacleServer interface {
	Register(context.Context, *RegisterRequest) (*Empty, error)
	GetSubnets(context.Context, *Empty) (*GetSubnetsResponse, error)
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	WaitForConnection(*Empty, Tentacle_WaitForConnectionServer) error
}

// UnimplementedTentacleServer can be embedded to have forward compatible implementations.
type UnimplementedTentacleServer struct {
}

func (*UnimplementedTentacleServer) Register(ctx context.Context, req *RegisterRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedTentacleServer) GetSubnets(ctx context.Context, req *Empty) (*GetSubnetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubnets not implemented")
}
func (*UnimplementedTentacleServer) Connect(ctx context.Context, req *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedTentacleServer) WaitForConnection(req *Empty, srv Tentacle_WaitForConnectionServer) error {
	return status.Errorf(codes.Unimplemented, "method WaitForConnection not implemented")
}

func RegisterTentacleServer(s *grpc.Server, srv TentacleServer) {
	s.RegisterService(&_Tentacle_serviceDesc, srv)
}

func _Tentacle_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TentacleServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Tentacle/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TentacleServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tentacle_GetSubnets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TentacleServer).GetSubnets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Tentacle/GetSubnets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TentacleServer).GetSubnets(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tentacle_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TentacleServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Tentacle/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TentacleServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tentacle_WaitForConnection_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TentacleServer).WaitForConnection(m, &tentacleWaitForConnectionServer{stream})
}

type Tentacle_WaitForConnectionServer interface {
	Send(*ConnectResponse) error
	grpc.ServerStream
}

type tentacleWaitForConnectionServer struct {
	grpc.ServerStream
}

func (x *tentacleWaitForConnectionServer) Send(m *ConnectResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Tentacle_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Tentacle",
	HandlerType: (*TentacleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Tentacle_Register_Handler,
		},
		{
			MethodName: "GetSubnets",
			Handler:    _Tentacle_GetSubnets_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _Tentacle_Connect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WaitForConnection",
			Handler:       _Tentacle_WaitForConnection_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protocol.proto",
}
