from com.coralogixapis.incidents.v1.incident_action.upsert_incident_state import upsert_incident_state_payload_pb2 as _upsert_incident_state_payload_pb2
from com.coralogixapis.incidents.v1.incident_action.upsert_incident_state import upsert_incident_state_type_pb2 as _upsert_incident_state_type_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class IncidentEventUpsertState(_message.Message):
    __slots__ = ("state_type", "payload", "is_muted")
    STATE_TYPE_FIELD_NUMBER: _ClassVar[int]
    PAYLOAD_FIELD_NUMBER: _ClassVar[int]
    IS_MUTED_FIELD_NUMBER: _ClassVar[int]
    state_type: _upsert_incident_state_type_pb2.UpsertIncidentStateType
    payload: _upsert_incident_state_payload_pb2.UpsertIncidentStatePayload
    is_muted: _wrappers_pb2.BoolValue
    def __init__(self, state_type: _Optional[_Union[_upsert_incident_state_type_pb2.UpsertIncidentStateType, str]] = ..., payload: _Optional[_Union[_upsert_incident_state_payload_pb2.UpsertIncidentStatePayload, _Mapping]] = ..., is_muted: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
