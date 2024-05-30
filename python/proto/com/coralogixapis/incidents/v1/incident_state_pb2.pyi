from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class IncidentState(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    INCIDENT_STATE_UNSPECIFIED: _ClassVar[IncidentState]
    INCIDENT_STATE_TRIGGERED: _ClassVar[IncidentState]
    INCIDENT_STATE_RESOLVED: _ClassVar[IncidentState]
INCIDENT_STATE_UNSPECIFIED: IncidentState
INCIDENT_STATE_TRIGGERED: IncidentState
INCIDENT_STATE_RESOLVED: IncidentState
