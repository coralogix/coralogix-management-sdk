// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: com/coralogix/tags/v1/tag_service.proto

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
	TagsService_CreateTag_FullMethodName         = "/com.coralogix.tags.v1.TagsService/CreateTag"
	TagsService_GetTags_FullMethodName           = "/com.coralogix.tags.v1.TagsService/GetTags"
	TagsService_UpdateTag_FullMethodName         = "/com.coralogix.tags.v1.TagsService/UpdateTag"
	TagsService_DeleteTag_FullMethodName         = "/com.coralogix.tags.v1.TagsService/DeleteTag"
	TagsService_GetTagSummary_FullMethodName     = "/com.coralogix.tags.v1.TagsService/GetTagSummary"
	TagsService_GetTagAlerts_FullMethodName      = "/com.coralogix.tags.v1.TagsService/GetTagAlerts"
	TagsService_GetTagErrorVolume_FullMethodName = "/com.coralogix.tags.v1.TagsService/GetTagErrorVolume"
)

// TagsServiceClient is the client API for TagsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagsServiceClient interface {
	CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*CreateTagResponse, error)
	GetTags(ctx context.Context, in *GetTagsRequest, opts ...grpc.CallOption) (*GetTagsResponse, error)
	UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*UpdateTagResponse, error)
	DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*DeleteTagResponse, error)
	GetTagSummary(ctx context.Context, in *GetTagSummaryRequest, opts ...grpc.CallOption) (*GetTagSummaryResponse, error)
	GetTagAlerts(ctx context.Context, in *GetTagAlertsRequest, opts ...grpc.CallOption) (*GetTagAlertsResponse, error)
	GetTagErrorVolume(ctx context.Context, in *GetTagErrorVolumeRequest, opts ...grpc.CallOption) (*GetTagErrorVolumeResponse, error)
}

type tagsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagsServiceClient(cc grpc.ClientConnInterface) TagsServiceClient {
	return &tagsServiceClient{cc}
}

func (c *tagsServiceClient) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*CreateTagResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTagResponse)
	err := c.cc.Invoke(ctx, TagsService_CreateTag_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) GetTags(ctx context.Context, in *GetTagsRequest, opts ...grpc.CallOption) (*GetTagsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTagsResponse)
	err := c.cc.Invoke(ctx, TagsService_GetTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*UpdateTagResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTagResponse)
	err := c.cc.Invoke(ctx, TagsService_UpdateTag_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*DeleteTagResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTagResponse)
	err := c.cc.Invoke(ctx, TagsService_DeleteTag_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) GetTagSummary(ctx context.Context, in *GetTagSummaryRequest, opts ...grpc.CallOption) (*GetTagSummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTagSummaryResponse)
	err := c.cc.Invoke(ctx, TagsService_GetTagSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) GetTagAlerts(ctx context.Context, in *GetTagAlertsRequest, opts ...grpc.CallOption) (*GetTagAlertsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTagAlertsResponse)
	err := c.cc.Invoke(ctx, TagsService_GetTagAlerts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsServiceClient) GetTagErrorVolume(ctx context.Context, in *GetTagErrorVolumeRequest, opts ...grpc.CallOption) (*GetTagErrorVolumeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTagErrorVolumeResponse)
	err := c.cc.Invoke(ctx, TagsService_GetTagErrorVolume_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagsServiceServer is the server API for TagsService service.
// All implementations must embed UnimplementedTagsServiceServer
// for forward compatibility.
type TagsServiceServer interface {
	CreateTag(context.Context, *CreateTagRequest) (*CreateTagResponse, error)
	GetTags(context.Context, *GetTagsRequest) (*GetTagsResponse, error)
	UpdateTag(context.Context, *UpdateTagRequest) (*UpdateTagResponse, error)
	DeleteTag(context.Context, *DeleteTagRequest) (*DeleteTagResponse, error)
	GetTagSummary(context.Context, *GetTagSummaryRequest) (*GetTagSummaryResponse, error)
	GetTagAlerts(context.Context, *GetTagAlertsRequest) (*GetTagAlertsResponse, error)
	GetTagErrorVolume(context.Context, *GetTagErrorVolumeRequest) (*GetTagErrorVolumeResponse, error)
	mustEmbedUnimplementedTagsServiceServer()
}

// UnimplementedTagsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTagsServiceServer struct{}

func (UnimplementedTagsServiceServer) CreateTag(context.Context, *CreateTagRequest) (*CreateTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}
func (UnimplementedTagsServiceServer) GetTags(context.Context, *GetTagsRequest) (*GetTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
}
func (UnimplementedTagsServiceServer) UpdateTag(context.Context, *UpdateTagRequest) (*UpdateTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}
func (UnimplementedTagsServiceServer) DeleteTag(context.Context, *DeleteTagRequest) (*DeleteTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}
func (UnimplementedTagsServiceServer) GetTagSummary(context.Context, *GetTagSummaryRequest) (*GetTagSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagSummary not implemented")
}
func (UnimplementedTagsServiceServer) GetTagAlerts(context.Context, *GetTagAlertsRequest) (*GetTagAlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagAlerts not implemented")
}
func (UnimplementedTagsServiceServer) GetTagErrorVolume(context.Context, *GetTagErrorVolumeRequest) (*GetTagErrorVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagErrorVolume not implemented")
}
func (UnimplementedTagsServiceServer) mustEmbedUnimplementedTagsServiceServer() {}
func (UnimplementedTagsServiceServer) testEmbeddedByValue()                     {}

// UnsafeTagsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagsServiceServer will
// result in compilation errors.
type UnsafeTagsServiceServer interface {
	mustEmbedUnimplementedTagsServiceServer()
}

func RegisterTagsServiceServer(s grpc.ServiceRegistrar, srv TagsServiceServer) {
	// If the following call pancis, it indicates UnimplementedTagsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TagsService_ServiceDesc, srv)
}

func _TagsService_CreateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).CreateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_CreateTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).CreateTag(ctx, req.(*CreateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_GetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).GetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_GetTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).GetTags(ctx, req.(*GetTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_UpdateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).UpdateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_UpdateTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).UpdateTag(ctx, req.(*UpdateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_DeleteTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).DeleteTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_DeleteTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).DeleteTag(ctx, req.(*DeleteTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_GetTagSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).GetTagSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_GetTagSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).GetTagSummary(ctx, req.(*GetTagSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_GetTagAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).GetTagAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_GetTagAlerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).GetTagAlerts(ctx, req.(*GetTagAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsService_GetTagErrorVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagErrorVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServiceServer).GetTagErrorVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsService_GetTagErrorVolume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServiceServer).GetTagErrorVolume(ctx, req.(*GetTagErrorVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TagsService_ServiceDesc is the grpc.ServiceDesc for TagsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TagsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.tags.v1.TagsService",
	HandlerType: (*TagsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTag",
			Handler:    _TagsService_CreateTag_Handler,
		},
		{
			MethodName: "GetTags",
			Handler:    _TagsService_GetTags_Handler,
		},
		{
			MethodName: "UpdateTag",
			Handler:    _TagsService_UpdateTag_Handler,
		},
		{
			MethodName: "DeleteTag",
			Handler:    _TagsService_DeleteTag_Handler,
		},
		{
			MethodName: "GetTagSummary",
			Handler:    _TagsService_GetTagSummary_Handler,
		},
		{
			MethodName: "GetTagAlerts",
			Handler:    _TagsService_GetTagAlerts_Handler,
		},
		{
			MethodName: "GetTagErrorVolume",
			Handler:    _TagsService_GetTagErrorVolume_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/tags/v1/tag_service.proto",
}
