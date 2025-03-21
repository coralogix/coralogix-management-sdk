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

use crate::{
    auth::AuthContext,
    error::{
        Result,
        SdkApiError,
        SdkError,
    },
    metadata::CallProperties,
    util::make_request_with_metadata,
};

use cx_api::proto::com::coralogixapis::aaa::organisations::v2::{
    CreateTeamInOrgRequest,
    CreateTeamInOrgResponse,
    DeleteTeamRequest,
    DeleteTeamResponse,
    GetTeamQuotaRequest,
    GetTeamQuotaResponse,
    GetTeamRequest,
    GetTeamResponse,
    ListTeamsRequest,
    ListTeamsResponse,
    MoveQuotaRequest,
    MoveQuotaResponse,
    SetDailyQuotaRequest,
    SetDailyQuotaResponse,
    TeamId,
    UpdateTeamRequest,
    UpdateTeamResponse,
    team_service_client::TeamServiceClient,
};
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

const TEAMS_FEATURE_GROUP_ID: &str = "aaa";

/// The Team API client.
/// Read more at https://coralogix.com/docs/user-team-management/ and https://coralogix.com/docs/user-guides/account-management/payment-and-billing/quota-management/
pub struct TeamsClient {
    metadata_map: MetadataMap,
    service_client: Mutex<TeamServiceClient<Channel>>,
}

impl TeamsClient {
    /// Creates a new client for the Teams API.
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
            metadata_map: request_metadata.to_metadata_map(),
            service_client: Mutex::new(TeamServiceClient::new(channel)),
        })
    }

    /// Creates a new Team in the organization.
    ///
    /// # Arguments
    /// * `team_name` - The name of the team.
    /// * `team_admins_email` - The email addresses of the team admins.
    /// * `daily_quota` - The daily quota for the team.
    pub async fn create(
        &self,
        team_name: String,
        team_admins_email: Vec<String>,
        daily_quota: Option<f64>,
    ) -> Result<CreateTeamInOrgResponse> {
        let request = make_request_with_metadata(
            CreateTeamInOrgRequest {
                team_name,
                team_admins_email,
                daily_quota,
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .create_team_in_org(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.aaa.organisations.v2.TeamService/CreateTeamInOrg"
                        .to_string(),
                    feature_group: TEAMS_FEATURE_GROUP_ID.into(),
                })
            })
    }

    /// Update the Team identified by its id.
    ///
    /// # Arguments
    /// * `team_id` - The id of the team to replace.
    /// * `team_name` - The name of the team.
    /// * `daily_quota` - The daily quota for the team.
    pub async fn replace(
        &self,
        team_id: u32,
        team_name: Option<String>,
        daily_quota: Option<f64>,
    ) -> Result<UpdateTeamResponse> {
        let request = make_request_with_metadata(
            UpdateTeamRequest {
                team_id: Some(TeamId { id: team_id }),
                team_name,
                daily_quota,
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .update_team(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.aaa.organisations.v2.TeamService/UpdateTeam"
                        .to_string(),
                    feature_group: TEAMS_FEATURE_GROUP_ID.into(),
                })
            })
    }

    /// Deletes the Team identified by its id.
    ///
    /// # Arguments
    /// * `team_id` - The id of the team to delete.
    pub async fn delete(&self, team_id: u32) -> Result<DeleteTeamResponse> {
        let request = make_request_with_metadata(
            DeleteTeamRequest {
                team_id: Some(TeamId { id: team_id }),
            },
            &self.metadata_map,
        );
        self.service_client
            .lock()
            .await
            .delete_team(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.aaa.organisations.v2.TeamService/DeleteTeam"
                        .to_string(),
                    feature_group: TEAMS_FEATURE_GROUP_ID.into(),
                })
            })
    }

    /// Retrieves the Team metadata.
    ///
    /// # Arguments
    /// * `team_id` - The id of the team to retrieve.
    pub async fn get(&self, team_id: u32) -> Result<GetTeamResponse> {
        let request = make_request_with_metadata(
            GetTeamRequest {
                team_id: Some(TeamId { id: team_id }),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_team(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.aaa.organisations.v2.TeamService/GetTeam"
                        .to_string(),
                    feature_group: TEAMS_FEATURE_GROUP_ID.into(),
                })
            })
    }

    /// Retrieves the Team quota.
    ///
    /// # Arguments
    /// * `team_id` - The id of the team to retrieve.
    pub async fn get_quota(&self, team_id: u32) -> Result<GetTeamQuotaResponse> {
        let request = make_request_with_metadata(
            GetTeamQuotaRequest {
                team_id: Some(TeamId { id: team_id }),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_team_quota(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.aaa.organisations.v2.TeamService/GetTeamQuota"
                        .to_string(),
                    feature_group: TEAMS_FEATURE_GROUP_ID.into(),
                })
            })
    }

    /// Sets the daily Team quota.
    ///
    /// # Arguments
    /// * `team_id` - The id of the team to retrieve.
    /// * `target_daily_quota` - The new daily quota for the team.
    pub async fn set_daily_quota(
        &self,
        team_id: u32,
        target_daily_quota: f32,
    ) -> Result<SetDailyQuotaResponse> {
        let request = make_request_with_metadata(
            SetDailyQuotaRequest {
                team_id: Some(TeamId { id: team_id }),
                target_daily_quota,
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .set_daily_quota(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.aaa.organisations.v2.TeamService/SetDailyQuota"
                        .to_string(),
                    feature_group: TEAMS_FEATURE_GROUP_ID.into(),
                })
            })
    }

    /// Moves the quota units from one team to the other.
    ///
    /// # Arguments
    /// * `source_team_id` - The id of the team to move the quota from.
    /// * `destination_team_id` - The id of the team to move the quota to.
    /// * `units_to_move` - The number of units to move.
    pub async fn move_quota(
        &self,
        source_team_id: u32,
        destination_team_id: u32,
        units_to_move: f64,
    ) -> Result<MoveQuotaResponse> {
        let request = make_request_with_metadata(
            MoveQuotaRequest {
                source_team: Some(TeamId { id: source_team_id }),
                destination_team: Some(TeamId {
                    id: destination_team_id,
                }),
                units_to_move,
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .move_quota(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.aaa.organisations.v2.TeamService/MoveQuota"
                        .to_string(),
                    feature_group: TEAMS_FEATURE_GROUP_ID.into(),
                })
            })
    }

    /// Retrieves a list of all [`TeamInfo`]s.
    ///     
    /// # Returns
    /// A list of all teams.
    pub async fn list(&self) -> Result<ListTeamsResponse> {
        let request = make_request_with_metadata(ListTeamsRequest {}, &self.metadata_map);

        self.service_client
            .lock()
            .await
            .list_teams(request)
            .await
            .map(|r| r.into_inner())
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogixapis.aaa.organisations.v2.TeamService/ListTeams"
                        .to_string(),
                    feature_group: TEAMS_FEATURE_GROUP_ID.into(),
                })
            })
    }
}
