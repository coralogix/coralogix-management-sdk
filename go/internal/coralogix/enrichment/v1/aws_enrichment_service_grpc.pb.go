// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
// source: com/coralogix/enrichment/v1/aws_enrichment_service.proto

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
	AwsEnrichmentService_GetSupportedAwsResourceTypes_FullMethodName = "/com.coralogix.enrichment.v1.AwsEnrichmentService/GetSupportedAwsResourceTypes"
)

// AwsEnrichmentServiceClient is the client API for AwsEnrichmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AwsEnrichmentServiceClient interface {
	GetSupportedAwsResourceTypes(ctx context.Context, in *GetSupportedAwsResourceTypesRequest, opts ...grpc.CallOption) (*GetSupportedAwsResourceTypesResponse, error)
}

type awsEnrichmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAwsEnrichmentServiceClient(cc grpc.ClientConnInterface) AwsEnrichmentServiceClient {
	return &awsEnrichmentServiceClient{cc}
}

func (c *awsEnrichmentServiceClient) GetSupportedAwsResourceTypes(ctx context.Context, in *GetSupportedAwsResourceTypesRequest, opts ...grpc.CallOption) (*GetSupportedAwsResourceTypesResponse, error) {
	out := new(GetSupportedAwsResourceTypesResponse)
	err := c.cc.Invoke(ctx, AwsEnrichmentService_GetSupportedAwsResourceTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AwsEnrichmentServiceServer is the server API for AwsEnrichmentService service.
// All implementations must embed UnimplementedAwsEnrichmentServiceServer
// for forward compatibility
type AwsEnrichmentServiceServer interface {
	GetSupportedAwsResourceTypes(context.Context, *GetSupportedAwsResourceTypesRequest) (*GetSupportedAwsResourceTypesResponse, error)
	mustEmbedUnimplementedAwsEnrichmentServiceServer()
}

// UnimplementedAwsEnrichmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAwsEnrichmentServiceServer struct {
}

func (UnimplementedAwsEnrichmentServiceServer) GetSupportedAwsResourceTypes(context.Context, *GetSupportedAwsResourceTypesRequest) (*GetSupportedAwsResourceTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportedAwsResourceTypes not implemented")
}
func (UnimplementedAwsEnrichmentServiceServer) mustEmbedUnimplementedAwsEnrichmentServiceServer() {}

// UnsafeAwsEnrichmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AwsEnrichmentServiceServer will
// result in compilation errors.
type UnsafeAwsEnrichmentServiceServer interface {
	mustEmbedUnimplementedAwsEnrichmentServiceServer()
}

func RegisterAwsEnrichmentServiceServer(s grpc.ServiceRegistrar, srv AwsEnrichmentServiceServer) {
	s.RegisterService(&AwsEnrichmentService_ServiceDesc, srv)
}

func _AwsEnrichmentService_GetSupportedAwsResourceTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSupportedAwsResourceTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwsEnrichmentServiceServer).GetSupportedAwsResourceTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AwsEnrichmentService_GetSupportedAwsResourceTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwsEnrichmentServiceServer).GetSupportedAwsResourceTypes(ctx, req.(*GetSupportedAwsResourceTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AwsEnrichmentService_ServiceDesc is the grpc.ServiceDesc for AwsEnrichmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AwsEnrichmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.enrichment.v1.AwsEnrichmentService",
	HandlerType: (*AwsEnrichmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSupportedAwsResourceTypes",
			Handler:    _AwsEnrichmentService_GetSupportedAwsResourceTypes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/enrichment/v1/aws_enrichment_service.proto",
}
