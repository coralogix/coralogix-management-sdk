from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class IncidentSeverity(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    INCIDENT_SEVERITY_UNSPECIFIED: _ClassVar[IncidentSeverity]
    INCIDENT_SEVERITY_INFO: _ClassVar[IncidentSeverity]
    INCIDENT_SEVERITY_WARNING: _ClassVar[IncidentSeverity]
    INCIDENT_SEVERITY_ERROR: _ClassVar[IncidentSeverity]
    INCIDENT_SEVERITY_CRITICAL: _ClassVar[IncidentSeverity]
INCIDENT_SEVERITY_UNSPECIFIED: IncidentSeverity
INCIDENT_SEVERITY_INFO: IncidentSeverity
INCIDENT_SEVERITY_WARNING: IncidentSeverity
INCIDENT_SEVERITY_ERROR: IncidentSeverity
INCIDENT_SEVERITY_CRITICAL: IncidentSeverity
