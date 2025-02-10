// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: com/coralogixapis/notification_center/presets/v1/presets_service.proto

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
	PresetsService_CreateCustomPreset_FullMethodName            = "/com.coralogixapis.notification_center.presets.v1.PresetsService/CreateCustomPreset"
	PresetsService_ReplaceCustomPreset_FullMethodName           = "/com.coralogixapis.notification_center.presets.v1.PresetsService/ReplaceCustomPreset"
	PresetsService_DeleteCustomPreset_FullMethodName            = "/com.coralogixapis.notification_center.presets.v1.PresetsService/DeleteCustomPreset"
	PresetsService_SetCustomPresetAsDefault_FullMethodName      = "/com.coralogixapis.notification_center.presets.v1.PresetsService/SetCustomPresetAsDefault"
	PresetsService_SetPresetAsDefault_FullMethodName            = "/com.coralogixapis.notification_center.presets.v1.PresetsService/SetPresetAsDefault"
	PresetsService_GetPreset_FullMethodName                     = "/com.coralogixapis.notification_center.presets.v1.PresetsService/GetPreset"
	PresetsService_ListPresetSummaries_FullMethodName           = "/com.coralogixapis.notification_center.presets.v1.PresetsService/ListPresetSummaries"
	PresetsService_BatchGetPresets_FullMethodName               = "/com.coralogixapis.notification_center.presets.v1.PresetsService/BatchGetPresets"
	PresetsService_GetDefaultPresetSummary_FullMethodName       = "/com.coralogixapis.notification_center.presets.v1.PresetsService/GetDefaultPresetSummary"
	PresetsService_GetSystemDefaultPresetSummary_FullMethodName = "/com.coralogixapis.notification_center.presets.v1.PresetsService/GetSystemDefaultPresetSummary"
)

// PresetsServiceClient is the client API for PresetsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Public API that allows retrieving available presets and managing custom presets.
type PresetsServiceClient interface {
	CreateCustomPreset(ctx context.Context, in *CreateCustomPresetRequest, opts ...grpc.CallOption) (*CreateCustomPresetResponse, error)
	ReplaceCustomPreset(ctx context.Context, in *ReplaceCustomPresetRequest, opts ...grpc.CallOption) (*ReplaceCustomPresetResponse, error)
	DeleteCustomPreset(ctx context.Context, in *DeleteCustomPresetRequest, opts ...grpc.CallOption) (*DeleteCustomPresetResponse, error)
	// Deprecated: Do not use.
	SetCustomPresetAsDefault(ctx context.Context, in *SetCustomPresetAsDefaultRequest, opts ...grpc.CallOption) (*SetCustomPresetAsDefaultResponse, error)
	SetPresetAsDefault(ctx context.Context, in *SetPresetAsDefaultRequest, opts ...grpc.CallOption) (*SetPresetAsDefaultResponse, error)
	GetPreset(ctx context.Context, in *GetPresetRequest, opts ...grpc.CallOption) (*GetPresetResponse, error)
	ListPresetSummaries(ctx context.Context, in *ListPresetSummariesRequest, opts ...grpc.CallOption) (*ListPresetSummariesResponse, error)
	BatchGetPresets(ctx context.Context, in *BatchGetPresetsRequest, opts ...grpc.CallOption) (*BatchGetPresetsResponse, error)
	GetDefaultPresetSummary(ctx context.Context, in *GetDefaultPresetSummaryRequest, opts ...grpc.CallOption) (*GetDefaultPresetSummaryResponse, error)
	GetSystemDefaultPresetSummary(ctx context.Context, in *GetSystemDefaultPresetSummaryRequest, opts ...grpc.CallOption) (*GetSystemDefaultPresetSummaryResponse, error)
}

type presetsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPresetsServiceClient(cc grpc.ClientConnInterface) PresetsServiceClient {
	return &presetsServiceClient{cc}
}

func (c *presetsServiceClient) CreateCustomPreset(ctx context.Context, in *CreateCustomPresetRequest, opts ...grpc.CallOption) (*CreateCustomPresetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCustomPresetResponse)
	err := c.cc.Invoke(ctx, PresetsService_CreateCustomPreset_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetsServiceClient) ReplaceCustomPreset(ctx context.Context, in *ReplaceCustomPresetRequest, opts ...grpc.CallOption) (*ReplaceCustomPresetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReplaceCustomPresetResponse)
	err := c.cc.Invoke(ctx, PresetsService_ReplaceCustomPreset_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetsServiceClient) DeleteCustomPreset(ctx context.Context, in *DeleteCustomPresetRequest, opts ...grpc.CallOption) (*DeleteCustomPresetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCustomPresetResponse)
	err := c.cc.Invoke(ctx, PresetsService_DeleteCustomPreset_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *presetsServiceClient) SetCustomPresetAsDefault(ctx context.Context, in *SetCustomPresetAsDefaultRequest, opts ...grpc.CallOption) (*SetCustomPresetAsDefaultResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetCustomPresetAsDefaultResponse)
	err := c.cc.Invoke(ctx, PresetsService_SetCustomPresetAsDefault_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetsServiceClient) SetPresetAsDefault(ctx context.Context, in *SetPresetAsDefaultRequest, opts ...grpc.CallOption) (*SetPresetAsDefaultResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetPresetAsDefaultResponse)
	err := c.cc.Invoke(ctx, PresetsService_SetPresetAsDefault_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetsServiceClient) GetPreset(ctx context.Context, in *GetPresetRequest, opts ...grpc.CallOption) (*GetPresetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPresetResponse)
	err := c.cc.Invoke(ctx, PresetsService_GetPreset_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetsServiceClient) ListPresetSummaries(ctx context.Context, in *ListPresetSummariesRequest, opts ...grpc.CallOption) (*ListPresetSummariesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPresetSummariesResponse)
	err := c.cc.Invoke(ctx, PresetsService_ListPresetSummaries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetsServiceClient) BatchGetPresets(ctx context.Context, in *BatchGetPresetsRequest, opts ...grpc.CallOption) (*BatchGetPresetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchGetPresetsResponse)
	err := c.cc.Invoke(ctx, PresetsService_BatchGetPresets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetsServiceClient) GetDefaultPresetSummary(ctx context.Context, in *GetDefaultPresetSummaryRequest, opts ...grpc.CallOption) (*GetDefaultPresetSummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDefaultPresetSummaryResponse)
	err := c.cc.Invoke(ctx, PresetsService_GetDefaultPresetSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetsServiceClient) GetSystemDefaultPresetSummary(ctx context.Context, in *GetSystemDefaultPresetSummaryRequest, opts ...grpc.CallOption) (*GetSystemDefaultPresetSummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSystemDefaultPresetSummaryResponse)
	err := c.cc.Invoke(ctx, PresetsService_GetSystemDefaultPresetSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PresetsServiceServer is the server API for PresetsService service.
// All implementations must embed UnimplementedPresetsServiceServer
// for forward compatibility.
//
// Public API that allows retrieving available presets and managing custom presets.
type PresetsServiceServer interface {
	CreateCustomPreset(context.Context, *CreateCustomPresetRequest) (*CreateCustomPresetResponse, error)
	ReplaceCustomPreset(context.Context, *ReplaceCustomPresetRequest) (*ReplaceCustomPresetResponse, error)
	DeleteCustomPreset(context.Context, *DeleteCustomPresetRequest) (*DeleteCustomPresetResponse, error)
	// Deprecated: Do not use.
	SetCustomPresetAsDefault(context.Context, *SetCustomPresetAsDefaultRequest) (*SetCustomPresetAsDefaultResponse, error)
	SetPresetAsDefault(context.Context, *SetPresetAsDefaultRequest) (*SetPresetAsDefaultResponse, error)
	GetPreset(context.Context, *GetPresetRequest) (*GetPresetResponse, error)
	ListPresetSummaries(context.Context, *ListPresetSummariesRequest) (*ListPresetSummariesResponse, error)
	BatchGetPresets(context.Context, *BatchGetPresetsRequest) (*BatchGetPresetsResponse, error)
	GetDefaultPresetSummary(context.Context, *GetDefaultPresetSummaryRequest) (*GetDefaultPresetSummaryResponse, error)
	GetSystemDefaultPresetSummary(context.Context, *GetSystemDefaultPresetSummaryRequest) (*GetSystemDefaultPresetSummaryResponse, error)
	mustEmbedUnimplementedPresetsServiceServer()
}

// UnimplementedPresetsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPresetsServiceServer struct{}

func (UnimplementedPresetsServiceServer) CreateCustomPreset(context.Context, *CreateCustomPresetRequest) (*CreateCustomPresetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomPreset not implemented")
}
func (UnimplementedPresetsServiceServer) ReplaceCustomPreset(context.Context, *ReplaceCustomPresetRequest) (*ReplaceCustomPresetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceCustomPreset not implemented")
}
func (UnimplementedPresetsServiceServer) DeleteCustomPreset(context.Context, *DeleteCustomPresetRequest) (*DeleteCustomPresetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomPreset not implemented")
}
func (UnimplementedPresetsServiceServer) SetCustomPresetAsDefault(context.Context, *SetCustomPresetAsDefaultRequest) (*SetCustomPresetAsDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCustomPresetAsDefault not implemented")
}
func (UnimplementedPresetsServiceServer) SetPresetAsDefault(context.Context, *SetPresetAsDefaultRequest) (*SetPresetAsDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPresetAsDefault not implemented")
}
func (UnimplementedPresetsServiceServer) GetPreset(context.Context, *GetPresetRequest) (*GetPresetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreset not implemented")
}
func (UnimplementedPresetsServiceServer) ListPresetSummaries(context.Context, *ListPresetSummariesRequest) (*ListPresetSummariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPresetSummaries not implemented")
}
func (UnimplementedPresetsServiceServer) BatchGetPresets(context.Context, *BatchGetPresetsRequest) (*BatchGetPresetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetPresets not implemented")
}
func (UnimplementedPresetsServiceServer) GetDefaultPresetSummary(context.Context, *GetDefaultPresetSummaryRequest) (*GetDefaultPresetSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultPresetSummary not implemented")
}
func (UnimplementedPresetsServiceServer) GetSystemDefaultPresetSummary(context.Context, *GetSystemDefaultPresetSummaryRequest) (*GetSystemDefaultPresetSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemDefaultPresetSummary not implemented")
}
func (UnimplementedPresetsServiceServer) mustEmbedUnimplementedPresetsServiceServer() {}
func (UnimplementedPresetsServiceServer) testEmbeddedByValue()                        {}

// UnsafePresetsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PresetsServiceServer will
// result in compilation errors.
type UnsafePresetsServiceServer interface {
	mustEmbedUnimplementedPresetsServiceServer()
}

func RegisterPresetsServiceServer(s grpc.ServiceRegistrar, srv PresetsServiceServer) {
	// If the following call pancis, it indicates UnimplementedPresetsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PresetsService_ServiceDesc, srv)
}

func _PresetsService_CreateCustomPreset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomPresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetsServiceServer).CreateCustomPreset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetsService_CreateCustomPreset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetsServiceServer).CreateCustomPreset(ctx, req.(*CreateCustomPresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresetsService_ReplaceCustomPreset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceCustomPresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetsServiceServer).ReplaceCustomPreset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetsService_ReplaceCustomPreset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetsServiceServer).ReplaceCustomPreset(ctx, req.(*ReplaceCustomPresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresetsService_DeleteCustomPreset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCustomPresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetsServiceServer).DeleteCustomPreset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetsService_DeleteCustomPreset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetsServiceServer).DeleteCustomPreset(ctx, req.(*DeleteCustomPresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresetsService_SetCustomPresetAsDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCustomPresetAsDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetsServiceServer).SetCustomPresetAsDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetsService_SetCustomPresetAsDefault_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetsServiceServer).SetCustomPresetAsDefault(ctx, req.(*SetCustomPresetAsDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresetsService_SetPresetAsDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPresetAsDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetsServiceServer).SetPresetAsDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetsService_SetPresetAsDefault_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetsServiceServer).SetPresetAsDefault(ctx, req.(*SetPresetAsDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresetsService_GetPreset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetsServiceServer).GetPreset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetsService_GetPreset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetsServiceServer).GetPreset(ctx, req.(*GetPresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresetsService_ListPresetSummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPresetSummariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetsServiceServer).ListPresetSummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetsService_ListPresetSummaries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetsServiceServer).ListPresetSummaries(ctx, req.(*ListPresetSummariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresetsService_BatchGetPresets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetPresetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetsServiceServer).BatchGetPresets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetsService_BatchGetPresets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetsServiceServer).BatchGetPresets(ctx, req.(*BatchGetPresetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresetsService_GetDefaultPresetSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultPresetSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetsServiceServer).GetDefaultPresetSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetsService_GetDefaultPresetSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetsServiceServer).GetDefaultPresetSummary(ctx, req.(*GetDefaultPresetSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresetsService_GetSystemDefaultPresetSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemDefaultPresetSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetsServiceServer).GetSystemDefaultPresetSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetsService_GetSystemDefaultPresetSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetsServiceServer).GetSystemDefaultPresetSummary(ctx, req.(*GetSystemDefaultPresetSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PresetsService_ServiceDesc is the grpc.ServiceDesc for PresetsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PresetsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.notification_center.presets.v1.PresetsService",
	HandlerType: (*PresetsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomPreset",
			Handler:    _PresetsService_CreateCustomPreset_Handler,
		},
		{
			MethodName: "ReplaceCustomPreset",
			Handler:    _PresetsService_ReplaceCustomPreset_Handler,
		},
		{
			MethodName: "DeleteCustomPreset",
			Handler:    _PresetsService_DeleteCustomPreset_Handler,
		},
		{
			MethodName: "SetCustomPresetAsDefault",
			Handler:    _PresetsService_SetCustomPresetAsDefault_Handler,
		},
		{
			MethodName: "SetPresetAsDefault",
			Handler:    _PresetsService_SetPresetAsDefault_Handler,
		},
		{
			MethodName: "GetPreset",
			Handler:    _PresetsService_GetPreset_Handler,
		},
		{
			MethodName: "ListPresetSummaries",
			Handler:    _PresetsService_ListPresetSummaries_Handler,
		},
		{
			MethodName: "BatchGetPresets",
			Handler:    _PresetsService_BatchGetPresets_Handler,
		},
		{
			MethodName: "GetDefaultPresetSummary",
			Handler:    _PresetsService_GetDefaultPresetSummary_Handler,
		},
		{
			MethodName: "GetSystemDefaultPresetSummary",
			Handler:    _PresetsService_GetSystemDefaultPresetSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/notification_center/presets/v1/presets_service.proto",
}
