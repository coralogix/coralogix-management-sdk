package examples

// import (
// 	"context"
// 	cxsdk "coralogix-management-sdk/go"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestApiKeys(t *testing.T) {
// 	region, err := cxsdk.CoralogixRegionFromEnv()
// 	assert.Nil(t, err)
// 	apiKey, err := cxsdk.CoralogixAPIKeyFromEnv()
// 	assert.Nil(t, err)
// 	creator := cxsdk.NewCallPropertiesCreator(region, apiKey)
// 	k := cxsdk.NewAPIKeysClient(creator)

// 	key, e := k.CreateAPIKey(context.Background(), &cxsdk.CreateAPIKeyRequest{
// 		Name: "My APM KEY",
// 		Owner: &cxsdk.Owner{
// 			Owner: &cxsdk.OwnerUserID{
// 				UserId: "4013254",
// 			},
// 		},
// 		KeyPermissions: &cxsdk.APIKeyPermissions{
// 			Presets:     []string{"APM"},
// 			Permissions: []string{},
// 		},
// 	})
// 	assert.Nil(t, e)

// 	newName := "new-name"
// 	k.UpdateAPIKey(context.Background(), &cxsdk.UpdateAPIKeyRequest{
// 		KeyId:   key.KeyId,
// 		NewName: &newName,
// 	})

// 	updated, _ := k.GetAPIKey(context.Background(), &cxsdk.GetAPIKeyRequest{
// 		KeyId: key.KeyId,
// 	})

// 	assert.Equal(t, updated.KeyInfo.Name, newName)

// 	k.DeleteAPIKey(context.Background(), &cxsdk.DeleteAPIKeyRequest{
// 		KeyId: key.KeyId,
// 	})
// }
