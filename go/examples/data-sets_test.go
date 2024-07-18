package examples

import (
	"context"
	cxsdk "coralogix-management-sdk/go"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestDataSets(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, apiKey)
	c := cxsdk.NewDataSetClient(creator)

	raw, e := os.ReadFile("date-to-day-of-the-week.csv")
	assert.Nil(t, e)
	textual := string(raw)

	data, e := c.CreateDataSet(context.Background(), &cxsdk.CreateDataSetRequest{
		Name:        wrapperspb.String("my-enrichment"),
		Description: wrapperspb.String("My custom enrichment description"),
		File: &cxsdk.File{
			Name:      wrapperspb.String("date-to-day-of-the-week"),
			Extension: wrapperspb.String("csv"),
			Content: &cxsdk.FileTextual{
				Textual: wrapperspb.String(textual),
			},
		},
	})
	assert.Nil(t, e)

	updated, e := c.UpdateDataSet(context.Background(), &cxsdk.UpdateDataSetRequest{
		CustomEnrichmentId: wrapperspb.UInt32(data.CustomEnrichment.Id),
		Description:        wrapperspb.String("My updated enrichment description"),
		File: &cxsdk.File{
			Name:      wrapperspb.String("date-to-day-of-the-week"),
			Extension: wrapperspb.String("csv"),
			Content: &cxsdk.FileTextual{
				Textual: wrapperspb.String(textual),
			},
		},
	})
	assert.Nil(t, e)
	fetched, e := c.GetDataSet(context.Background(), &cxsdk.GetDataSetRequest{
		Id: wrapperspb.UInt32(updated.CustomEnrichment.Id),
	})
	assert.Nil(t, e)

	assert.Equal(t, updated.CustomEnrichment.Description, fetched.CustomEnrichment.Description)
	assert.Equal(t, data.CustomEnrichment.Version+1, fetched.CustomEnrichment.Version)
	c.DeleteDataSet(context.Background(), &cxsdk.DeleteDataSetRequest{
		CustomEnrichmentId: wrapperspb.UInt32(fetched.CustomEnrichment.Id),
	})
}
