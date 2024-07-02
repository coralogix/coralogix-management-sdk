use std::str::FromStr;

use cx_sdk::{
    auth::{ApiKey, AuthData},
    com::coralogixapis::alerting::alert_scheduler_rule_protobuf::v1::{
        alert_scheduler_rule_service_client::AlertSchedulerRuleServiceClient,
        filter_by_alert_scheduler_rule_ids, ActiveTimeframe, AlertSchedulerRule,
        AlertSchedulerRuleIds, AlertSchedulerRuleWithActiveTimeframe,
        CreateAlertSchedulerRuleRequest, CreateBulkAlertSchedulerRuleRequest,
        DeleteAlertSchedulerRuleRequest, FilterByAlertSchedulerRuleIds,
        GetAlertSchedulerRuleRequest, GetBulkAlertSchedulerRuleRequest,
        UpdateAlertSchedulerRuleRequest, UpdateBulkAlertSchedulerRuleRequest,
    },
};
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
pub struct AlertSchedulerService {
    metadata_map: MetadataMap,
    service_client: Mutex<AlertSchedulerRuleServiceClient<Channel>>,
}

impl AlertSchedulerService {
    pub fn new(endpoint: &str, api_key: String) -> Self {
        let channel: Channel = Endpoint::from_str(endpoint).unwrap().connect_lazy();
        let api_key: ApiKey = api_key.into();
        let auth_data: AuthData = (&api_key).into();
        Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(AlertSchedulerRuleServiceClient::new(channel)),
        }
    }

    pub async fn create_alert_scheduler(
        &self,
        alert_scheduler_rule: AlertSchedulerRule,
    ) -> anyhow::Result<AlertSchedulerRule> {
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

    pub async fn create_bulk_alert_schedulers(
        &self,
        alert_scheduler_rules: Vec<AlertSchedulerRule>,
    ) -> anyhow::Result<Vec<AlertSchedulerRule>> {
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

    pub async fn update_alert_scheduler(
        &self,
        alert_scheduler_rule: AlertSchedulerRule,
    ) -> anyhow::Result<AlertSchedulerRule> {
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

    pub async fn update_bulk_alert_schedulers(
        &self,
        alert_scheduler_rules: Vec<AlertSchedulerRule>,
    ) -> anyhow::Result<Vec<AlertSchedulerRule>> {
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

    pub async fn get_alert_scheduler(&self, id: String) -> anyhow::Result<AlertSchedulerRule> {
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

    pub async fn get_bulk_alert_schedulers(
        &self,
        active_time_frame: ActiveTimeframe,
        alert_scheduler_rule_ids: Vec<String>,
        enabled: bool,
    ) -> anyhow::Result<Vec<AlertSchedulerRuleWithActiveTimeframe>> {
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
                .map_err(anyhow::Error::from)?;
            all_rules.extend(alert_scheduler_rules);
            if new_token.is_empty() {
                next_page_token = None;
            } else {
                next_page_token = Some(new_token);
            }
        }
        Ok(all_rules)
    }

    pub async fn delete_alert_scheduler(&self, id: String) -> anyhow::Result<()> {
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

#[cfg(test)]
mod tests {
    use cx_sdk::com::coralogixapis::alerting::alert_scheduler_rule_protobuf::v1::{
        schedule::Scheduler, timeframe::Until, AlertSchedulerRule, OneTime, Schedule,
        ScheduleOperation, Timeframe,
    };

    use crate::AlertSchedulerService;

    #[tokio::test]
    async fn test_actions_service() {
        let service = AlertSchedulerService::new(
            "https://ng-api-grpc.eu2.coralogix.com",
            "api-key".to_string(),
        );

        let alert_schduler_rule = AlertSchedulerRule {
            unique_identifier: None,
            id: None,
            name: String::from("example"),
            description: Some(String::from("example")),
            meta_labels: vec![],
            filter: None,
            schedule: Some(Schedule {
                schedule_operation: ScheduleOperation::Activate.into(),
                scheduler: Some(Scheduler::OneTime(OneTime {
                    timeframe: Some(Timeframe {
                        start_time: String::from("2021-01-04T00:00:00.000"),
                        timezone: "UTC".to_string(),
                        until: Some(Until::EndTime(String::from("2025-01-01T00:00:50.000"))),
                    }),
                })),
            }),
            enabled: true,
            created_at: None,
            updated_at: None,
        };

        let created_alert_scheduler_rule =
            service.create_alert_scheduler(alert_schduler_rule).await;

        assert!(created_alert_scheduler_rule.is_ok());

        let new_alert_scheduler_rule = AlertSchedulerRule {
            name: String::from("MyAlertUpdated"),
            ..created_alert_scheduler_rule.unwrap()
        };

        let updated_alert_scheduler_rule = service
            .update_alert_scheduler(new_alert_scheduler_rule)
            .await;

        assert!(updated_alert_scheduler_rule.is_ok());

        let retrieved_alert_scheduler_rule = service
            .get_alert_scheduler(updated_alert_scheduler_rule.unwrap().id.unwrap())
            .await;

        let retrieved_alert_scheduler_rule = retrieved_alert_scheduler_rule.unwrap();

        assert!(retrieved_alert_scheduler_rule.name == "MyAlertUpdated");

        let deletion_result = service
            .delete_alert_scheduler(retrieved_alert_scheduler_rule.id.unwrap())
            .await;

        assert!(deletion_result.is_ok());
    }
}
