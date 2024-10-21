// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.2
// source: com/coralogixapis/alerts/v3/alert_defs_service.proto

package v3

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
	AlertDefsService_GetAlertDef_FullMethodName     = "/com.coralogixapis.alerts.v3.AlertDefsService/GetAlertDef"
	AlertDefsService_CreateAlertDef_FullMethodName  = "/com.coralogixapis.alerts.v3.AlertDefsService/CreateAlertDef"
	AlertDefsService_ReplaceAlertDef_FullMethodName = "/com.coralogixapis.alerts.v3.AlertDefsService/ReplaceAlertDef"
	AlertDefsService_ListAlertDefs_FullMethodName   = "/com.coralogixapis.alerts.v3.AlertDefsService/ListAlertDefs"
	AlertDefsService_DownloadAlerts_FullMethodName  = "/com.coralogixapis.alerts.v3.AlertDefsService/DownloadAlerts"
	AlertDefsService_DeleteAlertDef_FullMethodName  = "/com.coralogixapis.alerts.v3.AlertDefsService/DeleteAlertDef"
	AlertDefsService_SetActive_FullMethodName       = "/com.coralogixapis.alerts.v3.AlertDefsService/SetActive"
)

// AlertDefsServiceClient is the client API for AlertDefsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertDefsServiceClient interface {
	// Get Alert Def by non changing ID AKA UniqueIdentifier
	GetAlertDef(ctx context.Context, in *GetAlertDefRequest, opts ...grpc.CallOption) (*GetAlertDefResponse, error)
	CreateAlertDef(ctx context.Context, in *CreateAlertDefRequest, opts ...grpc.CallOption) (*CreateAlertDefResponse, error)
	ReplaceAlertDef(ctx context.Context, in *ReplaceAlertDefRequest, opts ...grpc.CallOption) (*ReplaceAlertDefResponse, error)
	ListAlertDefs(ctx context.Context, in *ListAlertDefsRequest, opts ...grpc.CallOption) (*ListAlertDefsResponse, error)
	DownloadAlerts(ctx context.Context, in *DownloadAlertsRequest, opts ...grpc.CallOption) (AlertDefsService_DownloadAlertsClient, error)
	DeleteAlertDef(ctx context.Context, in *DeleteAlertDefRequest, opts ...grpc.CallOption) (*DeleteAlertDefResponse, error)
	SetActive(ctx context.Context, in *SetActiveRequest, opts ...grpc.CallOption) (*SetActiveResponse, error)
}

type alertDefsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertDefsServiceClient(cc grpc.ClientConnInterface) AlertDefsServiceClient {
	return &alertDefsServiceClient{cc}
}

func (c *alertDefsServiceClient) GetAlertDef(ctx context.Context, in *GetAlertDefRequest, opts ...grpc.CallOption) (*GetAlertDefResponse, error) {
	out := new(GetAlertDefResponse)
	err := c.cc.Invoke(ctx, AlertDefsService_GetAlertDef_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertDefsServiceClient) CreateAlertDef(ctx context.Context, in *CreateAlertDefRequest, opts ...grpc.CallOption) (*CreateAlertDefResponse, error) {
	out := new(CreateAlertDefResponse)
	err := c.cc.Invoke(ctx, AlertDefsService_CreateAlertDef_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertDefsServiceClient) ReplaceAlertDef(ctx context.Context, in *ReplaceAlertDefRequest, opts ...grpc.CallOption) (*ReplaceAlertDefResponse, error) {
	out := new(ReplaceAlertDefResponse)
	err := c.cc.Invoke(ctx, AlertDefsService_ReplaceAlertDef_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertDefsServiceClient) ListAlertDefs(ctx context.Context, in *ListAlertDefsRequest, opts ...grpc.CallOption) (*ListAlertDefsResponse, error) {
	out := new(ListAlertDefsResponse)
	err := c.cc.Invoke(ctx, AlertDefsService_ListAlertDefs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertDefsServiceClient) DownloadAlerts(ctx context.Context, in *DownloadAlertsRequest, opts ...grpc.CallOption) (AlertDefsService_DownloadAlertsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AlertDefsService_ServiceDesc.Streams[0], AlertDefsService_DownloadAlerts_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &alertDefsServiceDownloadAlertsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AlertDefsService_DownloadAlertsClient interface {
	Recv() (*DownloadAlertsResponse, error)
	grpc.ClientStream
}

type alertDefsServiceDownloadAlertsClient struct {
	grpc.ClientStream
}

func (x *alertDefsServiceDownloadAlertsClient) Recv() (*DownloadAlertsResponse, error) {
	m := new(DownloadAlertsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *alertDefsServiceClient) DeleteAlertDef(ctx context.Context, in *DeleteAlertDefRequest, opts ...grpc.CallOption) (*DeleteAlertDefResponse, error) {
	out := new(DeleteAlertDefResponse)
	err := c.cc.Invoke(ctx, AlertDefsService_DeleteAlertDef_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertDefsServiceClient) SetActive(ctx context.Context, in *SetActiveRequest, opts ...grpc.CallOption) (*SetActiveResponse, error) {
	out := new(SetActiveResponse)
	err := c.cc.Invoke(ctx, AlertDefsService_SetActive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertDefsServiceServer is the server API for AlertDefsService service.
// All implementations must embed UnimplementedAlertDefsServiceServer
// for forward compatibility
type AlertDefsServiceServer interface {
	// Get Alert Def by non changing ID AKA UniqueIdentifier
	GetAlertDef(context.Context, *GetAlertDefRequest) (*GetAlertDefResponse, error)
	CreateAlertDef(context.Context, *CreateAlertDefRequest) (*CreateAlertDefResponse, error)
	ReplaceAlertDef(context.Context, *ReplaceAlertDefRequest) (*ReplaceAlertDefResponse, error)
	ListAlertDefs(context.Context, *ListAlertDefsRequest) (*ListAlertDefsResponse, error)
	DownloadAlerts(*DownloadAlertsRequest, AlertDefsService_DownloadAlertsServer) error
	DeleteAlertDef(context.Context, *DeleteAlertDefRequest) (*DeleteAlertDefResponse, error)
	SetActive(context.Context, *SetActiveRequest) (*SetActiveResponse, error)
	mustEmbedUnimplementedAlertDefsServiceServer()
}

// UnimplementedAlertDefsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAlertDefsServiceServer struct {
}

func (UnimplementedAlertDefsServiceServer) GetAlertDef(context.Context, *GetAlertDefRequest) (*GetAlertDefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertDef not implemented")
}
func (UnimplementedAlertDefsServiceServer) CreateAlertDef(context.Context, *CreateAlertDefRequest) (*CreateAlertDefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlertDef not implemented")
}
func (UnimplementedAlertDefsServiceServer) ReplaceAlertDef(context.Context, *ReplaceAlertDefRequest) (*ReplaceAlertDefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceAlertDef not implemented")
}
func (UnimplementedAlertDefsServiceServer) ListAlertDefs(context.Context, *ListAlertDefsRequest) (*ListAlertDefsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlertDefs not implemented")
}
func (UnimplementedAlertDefsServiceServer) DownloadAlerts(*DownloadAlertsRequest, AlertDefsService_DownloadAlertsServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadAlerts not implemented")
}
func (UnimplementedAlertDefsServiceServer) DeleteAlertDef(context.Context, *DeleteAlertDefRequest) (*DeleteAlertDefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlertDef not implemented")
}
func (UnimplementedAlertDefsServiceServer) SetActive(context.Context, *SetActiveRequest) (*SetActiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetActive not implemented")
}
func (UnimplementedAlertDefsServiceServer) mustEmbedUnimplementedAlertDefsServiceServer() {}

// UnsafeAlertDefsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertDefsServiceServer will
// result in compilation errors.
type UnsafeAlertDefsServiceServer interface {
	mustEmbedUnimplementedAlertDefsServiceServer()
}

func RegisterAlertDefsServiceServer(s grpc.ServiceRegistrar, srv AlertDefsServiceServer) {
	s.RegisterService(&AlertDefsService_ServiceDesc, srv)
}

func _AlertDefsService_GetAlertDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertDefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertDefsServiceServer).GetAlertDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertDefsService_GetAlertDef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertDefsServiceServer).GetAlertDef(ctx, req.(*GetAlertDefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertDefsService_CreateAlertDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlertDefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertDefsServiceServer).CreateAlertDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertDefsService_CreateAlertDef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertDefsServiceServer).CreateAlertDef(ctx, req.(*CreateAlertDefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertDefsService_ReplaceAlertDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceAlertDefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertDefsServiceServer).ReplaceAlertDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertDefsService_ReplaceAlertDef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertDefsServiceServer).ReplaceAlertDef(ctx, req.(*ReplaceAlertDefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertDefsService_ListAlertDefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertDefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertDefsServiceServer).ListAlertDefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertDefsService_ListAlertDefs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertDefsServiceServer).ListAlertDefs(ctx, req.(*ListAlertDefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertDefsService_DownloadAlerts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadAlertsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AlertDefsServiceServer).DownloadAlerts(m, &alertDefsServiceDownloadAlertsServer{stream})
}

type AlertDefsService_DownloadAlertsServer interface {
	Send(*DownloadAlertsResponse) error
	grpc.ServerStream
}

type alertDefsServiceDownloadAlertsServer struct {
	grpc.ServerStream
}

func (x *alertDefsServiceDownloadAlertsServer) Send(m *DownloadAlertsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AlertDefsService_DeleteAlertDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAlertDefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertDefsServiceServer).DeleteAlertDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertDefsService_DeleteAlertDef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertDefsServiceServer).DeleteAlertDef(ctx, req.(*DeleteAlertDefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertDefsService_SetActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetActiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertDefsServiceServer).SetActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertDefsService_SetActive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertDefsServiceServer).SetActive(ctx, req.(*SetActiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AlertDefsService_ServiceDesc is the grpc.ServiceDesc for AlertDefsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlertDefsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.alerts.v3.AlertDefsService",
	HandlerType: (*AlertDefsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAlertDef",
			Handler:    _AlertDefsService_GetAlertDef_Handler,
		},
		{
			MethodName: "CreateAlertDef",
			Handler:    _AlertDefsService_CreateAlertDef_Handler,
		},
		{
			MethodName: "ReplaceAlertDef",
			Handler:    _AlertDefsService_ReplaceAlertDef_Handler,
		},
		{
			MethodName: "ListAlertDefs",
			Handler:    _AlertDefsService_ListAlertDefs_Handler,
		},
		{
			MethodName: "DeleteAlertDef",
			Handler:    _AlertDefsService_DeleteAlertDef_Handler,
		},
		{
			MethodName: "SetActive",
			Handler:    _AlertDefsService_SetActive_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadAlerts",
			Handler:       _AlertDefsService_DownloadAlerts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "com/coralogixapis/alerts/v3/alert_defs_service.proto",
}
