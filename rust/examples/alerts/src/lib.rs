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

    use std::vec;

    use cx_sdk::{
        CoralogixRegion,
        auth::AuthContext,
        client::{
            alerts::{
                self,
                ActivitySchedule,
                AlertDef,
                AlertDefNotificationGroup,
                AlertDefPriority,
                AlertDefProperties,
                AlertDefTargetSimple,
                AlertDefType,
                AlertsClient,
                DayOfWeek,
                FilterType,
                IntegrationType,
                LabelFilterType,
                LabelFilters,
                LogFilterOperationType,
                LogSeverity,
                LogsFilter,
                LogsSimpleFilter,
                LogsThresholdCondition,
                LogsThresholdConditionType,
                LogsThresholdRule,
                LogsThresholdType,
                LogsTimeWindow,
                LogsTimeWindowType,
                LogsTimeWindowValue,
                Recipients,
                Targets,
                TimeOfDay,
                TypeDefinition,
                integration_type,
            },
            alerts_scheduler::{
                AlertSchedulerClient,
                AlertSchedulerFilter,
                AlertSchedulerRule,
                AlertUniqueIds,
                OneTime,
                Schedule,
                ScheduleOperation,
                Scheduler,
                Timeframe,
                Until,
                WhichAlerts,
            },
        },
    };

    #[tokio::test]
    async fn test_alerts() {
        let alerts_service = AlertsClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
        )
        .unwrap();
        let alert = AlertDef {
            updated_time: None,
            created_time: None,
            alert_def_properties: Some(AlertDefProperties {
                name: Some("Standard alert example".to_string()),
                description: Some("Example of standard alert from terraform".to_string()),
                enabled: Some(true),
                priority: AlertDefPriority::P1.into(),
                deleted: None,
                r#type: AlertDefType::LogsThreshold.into(),
                group_by: vec![],
                incidents_settings: None,
                phantom_mode: Some(false),
                notification_group: Some(AlertDefNotificationGroup {
                    group_by_fields: vec![],
                    targets: Some(Targets::Simple(AlertDefTargetSimple {
                        integrations: vec![IntegrationType {
                            integration_type: Some(integration_type::IntegrationType::Recipients(
                                Recipients {
                                    emails: vec![String::from("example@coralogix.com")],
                                },
                            )),
                        }],
                    })),
                }),
                labels: [
                    ("alert_type".to_string(), "security".to_string()),
                    ("security_severity".to_string(), "high".to_string()),
                ]
                .into_iter()
                .collect(),
                schedule: Some(alerts::Schedule::ActiveOn(ActivitySchedule {
                    day_of_week: vec![DayOfWeek::Wednesday.into(), DayOfWeek::Thursday.into()],
                    start_time: Some(TimeOfDay {
                        hours: 8,
                        minutes: 30,
                    }),
                    end_time: Some(TimeOfDay {
                        hours: 20,
                        minutes: 30,
                    }),
                })),
                type_definition: Some(TypeDefinition::LogsThreshold(LogsThresholdType {
                    logs_filter: Some(LogsFilter {
                        filter_type: Some(FilterType::SimpleFilter(LogsSimpleFilter {
                            lucene_query: Some(String::from("remote_addr_enriched:/.*/")),
                            label_filters: Some(LabelFilters {
                                application_name: vec![LabelFilterType {
                                    value: Some(String::from("nginx")),
                                    operation: LogFilterOperationType::Includes.into(),
                                }],
                                subsystem_name: vec![LabelFilterType {
                                    value: Some(String::from("subsystem-name")),
                                    operation: LogFilterOperationType::StartsWith.into(),
                                }],
                                severities: vec![
                                    LogSeverity::Warning.into(),
                                    LogSeverity::Info.into(),
                                ],
                            }),
                        })),
                    }),
                    notification_payload_filter: vec![],
                    undetected_values_management: None,
                    rules: vec![LogsThresholdRule {
                        condition: Some(LogsThresholdCondition {
                            threshold: Some(10.0),
                            condition_type: LogsThresholdConditionType::MoreThanOrUnspecified
                                .into(),
                            time_window: Some(LogsTimeWindow {
                                r#type: Some(LogsTimeWindowType::LogsTimeWindowSpecificValue(
                                    LogsTimeWindowValue::Minutes5OrUnspecified.into(),
                                )),
                            }),
                        }),
                    }],
                })),
            }),
            id: None,
            alert_version_id: None,
        };

        let created_alert = alerts_service
            .create(alert.clone())
            .await
            .unwrap()
            .alert_def
            .unwrap();

        let retrieved_alert = alerts_service
            .get(created_alert.id.clone().unwrap())
            .await
            .unwrap()
            .alert_def;

        assert_eq!(retrieved_alert.unwrap(), created_alert);

        let updated_alert = AlertDef {
            alert_def_properties: Some(AlertDefProperties {
                description: Some("updated description".to_string()),
                ..created_alert.alert_def_properties.clone().unwrap()
            }),
            ..created_alert
        };

        let updated_alert = alerts_service
            .replace(updated_alert.clone())
            .await
            .unwrap()
            .alert_def
            .unwrap();

        assert!(
            updated_alert
                .alert_def_properties
                .unwrap()
                .description
                .unwrap()
                == "updated description"
        );

        alerts_service
            .delete(updated_alert.id.unwrap())
            .await
            .unwrap();
    }

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
