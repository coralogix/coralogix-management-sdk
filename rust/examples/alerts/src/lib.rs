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
                AlertDefOverride,
                AlertDefPriority,
                AlertDefProperties,
                AlertDefType,
                AlertDefWebhooksSettings,
                AlertsClient,
                BurnRateThreshold,
                BurnRateTypeSingle,
                DayOfWeek,
                DefaultLabels,
                DurationUnit,
                ErrorBudgetThreshold,
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
                NotifyOn,
                Recipients,
                RetriggeringPeriod,
                SloDefinition,
                SloThresholdCondition,
                SloThresholdRule,
                SloThresholdType,
                Threshold,
                TimeDuration,
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
            slos::{
                Metric,
                RequestBasedMetricSli,
                Sli,
                Slo,
                SloClient,
                SloTimeFrame,
                Window,
            },
        },
    };

    fn create_alert() -> AlertDef {
        AlertDef {
            updated_time: None,
            created_time: None,
            last_triggered_time: None,
            alert_def_properties: Some(AlertDefProperties {
                name: Some("Standard alert example".to_string()),
                description: Some("Example of standard alert from terraform".to_string()),
                enabled: Some(true),
                priority: AlertDefPriority::P1.into(),
                deleted: None,
                r#type: AlertDefType::LogsThreshold.into(),
                group_by_keys: vec![],
                incidents_settings: None,
                phantom_mode: Some(false),
                notification_group: Some(AlertDefNotificationGroup {
                    group_by_keys: vec![],
                    destinations: vec![],
                    webhooks: vec![AlertDefWebhooksSettings {
                        notify_on: Some(NotifyOn::TriggeredAndResolved.into()),
                        retriggering_period: Some(RetriggeringPeriod::Minutes(5)),
                        integration: Some(IntegrationType {
                            integration_type: Some(integration_type::IntegrationType::Recipients(
                                Recipients {
                                    emails: vec![String::from("example@coralogix.com")],
                                },
                            )),
                        }),
                    }],
                    router: None,
                }),
                entity_labels: [
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
                    evaluation_delay_ms: None,
                    rules: vec![LogsThresholdRule {
                        r#override: Some(AlertDefOverride {
                            priority: AlertDefPriority::P1.into(),
                        }),
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
                notification_group_excess: vec![],
            }),
            id: None,
            alert_version_id: None,
        }
    }

    fn create_burn_rate_slo_alert(slo_id: String) -> AlertDef {
        AlertDef {
            updated_time: None,
            created_time: None,
            last_triggered_time: None,
            alert_def_properties: Some(AlertDefProperties {
                name: Some("Standard alert example".to_string()),
                description: Some("Example of standard alert from terraform".to_string()),
                enabled: Some(true),
                priority: AlertDefPriority::P1.into(),
                deleted: None,
                r#type: AlertDefType::SloThreshold.into(),
                group_by_keys: vec![],
                incidents_settings: None,
                phantom_mode: Some(false),
                notification_group: Some(AlertDefNotificationGroup {
                    group_by_keys: vec![],
                    destinations: vec![],
                    webhooks: vec![AlertDefWebhooksSettings {
                        notify_on: Some(NotifyOn::TriggeredAndResolved.into()),
                        retriggering_period: Some(RetriggeringPeriod::Minutes(5)),
                        integration: Some(IntegrationType {
                            integration_type: Some(integration_type::IntegrationType::Recipients(
                                Recipients {
                                    emails: vec![String::from("example@coralogix.com")],
                                },
                            )),
                        }),
                    }],
                    router: None,
                }),
                entity_labels: [
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
                type_definition: Some(TypeDefinition::SloThreshold(SloThresholdType {
                    slo_definition: Some(SloDefinition {
                        slo_id: Some(slo_id),
                    }),
                    threshold: Some(Threshold::BurnRate(BurnRateThreshold {
                        rules: vec![SloThresholdRule {
                            r#override: Some(AlertDefOverride {
                                priority: AlertDefPriority::P1.into(),
                            }),
                            condition: Some(SloThresholdCondition {
                                threshold: Some(1.0),
                            }),
                        }],
                        r#type: Some(alerts::BurnRateThresholdType::Single(BurnRateTypeSingle {
                            time_duration: Some(TimeDuration {
                                duration: Some(1),
                                unit: DurationUnit::Hours.into(),
                            }),
                        })),
                    })),
                })),
                notification_group_excess: vec![],
            }),
            id: None,
            alert_version_id: None,
        }
    }

    fn create_error_budget_slo_alert(slo_id: String) -> AlertDef {
        AlertDef {
            updated_time: None,
            created_time: None,
            last_triggered_time: None,
            alert_def_properties: Some(AlertDefProperties {
                name: Some("Standard alert example".to_string()),
                description: Some("Example of standard alert from terraform".to_string()),
                enabled: Some(true),
                priority: AlertDefPriority::P1.into(),
                deleted: None,
                r#type: AlertDefType::SloThreshold.into(),
                group_by_keys: vec![],
                incidents_settings: None,
                phantom_mode: Some(false),
                notification_group: Some(AlertDefNotificationGroup {
                    group_by_keys: vec![],
                    destinations: vec![],
                    webhooks: vec![AlertDefWebhooksSettings {
                        notify_on: Some(NotifyOn::TriggeredAndResolved.into()),
                        retriggering_period: Some(RetriggeringPeriod::Minutes(5)),
                        integration: Some(IntegrationType {
                            integration_type: Some(integration_type::IntegrationType::Recipients(
                                Recipients {
                                    emails: vec![String::from("example@coralogix.com")],
                                },
                            )),
                        }),
                    }],
                    router: None,
                }),
                entity_labels: [
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
                type_definition: Some(TypeDefinition::SloThreshold(SloThresholdType {
                    slo_definition: Some(SloDefinition {
                        slo_id: Some(slo_id),
                    }),
                    threshold: Some(Threshold::ErrorBudget(ErrorBudgetThreshold {
                        rules: vec![SloThresholdRule {
                            r#override: Some(AlertDefOverride {
                                priority: AlertDefPriority::P3.into(),
                            }),
                            condition: Some(SloThresholdCondition {
                                threshold: Some(1.0),
                            }),
                        }],
                    })),
                })),
                notification_group_excess: vec![],
            }),
            id: None,
            alert_version_id: None,
        }
    }

    #[tokio::test]
    async fn test_alerts() {
        let alerts_client = AlertsClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
            Some(DefaultLabels::SdkVersion),
        )
        .unwrap();
        let alert = create_alert();

        let created_alert = alerts_client
            .create(alert.clone())
            .await
            .unwrap()
            .alert_def
            .unwrap();

        let retrieved_alert = alerts_client
            .get(created_alert.id.clone().unwrap())
            .await
            .unwrap()
            .alert_def;

        assert_eq!(retrieved_alert.unwrap(), created_alert);

        let retrieved_alerts = alerts_client.list().await.unwrap().alert_defs;

        assert!(!retrieved_alerts.is_empty());

        let updated_alert = AlertDef {
            alert_def_properties: Some(AlertDefProperties {
                description: Some("updated description".to_string()),
                ..created_alert.alert_def_properties.clone().unwrap()
            }),
            ..created_alert
        };

        let updated_alert = alerts_client
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

        alerts_client
            .delete(updated_alert.id.unwrap())
            .await
            .unwrap();
    }

    #[tokio::test]
    #[ignore = "Unstable test"]
    async fn test_burn_rate_slo_alert() {
        let alerts_client = AlertsClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
            Some(DefaultLabels::SdkVersion),
        )
        .unwrap();
        let slos_client = SloClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();
        let slo = Slo {
            id: None,
            name: "coralogix_rust_slo_example".into(),
            description: Some("description".to_string()),
            creator: None,
            labels: vec![("label1".to_string(), "value1".to_string())]
                .into_iter()
                .collect(),
            target_threshold_percentage: 95.0f32,
            create_time: None,
            update_time: None,
            sli: Some(Sli::RequestBasedMetricSli(RequestBasedMetricSli {
                good_events: Some(Metric {
                    query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)".to_string(),
                }),
                total_events: Some(Metric {
                    query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)".to_string(),
                }),
            })),
            window: Some(Window::SloTimeFrame(SloTimeFrame::SloTimeFrame7Days.into())),
            revision: None,
            grouping: None,
        };

        let create_slo_response = slos_client.create(slo.clone()).await.unwrap();
        let slo_id = create_slo_response.slo.unwrap().id.clone().unwrap();
        let alert = create_burn_rate_slo_alert(slo_id.clone());

        let created_alert = alerts_client
            .create(alert.clone())
            .await
            .unwrap()
            .alert_def
            .unwrap();

        let retrieved_alert = alerts_client
            .get(created_alert.id.clone().unwrap())
            .await
            .unwrap()
            .alert_def;

        assert_eq!(retrieved_alert.unwrap(), created_alert);

        let retrieved_alerts = alerts_client.list().await.unwrap().alert_defs;

        assert!(!retrieved_alerts.is_empty());

        let updated_alert = AlertDef {
            alert_def_properties: Some(AlertDefProperties {
                description: Some("updated description".to_string()),
                ..created_alert.alert_def_properties.clone().unwrap()
            }),
            ..created_alert
        };

        let updated_alert = alerts_client
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

        alerts_client
            .delete(updated_alert.id.unwrap())
            .await
            .unwrap();
        slos_client.delete(slo_id).await.unwrap();
    }

    #[tokio::test]
    #[ignore = "Unstable test"]
    async fn test_error_budget_slo_alert() {
        let alerts_client = AlertsClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
            Some(DefaultLabels::SdkVersion),
        )
        .unwrap();
        let slos_client = SloClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();
        let slo = Slo {
            id: None,
            name: "coralogix_rust_slo_example".into(),
            description: Some("description".to_string()),
            creator: None,
            labels: vec![("label1".to_string(), "value1".to_string())]
                .into_iter()
                .collect(),
            target_threshold_percentage: 95.0f32,
            create_time: None,
            update_time: None,
            sli: Some(Sli::RequestBasedMetricSli(RequestBasedMetricSli {
                good_events: Some(Metric {
                    query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)".to_string(),
                }),
                total_events: Some(Metric {
                    query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)".to_string(),
                }),
            })),
            window: Some(Window::SloTimeFrame(SloTimeFrame::SloTimeFrame7Days.into())),
            revision: None,
            grouping: None,
        };

        let create_slo_response = slos_client.create(slo.clone()).await.unwrap();
        let slo_id = create_slo_response.slo.unwrap().id.clone().unwrap();
        let alert = create_error_budget_slo_alert(slo_id.clone());

        let created_alert = alerts_client
            .create(alert.clone())
            .await
            .unwrap()
            .alert_def
            .unwrap();

        let retrieved_alert = alerts_client
            .get(created_alert.id.clone().unwrap())
            .await
            .unwrap()
            .alert_def;

        assert_eq!(retrieved_alert.unwrap(), created_alert);

        let retrieved_alerts = alerts_client.list().await.unwrap().alert_defs;

        assert!(!retrieved_alerts.is_empty());

        let updated_alert = AlertDef {
            alert_def_properties: Some(AlertDefProperties {
                description: Some("updated description".to_string()),
                ..created_alert.alert_def_properties.clone().unwrap()
            }),
            ..created_alert
        };

        let updated_alert = alerts_client
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

        alerts_client
            .delete(updated_alert.id.unwrap())
            .await
            .unwrap();
        slos_client.delete(slo_id).await.unwrap();
    }

    #[tokio::test]
    #[ignore = "Unstable test"]
    async fn test_alert_gets_deleted_on_slo_deletion() {
        let alerts_client = AlertsClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
            Some(DefaultLabels::SdkVersion),
        )
        .unwrap();
        let slos_client = SloClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();
        let slo = Slo {
            id: None,
            name: "coralogix_rust_slo_example".into(),
            description: Some("description".to_string()),
            creator: None,
            labels: vec![("label1".to_string(), "value1".to_string())]
                .into_iter()
                .collect(),
            target_threshold_percentage: 95.0f32,
            create_time: None,
            update_time: None,
            sli: Some(Sli::RequestBasedMetricSli(RequestBasedMetricSli {
                good_events: Some(Metric {
                    query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)".to_string(),
                }),
                total_events: Some(Metric {
                    query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)".to_string(),
                }),
            })),
            window: Some(Window::SloTimeFrame(SloTimeFrame::SloTimeFrame7Days.into())),
            revision: None,
            grouping: None,
        };

        let create_slo_response = slos_client.create(slo.clone()).await.unwrap();
        let slo_id = create_slo_response.slo.unwrap().id.clone().unwrap();
        let alert = create_burn_rate_slo_alert(slo_id.clone());

        let created_alert = alerts_client
            .create(alert.clone())
            .await
            .unwrap()
            .alert_def
            .unwrap();

        let retrieved_alert = alerts_client
            .get(created_alert.id.clone().unwrap())
            .await
            .unwrap()
            .alert_def;

        assert_eq!(retrieved_alert.unwrap(), created_alert);

        let retrieved_alerts = alerts_client.list().await.unwrap().alert_defs;

        assert!(!retrieved_alerts.is_empty());

        let updated_alert = AlertDef {
            alert_def_properties: Some(AlertDefProperties {
                description: Some("updated description".to_string()),
                ..created_alert.alert_def_properties.clone().unwrap()
            }),
            ..created_alert
        };

        let updated_alert = alerts_client
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

        slos_client.delete(slo_id).await.unwrap();

        let alert = alerts_client.get(updated_alert.id.unwrap()).await;

        assert!(
            alert.is_err(),
            "Alert should be deleted when SLO is deleted"
        );
    }

    #[tokio::test]
    async fn test_alerts_rm_schedule() {
        let alerts_client = AlertsClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
            Some(DefaultLabels::SdkVersion),
        )
        .unwrap();
        let alert = create_alert();

        let created_alert = alerts_client
            .create(alert.clone())
            .await
            .unwrap()
            .alert_def
            .unwrap();

        let retrieved_alert = alerts_client
            .get(created_alert.id.clone().unwrap())
            .await
            .unwrap()
            .alert_def;

        assert_eq!(retrieved_alert.unwrap(), created_alert);

        let retrieved_alerts = alerts_client.list().await.unwrap().alert_defs;

        assert!(!retrieved_alerts.is_empty());

        let updated_alert = AlertDef {
            alert_def_properties: Some(AlertDefProperties {
                schedule: None,
                ..created_alert.alert_def_properties.clone().unwrap()
            }),
            ..created_alert
        };

        let updated_alert = alerts_client
            .replace(updated_alert.clone())
            .await
            .unwrap()
            .alert_def
            .unwrap();

        assert!(
            updated_alert
                .alert_def_properties
                .unwrap()
                .schedule
                .is_none()
        );

        let retrieved_alert = alerts_client
            .get(updated_alert.id.clone().unwrap())
            .await
            .unwrap()
            .alert_def;

        assert!(
            retrieved_alert
                .unwrap()
                .alert_def_properties
                .unwrap()
                .schedule
                .is_none()
        );

        alerts_client
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
        let alerts_client = AlertsClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
            Some(DefaultLabels::SdkVersion),
        )
        .unwrap();
        let alert = create_alert();

        let created_alert = alerts_client
            .create(alert.clone())
            .await
            .unwrap()
            .alert_def
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
                    value: vec![created_alert.id.clone().unwrap()],
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
            id: None,
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

        alerts_client
            .delete(created_alert.id.unwrap())
            .await
            .unwrap();
    }
}
