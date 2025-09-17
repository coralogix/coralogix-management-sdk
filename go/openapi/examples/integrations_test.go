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
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

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

	params := []integrations.DeployedIntegrationInformationParametersInner{
		integrations.ParameterAsDeployedIntegrationInformationParametersInner(&integrations.Parameter{
			Key:         strPtr("ApplicationName"),
			StringValue: strPtr("cxsdk"),
		}),
		integrations.ParameterAsDeployedIntegrationInformationParametersInner(&integrations.Parameter{
			Key:         strPtr("SubsystemName"),
			StringValue: strPtr("aws-metrics-collector"),
		}),
		integrations.ParameterAsDeployedIntegrationInformationParametersInner(&integrations.Parameter{
			Key:         strPtr("AwsRoleArn"),
			StringValue: strPtr(role),
		}),
		integrations.ParameterAsDeployedIntegrationInformationParametersInner(&integrations.Parameter{
			Key:         strPtr("IntegrationName"),
			StringValue: strPtr("sdk-integration-setup"),
		}),
		integrations.ParameterAsDeployedIntegrationInformationParametersInner(&integrations.Parameter{
			Key:         strPtr("AwsRegion"),
			StringValue: strPtr(awsRegion),
		}),
		integrations.Parameter1AsDeployedIntegrationInformationParametersInner(&integrations.Parameter1{
			Key:          strPtr("WithAggregations"),
			BooleanValue: boolPtr(false),
		}),
		integrations.Parameter1AsDeployedIntegrationInformationParametersInner(&integrations.Parameter1{
			Key:          strPtr("EnrichWithTags"),
			BooleanValue: boolPtr(false),
		}),
		integrations.Parameter2AsDeployedIntegrationInformationParametersInner(&integrations.Parameter2{
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

	_, _, err := client.IntegrationServiceTestIntegration(context.Background()).
		TestIntegrationRequest(testReq).Execute()
	assert.Nil(t, err)
	if err != nil {
		log.Fatal(err)
	}

	saveReq := integrations.NewSaveIntegrationRequest(
		integrations.IntegrationMetadata{
			IntegrationKey: strPtr(name),
			Version:        strPtr(version),
			IntegrationParameters: &integrations.GenericIntegrationParameters{
				Parameters: params,
			},
		},
	)

	createResp, _, err := client.IntegrationServiceSaveIntegration(context.Background()).
		SaveIntegrationRequest(*saveReq).Execute()
	assert.Nil(t, err)

	_, _, err = client.
		IntegrationServiceDeleteIntegration(context.Background(), createResp.IntegrationId).
		Execute()
	assert.Nil(t, err)
}

func strPtr(s string) *string { return &s }

func boolPtr(b bool) *bool { return &b }
