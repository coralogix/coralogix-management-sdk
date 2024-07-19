package examples

import (
	"context"
	cxsdk "coralogix-management-sdk/go"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestWebhooks(t *testing.T) {
	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("slack-webhook"),
			Url:  wrapperspb.String("https://join.slack.com/example"),
			Type: cxsdk.WebhookType_SLACK,
			Config: &cxsdk.Slack{
				Slack: &cxsdk.SlackConfig{
					Attachments: []*cxsdk.SlackConfig_Attachment{
						{
							Type:     cxsdk.SlackConfig_LOGS,
							IsActive: wrapperspb.Bool(true),
						},
					},
					Digests: []*cxsdk.SlackConfig_Digest{
						{
							Type:     cxsdk.SlackConfig_FLOW_ANOMALIES,
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
			Type: cxsdk.WebhookType_GENERIC,
			Config: &cxsdk.GenericWebhook{
				GenericWebhook: &cxsdk.GenericWebhookConfig{
					Uuid:    wrapperspb.String(uuid.NewString()),
					Method:  cxsdk.GenericWebhookConfig_GET,
					Payload: wrapperspb.String("Hello from $ALERT_NAME, a coralogix alert"),
				},
			},
		},
	})

	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("pager-duty-webhook"),
			Url:  wrapperspb.String("https://example-url.com/"),
			Type: cxsdk.WebhookType_PAGERDUTY,
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
			Type: cxsdk.WebhookType_EMAIL_GROUP,
			Config: &cxsdk.EmailGroup{
				EmailGroup: &cxsdk.EmailGroupConfig{
					EmailAddresses: []*wrapperspb.StringValue{wrapperspb.String("user@example.com")},
				},
			},
		},
	})

	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("microsoft-teams-webhook"),
			Url:  wrapperspb.String("https://teams.microsoft.com/"),
			Type: cxsdk.WebhookType_MICROSOFT_TEAMS,
			Config: &cxsdk.MicrosoftTeams{
				MicrosoftTeams: &cxsdk.MicrosoftTeamsConfig{},
			},
		},
	})

	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("jira-webhook"),
			Url:  wrapperspb.String("https://my-jira.atlassian.net/jira/"),
			Type: cxsdk.WebhookType_JIRA,
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
			Type: cxsdk.WebhookType_OPSGENIE,
			Config: &cxsdk.Opsgenie{
				Opsgenie: &cxsdk.OpsgenieConfig{},
			},
		},
	})
	crud(t, &cxsdk.CreateOutgoingWebhookRequest{
		Data: &cxsdk.OutgoingWebhookInputData{
			Name: wrapperspb.String("demisto-webhook"),
			Url:  wrapperspb.String("https://example.com"),
			Type: cxsdk.WebhookType_DEMISTO,
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
			Type: cxsdk.WebhookType_SEND_LOG,
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
			Type: cxsdk.WebhookType_AWS_EVENT_BRIDGE,
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
}

func crud(t *testing.T, req *cxsdk.CreateOutgoingWebhookRequest) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, apiKey)
	c := cxsdk.NewWebhooksClient(creator)
	result, e := c.Create(context.Background(), req)
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
}
