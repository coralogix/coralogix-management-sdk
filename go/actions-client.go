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

	actions "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/actions/v2"
)

// ActionsClient is a client for the Coralogix Actions API.
type ActionsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// CreateActionRequest is a request to create an action.
type CreateActionRequest = actions.CreateActionRequest

// ReplaceActionRequest is a request to replace an action.
type ReplaceActionRequest = actions.ReplaceActionRequest

// DeleteActionRequest is a request to delete an action.
type DeleteActionRequest = actions.DeleteActionRequest

// GetActionRequest is a request to get an action.
type GetActionRequest = actions.GetActionRequest

// OrderActionsRequest is a request to order actions.
type OrderActionsRequest = actions.OrderActionsRequest

// Action is an action.
type Action = actions.Action

// SourceType is a type of source for an action.
type SourceType = actions.SourceType

const (
	// SourceTypeUnspecified is an unspecified source type.
	SourceTypeUnspecified = actions.SourceType_SOURCE_TYPE_UNSPECIFIED

	// SourceTypeLog is a webhook source type.
	SourceTypeLog = actions.SourceType_SOURCE_TYPE_LOG

	// SourceTypeDataMap is a data map source type.
	SourceTypeDataMap = actions.SourceType_SOURCE_TYPE_DATA_MAP
)

const (
	// GetActionRPC is the RPC for getting an action.
	GetActionRPC = actions.ActionsService_GetAction_FullMethodName

	// CreateActionRPC is the RPC for creating an action.
	CreateActionRPC = actions.ActionsService_CreateAction_FullMethodName

	// ReplaceActionRPC is the RPC for replacing an action.
	ReplaceActionRPC = actions.ActionsService_ReplaceAction_FullMethodName

	// DeleteActionRPC is the RPC for removing an action.
	DeleteActionRPC = actions.ActionsService_DeleteAction_FullMethodName

	// OrderActionsRPC is the RPC for ordering actions.
	OrderActionsRPC = actions.ActionsService_OrderActions_FullMethodName

	// ListActionsRPC is the RPC for listing actions.
	ListActionsRPC = actions.ActionsService_ListActions_FullMethodName

	// AtomicBatchExecuteActionsRPC is the RPC for executing actions atomically.
	AtomicBatchExecuteActionsRPC = actions.ActionsService_AtomicBatchExecuteActions_FullMethodName
)

// Create creates a new action.
func (a ActionsClient) Create(ctx context.Context, req *CreateActionRequest) (*actions.CreateActionResponse, error) {
	callProperties, err := a.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := actions.NewActionsServiceClient(conn)

	return client.CreateAction(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get gets an action.
func (a ActionsClient) Get(ctx context.Context, req *GetActionRequest) (*actions.GetActionResponse, error) {
	callProperties, err := a.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := actions.NewActionsServiceClient(conn)

	return client.GetAction(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Replace replaces an action.
func (a ActionsClient) Replace(ctx context.Context, req *ReplaceActionRequest) (*actions.ReplaceActionResponse, error) {
	callProperties, err := a.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := actions.NewActionsServiceClient(conn)

	return client.ReplaceAction(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes an action.
func (a ActionsClient) Delete(ctx context.Context, req *DeleteActionRequest) (*actions.DeleteActionResponse, error) {
	callProperties, err := a.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := actions.NewActionsServiceClient(conn)

	return client.DeleteAction(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Order sets the order of actions.
func (a ActionsClient) Order(ctx context.Context, req *actions.OrderActionsRequest) (*actions.OrderActionsResponse, error) {
	callProperties, err := a.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := actions.NewActionsServiceClient(conn)

	return client.OrderActions(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewActionsClient Creates a new actions client.
func NewActionsClient(c *CallPropertiesCreator) *ActionsClient {
	return &ActionsClient{callPropertiesCreator: c}
}
