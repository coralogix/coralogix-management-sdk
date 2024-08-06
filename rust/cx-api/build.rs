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

use std::path::PathBuf;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("cargo:rerun-if-changed=build.rs");
    println!("cargo:rerun-if-changed=../Cargo.lock");
    println!("cargo:rerun-if-changed=protofetch.toml");
    println!("cargo:rerun-if-changed=protofetch.lock");
    let mut project_root = PathBuf::from(std::env::var("CARGO_MANIFEST_DIR").unwrap());
    project_root.pop();
    project_root.pop();
    println!("project_root: {:?}", project_root);
    let building = &[
        #[cfg(feature = "alerts")]
        alerts_service(),
        #[cfg(feature = "alert_scheduler")]
        alerts_scheduler_service(),
        #[cfg(feature = "rules")]
        rules_service(),
        #[cfg(feature = "outgoing_webhooks")]
        outgoing_webhooks_service(),
        #[cfg(feature = "quota")]
        quota_service(),
        #[cfg(feature = "enrichment")]
        enrichment_service(),
        #[cfg(feature = "data_usage")]
        data_usage_service(),
        #[cfg(feature = "dashboards")]
        dashboards_service(),
        #[cfg(feature = "tags")]
        tags_service(),
        #[cfg(feature = "e2m")]
        e2m_service(),
        #[cfg(feature = "dataprime")]
        dataprime_service(),
        #[cfg(feature = "views")]
        views_service(),
        #[cfg(feature = "actions")]
        actions_service(),
        #[cfg(feature = "api_keys")]
        api_keys_service(),
        #[cfg(feature = "target")]
        target_service(),
        #[cfg(feature = "metrics_configurator")]
        metrics_configurator_service(),
        #[cfg(feature = "retentions")]
        retention_service(),
        #[cfg(feature = "team")]
        team_service(),
        #[cfg(feature = "scopes")]
        scopes_service(),
        #[cfg(feature = "slos")]
        slos_service(),
        #[cfg(feature = "groups")]
        groups_service(),
        #[cfg(feature = "integrations")]
        integrations_service(),
    ]
    .concat();

    // Replacing the original prost types with the WKT crate ones will add the serde attributes without having
    // to recompile those proto files which would lead to having to use the raw Google protobuf types.
    tonic_build::configure()
        .type_attribute(".", "#[derive(serde::Serialize, serde::Deserialize)]")
        .type_attribute(".", "#[serde(rename_all = \"snake_case\")]")
        .extern_path(".google.protobuf.Any", "::prost_wkt_types::Any")
        .extern_path(".google.protobuf.Struct", "::prost_wkt_types::Struct")
        .extern_path(".google.protobuf.Timestamp", "::prost_wkt_types::Timestamp")
        .extern_path(".google.protobuf.Duration", "::prost_wkt_types::Duration")
        .extern_path(".google.protobuf.Value", "::prost_wkt_types::Value")
        .include_file("coralogix_management_api_grpc.rs")
        .build_server(false)
        .build_client(true)
        .compile(building, &["../../proto/"])?;
    Ok(())
}

fn tags_service() -> &'static [&'static str] {
    &["../../proto/com/coralogix/tags/v1/tag_service.proto"]
}

fn alerts_service() -> &'static [&'static str] {
    &["../../proto/com/coralogixapis/alerts/v3/alert_defs_service.proto"]
}

fn alerts_scheduler_service() -> &'static [&'static str] {
    &["../../proto/com/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_rule_service.proto"]
}

fn rules_service() -> &'static [&'static str] {
    &["../../proto/com/coralogix/rules/v1/rule_groups_service.proto"]
}

fn outgoing_webhooks_service() -> &'static [&'static str] {
    &["../../proto/com/coralogix/outgoing_webhooks/v1/outgoing_webhook_service.proto"]
}

fn enrichment_service() -> &'static [&'static str] {
    &[
        "../../proto/com/coralogix/enrichment/v1/enrichment_service.proto",
        "../../proto/com/coralogix/enrichment/v1/custom_enrichment_service.proto",
    ]
}

fn data_usage_service() -> &'static [&'static str] {
    &["../../proto/com/coralogix/datausage/v2/data_usage_service.proto"]
}

fn quota_service() -> &'static [&'static str] {
    &["../../proto/com/coralogix/quota/v1/policies_service.proto"]
}

fn e2m_service() -> &'static [&'static str] {
    &["../../proto/com/coralogixapis/events2metrics/v2/events2metrics_service.proto"]
}

fn dashboards_service() -> &'static [&'static str] {
    &[
    "../../proto/com/coralogixapis/dashboards/v1/services/archive_logs_data_source_service.proto",
    "../../proto/com/coralogixapis/dashboards/v1/services/archive_spans_data_source_service.proto",
    "../../proto/com/coralogixapis/dashboards/v1/services/dashboard_catalog_service.proto",
    "../../proto/com/coralogixapis/dashboards/v1/services/dashboard_folders_service.proto",
    "../../proto/com/coralogixapis/dashboards/v1/services/dashboards_service.proto",
    "../../proto/com/coralogixapis/dashboards/v1/services/dataprime_data_source_service.proto",
    "../../proto/com/coralogixapis/dashboards/v1/services/logs_data_source_service.proto",
    "../../proto/com/coralogixapis/dashboards/v1/services/metrics_data_source_service.proto",
    "../../proto/com/coralogixapis/dashboards/v1/services/spans_data_source_service.proto",
    "../../proto/com/coralogixapis/dashboards/v1/services/team_settings_service.proto",
    ]
}

fn dataprime_service() -> &'static [&'static str] {
    &["../../proto/com/coralogixapis/dataprime/v1/query_service.proto"]
}

fn views_service() -> &'static [&'static str] {
    &[
        "../../proto/com/coralogixapis/views/v1/services/views_service.proto",
        "../../proto/com/coralogixapis/views/v1/services/views_folders_service.proto",
    ]
}

fn actions_service() -> &'static [&'static str] {
    &["../../proto/com/coralogixapis/actions/v2/actions_service.proto"]
}

fn api_keys_service() -> &'static [&'static str] {
    &["../../proto/com/coralogixapis/aaa/apikeys/v3/api_keys.proto"]
}

fn target_service() -> &'static [&'static str] {
    &["../../proto/com/coralogix/archive/v2/target_service.proto"]
}

fn metrics_configurator_service() -> &'static [&'static str] {
    &["../../proto/com/coralogix/metrics/metrics-configurator.proto"]
}

fn retention_service() -> &'static [&'static str] {
    &["../../proto/com/coralogix/archive/v1/retentions_service.proto"]
}

fn team_service() -> &'static [&'static str] {
    &["../../proto/com/coralogixapis/aaa/organisations/v2/team_service.proto"]
}

fn scopes_service() -> &'static [&'static str] {
    &["../../proto/com/coralogixapis/scopes/v1/scopes.proto"]
}

fn slos_service() -> &'static [&'static str] {
    &["../../proto/com/coralogixapis/apm/services/v1/service_slo_service.proto"]
}

fn groups_service() -> &'static [&'static str] {
    &["../../proto/com/coralogix/permissions/v1/team_permissions_mgmt_service.proto"]
}

fn integrations_service() -> &'static [&'static str] {
    &["../../proto/com/coralogix/integrations/v1/integration_service.proto"]
}
