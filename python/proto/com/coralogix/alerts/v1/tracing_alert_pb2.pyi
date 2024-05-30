from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class TracingAlert(_message.Message):
    __slots__ = ("condition_latency", "field_filters", "tag_filters")
    CONDITION_LATENCY_FIELD_NUMBER: _ClassVar[int]
    FIELD_FILTERS_FIELD_NUMBER: _ClassVar[int]
    TAG_FILTERS_FIELD_NUMBER: _ClassVar[int]
    condition_latency: int
    field_filters: _containers.RepeatedCompositeFieldContainer[FilterData]
    tag_filters: _containers.RepeatedCompositeFieldContainer[FilterData]
    def __init__(self, condition_latency: _Optional[int] = ..., field_filters: _Optional[_Iterable[_Union[FilterData, _Mapping]]] = ..., tag_filters: _Optional[_Iterable[_Union[FilterData, _Mapping]]] = ...) -> None: ...

class FilterData(_message.Message):
    __slots__ = ("field", "filters")
    FIELD_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    field: str
    filters: _containers.RepeatedCompositeFieldContainer[Filters]
    def __init__(self, field: _Optional[str] = ..., filters: _Optional[_Iterable[_Union[Filters, _Mapping]]] = ...) -> None: ...

class Filters(_message.Message):
    __slots__ = ("values", "operator")
    VALUES_FIELD_NUMBER: _ClassVar[int]
    OPERATOR_FIELD_NUMBER: _ClassVar[int]
    values: _containers.RepeatedScalarFieldContainer[str]
    operator: str
    def __init__(self, values: _Optional[_Iterable[str]] = ..., operator: _Optional[str] = ...) -> None: ...
