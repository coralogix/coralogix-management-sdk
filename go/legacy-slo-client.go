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
	"context"

	legacySlos "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/apm/services/v1"
)

// ServiceSlo is an SLO.
type ServiceSlo = legacySlos.ServiceSlo

// CreateLegacySloRequest is a request to create an SLO.
type CreateLegacySloRequest = legacySlos.CreateServiceSloRequest

// ServiceSloErrorSli is an SLO error SLI.
type ServiceSloErrorSli = legacySlos.ServiceSlo_ErrorSli

// ErrorSli is an SLO error SLI.
type ErrorSli = legacySlos.ErrorSli

// LatencySli is an SLO latency SLI.
type LatencySli = legacySlos.LatencySli

// ServiceSloLatencySli is a something.
type ServiceSloLatencySli = legacySlos.ServiceSlo_LatencySli

// GetLegacySloRequest is a request to get an SLO.
type GetLegacySloRequest = legacySlos.GetServiceSloRequest

// ReplaceLegacySloRequest is a request to replace an SLO.
type ReplaceLegacySloRequest = legacySlos.ReplaceServiceSloRequest

// DeleteLegacySloRequest is a request to delete an SLO.
type DeleteLegacySloRequest = legacySlos.DeleteServiceSloRequest

// SliFilter is an SLO filter.
type SliFilter = legacySlos.SliFilter

// ThresholdSymbol is an SLO threshold symbol.
type ThresholdSymbol = legacySlos.ThresholdSymbol

// SloThresholdSymbol values.
const (
	SloThresholdSymbolGreater        = legacySlos.ThresholdSymbol_THRESHOLD_SYMBOL_GREATER
	SloThresholdSymbolGreaterOrEqual = legacySlos.ThresholdSymbol_THRESHOLD_SYMBOL_GREATER_OR_EQUAL
	SloThresholdSymbolLess           = legacySlos.ThresholdSymbol_THRESHOLD_SYMBOL_LESS
	SloThresholdSymbolLessOrEqual    = legacySlos.ThresholdSymbol_THRESHOLD_SYMBOL_LESS_OR_EQUAL
	SloThresholdSymbolEqual          = legacySlos.ThresholdSymbol_THRESHOLD_SYMBOL_EQUAL
)

// SloPeriod is a SLO period.
type SloPeriod = legacySlos.SloPeriod

// SloStatus is a SLO status.
type SloStatus = legacySlos.SloStatus

// SLO period values.
const (
	SloPeriodUnspecified = legacySlos.SloPeriod_SLO_PERIOD_UNSPECIFIED
	SloPeriod7Days       = legacySlos.SloPeriod_SLO_PERIOD_7_DAYS
	SloPeriod14Days      = legacySlos.SloPeriod_SLO_PERIOD_14_DAYS
	SloPeriod30Days      = legacySlos.SloPeriod_SLO_PERIOD_30_DAYS
)

// CompareType is an SLO compare type.
type CompareType = legacySlos.CompareType

// SloCompareType values.
const (
	SloCompareTypeUnspecified = legacySlos.CompareType_COMPARE_TYPE_UNSPECIFIED
	SloCompareTypeIs          = legacySlos.CompareType_COMPARE_TYPE_IS
	SloCompareTypeStartsWith  = legacySlos.CompareType_COMPARE_TYPE_START_WITH
	SloCompareTypeEndsWith    = legacySlos.CompareType_COMPARE_TYPE_ENDS_WITH
	SloCompareTypeIncludes    = legacySlos.CompareType_COMPARE_TYPE_INCLUDES
)

// SloStatus values.
const (
	LegacySloStatusUnspecified = legacySlos.SloStatus_SLO_STATUS_UNSPECIFIED
	LegacySloStatusOk          = legacySlos.SloStatus_SLO_STATUS_OK
	LegacySloStatusBreached    = legacySlos.SloStatus_SLO_STATUS_BREACHED
)

// RPC names.
const (
	LegacySloGetRPC      = legacySlos.ServiceSloService_GetServiceSlo_FullMethodName
	LegacySloCreateRPC   = legacySlos.ServiceSloService_CreateServiceSlo_FullMethodName
	LegacySloReplaceRPC  = legacySlos.ServiceSloService_ReplaceServiceSlo_FullMethodName
	LegacySloDeleteRPC   = legacySlos.ServiceSloService_DeleteServiceSlo_FullMethodName
	SloListRPC           = legacySlos.ServiceSloService_ListServiceSlos_FullMethodName
	LegacySloBatchGetRPC = legacySlos.ServiceSloService_BatchGetServiceSlos_FullMethodName
)

// LegacySLOsClient is a client for the Coralogix SLOs API.
type LegacySLOsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new SLO.
func (c LegacySLOsClient) Create(ctx context.Context, req *legacySlos.CreateServiceSloRequest) (*legacySlos.CreateServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := legacySlos.NewServiceSloServiceClient(conn)

	response, err := client.CreateServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, SloCreateRPC, infraMonitoringFeatureGroupID)
	}
	return response, nil
}

// Get gets the specified SLO.
func (c LegacySLOsClient) Get(ctx context.Context, req *legacySlos.GetServiceSloRequest) (*legacySlos.GetServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := legacySlos.NewServiceSloServiceClient(conn)

	response, err := client.GetServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, SloGetRPC, infraMonitoringFeatureGroupID)
	}
	return response, nil
}

// Update updates the specified SLO.
func (c LegacySLOsClient) Update(ctx context.Context, req *legacySlos.ReplaceServiceSloRequest) (*legacySlos.ReplaceServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := legacySlos.NewServiceSloServiceClient(conn)

	response, err := client.ReplaceServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, SloReplaceRPC, infraMonitoringFeatureGroupID)
	}
	return response, nil
}

// Delete deletes the specified SLO.
func (c LegacySLOsClient) Delete(ctx context.Context, req *legacySlos.DeleteServiceSloRequest) (*legacySlos.DeleteServiceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := legacySlos.NewServiceSloServiceClient(conn)

	response, err := client.DeleteServiceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, SloDeleteRPC, infraMonitoringFeatureGroupID)
	}
	return response, nil
}

// GetBulk gets multiple SLOs in a single call.
func (c LegacySLOsClient) GetBulk(ctx context.Context, req *legacySlos.BatchGetServiceSlosRequest) (*legacySlos.BatchGetServiceSlosResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := legacySlos.NewServiceSloServiceClient(conn)

	response, err := client.BatchGetServiceSlos(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, SloBatchGetRPC, infraMonitoringFeatureGroupID)
	}
	return response, nil
}

// List lists all service SLOs.
func (c LegacySLOsClient) List(ctx context.Context, req *legacySlos.ListServiceSlosRequest) (*legacySlos.ListServiceSlosResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := legacySlos.NewServiceSloServiceClient(conn)

	response, err := client.ListServiceSlos(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, SloListRPC, infraMonitoringFeatureGroupID)
	}
	return response, nil
}

// NewLegacySLOsClient creates a new SLOs client.
func NewLegacySLOsClient(c *CallPropertiesCreator) *LegacySLOsClient {
	return &LegacySLOsClient{callPropertiesCreator: c}
}
