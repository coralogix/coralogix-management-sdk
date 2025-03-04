// Copyright 2025 Coralogix Ltd.
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

	views "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/views/v1"

	vservices "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/views/v1/services"
)

// CreateViewFolderRequest is a request to create a new view
type CreateViewFolderRequest = vservices.CreateViewFolderRequest

// CreateViewFolderResponse is a response to create a new view
type CreateViewFolderResponse = vservices.CreateViewFolderResponse

// GetViewFolderRequest is a request to get a view folder
type GetViewFolderRequest = vservices.GetViewFolderRequest

// GetViewFolderResponse is a response when getting a view folder
type GetViewFolderResponse = vservices.GetViewFolderResponse

// ListViewFoldersRequest is a request to list all views
type ListViewFoldersRequest = vservices.ListViewFoldersRequest

// ListViewsResponse is a response when listing views
type ListViewFoldersResponse = vservices.ListViewFoldersResponse

// DeleteViewFolderRequest is a request to delete a view folder
type DeleteViewFolderRequest = vservices.DeleteViewFolderRequest

// DeleteViewFolderResponse is a response when deleting a view folder
type DeleteViewFolderResponse = vservices.DeleteViewFolderResponse

// ReplaceViewFolderRequest is a request to replacing a view folder
type ReplaceViewFolderRequest = vservices.ReplaceViewFolderRequest

// ReplaceViewFolderResponse is a response to replacing a view folder
type ReplaceViewFolderResponse = vservices.ReplaceViewFolderResponse

type ViewFolder = views.ViewFolder

const viewFolderFeatureGroupID = "data-exploration"

// RPC Name values
const (
	CreateViewFolderRPC  = vservices.ViewsFoldersService_CreateViewFolder_FullMethodName
	DeleteViewFolderRPC  = vservices.ViewsFoldersService_DeleteViewFolder_FullMethodName
	GetViewFolderRPC     = vservices.ViewsFoldersService_GetViewFolder_FullMethodName
	ListViewFoldersRPC   = vservices.ViewsFoldersService_ListViewFolders_FullMethodName
	ReplaceViewFolderRPC = vservices.ViewsFoldersService_ReplaceViewFolder_FullMethodName
)

// ViewFoldersClient is a client for the view folders service
type ViewFoldersClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new view folder
func (c ViewFoldersClient) Create(ctx context.Context, req *CreateViewFolderRequest) (*CreateViewFolderResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := vservices.NewViewsFoldersServiceClient(conn)

	response, err := client.CreateViewFolder(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, CreateViewFolderRPC, viewFolderFeatureGroupID)
	}
	return response, nil
}

// Get gets a view folder by its ID
func (c ViewFoldersClient) Get(ctx context.Context, req *GetViewFolderRequest) (*GetViewFolderResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := vservices.NewViewsFoldersServiceClient(conn)

	response, err := client.GetViewFolder(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetViewFolderRPC, viewFolderFeatureGroupID)
	}
	return response, nil
}

// List lists all view folders
func (c ViewFoldersClient) List(ctx context.Context, req *ListViewFoldersRequest) (*ListViewFoldersResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := vservices.NewViewsFoldersServiceClient(conn)

	response, err := client.ListViewFolders(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ListViewFoldersRPC, viewFolderFeatureGroupID)
	}
	return response, nil
}

// Replace updates a view folder
func (c ViewFoldersClient) Replace(ctx context.Context, req *ReplaceViewFolderRequest) (*ReplaceViewFolderResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := vservices.NewViewsFoldersServiceClient(conn)

	response, err := client.ReplaceViewFolder(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ReplaceViewFolderRPC, viewFolderFeatureGroupID)
	}
	return response, nil
}

// Delete deletes a view folder
func (c ViewFoldersClient) Delete(ctx context.Context, req *DeleteViewFolderRequest) (*DeleteViewFolderResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := vservices.NewViewsFoldersServiceClient(conn)

	response, err := client.DeleteViewFolder(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, DeleteViewFolderRPC, viewFolderFeatureGroupID)
	}
	return response, nil
}

// NewViewFoldersClient creates a new ViewFoldersClient
func NewViewFoldersClient(c *CallPropertiesCreator) *ViewFoldersClient {
	return &ViewFoldersClient{callPropertiesCreator: c}
}
