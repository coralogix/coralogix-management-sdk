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
        SdkError,
    },
    metadata::CallProperties,
    util::make_request_with_metadata,
};

use cx_api::proto::com::coralogix::integrations::v1::{
    DeleteIntegrationRequest,
    DeleteIntegrationResponse,
    GenericIntegrationParameters,
    GetDeployedIntegrationRequest,
    GetDeployedIntegrationResponse,
    GetIntegrationDefinitionRequest,
    GetIntegrationDefinitionResponse,
    GetIntegrationDetailsRequest,
    GetIntegrationDetailsResponse,
    GetIntegrationsRequest,
    GetIntegrationsResponse,
    GetManagedIntegrationStatusRequest,
    GetManagedIntegrationStatusResponse,
    GetRumApplicationVersionDataRequest,
    GetRumApplicationVersionDataResponse,
    GetTemplateRequest,
    GetTemplateResponse,
    IntegrationMetadata,
    SaveIntegrationRequest,
    SaveIntegrationResponse,
    SyncRumDataRequest,
    SyncRumDataResponse,
    TestIntegrationRequest,
    TestIntegrationResponse,
    UpdateIntegrationRequest,
    UpdateIntegrationResponse,
    integration_metadata::SpecificData,
    integration_service_client::IntegrationServiceClient,
};

pub use cx_api::proto::com::coralogix::integrations::v1::{
    Parameter,
    integration_details::default_integration_details::*,
    integration_details::*,
    parameter::*,
    test_integration_result::Result as TestResult,
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

use crate::CoralogixRegion;

/// The Integration API client.
/// Read more at <https://coralogix.com/docs/user-team-management/>
pub struct IntegrationsClient {
    metadata_map: MetadataMap,
    service_client: Mutex<IntegrationServiceClient<Channel>>,
}

impl IntegrationsClient {
    /// Creates a new client for the Integrations API.
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
            metadata_map: request_metadata.to_metadata_map(),
            service_client: Mutex::new(IntegrationServiceClient::new(channel)),
        })
    }

    /// Creates a new Integration in the organization.
    ///
    /// # Arguments
    /// * `integration_key` - The key of the integration to update.
    /// * `version` - The version of the integration to update.
    /// * `parameters` - The parameters of the integration to update.
    pub async fn create(
        &self,
        integration_key: String,
        version: Option<String>,
        parameters: Option<Vec<Parameter>>,
    ) -> Result<SaveIntegrationResponse> {
        let request = make_request_with_metadata(
            SaveIntegrationRequest {
                metadata: Some(IntegrationMetadata {
                    integration_key: Some(integration_key),
                    version,
                    specific_data: parameters.map(|p| {
                        SpecificData::IntegrationParameters(GenericIntegrationParameters {
                            parameters: p,
                        })
                    }),
                }),
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .save_integration(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint:
                        "/com.coralogixapis.integrations.v1.IntegrationService/SaveIntegration"
                            .into(),
                })
            })
    }

    /// Update the Integration identified by its id.
    ///
    /// # Arguments
    /// * `id` - The id of the integration to update.
    /// * `integration_key` - The key of the integration to update.
    /// * `version` - The version of the integration to update.
    /// * `parameters` - The parameters of the integration to update.
    pub async fn update(
        &self,
        id: String,
        integration_key: String,
        version: Option<String>,
        parameters: Option<Vec<Parameter>>,
    ) -> Result<UpdateIntegrationResponse> {
        let request = make_request_with_metadata(
            UpdateIntegrationRequest {
                id: Some(id),
                metadata: Some(IntegrationMetadata {
                    integration_key: Some(integration_key),
                    version,
                    specific_data: parameters.map(|p| {
                        SpecificData::IntegrationParameters(GenericIntegrationParameters {
                            parameters: p,
                        })
                    }),
                }),
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .update_integration(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint:
                        "/com.coralogixapis.integrations.v1.IntegrationService/UpdateIntegration"
                            .into(),
                })
            })
    }

    /// Deletes the Integration identified by its id.
    ///
    /// # Arguments
    /// * `id` - The id of the integration to delete.
    pub async fn delete(&self, id: String) -> Result<DeleteIntegrationResponse> {
        let request = make_request_with_metadata(
            DeleteIntegrationRequest {
                integration_id: Some(id),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .delete_integration(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint:
                        "/com.coralogixapis.integrations.v1.IntegrationService/DeleteIntegration"
                            .into(),
                })
            })
    }

    /// Retrieves the Integration identified by its id.
    ///
    /// # Arguments
    /// * `id` - The id of the integration to retrieve.
    /// * `include_testing_revision` - Whether to include the testing revision.
    pub async fn get_details(
        &self,
        id: String,
        include_testing_revision: bool,
    ) -> Result<GetIntegrationDetailsResponse> {
        let request = make_request_with_metadata(
            GetIntegrationDetailsRequest {
                id: Some(id),
                include_testing_revision: Some(include_testing_revision),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_integration_details(request)
            .await
            .map(|r| r.into_inner())
            .map_err(
                |status| SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.integrations.v1.IntegrationService/GetIntegrationDetails"
                        .into(),
                }),
            )
    }

    /// Retrieves the Deployed Integration identified by its id.
    ///
    /// # Arguments
    /// * `id` - The id of the deployed integration to retrieve.
    pub async fn get(&self, id: String) -> Result<GetDeployedIntegrationResponse> {
        let request = make_request_with_metadata(
            GetDeployedIntegrationRequest {
                integration_id: Some(id),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_deployed_integration(request)
            .await
            .map(|r| r.into_inner())
            .map_err(
                |status| SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.integrations.v1.IntegrationService/GetDeployedIntegration"
                        .into(),
                }),
            )
    }

    /// Retrieves the Integration definition identified by its id.
    ///
    /// # Arguments
    /// * `id` - The id of the integration to retrieve.
    /// * `include_testing_revision` - Whether to include the testing revision.
    pub async fn get_definitions(
        &self,
        id: String,
        include_testing_revision: bool,
    ) -> Result<GetIntegrationDefinitionResponse> {
        let request = make_request_with_metadata(
            GetIntegrationDefinitionRequest {
                id: Some(id),
                include_testing_revision: Some(include_testing_revision),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_integration_definition(request)
            .await
            .map(|r| r.into_inner())
            .map_err(
                |status| SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.integrations.v1.IntegrationService/GetIntegrationDefinition"
                        .into(),
                }),
            )
    }

    /// Retrieves the Integration status identified by its id.
    ///
    /// # Arguments
    /// * `id` - The id of the integration to retrieve.
    pub async fn get_integration_status(
        &self,
        id: String,
    ) -> Result<GetManagedIntegrationStatusResponse> {
        let request = make_request_with_metadata(
            GetManagedIntegrationStatusRequest { integration_id: id },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_managed_integration_status(request)
            .await
            .map(|r| r.into_inner())
            .map_err(
                |status| SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.integrations.v1.IntegrationService/GetManagedIntegrationStatus"
                        .into(),
                }),
            )
    }

    /// Retrieves the Integration template identified by its id.
    ///
    /// # Arguments
    /// * `id` - The id of the integration to retrieve.
    pub async fn get_template(&self, id: String) -> Result<GetTemplateResponse> {
        let request = make_request_with_metadata(
            GetTemplateRequest {
                integration_id: Some(id),
                extra_params: Default::default(), // TODO: expose these parameters somehow?
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_template(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.integrations.v1.IntegrationService/GetTemplate"
                        .into(),
                })
            })
    }

    /// Retrieves the RUM application version data.
    ///
    /// # Arguments
    /// * `application_name` - The name of the application.
    pub async fn get_rum_application_version_data(
        &self,
        application_name: String,
    ) -> Result<GetRumApplicationVersionDataResponse> {
        let request = make_request_with_metadata(
            GetRumApplicationVersionDataRequest {
                application_name: Some(application_name),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_rum_application_version_data(request)
            .await
            .map(|r| r.into_inner())
            .map_err(
                |status| SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.integrations.v1.IntegrationService/GetRumApplicationVersionData"
                        .into(),
                }),
            )
    }

    /// Synchronizes the RUM data.
    ///
    /// # Arguments
    /// * `force` - Whether to force the synchronization.
    pub async fn sync_rum_data(&self, force: bool) -> Result<SyncRumDataResponse> {
        let request = make_request_with_metadata(
            SyncRumDataRequest { force: Some(force) },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .sync_rum_data(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.integrations.v1.IntegrationService/SyncRumData"
                        .into(),
                })
            })
    }

    /// Tests the Integration identified by its id.
    ///
    /// # Arguments
    /// * `id` - The id of the integration to test.
    /// * `integration_key` - The key of the integration to test.
    /// * `version` - The version of the integration to test.
    /// * `parameters` - The parameters of the integration to test.
    pub async fn test_integration(
        &self,
        integration_id: Option<String>,
        integration_key: String,
        version: Option<String>,
        parameters: Option<Vec<Parameter>>,
    ) -> Result<TestIntegrationResponse> {
        let request = make_request_with_metadata(
            TestIntegrationRequest {
                integration_id,
                integration_data: Some(IntegrationMetadata {
                    integration_key: Some(integration_key),
                    version,
                    specific_data: parameters.map(|p| {
                        SpecificData::IntegrationParameters(GenericIntegrationParameters {
                            parameters: p,
                        })
                    }),
                }),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .test_integration(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint:
                        "/com.coralogixapis.integrations.v1.IntegrationService/TestIntegration"
                            .into(),
                })
            })
    }

    /// Retrieves a list of all Integrations.
    ///     
    /// # Returns
    /// A list of all Integrations.
    pub async fn list(&self, include_testing_revision: bool) -> Result<GetIntegrationsResponse> {
        let request = make_request_with_metadata(
            GetIntegrationsRequest {
                include_testing_revision: Some(include_testing_revision),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_integrations(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint:
                        "/com.coralogixapis.integrations.v1.IntegrationService/GetIntegrations"
                            .into(),
                })
            })
    }
}
