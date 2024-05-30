from com.coralogixapis.incidents.v1.incident_event import incident_event_acknowledge_pb2 as _incident_event_acknowledge_pb2
from com.coralogixapis.incidents.v1.incident_event import incident_event_assign_pb2 as _incident_event_assign_pb2
from com.coralogixapis.incidents.v1.incident_event import incident_event_close_pb2 as _incident_event_close_pb2
from com.coralogixapis.incidents.v1.incident_event import incident_event_originator_administrative_pb2 as _incident_event_originator_administrative_pb2
from com.coralogixapis.incidents.v1.incident_event import incident_event_originator_operational_pb2 as _incident_event_originator_operational_pb2
from com.coralogixapis.incidents.v1.incident_event import incident_event_originator_type_pb2 as _incident_event_originator_type_pb2
from com.coralogixapis.incidents.v1.incident_event import incident_event_snooze_indicator_pb2 as _incident_event_snooze_indicator_pb2
from com.coralogixapis.incidents.v1.incident_event import incident_event_type_pb2 as _incident_event_type_pb2
from com.coralogixapis.incidents.v1.incident_event import incident_event_upsert_state_pb2 as _incident_event_upsert_state_pb2
from com.coralogixapis.incidents.v1.incident_event import incident_event_unassign_pb2 as _incident_event_unassign_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class IncidentEvent(_message.Message):
    __slots__ = ("id", "incident_event_type", "snooze_indicator", "assignment", "unassign", "upsert_state", "acknowledge", "close", "originator_type", "administrative_event", "operational_event")
    ID_FIELD_NUMBER: _ClassVar[int]
    INCIDENT_EVENT_TYPE_FIELD_NUMBER: _ClassVar[int]
    SNOOZE_INDICATOR_FIELD_NUMBER: _ClassVar[int]
    ASSIGNMENT_FIELD_NUMBER: _ClassVar[int]
    UNASSIGN_FIELD_NUMBER: _ClassVar[int]
    UPSERT_STATE_FIELD_NUMBER: _ClassVar[int]
    ACKNOWLEDGE_FIELD_NUMBER: _ClassVar[int]
    CLOSE_FIELD_NUMBER: _ClassVar[int]
    ORIGINATOR_TYPE_FIELD_NUMBER: _ClassVar[int]
    ADMINISTRATIVE_EVENT_FIELD_NUMBER: _ClassVar[int]
    OPERATIONAL_EVENT_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    incident_event_type: _incident_event_type_pb2.IncidentEventType
    snooze_indicator: _incident_event_snooze_indicator_pb2.IncidentEventSnoozeIndicator
    assignment: _incident_event_assign_pb2.IncidentEventAssign
    unassign: _incident_event_unassign_pb2.IncidentEventUnassign
    upsert_state: _incident_event_upsert_state_pb2.IncidentEventUpsertState
    acknowledge: _incident_event_acknowledge_pb2.IncidentEventAcknowledge
    close: _incident_event_close_pb2.IncidentEventClose
    originator_type: _incident_event_originator_type_pb2.OriginatorType
    administrative_event: _incident_event_originator_administrative_pb2.IncidentEventOriginatorAdministrative
    operational_event: _incident_event_originator_operational_pb2.IncidentEventOriginatorOperational
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., incident_event_type: _Optional[_Union[_incident_event_type_pb2.IncidentEventType, str]] = ..., snooze_indicator: _Optional[_Union[_incident_event_snooze_indicator_pb2.IncidentEventSnoozeIndicator, _Mapping]] = ..., assignment: _Optional[_Union[_incident_event_assign_pb2.IncidentEventAssign, _Mapping]] = ..., unassign: _Optional[_Union[_incident_event_unassign_pb2.IncidentEventUnassign, _Mapping]] = ..., upsert_state: _Optional[_Union[_incident_event_upsert_state_pb2.IncidentEventUpsertState, _Mapping]] = ..., acknowledge: _Optional[_Union[_incident_event_acknowledge_pb2.IncidentEventAcknowledge, _Mapping]] = ..., close: _Optional[_Union[_incident_event_close_pb2.IncidentEventClose, _Mapping]] = ..., originator_type: _Optional[_Union[_incident_event_originator_type_pb2.OriginatorType, str]] = ..., administrative_event: _Optional[_Union[_incident_event_originator_administrative_pb2.IncidentEventOriginatorAdministrative, _Mapping]] = ..., operational_event: _Optional[_Union[_incident_event_originator_operational_pb2.IncidentEventOriginatorOperational, _Mapping]] = ...) -> None: ...
