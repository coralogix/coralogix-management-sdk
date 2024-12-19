// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: com/coralogixapis/database_catalog/v1/blueprints_service.proto

package v1

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
	BlueprintsService_GetBlueprintsColumnsStream_FullMethodName = "/com.coralogixapis.database_catalog.v1.BlueprintsService/GetBlueprintsColumnsStream"
	BlueprintsService_GetBlueprintStats_FullMethodName          = "/com.coralogixapis.database_catalog.v1.BlueprintsService/GetBlueprintStats"
	BlueprintsService_GetBlueprintGraphsStream_FullMethodName   = "/com.coralogixapis.database_catalog.v1.BlueprintsService/GetBlueprintGraphsStream"
)

// BlueprintsServiceClient is the client API for BlueprintsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlueprintsServiceClient interface {
	GetBlueprintsColumnsStream(ctx context.Context, in *GetBlueprintsColumnsStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetBlueprintsColumnsStreamResponse], error)
	GetBlueprintStats(ctx context.Context, in *GetBlueprintStatsRequest, opts ...grpc.CallOption) (*GetBlueprintStatsResponse, error)
	GetBlueprintGraphsStream(ctx context.Context, in *GetBlueprintGraphsStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetBlueprintGraphsStreamResponse], error)
}

type blueprintsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlueprintsServiceClient(cc grpc.ClientConnInterface) BlueprintsServiceClient {
	return &blueprintsServiceClient{cc}
}

func (c *blueprintsServiceClient) GetBlueprintsColumnsStream(ctx context.Context, in *GetBlueprintsColumnsStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetBlueprintsColumnsStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BlueprintsService_ServiceDesc.Streams[0], BlueprintsService_GetBlueprintsColumnsStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetBlueprintsColumnsStreamRequest, GetBlueprintsColumnsStreamResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BlueprintsService_GetBlueprintsColumnsStreamClient = grpc.ServerStreamingClient[GetBlueprintsColumnsStreamResponse]

func (c *blueprintsServiceClient) GetBlueprintStats(ctx context.Context, in *GetBlueprintStatsRequest, opts ...grpc.CallOption) (*GetBlueprintStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBlueprintStatsResponse)
	err := c.cc.Invoke(ctx, BlueprintsService_GetBlueprintStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blueprintsServiceClient) GetBlueprintGraphsStream(ctx context.Context, in *GetBlueprintGraphsStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetBlueprintGraphsStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &BlueprintsService_ServiceDesc.Streams[1], BlueprintsService_GetBlueprintGraphsStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetBlueprintGraphsStreamRequest, GetBlueprintGraphsStreamResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BlueprintsService_GetBlueprintGraphsStreamClient = grpc.ServerStreamingClient[GetBlueprintGraphsStreamResponse]

// BlueprintsServiceServer is the server API for BlueprintsService service.
// All implementations must embed UnimplementedBlueprintsServiceServer
// for forward compatibility.
type BlueprintsServiceServer interface {
	GetBlueprintsColumnsStream(*GetBlueprintsColumnsStreamRequest, grpc.ServerStreamingServer[GetBlueprintsColumnsStreamResponse]) error
	GetBlueprintStats(context.Context, *GetBlueprintStatsRequest) (*GetBlueprintStatsResponse, error)
	GetBlueprintGraphsStream(*GetBlueprintGraphsStreamRequest, grpc.ServerStreamingServer[GetBlueprintGraphsStreamResponse]) error
	mustEmbedUnimplementedBlueprintsServiceServer()
}

// UnimplementedBlueprintsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBlueprintsServiceServer struct{}

func (UnimplementedBlueprintsServiceServer) GetBlueprintsColumnsStream(*GetBlueprintsColumnsStreamRequest, grpc.ServerStreamingServer[GetBlueprintsColumnsStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetBlueprintsColumnsStream not implemented")
}
func (UnimplementedBlueprintsServiceServer) GetBlueprintStats(context.Context, *GetBlueprintStatsRequest) (*GetBlueprintStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlueprintStats not implemented")
}
func (UnimplementedBlueprintsServiceServer) GetBlueprintGraphsStream(*GetBlueprintGraphsStreamRequest, grpc.ServerStreamingServer[GetBlueprintGraphsStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetBlueprintGraphsStream not implemented")
}
func (UnimplementedBlueprintsServiceServer) mustEmbedUnimplementedBlueprintsServiceServer() {}
func (UnimplementedBlueprintsServiceServer) testEmbeddedByValue()                           {}

// UnsafeBlueprintsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlueprintsServiceServer will
// result in compilation errors.
type UnsafeBlueprintsServiceServer interface {
	mustEmbedUnimplementedBlueprintsServiceServer()
}

func RegisterBlueprintsServiceServer(s grpc.ServiceRegistrar, srv BlueprintsServiceServer) {
	// If the following call pancis, it indicates UnimplementedBlueprintsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BlueprintsService_ServiceDesc, srv)
}

func _BlueprintsService_GetBlueprintsColumnsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBlueprintsColumnsStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlueprintsServiceServer).GetBlueprintsColumnsStream(m, &grpc.GenericServerStream[GetBlueprintsColumnsStreamRequest, GetBlueprintsColumnsStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BlueprintsService_GetBlueprintsColumnsStreamServer = grpc.ServerStreamingServer[GetBlueprintsColumnsStreamResponse]

func _BlueprintsService_GetBlueprintStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlueprintStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlueprintsServiceServer).GetBlueprintStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlueprintsService_GetBlueprintStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlueprintsServiceServer).GetBlueprintStats(ctx, req.(*GetBlueprintStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlueprintsService_GetBlueprintGraphsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBlueprintGraphsStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlueprintsServiceServer).GetBlueprintGraphsStream(m, &grpc.GenericServerStream[GetBlueprintGraphsStreamRequest, GetBlueprintGraphsStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type BlueprintsService_GetBlueprintGraphsStreamServer = grpc.ServerStreamingServer[GetBlueprintGraphsStreamResponse]

// BlueprintsService_ServiceDesc is the grpc.ServiceDesc for BlueprintsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlueprintsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.database_catalog.v1.BlueprintsService",
	HandlerType: (*BlueprintsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlueprintStats",
			Handler:    _BlueprintsService_GetBlueprintStats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlueprintsColumnsStream",
			Handler:       _BlueprintsService_GetBlueprintsColumnsStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetBlueprintGraphsStream",
			Handler:       _BlueprintsService_GetBlueprintGraphsStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "com/coralogixapis/database_catalog/v1/blueprints_service.proto",
}
