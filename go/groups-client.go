package cxsdk

import (
	"context"

	groups "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/permissions/v1"
)

// GroupsClient is a client for the Groups API
type GroupsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new group
func (c GroupsClient) Create(ctx context.Context, req *groups.CreateTeamGroupRequest) (*groups.CreateTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.CreateTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get retrieves a group by ID
func (c GroupsClient) Get(ctx context.Context, req *groups.GetTeamGroupRequest) (*groups.GetTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.GetTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// List retrieves all groups in the team
func (c GroupsClient) List(ctx context.Context, req *groups.GetTeamGroupsRequest) (*groups.GetTeamGroupsResponse, error) {
	callPoperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callPoperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.GetTeamGroups(callPoperties.Ctx, req, callPoperties.CallOptions...)
}

// Update updates a group
func (c GroupsClient) Update(ctx context.Context, req *groups.UpdateTeamGroupRequest) (*groups.UpdateTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.UpdateTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes a group by ID
func (c GroupsClient) Delete(ctx context.Context, req *groups.DeleteTeamGroupRequest) (*groups.DeleteTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.DeleteTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// AddUsers adds users to a group
func (c GroupsClient) AddUsers(ctx context.Context, req *groups.AddUsersToTeamGroupRequest) (*groups.AddUsersToTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.AddUsersToTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// RemoveUsers removes users from a group
func (c GroupsClient) RemoveUsers(ctx context.Context, req *groups.RemoveUsersFromTeamGroupRequest) (*groups.RemoveUsersFromTeamGroupResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := groups.NewTeamPermissionsMgmtServiceClient(conn)

	return client.RemoveUsersFromTeamGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewGroupsClient creates a new GroupsClient
func NewGroupsClient(callPropertiesCreator *CallPropertiesCreator) *GroupsClient {
	return &GroupsClient{callPropertiesCreator: callPropertiesCreator}
}
