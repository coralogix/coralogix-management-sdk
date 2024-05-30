from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from com.coralogix.dataprime.v1 import common_pb2 as _common_pb2
from com.coralogix.dataprime.v1 import histogram_pb2 as _histogram_pb2
from com.coralogix.dataprime.v1 import query_pb2 as _query_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class RunDataprimeQueryRequest(_message.Message):
    __slots__ = ("query_id", "query", "start_date", "end_date", "filters", "sort_model", "backend", "components", "tracing_metadata", "use_text_index", "requested_team_ids", "default_filter_horizon", "internal_client_id")
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    START_DATE_FIELD_NUMBER: _ClassVar[int]
    END_DATE_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    SORT_MODEL_FIELD_NUMBER: _ClassVar[int]
    BACKEND_FIELD_NUMBER: _ClassVar[int]
    COMPONENTS_FIELD_NUMBER: _ClassVar[int]
    TRACING_METADATA_FIELD_NUMBER: _ClassVar[int]
    USE_TEXT_INDEX_FIELD_NUMBER: _ClassVar[int]
    REQUESTED_TEAM_IDS_FIELD_NUMBER: _ClassVar[int]
    DEFAULT_FILTER_HORIZON_FIELD_NUMBER: _ClassVar[int]
    INTERNAL_CLIENT_ID_FIELD_NUMBER: _ClassVar[int]
    query_id: _wrappers_pb2.StringValue
    query: _query_pb2.QueryPayload
    start_date: _timestamp_pb2.Timestamp
    end_date: _timestamp_pb2.Timestamp
    filters: _containers.RepeatedCompositeFieldContainer[Filter]
    sort_model: _containers.RepeatedCompositeFieldContainer[SortModelItem]
    backend: _common_pb2.Backend
    components: _containers.RepeatedCompositeFieldContainer[RequestComponent]
    tracing_metadata: TracingMetadata
    use_text_index: bool
    requested_team_ids: _containers.RepeatedScalarFieldContainer[int]
    default_filter_horizon: Filter.Horizon
    internal_client_id: str
    def __init__(self, query_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., query: _Optional[_Union[_query_pb2.QueryPayload, _Mapping]] = ..., start_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., end_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., filters: _Optional[_Iterable[_Union[Filter, _Mapping]]] = ..., sort_model: _Optional[_Iterable[_Union[SortModelItem, _Mapping]]] = ..., backend: _Optional[_Union[_common_pb2.Backend, str]] = ..., components: _Optional[_Iterable[_Union[RequestComponent, _Mapping]]] = ..., tracing_metadata: _Optional[_Union[TracingMetadata, _Mapping]] = ..., use_text_index: bool = ..., requested_team_ids: _Optional[_Iterable[int]] = ..., default_filter_horizon: _Optional[_Union[Filter.Horizon, str]] = ..., internal_client_id: _Optional[str] = ...) -> None: ...

class RequestComponent(_message.Message):
    __slots__ = ("logs", "severity_histogram", "surrounding_logs", "count", "key_distribution", "histogram", "compared_histogram", "sidebar_filters_count", "team_id_counts")
    LOGS_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_HISTOGRAM_FIELD_NUMBER: _ClassVar[int]
    SURROUNDING_LOGS_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    KEY_DISTRIBUTION_FIELD_NUMBER: _ClassVar[int]
    HISTOGRAM_FIELD_NUMBER: _ClassVar[int]
    COMPARED_HISTOGRAM_FIELD_NUMBER: _ClassVar[int]
    SIDEBAR_FILTERS_COUNT_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_COUNTS_FIELD_NUMBER: _ClassVar[int]
    logs: Component.Logs
    severity_histogram: Component.SeverityHistogram
    surrounding_logs: Component.SurroundingLogs
    count: Component.Count
    key_distribution: Component.KeyDistribution
    histogram: Component.Histogram
    compared_histogram: Component.ComparedHistogram
    sidebar_filters_count: Component.SidebarFiltersCounts
    team_id_counts: Component.TeamIdCounts
    def __init__(self, logs: _Optional[_Union[Component.Logs, _Mapping]] = ..., severity_histogram: _Optional[_Union[Component.SeverityHistogram, _Mapping]] = ..., surrounding_logs: _Optional[_Union[Component.SurroundingLogs, _Mapping]] = ..., count: _Optional[_Union[Component.Count, _Mapping]] = ..., key_distribution: _Optional[_Union[Component.KeyDistribution, _Mapping]] = ..., histogram: _Optional[_Union[Component.Histogram, _Mapping]] = ..., compared_histogram: _Optional[_Union[Component.ComparedHistogram, _Mapping]] = ..., sidebar_filters_count: _Optional[_Union[Component.SidebarFiltersCounts, _Mapping]] = ..., team_id_counts: _Optional[_Union[Component.TeamIdCounts, _Mapping]] = ...) -> None: ...

class Component(_message.Message):
    __slots__ = ()
    class Logs(_message.Message):
        __slots__ = ("skip", "limit")
        SKIP_FIELD_NUMBER: _ClassVar[int]
        LIMIT_FIELD_NUMBER: _ClassVar[int]
        skip: _wrappers_pb2.Int64Value
        limit: _wrappers_pb2.Int64Value
        def __init__(self, skip: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...
    class Count(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class SeverityHistogram(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class SidebarFiltersCounts(_message.Message):
        __slots__ = ("count_unselected_values",)
        COUNT_UNSELECTED_VALUES_FIELD_NUMBER: _ClassVar[int]
        count_unselected_values: bool
        def __init__(self, count_unselected_values: bool = ...) -> None: ...
    class SurroundingLogs(_message.Message):
        __slots__ = ("log_id", "log_timestamp", "skip", "limit")
        LOG_ID_FIELD_NUMBER: _ClassVar[int]
        LOG_TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
        SKIP_FIELD_NUMBER: _ClassVar[int]
        LIMIT_FIELD_NUMBER: _ClassVar[int]
        log_id: _wrappers_pb2.StringValue
        log_timestamp: _timestamp_pb2.Timestamp
        skip: _wrappers_pb2.Int32Value
        limit: _wrappers_pb2.Int32Value
        def __init__(self, log_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., log_timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., skip: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    class KeyDistribution(_message.Message):
        __slots__ = ("key_distribution",)
        KEY_DISTRIBUTION_FIELD_NUMBER: _ClassVar[int]
        key_distribution: _common_pb2.UntypedKeypath
        def __init__(self, key_distribution: _Optional[_Union[_common_pb2.UntypedKeypath, _Mapping]] = ...) -> None: ...
    class Histogram(_message.Message):
        __slots__ = ("aggregation", "group_by")
        AGGREGATION_FIELD_NUMBER: _ClassVar[int]
        GROUP_BY_FIELD_NUMBER: _ClassVar[int]
        aggregation: _histogram_pb2.Aggregation
        group_by: _common_pb2.UntypedKeypath
        def __init__(self, aggregation: _Optional[_Union[_histogram_pb2.Aggregation, _Mapping]] = ..., group_by: _Optional[_Union[_common_pb2.UntypedKeypath, _Mapping]] = ...) -> None: ...
    class ComparedHistogram(_message.Message):
        __slots__ = ("aggregation", "group_by", "historical_shift_ms")
        AGGREGATION_FIELD_NUMBER: _ClassVar[int]
        GROUP_BY_FIELD_NUMBER: _ClassVar[int]
        HISTORICAL_SHIFT_MS_FIELD_NUMBER: _ClassVar[int]
        aggregation: _histogram_pb2.Aggregation
        group_by: _common_pb2.UntypedKeypath
        historical_shift_ms: _wrappers_pb2.Int64Value
        def __init__(self, aggregation: _Optional[_Union[_histogram_pb2.Aggregation, _Mapping]] = ..., group_by: _Optional[_Union[_common_pb2.UntypedKeypath, _Mapping]] = ..., historical_shift_ms: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...
    class TeamIdCounts(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    def __init__(self) -> None: ...

class Filter(_message.Message):
    __slots__ = ("field", "alternatives")
    class Horizon(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        HORIZON_BEFORE_QUERY_UNSPECIFIED: _ClassVar[Filter.Horizon]
        HORIZON_AFTER_QUERY: _ClassVar[Filter.Horizon]
    HORIZON_BEFORE_QUERY_UNSPECIFIED: Filter.Horizon
    HORIZON_AFTER_QUERY: Filter.Horizon
    FIELD_FIELD_NUMBER: _ClassVar[int]
    ALTERNATIVES_FIELD_NUMBER: _ClassVar[int]
    field: _common_pb2.UntypedKeypath
    alternatives: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, field: _Optional[_Union[_common_pb2.UntypedKeypath, _Mapping]] = ..., alternatives: _Optional[_Iterable[str]] = ...) -> None: ...

class SortModelItem(_message.Message):
    __slots__ = ("field", "ordering", "missing")
    class Ordering(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        ORDERING_ASC_UNSPECIFIED: _ClassVar[SortModelItem.Ordering]
        ORDERING_DESC: _ClassVar[SortModelItem.Ordering]
    ORDERING_ASC_UNSPECIFIED: SortModelItem.Ordering
    ORDERING_DESC: SortModelItem.Ordering
    class Missing(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        MISSING_LAST_UNSPECIFIED: _ClassVar[SortModelItem.Missing]
        MISSING_FIRST: _ClassVar[SortModelItem.Missing]
    MISSING_LAST_UNSPECIFIED: SortModelItem.Missing
    MISSING_FIRST: SortModelItem.Missing
    FIELD_FIELD_NUMBER: _ClassVar[int]
    ORDERING_FIELD_NUMBER: _ClassVar[int]
    MISSING_FIELD_NUMBER: _ClassVar[int]
    field: _common_pb2.UntypedKeypath
    ordering: SortModelItem.Ordering
    missing: SortModelItem.Missing
    def __init__(self, field: _Optional[_Union[_common_pb2.UntypedKeypath, _Mapping]] = ..., ordering: _Optional[_Union[SortModelItem.Ordering, str]] = ..., missing: _Optional[_Union[SortModelItem.Missing, str]] = ...) -> None: ...

class TracingMetadata(_message.Message):
    __slots__ = ("query_syntax", "query_text")
    class QuerySyntax(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        QUERY_SYNTAX_LUCENE_UNSPECIFIED: _ClassVar[TracingMetadata.QuerySyntax]
        QUERY_SYNTAX_DATAPRIME: _ClassVar[TracingMetadata.QuerySyntax]
    QUERY_SYNTAX_LUCENE_UNSPECIFIED: TracingMetadata.QuerySyntax
    QUERY_SYNTAX_DATAPRIME: TracingMetadata.QuerySyntax
    QUERY_SYNTAX_FIELD_NUMBER: _ClassVar[int]
    QUERY_TEXT_FIELD_NUMBER: _ClassVar[int]
    query_syntax: TracingMetadata.QuerySyntax
    query_text: _wrappers_pb2.StringValue
    def __init__(self, query_syntax: _Optional[_Union[TracingMetadata.QuerySyntax, str]] = ..., query_text: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
