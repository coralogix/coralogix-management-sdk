package cxsdk

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
	// dahboardsFolders  *DashboardsFoldersClient
	// users             *UsersClient
}

func (c *ClientSet) RuleGroups() *RuleGroupsClient {
	return c.ruleGroups
}

func (c *ClientSet) Alerts() *AlertsClient {
	return c.alerts
}

func (c *ClientSet) Enrichments() *EnrichmentsClient {
	return c.enrichments
}

func (c *ClientSet) DataSet() *DataSetClient {
	return c.dataSet
}

func (c *ClientSet) Dashboards() *DashboardsClient {
	return c.dashboards
}

func (c *ClientSet) Actions() *ActionsClient {
	return c.actions
}

func (c *ClientSet) TCOPolicies() *TCOPoliciesClient {
	return c.tcoPolicies
}

func (c *ClientSet) Webhooks() *WebhooksClient {
	return c.webhooks
}

func (c *ClientSet) Events2Metrics() *Events2MetricsClient {
	return c.events2Metrics
}

func (c *ClientSet) ArchiveRetentions() *ArchiveRetentionsClient {
	return c.archiveRetentions
}

func (c *ClientSet) ArchiveMetrics() *ArchiveMetricsClient {
	return c.archiveMetrics
}

func (c *ClientSet) ArchiveLogs() *ArchiveLogsClient {
	return c.archiveLogs
}

func (c *ClientSet) AlertSchedulers() *AlertsSchedulersClient {
	return c.alertsSchedulers
}

func (c *ClientSet) Teams() *TeamsClient {
	return c.teams
}

func (c *ClientSet) ApiKeys() *ApikeysClient {
	return c.apiKeys
}

func (c *ClientSet) SLOs() *SLOsClient {
	return c.slos
}

func (c *ClientSet) Scopes() *ScopesClient {
	return c.scopes
}

// func (c *ClientSet) DashboardsFolders() *DashboardsFoldersClient {
// 	return c.dahboardsFolders
// }

// func (c *ClientSet) Users() *UsersClient {
// 	return c.users
// }

func NewClientSet(targetUrl, apiKey string) *ClientSet {
	apikeyCPC := NewCallPropertiesCreator(targetUrl, apiKey)

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
		apiKeys: NewApiKeysClient(apikeyCPC),
		// users:             NewUsersClient(apikeyCPC),
	}
}
