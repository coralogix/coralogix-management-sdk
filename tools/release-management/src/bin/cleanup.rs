use std::path::PathBuf;

use clap::Parser;
use cx_sdk::{
    auth::AuthContext,
    client::{
        actions::ActionsClient,
        alerts::AlertsClient,
        alerts_scheduler::AlertsSchedulerClient,
        apikeys::ApiKeysClient,
        archive_logs::ArchiveLogsClient,
        archive_metrics::ArchiveMetricsClient,
        archive_retentions::ArchiveRetentionsClient,
        dashboards::DashboardsClient,
        datasets::DatasetsClient,
        enrichments::EnrichmentsClient,
        events2metrics::Events2MetricsClient,
        slos::SlosClient,
        tco_policies::TcoPoliciesClient,
        teams::TeamsClient,
        webhooks::WebhooksClient,
        scopes::ScopesClient,
        groups::GroupsClient,
        users::UsersClient,
        integrations::IntegrationsClient,
        dashboard_folders::DashboardFoldersClient,
        recording_rule_group_sets::RecordingRuleGroupSetsClient,
        saml::SamlClient,
        notifications::NotificationsClient,
        data_usage::DataUsageClient,
        contextual_data_integrations::ContextualDataIntegrationsClient,
        incidents::IncidentsClient,
        events::EventsClient,
        views::ViewsClient,
        views_folders::ViewsFoldersClient,
        extensions::ExtensionsClient,
        ip_access::IpAccessClient,
    },
    CoralogixRegion,
};
use git_cmd::Repo;

#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Args {
    #[arg(short, long, default_value = "protofetch.toml")]
    protofetch_descriptor: String,

    #[arg(short = 'l', long, default_value = "protofetch.lock")]
    protofetch_lockfile: String,

    #[arg(short, long)]
    root: String,

    #[arg(short, long, default_value = ".")]
    git_dir: String,

    #[arg(short = 'd', long, default_value = "proto")]
    protos_dir: String,

    #[arg(short = 'o', long)]
    github_repo: String,

    #[arg(short = 's', long, default_value = "true")]
    skip_hashes: bool,

    #[arg(short = 'f', long)]
    prefix_filter: Option<String>,
}

#[tokio::main]
async fn main() -> eyre::Result<()> {
    // Setup
    let subscriber = tracing_subscriber::FmtSubscriber::new();
    tracing::subscriber::set_global_default(subscriber)?;

    let args = Args::parse();

    // Instantiate all clients from `cx_sdk::client`
    let actions_client = ActionsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let alerts_client = AlertsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let alerts_scheduler_client = AlertsSchedulerClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let apikeys_client = ApiKeysClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let archive_logs_client = ArchiveLogsClient::new(
        CoralogixRegion::from_env().unwrap(),
        AuthContext::from_env(),
    )
    .unwrap();

    let archive_metrics_client = ArchiveMetricsClient::new(
        CoralogixRegion::from_env().unwrap(),
        AuthContext::from_env(),
    )
    .unwrap();

    let archive_retentions_client = ArchiveRetentionsClient::new(
        CoralogixRegion::from_env().unwrap(),
        AuthContext::from_env(),
    )
    .unwrap();

    let dashboards_client = DashboardsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let datasets_client = DatasetsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let enrichments_client = EnrichmentsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let events2metrics_client = Events2MetricsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let slos_client = SlosClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let tco_policies_client = TcoPoliciesClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let teams_client = TeamsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let webhooks_client = WebhooksClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let scopes_client = ScopesClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let groups_client = GroupsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let users_client = UsersClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let integrations_client = IntegrationsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let dashboard_folders_client = DashboardFoldersClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let recording_rule_group_sets_client = RecordingRuleGroupSetsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let saml_client = SamlClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let notifications_client = NotificationsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let data_usage_client = DataUsageClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let contextual_data_integrations_client = ContextualDataIntegrationsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let incidents_client = IncidentsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let events_client = EventsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let views_client = ViewsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let views_folders_client = ViewsFoldersClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let extensions_client = ExtensionsClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    let ip_access_client = IpAccessClient::new(
        AuthContext::from_env(),
        CoralogixRegion::from_env().unwrap(),
    )
    .unwrap();

    // Clients are now instantiated and can be used for further operations.
    
    Ok(())
}
