from com.coralogixapis.logs2metrics.v2 import logs_query_pb2 as _logs_query_pb2
from com.coralogixapis.spans2metrics.v2 import spans_query_pb2 as _spans_query_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.openapi.v1 import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class E2MType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    E2M_TYPE_UNSPECIFIED: _ClassVar[E2MType]
    E2M_TYPE_LOGS2METRICS: _ClassVar[E2MType]
    E2M_TYPE_SPANS2METRICS: _ClassVar[E2MType]
E2M_TYPE_UNSPECIFIED: E2MType
E2M_TYPE_LOGS2METRICS: E2MType
E2M_TYPE_SPANS2METRICS: E2MType

class E2M(_message.Message):
    __slots__ = ("id", "name", "description", "create_time", "update_time", "permutations", "metric_labels", "metric_fields", "type", "spans_query", "logs_query", "is_internal")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    CREATE_TIME_FIELD_NUMBER: _ClassVar[int]
    UPDATE_TIME_FIELD_NUMBER: _ClassVar[int]
    PERMUTATIONS_FIELD_NUMBER: _ClassVar[int]
    METRIC_LABELS_FIELD_NUMBER: _ClassVar[int]
    METRIC_FIELDS_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    SPANS_QUERY_FIELD_NUMBER: _ClassVar[int]
    LOGS_QUERY_FIELD_NUMBER: _ClassVar[int]
    IS_INTERNAL_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    create_time: _wrappers_pb2.StringValue
    update_time: _wrappers_pb2.StringValue
    permutations: E2MPermutations
    metric_labels: _containers.RepeatedCompositeFieldContainer[MetricLabel]
    metric_fields: _containers.RepeatedCompositeFieldContainer[MetricField]
    type: E2MType
    spans_query: _spans_query_pb2.SpansQuery
    logs_query: _logs_query_pb2.LogsQuery
    is_internal: _wrappers_pb2.BoolValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., create_time: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., update_time: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., permutations: _Optional[_Union[E2MPermutations, _Mapping]] = ..., metric_labels: _Optional[_Iterable[_Union[MetricLabel, _Mapping]]] = ..., metric_fields: _Optional[_Iterable[_Union[MetricField, _Mapping]]] = ..., type: _Optional[_Union[E2MType, str]] = ..., spans_query: _Optional[_Union[_spans_query_pb2.SpansQuery, _Mapping]] = ..., logs_query: _Optional[_Union[_logs_query_pb2.LogsQuery, _Mapping]] = ..., is_internal: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class E2MCreateParams(_message.Message):
    __slots__ = ("name", "description", "permutations_limit", "metric_labels", "metric_fields", "type", "spans_query", "logs_query")
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    PERMUTATIONS_LIMIT_FIELD_NUMBER: _ClassVar[int]
    METRIC_LABELS_FIELD_NUMBER: _ClassVar[int]
    METRIC_FIELDS_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    SPANS_QUERY_FIELD_NUMBER: _ClassVar[int]
    LOGS_QUERY_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    permutations_limit: _wrappers_pb2.Int32Value
    metric_labels: _containers.RepeatedCompositeFieldContainer[MetricLabel]
    metric_fields: _containers.RepeatedCompositeFieldContainer[MetricField]
    type: E2MType
    spans_query: _spans_query_pb2.SpansQuery
    logs_query: _logs_query_pb2.LogsQuery
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., permutations_limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., metric_labels: _Optional[_Iterable[_Union[MetricLabel, _Mapping]]] = ..., metric_fields: _Optional[_Iterable[_Union[MetricField, _Mapping]]] = ..., type: _Optional[_Union[E2MType, str]] = ..., spans_query: _Optional[_Union[_spans_query_pb2.SpansQuery, _Mapping]] = ..., logs_query: _Optional[_Union[_logs_query_pb2.LogsQuery, _Mapping]] = ...) -> None: ...

class E2MPermutations(_message.Message):
    __slots__ = ("limit", "has_exceeded_limit")
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    HAS_EXCEEDED_LIMIT_FIELD_NUMBER: _ClassVar[int]
    limit: int
    has_exceeded_limit: bool
    def __init__(self, limit: _Optional[int] = ..., has_exceeded_limit: bool = ...) -> None: ...

class MetricLabel(_message.Message):
    __slots__ = ("target_label", "source_field")
    TARGET_LABEL_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_FIELD_NUMBER: _ClassVar[int]
    target_label: _wrappers_pb2.StringValue
    source_field: _wrappers_pb2.StringValue
    def __init__(self, target_label: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., source_field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class MetricField(_message.Message):
    __slots__ = ("target_base_metric_name", "source_field", "aggregations")
    TARGET_BASE_METRIC_NAME_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_FIELD_NUMBER: _ClassVar[int]
    AGGREGATIONS_FIELD_NUMBER: _ClassVar[int]
    target_base_metric_name: _wrappers_pb2.StringValue
    source_field: _wrappers_pb2.StringValue
    aggregations: _containers.RepeatedCompositeFieldContainer[Aggregation]
    def __init__(self, target_base_metric_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., source_field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., aggregations: _Optional[_Iterable[_Union[Aggregation, _Mapping]]] = ...) -> None: ...

class Aggregation(_message.Message):
    __slots__ = ("enabled", "agg_type", "target_metric_name", "samples", "histogram")
    class AggType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        AGG_TYPE_UNSPECIFIED: _ClassVar[Aggregation.AggType]
        AGG_TYPE_MIN: _ClassVar[Aggregation.AggType]
        AGG_TYPE_MAX: _ClassVar[Aggregation.AggType]
        AGG_TYPE_COUNT: _ClassVar[Aggregation.AggType]
        AGG_TYPE_AVG: _ClassVar[Aggregation.AggType]
        AGG_TYPE_SUM: _ClassVar[Aggregation.AggType]
        AGG_TYPE_HISTOGRAM: _ClassVar[Aggregation.AggType]
        AGG_TYPE_SAMPLES: _ClassVar[Aggregation.AggType]
    AGG_TYPE_UNSPECIFIED: Aggregation.AggType
    AGG_TYPE_MIN: Aggregation.AggType
    AGG_TYPE_MAX: Aggregation.AggType
    AGG_TYPE_COUNT: Aggregation.AggType
    AGG_TYPE_AVG: Aggregation.AggType
    AGG_TYPE_SUM: Aggregation.AggType
    AGG_TYPE_HISTOGRAM: Aggregation.AggType
    AGG_TYPE_SAMPLES: Aggregation.AggType
    ENABLED_FIELD_NUMBER: _ClassVar[int]
    AGG_TYPE_FIELD_NUMBER: _ClassVar[int]
    TARGET_METRIC_NAME_FIELD_NUMBER: _ClassVar[int]
    SAMPLES_FIELD_NUMBER: _ClassVar[int]
    HISTOGRAM_FIELD_NUMBER: _ClassVar[int]
    enabled: bool
    agg_type: Aggregation.AggType
    target_metric_name: str
    samples: E2MAggSamples
    histogram: E2MAggHistogram
    def __init__(self, enabled: bool = ..., agg_type: _Optional[_Union[Aggregation.AggType, str]] = ..., target_metric_name: _Optional[str] = ..., samples: _Optional[_Union[E2MAggSamples, _Mapping]] = ..., histogram: _Optional[_Union[E2MAggHistogram, _Mapping]] = ...) -> None: ...

class E2MAggSamples(_message.Message):
    __slots__ = ("sample_type",)
    class SampleType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        SAMPLE_TYPE_UNSPECIFIED: _ClassVar[E2MAggSamples.SampleType]
        SAMPLE_TYPE_MIN: _ClassVar[E2MAggSamples.SampleType]
        SAMPLE_TYPE_MAX: _ClassVar[E2MAggSamples.SampleType]
    SAMPLE_TYPE_UNSPECIFIED: E2MAggSamples.SampleType
    SAMPLE_TYPE_MIN: E2MAggSamples.SampleType
    SAMPLE_TYPE_MAX: E2MAggSamples.SampleType
    SAMPLE_TYPE_FIELD_NUMBER: _ClassVar[int]
    sample_type: E2MAggSamples.SampleType
    def __init__(self, sample_type: _Optional[_Union[E2MAggSamples.SampleType, str]] = ...) -> None: ...

class E2MAggHistogram(_message.Message):
    __slots__ = ("buckets",)
    BUCKETS_FIELD_NUMBER: _ClassVar[int]
    buckets: _containers.RepeatedScalarFieldContainer[float]
    def __init__(self, buckets: _Optional[_Iterable[float]] = ...) -> None: ...
