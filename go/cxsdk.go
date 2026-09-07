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
	"fmt"
	"log"
	"os"
	"strings"
)

// REST URLs for the Coralogix regions. These are served from the same api.<domain>
// host as the OpenAPI endpoints, so a tenant only needs that one host resolvable.
const (
	RestUS1 = "https://api.coralogix.us"
	RestUS2 = "https://api.cx498.coralogix.com"
	RestUS3 = "https://api.us3.coralogix.com"
	RestEU1 = "https://api.coralogix.com"
	RestEU2 = "https://api.eu2.coralogix.com"
	RestAP1 = "https://api.app.coralogix.in"
	RestAP2 = "https://api.coralogixsg.com"
	RestAP3 = "https://api.ap3.coralogix.com"
)

// EnvCoralogxRegion is the name of the environment variable that contains the Coralogix region.
const EnvCoralogxRegion = "CORALOGIX_REGION"

// EnvCoralogixTeamLevelAPIKey is the name of the environment variable that contains the Coralogix Team API key.
const EnvCoralogixTeamLevelAPIKey = "CORALOGIX_TEAM_API_KEY"

// EnvCoralogixUserLevelAPIKey is the name of the environment variable that contains the Coralogix User API key.
const EnvCoralogixUserLevelAPIKey = "CORALOGIX_USER_API_KEY"

// CoralogixRegionFromEnv reads the Coralogix region from environment variables.
func CoralogixRegionFromEnv() (string, error) {
	regionIdentifier := strings.ToLower(os.Getenv(EnvCoralogxRegion))

	if regionIdentifier == "" {
		return "", fmt.Errorf("the %s environment variable is not set", EnvCoralogxRegion)
	}
	return regionIdentifier, nil
}

// normalizeCoralogixDomain reduces a Coralogix domain to its base form by
// stripping a leading "api." label. Callers identify a tenant by its API host
// (e.g. "api.eu2.coralogix.com"), while REST endpoints are derived by prefixing
// the base domain (e.g. "eu2.coralogix.com") with "api.".
func normalizeCoralogixDomain(domain string) string {
	return strings.TrimPrefix(domain, "api.")
}

// CoralogixRestEndpointFromRegion reads the Coralogix REST endpoint from a region
// shorthand or a custom domain.
func CoralogixRestEndpointFromRegion(regionIdentifier string) string {
	switch regionIdentifier {
	case "us1":
		return RestUS1
	case "us2":
		return RestUS2
	case "us3":
		return RestUS3
	case "eu1":
		return RestEU1
	case "eu2":
		return RestEU2
	case "ap1":
		return RestAP1
	case "ap2":
		return RestAP2
	case "ap3":
		return RestAP3
	default:
		return CoralogixRestEndpointFromDomain(regionIdentifier)
	}
}

// CoralogixRestEndpointFromDomain returns the REST endpoint for a custom
// Coralogix domain (e.g. "eu2.coralogix.com" or "mycompany.coralogix.com").
// A leading "api." is tolerated so the same domain value used to configure the
// OpenAPI client also yields the correct REST host.
func CoralogixRestEndpointFromDomain(domain string) string {
	return fmt.Sprintf("https://api.%s", normalizeCoralogixDomain(domain))
}

// AuthContext is a struct that holds the API keys for the Coralogix SDK.
type AuthContext struct {
	teamLevelAPIKey string
	userLevelAPIKey string
}

// NewAuthContext creates a new AuthContext.
func NewAuthContext(teamLevelAPIKey, userLevelAPIKey string) AuthContext {
	if teamLevelAPIKey == "" {
		log.Println("Warning: teamLevelAPIKey was not provided. Some functionality will not be available.")
	}
	if userLevelAPIKey == "" {
		log.Println("Warning: userLevelAPIKey was not provided. Some functionality will not be available.")
	}
	return AuthContext{
		teamLevelAPIKey: teamLevelAPIKey,
		userLevelAPIKey: userLevelAPIKey,
	}
}

// AuthContextFromEnv reads the Coralogix API keys from environment variables.
func AuthContextFromEnv() (AuthContext, error) {
	teamLevelAPIKey, err := CoralogixTeamsLevelAPIKeyFromEnv()
	if err != nil {
		return AuthContext{}, err
	}
	userLevelAPIKey, err := CoralogixUserLevelAPIKeyFromEnv()
	if err != nil {
		return AuthContext{}, err
	}
	if teamLevelAPIKey == "" {
		log.Println("Warning: teamLevelAPIKey is empty. Some functionality will not be available.")
	}
	if userLevelAPIKey == "" {
		log.Println("Warning: userLevelAPIKey is empty. Some functionality will not be available.")
	}

	return AuthContext{teamLevelAPIKey: teamLevelAPIKey, userLevelAPIKey: userLevelAPIKey}, nil
}

// CoralogixTeamsLevelAPIKeyFromEnv reads the Coralogix Team API key from environment variables.
func CoralogixTeamsLevelAPIKeyFromEnv() (string, error) {
	apiKey := os.Getenv(EnvCoralogixTeamLevelAPIKey)
	if apiKey == "" {
		return "", fmt.Errorf("the %s environment variable is not set", EnvCoralogixTeamLevelAPIKey)
	}
	return apiKey, nil
}

// CoralogixUserLevelAPIKeyFromEnv reads the Coralogix User API key from environment variables.
func CoralogixUserLevelAPIKeyFromEnv() (string, error) {
	apiKey := os.Getenv(EnvCoralogixUserLevelAPIKey)
	if apiKey == "" {
		return "", fmt.Errorf("the %s environment variable is not set", EnvCoralogixUserLevelAPIKey)
	}
	return apiKey, nil
}
