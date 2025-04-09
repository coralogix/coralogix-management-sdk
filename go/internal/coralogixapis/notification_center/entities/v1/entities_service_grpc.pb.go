// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: com/coralogixapis/notification_center/entities/v1/entities_service.proto

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
	EntitiesService_ListEntityTypes_FullMethodName    = "/com.coralogixapis.notification_center.entities.v1.EntitiesService/ListEntityTypes"
	EntitiesService_ListEntitySubTypes_FullMethodName = "/com.coralogixapis.notification_center.entities.v1.EntitiesService/ListEntitySubTypes"
)

// EntitiesServiceClient is the client API for EntitiesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Public API to query information on registered entities in the Notification Center
type EntitiesServiceClient interface {
	ListEntityTypes(ctx context.Context, in *ListEntityTypesRequest, opts ...grpc.CallOption) (*ListEntityTypesResponse, error)
	ListEntitySubTypes(ctx context.Context, in *ListEntitySubTypesRequest, opts ...grpc.CallOption) (*ListEntitySubTypesResponse, error)
}

type entitiesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEntitiesServiceClient(cc grpc.ClientConnInterface) EntitiesServiceClient {
	return &entitiesServiceClient{cc}
}

func (c *entitiesServiceClient) ListEntityTypes(ctx context.Context, in *ListEntityTypesRequest, opts ...grpc.CallOption) (*ListEntityTypesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListEntityTypesResponse)
	err := c.cc.Invoke(ctx, EntitiesService_ListEntityTypes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entitiesServiceClient) ListEntitySubTypes(ctx context.Context, in *ListEntitySubTypesRequest, opts ...grpc.CallOption) (*ListEntitySubTypesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListEntitySubTypesResponse)
	err := c.cc.Invoke(ctx, EntitiesService_ListEntitySubTypes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntitiesServiceServer is the server API for EntitiesService service.
// All implementations must embed UnimplementedEntitiesServiceServer
// for forward compatibility.
//
// Public API to query information on registered entities in the Notification Center
type EntitiesServiceServer interface {
	ListEntityTypes(context.Context, *ListEntityTypesRequest) (*ListEntityTypesResponse, error)
	ListEntitySubTypes(context.Context, *ListEntitySubTypesRequest) (*ListEntitySubTypesResponse, error)
	mustEmbedUnimplementedEntitiesServiceServer()
}

// UnimplementedEntitiesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEntitiesServiceServer struct{}

func (UnimplementedEntitiesServiceServer) ListEntityTypes(context.Context, *ListEntityTypesRequest) (*ListEntityTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntityTypes not implemented")
}
func (UnimplementedEntitiesServiceServer) ListEntitySubTypes(context.Context, *ListEntitySubTypesRequest) (*ListEntitySubTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntitySubTypes not implemented")
}
func (UnimplementedEntitiesServiceServer) mustEmbedUnimplementedEntitiesServiceServer() {}
func (UnimplementedEntitiesServiceServer) testEmbeddedByValue()                         {}

// UnsafeEntitiesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntitiesServiceServer will
// result in compilation errors.
type UnsafeEntitiesServiceServer interface {
	mustEmbedUnimplementedEntitiesServiceServer()
}

func RegisterEntitiesServiceServer(s grpc.ServiceRegistrar, srv EntitiesServiceServer) {
	// If the following call pancis, it indicates UnimplementedEntitiesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EntitiesService_ServiceDesc, srv)
}

func _EntitiesService_ListEntityTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntityTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntitiesServiceServer).ListEntityTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntitiesService_ListEntityTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntitiesServiceServer).ListEntityTypes(ctx, req.(*ListEntityTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntitiesService_ListEntitySubTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntitySubTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntitiesServiceServer).ListEntitySubTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EntitiesService_ListEntitySubTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntitiesServiceServer).ListEntitySubTypes(ctx, req.(*ListEntitySubTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EntitiesService_ServiceDesc is the grpc.ServiceDesc for EntitiesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntitiesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.notification_center.entities.v1.EntitiesService",
	HandlerType: (*EntitiesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEntityTypes",
			Handler:    _EntitiesService_ListEntityTypes_Handler,
		},
		{
			MethodName: "ListEntitySubTypes",
			Handler:    _EntitiesService_ListEntitySubTypes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/notification_center/entities/v1/entities_service.proto",
}
