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
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewEnrichmentsClient(cpc)

	enrichmentType := enrichments.EnrichmentType{
		EnrichmentTypeGeoIp: &enrichments.EnrichmentTypeGeoIp{
			GeoIp: &enrichments.GeoIpType{
				WithAsn: enrichments.PtrBool(true),
			},
		},
	}

	model := enrichments.NewEnrichmentRequestModel(
		enrichmentType,
		"coralogix.metadata.sdkId",
	)

	req := *enrichments.NewEnrichmentsCreationRequest([]enrichments.EnrichmentRequestModel{
		*model,
	})

	addResp, httpResp, err := client.
		EnrichmentServiceAddEnrichments(context.Background()).
		EnrichmentsCreationRequest(req).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotEmpty(t, addResp.GetEnrichments())

	ids := make([]int64, 0, len(addResp.GetEnrichments()))
	for _, e := range addResp.GetEnrichments() {
		ids = append(ids, e.GetId())
	}

	_, httpResp, err = client.
		EnrichmentServiceRemoveEnrichments(context.Background()).
		EnrichmentIds(ids).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestEnrichmentsAws(t *testing.T) {
	t.Skip("Skipping AWS")
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewEnrichmentsClient(cpc)

	enrichmentType := enrichments.EnrichmentType{
		EnrichmentTypeAws: &enrichments.EnrichmentTypeAws{
			Aws: &enrichments.AwsType{
				ResourceType: enrichments.PtrString("ec2"),
			},
		},
	}
	model := enrichments.NewEnrichmentRequestModel(enrichmentType, "coralogix.metadata.sdkId")
	req := *enrichments.NewEnrichmentsCreationRequest([]enrichments.EnrichmentRequestModel{*model})

	addResp, httpResp, err := client.
		EnrichmentServiceAddEnrichments(context.Background()).
		EnrichmentsCreationRequest(req).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotEmpty(t, addResp.GetEnrichments())

	ids := make([]int64, 0, len(addResp.GetEnrichments()))
	for _, e := range addResp.GetEnrichments() {
		ids = append(ids, e.GetId())
	}

	_, httpResp, err = client.
		EnrichmentServiceRemoveEnrichments(context.Background()).
		EnrichmentIds(ids).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestEnrichmentsCustom(t *testing.T) {
	t.Skip("Skipping Custom")
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewEnrichmentsClient(cpc)

	enrichmentType := enrichments.EnrichmentType{
		EnrichmentTypeCustomEnrichment: &enrichments.EnrichmentTypeCustomEnrichment{
			CustomEnrichment: &enrichments.CustomEnrichmentType{},
		},
	}
	model := enrichments.NewEnrichmentRequestModel(enrichmentType, "coralogix.metadata.sdkId")
	req := *enrichments.NewEnrichmentsCreationRequest([]enrichments.EnrichmentRequestModel{*model})

	addResp, httpResp, err := client.
		EnrichmentServiceAddEnrichments(context.Background()).
		EnrichmentsCreationRequest(req).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotEmpty(t, addResp.GetEnrichments())

	ids := make([]int64, 0, len(addResp.GetEnrichments()))
	for _, e := range addResp.GetEnrichments() {
		ids = append(ids, e.GetId())
	}

	_, httpResp, err = client.
		EnrichmentServiceRemoveEnrichments(context.Background()).
		EnrichmentIds(ids).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestEnrichmentsSuspiciousIp(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewEnrichmentsClient(cpc)
	enrichmentType := enrichments.EnrichmentType{
		EnrichmentTypeSuspiciousIp: &enrichments.EnrichmentTypeSuspiciousIp{
			SuspiciousIp: map[string]interface{}{},
		},
	}
	model := enrichments.NewEnrichmentRequestModel(enrichmentType, "coralogix.metadata.sdkId")
	req := *enrichments.NewEnrichmentsCreationRequest([]enrichments.EnrichmentRequestModel{*model})

	addResp, httpResp, err := client.
		EnrichmentServiceAddEnrichments(context.Background()).
		EnrichmentsCreationRequest(req).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotEmpty(t, addResp.GetEnrichments())

	ids := make([]int64, 0, len(addResp.GetEnrichments()))
	for _, e := range addResp.GetEnrichments() {
		ids = append(ids, e.GetId())
	}

	_, httpResp, err = client.
		EnrichmentServiceRemoveEnrichments(context.Background()).
		EnrichmentIds(ids).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}
