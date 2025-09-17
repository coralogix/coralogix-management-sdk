package cxsdk

import (
	"fmt"
	"github.com/google/uuid"
	"runtime"
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
