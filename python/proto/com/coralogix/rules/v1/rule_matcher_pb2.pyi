from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class RuleMatcher(_message.Message):
    __slots__ = ("application_name", "subsystem_name", "severity")
    APPLICATION_NAME_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_NAME_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    application_name: ApplicationNameConstraint
    subsystem_name: SubsystemNameConstraint
    severity: SeverityConstraint
    def __init__(self, application_name: _Optional[_Union[ApplicationNameConstraint, _Mapping]] = ..., subsystem_name: _Optional[_Union[SubsystemNameConstraint, _Mapping]] = ..., severity: _Optional[_Union[SeverityConstraint, _Mapping]] = ...) -> None: ...

class ApplicationNameConstraint(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: _wrappers_pb2.StringValue
    def __init__(self, value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class SubsystemNameConstraint(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: _wrappers_pb2.StringValue
    def __init__(self, value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class SeverityConstraint(_message.Message):
    __slots__ = ("value",)
    class Value(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        VALUE_DEBUG_OR_UNSPECIFIED: _ClassVar[SeverityConstraint.Value]
        VALUE_VERBOSE: _ClassVar[SeverityConstraint.Value]
        VALUE_INFO: _ClassVar[SeverityConstraint.Value]
        VALUE_WARNING: _ClassVar[SeverityConstraint.Value]
        VALUE_ERROR: _ClassVar[SeverityConstraint.Value]
        VALUE_CRITICAL: _ClassVar[SeverityConstraint.Value]
    VALUE_DEBUG_OR_UNSPECIFIED: SeverityConstraint.Value
    VALUE_VERBOSE: SeverityConstraint.Value
    VALUE_INFO: SeverityConstraint.Value
    VALUE_WARNING: SeverityConstraint.Value
    VALUE_ERROR: SeverityConstraint.Value
    VALUE_CRITICAL: SeverityConstraint.Value
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: SeverityConstraint.Value
    def __init__(self, value: _Optional[_Union[SeverityConstraint.Value, str]] = ...) -> None: ...
