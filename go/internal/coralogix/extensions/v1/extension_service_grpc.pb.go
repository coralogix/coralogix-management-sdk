// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: com/coralogix/extensions/v1/extension_service.proto

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
	ExtensionService_GetAllExtensions_FullMethodName = "/com.coralogix.extensions.v1.ExtensionService/GetAllExtensions"
	ExtensionService_GetExtension_FullMethodName     = "/com.coralogix.extensions.v1.ExtensionService/GetExtension"
)

// ExtensionServiceClient is the client API for ExtensionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExtensionServiceClient interface {
	GetAllExtensions(ctx context.Context, in *GetAllExtensionsRequest, opts ...grpc.CallOption) (*GetAllExtensionsResponse, error)
	GetExtension(ctx context.Context, in *GetExtensionRequest, opts ...grpc.CallOption) (*GetExtensionResponse, error)
}

type extensionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExtensionServiceClient(cc grpc.ClientConnInterface) ExtensionServiceClient {
	return &extensionServiceClient{cc}
}

func (c *extensionServiceClient) GetAllExtensions(ctx context.Context, in *GetAllExtensionsRequest, opts ...grpc.CallOption) (*GetAllExtensionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllExtensionsResponse)
	err := c.cc.Invoke(ctx, ExtensionService_GetAllExtensions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionServiceClient) GetExtension(ctx context.Context, in *GetExtensionRequest, opts ...grpc.CallOption) (*GetExtensionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetExtensionResponse)
	err := c.cc.Invoke(ctx, ExtensionService_GetExtension_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionServiceServer is the server API for ExtensionService service.
// All implementations must embed UnimplementedExtensionServiceServer
// for forward compatibility
type ExtensionServiceServer interface {
	GetAllExtensions(context.Context, *GetAllExtensionsRequest) (*GetAllExtensionsResponse, error)
	GetExtension(context.Context, *GetExtensionRequest) (*GetExtensionResponse, error)
	mustEmbedUnimplementedExtensionServiceServer()
}

// UnimplementedExtensionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExtensionServiceServer struct {
}

func (UnimplementedExtensionServiceServer) GetAllExtensions(context.Context, *GetAllExtensionsRequest) (*GetAllExtensionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllExtensions not implemented")
}
func (UnimplementedExtensionServiceServer) GetExtension(context.Context, *GetExtensionRequest) (*GetExtensionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExtension not implemented")
}
func (UnimplementedExtensionServiceServer) mustEmbedUnimplementedExtensionServiceServer() {}

// UnsafeExtensionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExtensionServiceServer will
// result in compilation errors.
type UnsafeExtensionServiceServer interface {
	mustEmbedUnimplementedExtensionServiceServer()
}

func RegisterExtensionServiceServer(s grpc.ServiceRegistrar, srv ExtensionServiceServer) {
	s.RegisterService(&ExtensionService_ServiceDesc, srv)
}

func _ExtensionService_GetAllExtensions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllExtensionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionServiceServer).GetAllExtensions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExtensionService_GetAllExtensions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionServiceServer).GetAllExtensions(ctx, req.(*GetAllExtensionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExtensionService_GetExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionServiceServer).GetExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExtensionService_GetExtension_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionServiceServer).GetExtension(ctx, req.(*GetExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExtensionService_ServiceDesc is the grpc.ServiceDesc for ExtensionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExtensionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.extensions.v1.ExtensionService",
	HandlerType: (*ExtensionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllExtensions",
			Handler:    _ExtensionService_GetAllExtensions_Handler,
		},
		{
			MethodName: "GetExtension",
			Handler:    _ExtensionService_GetExtension_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/extensions/v1/extension_service.proto",
}
