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
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// OpenAPIError is implemented by all SDK-generated GenericOpenAPIError structs.
// Each SDK package defines its own identical GenericOpenAPIError type,
// so we match them generically using this interface.
type OpenAPIError interface {
	error
	Body() []byte
	Model() interface{}
}

// APIError wraps SDK errors with HTTP metadata for richer context.
type APIError struct {
	Err        error
	StatusCode int
	Body       []byte
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e == nil {
		return "<nil>"
	}

	msg := strings.TrimSpace(e.Err.Error())
	body := strings.TrimSpace(string(e.Body))

	if body != "" {
		return fmt.Sprintf("%s — %s", msg, body)
	}
	return msg
}

// NewAPIError creates an APIError. It detects all SDK-generated GenericOpenAPIError types via interface matching.
func NewAPIError(resp *http.Response, err error) error {
	if err == nil {
		return nil
	}

	apiErr := &APIError{Err: err}

	if resp != nil {
		apiErr.StatusCode = resp.StatusCode
	}

	// Detect and extract from any SDK's GenericOpenAPIError
	if oapiErr, ok := err.(OpenAPIError); ok {
		apiErr.Body = oapiErr.Body()
	}

	return apiErr
}

// Code extracts the HTTP status code from an error, defaulting to 0 if unavailable.
func Code(err error) int {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode
	}
	return 0
}
