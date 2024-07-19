#[cfg(test)]
mod tests {
    use cx_sdk::{
        auth::ApiKey,
        client::webhooks::{
            generic_webhook_config, slack_config, AwsEventBridgeConfig, Config, DemistoConfig,
            EmailGroupConfig, GenericWebhookConfig, JiraConfig, MicrosoftTeamsConfig,
            OpsgenieConfig, PagerDutyConfig, SendLogConfig, SlackConfig, WebhookType,
            WebhooksClient,
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
                method: generic_webhook_config::MethodType::Post.into(),
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
            "microsoft-teams-webhook".into(),
            WebhookType::MicrosoftTeams,
            "https://teams.microsoft.com/".parse().unwrap(),
            Config::MicrosoftTeams(MicrosoftTeamsConfig {}),
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
    }

    async fn crud(name: String, t: WebhookType, url: String, config: Config) {
        let client = WebhooksClient::new(
            ApiKey::from_env().unwrap(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();
        let parsed_url: Url = url.parse().unwrap();

        let hook = client
            .create(t, Some(name.clone()), parsed_url.clone(), config.clone())
            .await
            .unwrap();
        let _ = client.get(hook.id.clone().unwrap()).await.unwrap();
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
        let _ = client.delete(hook.id.clone().unwrap()).await.unwrap();
    }
}
