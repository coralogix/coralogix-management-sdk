// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: com/coralogixapis/apm/requirements/v2/feature_requirements_service.proto

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
	FeatureRequirementsService_ValidateFeatureRequirements_FullMethodName        = "/com.coralogixapis.apm.requirements.v2.FeatureRequirementsService/ValidateFeatureRequirements"
	FeatureRequirementsService_ListFeatureRequirementsDataSources_FullMethodName = "/com.coralogixapis.apm.requirements.v2.FeatureRequirementsService/ListFeatureRequirementsDataSources"
)

// FeatureRequirementsServiceClient is the client API for FeatureRequirementsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeatureRequirementsServiceClient interface {
	ValidateFeatureRequirements(ctx context.Context, in *ValidateFeatureRequirementsRequest, opts ...grpc.CallOption) (*ValidateFeatureRequirementsResponse, error)
	ListFeatureRequirementsDataSources(ctx context.Context, in *ListFeatureRequirementsDataSourcesRequest, opts ...grpc.CallOption) (*ListFeatureRequirementsDataSourcesResponse, error)
}

type featureRequirementsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeatureRequirementsServiceClient(cc grpc.ClientConnInterface) FeatureRequirementsServiceClient {
	return &featureRequirementsServiceClient{cc}
}

func (c *featureRequirementsServiceClient) ValidateFeatureRequirements(ctx context.Context, in *ValidateFeatureRequirementsRequest, opts ...grpc.CallOption) (*ValidateFeatureRequirementsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateFeatureRequirementsResponse)
	err := c.cc.Invoke(ctx, FeatureRequirementsService_ValidateFeatureRequirements_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureRequirementsServiceClient) ListFeatureRequirementsDataSources(ctx context.Context, in *ListFeatureRequirementsDataSourcesRequest, opts ...grpc.CallOption) (*ListFeatureRequirementsDataSourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFeatureRequirementsDataSourcesResponse)
	err := c.cc.Invoke(ctx, FeatureRequirementsService_ListFeatureRequirementsDataSources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeatureRequirementsServiceServer is the server API for FeatureRequirementsService service.
// All implementations must embed UnimplementedFeatureRequirementsServiceServer
// for forward compatibility.
type FeatureRequirementsServiceServer interface {
	ValidateFeatureRequirements(context.Context, *ValidateFeatureRequirementsRequest) (*ValidateFeatureRequirementsResponse, error)
	ListFeatureRequirementsDataSources(context.Context, *ListFeatureRequirementsDataSourcesRequest) (*ListFeatureRequirementsDataSourcesResponse, error)
	mustEmbedUnimplementedFeatureRequirementsServiceServer()
}

// UnimplementedFeatureRequirementsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFeatureRequirementsServiceServer struct{}

func (UnimplementedFeatureRequirementsServiceServer) ValidateFeatureRequirements(context.Context, *ValidateFeatureRequirementsRequest) (*ValidateFeatureRequirementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateFeatureRequirements not implemented")
}
func (UnimplementedFeatureRequirementsServiceServer) ListFeatureRequirementsDataSources(context.Context, *ListFeatureRequirementsDataSourcesRequest) (*ListFeatureRequirementsDataSourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeatureRequirementsDataSources not implemented")
}
func (UnimplementedFeatureRequirementsServiceServer) mustEmbedUnimplementedFeatureRequirementsServiceServer() {
}
func (UnimplementedFeatureRequirementsServiceServer) testEmbeddedByValue() {}

// UnsafeFeatureRequirementsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeatureRequirementsServiceServer will
// result in compilation errors.
type UnsafeFeatureRequirementsServiceServer interface {
	mustEmbedUnimplementedFeatureRequirementsServiceServer()
}

func RegisterFeatureRequirementsServiceServer(s grpc.ServiceRegistrar, srv FeatureRequirementsServiceServer) {
	// If the following call pancis, it indicates UnimplementedFeatureRequirementsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FeatureRequirementsService_ServiceDesc, srv)
}

func _FeatureRequirementsService_ValidateFeatureRequirements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateFeatureRequirementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureRequirementsServiceServer).ValidateFeatureRequirements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureRequirementsService_ValidateFeatureRequirements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureRequirementsServiceServer).ValidateFeatureRequirements(ctx, req.(*ValidateFeatureRequirementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureRequirementsService_ListFeatureRequirementsDataSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFeatureRequirementsDataSourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureRequirementsServiceServer).ListFeatureRequirementsDataSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureRequirementsService_ListFeatureRequirementsDataSources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureRequirementsServiceServer).ListFeatureRequirementsDataSources(ctx, req.(*ListFeatureRequirementsDataSourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeatureRequirementsService_ServiceDesc is the grpc.ServiceDesc for FeatureRequirementsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeatureRequirementsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.apm.requirements.v2.FeatureRequirementsService",
	HandlerType: (*FeatureRequirementsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateFeatureRequirements",
			Handler:    _FeatureRequirementsService_ValidateFeatureRequirements_Handler,
		},
		{
			MethodName: "ListFeatureRequirementsDataSources",
			Handler:    _FeatureRequirementsService_ListFeatureRequirementsDataSources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/apm/requirements/v2/feature_requirements_service.proto",
}
