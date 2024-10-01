// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.28.2
// source: com/coralogixapis/events2metrics/v2/events2metrics_service.proto

package v2

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
	Events2MetricService_CreateE2M_FullMethodName             = "/com.coralogixapis.events2metrics.v2.Events2MetricService/CreateE2M"
	Events2MetricService_ListE2M_FullMethodName               = "/com.coralogixapis.events2metrics.v2.Events2MetricService/ListE2M"
	Events2MetricService_ReplaceE2M_FullMethodName            = "/com.coralogixapis.events2metrics.v2.Events2MetricService/ReplaceE2M"
	Events2MetricService_GetE2M_FullMethodName                = "/com.coralogixapis.events2metrics.v2.Events2MetricService/GetE2M"
	Events2MetricService_DeleteE2M_FullMethodName             = "/com.coralogixapis.events2metrics.v2.Events2MetricService/DeleteE2M"
	Events2MetricService_AtomicBatchExecuteE2M_FullMethodName = "/com.coralogixapis.events2metrics.v2.Events2MetricService/AtomicBatchExecuteE2M"
	Events2MetricService_ListLabelsCardinality_FullMethodName = "/com.coralogixapis.events2metrics.v2.Events2MetricService/ListLabelsCardinality"
	Events2MetricService_GetLimits_FullMethodName             = "/com.coralogixapis.events2metrics.v2.Events2MetricService/GetLimits"
)

// Events2MetricServiceClient is the client API for Events2MetricService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Events2MetricServiceClient interface {
	CreateE2M(ctx context.Context, in *CreateE2MRequest, opts ...grpc.CallOption) (*CreateE2MResponse, error)
	ListE2M(ctx context.Context, in *ListE2MRequest, opts ...grpc.CallOption) (*ListE2MResponse, error)
	ReplaceE2M(ctx context.Context, in *ReplaceE2MRequest, opts ...grpc.CallOption) (*ReplaceE2MResponse, error)
	GetE2M(ctx context.Context, in *GetE2MRequest, opts ...grpc.CallOption) (*GetE2MResponse, error)
	DeleteE2M(ctx context.Context, in *DeleteE2MRequest, opts ...grpc.CallOption) (*DeleteE2MResponse, error)
	AtomicBatchExecuteE2M(ctx context.Context, in *AtomicBatchExecuteE2MRequest, opts ...grpc.CallOption) (*AtomicBatchExecuteE2MResponse, error)
	ListLabelsCardinality(ctx context.Context, in *ListLabelsCardinalityRequest, opts ...grpc.CallOption) (*ListLabelsCardinalityResponse, error)
	GetLimits(ctx context.Context, in *GetLimitsRequest, opts ...grpc.CallOption) (*GetLimitsResponse, error)
}

type events2MetricServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEvents2MetricServiceClient(cc grpc.ClientConnInterface) Events2MetricServiceClient {
	return &events2MetricServiceClient{cc}
}

func (c *events2MetricServiceClient) CreateE2M(ctx context.Context, in *CreateE2MRequest, opts ...grpc.CallOption) (*CreateE2MResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateE2MResponse)
	err := c.cc.Invoke(ctx, Events2MetricService_CreateE2M_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *events2MetricServiceClient) ListE2M(ctx context.Context, in *ListE2MRequest, opts ...grpc.CallOption) (*ListE2MResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListE2MResponse)
	err := c.cc.Invoke(ctx, Events2MetricService_ListE2M_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *events2MetricServiceClient) ReplaceE2M(ctx context.Context, in *ReplaceE2MRequest, opts ...grpc.CallOption) (*ReplaceE2MResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplaceE2MResponse)
	err := c.cc.Invoke(ctx, Events2MetricService_ReplaceE2M_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *events2MetricServiceClient) GetE2M(ctx context.Context, in *GetE2MRequest, opts ...grpc.CallOption) (*GetE2MResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetE2MResponse)
	err := c.cc.Invoke(ctx, Events2MetricService_GetE2M_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *events2MetricServiceClient) DeleteE2M(ctx context.Context, in *DeleteE2MRequest, opts ...grpc.CallOption) (*DeleteE2MResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteE2MResponse)
	err := c.cc.Invoke(ctx, Events2MetricService_DeleteE2M_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *events2MetricServiceClient) AtomicBatchExecuteE2M(ctx context.Context, in *AtomicBatchExecuteE2MRequest, opts ...grpc.CallOption) (*AtomicBatchExecuteE2MResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AtomicBatchExecuteE2MResponse)
	err := c.cc.Invoke(ctx, Events2MetricService_AtomicBatchExecuteE2M_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *events2MetricServiceClient) ListLabelsCardinality(ctx context.Context, in *ListLabelsCardinalityRequest, opts ...grpc.CallOption) (*ListLabelsCardinalityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLabelsCardinalityResponse)
	err := c.cc.Invoke(ctx, Events2MetricService_ListLabelsCardinality_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *events2MetricServiceClient) GetLimits(ctx context.Context, in *GetLimitsRequest, opts ...grpc.CallOption) (*GetLimitsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLimitsResponse)
	err := c.cc.Invoke(ctx, Events2MetricService_GetLimits_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Events2MetricServiceServer is the server API for Events2MetricService service.
// All implementations must embed UnimplementedEvents2MetricServiceServer
// for forward compatibility
type Events2MetricServiceServer interface {
	CreateE2M(context.Context, *CreateE2MRequest) (*CreateE2MResponse, error)
	ListE2M(context.Context, *ListE2MRequest) (*ListE2MResponse, error)
	ReplaceE2M(context.Context, *ReplaceE2MRequest) (*ReplaceE2MResponse, error)
	GetE2M(context.Context, *GetE2MRequest) (*GetE2MResponse, error)
	DeleteE2M(context.Context, *DeleteE2MRequest) (*DeleteE2MResponse, error)
	AtomicBatchExecuteE2M(context.Context, *AtomicBatchExecuteE2MRequest) (*AtomicBatchExecuteE2MResponse, error)
	ListLabelsCardinality(context.Context, *ListLabelsCardinalityRequest) (*ListLabelsCardinalityResponse, error)
	GetLimits(context.Context, *GetLimitsRequest) (*GetLimitsResponse, error)
	mustEmbedUnimplementedEvents2MetricServiceServer()
}

// UnimplementedEvents2MetricServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEvents2MetricServiceServer struct {
}

func (UnimplementedEvents2MetricServiceServer) CreateE2M(context.Context, *CreateE2MRequest) (*CreateE2MResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateE2M not implemented")
}
func (UnimplementedEvents2MetricServiceServer) ListE2M(context.Context, *ListE2MRequest) (*ListE2MResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListE2M not implemented")
}
func (UnimplementedEvents2MetricServiceServer) ReplaceE2M(context.Context, *ReplaceE2MRequest) (*ReplaceE2MResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceE2M not implemented")
}
func (UnimplementedEvents2MetricServiceServer) GetE2M(context.Context, *GetE2MRequest) (*GetE2MResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetE2M not implemented")
}
func (UnimplementedEvents2MetricServiceServer) DeleteE2M(context.Context, *DeleteE2MRequest) (*DeleteE2MResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteE2M not implemented")
}
func (UnimplementedEvents2MetricServiceServer) AtomicBatchExecuteE2M(context.Context, *AtomicBatchExecuteE2MRequest) (*AtomicBatchExecuteE2MResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtomicBatchExecuteE2M not implemented")
}
func (UnimplementedEvents2MetricServiceServer) ListLabelsCardinality(context.Context, *ListLabelsCardinalityRequest) (*ListLabelsCardinalityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLabelsCardinality not implemented")
}
func (UnimplementedEvents2MetricServiceServer) GetLimits(context.Context, *GetLimitsRequest) (*GetLimitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLimits not implemented")
}
func (UnimplementedEvents2MetricServiceServer) mustEmbedUnimplementedEvents2MetricServiceServer() {}

// UnsafeEvents2MetricServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Events2MetricServiceServer will
// result in compilation errors.
type UnsafeEvents2MetricServiceServer interface {
	mustEmbedUnimplementedEvents2MetricServiceServer()
}

func RegisterEvents2MetricServiceServer(s grpc.ServiceRegistrar, srv Events2MetricServiceServer) {
	s.RegisterService(&Events2MetricService_ServiceDesc, srv)
}

func _Events2MetricService_CreateE2M_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateE2MRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Events2MetricServiceServer).CreateE2M(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Events2MetricService_CreateE2M_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Events2MetricServiceServer).CreateE2M(ctx, req.(*CreateE2MRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events2MetricService_ListE2M_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListE2MRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Events2MetricServiceServer).ListE2M(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Events2MetricService_ListE2M_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Events2MetricServiceServer).ListE2M(ctx, req.(*ListE2MRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events2MetricService_ReplaceE2M_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceE2MRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Events2MetricServiceServer).ReplaceE2M(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Events2MetricService_ReplaceE2M_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Events2MetricServiceServer).ReplaceE2M(ctx, req.(*ReplaceE2MRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events2MetricService_GetE2M_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetE2MRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Events2MetricServiceServer).GetE2M(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Events2MetricService_GetE2M_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Events2MetricServiceServer).GetE2M(ctx, req.(*GetE2MRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events2MetricService_DeleteE2M_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteE2MRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Events2MetricServiceServer).DeleteE2M(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Events2MetricService_DeleteE2M_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Events2MetricServiceServer).DeleteE2M(ctx, req.(*DeleteE2MRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events2MetricService_AtomicBatchExecuteE2M_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AtomicBatchExecuteE2MRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Events2MetricServiceServer).AtomicBatchExecuteE2M(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Events2MetricService_AtomicBatchExecuteE2M_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Events2MetricServiceServer).AtomicBatchExecuteE2M(ctx, req.(*AtomicBatchExecuteE2MRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events2MetricService_ListLabelsCardinality_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLabelsCardinalityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Events2MetricServiceServer).ListLabelsCardinality(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Events2MetricService_ListLabelsCardinality_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Events2MetricServiceServer).ListLabelsCardinality(ctx, req.(*ListLabelsCardinalityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events2MetricService_GetLimits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLimitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Events2MetricServiceServer).GetLimits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Events2MetricService_GetLimits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Events2MetricServiceServer).GetLimits(ctx, req.(*GetLimitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Events2MetricService_ServiceDesc is the grpc.ServiceDesc for Events2MetricService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Events2MetricService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.events2metrics.v2.Events2MetricService",
	HandlerType: (*Events2MetricServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateE2M",
			Handler:    _Events2MetricService_CreateE2M_Handler,
		},
		{
			MethodName: "ListE2M",
			Handler:    _Events2MetricService_ListE2M_Handler,
		},
		{
			MethodName: "ReplaceE2M",
			Handler:    _Events2MetricService_ReplaceE2M_Handler,
		},
		{
			MethodName: "GetE2M",
			Handler:    _Events2MetricService_GetE2M_Handler,
		},
		{
			MethodName: "DeleteE2M",
			Handler:    _Events2MetricService_DeleteE2M_Handler,
		},
		{
			MethodName: "AtomicBatchExecuteE2M",
			Handler:    _Events2MetricService_AtomicBatchExecuteE2M_Handler,
		},
		{
			MethodName: "ListLabelsCardinality",
			Handler:    _Events2MetricService_ListLabelsCardinality_Handler,
		},
		{
			MethodName: "GetLimits",
			Handler:    _Events2MetricService_GetLimits_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/events2metrics/v2/events2metrics_service.proto",
}
