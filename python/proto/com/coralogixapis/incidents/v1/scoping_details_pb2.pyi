from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ScopeDetails(_message.Message):
    __slots__ = ("subsystem_name", "application_name")
    SUBSYSTEM_NAME_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_NAME_FIELD_NUMBER: _ClassVar[int]
    subsystem_name: _wrappers_pb2.StringValue
    application_name: _wrappers_pb2.StringValue
    def __init__(self, subsystem_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., application_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
