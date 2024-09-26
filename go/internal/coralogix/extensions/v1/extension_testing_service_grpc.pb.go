// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.1
// source: com/coralogix/extensions/v1/extension_testing_service.proto

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
	ExtensionTestingService_InitializeTestingRevision_FullMethodName = "/com.coralogix.extensions.v1.ExtensionTestingService/InitializeTestingRevision"
	ExtensionTestingService_CleanupTestingRevision_FullMethodName    = "/com.coralogix.extensions.v1.ExtensionTestingService/CleanupTestingRevision"
	ExtensionTestingService_TestExtensionRevision_FullMethodName     = "/com.coralogix.extensions.v1.ExtensionTestingService/TestExtensionRevision"
)

// ExtensionTestingServiceClient is the client API for ExtensionTestingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExtensionTestingServiceClient interface {
	InitializeTestingRevision(ctx context.Context, in *InitializeTestingRevisionRequest, opts ...grpc.CallOption) (*InitializeTestingRevisionResponse, error)
	CleanupTestingRevision(ctx context.Context, in *CleanupTestingRevisionRequest, opts ...grpc.CallOption) (*CleanupTestingRevisionResponse, error)
	TestExtensionRevision(ctx context.Context, in *TestExtensionRevisionRequest, opts ...grpc.CallOption) (*TestExtensionRevisionResponse, error)
}

type extensionTestingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExtensionTestingServiceClient(cc grpc.ClientConnInterface) ExtensionTestingServiceClient {
	return &extensionTestingServiceClient{cc}
}

func (c *extensionTestingServiceClient) InitializeTestingRevision(ctx context.Context, in *InitializeTestingRevisionRequest, opts ...grpc.CallOption) (*InitializeTestingRevisionResponse, error) {
	out := new(InitializeTestingRevisionResponse)
	err := c.cc.Invoke(ctx, ExtensionTestingService_InitializeTestingRevision_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionTestingServiceClient) CleanupTestingRevision(ctx context.Context, in *CleanupTestingRevisionRequest, opts ...grpc.CallOption) (*CleanupTestingRevisionResponse, error) {
	out := new(CleanupTestingRevisionResponse)
	err := c.cc.Invoke(ctx, ExtensionTestingService_CleanupTestingRevision_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionTestingServiceClient) TestExtensionRevision(ctx context.Context, in *TestExtensionRevisionRequest, opts ...grpc.CallOption) (*TestExtensionRevisionResponse, error) {
	out := new(TestExtensionRevisionResponse)
	err := c.cc.Invoke(ctx, ExtensionTestingService_TestExtensionRevision_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionTestingServiceServer is the server API for ExtensionTestingService service.
// All implementations must embed UnimplementedExtensionTestingServiceServer
// for forward compatibility
type ExtensionTestingServiceServer interface {
	InitializeTestingRevision(context.Context, *InitializeTestingRevisionRequest) (*InitializeTestingRevisionResponse, error)
	CleanupTestingRevision(context.Context, *CleanupTestingRevisionRequest) (*CleanupTestingRevisionResponse, error)
	TestExtensionRevision(context.Context, *TestExtensionRevisionRequest) (*TestExtensionRevisionResponse, error)
	mustEmbedUnimplementedExtensionTestingServiceServer()
}

// UnimplementedExtensionTestingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExtensionTestingServiceServer struct {
}

func (UnimplementedExtensionTestingServiceServer) InitializeTestingRevision(context.Context, *InitializeTestingRevisionRequest) (*InitializeTestingRevisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitializeTestingRevision not implemented")
}
func (UnimplementedExtensionTestingServiceServer) CleanupTestingRevision(context.Context, *CleanupTestingRevisionRequest) (*CleanupTestingRevisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanupTestingRevision not implemented")
}
func (UnimplementedExtensionTestingServiceServer) TestExtensionRevision(context.Context, *TestExtensionRevisionRequest) (*TestExtensionRevisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestExtensionRevision not implemented")
}
func (UnimplementedExtensionTestingServiceServer) mustEmbedUnimplementedExtensionTestingServiceServer() {
}

// UnsafeExtensionTestingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExtensionTestingServiceServer will
// result in compilation errors.
type UnsafeExtensionTestingServiceServer interface {
	mustEmbedUnimplementedExtensionTestingServiceServer()
}

func RegisterExtensionTestingServiceServer(s grpc.ServiceRegistrar, srv ExtensionTestingServiceServer) {
	s.RegisterService(&ExtensionTestingService_ServiceDesc, srv)
}

func _ExtensionTestingService_InitializeTestingRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeTestingRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionTestingServiceServer).InitializeTestingRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExtensionTestingService_InitializeTestingRevision_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionTestingServiceServer).InitializeTestingRevision(ctx, req.(*InitializeTestingRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExtensionTestingService_CleanupTestingRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanupTestingRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionTestingServiceServer).CleanupTestingRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExtensionTestingService_CleanupTestingRevision_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionTestingServiceServer).CleanupTestingRevision(ctx, req.(*CleanupTestingRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExtensionTestingService_TestExtensionRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestExtensionRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionTestingServiceServer).TestExtensionRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExtensionTestingService_TestExtensionRevision_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionTestingServiceServer).TestExtensionRevision(ctx, req.(*TestExtensionRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExtensionTestingService_ServiceDesc is the grpc.ServiceDesc for ExtensionTestingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExtensionTestingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.extensions.v1.ExtensionTestingService",
	HandlerType: (*ExtensionTestingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitializeTestingRevision",
			Handler:    _ExtensionTestingService_InitializeTestingRevision_Handler,
		},
		{
			MethodName: "CleanupTestingRevision",
			Handler:    _ExtensionTestingService_CleanupTestingRevision_Handler,
		},
		{
			MethodName: "TestExtensionRevision",
			Handler:    _ExtensionTestingService_TestExtensionRevision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/extensions/v1/extension_testing_service.proto",
}
