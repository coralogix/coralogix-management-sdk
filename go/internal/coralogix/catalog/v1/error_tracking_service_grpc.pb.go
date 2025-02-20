// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: com/coralogixapis/service_catalog/v1/error-tracking/error_tracking_service.proto

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
	ErrorTrackingService_GetServiceErrorTrackingOverviewStream_FullMethodName = "/com.coralogixapis.service_catalog.v1.ErrorTrackingService/GetServiceErrorTrackingOverviewStream"
	ErrorTrackingService_GetServiceErrorTrackingListStream_FullMethodName     = "/com.coralogixapis.service_catalog.v1.ErrorTrackingService/GetServiceErrorTrackingListStream"
)

// ErrorTrackingServiceClient is the client API for ErrorTrackingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ErrorTrackingServiceClient interface {
	GetServiceErrorTrackingOverviewStream(ctx context.Context, in *GetServiceErrorTrackingOverviewRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetServiceErrorTrackingOverviewResponse], error)
	GetServiceErrorTrackingListStream(ctx context.Context, in *GetServiceErrorTrackingListRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetServiceErrorTrackingListResponse], error)
}

type errorTrackingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewErrorTrackingServiceClient(cc grpc.ClientConnInterface) ErrorTrackingServiceClient {
	return &errorTrackingServiceClient{cc}
}

func (c *errorTrackingServiceClient) GetServiceErrorTrackingOverviewStream(ctx context.Context, in *GetServiceErrorTrackingOverviewRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetServiceErrorTrackingOverviewResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ErrorTrackingService_ServiceDesc.Streams[0], ErrorTrackingService_GetServiceErrorTrackingOverviewStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetServiceErrorTrackingOverviewRequest, GetServiceErrorTrackingOverviewResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ErrorTrackingService_GetServiceErrorTrackingOverviewStreamClient = grpc.ServerStreamingClient[GetServiceErrorTrackingOverviewResponse]

func (c *errorTrackingServiceClient) GetServiceErrorTrackingListStream(ctx context.Context, in *GetServiceErrorTrackingListRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetServiceErrorTrackingListResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ErrorTrackingService_ServiceDesc.Streams[1], ErrorTrackingService_GetServiceErrorTrackingListStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetServiceErrorTrackingListRequest, GetServiceErrorTrackingListResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ErrorTrackingService_GetServiceErrorTrackingListStreamClient = grpc.ServerStreamingClient[GetServiceErrorTrackingListResponse]

// ErrorTrackingServiceServer is the server API for ErrorTrackingService service.
// All implementations must embed UnimplementedErrorTrackingServiceServer
// for forward compatibility.
type ErrorTrackingServiceServer interface {
	GetServiceErrorTrackingOverviewStream(*GetServiceErrorTrackingOverviewRequest, grpc.ServerStreamingServer[GetServiceErrorTrackingOverviewResponse]) error
	GetServiceErrorTrackingListStream(*GetServiceErrorTrackingListRequest, grpc.ServerStreamingServer[GetServiceErrorTrackingListResponse]) error
	mustEmbedUnimplementedErrorTrackingServiceServer()
}

// UnimplementedErrorTrackingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedErrorTrackingServiceServer struct{}

func (UnimplementedErrorTrackingServiceServer) GetServiceErrorTrackingOverviewStream(*GetServiceErrorTrackingOverviewRequest, grpc.ServerStreamingServer[GetServiceErrorTrackingOverviewResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetServiceErrorTrackingOverviewStream not implemented")
}
func (UnimplementedErrorTrackingServiceServer) GetServiceErrorTrackingListStream(*GetServiceErrorTrackingListRequest, grpc.ServerStreamingServer[GetServiceErrorTrackingListResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetServiceErrorTrackingListStream not implemented")
}
func (UnimplementedErrorTrackingServiceServer) mustEmbedUnimplementedErrorTrackingServiceServer() {}
func (UnimplementedErrorTrackingServiceServer) testEmbeddedByValue()                              {}

// UnsafeErrorTrackingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ErrorTrackingServiceServer will
// result in compilation errors.
type UnsafeErrorTrackingServiceServer interface {
	mustEmbedUnimplementedErrorTrackingServiceServer()
}

func RegisterErrorTrackingServiceServer(s grpc.ServiceRegistrar, srv ErrorTrackingServiceServer) {
	// If the following call pancis, it indicates UnimplementedErrorTrackingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ErrorTrackingService_ServiceDesc, srv)
}

func _ErrorTrackingService_GetServiceErrorTrackingOverviewStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetServiceErrorTrackingOverviewRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ErrorTrackingServiceServer).GetServiceErrorTrackingOverviewStream(m, &grpc.GenericServerStream[GetServiceErrorTrackingOverviewRequest, GetServiceErrorTrackingOverviewResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ErrorTrackingService_GetServiceErrorTrackingOverviewStreamServer = grpc.ServerStreamingServer[GetServiceErrorTrackingOverviewResponse]

func _ErrorTrackingService_GetServiceErrorTrackingListStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetServiceErrorTrackingListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ErrorTrackingServiceServer).GetServiceErrorTrackingListStream(m, &grpc.GenericServerStream[GetServiceErrorTrackingListRequest, GetServiceErrorTrackingListResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ErrorTrackingService_GetServiceErrorTrackingListStreamServer = grpc.ServerStreamingServer[GetServiceErrorTrackingListResponse]

// ErrorTrackingService_ServiceDesc is the grpc.ServiceDesc for ErrorTrackingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ErrorTrackingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.service_catalog.v1.ErrorTrackingService",
	HandlerType: (*ErrorTrackingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetServiceErrorTrackingOverviewStream",
			Handler:       _ErrorTrackingService_GetServiceErrorTrackingOverviewStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetServiceErrorTrackingListStream",
			Handler:       _ErrorTrackingService_GetServiceErrorTrackingListStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "com/coralogixapis/service_catalog/v1/error-tracking/error_tracking_service.proto",
}
