from com.coralogix.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.dashboards.v1.common import log_severity_level_pb2 as _log_severity_level_pb2
from com.coralogix.dashboards.v1.common import search_filter_pb2 as _search_filter_pb2
from com.coralogix.dashboards.v1.common import time_frame_pb2 as _time_frame_pb2
from com.coralogix.dashboards.v1.common import time_series_pb2 as _time_series_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SearchTopAbnormalErrorsRequest(_message.Message):
    __slots__ = ("filters", "time_frame", "limit", "aggregation_interval")
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_INTERVAL_FIELD_NUMBER: _ClassVar[int]
    filters: _containers.RepeatedCompositeFieldContainer[_search_filter_pb2.SearchFilter]
    time_frame: _time_frame_pb2.TimeFrame
    limit: _wrappers_pb2.Int32Value
    aggregation_interval: _duration_pb2.Duration
    def __init__(self, filters: _Optional[_Iterable[_Union[_search_filter_pb2.SearchFilter, _Mapping]]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., aggregation_interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ...) -> None: ...

class SearchTopNewerErrorsRequest(_message.Message):
    __slots__ = ("filters", "time_frame", "limit", "aggregation_interval")
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_INTERVAL_FIELD_NUMBER: _ClassVar[int]
    filters: _containers.RepeatedCompositeFieldContainer[_search_filter_pb2.SearchFilter]
    time_frame: _time_frame_pb2.TimeFrame
    limit: _wrappers_pb2.Int32Value
    aggregation_interval: _duration_pb2.Duration
    def __init__(self, filters: _Optional[_Iterable[_Union[_search_filter_pb2.SearchFilter, _Mapping]]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., aggregation_interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ...) -> None: ...

class SearchTopAbnormalErrorsResponse(_message.Message):
    __slots__ = ("new_templates", "abnormal_templates")
    class Template(_message.Message):
        __slots__ = ("template_id", "template_type", "log_severity_level", "actual_count", "normal_count", "time_series", "application", "subsystem", "log_example", "created_at")
        TEMPLATE_ID_FIELD_NUMBER: _ClassVar[int]
        TEMPLATE_TYPE_FIELD_NUMBER: _ClassVar[int]
        LOG_SEVERITY_LEVEL_FIELD_NUMBER: _ClassVar[int]
        ACTUAL_COUNT_FIELD_NUMBER: _ClassVar[int]
        NORMAL_COUNT_FIELD_NUMBER: _ClassVar[int]
        TIME_SERIES_FIELD_NUMBER: _ClassVar[int]
        APPLICATION_FIELD_NUMBER: _ClassVar[int]
        SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
        LOG_EXAMPLE_FIELD_NUMBER: _ClassVar[int]
        CREATED_AT_FIELD_NUMBER: _ClassVar[int]
        template_id: _wrappers_pb2.StringValue
        template_type: _wrappers_pb2.Int32Value
        log_severity_level: _log_severity_level_pb2.LogSeverityLevel
        actual_count: _wrappers_pb2.Int64Value
        normal_count: _wrappers_pb2.Int64Value
        time_series: _time_series_pb2.TimeSeries
        application: _wrappers_pb2.StringValue
        subsystem: _wrappers_pb2.StringValue
        log_example: _wrappers_pb2.StringValue
        created_at: _timestamp_pb2.Timestamp
        def __init__(self, template_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., template_type: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., log_severity_level: _Optional[_Union[_log_severity_level_pb2.LogSeverityLevel, str]] = ..., actual_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., normal_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., time_series: _Optional[_Union[_time_series_pb2.TimeSeries, _Mapping]] = ..., application: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., log_example: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
    NEW_TEMPLATES_FIELD_NUMBER: _ClassVar[int]
    ABNORMAL_TEMPLATES_FIELD_NUMBER: _ClassVar[int]
    new_templates: _containers.RepeatedCompositeFieldContainer[SearchTopAbnormalErrorsResponse.Template]
    abnormal_templates: _containers.RepeatedCompositeFieldContainer[SearchTopAbnormalErrorsResponse.Template]
    def __init__(self, new_templates: _Optional[_Iterable[_Union[SearchTopAbnormalErrorsResponse.Template, _Mapping]]] = ..., abnormal_templates: _Optional[_Iterable[_Union[SearchTopAbnormalErrorsResponse.Template, _Mapping]]] = ...) -> None: ...

class SearchTopNewerErrorsResponse(_message.Message):
    __slots__ = ("templates",)
    class Template(_message.Message):
        __slots__ = ("template_id", "template_type", "log_severity_level", "count", "time_series", "application", "subsystem", "log_example", "created_at")
        TEMPLATE_ID_FIELD_NUMBER: _ClassVar[int]
        TEMPLATE_TYPE_FIELD_NUMBER: _ClassVar[int]
        LOG_SEVERITY_LEVEL_FIELD_NUMBER: _ClassVar[int]
        COUNT_FIELD_NUMBER: _ClassVar[int]
        TIME_SERIES_FIELD_NUMBER: _ClassVar[int]
        APPLICATION_FIELD_NUMBER: _ClassVar[int]
        SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
        LOG_EXAMPLE_FIELD_NUMBER: _ClassVar[int]
        CREATED_AT_FIELD_NUMBER: _ClassVar[int]
        template_id: _wrappers_pb2.StringValue
        template_type: _wrappers_pb2.Int32Value
        log_severity_level: _log_severity_level_pb2.LogSeverityLevel
        count: _wrappers_pb2.Int64Value
        time_series: _time_series_pb2.TimeSeries
        application: _wrappers_pb2.StringValue
        subsystem: _wrappers_pb2.StringValue
        log_example: _wrappers_pb2.StringValue
        created_at: _timestamp_pb2.Timestamp
        def __init__(self, template_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., template_type: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., log_severity_level: _Optional[_Union[_log_severity_level_pb2.LogSeverityLevel, str]] = ..., count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., time_series: _Optional[_Union[_time_series_pb2.TimeSeries, _Mapping]] = ..., application: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., log_example: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
    TEMPLATES_FIELD_NUMBER: _ClassVar[int]
    templates: _containers.RepeatedCompositeFieldContainer[SearchTopNewerErrorsResponse.Template]
    def __init__(self, templates: _Optional[_Iterable[_Union[SearchTopNewerErrorsResponse.Template, _Mapping]]] = ...) -> None: ...
