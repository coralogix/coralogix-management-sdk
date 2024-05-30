from com.coralogix.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.dashboards.v1.common import search_filter_pb2 as _search_filter_pb2
from com.coralogix.dashboards.v1.common import time_frame_pb2 as _time_frame_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CountLogsRequest(_message.Message):
    __slots__ = ("filters", "time_frame", "interval", "severities")
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    INTERVAL_FIELD_NUMBER: _ClassVar[int]
    SEVERITIES_FIELD_NUMBER: _ClassVar[int]
    filters: _containers.RepeatedCompositeFieldContainer[_search_filter_pb2.SearchFilter]
    time_frame: _time_frame_pb2.TimeFrame
    interval: _duration_pb2.Duration
    severities: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.Int32Value]
    def __init__(self, filters: _Optional[_Iterable[_Union[_search_filter_pb2.SearchFilter, _Mapping]]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., severities: _Optional[_Iterable[_Union[_wrappers_pb2.Int32Value, _Mapping]]] = ...) -> None: ...

class CountLogsResponse(_message.Message):
    __slots__ = ("time_series", "total_count")
    TIME_SERIES_FIELD_NUMBER: _ClassVar[int]
    TOTAL_COUNT_FIELD_NUMBER: _ClassVar[int]
    time_series: HistogramTimeSeries
    total_count: _wrappers_pb2.Int64Value
    def __init__(self, time_series: _Optional[_Union[HistogramTimeSeries, _Mapping]] = ..., total_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class GetMetricsStatsRequest(_message.Message):
    __slots__ = ("filters", "time_frame", "interval")
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    INTERVAL_FIELD_NUMBER: _ClassVar[int]
    filters: _containers.RepeatedCompositeFieldContainer[_search_filter_pb2.SearchFilter]
    time_frame: _time_frame_pb2.TimeFrame
    interval: _duration_pb2.Duration
    def __init__(self, filters: _Optional[_Iterable[_Union[_search_filter_pb2.SearchFilter, _Mapping]]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ...) -> None: ...

class GetMetricsStatsResponse(_message.Message):
    __slots__ = ("time_series",)
    TIME_SERIES_FIELD_NUMBER: _ClassVar[int]
    time_series: HistogramTimeSeries
    def __init__(self, time_series: _Optional[_Union[HistogramTimeSeries, _Mapping]] = ...) -> None: ...

class CountSpansRequest(_message.Message):
    __slots__ = ("filters", "time_frame", "interval")
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    INTERVAL_FIELD_NUMBER: _ClassVar[int]
    filters: _containers.RepeatedCompositeFieldContainer[_search_filter_pb2.SearchFilter]
    time_frame: _time_frame_pb2.TimeFrame
    interval: _duration_pb2.Duration
    def __init__(self, filters: _Optional[_Iterable[_Union[_search_filter_pb2.SearchFilter, _Mapping]]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ...) -> None: ...

class CountSpansResponse(_message.Message):
    __slots__ = ("time_series", "total_count")
    TIME_SERIES_FIELD_NUMBER: _ClassVar[int]
    TOTAL_COUNT_FIELD_NUMBER: _ClassVar[int]
    time_series: HistogramTimeSeries
    total_count: _wrappers_pb2.Int64Value
    def __init__(self, time_series: _Optional[_Union[HistogramTimeSeries, _Mapping]] = ..., total_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class CountFailedXDRTestsRequest(_message.Message):
    __slots__ = ("filters", "time_frame", "interval")
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    INTERVAL_FIELD_NUMBER: _ClassVar[int]
    filters: _containers.RepeatedCompositeFieldContainer[_search_filter_pb2.SearchFilter]
    time_frame: _time_frame_pb2.TimeFrame
    interval: _duration_pb2.Duration
    def __init__(self, filters: _Optional[_Iterable[_Union[_search_filter_pb2.SearchFilter, _Mapping]]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ...) -> None: ...

class CountFailedXDRTestsResponse(_message.Message):
    __slots__ = ("time_series", "total_count")
    TIME_SERIES_FIELD_NUMBER: _ClassVar[int]
    TOTAL_COUNT_FIELD_NUMBER: _ClassVar[int]
    time_series: HistogramTimeSeries
    total_count: _wrappers_pb2.Int64Value
    def __init__(self, time_series: _Optional[_Union[HistogramTimeSeries, _Mapping]] = ..., total_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class HistogramTimeSeries(_message.Message):
    __slots__ = ("values",)
    VALUES_FIELD_NUMBER: _ClassVar[int]
    values: _containers.RepeatedCompositeFieldContainer[HistogramDataPoint]
    def __init__(self, values: _Optional[_Iterable[_Union[HistogramDataPoint, _Mapping]]] = ...) -> None: ...

class HistogramDataPoint(_message.Message):
    __slots__ = ("timestamp", "values")
    class ValuesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: _wrappers_pb2.DoubleValue
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ...) -> None: ...
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    VALUES_FIELD_NUMBER: _ClassVar[int]
    timestamp: _timestamp_pb2.Timestamp
    values: _containers.MessageMap[str, _wrappers_pb2.DoubleValue]
    def __init__(self, timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., values: _Optional[_Mapping[str, _wrappers_pb2.DoubleValue]] = ...) -> None: ...
