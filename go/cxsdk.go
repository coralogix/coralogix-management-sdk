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
	users             *UsersClient
	groups            *GroupsClient
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

// Users gets a UsersClient from the ClientSet.
func (c *ClientSet) Users() *UsersClient {
	return c.users
}

// Groups gets a GroupsClient from the ClientSet.
func (c *ClientSet) Groups() *GroupsClient {
	return c.groups
}

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
		groups:  NewGroupsClient(apikeyCPC),
		// users:             NewUsersClient(apikeyCPC),
	}
}

// EnvCoralogxRegion is the name of the environment variable that contains the Coralogix region.
const EnvCoralogxRegion = "CORALOGIX_REGION"

// EnvCoralogixAPIKey is the name of the environment variable that contains the Coralogix API key.
const EnvCoralogixAPIKey = "CORALOGIX_API_KEY"

// CoralogixGrpcEndpointFromEnv reads the Coralogix region from environment variables.
func CoralogixGrpcEndpointFromEnv() (string, error) {
	regionIdentifier := strings.ToLower(os.Getenv(EnvCoralogxRegion))
	if regionIdentifier == "" {
		return "", fmt.Errorf("the %s environment variable is not set", EnvCoralogxRegion)
	}

	switch regionIdentifier {
	case "us1":
		return GrpcUS1, nil
	case "us2":
		return GrpcUS2, nil
	case "eu1":
		return GrpcEU1, nil
	case "eu2":
		return GrpcEU2, nil
	case "ap1":
		return GrpcAP1, nil
	case "ap2":
		return GrpcAP2, nil
	default:
		return regionIdentifier, nil
	}
}

// CoralogixRestEndpointFromEnv reads the Coralogix REST endpoint from environment variables.
func CoralogixRestEndpointFromEnv() (string, error) {
	regionIdentifier := strings.ToLower(os.Getenv(EnvCoralogxRegion))
	if regionIdentifier == "" {
		return "", fmt.Errorf("the %s environment variable is not set", EnvCoralogxRegion)
	}

	switch regionIdentifier {
	case "us1":
		return RestUS1, nil
	case "us2":
		return RestUS2, nil
	case "eu1":
		return RestEU1, nil
	case "eu2":
		return RestEU2, nil
	case "ap1":
		return RestAP1, nil
	case "ap2":
		return RestAP2, nil
	default:
		return regionIdentifier, nil
	}
}

const (
	// GrpcUS1 is the URL for the us1 region.
	GrpcUS1 = "ng-api-grpc.coralogix.com:443"

	// GrpcUS2 is the URL for the us2 region.
	GrpcUS2 = "ng-api-grpc.cx498.coralogix.com:443"

	// GrpcEU1 is the URL for the eu1 region.
	GrpcEU1 = "ng-api-grpc.coralogix.com:443"

	// GrpcEU2 is the URL for the eu2 region.
	GrpcEU2 = "ng-api-grpc.eu2.coralogix.com:443"

	// GrpcAP1 is the URL for the ap1 region.
	GrpcAP1 = "ng-api-grpc.app.coralogix.in:443"

	// GrpcAP2 is the URL for the ap2 region.
	GrpcAP2 = "ng-api-grpc.coralogixsg.com:443"
)

const (
	// RestUS1 is the URL for the us1 region.
	RestUS1 = "https://ng-api-http.coralogix.com"

	// RestUS2 is the URL for the us2 region.
	RestUS2 = "https://ng-api-http.cx498.coralogix.com"

	// RestEU1 is the URL for the eu1 region.
	RestEU1 = "https://ng-api-http.coralogix.com"

	// RestEU2 is the URL for the eu2 region.
	RestEU2 = "https://ng-api-http.eu2.coralogix.com"

	// RestAP1 is the URL for the ap1 region.
	RestAP1 = "https://ng-api-http.app.coralogix.in"

	// RestAP2 is the URL for the ap2 region.
	RestAP2 = "https://ng-api-http.coralogixsg.com"
)

// CoralogixAPIKeyFromEnv reads the Coralogix API key from environment variables.
func CoralogixAPIKeyFromEnv() (string, error) {
	apiKey := os.Getenv(EnvCoralogixAPIKey)
	if apiKey == "" {
		return "", fmt.Errorf("the %s environment variable is not set", EnvCoralogixAPIKey)
	}
	return apiKey, nil
}
