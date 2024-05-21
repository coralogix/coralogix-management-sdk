use protofetch::{LockMode, Protofetch};


fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("cargo:rerun-if-changed=build.rs");
    println!("cargo:rerun-if-changed=../Cargo.lock");
    println!("cargo:rerun-if-changed=protofetch.toml");
    println!("cargo:rerun-if-changed=protofetch.lock");

    Protofetch::builder()
        .try_build()?
        .fetch(if is_ci::cached() {
            LockMode::Locked
        } else {
            LockMode::Update
        })?;
    tonic_build::configure()
        .include_file("coralogix_management_api_grpc.rs")
        .compile_well_known_types(true)
        .build_server(false)
        .build_client(true)
        .compile(
            &[
                "proto/com/coralogix/alerts/v2/alert_service.proto",
                "proto/com/coralogix/rules/v1/rule_groups_service.proto",
                "proto/com/coralogix/outgoing_webhooks/v1/outgoing_webhook_service.proto",
                "proto/com/coralogix/quota/v1/policies_service.proto",
                "proto/com/coralogix/enrichment/v1/enrichment_service.proto",
                "proto/com/coralogix/datausage/v2/data_usage_service.proto",
                "proto/com/coralogix/dashboards/v1/widgets/alerts_service.proto",
                "proto/com/coralogix/dashboards/v1/widgets/announcements_service.proto",
                "proto/com/coralogix/dashboards/v1/widgets/anomalies_service.proto",
                "proto/com/coralogix/dashboards/v1/widgets/logs_activity_service.proto",
                "proto/com/coralogix/dashboards/v1/widgets/platform_overview_service.proto",
                "proto/com/coralogix/dashboards/v1/widgets/top_abnormal_errors_service.proto",
                "proto/com/coralogix/dashboards/v1/widgets/top_errored_ratio_service.proto",
                "proto/com/coralogix/dashboards/v1/widgets/top_errors_service.proto",
                "proto/com/coralogix/tags/v1/tag_service.proto",
                "proto/com/coralogix/extensions/v1/extension_service.proto",
                "proto/com/coralogix/extensions/v1/extension_deployment_service.proto",
                "proto/com/coralogixapis/dashboards/v1/services/dashboards_service.proto",
                "proto/com/coralogixapis/dashboards/v1/services/dashboard_folders_service.proto",
                // "proto/com/coralogixapis/events2metrics/v2/events2metrics_service.proto",
                "proto/com/coralogixapis/dataprime/v1/query_service.proto",
                "proto/com/coralogixapis/incidents/v1/incidents_service.proto",
                "proto/com/coralogixapis/views/v1/services/views_service.proto",
                "proto/com/coralogixapis/views/v1/services/views_folders_service.proto",
                "proto/com/coralogixapis/scopes/v1/scopes.proto",
            ],
            &["proto/"],
        )?;
    Ok(())
}
