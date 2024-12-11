// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
// source: com/coralogix/integrations/v1/permission_check_service.proto

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
	PermissionCheckService_CheckApiKeyPermission_FullMethodName = "/com.coralogix.integrations.v1.PermissionCheckService/CheckApiKeyPermission"
)

// PermissionCheckServiceClient is the client API for PermissionCheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionCheckServiceClient interface {
	CheckApiKeyPermission(ctx context.Context, in *CheckApiKeyPermissionRequest, opts ...grpc.CallOption) (*CheckApiKeyPermissionResponse, error)
}

type permissionCheckServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionCheckServiceClient(cc grpc.ClientConnInterface) PermissionCheckServiceClient {
	return &permissionCheckServiceClient{cc}
}

func (c *permissionCheckServiceClient) CheckApiKeyPermission(ctx context.Context, in *CheckApiKeyPermissionRequest, opts ...grpc.CallOption) (*CheckApiKeyPermissionResponse, error) {
	out := new(CheckApiKeyPermissionResponse)
	err := c.cc.Invoke(ctx, PermissionCheckService_CheckApiKeyPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionCheckServiceServer is the server API for PermissionCheckService service.
// All implementations must embed UnimplementedPermissionCheckServiceServer
// for forward compatibility
type PermissionCheckServiceServer interface {
	CheckApiKeyPermission(context.Context, *CheckApiKeyPermissionRequest) (*CheckApiKeyPermissionResponse, error)
	mustEmbedUnimplementedPermissionCheckServiceServer()
}

// UnimplementedPermissionCheckServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionCheckServiceServer struct {
}

func (UnimplementedPermissionCheckServiceServer) CheckApiKeyPermission(context.Context, *CheckApiKeyPermissionRequest) (*CheckApiKeyPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckApiKeyPermission not implemented")
}
func (UnimplementedPermissionCheckServiceServer) mustEmbedUnimplementedPermissionCheckServiceServer() {
}

// UnsafePermissionCheckServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionCheckServiceServer will
// result in compilation errors.
type UnsafePermissionCheckServiceServer interface {
	mustEmbedUnimplementedPermissionCheckServiceServer()
}

func RegisterPermissionCheckServiceServer(s grpc.ServiceRegistrar, srv PermissionCheckServiceServer) {
	s.RegisterService(&PermissionCheckService_ServiceDesc, srv)
}

func _PermissionCheckService_CheckApiKeyPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckApiKeyPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionCheckServiceServer).CheckApiKeyPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionCheckService_CheckApiKeyPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionCheckServiceServer).CheckApiKeyPermission(ctx, req.(*CheckApiKeyPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionCheckService_ServiceDesc is the grpc.ServiceDesc for PermissionCheckService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionCheckService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.integrations.v1.PermissionCheckService",
	HandlerType: (*PermissionCheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckApiKeyPermission",
			Handler:    _PermissionCheckService_CheckApiKeyPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/integrations/v1/permission_check_service.proto",
}
