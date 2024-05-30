from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Group(_message.Message):
    __slots__ = ("field", "value", "groups")
    FIELD_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    GROUPS_FIELD_NUMBER: _ClassVar[int]
    field: FieldGroup
    value: _wrappers_pb2.DoubleValue
    groups: _containers.RepeatedCompositeFieldContainer[Group]
    def __init__(self, field: _Optional[_Union[FieldGroup, _Mapping]] = ..., value: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ..., groups: _Optional[_Iterable[_Union[Group, _Mapping]]] = ...) -> None: ...

class MultiGroup(_message.Message):
    __slots__ = ("fields", "values")
    FIELDS_FIELD_NUMBER: _ClassVar[int]
    VALUES_FIELD_NUMBER: _ClassVar[int]
    fields: _containers.RepeatedCompositeFieldContainer[FieldGroup]
    values: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.DoubleValue]
    def __init__(self, fields: _Optional[_Iterable[_Union[FieldGroup, _Mapping]]] = ..., values: _Optional[_Iterable[_Union[_wrappers_pb2.DoubleValue, _Mapping]]] = ...) -> None: ...

class FieldGroup(_message.Message):
    __slots__ = ("name", "value")
    NAME_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    value: _wrappers_pb2.StringValue
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
