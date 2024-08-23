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

	e2m "github.com/coralogix/coralogix-management-sdk/internal/coralogixapis/events2metrics/v2"
)

// Events2MetricsClient is a client for the Coralogix Events2Metrics API. Read more at https://coralogix.com/docs/events2metrics/
type Events2MetricsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create Creates a new metric.
func (e Events2MetricsClient) Create(ctx context.Context, req *e2m.CreateE2MRequest) (*e2m.CreateE2MResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := e2m.NewEvents2MetricServiceClient(conn)

	return client.CreateE2M(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get gets a metric.
func (e Events2MetricsClient) Get(ctx context.Context, req *e2m.GetE2MRequest) (*e2m.GetE2MResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := e2m.NewEvents2MetricServiceClient(conn)

	return client.GetE2M(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Replace replaces a metric.
func (e Events2MetricsClient) Replace(ctx context.Context, req *e2m.ReplaceE2MRequest) (*e2m.ReplaceE2MResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := e2m.NewEvents2MetricServiceClient(conn)

	return client.ReplaceE2M(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes a metric.
func (e Events2MetricsClient) Delete(ctx context.Context, req *e2m.DeleteE2MRequest) (*e2m.DeleteE2MResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := e2m.NewEvents2MetricServiceClient(conn)

	return client.DeleteE2M(callProperties.Ctx, req, callProperties.CallOptions...)
}

// ListLabelsCardinality lists the cardinality of labels for a given metric.
func (e Events2MetricsClient) ListLabelsCardinality(ctx context.Context, req *e2m.ListLabelsCardinalityRequest) (*e2m.ListLabelsCardinalityResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := e2m.NewEvents2MetricServiceClient(conn)

	return client.ListLabelsCardinality(callProperties.Ctx, req, callProperties.CallOptions...)
}

// List lists all metrics
func (e Events2MetricsClient) List(ctx context.Context) (*e2m.ListE2MResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := e2m.NewEvents2MetricServiceClient(conn)

	return client.ListE2M(callProperties.Ctx, &e2m.ListE2MRequest{}, callProperties.CallOptions...)
}

// NewEvents2MetricsClient creates a new Events2MetricsClient.
func NewEvents2MetricsClient(c *CallPropertiesCreator) *Events2MetricsClient {
	return &Events2MetricsClient{callPropertiesCreator: c}
}
