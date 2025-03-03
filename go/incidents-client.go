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
package cxsdk

import (
	"context"

	incidents "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/incidents/v1"
)

// GetIncidentRequest is a type for a request.
type GetIncidentRequest = incidents.GetIncidentRequest

// GetIncidentResponse is a type for a response.
type GetIncidentResponse = incidents.GetIncidentResponse

// BatchGetIncidentRequest is a type for a request.
type BatchGetIncidentRequest = incidents.BatchGetIncidentRequest

// BatchGetIncidentResponse is a type for a response.
type BatchGetIncidentResponse = incidents.BatchGetIncidentResponse

// ListIncidentsRequest is a type for a request.
type ListIncidentsRequest = incidents.ListIncidentsRequest

// ListIncidentsResponse is a type for a response.
type ListIncidentsResponse = incidents.ListIncidentsResponse

// ListIncidentAggregationsRequest is a type for a request.
type ListIncidentAggregationsRequest = incidents.ListIncidentAggregationsRequest

// ListIncidentAggregationsResponse is a type for a response.
type ListIncidentAggregationsResponse = incidents.ListIncidentAggregationsResponse

// GetFilterValuesRequest is a type for a request.
type GetFilterValuesRequest = incidents.GetFilterValuesRequest

// GetFilterValuesResponse is a type for a response.
type GetFilterValuesResponse = incidents.GetFilterValuesResponse

// AssignIncidentsRequest is a type for a request.
type AssignIncidentsRequest = incidents.AssignIncidentsRequest

// AssignIncidentsResponse is a type for a response.
type AssignIncidentsResponse = incidents.AssignIncidentsResponse

// UnassignIncidentsRequest is a type for a request.
type UnassignIncidentsRequest = incidents.UnassignIncidentsRequest

// UnassignIncidentsResponse is a type for a response.
type UnassignIncidentsResponse = incidents.UnassignIncidentsResponse

// AcknowledgeIncidentsRequest is a type for a request.
type AcknowledgeIncidentsRequest = incidents.AcknowledgeIncidentsRequest

// AcknowledgeIncidentsResponse is a type for a response.
type AcknowledgeIncidentsResponse = incidents.AcknowledgeIncidentsResponse

// CloseIncidentsRequest is a type for a request.
type CloseIncidentsRequest = incidents.CloseIncidentsRequest

// CloseIncidentsResponse is a type for a response.
type CloseIncidentsResponse = incidents.CloseIncidentsResponse

// GetIncidentEventsRequest is a type for a request.
type GetIncidentEventsRequest = incidents.GetIncidentEventsRequest

// GetIncidentEventsResponse is a type for a response.
type GetIncidentEventsResponse = incidents.GetIncidentEventsResponse

// ResolveIncidentsRequest is a type for a request.
type ResolveIncidentsRequest = incidents.ResolveIncidentsRequest

// ResolveIncidentsResponse is a type for a response.
type ResolveIncidentsResponse = incidents.ResolveIncidentsResponse

// GetIncidentUsingCorrelationKeyRequest is a type for a request.
type GetIncidentUsingCorrelationKeyRequest = incidents.GetIncidentUsingCorrelationKeyRequest

// GetIncidentUsingCorrelationKeyResponse is a type for a response.
type GetIncidentUsingCorrelationKeyResponse = incidents.GetIncidentUsingCorrelationKeyResponse

// ListIncidentEventsRequest is a type for a request.
type ListIncidentEventsRequest = incidents.ListIncidentEventsRequest

// ListIncidentEventsResponse is a type for a response.
type ListIncidentEventsResponse = incidents.ListIncidentEventsResponse

// ListIncidentEventsTotalCountRequest is a type for a request.
type ListIncidentEventsTotalCountRequest = incidents.ListIncidentEventsTotalCountRequest

// ListIncidentEventsTotalCountResponse is a type for a response.
type ListIncidentEventsTotalCountResponse = incidents.ListIncidentEventsTotalCountResponse

// ListIncidentEventsFilterValuesRequest is a type for a request.
type ListIncidentEventsFilterValuesRequest = incidents.ListIncidentEventsFilterValuesRequest

// ListIncidentEventsFilterValuesResponse is a type for a response.
type ListIncidentEventsFilterValuesResponse = incidents.ListIncidentEventsFilterValuesResponse

// IncidentEventQueryFilter is a type for a filter.
type IncidentEventQueryFilter = incidents.IncidentEventQueryFilter

// IncidentEventTimeRange is a type for a time range.
type IncidentEventTimeRange = incidents.TimeRange

// IncidentEventLabelsFilter is a type for a labels filter.
type IncidentEventLabelsFilter = incidents.LabelsFilter

// ContextualLabelValues is a type for a contextual label values.
type ContextualLabelValues = incidents.ContextualLabelValues

// IncidentEventMetaLabel is a type for a incident event meta label.
type IncidentEventMetaLabel = incidents.MetaLabel

// IncidentEventFilterOperator is a type for a incident event filter operator.
type IncidentEventFilterOperator = incidents.FilterOperator

// IncidentQueryFilter is a type for a incident query filter.
type IncidentQueryFilter = incidents.IncidentQueryFilter

// IncidentEventOrderByFieldType is a type for a incident event order by field.
type IncidentEventOrderByFieldType = incidents.IncidentEventOrderByFieldType

// These constants are the possible values for IncidentEventOrderByFieldType.
const (
	IncidentEventOrderByFieldTypeTimestampOrUnspecified = incidents.IncidentEventOrderByFieldType_INCIDENT_EVENT_ORDER_BY_FIELD_TYPE_TIMESTAMP_OR_UNSPECIFIED
)

// ListIncidentEventOrderBy is a type for a incident event order by.
type ListIncidentEventOrderBy = incidents.ListIncidentEventRequestOrderBy

// PaginationRequest is a type for a pagination request.
type PaginationRequest = incidents.PaginationRequest

// These constants are the possible values for ListIncidentEventOrderBy.
const (
	IncidentEventOrderByFieldName        = incidents.OrderByFields_ORDER_BY_FIELDS_NAME
	IncidentEventOrderByFieldSeverity    = incidents.OrderByFields_ORDER_BY_FIELDS_SEVERITY
	IncidentEventOrderByFieldID          = incidents.OrderByFields_ORDER_BY_FIELDS_ID
	IncidentEventOrderByFieldClosedTime  = incidents.OrderByFields_ORDER_BY_FIELDS_CLOSED_TIME
	IncidentEventOrderByFieldCreatedTime = incidents.OrderByFields_ORDER_BY_FIELDS_CREATED_TIME
	IncidentEventOrderByFieldUnspecified = incidents.OrderByFields_ORDER_BY_FIELDS_UNSPECIFIED
)

// These constants are the possible values for IncidentSeverity.
const (
	IncidentSeverityCritical    = incidents.IncidentSeverity_INCIDENT_SEVERITY_CRITICAL
	IncidentSeverityInfo        = incidents.IncidentSeverity_INCIDENT_SEVERITY_INFO
	IncidentSeverityLow         = incidents.IncidentSeverity_INCIDENT_SEVERITY_LOW
	IncidentSeverityUnspecified = incidents.IncidentSeverity_INCIDENT_SEVERITY_UNSPECIFIED
	IncidentSeverityWarning     = incidents.IncidentSeverity_INCIDENT_SEVERITY_WARNING
	IncidentSeverityError       = incidents.IncidentSeverity_INCIDENT_SEVERITY_ERROR
)

// These constants are the possible values for IncidentStatus.
const (
	IncidentStatusUnspecified  = incidents.IncidentStatus_INCIDENT_STATUS_UNSPECIFIED
	IncidentStatusAcknowledged = incidents.IncidentStatus_INCIDENT_STATUS_ACKNOWLEDGED
	IncidentStatusTriggered    = incidents.IncidentStatus_INCIDENT_STATUS_RESOLVED
	IncidentStatusResolved     = incidents.IncidentStatus_INCIDENT_STATUS_RESOLVED
)

// These constants are the possible values for IncidentEventFilterOperator.
const (
	IncidentEventFilterOperatorOrOrUnspecified = incidents.FilterOperator_FILTER_OPERATOR_OR_OR_UNSPECIFIED
	IncidentEventFilterOperatorAndAnd          = incidents.FilterOperator_FILTER_OPERATOR_AND
)

// These constants are the possible values for IncidentEventOrderByDirection.
const (
	IncidentEventOrderByDirectionAsc  = incidents.OrderByDirection_ORDER_BY_DIRECTION_ASC
	IncidentEventOrderByDirectionDesc = incidents.OrderByDirection_ORDER_BY_DIRECTION_DESC
)

const incidentsFeatureGroupID = "incidents"

// RPC method names
const (
	GetIncidentRPC                    = incidents.IncidentsService_GetIncident_FullMethodName
	BatchGetIncidentRPC               = incidents.IncidentsService_BatchGetIncident_FullMethodName
	ListIncidentsRPC                  = incidents.IncidentsService_ListIncidents_FullMethodName
	ListIncidentAggregationsRPC       = incidents.IncidentsService_ListIncidentAggregations_FullMethodName
	GetFilterValuesRPC                = incidents.IncidentsService_GetFilterValues_FullMethodName
	AssignIncidentsRPC                = incidents.IncidentsService_AssignIncidents_FullMethodName
	UnassignIncidentsRPC              = incidents.IncidentsService_UnassignIncidents_FullMethodName
	AcknowledgeIncidentsRPC           = incidents.IncidentsService_AcknowledgeIncidents_FullMethodName
	CloseIncidentsRPC                 = incidents.IncidentsService_CloseIncidents_FullMethodName
	GetIncidentEventsRPC              = incidents.IncidentsService_GetIncidentEvents_FullMethodName
	ResolveIncidentsRPC               = incidents.IncidentsService_ResolveIncidents_FullMethodName
	GetIncidentUsingCorrelationKeyRPC = incidents.IncidentsService_GetIncidentUsingCorrelationKey_FullMethodName
	ListIncidentEventsRPC             = incidents.IncidentsService_ListIncidentEvents_FullMethodName
	ListIncidentEventsTotalCountRPC   = incidents.IncidentsService_ListIncidentEventsTotalCount_FullMethodName
	ListIncidentEventsFilterValuesRPC = incidents.IncidentsService_ListIncidentEventsFilterValues_FullMethodName
)

// IncidentsClient is a client for the Incidents service
type IncidentsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// GetIncident gets an incident by its ID
func (c IncidentsClient) GetIncident(ctx context.Context, req *GetIncidentRequest) (*incidents.GetIncidentResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := incidents.NewIncidentsServiceClient(conn)

	response, err := client.GetIncident(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetIncidentRPC, incidentsFeatureGroupID)
	}
	return response, nil
}

// BatchGetIncident gets incidents by their IDs
func (c IncidentsClient) BatchGetIncident(ctx context.Context, req *BatchGetIncidentRequest) (*BatchGetIncidentResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := incidents.NewIncidentsServiceClient(conn)

	response, err := client.BatchGetIncident(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, BatchGetIncidentRPC, incidentsFeatureGroupID)
	}
	return response, nil
}

// ListIncidents lists incidents
func (c IncidentsClient) ListIncidents(ctx context.Context, req *ListIncidentsRequest) (*ListIncidentsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := incidents.NewIncidentsServiceClient(conn)

	response, err := client.ListIncidents(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ListIncidentsRPC, incidentsFeatureGroupID)
	}
	return response, nil
}

// ListIncidentAggregations lists incident aggregations
func (c IncidentsClient) ListIncidentAggregations(ctx context.Context, req *ListIncidentAggregationsRequest) (*ListIncidentAggregationsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := incidents.NewIncidentsServiceClient(conn)

	response, err := client.ListIncidentAggregations(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ListIncidentAggregationsRPC, incidentsFeatureGroupID)
	}
	return response, nil
}

// GetFilterValues gets filter values
func (c IncidentsClient) GetFilterValues(ctx context.Context, req *GetFilterValuesRequest) (*GetFilterValuesResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := incidents.NewIncidentsServiceClient(conn)

	response, err := client.GetFilterValues(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetFilterValuesRPC, incidentsFeatureGroupID)
	}
	return response, nil
}

// AssignIncidents assigns incidents to a user
func (c IncidentsClient) AssignIncidents(ctx context.Context, req *AssignIncidentsRequest) (*AssignIncidentsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := incidents.NewIncidentsServiceClient(conn)

	response, err := client.AssignIncidents(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, AssignIncidentsRPC, incidentsFeatureGroupID)
	}
	return response, nil
}

// UnassignIncidents unassigns incidents from a user
func (c IncidentsClient) UnassignIncidents(ctx context.Context, req *UnassignIncidentsRequest) (*UnassignIncidentsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := incidents.NewIncidentsServiceClient(conn)

	response, err := client.UnassignIncidents(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, UnassignIncidentsRPC, incidentsFeatureGroupID)
	}
	return response, nil
}

// AcknowledgeIncidents acknowledges incidents
func (c IncidentsClient) AcknowledgeIncidents(ctx context.Context, req *AcknowledgeIncidentsRequest) (*AcknowledgeIncidentsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := incidents.NewIncidentsServiceClient(conn)

	response, err := client.AcknowledgeIncidents(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, AcknowledgeIncidentsRPC, incidentsFeatureGroupID)
	}
	return response, nil
}

// CloseIncidents closes incidents
func (c IncidentsClient) CloseIncidents(ctx context.Context, req *CloseIncidentsRequest) (*CloseIncidentsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := incidents.NewIncidentsServiceClient(conn)

	response, err := client.CloseIncidents(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, CloseIncidentsRPC, incidentsFeatureGroupID)
	}
	return response, nil
}

// GetIncidentEvents gets incident events
func (c IncidentsClient) GetIncidentEvents(ctx context.Context, req *GetIncidentEventsRequest) (*GetIncidentEventsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := incidents.NewIncidentsServiceClient(conn)

	response, err := client.GetIncidentEvents(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetIncidentEventsRPC, incidentsFeatureGroupID)
	}
	return response, nil
}

// ResolveIncidents resolves incidents
func (c IncidentsClient) ResolveIncidents(ctx context.Context, req *ResolveIncidentsRequest) (*ResolveIncidentsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := incidents.NewIncidentsServiceClient(conn)

	response, err := client.ResolveIncidents(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ResolveIncidentsRPC, incidentsFeatureGroupID)
	}
	return response, nil
}

// GetIncidentUsingCorrelationKey gets an incident by its correlation key
func (c IncidentsClient) GetIncidentUsingCorrelationKey(ctx context.Context, req *GetIncidentUsingCorrelationKeyRequest) (*GetIncidentUsingCorrelationKeyResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := incidents.NewIncidentsServiceClient(conn)

	response, err := client.GetIncidentUsingCorrelationKey(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetIncidentUsingCorrelationKeyRPC, incidentsFeatureGroupID)
	}
	return response, nil
}

// ListIncidentEvents lists incident events
func (c IncidentsClient) ListIncidentEvents(ctx context.Context, req *ListIncidentEventsRequest) (*ListIncidentEventsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := incidents.NewIncidentsServiceClient(conn)

	response, err := client.ListIncidentEvents(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ListIncidentEventsRPC, incidentsFeatureGroupID)
	}
	return response, nil
}

// ListIncidentEventsTotalCount lists the total count of incident events
func (c IncidentsClient) ListIncidentEventsTotalCount(ctx context.Context, req *ListIncidentEventsTotalCountRequest) (*ListIncidentEventsTotalCountResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := incidents.NewIncidentsServiceClient(conn)

	response, err := client.ListIncidentEventsTotalCount(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ListIncidentEventsTotalCountRPC, incidentsFeatureGroupID)
	}
	return response, nil
}

// ListIncidentEventsFilterValues lists the filter values for incident events
func (c IncidentsClient) ListIncidentEventsFilterValues(ctx context.Context, req *ListIncidentEventsFilterValuesRequest) (*ListIncidentEventsFilterValuesResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := incidents.NewIncidentsServiceClient(conn)

	response, err := client.ListIncidentEventsFilterValues(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ListIncidentEventsFilterValuesRPC, incidentsFeatureGroupID)
	}
	return response, nil
}

// NewIncidentsClient creates a new IncidentsClient
func NewIncidentsClient(creator *CallPropertiesCreator) *IncidentsClient {
	return &IncidentsClient{
		callPropertiesCreator: creator,
	}
}
