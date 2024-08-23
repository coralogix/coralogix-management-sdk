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

	archiveRetention "github.com/coralogix/coralogix-management-sdk/internal/coralogix/archive/v1"
)

// ArchiveRetentionsClient is a client for the Coralogix Archive Retentions API.
type ArchiveRetentionsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Get gets the archive retentions.
func (c ArchiveRetentionsClient) Get(ctx context.Context, req *archiveRetention.GetRetentionsRequest) (*archiveRetention.GetRetentionsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveRetention.NewRetentionsServiceClient(conn)

	return client.GetRetentions(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Update updates the archive retentions.
func (c ArchiveRetentionsClient) Update(ctx context.Context, req *archiveRetention.UpdateRetentionsRequest) (*archiveRetention.UpdateRetentionsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveRetention.NewRetentionsServiceClient(conn)

	return client.UpdateRetentions(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Activate activates the archive retentions.
func (c ArchiveRetentionsClient) Activate(ctx context.Context, req *archiveRetention.ActivateRetentionsRequest) (*archiveRetention.ActivateRetentionsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveRetention.NewRetentionsServiceClient(conn)

	return client.ActivateRetentions(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetEnabled returns a boolean that signals whether archive retentions are enabled.
func (c ArchiveRetentionsClient) GetEnabled(ctx context.Context, req *archiveRetention.GetRetentionsEnabledRequest) (*archiveRetention.GetRetentionsEnabledResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveRetention.NewRetentionsServiceClient(conn)

	return client.GetRetentionsEnabled(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewArchiveRetentionsClient Creates a new archive retentions client.
func NewArchiveRetentionsClient(c *CallPropertiesCreator) *ArchiveRetentionsClient {
	return &ArchiveRetentionsClient{callPropertiesCreator: c}
}
