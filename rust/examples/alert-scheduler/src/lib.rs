#[cfg(test)]
mod tests {
    use cx_sdk::{
        auth::ApiKey,
        client::alerts_scheduler::{
            AlertSchedulerClient, AlertSchedulerFilter, AlertSchedulerRule, AlertUniqueIds,
            OneTime, Schedule, ScheduleOperation, Scheduler, Timeframe, Until, WhichAlerts,
        },
        CoralogixRegion,
    };

    #[tokio::test]
    async fn test_alert_scheduler_client() {
        let api_key = std::env::var("CORALOGIX_API_KEY").unwrap();
        let alert_scheduler_client = AlertSchedulerClient::new(
            CoralogixRegion::from_env().unwrap(),
            ApiKey::from(api_key.as_str()),
        );

        let alert_schduler_rule = AlertSchedulerRule {
            unique_identifier: None,
            id: None,
            name: String::from("example"),
            description: Some(String::from("example")),
            meta_labels: vec![],
            filter: Some(AlertSchedulerFilter {
                what_expression: "source logs | filter $d.cpodId:string == '122'".into(),
                which_alerts: Some(WhichAlerts::AlertUniqueIds(AlertUniqueIds {
                    value: vec!["07d5714a-c847-4afa-87db-5e5f6986688c".into()],
                })),
            }),
            schedule: Some(Schedule {
                schedule_operation: ScheduleOperation::Mute.into(),
                scheduler: Some(Scheduler::OneTime(OneTime {
                    timeframe: Some(Timeframe {
                        start_time: String::from("2021-01-04T00:00:00.000"),
                        timezone: "UTC+2".to_string(),
                        until: Some(Until::EndTime(String::from("2025-01-01T00:00:50.000"))),
                    }),
                })),
            }),
            enabled: true,
            created_at: None,
            updated_at: None,
        };

        let created_alert_scheduler_rule = alert_scheduler_client.create(alert_schduler_rule).await;
        if let Err(e) = &created_alert_scheduler_rule {
            let err = e.to_string();
            println!("Error: {:?}", err);
        }

        assert!(created_alert_scheduler_rule.is_ok());

        let new_alert_scheduler_rule = AlertSchedulerRule {
            name: String::from("MyAlertUpdated"),
            ..created_alert_scheduler_rule.unwrap()
        };

        let updated_alert_scheduler_rule = alert_scheduler_client
            .update(new_alert_scheduler_rule)
            .await;

        assert!(updated_alert_scheduler_rule.is_ok());

        let retrieved_alert_scheduler_rule = alert_scheduler_client
            .get(
                updated_alert_scheduler_rule
                    .unwrap()
                    .unique_identifier
                    .unwrap(),
            )
            .await;

        let retrieved_alert_scheduler_rule = retrieved_alert_scheduler_rule.unwrap();

        assert!(retrieved_alert_scheduler_rule.name == "MyAlertUpdated");

        let deletion_result = alert_scheduler_client
            .delete(retrieved_alert_scheduler_rule.unique_identifier.unwrap())
            .await;

        assert!(deletion_result.is_ok());
    }
}
