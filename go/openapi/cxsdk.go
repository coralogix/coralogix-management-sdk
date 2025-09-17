package cxsdk

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/google/uuid"
)

const (
	EnvCoralogixRegion = "CORALOGIX_REGION"
	EnvCoralogixAPIKey = "CORALOGIX_API_KEY"
)

const (
	OpenAPIEU1 = "https://api.coralogix.com/mgmt/openapi"
	OpenAPIEU2 = "https://api.eu2.coralogix.com/mgmt/openapi"
	OpenAPIUS1 = "https://api.coralogix.us/mgmt/openapi"
	OpenAPIUS2 = "https://api.cx498.coralogix.com/mgmt/openapi"
	OpenAPIAP2 = "https://api.coralogixsg.com/mgmt/openapi"
	OpenAPIAP3 = "https://api.ap3.coralogix.com/mgmt/openapi"
)

// ClientBuilder helps build a Client with flexible options.
type ClientBuilder struct {
	baseURL    string
	headers    map[string]string
	httpClient *http.Client
	sdkVersion string
}

// NewClientBuilder creates a new builder with defaults.
func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		headers:    make(map[string]string),
		httpClient: http.DefaultClient,
		sdkVersion: vanillaSdkVersion, // default version
	}
}

// WithRegion sets the API endpoint from region.
func (b *ClientBuilder) WithRegion(region string) *ClientBuilder {
	if b.baseURL == "" {
		b.baseURL = CoralogixRestEndpointFromRegion(region)
	}
	return b
}

// WithBaseURL sets a custom base URL (overrides region mapping).
func (b *ClientBuilder) WithBaseURL(url string) *ClientBuilder {
	b.baseURL = url
	return b
}

// WithAPIKey adds the standard Authorization header.
func (b *ClientBuilder) WithAPIKey(apiKey string) *ClientBuilder {
	b.headers["Authorization"] = fmt.Sprintf("Bearer %s", apiKey)
	return b
}

// WithHeader allows arbitrary header injection.
func (b *ClientBuilder) WithHeader(key, value string) *ClientBuilder {
	b.headers[key] = value
	return b
}

// WithTerraformSDKVersion sets the SDK version header to indicate the Terraform provider version.
func (b *ClientBuilder) WithTerraformSDKVersion(version string) *ClientBuilder {
	b.sdkVersion = "terraform-" + version
	return b
}

// WithOperatorVersion sets the SDK version header to indicate the Coralogix Operator version.
func (b *ClientBuilder) WithOperatorVersion(version string) *ClientBuilder {
	b.sdkVersion = "cxo-" + version
	return b
}

func (b *ClientBuilder) WithHttpClient(client *http.Client) *ClientBuilder {
	b.httpClient = client
	return b
}

// Build creates the Client.
func (b *ClientBuilder) Build() *APIClient {
	if b.baseURL == "" {
		panic("ClientBuilder: baseURL is not set (call WithRegion or WithBaseURL)")
	}

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{URL: b.baseURL, Description: "Coralogix endpoint: " + b.baseURL},
	}

	cfg.AddDefaultHeader(sdkVersionHeaderName, b.sdkVersion)
	cfg.AddDefaultHeader(sdkLanguageHeaderName, "go")
	cfg.AddDefaultHeader(sdkGoVersionHeaderName, runtime.Version())
	cfg.AddDefaultHeader(sdkCorrelationIDHeaderName, uuid.New().String())

	for k, v := range b.headers {
		cfg.AddDefaultHeader(k, v)
	}

	cfg.HTTPClient = b.httpClient

	return NewAPIClient(cfg)
}

// CoralogixRestEndpointFromRegion maps a region identifier to its REST base URL.
func CoralogixRestEndpointFromRegion(region string) string {
	switch strings.ToLower(region) {
	case "eu1":
		return OpenAPIEU1
	case "eu2":
		return OpenAPIEU2
	case "us1":
		return OpenAPIUS1
	case "us2":
		return OpenAPIUS2
	case "ap2":
		return OpenAPIAP2
	case "ap3":
		return OpenAPIAP3
	default:
		return fmt.Sprintf("https://api.%s.coralogix.com/mgmt/openapi", region)
	}
}
