// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: com/coralogix/quota/v1/policies_service.proto

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
	PoliciesService_GetPolicy_FullMethodName                   = "/com.coralogix.quota.v1.PoliciesService/GetPolicy"
	PoliciesService_CreatePolicy_FullMethodName                = "/com.coralogix.quota.v1.PoliciesService/CreatePolicy"
	PoliciesService_UpdatePolicy_FullMethodName                = "/com.coralogix.quota.v1.PoliciesService/UpdatePolicy"
	PoliciesService_GetCompanyPolicies_FullMethodName          = "/com.coralogix.quota.v1.PoliciesService/GetCompanyPolicies"
	PoliciesService_DeletePolicy_FullMethodName                = "/com.coralogix.quota.v1.PoliciesService/DeletePolicy"
	PoliciesService_ReorderPolicies_FullMethodName             = "/com.coralogix.quota.v1.PoliciesService/ReorderPolicies"
	PoliciesService_BulkTestLogPolicies_FullMethodName         = "/com.coralogix.quota.v1.PoliciesService/BulkTestLogPolicies"
	PoliciesService_TogglePolicy_FullMethodName                = "/com.coralogix.quota.v1.PoliciesService/TogglePolicy"
	PoliciesService_AtomicBatchCreatePolicy_FullMethodName     = "/com.coralogix.quota.v1.PoliciesService/AtomicBatchCreatePolicy"
	PoliciesService_AtomicOverwriteLogPolicies_FullMethodName  = "/com.coralogix.quota.v1.PoliciesService/AtomicOverwriteLogPolicies"
	PoliciesService_AtomicOverwriteSpanPolicies_FullMethodName = "/com.coralogix.quota.v1.PoliciesService/AtomicOverwriteSpanPolicies"
)

// PoliciesServiceClient is the client API for PoliciesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PoliciesServiceClient interface {
	GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*GetPolicyResponse, error)
	CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*CreatePolicyResponse, error)
	UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*UpdatePolicyResponse, error)
	GetCompanyPolicies(ctx context.Context, in *GetCompanyPoliciesRequest, opts ...grpc.CallOption) (*GetCompanyPoliciesResponse, error)
	DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*DeletePolicyResponse, error)
	ReorderPolicies(ctx context.Context, in *ReorderPoliciesRequest, opts ...grpc.CallOption) (*ReorderPoliciesResponse, error)
	BulkTestLogPolicies(ctx context.Context, in *BulkTestLogPoliciesRequest, opts ...grpc.CallOption) (*BulkTestLogPoliciesResponse, error)
	TogglePolicy(ctx context.Context, in *TogglePolicyRequest, opts ...grpc.CallOption) (*TogglePolicyResponse, error)
	AtomicBatchCreatePolicy(ctx context.Context, in *AtomicBatchCreatePolicyRequest, opts ...grpc.CallOption) (*AtomicBatchCreatePolicyResponse, error)
	AtomicOverwriteLogPolicies(ctx context.Context, in *AtomicOverwriteLogPoliciesRequest, opts ...grpc.CallOption) (*AtomicOverwriteLogPoliciesResponse, error)
	AtomicOverwriteSpanPolicies(ctx context.Context, in *AtomicOverwriteSpanPoliciesRequest, opts ...grpc.CallOption) (*AtomicOverwriteSpanPoliciesResponse, error)
}

type policiesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPoliciesServiceClient(cc grpc.ClientConnInterface) PoliciesServiceClient {
	return &policiesServiceClient{cc}
}

func (c *policiesServiceClient) GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*GetPolicyResponse, error) {
	out := new(GetPolicyResponse)
	err := c.cc.Invoke(ctx, PoliciesService_GetPolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policiesServiceClient) CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*CreatePolicyResponse, error) {
	out := new(CreatePolicyResponse)
	err := c.cc.Invoke(ctx, PoliciesService_CreatePolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policiesServiceClient) UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*UpdatePolicyResponse, error) {
	out := new(UpdatePolicyResponse)
	err := c.cc.Invoke(ctx, PoliciesService_UpdatePolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policiesServiceClient) GetCompanyPolicies(ctx context.Context, in *GetCompanyPoliciesRequest, opts ...grpc.CallOption) (*GetCompanyPoliciesResponse, error) {
	out := new(GetCompanyPoliciesResponse)
	err := c.cc.Invoke(ctx, PoliciesService_GetCompanyPolicies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policiesServiceClient) DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*DeletePolicyResponse, error) {
	out := new(DeletePolicyResponse)
	err := c.cc.Invoke(ctx, PoliciesService_DeletePolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policiesServiceClient) ReorderPolicies(ctx context.Context, in *ReorderPoliciesRequest, opts ...grpc.CallOption) (*ReorderPoliciesResponse, error) {
	out := new(ReorderPoliciesResponse)
	err := c.cc.Invoke(ctx, PoliciesService_ReorderPolicies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policiesServiceClient) BulkTestLogPolicies(ctx context.Context, in *BulkTestLogPoliciesRequest, opts ...grpc.CallOption) (*BulkTestLogPoliciesResponse, error) {
	out := new(BulkTestLogPoliciesResponse)
	err := c.cc.Invoke(ctx, PoliciesService_BulkTestLogPolicies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policiesServiceClient) TogglePolicy(ctx context.Context, in *TogglePolicyRequest, opts ...grpc.CallOption) (*TogglePolicyResponse, error) {
	out := new(TogglePolicyResponse)
	err := c.cc.Invoke(ctx, PoliciesService_TogglePolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policiesServiceClient) AtomicBatchCreatePolicy(ctx context.Context, in *AtomicBatchCreatePolicyRequest, opts ...grpc.CallOption) (*AtomicBatchCreatePolicyResponse, error) {
	out := new(AtomicBatchCreatePolicyResponse)
	err := c.cc.Invoke(ctx, PoliciesService_AtomicBatchCreatePolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policiesServiceClient) AtomicOverwriteLogPolicies(ctx context.Context, in *AtomicOverwriteLogPoliciesRequest, opts ...grpc.CallOption) (*AtomicOverwriteLogPoliciesResponse, error) {
	out := new(AtomicOverwriteLogPoliciesResponse)
	err := c.cc.Invoke(ctx, PoliciesService_AtomicOverwriteLogPolicies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policiesServiceClient) AtomicOverwriteSpanPolicies(ctx context.Context, in *AtomicOverwriteSpanPoliciesRequest, opts ...grpc.CallOption) (*AtomicOverwriteSpanPoliciesResponse, error) {
	out := new(AtomicOverwriteSpanPoliciesResponse)
	err := c.cc.Invoke(ctx, PoliciesService_AtomicOverwriteSpanPolicies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PoliciesServiceServer is the server API for PoliciesService service.
// All implementations must embed UnimplementedPoliciesServiceServer
// for forward compatibility
type PoliciesServiceServer interface {
	GetPolicy(context.Context, *GetPolicyRequest) (*GetPolicyResponse, error)
	CreatePolicy(context.Context, *CreatePolicyRequest) (*CreatePolicyResponse, error)
	UpdatePolicy(context.Context, *UpdatePolicyRequest) (*UpdatePolicyResponse, error)
	GetCompanyPolicies(context.Context, *GetCompanyPoliciesRequest) (*GetCompanyPoliciesResponse, error)
	DeletePolicy(context.Context, *DeletePolicyRequest) (*DeletePolicyResponse, error)
	ReorderPolicies(context.Context, *ReorderPoliciesRequest) (*ReorderPoliciesResponse, error)
	BulkTestLogPolicies(context.Context, *BulkTestLogPoliciesRequest) (*BulkTestLogPoliciesResponse, error)
	TogglePolicy(context.Context, *TogglePolicyRequest) (*TogglePolicyResponse, error)
	AtomicBatchCreatePolicy(context.Context, *AtomicBatchCreatePolicyRequest) (*AtomicBatchCreatePolicyResponse, error)
	AtomicOverwriteLogPolicies(context.Context, *AtomicOverwriteLogPoliciesRequest) (*AtomicOverwriteLogPoliciesResponse, error)
	AtomicOverwriteSpanPolicies(context.Context, *AtomicOverwriteSpanPoliciesRequest) (*AtomicOverwriteSpanPoliciesResponse, error)
	mustEmbedUnimplementedPoliciesServiceServer()
}

// UnimplementedPoliciesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPoliciesServiceServer struct {
}

func (UnimplementedPoliciesServiceServer) GetPolicy(context.Context, *GetPolicyRequest) (*GetPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicy not implemented")
}
func (UnimplementedPoliciesServiceServer) CreatePolicy(context.Context, *CreatePolicyRequest) (*CreatePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicy not implemented")
}
func (UnimplementedPoliciesServiceServer) UpdatePolicy(context.Context, *UpdatePolicyRequest) (*UpdatePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePolicy not implemented")
}
func (UnimplementedPoliciesServiceServer) GetCompanyPolicies(context.Context, *GetCompanyPoliciesRequest) (*GetCompanyPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyPolicies not implemented")
}
func (UnimplementedPoliciesServiceServer) DeletePolicy(context.Context, *DeletePolicyRequest) (*DeletePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicy not implemented")
}
func (UnimplementedPoliciesServiceServer) ReorderPolicies(context.Context, *ReorderPoliciesRequest) (*ReorderPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReorderPolicies not implemented")
}
func (UnimplementedPoliciesServiceServer) BulkTestLogPolicies(context.Context, *BulkTestLogPoliciesRequest) (*BulkTestLogPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkTestLogPolicies not implemented")
}
func (UnimplementedPoliciesServiceServer) TogglePolicy(context.Context, *TogglePolicyRequest) (*TogglePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TogglePolicy not implemented")
}
func (UnimplementedPoliciesServiceServer) AtomicBatchCreatePolicy(context.Context, *AtomicBatchCreatePolicyRequest) (*AtomicBatchCreatePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtomicBatchCreatePolicy not implemented")
}
func (UnimplementedPoliciesServiceServer) AtomicOverwriteLogPolicies(context.Context, *AtomicOverwriteLogPoliciesRequest) (*AtomicOverwriteLogPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtomicOverwriteLogPolicies not implemented")
}
func (UnimplementedPoliciesServiceServer) AtomicOverwriteSpanPolicies(context.Context, *AtomicOverwriteSpanPoliciesRequest) (*AtomicOverwriteSpanPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtomicOverwriteSpanPolicies not implemented")
}
func (UnimplementedPoliciesServiceServer) mustEmbedUnimplementedPoliciesServiceServer() {}

// UnsafePoliciesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PoliciesServiceServer will
// result in compilation errors.
type UnsafePoliciesServiceServer interface {
	mustEmbedUnimplementedPoliciesServiceServer()
}

func RegisterPoliciesServiceServer(s grpc.ServiceRegistrar, srv PoliciesServiceServer) {
	s.RegisterService(&PoliciesService_ServiceDesc, srv)
}

func _PoliciesService_GetPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoliciesServiceServer).GetPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoliciesService_GetPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoliciesServiceServer).GetPolicy(ctx, req.(*GetPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoliciesService_CreatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoliciesServiceServer).CreatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoliciesService_CreatePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoliciesServiceServer).CreatePolicy(ctx, req.(*CreatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoliciesService_UpdatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoliciesServiceServer).UpdatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoliciesService_UpdatePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoliciesServiceServer).UpdatePolicy(ctx, req.(*UpdatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoliciesService_GetCompanyPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoliciesServiceServer).GetCompanyPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoliciesService_GetCompanyPolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoliciesServiceServer).GetCompanyPolicies(ctx, req.(*GetCompanyPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoliciesService_DeletePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoliciesServiceServer).DeletePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoliciesService_DeletePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoliciesServiceServer).DeletePolicy(ctx, req.(*DeletePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoliciesService_ReorderPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReorderPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoliciesServiceServer).ReorderPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoliciesService_ReorderPolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoliciesServiceServer).ReorderPolicies(ctx, req.(*ReorderPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoliciesService_BulkTestLogPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkTestLogPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoliciesServiceServer).BulkTestLogPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoliciesService_BulkTestLogPolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoliciesServiceServer).BulkTestLogPolicies(ctx, req.(*BulkTestLogPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoliciesService_TogglePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TogglePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoliciesServiceServer).TogglePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoliciesService_TogglePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoliciesServiceServer).TogglePolicy(ctx, req.(*TogglePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoliciesService_AtomicBatchCreatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AtomicBatchCreatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoliciesServiceServer).AtomicBatchCreatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoliciesService_AtomicBatchCreatePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoliciesServiceServer).AtomicBatchCreatePolicy(ctx, req.(*AtomicBatchCreatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoliciesService_AtomicOverwriteLogPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AtomicOverwriteLogPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoliciesServiceServer).AtomicOverwriteLogPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoliciesService_AtomicOverwriteLogPolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoliciesServiceServer).AtomicOverwriteLogPolicies(ctx, req.(*AtomicOverwriteLogPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoliciesService_AtomicOverwriteSpanPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AtomicOverwriteSpanPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoliciesServiceServer).AtomicOverwriteSpanPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoliciesService_AtomicOverwriteSpanPolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoliciesServiceServer).AtomicOverwriteSpanPolicies(ctx, req.(*AtomicOverwriteSpanPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PoliciesService_ServiceDesc is the grpc.ServiceDesc for PoliciesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PoliciesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.quota.v1.PoliciesService",
	HandlerType: (*PoliciesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPolicy",
			Handler:    _PoliciesService_GetPolicy_Handler,
		},
		{
			MethodName: "CreatePolicy",
			Handler:    _PoliciesService_CreatePolicy_Handler,
		},
		{
			MethodName: "UpdatePolicy",
			Handler:    _PoliciesService_UpdatePolicy_Handler,
		},
		{
			MethodName: "GetCompanyPolicies",
			Handler:    _PoliciesService_GetCompanyPolicies_Handler,
		},
		{
			MethodName: "DeletePolicy",
			Handler:    _PoliciesService_DeletePolicy_Handler,
		},
		{
			MethodName: "ReorderPolicies",
			Handler:    _PoliciesService_ReorderPolicies_Handler,
		},
		{
			MethodName: "BulkTestLogPolicies",
			Handler:    _PoliciesService_BulkTestLogPolicies_Handler,
		},
		{
			MethodName: "TogglePolicy",
			Handler:    _PoliciesService_TogglePolicy_Handler,
		},
		{
			MethodName: "AtomicBatchCreatePolicy",
			Handler:    _PoliciesService_AtomicBatchCreatePolicy_Handler,
		},
		{
			MethodName: "AtomicOverwriteLogPolicies",
			Handler:    _PoliciesService_AtomicOverwriteLogPolicies_Handler,
		},
		{
			MethodName: "AtomicOverwriteSpanPolicies",
			Handler:    _PoliciesService_AtomicOverwriteSpanPolicies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/quota/v1/policies_service.proto",
}
