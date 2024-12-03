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

pub use crate::com::coralogixapis::actions::v2::Action;

use cx_api::proto::com::coralogix::enrichment::v1::{
    CreateCustomEnrichmentRequest,
    CreateCustomEnrichmentResponse,
    DeleteCustomEnrichmentRequest,
    DeleteCustomEnrichmentResponse,
    File,
    GetCustomEnrichmentRequest,
    GetCustomEnrichmentResponse,
    GetCustomEnrichmentsRequest,
    GetCustomEnrichmentsResponse,
    UpdateCustomEnrichmentRequest,
    UpdateCustomEnrichmentResponse,
    custom_enrichment_service_client::CustomEnrichmentServiceClient,
    file::Content,
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

pub use crate::com::coralogix::enrichment::v1::CustomEnrichment;

const DATASETS_FEATURE_GROUP_ID: &str = "logs";

/// The Custom Enrichments API client.
/// Read more at <https://coralogix.com/docs/custom-enrichment-api/>
pub struct DatasetClient {
    metadata_map: MetadataMap,
    service_client: Mutex<CustomEnrichmentServiceClient<Channel>>,
}

impl DatasetClient {
    /// Creates a new client for the Custom Enrichments API.
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
            service_client: Mutex::new(CustomEnrichmentServiceClient::new(channel)),
        })
    }

    /// Creates a new dataset.
    ///
    /// # Arguments
    /// * `name` - The name of the Custom Enrichment.
    /// * `description` - The description of the Custom Enrichment.
    /// * `data` - The data of the Custom Enrichment.
    pub async fn create(
        &self,
        name: String,
        description: Option<String>,
        data: Vec<u8>,
    ) -> Result<CreateCustomEnrichmentResponse> {
        let file = File {
            name: Some(name.clone()),
            content: Some(Content::Binary(data)),
            extension: Some("csv".into()),
        };
        let request = make_request_with_metadata(
            CreateCustomEnrichmentRequest {
                file: Some(file),
                name: Some(name.clone()),
                description,
            },
            &self.metadata_map,
        );
        {
            Ok(self
                .service_client
                .lock()
                .await
                .create_custom_enrichment(request)
                .await
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.enrichment.v1.CustomEnrichmentService/CreateCustomEnrichment".into(),
                        feature_group: DATASETS_FEATURE_GROUP_ID.into(),
                    }),
                )?
                .into_inner())
        }
    }

    /// Replaces the existing dataset identified by its id.
    ///
    /// # Arguments
    /// * `id` - The id of the Custom Enrichment.
    /// * `name` - The name of the Custom Enrichment.
    /// * `description` - The description of the Custom Enrichment.
    /// * `data` - The data of the Custom Enrichment.
    ///
    /// Note that the existing dataset will be replaced with a new version of the data.
    pub async fn replace(
        &self,
        id: u32,
        name: String,
        description: Option<String>,
        data: Vec<u8>,
    ) -> Result<UpdateCustomEnrichmentResponse> {
        let file = File {
            name: Some(name.clone()),
            content: Some(Content::Binary(data)),
            extension: Some("csv".into()),
        };
        let request = make_request_with_metadata(
            UpdateCustomEnrichmentRequest {
                custom_enrichment_id: Some(id),
                file: Some(file),
                name: Some(name.clone()),
                description,
            },
            &self.metadata_map,
        );
        {
            Ok(self
                .service_client
                .lock()
                .await
                .update_custom_enrichment(request)
                .await
                .map_err(
                    |status| SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.enrichment.v1.CustomEnrichmentService/UpdateCustomEnrichment".into(),
                        feature_group: DATASETS_FEATURE_GROUP_ID.into(),
                    }),
                )?
                .into_inner())
        }
    }

    /// Deletes the dataset identified by its id.
    ///
    /// # Arguments
    /// * `custom_enrichment_id` - The id of the Custom Enrichment.
    pub async fn delete(
        &self,
        custom_enrichment_id: u32,
    ) -> Result<DeleteCustomEnrichmentResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .delete_custom_enrichment(make_request_with_metadata(
                DeleteCustomEnrichmentRequest {
                    custom_enrichment_id: Some(custom_enrichment_id),
                },
                &self.metadata_map,
            ))
            .await
            .map_err(
                |status| SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.enrichment.v1.CustomEnrichmentService/DeleteCustomEnrichment".into(),
                    feature_group: DATASETS_FEATURE_GROUP_ID.into(),
                }),
            )?
            .into_inner())
    }

    /// Retrieves the dataset by id.
    ///
    /// # Arguments
    /// * `id` - The id of the Custom Enrichment.
    pub async fn get(&self, id: u32) -> Result<GetCustomEnrichmentResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_custom_enrichment(make_request_with_metadata(
                GetCustomEnrichmentRequest { id: Some(id) },
                &self.metadata_map,
            ))
            .await
            .map_err(
                |status| SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.enrichment.v1.CustomEnrichmentService/GetCustomEnrichment".into(),
                    feature_group: DATASETS_FEATURE_GROUP_ID.into(),
                }),
            )?
            .into_inner())
    }

    /// Retrieves all datasets.
    ///
    /// # Returns
    /// A list of all Custom Enrichments.
    pub async fn list(&self) -> Result<GetCustomEnrichmentsResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_custom_enrichments(make_request_with_metadata(
                GetCustomEnrichmentsRequest {},
                &self.metadata_map,
            ))
            .await
            .map_err(
                |status| SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.enrichment.v1.CustomEnrichmentService/GetCustomEnrichments".into(),
                    feature_group: DATASETS_FEATURE_GROUP_ID.into(),
                }),
            )?
            .into_inner())
    }
}
