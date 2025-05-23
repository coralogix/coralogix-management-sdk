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

use std::{
    collections::HashMap,
    str::FromStr,
};

use crate::{
    auth::AuthContext,
    com::coralogixapis::actions::v2::{
        CreateActionRequest,
        CreateActionResponse,
        DeleteActionRequest,
        GetActionRequest,
        ListActionsRequest,
        OrderActionsRequest,
        ReplaceActionRequest,
        ReplaceActionResponse,
        actions_service_client::ActionsServiceClient,
    },
    error::{
        Result,
        SdkApiError,
    },
    metadata::CallProperties,
    util::make_request_with_metadata,
};

pub use crate::com::coralogixapis::actions::v2::{
    Action,
    SourceType,
};

use cx_api::proto::com::coralogixapis::actions::v2::DeleteActionResponse;
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{
        Channel,
        ClientTlsConfig,
        Endpoint,
    },
};

use crate::CoralogixRegion;

const ACTIONS_FEATURE_GROUP_ID: &str = "actions";

/// The Actions API client.
/// Read more at <https://coralogix.com/docs/coralogix-action-extension/>
pub struct ActionsClient {
    teams_level_metadata_map: MetadataMap,
    service_client: Mutex<ActionsServiceClient<Channel>>,
}

impl ActionsClient {
    /// Creates a new client for the Actions API.
    ///
    /// # Arguments
    /// * `auth_context` - The [`AuthContext`] to use for authentication.
    /// * `region` - The [`CoralogixRegion`] to connect to.
    pub fn new(auth_context: AuthContext, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.grpc_endpoint().as_str())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();
        let request_metadata: CallProperties = (&auth_context.team_level_api_key).into();
        Ok(Self {
            teams_level_metadata_map: request_metadata.to_metadata_map(),
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
            &self.teams_level_metadata_map,
        );
        Ok(self
            .service_client
            .lock()
            .await
            .create_action(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.actions.v2.ActionsService/CreateAction".into(),
                feature_group: ACTIONS_FEATURE_GROUP_ID.into(),
            })?
            .into_inner())
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
            &self.teams_level_metadata_map,
        );
        Ok(self
            .service_client
            .lock()
            .await
            .replace_action(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.actions.v2.ActionsService/ReplaceAction".to_string(),
                feature_group: ACTIONS_FEATURE_GROUP_ID.to_string(),
            })?
            .into_inner())
    }

    /// Deletes the [`Action`] identified by its id.
    ///
    /// # Arguments
    /// * `action_id` - The id of the action to delete.
    pub async fn delete(&self, action_id: String) -> Result<DeleteActionResponse> {
        let request = make_request_with_metadata(
            DeleteActionRequest {
                id: Some(action_id),
            },
            &self.teams_level_metadata_map,
        );
        Ok(self
            .service_client
            .lock()
            .await
            .delete_action(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.actions.v2.ActionsService/DeleteAction".to_string(),
                feature_group: ACTIONS_FEATURE_GROUP_ID.to_string(),
            })?
            .into_inner())
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
            &self.teams_level_metadata_map,
        );

        Ok(self
            .service_client
            .lock()
            .await
            .get_action(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.actions.v2.ActionsService/GetAction".to_string(),
                feature_group: ACTIONS_FEATURE_GROUP_ID.to_string(),
            })?
            .into_inner()
            .action)
    }

    /// Retrieves a list of all [`Action`]s.
    ///     
    /// # Returns
    /// A list of all actions.
    pub async fn list(&self) -> Result<Vec<Action>> {
        let request =
            make_request_with_metadata(ListActionsRequest {}, &self.teams_level_metadata_map);

        Ok(self
            .service_client
            .lock()
            .await
            .list_actions(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.actions.v2.ActionsService/ListActions".to_string(),
                feature_group: ACTIONS_FEATURE_GROUP_ID.to_string(),
            })?
            .into_inner()
            .actions)
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
            &self.teams_level_metadata_map,
        );

        self.service_client
            .lock()
            .await
            .order_actions(request)
            .await
            .map_err(|status| SdkApiError {
                status,
                endpoint: "/com.coralogixapis.actions.v2.ActionsService/OrderActions".to_string(),
                feature_group: ACTIONS_FEATURE_GROUP_ID.to_string(),
            })?;
        Ok(())
    }
}
