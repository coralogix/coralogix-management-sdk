// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: com/coralogixapis/service_catalog/v1/error-tracking/error_tracking_service.proto

package error_tracking

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

const (
	ErrorTrackingService_GetServiceErrorTrackingOverviewStream_FullMethodName = "/com.coralogixapis.service_catalog.v1.ErrorTrackingService/GetServiceErrorTrackingOverviewStream"
	ErrorTrackingService_GetServiceErrorTrackingListStream_FullMethodName     = "/com.coralogixapis.service_catalog.v1.ErrorTrackingService/GetServiceErrorTrackingListStream"
)

// ErrorTrackingServiceClient is the client API for ErrorTrackingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ErrorTrackingServiceClient interface {
	GetServiceErrorTrackingOverviewStream(ctx context.Context, in *GetServiceErrorTrackingOverviewRequest, opts ...grpc.CallOption) (ErrorTrackingService_GetServiceErrorTrackingOverviewStreamClient, error)
	GetServiceErrorTrackingListStream(ctx context.Context, in *GetServiceErrorTrackingListRequest, opts ...grpc.CallOption) (ErrorTrackingService_GetServiceErrorTrackingListStreamClient, error)
}

type errorTrackingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewErrorTrackingServiceClient(cc grpc.ClientConnInterface) ErrorTrackingServiceClient {
	return &errorTrackingServiceClient{cc}
}

func (c *errorTrackingServiceClient) GetServiceErrorTrackingOverviewStream(ctx context.Context, in *GetServiceErrorTrackingOverviewRequest, opts ...grpc.CallOption) (ErrorTrackingService_GetServiceErrorTrackingOverviewStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ErrorTrackingService_ServiceDesc.Streams[0], ErrorTrackingService_GetServiceErrorTrackingOverviewStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &errorTrackingServiceGetServiceErrorTrackingOverviewStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ErrorTrackingService_GetServiceErrorTrackingOverviewStreamClient interface {
	Recv() (*GetServiceErrorTrackingOverviewResponse, error)
	grpc.ClientStream
}

type errorTrackingServiceGetServiceErrorTrackingOverviewStreamClient struct {
	grpc.ClientStream
}

func (x *errorTrackingServiceGetServiceErrorTrackingOverviewStreamClient) Recv() (*GetServiceErrorTrackingOverviewResponse, error) {
	m := new(GetServiceErrorTrackingOverviewResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *errorTrackingServiceClient) GetServiceErrorTrackingListStream(ctx context.Context, in *GetServiceErrorTrackingListRequest, opts ...grpc.CallOption) (ErrorTrackingService_GetServiceErrorTrackingListStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ErrorTrackingService_ServiceDesc.Streams[1], ErrorTrackingService_GetServiceErrorTrackingListStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &errorTrackingServiceGetServiceErrorTrackingListStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ErrorTrackingService_GetServiceErrorTrackingListStreamClient interface {
	Recv() (*GetServiceErrorTrackingListResponse, error)
	grpc.ClientStream
}

type errorTrackingServiceGetServiceErrorTrackingListStreamClient struct {
	grpc.ClientStream
}

func (x *errorTrackingServiceGetServiceErrorTrackingListStreamClient) Recv() (*GetServiceErrorTrackingListResponse, error) {
	m := new(GetServiceErrorTrackingListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ErrorTrackingServiceServer is the server API for ErrorTrackingService service.
// All implementations must embed UnimplementedErrorTrackingServiceServer
// for forward compatibility
type ErrorTrackingServiceServer interface {
	GetServiceErrorTrackingOverviewStream(*GetServiceErrorTrackingOverviewRequest, ErrorTrackingService_GetServiceErrorTrackingOverviewStreamServer) error
	GetServiceErrorTrackingListStream(*GetServiceErrorTrackingListRequest, ErrorTrackingService_GetServiceErrorTrackingListStreamServer) error
	mustEmbedUnimplementedErrorTrackingServiceServer()
}

// UnimplementedErrorTrackingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedErrorTrackingServiceServer struct {
}

func (UnimplementedErrorTrackingServiceServer) GetServiceErrorTrackingOverviewStream(*GetServiceErrorTrackingOverviewRequest, ErrorTrackingService_GetServiceErrorTrackingOverviewStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetServiceErrorTrackingOverviewStream not implemented")
}
func (UnimplementedErrorTrackingServiceServer) GetServiceErrorTrackingListStream(*GetServiceErrorTrackingListRequest, ErrorTrackingService_GetServiceErrorTrackingListStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetServiceErrorTrackingListStream not implemented")
}
func (UnimplementedErrorTrackingServiceServer) mustEmbedUnimplementedErrorTrackingServiceServer() {}

// UnsafeErrorTrackingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ErrorTrackingServiceServer will
// result in compilation errors.
type UnsafeErrorTrackingServiceServer interface {
	mustEmbedUnimplementedErrorTrackingServiceServer()
}

func RegisterErrorTrackingServiceServer(s grpc.ServiceRegistrar, srv ErrorTrackingServiceServer) {
	s.RegisterService(&ErrorTrackingService_ServiceDesc, srv)
}

func _ErrorTrackingService_GetServiceErrorTrackingOverviewStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetServiceErrorTrackingOverviewRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ErrorTrackingServiceServer).GetServiceErrorTrackingOverviewStream(m, &errorTrackingServiceGetServiceErrorTrackingOverviewStreamServer{stream})
}

type ErrorTrackingService_GetServiceErrorTrackingOverviewStreamServer interface {
	Send(*GetServiceErrorTrackingOverviewResponse) error
	grpc.ServerStream
}

type errorTrackingServiceGetServiceErrorTrackingOverviewStreamServer struct {
	grpc.ServerStream
}

func (x *errorTrackingServiceGetServiceErrorTrackingOverviewStreamServer) Send(m *GetServiceErrorTrackingOverviewResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ErrorTrackingService_GetServiceErrorTrackingListStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetServiceErrorTrackingListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ErrorTrackingServiceServer).GetServiceErrorTrackingListStream(m, &errorTrackingServiceGetServiceErrorTrackingListStreamServer{stream})
}

type ErrorTrackingService_GetServiceErrorTrackingListStreamServer interface {
	Send(*GetServiceErrorTrackingListResponse) error
	grpc.ServerStream
}

type errorTrackingServiceGetServiceErrorTrackingListStreamServer struct {
	grpc.ServerStream
}

func (x *errorTrackingServiceGetServiceErrorTrackingListStreamServer) Send(m *GetServiceErrorTrackingListResponse) error {
	return x.ServerStream.SendMsg(m)
}

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
