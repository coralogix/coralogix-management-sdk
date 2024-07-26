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
	archiveLogs "coralogix-management-sdk/go/internal/coralogix/archive/v2"
)

// ArchiveLogsClient is a client for the Coralogix Archive Logs API.
type ArchiveLogsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// UpdateArchiveLogs updates the archive logs target.
func (c ArchiveLogsClient) UpdateArchiveLogs(ctx context.Context, req *archiveLogs.SetTargetRequest) (*archiveLogs.SetTargetResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveLogs.NewTargetServiceClient(conn)

	return client.SetTarget(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetArchiveLogs gets the archive logs target.
func (c ArchiveLogsClient) GetArchiveLogs(ctx context.Context) (*archiveLogs.GetTargetResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveLogs.NewTargetServiceClient(conn)

	return client.GetTarget(callProperties.Ctx, &archiveLogs.GetTargetRequest{}, callProperties.CallOptions...)
}

// NewArchiveLogsClient creates a new archive logs client.
func NewArchiveLogsClient(c *CallPropertiesCreator) *ArchiveLogsClient {
	return &ArchiveLogsClient{callPropertiesCreator: c}
}
