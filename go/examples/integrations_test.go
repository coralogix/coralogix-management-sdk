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
	"log"
	"os"
	"testing"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func bp(k string, v bool) *cxsdk.IntegrationParameter {
	return &cxsdk.IntegrationParameter{
		Key: k,
		Value: &cxsdk.IntegrationParameterBooleanValue{
			BooleanValue: &wrapperspb.BoolValue{Value: v},
		},
	}
}

func sp(k string, v string) *cxsdk.IntegrationParameter {
	return &cxsdk.IntegrationParameter{
		Key: k,
		Value: &cxsdk.IntegrationParameterStringValue{
			StringValue: &wrapperspb.StringValue{Value: v},
		},
	}
}

func sl(k string, v []string) *cxsdk.IntegrationParameter {
	wrapped := make([]*wrapperspb.StringValue, len(v))
	for i, s := range v {
		wrapped[i] = &wrapperspb.StringValue{Value: s}
	}
	return &cxsdk.IntegrationParameter{
		Key: k,
		Value: &cxsdk.IntegrationParameterStringList{
			StringList: &cxsdk.IntegrationParameterStringListInner{
				Values: wrapped,
			},
		},
	}
}

func TestIntegration(t *testing.T) {
	awsRegion := os.Getenv("AWS_REGION")
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator, err := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	assertNilAndPrintError(t, err)

	c := cxsdk.NewIntegrationsClient(creator)
	defer creator.CloseConnection()
	role := os.Getenv("AWS_TEST_ROLE")

	name := "aws-metrics-collector"
	version := "0.1.0"
	params := []*cxsdk.IntegrationParameter{
		sp("ApplicationName", "cxsdk"),
		sp("SubsystemName", "aws-metrics-collector"),
		sl("MetricNamespaces", []string{"AWS/S3"}),
		sp("AwsRoleArn", role),
		sp("IntegrationName", "sdk-integration-setup"),
		sp("AwsRegion", awsRegion),
		bp("WithAggregations", false),
		bp("EnrichWithTags", true),
	}
	_, err = c.Test(context.Background(), &cxsdk.TestIntegrationRequest{
		IntegrationData: &cxsdk.IntegrationMetadata{
			IntegrationKey: &wrapperspb.StringValue{Value: name},
			Version:        &wrapperspb.StringValue{Value: version},
			SpecificData: &cxsdk.IntegrationMetadataIntegrationParameters{
				IntegrationParameters: &cxsdk.GenericIntegrationParameters{
					Parameters: params,
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	assertNilAndPrintError(t, err)

	createResponse, err := c.Create(context.Background(), &cxsdk.SaveIntegrationRequest{
		Metadata: &cxsdk.IntegrationMetadata{
			IntegrationKey: wrapperspb.String(name),
			Version:        wrapperspb.String(version),
			SpecificData: &cxsdk.IntegrationMetadataIntegrationParameters{
				IntegrationParameters: &cxsdk.GenericIntegrationParameters{
					Parameters: params,
				},
			},
		},
	})

	assertNilAndPrintError(t, err)

	details, listError := c.GetDetails(context.Background(), &cxsdk.GetIntegrationDetailsRequest{
		Id:                     wrapperspb.String(name),
		IncludeTestingRevision: wrapperspb.Bool(true),
	})
	if listError != nil {
		log.Fatal(listError.Error())
	}
	assertNilAndPrintError(t, listError)

	integrationDetail := details.IntegrationDetail.IntegrationTypeDetails.(*cxsdk.IntegrationDetailsDefault)
	found := false
	for _, d := range integrationDetail.Default.Registered {
		if d.Id.Value == createResponse.IntegrationId.Value {
			found = true
			break
		}
	}
	assert.True(t, found)

	deployedIntegration, getError := c.Get(context.Background(), &cxsdk.GetDeployedIntegrationRequest{
		IntegrationId: createResponse.IntegrationId,
	})

	if getError != nil {
		log.Fatal(getError.Error())
	}

	assertNilAndPrintError(t, getError)
	assert.NotNil(t, deployedIntegration)

	_, deletionError := c.Delete(context.Background(), &cxsdk.DeleteIntegrationRequest{
		IntegrationId: createResponse.IntegrationId,
	})

	assertNilAndPrintError(t, deletionError)

}

func TestWebhooks(t *testing.T) {
	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("slack-webhook"),
			Url:  wrapperspb.String("https://join.slack.com/example"),
			Type: cxsdk.WebhookTypeSlack,
			Config: &cxsdk.SlackWebhookInputData{
				Slack: &cxsdk.SlackConfig{
					Attachments: []*cxsdk.SlackConfigAttachment{
						{
							Type:     cxsdk.SlackConfigLogs,
							IsActive: wrapperspb.Bool(true),
						},
					},
					Digests: []*cxsdk.SlackConfigDigest{
						{
							Type:     cxsdk.SlackConfigFlowAnomalies,
							IsActive: wrapperspb.Bool(true),
						},
					},
				},
			},
		},
	})

	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("custom-webhook"),
			Url:  wrapperspb.String("https://example-url.com"),
			Type: cxsdk.WebhookTypeGeneric,
			Config: &cxsdk.GenericWebhookInputData{
				GenericWebhook: &cxsdk.GenericWebhookConfig{
					Uuid:    wrapperspb.String(uuid.NewString()),
					Method:  cxsdk.GenericWebhookConfigGet,
					Payload: wrapperspb.String("Hello from $ALERT_NAME, a coralogix alert"),
				},
			},
		},
	})

	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("pager-duty-webhook"),
			Url:  wrapperspb.String("https://example-url.com/"),
			Type: cxsdk.WebhookTypePagerduty,
			Config: &cxsdk.PagerDutyWebhookInputData{
				PagerDuty: &cxsdk.PagerDutyConfig{
					ServiceKey: wrapperspb.String("example-key"),
				},
			},
		},
	})

	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("email-group-webhook"),
			Url:  wrapperspb.String("https://example-url.com/"),
			Type: cxsdk.WebhookTypeEmailGroup,
			Config: &cxsdk.EmailGroupWebhookInputData{
				EmailGroup: &cxsdk.EmailGroupConfig{
					EmailAddresses: []*wrapperspb.StringValue{wrapperspb.String("user@example.com")},
				},
			},
		},
	})

	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("jira-webhook"),
			Url:  wrapperspb.String("https://my-jira.atlassian.net/jira/"),
			Type: cxsdk.WebhookTypeJira,
			Config: &cxsdk.JiraWebhookInputData{
				Jira: &cxsdk.JiraConfig{
					ProjectKey: wrapperspb.String("example-key"),
					Email:      wrapperspb.String("email@coralogix.com"),
					ApiToken:   wrapperspb.String("example-token"),
				},
			},
		},
	})

	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("opsgenie-webhook"),
			Url:  wrapperspb.String("https://example.opsgenie.com"),
			Type: cxsdk.WebhookTypeOpsgenie,
			Config: &cxsdk.OpsgenieWebhookInputData{
				Opsgenie: &cxsdk.OpsgenieConfig{},
			},
		},
	})
	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("demisto-webhook"),
			Url:  wrapperspb.String("https://example.com"),
			Type: cxsdk.WebhookTypeDemisto,
			Config: &cxsdk.DemistoWebhookInputData{
				Demisto: &cxsdk.DemistoConfig{
					Uuid:    wrapperspb.String(uuid.NewString()),
					Payload: wrapperspb.String("Hello from $ALERT_NAME, a coralogix alert"),
				},
			},
		},
	})

	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("sendlog-webhook"),
			Url:  wrapperspb.String("https://example.com"),
			Type: cxsdk.WebhookTypeSendLog,
			Config: &cxsdk.SendLogWebhookInputData{
				SendLog: &cxsdk.SendLogConfig{
					Uuid:    wrapperspb.String(uuid.NewString()),
					Payload: wrapperspb.String("Hello from $ALERT_NAME, a coralogix alert"),
				},
			},
		},
	})

	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("event-bridge-webhook"),
			Url:  wrapperspb.String("https://aws.amazon.com"),
			Type: cxsdk.WebhookTypeAwsEventBridge,
			Config: &cxsdk.AwsEventBridgeWebhookInputData{
				AwsEventBridge: &cxsdk.AwsEventBridgeConfig{
					EventBusArn: wrapperspb.String("arn:aws:events:us-east-1:123456789012:event-bus/default"),
					Detail:      wrapperspb.String("example-detail"),
					DetailType:  wrapperspb.String("example-detail-type"),
					Source:      wrapperspb.String("example-source"),
					RoleName:    wrapperspb.String("example-role"),
				},
			},
		},
	})

	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator, err := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	assertNilAndPrintError(t, err)

	c := cxsdk.NewWebhooksClient(creator)
	defer creator.CloseConnection()
	testResult, err := c.Test(context.Background(), &cxsdk.TestOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("custom-webhook"),
			Url:  wrapperspb.String("https://httpbun.org/status/200"),
			Type: cxsdk.WebhookTypeGeneric,
			Config: &cxsdk.GenericWebhookInputData{
				GenericWebhook: &cxsdk.GenericWebhookConfig{
					Uuid:    wrapperspb.String(uuid.NewString()),
					Method:  cxsdk.GenericWebhookConfigGet,
					Payload: nil,
				},
			},
		},
	})
	assertNilAndPrintError(t, err)
	if testResult.GetFailure() != nil {
		log.Fatal(testResult.GetFailure().ErrorMessage.Value)
	}
	assert.NotNil(t, testResult.GetSuccess())
	assert.Nil(t, testResult.GetFailure())
}

func crud(t *testing.T, req *cxsdk.CreateOutgoingWebhookRequest) {

	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator, err := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	assertNilAndPrintError(t, err)
	c := cxsdk.NewWebhooksClient(creator)
	defer creator.CloseConnection()
	result, e := c.Create(context.Background(), req)
	if e != nil {
		log.Fatal(e.Error())
	}
	assertNilAndPrintError(t, e)
	newName := fmt.Sprintf("%v-2", req.Data.Name.GetValue())
	_, e = c.Update(context.Background(),
		&cxsdk.UpdateOutgoingWebhookRequest{
			Id: result.Id.GetValue(),
			Data: &cxsdk.OutgoingWebhookInputData{
				Name:   wrapperspb.String(newName),
				Url:    req.Data.Url,
				Type:   req.Data.Type,
				Config: req.Data.Config,
			},
		})

	assertNilAndPrintError(t, e)
	_, e = c.Get(context.Background(), &cxsdk.GetOutgoingWebhookRequest{
		Id: result.Id,
	})
	assertNilAndPrintError(t, e)
	_, e = c.Delete(context.Background(), &cxsdk.DeleteOutgoingWebhookRequest{
		Id: result.Id,
	})
	assertNilAndPrintError(t, e)
}

func TestContextualDataIntegrations(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator, err := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	assertNilAndPrintError(t, err)
	c := cxsdk.NewContextualDataIntegrationsClient(creator)
	defer creator.CloseConnection()

	// Get all integrations first to check initial state
	allIntegrations, err := c.List(context.Background(), &cxsdk.GetContextualDataIntegrationsRequest{})
	assertNilAndPrintError(t, err)
	initialCount := len(allIntegrations.Integrations)

	// Create a new integration
	createResponse, err := c.Create(context.Background(), &cxsdk.SaveContextualDataIntegrationRequest{
		Metadata: &cxsdk.IntegrationMetadata{
			IntegrationKey: wrapperspb.String("AWS-Status-Logs"),
			Version:        wrapperspb.String("0.0.1"),
			SpecificData: &cxsdk.IntegrationMetadataIntegrationParameters{
				IntegrationParameters: &cxsdk.GenericIntegrationParameters{
					Parameters: []*cxsdk.IntegrationParameter{},
				},
			},
		},
	})
	assertNilAndPrintError(t, err)

	// Get integration definition
	_, err = c.GetDefinition(context.Background(), &cxsdk.GetContextualDataIntegrationDefinitionRequest{
		Id: &wrapperspb.StringValue{Value: "AWS-Status-Logs"},
	})
	assertNilAndPrintError(t, err)

	// Delete the integration
	_, err = c.Delete(context.Background(), &cxsdk.DeleteContextualDataIntegrationRequest{
		IntegrationId: createResponse.IntegrationId,
	})
	assertNilAndPrintError(t, err)

	// Verify deletion by getting all integrations again
	finalIntegrations, err := c.List(context.Background(), &cxsdk.GetContextualDataIntegrationsRequest{})
	assertNilAndPrintError(t, err)
	assert.Equal(t, initialCount, len(finalIntegrations.Integrations))
}
