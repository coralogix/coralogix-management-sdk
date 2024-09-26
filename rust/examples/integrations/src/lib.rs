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
        auth::AuthContext,
        client::integrations::{
            IntegrationsClient,
            Parameter,
            StringList,
            TestResult,
            Value,
        },
        CoralogixRegion,
    };

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
}
