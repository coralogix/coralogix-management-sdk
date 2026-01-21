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
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
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
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewWebhooksClient(cfg)

	_, httpResp, err := client.OutgoingWebhooksServiceListAllOutgoingWebhooks(context.Background()).Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	tests := []struct {
		name string
		data webhooks.OutgoingWebhookInputData
	}{
		{
			name: "slack-webhook",
			data: webhooks.OutgoingWebhookInputDataSlackAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataSlack{
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
			),
		},
		{
			name: "custom-webhook",
			data: webhooks.OutgoingWebhookInputDataGenericWebhookAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataGenericWebhook{
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
			),
		},
		{
			name: "pager-duty-webhook",
			data: webhooks.OutgoingWebhookInputDataPagerDutyAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataPagerDuty{
					Name: webhooks.PtrString("pager-duty-webhook"),
					Type: webhooks.WEBHOOKTYPE_PAGERDUTY.Ptr(),
					Url:  webhooks.PtrString("https://api.staging.coralogix.net/mgmt/testing/tools/httpbin/post/"),
					PagerDuty: &webhooks.PagerDutyConfig{
						ServiceKey: webhooks.PtrString("example-key"),
					},
				},
			),
		},
		{
			name: "email-group-webhook",
			data: webhooks.OutgoingWebhookInputDataEmailGroupAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataEmailGroup{
					Name: webhooks.PtrString("email-group-webhook"),
					Type: webhooks.WEBHOOKTYPE_EMAIL_GROUP.Ptr(),
					Url:  webhooks.PtrString("https://api.staging.coralogix.net/mgmt/testing/tools/httpbin/post/"),
					EmailGroup: &webhooks.EmailGroupConfig{
						EmailAddresses: []string{"user@example.com"}},
				},
			),
		},
		{
			name: "jira-webhook",
			data: webhooks.OutgoingWebhookInputDataJiraAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataJira{
					Name: webhooks.PtrString("jira-webhook"),
					Type: webhooks.WEBHOOKTYPE_JIRA.Ptr(),
					Url:  webhooks.PtrString("https://my-jira.atlassian.net/jira/"),
					Jira: &webhooks.JiraConfig{

						ProjectKey: webhooks.PtrString("example-key"),
						Email:      webhooks.PtrString("email@coralogix.com"),
						ApiToken:   webhooks.PtrString("example-token"),
					},
				},
			),
		},
		{
			name: "opsgenie-webhook",
			data: webhooks.OutgoingWebhookInputDataOpsgenieAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataOpsgenie{
					Name:     webhooks.PtrString("opsgenie-webhook"),
					Type:     webhooks.WEBHOOKTYPE_OPSGENIE.Ptr(),
					Url:      webhooks.PtrString("https://api.opsgenie.com"),
					Opsgenie: map[string]interface{}{},
				},
			),
		},
		{
			name: "demisto-webhook",
			data: webhooks.OutgoingWebhookInputDataDemistoAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataDemisto{
					Name: webhooks.PtrString("demisto-webhook"),
					Type: webhooks.WEBHOOKTYPE_DEMISTO.Ptr(),
					Url:  webhooks.PtrString("https://example.com"),
					Demisto: &webhooks.DemistoConfig{
						Uuid:    webhooks.PtrString(uuid.NewString()),
						Payload: webhooks.PtrString("Hello from $ALERT_NAME, a coralogix alert"),
					},
				},
			),
		},
		{
			name: "sendlog-webhook",
			data: webhooks.OutgoingWebhookInputDataSendLogAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataSendLog{
					Name: webhooks.PtrString("sendlog-webhook"),
					Type: webhooks.WEBHOOKTYPE_SEND_LOG.Ptr(),
					Url:  webhooks.PtrString("https://example.com"),
					SendLog: &webhooks.SendLogConfig{
						Uuid:    webhooks.PtrString(uuid.NewString()),
						Payload: webhooks.PtrString("Hello from $ALERT_NAME, a coralogix alert"),
					},
				},
			),
		},
		{
			name: "event-bridge-webhook",
			data: webhooks.OutgoingWebhookInputDataAwsEventBridgeAsOutgoingWebhookInputData(
				&webhooks.OutgoingWebhookInputDataAwsEventBridge{
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
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
			// updateReq.Data will be reused, just update the name
			switch v := updateReq.Data.GetActualInstance().(type) {
			case *webhooks.OutgoingWebhookInputDataSlack:
				v.Name = &newName
			case *webhooks.OutgoingWebhookInputDataGenericWebhook:
				v.Name = &newName
			case *webhooks.OutgoingWebhookInputDataJira:
				v.Name = &newName
			case *webhooks.OutgoingWebhookInputDataPagerDuty:
				v.Name = &newName
			case *webhooks.OutgoingWebhookInputDataDemisto:
				v.Name = &newName
			case *webhooks.OutgoingWebhookInputDataEmailGroup:
				v.Name = &newName
			case *webhooks.OutgoingWebhookInputDataOpsgenie:
				v.Name = &newName
			case *webhooks.OutgoingWebhookInputDataSendLog:
				v.Name = &newName
			case *webhooks.OutgoingWebhookInputDataAwsEventBridge:
				v.Name = &newName
			default:
				t.Fatalf("unknown webhook type: %T", v)
			}

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
