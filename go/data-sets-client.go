package cxsdk

import (
	"context"
	enrichment "coralogix-management-sdk/go/internal/coralogix/enrichment/v1"
)

type CreateDataSetRequest = enrichment.CreateCustomEnrichmentRequest
type GetDataSetRequest = enrichment.GetCustomEnrichmentRequest
type UpdateDataSetRequest = enrichment.UpdateCustomEnrichmentRequest
type DeleteDataSetRequest = enrichment.DeleteCustomEnrichmentRequest

type File = enrichment.File
type File_Binary = enrichment.File_Binary
type File_Textual = enrichment.File_Textual

type DataSetClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

func (d DataSetClient) CreatDataSet(ctx context.Context, req *CreateDataSetRequest) (*enrichment.CreateCustomEnrichmentResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()

	client := enrichment.NewCustomEnrichmentServiceClient(conn)

	return client.CreateCustomEnrichment(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (d DataSetClient) GetDataSet(ctx context.Context, req *GetDataSetRequest) (*enrichment.GetCustomEnrichmentResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()

	client := enrichment.NewCustomEnrichmentServiceClient(conn)

	return client.GetCustomEnrichment(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (d DataSetClient) UpdateDataSet(ctx context.Context, req *UpdateDataSetRequest) (*enrichment.UpdateCustomEnrichmentResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()

	client := enrichment.NewCustomEnrichmentServiceClient(conn)

	return client.UpdateCustomEnrichment(callProperties.Ctx, req, callProperties.CallOptions...)
}

func (d DataSetClient) DeleteDataSet(ctx context.Context, req *DeleteDataSetRequest) (*enrichment.DeleteCustomEnrichmentResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()

	client := enrichment.NewCustomEnrichmentServiceClient(conn)

	return client.DeleteCustomEnrichment(callProperties.Ctx, req, callProperties.CallOptions...)
}

func NewDataSetClient(c *CallPropertiesCreator) *DataSetClient {
	return &DataSetClient{callPropertiesCreator: c}
}
