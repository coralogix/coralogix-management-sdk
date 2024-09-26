// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.1
// source: com/coralogixapis/service_catalog/v1/overview_service.proto

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
	OverviewService_GetOverview_FullMethodName       = "/com.coralogixapis.service_catalog.v1.OverviewService/GetOverview"
	OverviewService_GetOverviewStream_FullMethodName = "/com.coralogixapis.service_catalog.v1.OverviewService/GetOverviewStream"
	OverviewService_GetApdex_FullMethodName          = "/com.coralogixapis.service_catalog.v1.OverviewService/GetApdex"
)

// OverviewServiceClient is the client API for OverviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OverviewServiceClient interface {
	GetOverview(ctx context.Context, in *GetOverviewRequest, opts ...grpc.CallOption) (*GetOverviewResponse, error)
	GetOverviewStream(ctx context.Context, in *GetOverviewStreamRequest, opts ...grpc.CallOption) (OverviewService_GetOverviewStreamClient, error)
	GetApdex(ctx context.Context, in *GetApdexRequest, opts ...grpc.CallOption) (*GetApdexResponse, error)
}

type overviewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOverviewServiceClient(cc grpc.ClientConnInterface) OverviewServiceClient {
	return &overviewServiceClient{cc}
}

func (c *overviewServiceClient) GetOverview(ctx context.Context, in *GetOverviewRequest, opts ...grpc.CallOption) (*GetOverviewResponse, error) {
	out := new(GetOverviewResponse)
	err := c.cc.Invoke(ctx, OverviewService_GetOverview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *overviewServiceClient) GetOverviewStream(ctx context.Context, in *GetOverviewStreamRequest, opts ...grpc.CallOption) (OverviewService_GetOverviewStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &OverviewService_ServiceDesc.Streams[0], OverviewService_GetOverviewStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &overviewServiceGetOverviewStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OverviewService_GetOverviewStreamClient interface {
	Recv() (*GetOverviewStreamResponse, error)
	grpc.ClientStream
}

type overviewServiceGetOverviewStreamClient struct {
	grpc.ClientStream
}

func (x *overviewServiceGetOverviewStreamClient) Recv() (*GetOverviewStreamResponse, error) {
	m := new(GetOverviewStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *overviewServiceClient) GetApdex(ctx context.Context, in *GetApdexRequest, opts ...grpc.CallOption) (*GetApdexResponse, error) {
	out := new(GetApdexResponse)
	err := c.cc.Invoke(ctx, OverviewService_GetApdex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OverviewServiceServer is the server API for OverviewService service.
// All implementations must embed UnimplementedOverviewServiceServer
// for forward compatibility
type OverviewServiceServer interface {
	GetOverview(context.Context, *GetOverviewRequest) (*GetOverviewResponse, error)
	GetOverviewStream(*GetOverviewStreamRequest, OverviewService_GetOverviewStreamServer) error
	GetApdex(context.Context, *GetApdexRequest) (*GetApdexResponse, error)
	mustEmbedUnimplementedOverviewServiceServer()
}

// UnimplementedOverviewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOverviewServiceServer struct {
}

func (UnimplementedOverviewServiceServer) GetOverview(context.Context, *GetOverviewRequest) (*GetOverviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOverview not implemented")
}
func (UnimplementedOverviewServiceServer) GetOverviewStream(*GetOverviewStreamRequest, OverviewService_GetOverviewStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOverviewStream not implemented")
}
func (UnimplementedOverviewServiceServer) GetApdex(context.Context, *GetApdexRequest) (*GetApdexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApdex not implemented")
}
func (UnimplementedOverviewServiceServer) mustEmbedUnimplementedOverviewServiceServer() {}

// UnsafeOverviewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OverviewServiceServer will
// result in compilation errors.
type UnsafeOverviewServiceServer interface {
	mustEmbedUnimplementedOverviewServiceServer()
}

func RegisterOverviewServiceServer(s grpc.ServiceRegistrar, srv OverviewServiceServer) {
	s.RegisterService(&OverviewService_ServiceDesc, srv)
}

func _OverviewService_GetOverview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOverviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverviewServiceServer).GetOverview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OverviewService_GetOverview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverviewServiceServer).GetOverview(ctx, req.(*GetOverviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OverviewService_GetOverviewStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetOverviewStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OverviewServiceServer).GetOverviewStream(m, &overviewServiceGetOverviewStreamServer{stream})
}

type OverviewService_GetOverviewStreamServer interface {
	Send(*GetOverviewStreamResponse) error
	grpc.ServerStream
}

type overviewServiceGetOverviewStreamServer struct {
	grpc.ServerStream
}

func (x *overviewServiceGetOverviewStreamServer) Send(m *GetOverviewStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _OverviewService_GetApdex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApdexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverviewServiceServer).GetApdex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OverviewService_GetApdex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverviewServiceServer).GetApdex(ctx, req.(*GetApdexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OverviewService_ServiceDesc is the grpc.ServiceDesc for OverviewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OverviewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.service_catalog.v1.OverviewService",
	HandlerType: (*OverviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOverview",
			Handler:    _OverviewService_GetOverview_Handler,
		},
		{
			MethodName: "GetApdex",
			Handler:    _OverviewService_GetApdex_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetOverviewStream",
			Handler:       _OverviewService_GetOverviewStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "com/coralogixapis/service_catalog/v1/overview_service.proto",
}
