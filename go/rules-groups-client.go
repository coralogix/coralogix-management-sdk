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

	rulesgroups "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/rules/v1"
)

// RuleGroupsClient is a client for the Coralogix Rules Groups API.
type RuleGroupsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new rule group.
func (r RuleGroupsClient) Create(ctx context.Context, req *rulesgroups.CreateRuleGroupRequest) (*rulesgroups.CreateRuleGroupResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := rulesgroups.NewRuleGroupsServiceClient(conn)

	return client.CreateRuleGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get gets a rule group.
func (r RuleGroupsClient) Get(ctx context.Context, req *rulesgroups.GetRuleGroupRequest) (*rulesgroups.GetRuleGroupResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := rulesgroups.NewRuleGroupsServiceClient(conn)

	return client.GetRuleGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Update updates a rule group.
func (r RuleGroupsClient) Update(ctx context.Context, req *rulesgroups.UpdateRuleGroupRequest) (*rulesgroups.UpdateRuleGroupResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()

	client := rulesgroups.NewRuleGroupsServiceClient(conn)

	return client.UpdateRuleGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes a rule group.
func (r RuleGroupsClient) Delete(ctx context.Context, req *rulesgroups.DeleteRuleGroupRequest) (*rulesgroups.DeleteRuleGroupResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
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
