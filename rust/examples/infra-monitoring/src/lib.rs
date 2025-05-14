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
    use std::collections::HashMap;

    use cx_sdk::{
        CoralogixRegion,
        auth::AuthContext,
        client::slos::{
            Metric,
            MetricSli,
            Sli,
            Slo,
            SloClient,
            SloTimeFrame,
            Window,
        },
    };

    #[tokio::test]
    async fn test_slos() {
        let slos_client = SloClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let slo = Slo {
            id: String::default(),
            name: "coralogix_rust_slo_example".into(),
            description: Some("description".to_string()),
            creator: String::default(),
            labels: HashMap::new(),
            target_threshold_percentage: 95,
            create_time: None,
            update_time: None,
            sli: Some(Sli::MetricSli(MetricSli {
                good_events: Some(Metric {
                    query: "avg(rate(cpu_usage_seconds_total[5m])) by instance".to_string(),
                }),
                total_events: Some(Metric {
                    query: "avg(rate(cpu_usage_seconds_total[5m])) by instance".to_string(),
                }),
                group_by_labels: vec![],
            })),
            window: Some(Window::SloTimeFrame(SloTimeFrame::SloTimeFrame7Days.into())),
        };

        let create_slo_response = slos_client.create(slo.clone()).await.unwrap();

        let updated_slo = Slo {
            name: "updated_rust_slo".to_string(),
            ..create_slo_response.slo.unwrap()
        };

        let slo_update_response = slos_client.update(updated_slo).await.unwrap();

        assert!(slo_update_response.slo.clone().unwrap().name == "updated_rust_slo");

        let _ = slos_client
            .delete(slo_update_response.slo.unwrap().id)
            .await
            .unwrap();
    }
}
