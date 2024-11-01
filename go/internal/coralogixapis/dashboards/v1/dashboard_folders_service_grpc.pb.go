// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: com/coralogixapis/dashboards/v1/services/dashboard_folders_service.proto

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
	DashboardFoldersService_ListDashboardFolders_FullMethodName   = "/com.coralogixapis.dashboards.v1.services.DashboardFoldersService/ListDashboardFolders"
	DashboardFoldersService_GetDashboardFolder_FullMethodName     = "/com.coralogixapis.dashboards.v1.services.DashboardFoldersService/GetDashboardFolder"
	DashboardFoldersService_CreateDashboardFolder_FullMethodName  = "/com.coralogixapis.dashboards.v1.services.DashboardFoldersService/CreateDashboardFolder"
	DashboardFoldersService_ReplaceDashboardFolder_FullMethodName = "/com.coralogixapis.dashboards.v1.services.DashboardFoldersService/ReplaceDashboardFolder"
	DashboardFoldersService_DeleteDashboardFolder_FullMethodName  = "/com.coralogixapis.dashboards.v1.services.DashboardFoldersService/DeleteDashboardFolder"
)

// DashboardFoldersServiceClient is the client API for DashboardFoldersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DashboardFoldersServiceClient interface {
	ListDashboardFolders(ctx context.Context, in *ListDashboardFoldersRequest, opts ...grpc.CallOption) (*ListDashboardFoldersResponse, error)
	GetDashboardFolder(ctx context.Context, in *GetDashboardFolderRequest, opts ...grpc.CallOption) (*GetDashboardFolderResponse, error)
	CreateDashboardFolder(ctx context.Context, in *CreateDashboardFolderRequest, opts ...grpc.CallOption) (*CreateDashboardFolderResponse, error)
	ReplaceDashboardFolder(ctx context.Context, in *ReplaceDashboardFolderRequest, opts ...grpc.CallOption) (*ReplaceDashboardFolderResponse, error)
	DeleteDashboardFolder(ctx context.Context, in *DeleteDashboardFolderRequest, opts ...grpc.CallOption) (*DeleteDashboardFolderResponse, error)
}

type dashboardFoldersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDashboardFoldersServiceClient(cc grpc.ClientConnInterface) DashboardFoldersServiceClient {
	return &dashboardFoldersServiceClient{cc}
}

func (c *dashboardFoldersServiceClient) ListDashboardFolders(ctx context.Context, in *ListDashboardFoldersRequest, opts ...grpc.CallOption) (*ListDashboardFoldersResponse, error) {
	out := new(ListDashboardFoldersResponse)
	err := c.cc.Invoke(ctx, DashboardFoldersService_ListDashboardFolders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardFoldersServiceClient) GetDashboardFolder(ctx context.Context, in *GetDashboardFolderRequest, opts ...grpc.CallOption) (*GetDashboardFolderResponse, error) {
	out := new(GetDashboardFolderResponse)
	err := c.cc.Invoke(ctx, DashboardFoldersService_GetDashboardFolder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardFoldersServiceClient) CreateDashboardFolder(ctx context.Context, in *CreateDashboardFolderRequest, opts ...grpc.CallOption) (*CreateDashboardFolderResponse, error) {
	out := new(CreateDashboardFolderResponse)
	err := c.cc.Invoke(ctx, DashboardFoldersService_CreateDashboardFolder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardFoldersServiceClient) ReplaceDashboardFolder(ctx context.Context, in *ReplaceDashboardFolderRequest, opts ...grpc.CallOption) (*ReplaceDashboardFolderResponse, error) {
	out := new(ReplaceDashboardFolderResponse)
	err := c.cc.Invoke(ctx, DashboardFoldersService_ReplaceDashboardFolder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardFoldersServiceClient) DeleteDashboardFolder(ctx context.Context, in *DeleteDashboardFolderRequest, opts ...grpc.CallOption) (*DeleteDashboardFolderResponse, error) {
	out := new(DeleteDashboardFolderResponse)
	err := c.cc.Invoke(ctx, DashboardFoldersService_DeleteDashboardFolder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DashboardFoldersServiceServer is the server API for DashboardFoldersService service.
// All implementations must embed UnimplementedDashboardFoldersServiceServer
// for forward compatibility
type DashboardFoldersServiceServer interface {
	ListDashboardFolders(context.Context, *ListDashboardFoldersRequest) (*ListDashboardFoldersResponse, error)
	GetDashboardFolder(context.Context, *GetDashboardFolderRequest) (*GetDashboardFolderResponse, error)
	CreateDashboardFolder(context.Context, *CreateDashboardFolderRequest) (*CreateDashboardFolderResponse, error)
	ReplaceDashboardFolder(context.Context, *ReplaceDashboardFolderRequest) (*ReplaceDashboardFolderResponse, error)
	DeleteDashboardFolder(context.Context, *DeleteDashboardFolderRequest) (*DeleteDashboardFolderResponse, error)
	mustEmbedUnimplementedDashboardFoldersServiceServer()
}

// UnimplementedDashboardFoldersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDashboardFoldersServiceServer struct {
}

func (UnimplementedDashboardFoldersServiceServer) ListDashboardFolders(context.Context, *ListDashboardFoldersRequest) (*ListDashboardFoldersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDashboardFolders not implemented")
}
func (UnimplementedDashboardFoldersServiceServer) GetDashboardFolder(context.Context, *GetDashboardFolderRequest) (*GetDashboardFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDashboardFolder not implemented")
}
func (UnimplementedDashboardFoldersServiceServer) CreateDashboardFolder(context.Context, *CreateDashboardFolderRequest) (*CreateDashboardFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDashboardFolder not implemented")
}
func (UnimplementedDashboardFoldersServiceServer) ReplaceDashboardFolder(context.Context, *ReplaceDashboardFolderRequest) (*ReplaceDashboardFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceDashboardFolder not implemented")
}
func (UnimplementedDashboardFoldersServiceServer) DeleteDashboardFolder(context.Context, *DeleteDashboardFolderRequest) (*DeleteDashboardFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDashboardFolder not implemented")
}
func (UnimplementedDashboardFoldersServiceServer) mustEmbedUnimplementedDashboardFoldersServiceServer() {
}

// UnsafeDashboardFoldersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DashboardFoldersServiceServer will
// result in compilation errors.
type UnsafeDashboardFoldersServiceServer interface {
	mustEmbedUnimplementedDashboardFoldersServiceServer()
}

func RegisterDashboardFoldersServiceServer(s grpc.ServiceRegistrar, srv DashboardFoldersServiceServer) {
	s.RegisterService(&DashboardFoldersService_ServiceDesc, srv)
}

func _DashboardFoldersService_ListDashboardFolders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDashboardFoldersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardFoldersServiceServer).ListDashboardFolders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardFoldersService_ListDashboardFolders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardFoldersServiceServer).ListDashboardFolders(ctx, req.(*ListDashboardFoldersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardFoldersService_GetDashboardFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDashboardFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardFoldersServiceServer).GetDashboardFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardFoldersService_GetDashboardFolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardFoldersServiceServer).GetDashboardFolder(ctx, req.(*GetDashboardFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardFoldersService_CreateDashboardFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDashboardFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardFoldersServiceServer).CreateDashboardFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardFoldersService_CreateDashboardFolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardFoldersServiceServer).CreateDashboardFolder(ctx, req.(*CreateDashboardFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardFoldersService_ReplaceDashboardFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceDashboardFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardFoldersServiceServer).ReplaceDashboardFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardFoldersService_ReplaceDashboardFolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardFoldersServiceServer).ReplaceDashboardFolder(ctx, req.(*ReplaceDashboardFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardFoldersService_DeleteDashboardFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDashboardFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardFoldersServiceServer).DeleteDashboardFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardFoldersService_DeleteDashboardFolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardFoldersServiceServer).DeleteDashboardFolder(ctx, req.(*DeleteDashboardFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DashboardFoldersService_ServiceDesc is the grpc.ServiceDesc for DashboardFoldersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DashboardFoldersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.dashboards.v1.services.DashboardFoldersService",
	HandlerType: (*DashboardFoldersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDashboardFolders",
			Handler:    _DashboardFoldersService_ListDashboardFolders_Handler,
		},
		{
			MethodName: "GetDashboardFolder",
			Handler:    _DashboardFoldersService_GetDashboardFolder_Handler,
		},
		{
			MethodName: "CreateDashboardFolder",
			Handler:    _DashboardFoldersService_CreateDashboardFolder_Handler,
		},
		{
			MethodName: "ReplaceDashboardFolder",
			Handler:    _DashboardFoldersService_ReplaceDashboardFolder_Handler,
		},
		{
			MethodName: "DeleteDashboardFolder",
			Handler:    _DashboardFoldersService_DeleteDashboardFolder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/dashboards/v1/services/dashboard_folders_service.proto",
}
