// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: com/coralogixapis/scopes/v1/scopes.proto

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
	ScopesService_GetTeamScopesByIds_FullMethodName = "/com.coralogixapis.scopes.v1.ScopesService/GetTeamScopesByIds"
	ScopesService_GetTeamScopes_FullMethodName      = "/com.coralogixapis.scopes.v1.ScopesService/GetTeamScopes"
	ScopesService_CreateScope_FullMethodName        = "/com.coralogixapis.scopes.v1.ScopesService/CreateScope"
	ScopesService_UpdateScope_FullMethodName        = "/com.coralogixapis.scopes.v1.ScopesService/UpdateScope"
	ScopesService_DeleteScope_FullMethodName        = "/com.coralogixapis.scopes.v1.ScopesService/DeleteScope"
)

// ScopesServiceClient is the client API for ScopesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScopesServiceClient interface {
	GetTeamScopesByIds(ctx context.Context, in *GetTeamScopesByIdsRequest, opts ...grpc.CallOption) (*GetScopesResponse, error)
	GetTeamScopes(ctx context.Context, in *GetTeamScopesRequest, opts ...grpc.CallOption) (*GetScopesResponse, error)
	CreateScope(ctx context.Context, in *CreateScopeRequest, opts ...grpc.CallOption) (*CreateScopeResponse, error)
	UpdateScope(ctx context.Context, in *UpdateScopeRequest, opts ...grpc.CallOption) (*UpdateScopeResponse, error)
	DeleteScope(ctx context.Context, in *DeleteScopeRequest, opts ...grpc.CallOption) (*DeleteScopeResponse, error)
}

type scopesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScopesServiceClient(cc grpc.ClientConnInterface) ScopesServiceClient {
	return &scopesServiceClient{cc}
}

func (c *scopesServiceClient) GetTeamScopesByIds(ctx context.Context, in *GetTeamScopesByIdsRequest, opts ...grpc.CallOption) (*GetScopesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetScopesResponse)
	err := c.cc.Invoke(ctx, ScopesService_GetTeamScopesByIds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scopesServiceClient) GetTeamScopes(ctx context.Context, in *GetTeamScopesRequest, opts ...grpc.CallOption) (*GetScopesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetScopesResponse)
	err := c.cc.Invoke(ctx, ScopesService_GetTeamScopes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scopesServiceClient) CreateScope(ctx context.Context, in *CreateScopeRequest, opts ...grpc.CallOption) (*CreateScopeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateScopeResponse)
	err := c.cc.Invoke(ctx, ScopesService_CreateScope_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scopesServiceClient) UpdateScope(ctx context.Context, in *UpdateScopeRequest, opts ...grpc.CallOption) (*UpdateScopeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateScopeResponse)
	err := c.cc.Invoke(ctx, ScopesService_UpdateScope_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scopesServiceClient) DeleteScope(ctx context.Context, in *DeleteScopeRequest, opts ...grpc.CallOption) (*DeleteScopeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteScopeResponse)
	err := c.cc.Invoke(ctx, ScopesService_DeleteScope_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScopesServiceServer is the server API for ScopesService service.
// All implementations must embed UnimplementedScopesServiceServer
// for forward compatibility.
type ScopesServiceServer interface {
	GetTeamScopesByIds(context.Context, *GetTeamScopesByIdsRequest) (*GetScopesResponse, error)
	GetTeamScopes(context.Context, *GetTeamScopesRequest) (*GetScopesResponse, error)
	CreateScope(context.Context, *CreateScopeRequest) (*CreateScopeResponse, error)
	UpdateScope(context.Context, *UpdateScopeRequest) (*UpdateScopeResponse, error)
	DeleteScope(context.Context, *DeleteScopeRequest) (*DeleteScopeResponse, error)
	mustEmbedUnimplementedScopesServiceServer()
}

// UnimplementedScopesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedScopesServiceServer struct{}

func (UnimplementedScopesServiceServer) GetTeamScopesByIds(context.Context, *GetTeamScopesByIdsRequest) (*GetScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamScopesByIds not implemented")
}
func (UnimplementedScopesServiceServer) GetTeamScopes(context.Context, *GetTeamScopesRequest) (*GetScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamScopes not implemented")
}
func (UnimplementedScopesServiceServer) CreateScope(context.Context, *CreateScopeRequest) (*CreateScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScope not implemented")
}
func (UnimplementedScopesServiceServer) UpdateScope(context.Context, *UpdateScopeRequest) (*UpdateScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScope not implemented")
}
func (UnimplementedScopesServiceServer) DeleteScope(context.Context, *DeleteScopeRequest) (*DeleteScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScope not implemented")
}
func (UnimplementedScopesServiceServer) mustEmbedUnimplementedScopesServiceServer() {}
func (UnimplementedScopesServiceServer) testEmbeddedByValue()                       {}

// UnsafeScopesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScopesServiceServer will
// result in compilation errors.
type UnsafeScopesServiceServer interface {
	mustEmbedUnimplementedScopesServiceServer()
}

func RegisterScopesServiceServer(s grpc.ServiceRegistrar, srv ScopesServiceServer) {
	// If the following call pancis, it indicates UnimplementedScopesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ScopesService_ServiceDesc, srv)
}

func _ScopesService_GetTeamScopesByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamScopesByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopesServiceServer).GetTeamScopesByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScopesService_GetTeamScopesByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopesServiceServer).GetTeamScopesByIds(ctx, req.(*GetTeamScopesByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScopesService_GetTeamScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopesServiceServer).GetTeamScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScopesService_GetTeamScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopesServiceServer).GetTeamScopes(ctx, req.(*GetTeamScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScopesService_CreateScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopesServiceServer).CreateScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScopesService_CreateScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopesServiceServer).CreateScope(ctx, req.(*CreateScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScopesService_UpdateScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopesServiceServer).UpdateScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScopesService_UpdateScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopesServiceServer).UpdateScope(ctx, req.(*UpdateScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScopesService_DeleteScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopesServiceServer).DeleteScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScopesService_DeleteScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopesServiceServer).DeleteScope(ctx, req.(*DeleteScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScopesService_ServiceDesc is the grpc.ServiceDesc for ScopesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScopesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.scopes.v1.ScopesService",
	HandlerType: (*ScopesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeamScopesByIds",
			Handler:    _ScopesService_GetTeamScopesByIds_Handler,
		},
		{
			MethodName: "GetTeamScopes",
			Handler:    _ScopesService_GetTeamScopes_Handler,
		},
		{
			MethodName: "CreateScope",
			Handler:    _ScopesService_CreateScope_Handler,
		},
		{
			MethodName: "UpdateScope",
			Handler:    _ScopesService_UpdateScope_Handler,
		},
		{
			MethodName: "DeleteScope",
			Handler:    _ScopesService_DeleteScope_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/scopes/v1/scopes.proto",
}
