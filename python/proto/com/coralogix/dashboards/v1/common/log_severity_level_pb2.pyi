from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class LogSeverityLevel(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    LOG_SEVERITY_LEVEL_UNSPECIFIED: _ClassVar[LogSeverityLevel]
    LOG_SEVERITY_LEVEL_DEBUG: _ClassVar[LogSeverityLevel]
    LOG_SEVERITY_LEVEL_VERBOSE: _ClassVar[LogSeverityLevel]
    LOG_SEVERITY_LEVEL_INFO: _ClassVar[LogSeverityLevel]
    LOG_SEVERITY_LEVEL_WARNING: _ClassVar[LogSeverityLevel]
    LOG_SEVERITY_LEVEL_ERROR: _ClassVar[LogSeverityLevel]
    LOG_SEVERITY_LEVEL_CRITICAL: _ClassVar[LogSeverityLevel]
LOG_SEVERITY_LEVEL_UNSPECIFIED: LogSeverityLevel
LOG_SEVERITY_LEVEL_DEBUG: LogSeverityLevel
LOG_SEVERITY_LEVEL_VERBOSE: LogSeverityLevel
LOG_SEVERITY_LEVEL_INFO: LogSeverityLevel
LOG_SEVERITY_LEVEL_WARNING: LogSeverityLevel
LOG_SEVERITY_LEVEL_ERROR: LogSeverityLevel
LOG_SEVERITY_LEVEL_CRITICAL: LogSeverityLevel
