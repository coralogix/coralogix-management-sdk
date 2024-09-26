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
        auth::AuthContext,
        client::alerts::{
            integration_type,
            ActivitySchedule,
            AlertDef,
            AlertDefNotificationGroup,
            AlertDefPriority,
            AlertDefProperties,
            AlertDefTargetSimple,
            AlertDefType,
            AlertsClient,
            DayOfWeek,
            EvaluationWindow,
            FilterType,
            IntegrationType,
            LabelFilterType,
            LabelFilters,
            LogFilterOperationType,
            LogSeverity,
            LogsFilter,
            LogsMoreThanTypeDefinition,
            LogsTimeWindow,
            LogsTimeWindowType,
            LogsTimeWindowValue,
            LuceneFilter,
            Recipients,
            Schedule,
            Targets,
            TimeOfDay,
            TypeDefinition,
        },
        CoralogixRegion,
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
                alert_def_type: AlertDefType::LogsMoreThan.into(),
                group_by: vec![],
                incidents_settings: None,
                phantom_mode: Some(false),
                notification_group: Some(AlertDefNotificationGroup {
                    group_by_fields: vec!["coralogix.metadata.sdkId".into(), "EventType".into()],
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
                schedule: Some(Schedule::ActiveOn(ActivitySchedule {
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
                type_definition: Some(TypeDefinition::LogsMoreThan(LogsMoreThanTypeDefinition {
                    logs_filter: Some(LogsFilter {
                        filter_type: Some(FilterType::LuceneFilter(LuceneFilter {
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
                    threshold: Some(5),
                    time_window: Some(LogsTimeWindow {
                        r#type: Some(LogsTimeWindowType::LogsTimeWindowSpecificValue(
                            LogsTimeWindowValue::Minutes30.into(),
                        )),
                    }),
                    evaluation_window: EvaluationWindow::RollingOrUnspecified.into(),
                    notification_payload_filter: vec![],
                })),
            }),
            id: None,
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
}
