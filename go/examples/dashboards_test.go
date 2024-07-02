package examples

import (
	"context"
	cxsdk "coralogix-management-sdk/go"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestDashboards(t *testing.T) {
	creator := cxsdk.NewCallPropertiesCreator("https://ng-api-grpc.coralogix.com", "my-secret-token")
	c := cxsdk.NewDashboardsClient(creator)
	dat, err := os.ReadFile("dashboard.json")
	assert.Nil(t, err)
	var d cxsdk.Dashboard
	assert.Nil(t, protojson.Unmarshal(dat, &d))
	_, e := c.CreateDashboard(context.Background(), &cxsdk.CreateDashboardRequest{
		Dashboard: &d,
	})
	assert.Nil(t, e)
	_, e = c.PinDashboard(context.Background(), &cxsdk.PinDashboardRequest{
		DashboardId: d.Id,
	})
	assert.Nil(t, e)
	_, e = c.UnpinDashboard(context.Background(), &cxsdk.UnpinDashboardRequest{
		DashboardId: d.Id,
	})
	assert.Nil(t, e)
	_, e = c.DeleteDashboard(context.Background(), &cxsdk.DeleteDashboardRequest{
		DashboardId: d.Id,
	})
	assert.Nil(t, e)
}
