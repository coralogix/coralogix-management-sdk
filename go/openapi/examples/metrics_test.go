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

package examples

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	metrics "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/metrics_data_archive_service"
)

func TestArchiveMetrics(t *testing.T) {
	metricsBucket := os.Getenv("METRICS_BUCKET")
	awsRegion := os.Getenv("AWS_REGION")
	if metricsBucket == "" || awsRegion == "" {
		t.Fatalf("METRICS_BUCKET or AWS_REGION environment variable are not set")
	}

	ctx := context.Background()
	region, _ := cxsdk.URLFromRegion(cxsdk.RegionFromEnv())
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		region,
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewArchiveMetricsClient(cpc)

	s3Config := metrics.S3Config{
		Bucket: &metricsBucket,
		Region: &awsRegion,
	}

	configureReq := metrics.MetricsConfiguratorPublicServiceConfigureTenantRequest{
		ConfigureTenantRequestS3: &metrics.ConfigureTenantRequestS3{
			S3: &s3Config,
		},
	}
	_, httpResp, err := client.
		MetricsConfiguratorPublicServiceConfigureTenant(ctx).
		MetricsConfiguratorPublicServiceConfigureTenantRequest(configureReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	validateReq := metrics.MetricsConfiguratorPublicServiceValidateBucketRequest{
		ValidateBucketRequestS3: &metrics.ValidateBucketRequestS3{
			S3: &s3Config,
		},
	}
	_, httpResp, err = client.
		MetricsConfiguratorPublicServiceValidateBucket(ctx).
		MetricsConfiguratorPublicServiceValidateBucketRequest(validateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	retentionDays := int64(30)
	updateReq := metrics.MetricsConfiguratorPublicServiceUpdateRequest{
		UpdateRequestS3: &metrics.UpdateRequestS3{
			RetentionDays: &retentionDays,
			S3:            &s3Config,
		},
	}

	_, httpResp, err = client.
		MetricsConfiguratorPublicServiceUpdate(ctx).
		MetricsConfiguratorPublicServiceUpdateRequest(updateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	getResp, httpResp, err := client.
		MetricsConfiguratorPublicServiceGetTenantConfig(ctx).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, getResp)

	if getResp != nil && getResp.TenantConfig != nil && getResp.TenantConfig.TenantConfigV2S3 != nil {
		require.Equal(t, metricsBucket, *getResp.TenantConfig.TenantConfigV2S3.S3.Bucket)
		require.Equal(t, awsRegion, *getResp.TenantConfig.TenantConfigV2S3.S3.Region)
	} else {
		t.Fatalf("expected S3 tenant config, got nil")
	}

	_, httpResp, err = client.
		MetricsConfiguratorPublicServiceEnableArchive(ctx).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		MetricsConfiguratorPublicServiceDisableArchive(ctx).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}
