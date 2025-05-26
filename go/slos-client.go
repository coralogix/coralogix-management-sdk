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

	slos "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/slo/v1"
)

// Slo is an SLO.
type Slo = slos.Slo

// CreateServiceSloRequest is a request to create an SLO.
type CreateServiceSloRequest = slos.CreateSloRequest

// GetServiceSloRequest is a request to get an SLO.
type GetServiceSloRequest = slos.GetSloRequest

// ReplaceServiceSloRequest is a request to replace an SLO.
type ReplaceServiceSloRequest = slos.ReplaceSloRequest

// DeleteServiceSloRequest is a request to delete an SLO.
type DeleteServiceSloRequest = slos.DeleteSloRequest

// SloFilter is an SLO filter.
type SloFilter = slos.SloFilter

// SloFilterField is an SLO filter field.
type SloFilterField = slos.SloFilterField

// SloConstantFilterField is an SLO constant filter.
type SloConstantFilterField = slos.SloFilterField_ConstFilter

// SloConstantFilterFieldEnum is an SLO constant filter field enum.
type SloConstantFilterFieldEnum = slos.SloConstantFilterField

// SloLabelNameFilterField is an SLO label name filter.
type SloLabelNameFilterField = slos.SloFilterField_LabelName

// SloFilterPredicate is an SLO filter predicate.
type SloFilterPredicate = slos.SloFilterPredicate

// SloFilterPredicateIs is an SLO filter predicate is.
type SloFilterPredicateIs = slos.SloFilterPredicate_Is

// IsSloFilterPredicate is an SLO filter is predicate.
type IsSloFilterPredicate = slos.IsFilterPredicate

// SloRequestBasedMetricSli is a metric SLI.
type SloRequestBasedMetricSli = slos.Slo_RequestBasedMetricSli

// SloWindowBasedMetricSli is a metric SLI.
type SloWindowBasedMetricSli = slos.Slo_WindowBasedMetricSli

// WindowBasedMetricSli is a window-based metric SLI.
type WindowBasedMetricSli = slos.WindowBasedMetricSli

// RequestBasedMetricSli is a metric SLI.
type RequestBasedMetricSli = slos.RequestBasedMetricSli

// Metric is a metric.
type Metric = slos.Metric

// SloTimeframe is an SLO time frame.
type SloTimeframe = slos.Slo_SloTimeFrame

// SloTimeframeEnum is an enum type for SLO time frames.
type SloTimeframeEnum = slos.SloTimeFrame

// ListSlosRequest is a request to list SLOs.
type ListSlosRequest = slos.ListSlosRequest

// ListSlosResponse is a response to a list SLOs request.
type ListSlosResponse = slos.ListSlosResponse

// SloGrouping is a type that represents the SLO grouping.
type SloGrouping = slos.Grouping

// SloFilters is a type that allows to filter SLOs.
type SloFilters = slos.SloFilters

// SloWindow is a type that represents the SLO window.
type SloWindow = slos.WindowSloWindow

// SloComparisonOperator is a type that represents the SLO comparison operator.
type SloComparisonOperator = slos.ComparisonOperator

// SloTimeframe variants
const (
	SloTimeframeUnspecified = slos.SloTimeFrame_SLO_TIME_FRAME_UNSPECIFIED
	SloTimeframe7Days       = slos.SloTimeFrame_SLO_TIME_FRAME_7_DAYS
	SloTimeframe14Days      = slos.SloTimeFrame_SLO_TIME_FRAME_14_DAYS
	SloTimeframe21Days      = slos.SloTimeFrame_SLO_TIME_FRAME_21_DAYS
	SloTimeframe28Days      = slos.SloTimeFrame_SLO_TIME_FRAME_28_DAYS
	SloTimeframe90Days      = slos.SloTimeFrame_SLO_TIME_FRAME_90_DAYS
)

// SloWindow variants
const (
	SloWindowUnspecified = slos.WindowSloWindow_WINDOW_SLO_WINDOW_UNSPECIFIED
	SloWindow1Minute     = slos.WindowSloWindow_WINDOW_SLO_WINDOW_1_MINUTE
	SloWindow5Minutes    = slos.WindowSloWindow_WINDOW_SLO_WINDOW_5_MINUTES
)

// SloFilterField variants
const (
	SloConstantFilterFieldUnspecified = slos.SloConstantFilterField_SLO_CONST_FILTER_FIELD_UNSPECIFIED
	SloConstantFilterFieldUserName    = slos.SloConstantFilterField_SLO_CONST_FILTER_FIELD_USER_NAME
)

// SloComparisonOperator variants
const (
	SloComparisonOperatorUnspecified         = slos.ComparisonOperator_COMPARISON_OPERATOR_UNSPECIFIED
	SloComparisonOperatorGreaterThan         = slos.ComparisonOperator_COMPARISON_OPERATOR_GREATER_THAN
	SloComparisonOperatorLessThan            = slos.ComparisonOperator_COMPARISON_OPERATOR_LESS_THAN
	SloComparisonOperatorGreaterThanOrEquals = slos.ComparisonOperator_COMPARISON_OPERATOR_GREATER_THAN_OR_EQUALS
	SloComparisonOperatorLessThanOrEquals    = slos.ComparisonOperator_COMPARISON_OPERATOR_LESS_THAN_OR_EQUALS
)

// SLO RPC names
const (
	SloCreateRPC   = slos.SlosService_CreateSlo_FullMethodName
	SloGetRPC      = slos.SlosService_GetSlo_FullMethodName
	SloReplaceRPC  = slos.SlosService_ReplaceSlo_FullMethodName
	SloDeleteRPC   = slos.SlosService_DeleteSlo_FullMethodName
	SloBatchGetRPC = slos.SlosService_BatchGetSlos_FullMethodName
	SlosListRPC    = slos.SlosService_ListSlos_FullMethodName
)

// SloStatus values.
const (
	SloStatusUnspecified = slos.SloStatus_SLO_STATUS_UNSPECIFIED
	SloStatusOk          = slos.SloStatus_SLO_STATUS_OK
	SloStatusBreached    = slos.SloStatus_SLO_STATUS_BREACHED
	SloStatusCritical    = slos.SloStatus_SLO_STATUS_CRITICAL
	SloStatusWarning     = slos.SloStatus_SLO_STATUS_WARNING
	SloStatusPending     = slos.SloStatus_SLO_STATUS_PENDING
)

const infraMonitoringFeatureGroupID = "infra-monitoring"

// SLOsClient is a client for the Coralogix SLOs API.
type SLOsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new SLO.
func (c SLOsClient) Create(ctx context.Context, req *slos.CreateSloRequest) (*slos.CreateSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewSlosServiceClient(conn)

	response, err := client.CreateSlo(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, SloCreateRPC, infraMonitoringFeatureGroupID)
	}
	return response, nil
}

// Get gets the specified SLO.
func (c SLOsClient) Get(ctx context.Context, req *slos.GetSloRequest) (*slos.GetSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewSlosServiceClient(conn)

	response, err := client.GetSlo(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, SloGetRPC, infraMonitoringFeatureGroupID)
	}
	return response, nil
}

// Update updates the specified SLO.
func (c SLOsClient) Update(ctx context.Context, req *slos.ReplaceSloRequest) (*slos.ReplaceSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewSlosServiceClient(conn)

	response, err := client.ReplaceSlo(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, SloReplaceRPC, infraMonitoringFeatureGroupID)
	}
	return response, nil
}

// Delete deletes the specified SLO.
func (c SLOsClient) Delete(ctx context.Context, req *slos.DeleteSloRequest) (*slos.DeleteSloResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewSlosServiceClient(conn)

	response, err := client.DeleteSlo(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, SloDeleteRPC, infraMonitoringFeatureGroupID)
	}
	return response, nil
}

// GetBulk gets multiple SLOs in a single call.
func (c SLOsClient) GetBulk(ctx context.Context, req *slos.BatchGetSlosRequest) (*slos.BatchGetSlosResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewSlosServiceClient(conn)

	response, err := client.BatchGetSlos(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, SloBatchGetRPC, infraMonitoringFeatureGroupID)
	}
	return response, nil
}

// List lists all service SLOs.
func (c SLOsClient) List(ctx context.Context, req *slos.ListSlosRequest) (*slos.ListSlosResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := slos.NewSlosServiceClient(conn)

	response, err := client.ListSlos(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, SlosListRPC, infraMonitoringFeatureGroupID)
	}
	return response, nil
}

// NewSLOsClient creates a new SLOs client.
func NewSLOsClient(c *CallPropertiesCreator) *SLOsClient {
	return &SLOsClient{callPropertiesCreator: c}
}
