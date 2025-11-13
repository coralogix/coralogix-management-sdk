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
	"runtime"

	"github.com/google/uuid"
)

// CallPropertiesCreator provides configuration for building API clients.
type CallPropertiesCreator interface {
	URL() string
	Headers() map[string]string
}

// SDKCallPropertiesCreator implements CallPropertiesCreator.
type SDKCallPropertiesCreator struct {
	apiKey  string
	url     string
	headers map[string]string
}

// URL returns the base URL for API calls.
func (c *SDKCallPropertiesCreator) URL() string {
	return c.url
}

// Headers returns the headers to include in API calls.
func (c *SDKCallPropertiesCreator) Headers() map[string]string {
	return c.headers
}

// NewSDKCallPropertiesCreator creates a new CallPropertiesCreator with defaults.
func NewSDKCallPropertiesCreator(url, apiKey string) *SDKCallPropertiesCreator {
	return &SDKCallPropertiesCreator{
		url:    url,
		apiKey: apiKey,
		headers: map[string]string{
			"Authorization":            fmt.Sprintf("Bearer %s", apiKey),
			sdkVersionHeaderName:       vanillaSdkVersion,
			sdkLanguageHeaderName:      "go",
			sdkGoVersionHeaderName:     runtime.Version(),
			sdkCorrelationIDHeaderName: uuid.New().String(),
		},
	}
}

// NewSDKCallPropertiesCreatorTerraform creates a new CallPropertiesCreator, setting the Terraform version header.
func NewSDKCallPropertiesCreatorTerraform(url, apiKey, version string) *SDKCallPropertiesCreator {
	return &SDKCallPropertiesCreator{
		url:    url,
		apiKey: apiKey,
		headers: map[string]string{
			"Authorization":            fmt.Sprintf("Bearer %s", apiKey),
			sdkVersionHeaderName:       "terraform-" + version,
			sdkLanguageHeaderName:      "go",
			sdkGoVersionHeaderName:     runtime.Version(),
			sdkCorrelationIDHeaderName: uuid.New().String(),
		},
	}
}

// NewSDKCallPropertiesCreatorOperator creates a new CallPropertiesCreator, setting the Operator version header.
func NewSDKCallPropertiesCreatorOperator(url, apiKey, version string) *SDKCallPropertiesCreator {
	return &SDKCallPropertiesCreator{
		url:    url,
		apiKey: apiKey,
		headers: map[string]string{
			"Authorization":            fmt.Sprintf("Bearer %s", apiKey),
			sdkVersionHeaderName:       "cxo-" + version,
			sdkLanguageHeaderName:      "go",
			sdkGoVersionHeaderName:     runtime.Version(),
			sdkCorrelationIDHeaderName: uuid.New().String(),
		},
	}
}
