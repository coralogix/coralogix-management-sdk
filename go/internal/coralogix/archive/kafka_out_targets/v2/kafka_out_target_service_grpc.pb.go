// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: com/coralogix/archive/kafka_out_targets/v2/kafka_out_target_service.proto

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
	KafkaOutTargetService_GetKafkaOutTargets_FullMethodName      = "/com.coralogix.archive.kafka_out_targets.v2.KafkaOutTargetService/GetKafkaOutTargets"
	KafkaOutTargetService_UpsertKafkaOutTarget_FullMethodName    = "/com.coralogix.archive.kafka_out_targets.v2.KafkaOutTargetService/UpsertKafkaOutTarget"
	KafkaOutTargetService_DeleteKafkaOutTarget_FullMethodName    = "/com.coralogix.archive.kafka_out_targets.v2.KafkaOutTargetService/DeleteKafkaOutTarget"
	KafkaOutTargetService_SendKafkaOutTestMessage_FullMethodName = "/com.coralogix.archive.kafka_out_targets.v2.KafkaOutTargetService/SendKafkaOutTestMessage"
)

// KafkaOutTargetServiceClient is the client API for KafkaOutTargetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KafkaOutTargetServiceClient interface {
	GetKafkaOutTargets(ctx context.Context, in *GetKafkaOutTargetsRequest, opts ...grpc.CallOption) (*GetKafkaOutTargetsResponse, error)
	UpsertKafkaOutTarget(ctx context.Context, in *UpsertKafkaOutTargetRequest, opts ...grpc.CallOption) (*UpsertKafkaOutTargetResponse, error)
	DeleteKafkaOutTarget(ctx context.Context, in *DeleteKafkaOutTargetRequest, opts ...grpc.CallOption) (*DeleteKafkaOutTargetResponse, error)
	SendKafkaOutTestMessage(ctx context.Context, in *SendKafkaOutTestMessageRequest, opts ...grpc.CallOption) (*SendKafkaOutTestMessageResponse, error)
}

type kafkaOutTargetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKafkaOutTargetServiceClient(cc grpc.ClientConnInterface) KafkaOutTargetServiceClient {
	return &kafkaOutTargetServiceClient{cc}
}

func (c *kafkaOutTargetServiceClient) GetKafkaOutTargets(ctx context.Context, in *GetKafkaOutTargetsRequest, opts ...grpc.CallOption) (*GetKafkaOutTargetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetKafkaOutTargetsResponse)
	err := c.cc.Invoke(ctx, KafkaOutTargetService_GetKafkaOutTargets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaOutTargetServiceClient) UpsertKafkaOutTarget(ctx context.Context, in *UpsertKafkaOutTargetRequest, opts ...grpc.CallOption) (*UpsertKafkaOutTargetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpsertKafkaOutTargetResponse)
	err := c.cc.Invoke(ctx, KafkaOutTargetService_UpsertKafkaOutTarget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaOutTargetServiceClient) DeleteKafkaOutTarget(ctx context.Context, in *DeleteKafkaOutTargetRequest, opts ...grpc.CallOption) (*DeleteKafkaOutTargetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteKafkaOutTargetResponse)
	err := c.cc.Invoke(ctx, KafkaOutTargetService_DeleteKafkaOutTarget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaOutTargetServiceClient) SendKafkaOutTestMessage(ctx context.Context, in *SendKafkaOutTestMessageRequest, opts ...grpc.CallOption) (*SendKafkaOutTestMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendKafkaOutTestMessageResponse)
	err := c.cc.Invoke(ctx, KafkaOutTargetService_SendKafkaOutTestMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KafkaOutTargetServiceServer is the server API for KafkaOutTargetService service.
// All implementations must embed UnimplementedKafkaOutTargetServiceServer
// for forward compatibility.
type KafkaOutTargetServiceServer interface {
	GetKafkaOutTargets(context.Context, *GetKafkaOutTargetsRequest) (*GetKafkaOutTargetsResponse, error)
	UpsertKafkaOutTarget(context.Context, *UpsertKafkaOutTargetRequest) (*UpsertKafkaOutTargetResponse, error)
	DeleteKafkaOutTarget(context.Context, *DeleteKafkaOutTargetRequest) (*DeleteKafkaOutTargetResponse, error)
	SendKafkaOutTestMessage(context.Context, *SendKafkaOutTestMessageRequest) (*SendKafkaOutTestMessageResponse, error)
	mustEmbedUnimplementedKafkaOutTargetServiceServer()
}

// UnimplementedKafkaOutTargetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKafkaOutTargetServiceServer struct{}

func (UnimplementedKafkaOutTargetServiceServer) GetKafkaOutTargets(context.Context, *GetKafkaOutTargetsRequest) (*GetKafkaOutTargetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKafkaOutTargets not implemented")
}
func (UnimplementedKafkaOutTargetServiceServer) UpsertKafkaOutTarget(context.Context, *UpsertKafkaOutTargetRequest) (*UpsertKafkaOutTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertKafkaOutTarget not implemented")
}
func (UnimplementedKafkaOutTargetServiceServer) DeleteKafkaOutTarget(context.Context, *DeleteKafkaOutTargetRequest) (*DeleteKafkaOutTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKafkaOutTarget not implemented")
}
func (UnimplementedKafkaOutTargetServiceServer) SendKafkaOutTestMessage(context.Context, *SendKafkaOutTestMessageRequest) (*SendKafkaOutTestMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendKafkaOutTestMessage not implemented")
}
func (UnimplementedKafkaOutTargetServiceServer) mustEmbedUnimplementedKafkaOutTargetServiceServer() {}
func (UnimplementedKafkaOutTargetServiceServer) testEmbeddedByValue()                               {}

// UnsafeKafkaOutTargetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KafkaOutTargetServiceServer will
// result in compilation errors.
type UnsafeKafkaOutTargetServiceServer interface {
	mustEmbedUnimplementedKafkaOutTargetServiceServer()
}

func RegisterKafkaOutTargetServiceServer(s grpc.ServiceRegistrar, srv KafkaOutTargetServiceServer) {
	// If the following call pancis, it indicates UnimplementedKafkaOutTargetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KafkaOutTargetService_ServiceDesc, srv)
}

func _KafkaOutTargetService_GetKafkaOutTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKafkaOutTargetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaOutTargetServiceServer).GetKafkaOutTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaOutTargetService_GetKafkaOutTargets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaOutTargetServiceServer).GetKafkaOutTargets(ctx, req.(*GetKafkaOutTargetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaOutTargetService_UpsertKafkaOutTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertKafkaOutTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaOutTargetServiceServer).UpsertKafkaOutTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaOutTargetService_UpsertKafkaOutTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaOutTargetServiceServer).UpsertKafkaOutTarget(ctx, req.(*UpsertKafkaOutTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaOutTargetService_DeleteKafkaOutTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKafkaOutTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaOutTargetServiceServer).DeleteKafkaOutTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaOutTargetService_DeleteKafkaOutTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaOutTargetServiceServer).DeleteKafkaOutTarget(ctx, req.(*DeleteKafkaOutTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaOutTargetService_SendKafkaOutTestMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendKafkaOutTestMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaOutTargetServiceServer).SendKafkaOutTestMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaOutTargetService_SendKafkaOutTestMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaOutTargetServiceServer).SendKafkaOutTestMessage(ctx, req.(*SendKafkaOutTestMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KafkaOutTargetService_ServiceDesc is the grpc.ServiceDesc for KafkaOutTargetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KafkaOutTargetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.archive.kafka_out_targets.v2.KafkaOutTargetService",
	HandlerType: (*KafkaOutTargetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKafkaOutTargets",
			Handler:    _KafkaOutTargetService_GetKafkaOutTargets_Handler,
		},
		{
			MethodName: "UpsertKafkaOutTarget",
			Handler:    _KafkaOutTargetService_UpsertKafkaOutTarget_Handler,
		},
		{
			MethodName: "DeleteKafkaOutTarget",
			Handler:    _KafkaOutTargetService_DeleteKafkaOutTarget_Handler,
		},
		{
			MethodName: "SendKafkaOutTestMessage",
			Handler:    _KafkaOutTargetService_SendKafkaOutTestMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/archive/kafka_out_targets/v2/kafka_out_target_service.proto",
}
