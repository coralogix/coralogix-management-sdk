// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.28.2
// source: com/coralogix/archive/dataset/v2/internal_dataset_service.proto

package v2

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
	InternalDatasetManagementService_GetAllDatasetRules_FullMethodName = "/com.coralogix.archive.dataset.v2.InternalDatasetManagementService/GetAllDatasetRules"
)

// InternalDatasetManagementServiceClient is the client API for InternalDatasetManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalDatasetManagementServiceClient interface {
	GetAllDatasetRules(ctx context.Context, in *GetAllDatasetRulesRequest, opts ...grpc.CallOption) (*GetAllDatasetRulesResponse, error)
}

type internalDatasetManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalDatasetManagementServiceClient(cc grpc.ClientConnInterface) InternalDatasetManagementServiceClient {
	return &internalDatasetManagementServiceClient{cc}
}

func (c *internalDatasetManagementServiceClient) GetAllDatasetRules(ctx context.Context, in *GetAllDatasetRulesRequest, opts ...grpc.CallOption) (*GetAllDatasetRulesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllDatasetRulesResponse)
	err := c.cc.Invoke(ctx, InternalDatasetManagementService_GetAllDatasetRules_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalDatasetManagementServiceServer is the server API for InternalDatasetManagementService service.
// All implementations must embed UnimplementedInternalDatasetManagementServiceServer
// for forward compatibility
type InternalDatasetManagementServiceServer interface {
	GetAllDatasetRules(context.Context, *GetAllDatasetRulesRequest) (*GetAllDatasetRulesResponse, error)
	mustEmbedUnimplementedInternalDatasetManagementServiceServer()
}

// UnimplementedInternalDatasetManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInternalDatasetManagementServiceServer struct {
}

func (UnimplementedInternalDatasetManagementServiceServer) GetAllDatasetRules(context.Context, *GetAllDatasetRulesRequest) (*GetAllDatasetRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDatasetRules not implemented")
}
func (UnimplementedInternalDatasetManagementServiceServer) mustEmbedUnimplementedInternalDatasetManagementServiceServer() {
}

// UnsafeInternalDatasetManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternalDatasetManagementServiceServer will
// result in compilation errors.
type UnsafeInternalDatasetManagementServiceServer interface {
	mustEmbedUnimplementedInternalDatasetManagementServiceServer()
}

func RegisterInternalDatasetManagementServiceServer(s grpc.ServiceRegistrar, srv InternalDatasetManagementServiceServer) {
	s.RegisterService(&InternalDatasetManagementService_ServiceDesc, srv)
}

func _InternalDatasetManagementService_GetAllDatasetRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllDatasetRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalDatasetManagementServiceServer).GetAllDatasetRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalDatasetManagementService_GetAllDatasetRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalDatasetManagementServiceServer).GetAllDatasetRules(ctx, req.(*GetAllDatasetRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InternalDatasetManagementService_ServiceDesc is the grpc.ServiceDesc for InternalDatasetManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InternalDatasetManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.archive.dataset.v2.InternalDatasetManagementService",
	HandlerType: (*InternalDatasetManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllDatasetRules",
			Handler:    _InternalDatasetManagementService_GetAllDatasetRules_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/archive/dataset/v2/internal_dataset_service.proto",
}
