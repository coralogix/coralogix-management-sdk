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

// ServiceSlo is an SLO.
type ServiceSlo = slos.ServiceSlo

// CreateServiceSloRequest is a request to create an SLO.
type CreateServiceSloRequest = slos.CreateServiceSloRequest

// ServiceSloErrorSli is an SLO error SLI.
type ServiceSloErrorSli = slos.ServiceSlo_ErrorSli

// GetServiceSloRequest is a request to get an SLO.
type GetServiceSloRequest = slos.GetServiceSloRequest

// ReplaceServiceSloRequest is a request to replace an SLO.
type ReplaceServiceSloRequest = slos.ReplaceServiceSloRequest

// DeleteServiceSloRequest is a request to delete an SLO.
type DeleteServiceSloRequest = slos.DeleteServiceSloRequest

// SLO period values.
const (
	SloPeriod7Days       = slos.SloPeriod_SLO_PERIOD_7_DAYS
	SloPeriodUnspecified = slos.SloPeriod_SLO_PERIOD_UNSPECIFIED
	SloPeriod14Days      = slos.SloPeriod_SLO_PERIOD_14_DAYS
	SloPeriod30Days      = slos.SloPeriod_SLO_PERIOD_30_DAYS
)

// SLOsClient is a client for the Coralogix SLOs API.
type SLOsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new SLO.
func (c SLOsClient) Create(ctx context.Context, req *slos.CreateServiceSloRequest) (*slos.CreateServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewServiceSloServiceClient(conn)

	return client.CreateServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get gets the specified SLO.
func (c SLOsClient) Get(ctx context.Context, req *slos.GetServiceSloRequest) (*slos.GetServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewServiceSloServiceClient(conn)

	return client.GetServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Update updates the specified SLO.
func (c SLOsClient) Update(ctx context.Context, req *slos.ReplaceServiceSloRequest) (*slos.ReplaceServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewServiceSloServiceClient(conn)

	return client.ReplaceServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes the specified SLO.
func (c SLOsClient) Delete(ctx context.Context, req *slos.DeleteServiceSloRequest) (*slos.DeleteServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewServiceSloServiceClient(conn)

	return client.DeleteServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetBulk gets multiple SLOs in a single call.
func (c SLOsClient) GetBulk(ctx context.Context, req *slos.BatchGetServiceSlosRequest) (*slos.BatchGetServiceSlosResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewServiceSloServiceClient(conn)

	return client.BatchGetServiceSlos(callProperties.Ctx, req, callProperties.CallOptions...)
}

// List lists all service SLOs.
func (c SLOsClient) List(ctx context.Context, req *slos.ListServiceSlosRequest) (*slos.ListServiceSlosResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewServiceSloServiceClient(conn)

	return client.ListServiceSlos(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewSLOsClient creates a new SLOs client.
func NewSLOsClient(c *CallPropertiesCreator) *SLOsClient {
	return &SLOsClient{callPropertiesCreator: c}
}
