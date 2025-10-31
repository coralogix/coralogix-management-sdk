package examples

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	targets "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/target_service"
)

func TestArchiveLogs(t *testing.T) {
	logsBucket := os.Getenv("LOGS_BUCKET")
	awsRegion := os.Getenv("AWS_REGION")
	if logsBucket == "" || awsRegion == "" {
		t.Fatalf("LOGS_BUCKET or AWS_REGION environment variable is not set")
	}
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)
	client := cxsdk.NewArchiveLogsTargetClient(cpc)

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
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	got, httpResp, err := client.
		S3TargetServiceGetTarget(context.Background()).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotNil(t, got)

	if got.Target.TargetS3 != nil {
		assert.Equal(t, logsBucket, got.Target.TargetS3.S3.Bucket)
		assert.Equal(t, awsRegion, *got.Target.TargetS3.S3.Region)
	} else {
		t.Fatalf("expected S3 target, got nil")
	}
}
