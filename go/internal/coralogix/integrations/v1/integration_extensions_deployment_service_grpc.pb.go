// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: com/coralogix/integrations/v1/integration_extensions_deployment_service.proto

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
	IntegrationExtensionsDeploymentService_DeployIntegrationExtensions_FullMethodName = "/com.coralogix.integrations.v1.IntegrationExtensionsDeploymentService/DeployIntegrationExtensions"
)

// IntegrationExtensionsDeploymentServiceClient is the client API for IntegrationExtensionsDeploymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntegrationExtensionsDeploymentServiceClient interface {
	DeployIntegrationExtensions(ctx context.Context, in *DeployIntegrationExtensionsRequest, opts ...grpc.CallOption) (*DeployIntegrationExtensionsResponse, error)
}

type integrationExtensionsDeploymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIntegrationExtensionsDeploymentServiceClient(cc grpc.ClientConnInterface) IntegrationExtensionsDeploymentServiceClient {
	return &integrationExtensionsDeploymentServiceClient{cc}
}

func (c *integrationExtensionsDeploymentServiceClient) DeployIntegrationExtensions(ctx context.Context, in *DeployIntegrationExtensionsRequest, opts ...grpc.CallOption) (*DeployIntegrationExtensionsResponse, error) {
	out := new(DeployIntegrationExtensionsResponse)
	err := c.cc.Invoke(ctx, IntegrationExtensionsDeploymentService_DeployIntegrationExtensions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntegrationExtensionsDeploymentServiceServer is the server API for IntegrationExtensionsDeploymentService service.
// All implementations must embed UnimplementedIntegrationExtensionsDeploymentServiceServer
// for forward compatibility
type IntegrationExtensionsDeploymentServiceServer interface {
	DeployIntegrationExtensions(context.Context, *DeployIntegrationExtensionsRequest) (*DeployIntegrationExtensionsResponse, error)
	mustEmbedUnimplementedIntegrationExtensionsDeploymentServiceServer()
}

// UnimplementedIntegrationExtensionsDeploymentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIntegrationExtensionsDeploymentServiceServer struct {
}

func (UnimplementedIntegrationExtensionsDeploymentServiceServer) DeployIntegrationExtensions(context.Context, *DeployIntegrationExtensionsRequest) (*DeployIntegrationExtensionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployIntegrationExtensions not implemented")
}
func (UnimplementedIntegrationExtensionsDeploymentServiceServer) mustEmbedUnimplementedIntegrationExtensionsDeploymentServiceServer() {
}

// UnsafeIntegrationExtensionsDeploymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntegrationExtensionsDeploymentServiceServer will
// result in compilation errors.
type UnsafeIntegrationExtensionsDeploymentServiceServer interface {
	mustEmbedUnimplementedIntegrationExtensionsDeploymentServiceServer()
}

func RegisterIntegrationExtensionsDeploymentServiceServer(s grpc.ServiceRegistrar, srv IntegrationExtensionsDeploymentServiceServer) {
	s.RegisterService(&IntegrationExtensionsDeploymentService_ServiceDesc, srv)
}

func _IntegrationExtensionsDeploymentService_DeployIntegrationExtensions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployIntegrationExtensionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationExtensionsDeploymentServiceServer).DeployIntegrationExtensions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationExtensionsDeploymentService_DeployIntegrationExtensions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationExtensionsDeploymentServiceServer).DeployIntegrationExtensions(ctx, req.(*DeployIntegrationExtensionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IntegrationExtensionsDeploymentService_ServiceDesc is the grpc.ServiceDesc for IntegrationExtensionsDeploymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IntegrationExtensionsDeploymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.integrations.v1.IntegrationExtensionsDeploymentService",
	HandlerType: (*IntegrationExtensionsDeploymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeployIntegrationExtensions",
			Handler:    _IntegrationExtensionsDeploymentService_DeployIntegrationExtensions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/integrations/v1/integration_extensions_deployment_service.proto",
}
