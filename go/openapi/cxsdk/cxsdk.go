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

package cxsdk

import (
	actions "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/actions_service"
	alerts "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/alert_definitions_service"
	alertscheduler "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/alert_scheduler_rule_service"
	apikeys "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/api_keys_service"
	connectors "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/connectors_service"
	customEnrichments "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/custom_enrichments_service"
	dashboardfolders "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/dashboard_folders_service"
	dashboards "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/dashboard_service"
	enrichments "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/enrichments_service"
	events2metrics "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/events2metrics_service"
	extensiondeployments "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/extension_deployment_service"
	extensions "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/extension_service"
	viewsfolders "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/folders_for_views_service"
	globalrouters "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/global_routers_service"
	integrations "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/integration_service"
	ipaccess "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/ip_access_service"
	archivemetrics "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/metrics_data_archive_service"
	webhooks "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/outgoing_webhooks_service"
	tcopolicies "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/policies_service"
	presets "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/presets_service"
	archiveretention "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/retentions_service"
	customroles "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/role_management_service"
	scopes "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/scopes_service"

	// slo (no plural) is the legacy service. slos (plural) is the new one.
	recordingrules "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/recording_rules_service"
	rulegroups "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/rule_groups_service"
	slos "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/slos_service"
	targets "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/target_service"
	groups "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/team_permissions_management_service"
	views "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/views_service"
)

// ClientSet provides access to all API clients.
type ClientSet struct {
	actions              *actions.ActionsServiceAPIService
	alerts               *alerts.AlertDefinitionsServiceAPIService
	alertScheduler       *alertscheduler.AlertSchedulerRuleServiceAPIService
	apiKeys              *apikeys.APIKeysServiceAPIService
	archiveLogs          *targets.TargetServiceAPIService
	archiveMetrics       *archivemetrics.MetricsDataArchiveServiceAPIService
	archiveRetentions    *archiveretention.RetentionsServiceAPIService
	connectors           *connectors.ConnectorsServiceAPIService
	ipAccess             *ipaccess.IPAccessServiceAPIService
	customRoles          *customroles.RoleManagementServiceAPIService
	dashboards           *dashboards.DashboardServiceAPIService
	dashboardFolders     *dashboardfolders.DashboardFoldersServiceAPIService
	enrichments          *enrichments.EnrichmentsServiceAPIService
	customEnrichments    *customEnrichments.CustomEnrichmentsServiceAPIService
	extensions           *extensions.ExtensionServiceAPIService
	events2metrics       *events2metrics.Events2MetricsServiceAPIService
	extensionDeployments *extensiondeployments.ExtensionDeploymentServiceAPIService
	groups               *groups.TeamPermissionsManagementServiceAPIService
	integrations         *integrations.IntegrationServiceAPIService
	globalRouters        *globalrouters.GlobalRoutersServiceAPIService
	presets              *presets.PresetsServiceAPIService
	scopes               *scopes.ScopesServiceAPIService
	slos                 *slos.SlosServiceAPIService
	recordingRules       *recordingrules.RecordingRulesServiceAPIService
	ruleGroups           *rulegroups.RuleGroupsServiceAPIService
	tcoPolicies          *tcopolicies.PoliciesServiceAPIService
	webhooks             *webhooks.OutgoingWebhooksServiceAPIService
	views                *views.ViewsServiceAPIService
	viewsFolders         *viewsfolders.FoldersForViewsServiceAPIService
}

// Actions returns the ActionsServiceAPIService client.
func (c *ClientSet) Actions() *actions.ActionsServiceAPIService {
	return c.actions
}

// Alerts returns the AlertDefinitionsServiceAPIService client.
func (c *ClientSet) Alerts() *alerts.AlertDefinitionsServiceAPIService {
	return c.alerts
}

// AlertScheduler returns the AlertSchedulerRuleServiceAPIService client.
func (c *ClientSet) AlertScheduler() *alertscheduler.AlertSchedulerRuleServiceAPIService {
	return c.alertScheduler
}

// ArchiveLogs returns the TargetServiceAPIService client.
func (c *ClientSet) ArchiveLogs() *targets.TargetServiceAPIService {
	return c.archiveLogs
}

// ArchiveMetrics returns the MetricsDataArchiveServiceAPIService client.
func (c *ClientSet) ArchiveMetrics() *archivemetrics.MetricsDataArchiveServiceAPIService {
	return c.archiveMetrics
}

// ArchiveRetentions returns the RetentionsServiceAPIService client.
func (c *ClientSet) ArchiveRetentions() *archiveretention.RetentionsServiceAPIService {
	return c.archiveRetentions
}

// APIKeys returns the APIKeysServiceAPIService client.
func (c *ClientSet) APIKeys() *apikeys.APIKeysServiceAPIService {
	return c.apiKeys
}

// Connectors returns the ConnectorsServiceAPIService client.
func (c *ClientSet) Connectors() *connectors.ConnectorsServiceAPIService {
	return c.connectors
}

// IPAccess returns the IPAccessServiceAPIService client.
func (c *ClientSet) IPAccess() *ipaccess.IPAccessServiceAPIService {
	return c.ipAccess
}

// Dashboards returns the DashboardServiceAPIService client.
func (c *ClientSet) Dashboards() *dashboards.DashboardServiceAPIService {
	return c.dashboards
}

// DashboardFolders returns the DashboardFoldersServiceAPIService client.
func (c *ClientSet) DashboardFolders() *dashboardfolders.DashboardFoldersServiceAPIService {
	return c.dashboardFolders
}

// Enrichments returns the EnrichmentsServiceAPIService client.
func (c *ClientSet) Enrichments() *enrichments.EnrichmentsServiceAPIService {
	return c.enrichments
}

// CustomEnrichments returns the CustomEnrichmentsServiceAPIService client.
func (c *ClientSet) CustomEnrichments() *customEnrichments.CustomEnrichmentsServiceAPIService {
	return c.customEnrichments
}

// Events2Metrics returns the Events2MetricsServiceAPIService client.
func (c *ClientSet) Events2Metrics() *events2metrics.Events2MetricsServiceAPIService {
	return c.events2metrics
}

// CustomRoles returns the RoleManagementServiceAPIService client.
func (c *ClientSet) CustomRoles() *customroles.RoleManagementServiceAPIService {
	return c.customRoles
}

// Extensions returns the ExtensionServiceAPIService client.
func (c *ClientSet) Extensions() *extensions.ExtensionServiceAPIService {
	return c.extensions
}

// ExtensionDeployments returns the ExtensionDeploymentServiceAPIService client.
func (c *ClientSet) ExtensionDeployments() *extensiondeployments.ExtensionDeploymentServiceAPIService {
	return c.extensionDeployments
}

// Groups returns the TeamPermissionsManagementServiceAPIService client.
func (c *ClientSet) Groups() *groups.TeamPermissionsManagementServiceAPIService {
	return c.groups
}

// Integrations returns the IntegrationServiceAPIService client.
func (c *ClientSet) Integrations() *integrations.IntegrationServiceAPIService {
	return c.integrations
}

// GlobalRouters returns the GlobalRoutersServiceAPIService client.
func (c *ClientSet) GlobalRouters() *globalrouters.GlobalRoutersServiceAPIService {
	return c.globalRouters
}

// Presets returns the PresetsServiceAPIService client.
func (c *ClientSet) Presets() *presets.PresetsServiceAPIService {
	return c.presets
}

// SLOs returns the SlosServiceAPIService client.
func (c *ClientSet) SLOs() *slos.SlosServiceAPIService {
	return c.slos
}

// RecordingRules returns the RecordingRulesServiceAPIService client.
func (c *ClientSet) RecordingRules() *recordingrules.RecordingRulesServiceAPIService {
	return c.recordingRules
}

// RuleGroups returns the RuleGroupsServiceAPIService client.
func (c *ClientSet) RuleGroups() *rulegroups.RuleGroupsServiceAPIService {
	return c.ruleGroups
}

// TCOPolicies returns the PoliciesServiceAPIService client.
func (c *ClientSet) TCOPolicies() *tcopolicies.PoliciesServiceAPIService {
	return c.tcoPolicies
}

// Scopes returns the ScopesServiceAPIService client.
func (c *ClientSet) Scopes() *scopes.ScopesServiceAPIService {
	return c.scopes
}

// Webhooks returns the OutgoingWebhooksServiceAPIService client.
func (c *ClientSet) Webhooks() *webhooks.OutgoingWebhooksServiceAPIService {
	return c.webhooks
}

// Views returns the ViewsServiceAPIService client.
func (c *ClientSet) Views() *views.ViewsServiceAPIService {
	return c.views
}

// ViewsFolders returns the FoldersForViewsServiceAPIService client.
func (c *ClientSet) ViewsFolders() *viewsfolders.FoldersForViewsServiceAPIService {
	return c.viewsFolders
}

// NewClientSet builds a ClientSet from CallPropertiesCreator.
func NewClientSet(c *Config) *ClientSet {
	return &ClientSet{
		actions:              NewActionsClient(c),
		alerts:               NewAlertsClient(c),
		alertScheduler:       NewAlertSchedulerClient(c),
		apiKeys:              NewAPIKeysClient(c),
		archiveMetrics:       NewArchiveMetricsClient(c),
		archiveLogs:          NewArchiveLogsClient(c),
		archiveRetentions:    NewArchiveRetentionsClient(c),
		ipAccess:             NewIPAccessClient(c),
		connectors:           NewConnectorsClient(c),
		customRoles:          NewCustomRolesClient(c),
		dashboards:           NewDashboardClient(c),
		dashboardFolders:     NewDashboardFoldersClient(c),
		enrichments:          NewEnrichmentsClient(c),
		customEnrichments:    NewCustomEnrichmentsClient(c),
		events2metrics:       NewEvents2MetricsClient(c),
		extensions:           NewExtensionsClient(c),
		extensionDeployments: NewExtensionDeploymentsClient(c),
		groups:               NewGroupsClient(c),
		integrations:         NewIntegrationsClient(c),
		globalRouters:        NewGlobalRoutersClient(c),
		presets:              NewPresetsClient(c),
		scopes:               NewScopesClient(c),
		slos:                 NewSLOsClient(c),
		recordingRules:       NewRecordingRulesClient(c),
		ruleGroups:           NewRuleGroupsClient(c),
		tcoPolicies:          NewTCOPoliciesClient(c),
		webhooks:             NewWebhooksClient(c),
		views:                NewViewsClient(c),
		viewsFolders:         NewViewsFoldersClient(c),
	}
}

// NewActionsClient builds a new ActionsServiceAPIService from CallPropertiesCreator.
func NewActionsClient(c *Config) *actions.ActionsServiceAPIService {
	cfg := actions.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = actions.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return actions.NewAPIClient(cfg).ActionsServiceAPI
}

// NewAlertsClient builds a new AlertDefinitionsServiceAPIService from CallPropertiesCreator.
func NewAlertsClient(c *Config) *alerts.AlertDefinitionsServiceAPIService {
	cfg := alerts.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}

	cfg.Servers = alerts.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return alerts.NewAPIClient(cfg).AlertDefinitionsServiceAPI
}

// NewAlertSchedulerClient builds a new AlertSchedulerRuleServiceAPIService from CallPropertiesCreator.
func NewAlertSchedulerClient(c *Config) *alertscheduler.AlertSchedulerRuleServiceAPIService {
	cfg := alertscheduler.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = alertscheduler.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return alertscheduler.NewAPIClient(cfg).AlertSchedulerRuleServiceAPI
}

// NewAPIKeysClient builds a new APIKeysServiceAPIService from CallPropertiesCreator.
func NewAPIKeysClient(c *Config) *apikeys.APIKeysServiceAPIService {
	cfg := apikeys.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = apikeys.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return apikeys.NewAPIClient(cfg).APIKeysServiceAPI
}

// NewArchiveMetricsClient builds a new MetricsDataArchiveServiceAPIService from CallPropertiesCreator.
func NewArchiveMetricsClient(c *Config) *archivemetrics.MetricsDataArchiveServiceAPIService {
	cfg := archivemetrics.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = archivemetrics.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return archivemetrics.NewAPIClient(cfg).MetricsDataArchiveServiceAPI
}

// NewIPAccessClient builds a new IPAccessServiceAPIService from CallPropertiesCreator.
func NewIPAccessClient(c *Config) *ipaccess.IPAccessServiceAPIService {
	cfg := ipaccess.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = ipaccess.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return ipaccess.NewAPIClient(cfg).IPAccessServiceAPI
}

// NewConnectorsClient builds a new ConnectorsServiceAPIService from CallPropertiesCreator.
func NewConnectorsClient(c *Config) *connectors.ConnectorsServiceAPIService {
	cfg := connectors.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = connectors.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return connectors.NewAPIClient(cfg).ConnectorsServiceAPI
}

// NewCustomRolesClient builds a new RoleManagementServiceAPIService from CallPropertiesCreator.
func NewCustomRolesClient(c *Config) *customroles.RoleManagementServiceAPIService {
	cfg := customroles.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = customroles.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return customroles.NewAPIClient(cfg).RoleManagementServiceAPI
}

// NewDashboardClient builds a new DashboardServiceAPIService from CallPropertiesCreator.
func NewDashboardClient(c *Config) *dashboards.DashboardServiceAPIService {
	cfg := dashboards.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = dashboards.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return dashboards.NewAPIClient(cfg).DashboardServiceAPI
}

// NewDashboardFoldersClient builds a new DashboardFoldersServiceAPIService from CallPropertiesCreator.
func NewDashboardFoldersClient(c *Config) *dashboardfolders.DashboardFoldersServiceAPIService {
	cfg := dashboardfolders.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = dashboardfolders.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return dashboardfolders.NewAPIClient(cfg).DashboardFoldersServiceAPI
}

// NewEnrichmentsClient builds a new EnrichmentsServiceAPIService from CallPropertiesCreator.
func NewEnrichmentsClient(c *Config) *enrichments.EnrichmentsServiceAPIService {
	cfg := enrichments.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = enrichments.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return enrichments.NewAPIClient(cfg).EnrichmentsServiceAPI
}

// NewCustomEnrichmentsClient builds a new CustomEnrichmentsServiceAPIService from CallPropertiesCreator.
func NewCustomEnrichmentsClient(c *Config) *customEnrichments.CustomEnrichmentsServiceAPIService {
	cfg := customEnrichments.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = customEnrichments.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return customEnrichments.NewAPIClient(cfg).CustomEnrichmentsServiceAPI
}

// NewEvents2MetricsClient builds a new Events2MetricsServiceAPIService from CallPropertiesCreator.
func NewEvents2MetricsClient(c *Config) *events2metrics.Events2MetricsServiceAPIService {
	cfg := events2metrics.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = events2metrics.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return events2metrics.NewAPIClient(cfg).Events2MetricsServiceAPI
}

// NewExtensionsClient builds a new ExtensionServiceAPIService from CallPropertiesCreator.
func NewExtensionsClient(c *Config) *extensions.ExtensionServiceAPIService {
	cfg := extensions.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = extensions.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return extensions.NewAPIClient(cfg).ExtensionServiceAPI
}

// NewExtensionDeploymentsClient builds a new ExtensionDeploymentServiceAPIService from CallPropertiesCreator.
func NewExtensionDeploymentsClient(c *Config) *extensiondeployments.ExtensionDeploymentServiceAPIService {
	cfg := extensiondeployments.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = extensiondeployments.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return extensiondeployments.NewAPIClient(cfg).ExtensionDeploymentServiceAPI
}

// NewGroupsClient builds a new TeamPermissionsManagementServiceAPIService from CallPropertiesCreator.
func NewGroupsClient(c *Config) *groups.TeamPermissionsManagementServiceAPIService {
	cfg := groups.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = groups.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return groups.NewAPIClient(cfg).TeamPermissionsManagementServiceAPI
}

// NewIntegrationsClient builds a new IntegrationServiceAPIService from CallPropertiesCreator.
func NewIntegrationsClient(c *Config) *integrations.IntegrationServiceAPIService {
	cfg := integrations.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = integrations.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return integrations.NewAPIClient(cfg).IntegrationServiceAPI
}

// NewGlobalRoutersClient builds a new GlobalRoutersServiceAPIService from CallPropertiesCreator.
func NewGlobalRoutersClient(c *Config) *globalrouters.GlobalRoutersServiceAPIService {
	cfg := globalrouters.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = globalrouters.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return globalrouters.NewAPIClient(cfg).GlobalRoutersServiceAPI
}

// NewPresetsClient builds a new PresetsServiceAPIService from CallPropertiesCreator.
func NewPresetsClient(c *Config) *presets.PresetsServiceAPIService {
	cfg := presets.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = presets.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return presets.NewAPIClient(cfg).PresetsServiceAPI
}

// NewScopesClient builds a new ScopesServiceAPIService from CallPropertiesCreator.
func NewScopesClient(c *Config) *scopes.ScopesServiceAPIService {
	cfg := scopes.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = scopes.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return scopes.NewAPIClient(cfg).ScopesServiceAPI
}

// NewSLOsClient builds a new SlosServiceAPIService from CallPropertiesCreator.
func NewSLOsClient(c *Config) *slos.SlosServiceAPIService {
	cfg := slos.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = slos.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return slos.NewAPIClient(cfg).SlosServiceAPI
}

// NewRecordingRulesClient builds a new RecordingRulesServiceAPIService from CallPropertiesCreator.
func NewRecordingRulesClient(c *Config) *recordingrules.RecordingRulesServiceAPIService {
	cfg := recordingrules.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = recordingrules.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return recordingrules.NewAPIClient(cfg).RecordingRulesServiceAPI
}

// NewRuleGroupsClient builds a new RuleGroupsServiceAPIService from CallPropertiesCreator.
func NewRuleGroupsClient(c *Config) *rulegroups.RuleGroupsServiceAPIService {
	cfg := rulegroups.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = rulegroups.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return rulegroups.NewAPIClient(cfg).RuleGroupsServiceAPI
}

// NewTCOPoliciesClient builds a new PoliciesServiceAPIService from CallPropertiesCreator.
func NewTCOPoliciesClient(c *Config) *tcopolicies.PoliciesServiceAPIService {
	cfg := tcopolicies.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = tcopolicies.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return tcopolicies.NewAPIClient(cfg).PoliciesServiceAPI
}

// NewWebhooksClient builds a new OutgoingWebhooksServiceAPIService from CallPropertiesCreator.
func NewWebhooksClient(c *Config) *webhooks.OutgoingWebhooksServiceAPIService {
	cfg := webhooks.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = webhooks.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return webhooks.NewAPIClient(cfg).OutgoingWebhooksServiceAPI
}

// NewViewsClient builds a new ViewsServiceAPIService from CallPropertiesCreator.
func NewViewsClient(c *Config) *views.ViewsServiceAPIService {
	cfg := views.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = views.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return views.NewAPIClient(cfg).ViewsServiceAPI
}

// NewViewsFoldersClient builds a new FoldersForViewsServiceAPIService from CallPropertiesCreator.
func NewViewsFoldersClient(c *Config) *viewsfolders.FoldersForViewsServiceAPIService {
	cfg := viewsfolders.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = viewsfolders.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return viewsfolders.NewAPIClient(cfg).FoldersForViewsServiceAPI
}

// NewArchiveLogsClient builds a new TargetServiceAPIService from CallPropertiesCreator.
func NewArchiveLogsClient(c *Config) *targets.TargetServiceAPIService {
	cfg := targets.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = targets.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return targets.NewAPIClient(cfg).TargetServiceAPI
}

// NewArchiveRetentionsClient builds a new RetentionsServiceAPIService from CallPropertiesCreator.
func NewArchiveRetentionsClient(c *Config) *archiveretention.RetentionsServiceAPIService {
	cfg := archiveretention.NewConfiguration()
	if c.httpClient != nil {
		cfg.HTTPClient = c.httpClient
	}
	cfg.Servers = archiveretention.ServerConfigurations{{URL: c.url}}
	for k, v := range c.headers {
		cfg.AddDefaultHeader(k, v)
	}
	return archiveretention.NewAPIClient(cfg).RetentionsServiceAPI
}
