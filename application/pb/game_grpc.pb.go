// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: application/proto/game.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GameService_CreateGame_FullMethodName                    = "/pb.GameService/CreateGame"
	GameService_CreateGameStream_FullMethodName              = "/pb.GameService/CreateGameStream"
	GameService_CreateGameStreamBidirectional_FullMethodName = "/pb.GameService/CreateGameStreamBidirectional"
	GameService_ListAllGames_FullMethodName                  = "/pb.GameService/ListAllGames"
	GameService_GetGame_FullMethodName                       = "/pb.GameService/GetGame"
)

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameServiceClient interface {
	CreateGame(ctx context.Context, in *CreateGameInput, opts ...grpc.CallOption) (*BaseGameOutput, error)
	CreateGameStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[CreateGameInput, ListAllGamesOutput], error)
	CreateGameStreamBidirectional(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[CreateGameInput, BaseGameOutput], error)
	ListAllGames(ctx context.Context, in *BlankInput, opts ...grpc.CallOption) (*ListAllGamesOutput, error)
	GetGame(ctx context.Context, in *GetGameInput, opts ...grpc.CallOption) (*BaseGameOutput, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) CreateGame(ctx context.Context, in *CreateGameInput, opts ...grpc.CallOption) (*BaseGameOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseGameOutput)
	err := c.cc.Invoke(ctx, GameService_CreateGame_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) CreateGameStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[CreateGameInput, ListAllGamesOutput], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GameService_ServiceDesc.Streams[0], GameService_CreateGameStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CreateGameInput, ListAllGamesOutput]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GameService_CreateGameStreamClient = grpc.ClientStreamingClient[CreateGameInput, ListAllGamesOutput]

func (c *gameServiceClient) CreateGameStreamBidirectional(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[CreateGameInput, BaseGameOutput], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GameService_ServiceDesc.Streams[1], GameService_CreateGameStreamBidirectional_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CreateGameInput, BaseGameOutput]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GameService_CreateGameStreamBidirectionalClient = grpc.BidiStreamingClient[CreateGameInput, BaseGameOutput]

func (c *gameServiceClient) ListAllGames(ctx context.Context, in *BlankInput, opts ...grpc.CallOption) (*ListAllGamesOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAllGamesOutput)
	err := c.cc.Invoke(ctx, GameService_ListAllGames_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetGame(ctx context.Context, in *GetGameInput, opts ...grpc.CallOption) (*BaseGameOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseGameOutput)
	err := c.cc.Invoke(ctx, GameService_GetGame_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
// All implementations must embed UnimplementedGameServiceServer
// for forward compatibility.
type GameServiceServer interface {
	CreateGame(context.Context, *CreateGameInput) (*BaseGameOutput, error)
	CreateGameStream(grpc.ClientStreamingServer[CreateGameInput, ListAllGamesOutput]) error
	CreateGameStreamBidirectional(grpc.BidiStreamingServer[CreateGameInput, BaseGameOutput]) error
	ListAllGames(context.Context, *BlankInput) (*ListAllGamesOutput, error)
	GetGame(context.Context, *GetGameInput) (*BaseGameOutput, error)
	mustEmbedUnimplementedGameServiceServer()
}

// UnimplementedGameServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGameServiceServer struct{}

func (UnimplementedGameServiceServer) CreateGame(context.Context, *CreateGameInput) (*BaseGameOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (UnimplementedGameServiceServer) CreateGameStream(grpc.ClientStreamingServer[CreateGameInput, ListAllGamesOutput]) error {
	return status.Errorf(codes.Unimplemented, "method CreateGameStream not implemented")
}
func (UnimplementedGameServiceServer) CreateGameStreamBidirectional(grpc.BidiStreamingServer[CreateGameInput, BaseGameOutput]) error {
	return status.Errorf(codes.Unimplemented, "method CreateGameStreamBidirectional not implemented")
}
func (UnimplementedGameServiceServer) ListAllGames(context.Context, *BlankInput) (*ListAllGamesOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllGames not implemented")
}
func (UnimplementedGameServiceServer) GetGame(context.Context, *GetGameInput) (*BaseGameOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGame not implemented")
}
func (UnimplementedGameServiceServer) mustEmbedUnimplementedGameServiceServer() {}
func (UnimplementedGameServiceServer) testEmbeddedByValue()                     {}

// UnsafeGameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServiceServer will
// result in compilation errors.
type UnsafeGameServiceServer interface {
	mustEmbedUnimplementedGameServiceServer()
}

func RegisterGameServiceServer(s grpc.ServiceRegistrar, srv GameServiceServer) {
	// If the following call pancis, it indicates UnimplementedGameServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GameService_ServiceDesc, srv)
}

func _GameService_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_CreateGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).CreateGame(ctx, req.(*CreateGameInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_CreateGameStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServiceServer).CreateGameStream(&grpc.GenericServerStream[CreateGameInput, ListAllGamesOutput]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GameService_CreateGameStreamServer = grpc.ClientStreamingServer[CreateGameInput, ListAllGamesOutput]

func _GameService_CreateGameStreamBidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServiceServer).CreateGameStreamBidirectional(&grpc.GenericServerStream[CreateGameInput, BaseGameOutput]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GameService_CreateGameStreamBidirectionalServer = grpc.BidiStreamingServer[CreateGameInput, BaseGameOutput]

func _GameService_ListAllGames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlankInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).ListAllGames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_ListAllGames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).ListAllGames(ctx, req.(*BlankInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_GetGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetGame(ctx, req.(*GetGameInput))
	}
	return interceptor(ctx, in, info, handler)
}

// GameService_ServiceDesc is the grpc.ServiceDesc for GameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGame",
			Handler:    _GameService_CreateGame_Handler,
		},
		{
			MethodName: "ListAllGames",
			Handler:    _GameService_ListAllGames_Handler,
		},
		{
			MethodName: "GetGame",
			Handler:    _GameService_GetGame_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateGameStream",
			Handler:       _GameService_CreateGameStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateGameStreamBidirectional",
			Handler:       _GameService_CreateGameStreamBidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "application/proto/game.proto",
}
