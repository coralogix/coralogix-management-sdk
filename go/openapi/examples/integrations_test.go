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
	"testing"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	integrations "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/integration_service"
)

func TestIntegration(t *testing.T) {
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)

	client := cxsdk.NewIntegrationsClient(cpc)
	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		t.Fatalf("env AWS_REGION not set")
	}
	role := os.Getenv("AWS_TEST_ROLE")
	if role == "" {
		t.Fatalf("env AWS_TEST_ROLE not set")
	}

	name := "aws-metrics-collector"
	version := "0.1.0"

	params := []integrations.Parameter{
		integrations.ParameterStringValueAsParameter(&integrations.ParameterStringValue{
			Key:         strPtr("ApplicationName"),
			StringValue: strPtr("cxsdk"),
		}),
		integrations.ParameterStringValueAsParameter(&integrations.ParameterStringValue{
			Key:         strPtr("SubsystemName"),
			StringValue: strPtr("aws-metrics-collector"),
		}),
		integrations.ParameterStringValueAsParameter(&integrations.ParameterStringValue{
			Key:         strPtr("AwsRoleArn"),
			StringValue: strPtr(role),
		}),
		integrations.ParameterStringValueAsParameter(&integrations.ParameterStringValue{
			Key:         strPtr("IntegrationName"),
			StringValue: strPtr("sdk-integration-setup"),
		}),
		integrations.ParameterStringValueAsParameter(&integrations.ParameterStringValue{
			Key:         strPtr("AwsRegion"),
			StringValue: strPtr(awsRegion),
		}),
		integrations.ParameterBooleanValueAsParameter(&integrations.ParameterBooleanValue{
			Key:          strPtr("WithAggregations"),
			BooleanValue: boolPtr(false),
		}),
		integrations.ParameterBooleanValueAsParameter(&integrations.ParameterBooleanValue{
			Key:          strPtr("EnrichWithTags"),
			BooleanValue: boolPtr(false),
		}),
		integrations.ParameterStringListAsParameter(&integrations.ParameterStringList{
			Key: strPtr("MetricNamespaces"),
			StringList: &integrations.StringList{
				Values: []string{"AWS/S3"},
			},
		}),
	}

	testReq := integrations.TestIntegrationRequest{
		IntegrationData: integrations.IntegrationMetadata{
			IntegrationKey: strPtr(name),
			Version:        strPtr(version),
			IntegrationParameters: &integrations.GenericIntegrationParameters{
				Parameters: params,
			},
		},
	}

	_, httpResp, err := client.IntegrationServiceTestIntegration(context.Background()).
		TestIntegrationRequest(testReq).Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	saveReq := integrations.NewSaveIntegrationRequest(
		integrations.IntegrationMetadata{
			IntegrationKey: strPtr(name),
			Version:        strPtr(version),
			IntegrationParameters: &integrations.GenericIntegrationParameters{
				Parameters: params,
			},
		},
	)

	createResp, httpResp, err := client.IntegrationServiceSaveIntegration(context.Background()).
		SaveIntegrationRequest(*saveReq).Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	_, httpResp, err = client.
		IntegrationServiceDeleteIntegration(context.Background(), createResp.IntegrationId).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
}

func strPtr(s string) *string { return &s }

func boolPtr(b bool) *bool { return &b }
