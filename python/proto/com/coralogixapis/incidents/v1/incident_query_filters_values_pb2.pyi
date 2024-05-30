from com.coralogixapis.incidents.v1 import incident_severity_pb2 as _incident_severity_pb2
from com.coralogixapis.incidents.v1 import incident_status_pb2 as _incident_status_pb2
from com.coralogixapis.incidents.v1 import incident_state_pb2 as _incident_state_pb2
from com.coralogixapis.incidents.v1 import meta_label_pb2 as _meta_label_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AssigneeWithCount(_message.Message):
    __slots__ = ("assignee", "count")
    ASSIGNEE_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    assignee: _wrappers_pb2.StringValue
    count: _wrappers_pb2.Int32Value
    def __init__(self, assignee: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., count: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class IncidentStatusWithCount(_message.Message):
    __slots__ = ("status", "count")
    STATUS_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    status: _incident_status_pb2.IncidentStatus
    count: _wrappers_pb2.Int32Value
    def __init__(self, status: _Optional[_Union[_incident_status_pb2.IncidentStatus, str]] = ..., count: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class IncidentStateWithCount(_message.Message):
    __slots__ = ("state", "count")
    STATE_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    state: _incident_state_pb2.IncidentState
    count: _wrappers_pb2.Int32Value
    def __init__(self, state: _Optional[_Union[_incident_state_pb2.IncidentState, str]] = ..., count: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class IncidentSeverityWithCount(_message.Message):
    __slots__ = ("severity", "count")
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    severity: _incident_severity_pb2.IncidentSeverity
    count: _wrappers_pb2.Int32Value
    def __init__(self, severity: _Optional[_Union[_incident_severity_pb2.IncidentSeverity, str]] = ..., count: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class IncidentMetaLabelsWithCount(_message.Message):
    __slots__ = ("meta_label", "count")
    META_LABEL_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    meta_label: _meta_label_pb2.MetaLabel
    count: _wrappers_pb2.Int32Value
    def __init__(self, meta_label: _Optional[_Union[_meta_label_pb2.MetaLabel, _Mapping]] = ..., count: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class IncidentQueryFiltersValues(_message.Message):
    __slots__ = ("assignee_with_count", "status_with_count", "state_with_count", "severity_with_count", "contextual_labels", "meta_labels_with_count")
    class ContextualLabelsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: ContextualLabelValuesWithCount
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[ContextualLabelValuesWithCount, _Mapping]] = ...) -> None: ...
    ASSIGNEE_WITH_COUNT_FIELD_NUMBER: _ClassVar[int]
    STATUS_WITH_COUNT_FIELD_NUMBER: _ClassVar[int]
    STATE_WITH_COUNT_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_WITH_COUNT_FIELD_NUMBER: _ClassVar[int]
    CONTEXTUAL_LABELS_FIELD_NUMBER: _ClassVar[int]
    META_LABELS_WITH_COUNT_FIELD_NUMBER: _ClassVar[int]
    assignee_with_count: _containers.RepeatedCompositeFieldContainer[AssigneeWithCount]
    status_with_count: _containers.RepeatedCompositeFieldContainer[IncidentStatusWithCount]
    state_with_count: _containers.RepeatedCompositeFieldContainer[IncidentStateWithCount]
    severity_with_count: _containers.RepeatedCompositeFieldContainer[IncidentSeverityWithCount]
    contextual_labels: _containers.MessageMap[str, ContextualLabelValuesWithCount]
    meta_labels_with_count: _containers.RepeatedCompositeFieldContainer[IncidentMetaLabelsWithCount]
    def __init__(self, assignee_with_count: _Optional[_Iterable[_Union[AssigneeWithCount, _Mapping]]] = ..., status_with_count: _Optional[_Iterable[_Union[IncidentStatusWithCount, _Mapping]]] = ..., state_with_count: _Optional[_Iterable[_Union[IncidentStateWithCount, _Mapping]]] = ..., severity_with_count: _Optional[_Iterable[_Union[IncidentSeverityWithCount, _Mapping]]] = ..., contextual_labels: _Optional[_Mapping[str, ContextualLabelValuesWithCount]] = ..., meta_labels_with_count: _Optional[_Iterable[_Union[IncidentMetaLabelsWithCount, _Mapping]]] = ...) -> None: ...

class ContextualLabelValuesWithCount(_message.Message):
    __slots__ = ("values_with_count",)
    VALUES_WITH_COUNT_FIELD_NUMBER: _ClassVar[int]
    values_with_count: _containers.RepeatedCompositeFieldContainer[ContextualLabelValueWithCount]
    def __init__(self, values_with_count: _Optional[_Iterable[_Union[ContextualLabelValueWithCount, _Mapping]]] = ...) -> None: ...

class ContextualLabelValueWithCount(_message.Message):
    __slots__ = ("contextual_label_value", "count")
    CONTEXTUAL_LABEL_VALUE_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    contextual_label_value: _wrappers_pb2.StringValue
    count: _wrappers_pb2.Int32Value
    def __init__(self, contextual_label_value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., count: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
