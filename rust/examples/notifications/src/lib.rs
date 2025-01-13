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
        CoralogixRegion,
        auth::AuthContext,
        client::notifications::{
            ConditionType,
            ConfigOverrides,
            Connector,
            ConnectorConfig,
            ConnectorConfigField,
            ConnectorType,
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
            name: "TestConnector".to_string(),
            description: "Connector for Notification Center testing.".to_string(),
            r#type: ConnectorType::GenericHttps as i32,
            entity_type: "alerts".to_string(),
            connector_configs: vec![ConnectorConfig {
                output_schema_id: "default".to_string(),
                fields: vec![
                    ConnectorConfigField {
                        field_name: "url".to_string(),
                        template: "https://example.coralogix.com".to_string(),
                    },
                    ConnectorConfigField {
                        field_name: "method".to_string(),
                        template: "post".to_string(),
                    },
                ],
            }],
            id: None,
            user_facing_id: None,
            team_id: None,
            create_time: None,
            update_time: None,
        }
    }

    fn create_test_preset() -> Preset {
        Preset {
            name: "TestPreset".to_string(),
            description: "Preset for Notification Center testing.".to_string(),
            preset_type: Some(2),
            connector_type: ConnectorType::GenericHttps as i32,
            config_overrides: vec![ConfigOverrides {
                output_schema_id: "default".to_string(),
                condition_type: Some(ConditionType {
                    condition: Some(condition_type::Condition::MatchEntityType(
                        MatchEntityTypeCondition {
                            entity_type: "alerts".to_string(),
                        },
                    )),
                }),
                message_config: Some(MessageConfig {
                    fields: vec![
                        MessageConfigField {
                            field_name: "headers".to_string(),
                            template: "{ \"example\": \"header_value\" }".to_string(),
                        },
                        MessageConfigField {
                            field_name: "body".to_string(),
                            template: "{ \"example\": \"body_value\" }".to_string(),
                        },
                    ],
                }),
            }],
            create_time: None,
            entity_type: "alerts".to_string(),
            id: None,
            user_facing_id: None,
            update_time: None,
            parent: None,
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
    #[ignore = "Presets are not yet supported"]
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
}
