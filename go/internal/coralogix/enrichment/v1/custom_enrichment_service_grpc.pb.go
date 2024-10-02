// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.28.2
// source: com/coralogix/enrichment/v1/custom_enrichment_service.proto

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
	CustomEnrichmentService_GetCustomEnrichment_FullMethodName        = "/com.coralogix.enrichment.v1.CustomEnrichmentService/GetCustomEnrichment"
	CustomEnrichmentService_GetCustomEnrichments_FullMethodName       = "/com.coralogix.enrichment.v1.CustomEnrichmentService/GetCustomEnrichments"
	CustomEnrichmentService_CreateCustomEnrichment_FullMethodName     = "/com.coralogix.enrichment.v1.CustomEnrichmentService/CreateCustomEnrichment"
	CustomEnrichmentService_UpdateCustomEnrichment_FullMethodName     = "/com.coralogix.enrichment.v1.CustomEnrichmentService/UpdateCustomEnrichment"
	CustomEnrichmentService_DeleteCustomEnrichment_FullMethodName     = "/com.coralogix.enrichment.v1.CustomEnrichmentService/DeleteCustomEnrichment"
	CustomEnrichmentService_SearchCustomEnrichmentData_FullMethodName = "/com.coralogix.enrichment.v1.CustomEnrichmentService/SearchCustomEnrichmentData"
)

// CustomEnrichmentServiceClient is the client API for CustomEnrichmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomEnrichmentServiceClient interface {
	GetCustomEnrichment(ctx context.Context, in *GetCustomEnrichmentRequest, opts ...grpc.CallOption) (*GetCustomEnrichmentResponse, error)
	GetCustomEnrichments(ctx context.Context, in *GetCustomEnrichmentsRequest, opts ...grpc.CallOption) (*GetCustomEnrichmentsResponse, error)
	CreateCustomEnrichment(ctx context.Context, in *CreateCustomEnrichmentRequest, opts ...grpc.CallOption) (*CreateCustomEnrichmentResponse, error)
	UpdateCustomEnrichment(ctx context.Context, in *UpdateCustomEnrichmentRequest, opts ...grpc.CallOption) (*UpdateCustomEnrichmentResponse, error)
	DeleteCustomEnrichment(ctx context.Context, in *DeleteCustomEnrichmentRequest, opts ...grpc.CallOption) (*DeleteCustomEnrichmentResponse, error)
	SearchCustomEnrichmentData(ctx context.Context, in *SearchCustomEnrichmentDataRequest, opts ...grpc.CallOption) (*SearchCustomEnrichmentDataResponse, error)
}

type customEnrichmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomEnrichmentServiceClient(cc grpc.ClientConnInterface) CustomEnrichmentServiceClient {
	return &customEnrichmentServiceClient{cc}
}

func (c *customEnrichmentServiceClient) GetCustomEnrichment(ctx context.Context, in *GetCustomEnrichmentRequest, opts ...grpc.CallOption) (*GetCustomEnrichmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCustomEnrichmentResponse)
	err := c.cc.Invoke(ctx, CustomEnrichmentService_GetCustomEnrichment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customEnrichmentServiceClient) GetCustomEnrichments(ctx context.Context, in *GetCustomEnrichmentsRequest, opts ...grpc.CallOption) (*GetCustomEnrichmentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCustomEnrichmentsResponse)
	err := c.cc.Invoke(ctx, CustomEnrichmentService_GetCustomEnrichments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customEnrichmentServiceClient) CreateCustomEnrichment(ctx context.Context, in *CreateCustomEnrichmentRequest, opts ...grpc.CallOption) (*CreateCustomEnrichmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCustomEnrichmentResponse)
	err := c.cc.Invoke(ctx, CustomEnrichmentService_CreateCustomEnrichment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customEnrichmentServiceClient) UpdateCustomEnrichment(ctx context.Context, in *UpdateCustomEnrichmentRequest, opts ...grpc.CallOption) (*UpdateCustomEnrichmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCustomEnrichmentResponse)
	err := c.cc.Invoke(ctx, CustomEnrichmentService_UpdateCustomEnrichment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customEnrichmentServiceClient) DeleteCustomEnrichment(ctx context.Context, in *DeleteCustomEnrichmentRequest, opts ...grpc.CallOption) (*DeleteCustomEnrichmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCustomEnrichmentResponse)
	err := c.cc.Invoke(ctx, CustomEnrichmentService_DeleteCustomEnrichment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customEnrichmentServiceClient) SearchCustomEnrichmentData(ctx context.Context, in *SearchCustomEnrichmentDataRequest, opts ...grpc.CallOption) (*SearchCustomEnrichmentDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchCustomEnrichmentDataResponse)
	err := c.cc.Invoke(ctx, CustomEnrichmentService_SearchCustomEnrichmentData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomEnrichmentServiceServer is the server API for CustomEnrichmentService service.
// All implementations must embed UnimplementedCustomEnrichmentServiceServer
// for forward compatibility
type CustomEnrichmentServiceServer interface {
	GetCustomEnrichment(context.Context, *GetCustomEnrichmentRequest) (*GetCustomEnrichmentResponse, error)
	GetCustomEnrichments(context.Context, *GetCustomEnrichmentsRequest) (*GetCustomEnrichmentsResponse, error)
	CreateCustomEnrichment(context.Context, *CreateCustomEnrichmentRequest) (*CreateCustomEnrichmentResponse, error)
	UpdateCustomEnrichment(context.Context, *UpdateCustomEnrichmentRequest) (*UpdateCustomEnrichmentResponse, error)
	DeleteCustomEnrichment(context.Context, *DeleteCustomEnrichmentRequest) (*DeleteCustomEnrichmentResponse, error)
	SearchCustomEnrichmentData(context.Context, *SearchCustomEnrichmentDataRequest) (*SearchCustomEnrichmentDataResponse, error)
	mustEmbedUnimplementedCustomEnrichmentServiceServer()
}

// UnimplementedCustomEnrichmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomEnrichmentServiceServer struct {
}

func (UnimplementedCustomEnrichmentServiceServer) GetCustomEnrichment(context.Context, *GetCustomEnrichmentRequest) (*GetCustomEnrichmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomEnrichment not implemented")
}
func (UnimplementedCustomEnrichmentServiceServer) GetCustomEnrichments(context.Context, *GetCustomEnrichmentsRequest) (*GetCustomEnrichmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomEnrichments not implemented")
}
func (UnimplementedCustomEnrichmentServiceServer) CreateCustomEnrichment(context.Context, *CreateCustomEnrichmentRequest) (*CreateCustomEnrichmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomEnrichment not implemented")
}
func (UnimplementedCustomEnrichmentServiceServer) UpdateCustomEnrichment(context.Context, *UpdateCustomEnrichmentRequest) (*UpdateCustomEnrichmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomEnrichment not implemented")
}
func (UnimplementedCustomEnrichmentServiceServer) DeleteCustomEnrichment(context.Context, *DeleteCustomEnrichmentRequest) (*DeleteCustomEnrichmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomEnrichment not implemented")
}
func (UnimplementedCustomEnrichmentServiceServer) SearchCustomEnrichmentData(context.Context, *SearchCustomEnrichmentDataRequest) (*SearchCustomEnrichmentDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCustomEnrichmentData not implemented")
}
func (UnimplementedCustomEnrichmentServiceServer) mustEmbedUnimplementedCustomEnrichmentServiceServer() {
}

// UnsafeCustomEnrichmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomEnrichmentServiceServer will
// result in compilation errors.
type UnsafeCustomEnrichmentServiceServer interface {
	mustEmbedUnimplementedCustomEnrichmentServiceServer()
}

func RegisterCustomEnrichmentServiceServer(s grpc.ServiceRegistrar, srv CustomEnrichmentServiceServer) {
	s.RegisterService(&CustomEnrichmentService_ServiceDesc, srv)
}

func _CustomEnrichmentService_GetCustomEnrichment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomEnrichmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomEnrichmentServiceServer).GetCustomEnrichment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomEnrichmentService_GetCustomEnrichment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomEnrichmentServiceServer).GetCustomEnrichment(ctx, req.(*GetCustomEnrichmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomEnrichmentService_GetCustomEnrichments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomEnrichmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomEnrichmentServiceServer).GetCustomEnrichments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomEnrichmentService_GetCustomEnrichments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomEnrichmentServiceServer).GetCustomEnrichments(ctx, req.(*GetCustomEnrichmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomEnrichmentService_CreateCustomEnrichment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomEnrichmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomEnrichmentServiceServer).CreateCustomEnrichment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomEnrichmentService_CreateCustomEnrichment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomEnrichmentServiceServer).CreateCustomEnrichment(ctx, req.(*CreateCustomEnrichmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomEnrichmentService_UpdateCustomEnrichment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomEnrichmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomEnrichmentServiceServer).UpdateCustomEnrichment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomEnrichmentService_UpdateCustomEnrichment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomEnrichmentServiceServer).UpdateCustomEnrichment(ctx, req.(*UpdateCustomEnrichmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomEnrichmentService_DeleteCustomEnrichment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCustomEnrichmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomEnrichmentServiceServer).DeleteCustomEnrichment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomEnrichmentService_DeleteCustomEnrichment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomEnrichmentServiceServer).DeleteCustomEnrichment(ctx, req.(*DeleteCustomEnrichmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomEnrichmentService_SearchCustomEnrichmentData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCustomEnrichmentDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomEnrichmentServiceServer).SearchCustomEnrichmentData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomEnrichmentService_SearchCustomEnrichmentData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomEnrichmentServiceServer).SearchCustomEnrichmentData(ctx, req.(*SearchCustomEnrichmentDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomEnrichmentService_ServiceDesc is the grpc.ServiceDesc for CustomEnrichmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomEnrichmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.enrichment.v1.CustomEnrichmentService",
	HandlerType: (*CustomEnrichmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomEnrichment",
			Handler:    _CustomEnrichmentService_GetCustomEnrichment_Handler,
		},
		{
			MethodName: "GetCustomEnrichments",
			Handler:    _CustomEnrichmentService_GetCustomEnrichments_Handler,
		},
		{
			MethodName: "CreateCustomEnrichment",
			Handler:    _CustomEnrichmentService_CreateCustomEnrichment_Handler,
		},
		{
			MethodName: "UpdateCustomEnrichment",
			Handler:    _CustomEnrichmentService_UpdateCustomEnrichment_Handler,
		},
		{
			MethodName: "DeleteCustomEnrichment",
			Handler:    _CustomEnrichmentService_DeleteCustomEnrichment_Handler,
		},
		{
			MethodName: "SearchCustomEnrichmentData",
			Handler:    _CustomEnrichmentService_SearchCustomEnrichmentData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/enrichment/v1/custom_enrichment_service.proto",
}
