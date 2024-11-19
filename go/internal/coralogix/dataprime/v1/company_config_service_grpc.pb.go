// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: com/coralogix/dataprime/v1/company_config_service.proto

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
	CompanyConfigService_GetCompanyConfig_FullMethodName = "/com.coralogix.dataprime.v1.CompanyConfigService/GetCompanyConfig"
)

// CompanyConfigServiceClient is the client API for CompanyConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyConfigServiceClient interface {
	GetCompanyConfig(ctx context.Context, in *GetCompanyConfigRequest, opts ...grpc.CallOption) (*GetCompanyConfigResponse, error)
}

type companyConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyConfigServiceClient(cc grpc.ClientConnInterface) CompanyConfigServiceClient {
	return &companyConfigServiceClient{cc}
}

func (c *companyConfigServiceClient) GetCompanyConfig(ctx context.Context, in *GetCompanyConfigRequest, opts ...grpc.CallOption) (*GetCompanyConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCompanyConfigResponse)
	err := c.cc.Invoke(ctx, CompanyConfigService_GetCompanyConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyConfigServiceServer is the server API for CompanyConfigService service.
// All implementations must embed UnimplementedCompanyConfigServiceServer
// for forward compatibility.
type CompanyConfigServiceServer interface {
	GetCompanyConfig(context.Context, *GetCompanyConfigRequest) (*GetCompanyConfigResponse, error)
	mustEmbedUnimplementedCompanyConfigServiceServer()
}

// UnimplementedCompanyConfigServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCompanyConfigServiceServer struct{}

func (UnimplementedCompanyConfigServiceServer) GetCompanyConfig(context.Context, *GetCompanyConfigRequest) (*GetCompanyConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyConfig not implemented")
}
func (UnimplementedCompanyConfigServiceServer) mustEmbedUnimplementedCompanyConfigServiceServer() {}
func (UnimplementedCompanyConfigServiceServer) testEmbeddedByValue()                              {}

// UnsafeCompanyConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyConfigServiceServer will
// result in compilation errors.
type UnsafeCompanyConfigServiceServer interface {
	mustEmbedUnimplementedCompanyConfigServiceServer()
}

func RegisterCompanyConfigServiceServer(s grpc.ServiceRegistrar, srv CompanyConfigServiceServer) {
	// If the following call pancis, it indicates UnimplementedCompanyConfigServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CompanyConfigService_ServiceDesc, srv)
}

func _CompanyConfigService_GetCompanyConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyConfigServiceServer).GetCompanyConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyConfigService_GetCompanyConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyConfigServiceServer).GetCompanyConfig(ctx, req.(*GetCompanyConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyConfigService_ServiceDesc is the grpc.ServiceDesc for CompanyConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.dataprime.v1.CompanyConfigService",
	HandlerType: (*CompanyConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompanyConfig",
			Handler:    _CompanyConfigService_GetCompanyConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/dataprime/v1/company_config_service.proto",
}
