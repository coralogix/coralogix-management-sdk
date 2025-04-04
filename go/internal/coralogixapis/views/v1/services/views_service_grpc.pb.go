// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: com/coralogixapis/views/v1/services/views_service.proto

package services

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
	ViewsService_CreateView_FullMethodName  = "/com.coralogixapis.views.v1.services.ViewsService/CreateView"
	ViewsService_ReplaceView_FullMethodName = "/com.coralogixapis.views.v1.services.ViewsService/ReplaceView"
	ViewsService_GetView_FullMethodName     = "/com.coralogixapis.views.v1.services.ViewsService/GetView"
	ViewsService_DeleteView_FullMethodName  = "/com.coralogixapis.views.v1.services.ViewsService/DeleteView"
	ViewsService_ListViews_FullMethodName   = "/com.coralogixapis.views.v1.services.ViewsService/ListViews"
)

// ViewsServiceClient is the client API for ViewsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ViewsServiceClient interface {
	CreateView(ctx context.Context, in *CreateViewRequest, opts ...grpc.CallOption) (*CreateViewResponse, error)
	ReplaceView(ctx context.Context, in *ReplaceViewRequest, opts ...grpc.CallOption) (*ReplaceViewResponse, error)
	GetView(ctx context.Context, in *GetViewRequest, opts ...grpc.CallOption) (*GetViewResponse, error)
	DeleteView(ctx context.Context, in *DeleteViewRequest, opts ...grpc.CallOption) (*DeleteViewResponse, error)
	ListViews(ctx context.Context, in *ListViewsRequest, opts ...grpc.CallOption) (*ListViewsResponse, error)
}

type viewsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewViewsServiceClient(cc grpc.ClientConnInterface) ViewsServiceClient {
	return &viewsServiceClient{cc}
}

func (c *viewsServiceClient) CreateView(ctx context.Context, in *CreateViewRequest, opts ...grpc.CallOption) (*CreateViewResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateViewResponse)
	err := c.cc.Invoke(ctx, ViewsService_CreateView_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewsServiceClient) ReplaceView(ctx context.Context, in *ReplaceViewRequest, opts ...grpc.CallOption) (*ReplaceViewResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplaceViewResponse)
	err := c.cc.Invoke(ctx, ViewsService_ReplaceView_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewsServiceClient) GetView(ctx context.Context, in *GetViewRequest, opts ...grpc.CallOption) (*GetViewResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetViewResponse)
	err := c.cc.Invoke(ctx, ViewsService_GetView_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewsServiceClient) DeleteView(ctx context.Context, in *DeleteViewRequest, opts ...grpc.CallOption) (*DeleteViewResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteViewResponse)
	err := c.cc.Invoke(ctx, ViewsService_DeleteView_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewsServiceClient) ListViews(ctx context.Context, in *ListViewsRequest, opts ...grpc.CallOption) (*ListViewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListViewsResponse)
	err := c.cc.Invoke(ctx, ViewsService_ListViews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ViewsServiceServer is the server API for ViewsService service.
// All implementations must embed UnimplementedViewsServiceServer
// for forward compatibility.
type ViewsServiceServer interface {
	CreateView(context.Context, *CreateViewRequest) (*CreateViewResponse, error)
	ReplaceView(context.Context, *ReplaceViewRequest) (*ReplaceViewResponse, error)
	GetView(context.Context, *GetViewRequest) (*GetViewResponse, error)
	DeleteView(context.Context, *DeleteViewRequest) (*DeleteViewResponse, error)
	ListViews(context.Context, *ListViewsRequest) (*ListViewsResponse, error)
	mustEmbedUnimplementedViewsServiceServer()
}

// UnimplementedViewsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedViewsServiceServer struct{}

func (UnimplementedViewsServiceServer) CreateView(context.Context, *CreateViewRequest) (*CreateViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateView not implemented")
}
func (UnimplementedViewsServiceServer) ReplaceView(context.Context, *ReplaceViewRequest) (*ReplaceViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceView not implemented")
}
func (UnimplementedViewsServiceServer) GetView(context.Context, *GetViewRequest) (*GetViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetView not implemented")
}
func (UnimplementedViewsServiceServer) DeleteView(context.Context, *DeleteViewRequest) (*DeleteViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteView not implemented")
}
func (UnimplementedViewsServiceServer) ListViews(context.Context, *ListViewsRequest) (*ListViewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListViews not implemented")
}
func (UnimplementedViewsServiceServer) mustEmbedUnimplementedViewsServiceServer() {}
func (UnimplementedViewsServiceServer) testEmbeddedByValue()                      {}

// UnsafeViewsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ViewsServiceServer will
// result in compilation errors.
type UnsafeViewsServiceServer interface {
	mustEmbedUnimplementedViewsServiceServer()
}

func RegisterViewsServiceServer(s grpc.ServiceRegistrar, srv ViewsServiceServer) {
	// If the following call pancis, it indicates UnimplementedViewsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ViewsService_ServiceDesc, srv)
}

func _ViewsService_CreateView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewsServiceServer).CreateView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViewsService_CreateView_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewsServiceServer).CreateView(ctx, req.(*CreateViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewsService_ReplaceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewsServiceServer).ReplaceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViewsService_ReplaceView_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewsServiceServer).ReplaceView(ctx, req.(*ReplaceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewsService_GetView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewsServiceServer).GetView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViewsService_GetView_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewsServiceServer).GetView(ctx, req.(*GetViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewsService_DeleteView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewsServiceServer).DeleteView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViewsService_DeleteView_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewsServiceServer).DeleteView(ctx, req.(*DeleteViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewsService_ListViews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListViewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewsServiceServer).ListViews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViewsService_ListViews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewsServiceServer).ListViews(ctx, req.(*ListViewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ViewsService_ServiceDesc is the grpc.ServiceDesc for ViewsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ViewsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.views.v1.services.ViewsService",
	HandlerType: (*ViewsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateView",
			Handler:    _ViewsService_CreateView_Handler,
		},
		{
			MethodName: "ReplaceView",
			Handler:    _ViewsService_ReplaceView_Handler,
		},
		{
			MethodName: "GetView",
			Handler:    _ViewsService_GetView_Handler,
		},
		{
			MethodName: "DeleteView",
			Handler:    _ViewsService_DeleteView_Handler,
		},
		{
			MethodName: "ListViews",
			Handler:    _ViewsService_ListViews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/views/v1/services/views_service.proto",
}
