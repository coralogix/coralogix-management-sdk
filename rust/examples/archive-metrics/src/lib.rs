use cx_sdk::{
    auth::{ApiKey, AuthData},
    com::coralogix::metrics::metrics_configurator::{
        configure_tenant_request::StorageConfig, internal_update_request,
        metrics_configurator_public_service_client::MetricsConfiguratorPublicServiceClient,
        metrics_configurator_service_client::MetricsConfiguratorServiceClient, update_request,
        validate_bucket_request, ConfigureTenantRequest, GetTenantConfigRequest,
        GetTenantConfigResponse, GetTenantConfigResponseV2, InternalUpdateRequest,
        ListHotStoreConfigsRequest, ListHotStoreConfigsResponse, ListTenantConfigsRequest,
        ListTenantConfigsResponse, MigrateTenantRequest, RetentionPolicyRequest, UpdateRequest,
        ValidateBucketRequest,
    },
};
use std::str::FromStr;
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, Endpoint},
    Request,
};

pub fn make_request_with_metadata<T>(request: T, new_metadata: &MetadataMap) -> Request<T> {
    let mut req = Request::new(request);
    let metadata = req.metadata_mut();
    *metadata = new_metadata.clone();
    req
}

//This would be the client-facing service.
pub struct MetricsArchiveService {
    metadata_map: MetadataMap,
    service_client: Mutex<MetricsConfiguratorPublicServiceClient<Channel>>,
}

impl MetricsArchiveService {
    pub fn new(endpoint: &str, api_key: String) -> Self {
        let channel: Channel = Endpoint::from_str(endpoint).unwrap().connect_lazy();
        let api_key: ApiKey = api_key.into();
        let auth_data: AuthData = (&api_key).into();
        Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(MetricsConfiguratorPublicServiceClient::new(channel)),
        }
    }

    pub async fn configure_tenant(
        &self,
        retention_policy: RetentionPolicyRequest,
        storage_config: StorageConfig,
    ) -> anyhow::Result<()> {
        let request = make_request_with_metadata(
            ConfigureTenantRequest {
                retention_policy: Some(retention_policy),
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

    pub async fn update_tenant(
        &self,
        retention_days: u32,
        storage_config: update_request::StorageConfig,
    ) -> anyhow::Result<()> {
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

    pub async fn validate_bucket(
        &self,
        storage_config: validate_bucket_request::StorageConfig,
    ) -> anyhow::Result<()> {
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

    pub async fn get_tenant_config(&self) -> anyhow::Result<GetTenantConfigResponseV2> {
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

    pub async fn enable_archive(&self) -> anyhow::Result<()> {
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

    pub async fn disable_archive(&self) -> anyhow::Result<()> {
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

//This is a service that we would only use internally.
pub struct MetricsArchiveInternalService {
    metadata_map: MetadataMap,
    service_client: Mutex<MetricsConfiguratorServiceClient<Channel>>,
}

impl MetricsArchiveInternalService {
    pub fn new(endpoint: &str, api_key: String) -> Self {
        let channel: Channel = Endpoint::from_str(endpoint).unwrap().connect_lazy();
        let api_key: ApiKey = api_key.into();
        let auth_data: AuthData = (&api_key).into();
        Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(MetricsConfiguratorServiceClient::new(channel)),
        }
    }

    pub async fn get_tenant_config(
        &self,
        tenant_id: u32,
    ) -> anyhow::Result<GetTenantConfigResponse> {
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

    pub async fn list_tenant_configs(&self) -> anyhow::Result<ListTenantConfigsResponse> {
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

    pub async fn list_hot_store_configs(&self) -> anyhow::Result<ListHotStoreConfigsResponse> {
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

    pub async fn migrate_tenant(&self, tenant_id: u32) -> anyhow::Result<()> {
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

    pub async fn update(
        &self,
        tenant_id: u32,
        retention_days: u32,
        storage_config: internal_update_request::StorageConfig,
    ) -> anyhow::Result<()> {
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

#[cfg(test)]
mod tests {
    use cx_sdk::com::coralogix::metrics::metrics_configurator::{
        configure_tenant_request, internal_update_request, tenant_config, tenant_config_v2,
        update_request, validate_bucket_request, RetentionPolicyRequest, S3Config,
    };

    #[tokio::test]
    async fn test_metrics_service_configurator() {
        let endpoint = "https://ng-api-grpc.eu2.coralogix.com";
        let api_key = "api_key".to_string();
        let metrics_service = super::MetricsArchiveService::new(endpoint, api_key);

        let retention_policy = RetentionPolicyRequest {
            raw_resolution: 3,
            five_minutes_resolution: 5,
            one_hour_resolution: 7,
        };

        let s3_config = S3Config {
            bucket: "my-cool-bucket".to_string(),
            region: "us-west-2".to_string(),
        };

        let storage_config = configure_tenant_request::StorageConfig::S3(s3_config.clone());
        let configure_tenant_result = metrics_service
            .configure_tenant(retention_policy, storage_config)
            .await;

        assert!(configure_tenant_result.is_ok());

        let update_tenant_result = metrics_service
            .update_tenant(7, update_request::StorageConfig::S3(s3_config.clone()))
            .await;

        assert!(update_tenant_result.is_ok());

        let validate_bucket_result = metrics_service
            .validate_bucket(validate_bucket_request::StorageConfig::S3(
                s3_config.clone(),
            ))
            .await;

        assert!(validate_bucket_result.is_ok());

        let get_tenant_config_result = metrics_service.get_tenant_config().await;
        assert!(get_tenant_config_result.is_ok());
        assert!(
            get_tenant_config_result
                .unwrap()
                .tenant_config
                .unwrap()
                .storage_config
                .unwrap()
                == tenant_config_v2::StorageConfig::S3(s3_config.clone())
        );

        let enable_archive_result = metrics_service.enable_archive().await;
        assert!(enable_archive_result.is_ok());

        let disable_archive_result = metrics_service.disable_archive().await;
        assert!(disable_archive_result.is_ok());
    }

    #[tokio::test]
    async fn test_metrics_internal_service_configurator() {
        let endpoint = "https://ng-api-grpc.eu2.coralogix.com";
        let api_key = "api_key".to_string();
        let metrics_internal_service = super::MetricsArchiveInternalService::new(endpoint, api_key);

        let list_tenant_configs_result = metrics_internal_service.list_tenant_configs().await;
        assert!(list_tenant_configs_result.is_ok());

        let tenant_id = list_tenant_configs_result.unwrap().tenant_configs[0].tenant_id;

        let get_tenant_config_result = metrics_internal_service.get_tenant_config(tenant_id).await;
        assert!(get_tenant_config_result.is_ok());

        let list_hot_store_configs_result = metrics_internal_service.list_hot_store_configs().await;
        assert!(list_hot_store_configs_result.is_ok());

        let migrate_tenant_result = metrics_internal_service.migrate_tenant(tenant_id).await;
        assert!(migrate_tenant_result.is_ok());

        let update_result = metrics_internal_service
            .update(
                tenant_id,
                7,
                internal_update_request::StorageConfig::S3(S3Config {
                    bucket: "my-other-cool-bucket".to_string(),
                    region: "us-west-2".to_string(),
                }),
            )
            .await;

        assert!(update_result.is_ok());

        let tenant = get_tenant_config_result.unwrap().tenant_config.unwrap();
        assert!(
            tenant.storage_config.unwrap()
                == tenant_config::StorageConfig::S3(S3Config {
                    bucket: "my-other-cool-bucket".to_string(),
                    region: "us-west-2".to_string()
                })
        );
    }
}
