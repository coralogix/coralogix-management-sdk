#[cfg(test)]
mod tests {
    use cx_sdk::{
        auth::ApiKey,
        client::archive_logs::{LogsArchiveClient, S3TargetSpec, TargetSpec, TargetSpecValidation},
        CoralogixRegion,
    };

    #[tokio::test]
    async fn test_logs_archive() {
        let service = LogsArchiveClient::new(
            CoralogixRegion::from_env().unwrap(),
            ApiKey::from_env().unwrap(),
        );

        let target_spec = S3TargetSpec {
            bucket: "yak-2-bucket".to_string(),
            region: Some("eu-north-1".to_string()),
        };
        service
            .validate_target(true, TargetSpecValidation::S3(target_spec.clone()))
            .await
            .unwrap();

        service
            .set_target(true, TargetSpec::S3(target_spec))
            .await
            .unwrap();

        let _ = service.get_target().await.unwrap();
    }
}
