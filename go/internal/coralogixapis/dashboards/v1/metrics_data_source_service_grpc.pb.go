// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
// source: com/coralogixapis/dashboards/v1/services/metrics_data_source_service.proto

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
	MetricsDataSourceService_SearchMetricsTimeSeries_FullMethodName        = "/com.coralogixapis.dashboards.v1.services.MetricsDataSourceService/SearchMetricsTimeSeries"
	MetricsDataSourceService_SearchMetricsTimeValues_FullMethodName        = "/com.coralogixapis.dashboards.v1.services.MetricsDataSourceService/SearchMetricsTimeValues"
	MetricsDataSourceService_SearchMetricsGroupedSeries_FullMethodName     = "/com.coralogixapis.dashboards.v1.services.MetricsDataSourceService/SearchMetricsGroupedSeries"
	MetricsDataSourceService_SearchMetricsGroupedTimeSeries_FullMethodName = "/com.coralogixapis.dashboards.v1.services.MetricsDataSourceService/SearchMetricsGroupedTimeSeries"
	MetricsDataSourceService_SearchMetricsEvents_FullMethodName            = "/com.coralogixapis.dashboards.v1.services.MetricsDataSourceService/SearchMetricsEvents"
	MetricsDataSourceService_SearchMetricsAnnotationEvents_FullMethodName  = "/com.coralogixapis.dashboards.v1.services.MetricsDataSourceService/SearchMetricsAnnotationEvents"
	MetricsDataSourceService_SearchMetricsGroupedValues_FullMethodName     = "/com.coralogixapis.dashboards.v1.services.MetricsDataSourceService/SearchMetricsGroupedValues"
)

// MetricsDataSourceServiceClient is the client API for MetricsDataSourceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsDataSourceServiceClient interface {
	SearchMetricsTimeSeries(ctx context.Context, in *SearchMetricsTimeSeriesRequest, opts ...grpc.CallOption) (*SearchMetricsTimeSeriesResponse, error)
	SearchMetricsTimeValues(ctx context.Context, in *SearchMetricsTimeValuesRequest, opts ...grpc.CallOption) (*SearchMetricsTimeValuesResponse, error)
	SearchMetricsGroupedSeries(ctx context.Context, in *SearchMetricsGroupedSeriesRequest, opts ...grpc.CallOption) (*SearchMetricsGroupedSeriesResponse, error)
	SearchMetricsGroupedTimeSeries(ctx context.Context, in *SearchMetricsGroupedTimeSeriesRequest, opts ...grpc.CallOption) (*SearchMetricsGroupedTimeSeriesResponse, error)
	SearchMetricsEvents(ctx context.Context, in *SearchMetricsEventsRequest, opts ...grpc.CallOption) (*SearchMetricsEventsResponse, error)
	SearchMetricsAnnotationEvents(ctx context.Context, in *SearchMetricsAnnotationEventsRequest, opts ...grpc.CallOption) (*SearchMetricsAnnotationEventsResponse, error)
	SearchMetricsGroupedValues(ctx context.Context, in *SearchMetricsGroupedValuesRequest, opts ...grpc.CallOption) (*SearchMetricsGroupedValuesResponse, error)
}

type metricsDataSourceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsDataSourceServiceClient(cc grpc.ClientConnInterface) MetricsDataSourceServiceClient {
	return &metricsDataSourceServiceClient{cc}
}

func (c *metricsDataSourceServiceClient) SearchMetricsTimeSeries(ctx context.Context, in *SearchMetricsTimeSeriesRequest, opts ...grpc.CallOption) (*SearchMetricsTimeSeriesResponse, error) {
	out := new(SearchMetricsTimeSeriesResponse)
	err := c.cc.Invoke(ctx, MetricsDataSourceService_SearchMetricsTimeSeries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsDataSourceServiceClient) SearchMetricsTimeValues(ctx context.Context, in *SearchMetricsTimeValuesRequest, opts ...grpc.CallOption) (*SearchMetricsTimeValuesResponse, error) {
	out := new(SearchMetricsTimeValuesResponse)
	err := c.cc.Invoke(ctx, MetricsDataSourceService_SearchMetricsTimeValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsDataSourceServiceClient) SearchMetricsGroupedSeries(ctx context.Context, in *SearchMetricsGroupedSeriesRequest, opts ...grpc.CallOption) (*SearchMetricsGroupedSeriesResponse, error) {
	out := new(SearchMetricsGroupedSeriesResponse)
	err := c.cc.Invoke(ctx, MetricsDataSourceService_SearchMetricsGroupedSeries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsDataSourceServiceClient) SearchMetricsGroupedTimeSeries(ctx context.Context, in *SearchMetricsGroupedTimeSeriesRequest, opts ...grpc.CallOption) (*SearchMetricsGroupedTimeSeriesResponse, error) {
	out := new(SearchMetricsGroupedTimeSeriesResponse)
	err := c.cc.Invoke(ctx, MetricsDataSourceService_SearchMetricsGroupedTimeSeries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsDataSourceServiceClient) SearchMetricsEvents(ctx context.Context, in *SearchMetricsEventsRequest, opts ...grpc.CallOption) (*SearchMetricsEventsResponse, error) {
	out := new(SearchMetricsEventsResponse)
	err := c.cc.Invoke(ctx, MetricsDataSourceService_SearchMetricsEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsDataSourceServiceClient) SearchMetricsAnnotationEvents(ctx context.Context, in *SearchMetricsAnnotationEventsRequest, opts ...grpc.CallOption) (*SearchMetricsAnnotationEventsResponse, error) {
	out := new(SearchMetricsAnnotationEventsResponse)
	err := c.cc.Invoke(ctx, MetricsDataSourceService_SearchMetricsAnnotationEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsDataSourceServiceClient) SearchMetricsGroupedValues(ctx context.Context, in *SearchMetricsGroupedValuesRequest, opts ...grpc.CallOption) (*SearchMetricsGroupedValuesResponse, error) {
	out := new(SearchMetricsGroupedValuesResponse)
	err := c.cc.Invoke(ctx, MetricsDataSourceService_SearchMetricsGroupedValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsDataSourceServiceServer is the server API for MetricsDataSourceService service.
// All implementations must embed UnimplementedMetricsDataSourceServiceServer
// for forward compatibility
type MetricsDataSourceServiceServer interface {
	SearchMetricsTimeSeries(context.Context, *SearchMetricsTimeSeriesRequest) (*SearchMetricsTimeSeriesResponse, error)
	SearchMetricsTimeValues(context.Context, *SearchMetricsTimeValuesRequest) (*SearchMetricsTimeValuesResponse, error)
	SearchMetricsGroupedSeries(context.Context, *SearchMetricsGroupedSeriesRequest) (*SearchMetricsGroupedSeriesResponse, error)
	SearchMetricsGroupedTimeSeries(context.Context, *SearchMetricsGroupedTimeSeriesRequest) (*SearchMetricsGroupedTimeSeriesResponse, error)
	SearchMetricsEvents(context.Context, *SearchMetricsEventsRequest) (*SearchMetricsEventsResponse, error)
	SearchMetricsAnnotationEvents(context.Context, *SearchMetricsAnnotationEventsRequest) (*SearchMetricsAnnotationEventsResponse, error)
	SearchMetricsGroupedValues(context.Context, *SearchMetricsGroupedValuesRequest) (*SearchMetricsGroupedValuesResponse, error)
	mustEmbedUnimplementedMetricsDataSourceServiceServer()
}

// UnimplementedMetricsDataSourceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsDataSourceServiceServer struct {
}

func (UnimplementedMetricsDataSourceServiceServer) SearchMetricsTimeSeries(context.Context, *SearchMetricsTimeSeriesRequest) (*SearchMetricsTimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMetricsTimeSeries not implemented")
}
func (UnimplementedMetricsDataSourceServiceServer) SearchMetricsTimeValues(context.Context, *SearchMetricsTimeValuesRequest) (*SearchMetricsTimeValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMetricsTimeValues not implemented")
}
func (UnimplementedMetricsDataSourceServiceServer) SearchMetricsGroupedSeries(context.Context, *SearchMetricsGroupedSeriesRequest) (*SearchMetricsGroupedSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMetricsGroupedSeries not implemented")
}
func (UnimplementedMetricsDataSourceServiceServer) SearchMetricsGroupedTimeSeries(context.Context, *SearchMetricsGroupedTimeSeriesRequest) (*SearchMetricsGroupedTimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMetricsGroupedTimeSeries not implemented")
}
func (UnimplementedMetricsDataSourceServiceServer) SearchMetricsEvents(context.Context, *SearchMetricsEventsRequest) (*SearchMetricsEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMetricsEvents not implemented")
}
func (UnimplementedMetricsDataSourceServiceServer) SearchMetricsAnnotationEvents(context.Context, *SearchMetricsAnnotationEventsRequest) (*SearchMetricsAnnotationEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMetricsAnnotationEvents not implemented")
}
func (UnimplementedMetricsDataSourceServiceServer) SearchMetricsGroupedValues(context.Context, *SearchMetricsGroupedValuesRequest) (*SearchMetricsGroupedValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMetricsGroupedValues not implemented")
}
func (UnimplementedMetricsDataSourceServiceServer) mustEmbedUnimplementedMetricsDataSourceServiceServer() {
}

// UnsafeMetricsDataSourceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsDataSourceServiceServer will
// result in compilation errors.
type UnsafeMetricsDataSourceServiceServer interface {
	mustEmbedUnimplementedMetricsDataSourceServiceServer()
}

func RegisterMetricsDataSourceServiceServer(s grpc.ServiceRegistrar, srv MetricsDataSourceServiceServer) {
	s.RegisterService(&MetricsDataSourceService_ServiceDesc, srv)
}

func _MetricsDataSourceService_SearchMetricsTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMetricsTimeSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsDataSourceServiceServer).SearchMetricsTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsDataSourceService_SearchMetricsTimeSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsDataSourceServiceServer).SearchMetricsTimeSeries(ctx, req.(*SearchMetricsTimeSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsDataSourceService_SearchMetricsTimeValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMetricsTimeValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsDataSourceServiceServer).SearchMetricsTimeValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsDataSourceService_SearchMetricsTimeValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsDataSourceServiceServer).SearchMetricsTimeValues(ctx, req.(*SearchMetricsTimeValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsDataSourceService_SearchMetricsGroupedSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMetricsGroupedSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsDataSourceServiceServer).SearchMetricsGroupedSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsDataSourceService_SearchMetricsGroupedSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsDataSourceServiceServer).SearchMetricsGroupedSeries(ctx, req.(*SearchMetricsGroupedSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsDataSourceService_SearchMetricsGroupedTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMetricsGroupedTimeSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsDataSourceServiceServer).SearchMetricsGroupedTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsDataSourceService_SearchMetricsGroupedTimeSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsDataSourceServiceServer).SearchMetricsGroupedTimeSeries(ctx, req.(*SearchMetricsGroupedTimeSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsDataSourceService_SearchMetricsEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMetricsEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsDataSourceServiceServer).SearchMetricsEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsDataSourceService_SearchMetricsEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsDataSourceServiceServer).SearchMetricsEvents(ctx, req.(*SearchMetricsEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsDataSourceService_SearchMetricsAnnotationEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMetricsAnnotationEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsDataSourceServiceServer).SearchMetricsAnnotationEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsDataSourceService_SearchMetricsAnnotationEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsDataSourceServiceServer).SearchMetricsAnnotationEvents(ctx, req.(*SearchMetricsAnnotationEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsDataSourceService_SearchMetricsGroupedValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMetricsGroupedValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsDataSourceServiceServer).SearchMetricsGroupedValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsDataSourceService_SearchMetricsGroupedValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsDataSourceServiceServer).SearchMetricsGroupedValues(ctx, req.(*SearchMetricsGroupedValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricsDataSourceService_ServiceDesc is the grpc.ServiceDesc for MetricsDataSourceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsDataSourceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.dashboards.v1.services.MetricsDataSourceService",
	HandlerType: (*MetricsDataSourceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchMetricsTimeSeries",
			Handler:    _MetricsDataSourceService_SearchMetricsTimeSeries_Handler,
		},
		{
			MethodName: "SearchMetricsTimeValues",
			Handler:    _MetricsDataSourceService_SearchMetricsTimeValues_Handler,
		},
		{
			MethodName: "SearchMetricsGroupedSeries",
			Handler:    _MetricsDataSourceService_SearchMetricsGroupedSeries_Handler,
		},
		{
			MethodName: "SearchMetricsGroupedTimeSeries",
			Handler:    _MetricsDataSourceService_SearchMetricsGroupedTimeSeries_Handler,
		},
		{
			MethodName: "SearchMetricsEvents",
			Handler:    _MetricsDataSourceService_SearchMetricsEvents_Handler,
		},
		{
			MethodName: "SearchMetricsAnnotationEvents",
			Handler:    _MetricsDataSourceService_SearchMetricsAnnotationEvents_Handler,
		},
		{
			MethodName: "SearchMetricsGroupedValues",
			Handler:    _MetricsDataSourceService_SearchMetricsGroupedValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/dashboards/v1/services/metrics_data_source_service.proto",
}
