
#[cfg(test)]
mod tests {
    use cx_sdk::{client::datasets::DatasetClient, CoralogixRegion};
    use tokio::fs;


    #[tokio::test]
    async fn test_datasets() {
        let client = DatasetClient::new("my-api-key".into(), CoralogixRegion::from_env().unwrap()).unwrap();
        let dataset = client
            .create(
                "my-enrichment".into(),
                Some("My custom enrichment description".to_string()),
                fs::read_to_string("date-to-day-of-the-week.csv")
                    .await
                    .unwrap(),
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
                    .unwrap(),
            )
            .await
            .unwrap()
            .custom_enrichment
            .unwrap();

        let fetched = client.get(dataset_id).await.unwrap().unwrap();
        assert_eq!(updated.description, fetched.description);
        assert_eq!(fetched.version, dataset.version + 1);

        let _ = client.delete(dataset_id).await.unwrap();
    }
}
