package cxsdk

import (
	"context"
	actions "coralogix-management-sdk/go/internal/coralogixapis/actions/v2"
)

type ActionsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

type CreateActionRequest = actions.CreateActionRequest
type ReplaceActionRequest = actions.ReplaceActionRequest
type DeleteActionRequest = actions.DeleteActionRequest
type GetActionRequest = actions.GetActionRequest

type Action = actions.Action

const (
	SourceType_SOURCE_TYPE_UNSPECIFIED = actions.SourceType_SOURCE_TYPE_UNSPECIFIED
	SourceType_SOURCE_TYPE_LOG         = actions.SourceType_SOURCE_TYPE_LOG
	SourceType_SOURCE_TYPE_DATA_MAP    = actions.SourceType_SOURCE_TYPE_DATA_MAP
)

func (a ActionsClient) CreateAction(ctx context.Context, req *CreateActionRequest) (*actions.CreateActionResponse, error) {
	callProperties, err := a.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := actions.NewActionsServiceClient(conn)

	return client.CreateAction(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (a ActionsClient) GetAction(ctx context.Context, req *GetActionRequest) (*actions.GetActionResponse, error) {
	callProperties, err := a.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := actions.NewActionsServiceClient(conn)

	return client.GetAction(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (a ActionsClient) UpdateAction(ctx context.Context, req *ReplaceActionRequest) (*actions.ReplaceActionResponse, error) {
	callProperties, err := a.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := actions.NewActionsServiceClient(conn)

	return client.ReplaceAction(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (a ActionsClient) DeleteAction(ctx context.Context, req *DeleteActionRequest) (*actions.DeleteActionResponse, error) {
	callProperties, err := a.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := actions.NewActionsServiceClient(conn)

	return client.DeleteAction(callProperties.Ctx, req, callProperties.CallOptions...)
}

func NewActionsClient(c *CallPropertiesCreator) *ActionsClient {
	return &ActionsClient{callPropertiesCreator: c}
}
