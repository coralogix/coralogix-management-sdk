// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: com/coralogix/global_mapping/v1/label_mappings_service.proto

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
	LabelMappingsService_GetLabelMappings_FullMethodName          = "/com.coralogix.global_mapping.v1.LabelMappingsService/GetLabelMappings"
	LabelMappingsService_RemoveLabelMappings_FullMethodName       = "/com.coralogix.global_mapping.v1.LabelMappingsService/RemoveLabelMappings"
	LabelMappingsService_UpsertLabelMappings_FullMethodName       = "/com.coralogix.global_mapping.v1.LabelMappingsService/UpsertLabelMappings"
	LabelMappingsService_GetLabels_FullMethodName                 = "/com.coralogix.global_mapping.v1.LabelMappingsService/GetLabels"
	LabelMappingsService_Extract_FullMethodName                   = "/com.coralogix.global_mapping.v1.LabelMappingsService/Extract"
	LabelMappingsService_GetLabelValues_FullMethodName            = "/com.coralogix.global_mapping.v1.LabelMappingsService/GetLabelValues"
	LabelMappingsService_GetCustomLabelMappings_FullMethodName    = "/com.coralogix.global_mapping.v1.LabelMappingsService/GetCustomLabelMappings"
	LabelMappingsService_UpsertCustomLabelMappings_FullMethodName = "/com.coralogix.global_mapping.v1.LabelMappingsService/UpsertCustomLabelMappings"
	LabelMappingsService_RemoveCustomLabelMappings_FullMethodName = "/com.coralogix.global_mapping.v1.LabelMappingsService/RemoveCustomLabelMappings"
	LabelMappingsService_GetLabelKeys_FullMethodName              = "/com.coralogix.global_mapping.v1.LabelMappingsService/GetLabelKeys"
)

// LabelMappingsServiceClient is the client API for LabelMappingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LabelMappingsServiceClient interface {
	GetLabelMappings(ctx context.Context, in *GetLabelMappingsRequest, opts ...grpc.CallOption) (*GetLabelMappingsResponse, error)
	RemoveLabelMappings(ctx context.Context, in *RemoveLabelMappingsRequest, opts ...grpc.CallOption) (*RemoveLabelMappingsResponse, error)
	UpsertLabelMappings(ctx context.Context, in *UpsertLabelMappingsRequest, opts ...grpc.CallOption) (*UpsertLabelMappingsResponse, error)
	GetLabels(ctx context.Context, in *GetLabelsRequest, opts ...grpc.CallOption) (*GetLabelsResponse, error)
	Extract(ctx context.Context, in *ExtractRequest, opts ...grpc.CallOption) (*ExtractResponse, error)
	GetLabelValues(ctx context.Context, in *GetLabelValuesRequest, opts ...grpc.CallOption) (*GetLabelValuesResponse, error)
	GetCustomLabelMappings(ctx context.Context, in *GetCustomLabelMappingsRequest, opts ...grpc.CallOption) (*GetCustomLabelMappingsResponse, error)
	UpsertCustomLabelMappings(ctx context.Context, in *UpsertCustomLabelMappingsRequest, opts ...grpc.CallOption) (*UpsertCustomLabelMappingsResponse, error)
	RemoveCustomLabelMappings(ctx context.Context, in *RemoveCustomLabelMappingsRequest, opts ...grpc.CallOption) (*RemoveCustomLabelMappingsResponse, error)
	GetLabelKeys(ctx context.Context, in *GetLabelKeysRequest, opts ...grpc.CallOption) (*GetLabelKeysResponse, error)
}

type labelMappingsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLabelMappingsServiceClient(cc grpc.ClientConnInterface) LabelMappingsServiceClient {
	return &labelMappingsServiceClient{cc}
}

func (c *labelMappingsServiceClient) GetLabelMappings(ctx context.Context, in *GetLabelMappingsRequest, opts ...grpc.CallOption) (*GetLabelMappingsResponse, error) {
	out := new(GetLabelMappingsResponse)
	err := c.cc.Invoke(ctx, LabelMappingsService_GetLabelMappings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelMappingsServiceClient) RemoveLabelMappings(ctx context.Context, in *RemoveLabelMappingsRequest, opts ...grpc.CallOption) (*RemoveLabelMappingsResponse, error) {
	out := new(RemoveLabelMappingsResponse)
	err := c.cc.Invoke(ctx, LabelMappingsService_RemoveLabelMappings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelMappingsServiceClient) UpsertLabelMappings(ctx context.Context, in *UpsertLabelMappingsRequest, opts ...grpc.CallOption) (*UpsertLabelMappingsResponse, error) {
	out := new(UpsertLabelMappingsResponse)
	err := c.cc.Invoke(ctx, LabelMappingsService_UpsertLabelMappings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelMappingsServiceClient) GetLabels(ctx context.Context, in *GetLabelsRequest, opts ...grpc.CallOption) (*GetLabelsResponse, error) {
	out := new(GetLabelsResponse)
	err := c.cc.Invoke(ctx, LabelMappingsService_GetLabels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelMappingsServiceClient) Extract(ctx context.Context, in *ExtractRequest, opts ...grpc.CallOption) (*ExtractResponse, error) {
	out := new(ExtractResponse)
	err := c.cc.Invoke(ctx, LabelMappingsService_Extract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelMappingsServiceClient) GetLabelValues(ctx context.Context, in *GetLabelValuesRequest, opts ...grpc.CallOption) (*GetLabelValuesResponse, error) {
	out := new(GetLabelValuesResponse)
	err := c.cc.Invoke(ctx, LabelMappingsService_GetLabelValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelMappingsServiceClient) GetCustomLabelMappings(ctx context.Context, in *GetCustomLabelMappingsRequest, opts ...grpc.CallOption) (*GetCustomLabelMappingsResponse, error) {
	out := new(GetCustomLabelMappingsResponse)
	err := c.cc.Invoke(ctx, LabelMappingsService_GetCustomLabelMappings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelMappingsServiceClient) UpsertCustomLabelMappings(ctx context.Context, in *UpsertCustomLabelMappingsRequest, opts ...grpc.CallOption) (*UpsertCustomLabelMappingsResponse, error) {
	out := new(UpsertCustomLabelMappingsResponse)
	err := c.cc.Invoke(ctx, LabelMappingsService_UpsertCustomLabelMappings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelMappingsServiceClient) RemoveCustomLabelMappings(ctx context.Context, in *RemoveCustomLabelMappingsRequest, opts ...grpc.CallOption) (*RemoveCustomLabelMappingsResponse, error) {
	out := new(RemoveCustomLabelMappingsResponse)
	err := c.cc.Invoke(ctx, LabelMappingsService_RemoveCustomLabelMappings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelMappingsServiceClient) GetLabelKeys(ctx context.Context, in *GetLabelKeysRequest, opts ...grpc.CallOption) (*GetLabelKeysResponse, error) {
	out := new(GetLabelKeysResponse)
	err := c.cc.Invoke(ctx, LabelMappingsService_GetLabelKeys_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LabelMappingsServiceServer is the server API for LabelMappingsService service.
// All implementations must embed UnimplementedLabelMappingsServiceServer
// for forward compatibility
type LabelMappingsServiceServer interface {
	GetLabelMappings(context.Context, *GetLabelMappingsRequest) (*GetLabelMappingsResponse, error)
	RemoveLabelMappings(context.Context, *RemoveLabelMappingsRequest) (*RemoveLabelMappingsResponse, error)
	UpsertLabelMappings(context.Context, *UpsertLabelMappingsRequest) (*UpsertLabelMappingsResponse, error)
	GetLabels(context.Context, *GetLabelsRequest) (*GetLabelsResponse, error)
	Extract(context.Context, *ExtractRequest) (*ExtractResponse, error)
	GetLabelValues(context.Context, *GetLabelValuesRequest) (*GetLabelValuesResponse, error)
	GetCustomLabelMappings(context.Context, *GetCustomLabelMappingsRequest) (*GetCustomLabelMappingsResponse, error)
	UpsertCustomLabelMappings(context.Context, *UpsertCustomLabelMappingsRequest) (*UpsertCustomLabelMappingsResponse, error)
	RemoveCustomLabelMappings(context.Context, *RemoveCustomLabelMappingsRequest) (*RemoveCustomLabelMappingsResponse, error)
	GetLabelKeys(context.Context, *GetLabelKeysRequest) (*GetLabelKeysResponse, error)
	mustEmbedUnimplementedLabelMappingsServiceServer()
}

// UnimplementedLabelMappingsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLabelMappingsServiceServer struct {
}

func (UnimplementedLabelMappingsServiceServer) GetLabelMappings(context.Context, *GetLabelMappingsRequest) (*GetLabelMappingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabelMappings not implemented")
}
func (UnimplementedLabelMappingsServiceServer) RemoveLabelMappings(context.Context, *RemoveLabelMappingsRequest) (*RemoveLabelMappingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveLabelMappings not implemented")
}
func (UnimplementedLabelMappingsServiceServer) UpsertLabelMappings(context.Context, *UpsertLabelMappingsRequest) (*UpsertLabelMappingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertLabelMappings not implemented")
}
func (UnimplementedLabelMappingsServiceServer) GetLabels(context.Context, *GetLabelsRequest) (*GetLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabels not implemented")
}
func (UnimplementedLabelMappingsServiceServer) Extract(context.Context, *ExtractRequest) (*ExtractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Extract not implemented")
}
func (UnimplementedLabelMappingsServiceServer) GetLabelValues(context.Context, *GetLabelValuesRequest) (*GetLabelValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabelValues not implemented")
}
func (UnimplementedLabelMappingsServiceServer) GetCustomLabelMappings(context.Context, *GetCustomLabelMappingsRequest) (*GetCustomLabelMappingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomLabelMappings not implemented")
}
func (UnimplementedLabelMappingsServiceServer) UpsertCustomLabelMappings(context.Context, *UpsertCustomLabelMappingsRequest) (*UpsertCustomLabelMappingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertCustomLabelMappings not implemented")
}
func (UnimplementedLabelMappingsServiceServer) RemoveCustomLabelMappings(context.Context, *RemoveCustomLabelMappingsRequest) (*RemoveCustomLabelMappingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCustomLabelMappings not implemented")
}
func (UnimplementedLabelMappingsServiceServer) GetLabelKeys(context.Context, *GetLabelKeysRequest) (*GetLabelKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabelKeys not implemented")
}
func (UnimplementedLabelMappingsServiceServer) mustEmbedUnimplementedLabelMappingsServiceServer() {}

// UnsafeLabelMappingsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LabelMappingsServiceServer will
// result in compilation errors.
type UnsafeLabelMappingsServiceServer interface {
	mustEmbedUnimplementedLabelMappingsServiceServer()
}

func RegisterLabelMappingsServiceServer(s grpc.ServiceRegistrar, srv LabelMappingsServiceServer) {
	s.RegisterService(&LabelMappingsService_ServiceDesc, srv)
}

func _LabelMappingsService_GetLabelMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelMappingsServiceServer).GetLabelMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelMappingsService_GetLabelMappings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelMappingsServiceServer).GetLabelMappings(ctx, req.(*GetLabelMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelMappingsService_RemoveLabelMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveLabelMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelMappingsServiceServer).RemoveLabelMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelMappingsService_RemoveLabelMappings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelMappingsServiceServer).RemoveLabelMappings(ctx, req.(*RemoveLabelMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelMappingsService_UpsertLabelMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertLabelMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelMappingsServiceServer).UpsertLabelMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelMappingsService_UpsertLabelMappings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelMappingsServiceServer).UpsertLabelMappings(ctx, req.(*UpsertLabelMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelMappingsService_GetLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelMappingsServiceServer).GetLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelMappingsService_GetLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelMappingsServiceServer).GetLabels(ctx, req.(*GetLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelMappingsService_Extract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelMappingsServiceServer).Extract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelMappingsService_Extract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelMappingsServiceServer).Extract(ctx, req.(*ExtractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelMappingsService_GetLabelValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelMappingsServiceServer).GetLabelValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelMappingsService_GetLabelValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelMappingsServiceServer).GetLabelValues(ctx, req.(*GetLabelValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelMappingsService_GetCustomLabelMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomLabelMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelMappingsServiceServer).GetCustomLabelMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelMappingsService_GetCustomLabelMappings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelMappingsServiceServer).GetCustomLabelMappings(ctx, req.(*GetCustomLabelMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelMappingsService_UpsertCustomLabelMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertCustomLabelMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelMappingsServiceServer).UpsertCustomLabelMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelMappingsService_UpsertCustomLabelMappings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelMappingsServiceServer).UpsertCustomLabelMappings(ctx, req.(*UpsertCustomLabelMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelMappingsService_RemoveCustomLabelMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCustomLabelMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelMappingsServiceServer).RemoveCustomLabelMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelMappingsService_RemoveCustomLabelMappings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelMappingsServiceServer).RemoveCustomLabelMappings(ctx, req.(*RemoveCustomLabelMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelMappingsService_GetLabelKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelMappingsServiceServer).GetLabelKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelMappingsService_GetLabelKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelMappingsServiceServer).GetLabelKeys(ctx, req.(*GetLabelKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LabelMappingsService_ServiceDesc is the grpc.ServiceDesc for LabelMappingsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LabelMappingsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.global_mapping.v1.LabelMappingsService",
	HandlerType: (*LabelMappingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLabelMappings",
			Handler:    _LabelMappingsService_GetLabelMappings_Handler,
		},
		{
			MethodName: "RemoveLabelMappings",
			Handler:    _LabelMappingsService_RemoveLabelMappings_Handler,
		},
		{
			MethodName: "UpsertLabelMappings",
			Handler:    _LabelMappingsService_UpsertLabelMappings_Handler,
		},
		{
			MethodName: "GetLabels",
			Handler:    _LabelMappingsService_GetLabels_Handler,
		},
		{
			MethodName: "Extract",
			Handler:    _LabelMappingsService_Extract_Handler,
		},
		{
			MethodName: "GetLabelValues",
			Handler:    _LabelMappingsService_GetLabelValues_Handler,
		},
		{
			MethodName: "GetCustomLabelMappings",
			Handler:    _LabelMappingsService_GetCustomLabelMappings_Handler,
		},
		{
			MethodName: "UpsertCustomLabelMappings",
			Handler:    _LabelMappingsService_UpsertCustomLabelMappings_Handler,
		},
		{
			MethodName: "RemoveCustomLabelMappings",
			Handler:    _LabelMappingsService_RemoveCustomLabelMappings_Handler,
		},
		{
			MethodName: "GetLabelKeys",
			Handler:    _LabelMappingsService_GetLabelKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/global_mapping/v1/label_mappings_service.proto",
}
