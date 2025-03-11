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
    use cx_sdk::CoralogixRegion;
    use cx_sdk::auth::AuthContext;
    use cx_sdk::client::data_usage::{
        EnrichmentMapping,
        EnrichmentType,
        Type,
    };
    use cx_sdk::client::enrichments::{
        EnrichmentsClient,
        GeoIpType,
    };

    #[tokio::test]
    async fn test_enrichments() {
        let client = EnrichmentsClient::new(
            AuthContext::from_env(),
            CoralogixRegion::from_env().unwrap(),
        )
        .unwrap();

        // Create field mappings for testing
        let field_mappings = vec![EnrichmentMapping {
            field_name: Some("coralogix.metadata.sdkId".to_string()),
            enrichment_type: Some(EnrichmentType {
                r#type: Some(Type::GeoIp(GeoIpType {
                    with_asn: Some(true),
                })),
            }),
            enriched_field_name: None,
            selected_columns: vec![],
        }];

        // Test adding enrichments
        let creation_response = client.add(field_mappings).await.unwrap();

        let enrichment_ids = creation_response
            .enrichments
            .iter()
            .map(|e| e.id)
            .collect::<Vec<_>>();

        //Test retrieving enrichments
        let enrichments_with_created_enrichments = client.list().await.unwrap();
        assert!(
            enrichments_with_created_enrichments
                .enrichments
                .iter()
                .any(|e| enrichment_ids.contains(&e.id))
        );

        //Test deleting enrichments
        let _ = client.delete(enrichment_ids.clone()).await.unwrap();

        let enrichments_after_deletion = client.list().await.unwrap();
        assert!(
            enrichments_after_deletion
                .enrichments
                .iter()
                .all(|e| !enrichment_ids.contains(&e.id))
        );
    }
}
