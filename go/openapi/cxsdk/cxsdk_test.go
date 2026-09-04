// Copyright 2026 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cxsdk

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestClientSetIdentityAndUsers(t *testing.T) {
	const testHeader = "shared-config"

	config := NewConfigBuilder().
		WithURL("https://example.test").
		WithHeader("X-Test-Header", testHeader).
		Build()
	config.httpClient = &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		assertTestHeader(t, request, testHeader)

		var body string
		switch request.URL.Path {
		case "/aaa/identity/v1/whoami":
			body = `{"teamId":123,"teamName":"test-team"}`
		case "/aaa/teams/v2/123/search":
			body = `{"users":[],"totalCount":0}`
		default:
			return nil, fmt.Errorf("unexpected request path: %s", request.URL.Path)
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(body)),
			Request:    request,
		}, nil
	})}
	clientSet := NewClientSet(config)

	if clientSet.Identity() == nil {
		t.Fatal("Identity() returned nil")
	}
	if clientSet.Users() == nil {
		t.Fatal("Users() returned nil")
	}

	whoami, _, err := clientSet.Identity().IdentityServiceWhoAmI(context.Background()).Execute()
	if err != nil {
		t.Fatalf("WhoAmI request failed: %v", err)
	}
	if whoami.GetTeamId() != 123 {
		t.Fatalf("WhoAmI team ID = %d, want 123", whoami.GetTeamId())
	}

	search, _, err := clientSet.Users().UsersMgmtServiceSearchUsers(context.Background(), 123).Execute()
	if err != nil {
		t.Fatalf("SearchUsers request failed: %v", err)
	}
	if len(search.GetUsers()) != 0 {
		t.Fatalf("SearchUsers returned %d users, want 0", len(search.GetUsers()))
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

func assertTestHeader(t *testing.T, request *http.Request, want string) {
	t.Helper()
	if got := request.Header.Get("X-Test-Header"); got != want {
		t.Errorf("X-Test-Header = %q, want %q", got, want)
	}
}
