// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: proto/blackjack/blackjack.proto

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
	ViewGames(ctx context.Context, in *ViewGamesRequest, opts ...grpc.CallOption) (*ViewGamesResponse, error)
	PlayGame(ctx context.Context, opts ...grpc.CallOption) (BlackJackService_PlayGameClient, error)
}

type blackJackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlackJackServiceClient(cc grpc.ClientConnInterface) BlackJackServiceClient {
	return &blackJackServiceClient{cc}
}

func (c *blackJackServiceClient) ViewGames(ctx context.Context, in *ViewGamesRequest, opts ...grpc.CallOption) (*ViewGamesResponse, error) {
	out := new(ViewGamesResponse)
	err := c.cc.Invoke(ctx, "/blackjack.BlackJackService/ViewGames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blackJackServiceClient) PlayGame(ctx context.Context, opts ...grpc.CallOption) (BlackJackService_PlayGameClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlackJackService_ServiceDesc.Streams[0], "/blackjack.BlackJackService/PlayGame", opts...)
	if err != nil {
		return nil, err
	}
	x := &blackJackServicePlayGameClient{stream}
	return x, nil
}

type BlackJackService_PlayGameClient interface {
	Send(*BlackjackRequest) error
	Recv() (*BlackjackResponse, error)
	grpc.ClientStream
}

type blackJackServicePlayGameClient struct {
	grpc.ClientStream
}

func (x *blackJackServicePlayGameClient) Send(m *BlackjackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *blackJackServicePlayGameClient) Recv() (*BlackjackResponse, error) {
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
	ViewGames(context.Context, *ViewGamesRequest) (*ViewGamesResponse, error)
	PlayGame(BlackJackService_PlayGameServer) error
	mustEmbedUnimplementedBlackJackServiceServer()
}

// UnimplementedBlackJackServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlackJackServiceServer struct {
}

func (UnimplementedBlackJackServiceServer) ViewGames(context.Context, *ViewGamesRequest) (*ViewGamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewGames not implemented")
}
func (UnimplementedBlackJackServiceServer) PlayGame(BlackJackService_PlayGameServer) error {
	return status.Errorf(codes.Unimplemented, "method PlayGame not implemented")
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

func _BlackJackService_ViewGames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewGamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlackJackServiceServer).ViewGames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blackjack.BlackJackService/ViewGames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlackJackServiceServer).ViewGames(ctx, req.(*ViewGamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlackJackService_PlayGame_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BlackJackServiceServer).PlayGame(&blackJackServicePlayGameServer{stream})
}

type BlackJackService_PlayGameServer interface {
	Send(*BlackjackResponse) error
	Recv() (*BlackjackRequest, error)
	grpc.ServerStream
}

type blackJackServicePlayGameServer struct {
	grpc.ServerStream
}

func (x *blackJackServicePlayGameServer) Send(m *BlackjackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *blackJackServicePlayGameServer) Recv() (*BlackjackRequest, error) {
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
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ViewGames",
			Handler:    _BlackJackService_ViewGames_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PlayGame",
			Handler:       _BlackJackService_PlayGame_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/blackjack/blackjack.proto",
}
