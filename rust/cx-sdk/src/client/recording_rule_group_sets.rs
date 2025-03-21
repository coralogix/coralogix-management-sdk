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

pub use cx_api::proto::com::coralogixapis::metrics_rule_manager::v1::{
    CreateRuleGroupSet,
    CreateRuleGroupSetResult,
    DeleteRuleGroupSet,
    FetchRuleGroupResult,
    FetchRuleGroupSet,
    InRule,
    InRuleGroup,
    OutRuleGroupSet,
    RuleGroupSetListing,
    UpdateRuleGroupSet,
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

use crate::{
    CoralogixRegion,
    auth::AuthContext,
    error::{
        Result,
        SdkApiError,
        SdkError,
    },
    metadata::CallProperties,
    util::make_request_with_metadata,
};

const RECORDING_RULES_FEATURE_GROUP_ID: &str = "recording-rules";

use cx_api::proto::com::coralogixapis::metrics_rule_manager::v1::rule_group_sets_client::RuleGroupSetsClient;

/// A client for the recording rule group sets service.
pub struct RecordingRuleGroupSetsClient {
    service_client: Mutex<RuleGroupSetsClient<Channel>>,
    metadata_map: MetadataMap,
}

impl RecordingRuleGroupSetsClient {
    /// Creates a new GroupsClient.
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
            service_client: Mutex::new(RuleGroupSetsClient::new(channel)),
        })
    }

    #[allow(clippy::too_many_arguments)]
    /// Creates a new group.
    ///
    /// # Arguments
    /// * `name` - The name of the group.
    /// * `groups` - The [`InRuleGroup`]s to create.
    pub async fn create(
        &self,
        name: String,
        groups: Vec<InRuleGroup>,
    ) -> Result<CreateRuleGroupSetResult> {
        let request = make_request_with_metadata(
            CreateRuleGroupSet {
                name: Some(name),
                groups,
            },
            &self.metadata_map,
        );
        Ok(self
            .service_client
            .lock()
            .await
            .create(request)
            .await
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogix.rule_manager.groups.RuleGroupSets/Create".to_string(),
                    feature_group: RECORDING_RULES_FEATURE_GROUP_ID.to_string(),
                })
            })?
            .into_inner())
    }

    /// Fetches a recording rule group set by id.
    ///
    /// # Arguments
    /// * `group_id` - The id of the recording rule group set to fetch.
    pub async fn get(&self, id: String) -> Result<OutRuleGroupSet> {
        let request = make_request_with_metadata(FetchRuleGroupSet { id }, &self.metadata_map);
        Ok(self
            .service_client
            .lock()
            .await
            .fetch(request)
            .await
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogix.rule_manager.groups.RuleGroupSets/Fetch".to_string(),
                    feature_group: RECORDING_RULES_FEATURE_GROUP_ID.to_string(),
                })
            })?
            .into_inner())
    }

    /// Fetches all recording rule group sets.
    pub async fn list(&self) -> Result<RuleGroupSetListing> {
        let request = make_request_with_metadata((), &self.metadata_map);
        Ok(self
            .service_client
            .lock()
            .await
            .list(request)
            .await
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogix.rule_manager.groups.RuleGroupSets/List".to_string(),
                    feature_group: RECORDING_RULES_FEATURE_GROUP_ID.to_string(),
                })
            })?
            .into_inner())
    }

    /// Updates a group.
    /// * `id` - The id of the rule group set to update.
    /// * `groups` - The [`InRuleGroup`]s to update.
    pub async fn update(
        &self,
        id: String,
        groups: Vec<InRuleGroup>,
        name: Option<String>,
    ) -> Result<()> {
        let request =
            make_request_with_metadata(UpdateRuleGroupSet { id, groups, name }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .update(request)
            .await
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogix.rule_manager.groups.RuleGroupSets/Update".to_string(),
                    feature_group: RECORDING_RULES_FEATURE_GROUP_ID.to_string(),
                })
            })?;
        Ok(())
    }

    /// Deletes a group.
    ///
    /// # Arguments
    /// * `id` - The id of the recording rule group set to delete.
    pub async fn delete(&self, id: String) -> Result<()> {
        let request = make_request_with_metadata(DeleteRuleGroupSet { id }, &self.metadata_map);
        self.service_client
            .lock()
            .await
            .delete(request)
            .await
            .map_err(|status| {
                SdkError::ApiError(SdkApiError {
                    status,
                    endpoint: "/com.coralogix.rule_manager.groups.RuleGroupSets/Delete".to_string(),
                    feature_group: RECORDING_RULES_FEATURE_GROUP_ID.to_string(),
                })
            })?;
        Ok(())
    }
}
