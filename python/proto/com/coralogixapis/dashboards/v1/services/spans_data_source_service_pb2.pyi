from com.coralogixapis.dashboards.v1.ast import annotation_pb2 as _annotation_pb2
from com.coralogixapis.dashboards.v1.ast import filter_pb2 as _filter_pb2
from com.coralogixapis.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogixapis.dashboards.v1.common import annotation_event_pb2 as _annotation_event_pb2
from com.coralogixapis.dashboards.v1.common import group_pb2 as _group_pb2
from com.coralogixapis.dashboards.v1.common import grouped_series_pb2 as _grouped_series_pb2
from com.coralogixapis.dashboards.v1.common import pagination_pb2 as _pagination_pb2
from com.coralogixapis.dashboards.v1.common import span_field_pb2 as _span_field_pb2
from com.coralogixapis.dashboards.v1.common import spans_aggregation_pb2 as _spans_aggregation_pb2
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

class SearchSpansTimeSeriesRequest(_message.Message):
    __slots__ = ("time_frame", "interval", "filters", "lucene_query", "group_by", "aggregations", "limit")
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    INTERVAL_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_FIELD_NUMBER: _ClassVar[int]
    AGGREGATIONS_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    interval: _duration_pb2.Duration
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.SpansFilter]
    lucene_query: _wrappers_pb2.StringValue
    group_by: _containers.RepeatedCompositeFieldContainer[_span_field_pb2.SpanField]
    aggregations: _containers.RepeatedCompositeFieldContainer[_spans_aggregation_pb2.SpansAggregation]
    limit: _wrappers_pb2.Int32Value
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.SpansFilter, _Mapping]]] = ..., lucene_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., group_by: _Optional[_Iterable[_Union[_span_field_pb2.SpanField, _Mapping]]] = ..., aggregations: _Optional[_Iterable[_Union[_spans_aggregation_pb2.SpansAggregation, _Mapping]]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class SearchSpansTimeSeriesResponse(_message.Message):
    __slots__ = ("time_series", "total")
    TIME_SERIES_FIELD_NUMBER: _ClassVar[int]
    TOTAL_FIELD_NUMBER: _ClassVar[int]
    time_series: _containers.RepeatedCompositeFieldContainer[_time_series_pb2.TimeSeries]
    total: _wrappers_pb2.Int64Value
    def __init__(self, time_series: _Optional[_Iterable[_Union[_time_series_pb2.TimeSeries, _Mapping]]] = ..., total: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class SearchSpansEventsRequest(_message.Message):
    __slots__ = ("time_frame", "filters", "lucene_query", "pagination")
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
    PAGINATION_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.SpansFilter]
    lucene_query: _wrappers_pb2.StringValue
    pagination: SearchSpansEventsRequest.Pagination
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.SpansFilter, _Mapping]]] = ..., lucene_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., pagination: _Optional[_Union[SearchSpansEventsRequest.Pagination, _Mapping]] = ...) -> None: ...

class SearchSpansEventsResponse(_message.Message):
    __slots__ = ("total", "spans_events")
    TOTAL_FIELD_NUMBER: _ClassVar[int]
    SPANS_EVENTS_FIELD_NUMBER: _ClassVar[int]
    total: _wrappers_pb2.Int64Value
    spans_events: _containers.RepeatedCompositeFieldContainer[SpansEvent]
    def __init__(self, total: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., spans_events: _Optional[_Iterable[_Union[SpansEvent, _Mapping]]] = ...) -> None: ...

class SearchSpansEventGroupsRequest(_message.Message):
    __slots__ = ("time_frame", "filters", "lucene_query", "group_by_fields", "aggregations", "pagination")
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_FIELDS_FIELD_NUMBER: _ClassVar[int]
    AGGREGATIONS_FIELD_NUMBER: _ClassVar[int]
    PAGINATION_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.SpansFilter]
    lucene_query: _wrappers_pb2.StringValue
    group_by_fields: _containers.RepeatedCompositeFieldContainer[_span_field_pb2.SpanField]
    aggregations: _containers.RepeatedCompositeFieldContainer[_spans_aggregation_pb2.SpansAggregation]
    pagination: _pagination_pb2.Pagination
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.SpansFilter, _Mapping]]] = ..., lucene_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., group_by_fields: _Optional[_Iterable[_Union[_span_field_pb2.SpanField, _Mapping]]] = ..., aggregations: _Optional[_Iterable[_Union[_spans_aggregation_pb2.SpansAggregation, _Mapping]]] = ..., pagination: _Optional[_Union[_pagination_pb2.Pagination, _Mapping]] = ...) -> None: ...

class SearchSpansEventGroupsResponse(_message.Message):
    __slots__ = ("groups",)
    GROUPS_FIELD_NUMBER: _ClassVar[int]
    groups: _containers.RepeatedCompositeFieldContainer[_group_pb2.MultiGroup]
    def __init__(self, groups: _Optional[_Iterable[_Union[_group_pb2.MultiGroup, _Mapping]]] = ...) -> None: ...

class SearchGroupedSpansSeriesRequest(_message.Message):
    __slots__ = ("time_frame", "filters", "group_by_fields", "aggregation", "lucene_query", "limits")
    class Limit(_message.Message):
        __slots__ = ("group_by_fields", "limit", "min_percentage")
        GROUP_BY_FIELDS_FIELD_NUMBER: _ClassVar[int]
        LIMIT_FIELD_NUMBER: _ClassVar[int]
        MIN_PERCENTAGE_FIELD_NUMBER: _ClassVar[int]
        group_by_fields: _containers.RepeatedCompositeFieldContainer[_span_field_pb2.SpanField]
        limit: _wrappers_pb2.Int32Value
        min_percentage: _wrappers_pb2.Int32Value
        def __init__(self, group_by_fields: _Optional[_Iterable[_Union[_span_field_pb2.SpanField, _Mapping]]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., min_percentage: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_FIELDS_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_FIELD_NUMBER: _ClassVar[int]
    LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
    LIMITS_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.SpansFilter]
    group_by_fields: _containers.RepeatedCompositeFieldContainer[_span_field_pb2.SpanField]
    aggregation: _spans_aggregation_pb2.SpansAggregation
    lucene_query: _wrappers_pb2.StringValue
    limits: _containers.RepeatedCompositeFieldContainer[SearchGroupedSpansSeriesRequest.Limit]
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.SpansFilter, _Mapping]]] = ..., group_by_fields: _Optional[_Iterable[_Union[_span_field_pb2.SpanField, _Mapping]]] = ..., aggregation: _Optional[_Union[_spans_aggregation_pb2.SpansAggregation, _Mapping]] = ..., lucene_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., limits: _Optional[_Iterable[_Union[SearchGroupedSpansSeriesRequest.Limit, _Mapping]]] = ...) -> None: ...

class SearchGroupedSpansSeriesResponse(_message.Message):
    __slots__ = ("series",)
    SERIES_FIELD_NUMBER: _ClassVar[int]
    series: _containers.RepeatedCompositeFieldContainer[_grouped_series_pb2.GroupedSeries]
    def __init__(self, series: _Optional[_Iterable[_Union[_grouped_series_pb2.GroupedSeries, _Mapping]]] = ...) -> None: ...

class SearchSpansGroupedTimeSeriesRequest(_message.Message):
    __slots__ = ("time_frame", "interval", "filters", "aggregation", "lucene_query", "group_by")
    class GroupBy(_message.Message):
        __slots__ = ("fields", "limit", "min_percentage")
        FIELDS_FIELD_NUMBER: _ClassVar[int]
        LIMIT_FIELD_NUMBER: _ClassVar[int]
        MIN_PERCENTAGE_FIELD_NUMBER: _ClassVar[int]
        fields: _containers.RepeatedCompositeFieldContainer[_span_field_pb2.SpanField]
        limit: _wrappers_pb2.Int32Value
        min_percentage: _wrappers_pb2.Int32Value
        def __init__(self, fields: _Optional[_Iterable[_Union[_span_field_pb2.SpanField, _Mapping]]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., min_percentage: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    INTERVAL_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_FIELD_NUMBER: _ClassVar[int]
    LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    interval: _duration_pb2.Duration
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.SpansFilter]
    aggregation: _spans_aggregation_pb2.SpansAggregation
    lucene_query: _wrappers_pb2.StringValue
    group_by: _containers.RepeatedCompositeFieldContainer[SearchSpansGroupedTimeSeriesRequest.GroupBy]
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.SpansFilter, _Mapping]]] = ..., aggregation: _Optional[_Union[_spans_aggregation_pb2.SpansAggregation, _Mapping]] = ..., lucene_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., group_by: _Optional[_Iterable[_Union[SearchSpansGroupedTimeSeriesRequest.GroupBy, _Mapping]]] = ...) -> None: ...

class SearchSpansGroupedTimeSeriesResponse(_message.Message):
    __slots__ = ("series",)
    SERIES_FIELD_NUMBER: _ClassVar[int]
    series: _containers.RepeatedCompositeFieldContainer[_grouped_series_pb2.GroupedSeries]
    def __init__(self, series: _Optional[_Iterable[_Union[_grouped_series_pb2.GroupedSeries, _Mapping]]] = ...) -> None: ...

class SpansEvent(_message.Message):
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
    metadata: SpansEvent.Metadata
    start_time: _timestamp_pb2.Timestamp
    duration: _duration_pb2.Duration
    tags: _containers.RepeatedCompositeFieldContainer[SpansEvent.Tag]
    process_tags: _containers.RepeatedCompositeFieldContainer[SpansEvent.Tag]
    logs: _containers.RepeatedCompositeFieldContainer[SpansEvent.Log]
    def __init__(self, span_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., trace_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., parent_span_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., metadata: _Optional[_Union[SpansEvent.Metadata, _Mapping]] = ..., start_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., duration: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., tags: _Optional[_Iterable[_Union[SpansEvent.Tag, _Mapping]]] = ..., process_tags: _Optional[_Iterable[_Union[SpansEvent.Tag, _Mapping]]] = ..., logs: _Optional[_Iterable[_Union[SpansEvent.Log, _Mapping]]] = ...) -> None: ...

class SearchSpansTimeValueRequest(_message.Message):
    __slots__ = ("time_frame", "filters", "lucene_query", "aggregation")
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.SpansFilter]
    lucene_query: _wrappers_pb2.StringValue
    aggregation: _spans_aggregation_pb2.SpansAggregation
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.SpansFilter, _Mapping]]] = ..., lucene_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., aggregation: _Optional[_Union[_spans_aggregation_pb2.SpansAggregation, _Mapping]] = ...) -> None: ...

class SearchSpansTimeValueResponse(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: _wrappers_pb2.DoubleValue
    def __init__(self, value: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ...) -> None: ...

class SearchSpansAnnotationEventsRequest(_message.Message):
    __slots__ = ("time_frame", "filters", "lucene_query", "strategy", "limit")
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
    STRATEGY_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.SpansFilter]
    lucene_query: _wrappers_pb2.StringValue
    strategy: _annotation_pb2.Annotation.SpansSource.Strategy
    limit: _wrappers_pb2.Int32Value
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.SpansFilter, _Mapping]]] = ..., lucene_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., strategy: _Optional[_Union[_annotation_pb2.Annotation.SpansSource.Strategy, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class SearchSpansAnnotationEventsResponse(_message.Message):
    __slots__ = ("annotation_events",)
    ANNOTATION_EVENTS_FIELD_NUMBER: _ClassVar[int]
    annotation_events: _containers.RepeatedCompositeFieldContainer[_annotation_event_pb2.AnnotationEvent]
    def __init__(self, annotation_events: _Optional[_Iterable[_Union[_annotation_event_pb2.AnnotationEvent, _Mapping]]] = ...) -> None: ...
