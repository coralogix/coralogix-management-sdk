package examples

import (
	"context"
	"testing"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
	v2 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/archive/v2"

	"github.com/stretchr/testify/assert"
)

func TestArchiveLogs(t *testing.T) {

	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, apiKey)
	c := cxsdk.NewArchiveLogsClient(creator)
	s3Region := "eu-north-1"
	_, setTargetError := c.Update(context.Background(), &v2.SetTargetRequest{
		TargetSpec: &v2.SetTargetRequest_S3{
			S3: &v2.S3TargetSpec{
				Bucket: "yak-2-bucket",
				Region: &s3Region,
			},
		},
	})
	assert.Nil(t, setTargetError)

	_, getTargetError := c.Get(context.Background())

	assert.Nil(t, getTargetError)
}
