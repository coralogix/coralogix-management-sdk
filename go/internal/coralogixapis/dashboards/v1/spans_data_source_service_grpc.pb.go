// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: com/coralogixapis/dashboards/v1/services/spans_data_source_service.proto

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
	SpansDataSourceService_SearchSpansTimeSeries_FullMethodName        = "/com.coralogixapis.dashboards.v1.services.SpansDataSourceService/SearchSpansTimeSeries"
	SpansDataSourceService_SearchSpansEvents_FullMethodName            = "/com.coralogixapis.dashboards.v1.services.SpansDataSourceService/SearchSpansEvents"
	SpansDataSourceService_SearchSpansEventGroups_FullMethodName       = "/com.coralogixapis.dashboards.v1.services.SpansDataSourceService/SearchSpansEventGroups"
	SpansDataSourceService_SearchGroupedSpansSeries_FullMethodName     = "/com.coralogixapis.dashboards.v1.services.SpansDataSourceService/SearchGroupedSpansSeries"
	SpansDataSourceService_SearchSpansGroupedTimeSeries_FullMethodName = "/com.coralogixapis.dashboards.v1.services.SpansDataSourceService/SearchSpansGroupedTimeSeries"
	SpansDataSourceService_SearchSpansTimeValue_FullMethodName         = "/com.coralogixapis.dashboards.v1.services.SpansDataSourceService/SearchSpansTimeValue"
	SpansDataSourceService_SearchSpansAnnotationEvents_FullMethodName  = "/com.coralogixapis.dashboards.v1.services.SpansDataSourceService/SearchSpansAnnotationEvents"
)

// SpansDataSourceServiceClient is the client API for SpansDataSourceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpansDataSourceServiceClient interface {
	SearchSpansTimeSeries(ctx context.Context, in *SearchSpansTimeSeriesRequest, opts ...grpc.CallOption) (*SearchSpansTimeSeriesResponse, error)
	SearchSpansEvents(ctx context.Context, in *SearchSpansEventsRequest, opts ...grpc.CallOption) (*SearchSpansEventsResponse, error)
	SearchSpansEventGroups(ctx context.Context, in *SearchSpansEventGroupsRequest, opts ...grpc.CallOption) (*SearchSpansEventGroupsResponse, error)
	SearchGroupedSpansSeries(ctx context.Context, in *SearchGroupedSpansSeriesRequest, opts ...grpc.CallOption) (*SearchGroupedSpansSeriesResponse, error)
	SearchSpansGroupedTimeSeries(ctx context.Context, in *SearchSpansGroupedTimeSeriesRequest, opts ...grpc.CallOption) (*SearchSpansGroupedTimeSeriesResponse, error)
	SearchSpansTimeValue(ctx context.Context, in *SearchSpansTimeValueRequest, opts ...grpc.CallOption) (*SearchSpansTimeValueResponse, error)
	SearchSpansAnnotationEvents(ctx context.Context, in *SearchSpansAnnotationEventsRequest, opts ...grpc.CallOption) (*SearchSpansAnnotationEventsResponse, error)
}

type spansDataSourceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpansDataSourceServiceClient(cc grpc.ClientConnInterface) SpansDataSourceServiceClient {
	return &spansDataSourceServiceClient{cc}
}

func (c *spansDataSourceServiceClient) SearchSpansTimeSeries(ctx context.Context, in *SearchSpansTimeSeriesRequest, opts ...grpc.CallOption) (*SearchSpansTimeSeriesResponse, error) {
	out := new(SearchSpansTimeSeriesResponse)
	err := c.cc.Invoke(ctx, SpansDataSourceService_SearchSpansTimeSeries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spansDataSourceServiceClient) SearchSpansEvents(ctx context.Context, in *SearchSpansEventsRequest, opts ...grpc.CallOption) (*SearchSpansEventsResponse, error) {
	out := new(SearchSpansEventsResponse)
	err := c.cc.Invoke(ctx, SpansDataSourceService_SearchSpansEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spansDataSourceServiceClient) SearchSpansEventGroups(ctx context.Context, in *SearchSpansEventGroupsRequest, opts ...grpc.CallOption) (*SearchSpansEventGroupsResponse, error) {
	out := new(SearchSpansEventGroupsResponse)
	err := c.cc.Invoke(ctx, SpansDataSourceService_SearchSpansEventGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spansDataSourceServiceClient) SearchGroupedSpansSeries(ctx context.Context, in *SearchGroupedSpansSeriesRequest, opts ...grpc.CallOption) (*SearchGroupedSpansSeriesResponse, error) {
	out := new(SearchGroupedSpansSeriesResponse)
	err := c.cc.Invoke(ctx, SpansDataSourceService_SearchGroupedSpansSeries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spansDataSourceServiceClient) SearchSpansGroupedTimeSeries(ctx context.Context, in *SearchSpansGroupedTimeSeriesRequest, opts ...grpc.CallOption) (*SearchSpansGroupedTimeSeriesResponse, error) {
	out := new(SearchSpansGroupedTimeSeriesResponse)
	err := c.cc.Invoke(ctx, SpansDataSourceService_SearchSpansGroupedTimeSeries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spansDataSourceServiceClient) SearchSpansTimeValue(ctx context.Context, in *SearchSpansTimeValueRequest, opts ...grpc.CallOption) (*SearchSpansTimeValueResponse, error) {
	out := new(SearchSpansTimeValueResponse)
	err := c.cc.Invoke(ctx, SpansDataSourceService_SearchSpansTimeValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spansDataSourceServiceClient) SearchSpansAnnotationEvents(ctx context.Context, in *SearchSpansAnnotationEventsRequest, opts ...grpc.CallOption) (*SearchSpansAnnotationEventsResponse, error) {
	out := new(SearchSpansAnnotationEventsResponse)
	err := c.cc.Invoke(ctx, SpansDataSourceService_SearchSpansAnnotationEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpansDataSourceServiceServer is the server API for SpansDataSourceService service.
// All implementations must embed UnimplementedSpansDataSourceServiceServer
// for forward compatibility
type SpansDataSourceServiceServer interface {
	SearchSpansTimeSeries(context.Context, *SearchSpansTimeSeriesRequest) (*SearchSpansTimeSeriesResponse, error)
	SearchSpansEvents(context.Context, *SearchSpansEventsRequest) (*SearchSpansEventsResponse, error)
	SearchSpansEventGroups(context.Context, *SearchSpansEventGroupsRequest) (*SearchSpansEventGroupsResponse, error)
	SearchGroupedSpansSeries(context.Context, *SearchGroupedSpansSeriesRequest) (*SearchGroupedSpansSeriesResponse, error)
	SearchSpansGroupedTimeSeries(context.Context, *SearchSpansGroupedTimeSeriesRequest) (*SearchSpansGroupedTimeSeriesResponse, error)
	SearchSpansTimeValue(context.Context, *SearchSpansTimeValueRequest) (*SearchSpansTimeValueResponse, error)
	SearchSpansAnnotationEvents(context.Context, *SearchSpansAnnotationEventsRequest) (*SearchSpansAnnotationEventsResponse, error)
	mustEmbedUnimplementedSpansDataSourceServiceServer()
}

// UnimplementedSpansDataSourceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSpansDataSourceServiceServer struct {
}

func (UnimplementedSpansDataSourceServiceServer) SearchSpansTimeSeries(context.Context, *SearchSpansTimeSeriesRequest) (*SearchSpansTimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSpansTimeSeries not implemented")
}
func (UnimplementedSpansDataSourceServiceServer) SearchSpansEvents(context.Context, *SearchSpansEventsRequest) (*SearchSpansEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSpansEvents not implemented")
}
func (UnimplementedSpansDataSourceServiceServer) SearchSpansEventGroups(context.Context, *SearchSpansEventGroupsRequest) (*SearchSpansEventGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSpansEventGroups not implemented")
}
func (UnimplementedSpansDataSourceServiceServer) SearchGroupedSpansSeries(context.Context, *SearchGroupedSpansSeriesRequest) (*SearchGroupedSpansSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchGroupedSpansSeries not implemented")
}
func (UnimplementedSpansDataSourceServiceServer) SearchSpansGroupedTimeSeries(context.Context, *SearchSpansGroupedTimeSeriesRequest) (*SearchSpansGroupedTimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSpansGroupedTimeSeries not implemented")
}
func (UnimplementedSpansDataSourceServiceServer) SearchSpansTimeValue(context.Context, *SearchSpansTimeValueRequest) (*SearchSpansTimeValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSpansTimeValue not implemented")
}
func (UnimplementedSpansDataSourceServiceServer) SearchSpansAnnotationEvents(context.Context, *SearchSpansAnnotationEventsRequest) (*SearchSpansAnnotationEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSpansAnnotationEvents not implemented")
}
func (UnimplementedSpansDataSourceServiceServer) mustEmbedUnimplementedSpansDataSourceServiceServer() {
}

// UnsafeSpansDataSourceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpansDataSourceServiceServer will
// result in compilation errors.
type UnsafeSpansDataSourceServiceServer interface {
	mustEmbedUnimplementedSpansDataSourceServiceServer()
}

func RegisterSpansDataSourceServiceServer(s grpc.ServiceRegistrar, srv SpansDataSourceServiceServer) {
	s.RegisterService(&SpansDataSourceService_ServiceDesc, srv)
}

func _SpansDataSourceService_SearchSpansTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchSpansTimeSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpansDataSourceServiceServer).SearchSpansTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpansDataSourceService_SearchSpansTimeSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpansDataSourceServiceServer).SearchSpansTimeSeries(ctx, req.(*SearchSpansTimeSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpansDataSourceService_SearchSpansEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchSpansEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpansDataSourceServiceServer).SearchSpansEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpansDataSourceService_SearchSpansEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpansDataSourceServiceServer).SearchSpansEvents(ctx, req.(*SearchSpansEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpansDataSourceService_SearchSpansEventGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchSpansEventGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpansDataSourceServiceServer).SearchSpansEventGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpansDataSourceService_SearchSpansEventGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpansDataSourceServiceServer).SearchSpansEventGroups(ctx, req.(*SearchSpansEventGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpansDataSourceService_SearchGroupedSpansSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchGroupedSpansSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpansDataSourceServiceServer).SearchGroupedSpansSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpansDataSourceService_SearchGroupedSpansSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpansDataSourceServiceServer).SearchGroupedSpansSeries(ctx, req.(*SearchGroupedSpansSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpansDataSourceService_SearchSpansGroupedTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchSpansGroupedTimeSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpansDataSourceServiceServer).SearchSpansGroupedTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpansDataSourceService_SearchSpansGroupedTimeSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpansDataSourceServiceServer).SearchSpansGroupedTimeSeries(ctx, req.(*SearchSpansGroupedTimeSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpansDataSourceService_SearchSpansTimeValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchSpansTimeValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpansDataSourceServiceServer).SearchSpansTimeValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpansDataSourceService_SearchSpansTimeValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpansDataSourceServiceServer).SearchSpansTimeValue(ctx, req.(*SearchSpansTimeValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpansDataSourceService_SearchSpansAnnotationEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchSpansAnnotationEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpansDataSourceServiceServer).SearchSpansAnnotationEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpansDataSourceService_SearchSpansAnnotationEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpansDataSourceServiceServer).SearchSpansAnnotationEvents(ctx, req.(*SearchSpansAnnotationEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpansDataSourceService_ServiceDesc is the grpc.ServiceDesc for SpansDataSourceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpansDataSourceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.dashboards.v1.services.SpansDataSourceService",
	HandlerType: (*SpansDataSourceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchSpansTimeSeries",
			Handler:    _SpansDataSourceService_SearchSpansTimeSeries_Handler,
		},
		{
			MethodName: "SearchSpansEvents",
			Handler:    _SpansDataSourceService_SearchSpansEvents_Handler,
		},
		{
			MethodName: "SearchSpansEventGroups",
			Handler:    _SpansDataSourceService_SearchSpansEventGroups_Handler,
		},
		{
			MethodName: "SearchGroupedSpansSeries",
			Handler:    _SpansDataSourceService_SearchGroupedSpansSeries_Handler,
		},
		{
			MethodName: "SearchSpansGroupedTimeSeries",
			Handler:    _SpansDataSourceService_SearchSpansGroupedTimeSeries_Handler,
		},
		{
			MethodName: "SearchSpansTimeValue",
			Handler:    _SpansDataSourceService_SearchSpansTimeValue_Handler,
		},
		{
			MethodName: "SearchSpansAnnotationEvents",
			Handler:    _SpansDataSourceService_SearchSpansAnnotationEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/dashboards/v1/services/spans_data_source_service.proto",
}
