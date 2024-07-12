#[cfg(test)]
mod tests {

    use cx_sdk::{
        auth::ApiKey,
        client::actions::{Action, ActionsClient, SourceType},
        CoralogixRegion,
    };

    #[tokio::test]
    async fn test_actions_client() {
        let api_key = std::env::var("CORALOGIX_LEGACY_KEY").unwrap();
        let client = ActionsClient::new(
            ApiKey::from(api_key.as_str()),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        delete_all_actions(&client).await;

        let action = Action {
            name: Some("google search action".to_string()),
            url: Some("https://www.google.com/search?q={{$p.selected_value}}".to_string()),
            is_private: Some(false),
            source_type: SourceType::Log.into(),
            application_names: vec!["app".to_string()],
            subsystem_names: vec!["sub".to_string()],
            id: None,
            is_hidden: Some(false),
            created_by: Some("luigi.taglialatela@coralogix.com".into()),
        };

        let create_action_result = client.create(action).await;
        if let Err(e) = &create_action_result {
            let err = e.to_string();
            println!("Error: {:?}", err);
        }

        assert!(create_action_result.is_ok());

        let created_action = create_action_result.unwrap().action.unwrap();
        let updated_action = Action {
            name: Some("updated action".to_string()),
            ..created_action
        };

        let replace_action_result = client.replace(updated_action).await;
        assert!(replace_action_result.is_ok());

        let replaced_action = replace_action_result.unwrap().action.unwrap();
        assert!(replaced_action.name.unwrap() == "updated action");

        let retrieved_action = client.get(replaced_action.id.clone().unwrap()).await;

        assert!(retrieved_action.is_ok());
        assert!(retrieved_action.unwrap().is_some());

        let delete_action_result = client.delete(replaced_action.id.unwrap()).await;
        assert!(delete_action_result.is_ok());

        let retrieved_actions = client.list().await;

        assert!(retrieved_actions.is_ok());
    }

    async fn delete_all_actions(actions_client: &ActionsClient) {
        let actions = actions_client.list().await.unwrap();
        for action in actions {
            actions_client.delete(action.id.unwrap()).await.unwrap();
        }
    }
}
