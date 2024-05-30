from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DashboardFolder(_message.Message):
    __slots__ = ("id", "name", "parent_id")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    PARENT_ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    parent_id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., parent_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
