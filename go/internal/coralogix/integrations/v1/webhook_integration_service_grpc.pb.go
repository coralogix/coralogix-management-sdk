// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: com/coralogix/integrations/v1/webhook_integration_service.proto

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
	WebhookIntegrationService_CreateWebhookIntegration_FullMethodName           = "/com.coralogix.integrations.v1.WebhookIntegrationService/CreateWebhookIntegration"
	WebhookIntegrationService_ListWebhookIntegrations_FullMethodName            = "/com.coralogix.integrations.v1.WebhookIntegrationService/ListWebhookIntegrations"
	WebhookIntegrationService_DeleteWebhookIntegration_FullMethodName           = "/com.coralogix.integrations.v1.WebhookIntegrationService/DeleteWebhookIntegration"
	WebhookIntegrationService_ToggleWebhookIntegrationActivation_FullMethodName = "/com.coralogix.integrations.v1.WebhookIntegrationService/ToggleWebhookIntegrationActivation"
	WebhookIntegrationService_CountWebhookIntegrations_FullMethodName           = "/com.coralogix.integrations.v1.WebhookIntegrationService/CountWebhookIntegrations"
)

// WebhookIntegrationServiceClient is the client API for WebhookIntegrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebhookIntegrationServiceClient interface {
	CreateWebhookIntegration(ctx context.Context, in *CreateWebhookIntegrationRequest, opts ...grpc.CallOption) (*CreateWebhookIntegrationResponse, error)
	ListWebhookIntegrations(ctx context.Context, in *ListWebhookIntegrationsRequest, opts ...grpc.CallOption) (*ListWebhookIntegrationsResponse, error)
	DeleteWebhookIntegration(ctx context.Context, in *DeleteWebhookIntegrationRequest, opts ...grpc.CallOption) (*DeleteWebhookIntegrationResponse, error)
	ToggleWebhookIntegrationActivation(ctx context.Context, in *ToggleWebhookIntegrationActivationRequest, opts ...grpc.CallOption) (*ToggleWebhookIntegrationActivationResponse, error)
	CountWebhookIntegrations(ctx context.Context, in *CountWebhookIntegrationsRequest, opts ...grpc.CallOption) (*CountWebhookIntegrationsResponse, error)
}

type webhookIntegrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebhookIntegrationServiceClient(cc grpc.ClientConnInterface) WebhookIntegrationServiceClient {
	return &webhookIntegrationServiceClient{cc}
}

func (c *webhookIntegrationServiceClient) CreateWebhookIntegration(ctx context.Context, in *CreateWebhookIntegrationRequest, opts ...grpc.CallOption) (*CreateWebhookIntegrationResponse, error) {
	out := new(CreateWebhookIntegrationResponse)
	err := c.cc.Invoke(ctx, WebhookIntegrationService_CreateWebhookIntegration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookIntegrationServiceClient) ListWebhookIntegrations(ctx context.Context, in *ListWebhookIntegrationsRequest, opts ...grpc.CallOption) (*ListWebhookIntegrationsResponse, error) {
	out := new(ListWebhookIntegrationsResponse)
	err := c.cc.Invoke(ctx, WebhookIntegrationService_ListWebhookIntegrations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookIntegrationServiceClient) DeleteWebhookIntegration(ctx context.Context, in *DeleteWebhookIntegrationRequest, opts ...grpc.CallOption) (*DeleteWebhookIntegrationResponse, error) {
	out := new(DeleteWebhookIntegrationResponse)
	err := c.cc.Invoke(ctx, WebhookIntegrationService_DeleteWebhookIntegration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookIntegrationServiceClient) ToggleWebhookIntegrationActivation(ctx context.Context, in *ToggleWebhookIntegrationActivationRequest, opts ...grpc.CallOption) (*ToggleWebhookIntegrationActivationResponse, error) {
	out := new(ToggleWebhookIntegrationActivationResponse)
	err := c.cc.Invoke(ctx, WebhookIntegrationService_ToggleWebhookIntegrationActivation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookIntegrationServiceClient) CountWebhookIntegrations(ctx context.Context, in *CountWebhookIntegrationsRequest, opts ...grpc.CallOption) (*CountWebhookIntegrationsResponse, error) {
	out := new(CountWebhookIntegrationsResponse)
	err := c.cc.Invoke(ctx, WebhookIntegrationService_CountWebhookIntegrations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebhookIntegrationServiceServer is the server API for WebhookIntegrationService service.
// All implementations must embed UnimplementedWebhookIntegrationServiceServer
// for forward compatibility
type WebhookIntegrationServiceServer interface {
	CreateWebhookIntegration(context.Context, *CreateWebhookIntegrationRequest) (*CreateWebhookIntegrationResponse, error)
	ListWebhookIntegrations(context.Context, *ListWebhookIntegrationsRequest) (*ListWebhookIntegrationsResponse, error)
	DeleteWebhookIntegration(context.Context, *DeleteWebhookIntegrationRequest) (*DeleteWebhookIntegrationResponse, error)
	ToggleWebhookIntegrationActivation(context.Context, *ToggleWebhookIntegrationActivationRequest) (*ToggleWebhookIntegrationActivationResponse, error)
	CountWebhookIntegrations(context.Context, *CountWebhookIntegrationsRequest) (*CountWebhookIntegrationsResponse, error)
	mustEmbedUnimplementedWebhookIntegrationServiceServer()
}

// UnimplementedWebhookIntegrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWebhookIntegrationServiceServer struct {
}

func (UnimplementedWebhookIntegrationServiceServer) CreateWebhookIntegration(context.Context, *CreateWebhookIntegrationRequest) (*CreateWebhookIntegrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWebhookIntegration not implemented")
}
func (UnimplementedWebhookIntegrationServiceServer) ListWebhookIntegrations(context.Context, *ListWebhookIntegrationsRequest) (*ListWebhookIntegrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWebhookIntegrations not implemented")
}
func (UnimplementedWebhookIntegrationServiceServer) DeleteWebhookIntegration(context.Context, *DeleteWebhookIntegrationRequest) (*DeleteWebhookIntegrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWebhookIntegration not implemented")
}
func (UnimplementedWebhookIntegrationServiceServer) ToggleWebhookIntegrationActivation(context.Context, *ToggleWebhookIntegrationActivationRequest) (*ToggleWebhookIntegrationActivationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleWebhookIntegrationActivation not implemented")
}
func (UnimplementedWebhookIntegrationServiceServer) CountWebhookIntegrations(context.Context, *CountWebhookIntegrationsRequest) (*CountWebhookIntegrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountWebhookIntegrations not implemented")
}
func (UnimplementedWebhookIntegrationServiceServer) mustEmbedUnimplementedWebhookIntegrationServiceServer() {
}

// UnsafeWebhookIntegrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebhookIntegrationServiceServer will
// result in compilation errors.
type UnsafeWebhookIntegrationServiceServer interface {
	mustEmbedUnimplementedWebhookIntegrationServiceServer()
}

func RegisterWebhookIntegrationServiceServer(s grpc.ServiceRegistrar, srv WebhookIntegrationServiceServer) {
	s.RegisterService(&WebhookIntegrationService_ServiceDesc, srv)
}

func _WebhookIntegrationService_CreateWebhookIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWebhookIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookIntegrationServiceServer).CreateWebhookIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebhookIntegrationService_CreateWebhookIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookIntegrationServiceServer).CreateWebhookIntegration(ctx, req.(*CreateWebhookIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebhookIntegrationService_ListWebhookIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWebhookIntegrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookIntegrationServiceServer).ListWebhookIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebhookIntegrationService_ListWebhookIntegrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookIntegrationServiceServer).ListWebhookIntegrations(ctx, req.(*ListWebhookIntegrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebhookIntegrationService_DeleteWebhookIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWebhookIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookIntegrationServiceServer).DeleteWebhookIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebhookIntegrationService_DeleteWebhookIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookIntegrationServiceServer).DeleteWebhookIntegration(ctx, req.(*DeleteWebhookIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebhookIntegrationService_ToggleWebhookIntegrationActivation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleWebhookIntegrationActivationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookIntegrationServiceServer).ToggleWebhookIntegrationActivation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebhookIntegrationService_ToggleWebhookIntegrationActivation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookIntegrationServiceServer).ToggleWebhookIntegrationActivation(ctx, req.(*ToggleWebhookIntegrationActivationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebhookIntegrationService_CountWebhookIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountWebhookIntegrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookIntegrationServiceServer).CountWebhookIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebhookIntegrationService_CountWebhookIntegrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookIntegrationServiceServer).CountWebhookIntegrations(ctx, req.(*CountWebhookIntegrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WebhookIntegrationService_ServiceDesc is the grpc.ServiceDesc for WebhookIntegrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebhookIntegrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.integrations.v1.WebhookIntegrationService",
	HandlerType: (*WebhookIntegrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWebhookIntegration",
			Handler:    _WebhookIntegrationService_CreateWebhookIntegration_Handler,
		},
		{
			MethodName: "ListWebhookIntegrations",
			Handler:    _WebhookIntegrationService_ListWebhookIntegrations_Handler,
		},
		{
			MethodName: "DeleteWebhookIntegration",
			Handler:    _WebhookIntegrationService_DeleteWebhookIntegration_Handler,
		},
		{
			MethodName: "ToggleWebhookIntegrationActivation",
			Handler:    _WebhookIntegrationService_ToggleWebhookIntegrationActivation_Handler,
		},
		{
			MethodName: "CountWebhookIntegrations",
			Handler:    _WebhookIntegrationService_CountWebhookIntegrations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/integrations/v1/webhook_integration_service.proto",
}
