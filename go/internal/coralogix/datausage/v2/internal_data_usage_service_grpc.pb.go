// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: com/coralogix/datausage/v2/internal_data_usage_service.proto

package v2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	InternalDataUsageService_ExportTeamsDetailedUsage_FullMethodName = "/com.coralogix.datausage.v2.InternalDataUsageService/ExportTeamsDetailedUsage"
)

// InternalDataUsageServiceClient is the client API for InternalDataUsageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalDataUsageServiceClient interface {
	// / Streams entries containing detailed data usage for all teams given a date range and resolution interval.
	// / Results are aggregated by pillar and priority
	// / Pillars i.e. [Logs, Metrics, Tracing, ... ]
	// / Priority i.e. [High, Medium, Low, ... ]
	ExportTeamsDetailedUsage(ctx context.Context, in *ExportTeamsDetailedUsageRequest, opts ...grpc.CallOption) (InternalDataUsageService_ExportTeamsDetailedUsageClient, error)
}

type internalDataUsageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalDataUsageServiceClient(cc grpc.ClientConnInterface) InternalDataUsageServiceClient {
	return &internalDataUsageServiceClient{cc}
}

func (c *internalDataUsageServiceClient) ExportTeamsDetailedUsage(ctx context.Context, in *ExportTeamsDetailedUsageRequest, opts ...grpc.CallOption) (InternalDataUsageService_ExportTeamsDetailedUsageClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &InternalDataUsageService_ServiceDesc.Streams[0], InternalDataUsageService_ExportTeamsDetailedUsage_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &internalDataUsageServiceExportTeamsDetailedUsageClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InternalDataUsageService_ExportTeamsDetailedUsageClient interface {
	Recv() (*ExportTeamsDetailedUsageResponse, error)
	grpc.ClientStream
}

type internalDataUsageServiceExportTeamsDetailedUsageClient struct {
	grpc.ClientStream
}

func (x *internalDataUsageServiceExportTeamsDetailedUsageClient) Recv() (*ExportTeamsDetailedUsageResponse, error) {
	m := new(ExportTeamsDetailedUsageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InternalDataUsageServiceServer is the server API for InternalDataUsageService service.
// All implementations must embed UnimplementedInternalDataUsageServiceServer
// for forward compatibility
type InternalDataUsageServiceServer interface {
	// / Streams entries containing detailed data usage for all teams given a date range and resolution interval.
	// / Results are aggregated by pillar and priority
	// / Pillars i.e. [Logs, Metrics, Tracing, ... ]
	// / Priority i.e. [High, Medium, Low, ... ]
	ExportTeamsDetailedUsage(*ExportTeamsDetailedUsageRequest, InternalDataUsageService_ExportTeamsDetailedUsageServer) error
	mustEmbedUnimplementedInternalDataUsageServiceServer()
}

// UnimplementedInternalDataUsageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInternalDataUsageServiceServer struct {
}

func (UnimplementedInternalDataUsageServiceServer) ExportTeamsDetailedUsage(*ExportTeamsDetailedUsageRequest, InternalDataUsageService_ExportTeamsDetailedUsageServer) error {
	return status.Errorf(codes.Unimplemented, "method ExportTeamsDetailedUsage not implemented")
}
func (UnimplementedInternalDataUsageServiceServer) mustEmbedUnimplementedInternalDataUsageServiceServer() {
}

// UnsafeInternalDataUsageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternalDataUsageServiceServer will
// result in compilation errors.
type UnsafeInternalDataUsageServiceServer interface {
	mustEmbedUnimplementedInternalDataUsageServiceServer()
}

func RegisterInternalDataUsageServiceServer(s grpc.ServiceRegistrar, srv InternalDataUsageServiceServer) {
	s.RegisterService(&InternalDataUsageService_ServiceDesc, srv)
}

func _InternalDataUsageService_ExportTeamsDetailedUsage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExportTeamsDetailedUsageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InternalDataUsageServiceServer).ExportTeamsDetailedUsage(m, &internalDataUsageServiceExportTeamsDetailedUsageServer{ServerStream: stream})
}

type InternalDataUsageService_ExportTeamsDetailedUsageServer interface {
	Send(*ExportTeamsDetailedUsageResponse) error
	grpc.ServerStream
}

type internalDataUsageServiceExportTeamsDetailedUsageServer struct {
	grpc.ServerStream
}

func (x *internalDataUsageServiceExportTeamsDetailedUsageServer) Send(m *ExportTeamsDetailedUsageResponse) error {
	return x.ServerStream.SendMsg(m)
}

// InternalDataUsageService_ServiceDesc is the grpc.ServiceDesc for InternalDataUsageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InternalDataUsageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.datausage.v2.InternalDataUsageService",
	HandlerType: (*InternalDataUsageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExportTeamsDetailedUsage",
			Handler:       _InternalDataUsageService_ExportTeamsDetailedUsage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "com/coralogix/datausage/v2/internal_data_usage_service.proto",
}
