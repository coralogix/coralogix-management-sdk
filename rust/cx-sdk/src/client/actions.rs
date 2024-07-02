use std::{collections::HashMap, str::FromStr};

use crate::{
    auth::{ApiKey, AuthData},
    com::coralogixapis::actions::v2::{
        actions_service_client::ActionsServiceClient, CreateActionRequest, CreateActionResponse,
        DeleteActionRequest, GetActionRequest, ListActionsRequest, OrderActionsRequest,
        ReplaceActionRequest, ReplaceActionResponse,
    },
    error::Result,
    util::make_request_with_metadata,
};

pub use crate::com::coralogixapis::actions::v2::Action;

use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, Endpoint},
};

use crate::CoralogixRegion;

/// The Actions API client.
/// Read more at [https://coralogix.com/docs/coralogix-action-extension/]()
pub struct ActionsClient {
    metadata_map: MetadataMap,
    service_client: Mutex<ActionsServiceClient<Channel>>,
}

impl ActionsClient {
    /// Creates a new client for the Actions API.
    ///
    /// # Arguments
    /// * `api_key` - The API key to use for authentication.
    /// * `region` - The region to connect to.
    pub fn new(api_key: ApiKey, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.endpoint().as_str())?.connect_lazy();
        let auth_data: AuthData = (&api_key).into();
        Ok(Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(ActionsServiceClient::new(channel)),
        })
    }

    /// Creates a new Action
    ///
    /// # Arguments
    /// * `action` - The [`Action`] to create.
    pub async fn create(&self, action: Action) -> Result<CreateActionResponse> {
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

    /// Replaces the existing [`Action`] identified by its id.
    ///
    /// # Arguments
    /// * `action` - The action to replace.
    pub async fn replace(&self, action: Action) -> Result<ReplaceActionResponse> {
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

    /// Deletes the [`Action`] identified by its id.
    ///
    /// # Arguments
    /// * `action_id` - The id of the action to delete.
    pub async fn delete(&self, action_id: String) -> Result<()> {
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

    /// Retrieves the [`Action`] by id.
    ///
    /// # Arguments
    /// * `action_id` - The id of the action to retrieve.
    pub async fn get(&self, action_id: String) -> Result<Option<Action>> {
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

    /// Retrieves a list of all [`Action`]s.
    ///     
    /// # Returns
    /// A list of all actions.
    pub async fn list(&self) -> Result<Vec<Action>> {
        let request = make_request_with_metadata(ListActionsRequest {}, &self.metadata_map);

        self.service_client
            .lock()
            .await
            .list_actions(request)
            .await
            .map(|r| r.into_inner().actions)
            .map_err(From::from)
    }

    /// Orders the actions.
    ///
    /// # Arguments
    /// * `private_actions_order` - The order of the private actions.
    /// * `shared_actions_order` - The order of the shared actions.
    ///
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
