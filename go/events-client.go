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

	events "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/events/v3"
)

const eventsFeatureGroupID = "events"

// RPC methods
const (
	GetEventRPC            = events.EventsService_GetEvent_FullMethodName
	BatchGetEventsRPC      = events.EventsService_BatchGetEvent_FullMethodName
	GetEventsStatisticsRPC = events.EventsService_GetEventsStatistics_FullMethodName
	ListEventsRPC          = events.EventsService_ListEvents_FullMethodName
	ListEventsCountRPC     = events.EventsService_ListEventsCount_FullMethodName
)

// GetEventRequest is the request for the GetEvent RPC
type GetEventRequest = events.GetEventRequest

// GetEventResponse is the response for the GetEvent RPC
type GetEventResponse = events.GetEventResponse

// BatchGetEventRequest is the request for the BatchGetEvent RPC
type BatchGetEventRequest = events.BatchGetEventRequest

// BatchGetEventResponse is the response for the BatchGetEvent RPC
type BatchGetEventResponse = events.BatchGetEventResponse

// GetEventsStatisticsRequest is the request for the GetEventsStatistics RPC
type GetEventsStatisticsRequest = events.GetEventsStatisticsRequest

// GetEventsStatisticsResponse is the response for the GetEventsStatistics RPC
type GetEventsStatisticsResponse = events.GetEventsStatisticsResponse

// ListEventsRequest is the request for the ListEvents RPC
type ListEventsRequest = events.ListEventsRequest

// ListEventsResponse is the response for the ListEvents RPC
type ListEventsResponse = events.ListEventsResponse

// ListEventsCountRequest is the request for the ListEventsCount RPC
type ListEventsCountRequest = events.ListEventsCountRequest

// ListEventsCountResponse is the response for the ListEventsCount RPC
type ListEventsCountResponse = events.ListEventsCountResponse

// OrderBy is the request for the OrderBy RPC
type OrderBy = events.OrderBy

// EventsPaginationRequest is the request for the Pagination RPC
type EventsPaginationRequest = events.PaginationRequest

// EventsQueryFilter is the request for the EventsQueryFilter RPC
type EventsQueryFilter = events.EventsQueryFilter

// EventsQueryFilterTimestampRange is the request for the EventsQueryFilterTimestampRange RPC
type EventsQueryFilterTimestampRange = events.TimestampRange

// EventsFilter is the request for the EventsFilter RPC
type EventsFilter = events.EventsFilter

// EventFilters is the request for the EventFilters RPC
type EventFilters = events.Filters

// EventFilterAndPathValues is the request for the EventFilterAndPathValues RPC
type EventFilterAndPathValues = events.FilterPathAndValues

// EventMultipleFilterAndPathValues is the request for the EventMultipleFilterAndPathValues RPC
type EventMultipleFilterAndPathValues = events.FilterPathAndValues_MultipleValues

// EventFilterAndPathValuesFilters is the request for the EventFilterAndPathValuesFilter RPC
type EventFilterAndPathValuesFilters = events.FilterPathAndValues_Filters

// EventMultipleValues is the request for the EventMultipleValues RPC
type EventMultipleValues = events.MultipleValues

// These are the constants for the FilterMatcher enum
const (
	EventsFilterMatcherTextOrUnspecified = events.FilterMatcher_FILTER_MATCHER_TEXT_OR_UNSPECIFIED
	EventsFilterMatcherRegexp            = events.FilterMatcher_FILTER_MATCHER_REGEXP
)

// These are the constants for the FilterOperator enum
const (
	EventsFilterOperatorAndOrUnspecified = events.FilterOperator_FILTER_OPERATOR_AND_OR_UNSPECIFIED
	EventsFilterOperatorOr               = events.FilterOperator_FILTER_OPERATOR_OR
)

// These are the constants for the OrderByField enum
const (
	EventsOrderByFieldsUnspecified = events.OrderByFields_ORDER_BY_FIELDS_UNSPECIFIED
	EventsOrderByFieldsTimestamp   = events.OrderByFields_ORDER_BY_FIELDS_TIMESTAMP
)

// These are the constants for the OrderByDirection enum
const (
	EventsOrderByDirectionUnspecified = events.OrderByDirection_ORDER_BY_DIRECTION_UNSPECIFIED
	EventsOrderByDirectionAsc         = events.OrderByDirection_ORDER_BY_DIRECTION_ASC
	EventsOrderByDirectionDesc        = events.OrderByDirection_ORDER_BY_DIRECTION_DESC
)

// EventsClient is a client for the Coralogix Events API
type EventsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Get gets a single event by ID
func (e EventsClient) Get(ctx context.Context, req *GetEventRequest) (*GetEventResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := events.NewEventsServiceClient(conn)

	response, err := client.GetEvent(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetEventRPC, eventsFeatureGroupID)
	}
	return response, nil
}

// BatchGet gets multiple events based on provided criteria
func (e EventsClient) BatchGet(ctx context.Context, req *BatchGetEventRequest) (*BatchGetEventResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := events.NewEventsServiceClient(conn)

	response, err := client.BatchGetEvent(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, BatchGetEventsRPC, eventsFeatureGroupID)
	}
	return response, nil
}

// GetStatistics gets statistics about events
func (e EventsClient) GetStatistics(ctx context.Context, req *GetEventsStatisticsRequest) (*GetEventsStatisticsResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := events.NewEventsServiceClient(conn)

	response, err := client.GetEventsStatistics(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetEventsStatisticsRPC, eventsFeatureGroupID)
	}
	return response, nil
}

// List lists events based on the provided filter criteria
func (e EventsClient) List(ctx context.Context, req *ListEventsRequest) (*ListEventsResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := events.NewEventsServiceClient(conn)

	response, err := client.ListEvents(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ListEventsRPC, eventsFeatureGroupID)
	}
	return response, nil
}

// ListCount gets the count of events matching the filter criteria
func (e EventsClient) ListCount(ctx context.Context, req *ListEventsCountRequest) (*ListEventsCountResponse, error) {
	callProperties, err := e.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := events.NewEventsServiceClient(conn)

	response, err := client.ListEventsCount(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ListEventsCountRPC, eventsFeatureGroupID)
	}
	return response, nil
}

// NewEventsClient creates a new events client
func NewEventsClient(c *CallPropertiesCreator) *EventsClient {
	return &EventsClient{callPropertiesCreator: c}
}
