// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.1
// source: com/coralogix/global_mapping/v1/measurements_service.proto

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
	MeasurementsService_GetMeasurements_FullMethodName        = "/com.coralogix.global_mapping.v1.MeasurementsService/GetMeasurements"
	MeasurementsService_UpsertCompanyProviders_FullMethodName = "/com.coralogix.global_mapping.v1.MeasurementsService/UpsertCompanyProviders"
	MeasurementsService_GetQueries_FullMethodName             = "/com.coralogix.global_mapping.v1.MeasurementsService/GetQueries"
	MeasurementsService_GetMeasurementsTable_FullMethodName   = "/com.coralogix.global_mapping.v1.MeasurementsService/GetMeasurementsTable"
)

// MeasurementsServiceClient is the client API for MeasurementsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeasurementsServiceClient interface {
	GetMeasurements(ctx context.Context, in *GetMeasurementsRequest, opts ...grpc.CallOption) (*GetMeasurementsResponse, error)
	UpsertCompanyProviders(ctx context.Context, in *UpsertCompanyProvidersRequest, opts ...grpc.CallOption) (*UpsertCompanyProvidersResponse, error)
	GetQueries(ctx context.Context, in *GetQueriesRequest, opts ...grpc.CallOption) (*GetQueriesResponse, error)
	GetMeasurementsTable(ctx context.Context, in *GetMeasurementsTableRequest, opts ...grpc.CallOption) (*GetMeasurementsTableResponse, error)
}

type measurementsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeasurementsServiceClient(cc grpc.ClientConnInterface) MeasurementsServiceClient {
	return &measurementsServiceClient{cc}
}

func (c *measurementsServiceClient) GetMeasurements(ctx context.Context, in *GetMeasurementsRequest, opts ...grpc.CallOption) (*GetMeasurementsResponse, error) {
	out := new(GetMeasurementsResponse)
	err := c.cc.Invoke(ctx, MeasurementsService_GetMeasurements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measurementsServiceClient) UpsertCompanyProviders(ctx context.Context, in *UpsertCompanyProvidersRequest, opts ...grpc.CallOption) (*UpsertCompanyProvidersResponse, error) {
	out := new(UpsertCompanyProvidersResponse)
	err := c.cc.Invoke(ctx, MeasurementsService_UpsertCompanyProviders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measurementsServiceClient) GetQueries(ctx context.Context, in *GetQueriesRequest, opts ...grpc.CallOption) (*GetQueriesResponse, error) {
	out := new(GetQueriesResponse)
	err := c.cc.Invoke(ctx, MeasurementsService_GetQueries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measurementsServiceClient) GetMeasurementsTable(ctx context.Context, in *GetMeasurementsTableRequest, opts ...grpc.CallOption) (*GetMeasurementsTableResponse, error) {
	out := new(GetMeasurementsTableResponse)
	err := c.cc.Invoke(ctx, MeasurementsService_GetMeasurementsTable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeasurementsServiceServer is the server API for MeasurementsService service.
// All implementations must embed UnimplementedMeasurementsServiceServer
// for forward compatibility
type MeasurementsServiceServer interface {
	GetMeasurements(context.Context, *GetMeasurementsRequest) (*GetMeasurementsResponse, error)
	UpsertCompanyProviders(context.Context, *UpsertCompanyProvidersRequest) (*UpsertCompanyProvidersResponse, error)
	GetQueries(context.Context, *GetQueriesRequest) (*GetQueriesResponse, error)
	GetMeasurementsTable(context.Context, *GetMeasurementsTableRequest) (*GetMeasurementsTableResponse, error)
	mustEmbedUnimplementedMeasurementsServiceServer()
}

// UnimplementedMeasurementsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMeasurementsServiceServer struct {
}

func (UnimplementedMeasurementsServiceServer) GetMeasurements(context.Context, *GetMeasurementsRequest) (*GetMeasurementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeasurements not implemented")
}
func (UnimplementedMeasurementsServiceServer) UpsertCompanyProviders(context.Context, *UpsertCompanyProvidersRequest) (*UpsertCompanyProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertCompanyProviders not implemented")
}
func (UnimplementedMeasurementsServiceServer) GetQueries(context.Context, *GetQueriesRequest) (*GetQueriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueries not implemented")
}
func (UnimplementedMeasurementsServiceServer) GetMeasurementsTable(context.Context, *GetMeasurementsTableRequest) (*GetMeasurementsTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeasurementsTable not implemented")
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

func _MeasurementsService_GetMeasurements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeasurementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasurementsServiceServer).GetMeasurements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeasurementsService_GetMeasurements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasurementsServiceServer).GetMeasurements(ctx, req.(*GetMeasurementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasurementsService_UpsertCompanyProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertCompanyProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasurementsServiceServer).UpsertCompanyProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeasurementsService_UpsertCompanyProviders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasurementsServiceServer).UpsertCompanyProviders(ctx, req.(*UpsertCompanyProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
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

// MeasurementsService_ServiceDesc is the grpc.ServiceDesc for MeasurementsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeasurementsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.global_mapping.v1.MeasurementsService",
	HandlerType: (*MeasurementsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMeasurements",
			Handler:    _MeasurementsService_GetMeasurements_Handler,
		},
		{
			MethodName: "UpsertCompanyProviders",
			Handler:    _MeasurementsService_UpsertCompanyProviders_Handler,
		},
		{
			MethodName: "GetQueries",
			Handler:    _MeasurementsService_GetQueries_Handler,
		},
		{
			MethodName: "GetMeasurementsTable",
			Handler:    _MeasurementsService_GetMeasurementsTable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/global_mapping/v1/measurements_service.proto",
}
