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

package examples

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	integrations "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/integration_service"
	webhooks "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/outgoing_webhooks_service"
)

func TestIntegration(t *testing.T) {
	cfg := newTestConfig()
	client := cxsdk.NewIntegrationsClient(cfg)

	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		t.Fatalf("AWS_REGION environment variable is not set")
	}
	role := os.Getenv("AWS_TEST_ROLE")
	if role == "" {
		t.Fatalf("AWS_TEST_ROLE environment variable is not set")
	}

	name := "aws-metrics-collector"
	version := "0.1.0"

	params := []integrations.Parameter{
		{
			Key:         strPtr("ApplicationName"),
			StringValue: strPtr("cxsdk"),
		},
		{
			Key:         strPtr("SubsystemName"),
			StringValue: strPtr("aws-metrics-collector"),
		},
		{
			Key:         strPtr("AwsRoleArn"),
			StringValue: &role,
		},
		{
			Key:         strPtr("IntegrationName"),
			StringValue: strPtr("sdk-integration-setup"),
		},
		{
			Key:         strPtr("AwsRegion"),
			StringValue: &awsRegion,
		},
		{
			Key:          strPtr("WithAggregations"),
			BooleanValue: boolPtr(false),
		},
		{
			Key:          strPtr("EnrichWithTags"),
			BooleanValue: boolPtr(false),
		},
		{
			Key: strPtr("MetricNamespaces"),
			StringList: &integrations.StringList{
				Values: []string{"AWS/S3"},
			},
		},
	}

	testReq := integrations.TestIntegrationRequest{
		IntegrationData: &integrations.IntegrationMetadata{
			IntegrationKey: strPtr(name),
			Version:        strPtr(version),
			IntegrationParameters: &integrations.GenericIntegrationParameters{
				Parameters: params,
			},
		},
	}

	_, httpResp, err := client.IntegrationServiceTestIntegration(context.Background()).
		TestIntegrationRequest(testReq).Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	saveReq := integrations.SaveIntegrationRequest{
		Metadata: &integrations.IntegrationMetadata{
			IntegrationKey: strPtr(name),
			Version:        strPtr(version),
			IntegrationParameters: &integrations.GenericIntegrationParameters{
				Parameters: params,
			},
		},
	}

	createResp, httpResp, err := client.IntegrationServiceSaveIntegration(context.Background()).
		SaveIntegrationRequest(saveReq).Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		IntegrationServiceDeleteIntegration(context.Background(), *createResp.IntegrationId).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestWebhooks(t *testing.T) {
	cfg := newTestConfig()
	client := cxsdk.NewWebhooksClient(cfg)

	_, httpResp, err := client.OutgoingWebhooksServiceListAllOutgoingWebhooks(context.Background()).Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	tests := []struct {
		name string
		data webhooks.OutgoingWebhookInputData
	}{
		{
			name: "slack-webhook",
			data: webhooks.OutgoingWebhookInputData{
				Name: webhooks.PtrString("slack-webhook"),
				Type: webhooks.WEBHOOKTYPE_SLACK.Ptr(),
				Url:  webhooks.PtrString("https://join.slack.com/example"),
				Slack: &webhooks.SlackConfig{
					Attachments: []webhooks.Attachment{
						{
							IsActive: boolPtr(true),
							Type:     webhooks.ATTACHMENTTYPE_LOGS.Ptr(),
						},
					},
					Digests: []webhooks.Digest{
						{
							IsActive: boolPtr(true),
							Type:     webhooks.DIGESTTYPE_FLOW_ANOMALIES.Ptr(),
						},
					},
				},
			},
		},
		{
			name: "custom-webhook",
			data: webhooks.OutgoingWebhookInputData{
				Name: webhooks.PtrString("custom-webhook"),
				Type: webhooks.WEBHOOKTYPE_GENERIC.Ptr(),
				Url:  webhooks.PtrString("https://example-url.com"),
				GenericWebhook: &webhooks.GenericWebhookConfig{
					Method: webhooks.METHODTYPE_GET.Ptr(),
					Payload: webhooks.PtrString(
						"Hello from $ALERT_NAME, a coralogix alert",
					),
					Uuid: webhooks.PtrString(uuid.NewString()),
				},
			},
		},
		{
			name: "pager-duty-webhook",
			data: webhooks.OutgoingWebhookInputData{
				Name: webhooks.PtrString("pager-duty-webhook"),
				Type: webhooks.WEBHOOKTYPE_PAGERDUTY.Ptr(),
				Url:  webhooks.PtrString("https://example-url.com/"),
				PagerDuty: &webhooks.PagerDutyConfig{
					ServiceKey: webhooks.PtrString("example-key"),
				},
			},
		},
		{
			name: "email-group-webhook",
			data: webhooks.OutgoingWebhookInputData{
				Name: webhooks.PtrString("email-group-webhook"),
				Type: webhooks.WEBHOOKTYPE_EMAIL_GROUP.Ptr(),
				Url:  webhooks.PtrString("https://example-url.com/"),
				EmailGroup: &webhooks.EmailGroupConfig{
					EmailAddresses: []string{"user@example.com"}},
			},
		},
		{
			name: "jira-webhook",
			data: webhooks.OutgoingWebhookInputData{
				Name: webhooks.PtrString("jira-webhook"),
				Type: webhooks.WEBHOOKTYPE_JIRA.Ptr(),
				Url:  webhooks.PtrString("https://my-jira.atlassian.net/jira/"),
				Jira: &webhooks.JiraConfig{

					ProjectKey: webhooks.PtrString("example-key"),
					Email:      webhooks.PtrString("email@coralogix.com"),
					ApiToken:   webhooks.PtrString("example-token"),
				},
			},
		},
		{
			name: "opsgenie-webhook",
			data: webhooks.OutgoingWebhookInputData{
				Name:     webhooks.PtrString("opsgenie-webhook"),
				Type:     webhooks.WEBHOOKTYPE_OPSGENIE.Ptr(),
				Url:      webhooks.PtrString("https://example.opsgenie.com"),
				Opsgenie: map[string]interface{}{},
			},
		},
		{
			name: "demisto-webhook",
			data: webhooks.OutgoingWebhookInputData{
				Name: webhooks.PtrString("demisto-webhook"),
				Type: webhooks.WEBHOOKTYPE_DEMISTO.Ptr(),
				Url:  webhooks.PtrString("https://example.com"),
				Demisto: &webhooks.DemistoConfig{
					Uuid:    webhooks.PtrString(uuid.NewString()),
					Payload: webhooks.PtrString("Hello from $ALERT_NAME, a coralogix alert"),
				},
			},
		},
		{
			name: "sendlog-webhook",
			data: webhooks.OutgoingWebhookInputData{
				Name: webhooks.PtrString("sendlog-webhook"),
				Type: webhooks.WEBHOOKTYPE_SEND_LOG.Ptr(),
				SendLog: &webhooks.SendLogConfig{
					Uuid:    webhooks.PtrString(uuid.NewString()),
					Payload: webhooks.PtrString("Hello from $ALERT_NAME, a coralogix alert"),
				},
			},
		},
		{
			name: "event-bridge-webhook",
			data: webhooks.OutgoingWebhookInputData{
				Name: webhooks.PtrString("event-bridge-webhook"),
				Type: webhooks.WEBHOOKTYPE_AWS_EVENT_BRIDGE.Ptr(),
				Url:  webhooks.PtrString("https://aws.amazon.com"),
				AwsEventBridge: &webhooks.AwsEventBridgeConfig{
					EventBusArn: webhooks.PtrString("arn:aws:events:us-east-1:123456789012:event-bus/default"),
					Detail:      webhooks.PtrString("example-detail"),
					DetailType:  webhooks.PtrString("example-detail-type"),
					Source:      webhooks.PtrString("example-source"),
					RoleName:    webhooks.PtrString("example-role"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "opsgenie-webhook" {
				t.Skip("Skipping opsgenie-webhook test: example.opsgenie.com host cannot be resolved")
			}
			req := webhooks.CreateOutgoingWebhookRequest{
				Data: &tt.data,
			}
			createResp, httpResp, err := client.
				OutgoingWebhooksServiceCreateOutgoingWebhook(context.Background()).
				CreateOutgoingWebhookRequest(req).
				Execute()
			require.NoError(t, cxsdk.NewAPIError(httpResp, err))
			require.NotEmpty(t, createResp.GetId())

			id := createResp.GetId()
			newName := fmt.Sprintf("%s-updated", tt.name)

			updateReq := webhooks.UpdateOutgoingWebhookRequest{
				Id:   &id,
				Data: &tt.data,
			}
			updateReq.Data.Name = &newName

			_, httpResp, err = client.
				OutgoingWebhooksServiceUpdateOutgoingWebhook(context.Background()).
				UpdateOutgoingWebhookRequest(updateReq).
				Execute()
			require.NoError(t, cxsdk.NewAPIError(httpResp, err))

			_, httpResp, err = client.
				OutgoingWebhooksServiceGetOutgoingWebhook(context.Background(), id).
				Execute()
			require.NoError(t, cxsdk.NewAPIError(httpResp, err))

			_, httpResp, err = client.
				OutgoingWebhooksServiceDeleteOutgoingWebhook(context.Background(), id).
				Execute()
			require.NoError(t, cxsdk.NewAPIError(httpResp, err))
		})
	}
}

func strPtr(s string) *string { return &s }

func boolPtr(b bool) *bool { return &b }
