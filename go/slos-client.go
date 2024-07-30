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

	slos "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm/services/v1"
)

// SLOsClient is a client for the Coralogix SLOs API.
type SLOsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// CreateSLO creates a new SLO.
func (c SLOsClient) CreateSLO(ctx context.Context, req *slos.CreateServiceSloRequest) (*slos.CreateServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewServiceSloServiceClient(conn)

	return client.CreateServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetSLO gets the specified SLO.
func (c SLOsClient) GetSLO(ctx context.Context, req *slos.GetServiceSloRequest) (*slos.GetServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewServiceSloServiceClient(conn)

	return client.GetServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
}

// UpdateSLO updates the specified SLO.
func (c SLOsClient) UpdateSLO(ctx context.Context, req *slos.ReplaceServiceSloRequest) (*slos.ReplaceServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewServiceSloServiceClient(conn)

	return client.ReplaceServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
}

// DeleteSLO deletes the specified SLO.
func (c SLOsClient) DeleteSLO(ctx context.Context, req *slos.DeleteServiceSloRequest) (*slos.DeleteServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewServiceSloServiceClient(conn)

	return client.DeleteServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewSLOsClient creates a new SLOs client.
func NewSLOsClient(c *CallPropertiesCreator) *SLOsClient {
	return &SLOsClient{callPropertiesCreator: c}
}
