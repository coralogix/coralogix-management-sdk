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
	"fmt"
	"os"
	"strings"

	actions "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/actions_service"
	alerts "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/alert_definitions_service"
	alertscheduler "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/alert_scheduler_rule_service"
	apikeys "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/api_keys_service"
	connectors "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/connectors_service"
	dashboardfolders "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/dashboard_folders_service"
	dashboards "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/dashboard_service"
	enrichments "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/enrichments_service"
	extensiondeployments "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/extension_deployment_service"
	extensions "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/extension_service"
	viewsfolders "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/folders_for_views_service"
	globalrouters "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/global_routers_service"
	integrations "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/integration_service"
	ipaccess "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/ip_access_service"
	presets "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/presets_service"
	saml "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/saml_configuration_service"
	groups "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/team_permissions_management_service"
	views "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/views_service"
	//slos "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/slo_service"
)

const (
	envCoralogixRegion = "CORALOGIX_REGION"
	envCoralogixAPIKey = "CORALOGIX_API_KEY"
)

const (
	sdkVersionHeaderName       = "x-cx-sdk-version"
	sdkLanguageHeaderName      = "x-cx-sdk-language"
	sdkGoVersionHeaderName     = "x-cx-go-version"
	sdkCorrelationIDHeaderName = "x-cx-correlation-id"
	vanillaSdkVersion          = "1.8.1"
)

// ClientSet provides access to all API clients.
type ClientSet struct {
	actions              *actions.ActionsServiceAPIService
	alerts               *alerts.AlertDefinitionsServiceAPIService
	alertScheduler       *alertscheduler.AlertSchedulerRuleServiceAPIService
	apiKeys              *apikeys.APIKeysServiceAPIService
	connectors           *connectors.ConnectorsServiceAPIService
	ipAccess             *ipaccess.IPAccessServiceAPIService
	dashboards           *dashboards.DashboardServiceAPIService
	dashboardFolders     *dashboardfolders.DashboardFoldersServiceAPIService
	enrichments          *enrichments.EnrichmentsServiceAPIService
	extensions           *extensions.ExtensionServiceAPIService
	extensionDeployments *extensiondeployments.ExtensionDeploymentServiceAPIService
	groups               *groups.TeamPermissionsManagementServiceAPIService
	integrations         *integrations.IntegrationServiceAPIService
	globalRouters        *globalrouters.GlobalRoutersServiceAPIService
	presets              *presets.PresetsServiceAPIService
	saml                 *saml.SAMLConfigurationServiceAPIService
	//slos             *slos.SLOServiceAPIService
	views        *views.ViewsServiceAPIService
	viewsFolders *viewsfolders.FoldersForViewsServiceAPIService
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
//func (c *ClientSet) SLOs() *slos.SLOServiceAPIService {
//	return c.slos
//}

// SAML returns the SAMLConfigurationServiceAPIService client.
func (c *ClientSet) SAML() *saml.SAMLConfigurationServiceAPIService {
	return c.saml
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
func NewClientSet(c CallPropertiesCreator) *ClientSet {
	return &ClientSet{
		actions:              NewActionsClient(c),
		alerts:               NewAlertsClient(c),
		alertScheduler:       NewAlertSchedulerClient(c),
		apiKeys:              NewAPIKeysClient(c),
		ipAccess:             NewIPAccessClient(c),
		connectors:           NewConnectorsClient(c),
		dashboards:           NewDashboardClient(c),
		dashboardFolders:     NewDashboardFoldersClient(c),
		enrichments:          NewEnrichmentsClient(c),
		extensions:           NewExtensionsClient(c),
		extensionDeployments: NewExtensionDeploymentsClient(c),
		groups:               NewGroupsClient(c),
		integrations:         NewIntegrationsClient(c),
		globalRouters:        NewGlobalRoutersClient(c),
		presets:              NewPresetsClient(c),
		saml:                 NewSAMLClient(c),
		//slos:             NewSLOsClient(c),
		views:        NewViewsClient(c),
		viewsFolders: NewViewsFoldersClient(c),
	}
}

// NewActionsClient builds a new ActionsServiceAPIService from CallPropertiesCreator.
func NewActionsClient(c CallPropertiesCreator) *actions.ActionsServiceAPIService {
	cfg := actions.NewConfiguration()
	cfg.Servers = actions.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return actions.NewAPIClient(cfg).ActionsServiceAPI
}

// NewAlertsClient builds a new AlertDefinitionsServiceAPIService from CallPropertiesCreator.
func NewAlertsClient(c CallPropertiesCreator) *alerts.AlertDefinitionsServiceAPIService {
	cfg := alerts.NewConfiguration()
	cfg.Servers = alerts.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return alerts.NewAPIClient(cfg).AlertDefinitionsServiceAPI
}

// NewAlertSchedulerClient builds a new AlertSchedulerRuleServiceAPIService from CallPropertiesCreator.
func NewAlertSchedulerClient(c CallPropertiesCreator) *alertscheduler.AlertSchedulerRuleServiceAPIService {
	cfg := alertscheduler.NewConfiguration()
	cfg.Servers = alertscheduler.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return alertscheduler.NewAPIClient(cfg).AlertSchedulerRuleServiceAPI
}

// NewAPIKeysClient builds a new APIKeysServiceAPIService from CallPropertiesCreator.
func NewAPIKeysClient(c CallPropertiesCreator) *apikeys.APIKeysServiceAPIService {
	cfg := apikeys.NewConfiguration()
	cfg.Servers = apikeys.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return apikeys.NewAPIClient(cfg).APIKeysServiceAPI
}

// NewIPAccessClient builds a new IPAccessServiceAPIService from CallPropertiesCreator.
func NewIPAccessClient(c CallPropertiesCreator) *ipaccess.IPAccessServiceAPIService {
	cfg := ipaccess.NewConfiguration()
	cfg.Servers = ipaccess.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return ipaccess.NewAPIClient(cfg).IPAccessServiceAPI
}

// NewConnectorsClient builds a new ConnectorsServiceAPIService from CallPropertiesCreator.
func NewConnectorsClient(c CallPropertiesCreator) *connectors.ConnectorsServiceAPIService {
	cfg := connectors.NewConfiguration()
	cfg.Servers = connectors.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return connectors.NewAPIClient(cfg).ConnectorsServiceAPI
}

// NewDashboardClient builds a new DashboardServiceAPIService from CallPropertiesCreator.
func NewDashboardClient(c CallPropertiesCreator) *dashboards.DashboardServiceAPIService {
	cfg := dashboards.NewConfiguration()
	cfg.Servers = dashboards.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return dashboards.NewAPIClient(cfg).DashboardServiceAPI
}

// NewDashboardFoldersClient builds a new DashboardFoldersServiceAPIService from CallPropertiesCreator.
func NewDashboardFoldersClient(c CallPropertiesCreator) *dashboardfolders.DashboardFoldersServiceAPIService {
	cfg := dashboardfolders.NewConfiguration()
	cfg.Servers = dashboardfolders.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return dashboardfolders.NewAPIClient(cfg).DashboardFoldersServiceAPI
}

// NewEnrichmentsClient builds a new EnrichmentsServiceAPIService from CallPropertiesCreator.
func NewEnrichmentsClient(c CallPropertiesCreator) *enrichments.EnrichmentsServiceAPIService {
	cfg := enrichments.NewConfiguration()
	cfg.Servers = enrichments.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return enrichments.NewAPIClient(cfg).EnrichmentsServiceAPI
}

// NewExtensionsClient builds a new ExtensionServiceAPIService from CallPropertiesCreator.
func NewExtensionsClient(c CallPropertiesCreator) *extensions.ExtensionServiceAPIService {
	cfg := extensions.NewConfiguration()
	cfg.Servers = extensions.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return extensions.NewAPIClient(cfg).ExtensionServiceAPI
}

// NewExtensionDeploymentsClient builds a new ExtensionDeploymentServiceAPIService from CallPropertiesCreator.
func NewExtensionDeploymentsClient(c CallPropertiesCreator) *extensiondeployments.ExtensionDeploymentServiceAPIService {
	cfg := extensiondeployments.NewConfiguration()
	cfg.Servers = extensiondeployments.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return extensiondeployments.NewAPIClient(cfg).ExtensionDeploymentServiceAPI
}

// NewGroupsClient builds a new TeamPermissionsManagementServiceAPIService from CallPropertiesCreator.
func NewGroupsClient(c CallPropertiesCreator) *groups.TeamPermissionsManagementServiceAPIService {
	cfg := groups.NewConfiguration()
	cfg.Servers = groups.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return groups.NewAPIClient(cfg).TeamPermissionsManagementServiceAPI
}

// NewIntegrationsClient builds a new IntegrationServiceAPIService from CallPropertiesCreator.
func NewIntegrationsClient(c CallPropertiesCreator) *integrations.IntegrationServiceAPIService {
	cfg := integrations.NewConfiguration()
	cfg.Servers = integrations.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return integrations.NewAPIClient(cfg).IntegrationServiceAPI
}

// NewGlobalRoutersClient builds a new GlobalRoutersServiceAPIService from CallPropertiesCreator.
func NewGlobalRoutersClient(c CallPropertiesCreator) *globalrouters.GlobalRoutersServiceAPIService {
	cfg := globalrouters.NewConfiguration()
	cfg.Servers = globalrouters.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return globalrouters.NewAPIClient(cfg).GlobalRoutersServiceAPI
}

// NewPresetsClient builds a new PresetsServiceAPIService from CallPropertiesCreator.
func NewPresetsClient(c CallPropertiesCreator) *presets.PresetsServiceAPIService {
	cfg := presets.NewConfiguration()
	cfg.Servers = presets.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return presets.NewAPIClient(cfg).PresetsServiceAPI
}

// NewSAMLClient builds a new SAMLConfigurationServiceAPIService from CallPropertiesCreator.
func NewSAMLClient(c CallPropertiesCreator) *saml.SAMLConfigurationServiceAPIService {
	cfg := saml.NewConfiguration()
	cfg.Servers = saml.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return saml.NewAPIClient(cfg).SAMLConfigurationServiceAPI
}

// NewSLOsClient builds a new SlosServiceAPIService from CallPropertiesCreator.
//func NewSLOsClient(c CallPropertiesCreator) *slos.SLOServiceAPIService {
//	cfg := slos.NewConfiguration()
//	cfg.Servers = slos.ServerConfigurations{{URL: c.URL()}}
//	for k, v := range c.Headers() {
//		cfg.AddDefaultHeader(k, v)
//	}
//	return slos.NewAPIClient(cfg).SLOServiceAPI
//}

// NewViewsClient builds a new ViewsServiceAPIService from CallPropertiesCreator.
func NewViewsClient(c CallPropertiesCreator) *views.ViewsServiceAPIService {
	cfg := views.NewConfiguration()
	cfg.Servers = views.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return views.NewAPIClient(cfg).ViewsServiceAPI
}

// NewViewsFoldersClient builds a new FoldersForViewsServiceAPIService from CallPropertiesCreator.
func NewViewsFoldersClient(c CallPropertiesCreator) *viewsfolders.FoldersForViewsServiceAPIService {
	cfg := viewsfolders.NewConfiguration()
	cfg.Servers = viewsfolders.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return viewsfolders.NewAPIClient(cfg).FoldersForViewsServiceAPI
}

// RegionFromEnv reads the Coralogix region from the environment variable.
func RegionFromEnv() string {
	region, ok := os.LookupEnv(envCoralogixRegion)
	if !ok {
		panic("Environment variable CORALOGIX_REGION is not set")
	}

	return region
}

// URLFromRegion returns the Coralogix OpenAPI URL for the given region.
func URLFromRegion(region string) string {
	switch strings.ToLower(region) {
	case "eu1":
		return "https://api.coralogix.com/mgmt/openapi/latest"
	case "eu2":
		return "https://api.eu2.coralogix.com/mgmt/openapi/latest"
	case "us1":
		return "https://api.coralogix.us/mgmt/openapi/latest"
	case "us2":
		return "https://api.cx498.coralogix.com/mgmt/openapi/latest"
	case "ap1":
		return "https://api.coralogix.in/mgmt/openapi/latest"
	case "ap2":
		return "https://api.coralogixsg.com/mgmt/openapi/latest"
	case "ap3":
		return "https://api.ap3.coralogix.com/mgmt/openapi/latest"
	default:
		panic(fmt.Sprintf("Unknown Coralogix region: %s", region))
	}
}

// URLFromDomain returns the Coralogix OpenAPI URL for the given custom domain.
func URLFromDomain(domain string) string {
	return fmt.Sprintf("https://api.%s/mgmt/openapi/latest", domain)
}

// APIKeyFromEnv reads the Coralogix API key from the environment variable.
func APIKeyFromEnv() string {
	apiKey, ok := os.LookupEnv(envCoralogixAPIKey)
	if !ok {
		panic("Environment variable CORALOGIX_API_KEY is not set")
	}

	return apiKey
}
