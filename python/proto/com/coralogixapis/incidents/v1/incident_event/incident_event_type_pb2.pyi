from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class IncidentEventType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    INCIDENT_EVENT_TYPE_UNSPECIFIED: _ClassVar[IncidentEventType]
    INCIDENT_EVENT_TYPE_UPSERT_STATE: _ClassVar[IncidentEventType]
    INCIDENT_EVENT_TYPE_OPEN: _ClassVar[IncidentEventType]
    INCIDENT_EVENT_TYPE_CLOSE: _ClassVar[IncidentEventType]
    INCIDENT_EVENT_TYPE_SNOOZE_INDICATOR: _ClassVar[IncidentEventType]
    INCIDENT_EVENT_TYPE_ASSIGN: _ClassVar[IncidentEventType]
    INCIDENT_EVENT_TYPE_UNASSIGN: _ClassVar[IncidentEventType]
    INCIDENT_EVENT_TYPE_ACKNOWLEDGE: _ClassVar[IncidentEventType]
INCIDENT_EVENT_TYPE_UNSPECIFIED: IncidentEventType
INCIDENT_EVENT_TYPE_UPSERT_STATE: IncidentEventType
INCIDENT_EVENT_TYPE_OPEN: IncidentEventType
INCIDENT_EVENT_TYPE_CLOSE: IncidentEventType
INCIDENT_EVENT_TYPE_SNOOZE_INDICATOR: IncidentEventType
INCIDENT_EVENT_TYPE_ASSIGN: IncidentEventType
INCIDENT_EVENT_TYPE_UNASSIGN: IncidentEventType
INCIDENT_EVENT_TYPE_ACKNOWLEDGE: IncidentEventType
