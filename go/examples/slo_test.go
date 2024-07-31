package examples

import (
	"context"
	"testing"

	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm/services/v1"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestSlos(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, apiKey)
	c := cxsdk.NewSLOsClient(creator)

	createSloResponse, err := c.Create(context.Background(), &v1.CreateServiceSloRequest{
		Slo: &v1.ServiceSlo{
			Name:             &wrapperspb.StringValue{Value: "coralogix_slo_example"},
			ServiceName:      &wrapperspb.StringValue{Value: "service_name"},
			Description:      &wrapperspb.StringValue{Value: "description"},
			TargetPercentage: &wrapperspb.UInt32Value{Value: 30},
			SliType:          &v1.ServiceSlo_ErrorSli{},
			Period:           v1.SloPeriod_SLO_PERIOD_7_DAYS,
		},
	})

	assert.Nil(t, err)

	_, retrievalError := c.Get(context.Background(), &v1.GetServiceSloRequest{
		Id: createSloResponse.Slo.Id,
	})

	assert.Nil(t, retrievalError)

	updateSloResponse, updateError := c.Update(context.Background(), &v1.ReplaceServiceSloRequest{
		Slo: &v1.ServiceSlo{
			Id:               createSloResponse.Slo.Id,
			Name:             &wrapperspb.StringValue{Value: "coralogix_slo_updated_example"},
			ServiceName:      &wrapperspb.StringValue{Value: "service_name"},
			Description:      &wrapperspb.StringValue{Value: "description"},
			TargetPercentage: &wrapperspb.UInt32Value{Value: 30},
			SliType:          &v1.ServiceSlo_ErrorSli{},
			Period:           v1.SloPeriod_SLO_PERIOD_7_DAYS,
		},
	})

	assert.Nil(t, updateError)

	assert.Equal(t, createSloResponse.Slo.Id.Value, updateSloResponse.Slo.Id.Value)

	_, deletionError := c.Delete(context.Background(), &v1.DeleteServiceSloRequest{
		Id: createSloResponse.Slo.Id,
	})

	assert.Nil(t, deletionError)
}

func TestSlosWithFilters(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, apiKey)
	c := cxsdk.NewSLOsClient(creator)

	createSloResponse, err := c.Create(context.Background(), &v1.CreateServiceSloRequest{
		Slo: &v1.ServiceSlo{
			Name:             &wrapperspb.StringValue{Value: "coralogix_slo_example"},
			ServiceName:      &wrapperspb.StringValue{Value: "service_name"},
			Description:      &wrapperspb.StringValue{Value: "description"},
			TargetPercentage: &wrapperspb.UInt32Value{Value: 30},
			SliType:          &v1.ServiceSlo_ErrorSli{},
			Period:           v1.SloPeriod_SLO_PERIOD_7_DAYS,
			Filters: []*v1.SliFilter{
				{Field: &wrapperspb.StringValue{Value: "severity"}, CompareType: v1.CompareType_COMPARE_TYPE_IS, FieldValues: []*wrapperspb.StringValue{{Value: "ERROR"}, {Value: "Warning"}}},
			},
		},
	})

	assert.Nil(t, err)

	_, retrievalError := c.Get(context.Background(), &v1.GetServiceSloRequest{
		Id: createSloResponse.Slo.Id,
	})

	assert.Nil(t, retrievalError)

	updateSloResponse, updateError := c.Update(context.Background(), &v1.ReplaceServiceSloRequest{
		Slo: &v1.ServiceSlo{
			Id:               createSloResponse.Slo.Id,
			Name:             &wrapperspb.StringValue{Value: "coralogix_slo_updated_example"},
			ServiceName:      &wrapperspb.StringValue{Value: "service_name"},
			Description:      &wrapperspb.StringValue{Value: "description"},
			TargetPercentage: &wrapperspb.UInt32Value{Value: 30},
			SliType:          &v1.ServiceSlo_ErrorSli{},
			Period:           v1.SloPeriod_SLO_PERIOD_7_DAYS,
		},
	})

	assert.Nil(t, updateError)

	assert.Equal(t, createSloResponse.Slo.Id.Value, updateSloResponse.Slo.Id.Value)

	_, deletionError := c.Delete(context.Background(), &v1.DeleteServiceSloRequest{
		Id: createSloResponse.Slo.Id,
	})

	assert.Nil(t, deletionError)
}
