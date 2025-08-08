// Copyright 2025 Coralogix Ltd.
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

func TestExtensions(t *testing.T) {
	t.Skip("Unstable test")
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	client := cxsdk.NewExtensionsClient(creator)

	getAllExtensionsResponse, err := client.GetAll(context.Background(), &cxsdk.GetAllExtensionsRequest{
		IncludeHiddenExtensions: &wrapperspb.BoolValue{Value: true},
	})
	assertNilAndPrintError(t, err)
	extensionsDescriptors := getAllExtensionsResponse.Extensions

	assert.NotEmpty(t, extensionsDescriptors)

	var extensionToDeployId *wrapperspb.StringValue

	for _, extension := range extensionsDescriptors {
		if extension.Name.Value == "Kubernetes Observability" {
			extensionToDeployId = extension.Id
			break
		}
	}

	assert.NotNil(t, extensionToDeployId)

	getDeployedExtensionsResponse, err := client.GetDeployed(context.Background(), &cxsdk.GetDeployedExtensionsRequest{})
	assertNilAndPrintError(t, err)

	for _, deployedExtension := range getDeployedExtensionsResponse.DeployedExtensions {
		if deployedExtension.Id.Value == extensionToDeployId.Value {
			client.Undeploy(context.Background(), &cxsdk.UndeployExtensionRequest{
				Id: deployedExtension.Id,
			})
		}
	}

	getExtensionReponse, err := client.Get(context.Background(), &cxsdk.GetExtensionRequest{
		Id: extensionToDeployId,
	})

	assertNilAndPrintError(t, err)

	extensionToDeploy := getExtensionReponse.Extension

	var itemIds []*wrapperspb.StringValue

	for _, revision := range extensionToDeploy.Revisions {
		if revision.Version.Value == "0.0.2" {
			for _, item := range revision.Items {
				if item.TargetDomain != cxsdk.TargetDomainParsingRule {
					itemIds = append(itemIds, item.Id)
				}
			}
		}
	}

	_, err = client.Deploy(context.Background(), &cxsdk.DeployExtensionRequest{
		Id:      extensionToDeployId,
		Version: &wrapperspb.StringValue{Value: "0.0.2"},
		ItemIds: itemIds,
	})

	assertNilAndPrintError(t, err)

}
