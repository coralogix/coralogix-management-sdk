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
        CoralogixRegion,
        auth::AuthContext,
        client::{
            archive_logs::{
                LogsArchiveClient,
                S3TargetSpec,
                TargetSpec,
                TargetSpecValidation,
            },
            datasets::DatasetClient,
        },
    };
    use tokio::fs;
    use uuid::Uuid;

    #[tokio::test]
    async fn test_logs_archive() {
        let aws_region = std::env::var("AWS_REGION").unwrap();
        let logs_bucket_name = std::env::var("LOGS_BUCKET").unwrap();
        let service = LogsArchiveClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
        )
        .unwrap();

        let target_spec = S3TargetSpec {
            bucket: logs_bucket_name,
            region: Some(aws_region),
        };

        service
            .set_target(true, TargetSpec::S3(target_spec))
            .await
            .unwrap();

        let _ = service.get_target().await.unwrap();
    }

    #[tokio::test]
    async fn test_datasets() {
        let client = DatasetClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();
        let name = Uuid::new_v4().to_string().replace("-", "");

        let dataset = client
            .create(
                name.clone(),
                Some("My custom enrichment description".to_string()),
                fs::read_to_string("date-to-day-of-the-week.csv")
                    .await
                    .unwrap()
                    .into_bytes(),
            )
            .await
            .unwrap()
            .custom_enrichment
            .unwrap();
        let dataset_id = dataset.id;
        let updated = client
            .replace(
                dataset_id,
                name.into(),
                Some("My updated enrichment description".to_string()),
                fs::read_to_string("date-to-day-of-the-week.csv")
                    .await
                    .unwrap()
                    .into_bytes(),
            )
            .await
            .unwrap()
            .custom_enrichment
            .unwrap();

        let fetched = client.get(dataset_id).await.unwrap();
        let custom_enrichment = fetched.custom_enrichment.unwrap();
        assert_eq!(updated.description, custom_enrichment.description);
        assert_eq!(custom_enrichment.version, dataset.version + 1);

        let _ = client.delete(dataset_id).await.unwrap();
    }
}
