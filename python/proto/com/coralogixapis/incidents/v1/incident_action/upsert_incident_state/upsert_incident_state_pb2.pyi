from com.coralogixapis.incidents.v1.incident_action.upsert_incident_state import upsert_incident_state_payload_pb2 as _upsert_incident_state_payload_pb2
from com.coralogixapis.incidents.v1.incident_action.upsert_incident_state import upsert_incident_state_type_pb2 as _upsert_incident_state_type_pb2
from com.coralogixapis.incidents.v1 import scoping_details_pb2 as _scoping_details_pb2
from com.coralogixapis.incidents.v1 import incident_severity_pb2 as _incident_severity_pb2
from com.coralogixapis.incidents.v1 import meta_label_pb2 as _meta_label_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UpsertIncidentState(_message.Message):
    __slots__ = ("correlation_key", "company_id", "state_type", "severity", "contextual_labels", "display_labels", "payload", "scope_details", "close_on_resolve_event", "is_muted", "meta_labels")
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
    CORRELATION_KEY_FIELD_NUMBER: _ClassVar[int]
    COMPANY_ID_FIELD_NUMBER: _ClassVar[int]
    STATE_TYPE_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    CONTEXTUAL_LABELS_FIELD_NUMBER: _ClassVar[int]
    DISPLAY_LABELS_FIELD_NUMBER: _ClassVar[int]
    PAYLOAD_FIELD_NUMBER: _ClassVar[int]
    SCOPE_DETAILS_FIELD_NUMBER: _ClassVar[int]
    CLOSE_ON_RESOLVE_EVENT_FIELD_NUMBER: _ClassVar[int]
    IS_MUTED_FIELD_NUMBER: _ClassVar[int]
    META_LABELS_FIELD_NUMBER: _ClassVar[int]
    correlation_key: _wrappers_pb2.StringValue
    company_id: _wrappers_pb2.Int32Value
    state_type: _upsert_incident_state_type_pb2.UpsertIncidentStateType
    severity: _incident_severity_pb2.IncidentSeverity
    contextual_labels: _containers.ScalarMap[str, str]
    display_labels: _containers.ScalarMap[str, str]
    payload: _upsert_incident_state_payload_pb2.UpsertIncidentStatePayload
    scope_details: _scoping_details_pb2.ScopeDetails
    close_on_resolve_event: _wrappers_pb2.BoolValue
    is_muted: _wrappers_pb2.BoolValue
    meta_labels: _containers.RepeatedCompositeFieldContainer[_meta_label_pb2.MetaLabel]
    def __init__(self, correlation_key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., company_id: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., state_type: _Optional[_Union[_upsert_incident_state_type_pb2.UpsertIncidentStateType, str]] = ..., severity: _Optional[_Union[_incident_severity_pb2.IncidentSeverity, str]] = ..., contextual_labels: _Optional[_Mapping[str, str]] = ..., display_labels: _Optional[_Mapping[str, str]] = ..., payload: _Optional[_Union[_upsert_incident_state_payload_pb2.UpsertIncidentStatePayload, _Mapping]] = ..., scope_details: _Optional[_Union[_scoping_details_pb2.ScopeDetails, _Mapping]] = ..., close_on_resolve_event: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., is_muted: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., meta_labels: _Optional[_Iterable[_Union[_meta_label_pb2.MetaLabel, _Mapping]]] = ...) -> None: ...
