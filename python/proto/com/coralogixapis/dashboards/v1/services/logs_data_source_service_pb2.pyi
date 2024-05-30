from com.coralogixapis.dashboards.v1.ast import annotation_pb2 as _annotation_pb2
from com.coralogixapis.dashboards.v1.ast import filter_pb2 as _filter_pb2
from com.coralogixapis.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogixapis.dashboards.v1.common import annotation_event_pb2 as _annotation_event_pb2
from com.coralogixapis.dashboards.v1.common import group_pb2 as _group_pb2
from com.coralogixapis.dashboards.v1.common import grouped_series_pb2 as _grouped_series_pb2
from com.coralogixapis.dashboards.v1.common import log_severity_level_pb2 as _log_severity_level_pb2
from com.coralogixapis.dashboards.v1.common import logs_aggregation_pb2 as _logs_aggregation_pb2
from com.coralogixapis.dashboards.v1.common import order_direction_pb2 as _order_direction_pb2
from com.coralogixapis.dashboards.v1.common import ordering_field_pb2 as _ordering_field_pb2
from com.coralogixapis.dashboards.v1.common import pagination_pb2 as _pagination_pb2
from com.coralogixapis.dashboards.v1.common import query_pb2 as _query_pb2
from com.coralogixapis.dashboards.v1.common import time_frame_pb2 as _time_frame_pb2
from com.coralogixapis.dashboards.v1.common import time_series_pb2 as _time_series_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SearchLogsTimeSeriesRequest(_message.Message):
    __slots__ = ("time_frame", "interval", "filters", "group_by", "aggregations", "limit", "lucene_query")
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    INTERVAL_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_FIELD_NUMBER: _ClassVar[int]
    AGGREGATIONS_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    interval: _duration_pb2.Duration
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.LogsFilter]
    group_by: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    aggregations: _containers.RepeatedCompositeFieldContainer[_logs_aggregation_pb2.LogsAggregation]
    limit: _wrappers_pb2.Int32Value
    lucene_query: _wrappers_pb2.StringValue
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.LogsFilter, _Mapping]]] = ..., group_by: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., aggregations: _Optional[_Iterable[_Union[_logs_aggregation_pb2.LogsAggregation, _Mapping]]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., lucene_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class SearchLogsTimeSeriesResponse(_message.Message):
    __slots__ = ("time_series", "total")
    TIME_SERIES_FIELD_NUMBER: _ClassVar[int]
    TOTAL_FIELD_NUMBER: _ClassVar[int]
    time_series: _containers.RepeatedCompositeFieldContainer[_time_series_pb2.TimeSeries]
    total: _wrappers_pb2.Int64Value
    def __init__(self, time_series: _Optional[_Iterable[_Union[_time_series_pb2.TimeSeries, _Mapping]]] = ..., total: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class SearchLogsEventsRequest(_message.Message):
    __slots__ = ("time_frame", "filters", "lucene_query", "order_by", "pagination")
    class Pagination(_message.Message):
        __slots__ = ("offset", "limit")
        OFFSET_FIELD_NUMBER: _ClassVar[int]
        LIMIT_FIELD_NUMBER: _ClassVar[int]
        offset: _wrappers_pb2.Int32Value
        limit: _wrappers_pb2.Int32Value
        def __init__(self, offset: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_NUMBER: _ClassVar[int]
    PAGINATION_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.LogsFilter]
    lucene_query: _wrappers_pb2.StringValue
    order_by: _containers.RepeatedCompositeFieldContainer[_ordering_field_pb2.OrderingField]
    pagination: SearchLogsEventsRequest.Pagination
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.LogsFilter, _Mapping]]] = ..., lucene_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., order_by: _Optional[_Iterable[_Union[_ordering_field_pb2.OrderingField, _Mapping]]] = ..., pagination: _Optional[_Union[SearchLogsEventsRequest.Pagination, _Mapping]] = ...) -> None: ...

class SearchLogsEventsResponse(_message.Message):
    __slots__ = ("total", "events")
    TOTAL_FIELD_NUMBER: _ClassVar[int]
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    total: _wrappers_pb2.Int64Value
    events: _containers.RepeatedCompositeFieldContainer[LogsEvent]
    def __init__(self, total: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., events: _Optional[_Iterable[_Union[LogsEvent, _Mapping]]] = ...) -> None: ...

class SearchLogsEventGroupsRequest(_message.Message):
    __slots__ = ("time_frame", "filters", "lucene_query", "group_by", "aggregations", "order_by", "pagination")
    class OrderBy(_message.Message):
        __slots__ = ("grouping", "aggregation", "order_direction")
        GROUPING_FIELD_NUMBER: _ClassVar[int]
        AGGREGATION_FIELD_NUMBER: _ClassVar[int]
        ORDER_DIRECTION_FIELD_NUMBER: _ClassVar[int]
        grouping: _wrappers_pb2.StringValue
        aggregation: _logs_aggregation_pb2.LogsAggregation
        order_direction: _order_direction_pb2.OrderDirection
        def __init__(self, grouping: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., aggregation: _Optional[_Union[_logs_aggregation_pb2.LogsAggregation, _Mapping]] = ..., order_direction: _Optional[_Union[_order_direction_pb2.OrderDirection, str]] = ...) -> None: ...
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_FIELD_NUMBER: _ClassVar[int]
    AGGREGATIONS_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_NUMBER: _ClassVar[int]
    PAGINATION_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.LogsFilter]
    lucene_query: _wrappers_pb2.StringValue
    group_by: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    aggregations: _containers.RepeatedCompositeFieldContainer[_logs_aggregation_pb2.LogsAggregation]
    order_by: _containers.RepeatedCompositeFieldContainer[SearchLogsEventGroupsRequest.OrderBy]
    pagination: _pagination_pb2.Pagination
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.LogsFilter, _Mapping]]] = ..., lucene_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., group_by: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., aggregations: _Optional[_Iterable[_Union[_logs_aggregation_pb2.LogsAggregation, _Mapping]]] = ..., order_by: _Optional[_Iterable[_Union[SearchLogsEventGroupsRequest.OrderBy, _Mapping]]] = ..., pagination: _Optional[_Union[_pagination_pb2.Pagination, _Mapping]] = ...) -> None: ...

class SearchLogsEventGroupsResponse(_message.Message):
    __slots__ = ("groups",)
    GROUPS_FIELD_NUMBER: _ClassVar[int]
    groups: _containers.RepeatedCompositeFieldContainer[_group_pb2.MultiGroup]
    def __init__(self, groups: _Optional[_Iterable[_Union[_group_pb2.MultiGroup, _Mapping]]] = ...) -> None: ...

class SearchGroupedLogsSeriesRequest(_message.Message):
    __slots__ = ("time_frame", "filters", "group_by_fields", "aggregation", "lucene_query", "limits")
    class Limit(_message.Message):
        __slots__ = ("group_by_fields", "limit", "min_percentage")
        GROUP_BY_FIELDS_FIELD_NUMBER: _ClassVar[int]
        LIMIT_FIELD_NUMBER: _ClassVar[int]
        MIN_PERCENTAGE_FIELD_NUMBER: _ClassVar[int]
        group_by_fields: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        limit: _wrappers_pb2.Int32Value
        min_percentage: _wrappers_pb2.Int32Value
        def __init__(self, group_by_fields: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., min_percentage: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_FIELDS_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_FIELD_NUMBER: _ClassVar[int]
    LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
    LIMITS_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.LogsFilter]
    group_by_fields: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    aggregation: _logs_aggregation_pb2.LogsAggregation
    lucene_query: _wrappers_pb2.StringValue
    limits: _containers.RepeatedCompositeFieldContainer[SearchGroupedLogsSeriesRequest.Limit]
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.LogsFilter, _Mapping]]] = ..., group_by_fields: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., aggregation: _Optional[_Union[_logs_aggregation_pb2.LogsAggregation, _Mapping]] = ..., lucene_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., limits: _Optional[_Iterable[_Union[SearchGroupedLogsSeriesRequest.Limit, _Mapping]]] = ...) -> None: ...

class SearchGroupedLogsSeriesResponse(_message.Message):
    __slots__ = ("series",)
    SERIES_FIELD_NUMBER: _ClassVar[int]
    series: _containers.RepeatedCompositeFieldContainer[_grouped_series_pb2.GroupedSeries]
    def __init__(self, series: _Optional[_Iterable[_Union[_grouped_series_pb2.GroupedSeries, _Mapping]]] = ...) -> None: ...

class SearchLogsGroupedTimeSeriesRequest(_message.Message):
    __slots__ = ("time_frame", "interval", "filters", "aggregation", "lucene_query", "group_by")
    class GroupBy(_message.Message):
        __slots__ = ("fields", "limit", "min_percentage")
        FIELDS_FIELD_NUMBER: _ClassVar[int]
        LIMIT_FIELD_NUMBER: _ClassVar[int]
        MIN_PERCENTAGE_FIELD_NUMBER: _ClassVar[int]
        fields: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        limit: _wrappers_pb2.Int32Value
        min_percentage: _wrappers_pb2.Int32Value
        def __init__(self, fields: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., min_percentage: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    INTERVAL_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_FIELD_NUMBER: _ClassVar[int]
    LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    interval: _duration_pb2.Duration
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.LogsFilter]
    aggregation: _logs_aggregation_pb2.LogsAggregation
    lucene_query: _wrappers_pb2.StringValue
    group_by: _containers.RepeatedCompositeFieldContainer[SearchLogsGroupedTimeSeriesRequest.GroupBy]
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.LogsFilter, _Mapping]]] = ..., aggregation: _Optional[_Union[_logs_aggregation_pb2.LogsAggregation, _Mapping]] = ..., lucene_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., group_by: _Optional[_Iterable[_Union[SearchLogsGroupedTimeSeriesRequest.GroupBy, _Mapping]]] = ...) -> None: ...

class SearchLogsGroupedTimeSeriesResponse(_message.Message):
    __slots__ = ("series",)
    SERIES_FIELD_NUMBER: _ClassVar[int]
    series: _containers.RepeatedCompositeFieldContainer[_grouped_series_pb2.GroupedSeries]
    def __init__(self, series: _Optional[_Iterable[_Union[_grouped_series_pb2.GroupedSeries, _Mapping]]] = ...) -> None: ...

class LogsEvent(_message.Message):
    __slots__ = ("log_id", "timestamp", "text", "json", "logs_metadata")
    LOG_ID_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    JSON_FIELD_NUMBER: _ClassVar[int]
    LOGS_METADATA_FIELD_NUMBER: _ClassVar[int]
    log_id: _wrappers_pb2.StringValue
    timestamp: _timestamp_pb2.Timestamp
    text: _wrappers_pb2.StringValue
    json: _struct_pb2.Struct
    logs_metadata: LogsMetadata
    def __init__(self, log_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., text: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., json: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., logs_metadata: _Optional[_Union[LogsMetadata, _Mapping]] = ...) -> None: ...

class LogsMetadata(_message.Message):
    __slots__ = ("application_name", "subsystem_name", "severity")
    APPLICATION_NAME_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_NAME_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    application_name: _wrappers_pb2.StringValue
    subsystem_name: _wrappers_pb2.StringValue
    severity: _log_severity_level_pb2.LogSeverityLevel
    def __init__(self, application_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., severity: _Optional[_Union[_log_severity_level_pb2.LogSeverityLevel, str]] = ...) -> None: ...

class SearchLogsTimeValueRequest(_message.Message):
    __slots__ = ("time_frame", "filters", "lucene_query", "aggregation")
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.LogsFilter]
    lucene_query: _wrappers_pb2.StringValue
    aggregation: _logs_aggregation_pb2.LogsAggregation
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.LogsFilter, _Mapping]]] = ..., lucene_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., aggregation: _Optional[_Union[_logs_aggregation_pb2.LogsAggregation, _Mapping]] = ...) -> None: ...

class SearchLogsTimeValueResponse(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: _wrappers_pb2.DoubleValue
    def __init__(self, value: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ...) -> None: ...

class SearchLogsAnnotationEventsRequest(_message.Message):
    __slots__ = ("time_frame", "dataprime_query", "strategy", "limit")
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    DATAPRIME_QUERY_FIELD_NUMBER: _ClassVar[int]
    STRATEGY_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    dataprime_query: _query_pb2.FullDataprimeQuery
    strategy: _annotation_pb2.Annotation.LogsSource.Strategy
    limit: _wrappers_pb2.Int32Value
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., dataprime_query: _Optional[_Union[_query_pb2.FullDataprimeQuery, _Mapping]] = ..., strategy: _Optional[_Union[_annotation_pb2.Annotation.LogsSource.Strategy, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class SearchLogsAnnotationEventsResponse(_message.Message):
    __slots__ = ("annotation_events",)
    ANNOTATION_EVENTS_FIELD_NUMBER: _ClassVar[int]
    annotation_events: _containers.RepeatedCompositeFieldContainer[_annotation_event_pb2.AnnotationEvent]
    def __init__(self, annotation_events: _Optional[_Iterable[_Union[_annotation_event_pb2.AnnotationEvent, _Mapping]]] = ...) -> None: ...
