package examples

import (
	"context"
	"testing"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestDataprimeQuery(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, apiKey)
	c := cxsdk.NewDataprimeClient(creator)
	data, e := c.Query(context.Background(), &cxsdk.QueryRequest{
		Query:    wrapperspb.String("filter log_obj.message ~ 'Hello world'"),
		Metadata: nil,
	})
	assert.Nil(t, e)
	assert.NotNil(t, data)
}
