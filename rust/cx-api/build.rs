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
const PROTOS_DIR: &str = "proto";

fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("cargo:rerun-if-changed=build.rs");
    println!("cargo:rerun-if-changed=../Cargo.lock");
    println!("cargo:rerun-if-changed=../../protofetch.toml");
    println!("cargo:rerun-if-changed=../../protofetch.lock");

    let mut project_root = PathBuf::from(std::env::var("CARGO_MANIFEST_DIR").unwrap());
    project_root.pop();
    project_root.pop();
    project_root.push(PROTOS_DIR);

    let root = project_root.into_os_string().into_string().unwrap();
    println!("root: {root:?}");
    let building = &[
        #[cfg(feature = "alerts")]
        alerts_service(&root),
        #[cfg(feature = "alert_scheduler")]
        alerts_scheduler_service(&root),
        #[cfg(feature = "rules")]
        rules_service(&root),
        #[cfg(feature = "outgoing_webhooks")]
        outgoing_webhooks_service(&root),
        #[cfg(feature = "quota")]
        quota_service(&root),
        #[cfg(feature = "enrichment")]
        enrichment_service(&root),
        #[cfg(feature = "data_usage")]
        data_usage_service(&root),
        #[cfg(feature = "dashboards")]
        dashboards_service(&root),
        #[cfg(feature = "e2m")]
        e2m_service(&root),
        #[cfg(feature = "views")]
        views_service(&root),
        #[cfg(feature = "actions")]
        actions_service(&root),
        #[cfg(feature = "api_keys")]
        api_keys_service(&root),
        #[cfg(feature = "target")]
        target_service(&root),
        #[cfg(feature = "metrics_configurator")]
        metrics_configurator_service(&root),
        #[cfg(feature = "retentions")]
        retention_service(&root),
        #[cfg(feature = "team")]
        team_service(&root),
        #[cfg(feature = "scopes")]
        scopes_service(&root),
        #[cfg(feature = "slos")]
        slos_service(&root),
        #[cfg(feature = "groups")]
        groups_service(&root),
        #[cfg(feature = "integrations")]
        integrations_service(&root),
        #[cfg(feature = "recording_rule_groups")]
        recording_rule_group_sets_service(&root),
        #[cfg(feature = "saml_configuration")]
        saml_configuration_service(&root),
        #[cfg(feature = "notifications")]
        notifications_service(&root),
        #[cfg(feature = "contextual_data_integrations")]
        contextual_data_integrations_service(&root),
        #[cfg(feature = "incidents")]
        incidents_service(&root),
        #[cfg(feature = "events")]
        events_service(&root),
        #[cfg(feature = "extensions")]
        extensions_service(&root),
        #[cfg(feature = "ip_access")]
        ip_access_service(&root),
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
        .out_dir("src/generated")
        .compile_protos(building, &[format!("{root}/")])?;
    Ok(())
}

fn alerts_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogixapis/alerts/v3/alert_defs_service.proto",
        root
    )]
}

fn alerts_scheduler_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_rule_service.proto",
        root
    )]
}

fn rules_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogix/rules/v1/rule_groups_service.proto",
        root
    )]
}

fn outgoing_webhooks_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogix/outgoing_webhooks/v1/outgoing_webhook_service.proto",
        root
    )]
}

fn enrichment_service(root: &str) -> Vec<String> {
    vec![
        format!(
            "{}/com/coralogix/enrichment/v1/enrichment_service.proto",
            root
        ),
        format!(
            "{}/com/coralogix/enrichment/v1/custom_enrichment_service.proto",
            root
        ),
    ]
}

fn data_usage_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogix/datausage/v2/data_usage_service.proto",
        root
    )]
}

fn quota_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogix/quota/v1/policies_service.proto",
        root
    )]
}

fn e2m_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogixapis/events2metrics/v2/events2metrics_service.proto",
        root
    )]
}

fn recording_rule_group_sets_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogixapis/metrics_rule_manager/v1/groups.proto",
        root
    )]
}

fn dashboards_service(root: &str) -> Vec<String> {
    vec![
        format!(
            "{}/com/coralogixapis/dashboards/v1/services/dashboard_catalog_service.proto",
            root
        ),
        format!(
            "{}/com/coralogixapis/dashboards/v1/services/dashboard_folders_service.proto",
            root
        ),
        format!(
            "{}/com/coralogixapis/dashboards/v1/services/dashboards_service.proto",
            root
        ),
    ]
}

fn views_service(root: &str) -> Vec<String> {
    vec![
        format!(
            "{}/com/coralogixapis/views/v1/services/views_service.proto",
            root
        ),
        format!(
            "{}/com/coralogixapis/views/v1/services/views_folders_service.proto",
            root
        ),
    ]
}

fn actions_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogixapis/actions/v2/actions_service.proto",
        root
    )]
}

fn api_keys_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogixapis/aaa/apikeys/v3/api_keys.proto",
        root
    )]
}

fn target_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogix/archive/v2/target_service.proto",
        root
    )]
}

fn metrics_configurator_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogix/metrics/metrics-configurator.proto",
        root
    )]
}

fn retention_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogix/archive/v1/retentions_service.proto",
        root
    )]
}

fn team_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogixapis/aaa/organisations/v2/team_service.proto",
        root
    )]
}

fn scopes_service(root: &str) -> Vec<String> {
    vec![format!("{}/com/coralogixapis/scopes/v1/scopes.proto", root)]
}

fn slos_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogixapis/slo/v1/slo_service.proto",
        root
    )]
}

fn groups_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogix/permissions/v1/team_permissions_mgmt_service.proto",
        root
    )]
}

fn integrations_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogix/integrations/v1/integration_service.proto",
        root
    )]
}

fn contextual_data_integrations_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogix/integrations/v1/contextual_data_integration_service.proto",
        root
    )]
}

fn saml_configuration_service(root: &str) -> Vec<String> {
    vec![format!("{}/com/coralogixapis/aaa/sso/v2/saml.proto", root)]
}

#[deprecated(
    since = "1.3.0",
    note = "This API is still in Alpha and will change. Use with care."
)]
fn notifications_service(root: &str) -> Vec<String> {
    vec![
        format!(
            "{}/com/coralogixapis/notification_center/connectors/v1/connectors_service.proto",
            root
        ),
        format!(
            "{}/com/coralogixapis/notification_center/entities/v1/entities_service.proto",
            root
        ),
        format!(
            "{}/com/coralogixapis/notification_center/notifications/v1/testing_service.proto",
            root
        ),
        format!(
            "{}/com/coralogixapis/notification_center/presets/v1/presets_service.proto",
            root
        ),
        format!(
            "{}/com/coralogixapis/notification_center/routers/v1/global_routers_service.proto",
            root
        ),
    ]
}

fn incidents_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogixapis/incidents/v1/incidents_service.proto",
        root
    )]
}

fn events_service(root: &str) -> Vec<String> {
    vec![format!(
        "{}/com/coralogixapis/events/v3/events_service.proto",
        root
    )]
}

fn extensions_service(root: &str) -> Vec<String> {
    vec![
        format!(
            "{}/com/coralogix/extensions/v1/extension_service.proto",
            root
        ),
        format!(
            "{}/com/coralogix/extensions/v1/extension_deployment_service.proto",
            root
        ),
        format!(
            "{}/com/coralogix/extensions/v1/extension_testing_service.proto",
            root
        ),
    ]
}

fn ip_access_service(root: &str) -> Vec<String> {
    vec![format!("{}/com/coralogixapis/aaa/v1/ip_access.proto", root)]
}
