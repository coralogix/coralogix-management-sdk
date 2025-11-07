package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	extdeploy "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/extension_deployment_service"
	extensions "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/extension_service"
)

func TestExtensions(t *testing.T) {
	t.Skip("Unstable test")
	cpc := cxsdk.NewSDKCallPropertiesCreator(
		cxsdk.URLFromRegion(cxsdk.RegionFromEnv()),
		cxsdk.APIKeyFromEnv(),
	)
	ctx := context.Background()

	extensionsClient := cxsdk.NewExtensionsClient(cpc)
	deployClient := cxsdk.NewExtensionDeploymentsClient(cpc)

	includeHidden := true
	getAllReq := extensions.GetAllExtensionsRequest{
		IncludeHiddenExtensions: &includeHidden,
	}

	allResp, httpResp, err := extensionsClient.
		ExtensionServiceGetAllExtensions(ctx).
		GetAllExtensionsRequest(getAllReq).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotEmpty(t, allResp.Extensions)

	var extensionID string
	for _, e := range allResp.Extensions {
		if e.GetName() == "Kubernetes Observability" {
			extensionID = e.GetId()
			break
		}
	}
	assert.NotEmpty(t, extensionID, "expected 'Kubernetes Observability' extension to exist")

	deployedResp, httpResp, err := deployClient.
		ExtensionDeploymentServiceGetDeployedExtensions(ctx).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))

	for _, deployed := range deployedResp.DeployedExtensions {
		if deployed.GetId() == extensionID {
			undeployReq := extdeploy.RevertDeploymentOfExtensionRequest{
				Id: extensions.PtrString(deployed.GetId()),
			}
			_, httpResp, err := deployClient.
				ExtensionDeploymentServiceUndeployExtension(ctx).
				RevertDeploymentOfExtensionRequest(undeployReq).
				Execute()
			assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
			break
		}
	}

	extDetails, httpResp, err := extensionsClient.
		ExtensionServiceGetExtension(ctx, extensionID).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.Equal(t, extensionID, extDetails.GetId())

	var itemIDs []string
	for _, rev := range extDetails.Revisions {
		if rev.GetVersion() == "0.0.2" {
			for _, item := range rev.Items {
				if item.GetTargetDomain() != "PARSING_RULE" {
					itemIDs = append(itemIDs, item.GetId())
				}
			}
			break
		}
	}
	assert.NotEmpty(t, itemIDs, "expected items for version 0.0.2")

	deployReq := extdeploy.ExtensionDeployment{
		Id:      &extensionID,
		Version: extensions.PtrString("0.0.2"),
		ItemIds: itemIDs,
	}

	deployResp, httpResp, err := deployClient.
		ExtensionDeploymentServiceDeployExtension(ctx).
		Id(*deployReq.Id).
		Version(*deployReq.Version).
		ItemIds(deployReq.ItemIds).
		Execute()
	assertNilAndPrintError(t, cxsdk.NewAPIError(httpResp, err))
	assert.NotNil(t, deployResp, "deployment response should not be nil")

	fmt.Printf("Deployed extension %s version %s successfully\n", extensionID, deployReq.Version)
}
