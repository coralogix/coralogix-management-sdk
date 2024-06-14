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
use tokio::fs;
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

pub struct DashboardService {
    metadata_map: MetadataMap,
    service_client: Mutex<DashboardsServiceClient<Channel>>,
}

impl DashboardService {
    pub fn new(endpoint: &str, api_key: String) -> Self {
        let channel: Channel = Endpoint::from_str(endpoint).unwrap().connect_lazy();
        let api_key: ApiKey = api_key.into();
        let auth_data: AuthData = (&api_key).into();
        Self {
            metadata_map: auth_data.to_metadata_map(),
            service_client: Mutex::new(DashboardsServiceClient::new(channel)),
        }
    }

    pub async fn create_dashboard(&self, dashboard: Dashboard) -> Result<CreateDashboardResponse> {
        let request: Request<CreateDashboardRequest> = make_request_with_metadata(
            CreateDashboardRequest {
                request_id: None,
                dashboard: Some(dashboard),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .create_dashboard(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    pub async fn delete_dashboard(&self, dashboard_id: String) -> Result<DeleteDashboardResponse> {
        let request: Request<DeleteDashboardRequest> = make_request_with_metadata(
            DeleteDashboardRequest {
                request_id: None,
                dashboard_id: Some(dashboard_id),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .delete_dashboard(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    pub async fn replace_dashboard(
        &self,
        dashboard: Dashboard,
    ) -> Result<ReplaceDashboardResponse> {
        let request: Request<ReplaceDashboardRequest> = make_request_with_metadata(
            ReplaceDashboardRequest {
                request_id: None,
                dashboard: Some(dashboard),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .replace_dashboard(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    pub async fn get_dashboard(&self, dashboard_id: String) -> Result<GetDashboardResponse> {
        let request: Request<GetDashboardRequest> = make_request_with_metadata(
            GetDashboardRequest {
                dashboard_id: Some(dashboard_id),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .get_dashboard(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    pub async fn pin_dashboard(&self, dashboard_id: String) -> Result<PinDashboardResponse> {
        let request: Request<PinDashboardRequest> = make_request_with_metadata(
            PinDashboardRequest {
                request_id: None,
                dashboard_id: Some(dashboard_id),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .pin_dashboard(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    async fn unpin_dashboard(&self, dashboard_id: String) -> Result<UnpinDashboardResponse> {
        let request: Request<UnpinDashboardRequest> = make_request_with_metadata(
            UnpinDashboardRequest {
                request_id: None,
                dashboard_id: Some(dashboard_id),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .unpin_dashboard(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }

    #[allow(dead_code)]
    async fn assign_dashboard_to_folder(
        &self,
        dashboard_id: String,
        folder_id: String,
    ) -> Result<AssignDashboardFolderResponse> {
        let request: Request<AssignDashboardFolderRequest> = make_request_with_metadata(
            AssignDashboardFolderRequest {
                request_id: None,
                dashboard_id: Some(dashboard_id),
                folder_id: Some(folder_id),
            },
            &self.metadata_map,
        );
        {
            let mut client = self.service_client.lock().await.clone();

            client
                .assign_dashboard_folder(request)
                .await
                .map(|r| r.into_inner())
                .map_err(From::from)
        }
    }
}

#[tokio::main]
async fn main() {
    let svc = DashboardService::new("https://ng-api-grpc.coralogix.com", "my-api-key".into());
    let id = "abcdefghijklmnopqrstx".to_string();
    let raw_dashboard = fs::read_to_string("dashboard.json").await.unwrap();
    let dashboard: Dashboard = serde_json::from_str(raw_dashboard.as_str()).unwrap();
    let _ = svc.create_dashboard(dashboard).await.unwrap();

    let actual_dashboard = svc.get_dashboard(id.clone()).await.unwrap();
    assert_eq!(actual_dashboard.dashboard.unwrap().id, Some(id.clone()));
    let _ = svc.pin_dashboard(id.clone()).await.unwrap();
    let _ = svc.unpin_dashboard(id.clone()).await.unwrap();

    let _ = svc.delete_dashboard(id.clone()).await.unwrap();
}
