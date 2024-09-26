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

use cx_api::proto::com::coralogix::metrics::metrics_configurator::{
    metrics_configurator_public_service_client::MetricsConfiguratorPublicServiceClient,
    metrics_configurator_service_client::MetricsConfiguratorServiceClient,
    ConfigureTenantRequest,
    GetTenantConfigRequest,
    GetTenantConfigResponse,
    GetTenantConfigResponseV2,
    InternalUpdateRequest,
    ListHotStoreConfigsRequest,
    ListHotStoreConfigsResponse,
    ListTenantConfigsRequest,
    ListTenantConfigsResponse,
    MigrateTenantRequest,
    UpdateRequest,
    ValidateBucketRequest,
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
    auth::AuthContext,
    error::Result,
    metadata::CallProperties,
    util::make_request_with_metadata,
    CoralogixRegion,
};

pub use cx_api::proto::com::coralogix::metrics::metrics_configurator::{
    configure_tenant_request::StorageConfig,
    internal_update_request::StorageConfig as InternalStorageConfigUpdate,
    tenant_config::StorageConfig as InternalStorageConfig,
    tenant_config_v2::StorageConfig as StorageConfigView,
    update_request::StorageConfig as StorageConfigUpdate,
    validate_bucket_request::StorageConfig as StorageConfigValidation,
    RetentionPolicyRequest,
    S3Config,
};

/// The metrics archive API client.
/// Read more at [https://coralogix.com/docs/archive-s3-bucket-forever/]()
pub struct MetricsArchiveClient {
    metadata_map: MetadataMap,
    service_client: Mutex<MetricsConfiguratorPublicServiceClient<Channel>>,
}

impl MetricsArchiveClient {
    /// Creates a new client for the Metrics Archive API.
    ///
    /// # Arguments
    /// * `auth_context` - The  to use for authentication.
    /// * `region` - The region to connect to.
    pub fn new(region: CoralogixRegion, auth_context: AuthContext) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(&region.grpc_endpoint())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();
        let request_metadata: CallProperties = (&auth_context.team_level_api_key).into();
        Ok(Self {
            metadata_map: request_metadata.to_metadata_map(),
            service_client: Mutex::new(MetricsConfiguratorPublicServiceClient::new(channel)),
        })
    }

    /// Configures the tenant.
    ///
    /// # Arguments
    /// * `retention_policy` - The retention policy to set.
    /// * `storage_config` - The storage configuration to set.
    pub async fn configure_tenant(
        &self,
        retention_policy: Option<RetentionPolicyRequest>,
        storage_config: StorageConfig,
    ) -> Result<()> {
        let request = make_request_with_metadata(
            ConfigureTenantRequest {
                retention_policy,
                storage_config: Some(storage_config),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .configure_tenant(request)
                .await
                .map(|_| ())
                .map_err(From::from)
        }
    }

    /// Updates the tenant configuration.
    ///
    /// # Arguments
    /// * `retention_days` - The retention days to set.
    /// * `storage_config` - The storage configuration to set.
    pub async fn update_tenant(
        &self,
        retention_days: u32,
        storage_config: StorageConfigUpdate,
    ) -> Result<()> {
        let request = make_request_with_metadata(
            UpdateRequest {
                retention_days: Some(retention_days),
                storage_config: Some(storage_config),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client.update(request).await.map(|_| ()).map_err(From::from)
        }
    }

    /// Validates a bucket configuration.
    ///
    /// # Arguments
    /// * `storage_config` - The storage configuration to validate.
    /// * `bucket_name` - The name of the bucket to validate.
    pub async fn validate_bucket(&self, storage_config: StorageConfigValidation) -> Result<()> {
        let request = make_request_with_metadata(
            ValidateBucketRequest {
                storage_config: Some(storage_config),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .validate_bucket(request)
                .await
                .map(|_| ())
                .map_err(From::from)
        }
    }

    /// Gets the tenant configuration.
    pub async fn get_tenant_config(&self) -> Result<GetTenantConfigResponseV2> {
        let request = make_request_with_metadata((), &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .get_tenant_config(request)
                .await
                .map(|response| response.into_inner())
                .map_err(From::from)
        }
    }

    /// Enables the archive.
    pub async fn enable_archive(&self) -> Result<()> {
        let request = make_request_with_metadata((), &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .enable_archive(request)
                .await
                .map(|_| ())
                .map_err(From::from)
        }
    }

    /// Disables the archive.
    pub async fn disable_archive(&self) -> Result<()> {
        let request = make_request_with_metadata((), &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .disable_archive(request)
                .await
                .map(|_| ())
                .map_err(From::from)
        }
    }
}

/// A service client for the internal metrics archive API. It's only for Coralogix internal use.
pub struct MetricsArchiveInternalClient {
    metadata_map: MetadataMap,
    service_client: Mutex<MetricsConfiguratorServiceClient<Channel>>,
}

impl MetricsArchiveInternalClient {
    /// Creates a new client for the internal Metrics Archive API.
    ///
    /// # Arguments
    /// * `auth_context` - The  to use for authentication.
    /// * `region` - The region to connect to.
    pub fn new(region: CoralogixRegion, auth_context: AuthContext) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(&region.grpc_endpoint())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();
        let request_metadata: CallProperties = (&auth_context.team_level_api_key).into();
        Ok(Self {
            metadata_map: request_metadata.to_metadata_map(),
            service_client: Mutex::new(MetricsConfiguratorServiceClient::new(channel)),
        })
    }

    /// Gets the tenant configuration.
    ///
    /// # Arguments
    /// * `tenant_id` - The ID of the tenant to get the configuration for.
    pub async fn get_tenant_config(&self, tenant_id: u32) -> Result<GetTenantConfigResponse> {
        let request =
            make_request_with_metadata(GetTenantConfigRequest { tenant_id }, &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .get_tenant_config(request)
                .await
                .map(|response| response.into_inner())
                .map_err(From::from)
        }
    }

    /// Lists the tenant configurations.
    pub async fn list_tenant_configs(&self) -> Result<ListTenantConfigsResponse> {
        let request = make_request_with_metadata(ListTenantConfigsRequest {}, &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .list_tenant_configs(request)
                .await
                .map(|response| response.into_inner())
                .map_err(From::from)
        }
    }

    /// Lists the hot store configurations.
    pub async fn list_hot_store_configs(&self) -> Result<ListHotStoreConfigsResponse> {
        let request = make_request_with_metadata(ListHotStoreConfigsRequest {}, &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .list_host_store_configs(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Migrates a tenant.
    ///
    /// # Arguments
    /// * `tenant_id` - The ID of the tenant to migrate.
    pub async fn migrate_tenant(&self, tenant_id: u32) -> Result<()> {
        let request =
            make_request_with_metadata(MigrateTenantRequest { tenant_id }, &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .migrate_tenant(request)
                .await
                .map(|_| ())
                .map_err(From::from)
        }
    }

    /// Updates the tenant configuration.
    ///
    /// # Arguments
    /// * `tenant_id` - The ID of the tenant to update.
    pub async fn update(
        &self,
        tenant_id: u32,
        retention_days: u32,
        storage_config: InternalStorageConfigUpdate,
    ) -> Result<()> {
        let request = make_request_with_metadata(
            InternalUpdateRequest {
                retention_days: Some(retention_days),
                storage_config: Some(storage_config),
                tenant_id,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client.update(request).await.map(|_| ()).map_err(From::from)
        }
    }
}
