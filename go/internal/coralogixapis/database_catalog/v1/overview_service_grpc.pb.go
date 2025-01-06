// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
// source: com/coralogixapis/database_catalog/v1/overview_service.proto

package v1

import (
	context "context"
	common "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm_shared/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DatabaseOverviewService_GetDatabaseOverviewStream_FullMethodName = "/com.coralogixapis.database_catalog.v1.DatabaseOverviewService/GetDatabaseOverviewStream"
	DatabaseOverviewService_GetApdex_FullMethodName                  = "/com.coralogixapis.database_catalog.v1.DatabaseOverviewService/GetApdex"
)

// DatabaseOverviewServiceClient is the client API for DatabaseOverviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseOverviewServiceClient interface {
	GetDatabaseOverviewStream(ctx context.Context, in *GetDatabaseOverviewStreamRequest, opts ...grpc.CallOption) (DatabaseOverviewService_GetDatabaseOverviewStreamClient, error)
	GetApdex(ctx context.Context, in *GetApdexRequest, opts ...grpc.CallOption) (*common.GetApdexResponse, error)
}

type databaseOverviewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseOverviewServiceClient(cc grpc.ClientConnInterface) DatabaseOverviewServiceClient {
	return &databaseOverviewServiceClient{cc}
}

func (c *databaseOverviewServiceClient) GetDatabaseOverviewStream(ctx context.Context, in *GetDatabaseOverviewStreamRequest, opts ...grpc.CallOption) (DatabaseOverviewService_GetDatabaseOverviewStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &DatabaseOverviewService_ServiceDesc.Streams[0], DatabaseOverviewService_GetDatabaseOverviewStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &databaseOverviewServiceGetDatabaseOverviewStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DatabaseOverviewService_GetDatabaseOverviewStreamClient interface {
	Recv() (*GetDatabaseOverviewStreamResponse, error)
	grpc.ClientStream
}

type databaseOverviewServiceGetDatabaseOverviewStreamClient struct {
	grpc.ClientStream
}

func (x *databaseOverviewServiceGetDatabaseOverviewStreamClient) Recv() (*GetDatabaseOverviewStreamResponse, error) {
	m := new(GetDatabaseOverviewStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *databaseOverviewServiceClient) GetApdex(ctx context.Context, in *GetApdexRequest, opts ...grpc.CallOption) (*common.GetApdexResponse, error) {
	out := new(common.GetApdexResponse)
	err := c.cc.Invoke(ctx, DatabaseOverviewService_GetApdex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseOverviewServiceServer is the server API for DatabaseOverviewService service.
// All implementations must embed UnimplementedDatabaseOverviewServiceServer
// for forward compatibility
type DatabaseOverviewServiceServer interface {
	GetDatabaseOverviewStream(*GetDatabaseOverviewStreamRequest, DatabaseOverviewService_GetDatabaseOverviewStreamServer) error
	GetApdex(context.Context, *GetApdexRequest) (*common.GetApdexResponse, error)
	mustEmbedUnimplementedDatabaseOverviewServiceServer()
}

// UnimplementedDatabaseOverviewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDatabaseOverviewServiceServer struct {
}

func (UnimplementedDatabaseOverviewServiceServer) GetDatabaseOverviewStream(*GetDatabaseOverviewStreamRequest, DatabaseOverviewService_GetDatabaseOverviewStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDatabaseOverviewStream not implemented")
}
func (UnimplementedDatabaseOverviewServiceServer) GetApdex(context.Context, *GetApdexRequest) (*common.GetApdexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApdex not implemented")
}
func (UnimplementedDatabaseOverviewServiceServer) mustEmbedUnimplementedDatabaseOverviewServiceServer() {
}

// UnsafeDatabaseOverviewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseOverviewServiceServer will
// result in compilation errors.
type UnsafeDatabaseOverviewServiceServer interface {
	mustEmbedUnimplementedDatabaseOverviewServiceServer()
}

func RegisterDatabaseOverviewServiceServer(s grpc.ServiceRegistrar, srv DatabaseOverviewServiceServer) {
	s.RegisterService(&DatabaseOverviewService_ServiceDesc, srv)
}

func _DatabaseOverviewService_GetDatabaseOverviewStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDatabaseOverviewStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DatabaseOverviewServiceServer).GetDatabaseOverviewStream(m, &databaseOverviewServiceGetDatabaseOverviewStreamServer{stream})
}

type DatabaseOverviewService_GetDatabaseOverviewStreamServer interface {
	Send(*GetDatabaseOverviewStreamResponse) error
	grpc.ServerStream
}

type databaseOverviewServiceGetDatabaseOverviewStreamServer struct {
	grpc.ServerStream
}

func (x *databaseOverviewServiceGetDatabaseOverviewStreamServer) Send(m *GetDatabaseOverviewStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DatabaseOverviewService_GetApdex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApdexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseOverviewServiceServer).GetApdex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseOverviewService_GetApdex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseOverviewServiceServer).GetApdex(ctx, req.(*GetApdexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatabaseOverviewService_ServiceDesc is the grpc.ServiceDesc for DatabaseOverviewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatabaseOverviewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.database_catalog.v1.DatabaseOverviewService",
	HandlerType: (*DatabaseOverviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApdex",
			Handler:    _DatabaseOverviewService_GetApdex_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDatabaseOverviewStream",
			Handler:       _DatabaseOverviewService_GetDatabaseOverviewStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "com/coralogixapis/database_catalog/v1/overview_service.proto",
}
