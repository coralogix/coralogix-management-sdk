// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: com/coralogixapis/aaa/apikeys/v3/api_keys.proto

package v3

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
	ApiKeysService_CreateApiKey_FullMethodName       = "/com.coralogixapis.aaa.apikeys.v3.ApiKeysService/CreateApiKey"
	ApiKeysService_GetApiKey_FullMethodName          = "/com.coralogixapis.aaa.apikeys.v3.ApiKeysService/GetApiKey"
	ApiKeysService_GetSendDataApiKeys_FullMethodName = "/com.coralogixapis.aaa.apikeys.v3.ApiKeysService/GetSendDataApiKeys"
	ApiKeysService_DeleteApiKey_FullMethodName       = "/com.coralogixapis.aaa.apikeys.v3.ApiKeysService/DeleteApiKey"
	ApiKeysService_UpdateApiKey_FullMethodName       = "/com.coralogixapis.aaa.apikeys.v3.ApiKeysService/UpdateApiKey"
)

// ApiKeysServiceClient is the client API for ApiKeysService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiKeysServiceClient interface {
	CreateApiKey(ctx context.Context, in *CreateApiKeyRequest, opts ...grpc.CallOption) (*CreateApiKeyResponse, error)
	GetApiKey(ctx context.Context, in *GetApiKeyRequest, opts ...grpc.CallOption) (*GetApiKeyResponse, error)
	GetSendDataApiKeys(ctx context.Context, in *GetSendDataApiKeysRequest, opts ...grpc.CallOption) (*GetSendDataApiKeysResponse, error)
	DeleteApiKey(ctx context.Context, in *DeleteApiKeyRequest, opts ...grpc.CallOption) (*DeleteApiKeyResponse, error)
	UpdateApiKey(ctx context.Context, in *UpdateApiKeyRequest, opts ...grpc.CallOption) (*UpdateApiKeyResponse, error)
}

type apiKeysServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiKeysServiceClient(cc grpc.ClientConnInterface) ApiKeysServiceClient {
	return &apiKeysServiceClient{cc}
}

func (c *apiKeysServiceClient) CreateApiKey(ctx context.Context, in *CreateApiKeyRequest, opts ...grpc.CallOption) (*CreateApiKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateApiKeyResponse)
	err := c.cc.Invoke(ctx, ApiKeysService_CreateApiKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeysServiceClient) GetApiKey(ctx context.Context, in *GetApiKeyRequest, opts ...grpc.CallOption) (*GetApiKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetApiKeyResponse)
	err := c.cc.Invoke(ctx, ApiKeysService_GetApiKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeysServiceClient) GetSendDataApiKeys(ctx context.Context, in *GetSendDataApiKeysRequest, opts ...grpc.CallOption) (*GetSendDataApiKeysResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSendDataApiKeysResponse)
	err := c.cc.Invoke(ctx, ApiKeysService_GetSendDataApiKeys_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeysServiceClient) DeleteApiKey(ctx context.Context, in *DeleteApiKeyRequest, opts ...grpc.CallOption) (*DeleteApiKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteApiKeyResponse)
	err := c.cc.Invoke(ctx, ApiKeysService_DeleteApiKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeysServiceClient) UpdateApiKey(ctx context.Context, in *UpdateApiKeyRequest, opts ...grpc.CallOption) (*UpdateApiKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateApiKeyResponse)
	err := c.cc.Invoke(ctx, ApiKeysService_UpdateApiKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiKeysServiceServer is the server API for ApiKeysService service.
// All implementations must embed UnimplementedApiKeysServiceServer
// for forward compatibility.
type ApiKeysServiceServer interface {
	CreateApiKey(context.Context, *CreateApiKeyRequest) (*CreateApiKeyResponse, error)
	GetApiKey(context.Context, *GetApiKeyRequest) (*GetApiKeyResponse, error)
	GetSendDataApiKeys(context.Context, *GetSendDataApiKeysRequest) (*GetSendDataApiKeysResponse, error)
	DeleteApiKey(context.Context, *DeleteApiKeyRequest) (*DeleteApiKeyResponse, error)
	UpdateApiKey(context.Context, *UpdateApiKeyRequest) (*UpdateApiKeyResponse, error)
	mustEmbedUnimplementedApiKeysServiceServer()
}

// UnimplementedApiKeysServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedApiKeysServiceServer struct{}

func (UnimplementedApiKeysServiceServer) CreateApiKey(context.Context, *CreateApiKeyRequest) (*CreateApiKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApiKey not implemented")
}
func (UnimplementedApiKeysServiceServer) GetApiKey(context.Context, *GetApiKeyRequest) (*GetApiKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiKey not implemented")
}
func (UnimplementedApiKeysServiceServer) GetSendDataApiKeys(context.Context, *GetSendDataApiKeysRequest) (*GetSendDataApiKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSendDataApiKeys not implemented")
}
func (UnimplementedApiKeysServiceServer) DeleteApiKey(context.Context, *DeleteApiKeyRequest) (*DeleteApiKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApiKey not implemented")
}
func (UnimplementedApiKeysServiceServer) UpdateApiKey(context.Context, *UpdateApiKeyRequest) (*UpdateApiKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApiKey not implemented")
}
func (UnimplementedApiKeysServiceServer) mustEmbedUnimplementedApiKeysServiceServer() {}
func (UnimplementedApiKeysServiceServer) testEmbeddedByValue()                        {}

// UnsafeApiKeysServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiKeysServiceServer will
// result in compilation errors.
type UnsafeApiKeysServiceServer interface {
	mustEmbedUnimplementedApiKeysServiceServer()
}

func RegisterApiKeysServiceServer(s grpc.ServiceRegistrar, srv ApiKeysServiceServer) {
	// If the following call pancis, it indicates UnimplementedApiKeysServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ApiKeysService_ServiceDesc, srv)
}

func _ApiKeysService_CreateApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeysServiceServer).CreateApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiKeysService_CreateApiKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeysServiceServer).CreateApiKey(ctx, req.(*CreateApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKeysService_GetApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeysServiceServer).GetApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiKeysService_GetApiKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeysServiceServer).GetApiKey(ctx, req.(*GetApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKeysService_GetSendDataApiKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSendDataApiKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeysServiceServer).GetSendDataApiKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiKeysService_GetSendDataApiKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeysServiceServer).GetSendDataApiKeys(ctx, req.(*GetSendDataApiKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKeysService_DeleteApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeysServiceServer).DeleteApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiKeysService_DeleteApiKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeysServiceServer).DeleteApiKey(ctx, req.(*DeleteApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKeysService_UpdateApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeysServiceServer).UpdateApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiKeysService_UpdateApiKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeysServiceServer).UpdateApiKey(ctx, req.(*UpdateApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiKeysService_ServiceDesc is the grpc.ServiceDesc for ApiKeysService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiKeysService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.aaa.apikeys.v3.ApiKeysService",
	HandlerType: (*ApiKeysServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApiKey",
			Handler:    _ApiKeysService_CreateApiKey_Handler,
		},
		{
			MethodName: "GetApiKey",
			Handler:    _ApiKeysService_GetApiKey_Handler,
		},
		{
			MethodName: "GetSendDataApiKeys",
			Handler:    _ApiKeysService_GetSendDataApiKeys_Handler,
		},
		{
			MethodName: "DeleteApiKey",
			Handler:    _ApiKeysService_DeleteApiKey_Handler,
		},
		{
			MethodName: "UpdateApiKey",
			Handler:    _ApiKeysService_UpdateApiKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/aaa/apikeys/v3/api_keys.proto",
}
