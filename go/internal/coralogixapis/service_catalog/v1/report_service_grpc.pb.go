// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.28.2
// source: com/coralogixapis/service_catalog/v1/report_service.proto

package v1

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
	ReportService_GetReport_FullMethodName = "/com.coralogixapis.service_catalog.v1.ReportService/GetReport"
)

// ReportServiceClient is the client API for ReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportServiceClient interface {
	GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*GetReportResponse, error)
}

type reportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportServiceClient(cc grpc.ClientConnInterface) ReportServiceClient {
	return &reportServiceClient{cc}
}

func (c *reportServiceClient) GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*GetReportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReportResponse)
	err := c.cc.Invoke(ctx, ReportService_GetReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportServiceServer is the server API for ReportService service.
// All implementations must embed UnimplementedReportServiceServer
// for forward compatibility
type ReportServiceServer interface {
	GetReport(context.Context, *GetReportRequest) (*GetReportResponse, error)
	mustEmbedUnimplementedReportServiceServer()
}

// UnimplementedReportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReportServiceServer struct {
}

func (UnimplementedReportServiceServer) GetReport(context.Context, *GetReportRequest) (*GetReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}
func (UnimplementedReportServiceServer) mustEmbedUnimplementedReportServiceServer() {}

// UnsafeReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportServiceServer will
// result in compilation errors.
type UnsafeReportServiceServer interface {
	mustEmbedUnimplementedReportServiceServer()
}

func RegisterReportServiceServer(s grpc.ServiceRegistrar, srv ReportServiceServer) {
	s.RegisterService(&ReportService_ServiceDesc, srv)
}

func _ReportService_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportService_GetReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).GetReport(ctx, req.(*GetReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportService_ServiceDesc is the grpc.ServiceDesc for ReportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.service_catalog.v1.ReportService",
	HandlerType: (*ReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReport",
			Handler:    _ReportService_GetReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/service_catalog/v1/report_service.proto",
}
