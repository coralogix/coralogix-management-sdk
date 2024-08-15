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
#[cfg(test)]
mod tests {
    use cx_sdk::{
        auth::AuthContext,
        client::webhooks::{
            self, generic_webhook_config, slack_config, AwsEventBridgeConfig, Config,
            DemistoConfig, EmailGroupConfig, GenericWebhookConfig, JiraConfig, OpsgenieConfig,
            PagerDutyConfig, SendLogConfig, SlackConfig, WebhookType, WebhooksClient,
        },
        CoralogixRegion,
    };
    use url::Url;
    use uuid::Uuid;

    #[tokio::test]
    async fn test_outgoing_webhooks_client() {
        crud(
            "slack-webhook".into(),
            WebhookType::Slack,
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
        .await;
        crud(
            "custom-webhook".into(),
            WebhookType::Generic,
            "https://example-url.com/".parse().unwrap(),
            Config::GenericWebhook(GenericWebhookConfig {
                uuid: Some(Uuid::new_v4().to_string()),
                method: generic_webhook_config::MethodType::Get.into(),
                headers: Default::default(),
                payload: Some("Hello from $ALERT_NAME, a coralogix alert".into()), // https://coralogix.com/docs/alert-webhooks/#placeholders
            }),
        )
        .await;

        crud(
            "pager-duty-webhook".into(),
            WebhookType::Pagerduty,
            "https://example-url.com/".parse().unwrap(),
            Config::PagerDuty(PagerDutyConfig {
                service_key: Some("service-key".into()),
            }),
        )
        .await;

        crud(
            "email-group-webhook".into(),
            WebhookType::EmailGroup,
            "https://example-url.com/".parse().unwrap(),
            Config::EmailGroup(EmailGroupConfig {
                email_addresses: vec!["user@example.com".into()],
            }),
        )
        .await;

        crud(
            "jira-webhook".into(),
            WebhookType::Jira,
            "https://my-jira.atlassian.net/jira/".parse().unwrap(),
            Config::Jira(JiraConfig {
                api_token: Some("token-token".into()),
                email: Some("email@coralogix.com".into()),
                project_key: Some("project-key".into()),
            }),
        )
        .await;

        crud(
            "opsgenie-webhook".into(),
            WebhookType::Opsgenie,
            "https://example.opsgenie.com/".parse().unwrap(),
            Config::Opsgenie(OpsgenieConfig {}),
        )
        .await;

        crud(
            "demisto-webhook".into(),
            WebhookType::Demisto,
            "https://example.com/".parse().unwrap(),
            Config::Demisto(DemistoConfig {
                uuid: Some(Uuid::new_v4().to_string()),
                payload: Some("Hello from $ALERT_NAME, a coralogix alert".into()),
            }),
        )
        .await;

        crud(
            "sendlog-webhook".into(),
            WebhookType::SendLog,
            "https://example.com/".parse().unwrap(),
            Config::SendLog(SendLogConfig {
                uuid: Some(Uuid::new_v4().to_string()),
                payload: Some("Hello from $ALERT_NAME, a coralogix alert".into()),
            }),
        )
        .await;

        crud(
            "event-bridge-webhook".into(),
            WebhookType::AwsEventBridge,
            "https://example.com/".parse().unwrap(),
            Config::AwsEventBridge(AwsEventBridgeConfig {
                event_bus_arn: Some(
                    "arn:aws:events:us-east-1:123456789012:event-bus/default".into(),
                ),
                detail: Some("example-detail".into()),
                detail_type: Some("example-detail-type".into()),
                source: Some("example-source".into()),
                role_name: Some("example-role-name".into()),
            }),
        )
        .await;
        let client = WebhooksClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let result = client
            .test_webhook(
                WebhookType::Generic,
                Some("custom-webhook".into()),
                "https://httpbin.org/status/200".parse().unwrap(),
                Config::GenericWebhook(GenericWebhookConfig {
                    uuid: Some(Uuid::new_v4().to_string()),
                    method: generic_webhook_config::MethodType::Get.into(),
                    headers: Default::default(),
                    payload: None,
                }),
            )
            .await
            .unwrap();
        match result.result.clone().unwrap() {
            webhooks::WebhookTestResult::Success(_) => {}
            _ => panic!("Test webhook failed: {:?}", result.result.unwrap()),
        }
    }

    async fn crud(name: String, t: WebhookType, url: String, config: Config) {
        let client = WebhooksClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();
        let parsed_url: Url = url.parse().unwrap();

        let hook = client
            .create(t, Some(name.clone()), parsed_url.clone(), config.clone())
            .await
            .unwrap();
        let _ = client
            .replace(
                hook.id.clone().unwrap(),
                Some(format!("{}-2", name)),
                t,
                parsed_url,
                config,
            )
            .await
            .unwrap();
        let _ = client.get(hook.id.clone().unwrap()).await.unwrap();
        let _ = client.delete(hook.id.clone().unwrap()).await.unwrap();
    }
}
