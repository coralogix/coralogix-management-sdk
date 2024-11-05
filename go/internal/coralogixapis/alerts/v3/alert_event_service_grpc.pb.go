// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.28.3
// source: com/coralogixapis/alerts/v3/event/alert_event_service.proto

package v3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AlertEventService_GetAlertEvent_FullMethodName       = "/com.coralogixapis.alerts.v3.AlertEventService/GetAlertEvent"
	AlertEventService_GetAlertEventsStats_FullMethodName = "/com.coralogixapis.alerts.v3.AlertEventService/GetAlertEventsStats"
)

// AlertEventServiceClient is the client API for AlertEventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertEventServiceClient interface {
	GetAlertEvent(ctx context.Context, in *GetAlertEventRequest, opts ...grpc.CallOption) (*GetAlertEventResponse, error)
	GetAlertEventsStats(ctx context.Context, in *GetAlertEventStatsRequest, opts ...grpc.CallOption) (*GetAlertEventStatsResponse, error)
}

type alertEventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertEventServiceClient(cc grpc.ClientConnInterface) AlertEventServiceClient {
	return &alertEventServiceClient{cc}
}

func (c *alertEventServiceClient) GetAlertEvent(ctx context.Context, in *GetAlertEventRequest, opts ...grpc.CallOption) (*GetAlertEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAlertEventResponse)
	err := c.cc.Invoke(ctx, AlertEventService_GetAlertEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertEventServiceClient) GetAlertEventsStats(ctx context.Context, in *GetAlertEventStatsRequest, opts ...grpc.CallOption) (*GetAlertEventStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAlertEventStatsResponse)
	err := c.cc.Invoke(ctx, AlertEventService_GetAlertEventsStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertEventServiceServer is the server API for AlertEventService service.
// All implementations must embed UnimplementedAlertEventServiceServer
// for forward compatibility
type AlertEventServiceServer interface {
	GetAlertEvent(context.Context, *GetAlertEventRequest) (*GetAlertEventResponse, error)
	GetAlertEventsStats(context.Context, *GetAlertEventStatsRequest) (*GetAlertEventStatsResponse, error)
	mustEmbedUnimplementedAlertEventServiceServer()
}

// UnimplementedAlertEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAlertEventServiceServer struct {
}

func (UnimplementedAlertEventServiceServer) GetAlertEvent(context.Context, *GetAlertEventRequest) (*GetAlertEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertEvent not implemented")
}
func (UnimplementedAlertEventServiceServer) GetAlertEventsStats(context.Context, *GetAlertEventStatsRequest) (*GetAlertEventStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertEventsStats not implemented")
}
func (UnimplementedAlertEventServiceServer) mustEmbedUnimplementedAlertEventServiceServer() {}

// UnsafeAlertEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertEventServiceServer will
// result in compilation errors.
type UnsafeAlertEventServiceServer interface {
	mustEmbedUnimplementedAlertEventServiceServer()
}

func RegisterAlertEventServiceServer(s grpc.ServiceRegistrar, srv AlertEventServiceServer) {
	s.RegisterService(&AlertEventService_ServiceDesc, srv)
}

func _AlertEventService_GetAlertEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertEventServiceServer).GetAlertEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertEventService_GetAlertEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertEventServiceServer).GetAlertEvent(ctx, req.(*GetAlertEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertEventService_GetAlertEventsStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertEventStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertEventServiceServer).GetAlertEventsStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertEventService_GetAlertEventsStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertEventServiceServer).GetAlertEventsStats(ctx, req.(*GetAlertEventStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AlertEventService_ServiceDesc is the grpc.ServiceDesc for AlertEventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlertEventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.alerts.v3.AlertEventService",
	HandlerType: (*AlertEventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAlertEvent",
			Handler:    _AlertEventService_GetAlertEvent_Handler,
		},
		{
			MethodName: "GetAlertEventsStats",
			Handler:    _AlertEventService_GetAlertEventsStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/alerts/v3/event/alert_event_service.proto",
}
