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

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"

	"github.com/stretchr/testify/assert"
)

func TestArchiveMetrics(t *testing.T) {
	metricsBucket := os.Getenv("METRICS_BUCKET")
	awsRegion := os.Getenv("AWS_REGION")
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewArchiveMetricsClient(creator)
	s3Config := &cxsdk.ArchiveS3Config{
		Bucket: metricsBucket,
		Region: awsRegion,
	}

	_, configureErr := c.ConfigureTenant(context.Background(), &cxsdk.ConfigureTenantRequest{
		StorageConfig: &cxsdk.ConfigureTenantRequestS3{
			S3: s3Config,
		},
	})
	assertNilAndPrintError(t, configureErr)

	_, validateErr := c.ValidateTarget(context.Background(), &cxsdk.ValidateBucketRequest{
		StorageConfig: &cxsdk.ValidateBucketRequestS3{
			S3: s3Config,
		},
	})
	assertNilAndPrintError(t, validateErr)

	days := uint32(2)
	_, updateErr := c.Update(context.Background(), &cxsdk.UpdateTenantRequest{
		RetentionDays: &days,
		StorageConfig: &cxsdk.UpdateRequestS3{
			S3: s3Config,
		},
	})
	assertNilAndPrintError(t, updateErr)

	config, getTenantError := c.Get(context.Background())

	assertNilAndPrintError(t, getTenantError)
	assert.Equal(t, config.TenantConfig.StorageConfig.(*cxsdk.TenantConfigV2S3).S3.Bucket, s3Config.Bucket)
	_, e := c.Enable(context.Background())
	assertNilAndPrintError(t, e)
	_, e = c.Disable(context.Background())
	assertNilAndPrintError(t, e)
}
