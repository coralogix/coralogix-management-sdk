from com.coralogixapis.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogixapis.dashboards.v1.common import group_pb2 as _group_pb2
from com.coralogixapis.dashboards.v1.common import group_limit_pb2 as _group_limit_pb2
from com.coralogixapis.dashboards.v1.common import grouped_series_pb2 as _grouped_series_pb2
from com.coralogixapis.dashboards.v1.common import pagination_pb2 as _pagination_pb2
from com.coralogixapis.dashboards.v1.common import query_pb2 as _query_pb2
from com.coralogixapis.dashboards.v1.common import time_frame_pb2 as _time_frame_pb2
from com.coralogixapis.dashboards.v1.common import time_series_pb2 as _time_series_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SearchArchiveSpansTimeSeriesRequest(_message.Message):
    __slots__ = ("widget_id", "time_frame", "limit", "request_params_hash", "query", "query_raw", "aggregation_keys")
    WIDGET_ID_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    REQUEST_PARAMS_HASH_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    QUERY_RAW_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_KEYS_FIELD_NUMBER: _ClassVar[int]
    widget_id: _wrappers_pb2.StringValue
    time_frame: _time_frame_pb2.TimeFrame
    limit: _wrappers_pb2.Int32Value
    request_params_hash: _wrappers_pb2.StringValue
    query: _query_pb2.SerializedDataprimeQuery
    query_raw: _query_pb2.DataprimeQuery
    aggregation_keys: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, widget_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., request_params_hash: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., query: _Optional[_Union[_query_pb2.SerializedDataprimeQuery, _Mapping]] = ..., query_raw: _Optional[_Union[_query_pb2.DataprimeQuery, _Mapping]] = ..., aggregation_keys: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class SearchArchiveSpansTimeSeriesResponse(_message.Message):
    __slots__ = ("time_series", "total")
    TIME_SERIES_FIELD_NUMBER: _ClassVar[int]
    TOTAL_FIELD_NUMBER: _ClassVar[int]
    time_series: _containers.RepeatedCompositeFieldContainer[_time_series_pb2.TimeSeries]
    total: _wrappers_pb2.Int64Value
    def __init__(self, time_series: _Optional[_Iterable[_Union[_time_series_pb2.TimeSeries, _Mapping]]] = ..., total: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class SearchArchiveSpansEventsRequest(_message.Message):
    __slots__ = ("widget_id", "time_frame", "pagination", "request_params_hash", "query", "query_raw")
    WIDGET_ID_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    PAGINATION_FIELD_NUMBER: _ClassVar[int]
    REQUEST_PARAMS_HASH_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    QUERY_RAW_FIELD_NUMBER: _ClassVar[int]
    widget_id: _wrappers_pb2.StringValue
    time_frame: _time_frame_pb2.TimeFrame
    pagination: _pagination_pb2.Pagination
    request_params_hash: _wrappers_pb2.StringValue
    query: _query_pb2.SerializedDataprimeQuery
    query_raw: _query_pb2.DataprimeQuery
    def __init__(self, widget_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., pagination: _Optional[_Union[_pagination_pb2.Pagination, _Mapping]] = ..., request_params_hash: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., query: _Optional[_Union[_query_pb2.SerializedDataprimeQuery, _Mapping]] = ..., query_raw: _Optional[_Union[_query_pb2.DataprimeQuery, _Mapping]] = ...) -> None: ...

class SearchArchiveSpansEventsResponse(_message.Message):
    __slots__ = ("events", "total_fetched")
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    TOTAL_FETCHED_FIELD_NUMBER: _ClassVar[int]
    events: _containers.RepeatedCompositeFieldContainer[ArchiveSpansEvent]
    total_fetched: _wrappers_pb2.Int32Value
    def __init__(self, events: _Optional[_Iterable[_Union[ArchiveSpansEvent, _Mapping]]] = ..., total_fetched: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class SearchArchiveSpansEventsCountRequest(_message.Message):
    __slots__ = ("widget_id", "time_frame", "request_params_hash", "query", "query_raw")
    WIDGET_ID_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    REQUEST_PARAMS_HASH_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    QUERY_RAW_FIELD_NUMBER: _ClassVar[int]
    widget_id: _wrappers_pb2.StringValue
    time_frame: _time_frame_pb2.TimeFrame
    request_params_hash: _wrappers_pb2.StringValue
    query: _query_pb2.SerializedDataprimeQuery
    query_raw: _query_pb2.DataprimeQuery
    def __init__(self, widget_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., request_params_hash: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., query: _Optional[_Union[_query_pb2.SerializedDataprimeQuery, _Mapping]] = ..., query_raw: _Optional[_Union[_query_pb2.DataprimeQuery, _Mapping]] = ...) -> None: ...

class SearchArchiveSpansEventsCountResponse(_message.Message):
    __slots__ = ("count",)
    COUNT_FIELD_NUMBER: _ClassVar[int]
    count: _wrappers_pb2.Int64Value
    def __init__(self, count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class ArchiveSpansEvent(_message.Message):
    __slots__ = ("span_id", "trace_id", "parent_span_id", "metadata", "start_time", "duration", "tags", "process_tags", "logs")
    class Metadata(_message.Message):
        __slots__ = ("application_name", "subsystem_name", "service_name", "operation_name")
        APPLICATION_NAME_FIELD_NUMBER: _ClassVar[int]
        SUBSYSTEM_NAME_FIELD_NUMBER: _ClassVar[int]
        SERVICE_NAME_FIELD_NUMBER: _ClassVar[int]
        OPERATION_NAME_FIELD_NUMBER: _ClassVar[int]
        application_name: _wrappers_pb2.StringValue
        subsystem_name: _wrappers_pb2.StringValue
        service_name: _wrappers_pb2.StringValue
        operation_name: _wrappers_pb2.StringValue
        def __init__(self, application_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., service_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., operation_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class Tag(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: _wrappers_pb2.StringValue
        value: _wrappers_pb2.StringValue
        def __init__(self, key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class Log(_message.Message):
        __slots__ = ("timestamp", "fields")
        class FieldsEntry(_message.Message):
            __slots__ = ("key", "value")
            KEY_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            key: str
            value: _wrappers_pb2.StringValue
            def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
        TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
        FIELDS_FIELD_NUMBER: _ClassVar[int]
        timestamp: _timestamp_pb2.Timestamp
        fields: _containers.MessageMap[str, _wrappers_pb2.StringValue]
        def __init__(self, timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., fields: _Optional[_Mapping[str, _wrappers_pb2.StringValue]] = ...) -> None: ...
    SPAN_ID_FIELD_NUMBER: _ClassVar[int]
    TRACE_ID_FIELD_NUMBER: _ClassVar[int]
    PARENT_SPAN_ID_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    START_TIME_FIELD_NUMBER: _ClassVar[int]
    DURATION_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    PROCESS_TAGS_FIELD_NUMBER: _ClassVar[int]
    LOGS_FIELD_NUMBER: _ClassVar[int]
    span_id: _wrappers_pb2.StringValue
    trace_id: _wrappers_pb2.StringValue
    parent_span_id: _wrappers_pb2.StringValue
    metadata: ArchiveSpansEvent.Metadata
    start_time: _timestamp_pb2.Timestamp
    duration: _duration_pb2.Duration
    tags: _containers.RepeatedCompositeFieldContainer[ArchiveSpansEvent.Tag]
    process_tags: _containers.RepeatedCompositeFieldContainer[ArchiveSpansEvent.Tag]
    logs: _containers.RepeatedCompositeFieldContainer[ArchiveSpansEvent.Log]
    def __init__(self, span_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., trace_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., parent_span_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., metadata: _Optional[_Union[ArchiveSpansEvent.Metadata, _Mapping]] = ..., start_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., duration: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., tags: _Optional[_Iterable[_Union[ArchiveSpansEvent.Tag, _Mapping]]] = ..., process_tags: _Optional[_Iterable[_Union[ArchiveSpansEvent.Tag, _Mapping]]] = ..., logs: _Optional[_Iterable[_Union[ArchiveSpansEvent.Log, _Mapping]]] = ...) -> None: ...

class SearchArchiveSpansEventGroupsRequest(_message.Message):
    __slots__ = ("widget_id", "time_frame", "request_params_hash", "query", "query_raw", "group_by_keys", "aggregation_keys", "pagination")
    WIDGET_ID_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    REQUEST_PARAMS_HASH_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    QUERY_RAW_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_KEYS_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_KEYS_FIELD_NUMBER: _ClassVar[int]
    PAGINATION_FIELD_NUMBER: _ClassVar[int]
    widget_id: _wrappers_pb2.StringValue
    time_frame: _time_frame_pb2.TimeFrame
    request_params_hash: _wrappers_pb2.StringValue
    query: _query_pb2.SerializedDataprimeQuery
    query_raw: _query_pb2.DataprimeQuery
    group_by_keys: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    aggregation_keys: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    pagination: _pagination_pb2.Pagination
    def __init__(self, widget_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., request_params_hash: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., query: _Optional[_Union[_query_pb2.SerializedDataprimeQuery, _Mapping]] = ..., query_raw: _Optional[_Union[_query_pb2.DataprimeQuery, _Mapping]] = ..., group_by_keys: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., aggregation_keys: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., pagination: _Optional[_Union[_pagination_pb2.Pagination, _Mapping]] = ...) -> None: ...

class SearchArchiveSpansEventGroupsResponse(_message.Message):
    __slots__ = ("groups",)
    GROUPS_FIELD_NUMBER: _ClassVar[int]
    groups: _containers.RepeatedCompositeFieldContainer[_group_pb2.MultiGroup]
    def __init__(self, groups: _Optional[_Iterable[_Union[_group_pb2.MultiGroup, _Mapping]]] = ...) -> None: ...

class SearchArchiveGroupedSpansSeriesRequest(_message.Message):
    __slots__ = ("widget_id", "time_frame", "request_params_hash", "query", "query_raw", "group_by_keys", "aggregation_key", "limits")
    WIDGET_ID_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    REQUEST_PARAMS_HASH_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    QUERY_RAW_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_KEYS_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_KEY_FIELD_NUMBER: _ClassVar[int]
    LIMITS_FIELD_NUMBER: _ClassVar[int]
    widget_id: _wrappers_pb2.StringValue
    time_frame: _time_frame_pb2.TimeFrame
    request_params_hash: _wrappers_pb2.StringValue
    query: _query_pb2.SerializedDataprimeQuery
    query_raw: _query_pb2.DataprimeQuery
    group_by_keys: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    aggregation_key: _wrappers_pb2.StringValue
    limits: _containers.RepeatedCompositeFieldContainer[_group_limit_pb2.GroupLimit]
    def __init__(self, widget_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., request_params_hash: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., query: _Optional[_Union[_query_pb2.SerializedDataprimeQuery, _Mapping]] = ..., query_raw: _Optional[_Union[_query_pb2.DataprimeQuery, _Mapping]] = ..., group_by_keys: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., aggregation_key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., limits: _Optional[_Iterable[_Union[_group_limit_pb2.GroupLimit, _Mapping]]] = ...) -> None: ...

class SearchArchiveGroupedSpansSeriesResponse(_message.Message):
    __slots__ = ("series",)
    SERIES_FIELD_NUMBER: _ClassVar[int]
    series: _containers.RepeatedCompositeFieldContainer[_grouped_series_pb2.GroupedSeries]
    def __init__(self, series: _Optional[_Iterable[_Union[_grouped_series_pb2.GroupedSeries, _Mapping]]] = ...) -> None: ...

class SearchArchiveGroupedSpansTimeSeriesRequest(_message.Message):
    __slots__ = ("widget_id", "time_frame", "request_params_hash", "query", "query_raw", "group_by_keys", "aggregation_key", "limits")
    WIDGET_ID_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    REQUEST_PARAMS_HASH_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    QUERY_RAW_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_KEYS_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_KEY_FIELD_NUMBER: _ClassVar[int]
    LIMITS_FIELD_NUMBER: _ClassVar[int]
    widget_id: _wrappers_pb2.StringValue
    time_frame: _time_frame_pb2.TimeFrame
    request_params_hash: _wrappers_pb2.StringValue
    query: _query_pb2.SerializedDataprimeQuery
    query_raw: _query_pb2.DataprimeQuery
    group_by_keys: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    aggregation_key: _wrappers_pb2.StringValue
    limits: _containers.RepeatedCompositeFieldContainer[_group_limit_pb2.GroupLimit]
    def __init__(self, widget_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., request_params_hash: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., query: _Optional[_Union[_query_pb2.SerializedDataprimeQuery, _Mapping]] = ..., query_raw: _Optional[_Union[_query_pb2.DataprimeQuery, _Mapping]] = ..., group_by_keys: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., aggregation_key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., limits: _Optional[_Iterable[_Union[_group_limit_pb2.GroupLimit, _Mapping]]] = ...) -> None: ...

class SearchArchiveGroupedSpansTimeSeriesResponse(_message.Message):
    __slots__ = ("series",)
    SERIES_FIELD_NUMBER: _ClassVar[int]
    series: _grouped_series_pb2.GroupedSeries
    def __init__(self, series: _Optional[_Union[_grouped_series_pb2.GroupedSeries, _Mapping]] = ...) -> None: ...

class SearchArchiveSpansTimeValueRequest(_message.Message):
    __slots__ = ("widget_id", "time_frame", "request_params_hash", "query", "query_raw", "aggregation_key")
    WIDGET_ID_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    REQUEST_PARAMS_HASH_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    QUERY_RAW_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_KEY_FIELD_NUMBER: _ClassVar[int]
    widget_id: _wrappers_pb2.StringValue
    time_frame: _time_frame_pb2.TimeFrame
    request_params_hash: _wrappers_pb2.StringValue
    query: _query_pb2.SerializedDataprimeQuery
    query_raw: _query_pb2.DataprimeQuery
    aggregation_key: _wrappers_pb2.StringValue
    def __init__(self, widget_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., request_params_hash: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., query: _Optional[_Union[_query_pb2.SerializedDataprimeQuery, _Mapping]] = ..., query_raw: _Optional[_Union[_query_pb2.DataprimeQuery, _Mapping]] = ..., aggregation_key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class SearchArchiveSpansTimeValueResponse(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: _wrappers_pb2.DoubleValue
    def __init__(self, value: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ...) -> None: ...
