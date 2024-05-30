from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.openapi.v1 import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Severity(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    SEVERITY_UNSPECIFIED: _ClassVar[Severity]
    SEVERITY_DEBUG: _ClassVar[Severity]
    SEVERITY_VERBOSE: _ClassVar[Severity]
    SEVERITY_INFO: _ClassVar[Severity]
    SEVERITY_WARNING: _ClassVar[Severity]
    SEVERITY_ERROR: _ClassVar[Severity]
    SEVERITY_CRITICAL: _ClassVar[Severity]
SEVERITY_UNSPECIFIED: Severity
SEVERITY_DEBUG: Severity
SEVERITY_VERBOSE: Severity
SEVERITY_INFO: Severity
SEVERITY_WARNING: Severity
SEVERITY_ERROR: Severity
SEVERITY_CRITICAL: Severity

class LogsQuery(_message.Message):
    __slots__ = ("lucene", "alias", "applicationname_filters", "subsystemname_filters", "severity_filters")
    LUCENE_FIELD_NUMBER: _ClassVar[int]
    ALIAS_FIELD_NUMBER: _ClassVar[int]
    APPLICATIONNAME_FILTERS_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEMNAME_FILTERS_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FILTERS_FIELD_NUMBER: _ClassVar[int]
    lucene: _wrappers_pb2.StringValue
    alias: _wrappers_pb2.StringValue
    applicationname_filters: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    subsystemname_filters: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    severity_filters: _containers.RepeatedScalarFieldContainer[Severity]
    def __init__(self, lucene: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., alias: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., applicationname_filters: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystemname_filters: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., severity_filters: _Optional[_Iterable[_Union[Severity, str]]] = ...) -> None: ...
