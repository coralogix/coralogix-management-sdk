// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: com/coralogix/datausage/v1/data_usage_service.proto

package v1

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
	DataUsageService_GetTeamsDailyUsage_FullMethodName          = "/com.coralogix.datausage.v1.DataUsageService/GetTeamsDailyUsage"
	DataUsageService_GetTeamsBlocks_FullMethodName              = "/com.coralogix.datausage.v1.DataUsageService/GetTeamsBlocks"
	DataUsageService_GetTeamsQuotaHistory_FullMethodName        = "/com.coralogix.datausage.v1.DataUsageService/GetTeamsQuotaHistory"
	DataUsageService_GetTeamsQuota_FullMethodName               = "/com.coralogix.datausage.v1.DataUsageService/GetTeamsQuota"
	DataUsageService_GetDataUsage_FullMethodName                = "/com.coralogix.datausage.v1.DataUsageService/GetDataUsage"
	DataUsageService_GetDetailedDataUsage_FullMethodName        = "/com.coralogix.datausage.v1.DataUsageService/GetDetailedDataUsage"
	DataUsageService_GetDetailedDataUsageChunked_FullMethodName = "/com.coralogix.datausage.v1.DataUsageService/GetDetailedDataUsageChunked"
)

// DataUsageServiceClient is the client API for DataUsageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataUsageServiceClient interface {
	// / Returns a list of quota (units) used per day for a list of teams in the requested days
	GetTeamsDailyUsage(ctx context.Context, in *GetTeamsDailyUsageRequest, opts ...grpc.CallOption) (*GetTeamsDailyUsageResponse, error)
	// / Returns a list of block events (can be more than one per day) for a list of teams in the requested days
	GetTeamsBlocks(ctx context.Context, in *GetTeamsBlocksRequest, opts ...grpc.CallOption) (*GetTeamsBlocksResponse, error)
	// / Returns a list of quota updates (can be more than one per day) for a list of teams in the requested days
	GetTeamsQuotaHistory(ctx context.Context, in *GetTeamsQuotaHistoryRequest, opts ...grpc.CallOption) (*GetTeamsQuotaHistoryResponse, error)
	// / Returns the quota plan for a given set of teams at the requested timestamp
	GetTeamsQuota(ctx context.Context, in *GetTeamsQuotaRequest, opts ...grpc.CallOption) (*GetTeamsQuotaResponse, error)
	// / Returns a list of data usage for a given set of teams, time_range and resolution
	GetDataUsage(ctx context.Context, in *GetDataUsageRequest, opts ...grpc.CallOption) (*GetDataUsageResponse, error)
	GetDetailedDataUsage(ctx context.Context, in *GetDetailedDataUsageRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetDetailedDataUsageResponse], error)
	GetDetailedDataUsageChunked(ctx context.Context, in *GetDetailedDataUsageChunkedRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetDetailedDataUsageChunkedResponse], error)
}

type dataUsageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataUsageServiceClient(cc grpc.ClientConnInterface) DataUsageServiceClient {
	return &dataUsageServiceClient{cc}
}

func (c *dataUsageServiceClient) GetTeamsDailyUsage(ctx context.Context, in *GetTeamsDailyUsageRequest, opts ...grpc.CallOption) (*GetTeamsDailyUsageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTeamsDailyUsageResponse)
	err := c.cc.Invoke(ctx, DataUsageService_GetTeamsDailyUsage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataUsageServiceClient) GetTeamsBlocks(ctx context.Context, in *GetTeamsBlocksRequest, opts ...grpc.CallOption) (*GetTeamsBlocksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTeamsBlocksResponse)
	err := c.cc.Invoke(ctx, DataUsageService_GetTeamsBlocks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataUsageServiceClient) GetTeamsQuotaHistory(ctx context.Context, in *GetTeamsQuotaHistoryRequest, opts ...grpc.CallOption) (*GetTeamsQuotaHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTeamsQuotaHistoryResponse)
	err := c.cc.Invoke(ctx, DataUsageService_GetTeamsQuotaHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataUsageServiceClient) GetTeamsQuota(ctx context.Context, in *GetTeamsQuotaRequest, opts ...grpc.CallOption) (*GetTeamsQuotaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTeamsQuotaResponse)
	err := c.cc.Invoke(ctx, DataUsageService_GetTeamsQuota_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataUsageServiceClient) GetDataUsage(ctx context.Context, in *GetDataUsageRequest, opts ...grpc.CallOption) (*GetDataUsageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDataUsageResponse)
	err := c.cc.Invoke(ctx, DataUsageService_GetDataUsage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataUsageServiceClient) GetDetailedDataUsage(ctx context.Context, in *GetDetailedDataUsageRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetDetailedDataUsageResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DataUsageService_ServiceDesc.Streams[0], DataUsageService_GetDetailedDataUsage_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetDetailedDataUsageRequest, GetDetailedDataUsageResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DataUsageService_GetDetailedDataUsageClient = grpc.ServerStreamingClient[GetDetailedDataUsageResponse]

func (c *dataUsageServiceClient) GetDetailedDataUsageChunked(ctx context.Context, in *GetDetailedDataUsageChunkedRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetDetailedDataUsageChunkedResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DataUsageService_ServiceDesc.Streams[1], DataUsageService_GetDetailedDataUsageChunked_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetDetailedDataUsageChunkedRequest, GetDetailedDataUsageChunkedResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DataUsageService_GetDetailedDataUsageChunkedClient = grpc.ServerStreamingClient[GetDetailedDataUsageChunkedResponse]

// DataUsageServiceServer is the server API for DataUsageService service.
// All implementations must embed UnimplementedDataUsageServiceServer
// for forward compatibility.
type DataUsageServiceServer interface {
	// / Returns a list of quota (units) used per day for a list of teams in the requested days
	GetTeamsDailyUsage(context.Context, *GetTeamsDailyUsageRequest) (*GetTeamsDailyUsageResponse, error)
	// / Returns a list of block events (can be more than one per day) for a list of teams in the requested days
	GetTeamsBlocks(context.Context, *GetTeamsBlocksRequest) (*GetTeamsBlocksResponse, error)
	// / Returns a list of quota updates (can be more than one per day) for a list of teams in the requested days
	GetTeamsQuotaHistory(context.Context, *GetTeamsQuotaHistoryRequest) (*GetTeamsQuotaHistoryResponse, error)
	// / Returns the quota plan for a given set of teams at the requested timestamp
	GetTeamsQuota(context.Context, *GetTeamsQuotaRequest) (*GetTeamsQuotaResponse, error)
	// / Returns a list of data usage for a given set of teams, time_range and resolution
	GetDataUsage(context.Context, *GetDataUsageRequest) (*GetDataUsageResponse, error)
	GetDetailedDataUsage(*GetDetailedDataUsageRequest, grpc.ServerStreamingServer[GetDetailedDataUsageResponse]) error
	GetDetailedDataUsageChunked(*GetDetailedDataUsageChunkedRequest, grpc.ServerStreamingServer[GetDetailedDataUsageChunkedResponse]) error
	mustEmbedUnimplementedDataUsageServiceServer()
}

// UnimplementedDataUsageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDataUsageServiceServer struct{}

func (UnimplementedDataUsageServiceServer) GetTeamsDailyUsage(context.Context, *GetTeamsDailyUsageRequest) (*GetTeamsDailyUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamsDailyUsage not implemented")
}
func (UnimplementedDataUsageServiceServer) GetTeamsBlocks(context.Context, *GetTeamsBlocksRequest) (*GetTeamsBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamsBlocks not implemented")
}
func (UnimplementedDataUsageServiceServer) GetTeamsQuotaHistory(context.Context, *GetTeamsQuotaHistoryRequest) (*GetTeamsQuotaHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamsQuotaHistory not implemented")
}
func (UnimplementedDataUsageServiceServer) GetTeamsQuota(context.Context, *GetTeamsQuotaRequest) (*GetTeamsQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamsQuota not implemented")
}
func (UnimplementedDataUsageServiceServer) GetDataUsage(context.Context, *GetDataUsageRequest) (*GetDataUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataUsage not implemented")
}
func (UnimplementedDataUsageServiceServer) GetDetailedDataUsage(*GetDetailedDataUsageRequest, grpc.ServerStreamingServer[GetDetailedDataUsageResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetDetailedDataUsage not implemented")
}
func (UnimplementedDataUsageServiceServer) GetDetailedDataUsageChunked(*GetDetailedDataUsageChunkedRequest, grpc.ServerStreamingServer[GetDetailedDataUsageChunkedResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetDetailedDataUsageChunked not implemented")
}
func (UnimplementedDataUsageServiceServer) mustEmbedUnimplementedDataUsageServiceServer() {}
func (UnimplementedDataUsageServiceServer) testEmbeddedByValue()                          {}

// UnsafeDataUsageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataUsageServiceServer will
// result in compilation errors.
type UnsafeDataUsageServiceServer interface {
	mustEmbedUnimplementedDataUsageServiceServer()
}

func RegisterDataUsageServiceServer(s grpc.ServiceRegistrar, srv DataUsageServiceServer) {
	// If the following call pancis, it indicates UnimplementedDataUsageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DataUsageService_ServiceDesc, srv)
}

func _DataUsageService_GetTeamsDailyUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamsDailyUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataUsageServiceServer).GetTeamsDailyUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataUsageService_GetTeamsDailyUsage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataUsageServiceServer).GetTeamsDailyUsage(ctx, req.(*GetTeamsDailyUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataUsageService_GetTeamsBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamsBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataUsageServiceServer).GetTeamsBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataUsageService_GetTeamsBlocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataUsageServiceServer).GetTeamsBlocks(ctx, req.(*GetTeamsBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataUsageService_GetTeamsQuotaHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamsQuotaHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataUsageServiceServer).GetTeamsQuotaHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataUsageService_GetTeamsQuotaHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataUsageServiceServer).GetTeamsQuotaHistory(ctx, req.(*GetTeamsQuotaHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataUsageService_GetTeamsQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamsQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataUsageServiceServer).GetTeamsQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataUsageService_GetTeamsQuota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataUsageServiceServer).GetTeamsQuota(ctx, req.(*GetTeamsQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataUsageService_GetDataUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataUsageServiceServer).GetDataUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataUsageService_GetDataUsage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataUsageServiceServer).GetDataUsage(ctx, req.(*GetDataUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataUsageService_GetDetailedDataUsage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDetailedDataUsageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataUsageServiceServer).GetDetailedDataUsage(m, &grpc.GenericServerStream[GetDetailedDataUsageRequest, GetDetailedDataUsageResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DataUsageService_GetDetailedDataUsageServer = grpc.ServerStreamingServer[GetDetailedDataUsageResponse]

func _DataUsageService_GetDetailedDataUsageChunked_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDetailedDataUsageChunkedRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataUsageServiceServer).GetDetailedDataUsageChunked(m, &grpc.GenericServerStream[GetDetailedDataUsageChunkedRequest, GetDetailedDataUsageChunkedResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DataUsageService_GetDetailedDataUsageChunkedServer = grpc.ServerStreamingServer[GetDetailedDataUsageChunkedResponse]

// DataUsageService_ServiceDesc is the grpc.ServiceDesc for DataUsageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataUsageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.datausage.v1.DataUsageService",
	HandlerType: (*DataUsageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeamsDailyUsage",
			Handler:    _DataUsageService_GetTeamsDailyUsage_Handler,
		},
		{
			MethodName: "GetTeamsBlocks",
			Handler:    _DataUsageService_GetTeamsBlocks_Handler,
		},
		{
			MethodName: "GetTeamsQuotaHistory",
			Handler:    _DataUsageService_GetTeamsQuotaHistory_Handler,
		},
		{
			MethodName: "GetTeamsQuota",
			Handler:    _DataUsageService_GetTeamsQuota_Handler,
		},
		{
			MethodName: "GetDataUsage",
			Handler:    _DataUsageService_GetDataUsage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDetailedDataUsage",
			Handler:       _DataUsageService_GetDetailedDataUsage_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetDetailedDataUsageChunked",
			Handler:       _DataUsageService_GetDetailedDataUsageChunked_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "com/coralogix/datausage/v1/data_usage_service.proto",
}
