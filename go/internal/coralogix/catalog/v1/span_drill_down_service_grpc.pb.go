// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.28.2
// source: com/coralogix/catalog/v1/span_drill_down_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	SpanDrillDownService_GetServiceSideModelGraphs_FullMethodName = "/com.coralogix.catalog.v1.SpanDrillDownService/GetServiceSideModelGraphs"
)

// SpanDrillDownServiceClient is the client API for SpanDrillDownService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpanDrillDownServiceClient interface {
	GetServiceSideModelGraphs(ctx context.Context, in *GetServiceSideModelGraphsRequest, opts ...grpc.CallOption) (SpanDrillDownService_GetServiceSideModelGraphsClient, error)
}

type spanDrillDownServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpanDrillDownServiceClient(cc grpc.ClientConnInterface) SpanDrillDownServiceClient {
	return &spanDrillDownServiceClient{cc}
}

func (c *spanDrillDownServiceClient) GetServiceSideModelGraphs(ctx context.Context, in *GetServiceSideModelGraphsRequest, opts ...grpc.CallOption) (SpanDrillDownService_GetServiceSideModelGraphsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SpanDrillDownService_ServiceDesc.Streams[0], SpanDrillDownService_GetServiceSideModelGraphs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &spanDrillDownServiceGetServiceSideModelGraphsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SpanDrillDownService_GetServiceSideModelGraphsClient interface {
	Recv() (*GetServiceSideModelGraphsResponse, error)
	grpc.ClientStream
}

type spanDrillDownServiceGetServiceSideModelGraphsClient struct {
	grpc.ClientStream
}

func (x *spanDrillDownServiceGetServiceSideModelGraphsClient) Recv() (*GetServiceSideModelGraphsResponse, error) {
	m := new(GetServiceSideModelGraphsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpanDrillDownServiceServer is the server API for SpanDrillDownService service.
// All implementations must embed UnimplementedSpanDrillDownServiceServer
// for forward compatibility
type SpanDrillDownServiceServer interface {
	GetServiceSideModelGraphs(*GetServiceSideModelGraphsRequest, SpanDrillDownService_GetServiceSideModelGraphsServer) error
	mustEmbedUnimplementedSpanDrillDownServiceServer()
}

// UnimplementedSpanDrillDownServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSpanDrillDownServiceServer struct {
}

func (UnimplementedSpanDrillDownServiceServer) GetServiceSideModelGraphs(*GetServiceSideModelGraphsRequest, SpanDrillDownService_GetServiceSideModelGraphsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetServiceSideModelGraphs not implemented")
}
func (UnimplementedSpanDrillDownServiceServer) mustEmbedUnimplementedSpanDrillDownServiceServer() {}

// UnsafeSpanDrillDownServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpanDrillDownServiceServer will
// result in compilation errors.
type UnsafeSpanDrillDownServiceServer interface {
	mustEmbedUnimplementedSpanDrillDownServiceServer()
}

func RegisterSpanDrillDownServiceServer(s grpc.ServiceRegistrar, srv SpanDrillDownServiceServer) {
	s.RegisterService(&SpanDrillDownService_ServiceDesc, srv)
}

func _SpanDrillDownService_GetServiceSideModelGraphs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetServiceSideModelGraphsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SpanDrillDownServiceServer).GetServiceSideModelGraphs(m, &spanDrillDownServiceGetServiceSideModelGraphsServer{ServerStream: stream})
}

type SpanDrillDownService_GetServiceSideModelGraphsServer interface {
	Send(*GetServiceSideModelGraphsResponse) error
	grpc.ServerStream
}

type spanDrillDownServiceGetServiceSideModelGraphsServer struct {
	grpc.ServerStream
}

func (x *spanDrillDownServiceGetServiceSideModelGraphsServer) Send(m *GetServiceSideModelGraphsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SpanDrillDownService_ServiceDesc is the grpc.ServiceDesc for SpanDrillDownService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpanDrillDownService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.catalog.v1.SpanDrillDownService",
	HandlerType: (*SpanDrillDownServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetServiceSideModelGraphs",
			Handler:       _SpanDrillDownService_GetServiceSideModelGraphs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "com/coralogix/catalog/v1/span_drill_down_service.proto",
}
