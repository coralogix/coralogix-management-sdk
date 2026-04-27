// Copyright 2025 Coralogix Ltd.
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

package examples

import (
	"os"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
)

// newTestConfig creates a Config for tests by checking environment variables in order:
// 1. CORALOGIX_URL - full URL to the API
// 2. CORALOGIX_DOMAIN - custom domain (e.g., "api.staging.com")
// 3. CORALOGIX_REGION - region identifier (e.g., "eu2", "us1")
//
// Empty values are ignored to allow fallback (e.g., CORALOGIX_URL="" with valid CORALOGIX_REGION).
func newTestConfig() *cxsdk.Config {
	builder := cxsdk.NewConfigBuilder().WithAPIKeyEnv()

	if url := os.Getenv("CORALOGIX_URL"); url != "" {
		return builder.WithURLEnv().Build()
	}

	if domain := os.Getenv("CORALOGIX_DOMAIN"); domain != "" {
		return builder.WithDomainEnv().Build()
	}

	if region := os.Getenv("CORALOGIX_REGION"); region != "" {
		return builder.WithRegionEnv().Build()
	}

	panic("one of CORALOGIX_URL, CORALOGIX_DOMAIN, or CORALOGIX_REGION environment variables must be set (non-empty)")
}
