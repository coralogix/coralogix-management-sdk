from com.coralogixapis.incidents.v1.incident_action.upsert_incident_state import upsert_incident_state_pb2 as _upsert_incident_state_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AlertIncidentActionType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ALERT_INCIDENT_ACTION_TYPE_UNSPECIFIED: _ClassVar[AlertIncidentActionType]
    ALERT_INCIDENT_ACTION_TYPE_CLOSE: _ClassVar[AlertIncidentActionType]
    ALERT_INCIDENT_ACTION_TYPE_ACKNOWLEDGE: _ClassVar[AlertIncidentActionType]
    ALERT_INCIDENT_ACTION_TYPE_ASSIGN: _ClassVar[AlertIncidentActionType]

class IncidentActionType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    INCIDENT_ACTION_TYPE_UNSPECIFIED: _ClassVar[IncidentActionType]
    INCIDENT_ACTION_TYPE_CLOSE: _ClassVar[IncidentActionType]
    INCIDENT_ACTION_TYPE_ACKNOWLEDGE: _ClassVar[IncidentActionType]
    INCIDENT_ACTION_TYPE_ASSIGN: _ClassVar[IncidentActionType]
    INCIDENT_ACTION_TYPE_UPSERT_STATE: _ClassVar[IncidentActionType]
ALERT_INCIDENT_ACTION_TYPE_UNSPECIFIED: AlertIncidentActionType
ALERT_INCIDENT_ACTION_TYPE_CLOSE: AlertIncidentActionType
ALERT_INCIDENT_ACTION_TYPE_ACKNOWLEDGE: AlertIncidentActionType
ALERT_INCIDENT_ACTION_TYPE_ASSIGN: AlertIncidentActionType
INCIDENT_ACTION_TYPE_UNSPECIFIED: IncidentActionType
INCIDENT_ACTION_TYPE_CLOSE: IncidentActionType
INCIDENT_ACTION_TYPE_ACKNOWLEDGE: IncidentActionType
INCIDENT_ACTION_TYPE_ASSIGN: IncidentActionType
INCIDENT_ACTION_TYPE_UPSERT_STATE: IncidentActionType

class IncidentAction(_message.Message):
    __slots__ = ("action_type", "upsert_state", "action_timestamp")
    ACTION_TYPE_FIELD_NUMBER: _ClassVar[int]
    UPSERT_STATE_FIELD_NUMBER: _ClassVar[int]
    ACTION_TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    action_type: IncidentActionType
    upsert_state: _upsert_incident_state_pb2.UpsertIncidentState
    action_timestamp: _timestamp_pb2.Timestamp
    def __init__(self, action_type: _Optional[_Union[IncidentActionType, str]] = ..., upsert_state: _Optional[_Union[_upsert_incident_state_pb2.UpsertIncidentState, _Mapping]] = ..., action_timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
