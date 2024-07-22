package cxsdk

import (
	"context"
	dataprime "coralogix-management-sdk/go/internal/coralogixapis/dataprime/v1"
)

// DataprimeClient is a client for the Coralogix Dataprime API.
type DataprimeClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Runs a query.
func (c DataprimeClient) Query(ctx context.Context, req *dataprime.QueryRequest) (dataprime.DataprimeQueryService_QueryClient, error) {
	callProperties, err := c.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := dataprime.NewDataprimeQueryServiceClient(conn)

	return client.Query(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewDataprimeClient creates a new dataprime client.
func NewDataprimeClient(c *CallPropertiesCreator) *DataprimeClient {
	return &DataprimeClient{callPropertiesCreator: c}
}
