// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: com/coralogix/archive/v1/target_service.proto

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
	TargetService_IsArchiveConfigured_FullMethodName = "/com.coralogix.archive.v1.TargetService/IsArchiveConfigured"
	TargetService_GetTarget_FullMethodName           = "/com.coralogix.archive.v1.TargetService/GetTarget"
	TargetService_SetTarget_FullMethodName           = "/com.coralogix.archive.v1.TargetService/SetTarget"
)

// TargetServiceClient is the client API for TargetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TargetServiceClient interface {
	IsArchiveConfigured(ctx context.Context, in *IsArchiveConfiguredRequest, opts ...grpc.CallOption) (*IsArchiveConfiguredResponse, error)
	GetTarget(ctx context.Context, in *GetTargetRequest, opts ...grpc.CallOption) (*GetTargetResponse, error)
	SetTarget(ctx context.Context, in *SetTargetRequest, opts ...grpc.CallOption) (*SetTargetResponse, error)
}

type targetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTargetServiceClient(cc grpc.ClientConnInterface) TargetServiceClient {
	return &targetServiceClient{cc}
}

func (c *targetServiceClient) IsArchiveConfigured(ctx context.Context, in *IsArchiveConfiguredRequest, opts ...grpc.CallOption) (*IsArchiveConfiguredResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsArchiveConfiguredResponse)
	err := c.cc.Invoke(ctx, TargetService_IsArchiveConfigured_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetServiceClient) GetTarget(ctx context.Context, in *GetTargetRequest, opts ...grpc.CallOption) (*GetTargetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTargetResponse)
	err := c.cc.Invoke(ctx, TargetService_GetTarget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *targetServiceClient) SetTarget(ctx context.Context, in *SetTargetRequest, opts ...grpc.CallOption) (*SetTargetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetTargetResponse)
	err := c.cc.Invoke(ctx, TargetService_SetTarget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TargetServiceServer is the server API for TargetService service.
// All implementations must embed UnimplementedTargetServiceServer
// for forward compatibility.
type TargetServiceServer interface {
	IsArchiveConfigured(context.Context, *IsArchiveConfiguredRequest) (*IsArchiveConfiguredResponse, error)
	GetTarget(context.Context, *GetTargetRequest) (*GetTargetResponse, error)
	SetTarget(context.Context, *SetTargetRequest) (*SetTargetResponse, error)
	mustEmbedUnimplementedTargetServiceServer()
}

// UnimplementedTargetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTargetServiceServer struct{}

func (UnimplementedTargetServiceServer) IsArchiveConfigured(context.Context, *IsArchiveConfiguredRequest) (*IsArchiveConfiguredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsArchiveConfigured not implemented")
}
func (UnimplementedTargetServiceServer) GetTarget(context.Context, *GetTargetRequest) (*GetTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTarget not implemented")
}
func (UnimplementedTargetServiceServer) SetTarget(context.Context, *SetTargetRequest) (*SetTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTarget not implemented")
}
func (UnimplementedTargetServiceServer) mustEmbedUnimplementedTargetServiceServer() {}
func (UnimplementedTargetServiceServer) testEmbeddedByValue()                       {}

// UnsafeTargetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TargetServiceServer will
// result in compilation errors.
type UnsafeTargetServiceServer interface {
	mustEmbedUnimplementedTargetServiceServer()
}

func RegisterTargetServiceServer(s grpc.ServiceRegistrar, srv TargetServiceServer) {
	// If the following call pancis, it indicates UnimplementedTargetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TargetService_ServiceDesc, srv)
}

func _TargetService_IsArchiveConfigured_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsArchiveConfiguredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetServiceServer).IsArchiveConfigured(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TargetService_IsArchiveConfigured_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetServiceServer).IsArchiveConfigured(ctx, req.(*IsArchiveConfiguredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TargetService_GetTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetServiceServer).GetTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TargetService_GetTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetServiceServer).GetTarget(ctx, req.(*GetTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TargetService_SetTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetServiceServer).SetTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TargetService_SetTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetServiceServer).SetTarget(ctx, req.(*SetTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TargetService_ServiceDesc is the grpc.ServiceDesc for TargetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TargetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.archive.v1.TargetService",
	HandlerType: (*TargetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsArchiveConfigured",
			Handler:    _TargetService_IsArchiveConfigured_Handler,
		},
		{
			MethodName: "GetTarget",
			Handler:    _TargetService_GetTarget_Handler,
		},
		{
			MethodName: "SetTarget",
			Handler:    _TargetService_SetTarget_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/archive/v1/target_service.proto",
}

const (
	InternalTargetService_IsArchiveConfigured_FullMethodName = "/com.coralogix.archive.v1.InternalTargetService/IsArchiveConfigured"
)

// InternalTargetServiceClient is the client API for InternalTargetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalTargetServiceClient interface {
	IsArchiveConfigured(ctx context.Context, in *InternalTargetServiceIsArchiveConfiguredRequest, opts ...grpc.CallOption) (*InternalTargetServiceIsArchiveConfiguredResponse, error)
}

type internalTargetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalTargetServiceClient(cc grpc.ClientConnInterface) InternalTargetServiceClient {
	return &internalTargetServiceClient{cc}
}

func (c *internalTargetServiceClient) IsArchiveConfigured(ctx context.Context, in *InternalTargetServiceIsArchiveConfiguredRequest, opts ...grpc.CallOption) (*InternalTargetServiceIsArchiveConfiguredResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InternalTargetServiceIsArchiveConfiguredResponse)
	err := c.cc.Invoke(ctx, InternalTargetService_IsArchiveConfigured_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalTargetServiceServer is the server API for InternalTargetService service.
// All implementations must embed UnimplementedInternalTargetServiceServer
// for forward compatibility.
type InternalTargetServiceServer interface {
	IsArchiveConfigured(context.Context, *InternalTargetServiceIsArchiveConfiguredRequest) (*InternalTargetServiceIsArchiveConfiguredResponse, error)
	mustEmbedUnimplementedInternalTargetServiceServer()
}

// UnimplementedInternalTargetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInternalTargetServiceServer struct{}

func (UnimplementedInternalTargetServiceServer) IsArchiveConfigured(context.Context, *InternalTargetServiceIsArchiveConfiguredRequest) (*InternalTargetServiceIsArchiveConfiguredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsArchiveConfigured not implemented")
}
func (UnimplementedInternalTargetServiceServer) mustEmbedUnimplementedInternalTargetServiceServer() {}
func (UnimplementedInternalTargetServiceServer) testEmbeddedByValue()                               {}

// UnsafeInternalTargetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternalTargetServiceServer will
// result in compilation errors.
type UnsafeInternalTargetServiceServer interface {
	mustEmbedUnimplementedInternalTargetServiceServer()
}

func RegisterInternalTargetServiceServer(s grpc.ServiceRegistrar, srv InternalTargetServiceServer) {
	// If the following call pancis, it indicates UnimplementedInternalTargetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InternalTargetService_ServiceDesc, srv)
}

func _InternalTargetService_IsArchiveConfigured_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalTargetServiceIsArchiveConfiguredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalTargetServiceServer).IsArchiveConfigured(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalTargetService_IsArchiveConfigured_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalTargetServiceServer).IsArchiveConfigured(ctx, req.(*InternalTargetServiceIsArchiveConfiguredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InternalTargetService_ServiceDesc is the grpc.ServiceDesc for InternalTargetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InternalTargetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.archive.v1.InternalTargetService",
	HandlerType: (*InternalTargetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsArchiveConfigured",
			Handler:    _InternalTargetService_IsArchiveConfigured_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/archive/v1/target_service.proto",
}
