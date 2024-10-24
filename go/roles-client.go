// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cxsdk

import (
	"context"

	roles "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/aaa/rbac/v2"
)

// CreateRoleRequest is a request type.
type CreateRoleRequest = roles.CreateRoleRequest

// UpdateRoleRequest is a request type.
type UpdateRoleRequest = roles.UpdateRoleRequest

// DeleteRoleRequest is a request type.
type DeleteRoleRequest = roles.DeleteRoleRequest

// GetCustomRoleRequest is a request type.
type GetCustomRoleRequest = roles.GetCustomRoleRequest

// ListCustomRolesRequest is a request type.
type ListCustomRolesRequest = roles.ListCustomRolesRequest

// ListSystemRolesRequest is a request type.
type ListSystemRolesRequest = roles.ListSystemRolesRequest

// RolePermissions is a type for permissions within the role.
type RolePermissions = roles.Permissions

// CustomRole is a custom role.
type CustomRole = roles.CustomRole

// CreateRoleRequestParentRoleName is a role specification for the create request.
type CreateRoleRequestParentRoleName = roles.CreateRoleRequest_ParentRoleName

// RPC names.
const (
	RolesListSystemRolesRPC = roles.RoleManagementService_ListSystemRoles_FullMethodName
	RolesListCustomRolesRPC = roles.RoleManagementService_ListCustomRoles_FullMethodName
	RolesGetCustomRoleRPC   = roles.RoleManagementService_GetCustomRole_FullMethodName
	RolesCreateRoleRPC      = roles.RoleManagementService_CreateRole_FullMethodName
	RolesUpdateRoleRPC      = roles.RoleManagementService_UpdateRole_FullMethodName
	RolesDeleteRoleRPC      = roles.RoleManagementService_DeleteRole_FullMethodName
)

// RolesClient is a client for roles.
type RolesClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new role.
func (r RolesClient) Create(ctx context.Context, req *CreateRoleRequest) (*roles.CreateRoleResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := roles.NewRoleManagementServiceClient(conn)

	return client.CreateRole(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Update updates a role.
func (r RolesClient) Update(ctx context.Context, req *UpdateRoleRequest) (*roles.UpdateRoleResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()

	client := roles.NewRoleManagementServiceClient(conn)

	return client.UpdateRole(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes a role.
func (r RolesClient) Delete(ctx context.Context, req *DeleteRoleRequest) (*roles.DeleteRoleResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()

	client := roles.NewRoleManagementServiceClient(conn)

	return client.DeleteRole(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get retrieves a role by ID.
func (r RolesClient) Get(ctx context.Context, req *GetCustomRoleRequest) (*roles.GetCustomRoleResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := roles.NewRoleManagementServiceClient(conn)

	return client.GetCustomRole(callProperties.Ctx, req, callProperties.CallOptions...)
}

// List retrieves all accessible roles.
func (r RolesClient) List(ctx context.Context, req *ListCustomRolesRequest) (*roles.ListCustomRolesResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := roles.NewRoleManagementServiceClient(conn)

	return client.ListCustomRoles(callProperties.Ctx, req, callProperties.CallOptions...)
}

// ListSystemRoles retrieves all system roles.
func (r RolesClient) ListSystemRoles(ctx context.Context, req *ListSystemRolesRequest) (*roles.ListSystemRolesResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := roles.NewRoleManagementServiceClient(conn)

	return client.ListSystemRoles(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewRolesClient creates a new RolesClient.
func NewRolesClient(c *CallPropertiesCreator) *RolesClient {
	return &RolesClient{callPropertiesCreator: c}
}
