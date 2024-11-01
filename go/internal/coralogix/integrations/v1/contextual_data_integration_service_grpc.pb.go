// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.28.3
// source: com/coralogix/integrations/v1/contextual_data_integration_service.proto

package v1

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
	ContextualDataIntegrationService_GetContextualDataIntegrations_FullMethodName          = "/com.coralogix.integrations.v1.ContextualDataIntegrationService/GetContextualDataIntegrations"
	ContextualDataIntegrationService_GetContextualDataIntegrationDefinition_FullMethodName = "/com.coralogix.integrations.v1.ContextualDataIntegrationService/GetContextualDataIntegrationDefinition"
	ContextualDataIntegrationService_GetContextualDataIntegrationDetails_FullMethodName    = "/com.coralogix.integrations.v1.ContextualDataIntegrationService/GetContextualDataIntegrationDetails"
	ContextualDataIntegrationService_TestContextualDataIntegration_FullMethodName          = "/com.coralogix.integrations.v1.ContextualDataIntegrationService/TestContextualDataIntegration"
	ContextualDataIntegrationService_SaveContextualDataIntegration_FullMethodName          = "/com.coralogix.integrations.v1.ContextualDataIntegrationService/SaveContextualDataIntegration"
	ContextualDataIntegrationService_UpdateContextualDataIntegration_FullMethodName        = "/com.coralogix.integrations.v1.ContextualDataIntegrationService/UpdateContextualDataIntegration"
	ContextualDataIntegrationService_DeleteContextualDataIntegration_FullMethodName        = "/com.coralogix.integrations.v1.ContextualDataIntegrationService/DeleteContextualDataIntegration"
)

// ContextualDataIntegrationServiceClient is the client API for ContextualDataIntegrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContextualDataIntegrationServiceClient interface {
	GetContextualDataIntegrations(ctx context.Context, in *GetContextualDataIntegrationsRequest, opts ...grpc.CallOption) (*GetContextualDataIntegrationsResponse, error)
	GetContextualDataIntegrationDefinition(ctx context.Context, in *GetContextualDataIntegrationDefinitionRequest, opts ...grpc.CallOption) (*GetContextualDataIntegrationDefinitionResponse, error)
	GetContextualDataIntegrationDetails(ctx context.Context, in *GetContextualDataIntegrationDetailsRequest, opts ...grpc.CallOption) (*GetContextualDataIntegrationDetailsResponse, error)
	TestContextualDataIntegration(ctx context.Context, in *TestContextualDataIntegrationRequest, opts ...grpc.CallOption) (*TestContextualDataIntegrationResponse, error)
	SaveContextualDataIntegration(ctx context.Context, in *SaveContextualDataIntegrationRequest, opts ...grpc.CallOption) (*SaveContextualDataIntegrationResponse, error)
	UpdateContextualDataIntegration(ctx context.Context, in *UpdateContextualDataIntegrationRequest, opts ...grpc.CallOption) (*UpdateContextualDataIntegrationResponse, error)
	DeleteContextualDataIntegration(ctx context.Context, in *DeleteContextualDataIntegrationRequest, opts ...grpc.CallOption) (*DeleteContextualDataIntegrationResponse, error)
}

type contextualDataIntegrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContextualDataIntegrationServiceClient(cc grpc.ClientConnInterface) ContextualDataIntegrationServiceClient {
	return &contextualDataIntegrationServiceClient{cc}
}

func (c *contextualDataIntegrationServiceClient) GetContextualDataIntegrations(ctx context.Context, in *GetContextualDataIntegrationsRequest, opts ...grpc.CallOption) (*GetContextualDataIntegrationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetContextualDataIntegrationsResponse)
	err := c.cc.Invoke(ctx, ContextualDataIntegrationService_GetContextualDataIntegrations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextualDataIntegrationServiceClient) GetContextualDataIntegrationDefinition(ctx context.Context, in *GetContextualDataIntegrationDefinitionRequest, opts ...grpc.CallOption) (*GetContextualDataIntegrationDefinitionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetContextualDataIntegrationDefinitionResponse)
	err := c.cc.Invoke(ctx, ContextualDataIntegrationService_GetContextualDataIntegrationDefinition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextualDataIntegrationServiceClient) GetContextualDataIntegrationDetails(ctx context.Context, in *GetContextualDataIntegrationDetailsRequest, opts ...grpc.CallOption) (*GetContextualDataIntegrationDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetContextualDataIntegrationDetailsResponse)
	err := c.cc.Invoke(ctx, ContextualDataIntegrationService_GetContextualDataIntegrationDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextualDataIntegrationServiceClient) TestContextualDataIntegration(ctx context.Context, in *TestContextualDataIntegrationRequest, opts ...grpc.CallOption) (*TestContextualDataIntegrationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestContextualDataIntegrationResponse)
	err := c.cc.Invoke(ctx, ContextualDataIntegrationService_TestContextualDataIntegration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextualDataIntegrationServiceClient) SaveContextualDataIntegration(ctx context.Context, in *SaveContextualDataIntegrationRequest, opts ...grpc.CallOption) (*SaveContextualDataIntegrationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveContextualDataIntegrationResponse)
	err := c.cc.Invoke(ctx, ContextualDataIntegrationService_SaveContextualDataIntegration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextualDataIntegrationServiceClient) UpdateContextualDataIntegration(ctx context.Context, in *UpdateContextualDataIntegrationRequest, opts ...grpc.CallOption) (*UpdateContextualDataIntegrationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateContextualDataIntegrationResponse)
	err := c.cc.Invoke(ctx, ContextualDataIntegrationService_UpdateContextualDataIntegration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contextualDataIntegrationServiceClient) DeleteContextualDataIntegration(ctx context.Context, in *DeleteContextualDataIntegrationRequest, opts ...grpc.CallOption) (*DeleteContextualDataIntegrationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteContextualDataIntegrationResponse)
	err := c.cc.Invoke(ctx, ContextualDataIntegrationService_DeleteContextualDataIntegration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContextualDataIntegrationServiceServer is the server API for ContextualDataIntegrationService service.
// All implementations must embed UnimplementedContextualDataIntegrationServiceServer
// for forward compatibility
type ContextualDataIntegrationServiceServer interface {
	GetContextualDataIntegrations(context.Context, *GetContextualDataIntegrationsRequest) (*GetContextualDataIntegrationsResponse, error)
	GetContextualDataIntegrationDefinition(context.Context, *GetContextualDataIntegrationDefinitionRequest) (*GetContextualDataIntegrationDefinitionResponse, error)
	GetContextualDataIntegrationDetails(context.Context, *GetContextualDataIntegrationDetailsRequest) (*GetContextualDataIntegrationDetailsResponse, error)
	TestContextualDataIntegration(context.Context, *TestContextualDataIntegrationRequest) (*TestContextualDataIntegrationResponse, error)
	SaveContextualDataIntegration(context.Context, *SaveContextualDataIntegrationRequest) (*SaveContextualDataIntegrationResponse, error)
	UpdateContextualDataIntegration(context.Context, *UpdateContextualDataIntegrationRequest) (*UpdateContextualDataIntegrationResponse, error)
	DeleteContextualDataIntegration(context.Context, *DeleteContextualDataIntegrationRequest) (*DeleteContextualDataIntegrationResponse, error)
	mustEmbedUnimplementedContextualDataIntegrationServiceServer()
}

// UnimplementedContextualDataIntegrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContextualDataIntegrationServiceServer struct {
}

func (UnimplementedContextualDataIntegrationServiceServer) GetContextualDataIntegrations(context.Context, *GetContextualDataIntegrationsRequest) (*GetContextualDataIntegrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContextualDataIntegrations not implemented")
}
func (UnimplementedContextualDataIntegrationServiceServer) GetContextualDataIntegrationDefinition(context.Context, *GetContextualDataIntegrationDefinitionRequest) (*GetContextualDataIntegrationDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContextualDataIntegrationDefinition not implemented")
}
func (UnimplementedContextualDataIntegrationServiceServer) GetContextualDataIntegrationDetails(context.Context, *GetContextualDataIntegrationDetailsRequest) (*GetContextualDataIntegrationDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContextualDataIntegrationDetails not implemented")
}
func (UnimplementedContextualDataIntegrationServiceServer) TestContextualDataIntegration(context.Context, *TestContextualDataIntegrationRequest) (*TestContextualDataIntegrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestContextualDataIntegration not implemented")
}
func (UnimplementedContextualDataIntegrationServiceServer) SaveContextualDataIntegration(context.Context, *SaveContextualDataIntegrationRequest) (*SaveContextualDataIntegrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveContextualDataIntegration not implemented")
}
func (UnimplementedContextualDataIntegrationServiceServer) UpdateContextualDataIntegration(context.Context, *UpdateContextualDataIntegrationRequest) (*UpdateContextualDataIntegrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContextualDataIntegration not implemented")
}
func (UnimplementedContextualDataIntegrationServiceServer) DeleteContextualDataIntegration(context.Context, *DeleteContextualDataIntegrationRequest) (*DeleteContextualDataIntegrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContextualDataIntegration not implemented")
}
func (UnimplementedContextualDataIntegrationServiceServer) mustEmbedUnimplementedContextualDataIntegrationServiceServer() {
}

// UnsafeContextualDataIntegrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContextualDataIntegrationServiceServer will
// result in compilation errors.
type UnsafeContextualDataIntegrationServiceServer interface {
	mustEmbedUnimplementedContextualDataIntegrationServiceServer()
}

func RegisterContextualDataIntegrationServiceServer(s grpc.ServiceRegistrar, srv ContextualDataIntegrationServiceServer) {
	s.RegisterService(&ContextualDataIntegrationService_ServiceDesc, srv)
}

func _ContextualDataIntegrationService_GetContextualDataIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContextualDataIntegrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextualDataIntegrationServiceServer).GetContextualDataIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextualDataIntegrationService_GetContextualDataIntegrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextualDataIntegrationServiceServer).GetContextualDataIntegrations(ctx, req.(*GetContextualDataIntegrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextualDataIntegrationService_GetContextualDataIntegrationDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContextualDataIntegrationDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextualDataIntegrationServiceServer).GetContextualDataIntegrationDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextualDataIntegrationService_GetContextualDataIntegrationDefinition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextualDataIntegrationServiceServer).GetContextualDataIntegrationDefinition(ctx, req.(*GetContextualDataIntegrationDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextualDataIntegrationService_GetContextualDataIntegrationDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContextualDataIntegrationDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextualDataIntegrationServiceServer).GetContextualDataIntegrationDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextualDataIntegrationService_GetContextualDataIntegrationDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextualDataIntegrationServiceServer).GetContextualDataIntegrationDetails(ctx, req.(*GetContextualDataIntegrationDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextualDataIntegrationService_TestContextualDataIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestContextualDataIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextualDataIntegrationServiceServer).TestContextualDataIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextualDataIntegrationService_TestContextualDataIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextualDataIntegrationServiceServer).TestContextualDataIntegration(ctx, req.(*TestContextualDataIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextualDataIntegrationService_SaveContextualDataIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveContextualDataIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextualDataIntegrationServiceServer).SaveContextualDataIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextualDataIntegrationService_SaveContextualDataIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextualDataIntegrationServiceServer).SaveContextualDataIntegration(ctx, req.(*SaveContextualDataIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextualDataIntegrationService_UpdateContextualDataIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContextualDataIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextualDataIntegrationServiceServer).UpdateContextualDataIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextualDataIntegrationService_UpdateContextualDataIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextualDataIntegrationServiceServer).UpdateContextualDataIntegration(ctx, req.(*UpdateContextualDataIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContextualDataIntegrationService_DeleteContextualDataIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContextualDataIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextualDataIntegrationServiceServer).DeleteContextualDataIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContextualDataIntegrationService_DeleteContextualDataIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextualDataIntegrationServiceServer).DeleteContextualDataIntegration(ctx, req.(*DeleteContextualDataIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContextualDataIntegrationService_ServiceDesc is the grpc.ServiceDesc for ContextualDataIntegrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContextualDataIntegrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.integrations.v1.ContextualDataIntegrationService",
	HandlerType: (*ContextualDataIntegrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContextualDataIntegrations",
			Handler:    _ContextualDataIntegrationService_GetContextualDataIntegrations_Handler,
		},
		{
			MethodName: "GetContextualDataIntegrationDefinition",
			Handler:    _ContextualDataIntegrationService_GetContextualDataIntegrationDefinition_Handler,
		},
		{
			MethodName: "GetContextualDataIntegrationDetails",
			Handler:    _ContextualDataIntegrationService_GetContextualDataIntegrationDetails_Handler,
		},
		{
			MethodName: "TestContextualDataIntegration",
			Handler:    _ContextualDataIntegrationService_TestContextualDataIntegration_Handler,
		},
		{
			MethodName: "SaveContextualDataIntegration",
			Handler:    _ContextualDataIntegrationService_SaveContextualDataIntegration_Handler,
		},
		{
			MethodName: "UpdateContextualDataIntegration",
			Handler:    _ContextualDataIntegrationService_UpdateContextualDataIntegration_Handler,
		},
		{
			MethodName: "DeleteContextualDataIntegration",
			Handler:    _ContextualDataIntegrationService_DeleteContextualDataIntegration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/integrations/v1/contextual_data_integration_service.proto",
}
