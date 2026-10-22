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

	permissions "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/aaa/rbac/v2"
)

// ListAllPermissionsRequest is a request type for listing all permissions.
type ListAllPermissionsRequest = permissions.ListAllPermissionsRequest

// ListAllPermissionsResponse is a response type containing all permissions.
type ListAllPermissionsResponse = permissions.ListAllPermissionsResponse

// RbacPermission is a single permission entry with its canonical expression and deprecated aliases.
type RbacPermission = permissions.Permission

const permissionsFeatureGroupID = "permissions"

// RPC names.
const (
	PermissionsListAllRPC = permissions.PermissionsService_ListAllPermissions_FullMethodName
)

// PermissionsClient is a client for the PermissionsService.
type PermissionsClient struct {
	callPropertiesCreator CallPropertiesCreator
}

// ListAll retrieves all permissions including their deprecated expression aliases.
func (c PermissionsClient) ListAll(ctx context.Context, req *ListAllPermissionsRequest) (*ListAllPermissionsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}
	conn := callProperties.Connection
	defer conn.Close()
	client := permissions.NewPermissionsServiceClient(conn)

	response, err := client.ListAllPermissions(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, PermissionsListAllRPC, permissionsFeatureGroupID)
	}
	return response, nil
}

// NewPermissionsClient creates a new PermissionsClient.
func NewPermissionsClient(c CallPropertiesCreator) *PermissionsClient {
	return &PermissionsClient{callPropertiesCreator: c}
}
