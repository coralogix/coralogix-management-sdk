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

func TestURLFromDomain(t *testing.T) {
	cases := map[string]string{
		// The base domain is prefixed with "api." to form the OpenAPI host.
		"acme.coralogix.com": "https://api.acme.coralogix.com/mgmt/openapi/5",
		// An already-prefixed domain is accepted as-is (no double "api.").
		"api.acme.coralogix.com":  "https://api.acme.coralogix.com/mgmt/openapi/5",
		"mycompany.coralogix.com": "https://api.mycompany.coralogix.com/mgmt/openapi/5",
	}
	for in, want := range cases {
		if got := URLFromDomain(in); got != want {
			t.Errorf("URLFromDomain(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestURLFromDomainMatchesURLFromRegion(t *testing.T) {
	// Passing a region's base domain to URLFromDomain must yield the same host
	// that URLFromRegion produces for that region, so the two configuration
	// paths are interchangeable.
	want, ok := URLFromRegion("eu2")
	if !ok {
		t.Fatal("URLFromRegion(\"eu2\") returned not-ok")
	}
	if got := URLFromDomain("eu2.coralogix.com"); got != want {
		t.Errorf("URLFromDomain(\"eu2.coralogix.com\") = %q, want %q (URLFromRegion parity)", got, want)
	}
}
