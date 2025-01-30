// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: com/coralogixapis/service_catalog/v1/apdex_service.proto

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
	ApdexService_GetApdexThreshold_FullMethodName          = "/com.coralogixapis.service_catalog.v1.ApdexService/GetApdexThreshold"
	ApdexService_ReplaceApdexThreshold_FullMethodName      = "/com.coralogixapis.service_catalog.v1.ApdexService/ReplaceApdexThreshold"
	ApdexService_BatchExecuteApdexThreshold_FullMethodName = "/com.coralogixapis.service_catalog.v1.ApdexService/BatchExecuteApdexThreshold"
)

// ApdexServiceClient is the client API for ApdexService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApdexServiceClient interface {
	GetApdexThreshold(ctx context.Context, in *GetApdexThresholdRequest, opts ...grpc.CallOption) (*GetApdexThresholdResponse, error)
	ReplaceApdexThreshold(ctx context.Context, in *ReplaceApdexThresholdRequest, opts ...grpc.CallOption) (*ReplaceApdexThresholdResponse, error)
	BatchExecuteApdexThreshold(ctx context.Context, in *BatchExecuteApdexThresholdRequest, opts ...grpc.CallOption) (*BatchExecuteApdexThresholdResponse, error)
}

type apdexServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApdexServiceClient(cc grpc.ClientConnInterface) ApdexServiceClient {
	return &apdexServiceClient{cc}
}

func (c *apdexServiceClient) GetApdexThreshold(ctx context.Context, in *GetApdexThresholdRequest, opts ...grpc.CallOption) (*GetApdexThresholdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetApdexThresholdResponse)
	err := c.cc.Invoke(ctx, ApdexService_GetApdexThreshold_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apdexServiceClient) ReplaceApdexThreshold(ctx context.Context, in *ReplaceApdexThresholdRequest, opts ...grpc.CallOption) (*ReplaceApdexThresholdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplaceApdexThresholdResponse)
	err := c.cc.Invoke(ctx, ApdexService_ReplaceApdexThreshold_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apdexServiceClient) BatchExecuteApdexThreshold(ctx context.Context, in *BatchExecuteApdexThresholdRequest, opts ...grpc.CallOption) (*BatchExecuteApdexThresholdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchExecuteApdexThresholdResponse)
	err := c.cc.Invoke(ctx, ApdexService_BatchExecuteApdexThreshold_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApdexServiceServer is the server API for ApdexService service.
// All implementations must embed UnimplementedApdexServiceServer
// for forward compatibility.
type ApdexServiceServer interface {
	GetApdexThreshold(context.Context, *GetApdexThresholdRequest) (*GetApdexThresholdResponse, error)
	ReplaceApdexThreshold(context.Context, *ReplaceApdexThresholdRequest) (*ReplaceApdexThresholdResponse, error)
	BatchExecuteApdexThreshold(context.Context, *BatchExecuteApdexThresholdRequest) (*BatchExecuteApdexThresholdResponse, error)
	mustEmbedUnimplementedApdexServiceServer()
}

// UnimplementedApdexServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedApdexServiceServer struct{}

func (UnimplementedApdexServiceServer) GetApdexThreshold(context.Context, *GetApdexThresholdRequest) (*GetApdexThresholdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApdexThreshold not implemented")
}
func (UnimplementedApdexServiceServer) ReplaceApdexThreshold(context.Context, *ReplaceApdexThresholdRequest) (*ReplaceApdexThresholdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceApdexThreshold not implemented")
}
func (UnimplementedApdexServiceServer) BatchExecuteApdexThreshold(context.Context, *BatchExecuteApdexThresholdRequest) (*BatchExecuteApdexThresholdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchExecuteApdexThreshold not implemented")
}
func (UnimplementedApdexServiceServer) mustEmbedUnimplementedApdexServiceServer() {}
func (UnimplementedApdexServiceServer) testEmbeddedByValue()                      {}

// UnsafeApdexServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApdexServiceServer will
// result in compilation errors.
type UnsafeApdexServiceServer interface {
	mustEmbedUnimplementedApdexServiceServer()
}

func RegisterApdexServiceServer(s grpc.ServiceRegistrar, srv ApdexServiceServer) {
	// If the following call pancis, it indicates UnimplementedApdexServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ApdexService_ServiceDesc, srv)
}

func _ApdexService_GetApdexThreshold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApdexThresholdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApdexServiceServer).GetApdexThreshold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApdexService_GetApdexThreshold_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApdexServiceServer).GetApdexThreshold(ctx, req.(*GetApdexThresholdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApdexService_ReplaceApdexThreshold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceApdexThresholdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApdexServiceServer).ReplaceApdexThreshold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApdexService_ReplaceApdexThreshold_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApdexServiceServer).ReplaceApdexThreshold(ctx, req.(*ReplaceApdexThresholdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApdexService_BatchExecuteApdexThreshold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchExecuteApdexThresholdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApdexServiceServer).BatchExecuteApdexThreshold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApdexService_BatchExecuteApdexThreshold_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApdexServiceServer).BatchExecuteApdexThreshold(ctx, req.(*BatchExecuteApdexThresholdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApdexService_ServiceDesc is the grpc.ServiceDesc for ApdexService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApdexService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.service_catalog.v1.ApdexService",
	HandlerType: (*ApdexServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApdexThreshold",
			Handler:    _ApdexService_GetApdexThreshold_Handler,
		},
		{
			MethodName: "ReplaceApdexThreshold",
			Handler:    _ApdexService_ReplaceApdexThreshold_Handler,
		},
		{
			MethodName: "BatchExecuteApdexThreshold",
			Handler:    _ApdexService_BatchExecuteApdexThreshold_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/service_catalog/v1/apdex_service.proto",
}
