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
		// A base domain is prefixed with "api.", matching URLFromRegion.
		"factset.coralogix.com": "https://api.factset.coralogix.com/mgmt/openapi/5",
		"eu2.coralogix.com":     "https://api.eu2.coralogix.com/mgmt/openapi/5",
		// A domain that already carries the "api." prefix is used as-is
		// (backward compatible, no double prefix).
		"api.factset.coralogix.com": "https://api.factset.coralogix.com/mgmt/openapi/5",
	}
	for in, want := range cases {
		if got := URLFromDomain(in); got != want {
			t.Errorf("URLFromDomain(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestURLFromDomainMatchesRegion(t *testing.T) {
	// The custom-domain path for a standard region must produce the same URL
	// as the region path, so both configuration styles are interchangeable.
	want, ok := URLFromRegion("eu2")
	if !ok {
		t.Fatal("URLFromRegion(eu2) reported invalid")
	}
	if got := URLFromDomain("eu2.coralogix.com"); got != want {
		t.Errorf("URLFromDomain(eu2.coralogix.com) = %q, want %q (URLFromRegion)", got, want)
	}
}
