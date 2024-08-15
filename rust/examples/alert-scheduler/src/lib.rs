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
        client::alerts_scheduler::{
            AlertSchedulerClient, AlertSchedulerFilter, AlertSchedulerRule, AlertUniqueIds,
            OneTime, Schedule, ScheduleOperation, Scheduler, Timeframe, Until, WhichAlerts,
        },
        CoralogixRegion,
    };

    #[tokio::test]
    async fn test_alert_scheduler_client() {
        let alert_scheduler_client = AlertSchedulerClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
        )
        .unwrap();

        let alert_schduler_rule = AlertSchedulerRule {
            unique_identifier: None,
            id: None,
            name: String::from("example"),
            description: Some(String::from("example")),
            meta_labels: vec![],
            filter: Some(AlertSchedulerFilter {
                what_expression: "source logs | filter $d.cpodId:string == '122'".into(),
                which_alerts: Some(WhichAlerts::AlertUniqueIds(AlertUniqueIds {
                    value: vec!["55a457ed-5f23-407a-a724-12d7fe533a4e".into()],
                })),
            }),
            schedule: Some(Schedule {
                schedule_operation: ScheduleOperation::Mute.into(),
                scheduler: Some(Scheduler::OneTime(OneTime {
                    timeframe: Some(Timeframe {
                        start_time: String::from("2021-01-04T00:00:00.000"),
                        timezone: "UTC+2".to_string(),
                        until: Some(Until::EndTime(String::from("2025-01-01T00:00:50.000"))),
                    }),
                })),
            }),
            enabled: true,
            created_at: None,
            updated_at: None,
        };

        let created_alert_scheduler_rule = alert_scheduler_client
            .create(alert_schduler_rule)
            .await
            .unwrap();

        let new_alert_scheduler_rule = AlertSchedulerRule {
            name: String::from("MyAlertUpdated"),
            ..created_alert_scheduler_rule
        };

        let updated_alert_scheduler_rule = alert_scheduler_client
            .update(new_alert_scheduler_rule)
            .await
            .unwrap();

        let retrieved_alert_scheduler_rule = alert_scheduler_client
            .get(updated_alert_scheduler_rule.unique_identifier.unwrap())
            .await;

        let retrieved_alert_scheduler_rule = retrieved_alert_scheduler_rule.unwrap();

        assert!(retrieved_alert_scheduler_rule.name == "MyAlertUpdated");

        alert_scheduler_client
            .delete(retrieved_alert_scheduler_rule.unique_identifier.unwrap())
            .await
            .unwrap();
    }
}
