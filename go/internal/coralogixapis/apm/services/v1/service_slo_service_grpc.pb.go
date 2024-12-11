// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
// source: com/coralogixapis/apm/services/v1/service_slo_service.proto

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
	ServiceSloService_GetServiceSlo_FullMethodName       = "/com.coralogixapis.apm.services.v1.ServiceSloService/GetServiceSlo"
	ServiceSloService_CreateServiceSlo_FullMethodName    = "/com.coralogixapis.apm.services.v1.ServiceSloService/CreateServiceSlo"
	ServiceSloService_ReplaceServiceSlo_FullMethodName   = "/com.coralogixapis.apm.services.v1.ServiceSloService/ReplaceServiceSlo"
	ServiceSloService_DeleteServiceSlo_FullMethodName    = "/com.coralogixapis.apm.services.v1.ServiceSloService/DeleteServiceSlo"
	ServiceSloService_ListServiceSlos_FullMethodName     = "/com.coralogixapis.apm.services.v1.ServiceSloService/ListServiceSlos"
	ServiceSloService_BatchGetServiceSlos_FullMethodName = "/com.coralogixapis.apm.services.v1.ServiceSloService/BatchGetServiceSlos"
)

// ServiceSloServiceClient is the client API for ServiceSloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceSloServiceClient interface {
	GetServiceSlo(ctx context.Context, in *GetServiceSloRequest, opts ...grpc.CallOption) (*GetServiceSloResponse, error)
	CreateServiceSlo(ctx context.Context, in *CreateServiceSloRequest, opts ...grpc.CallOption) (*CreateServiceSloResponse, error)
	ReplaceServiceSlo(ctx context.Context, in *ReplaceServiceSloRequest, opts ...grpc.CallOption) (*ReplaceServiceSloResponse, error)
	DeleteServiceSlo(ctx context.Context, in *DeleteServiceSloRequest, opts ...grpc.CallOption) (*DeleteServiceSloResponse, error)
	ListServiceSlos(ctx context.Context, in *ListServiceSlosRequest, opts ...grpc.CallOption) (*ListServiceSlosResponse, error)
	BatchGetServiceSlos(ctx context.Context, in *BatchGetServiceSlosRequest, opts ...grpc.CallOption) (*BatchGetServiceSlosResponse, error)
}

type serviceSloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceSloServiceClient(cc grpc.ClientConnInterface) ServiceSloServiceClient {
	return &serviceSloServiceClient{cc}
}

func (c *serviceSloServiceClient) GetServiceSlo(ctx context.Context, in *GetServiceSloRequest, opts ...grpc.CallOption) (*GetServiceSloResponse, error) {
	out := new(GetServiceSloResponse)
	err := c.cc.Invoke(ctx, ServiceSloService_GetServiceSlo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceSloServiceClient) CreateServiceSlo(ctx context.Context, in *CreateServiceSloRequest, opts ...grpc.CallOption) (*CreateServiceSloResponse, error) {
	out := new(CreateServiceSloResponse)
	err := c.cc.Invoke(ctx, ServiceSloService_CreateServiceSlo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceSloServiceClient) ReplaceServiceSlo(ctx context.Context, in *ReplaceServiceSloRequest, opts ...grpc.CallOption) (*ReplaceServiceSloResponse, error) {
	out := new(ReplaceServiceSloResponse)
	err := c.cc.Invoke(ctx, ServiceSloService_ReplaceServiceSlo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceSloServiceClient) DeleteServiceSlo(ctx context.Context, in *DeleteServiceSloRequest, opts ...grpc.CallOption) (*DeleteServiceSloResponse, error) {
	out := new(DeleteServiceSloResponse)
	err := c.cc.Invoke(ctx, ServiceSloService_DeleteServiceSlo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceSloServiceClient) ListServiceSlos(ctx context.Context, in *ListServiceSlosRequest, opts ...grpc.CallOption) (*ListServiceSlosResponse, error) {
	out := new(ListServiceSlosResponse)
	err := c.cc.Invoke(ctx, ServiceSloService_ListServiceSlos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceSloServiceClient) BatchGetServiceSlos(ctx context.Context, in *BatchGetServiceSlosRequest, opts ...grpc.CallOption) (*BatchGetServiceSlosResponse, error) {
	out := new(BatchGetServiceSlosResponse)
	err := c.cc.Invoke(ctx, ServiceSloService_BatchGetServiceSlos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceSloServiceServer is the server API for ServiceSloService service.
// All implementations must embed UnimplementedServiceSloServiceServer
// for forward compatibility
type ServiceSloServiceServer interface {
	GetServiceSlo(context.Context, *GetServiceSloRequest) (*GetServiceSloResponse, error)
	CreateServiceSlo(context.Context, *CreateServiceSloRequest) (*CreateServiceSloResponse, error)
	ReplaceServiceSlo(context.Context, *ReplaceServiceSloRequest) (*ReplaceServiceSloResponse, error)
	DeleteServiceSlo(context.Context, *DeleteServiceSloRequest) (*DeleteServiceSloResponse, error)
	ListServiceSlos(context.Context, *ListServiceSlosRequest) (*ListServiceSlosResponse, error)
	BatchGetServiceSlos(context.Context, *BatchGetServiceSlosRequest) (*BatchGetServiceSlosResponse, error)
	mustEmbedUnimplementedServiceSloServiceServer()
}

// UnimplementedServiceSloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceSloServiceServer struct {
}

func (UnimplementedServiceSloServiceServer) GetServiceSlo(context.Context, *GetServiceSloRequest) (*GetServiceSloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceSlo not implemented")
}
func (UnimplementedServiceSloServiceServer) CreateServiceSlo(context.Context, *CreateServiceSloRequest) (*CreateServiceSloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServiceSlo not implemented")
}
func (UnimplementedServiceSloServiceServer) ReplaceServiceSlo(context.Context, *ReplaceServiceSloRequest) (*ReplaceServiceSloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceServiceSlo not implemented")
}
func (UnimplementedServiceSloServiceServer) DeleteServiceSlo(context.Context, *DeleteServiceSloRequest) (*DeleteServiceSloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServiceSlo not implemented")
}
func (UnimplementedServiceSloServiceServer) ListServiceSlos(context.Context, *ListServiceSlosRequest) (*ListServiceSlosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServiceSlos not implemented")
}
func (UnimplementedServiceSloServiceServer) BatchGetServiceSlos(context.Context, *BatchGetServiceSlosRequest) (*BatchGetServiceSlosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetServiceSlos not implemented")
}
func (UnimplementedServiceSloServiceServer) mustEmbedUnimplementedServiceSloServiceServer() {}

// UnsafeServiceSloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceSloServiceServer will
// result in compilation errors.
type UnsafeServiceSloServiceServer interface {
	mustEmbedUnimplementedServiceSloServiceServer()
}

func RegisterServiceSloServiceServer(s grpc.ServiceRegistrar, srv ServiceSloServiceServer) {
	s.RegisterService(&ServiceSloService_ServiceDesc, srv)
}

func _ServiceSloService_GetServiceSlo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceSloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceSloServiceServer).GetServiceSlo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceSloService_GetServiceSlo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceSloServiceServer).GetServiceSlo(ctx, req.(*GetServiceSloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceSloService_CreateServiceSlo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceSloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceSloServiceServer).CreateServiceSlo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceSloService_CreateServiceSlo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceSloServiceServer).CreateServiceSlo(ctx, req.(*CreateServiceSloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceSloService_ReplaceServiceSlo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceServiceSloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceSloServiceServer).ReplaceServiceSlo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceSloService_ReplaceServiceSlo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceSloServiceServer).ReplaceServiceSlo(ctx, req.(*ReplaceServiceSloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceSloService_DeleteServiceSlo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceSloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceSloServiceServer).DeleteServiceSlo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceSloService_DeleteServiceSlo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceSloServiceServer).DeleteServiceSlo(ctx, req.(*DeleteServiceSloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceSloService_ListServiceSlos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServiceSlosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceSloServiceServer).ListServiceSlos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceSloService_ListServiceSlos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceSloServiceServer).ListServiceSlos(ctx, req.(*ListServiceSlosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceSloService_BatchGetServiceSlos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetServiceSlosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceSloServiceServer).BatchGetServiceSlos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceSloService_BatchGetServiceSlos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceSloServiceServer).BatchGetServiceSlos(ctx, req.(*BatchGetServiceSlosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceSloService_ServiceDesc is the grpc.ServiceDesc for ServiceSloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceSloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.apm.services.v1.ServiceSloService",
	HandlerType: (*ServiceSloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServiceSlo",
			Handler:    _ServiceSloService_GetServiceSlo_Handler,
		},
		{
			MethodName: "CreateServiceSlo",
			Handler:    _ServiceSloService_CreateServiceSlo_Handler,
		},
		{
			MethodName: "ReplaceServiceSlo",
			Handler:    _ServiceSloService_ReplaceServiceSlo_Handler,
		},
		{
			MethodName: "DeleteServiceSlo",
			Handler:    _ServiceSloService_DeleteServiceSlo_Handler,
		},
		{
			MethodName: "ListServiceSlos",
			Handler:    _ServiceSloService_ListServiceSlos_Handler,
		},
		{
			MethodName: "BatchGetServiceSlos",
			Handler:    _ServiceSloService_BatchGetServiceSlos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/apm/services/v1/service_slo_service.proto",
}
