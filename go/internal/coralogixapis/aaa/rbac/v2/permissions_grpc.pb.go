// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: src/com/coralogixapis/aaa/rbac/v2/permissions.proto

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
	PermissionsService_ListAllPermissions_FullMethodName = "/com.coralogixapis.aaa.rbac.v2.PermissionsService/ListAllPermissions"
)

// PermissionsServiceClient is the client API for PermissionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionsServiceClient interface {
	ListAllPermissions(ctx context.Context, in *ListAllPermissionsRequest, opts ...grpc.CallOption) (*ListAllPermissionsResponse, error)
}

type permissionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionsServiceClient(cc grpc.ClientConnInterface) PermissionsServiceClient {
	return &permissionsServiceClient{cc}
}

func (c *permissionsServiceClient) ListAllPermissions(ctx context.Context, in *ListAllPermissionsRequest, opts ...grpc.CallOption) (*ListAllPermissionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAllPermissionsResponse)
	err := c.cc.Invoke(ctx, PermissionsService_ListAllPermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionsServiceServer is the server API for PermissionsService service.
// All implementations must embed UnimplementedPermissionsServiceServer
// for forward compatibility.
type PermissionsServiceServer interface {
	ListAllPermissions(context.Context, *ListAllPermissionsRequest) (*ListAllPermissionsResponse, error)
	mustEmbedUnimplementedPermissionsServiceServer()
}

// UnimplementedPermissionsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPermissionsServiceServer struct{}

func (UnimplementedPermissionsServiceServer) ListAllPermissions(context.Context, *ListAllPermissionsRequest) (*ListAllPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllPermissions not implemented")
}
func (UnimplementedPermissionsServiceServer) mustEmbedUnimplementedPermissionsServiceServer() {}
func (UnimplementedPermissionsServiceServer) testEmbeddedByValue()                            {}

// UnsafePermissionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionsServiceServer will
// result in compilation errors.
type UnsafePermissionsServiceServer interface {
	mustEmbedUnimplementedPermissionsServiceServer()
}

func RegisterPermissionsServiceServer(s grpc.ServiceRegistrar, srv PermissionsServiceServer) {
	// If the following call pancis, it indicates UnimplementedPermissionsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PermissionsService_ServiceDesc, srv)
}

func _PermissionsService_ListAllPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServiceServer).ListAllPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionsService_ListAllPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServiceServer).ListAllPermissions(ctx, req.(*ListAllPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionsService_ServiceDesc is the grpc.ServiceDesc for PermissionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.aaa.rbac.v2.PermissionsService",
	HandlerType: (*PermissionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAllPermissions",
			Handler:    _PermissionsService_ListAllPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/com/coralogixapis/aaa/rbac/v2/permissions.proto",
}
