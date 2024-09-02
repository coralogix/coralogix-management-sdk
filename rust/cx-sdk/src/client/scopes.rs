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

pub use cx_api::proto::com::coralogixapis::scopes::v1::{EntityType, Filter};

use cx_api::proto::com::coralogixapis::scopes::v1::{
    scopes_service_client::ScopesServiceClient, CreateScopeRequest, CreateScopeResponse,
    DeleteScopeRequest, DeleteScopeResponse, GetScopesResponse, GetTeamScopesByIdsRequest,
    GetTeamScopesRequest, UpdateScopeRequest, UpdateScopeResponse,
};
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, ClientTlsConfig, Endpoint},
};

use crate::CoralogixRegion;

/// The Scopes Service client.
/// Read more at <https://coralogix.com/docs/scopes/>
pub struct ScopesClient {
    metadata_map: MetadataMap,
    service_client: Mutex<ScopesServiceClient<Channel>>,
}

impl ScopesClient {
    /// Creates a new client for the SLO.
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
            service_client: Mutex::new(ScopesServiceClient::new(channel)),
        })
    }

    /// Create a new Scope.
    ///
    /// # Arguments
    /// * `display_name` - The display name of the Scope.
    /// * `description` - The description of the Scope.
    /// * `filters` - The filters of the Scope.
    /// * `default_expression` - The default expression of the Scope.
    pub async fn create(
        &self,
        display_name: String,
        description: Option<String>,
        filters: Vec<Filter>,
        default_expression: String,
    ) -> Result<CreateScopeResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .create_scope(make_request_with_metadata(
                CreateScopeRequest {
                    display_name,
                    description,
                    filters,
                    default_expression,
                },
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    /// Update a Scope.
    ///
    /// # Arguments
    /// * `id` - The ID of the Scope to update.
    /// * `display_name` - The display name of the Scope.
    /// * `description` - The description of the Scope.
    /// * `filters` - The filters of the Scope.
    /// * `default_expression` - The default expression of the Scope.
    pub async fn update(
        &self,
        id: String,
        display_name: String,
        description: Option<String>,
        filters: Vec<Filter>,
        default_expression: String,
    ) -> Result<UpdateScopeResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .update_scope(make_request_with_metadata(
                UpdateScopeRequest {
                    id,
                    display_name,
                    description,
                    filters,
                    default_expression,
                },
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    /// Delete a Scope by its ID.
    ///
    /// # Arguments
    /// * `id` - The ID of the Scope to delete.
    pub async fn delete(&self, id: String) -> Result<DeleteScopeResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .delete_scope(make_request_with_metadata(
                DeleteScopeRequest { id },
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    /// Get a list of Scopes for your team by their IDs.
    ///
    /// # Arguments
    /// * `ids` - The IDs of the Scopes to get.
    pub async fn get(&self, ids: Vec<String>) -> Result<GetScopesResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_team_scopes_by_ids(make_request_with_metadata(
                GetTeamScopesByIdsRequest { ids },
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }

    /// Get a list of all Scopes for your team.
    pub async fn list(&self) -> Result<GetScopesResponse> {
        Ok(self
            .service_client
            .lock()
            .await
            .get_team_scopes(make_request_with_metadata(
                GetTeamScopesRequest {},
                &self.metadata_map,
            ))
            .await?
            .into_inner())
    }
}
