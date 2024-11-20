// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: com/coralogix/metrics/metrics-configurator.proto

package metrics

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MetricsConfiguratorPublicService_ConfigureTenant_FullMethodName = "/com.coralogix.metrics.metrics_configurator.MetricsConfiguratorPublicService/ConfigureTenant"
	MetricsConfiguratorPublicService_Update_FullMethodName          = "/com.coralogix.metrics.metrics_configurator.MetricsConfiguratorPublicService/Update"
	MetricsConfiguratorPublicService_ValidateBucket_FullMethodName  = "/com.coralogix.metrics.metrics_configurator.MetricsConfiguratorPublicService/ValidateBucket"
	MetricsConfiguratorPublicService_GetTenantConfig_FullMethodName = "/com.coralogix.metrics.metrics_configurator.MetricsConfiguratorPublicService/GetTenantConfig"
	MetricsConfiguratorPublicService_EnableArchive_FullMethodName   = "/com.coralogix.metrics.metrics_configurator.MetricsConfiguratorPublicService/EnableArchive"
	MetricsConfiguratorPublicService_DisableArchive_FullMethodName  = "/com.coralogix.metrics.metrics_configurator.MetricsConfiguratorPublicService/DisableArchive"
)

// MetricsConfiguratorPublicServiceClient is the client API for MetricsConfiguratorPublicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsConfiguratorPublicServiceClient interface {
	ConfigureTenant(ctx context.Context, in *ConfigureTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ValidateBucket(ctx context.Context, in *ValidateBucketRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTenantConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTenantConfigResponseV2, error)
	EnableArchive(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DisableArchive(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type metricsConfiguratorPublicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsConfiguratorPublicServiceClient(cc grpc.ClientConnInterface) MetricsConfiguratorPublicServiceClient {
	return &metricsConfiguratorPublicServiceClient{cc}
}

func (c *metricsConfiguratorPublicServiceClient) ConfigureTenant(ctx context.Context, in *ConfigureTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MetricsConfiguratorPublicService_ConfigureTenant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsConfiguratorPublicServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MetricsConfiguratorPublicService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsConfiguratorPublicServiceClient) ValidateBucket(ctx context.Context, in *ValidateBucketRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MetricsConfiguratorPublicService_ValidateBucket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsConfiguratorPublicServiceClient) GetTenantConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTenantConfigResponseV2, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTenantConfigResponseV2)
	err := c.cc.Invoke(ctx, MetricsConfiguratorPublicService_GetTenantConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsConfiguratorPublicServiceClient) EnableArchive(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MetricsConfiguratorPublicService_EnableArchive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsConfiguratorPublicServiceClient) DisableArchive(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MetricsConfiguratorPublicService_DisableArchive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsConfiguratorPublicServiceServer is the server API for MetricsConfiguratorPublicService service.
// All implementations must embed UnimplementedMetricsConfiguratorPublicServiceServer
// for forward compatibility.
type MetricsConfiguratorPublicServiceServer interface {
	ConfigureTenant(context.Context, *ConfigureTenantRequest) (*emptypb.Empty, error)
	Update(context.Context, *UpdateRequest) (*emptypb.Empty, error)
	ValidateBucket(context.Context, *ValidateBucketRequest) (*emptypb.Empty, error)
	GetTenantConfig(context.Context, *emptypb.Empty) (*GetTenantConfigResponseV2, error)
	EnableArchive(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	DisableArchive(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedMetricsConfiguratorPublicServiceServer()
}

// UnimplementedMetricsConfiguratorPublicServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMetricsConfiguratorPublicServiceServer struct{}

func (UnimplementedMetricsConfiguratorPublicServiceServer) ConfigureTenant(context.Context, *ConfigureTenantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureTenant not implemented")
}
func (UnimplementedMetricsConfiguratorPublicServiceServer) Update(context.Context, *UpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMetricsConfiguratorPublicServiceServer) ValidateBucket(context.Context, *ValidateBucketRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateBucket not implemented")
}
func (UnimplementedMetricsConfiguratorPublicServiceServer) GetTenantConfig(context.Context, *emptypb.Empty) (*GetTenantConfigResponseV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenantConfig not implemented")
}
func (UnimplementedMetricsConfiguratorPublicServiceServer) EnableArchive(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableArchive not implemented")
}
func (UnimplementedMetricsConfiguratorPublicServiceServer) DisableArchive(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableArchive not implemented")
}
func (UnimplementedMetricsConfiguratorPublicServiceServer) mustEmbedUnimplementedMetricsConfiguratorPublicServiceServer() {
}
func (UnimplementedMetricsConfiguratorPublicServiceServer) testEmbeddedByValue() {}

// UnsafeMetricsConfiguratorPublicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsConfiguratorPublicServiceServer will
// result in compilation errors.
type UnsafeMetricsConfiguratorPublicServiceServer interface {
	mustEmbedUnimplementedMetricsConfiguratorPublicServiceServer()
}

func RegisterMetricsConfiguratorPublicServiceServer(s grpc.ServiceRegistrar, srv MetricsConfiguratorPublicServiceServer) {
	// If the following call pancis, it indicates UnimplementedMetricsConfiguratorPublicServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MetricsConfiguratorPublicService_ServiceDesc, srv)
}

func _MetricsConfiguratorPublicService_ConfigureTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsConfiguratorPublicServiceServer).ConfigureTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsConfiguratorPublicService_ConfigureTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsConfiguratorPublicServiceServer).ConfigureTenant(ctx, req.(*ConfigureTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsConfiguratorPublicService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsConfiguratorPublicServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsConfiguratorPublicService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsConfiguratorPublicServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsConfiguratorPublicService_ValidateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsConfiguratorPublicServiceServer).ValidateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsConfiguratorPublicService_ValidateBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsConfiguratorPublicServiceServer).ValidateBucket(ctx, req.(*ValidateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsConfiguratorPublicService_GetTenantConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsConfiguratorPublicServiceServer).GetTenantConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsConfiguratorPublicService_GetTenantConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsConfiguratorPublicServiceServer).GetTenantConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsConfiguratorPublicService_EnableArchive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsConfiguratorPublicServiceServer).EnableArchive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsConfiguratorPublicService_EnableArchive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsConfiguratorPublicServiceServer).EnableArchive(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsConfiguratorPublicService_DisableArchive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsConfiguratorPublicServiceServer).DisableArchive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsConfiguratorPublicService_DisableArchive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsConfiguratorPublicServiceServer).DisableArchive(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricsConfiguratorPublicService_ServiceDesc is the grpc.ServiceDesc for MetricsConfiguratorPublicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsConfiguratorPublicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.metrics.metrics_configurator.MetricsConfiguratorPublicService",
	HandlerType: (*MetricsConfiguratorPublicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigureTenant",
			Handler:    _MetricsConfiguratorPublicService_ConfigureTenant_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MetricsConfiguratorPublicService_Update_Handler,
		},
		{
			MethodName: "ValidateBucket",
			Handler:    _MetricsConfiguratorPublicService_ValidateBucket_Handler,
		},
		{
			MethodName: "GetTenantConfig",
			Handler:    _MetricsConfiguratorPublicService_GetTenantConfig_Handler,
		},
		{
			MethodName: "EnableArchive",
			Handler:    _MetricsConfiguratorPublicService_EnableArchive_Handler,
		},
		{
			MethodName: "DisableArchive",
			Handler:    _MetricsConfiguratorPublicService_DisableArchive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/metrics/metrics-configurator.proto",
}

const (
	MetricsConfiguratorService_GetTenantConfig_FullMethodName      = "/com.coralogix.metrics.metrics_configurator.MetricsConfiguratorService/GetTenantConfig"
	MetricsConfiguratorService_ListTenantConfigs_FullMethodName    = "/com.coralogix.metrics.metrics_configurator.MetricsConfiguratorService/ListTenantConfigs"
	MetricsConfiguratorService_ListHostStoreConfigs_FullMethodName = "/com.coralogix.metrics.metrics_configurator.MetricsConfiguratorService/ListHostStoreConfigs"
	MetricsConfiguratorService_MigrateTenant_FullMethodName        = "/com.coralogix.metrics.metrics_configurator.MetricsConfiguratorService/MigrateTenant"
	MetricsConfiguratorService_Update_FullMethodName               = "/com.coralogix.metrics.metrics_configurator.MetricsConfiguratorService/Update"
)

// MetricsConfiguratorServiceClient is the client API for MetricsConfiguratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsConfiguratorServiceClient interface {
	GetTenantConfig(ctx context.Context, in *GetTenantConfigRequest, opts ...grpc.CallOption) (*GetTenantConfigResponse, error)
	ListTenantConfigs(ctx context.Context, in *ListTenantConfigsRequest, opts ...grpc.CallOption) (*ListTenantConfigsResponse, error)
	ListHostStoreConfigs(ctx context.Context, in *ListHotStoreConfigsRequest, opts ...grpc.CallOption) (*ListHotStoreConfigsResponse, error)
	MigrateTenant(ctx context.Context, in *MigrateTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *InternalUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type metricsConfiguratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsConfiguratorServiceClient(cc grpc.ClientConnInterface) MetricsConfiguratorServiceClient {
	return &metricsConfiguratorServiceClient{cc}
}

func (c *metricsConfiguratorServiceClient) GetTenantConfig(ctx context.Context, in *GetTenantConfigRequest, opts ...grpc.CallOption) (*GetTenantConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTenantConfigResponse)
	err := c.cc.Invoke(ctx, MetricsConfiguratorService_GetTenantConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsConfiguratorServiceClient) ListTenantConfigs(ctx context.Context, in *ListTenantConfigsRequest, opts ...grpc.CallOption) (*ListTenantConfigsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTenantConfigsResponse)
	err := c.cc.Invoke(ctx, MetricsConfiguratorService_ListTenantConfigs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsConfiguratorServiceClient) ListHostStoreConfigs(ctx context.Context, in *ListHotStoreConfigsRequest, opts ...grpc.CallOption) (*ListHotStoreConfigsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListHotStoreConfigsResponse)
	err := c.cc.Invoke(ctx, MetricsConfiguratorService_ListHostStoreConfigs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsConfiguratorServiceClient) MigrateTenant(ctx context.Context, in *MigrateTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MetricsConfiguratorService_MigrateTenant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsConfiguratorServiceClient) Update(ctx context.Context, in *InternalUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MetricsConfiguratorService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsConfiguratorServiceServer is the server API for MetricsConfiguratorService service.
// All implementations must embed UnimplementedMetricsConfiguratorServiceServer
// for forward compatibility.
type MetricsConfiguratorServiceServer interface {
	GetTenantConfig(context.Context, *GetTenantConfigRequest) (*GetTenantConfigResponse, error)
	ListTenantConfigs(context.Context, *ListTenantConfigsRequest) (*ListTenantConfigsResponse, error)
	ListHostStoreConfigs(context.Context, *ListHotStoreConfigsRequest) (*ListHotStoreConfigsResponse, error)
	MigrateTenant(context.Context, *MigrateTenantRequest) (*emptypb.Empty, error)
	Update(context.Context, *InternalUpdateRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMetricsConfiguratorServiceServer()
}

// UnimplementedMetricsConfiguratorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMetricsConfiguratorServiceServer struct{}

func (UnimplementedMetricsConfiguratorServiceServer) GetTenantConfig(context.Context, *GetTenantConfigRequest) (*GetTenantConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenantConfig not implemented")
}
func (UnimplementedMetricsConfiguratorServiceServer) ListTenantConfigs(context.Context, *ListTenantConfigsRequest) (*ListTenantConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTenantConfigs not implemented")
}
func (UnimplementedMetricsConfiguratorServiceServer) ListHostStoreConfigs(context.Context, *ListHotStoreConfigsRequest) (*ListHotStoreConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHostStoreConfigs not implemented")
}
func (UnimplementedMetricsConfiguratorServiceServer) MigrateTenant(context.Context, *MigrateTenantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MigrateTenant not implemented")
}
func (UnimplementedMetricsConfiguratorServiceServer) Update(context.Context, *InternalUpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMetricsConfiguratorServiceServer) mustEmbedUnimplementedMetricsConfiguratorServiceServer() {
}
func (UnimplementedMetricsConfiguratorServiceServer) testEmbeddedByValue() {}

// UnsafeMetricsConfiguratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsConfiguratorServiceServer will
// result in compilation errors.
type UnsafeMetricsConfiguratorServiceServer interface {
	mustEmbedUnimplementedMetricsConfiguratorServiceServer()
}

func RegisterMetricsConfiguratorServiceServer(s grpc.ServiceRegistrar, srv MetricsConfiguratorServiceServer) {
	// If the following call pancis, it indicates UnimplementedMetricsConfiguratorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MetricsConfiguratorService_ServiceDesc, srv)
}

func _MetricsConfiguratorService_GetTenantConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsConfiguratorServiceServer).GetTenantConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsConfiguratorService_GetTenantConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsConfiguratorServiceServer).GetTenantConfig(ctx, req.(*GetTenantConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsConfiguratorService_ListTenantConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsConfiguratorServiceServer).ListTenantConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsConfiguratorService_ListTenantConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsConfiguratorServiceServer).ListTenantConfigs(ctx, req.(*ListTenantConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsConfiguratorService_ListHostStoreConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHotStoreConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsConfiguratorServiceServer).ListHostStoreConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsConfiguratorService_ListHostStoreConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsConfiguratorServiceServer).ListHostStoreConfigs(ctx, req.(*ListHotStoreConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsConfiguratorService_MigrateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MigrateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsConfiguratorServiceServer).MigrateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsConfiguratorService_MigrateTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsConfiguratorServiceServer).MigrateTenant(ctx, req.(*MigrateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsConfiguratorService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsConfiguratorServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsConfiguratorService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsConfiguratorServiceServer).Update(ctx, req.(*InternalUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricsConfiguratorService_ServiceDesc is the grpc.ServiceDesc for MetricsConfiguratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsConfiguratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.metrics.metrics_configurator.MetricsConfiguratorService",
	HandlerType: (*MetricsConfiguratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTenantConfig",
			Handler:    _MetricsConfiguratorService_GetTenantConfig_Handler,
		},
		{
			MethodName: "ListTenantConfigs",
			Handler:    _MetricsConfiguratorService_ListTenantConfigs_Handler,
		},
		{
			MethodName: "ListHostStoreConfigs",
			Handler:    _MetricsConfiguratorService_ListHostStoreConfigs_Handler,
		},
		{
			MethodName: "MigrateTenant",
			Handler:    _MetricsConfiguratorService_MigrateTenant_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MetricsConfiguratorService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/metrics/metrics-configurator.proto",
}
