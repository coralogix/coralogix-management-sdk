package cxsdk

import (
	"context"
	webhooks "coralogix-management-sdk/go/internal/coralogix/outgoing_webhooks/v1"
)

type CreateOutgoingWebhookRequest = webhooks.CreateOutgoingWebhookRequest
type GetOutgoingWebhookRequest = webhooks.GetOutgoingWebhookRequest
type UpdateOutgoingWebhookRequest = webhooks.UpdateOutgoingWebhookRequest
type DeleteOutgoingWebhookRequest = webhooks.DeleteOutgoingWebhookRequest
type ListOutgoingWebhookTypesRequest = webhooks.ListOutgoingWebhookTypesRequest
type ListAllOutgoingWebhooksRequest = webhooks.ListAllOutgoingWebhooksRequest
type TestExistingOutgoingWebhookRequest = webhooks.TestExistingOutgoingWebhookRequest
type TestOutgoingWebhookRequest = webhooks.TestOutgoingWebhookRequest

type OutgoingWebhookInputData = webhooks.OutgoingWebhookInputData
type WebhookType = webhooks.WebhookType

const (
	WebhookType_UNKNOWN                 = webhooks.WebhookType_UNKNOWN
	WebhookType_GENERIC                 = webhooks.WebhookType_GENERIC
	WebhookType_SLACK                   = webhooks.WebhookType_SLACK
	WebhookType_PAGERDUTY               = webhooks.WebhookType_PAGERDUTY
	WebhookType_SEND_LOG                = webhooks.WebhookType_SEND_LOG
	WebhookType_EMAIL_GROUP             = webhooks.WebhookType_EMAIL_GROUP
	WebhookType_MICROSOFT_TEAMS         = webhooks.WebhookType_MICROSOFT_TEAMS
	WebhookType_JIRA                    = webhooks.WebhookType_JIRA
	WebhookType_OPSGENIE                = webhooks.WebhookType_OPSGENIE
	WebhookType_DEMISTO                 = webhooks.WebhookType_DEMISTO
	WebhookType_AWS_EVENT_BRIDGE        = webhooks.WebhookType_AWS_EVENT_BRIDGE
	WebhookType_IBM_EVENT_NOTIFICATIONS = webhooks.WebhookType_IBM_EVENT_NOTIFICATIONS
)

type GenericWebhook = webhooks.OutgoingWebhookInputData_GenericWebhook
type Slack = webhooks.OutgoingWebhookInputData_Slack
type PagerDuty = webhooks.OutgoingWebhookInputData_PagerDuty
type SendLog = webhooks.OutgoingWebhookInputData_SendLog
type EmailGroup = webhooks.OutgoingWebhookInputData_EmailGroup
type MicrosoftTeams = webhooks.OutgoingWebhookInputData_MicrosoftTeams
type Jira = webhooks.OutgoingWebhookInputData_Jira
type Opsgenie = webhooks.OutgoingWebhookInputData_Opsgenie
type Demisto = webhooks.OutgoingWebhookInputData_Demisto
type AwsEventBridge = webhooks.OutgoingWebhookInputData_AwsEventBridge
type IbmEventNotifications = webhooks.OutgoingWebhookInputData_IbmEventNotifications

const (
	GenericWebhookConfig_UNKNOWN = webhooks.GenericWebhookConfig_UNKNOWN
	GenericWebhookConfig_GET     = webhooks.GenericWebhookConfig_GET
	GenericWebhookConfig_POST    = webhooks.GenericWebhookConfig_POST
	GenericWebhookConfig_PUT     = webhooks.GenericWebhookConfig_PUT
)

type GenericWebhookConfig = webhooks.GenericWebhookConfig

const (
	// attachment
	SlackConfig_EMPTY           = webhooks.SlackConfig_EMPTY
	SlackConfig_METRIC_SNAPSHOT = webhooks.SlackConfig_METRIC_SNAPSHOT
	SlackConfig_LOGS            = webhooks.SlackConfig_LOGS
	// digests
	SlackConfig_UNKNOWN                 = webhooks.SlackConfig_UNKNOWN
	SlackConfig_ERROR_AND_CRITICAL_LOGS = webhooks.SlackConfig_ERROR_AND_CRITICAL_LOGS
	SlackConfig_FLOW_ANOMALIES          = webhooks.SlackConfig_FLOW_ANOMALIES
	SlackConfig_SPIKE_ANOMALIES         = webhooks.SlackConfig_SPIKE_ANOMALIES
	SlackConfig_DATA_USAGE              = webhooks.SlackConfig_DATA_USAGE
)

type SlackConfig_Digest = webhooks.SlackConfig_Digest
type SlackConfig_Attachment = webhooks.SlackConfig_Attachment
type SlackConfig = webhooks.SlackConfig

type PagerDutyConfig = webhooks.PagerDutyConfig
type SendLogConfig = webhooks.SendLogConfig
type EmailGroupConfig = webhooks.EmailGroupConfig
type MicrosoftTeamsConfig = webhooks.MicrosoftTeamsConfig
type JiraConfig = webhooks.JiraConfig
type OpsgenieConfig = webhooks.OpsgenieConfig
type DemistoConfig = webhooks.DemistoConfig
type AwsEventBridgeConfig = webhooks.AwsEventBridgeConfig
type IbmEventNotificationsConfig = webhooks.IbmEventNotificationsConfig

type TestOutgoingWebhookResponse_Success = webhooks.TestOutgoingWebhookResponse_Success
type TestOutgoingWebhookResponse_Failure = webhooks.TestOutgoingWebhookResponse_Failure

// WebhooksClient is a client for the Coralogix Webhooks API.
type WebhooksClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// CreateWebhook creates a new webhook.
func (c WebhooksClient) Create(ctx context.Context, req *CreateOutgoingWebhookRequest) (*webhooks.CreateOutgoingWebhookResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := webhooks.NewOutgoingWebhooksServiceClient(conn)

	return client.CreateOutgoingWebhook(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Gets the specified webhook.
func (c WebhooksClient) Get(ctx context.Context, req *GetOutgoingWebhookRequest) (*webhooks.GetOutgoingWebhookResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := webhooks.NewOutgoingWebhooksServiceClient(conn)

	return client.GetOutgoingWebhook(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Replaces the specified webhook.
func (c WebhooksClient) Replace(ctx context.Context, req *UpdateOutgoingWebhookRequest) (*webhooks.UpdateOutgoingWebhookResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := webhooks.NewOutgoingWebhooksServiceClient(conn)

	return client.UpdateOutgoingWebhook(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Deletes the specified webhook.
func (c WebhooksClient) Delete(ctx context.Context, req *DeleteOutgoingWebhookRequest) (*webhooks.DeleteOutgoingWebhookResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := webhooks.NewOutgoingWebhooksServiceClient(conn)

	return client.DeleteOutgoingWebhook(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Lists all webhook types.
func (c WebhooksClient) ListTypes(ctx context.Context, req *ListOutgoingWebhookTypesRequest) (*webhooks.ListOutgoingWebhookTypesResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := webhooks.NewOutgoingWebhooksServiceClient(conn)

	return client.ListOutgoingWebhookTypes(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Lists all webhooks.
func (c WebhooksClient) List(ctx context.Context, req *ListAllOutgoingWebhooksRequest) (*webhooks.ListAllOutgoingWebhooksResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := webhooks.NewOutgoingWebhooksServiceClient(conn)

	return client.ListAllOutgoingWebhooks(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Tests an existing webhook.
func (c WebhooksClient) TestById(ctx context.Context, req *TestExistingOutgoingWebhookRequest) (*webhooks.TestOutgoingWebhookResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := webhooks.NewOutgoingWebhooksServiceClient(conn)

	return client.TestExistingOutgoingWebhook(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Tests the included webhook.
func (c WebhooksClient) Test(ctx context.Context, req *TestOutgoingWebhookRequest) (*webhooks.TestOutgoingWebhookResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := webhooks.NewOutgoingWebhooksServiceClient(conn)

	return client.TestOutgoingWebhook(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewWebhooksClient creates a new webhooks client.
func NewWebhooksClient(c *CallPropertiesCreator) *WebhooksClient {
	return &WebhooksClient{callPropertiesCreator: c}
}
