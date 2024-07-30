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

	archiveMetrics "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/metrics"

	"google.golang.org/protobuf/types/known/emptypb"
)

// ArchiveMetricsClient is a client for the Coralogix Archive Metrics API.
type ArchiveMetricsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// UpdateArchiveMetrics updates the archive metrics configuration.
func (c ArchiveMetricsClient) UpdateArchiveMetrics(ctx context.Context, req *archiveMetrics.ConfigureTenantRequest) (*emptypb.Empty, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveMetrics.NewMetricsConfiguratorPublicServiceClient(conn)

	return client.ConfigureTenant(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetArchiveMetrics gets the archive metrics configuration.
func (c ArchiveMetricsClient) GetArchiveMetrics(ctx context.Context) (*archiveMetrics.GetTenantConfigResponseV2, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveMetrics.NewMetricsConfiguratorPublicServiceClient(conn)

	return client.GetTenantConfig(callProperties.Ctx, &emptypb.Empty{}, callProperties.CallOptions...)
}

// NewArchiveMetricsClient creates a new archive metrics client.
func NewArchiveMetricsClient(c *CallPropertiesCreator) *ArchiveMetricsClient {
	return &ArchiveMetricsClient{callPropertiesCreator: c}
}
