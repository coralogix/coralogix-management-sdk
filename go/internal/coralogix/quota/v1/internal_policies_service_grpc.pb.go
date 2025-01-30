// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: com/coralogix/quota/v1/internal_policies_service.proto

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
	InternalPoliciesService_GetPolicy_FullMethodName           = "/com.coralogix.quota.v1.InternalPoliciesService/GetPolicy"
	InternalPoliciesService_CreatePolicy_FullMethodName        = "/com.coralogix.quota.v1.InternalPoliciesService/CreatePolicy"
	InternalPoliciesService_UpdatePolicy_FullMethodName        = "/com.coralogix.quota.v1.InternalPoliciesService/UpdatePolicy"
	InternalPoliciesService_GetCompanyPolicies_FullMethodName  = "/com.coralogix.quota.v1.InternalPoliciesService/GetCompanyPolicies"
	InternalPoliciesService_DeletePolicy_FullMethodName        = "/com.coralogix.quota.v1.InternalPoliciesService/DeletePolicy"
	InternalPoliciesService_ReorderPolicies_FullMethodName     = "/com.coralogix.quota.v1.InternalPoliciesService/ReorderPolicies"
	InternalPoliciesService_BulkTestLogPolicies_FullMethodName = "/com.coralogix.quota.v1.InternalPoliciesService/BulkTestLogPolicies"
	InternalPoliciesService_TogglePolicy_FullMethodName        = "/com.coralogix.quota.v1.InternalPoliciesService/TogglePolicy"
)

// InternalPoliciesServiceClient is the client API for InternalPoliciesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalPoliciesServiceClient interface {
	GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*GetPolicyResponse, error)
	CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*CreatePolicyResponse, error)
	UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*UpdatePolicyResponse, error)
	GetCompanyPolicies(ctx context.Context, in *GetCompanyPoliciesRequest, opts ...grpc.CallOption) (*GetCompanyPoliciesResponse, error)
	DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*DeletePolicyResponse, error)
	ReorderPolicies(ctx context.Context, in *ReorderPoliciesRequest, opts ...grpc.CallOption) (*ReorderPoliciesResponse, error)
	BulkTestLogPolicies(ctx context.Context, in *BulkTestLogPoliciesRequest, opts ...grpc.CallOption) (*BulkTestLogPoliciesResponse, error)
	TogglePolicy(ctx context.Context, in *TogglePolicyRequest, opts ...grpc.CallOption) (*TogglePolicyResponse, error)
}

type internalPoliciesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalPoliciesServiceClient(cc grpc.ClientConnInterface) InternalPoliciesServiceClient {
	return &internalPoliciesServiceClient{cc}
}

func (c *internalPoliciesServiceClient) GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*GetPolicyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPolicyResponse)
	err := c.cc.Invoke(ctx, InternalPoliciesService_GetPolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalPoliciesServiceClient) CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*CreatePolicyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePolicyResponse)
	err := c.cc.Invoke(ctx, InternalPoliciesService_CreatePolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalPoliciesServiceClient) UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*UpdatePolicyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePolicyResponse)
	err := c.cc.Invoke(ctx, InternalPoliciesService_UpdatePolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalPoliciesServiceClient) GetCompanyPolicies(ctx context.Context, in *GetCompanyPoliciesRequest, opts ...grpc.CallOption) (*GetCompanyPoliciesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCompanyPoliciesResponse)
	err := c.cc.Invoke(ctx, InternalPoliciesService_GetCompanyPolicies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalPoliciesServiceClient) DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*DeletePolicyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePolicyResponse)
	err := c.cc.Invoke(ctx, InternalPoliciesService_DeletePolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalPoliciesServiceClient) ReorderPolicies(ctx context.Context, in *ReorderPoliciesRequest, opts ...grpc.CallOption) (*ReorderPoliciesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReorderPoliciesResponse)
	err := c.cc.Invoke(ctx, InternalPoliciesService_ReorderPolicies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalPoliciesServiceClient) BulkTestLogPolicies(ctx context.Context, in *BulkTestLogPoliciesRequest, opts ...grpc.CallOption) (*BulkTestLogPoliciesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BulkTestLogPoliciesResponse)
	err := c.cc.Invoke(ctx, InternalPoliciesService_BulkTestLogPolicies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalPoliciesServiceClient) TogglePolicy(ctx context.Context, in *TogglePolicyRequest, opts ...grpc.CallOption) (*TogglePolicyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TogglePolicyResponse)
	err := c.cc.Invoke(ctx, InternalPoliciesService_TogglePolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalPoliciesServiceServer is the server API for InternalPoliciesService service.
// All implementations must embed UnimplementedInternalPoliciesServiceServer
// for forward compatibility.
type InternalPoliciesServiceServer interface {
	GetPolicy(context.Context, *GetPolicyRequest) (*GetPolicyResponse, error)
	CreatePolicy(context.Context, *CreatePolicyRequest) (*CreatePolicyResponse, error)
	UpdatePolicy(context.Context, *UpdatePolicyRequest) (*UpdatePolicyResponse, error)
	GetCompanyPolicies(context.Context, *GetCompanyPoliciesRequest) (*GetCompanyPoliciesResponse, error)
	DeletePolicy(context.Context, *DeletePolicyRequest) (*DeletePolicyResponse, error)
	ReorderPolicies(context.Context, *ReorderPoliciesRequest) (*ReorderPoliciesResponse, error)
	BulkTestLogPolicies(context.Context, *BulkTestLogPoliciesRequest) (*BulkTestLogPoliciesResponse, error)
	TogglePolicy(context.Context, *TogglePolicyRequest) (*TogglePolicyResponse, error)
	mustEmbedUnimplementedInternalPoliciesServiceServer()
}

// UnimplementedInternalPoliciesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInternalPoliciesServiceServer struct{}

func (UnimplementedInternalPoliciesServiceServer) GetPolicy(context.Context, *GetPolicyRequest) (*GetPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicy not implemented")
}
func (UnimplementedInternalPoliciesServiceServer) CreatePolicy(context.Context, *CreatePolicyRequest) (*CreatePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicy not implemented")
}
func (UnimplementedInternalPoliciesServiceServer) UpdatePolicy(context.Context, *UpdatePolicyRequest) (*UpdatePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePolicy not implemented")
}
func (UnimplementedInternalPoliciesServiceServer) GetCompanyPolicies(context.Context, *GetCompanyPoliciesRequest) (*GetCompanyPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyPolicies not implemented")
}
func (UnimplementedInternalPoliciesServiceServer) DeletePolicy(context.Context, *DeletePolicyRequest) (*DeletePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicy not implemented")
}
func (UnimplementedInternalPoliciesServiceServer) ReorderPolicies(context.Context, *ReorderPoliciesRequest) (*ReorderPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReorderPolicies not implemented")
}
func (UnimplementedInternalPoliciesServiceServer) BulkTestLogPolicies(context.Context, *BulkTestLogPoliciesRequest) (*BulkTestLogPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkTestLogPolicies not implemented")
}
func (UnimplementedInternalPoliciesServiceServer) TogglePolicy(context.Context, *TogglePolicyRequest) (*TogglePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TogglePolicy not implemented")
}
func (UnimplementedInternalPoliciesServiceServer) mustEmbedUnimplementedInternalPoliciesServiceServer() {
}
func (UnimplementedInternalPoliciesServiceServer) testEmbeddedByValue() {}

// UnsafeInternalPoliciesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternalPoliciesServiceServer will
// result in compilation errors.
type UnsafeInternalPoliciesServiceServer interface {
	mustEmbedUnimplementedInternalPoliciesServiceServer()
}

func RegisterInternalPoliciesServiceServer(s grpc.ServiceRegistrar, srv InternalPoliciesServiceServer) {
	// If the following call pancis, it indicates UnimplementedInternalPoliciesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InternalPoliciesService_ServiceDesc, srv)
}

func _InternalPoliciesService_GetPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalPoliciesServiceServer).GetPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalPoliciesService_GetPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalPoliciesServiceServer).GetPolicy(ctx, req.(*GetPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalPoliciesService_CreatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalPoliciesServiceServer).CreatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalPoliciesService_CreatePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalPoliciesServiceServer).CreatePolicy(ctx, req.(*CreatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalPoliciesService_UpdatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalPoliciesServiceServer).UpdatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalPoliciesService_UpdatePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalPoliciesServiceServer).UpdatePolicy(ctx, req.(*UpdatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalPoliciesService_GetCompanyPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalPoliciesServiceServer).GetCompanyPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalPoliciesService_GetCompanyPolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalPoliciesServiceServer).GetCompanyPolicies(ctx, req.(*GetCompanyPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalPoliciesService_DeletePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalPoliciesServiceServer).DeletePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalPoliciesService_DeletePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalPoliciesServiceServer).DeletePolicy(ctx, req.(*DeletePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalPoliciesService_ReorderPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReorderPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalPoliciesServiceServer).ReorderPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalPoliciesService_ReorderPolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalPoliciesServiceServer).ReorderPolicies(ctx, req.(*ReorderPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalPoliciesService_BulkTestLogPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkTestLogPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalPoliciesServiceServer).BulkTestLogPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalPoliciesService_BulkTestLogPolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalPoliciesServiceServer).BulkTestLogPolicies(ctx, req.(*BulkTestLogPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalPoliciesService_TogglePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TogglePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalPoliciesServiceServer).TogglePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalPoliciesService_TogglePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalPoliciesServiceServer).TogglePolicy(ctx, req.(*TogglePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InternalPoliciesService_ServiceDesc is the grpc.ServiceDesc for InternalPoliciesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InternalPoliciesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.quota.v1.InternalPoliciesService",
	HandlerType: (*InternalPoliciesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPolicy",
			Handler:    _InternalPoliciesService_GetPolicy_Handler,
		},
		{
			MethodName: "CreatePolicy",
			Handler:    _InternalPoliciesService_CreatePolicy_Handler,
		},
		{
			MethodName: "UpdatePolicy",
			Handler:    _InternalPoliciesService_UpdatePolicy_Handler,
		},
		{
			MethodName: "GetCompanyPolicies",
			Handler:    _InternalPoliciesService_GetCompanyPolicies_Handler,
		},
		{
			MethodName: "DeletePolicy",
			Handler:    _InternalPoliciesService_DeletePolicy_Handler,
		},
		{
			MethodName: "ReorderPolicies",
			Handler:    _InternalPoliciesService_ReorderPolicies_Handler,
		},
		{
			MethodName: "BulkTestLogPolicies",
			Handler:    _InternalPoliciesService_BulkTestLogPolicies_Handler,
		},
		{
			MethodName: "TogglePolicy",
			Handler:    _InternalPoliciesService_TogglePolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/quota/v1/internal_policies_service.proto",
}
