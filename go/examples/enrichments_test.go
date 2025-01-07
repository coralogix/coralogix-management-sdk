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

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestEnrichments(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	client := cxsdk.NewEnrichmentClient(creator)

	withAsn := true

	enrichments, err := client.Add(context.Background(), &cxsdk.AddEnrichmentsRequest{
		RequestEnrichments: []*cxsdk.EnrichmentRequestModel{
			{
				FieldName: &wrapperspb.StringValue{Value: "coralogix.metadata.sdkId"},
				EnrichmentType: &cxsdk.EnrichmentType{
					Type: &cxsdk.EnrichmentTypeGeoIP{
						GeoIp: &cxsdk.GeoIPType{
							WithAsn: &withAsn,
						},
					},
				},
			},
		},
	})
	assertNilAndPrintError(t, err)
	assert.NotNil(t, enrichments)

	enrichmentIDs := []uint32{}
	for _, e := range enrichments.Enrichments {
		enrichmentIDs = append(enrichmentIDs, e.Id)
	}

	// List enrichments and verify creation
	listEnrichmentsResponse, err := client.List(context.Background(), &cxsdk.GetEnrichmentsRequest{})
	assertNilAndPrintError(t, err)
	assert.NotEmpty(t, listEnrichmentsResponse.Enrichments)

	// Verify our created enrichments exist in the list
	for _, e := range listEnrichmentsResponse.Enrichments {
		assert.Contains(t, enrichmentIDs, e.Id)
	}

	// Delete enrichments
	wrappedEnrichmentIDs := []*wrapperspb.UInt32Value{}
	for _, id := range enrichmentIDs {
		wrappedEnrichmentIDs = append(wrappedEnrichmentIDs, &wrapperspb.UInt32Value{Value: id})
	}
	err = client.Delete(context.Background(), &cxsdk.DeleteEnrichmentsRequest{
		EnrichmentIds: wrappedEnrichmentIDs,
	})
	assertNilAndPrintError(t, err)

	// Verify deletion
	listEnrichmentsResponseAfterDeletion, err := client.List(context.Background(), &cxsdk.GetEnrichmentsRequest{})
	assertNilAndPrintError(t, err)

	// Verify our enrichments no longer exist
	for _, e := range listEnrichmentsResponseAfterDeletion.Enrichments {
		assert.NotContains(t, enrichmentIDs, e.Id)
	}
}
