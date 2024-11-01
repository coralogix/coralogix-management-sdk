// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.28.3
// source: com/coralogix/extensions/v1/internal_onboarding_service.proto

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
	InternalOnboardingService_InternalDeployExtension_FullMethodName = "/com.coralogix.extensions.v1.InternalOnboardingService/InternalDeployExtension"
)

// InternalOnboardingServiceClient is the client API for InternalOnboardingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalOnboardingServiceClient interface {
	InternalDeployExtension(ctx context.Context, in *InternalDeployExtensionRequest, opts ...grpc.CallOption) (*InternalDeployExtensionResponse, error)
}

type internalOnboardingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalOnboardingServiceClient(cc grpc.ClientConnInterface) InternalOnboardingServiceClient {
	return &internalOnboardingServiceClient{cc}
}

func (c *internalOnboardingServiceClient) InternalDeployExtension(ctx context.Context, in *InternalDeployExtensionRequest, opts ...grpc.CallOption) (*InternalDeployExtensionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InternalDeployExtensionResponse)
	err := c.cc.Invoke(ctx, InternalOnboardingService_InternalDeployExtension_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalOnboardingServiceServer is the server API for InternalOnboardingService service.
// All implementations must embed UnimplementedInternalOnboardingServiceServer
// for forward compatibility
type InternalOnboardingServiceServer interface {
	InternalDeployExtension(context.Context, *InternalDeployExtensionRequest) (*InternalDeployExtensionResponse, error)
	mustEmbedUnimplementedInternalOnboardingServiceServer()
}

// UnimplementedInternalOnboardingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInternalOnboardingServiceServer struct {
}

func (UnimplementedInternalOnboardingServiceServer) InternalDeployExtension(context.Context, *InternalDeployExtensionRequest) (*InternalDeployExtensionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InternalDeployExtension not implemented")
}
func (UnimplementedInternalOnboardingServiceServer) mustEmbedUnimplementedInternalOnboardingServiceServer() {
}

// UnsafeInternalOnboardingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternalOnboardingServiceServer will
// result in compilation errors.
type UnsafeInternalOnboardingServiceServer interface {
	mustEmbedUnimplementedInternalOnboardingServiceServer()
}

func RegisterInternalOnboardingServiceServer(s grpc.ServiceRegistrar, srv InternalOnboardingServiceServer) {
	s.RegisterService(&InternalOnboardingService_ServiceDesc, srv)
}

func _InternalOnboardingService_InternalDeployExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalDeployExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalOnboardingServiceServer).InternalDeployExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalOnboardingService_InternalDeployExtension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalOnboardingServiceServer).InternalDeployExtension(ctx, req.(*InternalDeployExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InternalOnboardingService_ServiceDesc is the grpc.ServiceDesc for InternalOnboardingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InternalOnboardingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.extensions.v1.InternalOnboardingService",
	HandlerType: (*InternalOnboardingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InternalDeployExtension",
			Handler:    _InternalOnboardingService_InternalDeployExtension_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/extensions/v1/internal_onboarding_service.proto",
}
