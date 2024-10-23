// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.2
// source: com/coralogixapis/dashboards/v1/services/archive_logs_data_source_service.proto

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
	ArchiveLogsDataSourceService_SearchArchiveLogsTimeSeries_FullMethodName        = "/com.coralogixapis.dashboards.v1.services.ArchiveLogsDataSourceService/SearchArchiveLogsTimeSeries"
	ArchiveLogsDataSourceService_SearchArchiveLogsEvents_FullMethodName            = "/com.coralogixapis.dashboards.v1.services.ArchiveLogsDataSourceService/SearchArchiveLogsEvents"
	ArchiveLogsDataSourceService_SearchArchiveLogsEventsCount_FullMethodName       = "/com.coralogixapis.dashboards.v1.services.ArchiveLogsDataSourceService/SearchArchiveLogsEventsCount"
	ArchiveLogsDataSourceService_SearchArchiveLogsEventGroups_FullMethodName       = "/com.coralogixapis.dashboards.v1.services.ArchiveLogsDataSourceService/SearchArchiveLogsEventGroups"
	ArchiveLogsDataSourceService_SearchArchiveGroupedLogsSeries_FullMethodName     = "/com.coralogixapis.dashboards.v1.services.ArchiveLogsDataSourceService/SearchArchiveGroupedLogsSeries"
	ArchiveLogsDataSourceService_SearchArchiveGroupedLogsTimeSeries_FullMethodName = "/com.coralogixapis.dashboards.v1.services.ArchiveLogsDataSourceService/SearchArchiveGroupedLogsTimeSeries"
	ArchiveLogsDataSourceService_SearchArchiveLogsGroupedValues_FullMethodName     = "/com.coralogixapis.dashboards.v1.services.ArchiveLogsDataSourceService/SearchArchiveLogsGroupedValues"
)

// ArchiveLogsDataSourceServiceClient is the client API for ArchiveLogsDataSourceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArchiveLogsDataSourceServiceClient interface {
	SearchArchiveLogsTimeSeries(ctx context.Context, in *SearchArchiveLogsTimeSeriesRequest, opts ...grpc.CallOption) (*SearchArchiveLogsTimeSeriesResponse, error)
	SearchArchiveLogsEvents(ctx context.Context, in *SearchArchiveLogsEventsRequest, opts ...grpc.CallOption) (*SearchArchiveLogsEventsResponse, error)
	SearchArchiveLogsEventsCount(ctx context.Context, in *SearchArchiveLogsEventsCountRequest, opts ...grpc.CallOption) (*SearchArchiveLogsEventsCountResponse, error)
	SearchArchiveLogsEventGroups(ctx context.Context, in *SearchArchiveLogsEventGroupsRequest, opts ...grpc.CallOption) (*SearchArchiveLogsEventGroupsResponse, error)
	SearchArchiveGroupedLogsSeries(ctx context.Context, in *SearchArchiveGroupedLogsSeriesRequest, opts ...grpc.CallOption) (*SearchArchiveGroupedLogsSeriesResponse, error)
	SearchArchiveGroupedLogsTimeSeries(ctx context.Context, in *SearchArchiveGroupedLogsTimeSeriesRequest, opts ...grpc.CallOption) (*SearchArchiveGroupedLogsTimeSeriesResponse, error)
	SearchArchiveLogsGroupedValues(ctx context.Context, in *SearchArchiveLogsGroupedValuesRequest, opts ...grpc.CallOption) (*SearchArchiveLogsGroupedValuesResponse, error)
}

type archiveLogsDataSourceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArchiveLogsDataSourceServiceClient(cc grpc.ClientConnInterface) ArchiveLogsDataSourceServiceClient {
	return &archiveLogsDataSourceServiceClient{cc}
}

func (c *archiveLogsDataSourceServiceClient) SearchArchiveLogsTimeSeries(ctx context.Context, in *SearchArchiveLogsTimeSeriesRequest, opts ...grpc.CallOption) (*SearchArchiveLogsTimeSeriesResponse, error) {
	out := new(SearchArchiveLogsTimeSeriesResponse)
	err := c.cc.Invoke(ctx, ArchiveLogsDataSourceService_SearchArchiveLogsTimeSeries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *archiveLogsDataSourceServiceClient) SearchArchiveLogsEvents(ctx context.Context, in *SearchArchiveLogsEventsRequest, opts ...grpc.CallOption) (*SearchArchiveLogsEventsResponse, error) {
	out := new(SearchArchiveLogsEventsResponse)
	err := c.cc.Invoke(ctx, ArchiveLogsDataSourceService_SearchArchiveLogsEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *archiveLogsDataSourceServiceClient) SearchArchiveLogsEventsCount(ctx context.Context, in *SearchArchiveLogsEventsCountRequest, opts ...grpc.CallOption) (*SearchArchiveLogsEventsCountResponse, error) {
	out := new(SearchArchiveLogsEventsCountResponse)
	err := c.cc.Invoke(ctx, ArchiveLogsDataSourceService_SearchArchiveLogsEventsCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *archiveLogsDataSourceServiceClient) SearchArchiveLogsEventGroups(ctx context.Context, in *SearchArchiveLogsEventGroupsRequest, opts ...grpc.CallOption) (*SearchArchiveLogsEventGroupsResponse, error) {
	out := new(SearchArchiveLogsEventGroupsResponse)
	err := c.cc.Invoke(ctx, ArchiveLogsDataSourceService_SearchArchiveLogsEventGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *archiveLogsDataSourceServiceClient) SearchArchiveGroupedLogsSeries(ctx context.Context, in *SearchArchiveGroupedLogsSeriesRequest, opts ...grpc.CallOption) (*SearchArchiveGroupedLogsSeriesResponse, error) {
	out := new(SearchArchiveGroupedLogsSeriesResponse)
	err := c.cc.Invoke(ctx, ArchiveLogsDataSourceService_SearchArchiveGroupedLogsSeries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *archiveLogsDataSourceServiceClient) SearchArchiveGroupedLogsTimeSeries(ctx context.Context, in *SearchArchiveGroupedLogsTimeSeriesRequest, opts ...grpc.CallOption) (*SearchArchiveGroupedLogsTimeSeriesResponse, error) {
	out := new(SearchArchiveGroupedLogsTimeSeriesResponse)
	err := c.cc.Invoke(ctx, ArchiveLogsDataSourceService_SearchArchiveGroupedLogsTimeSeries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *archiveLogsDataSourceServiceClient) SearchArchiveLogsGroupedValues(ctx context.Context, in *SearchArchiveLogsGroupedValuesRequest, opts ...grpc.CallOption) (*SearchArchiveLogsGroupedValuesResponse, error) {
	out := new(SearchArchiveLogsGroupedValuesResponse)
	err := c.cc.Invoke(ctx, ArchiveLogsDataSourceService_SearchArchiveLogsGroupedValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArchiveLogsDataSourceServiceServer is the server API for ArchiveLogsDataSourceService service.
// All implementations must embed UnimplementedArchiveLogsDataSourceServiceServer
// for forward compatibility
type ArchiveLogsDataSourceServiceServer interface {
	SearchArchiveLogsTimeSeries(context.Context, *SearchArchiveLogsTimeSeriesRequest) (*SearchArchiveLogsTimeSeriesResponse, error)
	SearchArchiveLogsEvents(context.Context, *SearchArchiveLogsEventsRequest) (*SearchArchiveLogsEventsResponse, error)
	SearchArchiveLogsEventsCount(context.Context, *SearchArchiveLogsEventsCountRequest) (*SearchArchiveLogsEventsCountResponse, error)
	SearchArchiveLogsEventGroups(context.Context, *SearchArchiveLogsEventGroupsRequest) (*SearchArchiveLogsEventGroupsResponse, error)
	SearchArchiveGroupedLogsSeries(context.Context, *SearchArchiveGroupedLogsSeriesRequest) (*SearchArchiveGroupedLogsSeriesResponse, error)
	SearchArchiveGroupedLogsTimeSeries(context.Context, *SearchArchiveGroupedLogsTimeSeriesRequest) (*SearchArchiveGroupedLogsTimeSeriesResponse, error)
	SearchArchiveLogsGroupedValues(context.Context, *SearchArchiveLogsGroupedValuesRequest) (*SearchArchiveLogsGroupedValuesResponse, error)
	mustEmbedUnimplementedArchiveLogsDataSourceServiceServer()
}

// UnimplementedArchiveLogsDataSourceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArchiveLogsDataSourceServiceServer struct {
}

func (UnimplementedArchiveLogsDataSourceServiceServer) SearchArchiveLogsTimeSeries(context.Context, *SearchArchiveLogsTimeSeriesRequest) (*SearchArchiveLogsTimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchArchiveLogsTimeSeries not implemented")
}
func (UnimplementedArchiveLogsDataSourceServiceServer) SearchArchiveLogsEvents(context.Context, *SearchArchiveLogsEventsRequest) (*SearchArchiveLogsEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchArchiveLogsEvents not implemented")
}
func (UnimplementedArchiveLogsDataSourceServiceServer) SearchArchiveLogsEventsCount(context.Context, *SearchArchiveLogsEventsCountRequest) (*SearchArchiveLogsEventsCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchArchiveLogsEventsCount not implemented")
}
func (UnimplementedArchiveLogsDataSourceServiceServer) SearchArchiveLogsEventGroups(context.Context, *SearchArchiveLogsEventGroupsRequest) (*SearchArchiveLogsEventGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchArchiveLogsEventGroups not implemented")
}
func (UnimplementedArchiveLogsDataSourceServiceServer) SearchArchiveGroupedLogsSeries(context.Context, *SearchArchiveGroupedLogsSeriesRequest) (*SearchArchiveGroupedLogsSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchArchiveGroupedLogsSeries not implemented")
}
func (UnimplementedArchiveLogsDataSourceServiceServer) SearchArchiveGroupedLogsTimeSeries(context.Context, *SearchArchiveGroupedLogsTimeSeriesRequest) (*SearchArchiveGroupedLogsTimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchArchiveGroupedLogsTimeSeries not implemented")
}
func (UnimplementedArchiveLogsDataSourceServiceServer) SearchArchiveLogsGroupedValues(context.Context, *SearchArchiveLogsGroupedValuesRequest) (*SearchArchiveLogsGroupedValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchArchiveLogsGroupedValues not implemented")
}
func (UnimplementedArchiveLogsDataSourceServiceServer) mustEmbedUnimplementedArchiveLogsDataSourceServiceServer() {
}

// UnsafeArchiveLogsDataSourceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArchiveLogsDataSourceServiceServer will
// result in compilation errors.
type UnsafeArchiveLogsDataSourceServiceServer interface {
	mustEmbedUnimplementedArchiveLogsDataSourceServiceServer()
}

func RegisterArchiveLogsDataSourceServiceServer(s grpc.ServiceRegistrar, srv ArchiveLogsDataSourceServiceServer) {
	s.RegisterService(&ArchiveLogsDataSourceService_ServiceDesc, srv)
}

func _ArchiveLogsDataSourceService_SearchArchiveLogsTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchArchiveLogsTimeSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveLogsDataSourceServiceServer).SearchArchiveLogsTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArchiveLogsDataSourceService_SearchArchiveLogsTimeSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveLogsDataSourceServiceServer).SearchArchiveLogsTimeSeries(ctx, req.(*SearchArchiveLogsTimeSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArchiveLogsDataSourceService_SearchArchiveLogsEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchArchiveLogsEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveLogsDataSourceServiceServer).SearchArchiveLogsEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArchiveLogsDataSourceService_SearchArchiveLogsEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveLogsDataSourceServiceServer).SearchArchiveLogsEvents(ctx, req.(*SearchArchiveLogsEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArchiveLogsDataSourceService_SearchArchiveLogsEventsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchArchiveLogsEventsCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveLogsDataSourceServiceServer).SearchArchiveLogsEventsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArchiveLogsDataSourceService_SearchArchiveLogsEventsCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveLogsDataSourceServiceServer).SearchArchiveLogsEventsCount(ctx, req.(*SearchArchiveLogsEventsCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArchiveLogsDataSourceService_SearchArchiveLogsEventGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchArchiveLogsEventGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveLogsDataSourceServiceServer).SearchArchiveLogsEventGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArchiveLogsDataSourceService_SearchArchiveLogsEventGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveLogsDataSourceServiceServer).SearchArchiveLogsEventGroups(ctx, req.(*SearchArchiveLogsEventGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArchiveLogsDataSourceService_SearchArchiveGroupedLogsSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchArchiveGroupedLogsSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveLogsDataSourceServiceServer).SearchArchiveGroupedLogsSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArchiveLogsDataSourceService_SearchArchiveGroupedLogsSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveLogsDataSourceServiceServer).SearchArchiveGroupedLogsSeries(ctx, req.(*SearchArchiveGroupedLogsSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArchiveLogsDataSourceService_SearchArchiveGroupedLogsTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchArchiveGroupedLogsTimeSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveLogsDataSourceServiceServer).SearchArchiveGroupedLogsTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArchiveLogsDataSourceService_SearchArchiveGroupedLogsTimeSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveLogsDataSourceServiceServer).SearchArchiveGroupedLogsTimeSeries(ctx, req.(*SearchArchiveGroupedLogsTimeSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArchiveLogsDataSourceService_SearchArchiveLogsGroupedValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchArchiveLogsGroupedValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveLogsDataSourceServiceServer).SearchArchiveLogsGroupedValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArchiveLogsDataSourceService_SearchArchiveLogsGroupedValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveLogsDataSourceServiceServer).SearchArchiveLogsGroupedValues(ctx, req.(*SearchArchiveLogsGroupedValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArchiveLogsDataSourceService_ServiceDesc is the grpc.ServiceDesc for ArchiveLogsDataSourceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArchiveLogsDataSourceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.dashboards.v1.services.ArchiveLogsDataSourceService",
	HandlerType: (*ArchiveLogsDataSourceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchArchiveLogsTimeSeries",
			Handler:    _ArchiveLogsDataSourceService_SearchArchiveLogsTimeSeries_Handler,
		},
		{
			MethodName: "SearchArchiveLogsEvents",
			Handler:    _ArchiveLogsDataSourceService_SearchArchiveLogsEvents_Handler,
		},
		{
			MethodName: "SearchArchiveLogsEventsCount",
			Handler:    _ArchiveLogsDataSourceService_SearchArchiveLogsEventsCount_Handler,
		},
		{
			MethodName: "SearchArchiveLogsEventGroups",
			Handler:    _ArchiveLogsDataSourceService_SearchArchiveLogsEventGroups_Handler,
		},
		{
			MethodName: "SearchArchiveGroupedLogsSeries",
			Handler:    _ArchiveLogsDataSourceService_SearchArchiveGroupedLogsSeries_Handler,
		},
		{
			MethodName: "SearchArchiveGroupedLogsTimeSeries",
			Handler:    _ArchiveLogsDataSourceService_SearchArchiveGroupedLogsTimeSeries_Handler,
		},
		{
			MethodName: "SearchArchiveLogsGroupedValues",
			Handler:    _ArchiveLogsDataSourceService_SearchArchiveLogsGroupedValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/dashboards/v1/services/archive_logs_data_source_service.proto",
}
