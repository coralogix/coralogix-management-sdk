use protofetch::{LockMode, Protofetch};

fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("cargo:rerun-if-changed=build.rs");
    println!("cargo:rerun-if-changed=../Cargo.lock");
    println!("cargo:rerun-if-changed=protofetch.toml");
    println!("cargo:rerun-if-changed=protofetch.lock");

    Protofetch::builder()
        .try_build().unwrap()
        .fetch(if is_ci::cached() {
            LockMode::Locked
        } else {
            LockMode::Update
        })?;
    let building =           &[
        // #[cfg(feature = "alerts")]
        alerts_service(),
        // #[cfg(feature = "rules")]
        rules_service(),
        // #[cfg(feature = "outgoing_webhooks")]
        outgoing_webhooks_service(),
        // #[cfg(feature = "quota")]
        quota_service(),
        // #[cfg(feature = "enrichment")]
        enrichment_service(),
        // #[cfg(feature = "data_usage")]
        data_usage_service(),
        // #[cfg(feature = "dashboards")]
        dashboards_service(),
        // #[cfg(feature = "tags")]
        tags_service(),
        // #[cfg(feature = "extensions")]
        extensions_service(),
        // "proto/com/coralogixapis/events2metrics/v2/events2metrics_service.proto",
        // #[cfg(feature = "dataprime")]
        dataprime_service(),
        // #[cfg(feature = "incidents")]
        incidents_service(),
        // #[cfg(feature = "views")]
        views_service(),
        // #[cfg(feature = "scopes")]
        scopes_service(),
    ]
    .concat(); 

    tonic_build::configure()
        .include_file("coralogix_management_api_grpc.rs")
        .compile_well_known_types(true)
        .build_server(false)
        .build_client(true)
        .compile(
            building,
            &["proto/"],
        )?;
    Ok(())
}

fn tags_service() -> &'static [&'static str] {
    &["proto/com/coralogix/tags/v1/tag_service.proto"]
}

fn alerts_service() -> &'static [&'static str] {
    &["proto/com/coralogix/alerts/v2/alert_service.proto"]
}

fn rules_service() -> &'static [&'static str] {
    &["proto/com/coralogix/rules/v1/rules_service.proto"]
}

fn outgoing_webhooks_service() -> &'static [&'static str] {
    &["proto/com/coralogix/outgoing_webhooks/v1/outgoing_webhook_service.proto"]
}

fn enrichment_service() -> &'static [&'static str] {
    &["proto/com/coralogix/enrichment/v1/enrichment_service.proto"]
}

fn data_usage_service() -> &'static [&'static str] {
    &["proto/com/coralogix/datausage/v2/data_usage_service.proto"]
}

fn quota_service() -> &'static [&'static str] {
    &["proto/com/coralogix/quota/v1/policies_service.proto"]
}

fn dashboards_service() -> &'static [&'static str] {
    &[
        "proto/com/coralogix/dashboards/v1/widgets/alerts_service.proto",
        "proto/com/coralogix/dashboards/v1/widgets/announcements_service.proto",
        "proto/com/coralogix/dashboards/v1/widgets/anomalies_service.proto",
        "proto/com/coralogix/dashboards/v1/widgets/logs_activity_service.proto",
        "proto/com/coralogix/dashboards/v1/widgets/platform_overview_service.proto",
        "proto/com/coralogix/dashboards/v1/widgets/top_abnormal_errors_service.proto",
        "proto/com/coralogix/dashboards/v1/widgets/top_errored_ratio_service.proto",
        "proto/com/coralogix/dashboards/v1/widgets/top_errors_service.proto",
        "proto/com/coralogixapis/dashboards/v1/services/dashboards_service.proto",
        "proto/com/coralogixapis/dashboards/v1/services/dashboard_folders_service.proto",
    ]
}

fn extensions_service() -> &'static [&'static str] {
    &[
        "proto/com/coralogix/extensions/v1/extension_service.proto",
        "proto/com/coralogix/extensions/v1/extension_deployment_service.proto",
    ]
}

fn dataprime_service() -> &'static [&'static str] {
    &["proto/com/coralogixapis/dataprime/v1/query_service.proto"]
}

fn incidents_service() -> &'static [&'static str] {
    &["proto/com/coralogixapis/incidents/v1/incidents_service.proto"]
}

fn views_service() -> &'static [&'static str] {
    &[
        "proto/com/coralogixapis/views/v1/services/views_service.proto",
        "proto/com/coralogixapis/views/v1/services/views_folders_service.proto",
    ]
}

fn scopes_service() -> &'static [&'static str] {
    &["proto/com/coralogixapis/scopes/v1/scopes.proto"]
}
