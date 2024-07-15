package cxsdk

import (
	"context"
	archiveRetention "coralogix-management-sdk/go/internal/coralogix/archive/v1"
)

// ArchiveRetentionsClient is a client for the Coralogix Archive Retentions API.
type ArchiveRetentionsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// GetRetentions gets the archive retentions.
func (c ArchiveRetentionsClient) GetRetentions(ctx context.Context, req *archiveRetention.GetRetentionsRequest) (*archiveRetention.GetRetentionsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveRetention.NewRetentionsServiceClient(conn)

	return client.GetRetentions(callProperties.Ctx, req, callProperties.CallOptions...)
}

// UpdateRetentions updates the archive retentions.
func (c ArchiveRetentionsClient) UpdateRetentions(ctx context.Context, req *archiveRetention.UpdateRetentionsRequest) (*archiveRetention.UpdateRetentionsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveRetention.NewRetentionsServiceClient(conn)

	return client.UpdateRetentions(callProperties.Ctx, req, callProperties.CallOptions...)
}

// ActivateRetentions activates the archive retentions.
func (c ArchiveRetentionsClient) ActivateRetentions(ctx context.Context, req *archiveRetention.ActivateRetentionsRequest) (*archiveRetention.ActivateRetentionsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveRetention.NewRetentionsServiceClient(conn)

	return client.ActivateRetentions(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetRetentionsEnabled returns a boolean that signals whether archive retentions are enabled.
func (c ArchiveRetentionsClient) GetRetentionsEnabled(ctx context.Context, req *archiveRetention.GetRetentionsEnabledRequest) (*archiveRetention.GetRetentionsEnabledResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
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
