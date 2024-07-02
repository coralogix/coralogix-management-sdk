use anyhow::Result;
use cx_sdk::auth::{ApiKey, AuthData};
use cx_sdk::com::coralogixapis::dashboards::v1::ast::Dashboard;
use cx_sdk::com::coralogixapis::dashboards::v1::services::dashboards_service_client::DashboardsServiceClient;
use cx_sdk::com::coralogixapis::dashboards::v1::services::{
    AssignDashboardFolderRequest, AssignDashboardFolderResponse, CreateDashboardRequest,
    CreateDashboardResponse, DeleteDashboardRequest, DeleteDashboardResponse, GetDashboardRequest,
    GetDashboardResponse, PinDashboardRequest, PinDashboardResponse, ReplaceDashboardRequest,
    ReplaceDashboardResponse, UnpinDashboardRequest, UnpinDashboardResponse,
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


#[cfg(test)]
mod tests {
    #[tokio::test]
    async fn test_dashboard_service() {
        let svc =
            super::DashboardService::new("my-api-key".into(), CoralogixRegion::from_env().unwrap()).unwrap();
        let id = "abcdefghijklmnopqrstx".to_string();
        let raw_dashboard = tokio::fs::read_to_string("dashboard.json").await.unwrap();
        let dashboard: super::Dashboard = serde_json::from_str(raw_dashboard.as_str()).unwrap();
        let _ = svc.create_dashboard(dashboard).await.unwrap();

        let actual_dashboard = svc.get_dashboard(id.clone()).await.unwrap();
        assert_eq!(actual_dashboard.dashboard.unwrap().id, Some(id.clone()));
        let _ = svc.pin_dashboard(id.clone()).await.unwrap();
        let _ = svc.unpin_dashboard(id.clone()).await.unwrap();

        let _ = svc.delete_dashboard(id.clone()).await.unwrap();
    }
}
