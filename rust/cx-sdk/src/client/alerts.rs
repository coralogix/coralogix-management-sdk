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

use crate::{
    auth::{AuthContext, AuthData},
    error::Result,
    util::make_request_with_metadata,
    CoralogixRegion,
};
use cx_api::proto::com::coralogixapis::alerts::v3::alert_defs_service_client::AlertDefsServiceClient;
use cx_api::proto::com::coralogixapis::alerts::v3::{
    CreateAlertDefRequest, CreateAlertDefResponse, DeleteAlertDefRequest, GetAlertDefRequest,
    GetAlertDefResponse, ListAlertDefsRequest, ListAlertDefsResponse, ReplaceAlertDefRequest,
    ReplaceAlertDefResponse, SetActiveRequest,
};
use std::str::FromStr;
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{Channel, ClientTlsConfig, Endpoint},
};

pub use cx_api::proto::com::coralogixapis::alerts::v3::{
    alert_def_advanced_target_settings::RetriggeringPeriod,
    alert_def_notification_group::Targets,
    alert_def_properties::{Schedule, TypeDefinition},
    integration_type,
    logs_filter::FilterType,
    logs_time_window::Type as LogsTimeWindowType,
    ActivitySchedule, AlertDef, AlertDefAdvancedTargetSettings, AlertDefAdvancedTargets,
    AlertDefNotificationGroup, AlertDefPriority, AlertDefProperties, AlertDefTargetSimple,
    AlertDefType, DayOfWeek, EvaluationWindow, IntegrationType, LabelFilterType, LabelFilters,
    LogFilterOperationType, LogSeverity, LogsFilter, LogsMoreThanTypeDefinition, LogsTimeWindow,
    LogsTimeWindowValue, LuceneFilter, NotifyOn, Recipients, TimeOfDay,
};

/// The Alerts API client.
/// Read more at [https://coralogix.com/docs/coralogix-user-defined-alerts/]()
pub struct AlertsClient {
    metadata_map: MetadataMap,
    service_client: Mutex<AlertDefsServiceClient<Channel>>,
}

impl AlertsClient {
    /// Creates a new client for the Alerts API.
    /// # Arguments
    /// * `auth_context` - The API key to use for authentication.
    /// * `region` - The region to connect to.
    pub fn new(region: CoralogixRegion, auth_context: AuthContext) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(&region.grpc_endpoint())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();
        let auth_data: AuthData = (&auth_context.team_level_api_key).into();
        Ok(Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(AlertDefsServiceClient::new(channel)),
        })
    }

    /// Get an alert definition by ID.
    ///
    /// # Arguments
    /// * `alert_id` - The ID of the alert definition to get.
    pub async fn get(&self, alert_id: String) -> Result<GetAlertDefResponse> {
        let request = make_request_with_metadata(
            GetAlertDefRequest { id: Some(alert_id) },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .get_alert_def(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// List all alert definitions.
    pub async fn list(&self) -> Result<ListAlertDefsResponse> {
        let request = make_request_with_metadata(ListAlertDefsRequest {}, &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .list_alert_defs(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Create a new alert definition.
    /// # Arguments
    ///
    /// * `alert` - The [`AlertDef`] to create.
    pub async fn create(&self, alert: AlertDef) -> Result<CreateAlertDefResponse> {
        let request = make_request_with_metadata(
            CreateAlertDefRequest {
                alert_def_properties: alert.alert_def_properties,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .create_alert_def(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Replace an existing alert definition.
    /// # Arguments
    /// * `alert` - The [`AlertDef`] to replace.
    pub async fn replace(&self, alert: AlertDef) -> Result<ReplaceAlertDefResponse> {
        let request = make_request_with_metadata(
            ReplaceAlertDefRequest {
                alert_def_properties: alert.alert_def_properties,
                id: alert.id,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .replace_alert_def(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    /// Delete an alert definition by ID.
    /// # Arguments
    /// * `alert_id` - The ID of the alert definition to delete.
    pub async fn delete(&self, alert_id: String) -> Result<()> {
        let request = make_request_with_metadata(
            DeleteAlertDefRequest { id: Some(alert_id) },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();
            client
                .delete_alert_def(request)
                .await
                .map(|_| ())
                .map_err(From::from)
        }
    }

    /// Set the active status of an alert definition.
    /// # Arguments
    /// * `id` - The ID of the alert definition to set the active status of.
    /// * `active` - The active status to set.
    pub async fn set(&self, id: String, active: bool) -> Result<()> {
        let request = make_request_with_metadata(
            SetActiveRequest {
                id: Some(id),
                active: Some(active),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();
            client
                .set_active(request)
                .await
                .map(|_| ())
                .map_err(From::from)
        }
    }
}
