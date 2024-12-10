// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
// source: com/coralogixapis/metrics-rule-manager/v1/groups.proto

package golang

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RuleGroups_Save_FullMethodName   = "/rule_manager.groups.RuleGroups/Save"
	RuleGroups_Delete_FullMethodName = "/rule_manager.groups.RuleGroups/Delete"
	RuleGroups_List_FullMethodName   = "/rule_manager.groups.RuleGroups/List"
	RuleGroups_Fetch_FullMethodName  = "/rule_manager.groups.RuleGroups/Fetch"
)

// RuleGroupsClient is the client API for RuleGroups service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RuleGroupsClient interface {
	// Creates or updates a rule group.
	Save(ctx context.Context, in *InRuleGroup, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Deletes a rule group matching the given input.
	Delete(ctx context.Context, in *DeleteRuleGroup, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Lists all the rule groups.
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RuleGroupListing, error)
	// Fetches a rule group matching a given input.
	Fetch(ctx context.Context, in *FetchRuleGroup, opts ...grpc.CallOption) (*FetchRuleGroupResult, error)
}

type ruleGroupsClient struct {
	cc grpc.ClientConnInterface
}

func NewRuleGroupsClient(cc grpc.ClientConnInterface) RuleGroupsClient {
	return &ruleGroupsClient{cc}
}

func (c *ruleGroupsClient) Save(ctx context.Context, in *InRuleGroup, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RuleGroups_Save_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleGroupsClient) Delete(ctx context.Context, in *DeleteRuleGroup, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RuleGroups_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleGroupsClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RuleGroupListing, error) {
	out := new(RuleGroupListing)
	err := c.cc.Invoke(ctx, RuleGroups_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleGroupsClient) Fetch(ctx context.Context, in *FetchRuleGroup, opts ...grpc.CallOption) (*FetchRuleGroupResult, error) {
	out := new(FetchRuleGroupResult)
	err := c.cc.Invoke(ctx, RuleGroups_Fetch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuleGroupsServer is the server API for RuleGroups service.
// All implementations must embed UnimplementedRuleGroupsServer
// for forward compatibility
type RuleGroupsServer interface {
	// Creates or updates a rule group.
	Save(context.Context, *InRuleGroup) (*emptypb.Empty, error)
	// Deletes a rule group matching the given input.
	Delete(context.Context, *DeleteRuleGroup) (*emptypb.Empty, error)
	// Lists all the rule groups.
	List(context.Context, *emptypb.Empty) (*RuleGroupListing, error)
	// Fetches a rule group matching a given input.
	Fetch(context.Context, *FetchRuleGroup) (*FetchRuleGroupResult, error)
	mustEmbedUnimplementedRuleGroupsServer()
}

// UnimplementedRuleGroupsServer must be embedded to have forward compatible implementations.
type UnimplementedRuleGroupsServer struct {
}

func (UnimplementedRuleGroupsServer) Save(context.Context, *InRuleGroup) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedRuleGroupsServer) Delete(context.Context, *DeleteRuleGroup) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRuleGroupsServer) List(context.Context, *emptypb.Empty) (*RuleGroupListing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRuleGroupsServer) Fetch(context.Context, *FetchRuleGroup) (*FetchRuleGroupResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedRuleGroupsServer) mustEmbedUnimplementedRuleGroupsServer() {}

// UnsafeRuleGroupsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RuleGroupsServer will
// result in compilation errors.
type UnsafeRuleGroupsServer interface {
	mustEmbedUnimplementedRuleGroupsServer()
}

func RegisterRuleGroupsServer(s grpc.ServiceRegistrar, srv RuleGroupsServer) {
	s.RegisterService(&RuleGroups_ServiceDesc, srv)
}

func _RuleGroups_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InRuleGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleGroupsServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleGroups_Save_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleGroupsServer).Save(ctx, req.(*InRuleGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleGroups_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRuleGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleGroupsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleGroups_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleGroupsServer).Delete(ctx, req.(*DeleteRuleGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleGroups_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleGroupsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleGroups_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleGroupsServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleGroups_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRuleGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleGroupsServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleGroups_Fetch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleGroupsServer).Fetch(ctx, req.(*FetchRuleGroup))
	}
	return interceptor(ctx, in, info, handler)
}

// RuleGroups_ServiceDesc is the grpc.ServiceDesc for RuleGroups service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RuleGroups_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rule_manager.groups.RuleGroups",
	HandlerType: (*RuleGroupsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _RuleGroups_Save_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RuleGroups_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RuleGroups_List_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _RuleGroups_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/metrics-rule-manager/v1/groups.proto",
}

const (
	RuleGroupSets_Create_FullMethodName = "/rule_manager.groups.RuleGroupSets/Create"
	RuleGroupSets_Update_FullMethodName = "/rule_manager.groups.RuleGroupSets/Update"
	RuleGroupSets_List_FullMethodName   = "/rule_manager.groups.RuleGroupSets/List"
	RuleGroupSets_Fetch_FullMethodName  = "/rule_manager.groups.RuleGroupSets/Fetch"
	RuleGroupSets_Delete_FullMethodName = "/rule_manager.groups.RuleGroupSets/Delete"
)

// RuleGroupSetsClient is the client API for RuleGroupSets service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RuleGroupSetsClient interface {
	Create(ctx context.Context, in *CreateRuleGroupSet, opts ...grpc.CallOption) (*CreateRuleGroupSetResult, error)
	Update(ctx context.Context, in *UpdateRuleGroupSet, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RuleGroupSetListing, error)
	Fetch(ctx context.Context, in *FetchRuleGroupSet, opts ...grpc.CallOption) (*OutRuleGroupSet, error)
	Delete(ctx context.Context, in *DeleteRuleGroupSet, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type ruleGroupSetsClient struct {
	cc grpc.ClientConnInterface
}

func NewRuleGroupSetsClient(cc grpc.ClientConnInterface) RuleGroupSetsClient {
	return &ruleGroupSetsClient{cc}
}

func (c *ruleGroupSetsClient) Create(ctx context.Context, in *CreateRuleGroupSet, opts ...grpc.CallOption) (*CreateRuleGroupSetResult, error) {
	out := new(CreateRuleGroupSetResult)
	err := c.cc.Invoke(ctx, RuleGroupSets_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleGroupSetsClient) Update(ctx context.Context, in *UpdateRuleGroupSet, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RuleGroupSets_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleGroupSetsClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RuleGroupSetListing, error) {
	out := new(RuleGroupSetListing)
	err := c.cc.Invoke(ctx, RuleGroupSets_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleGroupSetsClient) Fetch(ctx context.Context, in *FetchRuleGroupSet, opts ...grpc.CallOption) (*OutRuleGroupSet, error) {
	out := new(OutRuleGroupSet)
	err := c.cc.Invoke(ctx, RuleGroupSets_Fetch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleGroupSetsClient) Delete(ctx context.Context, in *DeleteRuleGroupSet, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RuleGroupSets_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuleGroupSetsServer is the server API for RuleGroupSets service.
// All implementations must embed UnimplementedRuleGroupSetsServer
// for forward compatibility
type RuleGroupSetsServer interface {
	Create(context.Context, *CreateRuleGroupSet) (*CreateRuleGroupSetResult, error)
	Update(context.Context, *UpdateRuleGroupSet) (*emptypb.Empty, error)
	List(context.Context, *emptypb.Empty) (*RuleGroupSetListing, error)
	Fetch(context.Context, *FetchRuleGroupSet) (*OutRuleGroupSet, error)
	Delete(context.Context, *DeleteRuleGroupSet) (*emptypb.Empty, error)
	mustEmbedUnimplementedRuleGroupSetsServer()
}

// UnimplementedRuleGroupSetsServer must be embedded to have forward compatible implementations.
type UnimplementedRuleGroupSetsServer struct {
}

func (UnimplementedRuleGroupSetsServer) Create(context.Context, *CreateRuleGroupSet) (*CreateRuleGroupSetResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRuleGroupSetsServer) Update(context.Context, *UpdateRuleGroupSet) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRuleGroupSetsServer) List(context.Context, *emptypb.Empty) (*RuleGroupSetListing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRuleGroupSetsServer) Fetch(context.Context, *FetchRuleGroupSet) (*OutRuleGroupSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedRuleGroupSetsServer) Delete(context.Context, *DeleteRuleGroupSet) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRuleGroupSetsServer) mustEmbedUnimplementedRuleGroupSetsServer() {}

// UnsafeRuleGroupSetsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RuleGroupSetsServer will
// result in compilation errors.
type UnsafeRuleGroupSetsServer interface {
	mustEmbedUnimplementedRuleGroupSetsServer()
}

func RegisterRuleGroupSetsServer(s grpc.ServiceRegistrar, srv RuleGroupSetsServer) {
	s.RegisterService(&RuleGroupSets_ServiceDesc, srv)
}

func _RuleGroupSets_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRuleGroupSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleGroupSetsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleGroupSets_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleGroupSetsServer).Create(ctx, req.(*CreateRuleGroupSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleGroupSets_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRuleGroupSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleGroupSetsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleGroupSets_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleGroupSetsServer).Update(ctx, req.(*UpdateRuleGroupSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleGroupSets_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleGroupSetsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleGroupSets_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleGroupSetsServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleGroupSets_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRuleGroupSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleGroupSetsServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleGroupSets_Fetch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleGroupSetsServer).Fetch(ctx, req.(*FetchRuleGroupSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleGroupSets_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRuleGroupSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleGroupSetsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleGroupSets_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleGroupSetsServer).Delete(ctx, req.(*DeleteRuleGroupSet))
	}
	return interceptor(ctx, in, info, handler)
}

// RuleGroupSets_ServiceDesc is the grpc.ServiceDesc for RuleGroupSets service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RuleGroupSets_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rule_manager.groups.RuleGroupSets",
	HandlerType: (*RuleGroupSetsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RuleGroupSets_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RuleGroupSets_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RuleGroupSets_List_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _RuleGroupSets_Fetch_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RuleGroupSets_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/metrics-rule-manager/v1/groups.proto",
}
