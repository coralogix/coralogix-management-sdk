package cxsdk

import (
	"context"
	archiveMetrics "coralogix-management-sdk/go/internal/coralogix/metrics"

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
