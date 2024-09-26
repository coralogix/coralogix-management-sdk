// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.1
// source: com/coralogixapis/apm/requirements/v2/feature_requirements_internal_service.proto

package v2

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
	FeatureRequirementsInternalService_ValidateFeatureRequirementsInternal_FullMethodName        = "/com.coralogixapis.apm.requirements.v2.FeatureRequirementsInternalService/ValidateFeatureRequirementsInternal"
	FeatureRequirementsInternalService_ListFeatureRequirementsInternalDataSources_FullMethodName = "/com.coralogixapis.apm.requirements.v2.FeatureRequirementsInternalService/ListFeatureRequirementsInternalDataSources"
)

// FeatureRequirementsInternalServiceClient is the client API for FeatureRequirementsInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeatureRequirementsInternalServiceClient interface {
	ValidateFeatureRequirementsInternal(ctx context.Context, in *ValidateFeatureRequirementsInternalRequest, opts ...grpc.CallOption) (*ValidateFeatureRequirementsInternalResponse, error)
	ListFeatureRequirementsInternalDataSources(ctx context.Context, in *ListFeatureRequirementsInternalDataSourcesRequest, opts ...grpc.CallOption) (*ListFeatureRequirementsInternalDataSourcesResponse, error)
}

type featureRequirementsInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeatureRequirementsInternalServiceClient(cc grpc.ClientConnInterface) FeatureRequirementsInternalServiceClient {
	return &featureRequirementsInternalServiceClient{cc}
}

func (c *featureRequirementsInternalServiceClient) ValidateFeatureRequirementsInternal(ctx context.Context, in *ValidateFeatureRequirementsInternalRequest, opts ...grpc.CallOption) (*ValidateFeatureRequirementsInternalResponse, error) {
	out := new(ValidateFeatureRequirementsInternalResponse)
	err := c.cc.Invoke(ctx, FeatureRequirementsInternalService_ValidateFeatureRequirementsInternal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureRequirementsInternalServiceClient) ListFeatureRequirementsInternalDataSources(ctx context.Context, in *ListFeatureRequirementsInternalDataSourcesRequest, opts ...grpc.CallOption) (*ListFeatureRequirementsInternalDataSourcesResponse, error) {
	out := new(ListFeatureRequirementsInternalDataSourcesResponse)
	err := c.cc.Invoke(ctx, FeatureRequirementsInternalService_ListFeatureRequirementsInternalDataSources_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeatureRequirementsInternalServiceServer is the server API for FeatureRequirementsInternalService service.
// All implementations must embed UnimplementedFeatureRequirementsInternalServiceServer
// for forward compatibility
type FeatureRequirementsInternalServiceServer interface {
	ValidateFeatureRequirementsInternal(context.Context, *ValidateFeatureRequirementsInternalRequest) (*ValidateFeatureRequirementsInternalResponse, error)
	ListFeatureRequirementsInternalDataSources(context.Context, *ListFeatureRequirementsInternalDataSourcesRequest) (*ListFeatureRequirementsInternalDataSourcesResponse, error)
	mustEmbedUnimplementedFeatureRequirementsInternalServiceServer()
}

// UnimplementedFeatureRequirementsInternalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeatureRequirementsInternalServiceServer struct {
}

func (UnimplementedFeatureRequirementsInternalServiceServer) ValidateFeatureRequirementsInternal(context.Context, *ValidateFeatureRequirementsInternalRequest) (*ValidateFeatureRequirementsInternalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateFeatureRequirementsInternal not implemented")
}
func (UnimplementedFeatureRequirementsInternalServiceServer) ListFeatureRequirementsInternalDataSources(context.Context, *ListFeatureRequirementsInternalDataSourcesRequest) (*ListFeatureRequirementsInternalDataSourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeatureRequirementsInternalDataSources not implemented")
}
func (UnimplementedFeatureRequirementsInternalServiceServer) mustEmbedUnimplementedFeatureRequirementsInternalServiceServer() {
}

// UnsafeFeatureRequirementsInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeatureRequirementsInternalServiceServer will
// result in compilation errors.
type UnsafeFeatureRequirementsInternalServiceServer interface {
	mustEmbedUnimplementedFeatureRequirementsInternalServiceServer()
}

func RegisterFeatureRequirementsInternalServiceServer(s grpc.ServiceRegistrar, srv FeatureRequirementsInternalServiceServer) {
	s.RegisterService(&FeatureRequirementsInternalService_ServiceDesc, srv)
}

func _FeatureRequirementsInternalService_ValidateFeatureRequirementsInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateFeatureRequirementsInternalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureRequirementsInternalServiceServer).ValidateFeatureRequirementsInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureRequirementsInternalService_ValidateFeatureRequirementsInternal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureRequirementsInternalServiceServer).ValidateFeatureRequirementsInternal(ctx, req.(*ValidateFeatureRequirementsInternalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureRequirementsInternalService_ListFeatureRequirementsInternalDataSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFeatureRequirementsInternalDataSourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureRequirementsInternalServiceServer).ListFeatureRequirementsInternalDataSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureRequirementsInternalService_ListFeatureRequirementsInternalDataSources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureRequirementsInternalServiceServer).ListFeatureRequirementsInternalDataSources(ctx, req.(*ListFeatureRequirementsInternalDataSourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeatureRequirementsInternalService_ServiceDesc is the grpc.ServiceDesc for FeatureRequirementsInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeatureRequirementsInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.apm.requirements.v2.FeatureRequirementsInternalService",
	HandlerType: (*FeatureRequirementsInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateFeatureRequirementsInternal",
			Handler:    _FeatureRequirementsInternalService_ValidateFeatureRequirementsInternal_Handler,
		},
		{
			MethodName: "ListFeatureRequirementsInternalDataSources",
			Handler:    _FeatureRequirementsInternalService_ListFeatureRequirementsInternalDataSources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/apm/requirements/v2/feature_requirements_internal_service.proto",
}
