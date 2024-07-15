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
