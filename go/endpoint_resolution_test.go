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

import "testing"

func TestCoralogixGrpcEndpointFromRegion(t *testing.T) {
	cases := map[string]string{
		// Known region shorthands resolve to the documented gRPC hosts.
		"eu2": GrpcEU2,
		"us1": GrpcUS1,
		// A base custom/BYOC domain is prefixed with ng-api-grpc.
		"acme.coralogix.com": "ng-api-grpc.acme.coralogix.com:443",
		// The OpenAPI host form ("api.<domain>") is normalized to the base
		// gRPC host, not a double-prefixed "ng-api-grpc.api.<domain>".
		"api.acme.coralogix.com": "ng-api-grpc.acme.coralogix.com:443",
	}
	for in, want := range cases {
		if got := CoralogixGrpcEndpointFromRegion(in); got != want {
			t.Errorf("CoralogixGrpcEndpointFromRegion(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestCoralogixGrpcEndpointFromDomain(t *testing.T) {
	cases := map[string]string{
		"acme.coralogix.com":      "ng-api-grpc.acme.coralogix.com:443",
		"api.acme.coralogix.com":  "ng-api-grpc.acme.coralogix.com:443",
		"mycompany.coralogix.com": "ng-api-grpc.mycompany.coralogix.com:443",
	}
	for in, want := range cases {
		if got := CoralogixGrpcEndpointFromDomain(in); got != want {
			t.Errorf("CoralogixGrpcEndpointFromDomain(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestCoralogixRestEndpointFromRegion(t *testing.T) {
	cases := map[string]string{
		"eu2":                    RestEU2,
		"acme.coralogix.com":     "https://ng-api-http.acme.coralogix.com",
		"api.acme.coralogix.com": "https://ng-api-http.acme.coralogix.com",
	}
	for in, want := range cases {
		if got := CoralogixRestEndpointFromRegion(in); got != want {
			t.Errorf("CoralogixRestEndpointFromRegion(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestCoralogixRestEndpointFromDomain(t *testing.T) {
	cases := map[string]string{
		"acme.coralogix.com":     "https://ng-api-http.acme.coralogix.com",
		"api.acme.coralogix.com": "https://ng-api-http.acme.coralogix.com",
	}
	for in, want := range cases {
		if got := CoralogixRestEndpointFromDomain(in); got != want {
			t.Errorf("CoralogixRestEndpointFromDomain(%q) = %q, want %q", in, got, want)
		}
	}
}
