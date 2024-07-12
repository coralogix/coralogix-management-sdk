#[cfg(test)]
mod tests {
    use cx_sdk::{
        auth::ApiKey,
        client::archive_logs::{LogsArchiveClient, S3TargetSpec, TargetSpec, TargetSpecValidation},
        CoralogixRegion,
    };

    #[tokio::test]
    async fn test_logs_archive() {
        let api_key = std::env::var("CORALOGIX_API_KEY").unwrap();
        let logs_archive_service = LogsArchiveClient::new(
            CoralogixRegion::from_env().unwrap(),
            ApiKey::from(api_key.as_str()),
        );

        let validate_target_result = logs_archive_service
            .validate_target(
                true,
                TargetSpecValidation::S3(S3TargetSpec {
                    bucket: "yak-bucket-logs".to_string(),
                    region: Some("eu-west-1".to_string()),
                }),
            )
            .await;

        if let Err(e) = &validate_target_result {
            let error = e.to_string();
            println!("Error: {:?}", error);
        }

        assert!(validate_target_result.is_ok());

        let set_target_result = logs_archive_service
            .set_target(
                true,
                TargetSpec::S3(S3TargetSpec {
                    bucket: "yak-bucket-logs".to_string(),
                    region: Some("eu-west-1".to_string()),
                }),
            )
            .await;

        assert!(set_target_result.is_ok());

        let get_target_response = logs_archive_service.get_target().await;

        assert!(get_target_response.is_ok());
    }
}
