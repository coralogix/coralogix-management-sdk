package examples

import (
	"context"
	cxsdk "coralogix-management-sdk/go"
	v2 "coralogix-management-sdk/go/internal/coralogix/archive/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArchiveLogs(t *testing.T) {

	region, err := cxsdk.CoralogixGrpcEndpointFromEnv()
	assert.Nil(t, err)
	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, apiKey)
	c := cxsdk.NewArchiveLogsClient(creator)
	s3Region := "eu-north-1"
	_, setTargetError := c.UpdateArchiveLogs(context.Background(), &v2.SetTargetRequest{
		TargetSpec: &v2.SetTargetRequest_S3{
			S3: &v2.S3TargetSpec{
				Bucket: "yak-2-bucket",
				Region: &s3Region,
			},
		},
	})
	assert.Nil(t, setTargetError)

	_, getTargetError := c.GetArchiveLogs(context.Background())

	assert.Nil(t, getTargetError)
}
