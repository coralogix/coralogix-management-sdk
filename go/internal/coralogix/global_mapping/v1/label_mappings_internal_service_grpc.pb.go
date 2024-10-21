// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.2
// source: com/coralogix/global_mapping/v1/label_mappings_internal_service.proto

package v1

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
	LabelMappingsInternalService_UpsertLabelMappingsInternal_FullMethodName = "/com.coralogix.global_mapping.v1.LabelMappingsInternalService/UpsertLabelMappingsInternal"
	LabelMappingsInternalService_GetLabelMappingsInternal_FullMethodName    = "/com.coralogix.global_mapping.v1.LabelMappingsInternalService/GetLabelMappingsInternal"
	LabelMappingsInternalService_GetLabelValuesInternal_FullMethodName      = "/com.coralogix.global_mapping.v1.LabelMappingsInternalService/GetLabelValuesInternal"
)

// LabelMappingsInternalServiceClient is the client API for LabelMappingsInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LabelMappingsInternalServiceClient interface {
	UpsertLabelMappingsInternal(ctx context.Context, in *UpsertLabelMappingsInternalRequest, opts ...grpc.CallOption) (*UpsertLabelMappingsInternalResponse, error)
	GetLabelMappingsInternal(ctx context.Context, in *GetLabelMappingsInternalRequest, opts ...grpc.CallOption) (*GetLabelMappingsInternalResponse, error)
	GetLabelValuesInternal(ctx context.Context, in *GetLabelValuesInternalRequest, opts ...grpc.CallOption) (*GetLabelValuesInternalResponse, error)
}

type labelMappingsInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLabelMappingsInternalServiceClient(cc grpc.ClientConnInterface) LabelMappingsInternalServiceClient {
	return &labelMappingsInternalServiceClient{cc}
}

func (c *labelMappingsInternalServiceClient) UpsertLabelMappingsInternal(ctx context.Context, in *UpsertLabelMappingsInternalRequest, opts ...grpc.CallOption) (*UpsertLabelMappingsInternalResponse, error) {
	out := new(UpsertLabelMappingsInternalResponse)
	err := c.cc.Invoke(ctx, LabelMappingsInternalService_UpsertLabelMappingsInternal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelMappingsInternalServiceClient) GetLabelMappingsInternal(ctx context.Context, in *GetLabelMappingsInternalRequest, opts ...grpc.CallOption) (*GetLabelMappingsInternalResponse, error) {
	out := new(GetLabelMappingsInternalResponse)
	err := c.cc.Invoke(ctx, LabelMappingsInternalService_GetLabelMappingsInternal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelMappingsInternalServiceClient) GetLabelValuesInternal(ctx context.Context, in *GetLabelValuesInternalRequest, opts ...grpc.CallOption) (*GetLabelValuesInternalResponse, error) {
	out := new(GetLabelValuesInternalResponse)
	err := c.cc.Invoke(ctx, LabelMappingsInternalService_GetLabelValuesInternal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LabelMappingsInternalServiceServer is the server API for LabelMappingsInternalService service.
// All implementations must embed UnimplementedLabelMappingsInternalServiceServer
// for forward compatibility
type LabelMappingsInternalServiceServer interface {
	UpsertLabelMappingsInternal(context.Context, *UpsertLabelMappingsInternalRequest) (*UpsertLabelMappingsInternalResponse, error)
	GetLabelMappingsInternal(context.Context, *GetLabelMappingsInternalRequest) (*GetLabelMappingsInternalResponse, error)
	GetLabelValuesInternal(context.Context, *GetLabelValuesInternalRequest) (*GetLabelValuesInternalResponse, error)
	mustEmbedUnimplementedLabelMappingsInternalServiceServer()
}

// UnimplementedLabelMappingsInternalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLabelMappingsInternalServiceServer struct {
}

func (UnimplementedLabelMappingsInternalServiceServer) UpsertLabelMappingsInternal(context.Context, *UpsertLabelMappingsInternalRequest) (*UpsertLabelMappingsInternalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertLabelMappingsInternal not implemented")
}
func (UnimplementedLabelMappingsInternalServiceServer) GetLabelMappingsInternal(context.Context, *GetLabelMappingsInternalRequest) (*GetLabelMappingsInternalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabelMappingsInternal not implemented")
}
func (UnimplementedLabelMappingsInternalServiceServer) GetLabelValuesInternal(context.Context, *GetLabelValuesInternalRequest) (*GetLabelValuesInternalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabelValuesInternal not implemented")
}
func (UnimplementedLabelMappingsInternalServiceServer) mustEmbedUnimplementedLabelMappingsInternalServiceServer() {
}

// UnsafeLabelMappingsInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LabelMappingsInternalServiceServer will
// result in compilation errors.
type UnsafeLabelMappingsInternalServiceServer interface {
	mustEmbedUnimplementedLabelMappingsInternalServiceServer()
}

func RegisterLabelMappingsInternalServiceServer(s grpc.ServiceRegistrar, srv LabelMappingsInternalServiceServer) {
	s.RegisterService(&LabelMappingsInternalService_ServiceDesc, srv)
}

func _LabelMappingsInternalService_UpsertLabelMappingsInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertLabelMappingsInternalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelMappingsInternalServiceServer).UpsertLabelMappingsInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelMappingsInternalService_UpsertLabelMappingsInternal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelMappingsInternalServiceServer).UpsertLabelMappingsInternal(ctx, req.(*UpsertLabelMappingsInternalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelMappingsInternalService_GetLabelMappingsInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelMappingsInternalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelMappingsInternalServiceServer).GetLabelMappingsInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelMappingsInternalService_GetLabelMappingsInternal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelMappingsInternalServiceServer).GetLabelMappingsInternal(ctx, req.(*GetLabelMappingsInternalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelMappingsInternalService_GetLabelValuesInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelValuesInternalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelMappingsInternalServiceServer).GetLabelValuesInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelMappingsInternalService_GetLabelValuesInternal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelMappingsInternalServiceServer).GetLabelValuesInternal(ctx, req.(*GetLabelValuesInternalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LabelMappingsInternalService_ServiceDesc is the grpc.ServiceDesc for LabelMappingsInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LabelMappingsInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.global_mapping.v1.LabelMappingsInternalService",
	HandlerType: (*LabelMappingsInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertLabelMappingsInternal",
			Handler:    _LabelMappingsInternalService_UpsertLabelMappingsInternal_Handler,
		},
		{
			MethodName: "GetLabelMappingsInternal",
			Handler:    _LabelMappingsInternalService_GetLabelMappingsInternal_Handler,
		},
		{
			MethodName: "GetLabelValuesInternal",
			Handler:    _LabelMappingsInternalService_GetLabelValuesInternal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/global_mapping/v1/label_mappings_internal_service.proto",
}
