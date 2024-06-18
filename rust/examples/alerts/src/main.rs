use anyhow::Result;
use cx_sdk::auth::{ApiKey, AuthData};
use cx_sdk::com::coralogixapis::alerts::v3::alert_notification::{
    IntegrationType, RetriggeringPeriod,
};
use cx_sdk::com::coralogixapis::alerts::v3::alert_properties::{
    AlertSchedule, AlertTypeDefinition,
};
use cx_sdk::com::coralogixapis::alerts::v3::{
    alerts_service_client::AlertsServiceClient, AlertProperties,
};
use cx_sdk::com::coralogixapis::alerts::v3::{
    ActivitySchedule, Alert, AlertNotification, AlertNotificationGroup, AlertPriority,
    AlertQueryFilter, AlertType, BatchGetAlertRequest, BatchGetAlertResponse, CreateAlertRequest,
    CreateAlertResponse, DeleteAlertRequest, GetAlertRequest, GetAlertResponse, GetLimitsRequest,
    GetLimitsResponse, ListAlertsRequest, ListAlertsResponse, LogsMoreThanAlertTypeDefinition,
    NotifyOn, OrderBy, Recipients, ReplaceAlertRequest, ReplaceAlertResponse, SetActiveRequest,
    TimeOfDay, ValidateAlertRequest,
};
use std::collections::HashMap;
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
    service_client: Mutex<AlertsServiceClient<Channel>>,
}

impl AlertsService {
    pub fn new(endpoint: &str, api_key: String) -> Self {
        let channel: Channel = Endpoint::from_str(endpoint).unwrap().connect_lazy();
        let api_key: ApiKey = api_key.into();
        let auth_data: AuthData = (&api_key).into();
        Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(AlertsServiceClient::new(channel)),
        }
    }

    pub async fn get_alert(&self, alert_id: String) -> Result<GetAlertResponse> {
        let request =
            make_request_with_metadata(GetAlertRequest { id: Some(alert_id) }, &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .get_alert(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    pub async fn batch_get_alert(&self, alert_ids: Vec<String>) -> Result<BatchGetAlertResponse> {
        let request =
            make_request_with_metadata(BatchGetAlertRequest { ids: alert_ids }, &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .batch_get_alert(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    pub async fn list_alerts(
        &self,
        filter: Option<AlertQueryFilter>,
        order_bys: Vec<OrderBy>,
    ) -> Result<ListAlertsResponse> {
        let request =
            make_request_with_metadata(ListAlertsRequest { filter, order_bys }, &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .list_alerts(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    pub async fn create_alert(&self, alert: Alert) -> Result<CreateAlertResponse> {
        let request = make_request_with_metadata(
            CreateAlertRequest {
                alert_properties: alert.properties,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .create_alert(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    pub async fn replace_alert(&self, alert: Alert) -> Result<ReplaceAlertResponse> {
        let request = make_request_with_metadata(
            ReplaceAlertRequest {
                alert_properties: alert.properties,
                id: alert.id,
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .replace_alert(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    pub async fn delete_alert(&self, alert_id: String) -> Result<()> {
        let request = make_request_with_metadata(
            DeleteAlertRequest { id: Some(alert_id) },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();
            client
                .delete_alert(request)
                .await
                .map(|_| ())
                .map_err(From::from)
        }
    }

    pub async fn validate_alert(&self, alert: Alert) -> Result<()> {
        let request = make_request_with_metadata(
            ValidateAlertRequest { alert: Some(alert) },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .validate_alert(request)
                .await
                .map(|_| ())
                .map_err(From::from)
        }
    }

    pub async fn get_limits(&self) -> Result<GetLimitsResponse> {
        let request = make_request_with_metadata(GetLimitsRequest {}, &self.metadata_map);
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .get_limits(request)
                .await
                .map(|r| r.into_inner())
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

#[tokio::main]
async fn main() {
    let endpoint = "https://ng-api-grpc.eu2.coralogix.com";
    let api_key = "api-key".to_string();
    let alerts_service = AlertsService::new(endpoint, api_key);
    let alert = Alert {
        updated_time: None,
        created_time: None,
        properties: Some(AlertProperties {
            name: Some("my_alert".to_string()),
            description: Some("description".to_string()),
            enabled: Some(true),
            alert_priority: AlertPriority::P1.into(),
            alert_type: AlertType::LogsMoreThan.into(),
            alert_group_bys: vec![],
            incidents_settings: None,
            notification_group: Some(AlertNotificationGroup {
                group_by_fields: vec!["host".into()],
                notifications: vec![AlertNotification {
                    notify_on: Some(NotifyOn::TriggeredAndResolved.into()),
                    retriggering_period: Some(RetriggeringPeriod::Minutes(120)),
                    integration_type: Some(IntegrationType::Recipients(Recipients {
                        emails: vec![String::from("luigi.taglialatela@coralogix.com")],
                    })),
                }],
            }),
            labels: HashMap::new(),
            alert_schedule: Some(AlertSchedule::ActiveOn(ActivitySchedule {
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
            alert_type_definition: Some(AlertTypeDefinition::LogsMoreThan(
                LogsMoreThanAlertTypeDefinition {
                    logs_filter: None,
                    threshold: Some(100),
                    time_window: None,
                    evaluation_window: 120,
                    notification_payload_filter: vec![],
                },
            )),
        }),
        id: None,
    };

    let alert_validation_result = alerts_service.validate_alert(alert.clone()).await;
    assert!(alert_validation_result.is_ok());

    let created_alert = alerts_service
        .create_alert(alert.clone())
        .await
        .unwrap()
        .alert;

    let retrieved_alert = alerts_service
        .get_alert(created_alert.as_ref().unwrap().id.clone().unwrap())
        .await
        .unwrap()
        .alert;

    assert_eq!(retrieved_alert.unwrap(), created_alert.unwrap());

    let updated_alert = Alert {
        properties: Some(AlertProperties {
            description: Some("updated description".to_string()),
            ..alert.properties.clone().unwrap()
        }),
        ..alert
    };

    let updated_alert = alerts_service
        .replace_alert(updated_alert.clone())
        .await
        .unwrap()
        .alert
        .unwrap();

    assert!(updated_alert.properties.unwrap().description.unwrap() == "updated description");

    alerts_service
        .delete_alert(updated_alert.id.unwrap())
        .await
        .unwrap();
}
