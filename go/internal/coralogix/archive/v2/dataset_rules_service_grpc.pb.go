// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: com/coralogix/archive/dataset/v2/dataset_rules_service.proto

package v2

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
	DatasetRulesService_SetDatasetRule_FullMethodName = "/com.coralogix.archive.dataset.v2.DatasetRulesService/SetDatasetRule"
)

// DatasetRulesServiceClient is the client API for DatasetRulesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatasetRulesServiceClient interface {
	SetDatasetRule(ctx context.Context, in *SetDatasetRuleRequest, opts ...grpc.CallOption) (*SetDatasetRuleResponse, error)
}

type datasetRulesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatasetRulesServiceClient(cc grpc.ClientConnInterface) DatasetRulesServiceClient {
	return &datasetRulesServiceClient{cc}
}

func (c *datasetRulesServiceClient) SetDatasetRule(ctx context.Context, in *SetDatasetRuleRequest, opts ...grpc.CallOption) (*SetDatasetRuleResponse, error) {
	out := new(SetDatasetRuleResponse)
	err := c.cc.Invoke(ctx, DatasetRulesService_SetDatasetRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasetRulesServiceServer is the server API for DatasetRulesService service.
// All implementations must embed UnimplementedDatasetRulesServiceServer
// for forward compatibility
type DatasetRulesServiceServer interface {
	SetDatasetRule(context.Context, *SetDatasetRuleRequest) (*SetDatasetRuleResponse, error)
	mustEmbedUnimplementedDatasetRulesServiceServer()
}

// UnimplementedDatasetRulesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDatasetRulesServiceServer struct {
}

func (UnimplementedDatasetRulesServiceServer) SetDatasetRule(context.Context, *SetDatasetRuleRequest) (*SetDatasetRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDatasetRule not implemented")
}
func (UnimplementedDatasetRulesServiceServer) mustEmbedUnimplementedDatasetRulesServiceServer() {}

// UnsafeDatasetRulesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatasetRulesServiceServer will
// result in compilation errors.
type UnsafeDatasetRulesServiceServer interface {
	mustEmbedUnimplementedDatasetRulesServiceServer()
}

func RegisterDatasetRulesServiceServer(s grpc.ServiceRegistrar, srv DatasetRulesServiceServer) {
	s.RegisterService(&DatasetRulesService_ServiceDesc, srv)
}

func _DatasetRulesService_SetDatasetRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDatasetRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetRulesServiceServer).SetDatasetRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetRulesService_SetDatasetRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetRulesServiceServer).SetDatasetRule(ctx, req.(*SetDatasetRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatasetRulesService_ServiceDesc is the grpc.ServiceDesc for DatasetRulesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatasetRulesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.archive.dataset.v2.DatasetRulesService",
	HandlerType: (*DatasetRulesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetDatasetRule",
			Handler:    _DatasetRulesService_SetDatasetRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/archive/dataset/v2/dataset_rules_service.proto",
}
