// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DataUsageService_GetTeamsDailyUsage_FullMethodName          = "/com.coralogix.datausage.v1.DataUsageService/GetTeamsDailyUsage"
	DataUsageService_GetTeamsBlocks_FullMethodName              = "/com.coralogix.datausage.v1.DataUsageService/GetTeamsBlocks"
	DataUsageService_GetTeamsQuotaHistory_FullMethodName        = "/com.coralogix.datausage.v1.DataUsageService/GetTeamsQuotaHistory"
	DataUsageService_GetTeamsQuota_FullMethodName               = "/com.coralogix.datausage.v1.DataUsageService/GetTeamsQuota"
	DataUsageService_GetTeamsDetailedUsage_FullMethodName       = "/com.coralogix.datausage.v1.DataUsageService/GetTeamsDetailedUsage"
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
	// Deprecated: Do not use.
	GetTeamsDetailedUsage(ctx context.Context, in *GetTeamsDetailedUsageRequest, opts ...grpc.CallOption) (DataUsageService_GetTeamsDetailedUsageClient, error)
	// / Returns a list of data usage for a given set of teams, time_range and resolution
	GetDataUsage(ctx context.Context, in *GetDataUsageRequest, opts ...grpc.CallOption) (*GetDataUsageResponse, error)
	GetDetailedDataUsage(ctx context.Context, in *GetDetailedDataUsageRequest, opts ...grpc.CallOption) (DataUsageService_GetDetailedDataUsageClient, error)
	GetDetailedDataUsageChunked(ctx context.Context, in *GetDetailedDataUsageChunkedRequest, opts ...grpc.CallOption) (DataUsageService_GetDetailedDataUsageChunkedClient, error)
}

type dataUsageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataUsageServiceClient(cc grpc.ClientConnInterface) DataUsageServiceClient {
	return &dataUsageServiceClient{cc}
}

func (c *dataUsageServiceClient) GetTeamsDailyUsage(ctx context.Context, in *GetTeamsDailyUsageRequest, opts ...grpc.CallOption) (*GetTeamsDailyUsageResponse, error) {
	out := new(GetTeamsDailyUsageResponse)
	err := c.cc.Invoke(ctx, DataUsageService_GetTeamsDailyUsage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataUsageServiceClient) GetTeamsBlocks(ctx context.Context, in *GetTeamsBlocksRequest, opts ...grpc.CallOption) (*GetTeamsBlocksResponse, error) {
	out := new(GetTeamsBlocksResponse)
	err := c.cc.Invoke(ctx, DataUsageService_GetTeamsBlocks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataUsageServiceClient) GetTeamsQuotaHistory(ctx context.Context, in *GetTeamsQuotaHistoryRequest, opts ...grpc.CallOption) (*GetTeamsQuotaHistoryResponse, error) {
	out := new(GetTeamsQuotaHistoryResponse)
	err := c.cc.Invoke(ctx, DataUsageService_GetTeamsQuotaHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataUsageServiceClient) GetTeamsQuota(ctx context.Context, in *GetTeamsQuotaRequest, opts ...grpc.CallOption) (*GetTeamsQuotaResponse, error) {
	out := new(GetTeamsQuotaResponse)
	err := c.cc.Invoke(ctx, DataUsageService_GetTeamsQuota_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *dataUsageServiceClient) GetTeamsDetailedUsage(ctx context.Context, in *GetTeamsDetailedUsageRequest, opts ...grpc.CallOption) (DataUsageService_GetTeamsDetailedUsageClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataUsageService_ServiceDesc.Streams[0], DataUsageService_GetTeamsDetailedUsage_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &dataUsageServiceGetTeamsDetailedUsageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataUsageService_GetTeamsDetailedUsageClient interface {
	Recv() (*GetTeamsDetailedUsageResponse, error)
	grpc.ClientStream
}

type dataUsageServiceGetTeamsDetailedUsageClient struct {
	grpc.ClientStream
}

func (x *dataUsageServiceGetTeamsDetailedUsageClient) Recv() (*GetTeamsDetailedUsageResponse, error) {
	m := new(GetTeamsDetailedUsageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataUsageServiceClient) GetDataUsage(ctx context.Context, in *GetDataUsageRequest, opts ...grpc.CallOption) (*GetDataUsageResponse, error) {
	out := new(GetDataUsageResponse)
	err := c.cc.Invoke(ctx, DataUsageService_GetDataUsage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataUsageServiceClient) GetDetailedDataUsage(ctx context.Context, in *GetDetailedDataUsageRequest, opts ...grpc.CallOption) (DataUsageService_GetDetailedDataUsageClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataUsageService_ServiceDesc.Streams[1], DataUsageService_GetDetailedDataUsage_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &dataUsageServiceGetDetailedDataUsageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataUsageService_GetDetailedDataUsageClient interface {
	Recv() (*GetDetailedDataUsageResponse, error)
	grpc.ClientStream
}

type dataUsageServiceGetDetailedDataUsageClient struct {
	grpc.ClientStream
}

func (x *dataUsageServiceGetDetailedDataUsageClient) Recv() (*GetDetailedDataUsageResponse, error) {
	m := new(GetDetailedDataUsageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataUsageServiceClient) GetDetailedDataUsageChunked(ctx context.Context, in *GetDetailedDataUsageChunkedRequest, opts ...grpc.CallOption) (DataUsageService_GetDetailedDataUsageChunkedClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataUsageService_ServiceDesc.Streams[2], DataUsageService_GetDetailedDataUsageChunked_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &dataUsageServiceGetDetailedDataUsageChunkedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataUsageService_GetDetailedDataUsageChunkedClient interface {
	Recv() (*GetDetailedDataUsageChunkedResponse, error)
	grpc.ClientStream
}

type dataUsageServiceGetDetailedDataUsageChunkedClient struct {
	grpc.ClientStream
}

func (x *dataUsageServiceGetDetailedDataUsageChunkedClient) Recv() (*GetDetailedDataUsageChunkedResponse, error) {
	m := new(GetDetailedDataUsageChunkedResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataUsageServiceServer is the server API for DataUsageService service.
// All implementations must embed UnimplementedDataUsageServiceServer
// for forward compatibility
type DataUsageServiceServer interface {
	// / Returns a list of quota (units) used per day for a list of teams in the requested days
	GetTeamsDailyUsage(context.Context, *GetTeamsDailyUsageRequest) (*GetTeamsDailyUsageResponse, error)
	// / Returns a list of block events (can be more than one per day) for a list of teams in the requested days
	GetTeamsBlocks(context.Context, *GetTeamsBlocksRequest) (*GetTeamsBlocksResponse, error)
	// / Returns a list of quota updates (can be more than one per day) for a list of teams in the requested days
	GetTeamsQuotaHistory(context.Context, *GetTeamsQuotaHistoryRequest) (*GetTeamsQuotaHistoryResponse, error)
	// / Returns the quota plan for a given set of teams at the requested timestamp
	GetTeamsQuota(context.Context, *GetTeamsQuotaRequest) (*GetTeamsQuotaResponse, error)
	// Deprecated: Do not use.
	GetTeamsDetailedUsage(*GetTeamsDetailedUsageRequest, DataUsageService_GetTeamsDetailedUsageServer) error
	// / Returns a list of data usage for a given set of teams, time_range and resolution
	GetDataUsage(context.Context, *GetDataUsageRequest) (*GetDataUsageResponse, error)
	GetDetailedDataUsage(*GetDetailedDataUsageRequest, DataUsageService_GetDetailedDataUsageServer) error
	GetDetailedDataUsageChunked(*GetDetailedDataUsageChunkedRequest, DataUsageService_GetDetailedDataUsageChunkedServer) error
	mustEmbedUnimplementedDataUsageServiceServer()
}

// UnimplementedDataUsageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataUsageServiceServer struct {
}

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
func (UnimplementedDataUsageServiceServer) GetTeamsDetailedUsage(*GetTeamsDetailedUsageRequest, DataUsageService_GetTeamsDetailedUsageServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTeamsDetailedUsage not implemented")
}
func (UnimplementedDataUsageServiceServer) GetDataUsage(context.Context, *GetDataUsageRequest) (*GetDataUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataUsage not implemented")
}
func (UnimplementedDataUsageServiceServer) GetDetailedDataUsage(*GetDetailedDataUsageRequest, DataUsageService_GetDetailedDataUsageServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDetailedDataUsage not implemented")
}
func (UnimplementedDataUsageServiceServer) GetDetailedDataUsageChunked(*GetDetailedDataUsageChunkedRequest, DataUsageService_GetDetailedDataUsageChunkedServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDetailedDataUsageChunked not implemented")
}
func (UnimplementedDataUsageServiceServer) mustEmbedUnimplementedDataUsageServiceServer() {}

// UnsafeDataUsageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataUsageServiceServer will
// result in compilation errors.
type UnsafeDataUsageServiceServer interface {
	mustEmbedUnimplementedDataUsageServiceServer()
}

func RegisterDataUsageServiceServer(s grpc.ServiceRegistrar, srv DataUsageServiceServer) {
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

func _DataUsageService_GetTeamsDetailedUsage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTeamsDetailedUsageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataUsageServiceServer).GetTeamsDetailedUsage(m, &dataUsageServiceGetTeamsDetailedUsageServer{stream})
}

type DataUsageService_GetTeamsDetailedUsageServer interface {
	Send(*GetTeamsDetailedUsageResponse) error
	grpc.ServerStream
}

type dataUsageServiceGetTeamsDetailedUsageServer struct {
	grpc.ServerStream
}

func (x *dataUsageServiceGetTeamsDetailedUsageServer) Send(m *GetTeamsDetailedUsageResponse) error {
	return x.ServerStream.SendMsg(m)
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
	return srv.(DataUsageServiceServer).GetDetailedDataUsage(m, &dataUsageServiceGetDetailedDataUsageServer{stream})
}

type DataUsageService_GetDetailedDataUsageServer interface {
	Send(*GetDetailedDataUsageResponse) error
	grpc.ServerStream
}

type dataUsageServiceGetDetailedDataUsageServer struct {
	grpc.ServerStream
}

func (x *dataUsageServiceGetDetailedDataUsageServer) Send(m *GetDetailedDataUsageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DataUsageService_GetDetailedDataUsageChunked_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDetailedDataUsageChunkedRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataUsageServiceServer).GetDetailedDataUsageChunked(m, &dataUsageServiceGetDetailedDataUsageChunkedServer{stream})
}

type DataUsageService_GetDetailedDataUsageChunkedServer interface {
	Send(*GetDetailedDataUsageChunkedResponse) error
	grpc.ServerStream
}

type dataUsageServiceGetDetailedDataUsageChunkedServer struct {
	grpc.ServerStream
}

func (x *dataUsageServiceGetDetailedDataUsageChunkedServer) Send(m *GetDetailedDataUsageChunkedResponse) error {
	return x.ServerStream.SendMsg(m)
}

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
			StreamName:    "GetTeamsDetailedUsage",
			Handler:       _DataUsageService_GetTeamsDetailedUsage_Handler,
			ServerStreams: true,
		},
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
