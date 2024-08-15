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
use std::str::FromStr;

use cx_api::proto::com::coralogix::permissions::v1::{
    team_permissions_mgmt_service_client::TeamPermissionsMgmtServiceClient,
    AddUsersToTeamGroupRequest, CreateTeamGroupRequest, DeleteTeamGroupRequest,
    GetTeamGroupRequest, GetTeamGroupResponse, GetTeamGroupsRequest, GetTeamGroupsResponse,
    RemoveUsersFromTeamGroupRequest, ScopeFilters, UpdateTeamGroupRequest, UpdateTeamGroupResponse,
};
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, ClientTlsConfig, Endpoint},
};

use crate::{
    auth::{AuthContext, AuthData},
    error::Result,
    util::make_request_with_metadata,
    CoralogixRegion,
};

pub use cx_api::proto::com::coralogix::permissions::v1::{
    update_team_group_request::{RoleUpdates, UserUpdates},
    CreateTeamGroupResponse, RoleId, TeamGroupId, TeamId, UserId,
};

/// GroupsClient is a client for the groups service.
pub struct GroupsClient {
    service_client: Mutex<TeamPermissionsMgmtServiceClient<Channel>>,
    metadata_map: MetadataMap,
}

impl GroupsClient {
    /// Creates a new GroupsClient.
    ///
    /// # Arguments
    /// * `api_key' - The [`AuthContext`] to use for authentication.
    /// * `region` - The [`CoralogixRegion`] to connect to.
    pub fn new(auth_context: AuthContext, region: CoralogixRegion) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(region.grpc_endpoint().as_str())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();
        let auth_data: AuthData = (&auth_context.team_level_api_key).into();
        Ok(Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(TeamPermissionsMgmtServiceClient::new(channel)),
        })
    }

    #[allow(clippy::too_many_arguments)]
    /// Creates a new group.
    ///
    /// # Arguments
    /// * `name` - The name of the group.
    /// * `team_id` - The [`TeamId`] of the team the group belongs to.
    /// * `description` - The description of the group.
    /// * `external_id` - The external ID of the group.
    /// * `role_ids` - The [`RoleId`]s of the roles in the group.
    /// * `user_ids` - The [`UserId`]s of the users in the group.
    /// * `scope_filters` - The [`ScopeFilters`] of the group.
    /// * `next_gen_scope_id` - The next-gen scope ID of the group.
    pub async fn create(
        &self,
        name: String,
        team_id: TeamId,
        description: String,
        external_id: Option<String>,
        role_ids: Vec<RoleId>,
        user_ids: Vec<UserId>,
        scope_filters: Option<ScopeFilters>,
        next_gen_scope_id: Option<String>,
    ) -> Result<CreateTeamGroupResponse> {
        let request = make_request_with_metadata(
            CreateTeamGroupRequest {
                name,
                team_id: Some(team_id),
                description: Some(description),
                external_id,
                role_ids,
                user_ids,
                scope_filters,
                next_gen_scope_id,
            },
            &self.metadata_map,
        );
        Ok(self
            .service_client
            .lock()
            .await
            .create_team_group(request)
            .await?
            .into_inner())
    }

    /// Fetches the groups for an organization.
    ///
    /// # Arguments
    /// * `group_id` - The [`TeamGroupId`] of the group to fetch.
    pub async fn get(&self, group_id: TeamGroupId) -> Result<GetTeamGroupResponse> {
        let request = make_request_with_metadata(
            GetTeamGroupRequest {
                group_id: Some(group_id),
            },
            &self.metadata_map,
        );
        Ok(self
            .service_client
            .lock()
            .await
            .get_team_group(request)
            .await?
            .into_inner())
    }

    /// Fetches all groups for a team.
    ///
    /// # Arguments
    /// * `team_id` - The [`TeamId`] of the team to fetch groups for.
    pub async fn list(&self, team_id: TeamId) -> Result<GetTeamGroupsResponse> {
        let request = make_request_with_metadata(
            GetTeamGroupsRequest {
                team_id: Some(team_id),
            },
            &self.metadata_map,
        );
        Ok(self
            .service_client
            .lock()
            .await
            .get_team_groups(request)
            .await?
            .into_inner())
    }

    /// Adds users to a group.
    ///
    /// # Arguments
    /// * `group_id` - The [`TeamGroupId`] of the group to add roles to.
    /// * `user_ids` - The [`UserId`]s of the users to add to the group.
    pub async fn add_users(&self, group_id: TeamGroupId, user_ids: Vec<UserId>) -> Result<()> {
        let request = make_request_with_metadata(
            AddUsersToTeamGroupRequest {
                group_id: Some(group_id),
                user_ids,
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .add_users_to_team_group(request)
            .await?;
        Ok(())
    }

    #[allow(clippy::too_many_arguments)]
    /// Updates a group.
    /// * `group_id` - The [`TeamGroupId`] of the group to update.
    /// * `name` - The name of the group.
    /// * `team_id` - The [`TeamId`] of the team the group belongs to.
    /// * `description` - The description of the group.
    /// * `external_id` - The external ID of the group.
    /// * `role_updates` - The [`RoleUpdates`] to apply to the group.
    /// * `user_updates` - The [`UserUpdates`] to apply to the group.
    /// * `scope_filters` - The [`ScopeFilters`] of the group.
    /// * `next_gen_scope_id` - The next-gen scope ID of the group.
    pub async fn update(
        &self,
        group_id: TeamGroupId,
        name: String,
        description: String,
        external_id: Option<String>,
        role_updates: Option<RoleUpdates>,
        user_updates: Option<UserUpdates>,
        scope_filters: Option<ScopeFilters>,
        next_gen_scope_id: Option<String>,
    ) -> Result<UpdateTeamGroupResponse> {
        let request = make_request_with_metadata(
            UpdateTeamGroupRequest {
                name,
                description: Some(description),
                external_id,
                group_id: Some(group_id),
                role_updates,
                user_updates,
                scope_filters,
                next_gen_scope_id,
            },
            &self.metadata_map,
        );
        Ok(self
            .service_client
            .lock()
            .await
            .update_team_group(request)
            .await?
            .into_inner())
    }

    /// Removes users from a group.
    ///
    /// # Arguments
    /// * `group_id` - The [`TeamGroupId`] of the group to remove roles from.
    /// * `user_ids` - The [`UserId`]s of the users to remove from the group.
    pub async fn remove_users(&self, group_id: TeamGroupId, user_ids: Vec<UserId>) -> Result<()> {
        let request = make_request_with_metadata(
            RemoveUsersFromTeamGroupRequest {
                group_id: Some(group_id),
                user_ids,
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .remove_users_from_team_group(request)
            .await?;
        Ok(())
    }

    /// Deletes a group.
    ///
    /// # Arguments
    /// * `group_id` - The [`TeamGroupId`] of the group to delete.
    pub async fn delete(&self, group_id: TeamGroupId) -> Result<()> {
        let request = make_request_with_metadata(
            DeleteTeamGroupRequest {
                group_id: Some(group_id),
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .delete_team_group(request)
            .await?;
        Ok(())
    }
}
