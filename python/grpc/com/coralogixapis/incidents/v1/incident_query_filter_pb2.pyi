from com.coralogixapis.incidents.v1 import incident_severity_pb2 as _incident_severity_pb2
from com.coralogixapis.incidents.v1 import incident_query_pb2 as _incident_query_pb2
from com.coralogixapis.incidents.v1 import incident_status_pb2 as _incident_status_pb2
from com.coralogixapis.incidents.v1 import incident_state_pb2 as _incident_state_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogixapis.incidents.v1 import meta_label_pb2 as _meta_label_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class IncidentQueryFilter(_message.Message):
    __slots__ = ("assignee", "status", "state", "severity", "contextual_labels", "start_time", "end_time", "search_query", "application_name", "subsystem_name", "is_muted", "created_at_range", "incident_duration_range", "meta_labels")
    class ContextualLabelsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: ContextualLabelValues
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[ContextualLabelValues, _Mapping]] = ...) -> None: ...
    ASSIGNEE_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    STATE_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    CONTEXTUAL_LABELS_FIELD_NUMBER: _ClassVar[int]
    START_TIME_FIELD_NUMBER: _ClassVar[int]
    END_TIME_FIELD_NUMBER: _ClassVar[int]
    SEARCH_QUERY_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_NAME_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_NAME_FIELD_NUMBER: _ClassVar[int]
    IS_MUTED_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_RANGE_FIELD_NUMBER: _ClassVar[int]
    INCIDENT_DURATION_RANGE_FIELD_NUMBER: _ClassVar[int]
    META_LABELS_FIELD_NUMBER: _ClassVar[int]
    assignee: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    status: _containers.RepeatedScalarFieldContainer[_incident_status_pb2.IncidentStatus]
    state: _containers.RepeatedScalarFieldContainer[_incident_state_pb2.IncidentState]
    severity: _containers.RepeatedScalarFieldContainer[_incident_severity_pb2.IncidentSeverity]
    contextual_labels: _containers.MessageMap[str, ContextualLabelValues]
    start_time: _timestamp_pb2.Timestamp
    end_time: _timestamp_pb2.Timestamp
    search_query: _incident_query_pb2.IncidentSearchQuery
    application_name: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    subsystem_name: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    is_muted: _wrappers_pb2.BoolValue
    created_at_range: TimeRange
    incident_duration_range: TimeRange
    meta_labels: _containers.RepeatedCompositeFieldContainer[_meta_label_pb2.MetaLabel]
    def __init__(self, assignee: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., status: _Optional[_Iterable[_Union[_incident_status_pb2.IncidentStatus, str]]] = ..., state: _Optional[_Iterable[_Union[_incident_state_pb2.IncidentState, str]]] = ..., severity: _Optional[_Iterable[_Union[_incident_severity_pb2.IncidentSeverity, str]]] = ..., contextual_labels: _Optional[_Mapping[str, ContextualLabelValues]] = ..., start_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., end_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., search_query: _Optional[_Union[_incident_query_pb2.IncidentSearchQuery, _Mapping]] = ..., application_name: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystem_name: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., is_muted: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., created_at_range: _Optional[_Union[TimeRange, _Mapping]] = ..., incident_duration_range: _Optional[_Union[TimeRange, _Mapping]] = ..., meta_labels: _Optional[_Iterable[_Union[_meta_label_pb2.MetaLabel, _Mapping]]] = ...) -> None: ...

class TimeRange(_message.Message):
    __slots__ = ("start_time", "end_time")
    START_TIME_FIELD_NUMBER: _ClassVar[int]
    END_TIME_FIELD_NUMBER: _ClassVar[int]
    start_time: _timestamp_pb2.Timestamp
    end_time: _timestamp_pb2.Timestamp
    def __init__(self, start_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., end_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class ContextualLabelValues(_message.Message):
    __slots__ = ("contextual_label_values",)
    CONTEXTUAL_LABEL_VALUES_FIELD_NUMBER: _ClassVar[int]
    contextual_label_values: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, contextual_label_values: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...
