// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cxsdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"runtime"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RestClient for Coralogix API
type RestClient struct {
	url           string
	apiKey        string
	client        *http.Client
	correlationID string
}

// NewRestClient creates a new RestClient
func NewRestClient(url string, apiKey string) *RestClient {
	correlationID := uuid.New().String()
	return &RestClient{url, apiKey, &http.Client{}, correlationID}
}

// Request executes request to Coralogix API
func (c *RestClient) Request(ctx context.Context, method, path, contentType string, body interface{}, headers ...string) (string, error) {
	var request *http.Request
	if body != nil {
		bodyReader := bytes.NewBuffer([]byte(body.(string)))
		var err error
		request, err = http.NewRequest(method, c.url+path, bodyReader)
		if err != nil {
			return "", err
		}

		request.Header.Set("Content-Type", contentType)
	} else {
		request, _ = http.NewRequest(method, c.url+path, nil)
	}

	request = request.WithContext(ctx)
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Authorization", "Bearer "+c.apiKey)
	request.Header.Set(sdkVersionHeaderName, vanillaSdkVersion)
	request.Header.Set(sdkLanguageHeaderName, "go")
	request.Header.Set(sdkGoVersionHeaderName, runtime.Version())
	request.Header.Set(sdkCorrelationIDHeaderName, c.correlationID)
	if len(headers)%2 != 0 {
		return "", fmt.Errorf("invalid headers, must be key-value pairs")
	}
	for i := 0; i < len(headers); i += 2 {
		request.Header.Set(headers[i], headers[i+1])
	}

	response, err := c.client.Do(request)
	if err != nil {
		return "", status.Convert(err).Err()
	}
	defer response.Body.Close()

	if response.StatusCode >= http.StatusOK && response.StatusCode < http.StatusMultipleChoices {

		bodyResp, err := io.ReadAll(response.Body)
		if err != nil {
			return "", err
		}

		return string(bodyResp), nil
	}

	if response.StatusCode == http.StatusNotFound {
		return "", status.Error(codes.NotFound, "Not found")
	}

	responseBody, err := httputil.DumpResponse(response, true)
	if err != nil {
		return "", status.Convert(err).Err()
	}

	return "", fmt.Errorf("API Error: %s. Status code: %s", string(responseBody), response.Status)
}

// Get executes GET request to Coralogix API
func (c *RestClient) Get(ctx context.Context, path string, headers ...string) (string, error) {
	return c.Request(ctx, "GET", path, "", nil, headers...)
}

// Post executes POST request to Coralogix API
func (c *RestClient) Post(ctx context.Context, path, contentType, body string, headers ...string) (string, error) {
	return c.Request(ctx, "POST", path, contentType, body, headers...)
}

// Put executes PUT request to Coralogix API
func (c *RestClient) Put(ctx context.Context, path, contentType, body string, headers ...string) (string, error) {
	return c.Request(ctx, "PUT", path, contentType, body, headers...)
}

// Delete executes DELETE request to Coralogix API
func (c *RestClient) Delete(ctx context.Context, path string, headers ...string) (string, error) {
	return c.Request(ctx, "DELETE", path, "", nil, headers...)
}
