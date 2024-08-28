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

// CreateTeamGroupRequest is a request to create a new group
type CreateTeamGroupRequest = groups.CreateTeamGroupRequest

// CreateTeamGroupResponse is a response to creating a new group
type CreateTeamGroupResponse = groups.CreateTeamGroupResponse

// DeleteTeamGroupRequest is a request to delete a group by its ID
type DeleteTeamGroupRequest = groups.DeleteTeamGroupRequest

// UpdateTeamGroupRequest is a request to update a group by its ID
type UpdateTeamGroupRequest = groups.UpdateTeamGroupRequest

// GetTeamGroupRequest is a request to get a group by its ID
type GetTeamGroupRequest = groups.GetTeamGroupRequest

// GetTeamGroupsRequest is a request to list all groups in the team
type GetTeamGroupsRequest = groups.GetTeamGroupsRequest

// GetTeamGroupResponse is a response to getting a group by its ID
type GetTeamGroupResponse = groups.GetTeamGroupResponse

// GroupsTeam is a team
type GroupsTeamGroup = groups.TeamGroup

// GroupsTeamID is an ID for a group in a Team
type GroupsTeamID = groups.TeamId

// GroupsTeamGroupID is an ID for a group in a Team
type GroupsTeamGroupID = groups.TeamGroupId

// GroupsRoleID is an ID for a group in a Team
type GroupsRoleID = groups.RoleId

// GroupsUserID is an ID for a user
type GroupsUserID = groups.UserId

// AddUsersToGroup for bulk adding users to groups
type AddUsersToGroup = groups.AddUsersToTeamGroupsRequest_AddUsersToTeamGroup

// RemoveUsersFromGroup for bulk removing users from groups
type RemoveUsersFromGroup = groups.RemoveUsersFromTeamGroupsRequest_RemoveUsersFromTeamGroup

// GetTeamGroupRpc is the name of the RPC to get a group by its ID
const GetTeamGroupRpc = groups.TeamPermissionsMgmtService_GetTeamGroup_FullMethodName

// GetTeamGroupsRpc is the name of the RPC to list all groups in the team
const GetTeamGroupByNameRpc = groups.TeamPermissionsMgmtService_GetTeamGroupByName_FullMethodName

// GetTeamGroupsRpc is the name of the RPC to list all groups in the team
const GetTeamGroupsRpc = groups.TeamPermissionsMgmtService_GetTeamGroups_FullMethodName

// CreateTeamGroupRpc is the name of the RPC to create a new group
const CreateTeamGroupRpc = groups.TeamPermissionsMgmtService_CreateTeamGroup_FullMethodName

// UpdateTeamGroupRpc is the name of the RPC to update a group
const UpdateTeamGroupRpc = groups.TeamPermissionsMgmtService_UpdateTeamGroup_FullMethodName

// DeleteTeamGroupRpc is the name of the RPC to delete a group by its ID
const DeleteTeamGroupRpc = groups.TeamPermissionsMgmtService_DeleteTeamGroup_FullMethodName

// GetGroupUsersRpc is the name of the RPC to get all users in a group
const GetGroupUsersRpc = groups.TeamPermissionsMgmtService_GetGroupUsers_FullMethodName

// AddUsersToTeamGroupRpc is the name of the RPC to add users to a group
const AddUsersToTeamGroupRpc = groups.TeamPermissionsMgmtService_AddUsersToTeamGroup_FullMethodName

// AddUsersToTeamGroupsRpc is the name of the RPC to remove users from a group
const AddUsersToTeamGroupsRpc = groups.TeamPermissionsMgmtService_AddUsersToTeamGroups_FullMethodName

// RemoveUsersFromTeamGroupRpc is the name of the RPC to remove users from a group
const RemoveUsersFromTeamGroupRpc = groups.TeamPermissionsMgmtService_RemoveUsersFromTeamGroup_FullMethodName

// RemoveUsersFromTeamGroupsRpc is the name of the RPC to remove users from multiple groups
const RemoveUsersFromTeamGroupsRpc = groups.TeamPermissionsMgmtService_RemoveUsersFromTeamGroups_FullMethodName

// SetTeamGroupScopeRpc is the name of the RPC to set the scope of a group
const SetTeamGroupScopeRpc = groups.TeamPermissionsMgmtService_SetTeamGroupScope_FullMethodName

// GetTeamGroupScopeRpc is the name of the RPC to get the scope of a group
const GetTeamGroupScopeRpc = groups.TeamPermissionsMgmtService_GetTeamGroupScope_FullMethodName

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
func (c GroupsClient) Get(ctx context.Context, req *GetTeamGroupRequest) (*groups.GetTeamGroupResponse, error) {
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

// AddUsersBulk adds users to a group
func (c GroupsClient) AddUsersBulk(ctx context.Context, req *groups.AddUsersToTeamGroupsRequest) (*groups.AddUsersToTeamGroupsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.AddUsersToTeamGroups(callProperties.Ctx, req, callProperties.CallOptions...)
}

// RemoveUsers removes users from a group
func (c GroupsClient) RemoveUsersBulk(ctx context.Context, req *groups.RemoveUsersFromTeamGroupsRequest) (*groups.RemoveUsersFromTeamGroupsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.RemoveUsersFromTeamGroups(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewGroupsClient creates a new GroupsClient
func NewGroupsClient(callPropertiesCreator *CallPropertiesCreator) *GroupsClient {
	return &GroupsClient{callPropertiesCreator: callPropertiesCreator}
}
