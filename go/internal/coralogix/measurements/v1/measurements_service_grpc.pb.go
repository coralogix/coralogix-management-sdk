// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.28.3
// source: com/coralogix/measurements/v1/measurements_service.proto

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
	MeasurementsService_GetQueries_FullMethodName           = "/com.coralogix.measurements.v1.MeasurementsService/GetQueries"
	MeasurementsService_GetMeasurementsTable_FullMethodName = "/com.coralogix.measurements.v1.MeasurementsService/GetMeasurementsTable"
	MeasurementsService_GetHierarchy_FullMethodName         = "/com.coralogix.measurements.v1.MeasurementsService/GetHierarchy"
)

// MeasurementsServiceClient is the client API for MeasurementsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeasurementsServiceClient interface {
	GetQueries(ctx context.Context, in *GetQueriesRequest, opts ...grpc.CallOption) (*GetQueriesResponse, error)
	GetMeasurementsTable(ctx context.Context, in *GetMeasurementsTableRequest, opts ...grpc.CallOption) (*GetMeasurementsTableResponse, error)
	GetHierarchy(ctx context.Context, in *GetHierarchyRequest, opts ...grpc.CallOption) (*GetHierarchyResponse, error)
}

type measurementsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeasurementsServiceClient(cc grpc.ClientConnInterface) MeasurementsServiceClient {
	return &measurementsServiceClient{cc}
}

func (c *measurementsServiceClient) GetQueries(ctx context.Context, in *GetQueriesRequest, opts ...grpc.CallOption) (*GetQueriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetQueriesResponse)
	err := c.cc.Invoke(ctx, MeasurementsService_GetQueries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measurementsServiceClient) GetMeasurementsTable(ctx context.Context, in *GetMeasurementsTableRequest, opts ...grpc.CallOption) (*GetMeasurementsTableResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMeasurementsTableResponse)
	err := c.cc.Invoke(ctx, MeasurementsService_GetMeasurementsTable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measurementsServiceClient) GetHierarchy(ctx context.Context, in *GetHierarchyRequest, opts ...grpc.CallOption) (*GetHierarchyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHierarchyResponse)
	err := c.cc.Invoke(ctx, MeasurementsService_GetHierarchy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeasurementsServiceServer is the server API for MeasurementsService service.
// All implementations must embed UnimplementedMeasurementsServiceServer
// for forward compatibility
type MeasurementsServiceServer interface {
	GetQueries(context.Context, *GetQueriesRequest) (*GetQueriesResponse, error)
	GetMeasurementsTable(context.Context, *GetMeasurementsTableRequest) (*GetMeasurementsTableResponse, error)
	GetHierarchy(context.Context, *GetHierarchyRequest) (*GetHierarchyResponse, error)
	mustEmbedUnimplementedMeasurementsServiceServer()
}

// UnimplementedMeasurementsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMeasurementsServiceServer struct {
}

func (UnimplementedMeasurementsServiceServer) GetQueries(context.Context, *GetQueriesRequest) (*GetQueriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueries not implemented")
}
func (UnimplementedMeasurementsServiceServer) GetMeasurementsTable(context.Context, *GetMeasurementsTableRequest) (*GetMeasurementsTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeasurementsTable not implemented")
}
func (UnimplementedMeasurementsServiceServer) GetHierarchy(context.Context, *GetHierarchyRequest) (*GetHierarchyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHierarchy not implemented")
}
func (UnimplementedMeasurementsServiceServer) mustEmbedUnimplementedMeasurementsServiceServer() {}

// UnsafeMeasurementsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeasurementsServiceServer will
// result in compilation errors.
type UnsafeMeasurementsServiceServer interface {
	mustEmbedUnimplementedMeasurementsServiceServer()
}

func RegisterMeasurementsServiceServer(s grpc.ServiceRegistrar, srv MeasurementsServiceServer) {
	s.RegisterService(&MeasurementsService_ServiceDesc, srv)
}

func _MeasurementsService_GetQueries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQueriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasurementsServiceServer).GetQueries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeasurementsService_GetQueries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasurementsServiceServer).GetQueries(ctx, req.(*GetQueriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasurementsService_GetMeasurementsTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeasurementsTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasurementsServiceServer).GetMeasurementsTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeasurementsService_GetMeasurementsTable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasurementsServiceServer).GetMeasurementsTable(ctx, req.(*GetMeasurementsTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasurementsService_GetHierarchy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHierarchyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasurementsServiceServer).GetHierarchy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeasurementsService_GetHierarchy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasurementsServiceServer).GetHierarchy(ctx, req.(*GetHierarchyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MeasurementsService_ServiceDesc is the grpc.ServiceDesc for MeasurementsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeasurementsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.measurements.v1.MeasurementsService",
	HandlerType: (*MeasurementsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQueries",
			Handler:    _MeasurementsService_GetQueries_Handler,
		},
		{
			MethodName: "GetMeasurementsTable",
			Handler:    _MeasurementsService_GetMeasurementsTable_Handler,
		},
		{
			MethodName: "GetHierarchy",
			Handler:    _MeasurementsService_GetHierarchy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/measurements/v1/measurements_service.proto",
}
