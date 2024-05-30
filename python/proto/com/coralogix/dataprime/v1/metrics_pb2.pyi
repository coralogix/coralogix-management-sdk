from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class MetricType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    METRIC_TYPE_UNSPECIFIED: _ClassVar[MetricType]
    METRIC_TYPE_TIMESTAMP: _ClassVar[MetricType]
    METRIC_TYPE_COUNT: _ClassVar[MetricType]
    METRIC_TYPE_GAUGE: _ClassVar[MetricType]
    METRIC_TYPE_TIME: _ClassVar[MetricType]
METRIC_TYPE_UNSPECIFIED: MetricType
METRIC_TYPE_TIMESTAMP: MetricType
METRIC_TYPE_COUNT: MetricType
METRIC_TYPE_GAUGE: MetricType
METRIC_TYPE_TIME: MetricType

class QueryMetrics(_message.Message):
    __slots__ = ("query_id", "stage_metrics")
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    STAGE_METRICS_FIELD_NUMBER: _ClassVar[int]
    query_id: str
    stage_metrics: _containers.RepeatedCompositeFieldContainer[StageMetrics]
    def __init__(self, query_id: _Optional[str] = ..., stage_metrics: _Optional[_Iterable[_Union[StageMetrics, _Mapping]]] = ...) -> None: ...

class StageMetrics(_message.Message):
    __slots__ = ("stage_id", "partition_id", "operator_metrics")
    STAGE_ID_FIELD_NUMBER: _ClassVar[int]
    PARTITION_ID_FIELD_NUMBER: _ClassVar[int]
    OPERATOR_METRICS_FIELD_NUMBER: _ClassVar[int]
    stage_id: int
    partition_id: int
    operator_metrics: _containers.RepeatedCompositeFieldContainer[OperatorMetrics]
    def __init__(self, stage_id: _Optional[int] = ..., partition_id: _Optional[int] = ..., operator_metrics: _Optional[_Iterable[_Union[OperatorMetrics, _Mapping]]] = ...) -> None: ...

class OperatorMetrics(_message.Message):
    __slots__ = ("operator_name", "metrics")
    OPERATOR_NAME_FIELD_NUMBER: _ClassVar[int]
    METRICS_FIELD_NUMBER: _ClassVar[int]
    operator_name: str
    metrics: _containers.RepeatedCompositeFieldContainer[Metric]
    def __init__(self, operator_name: _Optional[str] = ..., metrics: _Optional[_Iterable[_Union[Metric, _Mapping]]] = ...) -> None: ...

class MetricLabel(_message.Message):
    __slots__ = ("name", "value")
    NAME_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    name: str
    value: str
    def __init__(self, name: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...

class Metric(_message.Message):
    __slots__ = ("value", "name", "metric_type", "labels")
    VALUE_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    METRIC_TYPE_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    value: int
    name: str
    metric_type: MetricType
    labels: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, value: _Optional[int] = ..., name: _Optional[str] = ..., metric_type: _Optional[_Union[MetricType, str]] = ..., labels: _Optional[_Iterable[str]] = ...) -> None: ...

class GetMetricsRequest(_message.Message):
    __slots__ = ("query_id",)
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    query_id: str
    def __init__(self, query_id: _Optional[str] = ...) -> None: ...

class GetMetricsResponse(_message.Message):
    __slots__ = ("metrics",)
    METRICS_FIELD_NUMBER: _ClassVar[int]
    metrics: QueryMetrics
    def __init__(self, metrics: _Optional[_Union[QueryMetrics, _Mapping]] = ...) -> None: ...
