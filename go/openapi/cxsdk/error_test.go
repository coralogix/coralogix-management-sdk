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
	"errors"
	"net/http"
	"testing"
)

func TestIsResourceNotFoundBody(t *testing.T) {
	cases := []struct {
		name string
		body string
		want bool
	}{
		// Legacy plain-text bodies.
		{"text resource not found", "Not Found: Scope not found", true},
		{"text endpoint not found", "Not Found: Not Found", false},
		{"text empty message", "Not Found: ", false},
		{"text unrelated", "Bad Request: something", false},
		// New JSON error bodies ({"code","message"}).
		{"json resource not found", `{"code":404,"message":"Not Found: Scope not found"}`, true},
		{"json endpoint not found (grpc-gateway default)", `{"code":5,"message":"Not Found"}`, false},
		{"json generic not found", `{"code":404,"message":"Not Found: Not Found"}`, false},
		// Degenerate inputs fall back to the raw-body check.
		{"empty", "", false},
		{"malformed json", `{"code":404,`, false},
		{"json without message", `{"code":404}`, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := isResourceNotFoundBody([]byte(tc.body)); got != tc.want {
				t.Errorf("isResourceNotFoundBody(%q) = %v, want %v", tc.body, got, tc.want)
			}
		})
	}
}

func TestIsNotFound(t *testing.T) {
	cases := []struct {
		name       string
		statusCode int
		body       string
		want       bool
	}{
		{"404 json resource", http.StatusNotFound, `{"code":404,"message":"Not Found: Scope not found"}`, true},
		{"404 text resource", http.StatusNotFound, "Not Found: Scope not found", true},
		{"404 json endpoint", http.StatusNotFound, `{"code":5,"message":"Not Found"}`, false},
		{"non-404 status", http.StatusInternalServerError, `{"code":500,"message":"Not Found: Scope not found"}`, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := &APIError{
				Err:        errors.New(tc.body),
				StatusCode: tc.statusCode,
				Body:       []byte(tc.body),
			}
			if got := IsNotFound(err); got != tc.want {
				t.Errorf("IsNotFound(status=%d, body=%q) = %v, want %v", tc.statusCode, tc.body, got, tc.want)
			}
		})
	}
}
