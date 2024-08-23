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
	"testing"

	cxsdk "github.com/coralogix/coralogix-management-sdk"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestWebhooks(t *testing.T) {
	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("slack-webhook"),
			Url:  wrapperspb.String("https://join.slack.com/example"),
			Type: cxsdk.WebhookTypeSlack,
			Config: &cxsdk.Slack{
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
			Config: &cxsdk.GenericWebhook{
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
			Config: &cxsdk.PagerDuty{
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
			Config: &cxsdk.EmailGroup{
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
			Config: &cxsdk.Jira{
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
			Config: &cxsdk.Opsgenie{
				Opsgenie: &cxsdk.OpsgenieConfig{},
			},
		},
	})
	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("demisto-webhook"),
			Url:  wrapperspb.String("https://example.com"),
			Type: cxsdk.WebhookTypeDemisto,
			Config: &cxsdk.Demisto{
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
			Config: &cxsdk.SendLog{
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
			Config: &cxsdk.AwsEventBridge{
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
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)

	c := cxsdk.NewWebhooksClient(creator)
	testResult, err := c.Test(context.Background(), &cxsdk.TestOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("custom-webhook"),
			Url:  wrapperspb.String("https://httpbin.org/status/200"),
			Type: cxsdk.WebhookTypeGeneric,
			Config: &cxsdk.GenericWebhook{
				GenericWebhook: &cxsdk.GenericWebhookConfig{
					Uuid:    wrapperspb.String(uuid.NewString()),
					Method:  cxsdk.GenericWebhookConfigGet,
					Payload: nil,
				},
			},
		},
	})
	assert.Nil(t, err)
	if testResult.GetFailure() != nil {
		log.Fatal(testResult.GetFailure().ErrorMessage.Value)
	}
	assert.NotNil(t, testResult.GetSuccess())
	assert.Nil(t, testResult.GetFailure())

}

func crud(t *testing.T, req *cxsdk.CreateOutgoingWebhookRequest) {

	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewWebhooksClient(creator)
	result, e := c.Create(context.Background(), req)
	if e != nil {
		log.Fatal(e.Error())
	}
	assert.Nil(t, e)
	newName := fmt.Sprintf("%v-2", req.Data.Name.GetValue())
	_, e = c.Replace(context.Background(),
		&cxsdk.UpdateOutgoingWebhookRequest{
			Id: result.Id.GetValue(),
			Data: &cxsdk.OutgoingWebhookInputData{
				Name:   wrapperspb.String(newName),
				Url:    req.Data.Url,
				Type:   req.Data.Type,
				Config: req.Data.Config,
			},
		})

	assert.Nil(t, e)
	_, e = c.Get(context.Background(), &cxsdk.GetOutgoingWebhookRequest{
		Id: result.Id,
	})
	assert.Nil(t, e)
	_, e = c.Delete(context.Background(), &cxsdk.DeleteOutgoingWebhookRequest{
		Id: result.Id,
	})
	assert.Nil(t, e)
}
