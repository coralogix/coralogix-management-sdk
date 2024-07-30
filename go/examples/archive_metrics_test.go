package examples

// import (
// 	"context"
// 	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
// 	"github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/metrics"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestArchiveMetrics(t *testing.T) {

// 	region, err := cxsdk.CoralogixRegionFromEnv()
// 	assert.Nil(t, err)
// 	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
// 	assert.Nil(t, err)
// 	creator := cxsdk.NewCallPropertiesCreator(region, apiKey)
// 	c := cxsdk.NewArchiveMetricsClient(creator)

// 	_, configureTenantError := c.UpdateArchiveMetrics(context.Background(), &metrics.ConfigureTenantRequest{
// 		RetentionPolicy: nil,
// 		StorageConfig: &metrics.ConfigureTenantRequest_S3{
// 			S3: &metrics.S3Config{
// 				Bucket: "coralogix-c4c-eu2-prometheus-data",
// 				Region: "eu-west-1",
// 			},
// 		},
// 	})

// 	assert.Nil(t, configureTenantError)

// 	_, getTenantError := c.GetArchiveMetrics(context.Background())
// 	assert.Nil(t, getTenantError)
// }
