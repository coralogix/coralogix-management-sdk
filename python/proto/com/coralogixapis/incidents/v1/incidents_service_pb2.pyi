from com.coralogixapis.incidents.v1 import assignee_pb2 as _assignee_pb2
from com.coralogixapis.incidents.v1 import incident_pb2 as _incident_pb2
from com.coralogixapis.incidents.v1.incident_event import incident_event_pb2 as _incident_event_pb2
from com.coralogixapis.incidents.v1 import incident_query_pb2 as _incident_query_pb2
from com.coralogixapis.incidents.v1 import incident_query_filter_pb2 as _incident_query_filter_pb2
from com.coralogixapis.incidents.v1 import incident_query_filters_values_pb2 as _incident_query_filters_values_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor
AUDIT_LOG_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
audit_log_description: _descriptor.FieldDescriptor

class AuditLogDescription(_message.Message):
    __slots__ = ("description",)
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    description: str
    def __init__(self, description: _Optional[str] = ...) -> None: ...

class GetIncidentRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetIncidentResponse(_message.Message):
    __slots__ = ("incident",)
    INCIDENT_FIELD_NUMBER: _ClassVar[int]
    incident: _incident_pb2.Incident
    def __init__(self, incident: _Optional[_Union[_incident_pb2.Incident, _Mapping]] = ...) -> None: ...

class ListIncidentsResponse(_message.Message):
    __slots__ = ("incidents", "pagination")
    INCIDENTS_FIELD_NUMBER: _ClassVar[int]
    PAGINATION_FIELD_NUMBER: _ClassVar[int]
    incidents: _containers.RepeatedCompositeFieldContainer[_incident_pb2.Incident]
    pagination: PaginationResponse
    def __init__(self, incidents: _Optional[_Iterable[_Union[_incident_pb2.Incident, _Mapping]]] = ..., pagination: _Optional[_Union[PaginationResponse, _Mapping]] = ...) -> None: ...

class ListIncidentsRequest(_message.Message):
    __slots__ = ("filter", "pagination", "order_bys")
    FILTER_FIELD_NUMBER: _ClassVar[int]
    PAGINATION_FIELD_NUMBER: _ClassVar[int]
    ORDER_BYS_FIELD_NUMBER: _ClassVar[int]
    filter: _incident_query_filter_pb2.IncidentQueryFilter
    pagination: PaginationRequest
    order_bys: _containers.RepeatedCompositeFieldContainer[_incident_query_pb2.OrderBy]
    def __init__(self, filter: _Optional[_Union[_incident_query_filter_pb2.IncidentQueryFilter, _Mapping]] = ..., pagination: _Optional[_Union[PaginationRequest, _Mapping]] = ..., order_bys: _Optional[_Iterable[_Union[_incident_query_pb2.OrderBy, _Mapping]]] = ...) -> None: ...

class BatchGetIncidentRequest(_message.Message):
    __slots__ = ("ids",)
    IDS_FIELD_NUMBER: _ClassVar[int]
    ids: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, ids: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class BatchGetIncidentResponse(_message.Message):
    __slots__ = ("incidents", "not_found_ids")
    class IncidentsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: _incident_pb2.Incident
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[_incident_pb2.Incident, _Mapping]] = ...) -> None: ...
    INCIDENTS_FIELD_NUMBER: _ClassVar[int]
    NOT_FOUND_IDS_FIELD_NUMBER: _ClassVar[int]
    incidents: _containers.MessageMap[str, _incident_pb2.Incident]
    not_found_ids: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, incidents: _Optional[_Mapping[str, _incident_pb2.Incident]] = ..., not_found_ids: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class DeleteIncidentRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class PaginationRequest(_message.Message):
    __slots__ = ("page_size", "page_token")
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    PAGE_TOKEN_FIELD_NUMBER: _ClassVar[int]
    page_size: _wrappers_pb2.UInt32Value
    page_token: _wrappers_pb2.StringValue
    def __init__(self, page_size: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., page_token: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class PaginationResponse(_message.Message):
    __slots__ = ("total_size", "next_page_token")
    TOTAL_SIZE_FIELD_NUMBER: _ClassVar[int]
    NEXT_PAGE_TOKEN_FIELD_NUMBER: _ClassVar[int]
    total_size: _wrappers_pb2.UInt32Value
    next_page_token: _wrappers_pb2.StringValue
    def __init__(self, total_size: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., next_page_token: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class ListIncidentAggregationsRequest(_message.Message):
    __slots__ = ("filter", "group_bys", "pagination")
    FILTER_FIELD_NUMBER: _ClassVar[int]
    GROUP_BYS_FIELD_NUMBER: _ClassVar[int]
    PAGINATION_FIELD_NUMBER: _ClassVar[int]
    filter: _incident_query_filter_pb2.IncidentQueryFilter
    group_bys: _containers.RepeatedCompositeFieldContainer[_incident_query_pb2.GroupBy]
    pagination: PaginationRequest
    def __init__(self, filter: _Optional[_Union[_incident_query_filter_pb2.IncidentQueryFilter, _Mapping]] = ..., group_bys: _Optional[_Iterable[_Union[_incident_query_pb2.GroupBy, _Mapping]]] = ..., pagination: _Optional[_Union[PaginationRequest, _Mapping]] = ...) -> None: ...

class ListIncidentAggregationsResponse(_message.Message):
    __slots__ = ("incident_aggs", "pagination")
    INCIDENT_AGGS_FIELD_NUMBER: _ClassVar[int]
    PAGINATION_FIELD_NUMBER: _ClassVar[int]
    incident_aggs: _containers.RepeatedCompositeFieldContainer[_incident_pb2.IncidentAggregation]
    pagination: PaginationResponse
    def __init__(self, incident_aggs: _Optional[_Iterable[_Union[_incident_pb2.IncidentAggregation, _Mapping]]] = ..., pagination: _Optional[_Union[PaginationResponse, _Mapping]] = ...) -> None: ...

class GetFilterValuesRequest(_message.Message):
    __slots__ = ("filter",)
    FILTER_FIELD_NUMBER: _ClassVar[int]
    filter: _incident_query_filter_pb2.IncidentQueryFilter
    def __init__(self, filter: _Optional[_Union[_incident_query_filter_pb2.IncidentQueryFilter, _Mapping]] = ...) -> None: ...

class GetFilterValuesResponse(_message.Message):
    __slots__ = ("filters_values",)
    FILTERS_VALUES_FIELD_NUMBER: _ClassVar[int]
    filters_values: _incident_query_filters_values_pb2.IncidentQueryFiltersValues
    def __init__(self, filters_values: _Optional[_Union[_incident_query_filters_values_pb2.IncidentQueryFiltersValues, _Mapping]] = ...) -> None: ...

class GetIncidentEventsRequest(_message.Message):
    __slots__ = ("incident_id",)
    INCIDENT_ID_FIELD_NUMBER: _ClassVar[int]
    incident_id: _wrappers_pb2.StringValue
    def __init__(self, incident_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetIncidentEventsResponse(_message.Message):
    __slots__ = ("incident_events",)
    INCIDENT_EVENTS_FIELD_NUMBER: _ClassVar[int]
    incident_events: _containers.RepeatedCompositeFieldContainer[_incident_event_pb2.IncidentEvent]
    def __init__(self, incident_events: _Optional[_Iterable[_Union[_incident_event_pb2.IncidentEvent, _Mapping]]] = ...) -> None: ...

class AssignIncidentsRequest(_message.Message):
    __slots__ = ("incident_ids", "assigned_to")
    INCIDENT_IDS_FIELD_NUMBER: _ClassVar[int]
    ASSIGNED_TO_FIELD_NUMBER: _ClassVar[int]
    incident_ids: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    assigned_to: _assignee_pb2.UserDetails
    def __init__(self, incident_ids: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., assigned_to: _Optional[_Union[_assignee_pb2.UserDetails, _Mapping]] = ...) -> None: ...

class UnassignIncidentsRequest(_message.Message):
    __slots__ = ("incident_ids",)
    INCIDENT_IDS_FIELD_NUMBER: _ClassVar[int]
    incident_ids: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, incident_ids: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class UnassignIncidentsResponse(_message.Message):
    __slots__ = ("incidents",)
    INCIDENTS_FIELD_NUMBER: _ClassVar[int]
    incidents: _containers.RepeatedCompositeFieldContainer[_incident_pb2.Incident]
    def __init__(self, incidents: _Optional[_Iterable[_Union[_incident_pb2.Incident, _Mapping]]] = ...) -> None: ...

class AssignIncidentsResponse(_message.Message):
    __slots__ = ("incidents",)
    INCIDENTS_FIELD_NUMBER: _ClassVar[int]
    incidents: _containers.RepeatedCompositeFieldContainer[_incident_pb2.Incident]
    def __init__(self, incidents: _Optional[_Iterable[_Union[_incident_pb2.Incident, _Mapping]]] = ...) -> None: ...

class AcknowledgeIncidentsRequest(_message.Message):
    __slots__ = ("incident_ids",)
    INCIDENT_IDS_FIELD_NUMBER: _ClassVar[int]
    incident_ids: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, incident_ids: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class AcknowledgeIncidentsResponse(_message.Message):
    __slots__ = ("incidents",)
    INCIDENTS_FIELD_NUMBER: _ClassVar[int]
    incidents: _containers.RepeatedCompositeFieldContainer[_incident_pb2.Incident]
    def __init__(self, incidents: _Optional[_Iterable[_Union[_incident_pb2.Incident, _Mapping]]] = ...) -> None: ...

class CloseIncidentsRequest(_message.Message):
    __slots__ = ("incident_ids",)
    INCIDENT_IDS_FIELD_NUMBER: _ClassVar[int]
    incident_ids: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, incident_ids: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class CloseIncidentsResponse(_message.Message):
    __slots__ = ("incidents",)
    INCIDENTS_FIELD_NUMBER: _ClassVar[int]
    incidents: _containers.RepeatedCompositeFieldContainer[_incident_pb2.Incident]
    def __init__(self, incidents: _Optional[_Iterable[_Union[_incident_pb2.Incident, _Mapping]]] = ...) -> None: ...

class ResolveIncidentsRequest(_message.Message):
    __slots__ = ("incident_ids",)
    INCIDENT_IDS_FIELD_NUMBER: _ClassVar[int]
    incident_ids: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, incident_ids: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class ResolveIncidentsResponse(_message.Message):
    __slots__ = ("incidents",)
    INCIDENTS_FIELD_NUMBER: _ClassVar[int]
    incidents: _containers.RepeatedCompositeFieldContainer[_incident_pb2.Incident]
    def __init__(self, incidents: _Optional[_Iterable[_Union[_incident_pb2.Incident, _Mapping]]] = ...) -> None: ...

class GetIncidentUsingCorrelationKeyRequest(_message.Message):
    __slots__ = ("correlation_key", "incident_point_in_time")
    CORRELATION_KEY_FIELD_NUMBER: _ClassVar[int]
    INCIDENT_POINT_IN_TIME_FIELD_NUMBER: _ClassVar[int]
    correlation_key: _wrappers_pb2.StringValue
    incident_point_in_time: _timestamp_pb2.Timestamp
    def __init__(self, correlation_key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., incident_point_in_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class GetIncidentUsingCorrelationKeyResponse(_message.Message):
    __slots__ = ("incident",)
    INCIDENT_FIELD_NUMBER: _ClassVar[int]
    incident: _incident_pb2.Incident
    def __init__(self, incident: _Optional[_Union[_incident_pb2.Incident, _Mapping]] = ...) -> None: ...
