package cxsdk

import (
	"context"
	teams "coralogix-management-sdk/go/internal/coralogixapis/aaa/organisations/v2"
)

// CreateTeamInOrgRequest is a request to create a team.
type CreateTeamInOrgRequest = teams.CreateTeamInOrgRequest

// UpdateTeamRequest is a request to update a team.
type UpdateTeamRequest = teams.UpdateTeamRequest

// GetTeamRequest is a request to get a team.
type GetTeamRequest = teams.GetTeamRequest

// DeleteTeamRequest is a request to delete a team.
type DeleteTeamRequest = teams.DeleteTeamRequest

// SetDailyQuotaRequest is a request to set the daily quota for a team.
type SetDailyQuotaRequest = teams.SetDailyQuotaRequest

// GetTeamQuotaRequest is a request to get the quota for a team.
type GetTeamQuotaRequest = teams.GetTeamQuotaRequest

// MoveQuotaRequest is a request to move the quota from one team to another.
type MoveQuotaRequest = teams.MoveQuotaRequest

// TeamID identifies a team.
type TeamID = teams.TeamId

// TeamsClient is a client for the Coralogix Teams API.
type TeamsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// CreateTeam creates a new team.
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

// UpdateTeam updates a team.
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

// GetTeam gets a team.
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

// DeleteTeam deletes a team.
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

// SetDailyQuota sets the daily quota for a team.
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

// GetQuota gets the quota for a team.
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

// MoveQuota moves the quota from one team to another.
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

// NewTeamsClient creates a new teams client.
func NewTeamsClient(c *CallPropertiesCreator) *TeamsClient {
	return &TeamsClient{callPropertiesCreator: c}
}
