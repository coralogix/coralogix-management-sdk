from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Pagination(_message.Message):
    __slots__ = ("offset", "limit")
    OFFSET_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    offset: _wrappers_pb2.Int32Value
    limit: _wrappers_pb2.Int32Value
    def __init__(self, offset: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
