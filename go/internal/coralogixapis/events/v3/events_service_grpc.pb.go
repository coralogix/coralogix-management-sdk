// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: com/coralogixapis/events/v3/events_service.proto

package v3

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
	EventsService_GetEvent_FullMethodName            = "/com.coralogixapis.events.v3.EventsService/GetEvent"
	EventsService_BatchGetEvent_FullMethodName       = "/com.coralogixapis.events.v3.EventsService/BatchGetEvent"
	EventsService_ListEvents_FullMethodName          = "/com.coralogixapis.events.v3.EventsService/ListEvents"
	EventsService_ListEventsCount_FullMethodName     = "/com.coralogixapis.events.v3.EventsService/ListEventsCount"
	EventsService_GetEventsStatistics_FullMethodName = "/com.coralogixapis.events.v3.EventsService/GetEventsStatistics"
)

// EventsServiceClient is the client API for EventsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventsServiceClient interface {
	GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error)
	BatchGetEvent(ctx context.Context, in *BatchGetEventRequest, opts ...grpc.CallOption) (*BatchGetEventResponse, error)
	ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error)
	ListEventsCount(ctx context.Context, in *ListEventsCountRequest, opts ...grpc.CallOption) (*ListEventsCountResponse, error)
	GetEventsStatistics(ctx context.Context, in *GetEventsStatisticsRequest, opts ...grpc.CallOption) (*GetEventsStatisticsResponse, error)
}

type eventsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventsServiceClient(cc grpc.ClientConnInterface) EventsServiceClient {
	return &eventsServiceClient{cc}
}

func (c *eventsServiceClient) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEventResponse)
	err := c.cc.Invoke(ctx, EventsService_GetEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsServiceClient) BatchGetEvent(ctx context.Context, in *BatchGetEventRequest, opts ...grpc.CallOption) (*BatchGetEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchGetEventResponse)
	err := c.cc.Invoke(ctx, EventsService_BatchGetEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsServiceClient) ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListEventsResponse)
	err := c.cc.Invoke(ctx, EventsService_ListEvents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsServiceClient) ListEventsCount(ctx context.Context, in *ListEventsCountRequest, opts ...grpc.CallOption) (*ListEventsCountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListEventsCountResponse)
	err := c.cc.Invoke(ctx, EventsService_ListEventsCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsServiceClient) GetEventsStatistics(ctx context.Context, in *GetEventsStatisticsRequest, opts ...grpc.CallOption) (*GetEventsStatisticsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEventsStatisticsResponse)
	err := c.cc.Invoke(ctx, EventsService_GetEventsStatistics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventsServiceServer is the server API for EventsService service.
// All implementations must embed UnimplementedEventsServiceServer
// for forward compatibility.
type EventsServiceServer interface {
	GetEvent(context.Context, *GetEventRequest) (*GetEventResponse, error)
	BatchGetEvent(context.Context, *BatchGetEventRequest) (*BatchGetEventResponse, error)
	ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error)
	ListEventsCount(context.Context, *ListEventsCountRequest) (*ListEventsCountResponse, error)
	GetEventsStatistics(context.Context, *GetEventsStatisticsRequest) (*GetEventsStatisticsResponse, error)
	mustEmbedUnimplementedEventsServiceServer()
}

// UnimplementedEventsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEventsServiceServer struct{}

func (UnimplementedEventsServiceServer) GetEvent(context.Context, *GetEventRequest) (*GetEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedEventsServiceServer) BatchGetEvent(context.Context, *BatchGetEventRequest) (*BatchGetEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetEvent not implemented")
}
func (UnimplementedEventsServiceServer) ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEvents not implemented")
}
func (UnimplementedEventsServiceServer) ListEventsCount(context.Context, *ListEventsCountRequest) (*ListEventsCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEventsCount not implemented")
}
func (UnimplementedEventsServiceServer) GetEventsStatistics(context.Context, *GetEventsStatisticsRequest) (*GetEventsStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsStatistics not implemented")
}
func (UnimplementedEventsServiceServer) mustEmbedUnimplementedEventsServiceServer() {}
func (UnimplementedEventsServiceServer) testEmbeddedByValue()                       {}

// UnsafeEventsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventsServiceServer will
// result in compilation errors.
type UnsafeEventsServiceServer interface {
	mustEmbedUnimplementedEventsServiceServer()
}

func RegisterEventsServiceServer(s grpc.ServiceRegistrar, srv EventsServiceServer) {
	// If the following call pancis, it indicates UnimplementedEventsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EventsService_ServiceDesc, srv)
}

func _EventsService_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventsService_GetEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).GetEvent(ctx, req.(*GetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsService_BatchGetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).BatchGetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventsService_BatchGetEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).BatchGetEvent(ctx, req.(*BatchGetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsService_ListEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).ListEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventsService_ListEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).ListEvents(ctx, req.(*ListEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsService_ListEventsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventsCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).ListEventsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventsService_ListEventsCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).ListEventsCount(ctx, req.(*ListEventsCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsService_GetEventsStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).GetEventsStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventsService_GetEventsStatistics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).GetEventsStatistics(ctx, req.(*GetEventsStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventsService_ServiceDesc is the grpc.ServiceDesc for EventsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.events.v3.EventsService",
	HandlerType: (*EventsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvent",
			Handler:    _EventsService_GetEvent_Handler,
		},
		{
			MethodName: "BatchGetEvent",
			Handler:    _EventsService_BatchGetEvent_Handler,
		},
		{
			MethodName: "ListEvents",
			Handler:    _EventsService_ListEvents_Handler,
		},
		{
			MethodName: "ListEventsCount",
			Handler:    _EventsService_ListEventsCount_Handler,
		},
		{
			MethodName: "GetEventsStatistics",
			Handler:    _EventsService_GetEventsStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/events/v3/events_service.proto",
}
