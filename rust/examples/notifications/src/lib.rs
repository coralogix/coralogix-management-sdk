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

// 1. Create Connectors
// User creates connector: ConnectorsService/CreateConnector API
// User send test notification to confirm that connector configuration is correct: TestingService/TestExistingConnector API
// 2. Create Custom Presets [Optional - system presets can be used out of box]
// User create custom preset: PresetsService/CreateCustomPreset API
// User test custom preset: TestingService/TestExistingPreset API
// 3. Set Preset as default [Optional - there is always default provided by the system]
// User sets selected preset as default: PresetsService/SetPresetAsDefault API
// 4. Option 1: User use configured connectors and presets as a destinations in the Alerts v3 API
// User test destination using TestingService/TestDestination
// 4. Option 2: User configure global router and use router in Alerts v3 API
// User test routing condition (if the syntax is correct): TestingService/TestRoutingConditionValid
// User create Global Router using GlobalRoutersService/CreateGlobalRouter
// User use Global router in Alerts v3 notifications configuration

#[cfg(test)]
mod tests {

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
            GlobalRouter,
            MatchEntityTypeCondition,
            MessageConfig,
            MessageConfigField,
            NotificationsClient,
            Preset,
            condition_type,
        },
    };

    fn create_test_connector() -> Connector {
        Connector {
            name: "TestConnector".into(),
            description: "Connector for Notification Center testing.".into(),
            r#type: ConnectorType::GenericHttps as i32,
            id: None,
            connector_config: Some(ConnectorConfig {
                fields: vec![
                    ConnectorConfigField {
                        field_name: "url".into(),
                        value: "https://httpbin.org/post".into(),
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
            config_overrides: vec![],
        }
    }

    fn create_test_preset() -> Preset {
        Preset {
            name: "TestPreset".into(),
            description: "Preset for Notification Center testing.".into(),
            preset_type: Some(2),
            connector_type: ConnectorType::GenericHttps as i32,
            config_overrides: vec![ConfigOverrides {
                condition_type: Some(ConditionType {
                    condition: Some(condition_type::Condition::MatchEntityType(
                        MatchEntityTypeCondition {
                            ..Default::default()
                        },
                    )),
                }),
                message_config: Some(MessageConfig {
                    fields: vec![
                        MessageConfigField {
                            field_name: "headers".into(),
                            template: "{ \"example\": \"header_value\" }".into(),
                        },
                        MessageConfigField {
                            field_name: "body".into(),
                            template: "{ \"example\": \"body_value\" }".into(),
                        },
                    ],
                }),
                payload_type: "web".into(),
            }],
            create_time: None,
            entity_type: EntityType::Alerts.into(),
            id: None,
            update_time: None,
            parent: None,
            ..Default::default()
        }
    }

    #[tokio::test]
    async fn test_connectors() {
        let notifications_client = NotificationsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let connector = create_test_connector();

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
    async fn test_presets() {
        let notifications_client = NotificationsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let preset = create_test_preset();

        let create_response = notifications_client
            .create_custom_preset(preset.clone())
            .await
            .unwrap();
        let preset_id = create_response.preset.unwrap().id.unwrap();

        let retrieved_preset = notifications_client
            .get_preset(preset_id.clone())
            .await
            .unwrap()
            .preset;

        assert_eq!(retrieved_preset.unwrap().name, preset.name);

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

        let list_response = notifications_client
            .list_global_routers(EntityType::Alerts)
            .await
            .unwrap();

        let existing_router_id = list_response.routers.first().and_then(|r| r.id.clone());

        let global_router = GlobalRouter {
            id: existing_router_id.clone(),
            name: "TestGlobalRouter".to_string(),
            entity_type: EntityType::Alerts.into(),
            description: "Global Router for Notification Center testing.".to_string(),
            create_time: None,
            update_time: None,
            rules: vec![],
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
            .get_global_router(router_id)
            .await
            .unwrap()
            .router;

        assert_eq!(retrieved_router.unwrap().name, global_router.name);
    }
}
