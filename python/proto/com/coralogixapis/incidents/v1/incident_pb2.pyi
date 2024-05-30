from com.coralogixapis.incidents.v1 import assignee_pb2 as _assignee_pb2
from com.coralogixapis.incidents.v1.incident_event import incident_event_pb2 as _incident_event_pb2
from com.coralogixapis.incidents.v1 import incident_severity_pb2 as _incident_severity_pb2
from com.coralogixapis.incidents.v1 import incident_state_pb2 as _incident_state_pb2
from com.coralogixapis.incidents.v1 import incident_status_pb2 as _incident_status_pb2
from com.coralogixapis.incidents.v1 import meta_label_pb2 as _meta_label_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class IncidentFields(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    INCIDENTS_FIELDS_UNSPECIFIED: _ClassVar[IncidentFields]
    INCIDENTS_FIELDS_ID: _ClassVar[IncidentFields]
    INCIDENTS_FIELDS_SEVERITY: _ClassVar[IncidentFields]
    INCIDENTS_FIELDS_NAME: _ClassVar[IncidentFields]
    INCIDENTS_FIELDS_CREATED_TIME: _ClassVar[IncidentFields]
    INCIDENTS_FIELDS_CLOSED_TIME: _ClassVar[IncidentFields]
    INCIDENTS_FIELDS_STATE: _ClassVar[IncidentFields]
    INCIDENTS_FIELDS_STATUS: _ClassVar[IncidentFields]
    INCIDENTS_FIELDS_LAST_STATE_UPDATE_TIME: _ClassVar[IncidentFields]
    INCIDENTS_FIELDS_APPLICATION_NAME: _ClassVar[IncidentFields]
    INCIDENTS_FIELDS_SUBSYSTEM_NAME: _ClassVar[IncidentFields]
    INCIDENTS_FIELDS_DURATION: _ClassVar[IncidentFields]
INCIDENTS_FIELDS_UNSPECIFIED: IncidentFields
INCIDENTS_FIELDS_ID: IncidentFields
INCIDENTS_FIELDS_SEVERITY: IncidentFields
INCIDENTS_FIELDS_NAME: IncidentFields
INCIDENTS_FIELDS_CREATED_TIME: IncidentFields
INCIDENTS_FIELDS_CLOSED_TIME: IncidentFields
INCIDENTS_FIELDS_STATE: IncidentFields
INCIDENTS_FIELDS_STATUS: IncidentFields
INCIDENTS_FIELDS_LAST_STATE_UPDATE_TIME: IncidentFields
INCIDENTS_FIELDS_APPLICATION_NAME: IncidentFields
INCIDENTS_FIELDS_SUBSYSTEM_NAME: IncidentFields
INCIDENTS_FIELDS_DURATION: IncidentFields

class Incident(_message.Message):
    __slots__ = ("id", "name", "state", "status", "assignments", "description", "severity", "contextual_labels", "display_labels", "events", "created_at", "closed_at", "last_state_update_time", "last_state_update_key", "is_muted", "meta_labels", "duration")
    class ContextualLabelsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    class DisplayLabelsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    STATE_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    ASSIGNMENTS_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    CONTEXTUAL_LABELS_FIELD_NUMBER: _ClassVar[int]
    DISPLAY_LABELS_FIELD_NUMBER: _ClassVar[int]
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    CLOSED_AT_FIELD_NUMBER: _ClassVar[int]
    LAST_STATE_UPDATE_TIME_FIELD_NUMBER: _ClassVar[int]
    LAST_STATE_UPDATE_KEY_FIELD_NUMBER: _ClassVar[int]
    IS_MUTED_FIELD_NUMBER: _ClassVar[int]
    META_LABELS_FIELD_NUMBER: _ClassVar[int]
    DURATION_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    state: _incident_state_pb2.IncidentState
    status: _incident_status_pb2.IncidentStatus
    assignments: _containers.RepeatedCompositeFieldContainer[_assignee_pb2.Assignment]
    description: _wrappers_pb2.StringValue
    severity: _incident_severity_pb2.IncidentSeverity
    contextual_labels: _containers.ScalarMap[str, str]
    display_labels: _containers.ScalarMap[str, str]
    events: _containers.RepeatedCompositeFieldContainer[_incident_event_pb2.IncidentEvent]
    created_at: _timestamp_pb2.Timestamp
    closed_at: _timestamp_pb2.Timestamp
    last_state_update_time: _timestamp_pb2.Timestamp
    last_state_update_key: _wrappers_pb2.StringValue
    is_muted: _wrappers_pb2.BoolValue
    meta_labels: _containers.RepeatedCompositeFieldContainer[_meta_label_pb2.MetaLabel]
    duration: _duration_pb2.Duration
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., state: _Optional[_Union[_incident_state_pb2.IncidentState, str]] = ..., status: _Optional[_Union[_incident_status_pb2.IncidentStatus, str]] = ..., assignments: _Optional[_Iterable[_Union[_assignee_pb2.Assignment, _Mapping]]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., severity: _Optional[_Union[_incident_severity_pb2.IncidentSeverity, str]] = ..., contextual_labels: _Optional[_Mapping[str, str]] = ..., display_labels: _Optional[_Mapping[str, str]] = ..., events: _Optional[_Iterable[_Union[_incident_event_pb2.IncidentEvent, _Mapping]]] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., closed_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., last_state_update_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., last_state_update_key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., is_muted: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., meta_labels: _Optional[_Iterable[_Union[_meta_label_pb2.MetaLabel, _Mapping]]] = ..., duration: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ...) -> None: ...

class IncidentFieldOneOf(_message.Message):
    __slots__ = ("id", "severity", "name", "created_at", "closed_at", "state", "status", "last_state_update_time", "application_name", "subsystem_name", "duration")
    ID_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    CLOSED_AT_FIELD_NUMBER: _ClassVar[int]
    STATE_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    LAST_STATE_UPDATE_TIME_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_NAME_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_NAME_FIELD_NUMBER: _ClassVar[int]
    DURATION_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    severity: _incident_severity_pb2.IncidentSeverity
    name: _wrappers_pb2.StringValue
    created_at: _timestamp_pb2.Timestamp
    closed_at: _timestamp_pb2.Timestamp
    state: _incident_state_pb2.IncidentState
    status: _incident_status_pb2.IncidentStatus
    last_state_update_time: _timestamp_pb2.Timestamp
    application_name: _wrappers_pb2.StringValue
    subsystem_name: _wrappers_pb2.StringValue
    duration: _duration_pb2.Duration
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., severity: _Optional[_Union[_incident_severity_pb2.IncidentSeverity, str]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., closed_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., state: _Optional[_Union[_incident_state_pb2.IncidentState, str]] = ..., status: _Optional[_Union[_incident_status_pb2.IncidentStatus, str]] = ..., last_state_update_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., application_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., duration: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ...) -> None: ...

class ContextualLabels(_message.Message):
    __slots__ = ("field_name", "field_value")
    FIELD_NAME_FIELD_NUMBER: _ClassVar[int]
    FIELD_VALUE_FIELD_NUMBER: _ClassVar[int]
    field_name: _wrappers_pb2.StringValue
    field_value: _wrappers_pb2.StringValue
    def __init__(self, field_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., field_value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GroupByValues(_message.Message):
    __slots__ = ("incident_field", "contextual_labels")
    INCIDENT_FIELD_FIELD_NUMBER: _ClassVar[int]
    CONTEXTUAL_LABELS_FIELD_NUMBER: _ClassVar[int]
    incident_field: IncidentFieldOneOf
    contextual_labels: ContextualLabels
    def __init__(self, incident_field: _Optional[_Union[IncidentFieldOneOf, _Mapping]] = ..., contextual_labels: _Optional[_Union[ContextualLabels, _Mapping]] = ...) -> None: ...

class IncidentAggregation(_message.Message):
    __slots__ = ("group_bys_value", "agg_state_count", "agg_status_count", "agg_severity_count", "agg_assignments_count", "first_created_at", "last_closed_at", "all_values_count", "list_incidents_id", "last_state_update_time", "agg_meta_labels_count")
    GROUP_BYS_VALUE_FIELD_NUMBER: _ClassVar[int]
    AGG_STATE_COUNT_FIELD_NUMBER: _ClassVar[int]
    AGG_STATUS_COUNT_FIELD_NUMBER: _ClassVar[int]
    AGG_SEVERITY_COUNT_FIELD_NUMBER: _ClassVar[int]
    AGG_ASSIGNMENTS_COUNT_FIELD_NUMBER: _ClassVar[int]
    FIRST_CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    LAST_CLOSED_AT_FIELD_NUMBER: _ClassVar[int]
    ALL_VALUES_COUNT_FIELD_NUMBER: _ClassVar[int]
    LIST_INCIDENTS_ID_FIELD_NUMBER: _ClassVar[int]
    LAST_STATE_UPDATE_TIME_FIELD_NUMBER: _ClassVar[int]
    AGG_META_LABELS_COUNT_FIELD_NUMBER: _ClassVar[int]
    group_bys_value: _containers.RepeatedCompositeFieldContainer[GroupByValues]
    agg_state_count: _containers.RepeatedCompositeFieldContainer[IncidentStateCount]
    agg_status_count: _containers.RepeatedCompositeFieldContainer[IncidentStatusCount]
    agg_severity_count: _containers.RepeatedCompositeFieldContainer[IncidentSeverityCount]
    agg_assignments_count: _containers.RepeatedCompositeFieldContainer[IncidentAssignmentCount]
    first_created_at: _timestamp_pb2.Timestamp
    last_closed_at: _timestamp_pb2.Timestamp
    all_values_count: _wrappers_pb2.UInt32Value
    list_incidents_id: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    last_state_update_time: _timestamp_pb2.Timestamp
    agg_meta_labels_count: _containers.RepeatedCompositeFieldContainer[IncidentMetaLabelsCount]
    def __init__(self, group_bys_value: _Optional[_Iterable[_Union[GroupByValues, _Mapping]]] = ..., agg_state_count: _Optional[_Iterable[_Union[IncidentStateCount, _Mapping]]] = ..., agg_status_count: _Optional[_Iterable[_Union[IncidentStatusCount, _Mapping]]] = ..., agg_severity_count: _Optional[_Iterable[_Union[IncidentSeverityCount, _Mapping]]] = ..., agg_assignments_count: _Optional[_Iterable[_Union[IncidentAssignmentCount, _Mapping]]] = ..., first_created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., last_closed_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., all_values_count: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., list_incidents_id: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., last_state_update_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., agg_meta_labels_count: _Optional[_Iterable[_Union[IncidentMetaLabelsCount, _Mapping]]] = ...) -> None: ...

class IncidentMetaLabelsCount(_message.Message):
    __slots__ = ("meta_label", "count")
    META_LABEL_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    meta_label: _meta_label_pb2.MetaLabel
    count: _wrappers_pb2.UInt32Value
    def __init__(self, meta_label: _Optional[_Union[_meta_label_pb2.MetaLabel, _Mapping]] = ..., count: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...

class IncidentAssignmentCount(_message.Message):
    __slots__ = ("assigned_to", "count")
    ASSIGNED_TO_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    assigned_to: _assignee_pb2.UserDetails
    count: _wrappers_pb2.UInt32Value
    def __init__(self, assigned_to: _Optional[_Union[_assignee_pb2.UserDetails, _Mapping]] = ..., count: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...

class IncidentSeverityCount(_message.Message):
    __slots__ = ("severity", "count")
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    severity: _incident_severity_pb2.IncidentSeverity
    count: _wrappers_pb2.UInt32Value
    def __init__(self, severity: _Optional[_Union[_incident_severity_pb2.IncidentSeverity, str]] = ..., count: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...

class IncidentStatusCount(_message.Message):
    __slots__ = ("status", "count")
    STATUS_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    status: _incident_status_pb2.IncidentStatus
    count: _wrappers_pb2.UInt32Value
    def __init__(self, status: _Optional[_Union[_incident_status_pb2.IncidentStatus, str]] = ..., count: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...

class IncidentStateCount(_message.Message):
    __slots__ = ("state", "count")
    STATE_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    state: _incident_state_pb2.IncidentState
    count: _wrappers_pb2.UInt32Value
    def __init__(self, state: _Optional[_Union[_incident_state_pb2.IncidentState, str]] = ..., count: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...
