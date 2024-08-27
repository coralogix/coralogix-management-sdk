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
    auth::AuthContext, error::Result, metadata::CallProperties, util::make_request_with_metadata,
};

pub use crate::com::coralogixapis::actions::v2::Action;

use cx_api::proto::com::coralogix::enrichment::v1::{
    enrichment_service_client::EnrichmentServiceClient, AddEnrichmentsRequest,
    AddEnrichmentsResponse, AtomicOverwriteEnrichmentsRequest, AtomicOverwriteEnrichmentsResponse,
    GetEnrichmentLimitRequest, GetEnrichmentLimitResponse, GetEnrichmentsRequest,
    GetEnrichmentsResponse, RemoveEnrichmentsRequest, RemoveEnrichmentsResponse,
};
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, ClientTlsConfig, Endpoint},
};

use crate::CoralogixRegion;

pub use crate::com::coralogix::enrichment::v1::{
    enrichment_type::Type, EnrichmentFieldDefinition, EnrichmentRequestModel as EnrichmentMapping,
    EnrichmentType,
};

/// The Custom Enrichments API client.
/// Read more at <https://coralogix.com/docs/custom-enrichment-api/>
pub struct EnrichmentsClient {
    metadata_map: MetadataMap,
    service_client: Mutex<EnrichmentServiceClient<Channel>>,
}

impl EnrichmentsClient {
    /// Creates a new client for the Enrichments API.
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
            service_client: Mutex::new(EnrichmentServiceClient::new(channel)),
        })
    }

    /// Adds enrichment mappings.
    ///
    /// # Arguments
    /// * `enrichment_mapping` - The enrichment definitions to create.
    pub async fn add(
        &self,
        enrichment_mapping: Vec<EnrichmentMapping>,
    ) -> Result<AddEnrichmentsResponse> {
        self.service_client
            .lock()
            .await
            .add_enrichments(make_request_with_metadata(
                AddEnrichmentsRequest {
                    request_enrichments: enrichment_mapping,
                },
                &self.metadata_map,
            ))
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Updates the mappings for an enrichment.
    ///
    /// # Arguments
    /// * `enrichment_type` - The type of enrichment to update.
    /// * `field_mappings` - The associated field mappings.
    pub async fn update(
        &self,
        enrichment_type: EnrichmentType,
        field_mappings: Vec<EnrichmentFieldDefinition>,
    ) -> Result<AtomicOverwriteEnrichmentsResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .atomic_overwrite_enrichments(make_request_with_metadata(
                AtomicOverwriteEnrichmentsRequest {
                    enrichment_type: Some(enrichment_type),
                    enrichment_fields: field_mappings,
                },
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    /// Removes all enrichment mappings from the provided id.
    ///
    /// # Arguments
    /// * `ids` - The ids to remove mappings from.
    pub async fn remove_all(&self, dataset_ids: Vec<u32>) -> Result<RemoveEnrichmentsResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .remove_enrichments(make_request_with_metadata(
                RemoveEnrichmentsRequest {
                    enrichment_ids: dataset_ids,
                },
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    /// Retrieves the all enrichments.
    pub async fn list(&self) -> Result<GetEnrichmentsResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_enrichments(make_request_with_metadata(
                GetEnrichmentsRequest {},
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    /// Retrieves the Enrichment limits.
    pub async fn get_limits(&self) -> Result<GetEnrichmentLimitResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_enrichment_limit(make_request_with_metadata(
                GetEnrichmentLimitRequest {},
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }
}
