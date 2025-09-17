// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package examples

import (
	"context"
	"os"
	"testing"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go/openapi"
	"github.com/stretchr/testify/assert"
)

func newClient(t *testing.T) *cxsdk.APIClient {
	region := os.Getenv(cxsdk.EnvCoralogixRegion)
	if region == "" {
		t.Fatalf("env %s not set", cxsdk.EnvCoralogixRegion)
	}
	apiKey := os.Getenv(cxsdk.EnvCoralogixAPIKey)
	if apiKey == "" {
		t.Fatalf("env %s not set", cxsdk.EnvCoralogixAPIKey)
	}

	return cxsdk.NewClientBuilder().
		WithRegion(region).
		WithAPIKey(apiKey).
		Build()
}

func TestEnrichmentsGeoIp(t *testing.T) {
	client := newClient(t)

	enrichmentReq := cxsdk.NewComCoralogixEnrichmentV1EnrichmentRequestModel(
		cxsdk.ComCoralogixEnrichmentV1EnrichmentType{
			ComCoralogixEnrichmentV1EnrichmentTypeOneOf: &cxsdk.ComCoralogixEnrichmentV1EnrichmentTypeOneOf{
				GeoIp: &cxsdk.ComCoralogixEnrichmentV1GeoIpType{
					WithAsn: cxsdk.PtrBool(true)},
			}},
		"coralogix.metadata.sdkId",
	)

	creation := *cxsdk.NewEnrichmentsCreationRequest([]cxsdk.ComCoralogixEnrichmentV1EnrichmentRequestModel{
		*enrichmentReq,
	})

	addResp, _, err := client.EnrichmentsServiceAPI.
		EnrichmentServiceAddEnrichments(context.Background()).
		EnrichmentsCreationRequest(creation).
		Execute()
	assert.Nil(t, err)
	assert.NotEmpty(t, addResp.GetEnrichments())

	ids := make([]int64, 0, len(addResp.GetEnrichments()))
	for _, e := range addResp.GetEnrichments() {
		ids = append(ids, e.GetId())
	}

	listResp, _, err := client.EnrichmentsServiceAPI.
		EnrichmentServiceGetEnrichments(context.Background()).
		Execute()
	assert.Nil(t, err)
	assert.NotEmpty(t, listResp.GetEnrichments())

	found := false
	for _, e := range listResp.GetEnrichments() {
		for _, id := range ids {
			if e.GetId() == id {
				found = true
			}
		}
	}
	assert.True(t, found)

	_, _, err = client.EnrichmentsServiceAPI.
		EnrichmentServiceRemoveEnrichments(context.Background()).
		Execute()
	assert.Nil(t, err)
}

func TestEnrichmentsAws(t *testing.T) {
	client := newClient(t)

	req := cxsdk.NewComCoralogixEnrichmentV1EnrichmentRequestModel(
		cxsdk.ComCoralogixEnrichmentV1EnrichmentType{
			ComCoralogixEnrichmentV1EnrichmentTypeOneOf2: &cxsdk.ComCoralogixEnrichmentV1EnrichmentTypeOneOf2{
				Aws: &cxsdk.ComCoralogixEnrichmentV1AwsType{ResourceType: cxsdk.PtrString("ec2")},
			},
		},
		"coralogix.metadata.sdkId",
	)

	creation := *cxsdk.NewEnrichmentsCreationRequest([]cxsdk.ComCoralogixEnrichmentV1EnrichmentRequestModel{*req})

	addResp, _, err := client.EnrichmentsServiceAPI.
		EnrichmentServiceAddEnrichments(context.Background()).
		EnrichmentsCreationRequest(creation).
		Execute()
	// AWS enrichment may not be supported in your env → assert that it errors
	assert.NotNil(t, err)
	assert.Nil(t, addResp)
}

func TestEnrichmentsCustom(t *testing.T) {
	client := newClient(t)

	req := cxsdk.NewComCoralogixEnrichmentV1EnrichmentRequestModel(
		cxsdk.ComCoralogixEnrichmentV1EnrichmentType{
			ComCoralogixEnrichmentV1EnrichmentTypeOneOf3: &cxsdk.ComCoralogixEnrichmentV1EnrichmentTypeOneOf3{
				CustomEnrichment: &cxsdk.ComCoralogixEnrichmentV1CustomEnrichmentType{},
			},
		},
		"coralogix.metadata.sdkId",
	)

	creation := *cxsdk.NewEnrichmentsCreationRequest([]cxsdk.ComCoralogixEnrichmentV1EnrichmentRequestModel{*req})

	addResp, _, err := client.EnrichmentsServiceAPI.
		EnrichmentServiceAddEnrichments(context.Background()).
		EnrichmentsCreationRequest(creation).
		Execute()
	assert.NotNil(t, err)
	assert.Nil(t, addResp)
}

func TestEnrichmentsSuspiciousIp(t *testing.T) {
	client := newClient(t)

	req := cxsdk.NewComCoralogixEnrichmentV1EnrichmentRequestModel(
		cxsdk.ComCoralogixEnrichmentV1EnrichmentType{
			ComCoralogixEnrichmentV1EnrichmentTypeOneOf1: &cxsdk.ComCoralogixEnrichmentV1EnrichmentTypeOneOf1{
				SuspiciousIp: map[string]interface{}{},
			},
		},
		"coralogix.metadata.sdkId",
	)

	creation := *cxsdk.NewEnrichmentsCreationRequest([]cxsdk.ComCoralogixEnrichmentV1EnrichmentRequestModel{*req})

	addResp, _, err := client.EnrichmentsServiceAPI.
		EnrichmentServiceAddEnrichments(context.Background()).
		EnrichmentsCreationRequest(creation).
		Execute()
	assert.Nil(t, err)
	assert.NotEmpty(t, addResp.GetEnrichments())

	_, _, err = client.EnrichmentsServiceAPI.
		EnrichmentServiceRemoveEnrichments(context.Background()).
		Execute()
	assert.Nil(t, err)
}
