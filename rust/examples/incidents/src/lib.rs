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
        CoralogixRegion,
        auth::AuthContext,
        client::incidents::{
            FilterOperator,
            IncidentEventOrderByFieldType,
            IncidentEventQueryFilter,
            IncidentsClient,
            LabelsFilter,
            ListIncidentEventRequestOrderBy,
            OrderByDirection,
            PaginationRequest,
            TimeRange,
        },
    };
    use std::{
        collections::HashMap,
        hash::Hash,
    };

    #[tokio::test]
    async fn test_incidents() -> Result<(), Box<dyn std::error::Error>> {
        // Setup client
        let region = CoralogixRegion::from_env()?;
        let auth_context = AuthContext::from_env();
        let client = IncidentsClient::new(auth_context, region).await?;

        // List incidents to get initial state
        let list_response = client.list_incidents(None, None, vec![]).await?;
        let initial_count = list_response.incidents.len();

        // Get incident events
        let now = chrono::Utc::now().to_rfc3339();
        let yesterday = (chrono::Utc::now() - chrono::Days::new(1)).to_rfc3339();
        let _ = client
            .list_incident_events(
                Some(IncidentEventQueryFilter {
                    display_labels: HashMap::new(),
                    name: Some("test".to_string()),
                    is_muted: Some(false),
                    labels: Some(LabelsFilter {
                        meta_labels: vec![],
                        operator: FilterOperator::OrOrUnspecified.into(),
                    }),
                    timestamp: Some(TimeRange {
                        start_time: Some(yesterday.parse().unwrap()),
                        end_time: Some(now.parse().unwrap()),
                    }),
                    status: vec![],
                    severity: vec![],
                    contextual_labels: HashMap::new(),
                }),
                Some(PaginationRequest {
                    page_size: Some(10),
                    page_token: None,
                }),
                Some(ListIncidentEventRequestOrderBy {
                    field: IncidentEventOrderByFieldType::TimestampOrUnspecified.into(),
                    direction: OrderByDirection::Desc.into(),
                }),
            )
            .await?;

        // Get events count
        let _ = client
            .list_incident_events_total_count(Some(IncidentEventQueryFilter {
                display_labels: HashMap::new(),
                name: Some("test".to_string()),
                is_muted: Some(false),
                labels: Some(LabelsFilter {
                    meta_labels: vec![],
                    operator: FilterOperator::OrOrUnspecified.into(),
                }),
                timestamp: Some(TimeRange {
                    start_time: Some(yesterday.to_string().parse().unwrap()),
                    end_time: Some(now.to_string().parse().unwrap()),
                }),
                status: vec![],
                severity: vec![],
                contextual_labels: HashMap::new(),
            }))
            .await?;

        // Get filter values
        let _ = client
            .list_incident_events_filter_values(Some(IncidentEventQueryFilter {
                display_labels: HashMap::new(),
                name: Some("test".to_string()),
                is_muted: Some(false),
                labels: Some(LabelsFilter {
                    meta_labels: vec![],
                    operator: FilterOperator::OrOrUnspecified.into(),
                }),
                timestamp: Some(TimeRange {
                    start_time: Some(yesterday.to_string().parse().unwrap()),
                    end_time: Some(now.to_string().parse().unwrap()),
                }),
                status: vec![],
                severity: vec![],
                contextual_labels: HashMap::new(),
            }))
            .await?;

        // If there are incidents, test operations on the first one
        if initial_count > 0 {
            if let Some(incident) = list_response.incidents.first() {
                let id = incident.id.clone().unwrap();
                // Get single incident
                let get_response = client.get_incident(&id).await?;
                assert_eq!(incident.id, get_response.incident.unwrap().id);

                // Test batch get
                let batch_response = client.batch_get_incident(vec![id.clone()]).await?;
                assert_eq!(1, batch_response.incidents.len());

                // Test incident actions
                let _ = client.resolve_incidents(vec![id.clone()]).await?;
            }
        }

        // Test aggregations
        let _ = client
            .list_incident_aggregations(None, vec![], None)
            .await?;

        Ok(())
    }
}
