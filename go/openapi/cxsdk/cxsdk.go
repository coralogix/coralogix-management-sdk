package cxsdk

import (
	"fmt"
	views "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/views_service"
	"os"
	"strings"

	actions "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/actions_service"
	apikeys "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/api_keys_service"
	dashboardfolders "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/dashboard_folders_service"
	dashboards "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/dashboard_service"
	enrichments "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/enrichments_service"
	integrations "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/integration_service"
	ipaccess "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/ip_access_service"
	slos "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/slos_service"
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
	apiKeys          *apikeys.APIKeysServiceAPIService
	ipAccess         *ipaccess.IPAccessServiceAPIService
	dashboards       *dashboards.DashboardServiceAPIService
	dashboardFolders *dashboardfolders.DashboardFoldersServiceAPIService
	enrichments      *enrichments.EnrichmentsServiceAPIService
	integrations     *integrations.IntegrationServiceAPIService
	slos             *slos.SlosServiceAPIService
	views            *views.ViewsServiceAPIService
}

// Actions returns the ActionsServiceAPIService client.
func (c *ClientSet) Actions() *actions.ActionsServiceAPIService {
	return c.actions
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
func (c *ClientSet) SLOs() *slos.SlosServiceAPIService {
	return c.slos
}

// Views returns the ViewsServiceAPIService client.
func (c *ClientSet) Views() *views.ViewsServiceAPIService {
	return c.views
}

// NewClientSet builds a ClientSet from CallPropertiesCreator.
func NewClientSet(c CallPropertiesCreator) *ClientSet {
	return &ClientSet{
		actions:          NewActionsClient(c),
		apiKeys:          NewAPIKeysClient(c),
		ipAccess:         NewIPAccessClient(c),
		dashboards:       NewDashboardClient(c),
		dashboardFolders: NewDashboardFoldersClient(c),
		enrichments:      NewEnrichmentsClient(c),
		integrations:     NewIntegrationsClient(c),
		slos:             NewSLOsClient(c),
		views:            NewViewsClient(c),
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

// NewSLOsClient builds a new SlosServiceAPIService from CallPropertiesCreator.
func NewSLOsClient(c CallPropertiesCreator) *slos.SlosServiceAPIService {
	cfg := slos.NewConfiguration()
	cfg.Servers = slos.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return slos.NewAPIClient(cfg).SlosServiceAPI
}

// NewViewsClient builds a new ViewsServiceAPIService from CallPropertiesCreator.
func NewViewsClient(c CallPropertiesCreator) *views.ViewsServiceAPIService {
	cfg := views.NewConfiguration()
	cfg.Servers = views.ServerConfigurations{{URL: c.URL()}}
	for k, v := range c.Headers() {
		cfg.AddDefaultHeader(k, v)
	}
	return views.NewAPIClient(cfg).ViewsServiceAPI
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
		return "https://api.coralogix.com/mgmt/openapi"
	case "eu2":
		return "https://api.eu2.coralogix.com/mgmt/openapi"
	case "us1":
		return "https://api.coralogix.us/mgmt/openapi"
	case "us2":
		return "https://api.cx498.coralogix.com/mgmt/openapi"
	case "ap1":
		return "https://api.coralogix.in/mgmt/openapi"
	case "ap2":
		return "https://api.coralogixsg.com/mgmt/openapi"
	case "ap3":
		return "https://api.ap3.coralogix.com/mgmt/openapi"
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
