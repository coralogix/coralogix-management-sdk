#[cfg(test)]
mod tests {
    use cx_sdk::{
        auth::ApiKey,
        client::webhooks::{
            generic_webhook_config::{self, *},
            slack_config, Config, DigestType, GenericWebhookConfig, SlackConfig, WebhookType,
            WebhooksClient,
        },
        CoralogixRegion,
    };
    use uuid::Uuid;

    #[tokio::test]
    async fn test_outgoing_webhooks_client() {
        let client = WebhooksClient::new(
            ApiKey::from_env().unwrap(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();
        let _ = client
            .create(
                WebhookType::Slack,
                Some("slack-webhook".into()),
                "https://join.slack.com/example".parse().unwrap(),
                Config::Slack(SlackConfig {
                    digests: vec![slack_config::Digest {
                        r#type: slack_config::DigestType::FlowAnomalies.into(),
                        is_active: Some(true),
                    }],
                    attachments: vec![slack_config::Attachment {
                        r#type: slack_config::AttachmentType::MetricSnapshot.into(),
                        is_active: Some(true),
                    }],
                }),
            )
            .await
            .unwrap();
        let _ = client
            .create(
                WebhookType::Generic,
                Some("custom-webhook".into()),
                "https://example-url.com/".parse().unwrap(),
                Config::GenericWebhook(GenericWebhookConfig {
                    uuid: None,
                    method: generic_webhook_config::MethodType::Post.into(),
                    headers: Default::default(),
                    payload: Some("Hello from $ALERT_NAME, a coralogix alert".into()), // https://coralogix.com/docs/alert-webhooks/#placeholders
                }),
            )
            .await
            .unwrap();
        let _ = client.delete(id.clone()).await.unwrap();
    }
}

//   data "coralogix_webhook" "imported_webhook" {
//     id = coralogix_webhook.slack_webhook.id
//   }

//   resource "coralogix_webhook" "custom_webhook" {
//     name   = "custom-webhook"
//     custom = {
//       method  = "post"
//       headers = { "Content-Type" : "application/json" }
//       url     = "https://example-url.com/"
//     }
//   }

//   resource "coralogix_webhook" "pager_duty_webhook" {
//     name       = "pagerduty-webhook"
//     pager_duty = {
//       service_key = "service-key"
//     }
//   }

//   resource "coralogix_webhook" "email_group_webhook" {
//     name        = "email-group-webhook"
//     email_group = {
//       emails = ["user@example.com"]
//     }
//   }

//   resource "coralogix_webhook" "microsoft_teams_webhook" {
//     name            = "microsoft-teams-webhook"
//     microsoft_teams = {
//       url = "https://example-url.com/"
//     }
//   }

//   resource "coralogix_webhook" "jira_webhook" {
//     name = "jira-webhook"
//     jira = {
//       api_token   = "api-token"
//       email       = "example@coralogix.com"
//       project_key = "project-key"
//       url         = "https://coralogix.atlassian.net/jira/your-work"
//     }
//   }

//   resource "coralogix_webhook" "opsgenie_webhook" {
//     name     = "opsgenie-webhook"
//     opsgenie = {
//       url = "https://example-url.com/"
//     }
//   }

//   resource "coralogix_webhook" "demisto_webhook" {
//     name    = "demisto-webhook"
//     demisto = {
//       url = "https://example-url.com/"
//     }
//   }

//   resource "coralogix_webhook" "sendlog_webhook" {
//     name    = "sendlog-webhook"
//     sendlog = {
//       url = "https://example-url.com/"
//     }
//   }

//   resource "coralogix_webhook" "event_bridge_webhook" {
//     name         = "event_bridge_webhook"
//     event_bridge = {
//       event_bus_arn = "arn:aws:events:us-east-1:123456789012:event-bus/default"
//       detail        = "example_detail"
//       detail_type   = "example_detail_type"
//       source        = "example_source"
//       role_name     = "example_role_name"
//     }
//   }

//   //example of how to use webhooks that was created via terraform
//   resource "coralogix_alert" "standard_alert" {
//     name        = "Standard alert example"
//     description = "Example of standard alert from terraform"
//     severity    = "Critical"

//     notifications_group {
//       notification {
//         integration_id              = coralogix_webhook.slack_webhook.external_id
//         retriggering_period_minutes = 60
//         notify_on = "Triggered_only"
//       }
//     }

//     standard {
//       search_query = "remote_addr_enriched:/.*/"
//       condition {
//         immediately = true
//       }
//     }
//   }
