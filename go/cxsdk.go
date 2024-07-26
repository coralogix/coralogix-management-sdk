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
)

// ClientSet is a set of clients for the Coralogix SDK.
type ClientSet struct {
	ruleGroups        *RuleGroupsClient
	alerts            *AlertsClient
	enrichments       *EnrichmentsClient
	dataSet           *DataSetClient
	dashboards        *DashboardsClient
	actions           *ActionsClient
	tcoPolicies       *TCOPoliciesClient
	webhooks          *WebhooksClient
	events2Metrics    *Events2MetricsClient
	archiveRetentions *ArchiveRetentionsClient
	archiveMetrics    *ArchiveMetricsClient
	archiveLogs       *ArchiveLogsClient
	alertsSchedulers  *AlertsSchedulersClient
	teams             *TeamsClient
	slos              *SLOsClient
	scopes            *ScopesClient
	apiKeys           *ApikeysClient
	slis              *SLIClient
}

// RuleGroups gets a RuleGroupsClient from the ClientSet.
func (c *ClientSet) RuleGroups() *RuleGroupsClient {
	return c.ruleGroups
}

// Alerts gets an AlertsClient from the ClientSet.
func (c *ClientSet) Alerts() *AlertsClient {
	return c.alerts
}

// Enrichments gets an EnrichmentsClient from the ClientSet.
func (c *ClientSet) Enrichments() *EnrichmentsClient {
	return c.enrichments
}

// DataSet gets a DataSetClient from the ClientSet.
func (c *ClientSet) DataSet() *DataSetClient {
	return c.dataSet
}

// func (c *ClientSet) Dashboards() *DashboardsClient {
// 	return c.dashboards
// }

// Actions gets an ActionsClient from the ClientSet.
func (c *ClientSet) Actions() *ActionsClient {
	return c.actions
}

// TCOPolicies gets a TCOPoliciesClient from the ClientSet.
func (c *ClientSet) TCOPolicies() *TCOPoliciesClient {
	return c.tcoPolicies
}

// Webhooks gets a WebhooksClient from the ClientSet.
func (c *ClientSet) Webhooks() *WebhooksClient {
	return c.webhooks
}

// Events2Metrics gets an Events2MetricsClient from the ClientSet.
func (c *ClientSet) Events2Metrics() *Events2MetricsClient {
	return c.events2Metrics
}

// SLIs gets an SLIClient from the ClientSet.
func (c *ClientSet) SLIs() *SLIClient {
	return c.slis
}

// ArchiveRetentions gets an ArchiveRetentionsClient from the ClientSet.
func (c *ClientSet) ArchiveRetentions() *ArchiveRetentionsClient {
	return c.archiveRetentions
}

// ArchiveMetrics gets an ArchiveMetricsClient from the ClientSet.
func (c *ClientSet) ArchiveMetrics() *ArchiveMetricsClient {
	return c.archiveMetrics
}

// ArchiveLogs gets an ArchiveLogsClient from the ClientSet.
func (c *ClientSet) ArchiveLogs() *ArchiveLogsClient {
	return c.archiveLogs
}

// AlertSchedulers Gets an AlertsSchedulersClient from the ClientSet.
func (c *ClientSet) AlertSchedulers() *AlertsSchedulersClient {
	return c.alertsSchedulers
}

// Teams gets a TeamsClient from the ClientSet.
func (c *ClientSet) Teams() *TeamsClient {
	return c.teams
}

// APIKeys gets an SLOsClient from the ClientSet.
func (c *ClientSet) APIKeys() *ApikeysClient {
	return c.apiKeys
}

// SLOs gets an SLOsClient from the ClientSet.
func (c *ClientSet) SLOs() *SLOsClient {
	return c.slos
}

// Scopes gets a ScopesClient from the ClientSet.
func (c *ClientSet) Scopes() *ScopesClient {
	return c.scopes
}

// func (c *ClientSet) DashboardsFolders() *DashboardsFoldersClient {
// 	return c.dahboardsFolders
// }

// func (c *ClientSet) Users() *UsersClient {
// 	return c.users
// }

// NewClientSet Creates a new ClientSet.
func NewClientSet(targetURL, apiKey string) *ClientSet {
	apikeyCPC := NewCallPropertiesCreator(targetURL, apiKey)

	return &ClientSet{
		ruleGroups:        NewRuleGroupsClient(apikeyCPC),
		alerts:            NewAlertsClient(apikeyCPC),
		events2Metrics:    NewEvents2MetricsClient(apikeyCPC),
		enrichments:       NewEnrichmentClient(apikeyCPC),
		dataSet:           NewDataSetClient(apikeyCPC),
		dashboards:        NewDashboardsClient(apikeyCPC),
		actions:           NewActionsClient(apikeyCPC),
		tcoPolicies:       NewTCOPoliciesClient(apikeyCPC),
		webhooks:          NewWebhooksClient(apikeyCPC),
		archiveRetentions: NewArchiveRetentionsClient(apikeyCPC),
		archiveMetrics:    NewArchiveMetricsClient(apikeyCPC),
		archiveLogs:       NewArchiveLogsClient(apikeyCPC),
		alertsSchedulers:  NewAlertsSchedulersClient(apikeyCPC),
		teams:             NewTeamsClient(apikeyCPC),
		slos:              NewSLOsClient(apikeyCPC),
		scopes:            NewScopesClient(apikeyCPC),
		// dahboardsFolders:  NewDashboardsFoldersClient(apikeyCPC),
		apiKeys: NewAPIKeysClient(apikeyCPC),
		// users:             NewUsersClient(apikeyCPC),
	}
}

// EnvCoralogxRegion is the name of the environment variable that contains the Coralogix region.
const EnvCoralogxRegion = "CORALOGIX_REGION"

// EnvCoralogixAPIKey is the name of the environment variable that contains the Coralogix API key.
const EnvCoralogixAPIKey = "CORALOGIX_API_KEY"

// CoralogixRegionFromEnv reads the Coralogix region from environment variables.
func CoralogixRegionFromEnv() (string, error) {
	regionIdentifier := strings.ToLower(os.Getenv(EnvCoralogxRegion))
	if regionIdentifier == "" {
		return "", fmt.Errorf("the %s environment variable is not set", EnvCoralogxRegion)
	}

	switch regionIdentifier {
	case "us1":
		return RegionUS1, nil
	case "us2":
		return RegionUS2, nil
	case "eu1":
		return RegionEU1, nil
	case "eu2":
		return RegionEU2, nil
	case "ap1":
		return RegionAP1, nil
	case "ap2":
		return RegionAP2, nil
	default:
		return regionIdentifier, nil
	}
}

const (
	// RegionUS1 is the URL for the us1 region.
	RegionUS1 = "ng-api-grpc.coralogix.com:443"

	// RegionUS2 is the URL for the us2 region.
	RegionUS2 = "ng-api-grpc.cx498.coralogix.com:443"

	// RegionEU1 is the URL for the eu1 region.
	RegionEU1 = "ng-api-grpc.coralogix.com:443"

	// RegionEU2 is the URL for the eu2 region.
	RegionEU2 = "ng-api-grpc.eu2.coralogix.com:443"

	// RegionAP1 is the URL for the ap1 region.
	RegionAP1 = "ng-api-grpc.app.coralogix.in:443"

	// RegionAP2 is the URL for the ap2 region.
	RegionAP2 = "ng-api-grpc.coralogixsg.com:443"
)

// CoralogixAPIKeyFromEnv reads the Coralogix API key from environment variables.
func CoralogixAPIKeyFromEnv() (string, error) {
	apiKey := os.Getenv(EnvCoralogixAPIKey)
	if apiKey == "" {
		return "", fmt.Errorf("the %s environment variable is not set", EnvCoralogixAPIKey)
	}
	return apiKey, nil
}
