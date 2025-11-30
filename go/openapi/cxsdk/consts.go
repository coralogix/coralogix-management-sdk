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

const (
	envCoralogixAPIKey = "CORALOGIX_API_KEY"
	envCoralogixRegion = "CORALOGIX_REGION"
	envCoralogixDomain = "CORALOGIX_DOMAIN"
	envCoralogixURL    = "CORALOGIX_URL"

	authorizationHeaderName    = "Authorization"
	sdkVersionHeaderName       = "x-cx-sdk-version"
	sdkLanguageHeaderName      = "x-cx-sdk-language"
	sdkGoVersionHeaderName     = "x-cx-go-version"
	sdkCorrelationIDHeaderName = "x-cx-correlation-id"
	vanillaSdkVersion          = "1.12.0"
)
