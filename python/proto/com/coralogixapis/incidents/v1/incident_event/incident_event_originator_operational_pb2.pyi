from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class IncidentEventOriginatorOperational(_message.Message):
    __slots__ = ("system_name",)
    SYSTEM_NAME_FIELD_NUMBER: _ClassVar[int]
    system_name: _wrappers_pb2.StringValue
    def __init__(self, system_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
