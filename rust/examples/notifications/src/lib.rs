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

    use std::hash::Hash;
    use std::vec;

    use cx_sdk::client::alerts::{
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
        NotificationDestination,
        NotificationRouter,
        NotifyOn,
        Recipients,
        RetriggeringPeriod,
        TimeOfDay,
        TypeDefinition,
        integration_type,
    };

    use cx_sdk::client::notifications::{
        RoutingRule,
        RoutingTarget,
    };
    use cx_sdk::{
        CoralogixRegion,
        auth::AuthContext,
        client::notifications::{
            ConditionType,
            ConfigOverrides,
            Connector,
            ConnectorConfig,
            ConnectorConfigField,
            ConnectorType,
            EntityType,
            EntityTypeConfigOverrides,
            GlobalRouter,
            MatchEntityTypeAndSubTypeCondition,
            MessageConfig,
            MessageConfigField,
            NotificationsClient,
            Preset,
            PresetType,
            TemplatedConnectorConfigField,
            condition_type::Condition,
            test_result::{
                self,
            },
        },
    };

    fn create_test_https_connector(name: String) -> Connector {
        Connector {
            name,
            description: "Connector for Notification Center testing.".into(),
            r#type: ConnectorType::GenericHttps.into(),
            id: None,
            connector_config: Some(ConnectorConfig {
                fields: vec![
                    ConnectorConfigField {
                        field_name: "url".into(),
                        value: "https://httpbun.org/post".into(),
                    },
                    ConnectorConfigField {
                        field_name: "method".into(),
                        value: "post".into(),
                    },
                ],
            }),
            team_id: None,
            create_time: None,
            update_time: None,
            config_overrides: vec![EntityTypeConfigOverrides {
                deprecated_entity_type: "alerts".into(),
                entity_type: EntityType::Alerts.into(),
                fields: vec![TemplatedConnectorConfigField {
                    field_name: "additionalBodyFields".into(),
                    template: "{\"priority\": \"{{alertDef.priority}}\"}".into(),
                }],
            }],
            user_defined_id: None,
        }
    }

    fn create_test_slack_connector() -> Connector {
        Connector {
            name: "TestSlackConnectorRust".into(),
            description: "Connector for Notification Center testing.".into(),
            r#type: ConnectorType::Slack.into(),
            id: None,
            connector_config: Some(ConnectorConfig {
                fields: vec![
                    ConnectorConfigField {
                        field_name: "integrationId".into(),
                        value: "c6e72871-388b-4776-a7de-8b7d17ed4828".into(),
                    },
                    ConnectorConfigField {
                        field_name: "channel".into(),
                        value: "luigis-testing-grounds".into(),
                    },
                    ConnectorConfigField {
                        field_name: "fallbackChannel".into(),
                        value: "luigis-testing-grounds".into(),
                    },
                ],
            }),
            team_id: None,
            create_time: None,
            update_time: None,
            config_overrides: vec![EntityTypeConfigOverrides {
                deprecated_entity_type: "alerts".into(),
                entity_type: EntityType::Alerts.into(),
                fields: vec![TemplatedConnectorConfigField {
                    field_name: "channel".into(),
                    template: "{{alertDef.priority}}".into(),
                }],
            }],
            user_defined_id: None,
        }
    }

    fn create_test_pagerduty_connector() -> Connector {
        Connector {
            name: "TestPagerDutyConnectorRust".into(),
            description: "Connector for Notification Center testing.".into(),
            r#type: ConnectorType::Pagerduty.into(),
            id: None,
            connector_config: Some(ConnectorConfig {
                fields: vec![ConnectorConfigField {
                    field_name: "integrationKey".into(),
                    value: "test".into(),
                }],
            }),
            team_id: None,
            create_time: None,
            update_time: None,
            config_overrides: vec![EntityTypeConfigOverrides {
                deprecated_entity_type: "alerts".into(),
                entity_type: EntityType::Alerts.into(),
                fields: vec![TemplatedConnectorConfigField {
                    field_name: "integrationKey".into(),
                    template: "integration_{{alertDef.priority}}".into(),
                }],
            }],
            user_defined_id: None,
        }
    }

    fn create_https_preset(name: String) -> Preset {
        Preset {
            name,
            description: "Preset for Notification Center testing.".into(),
            preset_type: Some(PresetType::Custom.into()),
            connector_type: ConnectorType::GenericHttps as i32,
            config_overrides: vec![ConfigOverrides {
                condition_type: Some(ConditionType {
                    condition: Some(Condition::MatchEntityTypeAndSubType(
                        MatchEntityTypeAndSubTypeCondition {
                            deprecated_entity_type: "alerts".into(),
                            entity_sub_type: "logsImmediateResolved".into(),
                            entity_type: EntityType::Alerts.into(),
                        },
                    )),
                }),
                message_config: Some(MessageConfig {
                    fields: vec![
                        MessageConfigField {
                            field_name: "headers".into(),
                            template: "{}".into(),
                        },
                        MessageConfigField {
                            field_name: "body".into(),
                            template: "{ \"groupingKey\": \"{{alert.groupingKey}}\", \"status\": \"{{alert.status}}\", \"groups\": \"{{alert.groups}}\" }".into(),
                        },
                    ],
                }),
                payload_type: None,
            }],
            create_time: None,
            entity_type: EntityType::Alerts.into(),
            id: None,
            update_time: None,
            parent_id: Some("preset_system_generic_https_alerts_empty".into()),
            ..Default::default()
        }
    }

    fn create_slack_preset() -> Preset {
        Preset {
            name: "TestSlackPreset".into(),
            description: "Preset for Notification Center testing.".into(),
            preset_type: Some(PresetType::Custom.into()),
            connector_type: ConnectorType::Slack as i32,
            config_overrides: vec![ConfigOverrides {
                condition_type: Some(ConditionType {
                    condition: Some(Condition::MatchEntityTypeAndSubType(
                        MatchEntityTypeAndSubTypeCondition {
                            deprecated_entity_type: "alerts".into(),
                            entity_sub_type: "logsImmediateResolved".into(),
                            entity_type: EntityType::Alerts.into(),
                        },
                    )),
                }),
                message_config: Some(MessageConfig {
                    fields: vec![
                        MessageConfigField {
                            field_name: "title".into(),
                            template: "{{alert.status}} {{alertDef.priority}} - {{alertDef.name}}".into(),
                        },
                        MessageConfigField {
                            field_name: "description".into(),
                            template: "{{alert.timestamp |  date(format=\"%Y-%m-%d %H:%M\")}}\nWe've detected a log that matches the query.".into(),
                        },
                    ],
                }),
                payload_type: None,
            }],
            create_time: None,
            entity_type: EntityType::Alerts.into(),
            id: None,
            update_time: None,
            parent_id: Some("preset_system_slack_alerts_empty".into()),
            ..Default::default()
        }
    }

    fn create_pagerduty_preset() -> Preset {
        Preset {
            name: "TestPagerDutyPreset".into(),
            description: "Preset for Notification Center testing.".into(),
            preset_type: Some(PresetType::Custom.into()),
            connector_type: ConnectorType::Pagerduty as i32,
            config_overrides: vec![ConfigOverrides {
                condition_type: Some(ConditionType {
                    condition: Some(Condition::MatchEntityTypeAndSubType(
                        MatchEntityTypeAndSubTypeCondition {
                            deprecated_entity_type: "alerts".into(),
                            entity_sub_type: "logsImmediateTriggered".into(),
                            entity_type: EntityType::Alerts.into(),
                        },
                    )),
                }),
                message_config: Some(MessageConfig {
                    fields: vec![MessageConfigField {
                        field_name: "summary".into(),
                        template: "{{ alertDef.name }}".into(),
                    }],
                }),
                payload_type: None,
            }],
            create_time: None,
            entity_type: EntityType::Alerts.into(),
            id: None,
            update_time: None,
            parent_id: Some("preset_system_pagerduty_alerts_basic".into()),
            ..Default::default()
        }
    }

    #[tokio::test]
    async fn test_https_connector() {
        let notifications_client = NotificationsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let connector = create_test_https_connector("TestHttpsConnectorRust".into());

        let fields = connector
            .connector_config
            .as_ref()
            .map_or(vec![], |f| f.fields.clone());

        let success: bool = match notifications_client
            .test_connector_config(
                (&connector.r#type()).clone(),
                fields,
                EntityType::Alerts,
                "generic_https_default".into(),
            )
            .await
            .unwrap()
            .result
            .unwrap()
            .result
            .unwrap()
        {
            test_result::Result::Success(_) => true,
            test_result::Result::Failure(f) => false,
        };
        assert!(success);

        let create_response = notifications_client
            .create_connector(connector.clone())
            .await
            .unwrap();
        let connector_id = create_response.connector.unwrap().id.unwrap();

        let success: bool = match notifications_client
            .test_existing_connector(connector_id.clone(), "generic_https_default".into())
            .await
            .unwrap()
            .result
            .unwrap()
            .result
            .unwrap()
        {
            test_result::Result::Success(_) => true,
            test_result::Result::Failure(f) => false,
        };

        assert!(success);
        let retrieved_connector = notifications_client
            .get_connector(connector_id.clone())
            .await
            .unwrap()
            .connector;

        assert_eq!(retrieved_connector.unwrap().name, connector.name);

        notifications_client
            .delete_connector(connector_id)
            .await
            .unwrap();
    }

    #[tokio::test]
    async fn test_slack_connector() {
        let notifications_client = NotificationsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let connector = create_test_slack_connector();

        let fields = connector
            .connector_config
            .as_ref()
            .map_or(vec![], |f| f.fields.clone());

        let success: bool = match notifications_client
            .test_connector_config(
                (&connector.r#type()).clone(),
                fields,
                EntityType::Alerts,
                "slack_raw".into(),
            )
            .await
            .unwrap()
            .result
            .unwrap()
            .result
            .unwrap()
        {
            test_result::Result::Success(_) => true,
            test_result::Result::Failure(f) => false,
        };
        assert!(success);

        let create_response = notifications_client
            .create_connector(connector.clone())
            .await
            .unwrap();
        let connector_id = create_response.connector.unwrap().id.unwrap();

        let success: bool = match notifications_client
            .test_existing_connector(connector_id.clone(), "slack_raw".into())
            .await
            .unwrap()
            .result
            .unwrap()
            .result
            .unwrap()
        {
            test_result::Result::Success(_) => true,
            test_result::Result::Failure(f) => false,
        };

        assert!(success);
        let retrieved_connector = notifications_client
            .get_connector(connector_id.clone())
            .await
            .unwrap()
            .connector;

        assert_eq!(retrieved_connector.unwrap().name, connector.name);

        notifications_client
            .delete_connector(connector_id)
            .await
            .unwrap();
    }

    #[tokio::test]
    async fn test_pagerduty_connector() {
        let notifications_client = NotificationsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let connector = create_test_pagerduty_connector();

        let create_response = notifications_client
            .create_connector(connector.clone())
            .await
            .unwrap();
        let connector_id = create_response.connector.unwrap().id.unwrap();

        let retrieved_connector = notifications_client
            .get_connector(connector_id.clone())
            .await
            .unwrap()
            .connector;

        assert_eq!(retrieved_connector.unwrap().name, connector.name);

        notifications_client
            .delete_connector(connector_id)
            .await
            .unwrap();
    }

    #[tokio::test]
    async fn test_https_preset() {
        let notifications_client = NotificationsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let preset = create_https_preset("TestHttpsPresetRust".into());

        let create_preset_response = notifications_client
            .create_custom_preset(preset.clone())
            .await
            .unwrap();

        let preset_id = create_preset_response.preset.unwrap().id.unwrap();
        let retrieved_preset = notifications_client
            .get_preset(preset_id.clone())
            .await
            .unwrap()
            .preset;

        assert_eq!(retrieved_preset.unwrap().name, preset.name);

        let _ = notifications_client
            .set_preset_as_default(preset_id.clone())
            .await
            .unwrap();
        let default_preset = notifications_client
            .get_default_preset_summary(
                ConnectorType::GenericHttps.into(),
                EntityType::Alerts.into(),
            )
            .await
            .unwrap()
            .preset_summary;

        assert_eq!(default_preset.unwrap().name, preset.name);

        notifications_client
            .delete_custom_preset(preset_id)
            .await
            .unwrap();
    }

    #[tokio::test]
    async fn test_slack_preset() {
        let notifications_client = NotificationsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let preset = create_slack_preset();

        let create_preset_response = notifications_client
            .create_custom_preset(preset.clone())
            .await
            .unwrap();

        let preset_id = create_preset_response.preset.unwrap().id.unwrap();
        let retrieved_preset = notifications_client
            .get_preset(preset_id.clone())
            .await
            .unwrap()
            .preset;

        assert_eq!(retrieved_preset.unwrap().name, preset.name);

        let _ = notifications_client
            .set_preset_as_default(preset_id.clone())
            .await
            .unwrap();
        let default_preset = notifications_client
            .get_default_preset_summary(ConnectorType::Slack.into(), EntityType::Alerts.into())
            .await
            .unwrap()
            .preset_summary;

        assert_eq!(default_preset.unwrap().name, preset.name);

        notifications_client
            .delete_custom_preset(preset_id)
            .await
            .unwrap();
    }

    #[tokio::test]
    async fn test_pagerduty_preset() {
        let notifications_client = NotificationsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let preset = create_pagerduty_preset();

        let create_preset_response = notifications_client
            .create_custom_preset(preset.clone())
            .await
            .unwrap();

        let preset_id = create_preset_response.preset.unwrap().id.unwrap();
        let retrieved_preset = notifications_client
            .get_preset(preset_id.clone())
            .await
            .unwrap()
            .preset;

        assert_eq!(retrieved_preset.unwrap().name, preset.name);

        let _ = notifications_client
            .set_preset_as_default(preset_id.clone())
            .await
            .unwrap();
        let default_preset = notifications_client
            .get_default_preset_summary(ConnectorType::Pagerduty.into(), EntityType::Alerts.into())
            .await
            .unwrap()
            .preset_summary;

        assert_eq!(default_preset.unwrap().name, preset.name);

        notifications_client
            .delete_custom_preset(preset_id)
            .await
            .unwrap();
    }

    #[tokio::test]
    async fn test_global_router() {
        let notifications_client = NotificationsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let name = uuid::Uuid::new_v4().to_string();
        let connector = create_test_https_connector(format!("TestHttpsConnectorRustGlobalRouter-%s", name));
        let create_response = notifications_client
            .create_connector(connector.clone())
            .await
            .unwrap();
        let connector_id = create_response.connector.unwrap().id.unwrap();
        let preset = create_https_preset("TestHttpsPresetRustGlobalRouter".into());
        let create_preset_response = notifications_client
            .create_custom_preset(preset.clone())
            .await
            .unwrap();
        let preset_id = create_preset_response.preset.unwrap().id.unwrap();

        let global_router = GlobalRouter {
            id: Some("router_default".into()),
            name: "TestGlobalRouter".to_string(),
            entity_type: EntityType::Alerts.into(),
            description: "Global Router for Notification Center testing.".to_string(),
            create_time: None,
            update_time: None,
            rules: vec![RoutingRule {
                condition: "alertDef.priority == \"P1\"".to_string(),
                targets: vec![RoutingTarget {
                    connector_id: connector_id.clone(),
                    preset_id: Some(preset_id.clone()),
                    ..Default::default()
                }],
                name: Some("TestRoutingRule".to_string()),
                ..Default::default()
            }],
            fallback: vec![],
            entity_labels: std::collections::HashMap::new(),
            ..Default::default()
        };

        let create_or_replace_response = notifications_client
            .create_or_replace_global_router(global_router.clone())
            .await
            .unwrap();

        let router_id = create_or_replace_response.router.unwrap().id.unwrap();

        let retrieved_router = notifications_client
            .get_global_router(router_id.clone())
            .await
            .unwrap()
            .router;

        assert_eq!(retrieved_router.unwrap().name, global_router.name);

        notifications_client
            .delete_global_router(router_id)
            .await
            .unwrap();

        notifications_client
            .delete_connector(connector_id)
            .await
            .unwrap();
        notifications_client
            .delete_custom_preset(preset_id)
            .await
            .unwrap();
    }

    #[tokio::test]
    async fn test_alert_with_destinations() {
        let notifications_client = NotificationsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();
        let alerts_client = AlertsClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
            None,
        )
        .unwrap();

        let connector =
            create_test_https_connector("TestHttpsConnectorRustAlertWithDestination".into());
        let create_response = notifications_client
            .create_connector(connector.clone())
            .await
            .unwrap();
        let connector_id = create_response.connector.unwrap().id.unwrap();

        let preset = create_https_preset("TestHttpsPresetRustAlertWithDestination".into());
        let create_preset_response = notifications_client
            .create_custom_preset(preset.clone())
            .await
            .unwrap();
        let preset_id = create_preset_response.preset.unwrap().id.unwrap();

        let notification_destination = NotificationDestination {
            connector_id: connector_id.clone(),
            preset_id: Some(preset_id.clone()),
            notify_on: NotifyOn::TriggeredAndResolved.into(),
            triggered_routing_overrides: None,
            resolved_route_overrides: None,
        };

        let alert_def = create_alert_with_destinations(vec![notification_destination]);
        let alert_def_response = alerts_client.create(alert_def.clone()).await.unwrap();

        notifications_client
            .test_destination(
                EntityType::Alerts,
                vec![],
                preset_id.clone(),
                connector_id.clone(),
                vec![],
                "generic_https_default".into(),
            )
            .await
            .unwrap();

        alerts_client
            .delete(alert_def_response.alert_def.unwrap().id.unwrap())
            .await
            .unwrap();

        notifications_client
            .delete_connector(connector_id)
            .await
            .unwrap();

        notifications_client
            .delete_custom_preset(preset_id)
            .await
            .unwrap();
    }

    #[tokio::test]
    async fn test_alert_with_router() {
        let notifications_client = NotificationsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();
        let alerts_client = AlertsClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
            None,
        )
        .unwrap();

        let connector = create_test_https_connector("TestHttpsConnectorRustAlertWithRouter".into());
        let create_response = notifications_client
            .create_connector(connector.clone())
            .await
            .unwrap();
        let connector_id = create_response.connector.unwrap().id.unwrap();

        let preset = create_https_preset("TestHttpsPresetRustAlertWithRouter".into());
        let create_preset_response = notifications_client
            .create_custom_preset(preset.clone())
            .await
            .unwrap();
        let preset_id = create_preset_response.preset.unwrap().id.unwrap();

        let global_router = GlobalRouter {
            id: Some("router_default".into()),
            name: "TestGlobalRouter".to_string(),
            entity_type: EntityType::Alerts.into(),
            description: "Global Router for Notification Center testing.".to_string(),
            create_time: None,
            update_time: None,
            rules: vec![RoutingRule {
                condition: "alertDef.priority == \"P1\"".to_string(),
                targets: vec![RoutingTarget {
                    connector_id: connector_id.clone(),
                    preset_id: Some(preset_id.clone()),
                    ..Default::default()
                }],
                name: Some("TestRoutingRule".to_string()),
                ..Default::default()
            }],
            ..Default::default()
        };

        notifications_client
            .create_or_replace_global_router(global_router.clone())
            .await
            .unwrap();

        let alert_def = create_alert_with_router();
        let alert_def_response = alerts_client.create(alert_def.clone()).await.unwrap();

        alerts_client
            .delete(alert_def_response.alert_def.unwrap().id.unwrap())
            .await
            .unwrap();

        notifications_client
            .delete_connector(connector_id)
            .await
            .unwrap();

        notifications_client
            .delete_custom_preset(preset_id)
            .await
            .unwrap();

        notifications_client
            .delete_global_router("router_default".into())
            .await
            .unwrap();
    }

    fn create_alert_with_destinations(destinations: Vec<NotificationDestination>) -> AlertDef {
        AlertDef {
            updated_time: None,
            created_time: None,
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
                    destinations,
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

    fn create_alert_with_router() -> AlertDef {
        AlertDef {
            updated_time: None,
            created_time: None,
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
                    router: Some(NotificationRouter {
                        id: "router_default".into(),
                        notify_on: Some(NotifyOn::TriggeredAndResolved.into()),
                    }),
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
}
