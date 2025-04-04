// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: com/coralogix/global_mapping/v1/data_source_service.proto

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
	CompanyDataSourcesService_GetCompanyDataSources_FullMethodName = "/com.coralogix.global_mapping.v1.CompanyDataSourcesService/GetCompanyDataSources"
)

// CompanyDataSourcesServiceClient is the client API for CompanyDataSourcesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyDataSourcesServiceClient interface {
	GetCompanyDataSources(ctx context.Context, in *GetCompanyDataSourcesRequest, opts ...grpc.CallOption) (*GetCompanyDataSourcesResponse, error)
}

type companyDataSourcesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyDataSourcesServiceClient(cc grpc.ClientConnInterface) CompanyDataSourcesServiceClient {
	return &companyDataSourcesServiceClient{cc}
}

func (c *companyDataSourcesServiceClient) GetCompanyDataSources(ctx context.Context, in *GetCompanyDataSourcesRequest, opts ...grpc.CallOption) (*GetCompanyDataSourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCompanyDataSourcesResponse)
	err := c.cc.Invoke(ctx, CompanyDataSourcesService_GetCompanyDataSources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyDataSourcesServiceServer is the server API for CompanyDataSourcesService service.
// All implementations must embed UnimplementedCompanyDataSourcesServiceServer
// for forward compatibility.
type CompanyDataSourcesServiceServer interface {
	GetCompanyDataSources(context.Context, *GetCompanyDataSourcesRequest) (*GetCompanyDataSourcesResponse, error)
	mustEmbedUnimplementedCompanyDataSourcesServiceServer()
}

// UnimplementedCompanyDataSourcesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCompanyDataSourcesServiceServer struct{}

func (UnimplementedCompanyDataSourcesServiceServer) GetCompanyDataSources(context.Context, *GetCompanyDataSourcesRequest) (*GetCompanyDataSourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyDataSources not implemented")
}
func (UnimplementedCompanyDataSourcesServiceServer) mustEmbedUnimplementedCompanyDataSourcesServiceServer() {
}
func (UnimplementedCompanyDataSourcesServiceServer) testEmbeddedByValue() {}

// UnsafeCompanyDataSourcesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyDataSourcesServiceServer will
// result in compilation errors.
type UnsafeCompanyDataSourcesServiceServer interface {
	mustEmbedUnimplementedCompanyDataSourcesServiceServer()
}

func RegisterCompanyDataSourcesServiceServer(s grpc.ServiceRegistrar, srv CompanyDataSourcesServiceServer) {
	// If the following call pancis, it indicates UnimplementedCompanyDataSourcesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CompanyDataSourcesService_ServiceDesc, srv)
}

func _CompanyDataSourcesService_GetCompanyDataSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyDataSourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyDataSourcesServiceServer).GetCompanyDataSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyDataSourcesService_GetCompanyDataSources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyDataSourcesServiceServer).GetCompanyDataSources(ctx, req.(*GetCompanyDataSourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyDataSourcesService_ServiceDesc is the grpc.ServiceDesc for CompanyDataSourcesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyDataSourcesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.global_mapping.v1.CompanyDataSourcesService",
	HandlerType: (*CompanyDataSourcesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompanyDataSources",
			Handler:    _CompanyDataSourcesService_GetCompanyDataSources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/global_mapping/v1/data_source_service.proto",
}
