from com.coralogixapis.dashboards.v1.ast import filter_pb2 as _filter_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import data_mode_type_pb2 as _data_mode_type_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import queries_pb2 as _queries_pb2
from com.coralogixapis.dashboards.v1.common import logs_aggregation_pb2 as _logs_aggregation_pb2
from com.coralogixapis.dashboards.v1.common import observation_field_pb2 as _observation_field_pb2
from com.coralogixapis.dashboards.v1.common import ordering_field_pb2 as _ordering_field_pb2
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

class RowStyle(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ROW_STYLE_UNSPECIFIED: _ClassVar[RowStyle]
    ROW_STYLE_ONE_LINE: _ClassVar[RowStyle]
    ROW_STYLE_TWO_LINE: _ClassVar[RowStyle]
    ROW_STYLE_CONDENSED: _ClassVar[RowStyle]
    ROW_STYLE_JSON: _ClassVar[RowStyle]
    ROW_STYLE_LIST: _ClassVar[RowStyle]
ROW_STYLE_UNSPECIFIED: RowStyle
ROW_STYLE_ONE_LINE: RowStyle
ROW_STYLE_TWO_LINE: RowStyle
ROW_STYLE_CONDENSED: RowStyle
ROW_STYLE_JSON: RowStyle
ROW_STYLE_LIST: RowStyle

class DataTable(_message.Message):
    __slots__ = ("query", "results_per_page", "row_style", "columns", "order_by", "data_mode_type")
    class Query(_message.Message):
        __slots__ = ("logs", "spans", "metrics", "dataprime")
        LOGS_FIELD_NUMBER: _ClassVar[int]
        SPANS_FIELD_NUMBER: _ClassVar[int]
        METRICS_FIELD_NUMBER: _ClassVar[int]
        DATAPRIME_FIELD_NUMBER: _ClassVar[int]
        logs: DataTable.LogsQuery
        spans: DataTable.SpansQuery
        metrics: DataTable.MetricsQuery
        dataprime: DataTable.DataprimeQuery
        def __init__(self, logs: _Optional[_Union[DataTable.LogsQuery, _Mapping]] = ..., spans: _Optional[_Union[DataTable.SpansQuery, _Mapping]] = ..., metrics: _Optional[_Union[DataTable.MetricsQuery, _Mapping]] = ..., dataprime: _Optional[_Union[DataTable.DataprimeQuery, _Mapping]] = ...) -> None: ...
    class LogsQuery(_message.Message):
        __slots__ = ("lucene_query", "filters", "grouping")
        class Grouping(_message.Message):
            __slots__ = ("group_by", "aggregations", "group_bys")
            GROUP_BY_FIELD_NUMBER: _ClassVar[int]
            AGGREGATIONS_FIELD_NUMBER: _ClassVar[int]
            GROUP_BYS_FIELD_NUMBER: _ClassVar[int]
            group_by: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
            aggregations: _containers.RepeatedCompositeFieldContainer[DataTable.LogsQuery.Aggregation]
            group_bys: _containers.RepeatedCompositeFieldContainer[_observation_field_pb2.ObservationField]
            def __init__(self, group_by: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., aggregations: _Optional[_Iterable[_Union[DataTable.LogsQuery.Aggregation, _Mapping]]] = ..., group_bys: _Optional[_Iterable[_Union[_observation_field_pb2.ObservationField, _Mapping]]] = ...) -> None: ...
        class Aggregation(_message.Message):
            __slots__ = ("id", "name", "is_visible", "aggregation")
            ID_FIELD_NUMBER: _ClassVar[int]
            NAME_FIELD_NUMBER: _ClassVar[int]
            IS_VISIBLE_FIELD_NUMBER: _ClassVar[int]
            AGGREGATION_FIELD_NUMBER: _ClassVar[int]
            id: _wrappers_pb2.StringValue
            name: _wrappers_pb2.StringValue
            is_visible: _wrappers_pb2.BoolValue
            aggregation: _logs_aggregation_pb2.LogsAggregation
            def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., is_visible: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., aggregation: _Optional[_Union[_logs_aggregation_pb2.LogsAggregation, _Mapping]] = ...) -> None: ...
        LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        GROUPING_FIELD_NUMBER: _ClassVar[int]
        lucene_query: _queries_pb2.LuceneQuery
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.LogsFilter]
        grouping: DataTable.LogsQuery.Grouping
        def __init__(self, lucene_query: _Optional[_Union[_queries_pb2.LuceneQuery, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.LogsFilter, _Mapping]]] = ..., grouping: _Optional[_Union[DataTable.LogsQuery.Grouping, _Mapping]] = ...) -> None: ...
    class SpansQuery(_message.Message):
        __slots__ = ("lucene_query", "filters", "grouping")
        class Grouping(_message.Message):
            __slots__ = ("group_by", "aggregations")
            GROUP_BY_FIELD_NUMBER: _ClassVar[int]
            AGGREGATIONS_FIELD_NUMBER: _ClassVar[int]
            group_by: _containers.RepeatedCompositeFieldContainer[_span_field_pb2.SpanField]
            aggregations: _containers.RepeatedCompositeFieldContainer[DataTable.SpansQuery.Aggregation]
            def __init__(self, group_by: _Optional[_Iterable[_Union[_span_field_pb2.SpanField, _Mapping]]] = ..., aggregations: _Optional[_Iterable[_Union[DataTable.SpansQuery.Aggregation, _Mapping]]] = ...) -> None: ...
        class Aggregation(_message.Message):
            __slots__ = ("id", "name", "is_visible", "aggregation")
            ID_FIELD_NUMBER: _ClassVar[int]
            NAME_FIELD_NUMBER: _ClassVar[int]
            IS_VISIBLE_FIELD_NUMBER: _ClassVar[int]
            AGGREGATION_FIELD_NUMBER: _ClassVar[int]
            id: _wrappers_pb2.StringValue
            name: _wrappers_pb2.StringValue
            is_visible: _wrappers_pb2.BoolValue
            aggregation: _spans_aggregation_pb2.SpansAggregation
            def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., is_visible: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., aggregation: _Optional[_Union[_spans_aggregation_pb2.SpansAggregation, _Mapping]] = ...) -> None: ...
        LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        GROUPING_FIELD_NUMBER: _ClassVar[int]
        lucene_query: _queries_pb2.LuceneQuery
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.SpansFilter]
        grouping: DataTable.SpansQuery.Grouping
        def __init__(self, lucene_query: _Optional[_Union[_queries_pb2.LuceneQuery, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.SpansFilter, _Mapping]]] = ..., grouping: _Optional[_Union[DataTable.SpansQuery.Grouping, _Mapping]] = ...) -> None: ...
    class MetricsQuery(_message.Message):
        __slots__ = ("promql_query", "filters")
        PROMQL_QUERY_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        promql_query: _queries_pb2.PromQlQuery
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.MetricsFilter]
        def __init__(self, promql_query: _Optional[_Union[_queries_pb2.PromQlQuery, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.MetricsFilter, _Mapping]]] = ...) -> None: ...
    class DataprimeQuery(_message.Message):
        __slots__ = ("dataprime_query", "filters")
        DATAPRIME_QUERY_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        dataprime_query: _query_pb2.DataprimeQuery
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.Source]
        def __init__(self, dataprime_query: _Optional[_Union[_query_pb2.DataprimeQuery, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.Source, _Mapping]]] = ...) -> None: ...
    class Column(_message.Message):
        __slots__ = ("field", "width")
        FIELD_FIELD_NUMBER: _ClassVar[int]
        WIDTH_FIELD_NUMBER: _ClassVar[int]
        field: _wrappers_pb2.StringValue
        width: _wrappers_pb2.Int32Value
        def __init__(self, field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., width: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    QUERY_FIELD_NUMBER: _ClassVar[int]
    RESULTS_PER_PAGE_FIELD_NUMBER: _ClassVar[int]
    ROW_STYLE_FIELD_NUMBER: _ClassVar[int]
    COLUMNS_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_NUMBER: _ClassVar[int]
    DATA_MODE_TYPE_FIELD_NUMBER: _ClassVar[int]
    query: DataTable.Query
    results_per_page: _wrappers_pb2.Int32Value
    row_style: RowStyle
    columns: _containers.RepeatedCompositeFieldContainer[DataTable.Column]
    order_by: _ordering_field_pb2.OrderingField
    data_mode_type: _data_mode_type_pb2.DataModeType
    def __init__(self, query: _Optional[_Union[DataTable.Query, _Mapping]] = ..., results_per_page: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., row_style: _Optional[_Union[RowStyle, str]] = ..., columns: _Optional[_Iterable[_Union[DataTable.Column, _Mapping]]] = ..., order_by: _Optional[_Union[_ordering_field_pb2.OrderingField, _Mapping]] = ..., data_mode_type: _Optional[_Union[_data_mode_type_pb2.DataModeType, str]] = ...) -> None: ...
