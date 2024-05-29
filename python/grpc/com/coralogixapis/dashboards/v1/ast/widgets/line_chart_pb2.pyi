from com.coralogixapis.dashboards.v1.ast import filter_pb2 as _filter_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import data_mode_type_pb2 as _data_mode_type_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import legend_pb2 as _legend_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import queries_pb2 as _queries_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import scale_pb2 as _scale_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import units_pb2 as _units_pb2
from com.coralogixapis.dashboards.v1.common import logs_aggregation_pb2 as _logs_aggregation_pb2
from com.coralogixapis.dashboards.v1.common import observation_field_pb2 as _observation_field_pb2
from com.coralogixapis.dashboards.v1.common import query_pb2 as _query_pb2
from com.coralogixapis.dashboards.v1.common import span_field_pb2 as _span_field_pb2
from com.coralogixapis.dashboards.v1.common import spans_aggregation_pb2 as _spans_aggregation_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class LineChart(_message.Message):
    __slots__ = ("legend", "tooltip", "query_definitions", "stacked_line")
    class TooltipType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        TOOLTIP_TYPE_UNSPECIFIED: _ClassVar[LineChart.TooltipType]
        TOOLTIP_TYPE_ALL: _ClassVar[LineChart.TooltipType]
        TOOLTIP_TYPE_SINGLE: _ClassVar[LineChart.TooltipType]
    TOOLTIP_TYPE_UNSPECIFIED: LineChart.TooltipType
    TOOLTIP_TYPE_ALL: LineChart.TooltipType
    TOOLTIP_TYPE_SINGLE: LineChart.TooltipType
    class StackedLine(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        STACKED_LINE_UNSPECIFIED: _ClassVar[LineChart.StackedLine]
        STACKED_LINE_ABSOLUTE: _ClassVar[LineChart.StackedLine]
        STACKED_LINE_RELATIVE: _ClassVar[LineChart.StackedLine]
    STACKED_LINE_UNSPECIFIED: LineChart.StackedLine
    STACKED_LINE_ABSOLUTE: LineChart.StackedLine
    STACKED_LINE_RELATIVE: LineChart.StackedLine
    class QueryDefinition(_message.Message):
        __slots__ = ("id", "query", "series_name_template", "series_count_limit", "unit", "scale_type", "name", "is_visible", "color_scheme", "resolution", "data_mode_type")
        ID_FIELD_NUMBER: _ClassVar[int]
        QUERY_FIELD_NUMBER: _ClassVar[int]
        SERIES_NAME_TEMPLATE_FIELD_NUMBER: _ClassVar[int]
        SERIES_COUNT_LIMIT_FIELD_NUMBER: _ClassVar[int]
        UNIT_FIELD_NUMBER: _ClassVar[int]
        SCALE_TYPE_FIELD_NUMBER: _ClassVar[int]
        NAME_FIELD_NUMBER: _ClassVar[int]
        IS_VISIBLE_FIELD_NUMBER: _ClassVar[int]
        COLOR_SCHEME_FIELD_NUMBER: _ClassVar[int]
        RESOLUTION_FIELD_NUMBER: _ClassVar[int]
        DATA_MODE_TYPE_FIELD_NUMBER: _ClassVar[int]
        id: _wrappers_pb2.StringValue
        query: LineChart.Query
        series_name_template: _wrappers_pb2.StringValue
        series_count_limit: _wrappers_pb2.Int64Value
        unit: _units_pb2.Unit
        scale_type: _scale_pb2.ScaleType
        name: _wrappers_pb2.StringValue
        is_visible: _wrappers_pb2.BoolValue
        color_scheme: _wrappers_pb2.StringValue
        resolution: LineChart.Resolution
        data_mode_type: _data_mode_type_pb2.DataModeType
        def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., query: _Optional[_Union[LineChart.Query, _Mapping]] = ..., series_name_template: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., series_count_limit: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., unit: _Optional[_Union[_units_pb2.Unit, str]] = ..., scale_type: _Optional[_Union[_scale_pb2.ScaleType, str]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., is_visible: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., color_scheme: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., resolution: _Optional[_Union[LineChart.Resolution, _Mapping]] = ..., data_mode_type: _Optional[_Union[_data_mode_type_pb2.DataModeType, str]] = ...) -> None: ...
    class Query(_message.Message):
        __slots__ = ("logs", "metrics", "spans", "dataprime")
        LOGS_FIELD_NUMBER: _ClassVar[int]
        METRICS_FIELD_NUMBER: _ClassVar[int]
        SPANS_FIELD_NUMBER: _ClassVar[int]
        DATAPRIME_FIELD_NUMBER: _ClassVar[int]
        logs: LineChart.LogsQuery
        metrics: LineChart.MetricsQuery
        spans: LineChart.SpansQuery
        dataprime: LineChart.DataprimeQuery
        def __init__(self, logs: _Optional[_Union[LineChart.LogsQuery, _Mapping]] = ..., metrics: _Optional[_Union[LineChart.MetricsQuery, _Mapping]] = ..., spans: _Optional[_Union[LineChart.SpansQuery, _Mapping]] = ..., dataprime: _Optional[_Union[LineChart.DataprimeQuery, _Mapping]] = ...) -> None: ...
    class Tooltip(_message.Message):
        __slots__ = ("show_labels", "type")
        SHOW_LABELS_FIELD_NUMBER: _ClassVar[int]
        TYPE_FIELD_NUMBER: _ClassVar[int]
        show_labels: _wrappers_pb2.BoolValue
        type: LineChart.TooltipType
        def __init__(self, show_labels: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., type: _Optional[_Union[LineChart.TooltipType, str]] = ...) -> None: ...
    class LogsQuery(_message.Message):
        __slots__ = ("lucene_query", "group_by", "aggregations", "filters", "group_bys")
        LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
        GROUP_BY_FIELD_NUMBER: _ClassVar[int]
        AGGREGATIONS_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        GROUP_BYS_FIELD_NUMBER: _ClassVar[int]
        lucene_query: _queries_pb2.LuceneQuery
        group_by: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        aggregations: _containers.RepeatedCompositeFieldContainer[_logs_aggregation_pb2.LogsAggregation]
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.LogsFilter]
        group_bys: _containers.RepeatedCompositeFieldContainer[_observation_field_pb2.ObservationField]
        def __init__(self, lucene_query: _Optional[_Union[_queries_pb2.LuceneQuery, _Mapping]] = ..., group_by: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., aggregations: _Optional[_Iterable[_Union[_logs_aggregation_pb2.LogsAggregation, _Mapping]]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.LogsFilter, _Mapping]]] = ..., group_bys: _Optional[_Iterable[_Union[_observation_field_pb2.ObservationField, _Mapping]]] = ...) -> None: ...
    class MetricsQuery(_message.Message):
        __slots__ = ("promql_query", "filters")
        PROMQL_QUERY_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        promql_query: _queries_pb2.PromQlQuery
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.MetricsFilter]
        def __init__(self, promql_query: _Optional[_Union[_queries_pb2.PromQlQuery, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.MetricsFilter, _Mapping]]] = ...) -> None: ...
    class SpansQuery(_message.Message):
        __slots__ = ("lucene_query", "group_by", "aggregations", "filters")
        LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
        GROUP_BY_FIELD_NUMBER: _ClassVar[int]
        AGGREGATIONS_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        lucene_query: _queries_pb2.LuceneQuery
        group_by: _containers.RepeatedCompositeFieldContainer[_span_field_pb2.SpanField]
        aggregations: _containers.RepeatedCompositeFieldContainer[_spans_aggregation_pb2.SpansAggregation]
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.SpansFilter]
        def __init__(self, lucene_query: _Optional[_Union[_queries_pb2.LuceneQuery, _Mapping]] = ..., group_by: _Optional[_Iterable[_Union[_span_field_pb2.SpanField, _Mapping]]] = ..., aggregations: _Optional[_Iterable[_Union[_spans_aggregation_pb2.SpansAggregation, _Mapping]]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.SpansFilter, _Mapping]]] = ...) -> None: ...
    class DataprimeQuery(_message.Message):
        __slots__ = ("dataprime_query", "filters")
        DATAPRIME_QUERY_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        dataprime_query: _query_pb2.DataprimeQuery
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.Source]
        def __init__(self, dataprime_query: _Optional[_Union[_query_pb2.DataprimeQuery, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.Source, _Mapping]]] = ...) -> None: ...
    class Resolution(_message.Message):
        __slots__ = ("interval", "buckets_presented")
        INTERVAL_FIELD_NUMBER: _ClassVar[int]
        BUCKETS_PRESENTED_FIELD_NUMBER: _ClassVar[int]
        interval: _duration_pb2.Duration
        buckets_presented: _wrappers_pb2.Int32Value
        def __init__(self, interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., buckets_presented: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    LEGEND_FIELD_NUMBER: _ClassVar[int]
    TOOLTIP_FIELD_NUMBER: _ClassVar[int]
    QUERY_DEFINITIONS_FIELD_NUMBER: _ClassVar[int]
    STACKED_LINE_FIELD_NUMBER: _ClassVar[int]
    legend: _legend_pb2.Legend
    tooltip: LineChart.Tooltip
    query_definitions: _containers.RepeatedCompositeFieldContainer[LineChart.QueryDefinition]
    stacked_line: LineChart.StackedLine
    def __init__(self, legend: _Optional[_Union[_legend_pb2.Legend, _Mapping]] = ..., tooltip: _Optional[_Union[LineChart.Tooltip, _Mapping]] = ..., query_definitions: _Optional[_Iterable[_Union[LineChart.QueryDefinition, _Mapping]]] = ..., stacked_line: _Optional[_Union[LineChart.StackedLine, str]] = ...) -> None: ...
