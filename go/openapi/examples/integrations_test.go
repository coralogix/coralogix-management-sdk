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

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	integrations "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/integration_service"
	webhooks "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/outgoing_webhooks_service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewIntegrationsClient(cpc)
	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		t.Fatalf("env AWS_REGION not set")
	}
	role := os.Getenv("AWS_TEST_ROLE")
	if role == "" {
		t.Fatalf("env AWS_TEST_ROLE not set")
	}

	name := "aws-metrics-collector"
	version := "0.1.0"

	params := []integrations.Parameter{
		integrations.ParameterStringValueAsParameter(&integrations.ParameterStringValue{
			Key:         strPtr("ApplicationName"),
			StringValue: strPtr("cxsdk"),
		}),
		integrations.ParameterStringValueAsParameter(&integrations.ParameterStringValue{
			Key:         strPtr("SubsystemName"),
			StringValue: strPtr("aws-metrics-collector"),
		}),
		integrations.ParameterStringValueAsParameter(&integrations.ParameterStringValue{
			Key:         strPtr("AwsRoleArn"),
			StringValue: strPtr(role),
		}),
		integrations.ParameterStringValueAsParameter(&integrations.ParameterStringValue{
			Key:         strPtr("IntegrationName"),
			StringValue: strPtr("sdk-integration-setup"),
		}),
		integrations.ParameterStringValueAsParameter(&integrations.ParameterStringValue{
			Key:         strPtr("AwsRegion"),
			StringValue: strPtr(awsRegion),
		}),
		integrations.ParameterBooleanValueAsParameter(&integrations.ParameterBooleanValue{
			Key:          strPtr("WithAggregations"),
			BooleanValue: boolPtr(false),
		}),
		integrations.ParameterBooleanValueAsParameter(&integrations.ParameterBooleanValue{
			Key:          strPtr("EnrichWithTags"),
			BooleanValue: boolPtr(false),
		}),
		integrations.ParameterStringListAsParameter(&integrations.ParameterStringList{
			Key: strPtr("MetricNamespaces"),
			StringList: &integrations.StringList{
				Values: []string{"AWS/S3"},
			},
		}),
	}

	testReq := integrations.TestIntegrationRequest{
		IntegrationData: integrations.IntegrationMetadata{
			IntegrationKey: strPtr(name),
			Version:        strPtr(version),
			IntegrationParameters: &integrations.GenericIntegrationParameters{
				Parameters: params,
			},
		},
	}

	_, httpResp, err := client.IntegrationServiceTestIntegration(context.Background()).
		TestIntegrationRequest(testReq).Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	saveReq := integrations.NewSaveIntegrationRequest(
		integrations.IntegrationMetadata{
			IntegrationKey: strPtr(name),
			Version:        strPtr(version),
			IntegrationParameters: &integrations.GenericIntegrationParameters{
				Parameters: params,
			},
		},
	)

	createResp, httpResp, err := client.IntegrationServiceSaveIntegration(context.Background()).
		SaveIntegrationRequest(*saveReq).Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		IntegrationServiceDeleteIntegration(context.Background(), createResp.IntegrationId).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestWebhooks(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewWebhooksClient(cpc)

	tests := []struct {
		name string
		data webhooks.OutgoingWebhookInputData
	}{
		{
			name: "slack-webhook",
			data: webhooks.OutgoingWebhookInputDataSlackAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataSlack{
					Name: "slack-webhook",
					Type: webhooks.WEBHOOKTYPE_SLACK,
					Url:  webhooks.PtrString("https://join.slack.com/example"),
					Slack: &webhooks.SlackConfig{
						Attachments: []webhooks.Attachment{
							{
								IsActive: boolPtr(true),
								Type:     webhooks.ATTACHMENTTYPE_LOGS,
							},
						},
						Digests: []webhooks.Digest{
							{
								IsActive: boolPtr(true),
								Type:     webhooks.DIGESTTYPE_FLOW_ANOMALIES,
							},
						},
					},
				},
			),
		},
		{
			name: "custom-webhook",
			data: webhooks.OutgoingWebhookInputDataGenericWebhookAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataGenericWebhook{
					Name: "custom-webhook",
					Type: webhooks.WEBHOOKTYPE_GENERIC,
					Url:  webhooks.PtrString("https://example-url.com"),
					GenericWebhook: &webhooks.GenericWebhookConfig{
						Method: webhooks.METHODTYPE_GET,
						Payload: webhooks.PtrString(
							"Hello from $ALERT_NAME, a coralogix alert",
						),
						Uuid: uuid.NewString(),
					},
				},
			),
		},
		{
			name: "pager-duty-webhook",
			data: webhooks.OutgoingWebhookInputDataPagerDutyAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataPagerDuty{
					Name: "pager-duty-webhook",
					Type: webhooks.WEBHOOKTYPE_PAGERDUTY,
					Url:  webhooks.PtrString("https://example-url.com/"),
					PagerDuty: &webhooks.PagerDutyConfig{
						ServiceKey: "example-key",
					},
				},
			),
		},
		{
			name: "email-group-webhook",
			data: webhooks.OutgoingWebhookInputDataEmailGroupAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataEmailGroup{
					Name: "email-group-webhook",
					Type: webhooks.WEBHOOKTYPE_EMAIL_GROUP,
					Url:  webhooks.PtrString("https://example-url.com/"),
					EmailGroup: &webhooks.EmailGroupConfig{
						EmailAddresses: []string{"user@example.com"}},
				},
			),
		},
		{
			name: "jira-webhook",
			data: webhooks.OutgoingWebhookInputDataJiraAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataJira{
					Name: "jira-webhook",
					Type: webhooks.WEBHOOKTYPE_JIRA,
					Url:  webhooks.PtrString("https://my-jira.atlassian.net/jira/"),
					Jira: &webhooks.JiraConfig{

						ProjectKey: "example-key",
						Email:      "email@coralogix.com",
						ApiToken:   "example-token",
					},
				},
			),
		},
		{
			name: "opsgenie-webhook",
			data: webhooks.OutgoingWebhookInputDataOpsgenieAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataOpsgenie{
					Name:     "opsgenie-webhook",
					Type:     webhooks.WEBHOOKTYPE_OPSGENIE,
					Url:      webhooks.PtrString("https://example.opsgenie.com"),
					Opsgenie: map[string]interface{}{},
				},
			),
		},
		{
			name: "demisto-webhook",
			data: webhooks.OutgoingWebhookInputDataDemistoAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataDemisto{
					Name: "demisto-webhook",
					Type: webhooks.WEBHOOKTYPE_DEMISTO,
					Url:  webhooks.PtrString("https://example.com"),
					Demisto: &webhooks.DemistoConfig{
						Uuid:    uuid.NewString(),
						Payload: "Hello from $ALERT_NAME, a coralogix alert",
					},
				},
			),
		},
		{
			name: "sendlog-webhook",
			data: webhooks.OutgoingWebhookInputDataSendLogAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataSendLog{
					Name: "sendlog-webhook",
					Type: webhooks.WEBHOOKTYPE_SEND_LOG,
					Url:  webhooks.PtrString("https://example.com"),
					SendLog: &webhooks.SendLogConfig{
						Uuid:    uuid.NewString(),
						Payload: "Hello from $ALERT_NAME, a coralogix alert",
					},
				},
			),
		},
		{
			name: "event-bridge-webhook",
			data: webhooks.OutgoingWebhookInputDataAwsEventBridgeAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataAwsEventBridge{
					Name: "event-bridge-webhook",
					Type: webhooks.WEBHOOKTYPE_AWS_EVENT_BRIDGE,
					Url:  webhooks.PtrString("https://aws.amazon.com"),
					AwsEventBridge: &webhooks.AwsEventBridgeConfig{
						EventBusArn: "arn:aws:events:us-east-1:123456789012:event-bus/default",
						Detail:      "example-detail",
						DetailType:  "example-detail-type",
						Source:      "example-source",
						RoleName:    "example-role",
					},
				},
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := *webhooks.NewCreateOutgoingWebhookRequest(tt.data)
			createResp, httpResp, err := client.
				OutgoingWebhooksServiceCreateOutgoingWebhook(context.Background()).
				CreateOutgoingWebhookRequest(req).
				Execute()
			assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
			assert.NotEmpty(t, createResp.GetId())

			id := createResp.GetId()
			newName := fmt.Sprintf("%s-updated", tt.name)

			updateReq := webhooks.UpdateOutgoingWebhookRequest{
				Id:   id,
				Data: tt.data,
			}
			// updateReq.Data will be reused, just update the name
			switch v := updateReq.Data.GetActualInstance().(type) {
			case *webhooks.OutgoingWebhookInputDataSlack:
				v.Name = newName
			case *webhooks.OutgoingWebhookInputDataGenericWebhook:
				v.Name = newName
			case *webhooks.OutgoingWebhookInputDataJira:
				v.Name = newName
			case *webhooks.OutgoingWebhookInputDataPagerDuty:
				v.Name = newName
			case *webhooks.OutgoingWebhookInputDataDemisto:
				v.Name = newName
			case *webhooks.OutgoingWebhookInputDataEmailGroup:
				v.Name = newName
			case *webhooks.OutgoingWebhookInputDataOpsgenie:
				v.Name = newName
			case *webhooks.OutgoingWebhookInputDataSendLog:
				v.Name = newName
			case *webhooks.OutgoingWebhookInputDataAwsEventBridge:
				v.Name = newName
			default:
				t.Fatalf("unknown webhook type: %T", v)
			}

			_, httpResp, err = client.
				OutgoingWebhooksServiceUpdateOutgoingWebhook(context.Background()).
				UpdateOutgoingWebhookRequest(updateReq).
				Execute()
			assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

			_, httpResp, err = client.
				OutgoingWebhooksServiceGetOutgoingWebhook(context.Background(), id).
				Execute()
			assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

			_, httpResp, err = client.
				OutgoingWebhooksServiceDeleteOutgoingWebhook(context.Background(), id).
				Execute()
			assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
		})
	}
}

func strPtr(s string) *string { return &s }

func boolPtr(b bool) *bool { return &b }
