// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: com/coralogixapis/notification_center/notifications/v1/testing_service.proto

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
	TestingService_TestConnectorConfig_FullMethodName       = "/com.coralogixapis.notification_center.notifications.v1.TestingService/TestConnectorConfig"
	TestingService_TestExistingConnector_FullMethodName     = "/com.coralogixapis.notification_center.notifications.v1.TestingService/TestExistingConnector"
	TestingService_TestPresetConfig_FullMethodName          = "/com.coralogixapis.notification_center.notifications.v1.TestingService/TestPresetConfig"
	TestingService_TestExistingPreset_FullMethodName        = "/com.coralogixapis.notification_center.notifications.v1.TestingService/TestExistingPreset"
	TestingService_TestDestination_FullMethodName           = "/com.coralogixapis.notification_center.notifications.v1.TestingService/TestDestination"
	TestingService_TestTemplateRender_FullMethodName        = "/com.coralogixapis.notification_center.notifications.v1.TestingService/TestTemplateRender"
	TestingService_TestRoutingConditionValid_FullMethodName = "/com.coralogixapis.notification_center.notifications.v1.TestingService/TestRoutingConditionValid"
)

// TestingServiceClient is the client API for TestingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Public API for testing Notification Center configurations including connectors, presets, and templates.
type TestingServiceClient interface {
	TestConnectorConfig(ctx context.Context, in *TestConnectorConfigRequest, opts ...grpc.CallOption) (*TestConnectorConfigResponse, error)
	TestExistingConnector(ctx context.Context, in *TestExistingConnectorRequest, opts ...grpc.CallOption) (*TestExistingConnectorResponse, error)
	TestPresetConfig(ctx context.Context, in *TestPresetConfigRequest, opts ...grpc.CallOption) (*TestPresetConfigResponse, error)
	TestExistingPreset(ctx context.Context, in *TestExistingPresetRequest, opts ...grpc.CallOption) (*TestExistingPresetResponse, error)
	TestDestination(ctx context.Context, in *TestDestinationRequest, opts ...grpc.CallOption) (*TestDestinationResponse, error)
	TestTemplateRender(ctx context.Context, in *TestTemplateRenderRequest, opts ...grpc.CallOption) (*TestTemplateRenderResponse, error)
	TestRoutingConditionValid(ctx context.Context, in *TestRoutingConditionValidRequest, opts ...grpc.CallOption) (*TestRoutingConditionValidResponse, error)
}

type testingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestingServiceClient(cc grpc.ClientConnInterface) TestingServiceClient {
	return &testingServiceClient{cc}
}

func (c *testingServiceClient) TestConnectorConfig(ctx context.Context, in *TestConnectorConfigRequest, opts ...grpc.CallOption) (*TestConnectorConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestConnectorConfigResponse)
	err := c.cc.Invoke(ctx, TestingService_TestConnectorConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testingServiceClient) TestExistingConnector(ctx context.Context, in *TestExistingConnectorRequest, opts ...grpc.CallOption) (*TestExistingConnectorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestExistingConnectorResponse)
	err := c.cc.Invoke(ctx, TestingService_TestExistingConnector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testingServiceClient) TestPresetConfig(ctx context.Context, in *TestPresetConfigRequest, opts ...grpc.CallOption) (*TestPresetConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestPresetConfigResponse)
	err := c.cc.Invoke(ctx, TestingService_TestPresetConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testingServiceClient) TestExistingPreset(ctx context.Context, in *TestExistingPresetRequest, opts ...grpc.CallOption) (*TestExistingPresetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestExistingPresetResponse)
	err := c.cc.Invoke(ctx, TestingService_TestExistingPreset_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testingServiceClient) TestDestination(ctx context.Context, in *TestDestinationRequest, opts ...grpc.CallOption) (*TestDestinationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestDestinationResponse)
	err := c.cc.Invoke(ctx, TestingService_TestDestination_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testingServiceClient) TestTemplateRender(ctx context.Context, in *TestTemplateRenderRequest, opts ...grpc.CallOption) (*TestTemplateRenderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestTemplateRenderResponse)
	err := c.cc.Invoke(ctx, TestingService_TestTemplateRender_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testingServiceClient) TestRoutingConditionValid(ctx context.Context, in *TestRoutingConditionValidRequest, opts ...grpc.CallOption) (*TestRoutingConditionValidResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestRoutingConditionValidResponse)
	err := c.cc.Invoke(ctx, TestingService_TestRoutingConditionValid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestingServiceServer is the server API for TestingService service.
// All implementations must embed UnimplementedTestingServiceServer
// for forward compatibility.
//
// Public API for testing Notification Center configurations including connectors, presets, and templates.
type TestingServiceServer interface {
	TestConnectorConfig(context.Context, *TestConnectorConfigRequest) (*TestConnectorConfigResponse, error)
	TestExistingConnector(context.Context, *TestExistingConnectorRequest) (*TestExistingConnectorResponse, error)
	TestPresetConfig(context.Context, *TestPresetConfigRequest) (*TestPresetConfigResponse, error)
	TestExistingPreset(context.Context, *TestExistingPresetRequest) (*TestExistingPresetResponse, error)
	TestDestination(context.Context, *TestDestinationRequest) (*TestDestinationResponse, error)
	TestTemplateRender(context.Context, *TestTemplateRenderRequest) (*TestTemplateRenderResponse, error)
	TestRoutingConditionValid(context.Context, *TestRoutingConditionValidRequest) (*TestRoutingConditionValidResponse, error)
	mustEmbedUnimplementedTestingServiceServer()
}

// UnimplementedTestingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTestingServiceServer struct{}

func (UnimplementedTestingServiceServer) TestConnectorConfig(context.Context, *TestConnectorConfigRequest) (*TestConnectorConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestConnectorConfig not implemented")
}
func (UnimplementedTestingServiceServer) TestExistingConnector(context.Context, *TestExistingConnectorRequest) (*TestExistingConnectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestExistingConnector not implemented")
}
func (UnimplementedTestingServiceServer) TestPresetConfig(context.Context, *TestPresetConfigRequest) (*TestPresetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestPresetConfig not implemented")
}
func (UnimplementedTestingServiceServer) TestExistingPreset(context.Context, *TestExistingPresetRequest) (*TestExistingPresetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestExistingPreset not implemented")
}
func (UnimplementedTestingServiceServer) TestDestination(context.Context, *TestDestinationRequest) (*TestDestinationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestDestination not implemented")
}
func (UnimplementedTestingServiceServer) TestTemplateRender(context.Context, *TestTemplateRenderRequest) (*TestTemplateRenderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestTemplateRender not implemented")
}
func (UnimplementedTestingServiceServer) TestRoutingConditionValid(context.Context, *TestRoutingConditionValidRequest) (*TestRoutingConditionValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestRoutingConditionValid not implemented")
}
func (UnimplementedTestingServiceServer) mustEmbedUnimplementedTestingServiceServer() {}
func (UnimplementedTestingServiceServer) testEmbeddedByValue()                        {}

// UnsafeTestingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestingServiceServer will
// result in compilation errors.
type UnsafeTestingServiceServer interface {
	mustEmbedUnimplementedTestingServiceServer()
}

func RegisterTestingServiceServer(s grpc.ServiceRegistrar, srv TestingServiceServer) {
	// If the following call pancis, it indicates UnimplementedTestingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TestingService_ServiceDesc, srv)
}

func _TestingService_TestConnectorConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestConnectorConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestingServiceServer).TestConnectorConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestingService_TestConnectorConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestingServiceServer).TestConnectorConfig(ctx, req.(*TestConnectorConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestingService_TestExistingConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestExistingConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestingServiceServer).TestExistingConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestingService_TestExistingConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestingServiceServer).TestExistingConnector(ctx, req.(*TestExistingConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestingService_TestPresetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestPresetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestingServiceServer).TestPresetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestingService_TestPresetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestingServiceServer).TestPresetConfig(ctx, req.(*TestPresetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestingService_TestExistingPreset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestExistingPresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestingServiceServer).TestExistingPreset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestingService_TestExistingPreset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestingServiceServer).TestExistingPreset(ctx, req.(*TestExistingPresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestingService_TestDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestDestinationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestingServiceServer).TestDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestingService_TestDestination_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestingServiceServer).TestDestination(ctx, req.(*TestDestinationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestingService_TestTemplateRender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestTemplateRenderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestingServiceServer).TestTemplateRender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestingService_TestTemplateRender_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestingServiceServer).TestTemplateRender(ctx, req.(*TestTemplateRenderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestingService_TestRoutingConditionValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRoutingConditionValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestingServiceServer).TestRoutingConditionValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestingService_TestRoutingConditionValid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestingServiceServer).TestRoutingConditionValid(ctx, req.(*TestRoutingConditionValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestingService_ServiceDesc is the grpc.ServiceDesc for TestingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.notification_center.notifications.v1.TestingService",
	HandlerType: (*TestingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestConnectorConfig",
			Handler:    _TestingService_TestConnectorConfig_Handler,
		},
		{
			MethodName: "TestExistingConnector",
			Handler:    _TestingService_TestExistingConnector_Handler,
		},
		{
			MethodName: "TestPresetConfig",
			Handler:    _TestingService_TestPresetConfig_Handler,
		},
		{
			MethodName: "TestExistingPreset",
			Handler:    _TestingService_TestExistingPreset_Handler,
		},
		{
			MethodName: "TestDestination",
			Handler:    _TestingService_TestDestination_Handler,
		},
		{
			MethodName: "TestTemplateRender",
			Handler:    _TestingService_TestTemplateRender_Handler,
		},
		{
			MethodName: "TestRoutingConditionValid",
			Handler:    _TestingService_TestRoutingConditionValid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/notification_center/notifications/v1/testing_service.proto",
}
