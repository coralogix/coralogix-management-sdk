from com.coralogixapis.dashboards.v1.ast import annotation_pb2 as _annotation_pb2
from com.coralogixapis.dashboards.v1.ast import filter_pb2 as _filter_pb2
from com.coralogixapis.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogixapis.dashboards.v1.common import annotation_event_pb2 as _annotation_event_pb2
from com.coralogixapis.dashboards.v1.common import grouped_series_pb2 as _grouped_series_pb2
from com.coralogixapis.dashboards.v1.common import labelled_value_pb2 as _labelled_value_pb2
from com.coralogixapis.dashboards.v1.common import order_direction_pb2 as _order_direction_pb2
from com.coralogixapis.dashboards.v1.common import time_frame_pb2 as _time_frame_pb2
from com.coralogixapis.dashboards.v1.common import time_series_pb2 as _time_series_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SearchMetricsTimeSeriesRequest(_message.Message):
    __slots__ = ("time_frame", "interval", "promql_query", "limit", "filters")
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    INTERVAL_FIELD_NUMBER: _ClassVar[int]
    PROMQL_QUERY_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    interval: _duration_pb2.Duration
    promql_query: _wrappers_pb2.StringValue
    limit: _wrappers_pb2.Int32Value
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.MetricsFilter]
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., promql_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.MetricsFilter, _Mapping]]] = ...) -> None: ...

class SearchMetricsTimeSeriesResponse(_message.Message):
    __slots__ = ("time_series", "is_limit_exceeded", "total")
    TIME_SERIES_FIELD_NUMBER: _ClassVar[int]
    IS_LIMIT_EXCEEDED_FIELD_NUMBER: _ClassVar[int]
    TOTAL_FIELD_NUMBER: _ClassVar[int]
    time_series: _containers.RepeatedCompositeFieldContainer[_time_series_pb2.TimeSeries]
    is_limit_exceeded: bool
    total: _wrappers_pb2.Int64Value
    def __init__(self, time_series: _Optional[_Iterable[_Union[_time_series_pb2.TimeSeries, _Mapping]]] = ..., is_limit_exceeded: bool = ..., total: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class SearchMetricsTimeValuesRequest(_message.Message):
    __slots__ = ("time_frame", "promql_query", "limit", "filters")
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    PROMQL_QUERY_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    promql_query: _wrappers_pb2.StringValue
    limit: _wrappers_pb2.Int32Value
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.MetricsFilter]
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., promql_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.MetricsFilter, _Mapping]]] = ...) -> None: ...

class SearchMetricsTimeValuesResponse(_message.Message):
    __slots__ = ("values", "is_limit_exceeded", "total")
    VALUES_FIELD_NUMBER: _ClassVar[int]
    IS_LIMIT_EXCEEDED_FIELD_NUMBER: _ClassVar[int]
    TOTAL_FIELD_NUMBER: _ClassVar[int]
    values: _containers.RepeatedCompositeFieldContainer[_labelled_value_pb2.LabelledValue]
    is_limit_exceeded: bool
    total: _wrappers_pb2.Int64Value
    def __init__(self, values: _Optional[_Iterable[_Union[_labelled_value_pb2.LabelledValue, _Mapping]]] = ..., is_limit_exceeded: bool = ..., total: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class SearchMetricsGroupedSeriesRequest(_message.Message):
    __slots__ = ("time_frame", "promql_query", "group_by", "limits", "filters")
    class Limit(_message.Message):
        __slots__ = ("group_by", "limit", "min_percentage")
        GROUP_BY_FIELD_NUMBER: _ClassVar[int]
        LIMIT_FIELD_NUMBER: _ClassVar[int]
        MIN_PERCENTAGE_FIELD_NUMBER: _ClassVar[int]
        group_by: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        limit: _wrappers_pb2.Int32Value
        min_percentage: _wrappers_pb2.Int32Value
        def __init__(self, group_by: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., min_percentage: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    PROMQL_QUERY_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_FIELD_NUMBER: _ClassVar[int]
    LIMITS_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    promql_query: _wrappers_pb2.StringValue
    group_by: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    limits: _containers.RepeatedCompositeFieldContainer[SearchMetricsGroupedSeriesRequest.Limit]
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.MetricsFilter]
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., promql_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., group_by: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., limits: _Optional[_Iterable[_Union[SearchMetricsGroupedSeriesRequest.Limit, _Mapping]]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.MetricsFilter, _Mapping]]] = ...) -> None: ...

class SearchMetricsGroupedSeriesResponse(_message.Message):
    __slots__ = ("series",)
    SERIES_FIELD_NUMBER: _ClassVar[int]
    series: _containers.RepeatedCompositeFieldContainer[_grouped_series_pb2.GroupedSeries]
    def __init__(self, series: _Optional[_Iterable[_Union[_grouped_series_pb2.GroupedSeries, _Mapping]]] = ...) -> None: ...

class SearchMetricsGroupedTimeSeriesRequest(_message.Message):
    __slots__ = ("time_frame", "interval", "promql_query", "group_by", "filters")
    class GroupBy(_message.Message):
        __slots__ = ("labels", "limit", "min_percentage")
        LABELS_FIELD_NUMBER: _ClassVar[int]
        LIMIT_FIELD_NUMBER: _ClassVar[int]
        MIN_PERCENTAGE_FIELD_NUMBER: _ClassVar[int]
        labels: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        limit: _wrappers_pb2.Int32Value
        min_percentage: _wrappers_pb2.Int32Value
        def __init__(self, labels: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., min_percentage: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    INTERVAL_FIELD_NUMBER: _ClassVar[int]
    PROMQL_QUERY_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    interval: _duration_pb2.Duration
    promql_query: _wrappers_pb2.StringValue
    group_by: _containers.RepeatedCompositeFieldContainer[SearchMetricsGroupedTimeSeriesRequest.GroupBy]
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.MetricsFilter]
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., promql_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., group_by: _Optional[_Iterable[_Union[SearchMetricsGroupedTimeSeriesRequest.GroupBy, _Mapping]]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.MetricsFilter, _Mapping]]] = ...) -> None: ...

class SearchMetricsGroupedTimeSeriesResponse(_message.Message):
    __slots__ = ("series",)
    SERIES_FIELD_NUMBER: _ClassVar[int]
    series: _containers.RepeatedCompositeFieldContainer[_grouped_series_pb2.GroupedSeries]
    def __init__(self, series: _Optional[_Iterable[_Union[_grouped_series_pb2.GroupedSeries, _Mapping]]] = ...) -> None: ...

class SearchMetricsEventsRequest(_message.Message):
    __slots__ = ("time_frame", "time_series_interval", "promql_query", "pagination", "filters", "order_by")
    class OrderBy(_message.Message):
        __slots__ = ("direction", "label", "agg_type")
        DIRECTION_FIELD_NUMBER: _ClassVar[int]
        LABEL_FIELD_NUMBER: _ClassVar[int]
        AGG_TYPE_FIELD_NUMBER: _ClassVar[int]
        direction: _order_direction_pb2.OrderDirection
        label: _wrappers_pb2.StringValue
        agg_type: MetricsEvent.AggregationType
        def __init__(self, direction: _Optional[_Union[_order_direction_pb2.OrderDirection, str]] = ..., label: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., agg_type: _Optional[_Union[MetricsEvent.AggregationType, str]] = ...) -> None: ...
    class Pagination(_message.Message):
        __slots__ = ("offset", "limit")
        OFFSET_FIELD_NUMBER: _ClassVar[int]
        LIMIT_FIELD_NUMBER: _ClassVar[int]
        offset: _wrappers_pb2.Int32Value
        limit: _wrappers_pb2.Int32Value
        def __init__(self, offset: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    TIME_SERIES_INTERVAL_FIELD_NUMBER: _ClassVar[int]
    PROMQL_QUERY_FIELD_NUMBER: _ClassVar[int]
    PAGINATION_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    time_series_interval: _duration_pb2.Duration
    promql_query: _wrappers_pb2.StringValue
    pagination: SearchMetricsEventsRequest.Pagination
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.MetricsFilter]
    order_by: SearchMetricsEventsRequest.OrderBy
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., time_series_interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., promql_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., pagination: _Optional[_Union[SearchMetricsEventsRequest.Pagination, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.MetricsFilter, _Mapping]]] = ..., order_by: _Optional[_Union[SearchMetricsEventsRequest.OrderBy, _Mapping]] = ...) -> None: ...

class SearchMetricsEventsResponse(_message.Message):
    __slots__ = ("labels", "metrics_events", "total")
    LABELS_FIELD_NUMBER: _ClassVar[int]
    METRICS_EVENTS_FIELD_NUMBER: _ClassVar[int]
    TOTAL_FIELD_NUMBER: _ClassVar[int]
    labels: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    metrics_events: _containers.RepeatedCompositeFieldContainer[MetricsEvent]
    total: _wrappers_pb2.Int64Value
    def __init__(self, labels: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., metrics_events: _Optional[_Iterable[_Union[MetricsEvent, _Mapping]]] = ..., total: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class MetricsEvent(_message.Message):
    __slots__ = ("aggregations", "time_series")
    class AggregationType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        AGGREGATION_TYPE_UNSPECIFIED: _ClassVar[MetricsEvent.AggregationType]
        AGGREGATION_TYPE_LAST: _ClassVar[MetricsEvent.AggregationType]
        AGGREGATION_TYPE_MIN: _ClassVar[MetricsEvent.AggregationType]
        AGGREGATION_TYPE_MAX: _ClassVar[MetricsEvent.AggregationType]
        AGGREGATION_TYPE_AVG: _ClassVar[MetricsEvent.AggregationType]
        AGGREGATION_TYPE_SUM: _ClassVar[MetricsEvent.AggregationType]
    AGGREGATION_TYPE_UNSPECIFIED: MetricsEvent.AggregationType
    AGGREGATION_TYPE_LAST: MetricsEvent.AggregationType
    AGGREGATION_TYPE_MIN: MetricsEvent.AggregationType
    AGGREGATION_TYPE_MAX: MetricsEvent.AggregationType
    AGGREGATION_TYPE_AVG: MetricsEvent.AggregationType
    AGGREGATION_TYPE_SUM: MetricsEvent.AggregationType
    class Aggregation(_message.Message):
        __slots__ = ("aggregation_type", "value")
        AGGREGATION_TYPE_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        aggregation_type: MetricsEvent.AggregationType
        value: _wrappers_pb2.DoubleValue
        def __init__(self, aggregation_type: _Optional[_Union[MetricsEvent.AggregationType, str]] = ..., value: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ...) -> None: ...
    AGGREGATIONS_FIELD_NUMBER: _ClassVar[int]
    TIME_SERIES_FIELD_NUMBER: _ClassVar[int]
    aggregations: _containers.RepeatedCompositeFieldContainer[MetricsEvent.Aggregation]
    time_series: _time_series_pb2.TimeSeries
    def __init__(self, aggregations: _Optional[_Iterable[_Union[MetricsEvent.Aggregation, _Mapping]]] = ..., time_series: _Optional[_Union[_time_series_pb2.TimeSeries, _Mapping]] = ...) -> None: ...

class SearchMetricsAnnotationEventsRequest(_message.Message):
    __slots__ = ("time_frame", "interval", "filters", "promql_query", "strategy", "limit")
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    INTERVAL_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    PROMQL_QUERY_FIELD_NUMBER: _ClassVar[int]
    STRATEGY_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    time_frame: _time_frame_pb2.TimeFrame
    interval: _duration_pb2.Duration
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.MetricsFilter]
    promql_query: _wrappers_pb2.StringValue
    strategy: _annotation_pb2.Annotation.MetricsSource.Strategy
    limit: _wrappers_pb2.Int32Value
    def __init__(self, time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.MetricsFilter, _Mapping]]] = ..., promql_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., strategy: _Optional[_Union[_annotation_pb2.Annotation.MetricsSource.Strategy, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class SearchMetricsAnnotationEventsResponse(_message.Message):
    __slots__ = ("annotation_events",)
    ANNOTATION_EVENTS_FIELD_NUMBER: _ClassVar[int]
    annotation_events: _containers.RepeatedCompositeFieldContainer[_annotation_event_pb2.AnnotationEvent]
    def __init__(self, annotation_events: _Optional[_Iterable[_Union[_annotation_event_pb2.AnnotationEvent, _Mapping]]] = ...) -> None: ...
