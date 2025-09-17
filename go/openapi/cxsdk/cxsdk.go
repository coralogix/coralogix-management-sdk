package cxsdk

import (
	"fmt"
	"os"
	"strings"

	actions "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/actions_service"
	alerts "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/alert_definitions_service"
	apikeys "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/api_keys_service"
	dashboardfolders "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/dashboard_folders_service"
	dashboards "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/dashboard_service"
	enrichments "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/enrichments_service"
	viewsfolders "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/folders_for_views_service"
	integrations "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/integration_service"
	ipaccess "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/ip_access_service"
	saml "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/saml_configuration_service"
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
	actions          *actions.ActionsServiceAPIService
	alerts           *alerts.AlertDefinitionsServiceAPIService
	apiKeys          *apikeys.APIKeysServiceAPIService
	ipAccess         *ipaccess.IPAccessServiceAPIService
	dashboards       *dashboards.DashboardServiceAPIService
	dashboardFolders *dashboardfolders.DashboardFoldersServiceAPIService
	enrichments      *enrichments.EnrichmentsServiceAPIService
	integrations     *integrations.IntegrationServiceAPIService
	saml             *saml.SAMLConfigurationServiceAPIService
	//slos             *slos.SLOServiceAPIService
	views        *views.ViewsServiceAPIService
	viewsfolders *viewsfolders.FoldersForViewsServiceAPIService
}

// Actions returns the ActionsServiceAPIService client.
func (c *ClientSet) Actions() *actions.ActionsServiceAPIService {
	return c.actions
}

// Alerts returns the AlertDefinitionsServiceAPIService client.
func (c *ClientSet) Alerts() *alerts.AlertDefinitionsServiceAPIService {
	return c.alerts
}

// APIKeys returns the APIKeysServiceAPIService client.
func (c *ClientSet) APIKeys() *apikeys.APIKeysServiceAPIService {
	return c.apiKeys
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

// Integrations returns the IntegrationServiceAPIService client.
func (c *ClientSet) Integrations() *integrations.IntegrationServiceAPIService {
	return c.integrations
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
	return c.viewsfolders
}

// NewClientSet builds a ClientSet from CallPropertiesCreator.
func NewClientSet(c CallPropertiesCreator) *ClientSet {
	return &ClientSet{
		actions:          NewActionsClient(c),
		alerts:           NewAlertsClient(c),
		apiKeys:          NewAPIKeysClient(c),
		ipAccess:         NewIPAccessClient(c),
		dashboards:       NewDashboardClient(c),
		dashboardFolders: NewDashboardFoldersClient(c),
		enrichments:      NewEnrichmentsClient(c),
		integrations:     NewIntegrationsClient(c),
		saml:             NewSAMLClient(c),
		//slos:             NewSLOsClient(c),
		views:        NewViewsClient(c),
		viewsfolders: NewViewsFoldersClient(c),
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

// NewIntegrationsClient builds a new IntegrationServiceAPIService from CallPropertiesCreator.
func NewIntegrationsClient(c CallPropertiesCreator) *integrations.IntegrationServiceAPIService {
	cfg := integrations.NewConfiguration()
	cfg.Servers = integrations.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return integrations.NewAPIClient(cfg).IntegrationServiceAPI
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

// APIKeyFromEnv reads the Coralogix API key from the environment variable.
func APIKeyFromEnv() string {
	apiKey, ok := os.LookupEnv(envCoralogixAPIKey)
	if !ok {
		panic("Environment variable CORALOGIX_API_KEY is not set")
	}

	return apiKey
}
