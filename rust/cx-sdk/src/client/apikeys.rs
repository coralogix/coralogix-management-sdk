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

use cx_api::proto::com::coralogixapis::aaa::apikeys::v3::{
    CreateApiKeyRequest,
    CreateApiKeyResponse,
    DeleteApiKeyRequest,
    DeleteApiKeyResponse,
    GetApiKeyRequest,
    GetApiKeyResponse,
    Owner as OwnerWrapper,
    UpdateApiKeyRequest,
    UpdateApiKeyResponse,
    api_keys_service_client::ApiKeysServiceClient,
    create_api_key_request::KeyPermissions,
    update_api_key_request::{
        Permissions,
        Presets,
    },
};
use std::str::FromStr;

use crate::{
    CoralogixRegion,
    auth::AuthContext,
    error::Result,
    metadata::CallProperties,
    util::make_request_with_metadata,
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

pub use crate::com::coralogixapis::aaa::apikeys::v3::owner::Owner;

/// The API Keys API client.
/// Read more at <https://coralogix.com/docs/api-keys/>
pub struct ApiKeysClient {
    metadata_map: MetadataMap,
    service_client: Mutex<ApiKeysServiceClient<Channel>>,
}

impl ApiKeysClient {
    /// Creates a new client for the APIKeys API.
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
            service_client: Mutex::new(ApiKeysServiceClient::new(channel)),
        })
    }

    /// Creates a new API Key
    ///
    /// # Arguments
    /// * `name` - The name of the API key.
    /// * `owner` - The [`Owner`] of the API key.
    /// * `presets` - The presets of the API key.
    /// * `permissions` - The permissions of the API key.
    /// * `hashed` - Whether the API key should be encrypted.
    ///
    /// Note that when the API key is hashed, it will not be possible to retrieve it later.
    pub async fn create(
        &self,
        name: String,
        owner: Option<Owner>,
        presets: Vec<String>,
        permissions: Vec<String>,
        hashed: bool,
    ) -> Result<CreateApiKeyResponse> {
        let request = make_request_with_metadata(
            CreateApiKeyRequest {
                hashed,
                name,
                owner: owner.map(|o| OwnerWrapper { owner: Some(o) }),
                key_permissions: Some(KeyPermissions {
                    presets,
                    permissions,
                }),
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .create_api_key(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Updates an API key.
    ///
    /// # Arguments
    /// * `key_id` - The ID of the API key to update.
    /// * `is_active` - Whether the API key should be active.
    /// * `new_name` - The new name of the API key.
    /// * `presets` - The new presets of the API key.
    /// * `permissions` - The new permissions of the API key.
    pub async fn update(
        &self,
        key_id: String,
        is_active: Option<bool>,
        new_name: Option<String>,
        presets: Option<Vec<String>>,
        permissions: Option<Vec<String>>,
    ) -> Result<UpdateApiKeyResponse> {
        let request = make_request_with_metadata(
            UpdateApiKeyRequest {
                key_id,
                new_name,
                is_active,
                presets: presets.map(|p| Presets { presets: p }),
                permissions: permissions.map(|p| Permissions { permissions: p }),
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .update_api_key(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Deletes an API key.
    ///
    /// # Arguments
    /// * `key_id` - The ID of the API key to delete.
    pub async fn delete(&self, key_id: String) -> Result<DeleteApiKeyResponse> {
        let request =
            make_request_with_metadata(DeleteApiKeyRequest { key_id }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .delete_api_key(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    /// Retrieves an API key by its ID.
    ///
    /// # Arguments
    /// * `key_id` - The ID of the API key to retrieve.
    pub async fn get(&self, key_id: String) -> Result<GetApiKeyResponse> {
        let request = make_request_with_metadata(GetApiKeyRequest { key_id }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .get_api_key(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }
}
