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
	"testing"

	v1 "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm/services/v1"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestSlos(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewSLOsClient(creator)

	createSloResponse, err := c.Create(context.Background(), &cxsdk.CreateServiceSloRequest{
		Slo: &cxsdk.ServiceSlo{
			Name:             wrapperspb.String("coralogix_slo_example"),
			ServiceName:      wrapperspb.String("service_name"),
			Description:      wrapperspb.String("description"),
			TargetPercentage: &wrapperspb.UInt32Value{Value: 30},
			SliType:          &cxsdk.ServiceSloErrorSli{},
			Period:           cxsdk.SloPeriod7Days,
		},
	})

	assert.Nil(t, err)

	_, retrievalError := c.Get(context.Background(), &cxsdk.GetServiceSloRequest{
		Id: createSloResponse.Slo.Id,
	})

	assert.Nil(t, retrievalError)

	updateSloResponse, updateError := c.Update(context.Background(), &cxsdk.ReplaceServiceSloRequest{
		Slo: &cxsdk.ServiceSlo{
			Id:               createSloResponse.Slo.Id,
			Name:             &wrapperspb.StringValue{Value: "coralogix_slo_updated_example"},
			ServiceName:      &wrapperspb.StringValue{Value: "service_name"},
			Description:      &wrapperspb.StringValue{Value: "description"},
			TargetPercentage: &wrapperspb.UInt32Value{Value: 30},
			SliType:          &cxsdk.ServiceSloErrorSli{},
			Period:           v1.SloPeriod_SLO_PERIOD_7_DAYS,
		},
	})

	assert.Nil(t, updateError)

	assert.Equal(t, createSloResponse.Slo.Id.Value, updateSloResponse.Slo.Id.Value)

	_, deletionError := c.Delete(context.Background(), &cxsdk.DeleteServiceSloRequest{
		Id: createSloResponse.Slo.Id,
	})

	assert.Nil(t, deletionError)
}

func TestSlosWithFilters(t *testing.T) {

	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewSLOsClient(creator)

	createSloResponse, err := c.Create(context.Background(), &cxsdk.CreateServiceSloRequest{
		Slo: &cxsdk.ServiceSlo{
			Name:             &wrapperspb.StringValue{Value: "coralogix_slo_example"},
			ServiceName:      &wrapperspb.StringValue{Value: "service_name"},
			Description:      &wrapperspb.StringValue{Value: "description"},
			TargetPercentage: &wrapperspb.UInt32Value{Value: 30},
			SliType:          &cxsdk.ServiceSloErrorSli{},
			Period:           v1.SloPeriod_SLO_PERIOD_7_DAYS,
			Filters: []*v1.SliFilter{
				{Field: &wrapperspb.StringValue{Value: "severity"}, CompareType: v1.CompareType_COMPARE_TYPE_IS, FieldValues: []*wrapperspb.StringValue{{Value: "ERROR"}, {Value: "Warning"}}},
			},
		},
	})

	assert.Nil(t, err)

	_, retrievalError := c.Get(context.Background(), &cxsdk.GetServiceSloRequest{
		Id: createSloResponse.Slo.Id,
	})

	assert.Nil(t, retrievalError)

	updateSloResponse, updateError := c.Update(context.Background(), &cxsdk.ReplaceServiceSloRequest{
		Slo: &cxsdk.ServiceSlo{
			Id:               createSloResponse.Slo.Id,
			Name:             &wrapperspb.StringValue{Value: "coralogix_slo_updated_example"},
			ServiceName:      &wrapperspb.StringValue{Value: "service_name"},
			Description:      &wrapperspb.StringValue{Value: "description"},
			TargetPercentage: &wrapperspb.UInt32Value{Value: 30},
			SliType:          &cxsdk.ServiceSloErrorSli{},
			Period:           v1.SloPeriod_SLO_PERIOD_7_DAYS,
		},
	})

	assert.Nil(t, updateError)

	assert.Equal(t, createSloResponse.Slo.Id.Value, updateSloResponse.Slo.Id.Value)

	_, deletionError := c.Delete(context.Background(), &cxsdk.DeleteServiceSloRequest{
		Id: createSloResponse.Slo.Id,
	})

	assert.Nil(t, deletionError)
}
