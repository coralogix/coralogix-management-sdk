from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Backend(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    BACKEND_ARCHIVE_UNSPECIFIED: _ClassVar[Backend]
    BACKEND_OPENSEARCH: _ClassVar[Backend]
BACKEND_ARCHIVE_UNSPECIFIED: Backend
BACKEND_OPENSEARCH: Backend

class UntypedKeypath(_message.Message):
    __slots__ = ("root", "path_elements")
    class Root(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        ROOT_EVENT_LABELS_UNSPECIFIED: _ClassVar[UntypedKeypath.Root]
        ROOT_EVENT_METADATA: _ClassVar[UntypedKeypath.Root]
        ROOT_USER_DATA: _ClassVar[UntypedKeypath.Root]
    ROOT_EVENT_LABELS_UNSPECIFIED: UntypedKeypath.Root
    ROOT_EVENT_METADATA: UntypedKeypath.Root
    ROOT_USER_DATA: UntypedKeypath.Root
    ROOT_FIELD_NUMBER: _ClassVar[int]
    PATH_ELEMENTS_FIELD_NUMBER: _ClassVar[int]
    root: UntypedKeypath.Root
    path_elements: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, root: _Optional[_Union[UntypedKeypath.Root, str]] = ..., path_elements: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class KeyValue(_message.Message):
    __slots__ = ("key", "value")
    KEY_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    key: _wrappers_pb2.StringValue
    value: _wrappers_pb2.StringValue
    def __init__(self, key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class Interval(_message.Message):
    __slots__ = ("to",)
    FROM_FIELD_NUMBER: _ClassVar[int]
    TO_FIELD_NUMBER: _ClassVar[int]
    to: _wrappers_pb2.Int64Value
    def __init__(self, to: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., **kwargs) -> None: ...
