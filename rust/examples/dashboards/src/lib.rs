#[cfg(test)]
mod tests {
    use cx_sdk::{
        client::dashboards::{Dashboard, DashboardsClient},
        CoralogixRegion,
    };

    #[tokio::test]
    async fn test_dashboard_client() {
        let client =
            DashboardsClient::new("my-api-key".into(), CoralogixRegion::from_env().unwrap())
                .unwrap();
        let id = "abcdefghijklmnopqrstx".to_string();
        let raw_dashboard = tokio::fs::read_to_string("dashboard.json").await.unwrap();
        let dashboard: Dashboard = serde_json::from_str(raw_dashboard.as_str()).unwrap();
        let _ = client.create(dashboard).await.unwrap();

        let actual_dashboard = client.get(id.clone()).await.unwrap();
        assert_eq!(actual_dashboard.dashboard.unwrap().id, Some(id.clone()));
        let _ = client.pin_dashboard(id.clone()).await.unwrap();
        let _ = client.unpin_dashboard(id.clone()).await.unwrap();

        let _ = client.delete(id.clone()).await.unwrap();
    }
}
