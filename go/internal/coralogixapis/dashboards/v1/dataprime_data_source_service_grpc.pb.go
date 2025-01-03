// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
// source: com/coralogixapis/dashboards/v1/services/dataprime_data_source_service.proto

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
	DataprimeDataSourceService_SearchDataprime_FullMethodName        = "/com.coralogixapis.dashboards.v1.services.DataprimeDataSourceService/SearchDataprime"
	DataprimeDataSourceService_SearchDataprimeArchive_FullMethodName = "/com.coralogixapis.dashboards.v1.services.DataprimeDataSourceService/SearchDataprimeArchive"
)

// DataprimeDataSourceServiceClient is the client API for DataprimeDataSourceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataprimeDataSourceServiceClient interface {
	SearchDataprime(ctx context.Context, in *SearchDataprimeRequest, opts ...grpc.CallOption) (*SearchDataprimeResponse, error)
	SearchDataprimeArchive(ctx context.Context, in *SearchDataprimeArchiveRequest, opts ...grpc.CallOption) (*SearchDataprimeArchiveResponse, error)
}

type dataprimeDataSourceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataprimeDataSourceServiceClient(cc grpc.ClientConnInterface) DataprimeDataSourceServiceClient {
	return &dataprimeDataSourceServiceClient{cc}
}

func (c *dataprimeDataSourceServiceClient) SearchDataprime(ctx context.Context, in *SearchDataprimeRequest, opts ...grpc.CallOption) (*SearchDataprimeResponse, error) {
	out := new(SearchDataprimeResponse)
	err := c.cc.Invoke(ctx, DataprimeDataSourceService_SearchDataprime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataprimeDataSourceServiceClient) SearchDataprimeArchive(ctx context.Context, in *SearchDataprimeArchiveRequest, opts ...grpc.CallOption) (*SearchDataprimeArchiveResponse, error) {
	out := new(SearchDataprimeArchiveResponse)
	err := c.cc.Invoke(ctx, DataprimeDataSourceService_SearchDataprimeArchive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataprimeDataSourceServiceServer is the server API for DataprimeDataSourceService service.
// All implementations must embed UnimplementedDataprimeDataSourceServiceServer
// for forward compatibility
type DataprimeDataSourceServiceServer interface {
	SearchDataprime(context.Context, *SearchDataprimeRequest) (*SearchDataprimeResponse, error)
	SearchDataprimeArchive(context.Context, *SearchDataprimeArchiveRequest) (*SearchDataprimeArchiveResponse, error)
	mustEmbedUnimplementedDataprimeDataSourceServiceServer()
}

// UnimplementedDataprimeDataSourceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataprimeDataSourceServiceServer struct {
}

func (UnimplementedDataprimeDataSourceServiceServer) SearchDataprime(context.Context, *SearchDataprimeRequest) (*SearchDataprimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchDataprime not implemented")
}
func (UnimplementedDataprimeDataSourceServiceServer) SearchDataprimeArchive(context.Context, *SearchDataprimeArchiveRequest) (*SearchDataprimeArchiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchDataprimeArchive not implemented")
}
func (UnimplementedDataprimeDataSourceServiceServer) mustEmbedUnimplementedDataprimeDataSourceServiceServer() {
}

// UnsafeDataprimeDataSourceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataprimeDataSourceServiceServer will
// result in compilation errors.
type UnsafeDataprimeDataSourceServiceServer interface {
	mustEmbedUnimplementedDataprimeDataSourceServiceServer()
}

func RegisterDataprimeDataSourceServiceServer(s grpc.ServiceRegistrar, srv DataprimeDataSourceServiceServer) {
	s.RegisterService(&DataprimeDataSourceService_ServiceDesc, srv)
}

func _DataprimeDataSourceService_SearchDataprime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchDataprimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataprimeDataSourceServiceServer).SearchDataprime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataprimeDataSourceService_SearchDataprime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataprimeDataSourceServiceServer).SearchDataprime(ctx, req.(*SearchDataprimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataprimeDataSourceService_SearchDataprimeArchive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchDataprimeArchiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataprimeDataSourceServiceServer).SearchDataprimeArchive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataprimeDataSourceService_SearchDataprimeArchive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataprimeDataSourceServiceServer).SearchDataprimeArchive(ctx, req.(*SearchDataprimeArchiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataprimeDataSourceService_ServiceDesc is the grpc.ServiceDesc for DataprimeDataSourceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataprimeDataSourceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.dashboards.v1.services.DataprimeDataSourceService",
	HandlerType: (*DataprimeDataSourceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchDataprime",
			Handler:    _DataprimeDataSourceService_SearchDataprime_Handler,
		},
		{
			MethodName: "SearchDataprimeArchive",
			Handler:    _DataprimeDataSourceService_SearchDataprimeArchive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/dashboards/v1/services/dataprime_data_source_service.proto",
}
