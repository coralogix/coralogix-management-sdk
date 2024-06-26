package examples

import (
	"context"
	cxsdk "coralogix-management-sdk/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiKeys(t *testing.T) {
	creator := cxsdk.NewCallPropertiesCreator("https://ng-api-grpc.coralogix.com", "my-secret-token")
	k := cxsdk.NewApiKeysClient(creator)

	key, _ := k.CreateApiKey(context.Background(), &cxsdk.CreateApiKeyRequest{
		Name: "My APM KEY",
		Owner: &cxsdk.Owner{
			Owner: &cxsdk.Owner_UserId{
				UserId: "4013254",
			},
		},
		KeyPermissions: &cxsdk.ApiKeyPermissions{
			Presets:     []string{"APM"},
			Permissions: []string{},
		},
	})

	newName := "new-name"
	k.UpdateApiKey(context.Background(), &cxsdk.UpdateApiKeyRequest{
		KeyId:   key.KeyId,
		NewName: &newName,
	})

	updated, _ := k.GetApiKey(context.Background(), &cxsdk.GetApiKeyRequest{
		KeyId: key.KeyId,
	})

	assert.Equal(t, updated.KeyInfo.Name, newName)

	k.DeleteApiKey(context.Background(), &cxsdk.DeleteApiKeyRequest{
		KeyId: key.KeyId,
	})
}
