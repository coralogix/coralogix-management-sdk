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

use cx_sdk::{
    CoralogixRegion,
    auth::AuthContext,
    client::{
        actions::ActionsClient,
        alerts::AlertsClient,
        contextual_data_integrations::ContextualDataIntegrationsClient,
        dashboard_folders::DashboardFoldersClient,
        dashboards::DashboardsClient,
        datasets::DatasetClient,
        enrichments::EnrichmentsClient,
        events2metrics::Events2MetricsClient,
        extensions::ExtensionsClient,
        integrations::IntegrationsClient,
        ip_access::IpAccessClient,
        notifications::{ConnectorType, NotificationsClient},
        recording_rule_group_sets::RecordingRuleGroupSetsClient,
        scopes::ScopesClient,
        slos::SloClient,
        tco_policies::TcoPoliciesClient,
        users::UsersClient,
        views::ViewsClient,
        views_folders::ViewFoldersClient,
        webhooks::WebhooksClient,
    },
};
use tracing::{Level, info};

#[tokio::main]
async fn main() -> eyre::Result<()> {
    // Setup
    let subscriber = tracing_subscriber::FmtSubscriber::new();
    tracing::subscriber::set_global_default(subscriber)?;
    // Actions
    delete_all(
        "Actions".into(),
        Box::new(|| async {
            let actions_client = ActionsClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();

            for id in actions_client
                .list()
                .await
                .unwrap_or_default()
                .into_iter()
                .map(|v| v.id.unwrap())
            {
                info!("{:?}", actions_client.delete(id).await);
            }
        }),
    )
    .await;

    // Alerts
    delete_all(
        "Alerts".into(),
        Box::new(|| async {
            let alerts_client = AlertsClient::new(
                CoralogixRegion::from_env().unwrap(),
                AuthContext::from_env(),
                None,
            )
            .unwrap();

            for id in alerts_client
                .list()
                .await
                .unwrap_or_default()
                .alert_defs
                .into_iter()
                .map(|v| v.id.unwrap())
            {
                info!("{:?}", alerts_client.delete(id).await);
            }
        }),
    )
    .await;

    // Custom Enrichments
    delete_all(
        "Custom Enrichments".into(),
        Box::new(|| async {
            let datasets_client = DatasetClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();

            for id in datasets_client
                .list()
                .await
                .unwrap_or_default()
                .custom_enrichments
                .into_iter()
                .map(|v| v.id)
            {
                info!("{:?}", datasets_client.delete(id).await);
            }
        }),
    )
    .await;

    // All Enrichment types

    delete_all(
        "Other Enrichments".into(),
        Box::new(|| async {
            let enrichments_client = EnrichmentsClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();

            let all_enrichments = enrichments_client
                .list()
                .await
                .unwrap_or_default()
                .enrichments
                .into_iter()
                .map(|v| v.id)
                .collect();

            info!("{:?}", enrichments_client.delete(all_enrichments).await);
        }),
    )
    .await;

    // E2M
    delete_all(
        "Events2Metrics".into(),
        Box::new(|| async {
            let events2metrics_client = Events2MetricsClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();

            for id in events2metrics_client
                .list()
                .await
                .unwrap_or_default()
                .e2m
                .into_iter()
                .map(|v| v.id.unwrap())
            {
                info!("{:?}", events2metrics_client.delete(id).await);
            }
        }),
    )
    .await;

    // New SLO
    delete_all(
        "SLOs".into(),
        Box::new(|| async {
            let slos_client = SloClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();

            for id in slos_client
                .list(None)
                .await
                .unwrap_or_default()
                .slos
                .into_iter()
                .map(|v| v.id.unwrap())
            {
                info!("{:?}", slos_client.delete(id).await);
            }
        }),
    )
    .await;

    // TCOs
    delete_all(
        "TCOs".into(),
        Box::new(|| async {
            let tco_policies_client = TcoPoliciesClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();

            for id in tco_policies_client
                .list(cx_sdk::client::tco_policies::SourceType::Unspecified, false)
                .await
                .unwrap_or_default()
                .policies
                .into_iter()
                .map(|v| v.id.unwrap())
            {
                info!("{:?}", tco_policies_client.delete(id).await);
            }
        }),
    )
    .await;

    // Webhooks
    delete_all(
        "Webhooks".into(),
        Box::new(|| async {
            let webhooks_client = WebhooksClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();

            for id in webhooks_client
                .list()
                .await
                .unwrap_or_default()
                .deployed
                .into_iter()
                .map(|v| v.id.unwrap())
            {
                info!("{:?}", webhooks_client.delete(id).await);
            }
        }),
    )
    .await;

    // Notification center
    delete_all(
        "Notification Center".into(),
        Box::new(|| async {
            let notifications_client = NotificationsClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();

            for id in notifications_client
                .list_connectors(ConnectorType::GenericHttps, None)
                .await
                .unwrap_or_default()
                .connectors
                .into_iter()
                .map(|v| v.id.unwrap())
            {
                info!("{:?}", notifications_client.delete_connector(id).await);
            }

            for id in notifications_client
                .list_preset_summaries(
                    Some(ConnectorType::GenericHttps),
                    cx_sdk::client::notifications::EntityType::Alerts,
                )
                .await
                .unwrap_or_default()
                .preset_summaries
                .into_iter()
                .map(|v| v.id)
            {
                info!("{:?}", notifications_client.delete_custom_preset(id).await);
            }
        }),
    )
    .await;

    // Scopes
    delete_all(
        "Scopes".into(),
        Box::new(|| async {
            let scopes_client = ScopesClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();

            for id in scopes_client
                .list()
                .await
                .unwrap_or_default()
                .scopes
                .into_iter()
                .map(|v| v.id)
            {
                info!("{:?}", scopes_client.delete(id).await);
            }
        }),
    )
    .await;
    // Users
    delete_all(
        "Users".into(),
        Box::new(|| async {
            let users_client = UsersClient::new(
                CoralogixRegion::from_env().unwrap(),
                AuthContext::from_env(),
            );

            for id in users_client
                .list()
                .await
                .unwrap_or_default()
                .into_iter()
                .filter(|f| f.active)
                .map(|v| v.id.unwrap())
            {
                info!("{:?}", users_client.delete(&id).await);
            }
        }),
    )
    .await;

    // Dashboards
    delete_all(
        "Dashboards".into(),
        Box::new(|| async {
            let dashboards_client = DashboardsClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();

            for id in dashboards_client
                .list()
                .await
                .unwrap_or_default()
                .items
                .into_iter()
                .map(|v| v.id.unwrap())
            {
                info!("{:?}", dashboards_client.delete(id).await);
            }
        }),
    )
    .await;

    delete_all(
        "Dashboard Folders".into(),
        Box::new(|| async {
            let dashboard_folders_client = DashboardFoldersClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();

            for id in dashboard_folders_client
                .list()
                .await
                .unwrap_or_default()
                .folder
                .into_iter()
                .map(|v| v.id.unwrap())
            {
                info!("{:?}", dashboard_folders_client.delete(id).await);
            }
        }),
    )
    .await;

    // Recording Rules
    delete_all(
        "Recording Rules".into(),
        Box::new(|| async {
            let recording_rule_group_sets_client = RecordingRuleGroupSetsClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();

            for id in recording_rule_group_sets_client
                .list()
                .await
                .unwrap_or_default()
                .sets
                .into_iter()
                .map(|v| v.id)
            {
                info!("{:?}", recording_rule_group_sets_client.delete(id).await);
            }
        }),
    )
    .await;

    // Contextual data
    delete_all(
        "Contextual data".into(),
        Box::new(|| async {
            // Contextual data
            let contextual_data_integrations_client = ContextualDataIntegrationsClient::new(
                CoralogixRegion::from_env().unwrap(),
                AuthContext::from_env(),
            )
            .unwrap();

            for id in contextual_data_integrations_client
                .list(true)
                .await
                .unwrap_or_default()
                .integrations
                .into_iter()
                .map(|v| v.integration.unwrap().id.unwrap())
            {
                if let Ok(all_instances) = contextual_data_integrations_client
                    .get_details(id, true)
                    .await
                {
                    match all_instances
                        .integration_detail
                        .unwrap()
                        .integration_type_details
                    {
                        Some(cx_sdk::client::integrations::IntegrationTypeDetails::Default(
                            details,
                        )) => {
                            for d in details.registered {
                                info!(
                                    "{:?}",
                                    contextual_data_integrations_client
                                        .delete(d.id.unwrap())
                                        .await
                                );
                            }
                        }
                        _ => {}
                    }
                }
            }
        }),
    )
    .await;

    // integrations
    delete_all(
        "Integrations".into(),
        Box::new(|| async {
            let integrations_client = IntegrationsClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();
            for id in integrations_client
                .list(true)
                .await
                .unwrap_or_default()
                .integrations
                .into_iter()
                .map(|e| e.integration.unwrap().id.unwrap())
            {
                if let Ok(all_instances) = integrations_client.get_details(id, true).await {
                    match all_instances
                        .integration_detail
                        .unwrap()
                        .integration_type_details
                        .unwrap()
                    {
                        cx_sdk::client::integrations::IntegrationTypeDetails::Default(details) => {
                            for d in details.registered {
                                info!("{:?}", integrations_client.delete(d.id.unwrap()).await);
                            }
                        }
                    }
                }
            }
        }),
    )
    .await;

    // extensions
    delete_all(
        "Extensions".into(),
        Box::new(|| async {
            let extensions_client = ExtensionsClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();
            for id in extensions_client
                .get_deployed()
                .await
                .unwrap_or_default()
                .deployed_extensions
                .into_iter()
                .map(|e| e.id)
            {
                info!(
                    "{:?}",
                    extensions_client.undeploy(id.unwrap(), vec![]).await
                );
            }
        }),
    )
    .await;

    // views
    delete_all(
        "Views".into(),
        Box::new(|| async {
            let views_client = ViewsClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();

            for id in views_client
                .list()
                .await
                .unwrap_or_default()
                .into_iter()
                .map(|v| v.id.unwrap())
            {
                info!("{:?}", views_client.delete(id).await);
            }
        }),
    )
    .await;
    // view folders
    delete_all(
        "View Folders".into(),
        Box::new(|| async {
            let view_folders_client = ViewFoldersClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();

            for id in view_folders_client
                .list()
                .await
                .unwrap_or_default()
                .into_iter()
                .map(|v| v.id.unwrap())
            {
                info!("{:?}", view_folders_client.delete(id).await);
            }
        }),
    )
    .await;

    // IP Access
    delete_all(
        "IPAccess".into(),
        Box::new(|| async {
            let ip_access_client = IpAccessClient::new(
                AuthContext::from_env(),
                CoralogixRegion::from_env().unwrap(),
            )
            .unwrap();

            info!("{:?}", ip_access_client.delete(None).await);
        }),
    )
    .await;
    Ok(())
}

#[tracing::instrument(level=Level::INFO, skip_all, fields(entity=%name))]
async fn delete_all<F: Future<Output = ()> + Sized>(name: String, f: Box<fn() -> F>) {
    f().await;
    info!("done");
}
