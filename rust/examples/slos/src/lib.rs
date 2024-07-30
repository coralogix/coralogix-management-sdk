#[cfg(test)]
mod tests {
    use cx_sdk::{
        auth::ApiKey,
        client::slos::{
            CompareType, ErrorSli, ServiceSlo, SliFilter, SliType, SloClient, SloPeriod, SloStatus,
        },
        CoralogixRegion,
    };

    #[tokio::test]
    async fn slos_test() {
        let slos_client = SloClient::new(
            ApiKey::from_env().unwrap(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let slo = ServiceSlo {
            id: None,
            name: Some("coralogix_slo_example".to_string()),
            service_name: Some("service_name".to_string()),
            status: SloStatus::Ok.into(),
            description: Some("description".to_string()),
            target_percentage: Some(30),
            created_at: None,
            remaining_error_budget_percentage: None,
            filters: vec![],
            period: SloPeriod::SloPeriod7Days.into(),
            sli_type: Some(SliType::ErrorSli(ErrorSli {})),
        };

        let create_slo_response = slos_client.create(slo.clone()).await.unwrap();

        let _ = slos_client
            .get(create_slo_response.slo.clone().unwrap().id.unwrap())
            .await
            .unwrap();

        let updated_slo = ServiceSlo {
            name: Some("updated_slo".to_string()),
            ..create_slo_response.slo.unwrap()
        };

        let slo_update_response = slos_client.update(updated_slo).await.unwrap();

        assert!(slo_update_response.slo.clone().unwrap().name.unwrap() == "updated_slo");

        let _ = slos_client
            .delete(slo_update_response.slo.unwrap().id.unwrap())
            .await
            .unwrap();
    }

    #[tokio::test]
    async fn slos_test_with_filters() {
        let slos_client = SloClient::new(
            ApiKey::from_env().unwrap(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let slo = ServiceSlo {
            id: None,
            name: Some("coralogix_slo_example".to_string()),
            service_name: Some("service_name".to_string()),
            status: SloStatus::Ok.into(),
            description: Some("description".to_string()),
            target_percentage: Some(30),
            created_at: None,
            remaining_error_budget_percentage: None,
            filters: vec![SliFilter {
                field: Some("severity".to_string()),
                compare_type: CompareType::Is.into(),
                field_values: vec!["error".to_string(), "warning".to_string()],
            }],
            period: SloPeriod::SloPeriod7Days.into(),
            sli_type: Some(SliType::ErrorSli(ErrorSli {})),
        };

        let create_slo_response = slos_client.create(slo.clone()).await.unwrap();

        let _ = slos_client
            .get(create_slo_response.slo.clone().unwrap().id.unwrap())
            .await
            .unwrap();

        let updated_slo = ServiceSlo {
            name: Some("updated_slo".to_string()),
            ..create_slo_response.slo.unwrap()
        };

        let slo_update_response = slos_client.update(updated_slo).await.unwrap();

        assert!(slo_update_response.slo.clone().unwrap().name.unwrap() == "updated_slo");

        let _ = slos_client
            .delete(slo_update_response.slo.unwrap().id.unwrap())
            .await
            .unwrap();
    }
}
