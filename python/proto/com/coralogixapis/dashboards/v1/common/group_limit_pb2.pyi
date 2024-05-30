from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GroupLimit(_message.Message):
    __slots__ = ("group_by_fields", "limit", "min_percentage")
    GROUP_BY_FIELDS_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    MIN_PERCENTAGE_FIELD_NUMBER: _ClassVar[int]
    group_by_fields: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    limit: _wrappers_pb2.Int32Value
    min_percentage: _wrappers_pb2.Int32Value
    def __init__(self, group_by_fields: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., min_percentage: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
