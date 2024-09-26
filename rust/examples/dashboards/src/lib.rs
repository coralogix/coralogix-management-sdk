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

#[cfg(test)]
mod tests {
    use cx_sdk::{
        auth::AuthContext,
        client::dashboards::{
            Dashboard,
            DashboardsClient,
        },
        CoralogixRegion,
    };

    #[tokio::test]
    async fn test_dashboard_client() {
        let client = DashboardsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();
        let raw_dashboard = tokio::fs::read_to_string("dashboard.json").await.unwrap();
        let dashboard: Dashboard = serde_json::from_str(raw_dashboard.as_str()).unwrap();
        let id = dashboard.id.as_ref().unwrap().clone();
        if (client.get(id.clone()).await).is_ok() {
            let _ = client.delete(id.clone()).await.unwrap();
        }
        let _ = client.create(dashboard).await.unwrap();

        let actual_dashboard = client.get(id.clone()).await.unwrap();
        assert_eq!(actual_dashboard.dashboard.unwrap().id, Some(id.clone()));
        let _ = client.pin(id.clone()).await.unwrap();
        let _ = client.unpin(id.clone()).await.unwrap();

        let dashboards = client.list().await.unwrap();
        assert!(!dashboards.items.is_empty());
        assert!(dashboards.items.iter().any(|d| d.id == Some(id.clone())));
        let _ = client.delete(id.clone()).await.unwrap();
    }
}
