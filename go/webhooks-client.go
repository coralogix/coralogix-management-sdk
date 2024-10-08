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

	webhooks "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/outgoing_webhooks/v1"
)

// OutgoingWebhook is a webhook.
type OutgoingWebhook = webhooks.OutgoingWebhook

// CreateOutgoingWebhookRequest is a request to create a new webhook.
type CreateOutgoingWebhookRequest = webhooks.CreateOutgoingWebhookRequest

// CreateOutgoingWebhookResponse is a response for creating a new webhook.
type CreateOutgoingWebhookResponse = webhooks.CreateOutgoingWebhookResponse

// GetOutgoingWebhookRequest is a request to get a webhook by its ID.
type GetOutgoingWebhookRequest = webhooks.GetOutgoingWebhookRequest

// GetOutgoingWebhookResponse is a response for getting a webhook.
type GetOutgoingWebhookResponse = webhooks.GetOutgoingWebhookResponse

// UpdateOutgoingWebhookRequest is a request to update a webhook.
type UpdateOutgoingWebhookRequest = webhooks.UpdateOutgoingWebhookRequest

// UpdateOutgoingWebhookResponse is a response for updating a webhook.
type UpdateOutgoingWebhookResponse = webhooks.UpdateOutgoingWebhookResponse

// DeleteOutgoingWebhookRequest is a request to delete a webhook.
type DeleteOutgoingWebhookRequest = webhooks.DeleteOutgoingWebhookRequest

// DeleteOutgoingWebhookResponse is a response for deleting a webhook.
type DeleteOutgoingWebhookResponse = webhooks.DeleteOutgoingWebhookResponse

// ListOutgoingWebhookTypesRequest is a request to list all webhook types.
type ListOutgoingWebhookTypesRequest = webhooks.ListOutgoingWebhookTypesRequest

// ListOutgoingWebhookTypesResponse is a response for listing all webhook types.
type ListOutgoingWebhookTypesResponse = webhooks.ListOutgoingWebhookTypesResponse

// ListAllOutgoingWebhooksRequest is a request to list all webhooks.
type ListAllOutgoingWebhooksRequest = webhooks.ListAllOutgoingWebhooksRequest

// ListAllOutgoingWebhooksResponse is a response for listing all webhooks.
type ListAllOutgoingWebhooksResponse = webhooks.ListAllOutgoingWebhooksResponse

// TestExistingOutgoingWebhookRequest is a request to test an existing webhook.
type TestExistingOutgoingWebhookRequest = webhooks.TestExistingOutgoingWebhookRequest

// TestOutgoingWebhookRequest is a request to test a webhook.
type TestOutgoingWebhookRequest = webhooks.TestOutgoingWebhookRequest

// OutgoingWebhookInputData is the wrapped data for the webhooks.
type OutgoingWebhookInputData = webhooks.OutgoingWebhookInputData

// WebhookType ... is the type of webhook
type WebhookType = webhooks.WebhookType

const (
	// WebhookTypeUnknown is an unknown webhook type.
	WebhookTypeUnknown = webhooks.WebhookType_UNKNOWN
	// WebhookTypeGeneric is a generic webhook type.
	WebhookTypeGeneric = webhooks.WebhookType_GENERIC
	// WebhookTypeSlack is a Slack webhook type.
	WebhookTypeSlack = webhooks.WebhookType_SLACK
	// WebhookTypePagerduty is a PagerDuty webhook type.
	WebhookTypePagerduty = webhooks.WebhookType_PAGERDUTY
	// WebhookTypeSendLog is a SendLog webhook type.
	WebhookTypeSendLog = webhooks.WebhookType_SEND_LOG
	// WebhookTypeEmailGroup is an EmailGroup webhook type.
	WebhookTypeEmailGroup = webhooks.WebhookType_EMAIL_GROUP
	// WebhookTypeMicrosoftTeams is a MicrosoftTeams webhook type.
	WebhookTypeMicrosoftTeams = webhooks.WebhookType_MICROSOFT_TEAMS
	// WebhookTypeJira is a Jira webhook type.
	WebhookTypeJira = webhooks.WebhookType_JIRA
	// WebhookTypeOpsgenie is an Opsgenie webhook type.
	WebhookTypeOpsgenie = webhooks.WebhookType_OPSGENIE
	// WebhookTypeDemisto is a Demisto webhook type.
	WebhookTypeDemisto = webhooks.WebhookType_DEMISTO
	// WebhookTypeAwsEventBridge is an AWS EventBridge webhook type.
	WebhookTypeAwsEventBridge = webhooks.WebhookType_AWS_EVENT_BRIDGE
	// WebhookTypeIbmEventNotifications is an IBM Event Notifications webhook type.
	WebhookTypeIbmEventNotifications = webhooks.WebhookType_IBM_EVENT_NOTIFICATIONS
)

// GenericWebhookInputData is a generic webhook input type.
type GenericWebhookInputData = webhooks.OutgoingWebhookInputData_GenericWebhook

// SlackWebhookInputData is a Slack webhook input type.
type SlackWebhookInputData = webhooks.OutgoingWebhookInputData_Slack

// PagerDutyWebhookInputData is a PagerDuty webhook input type.
type PagerDutyWebhookInputData = webhooks.OutgoingWebhookInputData_PagerDuty

// SendLogWebhookInputData is a SendLog webhook input type.
type SendLogWebhookInputData = webhooks.OutgoingWebhookInputData_SendLog

// EmailGroupWebhookInputData is an EmailGroup webhook input type.
type EmailGroupWebhookInputData = webhooks.OutgoingWebhookInputData_EmailGroup

// MicrosoftTeamsWebhookInputData is a MicrosoftTeams webhook input type.
type MicrosoftTeamsWebhookInputData = webhooks.OutgoingWebhookInputData_MicrosoftTeams

// JiraWebhookInputData is a Jira webhook input type.
type JiraWebhookInputData = webhooks.OutgoingWebhookInputData_Jira

// OpsgenieWebhookInputData is an Opsgenie webhook input type.
type OpsgenieWebhookInputData = webhooks.OutgoingWebhookInputData_Opsgenie

// DemistoWebhookInputData is a Demisto webhook input type.
type DemistoWebhookInputData = webhooks.OutgoingWebhookInputData_Demisto

// AwsEventBridgeWebhookInputData is an AWS EventBridge webhook input type.
type AwsEventBridgeWebhookInputData = webhooks.OutgoingWebhookInputData_AwsEventBridge

// IbmEventNotificationsWebhookInputData is an IBM Event Notifications webhook input type.
type IbmEventNotificationsWebhookInputData = webhooks.OutgoingWebhookInputData_IbmEventNotifications

const (
	// GenericWebhookConfigUnknown is an unknown config
	GenericWebhookConfigUnknown = webhooks.GenericWebhookConfig_UNKNOWN
	// GenericWebhookConfigGet is a GET request
	GenericWebhookConfigGet = webhooks.GenericWebhookConfig_GET
	// GenericWebhookConfigPost is a POST request
	GenericWebhookConfigPost = webhooks.GenericWebhookConfig_POST
	// GenericWebhookConfigPut is a PUT request
	GenericWebhookConfigPut = webhooks.GenericWebhookConfig_PUT
)

// GenericWebhookConfig is the configuration for a generic webhook.
type GenericWebhookConfig = webhooks.GenericWebhookConfig

// GenericWebhookConfigMethodType is the type of the method.
type GenericWebhookConfigMethodType = webhooks.GenericWebhookConfig_MethodType

const (
	// SlackConfigEmpty is an empty attachment
	SlackConfigEmpty = webhooks.SlackConfig_EMPTY
	// SlackConfigMetricSnapshot for metrics
	SlackConfigMetricSnapshot = webhooks.SlackConfig_METRIC_SNAPSHOT
	// SlackConfigLogs for logs
	SlackConfigLogs = webhooks.SlackConfig_LOGS
	// SlackConfigUnknown unknown digest
	SlackConfigUnknown = webhooks.SlackConfig_UNKNOWN
	// SlackConfigErrorAndCriticalLogs for error and critical logs
	SlackConfigErrorAndCriticalLogs = webhooks.SlackConfig_ERROR_AND_CRITICAL_LOGS
	// SlackConfigFlowAnomalies for flow anomalies
	SlackConfigFlowAnomalies = webhooks.SlackConfig_FLOW_ANOMALIES
	// SlackConfigSpikeAnomalies for spike anomalies
	SlackConfigSpikeAnomalies = webhooks.SlackConfig_SPIKE_ANOMALIES
	// SlackConfigDataUsage for data usage
	SlackConfigDataUsage = webhooks.SlackConfig_DATA_USAGE
)

// SlackConfigDigest is a config
type SlackConfigDigest = webhooks.SlackConfig_Digest

// SlackConfigDigestType is a config type
type SlackConfigDigestType = webhooks.SlackConfig_DigestType

// SlackConfigAttachment is a config
type SlackConfigAttachment = webhooks.SlackConfig_Attachment

// SlackConfigAttachmentType is a config type
type SlackConfigAttachmentType = webhooks.SlackConfig_AttachmentType

// SlackConfig is a config
type SlackConfig = webhooks.SlackConfig

// PagerDutyConfig configures the service
type PagerDutyConfig = webhooks.PagerDutyConfig

// SendLogConfig configures the service
type SendLogConfig = webhooks.SendLogConfig

// EmailGroupConfig configures the service
type EmailGroupConfig = webhooks.EmailGroupConfig

// MicrosoftTeamsConfig configures the service
type MicrosoftTeamsConfig = webhooks.MicrosoftTeamsConfig

// JiraConfig configures the service
type JiraConfig = webhooks.JiraConfig

// OpsgenieConfig configures the service
type OpsgenieConfig = webhooks.OpsgenieConfig

// DemistoConfig configures the service
type DemistoConfig = webhooks.DemistoConfig

// AwsEventBridgeConfig configures the service
type AwsEventBridgeConfig = webhooks.AwsEventBridgeConfig

// IbmEventNotificationsConfig configures the service
type IbmEventNotificationsConfig = webhooks.IbmEventNotificationsConfig

// TestOutgoingWebhookSuccess response for testing a webhook
type TestOutgoingWebhookSuccess = webhooks.TestOutgoingWebhookResponse_Success

// TestOutgoingWebhookFailure response for testing a webhook
type TestOutgoingWebhookFailure = webhooks.TestOutgoingWebhookResponse_Failure

// GenericWebhook is a generic webhook.
type GenericWebhook = webhooks.OutgoingWebhook_GenericWebhook

// SlackWebhook is a Slack webhook.
type SlackWebhook = webhooks.OutgoingWebhook_Slack

// PagerDutyWebhook is a PagerDuty webhook.
type PagerDutyWebhook = webhooks.OutgoingWebhook_PagerDuty

// SendLogWebhook is a SendLog webhook.
type SendLogWebhook = webhooks.OutgoingWebhook_SendLog

// EmailGroupWebhook is an EmailGroup webhook.
type EmailGroupWebhook = webhooks.OutgoingWebhook_EmailGroup

// MicrosoftTeamsWebhook is a MicrosoftTeams webhook.
type MicrosoftTeamsWebhook = webhooks.OutgoingWebhook_MicrosoftTeams

// JiraWebhook is a Jira webhook.
type JiraWebhook = webhooks.OutgoingWebhook_Jira

// OpsgenieWebhook is an Opsgenie webhook.
type OpsgenieWebhook = webhooks.OutgoingWebhook_Opsgenie

// DemistoWebhook is a Demisto webhook.
type DemistoWebhook = webhooks.OutgoingWebhook_Demisto

// AwsEventBridgeWebhook is an AWS EventBridge webhook.
type AwsEventBridgeWebhook = webhooks.OutgoingWebhook_AwsEventBridge

// WebhooksClient is a client for the Coralogix Webhooks API.
type WebhooksClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new webhook.
func (c WebhooksClient) Create(ctx context.Context, req *CreateOutgoingWebhookRequest) (*webhooks.CreateOutgoingWebhookResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := webhooks.NewOutgoingWebhooksServiceClient(conn)

	return client.CreateOutgoingWebhook(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get gets the specified webhook.
func (c WebhooksClient) Get(ctx context.Context, req *GetOutgoingWebhookRequest) (*webhooks.GetOutgoingWebhookResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := webhooks.NewOutgoingWebhooksServiceClient(conn)

	return client.GetOutgoingWebhook(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Update updates the specified webhook.
func (c WebhooksClient) Update(ctx context.Context, req *UpdateOutgoingWebhookRequest) (*webhooks.UpdateOutgoingWebhookResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := webhooks.NewOutgoingWebhooksServiceClient(conn)

	return client.UpdateOutgoingWebhook(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes the specified webhook.
func (c WebhooksClient) Delete(ctx context.Context, req *DeleteOutgoingWebhookRequest) (*webhooks.DeleteOutgoingWebhookResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := webhooks.NewOutgoingWebhooksServiceClient(conn)

	return client.DeleteOutgoingWebhook(callProperties.Ctx, req, callProperties.CallOptions...)
}

// ListTypes lists all webhook types.
func (c WebhooksClient) ListTypes(ctx context.Context, req *ListOutgoingWebhookTypesRequest) (*webhooks.ListOutgoingWebhookTypesResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := webhooks.NewOutgoingWebhooksServiceClient(conn)

	return client.ListOutgoingWebhookTypes(callProperties.Ctx, req, callProperties.CallOptions...)
}

// List lists all webhooks.
func (c WebhooksClient) List(ctx context.Context, req *ListAllOutgoingWebhooksRequest) (*webhooks.ListAllOutgoingWebhooksResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := webhooks.NewOutgoingWebhooksServiceClient(conn)

	return client.ListAllOutgoingWebhooks(callProperties.Ctx, req, callProperties.CallOptions...)
}

// TestByID tests an existing webhook.
func (c WebhooksClient) TestByID(ctx context.Context, req *TestExistingOutgoingWebhookRequest) (*webhooks.TestOutgoingWebhookResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := webhooks.NewOutgoingWebhooksServiceClient(conn)

	return client.TestExistingOutgoingWebhook(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Test tests the included webhook.
func (c WebhooksClient) Test(ctx context.Context, req *TestOutgoingWebhookRequest) (*webhooks.TestOutgoingWebhookResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
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
