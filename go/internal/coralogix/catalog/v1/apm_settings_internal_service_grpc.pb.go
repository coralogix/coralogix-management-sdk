// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: com/coralogixapis/service_catalog/v1/apm_settings_internal_service.proto

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
	ApmSettingsInternalService_ValidateApmSource_FullMethodName = "/com.coralogixapis.service_catalog.v1.ApmSettingsInternalService/ValidateApmSource"
)

// ApmSettingsInternalServiceClient is the client API for ApmSettingsInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApmSettingsInternalServiceClient interface {
	ValidateApmSource(ctx context.Context, in *ValidateApmSourceRequest, opts ...grpc.CallOption) (*ValidateApmSourceResponse, error)
}

type apmSettingsInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApmSettingsInternalServiceClient(cc grpc.ClientConnInterface) ApmSettingsInternalServiceClient {
	return &apmSettingsInternalServiceClient{cc}
}

func (c *apmSettingsInternalServiceClient) ValidateApmSource(ctx context.Context, in *ValidateApmSourceRequest, opts ...grpc.CallOption) (*ValidateApmSourceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateApmSourceResponse)
	err := c.cc.Invoke(ctx, ApmSettingsInternalService_ValidateApmSource_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApmSettingsInternalServiceServer is the server API for ApmSettingsInternalService service.
// All implementations must embed UnimplementedApmSettingsInternalServiceServer
// for forward compatibility.
type ApmSettingsInternalServiceServer interface {
	ValidateApmSource(context.Context, *ValidateApmSourceRequest) (*ValidateApmSourceResponse, error)
	mustEmbedUnimplementedApmSettingsInternalServiceServer()
}

// UnimplementedApmSettingsInternalServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedApmSettingsInternalServiceServer struct{}

func (UnimplementedApmSettingsInternalServiceServer) ValidateApmSource(context.Context, *ValidateApmSourceRequest) (*ValidateApmSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateApmSource not implemented")
}
func (UnimplementedApmSettingsInternalServiceServer) mustEmbedUnimplementedApmSettingsInternalServiceServer() {
}
func (UnimplementedApmSettingsInternalServiceServer) testEmbeddedByValue() {}

// UnsafeApmSettingsInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApmSettingsInternalServiceServer will
// result in compilation errors.
type UnsafeApmSettingsInternalServiceServer interface {
	mustEmbedUnimplementedApmSettingsInternalServiceServer()
}

func RegisterApmSettingsInternalServiceServer(s grpc.ServiceRegistrar, srv ApmSettingsInternalServiceServer) {
	// If the following call pancis, it indicates UnimplementedApmSettingsInternalServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ApmSettingsInternalService_ServiceDesc, srv)
}

func _ApmSettingsInternalService_ValidateApmSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateApmSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApmSettingsInternalServiceServer).ValidateApmSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApmSettingsInternalService_ValidateApmSource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApmSettingsInternalServiceServer).ValidateApmSource(ctx, req.(*ValidateApmSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApmSettingsInternalService_ServiceDesc is the grpc.ServiceDesc for ApmSettingsInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApmSettingsInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.service_catalog.v1.ApmSettingsInternalService",
	HandlerType: (*ApmSettingsInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateApmSource",
			Handler:    _ApmSettingsInternalService_ValidateApmSource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/service_catalog/v1/apm_settings_internal_service.proto",
}