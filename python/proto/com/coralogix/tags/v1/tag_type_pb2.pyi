from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Type(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    TYPE_UNSPECIFIED: _ClassVar[Type]
    TYPE_CUSTOM_EVENT: _ClassVar[Type]
    TYPE_BITBUCKET: _ClassVar[Type]
    TYPE_GITLAB: _ClassVar[Type]
    TYPE_TFS: _ClassVar[Type]
    TYPE_HEROKU: _ClassVar[Type]
TYPE_UNSPECIFIED: Type
TYPE_CUSTOM_EVENT: Type
TYPE_BITBUCKET: Type
TYPE_GITLAB: Type
TYPE_TFS: Type
TYPE_HEROKU: Type

class TagType(_message.Message):
    __slots__ = ("id", "type")
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.Int64Value
    type: Type
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., type: _Optional[_Union[Type, str]] = ...) -> None: ...
