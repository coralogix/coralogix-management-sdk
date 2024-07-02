
#[cfg(test)]
mod tests {
    use std::{collections::HashMap, str::FromStr};

    use anyhow::Result;
    use cx_sdk::{
        auth::{ApiKey, AuthData},
        client::{actions::Action, ActionsClient}, CoralogixRegion
    };
    use tokio::sync::Mutex;
    
    #[tokio::test]
    async fn test_actions_client() {
        let client = ActionsClient::new(
            "api-key".to_string().into(),
            CoralogixRegion::from_env().unwrap(),
        ).unwrap();

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
            println!("Error: {:?}", e);
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

        let retrieved_action = client
            .get(replaced_action.id.clone().unwrap())
            .await;

        assert!(retrieved_action.is_ok());
        assert!(retrieved_action.unwrap().is_some());

        let delete_action_result = client.delete(replaced_action.id.unwrap()).await;
        assert!(delete_action_result.is_ok());

        let retrieved_actions = client.list().await;

        assert!(retrieved_actions.is_ok());
        assert!(!retrieved_actions.unwrap().is_empty());
    }
}
