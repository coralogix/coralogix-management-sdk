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

// TeamGroupID is the ID of a team group.
type TeamGroupID = groups.TeamGroupId

// GroupScope is a type for a group scope.
type GroupScope = groups.Scope

// TeamGroup is a type for a group.
type TeamGroup = groups.TeamGroup

// GroupsUser is a type for a user in a Groups context.
type GroupsUser = groups.User

// RoleID is a type for a role ID to use in a group.
type RoleID = groups.RoleId

// FilterType is a type for a filter type in a group.
type FilterType = groups.FilterType

const (
	FilterTypeUnspecified = groups.FilterType_FILTER_TYPE_UNSPECIFIED
	FilterTypeStartsWith  = groups.FilterType_FILTER_TYPE_STARTS_WITH
	FilterTypeEndsWith    = groups.FilterType_FILTER_TYPE_ENDS_WITH
	FilterTypeContains    = groups.FilterType_FILTER_TYPE_CONTAINS
	FilterTypeExact       = groups.FilterType_FILTER_TYPE_EXACT
)

// ScopeFilters is a type for scope filters used in group operations.
type ScopeFilters = groups.ScopeFilters

// GroupScopeFilter is a filter for a group scope.
type GroupScopeFilter = groups.ScopeFilter

// ScopeID is a type for a scope ID to use in a group.
type ScopeID = groups.ScopeId

// Role is a type for a role in a group.
type Role = groups.Role

// UserID is a type for a user ID to use in a group.
type UserID = groups.UserId

// GroupsTeamID is a type for a team ID to use in a group.
type GroupsTeamID = groups.TeamId

// UpdateTeamGroupRequestRoleUpdates is a type for a request.
type UpdateTeamGroupRequestRoleUpdates = groups.UpdateTeamGroupRequest_RoleUpdates

// UpdateTeamGroupRequestUserUpdates is a type for a request.
type UpdateTeamGroupRequestUserUpdates = groups.UpdateTeamGroupRequest_UserUpdates

// GetTeamGroupRequest is a type for a request.
type GetTeamGroupRequest = groups.GetTeamGroupRequest

// GetTeamGroupResponse is a type for a response.
type GetTeamGroupResponse = groups.GetTeamGroupResponse

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

// GetGroupUsersResponseNoMorePages is a constant for the response when there are no more pages of users in a group.
type GetGroupUsersResponseNoMorePages = groups.GetGroupUsersResponse_NoMorePages_

// GetGroupUsersResponseToken is a type for the next page token in a group users response.
type GetGroupUsersResponseToken = groups.GetGroupUsersResponse_Token

const groupsFeatureGroupID = "aaa"

// RPC Values
const (
	CreateTeamGroupRPC          = groups.TeamPermissionsMgmtService_CreateTeamGroup_FullMethodName
	GetTeamGroupRPC             = groups.TeamPermissionsMgmtService_GetTeamGroup_FullMethodName
	GetTeamGroupsRPC            = groups.TeamPermissionsMgmtService_GetTeamGroups_FullMethodName
	UpdateTeamGroupRPC          = groups.TeamPermissionsMgmtService_UpdateTeamGroup_FullMethodName
	DeleteTeamGroupRPC          = groups.TeamPermissionsMgmtService_DeleteTeamGroup_FullMethodName
	AddUsersToTeamGroupRPC      = groups.TeamPermissionsMgmtService_AddUsersToTeamGroup_FullMethodName
	RemoveUsersFromTeamGroupRPC = groups.TeamPermissionsMgmtService_RemoveUsersFromTeamGroup_FullMethodName
	GetGroupUsersRPC            = groups.TeamPermissionsMgmtService_GetGroupUsers_FullMethodName
)

// GroupsClient is a client for the Groups API
type GroupsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new group
func (c GroupsClient) Create(ctx context.Context, req *CreateTeamGroupRequest) (*groups.CreateTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	response, err := client.CreateTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, CreateTeamGroupRPC, groupsFeatureGroupID)
	}
	return response, nil
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

	response, err := client.GetTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetTeamGroupRPC, groupsFeatureGroupID)
	}
	return response, nil
}

// List retrieves all groups in the team
func (c GroupsClient) List(ctx context.Context, req *GetTeamGroupsRequest) (*groups.GetTeamGroupsResponse, error) {
	callPoperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callPoperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	response, err := client.GetTeamGroups(callPoperties.Ctx, req, callPoperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetTeamGroupsRPC, groupsFeatureGroupID)
	}
	return response, nil
}

// Update updates a group
func (c GroupsClient) Update(ctx context.Context, req *UpdateTeamGroupRequest) (*groups.UpdateTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	response, err := client.UpdateTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, UpdateTeamGroupRPC, groupsFeatureGroupID)
	}
	return response, nil
}

// Delete deletes a group by ID
func (c GroupsClient) Delete(ctx context.Context, req *DeleteTeamGroupRequest) (*groups.DeleteTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	response, err := client.DeleteTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, DeleteTeamGroupRPC, groupsFeatureGroupID)
	}
	return response, nil
}

// AddUsers adds users to a group
func (c GroupsClient) AddUsers(ctx context.Context, req *AddUsersToTeamGroupRequest) (*groups.AddUsersToTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	response, err := client.AddUsersToTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, AddUsersToTeamGroupRPC, groupsFeatureGroupID)
	}
	return response, nil
}

// RemoveUsers removes users from a group
func (c GroupsClient) RemoveUsers(ctx context.Context, req *RemoveUsersFromTeamGroupRequest) (*groups.RemoveUsersFromTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	response, err := client.RemoveUsersFromTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, RemoveUsersFromTeamGroupRPC, groupsFeatureGroupID)
	}
	return response, nil
}

// GetUsers retrieves all users in a group
func (c GroupsClient) GetUsers(ctx context.Context, req *GetGroupUsersRequest) (*groups.GetGroupUsersResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	response, err := client.GetGroupUsers(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetGroupUsersRPC, groupsFeatureGroupID)
	}
	return response, nil
}

// NewGroupsClient creates a new GroupsClient
func NewGroupsClient(callPropertiesCreator *CallPropertiesCreator) *GroupsClient {
	return &GroupsClient{callPropertiesCreator: callPropertiesCreator}
}
