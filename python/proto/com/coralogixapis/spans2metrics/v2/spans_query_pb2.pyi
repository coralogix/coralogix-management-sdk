from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SpansQuery(_message.Message):
    __slots__ = ("lucene", "applicationname_filters", "subsystemname_filters", "action_filters", "service_filters")
    LUCENE_FIELD_NUMBER: _ClassVar[int]
    APPLICATIONNAME_FILTERS_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEMNAME_FILTERS_FIELD_NUMBER: _ClassVar[int]
    ACTION_FILTERS_FIELD_NUMBER: _ClassVar[int]
    SERVICE_FILTERS_FIELD_NUMBER: _ClassVar[int]
    lucene: _wrappers_pb2.StringValue
    applicationname_filters: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    subsystemname_filters: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    action_filters: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    service_filters: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, lucene: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., applicationname_filters: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystemname_filters: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., action_filters: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., service_filters: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...
