// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

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
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDatasetLocationsResponse)
	err := c.cc.Invoke(ctx, DatasetLocationsService_GetDatasetLocations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetLocationsServiceClient) GetDatasetList(ctx context.Context, in *GetDatasetListRequest, opts ...grpc.CallOption) (*GetDatasetListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDatasetListResponse)
	err := c.cc.Invoke(ctx, DatasetLocationsService_GetDatasetList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasetLocationsServiceServer is the server API for DatasetLocationsService service.
// All implementations must embed UnimplementedDatasetLocationsServiceServer
// for forward compatibility.
type DatasetLocationsServiceServer interface {
	GetDatasetLocations(context.Context, *GetDatasetLocationsRequest) (*GetDatasetLocationsResponse, error)
	GetDatasetList(context.Context, *GetDatasetListRequest) (*GetDatasetListResponse, error)
	mustEmbedUnimplementedDatasetLocationsServiceServer()
}

// UnimplementedDatasetLocationsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDatasetLocationsServiceServer struct{}

func (UnimplementedDatasetLocationsServiceServer) GetDatasetLocations(context.Context, *GetDatasetLocationsRequest) (*GetDatasetLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetLocations not implemented")
}
func (UnimplementedDatasetLocationsServiceServer) GetDatasetList(context.Context, *GetDatasetListRequest) (*GetDatasetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetList not implemented")
}
func (UnimplementedDatasetLocationsServiceServer) mustEmbedUnimplementedDatasetLocationsServiceServer() {
}
func (UnimplementedDatasetLocationsServiceServer) testEmbeddedByValue() {}

// UnsafeDatasetLocationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatasetLocationsServiceServer will
// result in compilation errors.
type UnsafeDatasetLocationsServiceServer interface {
	mustEmbedUnimplementedDatasetLocationsServiceServer()
}

func RegisterDatasetLocationsServiceServer(s grpc.ServiceRegistrar, srv DatasetLocationsServiceServer) {
	// If the following call pancis, it indicates UnimplementedDatasetLocationsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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
	InternalDatasetLocationsService_GetDatasetLocations_FullMethodName     = "/com.coralogix.archive.dataset.v2.InternalDatasetLocationsService/GetDatasetLocations"
	InternalDatasetLocationsService_GetDatasetList_FullMethodName          = "/com.coralogix.archive.dataset.v2.InternalDatasetLocationsService/GetDatasetList"
	InternalDatasetLocationsService_GetLocationsByStableIds_FullMethodName = "/com.coralogix.archive.dataset.v2.InternalDatasetLocationsService/GetLocationsByStableIds"
	InternalDatasetLocationsService_UpdateDatasetLocation_FullMethodName   = "/com.coralogix.archive.dataset.v2.InternalDatasetLocationsService/UpdateDatasetLocation"
	InternalDatasetLocationsService_DeleteDatasetLocations_FullMethodName  = "/com.coralogix.archive.dataset.v2.InternalDatasetLocationsService/DeleteDatasetLocations"
)

// InternalDatasetLocationsServiceClient is the client API for InternalDatasetLocationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalDatasetLocationsServiceClient interface {
	GetDatasetLocations(ctx context.Context, in *InternalDatasetLocationsServiceGetDatasetLocationsRequest, opts ...grpc.CallOption) (*InternalDatasetLocationsServiceGetDatasetLocationsResponse, error)
	GetDatasetList(ctx context.Context, in *InternalDatasetLocationsServiceGetDatasetListRequest, opts ...grpc.CallOption) (*InternalDatasetLocationsServiceGetDatasetListResponse, error)
	GetLocationsByStableIds(ctx context.Context, in *InternalDatasetLocationsServiceGetLocationsByStableIdsRequest, opts ...grpc.CallOption) (*InternalDatasetLocationsServiceGetLocationsByStableIdsResponse, error)
	UpdateDatasetLocation(ctx context.Context, in *InternalDatasetLocationsServiceUpdateDatasetLocationRequest, opts ...grpc.CallOption) (*InternalDatasetLocationsServiceUpdateDatasetLocationResponse, error)
	DeleteDatasetLocations(ctx context.Context, in *InternalDatasetLocationsServiceDeleteDatasetLocationsRequest, opts ...grpc.CallOption) (*InternalDatasetLocationsServiceDeleteDatasetLocationsResponse, error)
}

type internalDatasetLocationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalDatasetLocationsServiceClient(cc grpc.ClientConnInterface) InternalDatasetLocationsServiceClient {
	return &internalDatasetLocationsServiceClient{cc}
}

func (c *internalDatasetLocationsServiceClient) GetDatasetLocations(ctx context.Context, in *InternalDatasetLocationsServiceGetDatasetLocationsRequest, opts ...grpc.CallOption) (*InternalDatasetLocationsServiceGetDatasetLocationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InternalDatasetLocationsServiceGetDatasetLocationsResponse)
	err := c.cc.Invoke(ctx, InternalDatasetLocationsService_GetDatasetLocations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalDatasetLocationsServiceClient) GetDatasetList(ctx context.Context, in *InternalDatasetLocationsServiceGetDatasetListRequest, opts ...grpc.CallOption) (*InternalDatasetLocationsServiceGetDatasetListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InternalDatasetLocationsServiceGetDatasetListResponse)
	err := c.cc.Invoke(ctx, InternalDatasetLocationsService_GetDatasetList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalDatasetLocationsServiceClient) GetLocationsByStableIds(ctx context.Context, in *InternalDatasetLocationsServiceGetLocationsByStableIdsRequest, opts ...grpc.CallOption) (*InternalDatasetLocationsServiceGetLocationsByStableIdsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InternalDatasetLocationsServiceGetLocationsByStableIdsResponse)
	err := c.cc.Invoke(ctx, InternalDatasetLocationsService_GetLocationsByStableIds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalDatasetLocationsServiceClient) UpdateDatasetLocation(ctx context.Context, in *InternalDatasetLocationsServiceUpdateDatasetLocationRequest, opts ...grpc.CallOption) (*InternalDatasetLocationsServiceUpdateDatasetLocationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InternalDatasetLocationsServiceUpdateDatasetLocationResponse)
	err := c.cc.Invoke(ctx, InternalDatasetLocationsService_UpdateDatasetLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalDatasetLocationsServiceClient) DeleteDatasetLocations(ctx context.Context, in *InternalDatasetLocationsServiceDeleteDatasetLocationsRequest, opts ...grpc.CallOption) (*InternalDatasetLocationsServiceDeleteDatasetLocationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InternalDatasetLocationsServiceDeleteDatasetLocationsResponse)
	err := c.cc.Invoke(ctx, InternalDatasetLocationsService_DeleteDatasetLocations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalDatasetLocationsServiceServer is the server API for InternalDatasetLocationsService service.
// All implementations must embed UnimplementedInternalDatasetLocationsServiceServer
// for forward compatibility.
type InternalDatasetLocationsServiceServer interface {
	GetDatasetLocations(context.Context, *InternalDatasetLocationsServiceGetDatasetLocationsRequest) (*InternalDatasetLocationsServiceGetDatasetLocationsResponse, error)
	GetDatasetList(context.Context, *InternalDatasetLocationsServiceGetDatasetListRequest) (*InternalDatasetLocationsServiceGetDatasetListResponse, error)
	GetLocationsByStableIds(context.Context, *InternalDatasetLocationsServiceGetLocationsByStableIdsRequest) (*InternalDatasetLocationsServiceGetLocationsByStableIdsResponse, error)
	UpdateDatasetLocation(context.Context, *InternalDatasetLocationsServiceUpdateDatasetLocationRequest) (*InternalDatasetLocationsServiceUpdateDatasetLocationResponse, error)
	DeleteDatasetLocations(context.Context, *InternalDatasetLocationsServiceDeleteDatasetLocationsRequest) (*InternalDatasetLocationsServiceDeleteDatasetLocationsResponse, error)
	mustEmbedUnimplementedInternalDatasetLocationsServiceServer()
}

// UnimplementedInternalDatasetLocationsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInternalDatasetLocationsServiceServer struct{}

func (UnimplementedInternalDatasetLocationsServiceServer) GetDatasetLocations(context.Context, *InternalDatasetLocationsServiceGetDatasetLocationsRequest) (*InternalDatasetLocationsServiceGetDatasetLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetLocations not implemented")
}
func (UnimplementedInternalDatasetLocationsServiceServer) GetDatasetList(context.Context, *InternalDatasetLocationsServiceGetDatasetListRequest) (*InternalDatasetLocationsServiceGetDatasetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetList not implemented")
}
func (UnimplementedInternalDatasetLocationsServiceServer) GetLocationsByStableIds(context.Context, *InternalDatasetLocationsServiceGetLocationsByStableIdsRequest) (*InternalDatasetLocationsServiceGetLocationsByStableIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocationsByStableIds not implemented")
}
func (UnimplementedInternalDatasetLocationsServiceServer) UpdateDatasetLocation(context.Context, *InternalDatasetLocationsServiceUpdateDatasetLocationRequest) (*InternalDatasetLocationsServiceUpdateDatasetLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDatasetLocation not implemented")
}
func (UnimplementedInternalDatasetLocationsServiceServer) DeleteDatasetLocations(context.Context, *InternalDatasetLocationsServiceDeleteDatasetLocationsRequest) (*InternalDatasetLocationsServiceDeleteDatasetLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDatasetLocations not implemented")
}
func (UnimplementedInternalDatasetLocationsServiceServer) mustEmbedUnimplementedInternalDatasetLocationsServiceServer() {
}
func (UnimplementedInternalDatasetLocationsServiceServer) testEmbeddedByValue() {}

// UnsafeInternalDatasetLocationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternalDatasetLocationsServiceServer will
// result in compilation errors.
type UnsafeInternalDatasetLocationsServiceServer interface {
	mustEmbedUnimplementedInternalDatasetLocationsServiceServer()
}

func RegisterInternalDatasetLocationsServiceServer(s grpc.ServiceRegistrar, srv InternalDatasetLocationsServiceServer) {
	// If the following call pancis, it indicates UnimplementedInternalDatasetLocationsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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

func _InternalDatasetLocationsService_GetLocationsByStableIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalDatasetLocationsServiceGetLocationsByStableIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalDatasetLocationsServiceServer).GetLocationsByStableIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalDatasetLocationsService_GetLocationsByStableIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalDatasetLocationsServiceServer).GetLocationsByStableIds(ctx, req.(*InternalDatasetLocationsServiceGetLocationsByStableIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalDatasetLocationsService_UpdateDatasetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalDatasetLocationsServiceUpdateDatasetLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalDatasetLocationsServiceServer).UpdateDatasetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalDatasetLocationsService_UpdateDatasetLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalDatasetLocationsServiceServer).UpdateDatasetLocation(ctx, req.(*InternalDatasetLocationsServiceUpdateDatasetLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalDatasetLocationsService_DeleteDatasetLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalDatasetLocationsServiceDeleteDatasetLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalDatasetLocationsServiceServer).DeleteDatasetLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalDatasetLocationsService_DeleteDatasetLocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalDatasetLocationsServiceServer).DeleteDatasetLocations(ctx, req.(*InternalDatasetLocationsServiceDeleteDatasetLocationsRequest))
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
		{
			MethodName: "GetLocationsByStableIds",
			Handler:    _InternalDatasetLocationsService_GetLocationsByStableIds_Handler,
		},
		{
			MethodName: "UpdateDatasetLocation",
			Handler:    _InternalDatasetLocationsService_UpdateDatasetLocation_Handler,
		},
		{
			MethodName: "DeleteDatasetLocations",
			Handler:    _InternalDatasetLocationsService_DeleteDatasetLocations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/archive/dataset/v2/dataset_locations_service.proto",
}
