from com.coralogixapis.dashboards.v1.ast import filter_pb2 as _filter_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import data_mode_type_pb2 as _data_mode_type_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import queries_pb2 as _queries_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import units_pb2 as _units_pb2
from com.coralogixapis.dashboards.v1.common import logs_aggregation_pb2 as _logs_aggregation_pb2
from com.coralogixapis.dashboards.v1.common import observation_field_pb2 as _observation_field_pb2
from com.coralogixapis.dashboards.v1.common import query_pb2 as _query_pb2
from com.coralogixapis.dashboards.v1.common import span_field_pb2 as _span_field_pb2
from com.coralogixapis.dashboards.v1.common import spans_aggregation_pb2 as _spans_aggregation_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class PieChart(_message.Message):
    __slots__ = ("query", "max_slices_per_chart", "min_slice_percentage", "stack_definition", "label_definition", "show_legend", "group_name_template", "unit", "color_scheme", "data_mode_type")
    class LabelSource(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        LABEL_SOURCE_UNSPECIFIED: _ClassVar[PieChart.LabelSource]
        LABEL_SOURCE_INNER: _ClassVar[PieChart.LabelSource]
        LABEL_SOURCE_STACK: _ClassVar[PieChart.LabelSource]
    LABEL_SOURCE_UNSPECIFIED: PieChart.LabelSource
    LABEL_SOURCE_INNER: PieChart.LabelSource
    LABEL_SOURCE_STACK: PieChart.LabelSource
    class Query(_message.Message):
        __slots__ = ("logs", "spans", "metrics", "dataprime")
        LOGS_FIELD_NUMBER: _ClassVar[int]
        SPANS_FIELD_NUMBER: _ClassVar[int]
        METRICS_FIELD_NUMBER: _ClassVar[int]
        DATAPRIME_FIELD_NUMBER: _ClassVar[int]
        logs: PieChart.LogsQuery
        spans: PieChart.SpansQuery
        metrics: PieChart.MetricsQuery
        dataprime: PieChart.DataprimeQuery
        def __init__(self, logs: _Optional[_Union[PieChart.LogsQuery, _Mapping]] = ..., spans: _Optional[_Union[PieChart.SpansQuery, _Mapping]] = ..., metrics: _Optional[_Union[PieChart.MetricsQuery, _Mapping]] = ..., dataprime: _Optional[_Union[PieChart.DataprimeQuery, _Mapping]] = ...) -> None: ...
    class StackDefinition(_message.Message):
        __slots__ = ("max_slices_per_stack", "stack_name_template")
        MAX_SLICES_PER_STACK_FIELD_NUMBER: _ClassVar[int]
        STACK_NAME_TEMPLATE_FIELD_NUMBER: _ClassVar[int]
        max_slices_per_stack: _wrappers_pb2.Int32Value
        stack_name_template: _wrappers_pb2.StringValue
        def __init__(self, max_slices_per_stack: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., stack_name_template: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class LabelDefinition(_message.Message):
        __slots__ = ("label_source", "is_visible", "show_name", "show_value", "show_percentage")
        LABEL_SOURCE_FIELD_NUMBER: _ClassVar[int]
        IS_VISIBLE_FIELD_NUMBER: _ClassVar[int]
        SHOW_NAME_FIELD_NUMBER: _ClassVar[int]
        SHOW_VALUE_FIELD_NUMBER: _ClassVar[int]
        SHOW_PERCENTAGE_FIELD_NUMBER: _ClassVar[int]
        label_source: PieChart.LabelSource
        is_visible: _wrappers_pb2.BoolValue
        show_name: _wrappers_pb2.BoolValue
        show_value: _wrappers_pb2.BoolValue
        show_percentage: _wrappers_pb2.BoolValue
        def __init__(self, label_source: _Optional[_Union[PieChart.LabelSource, str]] = ..., is_visible: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., show_name: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., show_value: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., show_percentage: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
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
    MAX_SLICES_PER_CHART_FIELD_NUMBER: _ClassVar[int]
    MIN_SLICE_PERCENTAGE_FIELD_NUMBER: _ClassVar[int]
    STACK_DEFINITION_FIELD_NUMBER: _ClassVar[int]
    LABEL_DEFINITION_FIELD_NUMBER: _ClassVar[int]
    SHOW_LEGEND_FIELD_NUMBER: _ClassVar[int]
    GROUP_NAME_TEMPLATE_FIELD_NUMBER: _ClassVar[int]
    UNIT_FIELD_NUMBER: _ClassVar[int]
    COLOR_SCHEME_FIELD_NUMBER: _ClassVar[int]
    DATA_MODE_TYPE_FIELD_NUMBER: _ClassVar[int]
    query: PieChart.Query
    max_slices_per_chart: _wrappers_pb2.Int32Value
    min_slice_percentage: _wrappers_pb2.Int32Value
    stack_definition: PieChart.StackDefinition
    label_definition: PieChart.LabelDefinition
    show_legend: _wrappers_pb2.BoolValue
    group_name_template: _wrappers_pb2.StringValue
    unit: _units_pb2.Unit
    color_scheme: _wrappers_pb2.StringValue
    data_mode_type: _data_mode_type_pb2.DataModeType
    def __init__(self, query: _Optional[_Union[PieChart.Query, _Mapping]] = ..., max_slices_per_chart: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., min_slice_percentage: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., stack_definition: _Optional[_Union[PieChart.StackDefinition, _Mapping]] = ..., label_definition: _Optional[_Union[PieChart.LabelDefinition, _Mapping]] = ..., show_legend: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., group_name_template: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., unit: _Optional[_Union[_units_pb2.Unit, str]] = ..., color_scheme: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., data_mode_type: _Optional[_Union[_data_mode_type_pb2.DataModeType, str]] = ...) -> None: ...
