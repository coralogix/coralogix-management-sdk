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

	alertsSchedulers "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1"

	metaLabels "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/alerting/meta_labels_protobuf/v1"
)

// AlertSchedulerRule is an alert scheduler rule.
type AlertSchedulerRule = alertsSchedulers.AlertSchedulerRule

// AlertSchedulerFilter is a filter for an alert scheduler rule.
type AlertSchedulerFilter = alertsSchedulers.Filter

// AlertSchedulerFilterUniqueIDs is a filter for an alert scheduler rule that contains unique IDs.
type AlertSchedulerFilterUniqueIDs = alertsSchedulers.Filter_AlertUniqueIds

// AlertUniqueIDs is a list of alert unique IDs.
type AlertUniqueIDs = alertsSchedulers.AlertUniqueIds

// CreateAlertSchedulerRuleRequest is a request to create an alert scheduler rule.
type CreateAlertSchedulerRuleRequest = alertsSchedulers.CreateAlertSchedulerRuleRequest

// UpdateAlertSchedulerRuleRequest is a request to update an alert scheduler rule.
type UpdateAlertSchedulerRuleRequest = alertsSchedulers.UpdateAlertSchedulerRuleRequest

// GetAlertSchedulerRuleRequest is a request to get an alert scheduler rule.
type GetAlertSchedulerRuleRequest = alertsSchedulers.GetAlertSchedulerRuleRequest

// DeleteAlertSchedulerRuleRequest is a request to delete an alert scheduler rule.
type DeleteAlertSchedulerRuleRequest = alertsSchedulers.DeleteAlertSchedulerRuleRequest

// Schedule is a schedule.
type Schedule = alertsSchedulers.Schedule

// ScheduleOneTime is a one-time schedule container.
type ScheduleOneTime = alertsSchedulers.Schedule_OneTime

// OneTime is a one-time schedule.
type OneTime = alertsSchedulers.OneTime

// Timeframe is a timeframe.
type Timeframe = alertsSchedulers.Timeframe

// TimeframeEndTime is the end time of a timeframe.
type TimeframeEndTime = alertsSchedulers.Timeframe_EndTime

// MetaLabel is a piece of metadata.
type MetaLabel = metaLabels.MetaLabel

// ScheduleOperationScheduleOperationActivate is a ScheduleOperation that activates an alert scheduler.
const ScheduleOperationScheduleOperationActivate = alertsSchedulers.ScheduleOperation_SCHEDULE_OPERATION_ACTIVATE

// ScheduleOperationScheduleOperationUnspecified is an unspecified ScheduleOperation.
const ScheduleOperationScheduleOperationUnspecified = alertsSchedulers.ScheduleOperation_SCHEDULE_OPERATION_UNSPECIFIED

// ScheduleOperationScheduleOperationMute is a ScheduleOperation that mutes an alert scheduler.
const ScheduleOperationScheduleOperationMute = alertsSchedulers.ScheduleOperation_SCHEDULE_OPERATION_MUTE

// AlertsSchedulersClient is a client for the Coralogix Alerts API.
type AlertsSchedulersClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new alert scheduler.
func (c AlertsSchedulersClient) Create(ctx context.Context, req *alertsSchedulers.CreateAlertSchedulerRuleRequest) (*alertsSchedulers.CreateAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alertsSchedulers.NewAlertSchedulerRuleServiceClient(conn)

	return client.CreateAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get gets an alert scheduler.
func (c AlertsSchedulersClient) Get(ctx context.Context, req *alertsSchedulers.GetAlertSchedulerRuleRequest) (*alertsSchedulers.GetAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alertsSchedulers.NewAlertSchedulerRuleServiceClient(conn)

	return client.GetAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Update updates an alert scheduler.
func (c AlertsSchedulersClient) Update(ctx context.Context, req *alertsSchedulers.UpdateAlertSchedulerRuleRequest) (*alertsSchedulers.UpdateAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alertsSchedulers.NewAlertSchedulerRuleServiceClient(conn)

	return client.UpdateAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes an alert scheduler.
func (c AlertsSchedulersClient) Delete(ctx context.Context, req *alertsSchedulers.DeleteAlertSchedulerRuleRequest) (*alertsSchedulers.DeleteAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alertsSchedulers.NewAlertSchedulerRuleServiceClient(conn)

	return client.DeleteAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

// CreateBulk creates multiple alert schedulers in bulk.
func (c AlertsSchedulersClient) CreateBulk(ctx context.Context, req *alertsSchedulers.CreateBulkAlertSchedulerRuleRequest) (*alertsSchedulers.CreateBulkAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alertsSchedulers.NewAlertSchedulerRuleServiceClient(conn)

	return client.CreateBulkAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

// UpdateBulk updates multiple alert schedulers in bulk.
func (c AlertsSchedulersClient) UpdateBulk(ctx context.Context, req *alertsSchedulers.UpdateBulkAlertSchedulerRuleRequest) (*alertsSchedulers.UpdateBulkAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alertsSchedulers.NewAlertSchedulerRuleServiceClient(conn)

	return client.UpdateBulkAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetBulk gets multiple alert schedulers in bulk.
func (c AlertsSchedulersClient) GetBulk(ctx context.Context, req *alertsSchedulers.GetBulkAlertSchedulerRuleRequest) (*alertsSchedulers.GetBulkAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alertsSchedulers.NewAlertSchedulerRuleServiceClient(conn)

	return client.GetBulkAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewAlertsSchedulersClient creates a new alerts scheduler client.
func NewAlertsSchedulersClient(c *CallPropertiesCreator) *AlertsSchedulersClient {
	return &AlertsSchedulersClient{callPropertiesCreator: c}
}
