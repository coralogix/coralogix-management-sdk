// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: com/coralogix/catalog/v1/service_catalog_sevice.proto

package v1

import (
	context "context"
	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/schemastore/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ServiceCatalogService_GetServices_FullMethodName                 = "/com.coralogix.catalog.v1.ServiceCatalogService/GetServices"
	ServiceCatalogService_GetServicesColumnsStream_FullMethodName    = "/com.coralogix.catalog.v1.ServiceCatalogService/GetServicesColumnsStream"
	ServiceCatalogService_GetService_FullMethodName                  = "/com.coralogix.catalog.v1.ServiceCatalogService/GetService"
	ServiceCatalogService_GetServiceCatalogFilters_FullMethodName    = "/com.coralogix.catalog.v1.ServiceCatalogService/GetServiceCatalogFilters"
	ServiceCatalogService_UpdateServiceCatalogFilters_FullMethodName = "/com.coralogix.catalog.v1.ServiceCatalogService/UpdateServiceCatalogFilters"
	ServiceCatalogService_GetTracingLabels_FullMethodName            = "/com.coralogix.catalog.v1.ServiceCatalogService/GetTracingLabels"
	ServiceCatalogService_GetSpanMetricBuckets_FullMethodName        = "/com.coralogix.catalog.v1.ServiceCatalogService/GetSpanMetricBuckets"
	ServiceCatalogService_GetAffectedAlerts_FullMethodName           = "/com.coralogix.catalog.v1.ServiceCatalogService/GetAffectedAlerts"
)

// ServiceCatalogServiceClient is the client API for ServiceCatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceCatalogServiceClient interface {
	GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error)
	GetServicesColumnsStream(ctx context.Context, in *GetServicesColumnsStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetServicesColumnsStreamResponse], error)
	GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error)
	GetServiceCatalogFilters(ctx context.Context, in *GetServiceCatalogFiltersRequest, opts ...grpc.CallOption) (*GetServiceCatalogFiltersResponse, error)
	UpdateServiceCatalogFilters(ctx context.Context, in *UpdateServiceCatalogFiltersRequest, opts ...grpc.CallOption) (*UpdateServiceCatalogFiltersResponse, error)
	GetTracingLabels(ctx context.Context, in *v1.TracingFieldsRequest, opts ...grpc.CallOption) (*v1.TracingFieldsResponse, error)
	GetSpanMetricBuckets(ctx context.Context, in *GetSpanMetricBucketsRequest, opts ...grpc.CallOption) (*GetSpanMetricBucketsResponse, error)
	GetAffectedAlerts(ctx context.Context, in *GetAffectedAlertsRequest, opts ...grpc.CallOption) (*GetAffectedAlertsResponse, error)
}

type serviceCatalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceCatalogServiceClient(cc grpc.ClientConnInterface) ServiceCatalogServiceClient {
	return &serviceCatalogServiceClient{cc}
}

func (c *serviceCatalogServiceClient) GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServicesResponse)
	err := c.cc.Invoke(ctx, ServiceCatalogService_GetServices_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceCatalogServiceClient) GetServicesColumnsStream(ctx context.Context, in *GetServicesColumnsStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetServicesColumnsStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ServiceCatalogService_ServiceDesc.Streams[0], ServiceCatalogService_GetServicesColumnsStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetServicesColumnsStreamRequest, GetServicesColumnsStreamResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ServiceCatalogService_GetServicesColumnsStreamClient = grpc.ServerStreamingClient[GetServicesColumnsStreamResponse]

func (c *serviceCatalogServiceClient) GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServiceResponse)
	err := c.cc.Invoke(ctx, ServiceCatalogService_GetService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceCatalogServiceClient) GetServiceCatalogFilters(ctx context.Context, in *GetServiceCatalogFiltersRequest, opts ...grpc.CallOption) (*GetServiceCatalogFiltersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServiceCatalogFiltersResponse)
	err := c.cc.Invoke(ctx, ServiceCatalogService_GetServiceCatalogFilters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceCatalogServiceClient) UpdateServiceCatalogFilters(ctx context.Context, in *UpdateServiceCatalogFiltersRequest, opts ...grpc.CallOption) (*UpdateServiceCatalogFiltersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateServiceCatalogFiltersResponse)
	err := c.cc.Invoke(ctx, ServiceCatalogService_UpdateServiceCatalogFilters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceCatalogServiceClient) GetTracingLabels(ctx context.Context, in *v1.TracingFieldsRequest, opts ...grpc.CallOption) (*v1.TracingFieldsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.TracingFieldsResponse)
	err := c.cc.Invoke(ctx, ServiceCatalogService_GetTracingLabels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceCatalogServiceClient) GetSpanMetricBuckets(ctx context.Context, in *GetSpanMetricBucketsRequest, opts ...grpc.CallOption) (*GetSpanMetricBucketsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSpanMetricBucketsResponse)
	err := c.cc.Invoke(ctx, ServiceCatalogService_GetSpanMetricBuckets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceCatalogServiceClient) GetAffectedAlerts(ctx context.Context, in *GetAffectedAlertsRequest, opts ...grpc.CallOption) (*GetAffectedAlertsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAffectedAlertsResponse)
	err := c.cc.Invoke(ctx, ServiceCatalogService_GetAffectedAlerts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceCatalogServiceServer is the server API for ServiceCatalogService service.
// All implementations must embed UnimplementedServiceCatalogServiceServer
// for forward compatibility.
type ServiceCatalogServiceServer interface {
	GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error)
	GetServicesColumnsStream(*GetServicesColumnsStreamRequest, grpc.ServerStreamingServer[GetServicesColumnsStreamResponse]) error
	GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error)
	GetServiceCatalogFilters(context.Context, *GetServiceCatalogFiltersRequest) (*GetServiceCatalogFiltersResponse, error)
	UpdateServiceCatalogFilters(context.Context, *UpdateServiceCatalogFiltersRequest) (*UpdateServiceCatalogFiltersResponse, error)
	GetTracingLabels(context.Context, *v1.TracingFieldsRequest) (*v1.TracingFieldsResponse, error)
	GetSpanMetricBuckets(context.Context, *GetSpanMetricBucketsRequest) (*GetSpanMetricBucketsResponse, error)
	GetAffectedAlerts(context.Context, *GetAffectedAlertsRequest) (*GetAffectedAlertsResponse, error)
	mustEmbedUnimplementedServiceCatalogServiceServer()
}

// UnimplementedServiceCatalogServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServiceCatalogServiceServer struct{}

func (UnimplementedServiceCatalogServiceServer) GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (UnimplementedServiceCatalogServiceServer) GetServicesColumnsStream(*GetServicesColumnsStreamRequest, grpc.ServerStreamingServer[GetServicesColumnsStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetServicesColumnsStream not implemented")
}
func (UnimplementedServiceCatalogServiceServer) GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetService not implemented")
}
func (UnimplementedServiceCatalogServiceServer) GetServiceCatalogFilters(context.Context, *GetServiceCatalogFiltersRequest) (*GetServiceCatalogFiltersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceCatalogFilters not implemented")
}
func (UnimplementedServiceCatalogServiceServer) UpdateServiceCatalogFilters(context.Context, *UpdateServiceCatalogFiltersRequest) (*UpdateServiceCatalogFiltersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServiceCatalogFilters not implemented")
}
func (UnimplementedServiceCatalogServiceServer) GetTracingLabels(context.Context, *v1.TracingFieldsRequest) (*v1.TracingFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTracingLabels not implemented")
}
func (UnimplementedServiceCatalogServiceServer) GetSpanMetricBuckets(context.Context, *GetSpanMetricBucketsRequest) (*GetSpanMetricBucketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpanMetricBuckets not implemented")
}
func (UnimplementedServiceCatalogServiceServer) GetAffectedAlerts(context.Context, *GetAffectedAlertsRequest) (*GetAffectedAlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAffectedAlerts not implemented")
}
func (UnimplementedServiceCatalogServiceServer) mustEmbedUnimplementedServiceCatalogServiceServer() {}
func (UnimplementedServiceCatalogServiceServer) testEmbeddedByValue()                               {}

// UnsafeServiceCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceCatalogServiceServer will
// result in compilation errors.
type UnsafeServiceCatalogServiceServer interface {
	mustEmbedUnimplementedServiceCatalogServiceServer()
}

func RegisterServiceCatalogServiceServer(s grpc.ServiceRegistrar, srv ServiceCatalogServiceServer) {
	// If the following call pancis, it indicates UnimplementedServiceCatalogServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ServiceCatalogService_ServiceDesc, srv)
}

func _ServiceCatalogService_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCatalogServiceServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceCatalogService_GetServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCatalogServiceServer).GetServices(ctx, req.(*GetServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceCatalogService_GetServicesColumnsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetServicesColumnsStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceCatalogServiceServer).GetServicesColumnsStream(m, &grpc.GenericServerStream[GetServicesColumnsStreamRequest, GetServicesColumnsStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ServiceCatalogService_GetServicesColumnsStreamServer = grpc.ServerStreamingServer[GetServicesColumnsStreamResponse]

func _ServiceCatalogService_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCatalogServiceServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceCatalogService_GetService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCatalogServiceServer).GetService(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceCatalogService_GetServiceCatalogFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceCatalogFiltersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCatalogServiceServer).GetServiceCatalogFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceCatalogService_GetServiceCatalogFilters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCatalogServiceServer).GetServiceCatalogFilters(ctx, req.(*GetServiceCatalogFiltersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceCatalogService_UpdateServiceCatalogFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceCatalogFiltersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCatalogServiceServer).UpdateServiceCatalogFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceCatalogService_UpdateServiceCatalogFilters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCatalogServiceServer).UpdateServiceCatalogFilters(ctx, req.(*UpdateServiceCatalogFiltersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceCatalogService_GetTracingLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.TracingFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCatalogServiceServer).GetTracingLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceCatalogService_GetTracingLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCatalogServiceServer).GetTracingLabels(ctx, req.(*v1.TracingFieldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceCatalogService_GetSpanMetricBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpanMetricBucketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCatalogServiceServer).GetSpanMetricBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceCatalogService_GetSpanMetricBuckets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCatalogServiceServer).GetSpanMetricBuckets(ctx, req.(*GetSpanMetricBucketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceCatalogService_GetAffectedAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAffectedAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceCatalogServiceServer).GetAffectedAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceCatalogService_GetAffectedAlerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceCatalogServiceServer).GetAffectedAlerts(ctx, req.(*GetAffectedAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceCatalogService_ServiceDesc is the grpc.ServiceDesc for ServiceCatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceCatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.catalog.v1.ServiceCatalogService",
	HandlerType: (*ServiceCatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServices",
			Handler:    _ServiceCatalogService_GetServices_Handler,
		},
		{
			MethodName: "GetService",
			Handler:    _ServiceCatalogService_GetService_Handler,
		},
		{
			MethodName: "GetServiceCatalogFilters",
			Handler:    _ServiceCatalogService_GetServiceCatalogFilters_Handler,
		},
		{
			MethodName: "UpdateServiceCatalogFilters",
			Handler:    _ServiceCatalogService_UpdateServiceCatalogFilters_Handler,
		},
		{
			MethodName: "GetTracingLabels",
			Handler:    _ServiceCatalogService_GetTracingLabels_Handler,
		},
		{
			MethodName: "GetSpanMetricBuckets",
			Handler:    _ServiceCatalogService_GetSpanMetricBuckets_Handler,
		},
		{
			MethodName: "GetAffectedAlerts",
			Handler:    _ServiceCatalogService_GetAffectedAlerts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetServicesColumnsStream",
			Handler:       _ServiceCatalogService_GetServicesColumnsStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "com/coralogix/catalog/v1/service_catalog_sevice.proto",
}
