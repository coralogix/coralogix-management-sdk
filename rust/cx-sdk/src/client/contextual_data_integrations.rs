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

use std::{
    str::FromStr,
    sync::Arc,
};

use tokio::sync::Mutex;
use tonic::transport::{
    Channel,
    ClientTlsConfig,
    Endpoint,
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

const CONTEXTUAL_DATA_INTEGRATIONS_FEATURE_GROUP_ID: &str = "contextual_data_integrations";

use cx_api::proto::com::coralogix::integrations::v1::{
    DeleteContextualDataIntegrationRequest,
    DeleteContextualDataIntegrationResponse,
    GetContextualDataIntegrationDefinitionRequest,
    GetContextualDataIntegrationDefinitionResponse,
    GetContextualDataIntegrationDetailsRequest,
    GetContextualDataIntegrationDetailsResponse,
    GetContextualDataIntegrationsRequest,
    GetContextualDataIntegrationsResponse,
    SaveContextualDataIntegrationRequest,
    SaveContextualDataIntegrationResponse,
    TestContextualDataIntegrationRequest,
    TestContextualDataIntegrationResponse,
    UpdateContextualDataIntegrationRequest,
    UpdateContextualDataIntegrationResponse,
    contextual_data_integration_service_client::ContextualDataIntegrationServiceClient,
};

pub use cx_api::proto::com::coralogix::integrations::v1::{
    GenericIntegrationParameters,
    IntegrationMetadata,
    integration_metadata::SpecificData,
};

/// Client for managing Contextual Data Integrations.
pub struct ContextualDataIntegrationsClient {
    metadata_map: tonic::metadata::MetadataMap,
    service_client: Arc<Mutex<ContextualDataIntegrationServiceClient<Channel>>>,
}

impl ContextualDataIntegrationsClient {
    /// Creates a new client for the Contextual Data Integrations API.
    pub fn new(region: CoralogixRegion, auth_context: AuthContext) -> Result<Self> {
        let channel = Endpoint::from_str(&region.grpc_endpoint())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();

        let request_metadata: CallProperties = (&auth_context.team_level_api_key).into();
        Ok(Self {
            metadata_map: request_metadata.to_metadata_map(),
            service_client: Arc::new(Mutex::new(ContextualDataIntegrationServiceClient::new(
                channel,
            ))),
        })
    }

    /// Get all contextual data integrations
    ///
    /// # Arguments
    /// * `include_testing_integrations` - Whether to include testing integrations.
    pub async fn list(
        &self,
        include_testing_integrations: bool,
    ) -> Result<GetContextualDataIntegrationsResponse> {
        let request = make_request_with_metadata(
            GetContextualDataIntegrationsRequest {
                include_testing_integrations: Some(include_testing_integrations),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_contextual_data_integrations(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogix.integrations.v1.ContextualDataIntegrationService/GetContextualDataIntegrations".into(),
                feature_group: CONTEXTUAL_DATA_INTEGRATIONS_FEATURE_GROUP_ID.into(),
            }))
    }

    /// Get contextual data integration definition
    ///
    /// # Arguments
    /// * `id` - The id of the integration to get the definition for.
    /// * `include_testing_integrations` - Whether to include testing integrations.
    pub async fn get_definition(
        &self,
        id: String,
        include_testing_integrations: bool,
    ) -> Result<GetContextualDataIntegrationDefinitionResponse> {
        let request = make_request_with_metadata(
            GetContextualDataIntegrationDefinitionRequest {
                id: Some(id),
                include_testing_integrations: Some(include_testing_integrations),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_contextual_data_integration_definition(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogix.integrations.v1.ContextualDataIntegrationService/GetContextualDataIntegrationDefinition".into(),
                feature_group: CONTEXTUAL_DATA_INTEGRATIONS_FEATURE_GROUP_ID.into(),
            }))
    }

    /// Get contextual data integration details
    ///
    /// # Arguments
    /// * `id` - The id of the integration to get the details for.
    /// * `include_testing_revisions` - Whether to include testing revisions.
    pub async fn get_details(
        &self,
        id: String,
        include_testing_revisions: bool,
    ) -> Result<GetContextualDataIntegrationDetailsResponse> {
        let request = make_request_with_metadata(
            GetContextualDataIntegrationDetailsRequest {
                id: Some(id),
                include_testing_revisions: Some(include_testing_revisions),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_contextual_data_integration_details(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogix.integrations.v1.ContextualDataIntegrationService/GetContextualDataIntegrationDetails".into(),
                feature_group: CONTEXTUAL_DATA_INTEGRATIONS_FEATURE_GROUP_ID.into(),
            }))
    }

    /// Test contextual data integration
    pub async fn test(
        &self,
        integration_data: IntegrationMetadata,
        integration_id: String,
    ) -> Result<TestContextualDataIntegrationResponse> {
        let request = make_request_with_metadata(
            TestContextualDataIntegrationRequest {
                integration_data: Some(integration_data),
                integration_id: Some(integration_id),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .test_contextual_data_integration(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogix.integrations.v1.ContextualDataIntegrationService/TestContextualDataIntegration".into(),
                feature_group: CONTEXTUAL_DATA_INTEGRATIONS_FEATURE_GROUP_ID.into(),
            }))
    }

    /// Save contextual data integration
    ///
    /// # Arguments
    /// * `metadata` - The metadata of the integration to save.
    pub async fn create(
        &self,
        metadata: IntegrationMetadata,
    ) -> Result<SaveContextualDataIntegrationResponse> {
        let request = make_request_with_metadata(
            SaveContextualDataIntegrationRequest {
                metadata: Some(metadata),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .save_contextual_data_integration(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogix.integrations.v1.ContextualDataIntegrationService/SaveContextualDataIntegration".into(),
                feature_group: CONTEXTUAL_DATA_INTEGRATIONS_FEATURE_GROUP_ID.into(),
            }))
    }

    /// Update contextual data integration
    ///
    /// # Arguments
    /// * `integration_id` - The id of the integration to update.
    /// * `metadata` - The metadata of the integration to update.
    pub async fn update(
        &self,
        integration_id: String,
        metadata: IntegrationMetadata,
    ) -> Result<UpdateContextualDataIntegrationResponse> {
        let request = make_request_with_metadata(
            UpdateContextualDataIntegrationRequest {
                integration_id: Some(integration_id),
                metadata: Some(metadata),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .update_contextual_data_integration(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogix.integrations.v1.ContextualDataIntegrationService/UpdateContextualDataIntegration".into(),
                feature_group: CONTEXTUAL_DATA_INTEGRATIONS_FEATURE_GROUP_ID.into(),
            }))
    }

    /// Delete contextual data integration
    pub async fn delete(
        &self,
        integration_id: String,
    ) -> Result<DeleteContextualDataIntegrationResponse> {
        let request = make_request_with_metadata(
            DeleteContextualDataIntegrationRequest {
                integration_id: Some(integration_id),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .delete_contextual_data_integration(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| SdkError::ApiError(SdkApiError {
                status,
                endpoint: "/com.coralogix.integrations.v1.ContextualDataIntegrationService/DeleteContextualDataIntegration".into(),
                feature_group: CONTEXTUAL_DATA_INTEGRATIONS_FEATURE_GROUP_ID.into(),
            }))
    }
}
