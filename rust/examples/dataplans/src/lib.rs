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
    use std::time::Duration;

    use cx_sdk::{
        auth::AuthContext,
        client::data_usage::{DataUsageClient, DateRange, ScopesFilter},
        CoralogixRegion,
    };

    #[tokio::test]
    async fn test_data_usage() {
        let client = DataUsageClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        let mut logs_stream = client
            .get_logs_count(
                Some(DateRange {
                    from_date: Some(String::from("2021-01-01T00:00:00Z").parse().unwrap()),
                    to_date: Some(String::from("2021-01-02T00:00:00Z").parse().unwrap()),
                }),
                Some(Duration::from_secs(3601)),
                Some(ScopesFilter {
                    application: vec!["myapp".into()],
                    subsystem: vec!["myapp".into()],
                }),
            )
            .await
            .unwrap();

        while let Some(_) = logs_stream.message().await.unwrap() {}

        let mut spans_count = client
            .get_spans_count(
                Some(DateRange {
                    from_date: Some(String::from("2021-01-01T00:00:00Z").parse().unwrap()),
                    to_date: Some(String::from("2021-01-02T00:00:00Z").parse().unwrap()),
                }),
                Some(Duration::from_secs(3601)),
                Some(ScopesFilter {
                    application: vec!["myapp".into()],
                    subsystem: vec!["myapp".into()],
                }),
            )
            .await
            .unwrap();

        while let Some(_) = spans_count.message().await.unwrap() {}

        let mut data_usage_stream = client
            .get_data_usage(
                Some(DateRange {
                    from_date: Some(String::from("2021-01-01T00:00:00Z").parse().unwrap()),
                    to_date: Some(String::from("2021-01-02T00:00:00Z").parse().unwrap()),
                }),
                Some(Duration::from_secs(3601)),
                vec![],
                vec![],
            )
            .await
            .unwrap();
        while let Some(_) = data_usage_stream.message().await.unwrap() {}

        let _ = client
            .update_data_usage_metrics_export_status(true)
            .await
            .unwrap();

        let data_usage_metrics_export_status_response =
            client.get_data_usage_metrics_export_status().await.unwrap();

        assert_eq!(data_usage_metrics_export_status_response.enabled, true);
    }
}
