// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: com/coralogixapis/dashboards/v1/services/dashboards_service.proto

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
	DashboardsService_CreateDashboard_FullMethodName         = "/com.coralogixapis.dashboards.v1.services.DashboardsService/CreateDashboard"
	DashboardsService_ReplaceDashboard_FullMethodName        = "/com.coralogixapis.dashboards.v1.services.DashboardsService/ReplaceDashboard"
	DashboardsService_DeleteDashboard_FullMethodName         = "/com.coralogixapis.dashboards.v1.services.DashboardsService/DeleteDashboard"
	DashboardsService_GetDashboard_FullMethodName            = "/com.coralogixapis.dashboards.v1.services.DashboardsService/GetDashboard"
	DashboardsService_PinDashboard_FullMethodName            = "/com.coralogixapis.dashboards.v1.services.DashboardsService/PinDashboard"
	DashboardsService_UnpinDashboard_FullMethodName          = "/com.coralogixapis.dashboards.v1.services.DashboardsService/UnpinDashboard"
	DashboardsService_ReplaceDefaultDashboard_FullMethodName = "/com.coralogixapis.dashboards.v1.services.DashboardsService/ReplaceDefaultDashboard"
	DashboardsService_AssignDashboardFolder_FullMethodName   = "/com.coralogixapis.dashboards.v1.services.DashboardsService/AssignDashboardFolder"
)

// DashboardsServiceClient is the client API for DashboardsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DashboardsServiceClient interface {
	CreateDashboard(ctx context.Context, in *CreateDashboardRequest, opts ...grpc.CallOption) (*CreateDashboardResponse, error)
	ReplaceDashboard(ctx context.Context, in *ReplaceDashboardRequest, opts ...grpc.CallOption) (*ReplaceDashboardResponse, error)
	DeleteDashboard(ctx context.Context, in *DeleteDashboardRequest, opts ...grpc.CallOption) (*DeleteDashboardResponse, error)
	GetDashboard(ctx context.Context, in *GetDashboardRequest, opts ...grpc.CallOption) (*GetDashboardResponse, error)
	PinDashboard(ctx context.Context, in *PinDashboardRequest, opts ...grpc.CallOption) (*PinDashboardResponse, error)
	UnpinDashboard(ctx context.Context, in *UnpinDashboardRequest, opts ...grpc.CallOption) (*UnpinDashboardResponse, error)
	ReplaceDefaultDashboard(ctx context.Context, in *ReplaceDefaultDashboardRequest, opts ...grpc.CallOption) (*ReplaceDefaultDashboardResponse, error)
	AssignDashboardFolder(ctx context.Context, in *AssignDashboardFolderRequest, opts ...grpc.CallOption) (*AssignDashboardFolderResponse, error)
}

type dashboardsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDashboardsServiceClient(cc grpc.ClientConnInterface) DashboardsServiceClient {
	return &dashboardsServiceClient{cc}
}

func (c *dashboardsServiceClient) CreateDashboard(ctx context.Context, in *CreateDashboardRequest, opts ...grpc.CallOption) (*CreateDashboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDashboardResponse)
	err := c.cc.Invoke(ctx, DashboardsService_CreateDashboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsServiceClient) ReplaceDashboard(ctx context.Context, in *ReplaceDashboardRequest, opts ...grpc.CallOption) (*ReplaceDashboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplaceDashboardResponse)
	err := c.cc.Invoke(ctx, DashboardsService_ReplaceDashboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsServiceClient) DeleteDashboard(ctx context.Context, in *DeleteDashboardRequest, opts ...grpc.CallOption) (*DeleteDashboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDashboardResponse)
	err := c.cc.Invoke(ctx, DashboardsService_DeleteDashboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsServiceClient) GetDashboard(ctx context.Context, in *GetDashboardRequest, opts ...grpc.CallOption) (*GetDashboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDashboardResponse)
	err := c.cc.Invoke(ctx, DashboardsService_GetDashboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsServiceClient) PinDashboard(ctx context.Context, in *PinDashboardRequest, opts ...grpc.CallOption) (*PinDashboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PinDashboardResponse)
	err := c.cc.Invoke(ctx, DashboardsService_PinDashboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsServiceClient) UnpinDashboard(ctx context.Context, in *UnpinDashboardRequest, opts ...grpc.CallOption) (*UnpinDashboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnpinDashboardResponse)
	err := c.cc.Invoke(ctx, DashboardsService_UnpinDashboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsServiceClient) ReplaceDefaultDashboard(ctx context.Context, in *ReplaceDefaultDashboardRequest, opts ...grpc.CallOption) (*ReplaceDefaultDashboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplaceDefaultDashboardResponse)
	err := c.cc.Invoke(ctx, DashboardsService_ReplaceDefaultDashboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardsServiceClient) AssignDashboardFolder(ctx context.Context, in *AssignDashboardFolderRequest, opts ...grpc.CallOption) (*AssignDashboardFolderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AssignDashboardFolderResponse)
	err := c.cc.Invoke(ctx, DashboardsService_AssignDashboardFolder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DashboardsServiceServer is the server API for DashboardsService service.
// All implementations must embed UnimplementedDashboardsServiceServer
// for forward compatibility.
type DashboardsServiceServer interface {
	CreateDashboard(context.Context, *CreateDashboardRequest) (*CreateDashboardResponse, error)
	ReplaceDashboard(context.Context, *ReplaceDashboardRequest) (*ReplaceDashboardResponse, error)
	DeleteDashboard(context.Context, *DeleteDashboardRequest) (*DeleteDashboardResponse, error)
	GetDashboard(context.Context, *GetDashboardRequest) (*GetDashboardResponse, error)
	PinDashboard(context.Context, *PinDashboardRequest) (*PinDashboardResponse, error)
	UnpinDashboard(context.Context, *UnpinDashboardRequest) (*UnpinDashboardResponse, error)
	ReplaceDefaultDashboard(context.Context, *ReplaceDefaultDashboardRequest) (*ReplaceDefaultDashboardResponse, error)
	AssignDashboardFolder(context.Context, *AssignDashboardFolderRequest) (*AssignDashboardFolderResponse, error)
	mustEmbedUnimplementedDashboardsServiceServer()
}

// UnimplementedDashboardsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDashboardsServiceServer struct{}

func (UnimplementedDashboardsServiceServer) CreateDashboard(context.Context, *CreateDashboardRequest) (*CreateDashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDashboard not implemented")
}
func (UnimplementedDashboardsServiceServer) ReplaceDashboard(context.Context, *ReplaceDashboardRequest) (*ReplaceDashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceDashboard not implemented")
}
func (UnimplementedDashboardsServiceServer) DeleteDashboard(context.Context, *DeleteDashboardRequest) (*DeleteDashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDashboard not implemented")
}
func (UnimplementedDashboardsServiceServer) GetDashboard(context.Context, *GetDashboardRequest) (*GetDashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDashboard not implemented")
}
func (UnimplementedDashboardsServiceServer) PinDashboard(context.Context, *PinDashboardRequest) (*PinDashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PinDashboard not implemented")
}
func (UnimplementedDashboardsServiceServer) UnpinDashboard(context.Context, *UnpinDashboardRequest) (*UnpinDashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnpinDashboard not implemented")
}
func (UnimplementedDashboardsServiceServer) ReplaceDefaultDashboard(context.Context, *ReplaceDefaultDashboardRequest) (*ReplaceDefaultDashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceDefaultDashboard not implemented")
}
func (UnimplementedDashboardsServiceServer) AssignDashboardFolder(context.Context, *AssignDashboardFolderRequest) (*AssignDashboardFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignDashboardFolder not implemented")
}
func (UnimplementedDashboardsServiceServer) mustEmbedUnimplementedDashboardsServiceServer() {}
func (UnimplementedDashboardsServiceServer) testEmbeddedByValue()                           {}

// UnsafeDashboardsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DashboardsServiceServer will
// result in compilation errors.
type UnsafeDashboardsServiceServer interface {
	mustEmbedUnimplementedDashboardsServiceServer()
}

func RegisterDashboardsServiceServer(s grpc.ServiceRegistrar, srv DashboardsServiceServer) {
	// If the following call pancis, it indicates UnimplementedDashboardsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DashboardsService_ServiceDesc, srv)
}

func _DashboardsService_CreateDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServiceServer).CreateDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardsService_CreateDashboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServiceServer).CreateDashboard(ctx, req.(*CreateDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardsService_ReplaceDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServiceServer).ReplaceDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardsService_ReplaceDashboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServiceServer).ReplaceDashboard(ctx, req.(*ReplaceDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardsService_DeleteDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServiceServer).DeleteDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardsService_DeleteDashboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServiceServer).DeleteDashboard(ctx, req.(*DeleteDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardsService_GetDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServiceServer).GetDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardsService_GetDashboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServiceServer).GetDashboard(ctx, req.(*GetDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardsService_PinDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PinDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServiceServer).PinDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardsService_PinDashboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServiceServer).PinDashboard(ctx, req.(*PinDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardsService_UnpinDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnpinDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServiceServer).UnpinDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardsService_UnpinDashboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServiceServer).UnpinDashboard(ctx, req.(*UnpinDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardsService_ReplaceDefaultDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceDefaultDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServiceServer).ReplaceDefaultDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardsService_ReplaceDefaultDashboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServiceServer).ReplaceDefaultDashboard(ctx, req.(*ReplaceDefaultDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardsService_AssignDashboardFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignDashboardFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardsServiceServer).AssignDashboardFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardsService_AssignDashboardFolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardsServiceServer).AssignDashboardFolder(ctx, req.(*AssignDashboardFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DashboardsService_ServiceDesc is the grpc.ServiceDesc for DashboardsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DashboardsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.dashboards.v1.services.DashboardsService",
	HandlerType: (*DashboardsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDashboard",
			Handler:    _DashboardsService_CreateDashboard_Handler,
		},
		{
			MethodName: "ReplaceDashboard",
			Handler:    _DashboardsService_ReplaceDashboard_Handler,
		},
		{
			MethodName: "DeleteDashboard",
			Handler:    _DashboardsService_DeleteDashboard_Handler,
		},
		{
			MethodName: "GetDashboard",
			Handler:    _DashboardsService_GetDashboard_Handler,
		},
		{
			MethodName: "PinDashboard",
			Handler:    _DashboardsService_PinDashboard_Handler,
		},
		{
			MethodName: "UnpinDashboard",
			Handler:    _DashboardsService_UnpinDashboard_Handler,
		},
		{
			MethodName: "ReplaceDefaultDashboard",
			Handler:    _DashboardsService_ReplaceDefaultDashboard_Handler,
		},
		{
			MethodName: "AssignDashboardFolder",
			Handler:    _DashboardsService_AssignDashboardFolder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/dashboards/v1/services/dashboards_service.proto",
}
