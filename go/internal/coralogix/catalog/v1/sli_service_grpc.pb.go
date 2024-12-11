// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
// source: com/coralogix/catalog/v1/sli_service.proto

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
	SliService_GetSlis_FullMethodName             = "/com.coralogix.catalog.v1.SliService/GetSlis"
	SliService_CreateSli_FullMethodName           = "/com.coralogix.catalog.v1.SliService/CreateSli"
	SliService_CreateSlis_FullMethodName          = "/com.coralogix.catalog.v1.SliService/CreateSlis"
	SliService_UpdateSli_FullMethodName           = "/com.coralogix.catalog.v1.SliService/UpdateSli"
	SliService_UpdateSlis_FullMethodName          = "/com.coralogix.catalog.v1.SliService/UpdateSlis"
	SliService_DeleteSli_FullMethodName           = "/com.coralogix.catalog.v1.SliService/DeleteSli"
	SliService_GetSliStatusHistory_FullMethodName = "/com.coralogix.catalog.v1.SliService/GetSliStatusHistory"
	SliService_GetSliE2MQuery_FullMethodName      = "/com.coralogix.catalog.v1.SliService/GetSliE2mQuery"
)

// SliServiceClient is the client API for SliService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SliServiceClient interface {
	GetSlis(ctx context.Context, in *GetSlisRequest, opts ...grpc.CallOption) (*GetSlisResponse, error)
	CreateSli(ctx context.Context, in *CreateSliRequest, opts ...grpc.CallOption) (*CreateSliResponse, error)
	CreateSlis(ctx context.Context, in *CreateSlisRequest, opts ...grpc.CallOption) (*CreateSlisResponse, error)
	UpdateSli(ctx context.Context, in *UpdateSliRequest, opts ...grpc.CallOption) (*UpdateSliResponse, error)
	UpdateSlis(ctx context.Context, in *UpdateSlisRequest, opts ...grpc.CallOption) (*UpdateSlisResponse, error)
	DeleteSli(ctx context.Context, in *DeleteSliRequest, opts ...grpc.CallOption) (*DeleteSliResponse, error)
	GetSliStatusHistory(ctx context.Context, in *GetSliStatusHistoryRequest, opts ...grpc.CallOption) (*GetSliStatusHistoryResponse, error)
	GetSliE2MQuery(ctx context.Context, in *GetSliE2MQueryRequest, opts ...grpc.CallOption) (*GetSliE2MQueryResponse, error)
}

type sliServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSliServiceClient(cc grpc.ClientConnInterface) SliServiceClient {
	return &sliServiceClient{cc}
}

func (c *sliServiceClient) GetSlis(ctx context.Context, in *GetSlisRequest, opts ...grpc.CallOption) (*GetSlisResponse, error) {
	out := new(GetSlisResponse)
	err := c.cc.Invoke(ctx, SliService_GetSlis_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sliServiceClient) CreateSli(ctx context.Context, in *CreateSliRequest, opts ...grpc.CallOption) (*CreateSliResponse, error) {
	out := new(CreateSliResponse)
	err := c.cc.Invoke(ctx, SliService_CreateSli_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sliServiceClient) CreateSlis(ctx context.Context, in *CreateSlisRequest, opts ...grpc.CallOption) (*CreateSlisResponse, error) {
	out := new(CreateSlisResponse)
	err := c.cc.Invoke(ctx, SliService_CreateSlis_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sliServiceClient) UpdateSli(ctx context.Context, in *UpdateSliRequest, opts ...grpc.CallOption) (*UpdateSliResponse, error) {
	out := new(UpdateSliResponse)
	err := c.cc.Invoke(ctx, SliService_UpdateSli_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sliServiceClient) UpdateSlis(ctx context.Context, in *UpdateSlisRequest, opts ...grpc.CallOption) (*UpdateSlisResponse, error) {
	out := new(UpdateSlisResponse)
	err := c.cc.Invoke(ctx, SliService_UpdateSlis_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sliServiceClient) DeleteSli(ctx context.Context, in *DeleteSliRequest, opts ...grpc.CallOption) (*DeleteSliResponse, error) {
	out := new(DeleteSliResponse)
	err := c.cc.Invoke(ctx, SliService_DeleteSli_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sliServiceClient) GetSliStatusHistory(ctx context.Context, in *GetSliStatusHistoryRequest, opts ...grpc.CallOption) (*GetSliStatusHistoryResponse, error) {
	out := new(GetSliStatusHistoryResponse)
	err := c.cc.Invoke(ctx, SliService_GetSliStatusHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sliServiceClient) GetSliE2MQuery(ctx context.Context, in *GetSliE2MQueryRequest, opts ...grpc.CallOption) (*GetSliE2MQueryResponse, error) {
	out := new(GetSliE2MQueryResponse)
	err := c.cc.Invoke(ctx, SliService_GetSliE2MQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SliServiceServer is the server API for SliService service.
// All implementations must embed UnimplementedSliServiceServer
// for forward compatibility
type SliServiceServer interface {
	GetSlis(context.Context, *GetSlisRequest) (*GetSlisResponse, error)
	CreateSli(context.Context, *CreateSliRequest) (*CreateSliResponse, error)
	CreateSlis(context.Context, *CreateSlisRequest) (*CreateSlisResponse, error)
	UpdateSli(context.Context, *UpdateSliRequest) (*UpdateSliResponse, error)
	UpdateSlis(context.Context, *UpdateSlisRequest) (*UpdateSlisResponse, error)
	DeleteSli(context.Context, *DeleteSliRequest) (*DeleteSliResponse, error)
	GetSliStatusHistory(context.Context, *GetSliStatusHistoryRequest) (*GetSliStatusHistoryResponse, error)
	GetSliE2MQuery(context.Context, *GetSliE2MQueryRequest) (*GetSliE2MQueryResponse, error)
	mustEmbedUnimplementedSliServiceServer()
}

// UnimplementedSliServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSliServiceServer struct {
}

func (UnimplementedSliServiceServer) GetSlis(context.Context, *GetSlisRequest) (*GetSlisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSlis not implemented")
}
func (UnimplementedSliServiceServer) CreateSli(context.Context, *CreateSliRequest) (*CreateSliResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSli not implemented")
}
func (UnimplementedSliServiceServer) CreateSlis(context.Context, *CreateSlisRequest) (*CreateSlisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSlis not implemented")
}
func (UnimplementedSliServiceServer) UpdateSli(context.Context, *UpdateSliRequest) (*UpdateSliResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSli not implemented")
}
func (UnimplementedSliServiceServer) UpdateSlis(context.Context, *UpdateSlisRequest) (*UpdateSlisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSlis not implemented")
}
func (UnimplementedSliServiceServer) DeleteSli(context.Context, *DeleteSliRequest) (*DeleteSliResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSli not implemented")
}
func (UnimplementedSliServiceServer) GetSliStatusHistory(context.Context, *GetSliStatusHistoryRequest) (*GetSliStatusHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSliStatusHistory not implemented")
}
func (UnimplementedSliServiceServer) GetSliE2MQuery(context.Context, *GetSliE2MQueryRequest) (*GetSliE2MQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSliE2MQuery not implemented")
}
func (UnimplementedSliServiceServer) mustEmbedUnimplementedSliServiceServer() {}

// UnsafeSliServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SliServiceServer will
// result in compilation errors.
type UnsafeSliServiceServer interface {
	mustEmbedUnimplementedSliServiceServer()
}

func RegisterSliServiceServer(s grpc.ServiceRegistrar, srv SliServiceServer) {
	s.RegisterService(&SliService_ServiceDesc, srv)
}

func _SliService_GetSlis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSlisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliServiceServer).GetSlis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SliService_GetSlis_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliServiceServer).GetSlis(ctx, req.(*GetSlisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SliService_CreateSli_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSliRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliServiceServer).CreateSli(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SliService_CreateSli_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliServiceServer).CreateSli(ctx, req.(*CreateSliRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SliService_CreateSlis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSlisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliServiceServer).CreateSlis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SliService_CreateSlis_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliServiceServer).CreateSlis(ctx, req.(*CreateSlisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SliService_UpdateSli_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSliRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliServiceServer).UpdateSli(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SliService_UpdateSli_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliServiceServer).UpdateSli(ctx, req.(*UpdateSliRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SliService_UpdateSlis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSlisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliServiceServer).UpdateSlis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SliService_UpdateSlis_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliServiceServer).UpdateSlis(ctx, req.(*UpdateSlisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SliService_DeleteSli_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSliRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliServiceServer).DeleteSli(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SliService_DeleteSli_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliServiceServer).DeleteSli(ctx, req.(*DeleteSliRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SliService_GetSliStatusHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSliStatusHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliServiceServer).GetSliStatusHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SliService_GetSliStatusHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliServiceServer).GetSliStatusHistory(ctx, req.(*GetSliStatusHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SliService_GetSliE2MQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSliE2MQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliServiceServer).GetSliE2MQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SliService_GetSliE2MQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliServiceServer).GetSliE2MQuery(ctx, req.(*GetSliE2MQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SliService_ServiceDesc is the grpc.ServiceDesc for SliService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SliService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.catalog.v1.SliService",
	HandlerType: (*SliServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSlis",
			Handler:    _SliService_GetSlis_Handler,
		},
		{
			MethodName: "CreateSli",
			Handler:    _SliService_CreateSli_Handler,
		},
		{
			MethodName: "CreateSlis",
			Handler:    _SliService_CreateSlis_Handler,
		},
		{
			MethodName: "UpdateSli",
			Handler:    _SliService_UpdateSli_Handler,
		},
		{
			MethodName: "UpdateSlis",
			Handler:    _SliService_UpdateSlis_Handler,
		},
		{
			MethodName: "DeleteSli",
			Handler:    _SliService_DeleteSli_Handler,
		},
		{
			MethodName: "GetSliStatusHistory",
			Handler:    _SliService_GetSliStatusHistory_Handler,
		},
		{
			MethodName: "GetSliE2mQuery",
			Handler:    _SliService_GetSliE2MQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/catalog/v1/sli_service.proto",
}
