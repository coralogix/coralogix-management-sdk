// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: com/coralogixapis/dataprime/v1/query_service.proto

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
	DataprimeQueryService_Query_FullMethodName = "/com.coralogixapis.dataprime.v1.DataprimeQueryService/Query"
)

// DataprimeQueryServiceClient is the client API for DataprimeQueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataprimeQueryServiceClient interface {
	// method to run dataprime text queries
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (DataprimeQueryService_QueryClient, error)
}

type dataprimeQueryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataprimeQueryServiceClient(cc grpc.ClientConnInterface) DataprimeQueryServiceClient {
	return &dataprimeQueryServiceClient{cc}
}

func (c *dataprimeQueryServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (DataprimeQueryService_QueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataprimeQueryService_ServiceDesc.Streams[0], DataprimeQueryService_Query_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &dataprimeQueryServiceQueryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataprimeQueryService_QueryClient interface {
	Recv() (*QueryResponse, error)
	grpc.ClientStream
}

type dataprimeQueryServiceQueryClient struct {
	grpc.ClientStream
}

func (x *dataprimeQueryServiceQueryClient) Recv() (*QueryResponse, error) {
	m := new(QueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataprimeQueryServiceServer is the server API for DataprimeQueryService service.
// All implementations must embed UnimplementedDataprimeQueryServiceServer
// for forward compatibility
type DataprimeQueryServiceServer interface {
	// method to run dataprime text queries
	Query(*QueryRequest, DataprimeQueryService_QueryServer) error
	mustEmbedUnimplementedDataprimeQueryServiceServer()
}

// UnimplementedDataprimeQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataprimeQueryServiceServer struct {
}

func (UnimplementedDataprimeQueryServiceServer) Query(*QueryRequest, DataprimeQueryService_QueryServer) error {
	return status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedDataprimeQueryServiceServer) mustEmbedUnimplementedDataprimeQueryServiceServer() {}

// UnsafeDataprimeQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataprimeQueryServiceServer will
// result in compilation errors.
type UnsafeDataprimeQueryServiceServer interface {
	mustEmbedUnimplementedDataprimeQueryServiceServer()
}

func RegisterDataprimeQueryServiceServer(s grpc.ServiceRegistrar, srv DataprimeQueryServiceServer) {
	s.RegisterService(&DataprimeQueryService_ServiceDesc, srv)
}

func _DataprimeQueryService_Query_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataprimeQueryServiceServer).Query(m, &dataprimeQueryServiceQueryServer{stream})
}

type DataprimeQueryService_QueryServer interface {
	Send(*QueryResponse) error
	grpc.ServerStream
}

type dataprimeQueryServiceQueryServer struct {
	grpc.ServerStream
}

func (x *dataprimeQueryServiceQueryServer) Send(m *QueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

// DataprimeQueryService_ServiceDesc is the grpc.ServiceDesc for DataprimeQueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataprimeQueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.dataprime.v1.DataprimeQueryService",
	HandlerType: (*DataprimeQueryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Query",
			Handler:       _DataprimeQueryService_Query_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "com/coralogixapis/dataprime/v1/query_service.proto",
}
