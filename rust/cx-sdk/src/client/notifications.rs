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

use std::str::FromStr;

use cx_api::proto::com::coralogixapis::notification_center::{
    connectors::v1::{
        BatchGetConnectorsRequest,
        BatchGetConnectorsResponse,
        CreateConnectorRequest,
        CreateConnectorResponse,
        DeleteConnectorRequest,
        DeleteConnectorResponse,
        GetConnectorRequest,
        GetConnectorResponse,
        GetConnectorTypeSummariesRequest,
        GetConnectorTypeSummariesResponse,
        ListConnectorsRequest,
        ListConnectorsResponse,
        ReplaceConnectorRequest,
        ReplaceConnectorResponse,
        connectors_service_client::ConnectorsServiceClient,
    },
    notifications::v1::{
        TestConnectorConfigRequest,
        TestConnectorConfigResponse,
        TestExistingConnectorRequest,
        TestExistingConnectorResponse,
        TestPresetConfigRequest,
        TestPresetConfigResponse,
        TestTemplateRenderRequest,
        TestTemplateRenderResponse,
        testing_service_client::TestingServiceClient,
    },
    presets::v1::{
        BatchGetPresetsRequest,
        BatchGetPresetsResponse,
        CreateCustomPresetRequest,
        CreateCustomPresetResponse,
        DeleteCustomPresetRequest,
        DeleteCustomPresetResponse,
        GetDefaultPresetSummaryRequest,
        GetDefaultPresetSummaryResponse,
        GetPresetRequest,
        GetPresetResponse,
        GetSystemDefaultPresetSummaryRequest,
        GetSystemDefaultPresetSummaryResponse,
        ListPresetSummariesRequest,
        ListPresetSummariesResponse,
        PresetIdentifier,
        ReplaceCustomPresetRequest,
        ReplaceCustomPresetResponse,
        SetPresetAsDefaultRequest,
        SetPresetAsDefaultResponse,
        preset_identifier,
        presets_service_client::PresetsServiceClient,
    },
    routers::v1::{
        BatchGetGlobalRoutersRequest,
        BatchGetGlobalRoutersResponse,
        CreateGlobalRouterRequest,
        CreateGlobalRouterResponse,
        CreateOrReplaceGlobalRouterRequest,
        CreateOrReplaceGlobalRouterResponse,
        DeleteGlobalRouterRequest,
        DeleteGlobalRouterResponse,
        GetGlobalRouterRequest,
        GetGlobalRouterResponse,
        ListGlobalRoutersRequest,
        ListGlobalRoutersResponse,
        ReplaceGlobalRouterRequest,
        ReplaceGlobalRouterResponse,
        global_routers_service_client::GlobalRoutersServiceClient,
    },
};

pub use cx_api::proto::com::coralogixapis::notification_center::{
    ConditionType,
    ConfigOverrides,
    ConnectorConfigField,
    ConnectorType,
    MatchEntityTypeCondition,
    MessageConfig,
    MessageConfigField,
    OrderBy,
    condition_type,
    connectors::v1::{
        Connector,
        ConnectorConfig,
    },
    presets::v1::Preset,
    routers::v1::GlobalRouter,
    routing::{
        GlobalRouterIdentifier,
        global_router_identifier,
    },
};

use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{
        Channel,
        ClientTlsConfig,
        Endpoint,
    },
};

use crate::{
    CoralogixRegion,
    auth::AuthContext,
    error::{
        Result,
        SdkApiError,
        SdkError,
    },
    metadata::CallProperties,
    util::make_request_with_metadata,
};

const NOTIFICATIONS_FEATURE_GROUP_ID: &str = "notifications";

/// Parameters for testing a preset configuration.
///
/// This struct groups the parameters required for the `test_preset_config` function.
pub struct TestPresetConfigParams {
    /// The entity type to test with.
    pub entity_type: String,
    /// The entity sub-type to test with (optional).
    pub entity_sub_type: Option<String>,
    /// The ID of the connector to test with.
    pub connector_id: String,
    /// The ID of the output schema to test with.
    pub output_schema_id: String,
    /// The message configuration fields to use in the test.
    pub message_config_fields: Vec<MessageConfigField>,
    /// The ID of the preset to test with.
    pub preset_user_facing_id: String,
    /// The configuration overrides to apply during the test.
    pub config_overrides: Vec<ConfigOverrides>,
}

/// A client for interacting with the Coralogix Notification Center APIs.
pub struct NotificationsClient {
    metadata_map: MetadataMap,
    connectors_client: Mutex<ConnectorsServiceClient<Channel>>,
    presets_client: Mutex<PresetsServiceClient<Channel>>,
    global_routers_client: Mutex<GlobalRoutersServiceClient<Channel>>,
    testing_client: Mutex<TestingServiceClient<Channel>>,
}

impl NotificationsClient {
    /// Creates a new client for the Notification Center.
    /// # Arguments
    /// * `auth_context` - The [`AuthContext`] to use for authentication.
    /// * `region` - The [`CoralogixRegion`] to connect to.
    pub fn new(auth_context: AuthContext, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.grpc_endpoint().as_str())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();
        let request_metadata: CallProperties = (&auth_context.team_level_api_key).into();
        Ok(Self {
            metadata_map: request_metadata.to_metadata_map(),
            connectors_client: Mutex::new(ConnectorsServiceClient::new(channel.clone())),
            presets_client: Mutex::new(PresetsServiceClient::new(channel.clone())),
            global_routers_client: Mutex::new(GlobalRoutersServiceClient::new(channel.clone())),
            testing_client: Mutex::new(TestingServiceClient::new(channel)),
        })
    }

    /// Get a connector by ID.
    /// # Arguments
    /// * `connector_id` - The ID of the connector to get.
    pub async fn get_connector(&self, connector_id: String) -> Result<GetConnectorResponse> {
        let request = make_request_with_metadata(
            GetConnectorRequest { id: connector_id },
            &self.metadata_map,
        );
        {
            let mut client = self.connectors_client.lock().await.clone();

            client
                .get_connector(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/GetConnector".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// List all connectors.
    /// # Arguments
    /// * `connector_type` - The [`ConnectorType`] to filter by.
    /// * `order_bys` - The [`OrderBy`]s to use for sorting.
    /// * `entity_type` - The entity type to filter by.
    pub async fn list_connectors(
        &self,
        connector_type: ConnectorType,
        order_bys: Vec<OrderBy>,
    ) -> Result<ListConnectorsResponse> {
        let request = make_request_with_metadata(
            ListConnectorsRequest {
                connector_type: connector_type as i32,
                order_bys,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.connectors_client.lock().await.clone();

            client
                .list_connectors(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/ListConnectors".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Create a new connector.
    /// # Arguments
    /// * `connector` - The [`Connector`] to create.
    pub async fn create_connector(&self, connector: Connector) -> Result<CreateConnectorResponse> {
        let request = make_request_with_metadata(
            CreateConnectorRequest {
                connector: Some(connector),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.connectors_client.lock().await.clone();

            client
                .create_connector(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/CreateConnector".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Replace an existing connector.
    /// # Arguments
    /// * `connector` - The [`Connector`] to replace.
    pub async fn replace_connector(
        &self,
        connector: Connector,
    ) -> Result<ReplaceConnectorResponse> {
        let request = make_request_with_metadata(
            ReplaceConnectorRequest {
                connector: Some(connector),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.connectors_client.lock().await.clone();

            client
                .replace_connector(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/ReplaceConnector".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Delete a connector by ID.
    /// # Arguments
    /// * `connector_id` - The ID of the connector to delete.
    pub async fn delete_connector(&self, connector_id: String) -> Result<DeleteConnectorResponse> {
        let request = make_request_with_metadata(
            DeleteConnectorRequest { id: connector_id },
            &self.metadata_map,
        );
        {
            let mut client = self.connectors_client.lock().await.clone();

            client
                .delete_connector(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/DeleteConnector".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Batch get connectors by ID.
    /// # Arguments
    /// * `connector_ids` - The IDs of the connectors to get.
    pub async fn batch_get_connectors(
        &self,
        connector_ids: Vec<String>,
    ) -> Result<BatchGetConnectorsResponse> {
        let request = make_request_with_metadata(
            BatchGetConnectorsRequest { ids: connector_ids },
            &self.metadata_map,
        );
        {
            let mut client = self.connectors_client.lock().await.clone();

            client
                .batch_get_connectors(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/BatchGetConnectors".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Get summaries of connector types.
    pub async fn get_connector_type_summaries(&self) -> Result<GetConnectorTypeSummariesResponse> {
        let request =
            make_request_with_metadata(GetConnectorTypeSummariesRequest {}, &self.metadata_map);
        {
            let mut client = self.connectors_client.lock().await.clone();

            client
                .get_connector_type_summaries(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.connectors.v1.ConnectorsService/GetConnectorTypeSummaries".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Create a new custom preset.
    /// # Arguments
    /// * `preset` - The [`Preset`] to create.
    pub async fn create_custom_preset(&self, preset: Preset) -> Result<CreateCustomPresetResponse> {
        let request = make_request_with_metadata(
            CreateCustomPresetRequest {
                preset: Some(preset),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.presets_client.lock().await.clone();

            client
                .create_custom_preset(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.presets.v1.PresetsService/CreateCustomPreset".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Replace an existing custom preset.
    /// # Arguments
    /// * `preset` - The [`Preset`] to replace.
    pub async fn replace_custom_preset(
        &self,
        preset: Preset,
    ) -> Result<ReplaceCustomPresetResponse> {
        let request = make_request_with_metadata(
            ReplaceCustomPresetRequest {
                preset: Some(preset),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.presets_client.lock().await.clone();

            client
                .replace_custom_preset(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.presets.v1.PresetsService/ReplaceCustomPreset".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Delete a custom preset by ID.
    /// # Arguments
    /// * `preset_id` - The ID of the preset to delete.
    pub async fn delete_custom_preset(
        &self,
        preset_user_facing_id: String,
    ) -> Result<DeleteCustomPresetResponse> {
        let request = make_request_with_metadata(
            DeleteCustomPresetRequest {
                identifier: Some(PresetIdentifier {
                    value: Some(preset_identifier::Value::UserFacingId(
                        preset_user_facing_id,
                    )),
                }),
                deprecated_identifier: None,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.presets_client.lock().await.clone();

            client
                .delete_custom_preset(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.presets.v1.PresetsService/DeleteCustomPreset".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Set a preset as the default.
    /// # Arguments
    /// * `preset_user_facing_id` - The user-facing ID of the preset to set as the default.
    pub async fn set_preset_as_default(
        &self,
        preset_user_facing_id: String,
    ) -> Result<SetPresetAsDefaultResponse> {
        let request = make_request_with_metadata(
            SetPresetAsDefaultRequest {
                identifier: Some(PresetIdentifier {
                    value: Some(preset_identifier::Value::UserFacingId(
                        preset_user_facing_id,
                    )),
                }),
            },
            &self.metadata_map,
        );

        let mut client = self.presets_client.lock().await.clone();

        client
            .set_preset_as_default(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.notification_center.presets.v1.PresetsService/SetPresetAsDefault".into(),
                    feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into(),
                })
            })
    }

    /// Get a preset by ID.
    /// # Arguments
    /// * `preset_id` - The ID of the preset to get.
    pub async fn get_preset(&self, preset_user_facing_id: String) -> Result<GetPresetResponse> {
        let request = make_request_with_metadata(
            GetPresetRequest {
                identifier: Some(PresetIdentifier {
                    value: Some(preset_identifier::Value::UserFacingId(
                        preset_user_facing_id,
                    )),
                }),
                deprecated_identifier: None,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.presets_client.lock().await.clone();

            client
                .get_preset(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.presets.v1.PresetsService/GetPreset".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// List summaries of presets.
    /// # Arguments
    /// * `connector_type` - The [`ConnectorType`] to list presets for.
    /// * `entity_type` - The entity type to list presets for.
    /// * `order_bys` - The [`OrderBy`]s to use for sorting.
    pub async fn list_preset_summaries(
        &self,
        connector_type: ConnectorType,
        entity_type: String,
        order_bys: Vec<OrderBy>,
    ) -> Result<ListPresetSummariesResponse> {
        let request = make_request_with_metadata(
            ListPresetSummariesRequest {
                connector_type: connector_type as i32,
                entity_type,
                order_bys,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.presets_client.lock().await.clone();

            client
                .list_preset_summaries(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.presets.v1.PresetsService/ListPresetSummaries".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Batch get presets by ID.
    /// # Arguments
    /// * `preset_ids` - The IDs of the presets to get.
    pub async fn batch_get_presets(
        &self,
        preset_ids: Vec<String>,
    ) -> Result<BatchGetPresetsResponse> {
        let request = make_request_with_metadata(
            BatchGetPresetsRequest { ids: preset_ids },
            &self.metadata_map,
        );
        {
            let mut client = self.presets_client.lock().await.clone();

            client
                .batch_get_presets(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.presets.v1.PresetsService/BatchGetPresets".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Get the default preset summary.
    /// # Arguments
    /// * `connector_type` - The [`ConnectorType`] to get the default preset summary for.
    /// * `entity_type` - The entity type to get the default preset summary for.
    pub async fn get_default_preset_summary(
        &self,
        connector_type: ConnectorType,
        entity_type: String,
    ) -> Result<GetDefaultPresetSummaryResponse> {
        let request = make_request_with_metadata(
            GetDefaultPresetSummaryRequest {
                connector_type: connector_type as i32,
                entity_type,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.presets_client.lock().await.clone();

            client
                .get_default_preset_summary(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.presets.v1.PresetsService/GetDefaultPresetSummary".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Get the system default preset summary.
    /// # Arguments
    /// * `connector_type` - The [`ConnectorType`] to get the system default preset summary for.
    /// * `entity_type` - The entity type to get the system default preset summary for.
    pub async fn get_system_default_preset_summary(
        &self,
        connector_type: ConnectorType,
        entity_type: String,
    ) -> Result<GetSystemDefaultPresetSummaryResponse> {
        let request = make_request_with_metadata(
            GetSystemDefaultPresetSummaryRequest {
                connector_type: connector_type as i32,
                entity_type,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.presets_client.lock().await.clone();

            client
                .get_system_default_preset_summary(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.presets.v1.PresetsService/GetSystemDefaultPresetSummary".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Create a new global router.
    /// # Arguments
    /// * `global_router` - The [`GlobalRouter`] to create.
    pub async fn create_global_router(
        &self,
        global_router: GlobalRouter,
    ) -> Result<CreateGlobalRouterResponse> {
        let request = make_request_with_metadata(
            CreateGlobalRouterRequest {
                router: Some(global_router),
            },
            &self.metadata_map,
        );
        let mut client = self.global_routers_client.lock().await.clone();

        client
            .create_global_router(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogixapis.notification_center.routers.v1.GlobalRoutersService/CreateGlobalRouter".into(),
                feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into(),
            }))
    }

    /// Create or replace a global router.
    /// # Arguments
    /// * `global_router` - The [`GlobalRouter`] to create or replace.
    pub async fn create_or_replace_global_router(
        &self,
        global_router: GlobalRouter,
    ) -> Result<CreateOrReplaceGlobalRouterResponse> {
        let request = make_request_with_metadata(
            CreateOrReplaceGlobalRouterRequest {
                router: Some(global_router),
            },
            &self.metadata_map,
        );
        let mut client = self.global_routers_client.lock().await.clone();

        client
            .create_or_replace_global_router(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogixapis.notification_center.routers.v1.GlobalRoutersService/CreateOrReplaceGlobalRouter".into(),
                feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into(),
            }))
    }

    /// Replace an existing global router.
    /// # Arguments
    /// * `global_router` - The [`GlobalRouter`] to replace.
    pub async fn replace_global_router(
        &self,
        global_router: GlobalRouter,
    ) -> Result<ReplaceGlobalRouterResponse> {
        let request = make_request_with_metadata(
            ReplaceGlobalRouterRequest {
                router: Some(global_router),
            },
            &self.metadata_map,
        );
        let mut client = self.global_routers_client.lock().await.clone();

        client
            .replace_global_router(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogixapis.notification_center.routers.v1.GlobalRoutersService/ReplaceGlobalRouter".into(),
                feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into(),
            }))
    }

    /// Delete a global router by identifier.
    /// # Arguments
    /// * `global_router_identifier` - The identifier of the global router to delete.
    pub async fn delete_global_router(
        &self,
        global_router_identifier: GlobalRouterIdentifier,
    ) -> Result<DeleteGlobalRouterResponse> {
        let request = make_request_with_metadata(
            DeleteGlobalRouterRequest {
                identifier: Some(global_router_identifier),
            },
            &self.metadata_map,
        );
        let mut client = self.global_routers_client.lock().await.clone();

        client
            .delete_global_router(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogixapis.notification_center.routers.v1.GlobalRoutersService/DeleteGlobalRouter".into(),
                feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into(),
            }))
    }

    /// Get a global router by identifier.
    /// # Arguments
    /// * `global_router_identifier` - The identifier of the global router to get.
    pub async fn get_global_router(
        &self,
        global_router_identifier: GlobalRouterIdentifier,
    ) -> Result<GetGlobalRouterResponse> {
        let request = make_request_with_metadata(
            GetGlobalRouterRequest {
                identifier: Some(global_router_identifier),
            },
            &self.metadata_map,
        );
        let mut client = self.global_routers_client.lock().await.clone();

        client
            .get_global_router(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogixapis.notification_center.routers.v1.GlobalRoutersService/GetGlobalRouter".into(),
                feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into(),
            }))
    }

    /// List all global routers.
    /// # Arguments
    /// * `entity_type` - The entity type to filter global routers by.
    pub async fn list_global_routers(
        &self,
        entity_type: String,
    ) -> Result<ListGlobalRoutersResponse> {
        let request = make_request_with_metadata(
            ListGlobalRoutersRequest {
                entity_type: Some(entity_type),
            },
            &self.metadata_map,
        );
        let mut client = self.global_routers_client.lock().await.clone();

        client
            .list_global_routers(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogixapis.notification_center.routers.v1.GlobalRoutersService/ListGlobalRouters".into(),
                feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into(),
            }))
    }

    /// Batch get global routers by identifiers.
    /// # Arguments
    /// * `global_router_identifiers` - The identifiers of the global routers to retrieve.
    pub async fn batch_get_global_routers(
        &self,
        global_router_identifiers: Vec<GlobalRouterIdentifier>,
    ) -> Result<BatchGetGlobalRoutersResponse> {
        let request = make_request_with_metadata(
            BatchGetGlobalRoutersRequest {
                ids: global_router_identifiers,
            },
            &self.metadata_map,
        );
        let mut client = self.global_routers_client.lock().await.clone();

        client
            .batch_get_global_routers(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogixapis.notification_center.routers.v1.GlobalRoutersService/BatchGetGlobalRouters".into(),
                feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into(),
            }))
    }

    /// Test a connector configuration.
    /// # Arguments
    /// * `connector_type` - The [`ConnectorType`] to test.
    /// * `output_schema_id` - The ID of the output schema to test with.
    /// * `connector_config_fields` - The [ConnectorConfigField]s to test with.
    /// * `entity_type` - The entity type to test with.
    pub async fn test_connector_config(
        &self,
        connector_type: ConnectorType,
        output_schema_id: String,
        connector_config_fields: Vec<ConnectorConfigField>,
        entity_type: String,
    ) -> Result<TestConnectorConfigResponse> {
        let request = make_request_with_metadata(
            TestConnectorConfigRequest {
                r#type: connector_type as i32,
                output_schema_id,
                fields: connector_config_fields,
                entity_type: Some(entity_type),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.testing_client.lock().await.clone();

            client
                .test_connector_config(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.notifications.v1.TestingService/TestConnectorConfig".into(),
                                                feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Test an existing connector.
    /// # Arguments
    /// * `connector_id` - The ID of the connector to test.
    /// * `output_schema_id` - The ID of the output schema to test with.
    pub async fn test_existing_connector(
        &self,
        connector_id: String,
        output_schema_id: String,
    ) -> Result<TestExistingConnectorResponse> {
        let request = make_request_with_metadata(
            TestExistingConnectorRequest {
                connector_id,
                output_schema_id,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.testing_client.lock().await.clone();

            client
                .test_existing_connector(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.notifications.v1.TestingService/TestExistingConnector".into(),
                                                feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Test a preset configuration.
    /// # Arguments
    /// * `params` - The parameters for testing the preset configuration.
    pub async fn test_preset_config(
        &self,
        params: TestPresetConfigParams,
    ) -> Result<TestPresetConfigResponse> {
        let request = make_request_with_metadata(
            TestPresetConfigRequest {
                entity_type: params.entity_type,
                entity_sub_type: params.entity_sub_type,
                connector_id: params.connector_id,
                preset_id: params.preset_user_facing_id,
                config_overrides: params.config_overrides,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.testing_client.lock().await.clone();

            client
                .test_preset_config(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.notifications.v1.TestingService/TestPresetConfig".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }

    /// Test rendering a template.
    /// # Arguments
    /// * `entity_type` - The entity type to test with.
    /// * `entity_sub_type` - The entity sub type to test with.
    /// * `template` - The template to render.
    pub async fn test_template_render(
        &self,
        entity_type: String,
        entity_sub_type: Option<String>,
        template: String,
    ) -> Result<TestTemplateRenderResponse> {
        let request = make_request_with_metadata(
            TestTemplateRenderRequest {
                entity_type,
                entity_sub_type,
                template,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.testing_client.lock().await.clone();

            client
                .test_template_render(request)
                .await
                .map(|r| r.into_inner())
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.notification_center.notifications.v1.TestingService/TestTemplateRender".into(),
                        feature_group: NOTIFICATIONS_FEATURE_GROUP_ID.into()
                    },
                ))
        }
    }
}
