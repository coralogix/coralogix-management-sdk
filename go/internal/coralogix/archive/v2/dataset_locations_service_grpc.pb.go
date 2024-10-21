// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.2
// source: com/coralogix/archive/dataset/v2/dataset_locations_service.proto

package v2

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
	DatasetLocationsService_GetDatasetLocations_FullMethodName = "/com.coralogix.archive.dataset.v2.DatasetLocationsService/GetDatasetLocations"
	DatasetLocationsService_GetDatasetList_FullMethodName      = "/com.coralogix.archive.dataset.v2.DatasetLocationsService/GetDatasetList"
)

// DatasetLocationsServiceClient is the client API for DatasetLocationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatasetLocationsServiceClient interface {
	GetDatasetLocations(ctx context.Context, in *GetDatasetLocationsRequest, opts ...grpc.CallOption) (*GetDatasetLocationsResponse, error)
	GetDatasetList(ctx context.Context, in *GetDatasetListRequest, opts ...grpc.CallOption) (*GetDatasetListResponse, error)
}

type datasetLocationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatasetLocationsServiceClient(cc grpc.ClientConnInterface) DatasetLocationsServiceClient {
	return &datasetLocationsServiceClient{cc}
}

func (c *datasetLocationsServiceClient) GetDatasetLocations(ctx context.Context, in *GetDatasetLocationsRequest, opts ...grpc.CallOption) (*GetDatasetLocationsResponse, error) {
	out := new(GetDatasetLocationsResponse)
	err := c.cc.Invoke(ctx, DatasetLocationsService_GetDatasetLocations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetLocationsServiceClient) GetDatasetList(ctx context.Context, in *GetDatasetListRequest, opts ...grpc.CallOption) (*GetDatasetListResponse, error) {
	out := new(GetDatasetListResponse)
	err := c.cc.Invoke(ctx, DatasetLocationsService_GetDatasetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasetLocationsServiceServer is the server API for DatasetLocationsService service.
// All implementations must embed UnimplementedDatasetLocationsServiceServer
// for forward compatibility
type DatasetLocationsServiceServer interface {
	GetDatasetLocations(context.Context, *GetDatasetLocationsRequest) (*GetDatasetLocationsResponse, error)
	GetDatasetList(context.Context, *GetDatasetListRequest) (*GetDatasetListResponse, error)
	mustEmbedUnimplementedDatasetLocationsServiceServer()
}

// UnimplementedDatasetLocationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDatasetLocationsServiceServer struct {
}

func (UnimplementedDatasetLocationsServiceServer) GetDatasetLocations(context.Context, *GetDatasetLocationsRequest) (*GetDatasetLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetLocations not implemented")
}
func (UnimplementedDatasetLocationsServiceServer) GetDatasetList(context.Context, *GetDatasetListRequest) (*GetDatasetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetList not implemented")
}
func (UnimplementedDatasetLocationsServiceServer) mustEmbedUnimplementedDatasetLocationsServiceServer() {
}

// UnsafeDatasetLocationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatasetLocationsServiceServer will
// result in compilation errors.
type UnsafeDatasetLocationsServiceServer interface {
	mustEmbedUnimplementedDatasetLocationsServiceServer()
}

func RegisterDatasetLocationsServiceServer(s grpc.ServiceRegistrar, srv DatasetLocationsServiceServer) {
	s.RegisterService(&DatasetLocationsService_ServiceDesc, srv)
}

func _DatasetLocationsService_GetDatasetLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetLocationsServiceServer).GetDatasetLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetLocationsService_GetDatasetLocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetLocationsServiceServer).GetDatasetLocations(ctx, req.(*GetDatasetLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetLocationsService_GetDatasetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetLocationsServiceServer).GetDatasetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetLocationsService_GetDatasetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetLocationsServiceServer).GetDatasetList(ctx, req.(*GetDatasetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatasetLocationsService_ServiceDesc is the grpc.ServiceDesc for DatasetLocationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatasetLocationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.archive.dataset.v2.DatasetLocationsService",
	HandlerType: (*DatasetLocationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDatasetLocations",
			Handler:    _DatasetLocationsService_GetDatasetLocations_Handler,
		},
		{
			MethodName: "GetDatasetList",
			Handler:    _DatasetLocationsService_GetDatasetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/archive/dataset/v2/dataset_locations_service.proto",
}

const (
	InternalDatasetLocationsService_GetDatasetLocations_FullMethodName = "/com.coralogix.archive.dataset.v2.InternalDatasetLocationsService/GetDatasetLocations"
	InternalDatasetLocationsService_GetDatasetList_FullMethodName      = "/com.coralogix.archive.dataset.v2.InternalDatasetLocationsService/GetDatasetList"
)

// InternalDatasetLocationsServiceClient is the client API for InternalDatasetLocationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalDatasetLocationsServiceClient interface {
	GetDatasetLocations(ctx context.Context, in *InternalDatasetLocationsServiceGetDatasetLocationsRequest, opts ...grpc.CallOption) (*InternalDatasetLocationsServiceGetDatasetLocationsResponse, error)
	GetDatasetList(ctx context.Context, in *InternalDatasetLocationsServiceGetDatasetListRequest, opts ...grpc.CallOption) (*InternalDatasetLocationsServiceGetDatasetListResponse, error)
}

type internalDatasetLocationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalDatasetLocationsServiceClient(cc grpc.ClientConnInterface) InternalDatasetLocationsServiceClient {
	return &internalDatasetLocationsServiceClient{cc}
}

func (c *internalDatasetLocationsServiceClient) GetDatasetLocations(ctx context.Context, in *InternalDatasetLocationsServiceGetDatasetLocationsRequest, opts ...grpc.CallOption) (*InternalDatasetLocationsServiceGetDatasetLocationsResponse, error) {
	out := new(InternalDatasetLocationsServiceGetDatasetLocationsResponse)
	err := c.cc.Invoke(ctx, InternalDatasetLocationsService_GetDatasetLocations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalDatasetLocationsServiceClient) GetDatasetList(ctx context.Context, in *InternalDatasetLocationsServiceGetDatasetListRequest, opts ...grpc.CallOption) (*InternalDatasetLocationsServiceGetDatasetListResponse, error) {
	out := new(InternalDatasetLocationsServiceGetDatasetListResponse)
	err := c.cc.Invoke(ctx, InternalDatasetLocationsService_GetDatasetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalDatasetLocationsServiceServer is the server API for InternalDatasetLocationsService service.
// All implementations must embed UnimplementedInternalDatasetLocationsServiceServer
// for forward compatibility
type InternalDatasetLocationsServiceServer interface {
	GetDatasetLocations(context.Context, *InternalDatasetLocationsServiceGetDatasetLocationsRequest) (*InternalDatasetLocationsServiceGetDatasetLocationsResponse, error)
	GetDatasetList(context.Context, *InternalDatasetLocationsServiceGetDatasetListRequest) (*InternalDatasetLocationsServiceGetDatasetListResponse, error)
	mustEmbedUnimplementedInternalDatasetLocationsServiceServer()
}

// UnimplementedInternalDatasetLocationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInternalDatasetLocationsServiceServer struct {
}

func (UnimplementedInternalDatasetLocationsServiceServer) GetDatasetLocations(context.Context, *InternalDatasetLocationsServiceGetDatasetLocationsRequest) (*InternalDatasetLocationsServiceGetDatasetLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetLocations not implemented")
}
func (UnimplementedInternalDatasetLocationsServiceServer) GetDatasetList(context.Context, *InternalDatasetLocationsServiceGetDatasetListRequest) (*InternalDatasetLocationsServiceGetDatasetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetList not implemented")
}
func (UnimplementedInternalDatasetLocationsServiceServer) mustEmbedUnimplementedInternalDatasetLocationsServiceServer() {
}

// UnsafeInternalDatasetLocationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternalDatasetLocationsServiceServer will
// result in compilation errors.
type UnsafeInternalDatasetLocationsServiceServer interface {
	mustEmbedUnimplementedInternalDatasetLocationsServiceServer()
}

func RegisterInternalDatasetLocationsServiceServer(s grpc.ServiceRegistrar, srv InternalDatasetLocationsServiceServer) {
	s.RegisterService(&InternalDatasetLocationsService_ServiceDesc, srv)
}

func _InternalDatasetLocationsService_GetDatasetLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalDatasetLocationsServiceGetDatasetLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalDatasetLocationsServiceServer).GetDatasetLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalDatasetLocationsService_GetDatasetLocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalDatasetLocationsServiceServer).GetDatasetLocations(ctx, req.(*InternalDatasetLocationsServiceGetDatasetLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalDatasetLocationsService_GetDatasetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalDatasetLocationsServiceGetDatasetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalDatasetLocationsServiceServer).GetDatasetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalDatasetLocationsService_GetDatasetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalDatasetLocationsServiceServer).GetDatasetList(ctx, req.(*InternalDatasetLocationsServiceGetDatasetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InternalDatasetLocationsService_ServiceDesc is the grpc.ServiceDesc for InternalDatasetLocationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InternalDatasetLocationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.archive.dataset.v2.InternalDatasetLocationsService",
	HandlerType: (*InternalDatasetLocationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDatasetLocations",
			Handler:    _InternalDatasetLocationsService_GetDatasetLocations_Handler,
		},
		{
			MethodName: "GetDatasetList",
			Handler:    _InternalDatasetLocationsService_GetDatasetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/archive/dataset/v2/dataset_locations_service.proto",
}
