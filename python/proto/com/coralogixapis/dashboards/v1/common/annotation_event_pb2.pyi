from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AnnotationEvent(_message.Message):
    __slots__ = ("instant", "range")
    class Instant(_message.Message):
        __slots__ = ("timestamp", "labels", "payload")
        class LabelsEntry(_message.Message):
            __slots__ = ("key", "value")
            KEY_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            key: str
            value: str
            def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
        TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
        LABELS_FIELD_NUMBER: _ClassVar[int]
        PAYLOAD_FIELD_NUMBER: _ClassVar[int]
        timestamp: _timestamp_pb2.Timestamp
        labels: _containers.ScalarMap[str, str]
        payload: _struct_pb2.Struct
        def __init__(self, timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., labels: _Optional[_Mapping[str, str]] = ..., payload: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...
    class Range(_message.Message):
        __slots__ = ("start", "end", "labels", "payload")
        class LabelsEntry(_message.Message):
            __slots__ = ("key", "value")
            KEY_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            key: str
            value: str
            def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
        START_FIELD_NUMBER: _ClassVar[int]
        END_FIELD_NUMBER: _ClassVar[int]
        LABELS_FIELD_NUMBER: _ClassVar[int]
        PAYLOAD_FIELD_NUMBER: _ClassVar[int]
        start: _timestamp_pb2.Timestamp
        end: _timestamp_pb2.Timestamp
        labels: _containers.ScalarMap[str, str]
        payload: _struct_pb2.Struct
        def __init__(self, start: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., end: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., labels: _Optional[_Mapping[str, str]] = ..., payload: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...
    INSTANT_FIELD_NUMBER: _ClassVar[int]
    RANGE_FIELD_NUMBER: _ClassVar[int]
    instant: AnnotationEvent.Instant
    range: AnnotationEvent.Range
    def __init__(self, instant: _Optional[_Union[AnnotationEvent.Instant, _Mapping]] = ..., range: _Optional[_Union[AnnotationEvent.Range, _Mapping]] = ...) -> None: ...
