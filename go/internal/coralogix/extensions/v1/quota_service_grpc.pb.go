// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: com/coralogix/extensions/v1/quota_service.proto

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
	QuotaService_GetQuotas_FullMethodName = "/com.coralogix.extensions.v1.QuotaService/GetQuotas"
)

// QuotaServiceClient is the client API for QuotaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuotaServiceClient interface {
	GetQuotas(ctx context.Context, in *GetQuotasRequest, opts ...grpc.CallOption) (*GetQuotasResponse, error)
}

type quotaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuotaServiceClient(cc grpc.ClientConnInterface) QuotaServiceClient {
	return &quotaServiceClient{cc}
}

func (c *quotaServiceClient) GetQuotas(ctx context.Context, in *GetQuotasRequest, opts ...grpc.CallOption) (*GetQuotasResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetQuotasResponse)
	err := c.cc.Invoke(ctx, QuotaService_GetQuotas_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuotaServiceServer is the server API for QuotaService service.
// All implementations must embed UnimplementedQuotaServiceServer
// for forward compatibility.
type QuotaServiceServer interface {
	GetQuotas(context.Context, *GetQuotasRequest) (*GetQuotasResponse, error)
	mustEmbedUnimplementedQuotaServiceServer()
}

// UnimplementedQuotaServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQuotaServiceServer struct{}

func (UnimplementedQuotaServiceServer) GetQuotas(context.Context, *GetQuotasRequest) (*GetQuotasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuotas not implemented")
}
func (UnimplementedQuotaServiceServer) mustEmbedUnimplementedQuotaServiceServer() {}
func (UnimplementedQuotaServiceServer) testEmbeddedByValue()                      {}

// UnsafeQuotaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuotaServiceServer will
// result in compilation errors.
type UnsafeQuotaServiceServer interface {
	mustEmbedUnimplementedQuotaServiceServer()
}

func RegisterQuotaServiceServer(s grpc.ServiceRegistrar, srv QuotaServiceServer) {
	// If the following call pancis, it indicates UnimplementedQuotaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&QuotaService_ServiceDesc, srv)
}

func _QuotaService_GetQuotas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuotasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServiceServer).GetQuotas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuotaService_GetQuotas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServiceServer).GetQuotas(ctx, req.(*GetQuotasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuotaService_ServiceDesc is the grpc.ServiceDesc for QuotaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuotaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.extensions.v1.QuotaService",
	HandlerType: (*QuotaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuotas",
			Handler:    _QuotaService_GetQuotas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/extensions/v1/quota_service.proto",
}
