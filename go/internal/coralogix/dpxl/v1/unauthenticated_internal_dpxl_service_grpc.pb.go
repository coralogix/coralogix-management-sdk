// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.1
// source: com/coralogix/dpxl/v1/unauthenticated_internal_dpxl_service.proto

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
	UnauthenticatedInternalDpxlService_Compile_FullMethodName                      = "/com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService/Compile"
	UnauthenticatedInternalDpxlService_ToDpxlText_FullMethodName                   = "/com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService/ToDpxlText"
	UnauthenticatedInternalDpxlService_ToOpenSearchClause_FullMethodName           = "/com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService/ToOpenSearchClause"
	UnauthenticatedInternalDpxlService_ToSerializedDataprimeCoreAst_FullMethodName = "/com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService/ToSerializedDataprimeCoreAst"
	UnauthenticatedInternalDpxlService_Combine_FullMethodName                      = "/com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService/Combine"
	UnauthenticatedInternalDpxlService_CombineMany_FullMethodName                  = "/com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService/CombineMany"
	UnauthenticatedInternalDpxlService_CheckType_FullMethodName                    = "/com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService/CheckType"
)

// UnauthenticatedInternalDpxlServiceClient is the client API for UnauthenticatedInternalDpxlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UnauthenticatedInternalDpxlServiceClient interface {
	Compile(ctx context.Context, in *UnauthenticatedInternalDpxlServiceCompileRequest, opts ...grpc.CallOption) (*UnauthenticatedInternalDpxlServiceCompileResponse, error)
	ToDpxlText(ctx context.Context, in *UnauthenticatedInternalDpxlServiceToDpxlTextRequest, opts ...grpc.CallOption) (*UnauthenticatedInternalDpxlServiceToDpxlTextResponse, error)
	ToOpenSearchClause(ctx context.Context, in *UnauthenticatedInternalDpxlServiceToOpenSearchClauseRequest, opts ...grpc.CallOption) (*UnauthenticatedInternalDpxlServiceToOpenSearchClauseResponse, error)
	ToSerializedDataprimeCoreAst(ctx context.Context, in *UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstRequest, opts ...grpc.CallOption) (*UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstResponse, error)
	Combine(ctx context.Context, in *UnauthenticatedInternalDpxlServiceCombineRequest, opts ...grpc.CallOption) (*UnauthenticatedInternalDpxlServiceCombineResponse, error)
	CombineMany(ctx context.Context, in *UnauthenticatedInternalDpxlServiceCombineManyRequest, opts ...grpc.CallOption) (*UnauthenticatedInternalDpxlServiceCombineManyResponse, error)
	CheckType(ctx context.Context, in *UnauthenticatedInternalDpxlServiceCheckTypeRequest, opts ...grpc.CallOption) (*UnauthenticatedInternalDpxlServiceCheckTypeResponse, error)
}

type unauthenticatedInternalDpxlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUnauthenticatedInternalDpxlServiceClient(cc grpc.ClientConnInterface) UnauthenticatedInternalDpxlServiceClient {
	return &unauthenticatedInternalDpxlServiceClient{cc}
}

func (c *unauthenticatedInternalDpxlServiceClient) Compile(ctx context.Context, in *UnauthenticatedInternalDpxlServiceCompileRequest, opts ...grpc.CallOption) (*UnauthenticatedInternalDpxlServiceCompileResponse, error) {
	out := new(UnauthenticatedInternalDpxlServiceCompileResponse)
	err := c.cc.Invoke(ctx, UnauthenticatedInternalDpxlService_Compile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unauthenticatedInternalDpxlServiceClient) ToDpxlText(ctx context.Context, in *UnauthenticatedInternalDpxlServiceToDpxlTextRequest, opts ...grpc.CallOption) (*UnauthenticatedInternalDpxlServiceToDpxlTextResponse, error) {
	out := new(UnauthenticatedInternalDpxlServiceToDpxlTextResponse)
	err := c.cc.Invoke(ctx, UnauthenticatedInternalDpxlService_ToDpxlText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unauthenticatedInternalDpxlServiceClient) ToOpenSearchClause(ctx context.Context, in *UnauthenticatedInternalDpxlServiceToOpenSearchClauseRequest, opts ...grpc.CallOption) (*UnauthenticatedInternalDpxlServiceToOpenSearchClauseResponse, error) {
	out := new(UnauthenticatedInternalDpxlServiceToOpenSearchClauseResponse)
	err := c.cc.Invoke(ctx, UnauthenticatedInternalDpxlService_ToOpenSearchClause_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unauthenticatedInternalDpxlServiceClient) ToSerializedDataprimeCoreAst(ctx context.Context, in *UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstRequest, opts ...grpc.CallOption) (*UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstResponse, error) {
	out := new(UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstResponse)
	err := c.cc.Invoke(ctx, UnauthenticatedInternalDpxlService_ToSerializedDataprimeCoreAst_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unauthenticatedInternalDpxlServiceClient) Combine(ctx context.Context, in *UnauthenticatedInternalDpxlServiceCombineRequest, opts ...grpc.CallOption) (*UnauthenticatedInternalDpxlServiceCombineResponse, error) {
	out := new(UnauthenticatedInternalDpxlServiceCombineResponse)
	err := c.cc.Invoke(ctx, UnauthenticatedInternalDpxlService_Combine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unauthenticatedInternalDpxlServiceClient) CombineMany(ctx context.Context, in *UnauthenticatedInternalDpxlServiceCombineManyRequest, opts ...grpc.CallOption) (*UnauthenticatedInternalDpxlServiceCombineManyResponse, error) {
	out := new(UnauthenticatedInternalDpxlServiceCombineManyResponse)
	err := c.cc.Invoke(ctx, UnauthenticatedInternalDpxlService_CombineMany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unauthenticatedInternalDpxlServiceClient) CheckType(ctx context.Context, in *UnauthenticatedInternalDpxlServiceCheckTypeRequest, opts ...grpc.CallOption) (*UnauthenticatedInternalDpxlServiceCheckTypeResponse, error) {
	out := new(UnauthenticatedInternalDpxlServiceCheckTypeResponse)
	err := c.cc.Invoke(ctx, UnauthenticatedInternalDpxlService_CheckType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UnauthenticatedInternalDpxlServiceServer is the server API for UnauthenticatedInternalDpxlService service.
// All implementations must embed UnimplementedUnauthenticatedInternalDpxlServiceServer
// for forward compatibility
type UnauthenticatedInternalDpxlServiceServer interface {
	Compile(context.Context, *UnauthenticatedInternalDpxlServiceCompileRequest) (*UnauthenticatedInternalDpxlServiceCompileResponse, error)
	ToDpxlText(context.Context, *UnauthenticatedInternalDpxlServiceToDpxlTextRequest) (*UnauthenticatedInternalDpxlServiceToDpxlTextResponse, error)
	ToOpenSearchClause(context.Context, *UnauthenticatedInternalDpxlServiceToOpenSearchClauseRequest) (*UnauthenticatedInternalDpxlServiceToOpenSearchClauseResponse, error)
	ToSerializedDataprimeCoreAst(context.Context, *UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstRequest) (*UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstResponse, error)
	Combine(context.Context, *UnauthenticatedInternalDpxlServiceCombineRequest) (*UnauthenticatedInternalDpxlServiceCombineResponse, error)
	CombineMany(context.Context, *UnauthenticatedInternalDpxlServiceCombineManyRequest) (*UnauthenticatedInternalDpxlServiceCombineManyResponse, error)
	CheckType(context.Context, *UnauthenticatedInternalDpxlServiceCheckTypeRequest) (*UnauthenticatedInternalDpxlServiceCheckTypeResponse, error)
	mustEmbedUnimplementedUnauthenticatedInternalDpxlServiceServer()
}

// UnimplementedUnauthenticatedInternalDpxlServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUnauthenticatedInternalDpxlServiceServer struct {
}

func (UnimplementedUnauthenticatedInternalDpxlServiceServer) Compile(context.Context, *UnauthenticatedInternalDpxlServiceCompileRequest) (*UnauthenticatedInternalDpxlServiceCompileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compile not implemented")
}
func (UnimplementedUnauthenticatedInternalDpxlServiceServer) ToDpxlText(context.Context, *UnauthenticatedInternalDpxlServiceToDpxlTextRequest) (*UnauthenticatedInternalDpxlServiceToDpxlTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToDpxlText not implemented")
}
func (UnimplementedUnauthenticatedInternalDpxlServiceServer) ToOpenSearchClause(context.Context, *UnauthenticatedInternalDpxlServiceToOpenSearchClauseRequest) (*UnauthenticatedInternalDpxlServiceToOpenSearchClauseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToOpenSearchClause not implemented")
}
func (UnimplementedUnauthenticatedInternalDpxlServiceServer) ToSerializedDataprimeCoreAst(context.Context, *UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstRequest) (*UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToSerializedDataprimeCoreAst not implemented")
}
func (UnimplementedUnauthenticatedInternalDpxlServiceServer) Combine(context.Context, *UnauthenticatedInternalDpxlServiceCombineRequest) (*UnauthenticatedInternalDpxlServiceCombineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Combine not implemented")
}
func (UnimplementedUnauthenticatedInternalDpxlServiceServer) CombineMany(context.Context, *UnauthenticatedInternalDpxlServiceCombineManyRequest) (*UnauthenticatedInternalDpxlServiceCombineManyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CombineMany not implemented")
}
func (UnimplementedUnauthenticatedInternalDpxlServiceServer) CheckType(context.Context, *UnauthenticatedInternalDpxlServiceCheckTypeRequest) (*UnauthenticatedInternalDpxlServiceCheckTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckType not implemented")
}
func (UnimplementedUnauthenticatedInternalDpxlServiceServer) mustEmbedUnimplementedUnauthenticatedInternalDpxlServiceServer() {
}

// UnsafeUnauthenticatedInternalDpxlServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UnauthenticatedInternalDpxlServiceServer will
// result in compilation errors.
type UnsafeUnauthenticatedInternalDpxlServiceServer interface {
	mustEmbedUnimplementedUnauthenticatedInternalDpxlServiceServer()
}

func RegisterUnauthenticatedInternalDpxlServiceServer(s grpc.ServiceRegistrar, srv UnauthenticatedInternalDpxlServiceServer) {
	s.RegisterService(&UnauthenticatedInternalDpxlService_ServiceDesc, srv)
}

func _UnauthenticatedInternalDpxlService_Compile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnauthenticatedInternalDpxlServiceCompileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnauthenticatedInternalDpxlServiceServer).Compile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UnauthenticatedInternalDpxlService_Compile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnauthenticatedInternalDpxlServiceServer).Compile(ctx, req.(*UnauthenticatedInternalDpxlServiceCompileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnauthenticatedInternalDpxlService_ToDpxlText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnauthenticatedInternalDpxlServiceToDpxlTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnauthenticatedInternalDpxlServiceServer).ToDpxlText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UnauthenticatedInternalDpxlService_ToDpxlText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnauthenticatedInternalDpxlServiceServer).ToDpxlText(ctx, req.(*UnauthenticatedInternalDpxlServiceToDpxlTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnauthenticatedInternalDpxlService_ToOpenSearchClause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnauthenticatedInternalDpxlServiceToOpenSearchClauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnauthenticatedInternalDpxlServiceServer).ToOpenSearchClause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UnauthenticatedInternalDpxlService_ToOpenSearchClause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnauthenticatedInternalDpxlServiceServer).ToOpenSearchClause(ctx, req.(*UnauthenticatedInternalDpxlServiceToOpenSearchClauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnauthenticatedInternalDpxlService_ToSerializedDataprimeCoreAst_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnauthenticatedInternalDpxlServiceServer).ToSerializedDataprimeCoreAst(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UnauthenticatedInternalDpxlService_ToSerializedDataprimeCoreAst_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnauthenticatedInternalDpxlServiceServer).ToSerializedDataprimeCoreAst(ctx, req.(*UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnauthenticatedInternalDpxlService_Combine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnauthenticatedInternalDpxlServiceCombineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnauthenticatedInternalDpxlServiceServer).Combine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UnauthenticatedInternalDpxlService_Combine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnauthenticatedInternalDpxlServiceServer).Combine(ctx, req.(*UnauthenticatedInternalDpxlServiceCombineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnauthenticatedInternalDpxlService_CombineMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnauthenticatedInternalDpxlServiceCombineManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnauthenticatedInternalDpxlServiceServer).CombineMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UnauthenticatedInternalDpxlService_CombineMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnauthenticatedInternalDpxlServiceServer).CombineMany(ctx, req.(*UnauthenticatedInternalDpxlServiceCombineManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnauthenticatedInternalDpxlService_CheckType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnauthenticatedInternalDpxlServiceCheckTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnauthenticatedInternalDpxlServiceServer).CheckType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UnauthenticatedInternalDpxlService_CheckType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnauthenticatedInternalDpxlServiceServer).CheckType(ctx, req.(*UnauthenticatedInternalDpxlServiceCheckTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UnauthenticatedInternalDpxlService_ServiceDesc is the grpc.ServiceDesc for UnauthenticatedInternalDpxlService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UnauthenticatedInternalDpxlService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.dpxl.v1.UnauthenticatedInternalDpxlService",
	HandlerType: (*UnauthenticatedInternalDpxlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compile",
			Handler:    _UnauthenticatedInternalDpxlService_Compile_Handler,
		},
		{
			MethodName: "ToDpxlText",
			Handler:    _UnauthenticatedInternalDpxlService_ToDpxlText_Handler,
		},
		{
			MethodName: "ToOpenSearchClause",
			Handler:    _UnauthenticatedInternalDpxlService_ToOpenSearchClause_Handler,
		},
		{
			MethodName: "ToSerializedDataprimeCoreAst",
			Handler:    _UnauthenticatedInternalDpxlService_ToSerializedDataprimeCoreAst_Handler,
		},
		{
			MethodName: "Combine",
			Handler:    _UnauthenticatedInternalDpxlService_Combine_Handler,
		},
		{
			MethodName: "CombineMany",
			Handler:    _UnauthenticatedInternalDpxlService_CombineMany_Handler,
		},
		{
			MethodName: "CheckType",
			Handler:    _UnauthenticatedInternalDpxlService_CheckType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/dpxl/v1/unauthenticated_internal_dpxl_service.proto",
}
