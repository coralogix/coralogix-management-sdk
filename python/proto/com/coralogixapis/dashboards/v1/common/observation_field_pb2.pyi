from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DatasetScope(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    DATASET_SCOPE_UNSPECIFIED: _ClassVar[DatasetScope]
    DATASET_SCOPE_USER_DATA: _ClassVar[DatasetScope]
    DATASET_SCOPE_LABEL: _ClassVar[DatasetScope]
    DATASET_SCOPE_METADATA: _ClassVar[DatasetScope]
DATASET_SCOPE_UNSPECIFIED: DatasetScope
DATASET_SCOPE_USER_DATA: DatasetScope
DATASET_SCOPE_LABEL: DatasetScope
DATASET_SCOPE_METADATA: DatasetScope

class ObservationField(_message.Message):
    __slots__ = ("keypath", "scope")
    KEYPATH_FIELD_NUMBER: _ClassVar[int]
    SCOPE_FIELD_NUMBER: _ClassVar[int]
    keypath: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    scope: DatasetScope
    def __init__(self, keypath: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., scope: _Optional[_Union[DatasetScope, str]] = ...) -> None: ...
