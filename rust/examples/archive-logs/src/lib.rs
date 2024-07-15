#[cfg(test)]
mod tests {
    use cx_sdk::{
        auth::ApiKey,
        client::archive_logs::{LogsArchiveClient, S3TargetSpec, TargetSpec, TargetSpecValidation},
        CoralogixRegion,
    };

    #[tokio::test]
    async fn test_logs_archive() {
        let logs_archive_service = LogsArchiveClient::new(
            CoralogixRegion::from_env().unwrap(),
            ApiKey::from_env().unwrap(),
        );

        logs_archive_service
            .validate_target(
                true,
                TargetSpecValidation::S3(S3TargetSpec {
                    bucket: "coralogix-c4c-eu2-prometheus-data".to_string(),
                    region: Some("eu-west-1".to_string()),
                }),
            )
            .await
            .unwrap();

        logs_archive_service
            .set_target(
                true,
                TargetSpec::S3(S3TargetSpec {
                    bucket: "coralogix-c4c-eu2-prometheus-data".to_string(),
                    region: Some("eu-west-1".to_string()),
                }),
            )
            .await
            .unwrap();

        let _ = logs_archive_service.get_target().await.unwrap();
    }
}
