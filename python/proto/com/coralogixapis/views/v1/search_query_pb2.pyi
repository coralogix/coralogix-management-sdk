from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SearchQuery(_message.Message):
    __slots__ = ("query",)
    QUERY_FIELD_NUMBER: _ClassVar[int]
    query: _wrappers_pb2.StringValue
    def __init__(self, query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
