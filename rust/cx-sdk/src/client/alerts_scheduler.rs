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

use cx_api::proto::com::coralogixapis::alerting::alert_scheduler_rule_protobuf::v1::{
    alert_scheduler_rule_service_client::AlertSchedulerRuleServiceClient,
    filter_by_alert_scheduler_rule_ids, AlertSchedulerRuleIds,
    AlertSchedulerRuleWithActiveTimeframe, CreateAlertSchedulerRuleRequest,
    CreateBulkAlertSchedulerRuleRequest, DeleteAlertSchedulerRuleRequest,
    FilterByAlertSchedulerRuleIds, GetAlertSchedulerRuleRequest, GetBulkAlertSchedulerRuleRequest,
    UpdateAlertSchedulerRuleRequest, UpdateBulkAlertSchedulerRuleRequest,
};
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, ClientTlsConfig, Endpoint},
};

use crate::{
    auth::AuthContext,
    error::{Result, SdkError},
    metadata::CallProperties,
    util::make_request_with_metadata,
    CoralogixRegion,
};

pub use cx_api::proto::com::coralogixapis::alerting::alert_scheduler_rule_protobuf::v1::{
    filter::WhichAlerts, schedule::Scheduler, timeframe::Until, ActiveTimeframe,
    AlertSchedulerRule, AlertUniqueIds, Filter as AlertSchedulerFilter, OneTime, Schedule,
    ScheduleOperation, Timeframe,
};

/// The Alert Scheduler API client.
/// Read more at [https://coralogix.com/docs/coralogix-user-defined-alerts/]()
pub struct AlertSchedulerClient {
    metadata_map: MetadataMap,
    service_client: Mutex<AlertSchedulerRuleServiceClient<Channel>>,
}

impl AlertSchedulerClient {
    /// Creates a new client for the Alert Scheduler API.
    ///
    /// # Arguments
    /// * `auth_context` - The API key to use for authentication.
    /// * `region` - The region to connect to.
    pub fn new(region: CoralogixRegion, auth_context: AuthContext) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(&region.grpc_endpoint())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();
        let request_metadata: CallProperties = (&auth_context.team_level_api_key).into();
        Ok(Self {
            metadata_map: request_metadata.to_metadata_map(),
            service_client: Mutex::new(AlertSchedulerRuleServiceClient::new(channel)),
        })
    }

    /// Creates a new Alert Scheduler Rule
    ///
    /// # Arguments
    /// * `alert_scheduler_rule` - The [`AlertSchedulerRule`] to create.
    pub async fn create(
        &self,
        alert_scheduler_rule: AlertSchedulerRule,
    ) -> Result<AlertSchedulerRule> {
        let request = make_request_with_metadata(
            CreateAlertSchedulerRuleRequest {
                alert_scheduler_rule: Some(alert_scheduler_rule),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .create_alert_scheduler_rule(request)
            .await
            .map(|r| r.into_inner().alert_scheduler_rule.unwrap()) //There should not be a case where the result is successful but the alert_scheduler_rule is None
            .map_err(From::from)
    }

    /// Creates multiple Alert Scheduler Rules
    ///
    /// # Arguments
    /// * `alert_scheduler_rules` - The [`AlertSchedulerRule`]s to create.
    pub async fn create_bulk(
        &self,
        alert_scheduler_rules: Vec<AlertSchedulerRule>,
    ) -> Result<Vec<AlertSchedulerRule>> {
        let request = make_request_with_metadata(
            CreateBulkAlertSchedulerRuleRequest {
                create_alert_scheduler_rule_requests: alert_scheduler_rules
                    .into_iter()
                    .map(|alert_scheduler_rule| CreateAlertSchedulerRuleRequest {
                        alert_scheduler_rule: Some(alert_scheduler_rule),
                    })
                    .collect(),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .create_bulk_alert_scheduler_rule(request)
            .await
            .map(|r| {
                r.into_inner()
                    .create_suppression_responses
                    .into_iter()
                    .map(|response| response.alert_scheduler_rule.unwrap()) //There should not be a case where the result is successful but the alert_scheduler_rule is None
                    .collect()
            })
            .map_err(From::from)
    }

    /// Updates an existing Alert Scheduler Rule identified by its unique identifier.
    ///
    /// # Arguments
    /// * `alert_scheduler_rule` - The [`AlertSchedulerRule`] to update.
    pub async fn update(
        &self,
        alert_scheduler_rule: AlertSchedulerRule,
    ) -> Result<AlertSchedulerRule> {
        let request = make_request_with_metadata(
            UpdateAlertSchedulerRuleRequest {
                alert_scheduler_rule: Some(alert_scheduler_rule),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .update_alert_scheduler_rule(request)
            .await
            .map(|r| r.into_inner().alert_scheduler_rule.unwrap()) //There should not be a case where the result is successful but the alert_scheduler_rule is None
            .map_err(From::from)
    }

    /// Updates multiple existing Alert Scheduler Rules identified by their unique identifiers.
    /// # Arguments
    /// * `alert_scheduler_rules` - The [`AlertSchedulerRule`]s to update.
    pub async fn update_bulk(
        &self,
        alert_scheduler_rules: Vec<AlertSchedulerRule>,
    ) -> Result<Vec<AlertSchedulerRule>> {
        let request = make_request_with_metadata(
            UpdateBulkAlertSchedulerRuleRequest {
                update_alert_scheduler_rule_requests: alert_scheduler_rules
                    .into_iter()
                    .map(|alert_scheduler_rule| UpdateAlertSchedulerRuleRequest {
                        alert_scheduler_rule: Some(alert_scheduler_rule),
                    })
                    .collect(),
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .update_bulk_alert_scheduler_rule(request)
            .await
            .map(|r| {
                r.into_inner()
                    .update_suppression_responses
                    .into_iter()
                    .map(|response| response.alert_scheduler_rule.unwrap()) //There should not be a case where the result is successful but the alert_scheduler_rule is None
                    .collect()
            })
            .map_err(From::from)
    }

    /// Retrieves an Alert Scheduler Rule by its unique identifier.
    /// # Arguments
    /// * `id` - The unique identifier of the Alert Scheduler Rule to retrieve.
    pub async fn get(&self, id: String) -> Result<AlertSchedulerRule> {
        let request = make_request_with_metadata(
            GetAlertSchedulerRuleRequest {
                alert_scheduler_rule_id: id,
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .get_alert_scheduler_rule(request)
            .await
            .map(|r| r.into_inner().alert_scheduler_rule.unwrap())
            .map_err(From::from)
    }

    /// Retrieves multiple Alert Scheduler Rules by their unique identifiers.
    /// # Arguments
    /// * `active_time_frame` - The [`ActiveTimeframe`] to filter the Alert Scheduler Rules by.
    pub async fn get_bulk(
        &self,
        active_time_frame: ActiveTimeframe,
        alert_scheduler_rule_ids: Vec<String>,
        enabled: bool,
    ) -> Result<Vec<AlertSchedulerRuleWithActiveTimeframe>> {
        let mut all_rules = Vec::new();
        let mut next_page_token = Some("".to_string());
        while let Some(token) = next_page_token {
            let request = make_request_with_metadata(
            GetBulkAlertSchedulerRuleRequest {
                active_timeframe: Some(active_time_frame.clone()),
                enabled: Some(enabled),
                alert_scheduler_rules_ids: Some(FilterByAlertSchedulerRuleIds {
                    alert_scheduler_rule_ids: Some(filter_by_alert_scheduler_rule_ids::AlertSchedulerRuleIds::AlertSchedulerIds(
                        AlertSchedulerRuleIds {
                            alert_scheduler_rule_ids: alert_scheduler_rule_ids.clone(),
                        },
                    )),
                }),
                next_page_token: Some(token),
            },
            &self.metadata_map,
        );

            let (alert_scheduler_rules, new_token) = self
                .service_client
                .lock()
                .await
                .get_bulk_alert_scheduler_rule(request)
                .await
                .map(|r| {
                    let response = r.into_inner();
                    (response.alert_scheduler_rules, response.next_page_token)
                })
                .map_err(SdkError::from)?;
            all_rules.extend(alert_scheduler_rules);
            if new_token.is_empty() {
                next_page_token = None;
            } else {
                next_page_token = Some(new_token);
            }
        }
        Ok(all_rules)
    }

    /// Deletes an Alert Scheduler Rule by its unique identifier.
    /// # Arguments
    /// * `id` - The unique identifier of the Alert Scheduler Rule to delete.
    pub async fn delete(&self, id: String) -> Result<()> {
        let request = make_request_with_metadata(
            DeleteAlertSchedulerRuleRequest {
                alert_scheduler_rule_id: id,
            },
            &self.metadata_map,
        );

        self.service_client
            .lock()
            .await
            .delete_alert_scheduler_rule(request)
            .await
            .map(|_| ())
            .map_err(From::from)
    }
}
