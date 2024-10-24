// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.2
// source: com/coralogix/tags/v1/tag_external_service.proto

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
	TagsExternalService_CreateExternalTag_FullMethodName  = "/com.coralogix.tags.v1.TagsExternalService/CreateExternalTag"
	TagsExternalService_AddExternalTag_FullMethodName     = "/com.coralogix.tags.v1.TagsExternalService/AddExternalTag"
	TagsExternalService_CreateBitbucketTag_FullMethodName = "/com.coralogix.tags.v1.TagsExternalService/CreateBitbucketTag"
	TagsExternalService_CreateTfsTag_FullMethodName       = "/com.coralogix.tags.v1.TagsExternalService/CreateTfsTag"
	TagsExternalService_CreateGitlabTag_FullMethodName    = "/com.coralogix.tags.v1.TagsExternalService/CreateGitlabTag"
)

// TagsExternalServiceClient is the client API for TagsExternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagsExternalServiceClient interface {
	CreateExternalTag(ctx context.Context, in *CreateExternalTagRequest, opts ...grpc.CallOption) (*CreateExternalTagResponse, error)
	AddExternalTag(ctx context.Context, in *AddExternalTagRequest, opts ...grpc.CallOption) (*AddExternalTagResponse, error)
	CreateBitbucketTag(ctx context.Context, in *CreateBitbucketTagRequest, opts ...grpc.CallOption) (*CreateBitbucketTagResponse, error)
	CreateTfsTag(ctx context.Context, in *CreateTfsTagRequest, opts ...grpc.CallOption) (*CreateTfsTagResponse, error)
	CreateGitlabTag(ctx context.Context, in *CreateGitlabTagRequest, opts ...grpc.CallOption) (*CreateGitlabTagResponse, error)
}

type tagsExternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagsExternalServiceClient(cc grpc.ClientConnInterface) TagsExternalServiceClient {
	return &tagsExternalServiceClient{cc}
}

func (c *tagsExternalServiceClient) CreateExternalTag(ctx context.Context, in *CreateExternalTagRequest, opts ...grpc.CallOption) (*CreateExternalTagResponse, error) {
	out := new(CreateExternalTagResponse)
	err := c.cc.Invoke(ctx, TagsExternalService_CreateExternalTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsExternalServiceClient) AddExternalTag(ctx context.Context, in *AddExternalTagRequest, opts ...grpc.CallOption) (*AddExternalTagResponse, error) {
	out := new(AddExternalTagResponse)
	err := c.cc.Invoke(ctx, TagsExternalService_AddExternalTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsExternalServiceClient) CreateBitbucketTag(ctx context.Context, in *CreateBitbucketTagRequest, opts ...grpc.CallOption) (*CreateBitbucketTagResponse, error) {
	out := new(CreateBitbucketTagResponse)
	err := c.cc.Invoke(ctx, TagsExternalService_CreateBitbucketTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsExternalServiceClient) CreateTfsTag(ctx context.Context, in *CreateTfsTagRequest, opts ...grpc.CallOption) (*CreateTfsTagResponse, error) {
	out := new(CreateTfsTagResponse)
	err := c.cc.Invoke(ctx, TagsExternalService_CreateTfsTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsExternalServiceClient) CreateGitlabTag(ctx context.Context, in *CreateGitlabTagRequest, opts ...grpc.CallOption) (*CreateGitlabTagResponse, error) {
	out := new(CreateGitlabTagResponse)
	err := c.cc.Invoke(ctx, TagsExternalService_CreateGitlabTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagsExternalServiceServer is the server API for TagsExternalService service.
// All implementations must embed UnimplementedTagsExternalServiceServer
// for forward compatibility
type TagsExternalServiceServer interface {
	CreateExternalTag(context.Context, *CreateExternalTagRequest) (*CreateExternalTagResponse, error)
	AddExternalTag(context.Context, *AddExternalTagRequest) (*AddExternalTagResponse, error)
	CreateBitbucketTag(context.Context, *CreateBitbucketTagRequest) (*CreateBitbucketTagResponse, error)
	CreateTfsTag(context.Context, *CreateTfsTagRequest) (*CreateTfsTagResponse, error)
	CreateGitlabTag(context.Context, *CreateGitlabTagRequest) (*CreateGitlabTagResponse, error)
	mustEmbedUnimplementedTagsExternalServiceServer()
}

// UnimplementedTagsExternalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagsExternalServiceServer struct {
}

func (UnimplementedTagsExternalServiceServer) CreateExternalTag(context.Context, *CreateExternalTagRequest) (*CreateExternalTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExternalTag not implemented")
}
func (UnimplementedTagsExternalServiceServer) AddExternalTag(context.Context, *AddExternalTagRequest) (*AddExternalTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExternalTag not implemented")
}
func (UnimplementedTagsExternalServiceServer) CreateBitbucketTag(context.Context, *CreateBitbucketTagRequest) (*CreateBitbucketTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBitbucketTag not implemented")
}
func (UnimplementedTagsExternalServiceServer) CreateTfsTag(context.Context, *CreateTfsTagRequest) (*CreateTfsTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTfsTag not implemented")
}
func (UnimplementedTagsExternalServiceServer) CreateGitlabTag(context.Context, *CreateGitlabTagRequest) (*CreateGitlabTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGitlabTag not implemented")
}
func (UnimplementedTagsExternalServiceServer) mustEmbedUnimplementedTagsExternalServiceServer() {}

// UnsafeTagsExternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagsExternalServiceServer will
// result in compilation errors.
type UnsafeTagsExternalServiceServer interface {
	mustEmbedUnimplementedTagsExternalServiceServer()
}

func RegisterTagsExternalServiceServer(s grpc.ServiceRegistrar, srv TagsExternalServiceServer) {
	s.RegisterService(&TagsExternalService_ServiceDesc, srv)
}

func _TagsExternalService_CreateExternalTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExternalTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsExternalServiceServer).CreateExternalTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsExternalService_CreateExternalTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsExternalServiceServer).CreateExternalTag(ctx, req.(*CreateExternalTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsExternalService_AddExternalTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddExternalTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsExternalServiceServer).AddExternalTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsExternalService_AddExternalTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsExternalServiceServer).AddExternalTag(ctx, req.(*AddExternalTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsExternalService_CreateBitbucketTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBitbucketTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsExternalServiceServer).CreateBitbucketTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsExternalService_CreateBitbucketTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsExternalServiceServer).CreateBitbucketTag(ctx, req.(*CreateBitbucketTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsExternalService_CreateTfsTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTfsTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsExternalServiceServer).CreateTfsTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsExternalService_CreateTfsTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsExternalServiceServer).CreateTfsTag(ctx, req.(*CreateTfsTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsExternalService_CreateGitlabTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGitlabTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsExternalServiceServer).CreateGitlabTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagsExternalService_CreateGitlabTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsExternalServiceServer).CreateGitlabTag(ctx, req.(*CreateGitlabTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TagsExternalService_ServiceDesc is the grpc.ServiceDesc for TagsExternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TagsExternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.tags.v1.TagsExternalService",
	HandlerType: (*TagsExternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExternalTag",
			Handler:    _TagsExternalService_CreateExternalTag_Handler,
		},
		{
			MethodName: "AddExternalTag",
			Handler:    _TagsExternalService_AddExternalTag_Handler,
		},
		{
			MethodName: "CreateBitbucketTag",
			Handler:    _TagsExternalService_CreateBitbucketTag_Handler,
		},
		{
			MethodName: "CreateTfsTag",
			Handler:    _TagsExternalService_CreateTfsTag_Handler,
		},
		{
			MethodName: "CreateGitlabTag",
			Handler:    _TagsExternalService_CreateGitlabTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/tags/v1/tag_external_service.proto",
}
