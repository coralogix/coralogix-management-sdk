// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: com/coralogix/archive/dataset/v2/internal_dataset_service.proto

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
	InternalDatasetManagementService_GetAllDatasetRules_FullMethodName     = "/com.coralogix.archive.dataset.v2.InternalDatasetManagementService/GetAllDatasetRules"
	InternalDatasetManagementService_GetDatasetSchemaFields_FullMethodName = "/com.coralogix.archive.dataset.v2.InternalDatasetManagementService/GetDatasetSchemaFields"
	InternalDatasetManagementService_SetSchemaRule_FullMethodName          = "/com.coralogix.archive.dataset.v2.InternalDatasetManagementService/SetSchemaRule"
	InternalDatasetManagementService_GetSchemaRule_FullMethodName          = "/com.coralogix.archive.dataset.v2.InternalDatasetManagementService/GetSchemaRule"
	InternalDatasetManagementService_InternalWriteToDataset_FullMethodName = "/com.coralogix.archive.dataset.v2.InternalDatasetManagementService/InternalWriteToDataset"
)

// InternalDatasetManagementServiceClient is the client API for InternalDatasetManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalDatasetManagementServiceClient interface {
	GetAllDatasetRules(ctx context.Context, in *GetAllDatasetRulesRequest, opts ...grpc.CallOption) (*GetAllDatasetRulesResponse, error)
	GetDatasetSchemaFields(ctx context.Context, in *GetDatasetSchemaFieldsRequest, opts ...grpc.CallOption) (*GetDatasetSchemaFieldsResponse, error)
	SetSchemaRule(ctx context.Context, in *InternalDatasetManagementServiceSetSchemaRuleRequest, opts ...grpc.CallOption) (*InternalDatasetManagementServiceSetSchemaRuleResponse, error)
	GetSchemaRule(ctx context.Context, in *InternalDatasetManagementServiceGetSchemaRuleRequest, opts ...grpc.CallOption) (*InternalDatasetManagementServiceGetSchemaRuleResponse, error)
	InternalWriteToDataset(ctx context.Context, in *InternalWriteToDatasetRequest, opts ...grpc.CallOption) (*InternalWriteToDatasetResponse, error)
}

type internalDatasetManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalDatasetManagementServiceClient(cc grpc.ClientConnInterface) InternalDatasetManagementServiceClient {
	return &internalDatasetManagementServiceClient{cc}
}

func (c *internalDatasetManagementServiceClient) GetAllDatasetRules(ctx context.Context, in *GetAllDatasetRulesRequest, opts ...grpc.CallOption) (*GetAllDatasetRulesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllDatasetRulesResponse)
	err := c.cc.Invoke(ctx, InternalDatasetManagementService_GetAllDatasetRules_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalDatasetManagementServiceClient) GetDatasetSchemaFields(ctx context.Context, in *GetDatasetSchemaFieldsRequest, opts ...grpc.CallOption) (*GetDatasetSchemaFieldsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDatasetSchemaFieldsResponse)
	err := c.cc.Invoke(ctx, InternalDatasetManagementService_GetDatasetSchemaFields_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalDatasetManagementServiceClient) SetSchemaRule(ctx context.Context, in *InternalDatasetManagementServiceSetSchemaRuleRequest, opts ...grpc.CallOption) (*InternalDatasetManagementServiceSetSchemaRuleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InternalDatasetManagementServiceSetSchemaRuleResponse)
	err := c.cc.Invoke(ctx, InternalDatasetManagementService_SetSchemaRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalDatasetManagementServiceClient) GetSchemaRule(ctx context.Context, in *InternalDatasetManagementServiceGetSchemaRuleRequest, opts ...grpc.CallOption) (*InternalDatasetManagementServiceGetSchemaRuleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InternalDatasetManagementServiceGetSchemaRuleResponse)
	err := c.cc.Invoke(ctx, InternalDatasetManagementService_GetSchemaRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalDatasetManagementServiceClient) InternalWriteToDataset(ctx context.Context, in *InternalWriteToDatasetRequest, opts ...grpc.CallOption) (*InternalWriteToDatasetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InternalWriteToDatasetResponse)
	err := c.cc.Invoke(ctx, InternalDatasetManagementService_InternalWriteToDataset_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalDatasetManagementServiceServer is the server API for InternalDatasetManagementService service.
// All implementations must embed UnimplementedInternalDatasetManagementServiceServer
// for forward compatibility.
type InternalDatasetManagementServiceServer interface {
	GetAllDatasetRules(context.Context, *GetAllDatasetRulesRequest) (*GetAllDatasetRulesResponse, error)
	GetDatasetSchemaFields(context.Context, *GetDatasetSchemaFieldsRequest) (*GetDatasetSchemaFieldsResponse, error)
	SetSchemaRule(context.Context, *InternalDatasetManagementServiceSetSchemaRuleRequest) (*InternalDatasetManagementServiceSetSchemaRuleResponse, error)
	GetSchemaRule(context.Context, *InternalDatasetManagementServiceGetSchemaRuleRequest) (*InternalDatasetManagementServiceGetSchemaRuleResponse, error)
	InternalWriteToDataset(context.Context, *InternalWriteToDatasetRequest) (*InternalWriteToDatasetResponse, error)
	mustEmbedUnimplementedInternalDatasetManagementServiceServer()
}

// UnimplementedInternalDatasetManagementServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInternalDatasetManagementServiceServer struct{}

func (UnimplementedInternalDatasetManagementServiceServer) GetAllDatasetRules(context.Context, *GetAllDatasetRulesRequest) (*GetAllDatasetRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDatasetRules not implemented")
}
func (UnimplementedInternalDatasetManagementServiceServer) GetDatasetSchemaFields(context.Context, *GetDatasetSchemaFieldsRequest) (*GetDatasetSchemaFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetSchemaFields not implemented")
}
func (UnimplementedInternalDatasetManagementServiceServer) SetSchemaRule(context.Context, *InternalDatasetManagementServiceSetSchemaRuleRequest) (*InternalDatasetManagementServiceSetSchemaRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSchemaRule not implemented")
}
func (UnimplementedInternalDatasetManagementServiceServer) GetSchemaRule(context.Context, *InternalDatasetManagementServiceGetSchemaRuleRequest) (*InternalDatasetManagementServiceGetSchemaRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchemaRule not implemented")
}
func (UnimplementedInternalDatasetManagementServiceServer) InternalWriteToDataset(context.Context, *InternalWriteToDatasetRequest) (*InternalWriteToDatasetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InternalWriteToDataset not implemented")
}
func (UnimplementedInternalDatasetManagementServiceServer) mustEmbedUnimplementedInternalDatasetManagementServiceServer() {
}
func (UnimplementedInternalDatasetManagementServiceServer) testEmbeddedByValue() {}

// UnsafeInternalDatasetManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternalDatasetManagementServiceServer will
// result in compilation errors.
type UnsafeInternalDatasetManagementServiceServer interface {
	mustEmbedUnimplementedInternalDatasetManagementServiceServer()
}

func RegisterInternalDatasetManagementServiceServer(s grpc.ServiceRegistrar, srv InternalDatasetManagementServiceServer) {
	// If the following call pancis, it indicates UnimplementedInternalDatasetManagementServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InternalDatasetManagementService_ServiceDesc, srv)
}

func _InternalDatasetManagementService_GetAllDatasetRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllDatasetRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalDatasetManagementServiceServer).GetAllDatasetRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalDatasetManagementService_GetAllDatasetRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalDatasetManagementServiceServer).GetAllDatasetRules(ctx, req.(*GetAllDatasetRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalDatasetManagementService_GetDatasetSchemaFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetSchemaFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalDatasetManagementServiceServer).GetDatasetSchemaFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalDatasetManagementService_GetDatasetSchemaFields_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalDatasetManagementServiceServer).GetDatasetSchemaFields(ctx, req.(*GetDatasetSchemaFieldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalDatasetManagementService_SetSchemaRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalDatasetManagementServiceSetSchemaRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalDatasetManagementServiceServer).SetSchemaRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalDatasetManagementService_SetSchemaRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalDatasetManagementServiceServer).SetSchemaRule(ctx, req.(*InternalDatasetManagementServiceSetSchemaRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalDatasetManagementService_GetSchemaRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalDatasetManagementServiceGetSchemaRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalDatasetManagementServiceServer).GetSchemaRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalDatasetManagementService_GetSchemaRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalDatasetManagementServiceServer).GetSchemaRule(ctx, req.(*InternalDatasetManagementServiceGetSchemaRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalDatasetManagementService_InternalWriteToDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalWriteToDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalDatasetManagementServiceServer).InternalWriteToDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalDatasetManagementService_InternalWriteToDataset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalDatasetManagementServiceServer).InternalWriteToDataset(ctx, req.(*InternalWriteToDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InternalDatasetManagementService_ServiceDesc is the grpc.ServiceDesc for InternalDatasetManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InternalDatasetManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.archive.dataset.v2.InternalDatasetManagementService",
	HandlerType: (*InternalDatasetManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllDatasetRules",
			Handler:    _InternalDatasetManagementService_GetAllDatasetRules_Handler,
		},
		{
			MethodName: "GetDatasetSchemaFields",
			Handler:    _InternalDatasetManagementService_GetDatasetSchemaFields_Handler,
		},
		{
			MethodName: "SetSchemaRule",
			Handler:    _InternalDatasetManagementService_SetSchemaRule_Handler,
		},
		{
			MethodName: "GetSchemaRule",
			Handler:    _InternalDatasetManagementService_GetSchemaRule_Handler,
		},
		{
			MethodName: "InternalWriteToDataset",
			Handler:    _InternalDatasetManagementService_InternalWriteToDataset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/archive/dataset/v2/internal_dataset_service.proto",
}
