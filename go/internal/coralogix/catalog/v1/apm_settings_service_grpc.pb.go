// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
// source: com/coralogixapis/service_catalog/v1/apm_settings_service.proto

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
	ApmSettingsService_GetApmSettings_FullMethodName                        = "/com.coralogixapis.service_catalog.v1.ApmSettingsService/GetApmSettings"
	ApmSettingsService_ReplaceApmSettings_FullMethodName                    = "/com.coralogixapis.service_catalog.v1.ApmSettingsService/ReplaceApmSettings"
	ApmSettingsService_ValidateApmSource_FullMethodName                     = "/com.coralogixapis.service_catalog.v1.ApmSettingsService/ValidateApmSource"
	ApmSettingsService_ReplaceRetentionPeriod_FullMethodName                = "/com.coralogixapis.service_catalog.v1.ApmSettingsService/ReplaceRetentionPeriod"
	ApmSettingsService_GetApmSource_FullMethodName                          = "/com.coralogixapis.service_catalog.v1.ApmSettingsService/GetApmSource"
	ApmSettingsService_GetSpanMetricDimensions_FullMethodName               = "/com.coralogixapis.service_catalog.v1.ApmSettingsService/GetSpanMetricDimensions"
	ApmSettingsService_GetSpanMetricDimensionsByServiceNames_FullMethodName = "/com.coralogixapis.service_catalog.v1.ApmSettingsService/GetSpanMetricDimensionsByServiceNames"
)

// ApmSettingsServiceClient is the client API for ApmSettingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApmSettingsServiceClient interface {
	GetApmSettings(ctx context.Context, in *GetApmSettingsRequest, opts ...grpc.CallOption) (*GetApmSettingsResponse, error)
	ReplaceApmSettings(ctx context.Context, in *ReplaceApmSettingsRequest, opts ...grpc.CallOption) (*ReplaceApmSettingsResponse, error)
	ValidateApmSource(ctx context.Context, in *ValidateApmSourceRequest, opts ...grpc.CallOption) (*ValidateApmSourceResponse, error)
	ReplaceRetentionPeriod(ctx context.Context, in *ReplaceRetentionPeriodRequest, opts ...grpc.CallOption) (*ReplaceRetentionPeriodResponse, error)
	GetApmSource(ctx context.Context, in *GetApmSourceRequest, opts ...grpc.CallOption) (*GetApmSourceResponse, error)
	GetSpanMetricDimensions(ctx context.Context, in *GetSpanMetricDimensionsRequest, opts ...grpc.CallOption) (*GetSpanMetricDimensionsResponse, error)
	GetSpanMetricDimensionsByServiceNames(ctx context.Context, in *GetSpanMetricDimensionsByServiceNamesRequest, opts ...grpc.CallOption) (*GetSpanMetricDimensionsByServiceNamesResponse, error)
}

type apmSettingsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApmSettingsServiceClient(cc grpc.ClientConnInterface) ApmSettingsServiceClient {
	return &apmSettingsServiceClient{cc}
}

func (c *apmSettingsServiceClient) GetApmSettings(ctx context.Context, in *GetApmSettingsRequest, opts ...grpc.CallOption) (*GetApmSettingsResponse, error) {
	out := new(GetApmSettingsResponse)
	err := c.cc.Invoke(ctx, ApmSettingsService_GetApmSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apmSettingsServiceClient) ReplaceApmSettings(ctx context.Context, in *ReplaceApmSettingsRequest, opts ...grpc.CallOption) (*ReplaceApmSettingsResponse, error) {
	out := new(ReplaceApmSettingsResponse)
	err := c.cc.Invoke(ctx, ApmSettingsService_ReplaceApmSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apmSettingsServiceClient) ValidateApmSource(ctx context.Context, in *ValidateApmSourceRequest, opts ...grpc.CallOption) (*ValidateApmSourceResponse, error) {
	out := new(ValidateApmSourceResponse)
	err := c.cc.Invoke(ctx, ApmSettingsService_ValidateApmSource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apmSettingsServiceClient) ReplaceRetentionPeriod(ctx context.Context, in *ReplaceRetentionPeriodRequest, opts ...grpc.CallOption) (*ReplaceRetentionPeriodResponse, error) {
	out := new(ReplaceRetentionPeriodResponse)
	err := c.cc.Invoke(ctx, ApmSettingsService_ReplaceRetentionPeriod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apmSettingsServiceClient) GetApmSource(ctx context.Context, in *GetApmSourceRequest, opts ...grpc.CallOption) (*GetApmSourceResponse, error) {
	out := new(GetApmSourceResponse)
	err := c.cc.Invoke(ctx, ApmSettingsService_GetApmSource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apmSettingsServiceClient) GetSpanMetricDimensions(ctx context.Context, in *GetSpanMetricDimensionsRequest, opts ...grpc.CallOption) (*GetSpanMetricDimensionsResponse, error) {
	out := new(GetSpanMetricDimensionsResponse)
	err := c.cc.Invoke(ctx, ApmSettingsService_GetSpanMetricDimensions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apmSettingsServiceClient) GetSpanMetricDimensionsByServiceNames(ctx context.Context, in *GetSpanMetricDimensionsByServiceNamesRequest, opts ...grpc.CallOption) (*GetSpanMetricDimensionsByServiceNamesResponse, error) {
	out := new(GetSpanMetricDimensionsByServiceNamesResponse)
	err := c.cc.Invoke(ctx, ApmSettingsService_GetSpanMetricDimensionsByServiceNames_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApmSettingsServiceServer is the server API for ApmSettingsService service.
// All implementations must embed UnimplementedApmSettingsServiceServer
// for forward compatibility
type ApmSettingsServiceServer interface {
	GetApmSettings(context.Context, *GetApmSettingsRequest) (*GetApmSettingsResponse, error)
	ReplaceApmSettings(context.Context, *ReplaceApmSettingsRequest) (*ReplaceApmSettingsResponse, error)
	ValidateApmSource(context.Context, *ValidateApmSourceRequest) (*ValidateApmSourceResponse, error)
	ReplaceRetentionPeriod(context.Context, *ReplaceRetentionPeriodRequest) (*ReplaceRetentionPeriodResponse, error)
	GetApmSource(context.Context, *GetApmSourceRequest) (*GetApmSourceResponse, error)
	GetSpanMetricDimensions(context.Context, *GetSpanMetricDimensionsRequest) (*GetSpanMetricDimensionsResponse, error)
	GetSpanMetricDimensionsByServiceNames(context.Context, *GetSpanMetricDimensionsByServiceNamesRequest) (*GetSpanMetricDimensionsByServiceNamesResponse, error)
	mustEmbedUnimplementedApmSettingsServiceServer()
}

// UnimplementedApmSettingsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApmSettingsServiceServer struct {
}

func (UnimplementedApmSettingsServiceServer) GetApmSettings(context.Context, *GetApmSettingsRequest) (*GetApmSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApmSettings not implemented")
}
func (UnimplementedApmSettingsServiceServer) ReplaceApmSettings(context.Context, *ReplaceApmSettingsRequest) (*ReplaceApmSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceApmSettings not implemented")
}
func (UnimplementedApmSettingsServiceServer) ValidateApmSource(context.Context, *ValidateApmSourceRequest) (*ValidateApmSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateApmSource not implemented")
}
func (UnimplementedApmSettingsServiceServer) ReplaceRetentionPeriod(context.Context, *ReplaceRetentionPeriodRequest) (*ReplaceRetentionPeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceRetentionPeriod not implemented")
}
func (UnimplementedApmSettingsServiceServer) GetApmSource(context.Context, *GetApmSourceRequest) (*GetApmSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApmSource not implemented")
}
func (UnimplementedApmSettingsServiceServer) GetSpanMetricDimensions(context.Context, *GetSpanMetricDimensionsRequest) (*GetSpanMetricDimensionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpanMetricDimensions not implemented")
}
func (UnimplementedApmSettingsServiceServer) GetSpanMetricDimensionsByServiceNames(context.Context, *GetSpanMetricDimensionsByServiceNamesRequest) (*GetSpanMetricDimensionsByServiceNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpanMetricDimensionsByServiceNames not implemented")
}
func (UnimplementedApmSettingsServiceServer) mustEmbedUnimplementedApmSettingsServiceServer() {}

// UnsafeApmSettingsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApmSettingsServiceServer will
// result in compilation errors.
type UnsafeApmSettingsServiceServer interface {
	mustEmbedUnimplementedApmSettingsServiceServer()
}

func RegisterApmSettingsServiceServer(s grpc.ServiceRegistrar, srv ApmSettingsServiceServer) {
	s.RegisterService(&ApmSettingsService_ServiceDesc, srv)
}

func _ApmSettingsService_GetApmSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApmSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApmSettingsServiceServer).GetApmSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApmSettingsService_GetApmSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApmSettingsServiceServer).GetApmSettings(ctx, req.(*GetApmSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApmSettingsService_ReplaceApmSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceApmSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApmSettingsServiceServer).ReplaceApmSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApmSettingsService_ReplaceApmSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApmSettingsServiceServer).ReplaceApmSettings(ctx, req.(*ReplaceApmSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApmSettingsService_ValidateApmSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateApmSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApmSettingsServiceServer).ValidateApmSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApmSettingsService_ValidateApmSource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApmSettingsServiceServer).ValidateApmSource(ctx, req.(*ValidateApmSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApmSettingsService_ReplaceRetentionPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceRetentionPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApmSettingsServiceServer).ReplaceRetentionPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApmSettingsService_ReplaceRetentionPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApmSettingsServiceServer).ReplaceRetentionPeriod(ctx, req.(*ReplaceRetentionPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApmSettingsService_GetApmSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApmSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApmSettingsServiceServer).GetApmSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApmSettingsService_GetApmSource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApmSettingsServiceServer).GetApmSource(ctx, req.(*GetApmSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApmSettingsService_GetSpanMetricDimensions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpanMetricDimensionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApmSettingsServiceServer).GetSpanMetricDimensions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApmSettingsService_GetSpanMetricDimensions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApmSettingsServiceServer).GetSpanMetricDimensions(ctx, req.(*GetSpanMetricDimensionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApmSettingsService_GetSpanMetricDimensionsByServiceNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpanMetricDimensionsByServiceNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApmSettingsServiceServer).GetSpanMetricDimensionsByServiceNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApmSettingsService_GetSpanMetricDimensionsByServiceNames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApmSettingsServiceServer).GetSpanMetricDimensionsByServiceNames(ctx, req.(*GetSpanMetricDimensionsByServiceNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApmSettingsService_ServiceDesc is the grpc.ServiceDesc for ApmSettingsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApmSettingsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.service_catalog.v1.ApmSettingsService",
	HandlerType: (*ApmSettingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApmSettings",
			Handler:    _ApmSettingsService_GetApmSettings_Handler,
		},
		{
			MethodName: "ReplaceApmSettings",
			Handler:    _ApmSettingsService_ReplaceApmSettings_Handler,
		},
		{
			MethodName: "ValidateApmSource",
			Handler:    _ApmSettingsService_ValidateApmSource_Handler,
		},
		{
			MethodName: "ReplaceRetentionPeriod",
			Handler:    _ApmSettingsService_ReplaceRetentionPeriod_Handler,
		},
		{
			MethodName: "GetApmSource",
			Handler:    _ApmSettingsService_GetApmSource_Handler,
		},
		{
			MethodName: "GetSpanMetricDimensions",
			Handler:    _ApmSettingsService_GetSpanMetricDimensions_Handler,
		},
		{
			MethodName: "GetSpanMetricDimensionsByServiceNames",
			Handler:    _ApmSettingsService_GetSpanMetricDimensionsByServiceNames_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/service_catalog/v1/apm_settings_service.proto",
}
