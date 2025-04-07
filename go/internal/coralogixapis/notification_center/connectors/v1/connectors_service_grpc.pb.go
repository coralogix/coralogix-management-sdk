// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ConnectorsService_CreateConnector_FullMethodName            = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/CreateConnector"
	ConnectorsService_ReplaceConnector_FullMethodName           = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/ReplaceConnector"
	ConnectorsService_DeleteConnector_FullMethodName            = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/DeleteConnector"
	ConnectorsService_GetConnector_FullMethodName               = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/GetConnector"
	ConnectorsService_ListConnectors_FullMethodName             = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/ListConnectors"
	ConnectorsService_BatchGetConnectors_FullMethodName         = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/BatchGetConnectors"
	ConnectorsService_BatchGetConnectorSummaries_FullMethodName = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/BatchGetConnectorSummaries"
	ConnectorsService_ListConnectorSummaries_FullMethodName     = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/ListConnectorSummaries"
	ConnectorsService_GetConnectorTypeSummaries_FullMethodName  = "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/GetConnectorTypeSummaries"
)

// ConnectorsServiceClient is the client API for ConnectorsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Public API that allows managing connectors.
type ConnectorsServiceClient interface {
	CreateConnector(ctx context.Context, in *CreateConnectorRequest, opts ...grpc.CallOption) (*CreateConnectorResponse, error)
	ReplaceConnector(ctx context.Context, in *ReplaceConnectorRequest, opts ...grpc.CallOption) (*ReplaceConnectorResponse, error)
	DeleteConnector(ctx context.Context, in *DeleteConnectorRequest, opts ...grpc.CallOption) (*DeleteConnectorResponse, error)
	GetConnector(ctx context.Context, in *GetConnectorRequest, opts ...grpc.CallOption) (*GetConnectorResponse, error)
	ListConnectors(ctx context.Context, in *ListConnectorsRequest, opts ...grpc.CallOption) (*ListConnectorsResponse, error)
	BatchGetConnectors(ctx context.Context, in *BatchGetConnectorsRequest, opts ...grpc.CallOption) (*BatchGetConnectorsResponse, error)
	BatchGetConnectorSummaries(ctx context.Context, in *BatchGetConnectorSummariesRequest, opts ...grpc.CallOption) (*BatchGetConnectorSummariesResponse, error)
	ListConnectorSummaries(ctx context.Context, in *ListConnectorSummariesRequest, opts ...grpc.CallOption) (*ListConnectorSummariesResponse, error)
	GetConnectorTypeSummaries(ctx context.Context, in *GetConnectorTypeSummariesRequest, opts ...grpc.CallOption) (*GetConnectorTypeSummariesResponse, error)
}

type connectorsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorsServiceClient(cc grpc.ClientConnInterface) ConnectorsServiceClient {
	return &connectorsServiceClient{cc}
}

func (c *connectorsServiceClient) CreateConnector(ctx context.Context, in *CreateConnectorRequest, opts ...grpc.CallOption) (*CreateConnectorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateConnectorResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_CreateConnector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) ReplaceConnector(ctx context.Context, in *ReplaceConnectorRequest, opts ...grpc.CallOption) (*ReplaceConnectorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplaceConnectorResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_ReplaceConnector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) DeleteConnector(ctx context.Context, in *DeleteConnectorRequest, opts ...grpc.CallOption) (*DeleteConnectorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteConnectorResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_DeleteConnector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) GetConnector(ctx context.Context, in *GetConnectorRequest, opts ...grpc.CallOption) (*GetConnectorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetConnectorResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_GetConnector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) ListConnectors(ctx context.Context, in *ListConnectorsRequest, opts ...grpc.CallOption) (*ListConnectorsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListConnectorsResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_ListConnectors_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) BatchGetConnectors(ctx context.Context, in *BatchGetConnectorsRequest, opts ...grpc.CallOption) (*BatchGetConnectorsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchGetConnectorsResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_BatchGetConnectors_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) BatchGetConnectorSummaries(ctx context.Context, in *BatchGetConnectorSummariesRequest, opts ...grpc.CallOption) (*BatchGetConnectorSummariesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchGetConnectorSummariesResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_BatchGetConnectorSummaries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) ListConnectorSummaries(ctx context.Context, in *ListConnectorSummariesRequest, opts ...grpc.CallOption) (*ListConnectorSummariesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListConnectorSummariesResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_ListConnectorSummaries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorsServiceClient) GetConnectorTypeSummaries(ctx context.Context, in *GetConnectorTypeSummariesRequest, opts ...grpc.CallOption) (*GetConnectorTypeSummariesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetConnectorTypeSummariesResponse)
	err := c.cc.Invoke(ctx, ConnectorsService_GetConnectorTypeSummaries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorsServiceServer is the server API for ConnectorsService service.
// All implementations must embed UnimplementedConnectorsServiceServer
// for forward compatibility.
//
// Public API that allows managing connectors.
type ConnectorsServiceServer interface {
	CreateConnector(context.Context, *CreateConnectorRequest) (*CreateConnectorResponse, error)
	ReplaceConnector(context.Context, *ReplaceConnectorRequest) (*ReplaceConnectorResponse, error)
	DeleteConnector(context.Context, *DeleteConnectorRequest) (*DeleteConnectorResponse, error)
	GetConnector(context.Context, *GetConnectorRequest) (*GetConnectorResponse, error)
	ListConnectors(context.Context, *ListConnectorsRequest) (*ListConnectorsResponse, error)
	BatchGetConnectors(context.Context, *BatchGetConnectorsRequest) (*BatchGetConnectorsResponse, error)
	BatchGetConnectorSummaries(context.Context, *BatchGetConnectorSummariesRequest) (*BatchGetConnectorSummariesResponse, error)
	ListConnectorSummaries(context.Context, *ListConnectorSummariesRequest) (*ListConnectorSummariesResponse, error)
	GetConnectorTypeSummaries(context.Context, *GetConnectorTypeSummariesRequest) (*GetConnectorTypeSummariesResponse, error)
	mustEmbedUnimplementedConnectorsServiceServer()
}

// UnimplementedConnectorsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConnectorsServiceServer struct{}

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
func (UnimplementedConnectorsServiceServer) BatchGetConnectorSummaries(context.Context, *BatchGetConnectorSummariesRequest) (*BatchGetConnectorSummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetConnectorSummaries not implemented")
}
func (UnimplementedConnectorsServiceServer) ListConnectorSummaries(context.Context, *ListConnectorSummariesRequest) (*ListConnectorSummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConnectorSummaries not implemented")
}
func (UnimplementedConnectorsServiceServer) GetConnectorTypeSummaries(context.Context, *GetConnectorTypeSummariesRequest) (*GetConnectorTypeSummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectorTypeSummaries not implemented")
}
func (UnimplementedConnectorsServiceServer) mustEmbedUnimplementedConnectorsServiceServer() {}
func (UnimplementedConnectorsServiceServer) testEmbeddedByValue()                           {}

// UnsafeConnectorsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectorsServiceServer will
// result in compilation errors.
type UnsafeConnectorsServiceServer interface {
	mustEmbedUnimplementedConnectorsServiceServer()
}

func RegisterConnectorsServiceServer(s grpc.ServiceRegistrar, srv ConnectorsServiceServer) {
	// If the following call pancis, it indicates UnimplementedConnectorsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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

func _ConnectorsService_BatchGetConnectorSummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetConnectorSummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorsServiceServer).BatchGetConnectorSummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectorsService_BatchGetConnectorSummaries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorsServiceServer).BatchGetConnectorSummaries(ctx, req.(*BatchGetConnectorSummariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorsService_ListConnectorSummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConnectorSummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorsServiceServer).ListConnectorSummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectorsService_ListConnectorSummaries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorsServiceServer).ListConnectorSummaries(ctx, req.(*ListConnectorSummariesRequest))
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
			MethodName: "BatchGetConnectorSummaries",
			Handler:    _ConnectorsService_BatchGetConnectorSummaries_Handler,
		},
		{
			MethodName: "ListConnectorSummaries",
			Handler:    _ConnectorsService_ListConnectorSummaries_Handler,
		},
		{
			MethodName: "GetConnectorTypeSummaries",
			Handler:    _ConnectorsService_GetConnectorTypeSummaries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/notification_center/connectors/v1/connectors_service.proto",
}
