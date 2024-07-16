#[cfg(test)]
mod tests {

    use cx_sdk::{
        auth::ApiKey,
        client::scopes::{EntityType, Filter, ScopesClient},
        CoralogixRegion,
    };

    #[tokio::test]
    async fn test_scopes_client() {
        let client = ScopesClient::new(
            ApiKey::from_env().unwrap(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let create_result = client
            .create(
                "Test Data Access Rule".into(),
                Some("Data Access Rule intended for testing".into()),
                vec![Filter {
                    entity_type: EntityType::Logs.into(),
                    expression: "<v1> foo == 'bar'".into(),
                }],
                "<v1> foo == 'bar'".into(),
            )
            .await;
        if let Err(e) = &create_result {
            let err = e.to_string();
            println!("Error: {:?}", err);
        }

        assert!(create_result.is_ok());

        let scope = create_result.unwrap().scope.unwrap();

        let update_result = client
            .update(
                scope.id,
                "Updated Test Data Access Rule".into(),
                scope.description,
                scope.filters,
                scope.default_expression,
            )
            .await;
        assert!(update_result.is_ok());

        let new_scope = update_result.unwrap().scope.unwrap();
        assert!(new_scope.display_name == "Updated Test Data Access Rule");

        let retrieved = client.get(vec![new_scope.id.clone()]).await;

        assert!(retrieved.is_ok());

        let delete_result = client.delete(new_scope.id).await;
        assert!(delete_result.is_ok());
        let retrieved = client.list().await;

        assert!(retrieved.is_ok());
        assert!(retrieved.unwrap().scopes.is_empty());
    }
}
