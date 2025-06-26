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
    use std::env;

    use cx_sdk::{
        CoralogixRegion,
        auth::AuthContext,
        client::{
            contextual_data_integrations::{
                ContextualDataIntegrationsClient,
                GenericIntegrationParameters,
                IntegrationMetadata,
                SpecificData,
            },
            integrations::{
                IntegrationsClient,
                Parameter,
                StringList,
                TestResult,
                Value,
            },
            webhooks::{
                self,
                AwsEventBridgeConfig,
                Config,
                DemistoConfig,
                EmailGroupConfig,
                GenericWebhookConfig,
                JiraConfig,
                OpsgenieConfig,
                PagerDutyConfig,
                SendLogConfig,
                SlackConfig,
                WebhookType,
                WebhooksClient,
                generic_webhook_config,
                slack_config,
            },
        },
    };
    use url::Url;
    use uuid::Uuid;

    fn bp<S: Into<String>>(k: S, v: bool) -> Parameter {
        Parameter {
            key: k.into(),
            value: Some(Value::BooleanValue(v)),
        }
    }

    fn sp<S: Into<String>>(k: S, v: S) -> Parameter {
        Parameter {
            key: k.into(),
            value: Some(Value::StringValue(v.into())),
        }
    }

    fn sl<S: Into<String>>(k: S, v: Vec<S>) -> Parameter {
        Parameter {
            key: k.into(),
            value: Some(Value::StringList(StringList {
                values: v.into_iter().map(Into::into).collect(),
            })),
        }
    }

    #[tokio::test]
    #[ignore = "Cleanup is not working on error"]
    async fn test_integrations() {
        let aws_region = std::env::var("AWS_REGION").unwrap();
        let client = IntegrationsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let role = env::var("AWS_TEST_ROLE").unwrap();
        let _ = client.list(false).await.unwrap();
        let name: String = "aws-metrics-collector".into();
        let version: String = "0.1.0".into();
        let params = vec![
            sp("ApplicationName", "cxsdk"),
            sp("SubsystemName", "aws-metrics-collector"),
            sl("MetricNamespaces", vec!["AWS/S3"]),
            sp("AwsRoleArn", &role),
            sp("IntegrationName", "sdk-integration-setup"),
            sp("AwsRegion", &aws_region),
            bp("WithAggregations", false),
            bp("EnrichWithTags", true),
        ];
        // test before saving
        let testing = client
            .test_integration(
                None,
                name.clone(),
                Some(version.clone()),
                Some(params.clone()),
            )
            .await
            .unwrap()
            .result
            .unwrap()
            .result
            .unwrap();

        match testing {
            TestResult::Success(_) => {}
            TestResult::Failure(f) => {
                panic!(
                    "Testing the integration didn't suceed test failed: {:?}",
                    f.error_message
                );
            }
        }

        // AWS Metrics Collector Integration Setup
        let create = client
            .create(name.clone(), Some(version.clone()), Some(params.clone()))
            .await
            .unwrap();

        let _ = client.get_details(name.clone(), true).await.unwrap();

        let id = create.integration_id.unwrap();

        let deployed_integration = client.get(id.clone()).await.unwrap();

        assert!(deployed_integration.integration.unwrap().id.unwrap() == id.clone());

        let all_instances = client.get_details(name.clone(), true).await.unwrap();
        match all_instances
            .integration_detail
            .unwrap()
            .integration_type_details
            .unwrap()
        {
            cx_sdk::client::integrations::IntegrationTypeDetails::Default(details) => {
                assert!(details.registered.into_iter().any(|i| i.id.unwrap() == id))
            }
        }

        client.delete(id).await.unwrap();
    }

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
                "https://httpbun.org/status/200".parse().unwrap(),
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

    #[tokio::test]
    async fn test_contextual_data_integrations() {
        let integration_key = "AWS-Status-Logs";
        let client = ContextualDataIntegrationsClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
        )
        .unwrap();

        // Get the integration definition
        let _ = client
            .get_definition(integration_key.into(), true)
            .await
            .unwrap();

        // Create a new integration

        let created_integration = client
            .create(IntegrationMetadata {
                integration_key: Some(integration_key.into()),
                version: Some("0.0.1".into()),
                specific_data: Some(SpecificData::IntegrationParameters(
                    GenericIntegrationParameters { parameters: vec![] },
                )),
            })
            .await
            .unwrap();

        let integration_id = created_integration.integration_id.unwrap();

        // List integrations
        let integrations = client.list(true).await.unwrap().integrations;
        assert!(!integrations.is_empty());

        // Delete the integration
        client.delete(integration_id.clone()).await.unwrap();

        // Verify deletion
        let integrations_after_deletion = client.list(true).await.unwrap().integrations;
        assert!(
            !integrations_after_deletion
                .into_iter()
                .any(|i| i.integration.unwrap().id == Some(integration_id.clone()))
        );
    }
}
