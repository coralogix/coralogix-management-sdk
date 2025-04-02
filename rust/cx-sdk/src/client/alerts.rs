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

use crate::error::{
    SdkApiError,
    SdkError,
};
use crate::{
    CoralogixRegion,
    auth::AuthContext,
    error::Result,
    metadata::CallProperties,
    util::make_request_with_metadata,
};
use crate::{
    SDK_VERSION,
    SDK_VERSION_HEADER_NAME,
};
use cx_api::proto::com::coralogixapis::alerts::v3::alert_defs_service_client::AlertDefsServiceClient;
use cx_api::proto::com::coralogixapis::alerts::v3::{
    CreateAlertDefRequest,
    CreateAlertDefResponse,
    DeleteAlertDefRequest,
    GetAlertDefRequest,
    GetAlertDefResponse,
    ListAlertDefsRequest,
    ListAlertDefsResponse,
    ReplaceAlertDefRequest,
    ReplaceAlertDefResponse,
    SetActiveRequest,
};

use std::collections::HashMap;
use std::str::FromStr;
use tokio::sync::Mutex;
use tonic::{
    metadata::MetadataMap,
    transport::{
        Channel,
        ClientTlsConfig,
        Endpoint,
    },
};

pub use cx_api::proto::com::coralogixapis::alerts::v3::{
    ActivitySchedule,
    ActivitySchedule as AlertDefActivitySchedule,
    AlertDef,
    AlertDefIncidentSettings,
    AlertDefNotificationGroup,
    AlertDefOverride,
    AlertDefPriority,
    AlertDefProperties,
    AlertDefType,
    AlertDefWebhooksSettings,
    AlertsOp,
    AutoRetireTimeframe,
    DayOfWeek as AlertDayOfWeek,
    DayOfWeek,
    FlowStages,
    FlowStagesGroup,
    FlowStagesGroups,
    FlowStagesGroupsAlertDefs,
    FlowType,
    IntegrationType,
    LabelFilterType,
    LabelFilters,
    LogFilterOperationType,
    LogSeverity,
    LogsAnomalyCondition,
    LogsAnomalyConditionType,
    LogsAnomalyRule,
    LogsAnomalyType,
    LogsFilter,
    LogsImmediateType,
    LogsNewValueCondition,
    LogsNewValueRule,
    LogsNewValueTimeWindow,
    LogsNewValueTimeWindowValue,
    LogsNewValueType,
    LogsRatioCondition,
    LogsRatioRules,
    LogsRatioThresholdType,
    LogsRatioTimeWindow,
    LogsSimpleFilter,
    LogsThresholdCondition,
    LogsThresholdConditionType,
    LogsThresholdRule,
    LogsThresholdType,
    LogsTimeRelativeCondition,
    LogsTimeRelativeConditionType,
    LogsTimeRelativeRule,
    LogsTimeRelativeThresholdType,
    LogsTimeWindow,
    LogsTimeWindowValue,
    LogsUniqueCountCondition,
    LogsUniqueCountRule,
    LogsUniqueCountType,
    MetricAnomalyCondition,
    MetricAnomalyRule,
    MetricAnomalyType,
    MetricMissingValues,
    MetricThresholdType,
    MetricTimeWindow,
    MetricTimeWindowValue,
    NextOp,
    NotifyOn,
    Recipients,
    TimeOfDay,
    TimeframeType,
    TracingFilter,
    TracingFilterOperationType,
    TracingImmediateType,
    TracingSimpleFilter,
    TracingThresholdCondition,
    TracingThresholdRule,
    TracingThresholdType,
    TracingTimeWindow,
    TracingTimeWindowValue,
    UndetectedValuesManagement,
    alert_def_properties::{
        Schedule,
        TypeDefinition,
    },
    alert_def_webhooks_settings::*,
    integration_type,
    logs_filter::FilterType,
    logs_time_window::Type as LogsTimeWindowType,
    metric_missing_values::MissingValues,
};

const ALERTS_FEATURE_GROUP_ID: &str = "alerts";

/// Labels added by default to each Create/Replace request
pub enum DefaultLabels {
    /// SDK version under the key `x-cx-sdk-version``
    SdkVersion,
    /// Custom labels that are added by key-value
    Custom(HashMap<String, String>),
}

/// The Alerts API client.
/// Read more at [https://coralogix.com/docs/coralogix-user-defined-alerts/]()
///
pub struct AlertsClient {
    metadata_map: MetadataMap,
    service_client: Mutex<AlertDefsServiceClient<Channel>>,
    default_labels: HashMap<String, String>,
}

impl AlertsClient {
    /// Creates a new client for the Alerts API.
    /// # Arguments
    /// * `auth_context` - The API key to use for authentication.
    /// * `region` - The region to connect to.
    /// * `default_labels` - Labels to add to each request. Overwrites keys if they already exist.
    pub fn new(
        region: CoralogixRegion,
        auth_context: AuthContext,
        default_labels: Option<DefaultLabels>,
    ) -> Result<Self> {
        let channel: Channel = Endpoint::from_str(&region.grpc_endpoint())?
            .tls_config(ClientTlsConfig::new().with_native_roots())?
            .connect_lazy();
        let request_metadata: CallProperties = (&auth_context.team_level_api_key).into();

        let default_labels: HashMap<String, String> = match default_labels.as_ref() {
            Some(labels) => match labels {
                DefaultLabels::SdkVersion => default_label_hashmap(),
                DefaultLabels::Custom(hash_map) => hash_map.clone(),
            },
            None => HashMap::new(),
        };

        Ok(Self {
            metadata_map: request_metadata.to_metadata_map(),
            service_client: Mutex::new(AlertDefsServiceClient::new(channel)),
            default_labels,
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
                .map_err(|status| {
                    SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.alerts.v3.AlertDefsService/GetAlertDef"
                            .into(),
                        feature_group: ALERTS_FEATURE_GROUP_ID.into(),
                    })
                })
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
                .map_err(|status| {
                    SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.alerts.v3.AlertDefsService/ListAlertDefs"
                            .into(),
                        feature_group: ALERTS_FEATURE_GROUP_ID.into(),
                    })
                })
        }
    }

    /// Create a new alert definition.
    /// # Arguments
    ///
    /// * `alert` - The [`AlertDef`] to create.
    pub async fn create(&self, alert: AlertDef) -> Result<CreateAlertDefResponse> {
        let request = make_request_with_metadata(
            CreateAlertDefRequest {
                alert_def_properties: alert.alert_def_properties.clone(),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();
            let mut alert = alert;
            if let Some(alert_def) = &mut alert.alert_def_properties {
                let _ = self
                    .default_labels
                    .clone()
                    .into_iter()
                    .map(|(k, v)| alert_def.entity_labels.insert(k, v));
            }
            client
                .create_alert_def(request)
                .await
                .map(|r| r.into_inner())
                .map_err(|status| {
                    SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.alerts.v3.AlertDefsService/CreateAlertDef"
                            .into(),
                        feature_group: ALERTS_FEATURE_GROUP_ID.into(),
                    })
                })
        }
    }

    /// Replace an existing alert definition.
    /// # Arguments
    /// * `alert` - The [`AlertDef`] to replace.
    pub async fn replace(&self, alert: AlertDef) -> Result<ReplaceAlertDefResponse> {
        let mut alert = alert;
        if let Some(alert_def) = &mut alert.alert_def_properties {
            let _ = self
                .default_labels
                .clone()
                .into_iter()
                .map(|(k, v)| alert_def.entity_labels.insert(k, v));
        }

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
                .map_err(|status| {
                    SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.alerts.v3.AlertDefsService/ReplaceAlertDef"
                            .into(),
                        feature_group: ALERTS_FEATURE_GROUP_ID.into(),
                    })
                })
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
                .map_err(|status| {
                    SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.alerts.v3.AlertDefsService/DeleteAlertDef"
                            .into(),
                        feature_group: ALERTS_FEATURE_GROUP_ID.into(),
                    })
                })
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
                .map_err(|status| {
                    SdkError::ApiError(SdkApiError {
                        status,
                        endpoint: "/com.coralogixapis.alerts.v3.AlertDefsService/SetActive".into(),
                        feature_group: ALERTS_FEATURE_GROUP_ID.into(),
                    })
                })
        }
    }
}

fn default_label_hashmap() -> HashMap<String, String> {
    let mut hm = HashMap::new();
    hm.insert(SDK_VERSION_HEADER_NAME.to_string(), SDK_VERSION.to_string());
    hm
}
