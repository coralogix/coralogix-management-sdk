package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	extdeploy "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/extension_deployment_service"
	extensions "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/extension_service"
)

func TestExtensions(t *testing.T) {
	t.Skip("Unstable test")
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().Build()
	clientSet := cxsdk.NewClientSet(cfg)
	extensionsClient := clientSet.Extensions()
	deployClient := clientSet.ExtensionDeployments()

	ctx := context.Background()

	includeHidden := true
	getAllReq := extensions.GetAllExtensionsRequest{
		IncludeHiddenExtensions: &includeHidden,
	}

	allResp, httpResp, err := extensionsClient.
		ExtensionServiceGetAllExtensions(ctx).
		GetAllExtensionsRequest(getAllReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotEmpty(t, allResp.Extensions)

	var extensionID string
	for _, e := range allResp.Extensions {
		if e.GetName() == "Kubernetes Observability" {
			extensionID = e.GetId()
			break
		}
	}
	require.NotEmpty(t, extensionID, "expected 'Kubernetes Observability' extension to exist")

	deployedResp, httpResp, err := deployClient.
		ExtensionDeploymentServiceGetDeployedExtensions(ctx).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	for _, deployed := range deployedResp.DeployedExtensions {
		if deployed.GetId() == extensionID {
			undeployReq := extdeploy.RevertDeploymentOfExtensionRequest{
				Id: extensions.PtrString(deployed.GetId()),
			}
			_, httpResp, err := deployClient.
				ExtensionDeploymentServiceUndeployExtension(ctx).
				RevertDeploymentOfExtensionRequest(undeployReq).
				Execute()
			require.NoError(t, cxsdk.NewAPIError(httpResp, err))
			break
		}
	}

	extDetails, httpResp, err := extensionsClient.
		ExtensionServiceGetExtension(ctx, extensionID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, extensionID, extDetails.GetId())

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
	require.NotEmpty(t, itemIDs, "expected items for version 0.0.2")

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
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, deployResp, "deployment response should not be nil")

	fmt.Printf("Deployed extension %s version %s successfully\n", extensionID, *deployReq.Version)
}
