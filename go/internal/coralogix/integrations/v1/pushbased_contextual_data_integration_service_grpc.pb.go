// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
// source: com/coralogix/integrations/v1/pushbased_contextual_data_integration_service.proto

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
	PushBasedContextualDataIntegrationService_CreatePushBasedContextualDataIntegration_FullMethodName = "/com.coralogix.integrations.v1.PushBasedContextualDataIntegrationService/CreatePushBasedContextualDataIntegration"
	PushBasedContextualDataIntegrationService_UpdatePushBasedContextualDataIntegration_FullMethodName = "/com.coralogix.integrations.v1.PushBasedContextualDataIntegrationService/UpdatePushBasedContextualDataIntegration"
	PushBasedContextualDataIntegrationService_ListPushBasedContextualDataIntegrations_FullMethodName  = "/com.coralogix.integrations.v1.PushBasedContextualDataIntegrationService/ListPushBasedContextualDataIntegrations"
	PushBasedContextualDataIntegrationService_DeletePushBasedContextualDataIntegration_FullMethodName = "/com.coralogix.integrations.v1.PushBasedContextualDataIntegrationService/DeletePushBasedContextualDataIntegration"
	PushBasedContextualDataIntegrationService_CountPushBasedContextualDataIntegrations_FullMethodName = "/com.coralogix.integrations.v1.PushBasedContextualDataIntegrationService/CountPushBasedContextualDataIntegrations"
)

// PushBasedContextualDataIntegrationServiceClient is the client API for PushBasedContextualDataIntegrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushBasedContextualDataIntegrationServiceClient interface {
	CreatePushBasedContextualDataIntegration(ctx context.Context, in *CreatePushBasedIntegrationRequest, opts ...grpc.CallOption) (*CreatePushBasedIntegrationResponse, error)
	UpdatePushBasedContextualDataIntegration(ctx context.Context, in *UpdatePushBasedContextualDataIntegrationsRequest, opts ...grpc.CallOption) (*UpdatePushBasedContextualDataIntegrationsResponse, error)
	ListPushBasedContextualDataIntegrations(ctx context.Context, in *ListPushBasedContextualDataIntegrationsRequest, opts ...grpc.CallOption) (*ListPushBasedContextualDataIntegrationsResponse, error)
	DeletePushBasedContextualDataIntegration(ctx context.Context, in *DeletePushBasedContextualDataIntegrationRequest, opts ...grpc.CallOption) (*DeletePushBasedContextualDataIntegrationResponse, error)
	CountPushBasedContextualDataIntegrations(ctx context.Context, in *CountPushBasedContextualDataIntegrationsRequest, opts ...grpc.CallOption) (*CountPushBasedContextualDataIntegrationsResponse, error)
}

type pushBasedContextualDataIntegrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPushBasedContextualDataIntegrationServiceClient(cc grpc.ClientConnInterface) PushBasedContextualDataIntegrationServiceClient {
	return &pushBasedContextualDataIntegrationServiceClient{cc}
}

func (c *pushBasedContextualDataIntegrationServiceClient) CreatePushBasedContextualDataIntegration(ctx context.Context, in *CreatePushBasedIntegrationRequest, opts ...grpc.CallOption) (*CreatePushBasedIntegrationResponse, error) {
	out := new(CreatePushBasedIntegrationResponse)
	err := c.cc.Invoke(ctx, PushBasedContextualDataIntegrationService_CreatePushBasedContextualDataIntegration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushBasedContextualDataIntegrationServiceClient) UpdatePushBasedContextualDataIntegration(ctx context.Context, in *UpdatePushBasedContextualDataIntegrationsRequest, opts ...grpc.CallOption) (*UpdatePushBasedContextualDataIntegrationsResponse, error) {
	out := new(UpdatePushBasedContextualDataIntegrationsResponse)
	err := c.cc.Invoke(ctx, PushBasedContextualDataIntegrationService_UpdatePushBasedContextualDataIntegration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushBasedContextualDataIntegrationServiceClient) ListPushBasedContextualDataIntegrations(ctx context.Context, in *ListPushBasedContextualDataIntegrationsRequest, opts ...grpc.CallOption) (*ListPushBasedContextualDataIntegrationsResponse, error) {
	out := new(ListPushBasedContextualDataIntegrationsResponse)
	err := c.cc.Invoke(ctx, PushBasedContextualDataIntegrationService_ListPushBasedContextualDataIntegrations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushBasedContextualDataIntegrationServiceClient) DeletePushBasedContextualDataIntegration(ctx context.Context, in *DeletePushBasedContextualDataIntegrationRequest, opts ...grpc.CallOption) (*DeletePushBasedContextualDataIntegrationResponse, error) {
	out := new(DeletePushBasedContextualDataIntegrationResponse)
	err := c.cc.Invoke(ctx, PushBasedContextualDataIntegrationService_DeletePushBasedContextualDataIntegration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushBasedContextualDataIntegrationServiceClient) CountPushBasedContextualDataIntegrations(ctx context.Context, in *CountPushBasedContextualDataIntegrationsRequest, opts ...grpc.CallOption) (*CountPushBasedContextualDataIntegrationsResponse, error) {
	out := new(CountPushBasedContextualDataIntegrationsResponse)
	err := c.cc.Invoke(ctx, PushBasedContextualDataIntegrationService_CountPushBasedContextualDataIntegrations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushBasedContextualDataIntegrationServiceServer is the server API for PushBasedContextualDataIntegrationService service.
// All implementations must embed UnimplementedPushBasedContextualDataIntegrationServiceServer
// for forward compatibility
type PushBasedContextualDataIntegrationServiceServer interface {
	CreatePushBasedContextualDataIntegration(context.Context, *CreatePushBasedIntegrationRequest) (*CreatePushBasedIntegrationResponse, error)
	UpdatePushBasedContextualDataIntegration(context.Context, *UpdatePushBasedContextualDataIntegrationsRequest) (*UpdatePushBasedContextualDataIntegrationsResponse, error)
	ListPushBasedContextualDataIntegrations(context.Context, *ListPushBasedContextualDataIntegrationsRequest) (*ListPushBasedContextualDataIntegrationsResponse, error)
	DeletePushBasedContextualDataIntegration(context.Context, *DeletePushBasedContextualDataIntegrationRequest) (*DeletePushBasedContextualDataIntegrationResponse, error)
	CountPushBasedContextualDataIntegrations(context.Context, *CountPushBasedContextualDataIntegrationsRequest) (*CountPushBasedContextualDataIntegrationsResponse, error)
	mustEmbedUnimplementedPushBasedContextualDataIntegrationServiceServer()
}

// UnimplementedPushBasedContextualDataIntegrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPushBasedContextualDataIntegrationServiceServer struct {
}

func (UnimplementedPushBasedContextualDataIntegrationServiceServer) CreatePushBasedContextualDataIntegration(context.Context, *CreatePushBasedIntegrationRequest) (*CreatePushBasedIntegrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePushBasedContextualDataIntegration not implemented")
}
func (UnimplementedPushBasedContextualDataIntegrationServiceServer) UpdatePushBasedContextualDataIntegration(context.Context, *UpdatePushBasedContextualDataIntegrationsRequest) (*UpdatePushBasedContextualDataIntegrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePushBasedContextualDataIntegration not implemented")
}
func (UnimplementedPushBasedContextualDataIntegrationServiceServer) ListPushBasedContextualDataIntegrations(context.Context, *ListPushBasedContextualDataIntegrationsRequest) (*ListPushBasedContextualDataIntegrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPushBasedContextualDataIntegrations not implemented")
}
func (UnimplementedPushBasedContextualDataIntegrationServiceServer) DeletePushBasedContextualDataIntegration(context.Context, *DeletePushBasedContextualDataIntegrationRequest) (*DeletePushBasedContextualDataIntegrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePushBasedContextualDataIntegration not implemented")
}
func (UnimplementedPushBasedContextualDataIntegrationServiceServer) CountPushBasedContextualDataIntegrations(context.Context, *CountPushBasedContextualDataIntegrationsRequest) (*CountPushBasedContextualDataIntegrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountPushBasedContextualDataIntegrations not implemented")
}
func (UnimplementedPushBasedContextualDataIntegrationServiceServer) mustEmbedUnimplementedPushBasedContextualDataIntegrationServiceServer() {
}

// UnsafePushBasedContextualDataIntegrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushBasedContextualDataIntegrationServiceServer will
// result in compilation errors.
type UnsafePushBasedContextualDataIntegrationServiceServer interface {
	mustEmbedUnimplementedPushBasedContextualDataIntegrationServiceServer()
}

func RegisterPushBasedContextualDataIntegrationServiceServer(s grpc.ServiceRegistrar, srv PushBasedContextualDataIntegrationServiceServer) {
	s.RegisterService(&PushBasedContextualDataIntegrationService_ServiceDesc, srv)
}

func _PushBasedContextualDataIntegrationService_CreatePushBasedContextualDataIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePushBasedIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushBasedContextualDataIntegrationServiceServer).CreatePushBasedContextualDataIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PushBasedContextualDataIntegrationService_CreatePushBasedContextualDataIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushBasedContextualDataIntegrationServiceServer).CreatePushBasedContextualDataIntegration(ctx, req.(*CreatePushBasedIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushBasedContextualDataIntegrationService_UpdatePushBasedContextualDataIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePushBasedContextualDataIntegrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushBasedContextualDataIntegrationServiceServer).UpdatePushBasedContextualDataIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PushBasedContextualDataIntegrationService_UpdatePushBasedContextualDataIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushBasedContextualDataIntegrationServiceServer).UpdatePushBasedContextualDataIntegration(ctx, req.(*UpdatePushBasedContextualDataIntegrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushBasedContextualDataIntegrationService_ListPushBasedContextualDataIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPushBasedContextualDataIntegrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushBasedContextualDataIntegrationServiceServer).ListPushBasedContextualDataIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PushBasedContextualDataIntegrationService_ListPushBasedContextualDataIntegrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushBasedContextualDataIntegrationServiceServer).ListPushBasedContextualDataIntegrations(ctx, req.(*ListPushBasedContextualDataIntegrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushBasedContextualDataIntegrationService_DeletePushBasedContextualDataIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePushBasedContextualDataIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushBasedContextualDataIntegrationServiceServer).DeletePushBasedContextualDataIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PushBasedContextualDataIntegrationService_DeletePushBasedContextualDataIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushBasedContextualDataIntegrationServiceServer).DeletePushBasedContextualDataIntegration(ctx, req.(*DeletePushBasedContextualDataIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushBasedContextualDataIntegrationService_CountPushBasedContextualDataIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountPushBasedContextualDataIntegrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushBasedContextualDataIntegrationServiceServer).CountPushBasedContextualDataIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PushBasedContextualDataIntegrationService_CountPushBasedContextualDataIntegrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushBasedContextualDataIntegrationServiceServer).CountPushBasedContextualDataIntegrations(ctx, req.(*CountPushBasedContextualDataIntegrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PushBasedContextualDataIntegrationService_ServiceDesc is the grpc.ServiceDesc for PushBasedContextualDataIntegrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PushBasedContextualDataIntegrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.integrations.v1.PushBasedContextualDataIntegrationService",
	HandlerType: (*PushBasedContextualDataIntegrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePushBasedContextualDataIntegration",
			Handler:    _PushBasedContextualDataIntegrationService_CreatePushBasedContextualDataIntegration_Handler,
		},
		{
			MethodName: "UpdatePushBasedContextualDataIntegration",
			Handler:    _PushBasedContextualDataIntegrationService_UpdatePushBasedContextualDataIntegration_Handler,
		},
		{
			MethodName: "ListPushBasedContextualDataIntegrations",
			Handler:    _PushBasedContextualDataIntegrationService_ListPushBasedContextualDataIntegrations_Handler,
		},
		{
			MethodName: "DeletePushBasedContextualDataIntegration",
			Handler:    _PushBasedContextualDataIntegrationService_DeletePushBasedContextualDataIntegration_Handler,
		},
		{
			MethodName: "CountPushBasedContextualDataIntegrations",
			Handler:    _PushBasedContextualDataIntegrationService_CountPushBasedContextualDataIntegrations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/integrations/v1/pushbased_contextual_data_integration_service.proto",
}
