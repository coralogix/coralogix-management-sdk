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
    use std::collections::HashMap;

    use cx_sdk::{
        CoralogixRegion,
        auth::AuthContext,
        client::recording_rule_group_sets::{
            InRule,
            InRuleGroup,
            RecordingRuleGroupSetsClient,
        },
    };

    #[tokio::test]
    async fn test_recording_rule_group_sets() {
        let client = RecordingRuleGroupSetsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();
        let create_rule_group_set_response = client
            .create(
                format!("TestRecordingRuleGroupSet {}", chrono::Utc::now().timestamp_millis()).into(),
                vec![InRuleGroup {
                    name: "foo".into(),
                    interval: Some(180),
                    limit: Some(100),
                    rules: vec![InRule {
                        record: "ts3db_live_ingester_write_latency:3m".into(),
                        expr: "sum(rate(ts3db_live_ingester_write_latency_seconds_count{CX_LEVEL=\"staging\",pod=~\"ts3db-live-ingester.*\"}[2m])) by (pod)".into(),
                        labels: HashMap::new(),
                    }, InRule{ record: "job:http_requests_total:sum".into(), expr: "sum(rate(http_requests_total[5m])) by (job)".into(), labels: HashMap::new() }],
                    id: None, 
                    version: None,
                }, InRuleGroup {
                    name: "bar".into(),
                    interval: Some(180),
                    limit: Some(100),
                    rules: vec![InRule {
                        record: "ts3db_live_ingester_write_latency:3m".into(),
                        expr: "sum(rate(ts3db_live_ingester_write_latency_seconds_count{CX_LEVEL=\"staging\",pod=~\"ts3db-live-ingester.*\"}[2m])) by (pod)".into(),
                        labels: HashMap::new(),
                    }, InRule{ record: "job:http_requests_total:sum".into(), expr: "sum(rate(http_requests_total[5m])) by (job)".into(), labels: HashMap::new() }],
                    id: None, 
                    version: None,
                }],
            )
            .await.unwrap();

        let rule_group_set = client
            .get(create_rule_group_set_response.id.clone())
            .await
            .unwrap();

        assert!(rule_group_set.groups.len() == 2);

        client
            .delete(create_rule_group_set_response.id)
            .await
            .unwrap();
    }
}
