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

// CreateDataSetRequest is a request to create a data set.
type CreateDataSetRequest = enrichment.CreateCustomEnrichmentRequest

// GetDataSetRequest is a request to get a data set.
type GetDataSetRequest = enrichment.GetCustomEnrichmentRequest

// UpdateDataSetRequest is a request to update a data set.
type UpdateDataSetRequest = enrichment.UpdateCustomEnrichmentRequest

// DeleteDataSetRequest is a request to delete a data set.
type DeleteDataSetRequest = enrichment.DeleteCustomEnrichmentRequest

// File is a file. Can be either FileBinary or FileTextual.
type File = enrichment.File

// FileBinary is a binary file.
type FileBinary = enrichment.File_Binary

// FileTextual is a textual file.
type FileTextual = enrichment.File_Textual

// DataSetClient is a client for the Coralogix Data Sets API.
type DataSetClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// CreateDataSet creates a new data set.
func (d DataSetClient) CreateDataSet(ctx context.Context, req *CreateDataSetRequest) (*enrichment.CreateCustomEnrichmentResponse, error) {
	callProperties, err := d.callPropertiesCreator.GetCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()

	client := enrichment.NewCustomEnrichmentServiceClient(conn)

	return client.CreateCustomEnrichment(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetDataSet gets a data set.
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

// UpdateDataSet updates a data set.
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

// DeleteDataSet deletes a data set.
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

// NewDataSetClient creates a new data set client.
func NewDataSetClient(c *CallPropertiesCreator) *DataSetClient {
	return &DataSetClient{callPropertiesCreator: c}
}
