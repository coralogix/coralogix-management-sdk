// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.28.2
// source: com/coralogixapis/aaa/organisations/v2/team_service.proto

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
	TeamService_CreateTeamInOrg_FullMethodName = "/com.coralogixapis.aaa.organisations.v2.TeamService/CreateTeamInOrg"
	TeamService_MoveQuota_FullMethodName       = "/com.coralogixapis.aaa.organisations.v2.TeamService/MoveQuota"
	TeamService_GetTeamQuota_FullMethodName    = "/com.coralogixapis.aaa.organisations.v2.TeamService/GetTeamQuota"
	TeamService_SetDailyQuota_FullMethodName   = "/com.coralogixapis.aaa.organisations.v2.TeamService/SetDailyQuota"
	TeamService_UpdateTeam_FullMethodName      = "/com.coralogixapis.aaa.organisations.v2.TeamService/UpdateTeam"
	TeamService_GetTeam_FullMethodName         = "/com.coralogixapis.aaa.organisations.v2.TeamService/GetTeam"
	TeamService_DeleteTeam_FullMethodName      = "/com.coralogixapis.aaa.organisations.v2.TeamService/DeleteTeam"
	TeamService_ListTeams_FullMethodName       = "/com.coralogixapis.aaa.organisations.v2.TeamService/ListTeams"
)

// TeamServiceClient is the client API for TeamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamServiceClient interface {
	CreateTeamInOrg(ctx context.Context, in *CreateTeamInOrgRequest, opts ...grpc.CallOption) (*CreateTeamInOrgResponse, error)
	MoveQuota(ctx context.Context, in *MoveQuotaRequest, opts ...grpc.CallOption) (*MoveQuotaResponse, error)
	GetTeamQuota(ctx context.Context, in *GetTeamQuotaRequest, opts ...grpc.CallOption) (*GetTeamQuotaResponse, error)
	SetDailyQuota(ctx context.Context, in *SetDailyQuotaRequest, opts ...grpc.CallOption) (*SetDailyQuotaResponse, error)
	UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*UpdateTeamResponse, error)
	GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*GetTeamResponse, error)
	DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamResponse, error)
	ListTeams(ctx context.Context, in *ListTeamsRequest, opts ...grpc.CallOption) (*ListTeamsResponse, error)
}

type teamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamServiceClient(cc grpc.ClientConnInterface) TeamServiceClient {
	return &teamServiceClient{cc}
}

func (c *teamServiceClient) CreateTeamInOrg(ctx context.Context, in *CreateTeamInOrgRequest, opts ...grpc.CallOption) (*CreateTeamInOrgResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTeamInOrgResponse)
	err := c.cc.Invoke(ctx, TeamService_CreateTeamInOrg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) MoveQuota(ctx context.Context, in *MoveQuotaRequest, opts ...grpc.CallOption) (*MoveQuotaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MoveQuotaResponse)
	err := c.cc.Invoke(ctx, TeamService_MoveQuota_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetTeamQuota(ctx context.Context, in *GetTeamQuotaRequest, opts ...grpc.CallOption) (*GetTeamQuotaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTeamQuotaResponse)
	err := c.cc.Invoke(ctx, TeamService_GetTeamQuota_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) SetDailyQuota(ctx context.Context, in *SetDailyQuotaRequest, opts ...grpc.CallOption) (*SetDailyQuotaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetDailyQuotaResponse)
	err := c.cc.Invoke(ctx, TeamService_SetDailyQuota_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*UpdateTeamResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTeamResponse)
	err := c.cc.Invoke(ctx, TeamService_UpdateTeam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*GetTeamResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTeamResponse)
	err := c.cc.Invoke(ctx, TeamService_GetTeam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTeamResponse)
	err := c.cc.Invoke(ctx, TeamService_DeleteTeam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) ListTeams(ctx context.Context, in *ListTeamsRequest, opts ...grpc.CallOption) (*ListTeamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTeamsResponse)
	err := c.cc.Invoke(ctx, TeamService_ListTeams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamServiceServer is the server API for TeamService service.
// All implementations must embed UnimplementedTeamServiceServer
// for forward compatibility
type TeamServiceServer interface {
	CreateTeamInOrg(context.Context, *CreateTeamInOrgRequest) (*CreateTeamInOrgResponse, error)
	MoveQuota(context.Context, *MoveQuotaRequest) (*MoveQuotaResponse, error)
	GetTeamQuota(context.Context, *GetTeamQuotaRequest) (*GetTeamQuotaResponse, error)
	SetDailyQuota(context.Context, *SetDailyQuotaRequest) (*SetDailyQuotaResponse, error)
	UpdateTeam(context.Context, *UpdateTeamRequest) (*UpdateTeamResponse, error)
	GetTeam(context.Context, *GetTeamRequest) (*GetTeamResponse, error)
	DeleteTeam(context.Context, *DeleteTeamRequest) (*DeleteTeamResponse, error)
	ListTeams(context.Context, *ListTeamsRequest) (*ListTeamsResponse, error)
	mustEmbedUnimplementedTeamServiceServer()
}

// UnimplementedTeamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTeamServiceServer struct {
}

func (UnimplementedTeamServiceServer) CreateTeamInOrg(context.Context, *CreateTeamInOrgRequest) (*CreateTeamInOrgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeamInOrg not implemented")
}
func (UnimplementedTeamServiceServer) MoveQuota(context.Context, *MoveQuotaRequest) (*MoveQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveQuota not implemented")
}
func (UnimplementedTeamServiceServer) GetTeamQuota(context.Context, *GetTeamQuotaRequest) (*GetTeamQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamQuota not implemented")
}
func (UnimplementedTeamServiceServer) SetDailyQuota(context.Context, *SetDailyQuotaRequest) (*SetDailyQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDailyQuota not implemented")
}
func (UnimplementedTeamServiceServer) UpdateTeam(context.Context, *UpdateTeamRequest) (*UpdateTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeam not implemented")
}
func (UnimplementedTeamServiceServer) GetTeam(context.Context, *GetTeamRequest) (*GetTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeam not implemented")
}
func (UnimplementedTeamServiceServer) DeleteTeam(context.Context, *DeleteTeamRequest) (*DeleteTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeam not implemented")
}
func (UnimplementedTeamServiceServer) ListTeams(context.Context, *ListTeamsRequest) (*ListTeamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTeams not implemented")
}
func (UnimplementedTeamServiceServer) mustEmbedUnimplementedTeamServiceServer() {}

// UnsafeTeamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamServiceServer will
// result in compilation errors.
type UnsafeTeamServiceServer interface {
	mustEmbedUnimplementedTeamServiceServer()
}

func RegisterTeamServiceServer(s grpc.ServiceRegistrar, srv TeamServiceServer) {
	s.RegisterService(&TeamService_ServiceDesc, srv)
}

func _TeamService_CreateTeamInOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTeamInOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).CreateTeamInOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_CreateTeamInOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).CreateTeamInOrg(ctx, req.(*CreateTeamInOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_MoveQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).MoveQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_MoveQuota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).MoveQuota(ctx, req.(*MoveQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetTeamQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetTeamQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetTeamQuota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetTeamQuota(ctx, req.(*GetTeamQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_SetDailyQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDailyQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).SetDailyQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_SetDailyQuota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).SetDailyQuota(ctx, req.(*SetDailyQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_UpdateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).UpdateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_UpdateTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).UpdateTeam(ctx, req.(*UpdateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetTeam(ctx, req.(*GetTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_DeleteTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).DeleteTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_DeleteTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).DeleteTeam(ctx, req.(*DeleteTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_ListTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).ListTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_ListTeams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).ListTeams(ctx, req.(*ListTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeamService_ServiceDesc is the grpc.ServiceDesc for TeamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.coralogixapis.aaa.organisations.v2.TeamService",
	HandlerType: (*TeamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTeamInOrg",
			Handler:    _TeamService_CreateTeamInOrg_Handler,
		},
		{
			MethodName: "MoveQuota",
			Handler:    _TeamService_MoveQuota_Handler,
		},
		{
			MethodName: "GetTeamQuota",
			Handler:    _TeamService_GetTeamQuota_Handler,
		},
		{
			MethodName: "SetDailyQuota",
			Handler:    _TeamService_SetDailyQuota_Handler,
		},
		{
			MethodName: "UpdateTeam",
			Handler:    _TeamService_UpdateTeam_Handler,
		},
		{
			MethodName: "GetTeam",
			Handler:    _TeamService_GetTeam_Handler,
		},
		{
			MethodName: "DeleteTeam",
			Handler:    _TeamService_DeleteTeam_Handler,
		},
		{
			MethodName: "ListTeams",
			Handler:    _TeamService_ListTeams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/coralogixapis/aaa/organisations/v2/team_service.proto",
}
