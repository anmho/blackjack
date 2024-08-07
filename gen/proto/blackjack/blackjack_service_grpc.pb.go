// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: proto/blackjack/blackjack_service.proto

package pb_blackjack

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BlackJackServiceClient is the client API for BlackJackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlackJackServiceClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (BlackJackService_ConnectClient, error)
}

type blackJackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlackJackServiceClient(cc grpc.ClientConnInterface) BlackJackServiceClient {
	return &blackJackServiceClient{cc}
}

func (c *blackJackServiceClient) Connect(ctx context.Context, opts ...grpc.CallOption) (BlackJackService_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlackJackService_ServiceDesc.Streams[0], "/blackjack.BlackJackService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &blackJackServiceConnectClient{stream}
	return x, nil
}

type BlackJackService_ConnectClient interface {
	Send(*BlackjackRequest) error
	Recv() (*BlackjackResponse, error)
	grpc.ClientStream
}

type blackJackServiceConnectClient struct {
	grpc.ClientStream
}

func (x *blackJackServiceConnectClient) Send(m *BlackjackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *blackJackServiceConnectClient) Recv() (*BlackjackResponse, error) {
	m := new(BlackjackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlackJackServiceServer is the server API for BlackJackService service.
// All implementations must embed UnimplementedBlackJackServiceServer
// for forward compatibility
type BlackJackServiceServer interface {
	Connect(BlackJackService_ConnectServer) error
	mustEmbedUnimplementedBlackJackServiceServer()
}

// UnimplementedBlackJackServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlackJackServiceServer struct {
}

func (UnimplementedBlackJackServiceServer) Connect(BlackJackService_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedBlackJackServiceServer) mustEmbedUnimplementedBlackJackServiceServer() {}

// UnsafeBlackJackServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlackJackServiceServer will
// result in compilation errors.
type UnsafeBlackJackServiceServer interface {
	mustEmbedUnimplementedBlackJackServiceServer()
}

func RegisterBlackJackServiceServer(s grpc.ServiceRegistrar, srv BlackJackServiceServer) {
	s.RegisterService(&BlackJackService_ServiceDesc, srv)
}

func _BlackJackService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BlackJackServiceServer).Connect(&blackJackServiceConnectServer{stream})
}

type BlackJackService_ConnectServer interface {
	Send(*BlackjackResponse) error
	Recv() (*BlackjackRequest, error)
	grpc.ServerStream
}

type blackJackServiceConnectServer struct {
	grpc.ServerStream
}

func (x *blackJackServiceConnectServer) Send(m *BlackjackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *blackJackServiceConnectServer) Recv() (*BlackjackRequest, error) {
	m := new(BlackjackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlackJackService_ServiceDesc is the grpc.ServiceDesc for BlackJackService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlackJackService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blackjack.BlackJackService",
	HandlerType: (*BlackJackServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _BlackJackService_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/blackjack/blackjack_service.proto",
}
