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

use cx_api::proto::com::coralogix::archive::v1::{
    ActivateRetentionsRequest,
    ActivateRetentionsResponse,
    GetRetentionsEnabledRequest,
    GetRetentionsEnabledResponse,
    GetRetentionsRequest,
    GetRetentionsResponse,
    UpdateRetentionsRequest,
    UpdateRetentionsResponse,
    retentions_service_client::RetentionsServiceClient,
};
use std::str::FromStr;
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
    error::Result,
    metadata::CallProperties,
    util::make_request_with_metadata,
};

pub use cx_api::proto::com::coralogix::archive::v1::RetentionUpdateElement;

/// The Archive Retention API client.
/// Read more at <https://coralogix.com/docs/archive-retention/>
pub struct ArchiveRetentionClient {
    metadata_map: MetadataMap,
    service_client: Mutex<RetentionsServiceClient<Channel>>,
}

impl ArchiveRetentionClient {
    /// Creates a new client for the Archive Retention API.
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
            service_client: Mutex::new(RetentionsServiceClient::new(channel)),
        })
    }

    /// Retrieves the retention settings for a tenant.
    pub async fn get(&self) -> Result<GetRetentionsResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_retentions(make_request_with_metadata(
                GetRetentionsRequest {},
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    /// Retrieves the enabled retention settings for a tenant.
    pub async fn get_enabled(&self) -> Result<GetRetentionsEnabledResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_retentions_enabled(make_request_with_metadata(
                GetRetentionsEnabledRequest {},
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    /// Updates the retention settings for a tenant.
    pub async fn update(
        &self,
        retention_update_elements: Vec<RetentionUpdateElement>,
    ) -> Result<UpdateRetentionsResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .update_retentions(make_request_with_metadata(
                UpdateRetentionsRequest {
                    retention_update_elements,
                },
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    /// Activates all archive retentions
    pub async fn activate(&self) -> Result<ActivateRetentionsResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .activate_retentions(make_request_with_metadata(
                ActivateRetentionsRequest {},
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }
}
