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

	enrichment "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/enrichment/v1"
)

// Enrichment is a column mapping for enrichments.
type Enrichment = enrichment.Enrichment

// GetEnrichmentsRequest is a request for getting enrichments.
type GetEnrichmentsRequest = enrichment.GetEnrichmentsRequest

// GetEnrichmentLimitRequest is a request for getting the enrichment limits.
type GetEnrichmentLimitRequest = enrichment.GetEnrichmentLimitRequest

// AddEnrichmentsRequest is a request to add enrichments.
type AddEnrichmentsRequest = enrichment.AddEnrichmentsRequest

// RemoveEnrichmentsRequest is a request to remove enrichments.
type RemoveEnrichmentsRequest = enrichment.RemoveEnrichmentsRequest

// EnrichmentsClient is a client for the Coralogix Enrichments API.
type EnrichmentsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new enrichment.
func (e EnrichmentsClient) Create(ctx context.Context, req *AddEnrichmentsRequest) ([]*enrichment.Enrichment, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := enrichment.NewEnrichmentServiceClient(conn)

	resp, err := client.AddEnrichments(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, err
	}

	enrichments := resp.GetEnrichments()
	from := len(enrichments) - len(req.GetRequestEnrichments())
	to := len(enrichments)
	return enrichments[from:to], nil
}

// GetByType gets all enrichments of a certain type.
func (e EnrichmentsClient) GetByType(ctx context.Context, enrichmentType string) ([]*enrichment.Enrichment, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := enrichment.NewEnrichmentServiceClient(conn)

	resp, err := client.GetEnrichments(callProperties.Ctx, &GetEnrichmentsRequest{}, callProperties.CallOptions...)
	if err != nil {
		return nil, err
	}

	result := make([]*Enrichment, 0)
	for _, enrichment := range resp.GetEnrichments() {
		if enrichment.GetEnrichmentType().String() == enrichmentType+":{}" {
			result = append(result, enrichment)
		}
	}

	return result, nil
}

// Get gets all custom enrichments.
func (e EnrichmentsClient) Get(ctx context.Context, customEnrichmentID uint32) ([]*enrichment.Enrichment, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := enrichment.NewEnrichmentServiceClient(conn)

	resp, err := client.GetEnrichments(callProperties.Ctx, &GetEnrichmentsRequest{}, callProperties.CallOptions...)
	if err != nil {
		return nil, err
	}

	result := make([]*enrichment.Enrichment, 0)
	for _, enrichment := range resp.GetEnrichments() {
		if customEnrichment := enrichment.GetEnrichmentType().GetCustomEnrichment(); customEnrichment != nil && customEnrichment.GetId().GetValue() == customEnrichmentID {
			result = append(result, enrichment)
		}
	}

	return result, nil
}

// Delete deletes the specified enrichments.
func (e EnrichmentsClient) Delete(ctx context.Context, req *RemoveEnrichmentsRequest) error {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return err
	}

	conn := callProperties.Connection
	defer conn.Close()

	client := enrichment.NewEnrichmentServiceClient(conn)

	_, err = client.RemoveEnrichments(callProperties.Ctx, req, callProperties.CallOptions...)
	return err
}

// List returns all enrichments.
func (e EnrichmentsClient) List(ctx context.Context, req *GetEnrichmentsRequest) ([]*enrichment.Enrichment, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()

	client := enrichment.NewEnrichmentServiceClient(conn)

	resp, err := client.GetEnrichments(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, err
	}

	return resp.GetEnrichments(), nil
}

// GetLimits returns the enrichment limits.
func (e EnrichmentsClient) GetLimits(ctx context.Context) (*enrichment.GetEnrichmentLimitResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()

	client := enrichment.NewEnrichmentServiceClient(conn)

	return client.GetEnrichmentLimit(callProperties.Ctx, &GetEnrichmentLimitRequest{}, callProperties.CallOptions...)
}

// NewEnrichmentClient creates a new enrichments client.
func NewEnrichmentClient(c *CallPropertiesCreator) *EnrichmentsClient {
	return &EnrichmentsClient{callPropertiesCreator: c}
}
