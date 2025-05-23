// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: com/coralogixapis/notification_center/remote_actions/v1/remote_actions_service.proto

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
	RemoteActionsService_TriggerRemoteAction_FullMethodName = "/com.coralogixapis.notification_center.remote_actions.v1.RemoteActionsService/TriggerRemoteAction"
)

// RemoteActionsServiceClient is the client API for RemoteActionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Public API for Remote Actions execution.
// Note: This API is in development phase and subject to change.
type RemoteActionsServiceClient interface {
	TriggerRemoteAction(ctx context.Context, in *TriggerRemoteActionRequest, opts ...grpc.CallOption) (*TriggerRemoteActionResponse, error)
}

type remoteActionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteActionsServiceClient(cc grpc.ClientConnInterface) RemoteActionsServiceClient {
	return &remoteActionsServiceClient{cc}
}

func (c *remoteActionsServiceClient) TriggerRemoteAction(ctx context.Context, in *TriggerRemoteActionRequest, opts ...grpc.CallOption) (*TriggerRemoteActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TriggerRemoteActionResponse)
	err := c.cc.Invoke(ctx, RemoteActionsService_TriggerRemoteAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteActionsServiceServer is the server API for RemoteActionsService service.
// All implementations must embed UnimplementedRemoteActionsServiceServer
// for forward compatibility.
//
// Public API for Remote Actions execution.
// Note: This API is in development phase and subject to change.
type RemoteActionsServiceServer interface {
	TriggerRemoteAction(context.Context, *TriggerRemoteActionRequest) (*TriggerRemoteActionResponse, error)
	mustEmbedUnimplementedRemoteActionsServiceServer()
}

// UnimplementedRemoteActionsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRemoteActionsServiceServer struct{}

func (UnimplementedRemoteActionsServiceServer) TriggerRemoteAction(context.Context, *TriggerRemoteActionRequest) (*TriggerRemoteActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerRemoteAction not implemented")
}
func (UnimplementedRemoteActionsServiceServer) mustEmbedUnimplementedRemoteActionsServiceServer() {}
func (UnimplementedRemoteActionsServiceServer) testEmbeddedByValue()                              {}

// UnsafeRemoteActionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteActionsServiceServer will
// result in compilation errors.
type UnsafeRemoteActionsServiceServer interface {
	mustEmbedUnimplementedRemoteActionsServiceServer()
}

func RegisterRemoteActionsServiceServer(s grpc.ServiceRegistrar, srv RemoteActionsServiceServer) {
	// If the following call pancis, it indicates UnimplementedRemoteActionsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RemoteActionsService_ServiceDesc, srv)
}

func _RemoteActionsService_TriggerRemoteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerRemoteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteActionsServiceServer).TriggerRemoteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteActionsService_TriggerRemoteAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteActionsServiceServer).TriggerRemoteAction(ctx, req.(*TriggerRemoteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoteActionsService_ServiceDesc is the grpc.ServiceDesc for RemoteActionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoteActionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.notification_center.remote_actions.v1.RemoteActionsService",
	HandlerType: (*RemoteActionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TriggerRemoteAction",
			Handler:    _RemoteActionsService_TriggerRemoteAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/notification_center/remote_actions/v1/remote_actions_service.proto",
}
