// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: com/coralogix/global_mapping/v1/measurements_internal_service.proto

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
	MeasurementsInternalService_UpsertGlobalMeasurementsInternal_FullMethodName = "/com.coralogix.global_mapping.v1.MeasurementsInternalService/UpsertGlobalMeasurementsInternal"
	MeasurementsInternalService_RemoveGlobalMeasurementsInternal_FullMethodName = "/com.coralogix.global_mapping.v1.MeasurementsInternalService/RemoveGlobalMeasurementsInternal"
	MeasurementsInternalService_UpsertCompanyProvidersInternal_FullMethodName   = "/com.coralogix.global_mapping.v1.MeasurementsInternalService/UpsertCompanyProvidersInternal"
	MeasurementsInternalService_GetMeasurementsInternal_FullMethodName          = "/com.coralogix.global_mapping.v1.MeasurementsInternalService/GetMeasurementsInternal"
)

// MeasurementsInternalServiceClient is the client API for MeasurementsInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeasurementsInternalServiceClient interface {
	UpsertGlobalMeasurementsInternal(ctx context.Context, in *UpsertGlobalMeasurementsInternalRequest, opts ...grpc.CallOption) (*UpsertGlobalMeasurementsInternalResponse, error)
	RemoveGlobalMeasurementsInternal(ctx context.Context, in *RemoveGlobalMeasurementsInternalRequest, opts ...grpc.CallOption) (*RemoveGlobalMeasurementsInternalResponse, error)
	UpsertCompanyProvidersInternal(ctx context.Context, in *UpsertCompanyProvidersInternalRequest, opts ...grpc.CallOption) (*UpsertCompanyProvidersInternalResponse, error)
	GetMeasurementsInternal(ctx context.Context, in *GetMeasurementsInternalRequest, opts ...grpc.CallOption) (*GetMeasurementsInternalResponse, error)
}

type measurementsInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeasurementsInternalServiceClient(cc grpc.ClientConnInterface) MeasurementsInternalServiceClient {
	return &measurementsInternalServiceClient{cc}
}

func (c *measurementsInternalServiceClient) UpsertGlobalMeasurementsInternal(ctx context.Context, in *UpsertGlobalMeasurementsInternalRequest, opts ...grpc.CallOption) (*UpsertGlobalMeasurementsInternalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpsertGlobalMeasurementsInternalResponse)
	err := c.cc.Invoke(ctx, MeasurementsInternalService_UpsertGlobalMeasurementsInternal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measurementsInternalServiceClient) RemoveGlobalMeasurementsInternal(ctx context.Context, in *RemoveGlobalMeasurementsInternalRequest, opts ...grpc.CallOption) (*RemoveGlobalMeasurementsInternalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveGlobalMeasurementsInternalResponse)
	err := c.cc.Invoke(ctx, MeasurementsInternalService_RemoveGlobalMeasurementsInternal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measurementsInternalServiceClient) UpsertCompanyProvidersInternal(ctx context.Context, in *UpsertCompanyProvidersInternalRequest, opts ...grpc.CallOption) (*UpsertCompanyProvidersInternalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpsertCompanyProvidersInternalResponse)
	err := c.cc.Invoke(ctx, MeasurementsInternalService_UpsertCompanyProvidersInternal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measurementsInternalServiceClient) GetMeasurementsInternal(ctx context.Context, in *GetMeasurementsInternalRequest, opts ...grpc.CallOption) (*GetMeasurementsInternalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMeasurementsInternalResponse)
	err := c.cc.Invoke(ctx, MeasurementsInternalService_GetMeasurementsInternal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeasurementsInternalServiceServer is the server API for MeasurementsInternalService service.
// All implementations must embed UnimplementedMeasurementsInternalServiceServer
// for forward compatibility.
type MeasurementsInternalServiceServer interface {
	UpsertGlobalMeasurementsInternal(context.Context, *UpsertGlobalMeasurementsInternalRequest) (*UpsertGlobalMeasurementsInternalResponse, error)
	RemoveGlobalMeasurementsInternal(context.Context, *RemoveGlobalMeasurementsInternalRequest) (*RemoveGlobalMeasurementsInternalResponse, error)
	UpsertCompanyProvidersInternal(context.Context, *UpsertCompanyProvidersInternalRequest) (*UpsertCompanyProvidersInternalResponse, error)
	GetMeasurementsInternal(context.Context, *GetMeasurementsInternalRequest) (*GetMeasurementsInternalResponse, error)
	mustEmbedUnimplementedMeasurementsInternalServiceServer()
}

// UnimplementedMeasurementsInternalServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMeasurementsInternalServiceServer struct{}

func (UnimplementedMeasurementsInternalServiceServer) UpsertGlobalMeasurementsInternal(context.Context, *UpsertGlobalMeasurementsInternalRequest) (*UpsertGlobalMeasurementsInternalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertGlobalMeasurementsInternal not implemented")
}
func (UnimplementedMeasurementsInternalServiceServer) RemoveGlobalMeasurementsInternal(context.Context, *RemoveGlobalMeasurementsInternalRequest) (*RemoveGlobalMeasurementsInternalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGlobalMeasurementsInternal not implemented")
}
func (UnimplementedMeasurementsInternalServiceServer) UpsertCompanyProvidersInternal(context.Context, *UpsertCompanyProvidersInternalRequest) (*UpsertCompanyProvidersInternalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertCompanyProvidersInternal not implemented")
}
func (UnimplementedMeasurementsInternalServiceServer) GetMeasurementsInternal(context.Context, *GetMeasurementsInternalRequest) (*GetMeasurementsInternalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeasurementsInternal not implemented")
}
func (UnimplementedMeasurementsInternalServiceServer) mustEmbedUnimplementedMeasurementsInternalServiceServer() {
}
func (UnimplementedMeasurementsInternalServiceServer) testEmbeddedByValue() {}

// UnsafeMeasurementsInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeasurementsInternalServiceServer will
// result in compilation errors.
type UnsafeMeasurementsInternalServiceServer interface {
	mustEmbedUnimplementedMeasurementsInternalServiceServer()
}

func RegisterMeasurementsInternalServiceServer(s grpc.ServiceRegistrar, srv MeasurementsInternalServiceServer) {
	// If the following call pancis, it indicates UnimplementedMeasurementsInternalServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MeasurementsInternalService_ServiceDesc, srv)
}

func _MeasurementsInternalService_UpsertGlobalMeasurementsInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertGlobalMeasurementsInternalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasurementsInternalServiceServer).UpsertGlobalMeasurementsInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeasurementsInternalService_UpsertGlobalMeasurementsInternal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasurementsInternalServiceServer).UpsertGlobalMeasurementsInternal(ctx, req.(*UpsertGlobalMeasurementsInternalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasurementsInternalService_RemoveGlobalMeasurementsInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGlobalMeasurementsInternalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasurementsInternalServiceServer).RemoveGlobalMeasurementsInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeasurementsInternalService_RemoveGlobalMeasurementsInternal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasurementsInternalServiceServer).RemoveGlobalMeasurementsInternal(ctx, req.(*RemoveGlobalMeasurementsInternalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasurementsInternalService_UpsertCompanyProvidersInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertCompanyProvidersInternalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasurementsInternalServiceServer).UpsertCompanyProvidersInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeasurementsInternalService_UpsertCompanyProvidersInternal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasurementsInternalServiceServer).UpsertCompanyProvidersInternal(ctx, req.(*UpsertCompanyProvidersInternalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasurementsInternalService_GetMeasurementsInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeasurementsInternalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasurementsInternalServiceServer).GetMeasurementsInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeasurementsInternalService_GetMeasurementsInternal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasurementsInternalServiceServer).GetMeasurementsInternal(ctx, req.(*GetMeasurementsInternalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MeasurementsInternalService_ServiceDesc is the grpc.ServiceDesc for MeasurementsInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeasurementsInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.global_mapping.v1.MeasurementsInternalService",
	HandlerType: (*MeasurementsInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertGlobalMeasurementsInternal",
			Handler:    _MeasurementsInternalService_UpsertGlobalMeasurementsInternal_Handler,
		},
		{
			MethodName: "RemoveGlobalMeasurementsInternal",
			Handler:    _MeasurementsInternalService_RemoveGlobalMeasurementsInternal_Handler,
		},
		{
			MethodName: "UpsertCompanyProvidersInternal",
			Handler:    _MeasurementsInternalService_UpsertCompanyProvidersInternal_Handler,
		},
		{
			MethodName: "GetMeasurementsInternal",
			Handler:    _MeasurementsInternalService_GetMeasurementsInternal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/global_mapping/v1/measurements_internal_service.proto",
}
