package cxsdk

import (
	"context"
	webhooks "coralogix-management-sdk/go/internal/coralogix/outgoing_webhooks/v1"
)

// CreateOutgoingWebhookRequest is a request to create a new webhook.
type CreateOutgoingWebhookRequest = webhooks.CreateOutgoingWebhookRequest

// GetOutgoingWebhookRequest is a request to get a webhook by its ID.
type GetOutgoingWebhookRequest = webhooks.GetOutgoingWebhookRequest

// UpdateOutgoingWebhookRequest is a request to update a webhook.
type UpdateOutgoingWebhookRequest = webhooks.UpdateOutgoingWebhookRequest

// DeleteOutgoingWebhookRequest is a request to delete a webhook.
type DeleteOutgoingWebhookRequest = webhooks.DeleteOutgoingWebhookRequest

// ListOutgoingWebhookTypesRequest is a request to list all webhook types.
type ListOutgoingWebhookTypesRequest = webhooks.ListOutgoingWebhookTypesRequest

// ListAllOutgoingWebhooksRequest is a request to list all webhooks.
type ListAllOutgoingWebhooksRequest = webhooks.ListAllOutgoingWebhooksRequest

// TestExistingOutgoingWebhookRequest is a request to test an existing webhook.
type TestExistingOutgoingWebhookRequest = webhooks.TestExistingOutgoingWebhookRequest

// TestOutgoingWebhookRequest is a request to test a webhook.
type TestOutgoingWebhookRequest = webhooks.TestOutgoingWebhookRequest

// OutgoingWebhookInputData is the wrapped data for the webhooks.
type OutgoingWebhookInputData = webhooks.OutgoingWebhookInputData

// Webhook type
type WebhookType = webhooks.WebhookType

const (
	WebhookTypeUnknown               = webhooks.WebhookType_UNKNOWN
	WebhookTypeGeneric               = webhooks.WebhookType_GENERIC
	WebhookTypeSlack                 = webhooks.WebhookType_SLACK
	WebhookTypePagerduty             = webhooks.WebhookType_PAGERDUTY
	WebhookTypeSendLog               = webhooks.WebhookType_SEND_LOG
	WebhookTypeEmailGroup            = webhooks.WebhookType_EMAIL_GROUP
	WebhookTypeMicrosoftTeams        = webhooks.WebhookType_MICROSOFT_TEAMS
	WebhookTypeJira                  = webhooks.WebhookType_JIRA
	WebhookTypeOpsgenie              = webhooks.WebhookType_OPSGENIE
	WebhookTypeDemisto               = webhooks.WebhookType_DEMISTO
	WebhookTypeAwsEventBridge        = webhooks.WebhookType_AWS_EVENT_BRIDGE
	WebhookTypeIbmEventNotifications = webhooks.WebhookType_IBM_EVENT_NOTIFICATIONS
)

// Webhook Type: Generic
type GenericWebhook = webhooks.OutgoingWebhookInputData_GenericWebhook

// Webhook Type: Slack
type Slack = webhooks.OutgoingWebhookInputData_Slack

// Webhook Type: PagerDuty
type PagerDuty = webhooks.OutgoingWebhookInputData_PagerDuty

// Webhook Type: SendLog
type SendLog = webhooks.OutgoingWebhookInputData_SendLog

// Webhook Type: EmailGroup
type EmailGroup = webhooks.OutgoingWebhookInputData_EmailGroup

// Webhook Type: MicrosoftTeams
type MicrosoftTeams = webhooks.OutgoingWebhookInputData_MicrosoftTeams

// Webhook Type: Jira
type Jira = webhooks.OutgoingWebhookInputData_Jira

// Webhook Type: Opsgenie
type Opsgenie = webhooks.OutgoingWebhookInputData_Opsgenie

// Webhook Type: Demisto
type Demisto = webhooks.OutgoingWebhookInputData_Demisto

// Webhook Type: AwsEventBridge
type AwsEventBridge = webhooks.OutgoingWebhookInputData_AwsEventBridge

// Webhook Type: IbmEventNotifications
type IbmEventNotifications = webhooks.OutgoingWebhookInputData_IbmEventNotifications

const (
	GenericWebhookConfigUnknown = webhooks.GenericWebhookConfig_UNKNOWN
	GenericWebhookConfigGet     = webhooks.GenericWebhookConfig_GET
	GenericWebhookConfigPost    = webhooks.GenericWebhookConfig_POST
	GenericWebhookConfigPut     = webhooks.GenericWebhookConfig_PUT
)

// GenericWebhookConfig is the configuration for a generic webhook.
type GenericWebhookConfig = webhooks.GenericWebhookConfig

const (
	// Slack attachment
	SlackConfigEmpty          = webhooks.SlackConfig_EMPTY
	SlackConfigMetricSnapshot = webhooks.SlackConfig_METRIC_SNAPSHOT
	SlackConfigLogs           = webhooks.SlackConfig_LOGS
	// Slack digests
	SlackConfigUnknown              = webhooks.SlackConfig_UNKNOWN
	SlackConfigErrorAndCriticalLogs = webhooks.SlackConfig_ERROR_AND_CRITICAL_LOGS
	SlackConfigFlowAnomalies        = webhooks.SlackConfig_FLOW_ANOMALIES
	SlackConfigSpikeAnomalies       = webhooks.SlackConfig_SPIKE_ANOMALIES
	SlackConfigDataUsage            = webhooks.SlackConfig_DATA_USAGE
)

// Slack config of the summary/digest it contains
type SlackConfigDigest = webhooks.SlackConfig_Digest

// Slack config of the attachment
type SlackConfigAttachment = webhooks.SlackConfig_Attachment

// Slack config
type SlackConfig = webhooks.SlackConfig

// PagerDuty config
type PagerDutyConfig = webhooks.PagerDutyConfig

// SendLog config
type SendLogConfig = webhooks.SendLogConfig

// EmailGroup config
type EmailGroupConfig = webhooks.EmailGroupConfig

// MicrosoftTeams config
type MicrosoftTeamsConfig = webhooks.MicrosoftTeamsConfig

// Jira config
type JiraConfig = webhooks.JiraConfig

// Opsgenie config
type OpsgenieConfig = webhooks.OpsgenieConfig

// Demisto config
type DemistoConfig = webhooks.DemistoConfig

// AwsEventBridge config
type AwsEventBridgeConfig = webhooks.AwsEventBridgeConfig

// IbmEventNotifications config
type IbmEventNotificationsConfig = webhooks.IbmEventNotificationsConfig

// Success response for testing a webhook
type TestOutgoingWebhookSuccess = webhooks.TestOutgoingWebhookResponse_Success

// Fail response for testing a webhook
type TestOutgoingWebhookFailure = webhooks.TestOutgoingWebhookResponse_Failure

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
