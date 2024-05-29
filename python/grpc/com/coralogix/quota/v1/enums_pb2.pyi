from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class Priority(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    PRIORITY_TYPE_UNSPECIFIED: _ClassVar[Priority]
    PRIORITY_TYPE_BLOCK: _ClassVar[Priority]
    PRIORITY_TYPE_LOW: _ClassVar[Priority]
    PRIORITY_TYPE_MEDIUM: _ClassVar[Priority]
    PRIORITY_TYPE_HIGH: _ClassVar[Priority]

class Severity(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    SEVERITY_UNSPECIFIED: _ClassVar[Severity]
    SEVERITY_DEBUG: _ClassVar[Severity]
    SEVERITY_VERBOSE: _ClassVar[Severity]
    SEVERITY_INFO: _ClassVar[Severity]
    SEVERITY_WARNING: _ClassVar[Severity]
    SEVERITY_ERROR: _ClassVar[Severity]
    SEVERITY_CRITICAL: _ClassVar[Severity]

class SourceType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    SOURCE_TYPE_UNSPECIFIED: _ClassVar[SourceType]
    SOURCE_TYPE_LOGS: _ClassVar[SourceType]
    SOURCE_TYPE_SPANS: _ClassVar[SourceType]

class RuleTypeId(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    RULE_TYPE_ID_UNSPECIFIED: _ClassVar[RuleTypeId]
    RULE_TYPE_ID_IS: _ClassVar[RuleTypeId]
    RULE_TYPE_ID_IS_NOT: _ClassVar[RuleTypeId]
    RULE_TYPE_ID_START_WITH: _ClassVar[RuleTypeId]
    RULE_TYPE_ID_INCLUDES: _ClassVar[RuleTypeId]
PRIORITY_TYPE_UNSPECIFIED: Priority
PRIORITY_TYPE_BLOCK: Priority
PRIORITY_TYPE_LOW: Priority
PRIORITY_TYPE_MEDIUM: Priority
PRIORITY_TYPE_HIGH: Priority
SEVERITY_UNSPECIFIED: Severity
SEVERITY_DEBUG: Severity
SEVERITY_VERBOSE: Severity
SEVERITY_INFO: Severity
SEVERITY_WARNING: Severity
SEVERITY_ERROR: Severity
SEVERITY_CRITICAL: Severity
SOURCE_TYPE_UNSPECIFIED: SourceType
SOURCE_TYPE_LOGS: SourceType
SOURCE_TYPE_SPANS: SourceType
RULE_TYPE_ID_UNSPECIFIED: RuleTypeId
RULE_TYPE_ID_IS: RuleTypeId
RULE_TYPE_ID_IS_NOT: RuleTypeId
RULE_TYPE_ID_START_WITH: RuleTypeId
RULE_TYPE_ID_INCLUDES: RuleTypeId
