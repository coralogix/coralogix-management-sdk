from com.coralogixapis.dashboards.v1.ast import filter_pb2 as _filter_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import data_mode_type_pb2 as _data_mode_type_pb2
from com.coralogixapis.dashboards.v1.ast.widgets.common import queries_pb2 as _queries_pb2
from com.coralogixapis.dashboards.v1.common import logs_aggregation_pb2 as _logs_aggregation_pb2
from com.coralogixapis.dashboards.v1.common import query_pb2 as _query_pb2
from com.coralogixapis.dashboards.v1.common import spans_aggregation_pb2 as _spans_aggregation_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Gauge(_message.Message):
    __slots__ = ("query", "min", "max", "show_inner_arc", "show_outer_arc", "unit", "thresholds", "data_mode_type", "threshold_by")
    class Aggregation(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        AGGREGATION_UNSPECIFIED: _ClassVar[Gauge.Aggregation]
        AGGREGATION_LAST: _ClassVar[Gauge.Aggregation]
        AGGREGATION_MIN: _ClassVar[Gauge.Aggregation]
        AGGREGATION_MAX: _ClassVar[Gauge.Aggregation]
        AGGREGATION_AVG: _ClassVar[Gauge.Aggregation]
        AGGREGATION_SUM: _ClassVar[Gauge.Aggregation]
    AGGREGATION_UNSPECIFIED: Gauge.Aggregation
    AGGREGATION_LAST: Gauge.Aggregation
    AGGREGATION_MIN: Gauge.Aggregation
    AGGREGATION_MAX: Gauge.Aggregation
    AGGREGATION_AVG: Gauge.Aggregation
    AGGREGATION_SUM: Gauge.Aggregation
    class Unit(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        UNIT_UNSPECIFIED: _ClassVar[Gauge.Unit]
        UNIT_NUMBER: _ClassVar[Gauge.Unit]
        UNIT_PERCENT: _ClassVar[Gauge.Unit]
        UNIT_MICROSECONDS: _ClassVar[Gauge.Unit]
        UNIT_MILLISECONDS: _ClassVar[Gauge.Unit]
        UNIT_SECONDS: _ClassVar[Gauge.Unit]
        UNIT_BYTES: _ClassVar[Gauge.Unit]
        UNIT_KBYTES: _ClassVar[Gauge.Unit]
        UNIT_MBYTES: _ClassVar[Gauge.Unit]
        UNIT_GBYTES: _ClassVar[Gauge.Unit]
        UNIT_BYTES_IEC: _ClassVar[Gauge.Unit]
        UNIT_KIBYTES: _ClassVar[Gauge.Unit]
        UNIT_MIBYTES: _ClassVar[Gauge.Unit]
        UNIT_GIBYTES: _ClassVar[Gauge.Unit]
        UNIT_EUR_CENTS: _ClassVar[Gauge.Unit]
        UNIT_EUR: _ClassVar[Gauge.Unit]
        UNIT_USD_CENTS: _ClassVar[Gauge.Unit]
        UNIT_USD: _ClassVar[Gauge.Unit]
    UNIT_UNSPECIFIED: Gauge.Unit
    UNIT_NUMBER: Gauge.Unit
    UNIT_PERCENT: Gauge.Unit
    UNIT_MICROSECONDS: Gauge.Unit
    UNIT_MILLISECONDS: Gauge.Unit
    UNIT_SECONDS: Gauge.Unit
    UNIT_BYTES: Gauge.Unit
    UNIT_KBYTES: Gauge.Unit
    UNIT_MBYTES: Gauge.Unit
    UNIT_GBYTES: Gauge.Unit
    UNIT_BYTES_IEC: Gauge.Unit
    UNIT_KIBYTES: Gauge.Unit
    UNIT_MIBYTES: Gauge.Unit
    UNIT_GIBYTES: Gauge.Unit
    UNIT_EUR_CENTS: Gauge.Unit
    UNIT_EUR: Gauge.Unit
    UNIT_USD_CENTS: Gauge.Unit
    UNIT_USD: Gauge.Unit
    class ThresholdBy(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        THRESHOLD_BY_UNSPECIFIED: _ClassVar[Gauge.ThresholdBy]
        THRESHOLD_BY_VALUE: _ClassVar[Gauge.ThresholdBy]
        THRESHOLD_BY_BACKGROUND: _ClassVar[Gauge.ThresholdBy]
    THRESHOLD_BY_UNSPECIFIED: Gauge.ThresholdBy
    THRESHOLD_BY_VALUE: Gauge.ThresholdBy
    THRESHOLD_BY_BACKGROUND: Gauge.ThresholdBy
    class Query(_message.Message):
        __slots__ = ("metrics", "logs", "spans", "dataprime")
        METRICS_FIELD_NUMBER: _ClassVar[int]
        LOGS_FIELD_NUMBER: _ClassVar[int]
        SPANS_FIELD_NUMBER: _ClassVar[int]
        DATAPRIME_FIELD_NUMBER: _ClassVar[int]
        metrics: Gauge.MetricsQuery
        logs: Gauge.LogsQuery
        spans: Gauge.SpansQuery
        dataprime: Gauge.DataprimeQuery
        def __init__(self, metrics: _Optional[_Union[Gauge.MetricsQuery, _Mapping]] = ..., logs: _Optional[_Union[Gauge.LogsQuery, _Mapping]] = ..., spans: _Optional[_Union[Gauge.SpansQuery, _Mapping]] = ..., dataprime: _Optional[_Union[Gauge.DataprimeQuery, _Mapping]] = ...) -> None: ...
    class MetricsQuery(_message.Message):
        __slots__ = ("promql_query", "aggregation", "filters")
        PROMQL_QUERY_FIELD_NUMBER: _ClassVar[int]
        AGGREGATION_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        promql_query: _queries_pb2.PromQlQuery
        aggregation: Gauge.Aggregation
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.MetricsFilter]
        def __init__(self, promql_query: _Optional[_Union[_queries_pb2.PromQlQuery, _Mapping]] = ..., aggregation: _Optional[_Union[Gauge.Aggregation, str]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.MetricsFilter, _Mapping]]] = ...) -> None: ...
    class LogsQuery(_message.Message):
        __slots__ = ("lucene_query", "logs_aggregation", "aggregation", "filters")
        LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
        LOGS_AGGREGATION_FIELD_NUMBER: _ClassVar[int]
        AGGREGATION_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        lucene_query: _queries_pb2.LuceneQuery
        logs_aggregation: _logs_aggregation_pb2.LogsAggregation
        aggregation: Gauge.Aggregation
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.LogsFilter]
        def __init__(self, lucene_query: _Optional[_Union[_queries_pb2.LuceneQuery, _Mapping]] = ..., logs_aggregation: _Optional[_Union[_logs_aggregation_pb2.LogsAggregation, _Mapping]] = ..., aggregation: _Optional[_Union[Gauge.Aggregation, str]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.LogsFilter, _Mapping]]] = ...) -> None: ...
    class SpansQuery(_message.Message):
        __slots__ = ("lucene_query", "spans_aggregation", "aggregation", "filters")
        LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
        SPANS_AGGREGATION_FIELD_NUMBER: _ClassVar[int]
        AGGREGATION_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        lucene_query: _queries_pb2.LuceneQuery
        spans_aggregation: _spans_aggregation_pb2.SpansAggregation
        aggregation: Gauge.Aggregation
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.SpansFilter]
        def __init__(self, lucene_query: _Optional[_Union[_queries_pb2.LuceneQuery, _Mapping]] = ..., spans_aggregation: _Optional[_Union[_spans_aggregation_pb2.SpansAggregation, _Mapping]] = ..., aggregation: _Optional[_Union[Gauge.Aggregation, str]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.SpansFilter, _Mapping]]] = ...) -> None: ...
    class DataprimeQuery(_message.Message):
        __slots__ = ("dataprime_query", "filters")
        DATAPRIME_QUERY_FIELD_NUMBER: _ClassVar[int]
        FILTERS_FIELD_NUMBER: _ClassVar[int]
        dataprime_query: _query_pb2.DataprimeQuery
        filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter.Source]
        def __init__(self, dataprime_query: _Optional[_Union[_query_pb2.DataprimeQuery, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter.Source, _Mapping]]] = ...) -> None: ...
    class Threshold(_message.Message):
        __slots__ = ("color",)
        FROM_FIELD_NUMBER: _ClassVar[int]
        COLOR_FIELD_NUMBER: _ClassVar[int]
        color: _wrappers_pb2.StringValue
        def __init__(self, color: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., **kwargs) -> None: ...
    QUERY_FIELD_NUMBER: _ClassVar[int]
    MIN_FIELD_NUMBER: _ClassVar[int]
    MAX_FIELD_NUMBER: _ClassVar[int]
    SHOW_INNER_ARC_FIELD_NUMBER: _ClassVar[int]
    SHOW_OUTER_ARC_FIELD_NUMBER: _ClassVar[int]
    UNIT_FIELD_NUMBER: _ClassVar[int]
    THRESHOLDS_FIELD_NUMBER: _ClassVar[int]
    DATA_MODE_TYPE_FIELD_NUMBER: _ClassVar[int]
    THRESHOLD_BY_FIELD_NUMBER: _ClassVar[int]
    query: Gauge.Query
    min: _wrappers_pb2.DoubleValue
    max: _wrappers_pb2.DoubleValue
    show_inner_arc: _wrappers_pb2.BoolValue
    show_outer_arc: _wrappers_pb2.BoolValue
    unit: Gauge.Unit
    thresholds: _containers.RepeatedCompositeFieldContainer[Gauge.Threshold]
    data_mode_type: _data_mode_type_pb2.DataModeType
    threshold_by: Gauge.ThresholdBy
    def __init__(self, query: _Optional[_Union[Gauge.Query, _Mapping]] = ..., min: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ..., max: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ..., show_inner_arc: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., show_outer_arc: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., unit: _Optional[_Union[Gauge.Unit, str]] = ..., thresholds: _Optional[_Iterable[_Union[Gauge.Threshold, _Mapping]]] = ..., data_mode_type: _Optional[_Union[_data_mode_type_pb2.DataModeType, str]] = ..., threshold_by: _Optional[_Union[Gauge.ThresholdBy, str]] = ...) -> None: ...
