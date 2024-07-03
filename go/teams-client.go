package cxsdk

import (
	"context"
	teams "coralogix-management-sdk/go/internal/coralogixapis/aaa/organisations/v2"
)

type CreateTeamInOrgRequest = teams.CreateTeamInOrgRequest
type UpdateTeamRequest = teams.UpdateTeamRequest
type GetTeamRequest = teams.GetTeamRequest
type DeleteTeamRequest = teams.DeleteTeamRequest
type SetDailyQuotaRequest = teams.SetDailyQuotaRequest
type GetTeamQuotaRequest = teams.GetTeamQuotaRequest
type MoveQuotaRequest = teams.MoveQuotaRequest

type TeamId = teams.TeamId

type TeamsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

func (c TeamsClient) CreateTeam(ctx context.Context, req *CreateTeamInOrgRequest) (*teams.CreateTeamInOrgResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := teams.NewTeamServiceClient(conn)

	return client.CreateTeamInOrg(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (c TeamsClient) UpdateTeam(ctx context.Context, req *UpdateTeamRequest) (*teams.UpdateTeamResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := teams.NewTeamServiceClient(conn)

	return client.UpdateTeam(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (c TeamsClient) GetTeam(ctx context.Context, req *GetTeamRequest) (*teams.GetTeamResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := teams.NewTeamServiceClient(conn)

	return client.GetTeam(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (c TeamsClient) DeleteTeam(ctx context.Context, req *DeleteTeamRequest) (*teams.DeleteTeamResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := teams.NewTeamServiceClient(conn)

	return client.DeleteTeam(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (c TeamsClient) SetDailyQuota(ctx context.Context, req *SetDailyQuotaRequest) (*teams.SetDailyQuotaResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := teams.NewTeamServiceClient(conn)

	return client.SetDailyQuota(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (c TeamsClient) GetQuota(ctx context.Context, req *GetTeamQuotaRequest) (*teams.GetTeamQuotaResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := teams.NewTeamServiceClient(conn)

	return client.GetTeamQuota(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (c TeamsClient) MoveQuota(ctx context.Context, req *MoveQuotaRequest) (*teams.MoveQuotaResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := teams.NewTeamServiceClient(conn)

	return client.MoveQuota(callProperties.Ctx, req, callProperties.CallOptions...)
}

func NewTeamsClient(c *CallPropertiesCreator) *TeamsClient {
	return &TeamsClient{callPropertiesCreator: c}
}
