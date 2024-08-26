// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: com/coralogix/dataprime/v1/language_service.proto

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
	DataprimeLanguageService_Compile_FullMethodName = "/com.coralogix.dataprime.v1.DataprimeLanguageService/Compile"
	DataprimeLanguageService_Check_FullMethodName   = "/com.coralogix.dataprime.v1.DataprimeLanguageService/Check"
)

// DataprimeLanguageServiceClient is the client API for DataprimeLanguageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataprimeLanguageServiceClient interface {
	Compile(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileResponse, error)
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
}

type dataprimeLanguageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataprimeLanguageServiceClient(cc grpc.ClientConnInterface) DataprimeLanguageServiceClient {
	return &dataprimeLanguageServiceClient{cc}
}

func (c *dataprimeLanguageServiceClient) Compile(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileResponse, error) {
	out := new(CompileResponse)
	err := c.cc.Invoke(ctx, DataprimeLanguageService_Compile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataprimeLanguageServiceClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, DataprimeLanguageService_Check_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataprimeLanguageServiceServer is the server API for DataprimeLanguageService service.
// All implementations must embed UnimplementedDataprimeLanguageServiceServer
// for forward compatibility
type DataprimeLanguageServiceServer interface {
	Compile(context.Context, *CompileRequest) (*CompileResponse, error)
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	mustEmbedUnimplementedDataprimeLanguageServiceServer()
}

// UnimplementedDataprimeLanguageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataprimeLanguageServiceServer struct {
}

func (UnimplementedDataprimeLanguageServiceServer) Compile(context.Context, *CompileRequest) (*CompileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compile not implemented")
}
func (UnimplementedDataprimeLanguageServiceServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedDataprimeLanguageServiceServer) mustEmbedUnimplementedDataprimeLanguageServiceServer() {
}

// UnsafeDataprimeLanguageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataprimeLanguageServiceServer will
// result in compilation errors.
type UnsafeDataprimeLanguageServiceServer interface {
	mustEmbedUnimplementedDataprimeLanguageServiceServer()
}

func RegisterDataprimeLanguageServiceServer(s grpc.ServiceRegistrar, srv DataprimeLanguageServiceServer) {
	s.RegisterService(&DataprimeLanguageService_ServiceDesc, srv)
}

func _DataprimeLanguageService_Compile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataprimeLanguageServiceServer).Compile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataprimeLanguageService_Compile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataprimeLanguageServiceServer).Compile(ctx, req.(*CompileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataprimeLanguageService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataprimeLanguageServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataprimeLanguageService_Check_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataprimeLanguageServiceServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataprimeLanguageService_ServiceDesc is the grpc.ServiceDesc for DataprimeLanguageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataprimeLanguageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.dataprime.v1.DataprimeLanguageService",
	HandlerType: (*DataprimeLanguageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compile",
			Handler:    _DataprimeLanguageService_Compile_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _DataprimeLanguageService_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/dataprime/v1/language_service.proto",
}
