from com.coralogixapis.dashboards.v1.common import group_pb2 as _group_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GroupedSeries(_message.Message):
    __slots__ = ("groups",)
    GROUPS_FIELD_NUMBER: _ClassVar[int]
    groups: _containers.RepeatedCompositeFieldContainer[_group_pb2.Group]
    def __init__(self, groups: _Optional[_Iterable[_Union[_group_pb2.Group, _Mapping]]] = ...) -> None: ...
