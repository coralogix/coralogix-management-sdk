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

	cxsdk "github.com/coralogix/coralogix-management-sdk/go/openapi"
	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	region := os.Getenv(cxsdk.EnvCoralogixRegion)
	if region == "" {
		t.Fatalf("env %s not set", cxsdk.EnvCoralogixRegion)
	}
	apiKey := os.Getenv(cxsdk.EnvCoralogixAPIKey)
	if apiKey == "" {
		t.Fatalf("env %s not set", cxsdk.EnvCoralogixAPIKey)
	}
	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		t.Fatalf("env AWS_REGION not set")
	}
	role := os.Getenv("AWS_TEST_ROLE")
	if role == "" {
		t.Fatalf("env AWS_TEST_ROLE not set")
	}

	client := cxsdk.NewClientBuilder().
		WithRegion(region).
		WithAPIKey(apiKey).
		Build()

	name := "aws-metrics-collector"
	version := "0.1.0"

	params := []cxsdk.ComCoralogixIntegrationsV1Parameter{
		cxsdk.ParameterAsComCoralogixIntegrationsV1Parameter(&cxsdk.Parameter{
			Key:         strPtr("ApplicationName"),
			StringValue: strPtr("cxsdk"),
		}),
		cxsdk.ParameterAsComCoralogixIntegrationsV1Parameter(&cxsdk.Parameter{
			Key:         strPtr("SubsystemName"),
			StringValue: strPtr("aws-metrics-collector"),
		}),
		cxsdk.ParameterAsComCoralogixIntegrationsV1Parameter(&cxsdk.Parameter{
			Key:         strPtr("AwsRoleArn"),
			StringValue: strPtr(role),
		}),
		cxsdk.ParameterAsComCoralogixIntegrationsV1Parameter(&cxsdk.Parameter{
			Key:         strPtr("IntegrationName"),
			StringValue: strPtr("sdk-integration-setup"),
		}),
		cxsdk.ParameterAsComCoralogixIntegrationsV1Parameter(&cxsdk.Parameter{
			Key:         strPtr("AwsRegion"),
			StringValue: strPtr(awsRegion),
		}),
		cxsdk.Parameter1AsComCoralogixIntegrationsV1Parameter(&cxsdk.Parameter1{
			Key:          strPtr("WithAggregations"),
			BooleanValue: boolPtr(false),
		}),
		cxsdk.Parameter1AsComCoralogixIntegrationsV1Parameter(&cxsdk.Parameter1{
			Key:          strPtr("EnrichWithTags"),
			BooleanValue: boolPtr(false),
		}),
		cxsdk.Parameter2AsComCoralogixIntegrationsV1Parameter(&cxsdk.Parameter2{
			Key: strPtr("MetricNamespaces"),
			StringList: &cxsdk.ComCoralogixIntegrationsV1ParameterStringList{
				Values: []string{"AWS/S3"},
			},
		}),
	}

	testReq := cxsdk.TestIntegrationRequest{
		IntegrationData: cxsdk.ComCoralogixIntegrationsV1IntegrationMetadata{
			IntegrationKey: strPtr(name),
			Version:        strPtr(version),
			IntegrationParameters: &cxsdk.ComCoralogixIntegrationsV1GenericIntegrationParameters{
				Parameters: params,
			},
		},
	}

	_, _, err := client.IntegrationServiceAPI.IntegrationServiceTestIntegration(context.Background()).
		TestIntegrationRequest(testReq).Execute()
	assert.Nil(t, err)
	if err != nil {
		log.Fatal(err)
	}

	saveReq := cxsdk.NewSaveIntegrationRequest(
		cxsdk.ComCoralogixIntegrationsV1IntegrationMetadata{
			IntegrationKey: strPtr(name),
			Version:        strPtr(version),
			IntegrationParameters: &cxsdk.ComCoralogixIntegrationsV1GenericIntegrationParameters{
				Parameters: params,
			},
		},
	)

	createResp, _, err := client.IntegrationServiceAPI.IntegrationServiceSaveIntegration(context.Background()).
		SaveIntegrationRequest(*saveReq).Execute()
	assert.Nil(t, err)

	_, _, err = client.IntegrationServiceAPI.
		IntegrationServiceDeleteIntegration(context.Background(), createResp.IntegrationId).
		Execute()
	assert.Nil(t, err)
}

func strPtr(s string) *string { return &s }

func boolPtr(b bool) *bool { return &b }
