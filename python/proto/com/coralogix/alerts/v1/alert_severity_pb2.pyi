from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class AlertSeverity(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ALERT_SEVERITY_INFO_OR_UNSPECIFIED: _ClassVar[AlertSeverity]
    ALERT_SEVERITY_WARNING: _ClassVar[AlertSeverity]
    ALERT_SEVERITY_CRITICAL: _ClassVar[AlertSeverity]
    ALERT_SEVERITY_ERROR: _ClassVar[AlertSeverity]
ALERT_SEVERITY_INFO_OR_UNSPECIFIED: AlertSeverity
ALERT_SEVERITY_WARNING: AlertSeverity
ALERT_SEVERITY_CRITICAL: AlertSeverity
ALERT_SEVERITY_ERROR: AlertSeverity
