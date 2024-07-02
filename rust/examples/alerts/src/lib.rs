use anyhow::Result;
use cx_sdk::auth::{ApiKey, AuthData};
use cx_sdk::com::coralogixapis::alerts::v3::alert_defs_service_client::AlertDefsServiceClient;
use cx_sdk::com::coralogixapis::alerts::v3::{
    AlertDef, CreateAlertDefRequest, CreateAlertDefResponse, DeleteAlertDefRequest,
    GetAlertDefRequest, GetAlertDefResponse, ListAlertDefsRequest, ListAlertDefsResponse,
    ReplaceAlertDefRequest, ReplaceAlertDefResponse, SetActiveRequest,
};
use std::str::FromStr;
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

pub struct AlertsService {
    metadata_map: MetadataMap,
    service_client: Mutex<AlertDefsServiceClient<Channel>>,
}

impl AlertsService {
    pub fn new(endpoint: &str, api_key: String) -> Self {
        let channel: Channel = Endpoint::from_str(endpoint).unwrap().connect_lazy();
        let api_key: ApiKey = api_key.into();
        let auth_data: AuthData = (&api_key).into();
        Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(AlertDefsServiceClient::new(channel)),
        }
    }

    pub async fn get_alert(&self, alert_id: String) -> Result<GetAlertDefResponse> {
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

    pub async fn list_alerts(&self) -> Result<ListAlertDefsResponse> {
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

    pub async fn create_alert(&self, alert: AlertDef) -> Result<CreateAlertDefResponse> {
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

    pub async fn replace_alert(&self, alert: AlertDef) -> Result<ReplaceAlertDefResponse> {
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

    pub async fn delete_alert(&self, alert_id: String) -> Result<()> {
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

    pub async fn set_active(&self, id: String, active: bool) -> Result<()> {
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

#[cfg(test)]
mod tests {
    use std::collections::HashMap;

    use cx_sdk::com::coralogixapis::alerts::v3::{
        alert_def_notification::{IntegrationType, RetriggeringPeriod},
        alert_def_properties::{Schedule, TypeDefinition},
        ActivitySchedule, AlertDefNotification, AlertDefNotificationGroup, AlertDefPriority,
        AlertDefProperties, AlertDefType, LogsMoreThanTypeDefinition, NotifyOn, Recipients,
        TimeOfDay,
    };

    use super::*;

    #[tokio::test]
    async fn test_alerts() {
        let endpoint = "https://ng-api-grpc.eu2.coralogix.com";
        let api_key = "api-key".to_string();
        let alerts_service = AlertsService::new(endpoint, api_key);
        let alert = AlertDef {
            updated_time: None,
            created_time: None,
            alert_def_properties: Some(AlertDefProperties {
                name: Some("my_alert".to_string()),
                description: Some("description".to_string()),
                enabled: Some(true),
                priority: AlertDefPriority::P1.into(),
                alert_def_type: AlertDefType::LogsMoreThan.into(),
                group_by: vec![],
                incidents_settings: None,
                notification_group: Some(AlertDefNotificationGroup {
                    group_by_fields: vec!["host".into()],
                    notifications: vec![AlertDefNotification {
                        notify_on: Some(NotifyOn::TriggeredAndResolved.into()),
                        retriggering_period: Some(RetriggeringPeriod::Minutes(120)),
                        integration_type: Some(IntegrationType::Recipients(Recipients {
                            emails: vec![String::from("luigi.taglialatela@coralogix.com")],
                        })),
                    }],
                }),
                labels: HashMap::new(),
                schedule: Some(Schedule::ActiveOn(ActivitySchedule {
                    day_of_week: vec![1, 2, 3, 4, 5],
                    start_time: Some(TimeOfDay {
                        hours: 10,
                        minutes: 0,
                    }),
                    end_time: Some(TimeOfDay {
                        hours: 18,
                        minutes: 0,
                    }),
                })),
                type_definition: Some(TypeDefinition::LogsMoreThan(LogsMoreThanTypeDefinition {
                    logs_filter: None,
                    threshold: Some(100),
                    time_window: None,
                    evaluation_window: 120,
                    notification_payload_filter: vec![],
                })),
            }),
            id: None,
        };

        let created_alert = alerts_service
            .create_alert(alert.clone())
            .await
            .unwrap()
            .alert_def;

        let retrieved_alert = alerts_service
            .get_alert(created_alert.as_ref().unwrap().id.clone().unwrap())
            .await
            .unwrap()
            .alert_def;

        assert_eq!(retrieved_alert.unwrap(), created_alert.unwrap());

        let updated_alert = AlertDef {
            alert_def_properties: Some(AlertDefProperties {
                description: Some("updated description".to_string()),
                ..alert.alert_def_properties.clone().unwrap()
            }),
            ..alert
        };

        let updated_alert = alerts_service
            .replace_alert(updated_alert.clone())
            .await
            .unwrap()
            .alert_def
            .unwrap();

        assert!(
            updated_alert
                .alert_def_properties
                .unwrap()
                .description
                .unwrap()
                == "updated description"
        );

        alerts_service
            .delete_alert(updated_alert.id.unwrap())
            .await
            .unwrap();
    }
}
