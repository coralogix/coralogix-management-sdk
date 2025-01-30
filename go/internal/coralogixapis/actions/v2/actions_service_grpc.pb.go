// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: com/coralogixapis/actions/v2/actions_service.proto

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
	ActionsService_CreateAction_FullMethodName              = "/com.coralogixapis.actions.v2.ActionsService/CreateAction"
	ActionsService_ReplaceAction_FullMethodName             = "/com.coralogixapis.actions.v2.ActionsService/ReplaceAction"
	ActionsService_DeleteAction_FullMethodName              = "/com.coralogixapis.actions.v2.ActionsService/DeleteAction"
	ActionsService_GetAction_FullMethodName                 = "/com.coralogixapis.actions.v2.ActionsService/GetAction"
	ActionsService_ListActions_FullMethodName               = "/com.coralogixapis.actions.v2.ActionsService/ListActions"
	ActionsService_OrderActions_FullMethodName              = "/com.coralogixapis.actions.v2.ActionsService/OrderActions"
	ActionsService_AtomicBatchExecuteActions_FullMethodName = "/com.coralogixapis.actions.v2.ActionsService/AtomicBatchExecuteActions"
)

// ActionsServiceClient is the client API for ActionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActionsServiceClient interface {
	CreateAction(ctx context.Context, in *CreateActionRequest, opts ...grpc.CallOption) (*CreateActionResponse, error)
	ReplaceAction(ctx context.Context, in *ReplaceActionRequest, opts ...grpc.CallOption) (*ReplaceActionResponse, error)
	DeleteAction(ctx context.Context, in *DeleteActionRequest, opts ...grpc.CallOption) (*DeleteActionResponse, error)
	GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionResponse, error)
	ListActions(ctx context.Context, in *ListActionsRequest, opts ...grpc.CallOption) (*ListActionsResponse, error)
	OrderActions(ctx context.Context, in *OrderActionsRequest, opts ...grpc.CallOption) (*OrderActionsResponse, error)
	AtomicBatchExecuteActions(ctx context.Context, in *AtomicBatchExecuteActionsRequest, opts ...grpc.CallOption) (*AtomicBatchExecuteActionsResponse, error)
}

type actionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActionsServiceClient(cc grpc.ClientConnInterface) ActionsServiceClient {
	return &actionsServiceClient{cc}
}

func (c *actionsServiceClient) CreateAction(ctx context.Context, in *CreateActionRequest, opts ...grpc.CallOption) (*CreateActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_CreateAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) ReplaceAction(ctx context.Context, in *ReplaceActionRequest, opts ...grpc.CallOption) (*ReplaceActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplaceActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_ReplaceAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) DeleteAction(ctx context.Context, in *DeleteActionRequest, opts ...grpc.CallOption) (*DeleteActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_DeleteAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_GetAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) ListActions(ctx context.Context, in *ListActionsRequest, opts ...grpc.CallOption) (*ListActionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListActionsResponse)
	err := c.cc.Invoke(ctx, ActionsService_ListActions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) OrderActions(ctx context.Context, in *OrderActionsRequest, opts ...grpc.CallOption) (*OrderActionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderActionsResponse)
	err := c.cc.Invoke(ctx, ActionsService_OrderActions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) AtomicBatchExecuteActions(ctx context.Context, in *AtomicBatchExecuteActionsRequest, opts ...grpc.CallOption) (*AtomicBatchExecuteActionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AtomicBatchExecuteActionsResponse)
	err := c.cc.Invoke(ctx, ActionsService_AtomicBatchExecuteActions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionsServiceServer is the server API for ActionsService service.
// All implementations must embed UnimplementedActionsServiceServer
// for forward compatibility.
type ActionsServiceServer interface {
	CreateAction(context.Context, *CreateActionRequest) (*CreateActionResponse, error)
	ReplaceAction(context.Context, *ReplaceActionRequest) (*ReplaceActionResponse, error)
	DeleteAction(context.Context, *DeleteActionRequest) (*DeleteActionResponse, error)
	GetAction(context.Context, *GetActionRequest) (*GetActionResponse, error)
	ListActions(context.Context, *ListActionsRequest) (*ListActionsResponse, error)
	OrderActions(context.Context, *OrderActionsRequest) (*OrderActionsResponse, error)
	AtomicBatchExecuteActions(context.Context, *AtomicBatchExecuteActionsRequest) (*AtomicBatchExecuteActionsResponse, error)
	mustEmbedUnimplementedActionsServiceServer()
}

// UnimplementedActionsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedActionsServiceServer struct{}

func (UnimplementedActionsServiceServer) CreateAction(context.Context, *CreateActionRequest) (*CreateActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAction not implemented")
}
func (UnimplementedActionsServiceServer) ReplaceAction(context.Context, *ReplaceActionRequest) (*ReplaceActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceAction not implemented")
}
func (UnimplementedActionsServiceServer) DeleteAction(context.Context, *DeleteActionRequest) (*DeleteActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAction not implemented")
}
func (UnimplementedActionsServiceServer) GetAction(context.Context, *GetActionRequest) (*GetActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAction not implemented")
}
func (UnimplementedActionsServiceServer) ListActions(context.Context, *ListActionsRequest) (*ListActionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActions not implemented")
}
func (UnimplementedActionsServiceServer) OrderActions(context.Context, *OrderActionsRequest) (*OrderActionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderActions not implemented")
}
func (UnimplementedActionsServiceServer) AtomicBatchExecuteActions(context.Context, *AtomicBatchExecuteActionsRequest) (*AtomicBatchExecuteActionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtomicBatchExecuteActions not implemented")
}
func (UnimplementedActionsServiceServer) mustEmbedUnimplementedActionsServiceServer() {}
func (UnimplementedActionsServiceServer) testEmbeddedByValue()                        {}

// UnsafeActionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActionsServiceServer will
// result in compilation errors.
type UnsafeActionsServiceServer interface {
	mustEmbedUnimplementedActionsServiceServer()
}

func RegisterActionsServiceServer(s grpc.ServiceRegistrar, srv ActionsServiceServer) {
	// If the following call pancis, it indicates UnimplementedActionsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ActionsService_ServiceDesc, srv)
}

func _ActionsService_CreateAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).CreateAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_CreateAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).CreateAction(ctx, req.(*CreateActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_ReplaceAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).ReplaceAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_ReplaceAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).ReplaceAction(ctx, req.(*ReplaceActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_DeleteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).DeleteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_DeleteAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).DeleteAction(ctx, req.(*DeleteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_GetAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).GetAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_GetAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).GetAction(ctx, req.(*GetActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_ListActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).ListActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_ListActions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).ListActions(ctx, req.(*ListActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_OrderActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).OrderActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_OrderActions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).OrderActions(ctx, req.(*OrderActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_AtomicBatchExecuteActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AtomicBatchExecuteActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).AtomicBatchExecuteActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_AtomicBatchExecuteActions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).AtomicBatchExecuteActions(ctx, req.(*AtomicBatchExecuteActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActionsService_ServiceDesc is the grpc.ServiceDesc for ActionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.actions.v2.ActionsService",
	HandlerType: (*ActionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAction",
			Handler:    _ActionsService_CreateAction_Handler,
		},
		{
			MethodName: "ReplaceAction",
			Handler:    _ActionsService_ReplaceAction_Handler,
		},
		{
			MethodName: "DeleteAction",
			Handler:    _ActionsService_DeleteAction_Handler,
		},
		{
			MethodName: "GetAction",
			Handler:    _ActionsService_GetAction_Handler,
		},
		{
			MethodName: "ListActions",
			Handler:    _ActionsService_ListActions_Handler,
		},
		{
			MethodName: "OrderActions",
			Handler:    _ActionsService_OrderActions_Handler,
		},
		{
			MethodName: "AtomicBatchExecuteActions",
			Handler:    _ActionsService_AtomicBatchExecuteActions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/actions/v2/actions_service.proto",
}
