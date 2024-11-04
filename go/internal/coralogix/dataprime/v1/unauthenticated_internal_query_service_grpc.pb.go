// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: com/coralogix/dataprime/v1/unauthenticated_internal_query_service.proto

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
	UnauthenticatedInternalDataprimeQueryService_Query_FullMethodName = "/com.coralogix.dataprime.v1.UnauthenticatedInternalDataprimeQueryService/Query"
)

// UnauthenticatedInternalDataprimeQueryServiceClient is the client API for UnauthenticatedInternalDataprimeQueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UnauthenticatedInternalDataprimeQueryServiceClient interface {
	// method to run dataprime queries
	Query(ctx context.Context, in *UnauthenticatedInternalDataprimeQueryServiceQueryRequest, opts ...grpc.CallOption) (UnauthenticatedInternalDataprimeQueryService_QueryClient, error)
}

type unauthenticatedInternalDataprimeQueryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUnauthenticatedInternalDataprimeQueryServiceClient(cc grpc.ClientConnInterface) UnauthenticatedInternalDataprimeQueryServiceClient {
	return &unauthenticatedInternalDataprimeQueryServiceClient{cc}
}

func (c *unauthenticatedInternalDataprimeQueryServiceClient) Query(ctx context.Context, in *UnauthenticatedInternalDataprimeQueryServiceQueryRequest, opts ...grpc.CallOption) (UnauthenticatedInternalDataprimeQueryService_QueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &UnauthenticatedInternalDataprimeQueryService_ServiceDesc.Streams[0], UnauthenticatedInternalDataprimeQueryService_Query_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &unauthenticatedInternalDataprimeQueryServiceQueryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UnauthenticatedInternalDataprimeQueryService_QueryClient interface {
	Recv() (*UnauthenticatedInternalDataprimeQueryServiceQueryResponse, error)
	grpc.ClientStream
}

type unauthenticatedInternalDataprimeQueryServiceQueryClient struct {
	grpc.ClientStream
}

func (x *unauthenticatedInternalDataprimeQueryServiceQueryClient) Recv() (*UnauthenticatedInternalDataprimeQueryServiceQueryResponse, error) {
	m := new(UnauthenticatedInternalDataprimeQueryServiceQueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UnauthenticatedInternalDataprimeQueryServiceServer is the server API for UnauthenticatedInternalDataprimeQueryService service.
// All implementations must embed UnimplementedUnauthenticatedInternalDataprimeQueryServiceServer
// for forward compatibility
type UnauthenticatedInternalDataprimeQueryServiceServer interface {
	// method to run dataprime queries
	Query(*UnauthenticatedInternalDataprimeQueryServiceQueryRequest, UnauthenticatedInternalDataprimeQueryService_QueryServer) error
	mustEmbedUnimplementedUnauthenticatedInternalDataprimeQueryServiceServer()
}

// UnimplementedUnauthenticatedInternalDataprimeQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUnauthenticatedInternalDataprimeQueryServiceServer struct {
}

func (UnimplementedUnauthenticatedInternalDataprimeQueryServiceServer) Query(*UnauthenticatedInternalDataprimeQueryServiceQueryRequest, UnauthenticatedInternalDataprimeQueryService_QueryServer) error {
	return status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedUnauthenticatedInternalDataprimeQueryServiceServer) mustEmbedUnimplementedUnauthenticatedInternalDataprimeQueryServiceServer() {
}

// UnsafeUnauthenticatedInternalDataprimeQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UnauthenticatedInternalDataprimeQueryServiceServer will
// result in compilation errors.
type UnsafeUnauthenticatedInternalDataprimeQueryServiceServer interface {
	mustEmbedUnimplementedUnauthenticatedInternalDataprimeQueryServiceServer()
}

func RegisterUnauthenticatedInternalDataprimeQueryServiceServer(s grpc.ServiceRegistrar, srv UnauthenticatedInternalDataprimeQueryServiceServer) {
	s.RegisterService(&UnauthenticatedInternalDataprimeQueryService_ServiceDesc, srv)
}

func _UnauthenticatedInternalDataprimeQueryService_Query_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UnauthenticatedInternalDataprimeQueryServiceQueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UnauthenticatedInternalDataprimeQueryServiceServer).Query(m, &unauthenticatedInternalDataprimeQueryServiceQueryServer{stream})
}

type UnauthenticatedInternalDataprimeQueryService_QueryServer interface {
	Send(*UnauthenticatedInternalDataprimeQueryServiceQueryResponse) error
	grpc.ServerStream
}

type unauthenticatedInternalDataprimeQueryServiceQueryServer struct {
	grpc.ServerStream
}

func (x *unauthenticatedInternalDataprimeQueryServiceQueryServer) Send(m *UnauthenticatedInternalDataprimeQueryServiceQueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

// UnauthenticatedInternalDataprimeQueryService_ServiceDesc is the grpc.ServiceDesc for UnauthenticatedInternalDataprimeQueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UnauthenticatedInternalDataprimeQueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.dataprime.v1.UnauthenticatedInternalDataprimeQueryService",
	HandlerType: (*UnauthenticatedInternalDataprimeQueryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Query",
			Handler:       _UnauthenticatedInternalDataprimeQueryService_Query_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "com/coralogix/dataprime/v1/unauthenticated_internal_query_service.proto",
}
