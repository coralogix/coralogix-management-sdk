// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.1
// source: com/coralogixapis/apm/queries/k8s/v1/infra_filters_service.proto

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
	InfraFiltersService_ListFilters_FullMethodName = "/com.coralogixapis.apm.queries.k8s.v1.InfraFiltersService/ListFilters"
)

// InfraFiltersServiceClient is the client API for InfraFiltersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InfraFiltersServiceClient interface {
	ListFilters(ctx context.Context, in *ListFiltersRequest, opts ...grpc.CallOption) (*ListFiltersResponse, error)
}

type infraFiltersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInfraFiltersServiceClient(cc grpc.ClientConnInterface) InfraFiltersServiceClient {
	return &infraFiltersServiceClient{cc}
}

func (c *infraFiltersServiceClient) ListFilters(ctx context.Context, in *ListFiltersRequest, opts ...grpc.CallOption) (*ListFiltersResponse, error) {
	out := new(ListFiltersResponse)
	err := c.cc.Invoke(ctx, InfraFiltersService_ListFilters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfraFiltersServiceServer is the server API for InfraFiltersService service.
// All implementations must embed UnimplementedInfraFiltersServiceServer
// for forward compatibility
type InfraFiltersServiceServer interface {
	ListFilters(context.Context, *ListFiltersRequest) (*ListFiltersResponse, error)
	mustEmbedUnimplementedInfraFiltersServiceServer()
}

// UnimplementedInfraFiltersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInfraFiltersServiceServer struct {
}

func (UnimplementedInfraFiltersServiceServer) ListFilters(context.Context, *ListFiltersRequest) (*ListFiltersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFilters not implemented")
}
func (UnimplementedInfraFiltersServiceServer) mustEmbedUnimplementedInfraFiltersServiceServer() {}

// UnsafeInfraFiltersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfraFiltersServiceServer will
// result in compilation errors.
type UnsafeInfraFiltersServiceServer interface {
	mustEmbedUnimplementedInfraFiltersServiceServer()
}

func RegisterInfraFiltersServiceServer(s grpc.ServiceRegistrar, srv InfraFiltersServiceServer) {
	s.RegisterService(&InfraFiltersService_ServiceDesc, srv)
}

func _InfraFiltersService_ListFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFiltersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraFiltersServiceServer).ListFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InfraFiltersService_ListFilters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraFiltersServiceServer).ListFilters(ctx, req.(*ListFiltersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InfraFiltersService_ServiceDesc is the grpc.ServiceDesc for InfraFiltersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InfraFiltersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.apm.queries.k8s.v1.InfraFiltersService",
	HandlerType: (*InfraFiltersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFilters",
			Handler:    _InfraFiltersService_ListFilters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/apm/queries/k8s/v1/infra_filters_service.proto",
}
