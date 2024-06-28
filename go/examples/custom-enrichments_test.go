package examples

import (
	"context"
	cxsdk "coralogix-management-sdk/go"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestCustomEnrichments(t *testing.T) {
	creator := cxsdk.NewCallPropertiesCreator("https://ng-api-grpc.coralogix.com", "my-secret-token")
	c := cxsdk.NewDataSetClient(creator)

	raw, e := os.ReadFile("data-to-day-of-the-week.csv")
	assert.Nil(t, e)
	textual := string(raw)

	data, e := c.CreatDataSet(context.Background(), &cxsdk.CreateDataSetRequest{
		Name:        wrapperspb.String("My custom enrichment"),
		Description: wrapperspb.String("My custom enrichment description"),
		File: &cxsdk.File{
			Name:      wrapperspb.String("date-to-day-of-the-week"),
			Extension: wrapperspb.String("csv"),
			Content: &cxsdk.File_Textual{
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
			Content: &cxsdk.File_Textual{
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
	assert.Equal(t, data.CustomEnrichment.Version, fetched.CustomEnrichment.Version+1)
	c.DeleteDataSet(context.Background(), &cxsdk.DeleteDataSetRequest{
		CustomEnrichmentId: wrapperspb.UInt32(fetched.CustomEnrichment.Id),
	})
}
