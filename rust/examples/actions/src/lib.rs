use std::{collections::HashMap, str::FromStr};

use anyhow::Result;
use cx_sdk::{
    auth::{ApiKey, AuthData},
    com::coralogixapis::actions::v2::{
        actions_service_client::ActionsServiceClient, Action, CreateActionRequest,
        CreateActionResponse, DeleteActionRequest, GetActionRequest, ListActionsRequest,
        OrderActionsRequest, ReplaceActionRequest, ReplaceActionResponse,
    },
};
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, Endpoint},
    Request,
};

pub fn make_request_with_metadata<T>(request: T, new_metadata: &MetadataMap) -> Request<T> {
    let mut req = Request::new(request);
    let metadata = req.metadata_mut();
    *metadata = new_metadata.clone();
    req
}

pub struct ActionsService {
    metadata_map: MetadataMap,
    service_client: Mutex<ActionsServiceClient<Channel>>,
}

impl ActionsService {
    pub fn new(endpoint: &str, api_key: String) -> Self {
        let channel: Channel = Endpoint::from_str(endpoint).unwrap().connect_lazy();
        let api_key: ApiKey = api_key.into();
        let auth_data: AuthData = (&api_key).into();
        Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(ActionsServiceClient::new(channel)),
        }
    }

    pub async fn create_action(&self, action: Action) -> Result<CreateActionResponse> {
        let request = make_request_with_metadata(
            CreateActionRequest {
                name: action.name,
                url: action.url,
                is_private: action.is_private,
                source_type: action.source_type,
                application_names: action.application_names,
                subsystem_names: action.subsystem_names,
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .create_action(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    pub async fn replace_action(&self, action: Action) -> Result<ReplaceActionResponse> {
        let request = make_request_with_metadata(
            ReplaceActionRequest {
                action: Some(action),
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .replace_action(request)
            .await
            .map(|r| r.into_inner())
            .map_err(From::from)
    }

    pub async fn delete_action(&self, action_id: String) -> Result<()> {
        let request = make_request_with_metadata(
            DeleteActionRequest {
                id: Some(action_id),
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .delete_action(request)
            .await
            .map(|_| ())
            .map_err(From::from)
    }

    pub async fn get_action(&self, action_id: String) -> Result<Option<Action>> {
        let request = make_request_with_metadata(
            GetActionRequest {
                id: Some(action_id),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_action(request)
            .await
            .map(|r| r.into_inner().action)
            .map_err(From::from)
    }

    pub async fn list_actions(&self) -> Result<Vec<Action>> {
        let request = make_request_with_metadata(ListActionsRequest {}, &self.metadata_map);

        self.service_client
            .lock()
            .await
            .list_actions(request)
            .await
            .map(|r| r.into_inner().actions)
            .map_err(From::from)
    }

    pub async fn order_actions(
        &self,
        private_actions_order: HashMap<String, u32>,
        shared_actions_order: HashMap<String, u32>,
    ) -> Result<()> {
        let request = make_request_with_metadata(
            OrderActionsRequest {
                private_actions_order,
                shared_actions_order,
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .order_actions(request)
            .await
            .map(|_| ())
            .map_err(From::from)
    }
}

#[cfg(test)]
mod tests {
    use cx_sdk::com::coralogixapis::actions::v2::{Action, SourceType};

    use crate::ActionsService;

    #[tokio::test]
    async fn test_actions_service() {
        let service = ActionsService::new(
            "https://ng-api-grpc.eu2.coralogix.com",
            "api-key".to_string(),
        );

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

        let create_action_result = service.create_action(action).await;
        if let Err(e) = &create_action_result {
            println!("Error: {:?}", e);
        }

        assert!(create_action_result.is_ok());

        let created_action = create_action_result.unwrap().action.unwrap();
        let updated_action = Action {
            name: Some("updated action".to_string()),
            ..created_action
        };

        let replace_action_result = service.replace_action(updated_action).await;
        assert!(replace_action_result.is_ok());

        let replaced_action = replace_action_result.unwrap().action.unwrap();
        assert!(replaced_action.name.unwrap() == "updated action");

        let retrieved_action = service
            .get_action(replaced_action.id.clone().unwrap())
            .await;

        assert!(retrieved_action.is_ok());
        assert!(retrieved_action.unwrap().is_some());

        let delete_action_result = service.delete_action(replaced_action.id.unwrap()).await;
        assert!(delete_action_result.is_ok());

        let retrieved_actions = service.list_actions().await;

        assert!(retrieved_actions.is_ok());
        assert!(!retrieved_actions.unwrap().is_empty());
    }
}
