// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: com/coralogix/archive/v2/target_service.proto

package v2

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
	TargetService_GetTarget_FullMethodName      = "/com.coralogix.archive.v2.TargetService/GetTarget"
	TargetService_SetTarget_FullMethodName      = "/com.coralogix.archive.v2.TargetService/SetTarget"
	TargetService_ValidateTarget_FullMethodName = "/com.coralogix.archive.v2.TargetService/ValidateTarget"
)

// TargetServiceClient is the client API for TargetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TargetServiceClient interface {
	GetTarget(ctx context.Context, in *GetTargetRequest, opts ...grpc.CallOption) (*GetTargetResponse, error)
	SetTarget(ctx context.Context, in *SetTargetRequest, opts ...grpc.CallOption) (*SetTargetResponse, error)
	ValidateTarget(ctx context.Context, in *ValidateTargetRequest, opts ...grpc.CallOption) (*ValidateTargetResponse, error)
}

type targetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTargetServiceClient(cc grpc.ClientConnInterface) TargetServiceClient {
	return &targetServiceClient{cc}
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

func (c *targetServiceClient) ValidateTarget(ctx context.Context, in *ValidateTargetRequest, opts ...grpc.CallOption) (*ValidateTargetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateTargetResponse)
	err := c.cc.Invoke(ctx, TargetService_ValidateTarget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TargetServiceServer is the server API for TargetService service.
// All implementations must embed UnimplementedTargetServiceServer
// for forward compatibility.
type TargetServiceServer interface {
	GetTarget(context.Context, *GetTargetRequest) (*GetTargetResponse, error)
	SetTarget(context.Context, *SetTargetRequest) (*SetTargetResponse, error)
	ValidateTarget(context.Context, *ValidateTargetRequest) (*ValidateTargetResponse, error)
	mustEmbedUnimplementedTargetServiceServer()
}

// UnimplementedTargetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTargetServiceServer struct{}

func (UnimplementedTargetServiceServer) GetTarget(context.Context, *GetTargetRequest) (*GetTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTarget not implemented")
}
func (UnimplementedTargetServiceServer) SetTarget(context.Context, *SetTargetRequest) (*SetTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTarget not implemented")
}
func (UnimplementedTargetServiceServer) ValidateTarget(context.Context, *ValidateTargetRequest) (*ValidateTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateTarget not implemented")
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

func _TargetService_ValidateTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TargetServiceServer).ValidateTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TargetService_ValidateTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TargetServiceServer).ValidateTarget(ctx, req.(*ValidateTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TargetService_ServiceDesc is the grpc.ServiceDesc for TargetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TargetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.archive.v2.TargetService",
	HandlerType: (*TargetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTarget",
			Handler:    _TargetService_GetTarget_Handler,
		},
		{
			MethodName: "SetTarget",
			Handler:    _TargetService_SetTarget_Handler,
		},
		{
			MethodName: "ValidateTarget",
			Handler:    _TargetService_ValidateTarget_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/archive/v2/target_service.proto",
}
