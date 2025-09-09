// Copyright 2025 Coralogix Ltd.
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
        CoralogixRegion,
        auth::AuthContext,
        client::events::{
            EventsClient,
            EventsFilter,
            PaginationRequest,
            TimestampRange,
        },
    };

    #[tokio::test]
    async fn test_events() {
        let client = EventsClient::new(
            CoralogixRegion::from_env().unwrap(),
            AuthContext::from_env(),
        )
        .unwrap();

        let now = chrono::Utc::now();
        let yesterday = now - chrono::Duration::days(1);

        let filter = EventsFilter {
            cx_event_labels_filters: None,
            timestamp: Some(TimestampRange {
                from: Some(String::from(yesterday.to_rfc3339()).parse().unwrap()),
                to: Some(String::from(now.to_rfc3339()).parse().unwrap()),
            }),
            cx_event_types: vec![],
            cx_event_keys: vec![],
            cx_event_metadata_filters: None,
        };
        let pagination = PaginationRequest {
            page_size: Some(10),
            page_token: None,
        };

        // Create the event
        let _ = client
            .list(Some(filter.clone()), vec![], Some(pagination))
            .await
            .unwrap();

        let _ = client.list_count(Some(filter.clone())).await.unwrap();
        let _ = client.get_statistics(Some(filter.clone())).await.unwrap();
    }
}
