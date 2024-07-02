package cxsdk

import (
	"context"
	alertsSchedulers "coralogix-management-sdk/go/internal/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1"

	metaLabels "coralogix-management-sdk/go/internal/coralogixapis/alerting/meta_labels_protobuf/v1"
)

type AlertSchedulerRule = alertsSchedulers.AlertSchedulerRule
type CreateAlertSchedulerRuleRequest = alertsSchedulers.CreateAlertSchedulerRuleRequest
type UpdateAlertSchedulerRuleRequest = alertsSchedulers.UpdateAlertSchedulerRuleRequest
type GetAlertSchedulerRuleRequest = alertsSchedulers.GetAlertSchedulerRuleRequest
type DeleteAlertSchedulerRuleRequest = alertsSchedulers.DeleteAlertSchedulerRuleRequest
type Schedule = alertsSchedulers.Schedule
type Schedule_OneTime = alertsSchedulers.Schedule_OneTime
type OneTime = alertsSchedulers.OneTime
type Timeframe = alertsSchedulers.Timeframe
type Timeframe_EndTime = alertsSchedulers.Timeframe_EndTime
type MetaLabel = metaLabels.MetaLabel

const (
	ScheduleOperation_SCHEDULE_OPERATION_ACTIVATE    = alertsSchedulers.ScheduleOperation_SCHEDULE_OPERATION_ACTIVATE
	ScheduleOperation_SCHEDULE_OPERATION_UNSPECIFIED = alertsSchedulers.ScheduleOperation_SCHEDULE_OPERATION_UNSPECIFIED
	ScheduleOperation_SCHEDULE_OPERATION_MUTE        = alertsSchedulers.ScheduleOperation_SCHEDULE_OPERATION_MUTE
)

type AlertsSchedulersClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

func (c AlertsSchedulersClient) CreateAlertScheduler(ctx context.Context, req *alertsSchedulers.CreateAlertSchedulerRuleRequest) (*alertsSchedulers.CreateAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alertsSchedulers.NewAlertSchedulerRuleServiceClient(conn)

	return client.CreateAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (c AlertsSchedulersClient) GetAlertScheduler(ctx context.Context, req *alertsSchedulers.GetAlertSchedulerRuleRequest) (*alertsSchedulers.GetAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alertsSchedulers.NewAlertSchedulerRuleServiceClient(conn)

	return client.GetAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (c AlertsSchedulersClient) UpdateAlertScheduler(ctx context.Context, req *alertsSchedulers.UpdateAlertSchedulerRuleRequest) (*alertsSchedulers.UpdateAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alertsSchedulers.NewAlertSchedulerRuleServiceClient(conn)

	return client.UpdateAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (c AlertsSchedulersClient) DeleteAlertScheduler(ctx context.Context, req *alertsSchedulers.DeleteAlertSchedulerRuleRequest) (*alertsSchedulers.DeleteAlertSchedulerRuleResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := alertsSchedulers.NewAlertSchedulerRuleServiceClient(conn)

	return client.DeleteAlertSchedulerRule(callProperties.Ctx, req, callProperties.CallOptions...)
}

func NewAlertsSchedulersClient(c *CallPropertiesCreator) *AlertsSchedulersClient {
	return &AlertsSchedulersClient{callPropertiesCreator: c}
}
