from com.coralogixapis.dashboards.v1.common import observation_field_pb2 as _observation_field_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class LogsAggregation(_message.Message):
    __slots__ = ("count", "count_distinct", "sum", "average", "min", "max", "percentile")
    class Count(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class CountDistinct(_message.Message):
        __slots__ = ("field", "observation_field")
        FIELD_FIELD_NUMBER: _ClassVar[int]
        OBSERVATION_FIELD_FIELD_NUMBER: _ClassVar[int]
        field: _wrappers_pb2.StringValue
        observation_field: _observation_field_pb2.ObservationField
        def __init__(self, field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., observation_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
    class Sum(_message.Message):
        __slots__ = ("field", "observation_field")
        FIELD_FIELD_NUMBER: _ClassVar[int]
        OBSERVATION_FIELD_FIELD_NUMBER: _ClassVar[int]
        field: _wrappers_pb2.StringValue
        observation_field: _observation_field_pb2.ObservationField
        def __init__(self, field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., observation_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
    class Average(_message.Message):
        __slots__ = ("field", "observation_field")
        FIELD_FIELD_NUMBER: _ClassVar[int]
        OBSERVATION_FIELD_FIELD_NUMBER: _ClassVar[int]
        field: _wrappers_pb2.StringValue
        observation_field: _observation_field_pb2.ObservationField
        def __init__(self, field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., observation_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
    class Min(_message.Message):
        __slots__ = ("field", "observation_field")
        FIELD_FIELD_NUMBER: _ClassVar[int]
        OBSERVATION_FIELD_FIELD_NUMBER: _ClassVar[int]
        field: _wrappers_pb2.StringValue
        observation_field: _observation_field_pb2.ObservationField
        def __init__(self, field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., observation_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
    class Max(_message.Message):
        __slots__ = ("field", "observation_field")
        FIELD_FIELD_NUMBER: _ClassVar[int]
        OBSERVATION_FIELD_FIELD_NUMBER: _ClassVar[int]
        field: _wrappers_pb2.StringValue
        observation_field: _observation_field_pb2.ObservationField
        def __init__(self, field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., observation_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
    class Percentile(_message.Message):
        __slots__ = ("field", "percent", "observation_field")
        FIELD_FIELD_NUMBER: _ClassVar[int]
        PERCENT_FIELD_NUMBER: _ClassVar[int]
        OBSERVATION_FIELD_FIELD_NUMBER: _ClassVar[int]
        field: _wrappers_pb2.StringValue
        percent: _wrappers_pb2.DoubleValue
        observation_field: _observation_field_pb2.ObservationField
        def __init__(self, field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., percent: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ..., observation_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
    COUNT_FIELD_NUMBER: _ClassVar[int]
    COUNT_DISTINCT_FIELD_NUMBER: _ClassVar[int]
    SUM_FIELD_NUMBER: _ClassVar[int]
    AVERAGE_FIELD_NUMBER: _ClassVar[int]
    MIN_FIELD_NUMBER: _ClassVar[int]
    MAX_FIELD_NUMBER: _ClassVar[int]
    PERCENTILE_FIELD_NUMBER: _ClassVar[int]
    count: LogsAggregation.Count
    count_distinct: LogsAggregation.CountDistinct
    sum: LogsAggregation.Sum
    average: LogsAggregation.Average
    min: LogsAggregation.Min
    max: LogsAggregation.Max
    percentile: LogsAggregation.Percentile
    def __init__(self, count: _Optional[_Union[LogsAggregation.Count, _Mapping]] = ..., count_distinct: _Optional[_Union[LogsAggregation.CountDistinct, _Mapping]] = ..., sum: _Optional[_Union[LogsAggregation.Sum, _Mapping]] = ..., average: _Optional[_Union[LogsAggregation.Average, _Mapping]] = ..., min: _Optional[_Union[LogsAggregation.Min, _Mapping]] = ..., max: _Optional[_Union[LogsAggregation.Max, _Mapping]] = ..., percentile: _Optional[_Union[LogsAggregation.Percentile, _Mapping]] = ...) -> None: ...
