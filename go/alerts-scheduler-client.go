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

	scheduler "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1"

	metaLabels "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerting/meta_labels_protobuf/v1"
)

// AlertSchedulerRule is an alert scheduler rule.
type AlertSchedulerRule = scheduler.AlertSchedulerRule

// AlertSchedulerFilter is a filter for an alert scheduler rule.
type AlertSchedulerFilter = scheduler.Filter

// AlertSchedulerFilterUniqueIDs is a filter for an alert scheduler rule that contains unique IDs.
type AlertSchedulerFilterUniqueIDs = scheduler.Filter_AlertUniqueIds

// AlertUniqueIDs is a list of alert unique IDs.
type AlertUniqueIDs = scheduler.AlertUniqueIds

// CreateAlertSchedulerRuleRequest is a request to create an alert scheduler rule.
type CreateAlertSchedulerRuleRequest = scheduler.CreateAlertSchedulerRuleRequest

// UpdateAlertSchedulerRuleRequest is a request to update an alert scheduler rule.
type UpdateAlertSchedulerRuleRequest = scheduler.UpdateAlertSchedulerRuleRequest

// GetAlertSchedulerRuleRequest is a request to get an alert scheduler rule.
type GetAlertSchedulerRuleRequest = scheduler.GetAlertSchedulerRuleRequest

// DeleteAlertSchedulerRuleRequest is a request to delete an alert scheduler rule.
type DeleteAlertSchedulerRuleRequest = scheduler.DeleteAlertSchedulerRuleRequest

// Schedule is a schedule.
type Schedule = scheduler.Schedule

// ScheduleOneTime is a one-time schedule container.
type ScheduleOneTime = scheduler.Schedule_OneTime

// OneTime is a one-time schedule.
type OneTime = scheduler.OneTime

// Timeframe is a timeframe.
type Timeframe = scheduler.Timeframe

// TimeframeEndTime is the end time of a timeframe.
type TimeframeEndTime = scheduler.Timeframe_EndTime

// MetaLabel is a piece of metadata.
type MetaLabel = metaLabels.MetaLabel

const (
	// ScheduleOperationActivate is a ScheduleOperation that activates an alert scheduler.
	ScheduleOperationActivate = scheduler.ScheduleOperation_SCHEDULE_OPERATION_ACTIVATE

	// ScheduleOperationUnspecified is an unspecified ScheduleOperation.
	ScheduleOperationUnspecified = scheduler.ScheduleOperation_SCHEDULE_OPERATION_UNSPECIFIED

	// ScheduleOperationMute is a ScheduleOperation that mutes an alert scheduler.
	ScheduleOperationMute = scheduler.ScheduleOperation_SCHEDULE_OPERATION_MUTE
)

// RPC names.
const (
	CreateAlertSchedulerRuleRPC     = scheduler.AlertSchedulerRuleService_CreateAlertSchedulerRule_FullMethodName
	UpdateAlertSchedulerRuleRPC     = scheduler.AlertSchedulerRuleService_UpdateAlertSchedulerRule_FullMethodName
	GetAlertSchedulerRuleRPC        = scheduler.AlertSchedulerRuleService_GetAlertSchedulerRule_FullMethodName
	DeleteAlertSchedulerRuleRPC     = scheduler.AlertSchedulerRuleService_DeleteAlertSchedulerRule_FullMethodName
	GetBulkAlertSchedulerRuleRPC    = scheduler.AlertSchedulerRuleService_GetBulkAlertSchedulerRule_FullMethodName
	CreateBulkAlertSchedulerRuleRPC = scheduler.AlertSchedulerRuleService_CreateBulkAlertSchedulerRule_FullMethodName
	UpdateBulkAlertSchedulerRuleRPC = scheduler.AlertSchedulerRuleService_UpdateBulkAlertSchedulerRule_FullMethodName
)

// AlertSchedulerClient is a client for the Coralogix Alerts API.
type AlertSchedulerClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new alert scheduler.
func (c AlertSchedulerClient) Create(ctx context.Context, req *scheduler.CreateAlertSchedulerRuleRequest) (*scheduler.CreateAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := scheduler.NewAlertSchedulerRuleServiceClient(conn)

	return client.CreateAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get gets an alert scheduler.
func (c AlertSchedulerClient) Get(ctx context.Context, req *scheduler.GetAlertSchedulerRuleRequest) (*scheduler.GetAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := scheduler.NewAlertSchedulerRuleServiceClient(conn)

	return client.GetAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Update updates an alert scheduler.
func (c AlertSchedulerClient) Update(ctx context.Context, req *scheduler.UpdateAlertSchedulerRuleRequest) (*scheduler.UpdateAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := scheduler.NewAlertSchedulerRuleServiceClient(conn)

	return client.UpdateAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes an alert scheduler.
func (c AlertSchedulerClient) Delete(ctx context.Context, req *scheduler.DeleteAlertSchedulerRuleRequest) (*scheduler.DeleteAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := scheduler.NewAlertSchedulerRuleServiceClient(conn)

	return client.DeleteAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

// CreateBulk creates multiple alert schedulers in bulk.
func (c AlertSchedulerClient) CreateBulk(ctx context.Context, req *scheduler.CreateBulkAlertSchedulerRuleRequest) (*scheduler.CreateBulkAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := scheduler.NewAlertSchedulerRuleServiceClient(conn)

	return client.CreateBulkAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

// UpdateBulk updates multiple alert schedulers in bulk.
func (c AlertSchedulerClient) UpdateBulk(ctx context.Context, req *scheduler.UpdateBulkAlertSchedulerRuleRequest) (*scheduler.UpdateBulkAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := scheduler.NewAlertSchedulerRuleServiceClient(conn)

	return client.UpdateBulkAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetBulk gets multiple alert schedulers in bulk.
func (c AlertSchedulerClient) GetBulk(ctx context.Context, req *scheduler.GetBulkAlertSchedulerRuleRequest) (*scheduler.GetBulkAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := scheduler.NewAlertSchedulerRuleServiceClient(conn)

	return client.GetBulkAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewAlertSchedulerClient creates a new alerts scheduler client.
func NewAlertSchedulerClient(c *CallPropertiesCreator) *AlertSchedulerClient {
	return &AlertSchedulerClient{callPropertiesCreator: c}
}
