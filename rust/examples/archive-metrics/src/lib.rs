#[cfg(test)]
mod tests {
    use cx_sdk::{
        auth::ApiKey,
        client::archive_metrics::{
            MetricsArchiveClient, S3Config, StorageConfig, StorageConfigUpdate,
            StorageConfigValidation, StorageConfigView,
        },
        CoralogixRegion,
    };

    #[tokio::test]
    #[ignore]
    async fn test_metrics_service_configurator() {
        let api_key = std::env::var("CORALOGIX_ALERTS_RULES_TAGS_API_KEY").unwrap();
        let metrics_service = MetricsArchiveClient::new(
            CoralogixRegion::from_env().unwrap(),
            ApiKey::from(api_key.as_str()),
        );

        let s3_config = S3Config {
            bucket: "yak-bucket-metrics".to_string(),
            region: "eu-west-1".to_string(),
        };

        let storage_config = StorageConfig::S3(s3_config.clone());
        let configure_tenant_result = metrics_service.configure_tenant(None, storage_config).await;

        if let Err(e) = &configure_tenant_result {
            let err = e.to_string();
            println!("Error: {:?}", err);
        }

        assert!(configure_tenant_result.is_ok());

        let update_tenant_result = metrics_service
            .update_tenant(2, StorageConfigUpdate::S3(s3_config.clone()))
            .await;

        assert!(update_tenant_result.is_ok());

        let validate_bucket_result = metrics_service
            .validate_bucket(StorageConfigValidation::S3(s3_config.clone()))
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
                == StorageConfigView::S3(s3_config.clone())
        );

        let enable_archive_result = metrics_service.enable_archive().await;
        assert!(enable_archive_result.is_ok());

        let disable_archive_result = metrics_service.disable_archive().await;
        assert!(disable_archive_result.is_ok());
    }
}
