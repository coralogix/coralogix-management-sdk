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

    use std::{sync::Arc, time::Duration};

    use cx_sdk::{
        CoralogixRegion,
        auth::AuthContext,
        client::slos::{
            Field, IsFilterPredicate, Metric, Predicate, RequestBasedMetricSli, Sli, Slo,
            SloClient, SloFilter, SloFilterField, SloFilterPredicate, SloFilters, SloTimeFrame,
            Window,
        },
    };
    use tokio::{join, sync::Mutex, task::JoinSet};

    #[tokio::test(flavor = "multi_thread", worker_threads = 100)]
    async fn test_slos() {
        let slos_client = SloClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();
        let slos_mtx = Arc::new(Mutex::new(slos_client));
        let all_slos: Vec<Slo> = (0..10)
            .map(|i| {Slo {
                    id: None,
                    name: "coralogix_rust_slo_example".into(),
                    description: Some("description".to_string()),
                    creator: None,
                    labels: vec![("label1".to_string(), "value1".to_string())]
                        .into_iter()
                        .collect(),
                    target_threshold_percentage: 95.0f32,
                    create_time: None,
                    update_time: None,
                    sli: Some(Sli::RequestBasedMetricSli(RequestBasedMetricSli {
                        good_events: Some(Metric {
                            query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)"
                                .to_string(),
                        }),
                        total_events: Some(Metric {
                            query: "avg(rate(cpu_usage_seconds_total[5m])) by (instance)"
                                .to_string(),
                        }),
                    })),
                    window: Some(Window::SloTimeFrame(SloTimeFrame::SloTimeFrame7Days.into())),
                    revision: None,
                    grouping: None,
                }
            })
            .collect();
        let mut set = JoinSet::new();
        for (i, slo) in all_slos.into_iter().enumerate() {
            let c = slos_mtx.clone();
            set.spawn(async move {
                let id;
                {
                    let client = c.lock().await;
                    let create_slo_response = client.create(slo).await.unwrap();
                    id = create_slo_response.slo.unwrap().id.unwrap();
                }
                println!("{i}: Created SLO with {id}");

                tokio::time::sleep(Duration::from_secs(30)).await;
                {
                    let client = c.lock().await;
                    let _ = client.delete(id.clone()).await.unwrap(); 
                }
                println!("{i}: Deleted SLO with {id}");
            });
        }

        let _ = set.join_all().await;
        // let create_slo_response = slos_client.create(slo.clone()).await.unwrap();

        // let updated_slo = Slo {
        //     name: "updated_rust_slo".to_string(),
        //     ..create_slo_response.slo.unwrap()
        // };

        // let slo_update_response = slos_client.update(updated_slo).await.unwrap();

        // let slo_filters = SloFilters {
        //     filters: vec![SloFilter {
        //         field: Some(SloFilterField {
        //             field: Some(Field::LabelName("label1".to_string())),
        //         }),
        //         predicate: Some(SloFilterPredicate {
        //             predicate: Some(Predicate::Is(IsFilterPredicate {
        //                 is: vec!["value1".into()],
        //             })),
        //         }),
        //     }],
        // };
        // let list_slos_response = slos_client.list(Some(slo_filters)).await.unwrap();
        // let slos = list_slos_response.slos;
        // assert!(slos.len() > 0);

        // assert!(slo_update_response.slo.clone().unwrap().name == "updated_rust_slo");

        // let _ = slos_client
        //     .delete(slo_update_response.slo.unwrap().id.unwrap())
        //     .await
        //     .unwrap();
    }
}
