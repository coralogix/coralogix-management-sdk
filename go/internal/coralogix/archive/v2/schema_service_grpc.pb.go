// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: com/coralogix/archive/dataset/v2/schema_service.proto

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
	SchemaService_SetNamedSchema_FullMethodName = "/com.coralogix.archive.dataset.v2.SchemaService/SetNamedSchema"
	SchemaService_GetNamedSchema_FullMethodName = "/com.coralogix.archive.dataset.v2.SchemaService/GetNamedSchema"
)

// SchemaServiceClient is the client API for SchemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchemaServiceClient interface {
	SetNamedSchema(ctx context.Context, in *SetNamedSchemaRequest, opts ...grpc.CallOption) (*SetNamedSchemaResponse, error)
	GetNamedSchema(ctx context.Context, in *GetNamedSchemaRequest, opts ...grpc.CallOption) (*GetNamedSchemaResponse, error)
}

type schemaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchemaServiceClient(cc grpc.ClientConnInterface) SchemaServiceClient {
	return &schemaServiceClient{cc}
}

func (c *schemaServiceClient) SetNamedSchema(ctx context.Context, in *SetNamedSchemaRequest, opts ...grpc.CallOption) (*SetNamedSchemaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetNamedSchemaResponse)
	err := c.cc.Invoke(ctx, SchemaService_SetNamedSchema_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) GetNamedSchema(ctx context.Context, in *GetNamedSchemaRequest, opts ...grpc.CallOption) (*GetNamedSchemaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNamedSchemaResponse)
	err := c.cc.Invoke(ctx, SchemaService_GetNamedSchema_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaServiceServer is the server API for SchemaService service.
// All implementations must embed UnimplementedSchemaServiceServer
// for forward compatibility.
type SchemaServiceServer interface {
	SetNamedSchema(context.Context, *SetNamedSchemaRequest) (*SetNamedSchemaResponse, error)
	GetNamedSchema(context.Context, *GetNamedSchemaRequest) (*GetNamedSchemaResponse, error)
	mustEmbedUnimplementedSchemaServiceServer()
}

// UnimplementedSchemaServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSchemaServiceServer struct{}

func (UnimplementedSchemaServiceServer) SetNamedSchema(context.Context, *SetNamedSchemaRequest) (*SetNamedSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNamedSchema not implemented")
}
func (UnimplementedSchemaServiceServer) GetNamedSchema(context.Context, *GetNamedSchemaRequest) (*GetNamedSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNamedSchema not implemented")
}
func (UnimplementedSchemaServiceServer) mustEmbedUnimplementedSchemaServiceServer() {}
func (UnimplementedSchemaServiceServer) testEmbeddedByValue()                       {}

// UnsafeSchemaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchemaServiceServer will
// result in compilation errors.
type UnsafeSchemaServiceServer interface {
	mustEmbedUnimplementedSchemaServiceServer()
}

func RegisterSchemaServiceServer(s grpc.ServiceRegistrar, srv SchemaServiceServer) {
	// If the following call pancis, it indicates UnimplementedSchemaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SchemaService_ServiceDesc, srv)
}

func _SchemaService_SetNamedSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNamedSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).SetNamedSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_SetNamedSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).SetNamedSchema(ctx, req.(*SetNamedSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_GetNamedSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamedSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).GetNamedSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_GetNamedSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).GetNamedSchema(ctx, req.(*GetNamedSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SchemaService_ServiceDesc is the grpc.ServiceDesc for SchemaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchemaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.archive.dataset.v2.SchemaService",
	HandlerType: (*SchemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetNamedSchema",
			Handler:    _SchemaService_SetNamedSchema_Handler,
		},
		{
			MethodName: "GetNamedSchema",
			Handler:    _SchemaService_GetNamedSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/archive/dataset/v2/schema_service.proto",
}
