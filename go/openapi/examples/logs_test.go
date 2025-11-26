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
	targets "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/target_service"
)

func TestArchiveLogs(t *testing.T) {
	logsBucket := os.Getenv("LOGS_BUCKET")
	awsRegion := os.Getenv("AWS_REGION")
	if logsBucket == "" || awsRegion == "" {
		t.Fatalf("LOGS_BUCKET or AWS_REGION environment variable are not set")
	}
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewArchiveLogsClient(cfg)

	setTargetReq := targets.SetTargetResponse{
		S3: &targets.S3TargetSpec{
			Bucket: logsBucket,
			Region: &awsRegion,
		},
	}

	_, httpResp, err := client.
		S3TargetServiceSetTarget(context.Background()).
		SetTargetResponse(setTargetReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	got, httpResp, err := client.
		S3TargetServiceGetTarget(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, got)

	if got.Target.TargetS3 != nil {
		require.Equal(t, logsBucket, got.Target.TargetS3.S3.Bucket)
		require.Equal(t, awsRegion, *got.Target.TargetS3.S3.Region)
	} else {
		t.Fatalf("expected S3 target, got nil")
	}
}
