// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: com/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_rule_service.proto

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
	AlertSchedulerRuleService_GetAlertSchedulerRule_FullMethodName        = "/com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService/GetAlertSchedulerRule"
	AlertSchedulerRuleService_CreateAlertSchedulerRule_FullMethodName     = "/com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService/CreateAlertSchedulerRule"
	AlertSchedulerRuleService_UpdateAlertSchedulerRule_FullMethodName     = "/com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService/UpdateAlertSchedulerRule"
	AlertSchedulerRuleService_DeleteAlertSchedulerRule_FullMethodName     = "/com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService/DeleteAlertSchedulerRule"
	AlertSchedulerRuleService_GetBulkAlertSchedulerRule_FullMethodName    = "/com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService/GetBulkAlertSchedulerRule"
	AlertSchedulerRuleService_CreateBulkAlertSchedulerRule_FullMethodName = "/com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService/CreateBulkAlertSchedulerRule"
	AlertSchedulerRuleService_UpdateBulkAlertSchedulerRule_FullMethodName = "/com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService/UpdateBulkAlertSchedulerRule"
)

// AlertSchedulerRuleServiceClient is the client API for AlertSchedulerRuleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertSchedulerRuleServiceClient interface {
	GetAlertSchedulerRule(ctx context.Context, in *GetAlertSchedulerRuleRequest, opts ...grpc.CallOption) (*GetAlertSchedulerRuleResponse, error)
	CreateAlertSchedulerRule(ctx context.Context, in *CreateAlertSchedulerRuleRequest, opts ...grpc.CallOption) (*CreateAlertSchedulerRuleResponse, error)
	UpdateAlertSchedulerRule(ctx context.Context, in *UpdateAlertSchedulerRuleRequest, opts ...grpc.CallOption) (*UpdateAlertSchedulerRuleResponse, error)
	DeleteAlertSchedulerRule(ctx context.Context, in *DeleteAlertSchedulerRuleRequest, opts ...grpc.CallOption) (*DeleteAlertSchedulerRuleResponse, error)
	GetBulkAlertSchedulerRule(ctx context.Context, in *GetBulkAlertSchedulerRuleRequest, opts ...grpc.CallOption) (*GetBulkAlertSchedulerRuleResponse, error)
	CreateBulkAlertSchedulerRule(ctx context.Context, in *CreateBulkAlertSchedulerRuleRequest, opts ...grpc.CallOption) (*CreateBulkAlertSchedulerRuleResponse, error)
	UpdateBulkAlertSchedulerRule(ctx context.Context, in *UpdateBulkAlertSchedulerRuleRequest, opts ...grpc.CallOption) (*UpdateBulkAlertSchedulerRuleResponse, error)
}

type alertSchedulerRuleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertSchedulerRuleServiceClient(cc grpc.ClientConnInterface) AlertSchedulerRuleServiceClient {
	return &alertSchedulerRuleServiceClient{cc}
}

func (c *alertSchedulerRuleServiceClient) GetAlertSchedulerRule(ctx context.Context, in *GetAlertSchedulerRuleRequest, opts ...grpc.CallOption) (*GetAlertSchedulerRuleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAlertSchedulerRuleResponse)
	err := c.cc.Invoke(ctx, AlertSchedulerRuleService_GetAlertSchedulerRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertSchedulerRuleServiceClient) CreateAlertSchedulerRule(ctx context.Context, in *CreateAlertSchedulerRuleRequest, opts ...grpc.CallOption) (*CreateAlertSchedulerRuleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAlertSchedulerRuleResponse)
	err := c.cc.Invoke(ctx, AlertSchedulerRuleService_CreateAlertSchedulerRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertSchedulerRuleServiceClient) UpdateAlertSchedulerRule(ctx context.Context, in *UpdateAlertSchedulerRuleRequest, opts ...grpc.CallOption) (*UpdateAlertSchedulerRuleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAlertSchedulerRuleResponse)
	err := c.cc.Invoke(ctx, AlertSchedulerRuleService_UpdateAlertSchedulerRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertSchedulerRuleServiceClient) DeleteAlertSchedulerRule(ctx context.Context, in *DeleteAlertSchedulerRuleRequest, opts ...grpc.CallOption) (*DeleteAlertSchedulerRuleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteAlertSchedulerRuleResponse)
	err := c.cc.Invoke(ctx, AlertSchedulerRuleService_DeleteAlertSchedulerRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertSchedulerRuleServiceClient) GetBulkAlertSchedulerRule(ctx context.Context, in *GetBulkAlertSchedulerRuleRequest, opts ...grpc.CallOption) (*GetBulkAlertSchedulerRuleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBulkAlertSchedulerRuleResponse)
	err := c.cc.Invoke(ctx, AlertSchedulerRuleService_GetBulkAlertSchedulerRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertSchedulerRuleServiceClient) CreateBulkAlertSchedulerRule(ctx context.Context, in *CreateBulkAlertSchedulerRuleRequest, opts ...grpc.CallOption) (*CreateBulkAlertSchedulerRuleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBulkAlertSchedulerRuleResponse)
	err := c.cc.Invoke(ctx, AlertSchedulerRuleService_CreateBulkAlertSchedulerRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertSchedulerRuleServiceClient) UpdateBulkAlertSchedulerRule(ctx context.Context, in *UpdateBulkAlertSchedulerRuleRequest, opts ...grpc.CallOption) (*UpdateBulkAlertSchedulerRuleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBulkAlertSchedulerRuleResponse)
	err := c.cc.Invoke(ctx, AlertSchedulerRuleService_UpdateBulkAlertSchedulerRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertSchedulerRuleServiceServer is the server API for AlertSchedulerRuleService service.
// All implementations must embed UnimplementedAlertSchedulerRuleServiceServer
// for forward compatibility.
type AlertSchedulerRuleServiceServer interface {
	GetAlertSchedulerRule(context.Context, *GetAlertSchedulerRuleRequest) (*GetAlertSchedulerRuleResponse, error)
	CreateAlertSchedulerRule(context.Context, *CreateAlertSchedulerRuleRequest) (*CreateAlertSchedulerRuleResponse, error)
	UpdateAlertSchedulerRule(context.Context, *UpdateAlertSchedulerRuleRequest) (*UpdateAlertSchedulerRuleResponse, error)
	DeleteAlertSchedulerRule(context.Context, *DeleteAlertSchedulerRuleRequest) (*DeleteAlertSchedulerRuleResponse, error)
	GetBulkAlertSchedulerRule(context.Context, *GetBulkAlertSchedulerRuleRequest) (*GetBulkAlertSchedulerRuleResponse, error)
	CreateBulkAlertSchedulerRule(context.Context, *CreateBulkAlertSchedulerRuleRequest) (*CreateBulkAlertSchedulerRuleResponse, error)
	UpdateBulkAlertSchedulerRule(context.Context, *UpdateBulkAlertSchedulerRuleRequest) (*UpdateBulkAlertSchedulerRuleResponse, error)
	mustEmbedUnimplementedAlertSchedulerRuleServiceServer()
}

// UnimplementedAlertSchedulerRuleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAlertSchedulerRuleServiceServer struct{}

func (UnimplementedAlertSchedulerRuleServiceServer) GetAlertSchedulerRule(context.Context, *GetAlertSchedulerRuleRequest) (*GetAlertSchedulerRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertSchedulerRule not implemented")
}
func (UnimplementedAlertSchedulerRuleServiceServer) CreateAlertSchedulerRule(context.Context, *CreateAlertSchedulerRuleRequest) (*CreateAlertSchedulerRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlertSchedulerRule not implemented")
}
func (UnimplementedAlertSchedulerRuleServiceServer) UpdateAlertSchedulerRule(context.Context, *UpdateAlertSchedulerRuleRequest) (*UpdateAlertSchedulerRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAlertSchedulerRule not implemented")
}
func (UnimplementedAlertSchedulerRuleServiceServer) DeleteAlertSchedulerRule(context.Context, *DeleteAlertSchedulerRuleRequest) (*DeleteAlertSchedulerRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlertSchedulerRule not implemented")
}
func (UnimplementedAlertSchedulerRuleServiceServer) GetBulkAlertSchedulerRule(context.Context, *GetBulkAlertSchedulerRuleRequest) (*GetBulkAlertSchedulerRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBulkAlertSchedulerRule not implemented")
}
func (UnimplementedAlertSchedulerRuleServiceServer) CreateBulkAlertSchedulerRule(context.Context, *CreateBulkAlertSchedulerRuleRequest) (*CreateBulkAlertSchedulerRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBulkAlertSchedulerRule not implemented")
}
func (UnimplementedAlertSchedulerRuleServiceServer) UpdateBulkAlertSchedulerRule(context.Context, *UpdateBulkAlertSchedulerRuleRequest) (*UpdateBulkAlertSchedulerRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBulkAlertSchedulerRule not implemented")
}
func (UnimplementedAlertSchedulerRuleServiceServer) mustEmbedUnimplementedAlertSchedulerRuleServiceServer() {
}
func (UnimplementedAlertSchedulerRuleServiceServer) testEmbeddedByValue() {}

// UnsafeAlertSchedulerRuleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertSchedulerRuleServiceServer will
// result in compilation errors.
type UnsafeAlertSchedulerRuleServiceServer interface {
	mustEmbedUnimplementedAlertSchedulerRuleServiceServer()
}

func RegisterAlertSchedulerRuleServiceServer(s grpc.ServiceRegistrar, srv AlertSchedulerRuleServiceServer) {
	// If the following call pancis, it indicates UnimplementedAlertSchedulerRuleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AlertSchedulerRuleService_ServiceDesc, srv)
}

func _AlertSchedulerRuleService_GetAlertSchedulerRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertSchedulerRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertSchedulerRuleServiceServer).GetAlertSchedulerRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertSchedulerRuleService_GetAlertSchedulerRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertSchedulerRuleServiceServer).GetAlertSchedulerRule(ctx, req.(*GetAlertSchedulerRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertSchedulerRuleService_CreateAlertSchedulerRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlertSchedulerRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertSchedulerRuleServiceServer).CreateAlertSchedulerRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertSchedulerRuleService_CreateAlertSchedulerRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertSchedulerRuleServiceServer).CreateAlertSchedulerRule(ctx, req.(*CreateAlertSchedulerRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertSchedulerRuleService_UpdateAlertSchedulerRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlertSchedulerRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertSchedulerRuleServiceServer).UpdateAlertSchedulerRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertSchedulerRuleService_UpdateAlertSchedulerRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertSchedulerRuleServiceServer).UpdateAlertSchedulerRule(ctx, req.(*UpdateAlertSchedulerRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertSchedulerRuleService_DeleteAlertSchedulerRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAlertSchedulerRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertSchedulerRuleServiceServer).DeleteAlertSchedulerRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertSchedulerRuleService_DeleteAlertSchedulerRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertSchedulerRuleServiceServer).DeleteAlertSchedulerRule(ctx, req.(*DeleteAlertSchedulerRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertSchedulerRuleService_GetBulkAlertSchedulerRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBulkAlertSchedulerRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertSchedulerRuleServiceServer).GetBulkAlertSchedulerRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertSchedulerRuleService_GetBulkAlertSchedulerRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertSchedulerRuleServiceServer).GetBulkAlertSchedulerRule(ctx, req.(*GetBulkAlertSchedulerRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertSchedulerRuleService_CreateBulkAlertSchedulerRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBulkAlertSchedulerRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertSchedulerRuleServiceServer).CreateBulkAlertSchedulerRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertSchedulerRuleService_CreateBulkAlertSchedulerRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertSchedulerRuleServiceServer).CreateBulkAlertSchedulerRule(ctx, req.(*CreateBulkAlertSchedulerRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertSchedulerRuleService_UpdateBulkAlertSchedulerRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBulkAlertSchedulerRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertSchedulerRuleServiceServer).UpdateBulkAlertSchedulerRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertSchedulerRuleService_UpdateBulkAlertSchedulerRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertSchedulerRuleServiceServer).UpdateBulkAlertSchedulerRule(ctx, req.(*UpdateBulkAlertSchedulerRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AlertSchedulerRuleService_ServiceDesc is the grpc.ServiceDesc for AlertSchedulerRuleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlertSchedulerRuleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRuleService",
	HandlerType: (*AlertSchedulerRuleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAlertSchedulerRule",
			Handler:    _AlertSchedulerRuleService_GetAlertSchedulerRule_Handler,
		},
		{
			MethodName: "CreateAlertSchedulerRule",
			Handler:    _AlertSchedulerRuleService_CreateAlertSchedulerRule_Handler,
		},
		{
			MethodName: "UpdateAlertSchedulerRule",
			Handler:    _AlertSchedulerRuleService_UpdateAlertSchedulerRule_Handler,
		},
		{
			MethodName: "DeleteAlertSchedulerRule",
			Handler:    _AlertSchedulerRuleService_DeleteAlertSchedulerRule_Handler,
		},
		{
			MethodName: "GetBulkAlertSchedulerRule",
			Handler:    _AlertSchedulerRuleService_GetBulkAlertSchedulerRule_Handler,
		},
		{
			MethodName: "CreateBulkAlertSchedulerRule",
			Handler:    _AlertSchedulerRuleService_CreateBulkAlertSchedulerRule_Handler,
		},
		{
			MethodName: "UpdateBulkAlertSchedulerRule",
			Handler:    _AlertSchedulerRuleService_UpdateBulkAlertSchedulerRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_rule_service.proto",
}
