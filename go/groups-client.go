// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cxsdk

import (
	"context"

	groups "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/permissions/v1"
)

type TeamGroupId = groups.TeamGroupId

// GroupScope is a type for a group scope.
type GroupScope = groups.Scope

// TeamGroup is a type for a group.
type TeamGroup = groups.TeamGroup

// GroupsUser is a type for a user in a Groups context.
type GroupsUser = groups.User

type UpdateTeamGroupRequestRoleUpdates = groups.UpdateTeamGroupRequest_RoleUpdates
type UpdateTeamGroupRequestUserUpdates = groups.UpdateTeamGroupRequest_UserUpdates

// GetTeamGroupRequest is a type for a request.
type GetTeamGroupRequest = groups.GetTeamGroupRequest

// GetTeamGroupByNameRequest is a type for a request.
type GetTeamGroupByNameRequest = groups.GetTeamGroupByNameRequest

// GetTeamGroupsRequest is a type for a request.
type GetTeamGroupsRequest = groups.GetTeamGroupsRequest

// CreateTeamGroupRequest is a type for a request.
type CreateTeamGroupRequest = groups.CreateTeamGroupRequest

// UpdateTeamGroupRequest is a type for a request.
type UpdateTeamGroupRequest = groups.UpdateTeamGroupRequest

// DeleteTeamGroupRequest is a type for a request.
type DeleteTeamGroupRequest = groups.DeleteTeamGroupRequest

// GetGroupUsersRequest is a type for a request.
type GetGroupUsersRequest = groups.GetGroupUsersRequest

// AddUsersToTeamGroupRequest is a type for a request.
type AddUsersToTeamGroupRequest = groups.AddUsersToTeamGroupRequest

// AddUsersToTeamGroupsRequest is a type for a request.
type AddUsersToTeamGroupsRequest = groups.AddUsersToTeamGroupsRequest

// RemoveUsersFromTeamGroupRequest is a type for a request.
type RemoveUsersFromTeamGroupRequest = groups.RemoveUsersFromTeamGroupRequest

// RemoveUsersFromTeamGroupsRequest is a type for a request.
type RemoveUsersFromTeamGroupsRequest = groups.RemoveUsersFromTeamGroupsRequest

// SetTeamGroupScopeRequest is a type for a request.
type SetTeamGroupScopeRequest = groups.SetTeamGroupScopeRequest

// GetTeamGroupScopeRequest is a type for a request.
type GetTeamGroupScopeRequest = groups.GetTeamGroupScopeRequest

// RPC Values
const (
	CreateTeamGroupRpc          = groups.TeamPermissionsMgmtService_CreateTeamGroup_FullMethodName
	GetTeamGroupRpc             = groups.TeamPermissionsMgmtService_GetTeamGroup_FullMethodName
	GetTeamGroupsRpc            = groups.TeamPermissionsMgmtService_GetTeamGroups_FullMethodName
	UpdateTeamGroupRpc          = groups.TeamPermissionsMgmtService_UpdateTeamGroup_FullMethodName
	DeleteTeamGroupRpc          = groups.TeamPermissionsMgmtService_DeleteTeamGroup_FullMethodName
	AddUsersToTeamGroupRpc      = groups.TeamPermissionsMgmtService_AddUsersToTeamGroup_FullMethodName
	RemoveUsersFromTeamGroupRpc = groups.TeamPermissionsMgmtService_RemoveUsersFromTeamGroup_FullMethodName
)

// GroupsClient is a client for the Groups API
type GroupsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new group
func (c GroupsClient) Create(ctx context.Context, req *groups.CreateTeamGroupRequest) (*groups.CreateTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.CreateTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get retrieves a group by ID
func (c GroupsClient) Get(ctx context.Context, req *groups.GetTeamGroupRequest) (*groups.GetTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.GetTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// List retrieves all groups in the team
func (c GroupsClient) List(ctx context.Context, req *groups.GetTeamGroupsRequest) (*groups.GetTeamGroupsResponse, error) {
	callPoperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callPoperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.GetTeamGroups(callPoperties.Ctx, req, callPoperties.CallOptions...)
}

// Update updates a group
func (c GroupsClient) Update(ctx context.Context, req *groups.UpdateTeamGroupRequest) (*groups.UpdateTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.UpdateTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes a group by ID
func (c GroupsClient) Delete(ctx context.Context, req *groups.DeleteTeamGroupRequest) (*groups.DeleteTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.DeleteTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// AddUsers adds users to a group
func (c GroupsClient) AddUsers(ctx context.Context, req *groups.AddUsersToTeamGroupRequest) (*groups.AddUsersToTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.AddUsersToTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// RemoveUsers removes users from a group
func (c GroupsClient) RemoveUsers(ctx context.Context, req *groups.RemoveUsersFromTeamGroupRequest) (*groups.RemoveUsersFromTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.RemoveUsersFromTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewGroupsClient creates a new GroupsClient
func NewGroupsClient(callPropertiesCreator *CallPropertiesCreator) *GroupsClient {
	return &GroupsClient{callPropertiesCreator: callPropertiesCreator}
}
