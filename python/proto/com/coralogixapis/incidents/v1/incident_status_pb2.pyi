from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class IncidentStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    INCIDENT_STATUS_UNSPECIFIED: _ClassVar[IncidentStatus]
    INCIDENT_STATUS_TRIGGERED: _ClassVar[IncidentStatus]
    INCIDENT_STATUS_ACKNOWLEDGED: _ClassVar[IncidentStatus]
    INCIDENT_STATUS_RESOLVED: _ClassVar[IncidentStatus]
INCIDENT_STATUS_UNSPECIFIED: IncidentStatus
INCIDENT_STATUS_TRIGGERED: IncidentStatus
INCIDENT_STATUS_ACKNOWLEDGED: IncidentStatus
INCIDENT_STATUS_RESOLVED: IncidentStatus
