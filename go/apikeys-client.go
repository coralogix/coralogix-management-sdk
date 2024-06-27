package cxsdk

import (
	"context"
	apikeys "coralogix-management-sdk/go/internal/coralogixapis/aaa/apikeys/v3"
)

type ApikeysClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

type CreateApiKeyRequest = apikeys.CreateApiKeyRequest
type GetApiKeyRequest = apikeys.GetApiKeyRequest
type UpdateApiKeyRequest = apikeys.UpdateApiKeyRequest
type DeleteApiKeyRequest = apikeys.DeleteApiKeyRequest
type ApiKeyPermissions = apikeys.CreateApiKeyRequest_KeyPermissions

type Owner = apikeys.Owner
type Owner_UserId = apikeys.Owner_UserId
type Owner_TeamId = apikeys.Owner_TeamId
type Owner_OrganisationId = apikeys.Owner_OrganisationId

func (t ApikeysClient) CreateApiKey(ctx context.Context, req *apikeys.CreateApiKeyRequest) (*apikeys.CreateApiKeyResponse, error) {
	callProperties, err := t.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := apikeys.NewApiKeysServiceClient(conn)

	return client.CreateApiKey(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (t ApikeysClient) GetApiKey(ctx context.Context, req *apikeys.GetApiKeyRequest) (*apikeys.GetApiKeyResponse, error) {
	callProperties, err := t.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := apikeys.NewApiKeysServiceClient(conn)

	return client.GetApiKey(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (t ApikeysClient) UpdateApiKey(ctx context.Context, req *apikeys.UpdateApiKeyRequest) (*apikeys.UpdateApiKeyResponse, error) {
	callProperties, err := t.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := apikeys.NewApiKeysServiceClient(conn)

	return client.UpdateApiKey(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (t ApikeysClient) DeleteApiKey(ctx context.Context, req *apikeys.DeleteApiKeyRequest) (*apikeys.DeleteApiKeyResponse, error) {
	callProperties, err := t.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := apikeys.NewApiKeysServiceClient(conn)

	return client.DeleteApiKey(callProperties.Ctx, req, callProperties.CallOptions...)
}

func NewApiKeysClient(c *CallPropertiesCreator) *ApikeysClient {
	return &ApikeysClient{callPropertiesCreator: c}
}
