// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.1
// source: com/coralogixapis/service_catalog/v1/distributed_tracing_service.proto

package v1

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
	DistributedTracingService_GetTimeConsuming_FullMethodName = "/com.coralogixapis.service_catalog.v1.DistributedTracingService/GetTimeConsuming"
	DistributedTracingService_GetThroughput_FullMethodName    = "/com.coralogixapis.service_catalog.v1.DistributedTracingService/GetThroughput"
	DistributedTracingService_GetErrors_FullMethodName        = "/com.coralogixapis.service_catalog.v1.DistributedTracingService/GetErrors"
	DistributedTracingService_GetErrorsStream_FullMethodName  = "/com.coralogixapis.service_catalog.v1.DistributedTracingService/GetErrorsStream"
	DistributedTracingService_GetP99Latency_FullMethodName    = "/com.coralogixapis.service_catalog.v1.DistributedTracingService/GetP99Latency"
	DistributedTracingService_ListOperations_FullMethodName   = "/com.coralogixapis.service_catalog.v1.DistributedTracingService/ListOperations"
	DistributedTracingService_GetLatencyGraph_FullMethodName  = "/com.coralogixapis.service_catalog.v1.DistributedTracingService/GetLatencyGraph"
)

// DistributedTracingServiceClient is the client API for DistributedTracingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DistributedTracingServiceClient interface {
	GetTimeConsuming(ctx context.Context, in *GetTimeConsumingRequest, opts ...grpc.CallOption) (*GetTimeConsumingResponse, error)
	GetThroughput(ctx context.Context, in *GetThroughputRequest, opts ...grpc.CallOption) (*GetThroughputResponse, error)
	GetErrors(ctx context.Context, in *GetErrorsRequest, opts ...grpc.CallOption) (*GetErrorsResponse, error)
	GetErrorsStream(ctx context.Context, in *GetErrorsStreamRequest, opts ...grpc.CallOption) (DistributedTracingService_GetErrorsStreamClient, error)
	GetP99Latency(ctx context.Context, in *GetP99LatencyRequest, opts ...grpc.CallOption) (*GetP99LatencyResponse, error)
	ListOperations(ctx context.Context, in *ListOperationsRequest, opts ...grpc.CallOption) (*ListOperationsResponse, error)
	GetLatencyGraph(ctx context.Context, in *GetLatencyGraphRequest, opts ...grpc.CallOption) (DistributedTracingService_GetLatencyGraphClient, error)
}

type distributedTracingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDistributedTracingServiceClient(cc grpc.ClientConnInterface) DistributedTracingServiceClient {
	return &distributedTracingServiceClient{cc}
}

func (c *distributedTracingServiceClient) GetTimeConsuming(ctx context.Context, in *GetTimeConsumingRequest, opts ...grpc.CallOption) (*GetTimeConsumingResponse, error) {
	out := new(GetTimeConsumingResponse)
	err := c.cc.Invoke(ctx, DistributedTracingService_GetTimeConsuming_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedTracingServiceClient) GetThroughput(ctx context.Context, in *GetThroughputRequest, opts ...grpc.CallOption) (*GetThroughputResponse, error) {
	out := new(GetThroughputResponse)
	err := c.cc.Invoke(ctx, DistributedTracingService_GetThroughput_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedTracingServiceClient) GetErrors(ctx context.Context, in *GetErrorsRequest, opts ...grpc.CallOption) (*GetErrorsResponse, error) {
	out := new(GetErrorsResponse)
	err := c.cc.Invoke(ctx, DistributedTracingService_GetErrors_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedTracingServiceClient) GetErrorsStream(ctx context.Context, in *GetErrorsStreamRequest, opts ...grpc.CallOption) (DistributedTracingService_GetErrorsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &DistributedTracingService_ServiceDesc.Streams[0], DistributedTracingService_GetErrorsStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &distributedTracingServiceGetErrorsStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DistributedTracingService_GetErrorsStreamClient interface {
	Recv() (*GetErrorsStreamResponse, error)
	grpc.ClientStream
}

type distributedTracingServiceGetErrorsStreamClient struct {
	grpc.ClientStream
}

func (x *distributedTracingServiceGetErrorsStreamClient) Recv() (*GetErrorsStreamResponse, error) {
	m := new(GetErrorsStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *distributedTracingServiceClient) GetP99Latency(ctx context.Context, in *GetP99LatencyRequest, opts ...grpc.CallOption) (*GetP99LatencyResponse, error) {
	out := new(GetP99LatencyResponse)
	err := c.cc.Invoke(ctx, DistributedTracingService_GetP99Latency_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedTracingServiceClient) ListOperations(ctx context.Context, in *ListOperationsRequest, opts ...grpc.CallOption) (*ListOperationsResponse, error) {
	out := new(ListOperationsResponse)
	err := c.cc.Invoke(ctx, DistributedTracingService_ListOperations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedTracingServiceClient) GetLatencyGraph(ctx context.Context, in *GetLatencyGraphRequest, opts ...grpc.CallOption) (DistributedTracingService_GetLatencyGraphClient, error) {
	stream, err := c.cc.NewStream(ctx, &DistributedTracingService_ServiceDesc.Streams[1], DistributedTracingService_GetLatencyGraph_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &distributedTracingServiceGetLatencyGraphClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DistributedTracingService_GetLatencyGraphClient interface {
	Recv() (*GetLatencyGraphResponse, error)
	grpc.ClientStream
}

type distributedTracingServiceGetLatencyGraphClient struct {
	grpc.ClientStream
}

func (x *distributedTracingServiceGetLatencyGraphClient) Recv() (*GetLatencyGraphResponse, error) {
	m := new(GetLatencyGraphResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DistributedTracingServiceServer is the server API for DistributedTracingService service.
// All implementations must embed UnimplementedDistributedTracingServiceServer
// for forward compatibility
type DistributedTracingServiceServer interface {
	GetTimeConsuming(context.Context, *GetTimeConsumingRequest) (*GetTimeConsumingResponse, error)
	GetThroughput(context.Context, *GetThroughputRequest) (*GetThroughputResponse, error)
	GetErrors(context.Context, *GetErrorsRequest) (*GetErrorsResponse, error)
	GetErrorsStream(*GetErrorsStreamRequest, DistributedTracingService_GetErrorsStreamServer) error
	GetP99Latency(context.Context, *GetP99LatencyRequest) (*GetP99LatencyResponse, error)
	ListOperations(context.Context, *ListOperationsRequest) (*ListOperationsResponse, error)
	GetLatencyGraph(*GetLatencyGraphRequest, DistributedTracingService_GetLatencyGraphServer) error
	mustEmbedUnimplementedDistributedTracingServiceServer()
}

// UnimplementedDistributedTracingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDistributedTracingServiceServer struct {
}

func (UnimplementedDistributedTracingServiceServer) GetTimeConsuming(context.Context, *GetTimeConsumingRequest) (*GetTimeConsumingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeConsuming not implemented")
}
func (UnimplementedDistributedTracingServiceServer) GetThroughput(context.Context, *GetThroughputRequest) (*GetThroughputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThroughput not implemented")
}
func (UnimplementedDistributedTracingServiceServer) GetErrors(context.Context, *GetErrorsRequest) (*GetErrorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetErrors not implemented")
}
func (UnimplementedDistributedTracingServiceServer) GetErrorsStream(*GetErrorsStreamRequest, DistributedTracingService_GetErrorsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetErrorsStream not implemented")
}
func (UnimplementedDistributedTracingServiceServer) GetP99Latency(context.Context, *GetP99LatencyRequest) (*GetP99LatencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetP99Latency not implemented")
}
func (UnimplementedDistributedTracingServiceServer) ListOperations(context.Context, *ListOperationsRequest) (*ListOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedDistributedTracingServiceServer) GetLatencyGraph(*GetLatencyGraphRequest, DistributedTracingService_GetLatencyGraphServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLatencyGraph not implemented")
}
func (UnimplementedDistributedTracingServiceServer) mustEmbedUnimplementedDistributedTracingServiceServer() {
}

// UnsafeDistributedTracingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DistributedTracingServiceServer will
// result in compilation errors.
type UnsafeDistributedTracingServiceServer interface {
	mustEmbedUnimplementedDistributedTracingServiceServer()
}

func RegisterDistributedTracingServiceServer(s grpc.ServiceRegistrar, srv DistributedTracingServiceServer) {
	s.RegisterService(&DistributedTracingService_ServiceDesc, srv)
}

func _DistributedTracingService_GetTimeConsuming_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeConsumingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedTracingServiceServer).GetTimeConsuming(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DistributedTracingService_GetTimeConsuming_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedTracingServiceServer).GetTimeConsuming(ctx, req.(*GetTimeConsumingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedTracingService_GetThroughput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetThroughputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedTracingServiceServer).GetThroughput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DistributedTracingService_GetThroughput_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedTracingServiceServer).GetThroughput(ctx, req.(*GetThroughputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedTracingService_GetErrors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetErrorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedTracingServiceServer).GetErrors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DistributedTracingService_GetErrors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedTracingServiceServer).GetErrors(ctx, req.(*GetErrorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedTracingService_GetErrorsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetErrorsStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DistributedTracingServiceServer).GetErrorsStream(m, &distributedTracingServiceGetErrorsStreamServer{stream})
}

type DistributedTracingService_GetErrorsStreamServer interface {
	Send(*GetErrorsStreamResponse) error
	grpc.ServerStream
}

type distributedTracingServiceGetErrorsStreamServer struct {
	grpc.ServerStream
}

func (x *distributedTracingServiceGetErrorsStreamServer) Send(m *GetErrorsStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DistributedTracingService_GetP99Latency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetP99LatencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedTracingServiceServer).GetP99Latency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DistributedTracingService_GetP99Latency_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedTracingServiceServer).GetP99Latency(ctx, req.(*GetP99LatencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedTracingService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedTracingServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DistributedTracingService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedTracingServiceServer).ListOperations(ctx, req.(*ListOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedTracingService_GetLatencyGraph_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetLatencyGraphRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DistributedTracingServiceServer).GetLatencyGraph(m, &distributedTracingServiceGetLatencyGraphServer{stream})
}

type DistributedTracingService_GetLatencyGraphServer interface {
	Send(*GetLatencyGraphResponse) error
	grpc.ServerStream
}

type distributedTracingServiceGetLatencyGraphServer struct {
	grpc.ServerStream
}

func (x *distributedTracingServiceGetLatencyGraphServer) Send(m *GetLatencyGraphResponse) error {
	return x.ServerStream.SendMsg(m)
}

// DistributedTracingService_ServiceDesc is the grpc.ServiceDesc for DistributedTracingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DistributedTracingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.service_catalog.v1.DistributedTracingService",
	HandlerType: (*DistributedTracingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTimeConsuming",
			Handler:    _DistributedTracingService_GetTimeConsuming_Handler,
		},
		{
			MethodName: "GetThroughput",
			Handler:    _DistributedTracingService_GetThroughput_Handler,
		},
		{
			MethodName: "GetErrors",
			Handler:    _DistributedTracingService_GetErrors_Handler,
		},
		{
			MethodName: "GetP99Latency",
			Handler:    _DistributedTracingService_GetP99Latency_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _DistributedTracingService_ListOperations_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetErrorsStream",
			Handler:       _DistributedTracingService_GetErrorsStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetLatencyGraph",
			Handler:       _DistributedTracingService_GetLatencyGraph_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "com/coralogixapis/service_catalog/v1/distributed_tracing_service.proto",
}
