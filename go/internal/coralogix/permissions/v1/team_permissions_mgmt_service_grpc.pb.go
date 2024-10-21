// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.2
// source: com/coralogix/permissions/v1/team_permissions_mgmt_service.proto

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
	TeamPermissionsMgmtService_GetTeamGroup_FullMethodName              = "/com.coralogix.permissions.v1.TeamPermissionsMgmtService/GetTeamGroup"
	TeamPermissionsMgmtService_GetTeamGroupByName_FullMethodName        = "/com.coralogix.permissions.v1.TeamPermissionsMgmtService/GetTeamGroupByName"
	TeamPermissionsMgmtService_GetTeamGroups_FullMethodName             = "/com.coralogix.permissions.v1.TeamPermissionsMgmtService/GetTeamGroups"
	TeamPermissionsMgmtService_CreateTeamGroup_FullMethodName           = "/com.coralogix.permissions.v1.TeamPermissionsMgmtService/CreateTeamGroup"
	TeamPermissionsMgmtService_UpdateTeamGroup_FullMethodName           = "/com.coralogix.permissions.v1.TeamPermissionsMgmtService/UpdateTeamGroup"
	TeamPermissionsMgmtService_DeleteTeamGroup_FullMethodName           = "/com.coralogix.permissions.v1.TeamPermissionsMgmtService/DeleteTeamGroup"
	TeamPermissionsMgmtService_GetGroupUsers_FullMethodName             = "/com.coralogix.permissions.v1.TeamPermissionsMgmtService/GetGroupUsers"
	TeamPermissionsMgmtService_AddUsersToTeamGroup_FullMethodName       = "/com.coralogix.permissions.v1.TeamPermissionsMgmtService/AddUsersToTeamGroup"
	TeamPermissionsMgmtService_AddUsersToTeamGroups_FullMethodName      = "/com.coralogix.permissions.v1.TeamPermissionsMgmtService/AddUsersToTeamGroups"
	TeamPermissionsMgmtService_RemoveUsersFromTeamGroup_FullMethodName  = "/com.coralogix.permissions.v1.TeamPermissionsMgmtService/RemoveUsersFromTeamGroup"
	TeamPermissionsMgmtService_RemoveUsersFromTeamGroups_FullMethodName = "/com.coralogix.permissions.v1.TeamPermissionsMgmtService/RemoveUsersFromTeamGroups"
	TeamPermissionsMgmtService_SetTeamGroupScope_FullMethodName         = "/com.coralogix.permissions.v1.TeamPermissionsMgmtService/SetTeamGroupScope"
	TeamPermissionsMgmtService_GetTeamGroupScope_FullMethodName         = "/com.coralogix.permissions.v1.TeamPermissionsMgmtService/GetTeamGroupScope"
)

// TeamPermissionsMgmtServiceClient is the client API for TeamPermissionsMgmtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamPermissionsMgmtServiceClient interface {
	// / Fetch team group for given team id
	GetTeamGroup(ctx context.Context, in *GetTeamGroupRequest, opts ...grpc.CallOption) (*GetTeamGroupResponse, error)
	// / Fetch team group by name for authenticated team id
	GetTeamGroupByName(ctx context.Context, in *GetTeamGroupByNameRequest, opts ...grpc.CallOption) (*GetTeamGroupByNameResponse, error)
	// / Fetches all team groups linked with team
	GetTeamGroups(ctx context.Context, in *GetTeamGroupsRequest, opts ...grpc.CallOption) (*GetTeamGroupsResponse, error)
	// / Creates a new team group and optionally associates roles, users and a scope to it
	CreateTeamGroup(ctx context.Context, in *CreateTeamGroupRequest, opts ...grpc.CallOption) (*CreateTeamGroupResponse, error)
	// / Updates an existing team group details
	UpdateTeamGroup(ctx context.Context, in *UpdateTeamGroupRequest, opts ...grpc.CallOption) (*UpdateTeamGroupResponse, error)
	// / Deletes an existing team group
	DeleteTeamGroup(ctx context.Context, in *DeleteTeamGroupRequest, opts ...grpc.CallOption) (*DeleteTeamGroupResponse, error)
	// / Fetches all users assigned to group
	GetGroupUsers(ctx context.Context, in *GetGroupUsersRequest, opts ...grpc.CallOption) (*GetGroupUsersResponse, error)
	// / Add users to team groups
	AddUsersToTeamGroup(ctx context.Context, in *AddUsersToTeamGroupRequest, opts ...grpc.CallOption) (*AddUsersToTeamGroupResponse, error)
	// / Bulk Add users to team groups
	AddUsersToTeamGroups(ctx context.Context, in *AddUsersToTeamGroupsRequest, opts ...grpc.CallOption) (*AddUsersToTeamGroupsResponse, error)
	// / Remove user accounts to organisation group
	RemoveUsersFromTeamGroup(ctx context.Context, in *RemoveUsersFromTeamGroupRequest, opts ...grpc.CallOption) (*RemoveUsersFromTeamGroupResponse, error)
	// / Bulk Remove user accounts to organisation group
	RemoveUsersFromTeamGroups(ctx context.Context, in *RemoveUsersFromTeamGroupsRequest, opts ...grpc.CallOption) (*RemoveUsersFromTeamGroupsResponse, error)
	// / Sets team group scope. Replaces it if it already exists.
	SetTeamGroupScope(ctx context.Context, in *SetTeamGroupScopeRequest, opts ...grpc.CallOption) (*SetTeamGroupScopeResponse, error)
	// / Fetches team group scope
	GetTeamGroupScope(ctx context.Context, in *GetTeamGroupScopeRequest, opts ...grpc.CallOption) (*GetTeamGroupScopeResponse, error)
}

type teamPermissionsMgmtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamPermissionsMgmtServiceClient(cc grpc.ClientConnInterface) TeamPermissionsMgmtServiceClient {
	return &teamPermissionsMgmtServiceClient{cc}
}

func (c *teamPermissionsMgmtServiceClient) GetTeamGroup(ctx context.Context, in *GetTeamGroupRequest, opts ...grpc.CallOption) (*GetTeamGroupResponse, error) {
	out := new(GetTeamGroupResponse)
	err := c.cc.Invoke(ctx, TeamPermissionsMgmtService_GetTeamGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamPermissionsMgmtServiceClient) GetTeamGroupByName(ctx context.Context, in *GetTeamGroupByNameRequest, opts ...grpc.CallOption) (*GetTeamGroupByNameResponse, error) {
	out := new(GetTeamGroupByNameResponse)
	err := c.cc.Invoke(ctx, TeamPermissionsMgmtService_GetTeamGroupByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamPermissionsMgmtServiceClient) GetTeamGroups(ctx context.Context, in *GetTeamGroupsRequest, opts ...grpc.CallOption) (*GetTeamGroupsResponse, error) {
	out := new(GetTeamGroupsResponse)
	err := c.cc.Invoke(ctx, TeamPermissionsMgmtService_GetTeamGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamPermissionsMgmtServiceClient) CreateTeamGroup(ctx context.Context, in *CreateTeamGroupRequest, opts ...grpc.CallOption) (*CreateTeamGroupResponse, error) {
	out := new(CreateTeamGroupResponse)
	err := c.cc.Invoke(ctx, TeamPermissionsMgmtService_CreateTeamGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamPermissionsMgmtServiceClient) UpdateTeamGroup(ctx context.Context, in *UpdateTeamGroupRequest, opts ...grpc.CallOption) (*UpdateTeamGroupResponse, error) {
	out := new(UpdateTeamGroupResponse)
	err := c.cc.Invoke(ctx, TeamPermissionsMgmtService_UpdateTeamGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamPermissionsMgmtServiceClient) DeleteTeamGroup(ctx context.Context, in *DeleteTeamGroupRequest, opts ...grpc.CallOption) (*DeleteTeamGroupResponse, error) {
	out := new(DeleteTeamGroupResponse)
	err := c.cc.Invoke(ctx, TeamPermissionsMgmtService_DeleteTeamGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamPermissionsMgmtServiceClient) GetGroupUsers(ctx context.Context, in *GetGroupUsersRequest, opts ...grpc.CallOption) (*GetGroupUsersResponse, error) {
	out := new(GetGroupUsersResponse)
	err := c.cc.Invoke(ctx, TeamPermissionsMgmtService_GetGroupUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamPermissionsMgmtServiceClient) AddUsersToTeamGroup(ctx context.Context, in *AddUsersToTeamGroupRequest, opts ...grpc.CallOption) (*AddUsersToTeamGroupResponse, error) {
	out := new(AddUsersToTeamGroupResponse)
	err := c.cc.Invoke(ctx, TeamPermissionsMgmtService_AddUsersToTeamGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamPermissionsMgmtServiceClient) AddUsersToTeamGroups(ctx context.Context, in *AddUsersToTeamGroupsRequest, opts ...grpc.CallOption) (*AddUsersToTeamGroupsResponse, error) {
	out := new(AddUsersToTeamGroupsResponse)
	err := c.cc.Invoke(ctx, TeamPermissionsMgmtService_AddUsersToTeamGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamPermissionsMgmtServiceClient) RemoveUsersFromTeamGroup(ctx context.Context, in *RemoveUsersFromTeamGroupRequest, opts ...grpc.CallOption) (*RemoveUsersFromTeamGroupResponse, error) {
	out := new(RemoveUsersFromTeamGroupResponse)
	err := c.cc.Invoke(ctx, TeamPermissionsMgmtService_RemoveUsersFromTeamGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamPermissionsMgmtServiceClient) RemoveUsersFromTeamGroups(ctx context.Context, in *RemoveUsersFromTeamGroupsRequest, opts ...grpc.CallOption) (*RemoveUsersFromTeamGroupsResponse, error) {
	out := new(RemoveUsersFromTeamGroupsResponse)
	err := c.cc.Invoke(ctx, TeamPermissionsMgmtService_RemoveUsersFromTeamGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamPermissionsMgmtServiceClient) SetTeamGroupScope(ctx context.Context, in *SetTeamGroupScopeRequest, opts ...grpc.CallOption) (*SetTeamGroupScopeResponse, error) {
	out := new(SetTeamGroupScopeResponse)
	err := c.cc.Invoke(ctx, TeamPermissionsMgmtService_SetTeamGroupScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamPermissionsMgmtServiceClient) GetTeamGroupScope(ctx context.Context, in *GetTeamGroupScopeRequest, opts ...grpc.CallOption) (*GetTeamGroupScopeResponse, error) {
	out := new(GetTeamGroupScopeResponse)
	err := c.cc.Invoke(ctx, TeamPermissionsMgmtService_GetTeamGroupScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamPermissionsMgmtServiceServer is the server API for TeamPermissionsMgmtService service.
// All implementations must embed UnimplementedTeamPermissionsMgmtServiceServer
// for forward compatibility
type TeamPermissionsMgmtServiceServer interface {
	// / Fetch team group for given team id
	GetTeamGroup(context.Context, *GetTeamGroupRequest) (*GetTeamGroupResponse, error)
	// / Fetch team group by name for authenticated team id
	GetTeamGroupByName(context.Context, *GetTeamGroupByNameRequest) (*GetTeamGroupByNameResponse, error)
	// / Fetches all team groups linked with team
	GetTeamGroups(context.Context, *GetTeamGroupsRequest) (*GetTeamGroupsResponse, error)
	// / Creates a new team group and optionally associates roles, users and a scope to it
	CreateTeamGroup(context.Context, *CreateTeamGroupRequest) (*CreateTeamGroupResponse, error)
	// / Updates an existing team group details
	UpdateTeamGroup(context.Context, *UpdateTeamGroupRequest) (*UpdateTeamGroupResponse, error)
	// / Deletes an existing team group
	DeleteTeamGroup(context.Context, *DeleteTeamGroupRequest) (*DeleteTeamGroupResponse, error)
	// / Fetches all users assigned to group
	GetGroupUsers(context.Context, *GetGroupUsersRequest) (*GetGroupUsersResponse, error)
	// / Add users to team groups
	AddUsersToTeamGroup(context.Context, *AddUsersToTeamGroupRequest) (*AddUsersToTeamGroupResponse, error)
	// / Bulk Add users to team groups
	AddUsersToTeamGroups(context.Context, *AddUsersToTeamGroupsRequest) (*AddUsersToTeamGroupsResponse, error)
	// / Remove user accounts to organisation group
	RemoveUsersFromTeamGroup(context.Context, *RemoveUsersFromTeamGroupRequest) (*RemoveUsersFromTeamGroupResponse, error)
	// / Bulk Remove user accounts to organisation group
	RemoveUsersFromTeamGroups(context.Context, *RemoveUsersFromTeamGroupsRequest) (*RemoveUsersFromTeamGroupsResponse, error)
	// / Sets team group scope. Replaces it if it already exists.
	SetTeamGroupScope(context.Context, *SetTeamGroupScopeRequest) (*SetTeamGroupScopeResponse, error)
	// / Fetches team group scope
	GetTeamGroupScope(context.Context, *GetTeamGroupScopeRequest) (*GetTeamGroupScopeResponse, error)
	mustEmbedUnimplementedTeamPermissionsMgmtServiceServer()
}

// UnimplementedTeamPermissionsMgmtServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTeamPermissionsMgmtServiceServer struct {
}

func (UnimplementedTeamPermissionsMgmtServiceServer) GetTeamGroup(context.Context, *GetTeamGroupRequest) (*GetTeamGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamGroup not implemented")
}
func (UnimplementedTeamPermissionsMgmtServiceServer) GetTeamGroupByName(context.Context, *GetTeamGroupByNameRequest) (*GetTeamGroupByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamGroupByName not implemented")
}
func (UnimplementedTeamPermissionsMgmtServiceServer) GetTeamGroups(context.Context, *GetTeamGroupsRequest) (*GetTeamGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamGroups not implemented")
}
func (UnimplementedTeamPermissionsMgmtServiceServer) CreateTeamGroup(context.Context, *CreateTeamGroupRequest) (*CreateTeamGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeamGroup not implemented")
}
func (UnimplementedTeamPermissionsMgmtServiceServer) UpdateTeamGroup(context.Context, *UpdateTeamGroupRequest) (*UpdateTeamGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeamGroup not implemented")
}
func (UnimplementedTeamPermissionsMgmtServiceServer) DeleteTeamGroup(context.Context, *DeleteTeamGroupRequest) (*DeleteTeamGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeamGroup not implemented")
}
func (UnimplementedTeamPermissionsMgmtServiceServer) GetGroupUsers(context.Context, *GetGroupUsersRequest) (*GetGroupUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupUsers not implemented")
}
func (UnimplementedTeamPermissionsMgmtServiceServer) AddUsersToTeamGroup(context.Context, *AddUsersToTeamGroupRequest) (*AddUsersToTeamGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUsersToTeamGroup not implemented")
}
func (UnimplementedTeamPermissionsMgmtServiceServer) AddUsersToTeamGroups(context.Context, *AddUsersToTeamGroupsRequest) (*AddUsersToTeamGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUsersToTeamGroups not implemented")
}
func (UnimplementedTeamPermissionsMgmtServiceServer) RemoveUsersFromTeamGroup(context.Context, *RemoveUsersFromTeamGroupRequest) (*RemoveUsersFromTeamGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUsersFromTeamGroup not implemented")
}
func (UnimplementedTeamPermissionsMgmtServiceServer) RemoveUsersFromTeamGroups(context.Context, *RemoveUsersFromTeamGroupsRequest) (*RemoveUsersFromTeamGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUsersFromTeamGroups not implemented")
}
func (UnimplementedTeamPermissionsMgmtServiceServer) SetTeamGroupScope(context.Context, *SetTeamGroupScopeRequest) (*SetTeamGroupScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTeamGroupScope not implemented")
}
func (UnimplementedTeamPermissionsMgmtServiceServer) GetTeamGroupScope(context.Context, *GetTeamGroupScopeRequest) (*GetTeamGroupScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamGroupScope not implemented")
}
func (UnimplementedTeamPermissionsMgmtServiceServer) mustEmbedUnimplementedTeamPermissionsMgmtServiceServer() {
}

// UnsafeTeamPermissionsMgmtServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamPermissionsMgmtServiceServer will
// result in compilation errors.
type UnsafeTeamPermissionsMgmtServiceServer interface {
	mustEmbedUnimplementedTeamPermissionsMgmtServiceServer()
}

func RegisterTeamPermissionsMgmtServiceServer(s grpc.ServiceRegistrar, srv TeamPermissionsMgmtServiceServer) {
	s.RegisterService(&TeamPermissionsMgmtService_ServiceDesc, srv)
}

func _TeamPermissionsMgmtService_GetTeamGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamPermissionsMgmtServiceServer).GetTeamGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamPermissionsMgmtService_GetTeamGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamPermissionsMgmtServiceServer).GetTeamGroup(ctx, req.(*GetTeamGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamPermissionsMgmtService_GetTeamGroupByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamGroupByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamPermissionsMgmtServiceServer).GetTeamGroupByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamPermissionsMgmtService_GetTeamGroupByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamPermissionsMgmtServiceServer).GetTeamGroupByName(ctx, req.(*GetTeamGroupByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamPermissionsMgmtService_GetTeamGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamPermissionsMgmtServiceServer).GetTeamGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamPermissionsMgmtService_GetTeamGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamPermissionsMgmtServiceServer).GetTeamGroups(ctx, req.(*GetTeamGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamPermissionsMgmtService_CreateTeamGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTeamGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamPermissionsMgmtServiceServer).CreateTeamGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamPermissionsMgmtService_CreateTeamGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamPermissionsMgmtServiceServer).CreateTeamGroup(ctx, req.(*CreateTeamGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamPermissionsMgmtService_UpdateTeamGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamPermissionsMgmtServiceServer).UpdateTeamGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamPermissionsMgmtService_UpdateTeamGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamPermissionsMgmtServiceServer).UpdateTeamGroup(ctx, req.(*UpdateTeamGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamPermissionsMgmtService_DeleteTeamGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeamGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamPermissionsMgmtServiceServer).DeleteTeamGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamPermissionsMgmtService_DeleteTeamGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamPermissionsMgmtServiceServer).DeleteTeamGroup(ctx, req.(*DeleteTeamGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamPermissionsMgmtService_GetGroupUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamPermissionsMgmtServiceServer).GetGroupUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamPermissionsMgmtService_GetGroupUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamPermissionsMgmtServiceServer).GetGroupUsers(ctx, req.(*GetGroupUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamPermissionsMgmtService_AddUsersToTeamGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUsersToTeamGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamPermissionsMgmtServiceServer).AddUsersToTeamGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamPermissionsMgmtService_AddUsersToTeamGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamPermissionsMgmtServiceServer).AddUsersToTeamGroup(ctx, req.(*AddUsersToTeamGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamPermissionsMgmtService_AddUsersToTeamGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUsersToTeamGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamPermissionsMgmtServiceServer).AddUsersToTeamGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamPermissionsMgmtService_AddUsersToTeamGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamPermissionsMgmtServiceServer).AddUsersToTeamGroups(ctx, req.(*AddUsersToTeamGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamPermissionsMgmtService_RemoveUsersFromTeamGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUsersFromTeamGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamPermissionsMgmtServiceServer).RemoveUsersFromTeamGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamPermissionsMgmtService_RemoveUsersFromTeamGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamPermissionsMgmtServiceServer).RemoveUsersFromTeamGroup(ctx, req.(*RemoveUsersFromTeamGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamPermissionsMgmtService_RemoveUsersFromTeamGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUsersFromTeamGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamPermissionsMgmtServiceServer).RemoveUsersFromTeamGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamPermissionsMgmtService_RemoveUsersFromTeamGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamPermissionsMgmtServiceServer).RemoveUsersFromTeamGroups(ctx, req.(*RemoveUsersFromTeamGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamPermissionsMgmtService_SetTeamGroupScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTeamGroupScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamPermissionsMgmtServiceServer).SetTeamGroupScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamPermissionsMgmtService_SetTeamGroupScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamPermissionsMgmtServiceServer).SetTeamGroupScope(ctx, req.(*SetTeamGroupScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamPermissionsMgmtService_GetTeamGroupScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamGroupScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamPermissionsMgmtServiceServer).GetTeamGroupScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamPermissionsMgmtService_GetTeamGroupScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamPermissionsMgmtServiceServer).GetTeamGroupScope(ctx, req.(*GetTeamGroupScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeamPermissionsMgmtService_ServiceDesc is the grpc.ServiceDesc for TeamPermissionsMgmtService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeamPermissionsMgmtService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogix.permissions.v1.TeamPermissionsMgmtService",
	HandlerType: (*TeamPermissionsMgmtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeamGroup",
			Handler:    _TeamPermissionsMgmtService_GetTeamGroup_Handler,
		},
		{
			MethodName: "GetTeamGroupByName",
			Handler:    _TeamPermissionsMgmtService_GetTeamGroupByName_Handler,
		},
		{
			MethodName: "GetTeamGroups",
			Handler:    _TeamPermissionsMgmtService_GetTeamGroups_Handler,
		},
		{
			MethodName: "CreateTeamGroup",
			Handler:    _TeamPermissionsMgmtService_CreateTeamGroup_Handler,
		},
		{
			MethodName: "UpdateTeamGroup",
			Handler:    _TeamPermissionsMgmtService_UpdateTeamGroup_Handler,
		},
		{
			MethodName: "DeleteTeamGroup",
			Handler:    _TeamPermissionsMgmtService_DeleteTeamGroup_Handler,
		},
		{
			MethodName: "GetGroupUsers",
			Handler:    _TeamPermissionsMgmtService_GetGroupUsers_Handler,
		},
		{
			MethodName: "AddUsersToTeamGroup",
			Handler:    _TeamPermissionsMgmtService_AddUsersToTeamGroup_Handler,
		},
		{
			MethodName: "AddUsersToTeamGroups",
			Handler:    _TeamPermissionsMgmtService_AddUsersToTeamGroups_Handler,
		},
		{
			MethodName: "RemoveUsersFromTeamGroup",
			Handler:    _TeamPermissionsMgmtService_RemoveUsersFromTeamGroup_Handler,
		},
		{
			MethodName: "RemoveUsersFromTeamGroups",
			Handler:    _TeamPermissionsMgmtService_RemoveUsersFromTeamGroups_Handler,
		},
		{
			MethodName: "SetTeamGroupScope",
			Handler:    _TeamPermissionsMgmtService_SetTeamGroupScope_Handler,
		},
		{
			MethodName: "GetTeamGroupScope",
			Handler:    _TeamPermissionsMgmtService_GetTeamGroupScope_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogix/permissions/v1/team_permissions_mgmt_service.proto",
}
