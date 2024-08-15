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

#[cfg(test)]
mod tests {
    use cx_sdk::{
        auth::AuthContext,
        client::archive_metrics::{
            MetricsArchiveClient, S3Config, StorageConfig, StorageConfigUpdate,
            StorageConfigValidation, StorageConfigView,
        },
        CoralogixRegion,
    };

    #[tokio::test]
    async fn test_metrics_service_configurator() {
        let metrics_service = MetricsArchiveClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
        )
        .unwrap();

        let s3_config = S3Config {
            bucket: "yak-coralogix-bucket".to_string(),
            region: "eu-north-1".to_string(),
        };

        let storage_config = StorageConfig::S3(s3_config.clone());
        metrics_service
            .configure_tenant(None, storage_config)
            .await
            .unwrap();

        metrics_service
            .validate_bucket(StorageConfigValidation::S3(s3_config.clone()))
            .await
            .unwrap();

        metrics_service
            .update_tenant(2, StorageConfigUpdate::S3(s3_config.clone()))
            .await
            .unwrap();

        let get_tenant_config_result = metrics_service.get_tenant_config().await;
        assert!(
            get_tenant_config_result
                .unwrap()
                .tenant_config
                .unwrap()
                .storage_config
                .unwrap()
                == StorageConfigView::S3(s3_config.clone())
        );

        metrics_service.enable_archive().await.unwrap();
        metrics_service.disable_archive().await.unwrap();
    }
}
