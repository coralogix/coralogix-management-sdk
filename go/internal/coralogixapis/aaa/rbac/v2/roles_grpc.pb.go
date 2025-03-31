// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: com/coralogixapis/aaa/rbac/v2/roles.proto

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
	RoleManagementService_ListSystemRoles_FullMethodName = "/com.coralogixapis.aaa.rbac.v2.RoleManagementService/ListSystemRoles"
	RoleManagementService_ListCustomRoles_FullMethodName = "/com.coralogixapis.aaa.rbac.v2.RoleManagementService/ListCustomRoles"
	RoleManagementService_GetCustomRole_FullMethodName   = "/com.coralogixapis.aaa.rbac.v2.RoleManagementService/GetCustomRole"
	RoleManagementService_CreateRole_FullMethodName      = "/com.coralogixapis.aaa.rbac.v2.RoleManagementService/CreateRole"
	RoleManagementService_UpdateRole_FullMethodName      = "/com.coralogixapis.aaa.rbac.v2.RoleManagementService/UpdateRole"
	RoleManagementService_DeleteRole_FullMethodName      = "/com.coralogixapis.aaa.rbac.v2.RoleManagementService/DeleteRole"
)

// RoleManagementServiceClient is the client API for RoleManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleManagementServiceClient interface {
	ListSystemRoles(ctx context.Context, in *ListSystemRolesRequest, opts ...grpc.CallOption) (*ListSystemRolesResponse, error)
	ListCustomRoles(ctx context.Context, in *ListCustomRolesRequest, opts ...grpc.CallOption) (*ListCustomRolesResponse, error)
	GetCustomRole(ctx context.Context, in *GetCustomRoleRequest, opts ...grpc.CallOption) (*GetCustomRoleResponse, error)
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error)
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error)
	DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*DeleteRoleResponse, error)
}

type roleManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleManagementServiceClient(cc grpc.ClientConnInterface) RoleManagementServiceClient {
	return &roleManagementServiceClient{cc}
}

func (c *roleManagementServiceClient) ListSystemRoles(ctx context.Context, in *ListSystemRolesRequest, opts ...grpc.CallOption) (*ListSystemRolesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSystemRolesResponse)
	err := c.cc.Invoke(ctx, RoleManagementService_ListSystemRoles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleManagementServiceClient) ListCustomRoles(ctx context.Context, in *ListCustomRolesRequest, opts ...grpc.CallOption) (*ListCustomRolesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCustomRolesResponse)
	err := c.cc.Invoke(ctx, RoleManagementService_ListCustomRoles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleManagementServiceClient) GetCustomRole(ctx context.Context, in *GetCustomRoleRequest, opts ...grpc.CallOption) (*GetCustomRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCustomRoleResponse)
	err := c.cc.Invoke(ctx, RoleManagementService_GetCustomRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleManagementServiceClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRoleResponse)
	err := c.cc.Invoke(ctx, RoleManagementService_CreateRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleManagementServiceClient) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateRoleResponse)
	err := c.cc.Invoke(ctx, RoleManagementService_UpdateRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleManagementServiceClient) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*DeleteRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRoleResponse)
	err := c.cc.Invoke(ctx, RoleManagementService_DeleteRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleManagementServiceServer is the server API for RoleManagementService service.
// All implementations must embed UnimplementedRoleManagementServiceServer
// for forward compatibility.
type RoleManagementServiceServer interface {
	ListSystemRoles(context.Context, *ListSystemRolesRequest) (*ListSystemRolesResponse, error)
	ListCustomRoles(context.Context, *ListCustomRolesRequest) (*ListCustomRolesResponse, error)
	GetCustomRole(context.Context, *GetCustomRoleRequest) (*GetCustomRoleResponse, error)
	CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error)
	UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleResponse, error)
	DeleteRole(context.Context, *DeleteRoleRequest) (*DeleteRoleResponse, error)
	mustEmbedUnimplementedRoleManagementServiceServer()
}

// UnimplementedRoleManagementServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRoleManagementServiceServer struct{}

func (UnimplementedRoleManagementServiceServer) ListSystemRoles(context.Context, *ListSystemRolesRequest) (*ListSystemRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSystemRoles not implemented")
}
func (UnimplementedRoleManagementServiceServer) ListCustomRoles(context.Context, *ListCustomRolesRequest) (*ListCustomRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCustomRoles not implemented")
}
func (UnimplementedRoleManagementServiceServer) GetCustomRole(context.Context, *GetCustomRoleRequest) (*GetCustomRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomRole not implemented")
}
func (UnimplementedRoleManagementServiceServer) CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedRoleManagementServiceServer) UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedRoleManagementServiceServer) DeleteRole(context.Context, *DeleteRoleRequest) (*DeleteRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedRoleManagementServiceServer) mustEmbedUnimplementedRoleManagementServiceServer() {}
func (UnimplementedRoleManagementServiceServer) testEmbeddedByValue()                               {}

// UnsafeRoleManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleManagementServiceServer will
// result in compilation errors.
type UnsafeRoleManagementServiceServer interface {
	mustEmbedUnimplementedRoleManagementServiceServer()
}

func RegisterRoleManagementServiceServer(s grpc.ServiceRegistrar, srv RoleManagementServiceServer) {
	// If the following call pancis, it indicates UnimplementedRoleManagementServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RoleManagementService_ServiceDesc, srv)
}

func _RoleManagementService_ListSystemRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSystemRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleManagementServiceServer).ListSystemRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleManagementService_ListSystemRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleManagementServiceServer).ListSystemRoles(ctx, req.(*ListSystemRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleManagementService_ListCustomRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCustomRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleManagementServiceServer).ListCustomRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleManagementService_ListCustomRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleManagementServiceServer).ListCustomRoles(ctx, req.(*ListCustomRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleManagementService_GetCustomRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleManagementServiceServer).GetCustomRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleManagementService_GetCustomRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleManagementServiceServer).GetCustomRole(ctx, req.(*GetCustomRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleManagementService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleManagementServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleManagementService_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleManagementServiceServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleManagementService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleManagementServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleManagementService_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleManagementServiceServer).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleManagementService_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleManagementServiceServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleManagementService_DeleteRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleManagementServiceServer).DeleteRole(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleManagementService_ServiceDesc is the grpc.ServiceDesc for RoleManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.aaa.rbac.v2.RoleManagementService",
	HandlerType: (*RoleManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSystemRoles",
			Handler:    _RoleManagementService_ListSystemRoles_Handler,
		},
		{
			MethodName: "ListCustomRoles",
			Handler:    _RoleManagementService_ListCustomRoles_Handler,
		},
		{
			MethodName: "GetCustomRole",
			Handler:    _RoleManagementService_GetCustomRole_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _RoleManagementService_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _RoleManagementService_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _RoleManagementService_DeleteRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/aaa/rbac/v2/roles.proto",
}
