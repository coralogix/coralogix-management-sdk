// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
// source: com/coralogixapis/notification_center/connectors/v1/connectors_service.proto

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
	ConnectorsService_CreateConnector_FullMethodName           = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/CreateConnector"
	ConnectorsService_ReplaceConnector_FullMethodName          = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/ReplaceConnector"
	ConnectorsService_DeleteConnector_FullMethodName           = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/DeleteConnector"
	ConnectorsService_GetConnector_FullMethodName              = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/GetConnector"
	ConnectorsService_ListConnectors_FullMethodName            = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/ListConnectors"
	ConnectorsService_BatchGetConnectors_FullMethodName        = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/BatchGetConnectors"
	ConnectorsService_GetConnectorTypeSummaries_FullMethodName = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/GetConnectorTypeSummaries"
)

// ConnectorsServiceClient is the client API for ConnectorsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectorsServiceClient interface {
	CreateConnector(ctx context.Context, in *CreateConnectorRequest, opts ...grpc.CallOption) (*CreateConnectorResponse, error)
	ReplaceConnector(ctx context.Context, in *ReplaceConnectorRequest, opts ...grpc.CallOption) (*ReplaceConnectorResponse, error)
	DeleteConnector(ctx context.Context, in *DeleteConnectorRequest, opts ...grpc.CallOption) (*DeleteConnectorResponse, error)
	GetConnector(ctx context.Context, in *GetConnectorRequest, opts ...grpc.CallOption) (*GetConnectorResponse, error)
	ListConnectors(ctx context.Context, in *ListConnectorsRequest, opts ...grpc.CallOption) (*ListConnectorsResponse, error)
	BatchGetConnectors(ctx context.Context, in *BatchGetConnectorsRequest, opts ...grpc.CallOption) (*BatchGetConnectorsResponse, error)
	GetConnectorTypeSummaries(ctx context.Context, in *GetConnectorTypeSummariesRequest, opts ...grpc.CallOption) (*GetConnectorTypeSummariesResponse, error)
}

type connectorsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorsServiceClient(cc grpc.ClientConnInterface) ConnectorsServiceClient {
	return &connectorsServiceClient{cc}
}

func (c *connectorsServiceClient) CreateConnector(ctx context.Context, in *CreateConnectorRequest, opts ...grpc.CallOption) (*CreateConnectorResponse, error) {
	out := new(CreateConnectorResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_CreateConnector_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) ReplaceConnector(ctx context.Context, in *ReplaceConnectorRequest, opts ...grpc.CallOption) (*ReplaceConnectorResponse, error) {
	out := new(ReplaceConnectorResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_ReplaceConnector_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) DeleteConnector(ctx context.Context, in *DeleteConnectorRequest, opts ...grpc.CallOption) (*DeleteConnectorResponse, error) {
	out := new(DeleteConnectorResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_DeleteConnector_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) GetConnector(ctx context.Context, in *GetConnectorRequest, opts ...grpc.CallOption) (*GetConnectorResponse, error) {
	out := new(GetConnectorResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_GetConnector_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) ListConnectors(ctx context.Context, in *ListConnectorsRequest, opts ...grpc.CallOption) (*ListConnectorsResponse, error) {
	out := new(ListConnectorsResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_ListConnectors_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) BatchGetConnectors(ctx context.Context, in *BatchGetConnectorsRequest, opts ...grpc.CallOption) (*BatchGetConnectorsResponse, error) {
	out := new(BatchGetConnectorsResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_BatchGetConnectors_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) GetConnectorTypeSummaries(ctx context.Context, in *GetConnectorTypeSummariesRequest, opts ...grpc.CallOption) (*GetConnectorTypeSummariesResponse, error) {
	out := new(GetConnectorTypeSummariesResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_GetConnectorTypeSummaries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorsServiceServer is the server API for ConnectorsService service.
// All implementations must embed UnimplementedConnectorsServiceServer
// for forward compatibility
type ConnectorsServiceServer interface {
	CreateConnector(context.Context, *CreateConnectorRequest) (*CreateConnectorResponse, error)
	ReplaceConnector(context.Context, *ReplaceConnectorRequest) (*ReplaceConnectorResponse, error)
	DeleteConnector(context.Context, *DeleteConnectorRequest) (*DeleteConnectorResponse, error)
	GetConnector(context.Context, *GetConnectorRequest) (*GetConnectorResponse, error)
	ListConnectors(context.Context, *ListConnectorsRequest) (*ListConnectorsResponse, error)
	BatchGetConnectors(context.Context, *BatchGetConnectorsRequest) (*BatchGetConnectorsResponse, error)
	GetConnectorTypeSummaries(context.Context, *GetConnectorTypeSummariesRequest) (*GetConnectorTypeSummariesResponse, error)
	mustEmbedUnimplementedConnectorsServiceServer()
}

// UnimplementedConnectorsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConnectorsServiceServer struct {
}

func (UnimplementedConnectorsServiceServer) CreateConnector(context.Context, *CreateConnectorRequest) (*CreateConnectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConnector not implemented")
}
func (UnimplementedConnectorsServiceServer) ReplaceConnector(context.Context, *ReplaceConnectorRequest) (*ReplaceConnectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceConnector not implemented")
}
func (UnimplementedConnectorsServiceServer) DeleteConnector(context.Context, *DeleteConnectorRequest) (*DeleteConnectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConnector not implemented")
}
func (UnimplementedConnectorsServiceServer) GetConnector(context.Context, *GetConnectorRequest) (*GetConnectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnector not implemented")
}
func (UnimplementedConnectorsServiceServer) ListConnectors(context.Context, *ListConnectorsRequest) (*ListConnectorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConnectors not implemented")
}
func (UnimplementedConnectorsServiceServer) BatchGetConnectors(context.Context, *BatchGetConnectorsRequest) (*BatchGetConnectorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetConnectors not implemented")
}
func (UnimplementedConnectorsServiceServer) GetConnectorTypeSummaries(context.Context, *GetConnectorTypeSummariesRequest) (*GetConnectorTypeSummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectorTypeSummaries not implemented")
}
func (UnimplementedConnectorsServiceServer) mustEmbedUnimplementedConnectorsServiceServer() {}

// UnsafeConnectorsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectorsServiceServer will
// result in compilation errors.
type UnsafeConnectorsServiceServer interface {
	mustEmbedUnimplementedConnectorsServiceServer()
}

func RegisterConnectorsServiceServer(s grpc.ServiceRegistrar, srv ConnectorsServiceServer) {
	s.RegisterService(&ConnectorsService_ServiceDesc, srv)
}

func _ConnectorsService_CreateConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorsServiceServer).CreateConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectorsService_CreateConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorsServiceServer).CreateConnector(ctx, req.(*CreateConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorsService_ReplaceConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorsServiceServer).ReplaceConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectorsService_ReplaceConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorsServiceServer).ReplaceConnector(ctx, req.(*ReplaceConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorsService_DeleteConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorsServiceServer).DeleteConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectorsService_DeleteConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorsServiceServer).DeleteConnector(ctx, req.(*DeleteConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorsService_GetConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorsServiceServer).GetConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectorsService_GetConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorsServiceServer).GetConnector(ctx, req.(*GetConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorsService_ListConnectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConnectorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorsServiceServer).ListConnectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectorsService_ListConnectors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorsServiceServer).ListConnectors(ctx, req.(*ListConnectorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorsService_BatchGetConnectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetConnectorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorsServiceServer).BatchGetConnectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectorsService_BatchGetConnectors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorsServiceServer).BatchGetConnectors(ctx, req.(*BatchGetConnectorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorsService_GetConnectorTypeSummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectorTypeSummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorsServiceServer).GetConnectorTypeSummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectorsService_GetConnectorTypeSummaries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorsServiceServer).GetConnectorTypeSummaries(ctx, req.(*GetConnectorTypeSummariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectorsService_ServiceDesc is the grpc.ServiceDesc for ConnectorsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectorsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.notification_center.connectors.v1.ConnectorsService",
	HandlerType: (*ConnectorsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConnector",
			Handler:    _ConnectorsService_CreateConnector_Handler,
		},
		{
			MethodName: "ReplaceConnector",
			Handler:    _ConnectorsService_ReplaceConnector_Handler,
		},
		{
			MethodName: "DeleteConnector",
			Handler:    _ConnectorsService_DeleteConnector_Handler,
		},
		{
			MethodName: "GetConnector",
			Handler:    _ConnectorsService_GetConnector_Handler,
		},
		{
			MethodName: "ListConnectors",
			Handler:    _ConnectorsService_ListConnectors_Handler,
		},
		{
			MethodName: "BatchGetConnectors",
			Handler:    _ConnectorsService_BatchGetConnectors_Handler,
		},
		{
			MethodName: "GetConnectorTypeSummaries",
			Handler:    _ConnectorsService_GetConnectorTypeSummaries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/notification_center/connectors/v1/connectors_service.proto",
}
