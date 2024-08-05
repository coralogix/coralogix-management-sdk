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
	"os"
	"reflect"
	"testing"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func bp(k string, v bool) *cxsdk.IntegrationParameter {
	return &cxsdk.IntegrationParameter{
		Key: k,
		Value: &cxsdk.IntegrationParameterBooleanValue{
			BooleanValue: &wrapperspb.BoolValue{Value: v},
		},
	}
}

func sp(k string, v string) *cxsdk.IntegrationParameter {
	return &cxsdk.IntegrationParameter{
		Key: k,
		Value: &cxsdk.IntegrationParameterStringValue{
			StringValue: &wrapperspb.StringValue{Value: v},
		},
	}
}

func sl(k string, v []string) *cxsdk.IntegrationParameter {
	wrapped := make([]*wrapperspb.StringValue, len(v))
	for i, s := range v {
		wrapped[i] = &wrapperspb.StringValue{Value: s}
	}
	return &cxsdk.IntegrationParameter{
		Key: k,
		Value: &cxsdk.IntegrationParameterStringList{
			StringList: &cxsdk.IntegrationParameterStringListInner{
				Values: wrapped,
			},
		},
	}
}

func TestIntegration(t *testing.T) {

	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, apiKey)
	c := cxsdk.NewIntegrationsClient(creator)
	role := os.Getenv("AWS_TEST_ROLE")

	name := "aws-metrics-collector"
	version := "0.1.0"
	params := []*cxsdk.IntegrationParameter{
		sp("ApplicationName", "cxsdk"),
		sp("SubsystemName", "aws-metrics-collector"),
		sl("MetricNamespaces", []string{"AWS/S3"}),
		sp("AwsRoleArn", role),
		sp("IntegrationName", "sdk-integration-setup"),
		sp("AwsRegion", "eu-north-1"),
		bp("WithAggregations", false),
		bp("EnrichWithTags", true),
	}
	testingResponse, err := c.Test(context.Background(), &cxsdk.TestIntegrationRequest{
		IntegrationData: &cxsdk.IntegrationMetadata{
			IntegrationKey: &wrapperspb.StringValue{Value: name},
			Version:        &wrapperspb.StringValue{Value: version},
			SpecificData: &cxsdk.IntegrationMetadataIntegrationParameters{
				IntegrationParameters: &cxsdk.GenericIntegrationParameters{
					Parameters: params,
				},
			},
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, reflect.TypeOf(testingResponse.Result.Result), reflect.TypeOf(cxsdk.IntegrationTestSuccess{}))

	createResponse, err := c.Create(context.Background(), &cxsdk.SaveIntegrationRequest{
		Metadata: &cxsdk.IntegrationMetadata{
			IntegrationKey: &wrapperspb.StringValue{Value: name},
			Version:        &wrapperspb.StringValue{Value: version},
			SpecificData: &cxsdk.IntegrationMetadataIntegrationParameters{
				IntegrationParameters: &cxsdk.GenericIntegrationParameters{
					Parameters: params,
				},
			},
		},
	})

	assert.Nil(t, err)

	details, retrievalError := c.Get(context.Background(), &cxsdk.GetIntegrationDetailsRequest{
		Id: createResponse.IntegrationId,
	})
	assert.Nil(t, retrievalError)
	integrationDetail := details.IntegrationDetail.IntegrationTypeDetails.(*cxsdk.IntegrationDetailsDefault)
	found := false
	for _, d := range integrationDetail.Default.Registered {
		if d.Id == createResponse.IntegrationId {
			found = true
			break
		}
	}
	assert.True(t, found)

	_, deletionError := c.Delete(context.Background(), &cxsdk.DeleteIntegrationRequest{
		IntegrationId: createResponse.IntegrationId,
	})

	assert.Nil(t, deletionError)
}
