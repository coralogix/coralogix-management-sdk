from com.coralogix.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.dashboards.v1.common import search_filter_pb2 as _search_filter_pb2
from com.coralogix.dashboards.v1.common import time_frame_pb2 as _time_frame_pb2
from com.coralogix.dashboards.v1.common import time_series_pb2 as _time_series_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SearchLogsActivityRequest(_message.Message):
    __slots__ = ("filters", "time_frame", "interval")
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    INTERVAL_FIELD_NUMBER: _ClassVar[int]
    filters: _containers.RepeatedCompositeFieldContainer[_search_filter_pb2.SearchFilter]
    time_frame: _time_frame_pb2.TimeFrame
    interval: _duration_pb2.Duration
    def __init__(self, filters: _Optional[_Iterable[_Union[_search_filter_pb2.SearchFilter, _Mapping]]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ...) -> None: ...

class SearchLogsActivityResponse(_message.Message):
    __slots__ = ("time_series", "tags")
    TIME_SERIES_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    time_series: _containers.RepeatedCompositeFieldContainer[_time_series_pb2.TimeSeries]
    tags: _containers.RepeatedCompositeFieldContainer[Tag]
    def __init__(self, time_series: _Optional[_Iterable[_Union[_time_series_pb2.TimeSeries, _Mapping]]] = ..., tags: _Optional[_Iterable[_Union[Tag, _Mapping]]] = ...) -> None: ...

class Tag(_message.Message):
    __slots__ = ("id", "occurred_on", "avatar", "title", "description")
    ID_FIELD_NUMBER: _ClassVar[int]
    OCCURRED_ON_FIELD_NUMBER: _ClassVar[int]
    AVATAR_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.Int32Value
    occurred_on: _timestamp_pb2.Timestamp
    avatar: _wrappers_pb2.StringValue
    title: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., occurred_on: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., avatar: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., title: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
