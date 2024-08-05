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
	"context"

	dashboards "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/dashboards/v1"
)

// DashboardsCatalogClient is a client for the Coralogix Dashboards API.
type DashboardsCatalogClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// GetDashboardCatalogRequest is a request to get a dashboard.
type GetDashboardCatalogRequest = dashboards.GetDashboardCatalogRequest

// Get gets a dashboards catalog.
func (d DashboardsCatalogClient) Get(ctx context.Context, req *dashboards.GetDashboardCatalogRequest) (*dashboards.GetDashboardCatalogResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dashboards.NewDashboardCatalogServiceClient(conn)

	return client.GetDashboardCatalog(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewDashboardsCatalogClient creates a new NewDashboardsCatalogClient.
func NewDashboardsCatalogClient(c *CallPropertiesCreator) *DashboardsCatalogClient {
	return &DashboardsCatalogClient{callPropertiesCreator: c}
}
