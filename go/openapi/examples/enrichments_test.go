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
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	enrichments "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/enrichments_service"
)

func TestEnrichmentsGeoIp(t *testing.T) {
	t.Skip("Skipping until we can delete by IDs")
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewEnrichmentsClient(cpc)

	enrichmentType := enrichments.EnrichmentEnrichmentType{
		EnrichmentEnrichmentTypeOneOf: &enrichments.EnrichmentEnrichmentTypeOneOf{
			GeoIp: &enrichments.EnrichmentEnrichmentTypeOneOfGeoIp{
				WithAsn: enrichments.PtrBool(true),
			},
		},
	}

	// Create enrichment prototype
	model := enrichments.NewEnrichmentPrototype(
		enrichmentType,
		"coralogix.metadata.sdkId",
	)

	req := *enrichments.NewEnrichmentsCreationRequest([]enrichments.EnrichmentPrototype{
		*model,
	})

	addResp, _, err := client.
		EnrichmentServiceAddEnrichments(context.Background()).
		EnrichmentsCreationRequest(req).
		Execute()
	assertNilAndPrintError(t, err)
	assert.NotEmpty(t, addResp.GetEnrichments())

	ids := make([]int64, 0, len(addResp.GetEnrichments()))
	for _, e := range addResp.GetEnrichments() {
		ids = append(ids, e.GetId())
	}

	listResp, _, err := client.EnrichmentServiceGetEnrichments(context.Background()).Execute()
	assertNilAndPrintError(t, err)
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

	// TODO: Use query params to set IDs
	//_, _, err = client.EnrichmentsServiceAPI.
	//	EnrichmentServiceRemoveEnrichments(context.Background()).
	//	Execute()
	//assertNilAndPrintError(t, err)
}

func TestEnrichmentsAws(t *testing.T) {
	t.Skip("Skipping until we can delete by IDs")
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewEnrichmentsClient(cpc)

	enrichmentType := enrichments.EnrichmentEnrichmentType{
		EnrichmentEnrichmentTypeOneOf2: &enrichments.EnrichmentEnrichmentTypeOneOf2{
			Aws: &enrichments.EnrichmentEnrichmentTypeOneOf2Aws{
				ResourceType: enrichments.PtrString("ec2"),
			},
		},
	}
	model := enrichments.NewEnrichmentPrototype(enrichmentType, "coralogix.metadata.sdkId")
	req := *enrichments.NewEnrichmentsCreationRequest([]enrichments.EnrichmentPrototype{*model})

	addResp, _, err := client.
		EnrichmentServiceAddEnrichments(context.Background()).
		EnrichmentsCreationRequest(req).
		Execute()
	assertNilAndPrintError(t, err)
	assert.NotEmpty(t, addResp.GetEnrichments())

	// TODO: Use query params to set IDs
	//_, _, err = client.EnrichmentsServiceAPI.
	//	EnrichmentServiceRemoveEnrichments(context.Background()).
	//	Execute()
	//assertNilAndPrintError(t, err)
}

func TestEnrichmentsCustom(t *testing.T) {
	t.Skip("Skipping until we can delete by IDs")
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewEnrichmentsClient(cpc)

	enrichmentType := enrichments.EnrichmentEnrichmentType{
		EnrichmentEnrichmentTypeOneOf3: &enrichments.EnrichmentEnrichmentTypeOneOf3{
			CustomEnrichment: &enrichments.EnrichmentEnrichmentTypeOneOf3CustomEnrichment{},
		},
	}
	model := enrichments.NewEnrichmentPrototype(enrichmentType, "coralogix.metadata.sdkId")
	req := *enrichments.NewEnrichmentsCreationRequest([]enrichments.EnrichmentPrototype{*model})

	addResp, _, err := client.
		EnrichmentServiceAddEnrichments(context.Background()).
		EnrichmentsCreationRequest(req).
		Execute()
	assertNilAndPrintError(t, err)
	assert.NotEmpty(t, addResp.GetEnrichments())

	// TODO: Use query params to set IDs
	//_, _, err = client.EnrichmentsServiceAPI.
	//	EnrichmentServiceRemoveEnrichments(context.Background()).
	//	Execute()
	//assertNilAndPrintError(t, err)
}

func TestEnrichmentsSuspiciousIp(t *testing.T) {
	t.Skip("Skipping until we can delete by IDs")

	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewEnrichmentsClient(cpc)
	enrichmentType := enrichments.EnrichmentEnrichmentType{
		EnrichmentEnrichmentTypeOneOf1: &enrichments.EnrichmentEnrichmentTypeOneOf1{
			SuspiciousIp: map[string]interface{}{},
		},
	}
	model := enrichments.NewEnrichmentPrototype(enrichmentType, "coralogix.metadata.sdkId")
	req := *enrichments.NewEnrichmentsCreationRequest([]enrichments.EnrichmentPrototype{*model})

	addResp, _, err := client.
		EnrichmentServiceAddEnrichments(context.Background()).
		EnrichmentsCreationRequest(req).
		Execute()
	assertNilAndPrintError(t, err)
	assert.NotEmpty(t, addResp.GetEnrichments())

	// TODO: Use query params to set IDs
	//_, _, err = client.EnrichmentsServiceAPI.
	//	EnrichmentServiceRemoveEnrichments(context.Background()).
	//	Execute()
	//assertNilAndPrintError(t, err)
}
