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

// CreateViewRequest is a request to create a new view
type CreateViewRequest = vservices.CreateViewRequest

// CreateViewResponse is a response to create a new view
type CreateViewResponse = vservices.CreateViewResponse

// GetViewRequest is a request to get a view
type GetViewRequest = vservices.GetViewRequest

// GetViewResponse is a response when getting a view
type GetViewResponse = vservices.GetViewResponse

// ListViewsRequest is a request to list all views
type ListViewsRequest = vservices.ListViewsRequest

// ListViewsResponse is a response when listing views
type ListViewsResponse = vservices.ListViewsResponse

// DeleteViewRequest is a request to delete a view
type DeleteViewRequest = vservices.DeleteViewRequest

// DeleteViewResponse is a response when deleting a view
type DeleteViewResponse = vservices.DeleteViewResponse

// ReplaceViewRequest is a request to replacing a view
type ReplaceViewRequest = vservices.ReplaceViewRequest

// ReplaceViewResponse is a response to replacing a view
type ReplaceViewResponse = vservices.ReplaceViewResponse

// SearchQuery is a type for managing view properties
type SearchQuery = views.SearchQuery

// TimeSelection is a type for managing view properties
type TimeSelection = views.TimeSelection

// SelectedFilters is a type for managing view properties
type SelectedFilters = views.SelectedFilters

// ViewFilter is a type for managing view properties
type ViewFilter = views.Filter

// ViewTimeSelectionQuick is a type for managing view properties
type ViewTimeSelectionQuick = views.TimeSelection_QuickSelection

// QuickTimeSelection is a type for managing view properties
type QuickTimeSelection = views.QuickTimeSelection

// ViewTimeSelectionCustom is a type for managing view properties
type ViewTimeSelectionCustom = views.TimeSelection_CustomSelection

// View is a type for managing view properties
type View = vservices.View

const viewFeatureGroupID = "data-exploration"

// RPC Name values
const (
	CreateViewRPC  = vservices.ViewsService_CreateView_FullMethodName
	DeleteViewRPC  = vservices.ViewsService_DeleteView_FullMethodName
	GetViewRPC     = vservices.ViewsService_GetView_FullMethodName
	ListViewsRPC   = vservices.ViewsService_ListViews_FullMethodName
	ReplaceViewRPC = vservices.ViewsService_ReplaceView_FullMethodName
)

// ViewsClient is a client for the views service
type ViewsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new view
func (c ViewsClient) Create(ctx context.Context, req *CreateViewRequest) (*CreateViewResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := vservices.NewViewsServiceClient(conn)

	response, err := client.CreateView(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, CreateViewRPC, viewFeatureGroupID)
	}
	return response, nil
}

// Get gets a view by its ID
func (c ViewsClient) Get(ctx context.Context, req *GetViewRequest) (*GetViewResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := vservices.NewViewsServiceClient(conn)

	response, err := client.GetView(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetViewRPC, viewFeatureGroupID)
	}
	return response, nil
}

// List lists all views for the current team
func (c ViewsClient) List(ctx context.Context, req *ListViewsRequest) (*ListViewsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := vservices.NewViewsServiceClient(conn)

	response, err := client.ListViews(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ListViewsRPC, viewFeatureGroupID)
	}
	return response, nil
}

// Replace updates a view
func (c ViewsClient) Replace(ctx context.Context, req *ReplaceViewRequest) (*ReplaceViewResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := vservices.NewViewsServiceClient(conn)

	response, err := client.ReplaceView(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ReplaceViewRPC, viewFeatureGroupID)
	}
	return response, nil
}

// Delete deletes a view
func (c ViewsClient) Delete(ctx context.Context, req *DeleteViewRequest) (*DeleteViewResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := vservices.NewViewsServiceClient(conn)

	response, err := client.DeleteView(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, DeleteViewRPC, viewFeatureGroupID)
	}
	return response, nil
}

// NewViewsClient creates a new ViewsClient
func NewViewsClient(c *CallPropertiesCreator) *ViewsClient {
	return &ViewsClient{callPropertiesCreator: c}
}
