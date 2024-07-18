package cxsdk

import (
	"context"
	rulesgroups "coralogix-management-sdk/go/internal/coralogix/rules/v1"
)

// RuleGroupsClient is a client for the Coralogix Rules Groups API.
type RuleGroupsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// CreateRuleGroup creates a new rule group.
func (r RuleGroupsClient) CreateRuleGroup(ctx context.Context, req *rulesgroups.CreateRuleGroupRequest) (*rulesgroups.CreateRuleGroupResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := rulesgroups.NewRuleGroupsServiceClient(conn)

	return client.CreateRuleGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetRuleGroup gets a rule group.
func (r RuleGroupsClient) GetRuleGroup(ctx context.Context, req *rulesgroups.GetRuleGroupRequest) (*rulesgroups.GetRuleGroupResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := rulesgroups.NewRuleGroupsServiceClient(conn)

	return client.GetRuleGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// UpdateRuleGroup updates a rule group.
func (r RuleGroupsClient) UpdateRuleGroup(ctx context.Context, req *rulesgroups.UpdateRuleGroupRequest) (*rulesgroups.UpdateRuleGroupResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()

	client := rulesgroups.NewRuleGroupsServiceClient(conn)

	return client.UpdateRuleGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// DeleteRuleGroup deletes a rule group.
func (r RuleGroupsClient) DeleteRuleGroup(ctx context.Context, req *rulesgroups.DeleteRuleGroupRequest) (*rulesgroups.DeleteRuleGroupResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()

	client := rulesgroups.NewRuleGroupsServiceClient(conn)

	return client.DeleteRuleGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewRuleGroupsClient creates a new rule groups client.
func NewRuleGroupsClient(c *CallPropertiesCreator) *RuleGroupsClient {
	return &RuleGroupsClient{callPropertiesCreator: c}
}
