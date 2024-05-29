from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Filter(_message.Message):
    __slots__ = ("name", "selected_values")
    class SelectedValuesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: bool
        def __init__(self, key: _Optional[str] = ..., value: bool = ...) -> None: ...
    NAME_FIELD_NUMBER: _ClassVar[int]
    SELECTED_VALUES_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    selected_values: _containers.ScalarMap[str, bool]
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., selected_values: _Optional[_Mapping[str, bool]] = ...) -> None: ...

class SelectedFilters(_message.Message):
    __slots__ = ("filters",)
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    filters: _containers.RepeatedCompositeFieldContainer[Filter]
    def __init__(self, filters: _Optional[_Iterable[_Union[Filter, _Mapping]]] = ...) -> None: ...
