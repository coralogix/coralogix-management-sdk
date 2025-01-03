// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
// source: com/coralogix/extensions/v1/extension_content_management_service.proto

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
	ExtensionContentManagementService_ImportAndReleaseExtensions_FullMethodName = "/com.coralogix.extensions.v1.ExtensionContentManagementService/ImportAndReleaseExtensions"
	ExtensionContentManagementService_ValidateExtensionItems_FullMethodName     = "/com.coralogix.extensions.v1.ExtensionContentManagementService/ValidateExtensionItems"
)

// ExtensionContentManagementServiceClient is the client API for ExtensionContentManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExtensionContentManagementServiceClient interface {
	ImportAndReleaseExtensions(ctx context.Context, opts ...grpc.CallOption) (ExtensionContentManagementService_ImportAndReleaseExtensionsClient, error)
	ValidateExtensionItems(ctx context.Context, opts ...grpc.CallOption) (ExtensionContentManagementService_ValidateExtensionItemsClient, error)
}

type extensionContentManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExtensionContentManagementServiceClient(cc grpc.ClientConnInterface) ExtensionContentManagementServiceClient {
	return &extensionContentManagementServiceClient{cc}
}

func (c *extensionContentManagementServiceClient) ImportAndReleaseExtensions(ctx context.Context, opts ...grpc.CallOption) (ExtensionContentManagementService_ImportAndReleaseExtensionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExtensionContentManagementService_ServiceDesc.Streams[0], ExtensionContentManagementService_ImportAndReleaseExtensions_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &extensionContentManagementServiceImportAndReleaseExtensionsClient{stream}
	return x, nil
}

type ExtensionContentManagementService_ImportAndReleaseExtensionsClient interface {
	Send(*ExtensionData) error
	CloseAndRecv() (*ImportAndReleaseExtensionsResponse, error)
	grpc.ClientStream
}

type extensionContentManagementServiceImportAndReleaseExtensionsClient struct {
	grpc.ClientStream
}

func (x *extensionContentManagementServiceImportAndReleaseExtensionsClient) Send(m *ExtensionData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *extensionContentManagementServiceImportAndReleaseExtensionsClient) CloseAndRecv() (*ImportAndReleaseExtensionsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ImportAndReleaseExtensionsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *extensionContentManagementServiceClient) ValidateExtensionItems(ctx context.Context, opts ...grpc.CallOption) (ExtensionContentManagementService_ValidateExtensionItemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExtensionContentManagementService_ServiceDesc.Streams[1], ExtensionContentManagementService_ValidateExtensionItems_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &extensionContentManagementServiceValidateExtensionItemsClient{stream}
	return x, nil
}

type ExtensionContentManagementService_ValidateExtensionItemsClient interface {
	Send(*ExtensionData) error
	CloseAndRecv() (*ValidateExtensionItemsResponse, error)
	grpc.ClientStream
}

type extensionContentManagementServiceValidateExtensionItemsClient struct {
	grpc.ClientStream
}

func (x *extensionContentManagementServiceValidateExtensionItemsClient) Send(m *ExtensionData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *extensionContentManagementServiceValidateExtensionItemsClient) CloseAndRecv() (*ValidateExtensionItemsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ValidateExtensionItemsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExtensionContentManagementServiceServer is the server API for ExtensionContentManagementService service.
// All implementations must embed UnimplementedExtensionContentManagementServiceServer
// for forward compatibility
type ExtensionContentManagementServiceServer interface {
	ImportAndReleaseExtensions(ExtensionContentManagementService_ImportAndReleaseExtensionsServer) error
	ValidateExtensionItems(ExtensionContentManagementService_ValidateExtensionItemsServer) error
	mustEmbedUnimplementedExtensionContentManagementServiceServer()
}

// UnimplementedExtensionContentManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExtensionContentManagementServiceServer struct {
}

func (UnimplementedExtensionContentManagementServiceServer) ImportAndReleaseExtensions(ExtensionContentManagementService_ImportAndReleaseExtensionsServer) error {
	return status.Errorf(codes.Unimplemented, "method ImportAndReleaseExtensions not implemented")
}
func (UnimplementedExtensionContentManagementServiceServer) ValidateExtensionItems(ExtensionContentManagementService_ValidateExtensionItemsServer) error {
	return status.Errorf(codes.Unimplemented, "method ValidateExtensionItems not implemented")
}
func (UnimplementedExtensionContentManagementServiceServer) mustEmbedUnimplementedExtensionContentManagementServiceServer() {
}

// UnsafeExtensionContentManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExtensionContentManagementServiceServer will
// result in compilation errors.
type UnsafeExtensionContentManagementServiceServer interface {
	mustEmbedUnimplementedExtensionContentManagementServiceServer()
}

func RegisterExtensionContentManagementServiceServer(s grpc.ServiceRegistrar, srv ExtensionContentManagementServiceServer) {
	s.RegisterService(&ExtensionContentManagementService_ServiceDesc, srv)
}

func _ExtensionContentManagementService_ImportAndReleaseExtensions_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExtensionContentManagementServiceServer).ImportAndReleaseExtensions(&extensionContentManagementServiceImportAndReleaseExtensionsServer{stream})
}

type ExtensionContentManagementService_ImportAndReleaseExtensionsServer interface {
	SendAndClose(*ImportAndReleaseExtensionsResponse) error
	Recv() (*ExtensionData, error)
	grpc.ServerStream
}

type extensionContentManagementServiceImportAndReleaseExtensionsServer struct {
	grpc.ServerStream
}

func (x *extensionContentManagementServiceImportAndReleaseExtensionsServer) SendAndClose(m *ImportAndReleaseExtensionsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *extensionContentManagementServiceImportAndReleaseExtensionsServer) Recv() (*ExtensionData, error) {
	m := new(ExtensionData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ExtensionContentManagementService_ValidateExtensionItems_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExtensionContentManagementServiceServer).ValidateExtensionItems(&extensionContentManagementServiceValidateExtensionItemsServer{stream})
}

type ExtensionContentManagementService_ValidateExtensionItemsServer interface {
	SendAndClose(*ValidateExtensionItemsResponse) error
	Recv() (*ExtensionData, error)
	grpc.ServerStream
}

type extensionContentManagementServiceValidateExtensionItemsServer struct {
	grpc.ServerStream
}

func (x *extensionContentManagementServiceValidateExtensionItemsServer) SendAndClose(m *ValidateExtensionItemsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *extensionContentManagementServiceValidateExtensionItemsServer) Recv() (*ExtensionData, error) {
	m := new(ExtensionData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExtensionContentManagementService_ServiceDesc is the grpc.ServiceDesc for ExtensionContentManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExtensionContentManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.extensions.v1.ExtensionContentManagementService",
	HandlerType: (*ExtensionContentManagementServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ImportAndReleaseExtensions",
			Handler:       _ExtensionContentManagementService_ImportAndReleaseExtensions_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ValidateExtensionItems",
			Handler:       _ExtensionContentManagementService_ValidateExtensionItems_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "com/coralogix/extensions/v1/extension_content_management_service.proto",
}
