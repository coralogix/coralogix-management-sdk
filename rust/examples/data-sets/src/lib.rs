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
    use cx_sdk::{auth::AuthContext, client::datasets::DatasetClient, CoralogixRegion};
    use tokio::fs;

    #[tokio::test]
    async fn test_datasets() {
        let client = DatasetClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();
        let dataset = client
            .create(
                "my-enrichment".into(),
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
                "my-enrichment".into(),
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
