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

package cxsdk

import (
	"fmt"

	"github.com/google/uuid"
)

// SDKCallPropertiesCreator holds region and API keys for REST clients such as SCIM Users.
type SDKCallPropertiesCreator struct {
	coraglogixRegion string
	teamsLevelAPIKey string
	userLevelAPIKey  string
	correlationID    string
	sdkVersion       string
}

// NewSDKCallPropertiesCreator creates a new SDKCallPropertiesCreator object.
func NewSDKCallPropertiesCreator(region string, authContext AuthContext) *SDKCallPropertiesCreator {
	return &SDKCallPropertiesCreator{
		coraglogixRegion: region,
		teamsLevelAPIKey: authContext.teamLevelAPIKey,
		userLevelAPIKey:  authContext.userLevelAPIKey,
		correlationID:    uuid.New().String(),
		sdkVersion:       fmt.Sprint("sdk-", vanillaSdkVersion),
	}
}

// NewSDKCallPropertiesCreatorTerraform creates a new SDKCallPropertiesCreator object, specifying which version of the Terraform Operator is being used.
func NewSDKCallPropertiesCreatorTerraform(region string, authContext AuthContext, terraformProviderVersion string) *SDKCallPropertiesCreator {
	return &SDKCallPropertiesCreator{
		coraglogixRegion: region,
		teamsLevelAPIKey: authContext.teamLevelAPIKey,
		userLevelAPIKey:  authContext.userLevelAPIKey,
		correlationID:    uuid.New().String(),
		sdkVersion:       fmt.Sprint("terraform-", terraformProviderVersion),
	}
}

// NewSDKCallPropertiesCreatorOperator creates a new SDKCallPropertiesCreator object, specifying which version of the Operator is being used.
func NewSDKCallPropertiesCreatorOperator(region string, authContext AuthContext, cxOperator string) *SDKCallPropertiesCreator {
	return &SDKCallPropertiesCreator{
		coraglogixRegion: region,
		teamsLevelAPIKey: authContext.teamLevelAPIKey,
		userLevelAPIKey:  authContext.userLevelAPIKey,
		correlationID:    uuid.New().String(),
		sdkVersion:       fmt.Sprint("cxo-", cxOperator),
	}
}

// NewSDKCallPropertiesCreatorCustomSdk creates a new SDKCallPropertiesCreator object with a custom version.
func NewSDKCallPropertiesCreatorCustomSdk(region string, authContext AuthContext, customSdkVersion string) *SDKCallPropertiesCreator {
	return &SDKCallPropertiesCreator{
		coraglogixRegion: region,
		teamsLevelAPIKey: authContext.teamLevelAPIKey,
		userLevelAPIKey:  authContext.userLevelAPIKey,
		correlationID:    uuid.New().String(),
		sdkVersion:       fmt.Sprint("custom-", customSdkVersion),
	}
}
