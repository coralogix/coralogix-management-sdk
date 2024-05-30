from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class PolicyOrder(_message.Message):
    __slots__ = ("order", "id")
    ORDER_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    order: _wrappers_pb2.Int32Value
    id: _wrappers_pb2.StringValue
    def __init__(self, order: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
