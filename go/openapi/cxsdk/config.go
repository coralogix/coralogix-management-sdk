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
	"log/slog"
	"net/http"
	gourl "net/url"
	"os"
	"runtime"
	"strings"

	"github.com/google/uuid"
)

// Config holds the configuration for the Coralogix OpenAPI SDK.
type Config struct {
	url        string
	headers    map[string]string
	httpClient *http.Client
}

// ConfigBuilder is used to build a Config.
type ConfigBuilder struct {
	url     string
	headers map[string]string
	logger  *slog.Logger
}

// NewConfigBuilder creates a new ConfigBuilder with default vanilla SDK version tag.
func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		headers: map[string]string{
			sdkVersionHeaderName:       "sdk-" + vanillaSdkVersion,
			sdkLanguageHeaderName:      "go",
			sdkGoVersionHeaderName:     runtime.Version(),
			sdkCorrelationIDHeaderName: uuid.New().String(),
		},
	}
}

// WithHeader adds a custom header to the ConfigBuilder.
func (b *ConfigBuilder) WithHeader(key, value string) *ConfigBuilder {
	b.headers[key] = value
	return b
}

// WithHeaders adds multiple custom headers to the ConfigBuilder.
func (b *ConfigBuilder) WithHeaders(headers map[string]string) *ConfigBuilder {
	for k, v := range headers {
		b.WithHeader(k, v)
	}
	return b
}

// WithAPIKey sets the API key for the ConfigBuilder.
func (b *ConfigBuilder) WithAPIKey(apiKey string) *ConfigBuilder {
	return b.WithHeader(authorizationHeaderName, fmt.Sprintf("Bearer %s", apiKey))
}

// WithAPIKeyEnv sets the API key from the CORALOGIX_API_KEY environment variable.
func (b *ConfigBuilder) WithAPIKeyEnv() *ConfigBuilder {
	apiKey, ok := os.LookupEnv(envCoralogixAPIKey)
	if !ok {
		panic("Environment variable CORALOGIX_API_KEY is not set")
	}
	return b.WithAPIKey(apiKey)
}

// WithVersion sets a generic version tag passed from the caller.
func (b *ConfigBuilder) WithVersion(version string) *ConfigBuilder {
	if version == "" {
		panic("version is empty")
	}
	return b.WithHeader(sdkVersionHeaderName, version)
}

// WithTerraformVersion sets the version tag to the Terraform provider version.
func (b *ConfigBuilder) WithTerraformVersion(version string) *ConfigBuilder {
	if version == "" {
		panic("Terraform version is empty")
	}
	return b.WithVersion("terraform-" + version)
}

// WithOperatorVersion sets the version tag to the Operator version.
func (b *ConfigBuilder) WithOperatorVersion(version string) *ConfigBuilder {
	if version == "" {
		panic("Operator version is empty")
	}
	return b.WithVersion("cxo-" + version)
}

// WithRegion sets the region for the ConfigBuilder.
func (b *ConfigBuilder) WithRegion(region string) *ConfigBuilder {
	if url, ok := URLFromRegion(region); ok {
		b.url = url
	}
	return b
}

// WithRegionEnv sets the region from the CORALOGIX_REGION environment variable.
func (b *ConfigBuilder) WithRegionEnv() *ConfigBuilder {
	region, ok := os.LookupEnv(envCoralogixRegion)
	if !ok {
		panic("Environment variable CORALOGIX_REGION is not set")
	}
	return b.WithRegion(region)
}

// WithDomain sets the custom domain for the ConfigBuilder.
func (b *ConfigBuilder) WithDomain(domain string) *ConfigBuilder {
	b.url = URLFromDomain(domain)
	return b
}

// WithDomainEnv sets the custom domain from the CORALOGIX_DOMAIN environment variable.
func (b *ConfigBuilder) WithDomainEnv() *ConfigBuilder {
	domain, ok := os.LookupEnv(envCoralogixDomain)
	if !ok {
		panic("Environment variable CORALOGIX_DOMAIN is not set")
	}
	return b.WithDomain(domain)
}

// WithURL sets the URL for the ConfigBuilder.
func (b *ConfigBuilder) WithURL(url string) *ConfigBuilder {
	if url == "" {
		panic("URL is empty")
	}
	b.url = url
	return b
}

// WithURLEnv sets the URL from the CORALOGIX_URL environment variable.
func (b *ConfigBuilder) WithURLEnv() *ConfigBuilder {
	url, ok := os.LookupEnv(envCoralogixURL)
	if !ok {
		panic("Environment variable CORALOGIX_URL is not set")
	}
	return b.WithURL(url)
}

// WithDefaultLogger sets a default JSON logger for the ConfigBuilder.
func (b *ConfigBuilder) WithDefaultLogger() *ConfigBuilder {
	b.logger = slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	)
	return b
}

// WithLogger sets a custom logger for the ConfigBuilder.
func (b *ConfigBuilder) WithLogger(logger *slog.Logger) *ConfigBuilder {
	b.logger = logger
	return b
}

// Build builds the Config.
func (b *ConfigBuilder) Build() *Config {
	var httpClient *http.Client
	if b.logger != nil {
		httpClient = &http.Client{
			Transport: NewLoggingTransport(b.logger),
		}
	}

	return &Config{
		url:        b.url,
		headers:    b.headers,
		httpClient: httpClient,
	}
}

// URLFromRegion returns the Coralogix OpenAPI URL for the given region.
func URLFromRegion(region string) (string, bool) {
	valid := true
	url := gourl.URL{
		Scheme: "https",
		Path:   "mgmt/openapi/3",
	}

	switch strings.ToLower(region) {
	case "eu1":
		url.Host = "api.eu1.coralogix.com"
	case "eu2":
		url.Host = "api.eu2.coralogix.com"
	case "us1":
		url.Host = "api.us1.coralogix.com"
	case "us2":
		url.Host = "api.us2.coralogix.com"
	case "ap1":
		url.Host = "api.ap1.coralogix.com"
	case "ap2":
		url.Host = "api.ap2.coralogix.com"
	case "ap3":
		url.Host = "api.ap3.coralogix.com"
	default:
		valid = false
	}
	return url.String(), valid
}

// URLFromDomain returns the Coralogix OpenAPI URL for the given custom domain.
func URLFromDomain(domain string) string {
	url := gourl.URL{
		Scheme: "https",
		Path:   "mgmt/openapi/3",
		Host:   domain,
	}
	return url.String()
}
