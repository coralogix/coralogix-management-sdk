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

	archiveMetrics "github.com/coralogix/coralogix-management-sdk/internal/coralogix/metrics"

	"google.golang.org/protobuf/types/known/emptypb"
)

// ConfigureTenantRequest is a request to configure the tenant.
type ConfigureTenantRequest = archiveMetrics.ConfigureTenantRequest

// UpdateTenantRequest is a request to update the tenant.
type UpdateTenantRequest = archiveMetrics.UpdateRequest

// ValidateBucketRequest is a request to validate the bucket.
type ValidateBucketRequest = archiveMetrics.ValidateBucketRequest

// GetTenantConfigRequest is a request to get the tenant configuration.
type GetTenantConfigRequest = archiveMetrics.GetTenantConfigRequest

// ConfigureTenantRequestS3 is a request to configure the tenant with S3.
type ConfigureTenantRequestS3 = archiveMetrics.ConfigureTenantRequest_S3

// ArchiveS3Config is an S3 configuration for the archive.
type ArchiveS3Config = archiveMetrics.S3Config

// ValidateBucketRequestS3 is a request to validate the S3 bucket.
type ValidateBucketRequestS3 = archiveMetrics.ValidateBucketRequest_S3

// ValidateBucketRequestIbm is a request to validate the IBM storage.
type ValidateBucketRequestIbm = archiveMetrics.ValidateBucketRequest_Ibm

// UpdateRequestS3 is a type to update the S3 bucket.
type UpdateRequestS3 = archiveMetrics.UpdateRequest_S3

// UpdateRequestIbm is a type to update the IBM storage.
type UpdateRequestIbm = archiveMetrics.UpdateRequest_Ibm

// TenantConfigV2S3 is a type to view the S3 bucket config.
type TenantConfigV2S3 = archiveMetrics.TenantConfigV2_S3

// TenantConfigV2Ibm is a type to view the IBM storage config.
type TenantConfigV2Ibm = archiveMetrics.TenantConfigV2_Ibm

// RetentionPolicyRequest is a request to set the retention policy.
type RetentionPolicyRequest = archiveMetrics.RetentionPolicyRequest

// ArchiveMetricsClient is a client for the Coralogix Archive Metrics API.
type ArchiveMetricsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Update updates the archive metrics configuration.
func (c ArchiveMetricsClient) Update(ctx context.Context, req *archiveMetrics.UpdateRequest) (*emptypb.Empty, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveMetrics.NewMetricsConfiguratorPublicServiceClient(conn)

	return client.Update(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get gets the archive metrics configuration.
func (c ArchiveMetricsClient) Get(ctx context.Context) (*archiveMetrics.GetTenantConfigResponseV2, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveMetrics.NewMetricsConfiguratorPublicServiceClient(conn)

	return client.GetTenantConfig(callProperties.Ctx, &emptypb.Empty{}, callProperties.CallOptions...)
}

// ConfigureTenant configures the archive metrics bucket.
func (c ArchiveMetricsClient) ConfigureTenant(ctx context.Context, req *archiveMetrics.ConfigureTenantRequest) (*emptypb.Empty, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveMetrics.NewMetricsConfiguratorPublicServiceClient(conn)

	return client.ConfigureTenant(callProperties.Ctx, req, callProperties.CallOptions...)
}

// ValidateTarget validates the archive metrics bucket.
func (c ArchiveMetricsClient) ValidateTarget(ctx context.Context, req *archiveMetrics.ValidateBucketRequest) (*emptypb.Empty, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveMetrics.NewMetricsConfiguratorPublicServiceClient(conn)

	return client.ValidateBucket(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Enable enables the metrics archive.
func (c ArchiveMetricsClient) Enable(ctx context.Context) (*emptypb.Empty, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveMetrics.NewMetricsConfiguratorPublicServiceClient(conn)

	return client.EnableArchive(callProperties.Ctx, &emptypb.Empty{}, callProperties.CallOptions...)
}

// Disable disables the metrics archive.
func (c ArchiveMetricsClient) Disable(ctx context.Context) (*emptypb.Empty, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := archiveMetrics.NewMetricsConfiguratorPublicServiceClient(conn)

	return client.DisableArchive(callProperties.Ctx, &emptypb.Empty{}, callProperties.CallOptions...)
}

// NewArchiveMetricsClient creates a new archive metrics client.
func NewArchiveMetricsClient(c *CallPropertiesCreator) *ArchiveMetricsClient {
	return &ArchiveMetricsClient{callPropertiesCreator: c}
}
