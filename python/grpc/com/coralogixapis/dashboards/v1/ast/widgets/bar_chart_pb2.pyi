from com.coralogixapis.dashboards.v1.ast import filter_pb2 as _filter_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import colors_by_pb2 as _colors_by_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import data_mode_type_pb2 as _data_mode_type_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import queries_pb2 as _queries_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import scale_pb2 as _scale_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import sort_by_pb2 as _sort_by_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import units_pb2 as _units_pb2
from com.coralogixapis.dashboards.v1.common import logs_aggregation_pb2 as _logs_aggregation_pb2
from com.coralogixapis.dashboards.v1.common import observation_field_pb2 as _observation_field_pb2
from com.coralogixapis.dashboards.v1.common import query_pb2 as _query_pb2
from com.coralogixapis.dashboards.v1.common import span_field_pb2 as _span_field_pb2
from com.coralogixapis.dashboards.v1.common import spans_aggregation_pb2 as _spans_aggregation_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class BarChart(_message.Message):
    __slots__ = ("query", "max_bars_per_chart", "group_name_template", "stack_definition", "scale_type", "colors_by", "x_axis", "unit", "sort_by", "color_scheme", "data_mode_type")
    class Query(_message.Message):
        __slots__ = ("logs", "spans", "metrics", "dataprime")
        LOGS_FIELD_NUMBER: _ClassVar[int]
        SPANS_FIELD_NUMBER: _ClassVar[int]
        METRICS_FIELD_NUMBER: _ClassVar[int]
        DATAPRIME_FIELD_NUMBER: _ClassVar[int]
        logs: BarChart.LogsQuery
        spans: BarChart.SpansQuery
        metrics: BarChart.MetricsQuery
        dataprime: BarChart.DataprimeQuery
        def __init__(self, logs: _Optional[_Union[BarChart.LogsQuery, _Mapping]] = ..., spans: _Optional[_Union[BarChart.SpansQuery, _Mapping]] = ..., metrics: _Optional[_Union[BarChart.MetricsQuery, _Mapping]] = ..., dataprime: _Optional[_Union[BarChart.DataprimeQuery, _Mapping]] = ...) -> None: ...
    class StackDefinition(_message.Message):
        __slots__ = ("max_slices_per_bar", "stack_name_template")
        MAX_SLICES_PER_BAR_FIELD_NUMBER: _ClassVar[int]
        STACK_NAME_TEMPLATE_FIELD_NUMBER: _ClassVar[int]
        max_slices_per_bar: _wrappers_pb2.Int32Value
        stack_name_template: _wrappers_pb2.StringValue
        def __init__(self, max_slices_per_bar: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., stack_name_template: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class XAxis(_message.Message):
        __slots__ = ("value", "time")
        class XAxisByValue(_message.Message):
            __slots__ = ()
            def __init__(self) -> None: ...
        class XAxisByTime(_message.Message):
            __slots__ = ("interval", "buckets_presented")
            INTERVAL_FIELD_NUMBER: _ClassVar[int]
            BUCKETS_PRESENTED_FIELD_NUMBER: _ClassVar[int]
            interval: _duration_pb2.Duration
            buckets_presented: _wrappers_pb2.Int32Value
            def __init__(self, interval: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., buckets_presented: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
        VALUE_FIELD_NUMBER: _ClassVar[int]
        TIME_FIELD_NUMBER: _ClassVar[int]
        value: BarChart.XAxis.XAxisByValue
        time: BarChart.XAxis.XAxisByTime
        def __init__(self, value: _Optional[_Union[BarChart.XAxis.XAxisByValue, _Mapping]] = ..., time: _Optional[_Union[BarChart.XAxis.XAxisByTime, _Mapping]] = ...) -> None: ...
    class LogsQuery(_message.Message):
        __slots__ = ("lucene_query", "aggregation", "filters", "group_names", "stacked_group_name", "group_names_fields", "stacked_group_name_field")
        LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
        AGGREGATION_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        GROUP_NAMES_FIELD_NUMBER: _ClassVar[int]
        STACKED_GROUP_NAME_FIELD_NUMBER: _ClassVar[int]
        GROUP_NAMES_FIELDS_FIELD_NUMBER: _ClassVar[int]
        STACKED_GROUP_NAME_FIELD_FIELD_NUMBER: _ClassVar[int]
        lucene_query: _queries_pb2.LuceneQuery
        aggregation: _logs_aggregation_pb2.LogsAggregation
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.LogsFilter]
        group_names: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        stacked_group_name: _wrappers_pb2.StringValue
        group_names_fields: _containers.RepeatedCompositeFieldContainer[_observation_field_pb2.ObservationField]
        stacked_group_name_field: _observation_field_pb2.ObservationField
        def __init__(self, lucene_query: _Optional[_Union[_queries_pb2.LuceneQuery, _Mapping]] = ..., aggregation: _Optional[_Union[_logs_aggregation_pb2.LogsAggregation, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.LogsFilter, _Mapping]]] = ..., group_names: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., stacked_group_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., group_names_fields: _Optional[_Iterable[_Union[_observation_field_pb2.ObservationField, _Mapping]]] = ..., stacked_group_name_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
    class SpansQuery(_message.Message):
        __slots__ = ("lucene_query", "aggregation", "filters", "group_names", "stacked_group_name")
        LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
        AGGREGATION_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        GROUP_NAMES_FIELD_NUMBER: _ClassVar[int]
        STACKED_GROUP_NAME_FIELD_NUMBER: _ClassVar[int]
        lucene_query: _queries_pb2.LuceneQuery
        aggregation: _spans_aggregation_pb2.SpansAggregation
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.SpansFilter]
        group_names: _containers.RepeatedCompositeFieldContainer[_span_field_pb2.SpanField]
        stacked_group_name: _span_field_pb2.SpanField
        def __init__(self, lucene_query: _Optional[_Union[_queries_pb2.LuceneQuery, _Mapping]] = ..., aggregation: _Optional[_Union[_spans_aggregation_pb2.SpansAggregation, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.SpansFilter, _Mapping]]] = ..., group_names: _Optional[_Iterable[_Union[_span_field_pb2.SpanField, _Mapping]]] = ..., stacked_group_name: _Optional[_Union[_span_field_pb2.SpanField, _Mapping]] = ...) -> None: ...
    class MetricsQuery(_message.Message):
        __slots__ = ("promql_query", "filters", "group_names", "stacked_group_name")
        PROMQL_QUERY_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        GROUP_NAMES_FIELD_NUMBER: _ClassVar[int]
        STACKED_GROUP_NAME_FIELD_NUMBER: _ClassVar[int]
        promql_query: _queries_pb2.PromQlQuery
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.MetricsFilter]
        group_names: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        stacked_group_name: _wrappers_pb2.StringValue
        def __init__(self, promql_query: _Optional[_Union[_queries_pb2.PromQlQuery, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.MetricsFilter, _Mapping]]] = ..., group_names: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., stacked_group_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class DataprimeQuery(_message.Message):
        __slots__ = ("dataprime_query", "filters", "group_names", "stacked_group_name")
        DATAPRIME_QUERY_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        GROUP_NAMES_FIELD_NUMBER: _ClassVar[int]
        STACKED_GROUP_NAME_FIELD_NUMBER: _ClassVar[int]
        dataprime_query: _query_pb2.DataprimeQuery
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.Source]
        group_names: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        stacked_group_name: _wrappers_pb2.StringValue
        def __init__(self, dataprime_query: _Optional[_Union[_query_pb2.DataprimeQuery, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.Source, _Mapping]]] = ..., group_names: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., stacked_group_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    QUERY_FIELD_NUMBER: _ClassVar[int]
    MAX_BARS_PER_CHART_FIELD_NUMBER: _ClassVar[int]
    GROUP_NAME_TEMPLATE_FIELD_NUMBER: _ClassVar[int]
    STACK_DEFINITION_FIELD_NUMBER: _ClassVar[int]
    SCALE_TYPE_FIELD_NUMBER: _ClassVar[int]
    COLORS_BY_FIELD_NUMBER: _ClassVar[int]
    X_AXIS_FIELD_NUMBER: _ClassVar[int]
    UNIT_FIELD_NUMBER: _ClassVar[int]
    SORT_BY_FIELD_NUMBER: _ClassVar[int]
    COLOR_SCHEME_FIELD_NUMBER: _ClassVar[int]
    DATA_MODE_TYPE_FIELD_NUMBER: _ClassVar[int]
    query: BarChart.Query
    max_bars_per_chart: _wrappers_pb2.Int32Value
    group_name_template: _wrappers_pb2.StringValue
    stack_definition: BarChart.StackDefinition
    scale_type: _scale_pb2.ScaleType
    colors_by: _colors_by_pb2.ColorsBy
    x_axis: BarChart.XAxis
    unit: _units_pb2.Unit
    sort_by: _sort_by_pb2.SortByType
    color_scheme: _wrappers_pb2.StringValue
    data_mode_type: _data_mode_type_pb2.DataModeType
    def __init__(self, query: _Optional[_Union[BarChart.Query, _Mapping]] = ..., max_bars_per_chart: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., group_name_template: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., stack_definition: _Optional[_Union[BarChart.StackDefinition, _Mapping]] = ..., scale_type: _Optional[_Union[_scale_pb2.ScaleType, str]] = ..., colors_by: _Optional[_Union[_colors_by_pb2.ColorsBy, _Mapping]] = ..., x_axis: _Optional[_Union[BarChart.XAxis, _Mapping]] = ..., unit: _Optional[_Union[_units_pb2.Unit, str]] = ..., sort_by: _Optional[_Union[_sort_by_pb2.SortByType, str]] = ..., color_scheme: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., data_mode_type: _Optional[_Union[_data_mode_type_pb2.DataModeType, str]] = ...) -> None: ...
