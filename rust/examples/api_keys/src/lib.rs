#[cfg(test)]
mod tests {
    use cx_sdk::{client::apikeys::{ApiKeysClient, Owner}, CoralogixRegion};
    #[tokio::test]
    async fn test_actions_client() {
        let client = ApiKeysClient::new(
            ApiKey::from_env().unwrap(),
            CoralogixRegion::from_env().unwrap(),
        ).unwrap();

        let create_result = client
            .create(
                "My APM KEY".to_string(),
                Some(Owner::UserId("4013254".to_string())),
                vec!["APM".to_string()],
                vec![],
                false,
            )
            .await;

        assert!(create_result.is_ok());

        let key_id = create_result.unwrap().key_id;

        let update_result = client
            .update(
                key_id.clone(),
                None,
                Some("new-name".to_string()),
                None,
                None,
            )
            .await;
        assert!(update_result.is_ok());

        let new_api_key = client.get(key_id.clone()).await;

        assert!(new_api_key.is_ok());
        assert_eq!(
            new_api_key.unwrap().key_info.map(|k| k.name),
            Some("new-name".to_string())
        );

        let delete_action_result = client.delete(key_id).await;
        assert!(delete_action_result.is_ok());
    }
}
