from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class MetaLabel(_message.Message):
    __slots__ = ("key", "value")
    KEY_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    key: _wrappers_pb2.StringValue
    value: _wrappers_pb2.StringValue
    def __init__(self, key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
