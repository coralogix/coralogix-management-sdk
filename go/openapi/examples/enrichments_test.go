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
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	"github.com/coralogix/coralogix-management-sdk/go/openapi/gen/custom_enrichments_service"
	enrichments "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/enrichments_service"
)

func TestEnrichmentsGeoIp(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewEnrichmentsClient(cfg)

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
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotEmpty(t, addResp.GetEnrichments())

	ids := make([]int64, 0, len(addResp.GetEnrichments()))
	for _, e := range addResp.GetEnrichments() {
		ids = append(ids, e.GetId())
	}

	_, httpResp, err = client.
		EnrichmentServiceRemoveEnrichments(context.Background()).
		EnrichmentIds(ids).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestEnrichmentsAws(t *testing.T) {
	t.Skip("Skipping AWS")
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewEnrichmentsClient(cfg)

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
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotEmpty(t, addResp.GetEnrichments())

	ids := make([]int64, 0, len(addResp.GetEnrichments()))
	for _, e := range addResp.GetEnrichments() {
		ids = append(ids, e.GetId())
	}

	_, httpResp, err = client.
		EnrichmentServiceRemoveEnrichments(context.Background()).
		EnrichmentIds(ids).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestEnrichmentsCustom(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	customEClient := cxsdk.NewCustomEnrichmentsClient(cfg)
	name := fmt.Sprintf("test-%v", uuid.NewString())

	ext := "csv"
	contents := "local_id,instance_type\nfoo1,t2.micro\nfoo2,t2.micro\nfoo3,t2.micro\nbar1,m3.large\n"

	data, httpResp, err := customEClient.
		CustomEnrichmentServiceCreateCustomEnrichment(context.Background()).
		Name(name).
		Description(name).
		File(custom_enrichments_service.
			CustomEnrichmentServiceCreateCustomEnrichmentFileParameter{
			FileTextual: &custom_enrichments_service.FileTextual{
				Extension: &ext,
				Name:      &name,
				Textual:   &contents,
			},
		}).
		Execute()

	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	client := cxsdk.NewEnrichmentsClient(cfg)

	enrichmentType := enrichments.EnrichmentType{
		EnrichmentTypeCustomEnrichment: &enrichments.EnrichmentTypeCustomEnrichment{
			CustomEnrichment: &enrichments.CustomEnrichmentType{
				Id: data.CustomEnrichment.Id,
			},
		},
	}
	model := enrichments.NewEnrichmentRequestModel(enrichmentType, "coralogix.metadata.sdkId")
	req := *enrichments.NewEnrichmentsCreationRequest([]enrichments.EnrichmentRequestModel{*model})

	addResp, httpResp, err := client.
		EnrichmentServiceAddEnrichments(context.Background()).
		EnrichmentsCreationRequest(req).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotEmpty(t, addResp.GetEnrichments())

	ids := make([]int64, 0, len(addResp.GetEnrichments()))
	for _, e := range addResp.GetEnrichments() {
		ids = append(ids, e.GetId())
	}

	_, httpResp, err = client.
		EnrichmentServiceRemoveEnrichments(context.Background()).
		EnrichmentIds(ids).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = customEClient.
		CustomEnrichmentServiceDeleteCustomEnrichment(context.Background(), *data.CustomEnrichment.Id).
		Execute()

	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func TestEnrichmentsSuspiciousIp(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	client := cxsdk.NewEnrichmentsClient(cfg)

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
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotEmpty(t, addResp.GetEnrichments())

	ids := make([]int64, 0, len(addResp.GetEnrichments()))
	for _, e := range addResp.GetEnrichments() {
		ids = append(ids, e.GetId())
	}

	_, httpResp, err = client.
		EnrichmentServiceRemoveEnrichments(context.Background()).
		EnrichmentIds(ids).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}
