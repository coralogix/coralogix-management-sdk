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

use crate::{
    auth::AuthContext,
    error::{
        Result,
        SdkApiError,
    },
    metadata::CallProperties,
    util::make_request_with_metadata,
};

pub use crate::com::coralogixapis::actions::v2::{
    Action,
    SourceType,
};

use cx_api::proto::com::coralogix::extensions::v1::{
    DeployExtensionRequest,
    DeployExtensionResponse,
    ExtensionDeployment,
    GetAllExtensionsRequest,
    GetAllExtensionsResponse,
    GetDeployedExtensionsRequest,
    GetDeployedExtensionsResponse,
    GetExtensionRequest,
    GetExtensionResponse,
    UndeployExtensionRequest,
    UndeployExtensionResponse,
    UpdateExtensionRequest,
    UpdateExtensionResponse,
    extension_deployment_service_client::ExtensionDeploymentServiceClient,
    extension_service_client::ExtensionServiceClient,
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

pub use crate::com::coralogix::extensions::v1::{
    Extension,
    TargetDomain,
    get_all_extensions_request::Filter as ExtensionFilter,
};

use crate::CoralogixRegion;

const EXTENSIONS_FEATURE_GROUP_ID: &str = "extensions";

/// The Actions API client.
/// Read more at <https://coralogix.com/docs/coralogix-action-extension/>
pub struct ExtensionsClient {
    teams_level_metadata_map: MetadataMap,
    extensions_service_client: Mutex<ExtensionServiceClient<Channel>>,
    extensions_deployment_service_client: Mutex<ExtensionDeploymentServiceClient<Channel>>,
}

impl ExtensionsClient {
    /// Creates a new client for the Extensions API.
    ///
    /// # Arguments
    /// * `auth_context` - The [`AuthContext`] to use for authentication.
    /// * `region` - The [`CoralogixRegion`] to connect to.
    pub fn new(auth_context: AuthContext, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.grpc_endpoint().as_str())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();
        let request_metadata: CallProperties = (&auth_context.team_level_api_key).into();
        Ok(Self {
            teams_level_metadata_map: request_metadata.to_metadata_map(),
            extensions_service_client: Mutex::new(
                ExtensionServiceClient::new(channel.clone())
                    .max_decoding_message_size(50 * 1024 * 1024)
                    .max_encoding_message_size(50 * 1024 * 1024),
            ),
            extensions_deployment_service_client: Mutex::new(
                ExtensionDeploymentServiceClient::new(channel)
                    .max_decoding_message_size(50 * 1024 * 1024)
                    .max_encoding_message_size(50 * 1024 * 1024),
            ),
        })
    }

    /// Get all extensions
    ///
    /// # Arguments
    /// * `filter` - The [`ExtensionFilter`] used to filter the extensions.
    /// * `include_hidden_extensions` - Whether to include hidden extensions.
    pub async fn get_all(
        &self,
        filter: Option<ExtensionFilter>,
        include_hidden_extensions: bool,
    ) -> Result<GetAllExtensionsResponse> {
        let request = make_request_with_metadata(
            GetAllExtensionsRequest {
                filter,
                include_hidden_extensions: Some(include_hidden_extensions),
            },
            &self.teams_level_metadata_map,
        );
        Ok(self
            .extensions_service_client
            .lock()
            .await
            .get_all_extensions(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogix.extensions.v1.ExtensionService/GetAllExtensions".into(),
                feature_group: EXTENSIONS_FEATURE_GROUP_ID.into(),
            })?
            .into_inner())
    }

    /// Get an extension by ID
    ///
    /// # Arguments
    /// * `id` - The ID of the extension to get.
    pub async fn get(&self, id: String) -> Result<GetExtensionResponse> {
        let request = make_request_with_metadata(
            GetExtensionRequest {
                id: Some(id),
                include_dashboard_binaries: None,
                include_testing_revision: None,
            },
            &self.teams_level_metadata_map,
        );
        //let client = self.service_client.lock().await.descri
        Ok(self
            .extensions_service_client
            .lock()
            .await
            .get_extension(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogix.extensions.v1.ExtensionService/GetExtension".into(),
                feature_group: EXTENSIONS_FEATURE_GROUP_ID.into(),
            })?
            .into_inner())
    }

    /// Get all deplyoyed extensions
    pub async fn get_deployed(&self) -> Result<GetDeployedExtensionsResponse> {
        let request = make_request_with_metadata(
            GetDeployedExtensionsRequest {},
            &self.teams_level_metadata_map,
        );
        Ok(self
            .extensions_deployment_service_client
            .lock()
            .await
            .get_deployed_extensions(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint:
                    "/com.coralogix.extensions.v1.ExtensionDeploymentService/GetDeployedExtensions"
                        .into(),
                feature_group: EXTENSIONS_FEATURE_GROUP_ID.into(),
            })?
            .into_inner())
    }

    /// Deploy an extension
    ///
    /// # Arguments
    /// * `id` - The ID of the extension to deploy.
    /// * `version` - The version of the extension to deploy.
    /// * `applications` - The applications to deploy the extension to.
    /// * `subsystems` - The subsystems to deploy the extension to.
    /// * `item_ids` - The item IDs to deploy the extension to.
    /// * `extension_deployment` - The extension deployment to deploy.
    /// * `extension_deployment` - The extension deployment to deploy.
    /// * `extension_deployment` - The extension deployment to deploy.
    pub async fn deploy(
        &self,
        id: String,
        version: String,
        applications: Vec<String>,
        subsystems: Vec<String>,
        item_ids: Vec<String>,
        extension_deployment: Option<ExtensionDeployment>,
    ) -> Result<DeployExtensionResponse> {
        let request = make_request_with_metadata(
            DeployExtensionRequest {
                id: Some(id),
                version: Some(version),
                item_ids,
                applications,
                subsystems,
                extension_deployment,
            },
            &self.teams_level_metadata_map,
        );
        Ok(self
            .extensions_deployment_service_client
            .lock()
            .await
            .deploy_extension(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogix.extensions.v1.ExtensionDeploymentService/DeployExtension"
                    .into(),
                feature_group: EXTENSIONS_FEATURE_GROUP_ID.into(),
            })?
            .into_inner())
    }

    /// Undeploy an extension
    ///
    /// # Arguments
    /// * `id` - The ID of the extension to undeploy.
    /// * `kept_extension_items` - The extension items to keep.
    pub async fn undeploy(
        &self,
        id: String,
        kept_extension_items: Vec<String>,
    ) -> Result<UndeployExtensionResponse> {
        let request = make_request_with_metadata(
            UndeployExtensionRequest {
                id: Some(id),
                kept_extension_items,
            },
            &self.teams_level_metadata_map,
        );
        Ok(self
            .extensions_deployment_service_client
            .lock()
            .await
            .undeploy_extension(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint:
                    "/com.coralogix.extensions.v1.ExtensionDeploymentService/UndeployExtension"
                        .into(),
                feature_group: EXTENSIONS_FEATURE_GROUP_ID.into(),
            })?
            .into_inner())
    }

    /// Update a deployed extension
    ///
    /// # Arguments
    /// * `id` - The ID of the extension to update.
    /// * `version` - The version of the extension to update.
    /// * `item_ids` - The item IDs to update the extension to.
    /// * `applications` - The applications to update the extension to.
    /// * `subsystems` - The subsystems to update the extension to.
    /// * `extension_deployment` - The extension deployment to update.
    pub async fn update(
        &self,
        id: String,
        version: String,
        item_ids: Vec<String>,
        applications: Vec<String>,
        subsystems: Vec<String>,
        extension_deployment: Option<ExtensionDeployment>,
    ) -> Result<UpdateExtensionResponse> {
        let request = make_request_with_metadata(
            UpdateExtensionRequest {
                id: Some(id),
                version: Some(version),
                item_ids,
                applications,
                subsystems,
                extension_deployment,
            },
            &self.teams_level_metadata_map,
        );
        Ok(self
            .extensions_deployment_service_client
            .lock()
            .await
            .update_extension(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint:
                    "/com.coralogix.extensions.v1.ExtensionDeploymentService/UpdateExtensions"
                        .into(),
                feature_group: EXTENSIONS_FEATURE_GROUP_ID.into(),
            })?
            .into_inner())
    }
}
