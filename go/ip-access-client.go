// Copyright 2025 Coralogix Ltd.
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
	"context"

	ipaccess "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/aaa/v1"
)

type (
	// CreateCompanyIPAccessSettingsRequest creates a request to create company IP access settings.
	CreateCompanyIPAccessSettingsRequest = ipaccess.CreateCompanyIpAccessSettingsRequest
	// CreateCompanyIPAccessSettingsResponse creates a response for creating company IP access settings.
	CreateCompanyIPAccessSettingsResponse = ipaccess.CreateCompanyIpAccessSettingsResponse
	// ReplaceCompanyIPAccessSettingsRequest replaces company IP access settings.
	ReplaceCompanyIPAccessSettingsRequest = ipaccess.ReplaceCompanyIpAccessSettingsRequest
	// ReplaceCompanyIPAccessSettingsResponse replaces company IP access settings response.
	ReplaceCompanyIPAccessSettingsResponse = ipaccess.ReplaceCompanyIpAccessSettingsResponse
	// GetCompanyIPAccessSettingsRequest retrieves company IP access settings.
	GetCompanyIPAccessSettingsRequest = ipaccess.GetCompanyIpAccessSettingsRequest
	// GetCompanyIPAccessSettingsResponse retrieves company IP access settings response.
	GetCompanyIPAccessSettingsResponse = ipaccess.GetCompanyIpAccessSettingsResponse
	// DeleteCompanyIPAccessSettingsRequest deletes company IP access settings.
	DeleteCompanyIPAccessSettingsRequest = ipaccess.DeleteCompanyIpAccessSettingsRequest
	// DeleteCompanyIPAccessSettingsResponse deletes company IP access settings response.
	DeleteCompanyIPAccessSettingsResponse = ipaccess.DeleteCompanyIpAccessSettingsResponse
	// CoralogixCustomerSupportAccess represents the access level for Coralogix customer support.
	CoralogixCustomerSupportAccess = ipaccess.CoralogixCustomerSupportAccess
	// IPAccess represents the IP access settings for a company.
	IPAccess = ipaccess.IpAccess
	// CompanyIPAccessSettings represents the company IP access settings.
	CompanyIPAccessSettings = ipaccess.CompanyIpAccessSettings
)

// CoralogixCustomerSupportAccess values.
const (
	CoralogixCustomerSupportAccessUnspecified = ipaccess.CoralogixCustomerSupportAccess_CORALOGIX_CUSTOMER_SUPPORT_ACCESS_UNSPECIFIED
	CoralogixCustomerSupportAccessDisabled    = ipaccess.CoralogixCustomerSupportAccess_CORALOGIX_CUSTOMER_SUPPORT_ACCESS_DISABLED
	CoralogixCustomerSupportAccessEnabled     = ipaccess.CoralogixCustomerSupportAccess_CORALOGIX_CUSTOMER_SUPPORT_ACCESS_ENABLED
)

const ipAccessFeatureGroupID = "aaa"

// RPC method names
const (
	CreateCompanyIPAccessSettingsRPC  = ipaccess.IpAccessService_CreateCompanyIpAccessSettings_FullMethodName
	ReplaceCompanyIPAccessSettingsRPC = ipaccess.IpAccessService_ReplaceCompanyIpAccessSettings_FullMethodName
	GetCompanyIPAccessSettingsRPC     = ipaccess.IpAccessService_GetCompanyIpAccessSettings_FullMethodName
	DeleteCompanyIPAccessSettingsRPC  = ipaccess.IpAccessService_DeleteCompanyIpAccessSettings_FullMethodName
)

// Create creates company IP access settings.
func (c IPAccessClient) Create(ctx context.Context, req *CreateCompanyIPAccessSettingsRequest) (*CreateCompanyIPAccessSettingsResponse, error) {
	callProps, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProps.Connection
	defer conn.Close()
	client := ipaccess.NewIpAccessServiceClient(conn)

	res, err := client.CreateCompanyIpAccessSettings(callProps.Ctx, req, callProps.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, CreateCompanyIPAccessSettingsRPC, ipAccessFeatureGroupID)
	}
	return res, nil
}

// Replace replaces company IP access settings.
func (c IPAccessClient) Replace(ctx context.Context, req *ReplaceCompanyIPAccessSettingsRequest) (*ReplaceCompanyIPAccessSettingsResponse, error) {
	callProps, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProps.Connection
	defer conn.Close()
	client := ipaccess.NewIpAccessServiceClient(conn)

	res, err := client.ReplaceCompanyIpAccessSettings(callProps.Ctx, req, callProps.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ReplaceCompanyIPAccessSettingsRPC, ipAccessFeatureGroupID)
	}
	return res, nil
}

// Get retrieves company IP access settings.
func (c IPAccessClient) Get(ctx context.Context, req *GetCompanyIPAccessSettingsRequest) (*GetCompanyIPAccessSettingsResponse, error) {
	callProps, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProps.Connection
	defer conn.Close()
	client := ipaccess.NewIpAccessServiceClient(conn)

	res, err := client.GetCompanyIpAccessSettings(callProps.Ctx, req, callProps.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetCompanyIPAccessSettingsRPC, ipAccessFeatureGroupID)
	}
	return res, nil
}

// Delete deletes company IP access settings.
func (c IPAccessClient) Delete(ctx context.Context, req *DeleteCompanyIPAccessSettingsRequest) (*DeleteCompanyIPAccessSettingsResponse, error) {
	callProps, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProps.Connection
	defer conn.Close()
	client := ipaccess.NewIpAccessServiceClient(conn)

	res, err := client.DeleteCompanyIpAccessSettings(callProps.Ctx, req, callProps.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, DeleteCompanyIPAccessSettingsRPC, ipAccessFeatureGroupID)
	}
	return res, nil
}

// IPAccessClient is a client for the Coralogix IP Access API.
type IPAccessClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// NewIPAccessClient creates a new IP Access client.
func NewIPAccessClient(c *CallPropertiesCreator) *IPAccessClient {
	return &IPAccessClient{callPropertiesCreator: c}
}
