#[cfg(test)]
mod tests {
    use cx_sdk::{auth::ApiKey, client::datasets::DatasetClient, CoralogixRegion};
    use tokio::fs;

    #[tokio::test]
    async fn test_datasets() {
        let client = DatasetClient::new(
            ApiKey::from_env().unwrap(),
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
