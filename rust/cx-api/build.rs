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
        #[cfg(feature = "extensions")]
        extensions_service(),
        // "../../proto/com/coralogixapis/events2metrics/v2/events2metrics_service.proto",
        #[cfg(feature = "dataprime")]
        dataprime_service(),
        #[cfg(feature = "views")]
        views_service(),
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
    &["../../proto/com/coralogixapis/alerts/v3/alerts_service.proto"]
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

fn dashboards_service() -> &'static [&'static str] {
    &[
    "../../proto/com/coralogixapis/dashboards/v1/ast/annotation.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/dashboard.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/filter.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/folder_path.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/layout.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/variable.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/widget.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/widgets/bar_chart.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/widgets/common/colors_by.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/widgets/common/data_mode_type.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/widgets/common/legend.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/widgets/common/queries.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/widgets/common/scale.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/widgets/common/sort_by.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/widgets/common/units.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/widgets/data_table.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/widgets/gauge.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/widgets/horizontal_bar_chart.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/widgets/line_chart.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/widgets/markdown.proto",
    "../../proto/com/coralogixapis/dashboards/v1/ast/widgets/pie_chart.proto",
    "../../proto/com/coralogixapis/dashboards/v1/audit_log.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/annotation_event.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/dataprime_result.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/folder.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/group.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/group_limit.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/grouped_series.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/labelled_value.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/log_severity_level.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/logs_aggregation.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/observation_field.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/order_direction.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/ordering_field.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/pagination.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/query.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/span_field.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/spans_aggregation.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/time_frame.proto",
    "../../proto/com/coralogixapis/dashboards/v1/common/time_series.proto",
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
    "../../proto/com/coralogixapis/dashboards/v1/types.proto",
    ]
}

fn extensions_service() -> &'static [&'static str] {
    &[
        "../../proto/com/coralogix/extensions/v1/extension_service.proto",
        "../../proto/com/coralogix/extensions/v1/extension_deployment_service.proto",
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
