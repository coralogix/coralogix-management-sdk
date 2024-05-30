from com.coralogix.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.dashboards.v1.common import search_filter_pb2 as _search_filter_pb2
from com.coralogix.dashboards.v1.common import time_frame_pb2 as _time_frame_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class FetchTopErroredRatioRequest(_message.Message):
    __slots__ = ("filters", "time_frame", "limit")
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    filters: _containers.RepeatedCompositeFieldContainer[_search_filter_pb2.SearchFilter]
    time_frame: _time_frame_pb2.TimeFrame
    limit: _wrappers_pb2.Int32Value
    def __init__(self, filters: _Optional[_Iterable[_Union[_search_filter_pb2.SearchFilter, _Mapping]]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class FetchTopErroredRatioResponse(_message.Message):
    __slots__ = ("ratios", "count_all_logs")
    RATIOS_FIELD_NUMBER: _ClassVar[int]
    COUNT_ALL_LOGS_FIELD_NUMBER: _ClassVar[int]
    ratios: _containers.RepeatedCompositeFieldContainer[TopErroredRatio]
    count_all_logs: _wrappers_pb2.Int64Value
    def __init__(self, ratios: _Optional[_Iterable[_Union[TopErroredRatio, _Mapping]]] = ..., count_all_logs: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class TopErroredRatio(_message.Message):
    __slots__ = ("application", "subsystem", "counts_per_severity", "count_all_severities")
    class CountsPerSeverityEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: int
        value: _wrappers_pb2.Int64Value
        def __init__(self, key: _Optional[int] = ..., value: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
    COUNTS_PER_SEVERITY_FIELD_NUMBER: _ClassVar[int]
    COUNT_ALL_SEVERITIES_FIELD_NUMBER: _ClassVar[int]
    application: _wrappers_pb2.StringValue
    subsystem: _wrappers_pb2.StringValue
    counts_per_severity: _containers.MessageMap[int, _wrappers_pb2.Int64Value]
    count_all_severities: _wrappers_pb2.Int64Value
    def __init__(self, application: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., counts_per_severity: _Optional[_Mapping[int, _wrappers_pb2.Int64Value]] = ..., count_all_severities: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...
