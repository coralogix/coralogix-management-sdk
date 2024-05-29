from com.coralogix.quota.v1 import enums_pb2 as _enums_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class LogRules(_message.Message):
    __slots__ = ("severities",)
    SEVERITIES_FIELD_NUMBER: _ClassVar[int]
    severities: _containers.RepeatedScalarFieldContainer[_enums_pb2.Severity]
    def __init__(self, severities: _Optional[_Iterable[_Union[_enums_pb2.Severity, str]]] = ...) -> None: ...
