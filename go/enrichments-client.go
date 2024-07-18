package cxsdk

import (
	"context"
	enrichment "coralogix-management-sdk/go/internal/coralogix/enrichment/v1"
)

// EnrichmentsClient is a client for the Coralogix Enrichments API.
type EnrichmentsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// CreateEnrichments creates a new enrichment.
func (e EnrichmentsClient) CreateEnrichments(ctx context.Context, req *enrichment.AddEnrichmentsRequest) ([]*enrichment.Enrichment, error) {
	callProperties, err := e.callPropertiesCreator.GetCallProperties(ctx)
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

// GetEnrichmentsByType gets all enrichments of a certain type.
func (e EnrichmentsClient) GetEnrichmentsByType(ctx context.Context, enrichmentType string) ([]*enrichment.Enrichment, error) {
	callProperties, err := e.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := enrichment.NewEnrichmentServiceClient(conn)

	resp, err := client.GetEnrichments(callProperties.Ctx, &enrichment.GetEnrichmentsRequest{}, callProperties.CallOptions...)
	if err != nil {
		return nil, err
	}

	result := make([]*enrichment.Enrichment, 0)
	for _, enrichment := range resp.GetEnrichments() {
		if enrichment.GetEnrichmentType().String() == enrichmentType+":{}" {
			result = append(result, enrichment)
		}
	}

	return result, nil
}

// GetCustomEnrichments gets all custom enrichments.
func (e EnrichmentsClient) GetCustomEnrichments(ctx context.Context, customEnrichmentID uint32) ([]*enrichment.Enrichment, error) {
	callProperties, err := e.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := enrichment.NewEnrichmentServiceClient(conn)

	resp, err := client.GetEnrichments(callProperties.Ctx, &enrichment.GetEnrichmentsRequest{}, callProperties.CallOptions...)
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

// DeleteEnrichments deletes the specified enrichments.
func (e EnrichmentsClient) DeleteEnrichments(ctx context.Context, req *enrichment.RemoveEnrichmentsRequest) error {
	callProperties, err := e.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return err
	}

	conn := callProperties.Connection
	defer conn.Close()

	client := enrichment.NewEnrichmentServiceClient(conn)

	_, err = client.RemoveEnrichments(callProperties.Ctx, req, callProperties.CallOptions...)
	return err
}

// NewEnrichmentClient creates a new enrichments client.
func NewEnrichmentClient(c *CallPropertiesCreator) *EnrichmentsClient {
	return &EnrichmentsClient{callPropertiesCreator: c}
}
