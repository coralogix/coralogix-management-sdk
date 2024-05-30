from com.coralogix.datausage.v2 import data_usage_pb2 as _data_usage_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AggregateBy(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    AGGREGATE_BY_UNSPECIFIED: _ClassVar[AggregateBy]
    AGGREGATE_BY_APPLICATION: _ClassVar[AggregateBy]
    AGGREGATE_BY_SUBSYSTEM: _ClassVar[AggregateBy]
    AGGREGATE_BY_PILLAR: _ClassVar[AggregateBy]
    AGGREGATE_BY_PRIORITY: _ClassVar[AggregateBy]
AGGREGATE_BY_UNSPECIFIED: AggregateBy
AGGREGATE_BY_APPLICATION: AggregateBy
AGGREGATE_BY_SUBSYSTEM: AggregateBy
AGGREGATE_BY_PILLAR: AggregateBy
AGGREGATE_BY_PRIORITY: AggregateBy

class GetTeamDetailedDataUsageRequest(_message.Message):
    __slots__ = ("date_range", "resolution", "aggregate")
    DATE_RANGE_FIELD_NUMBER: _ClassVar[int]
    RESOLUTION_FIELD_NUMBER: _ClassVar[int]
    AGGREGATE_FIELD_NUMBER: _ClassVar[int]
    date_range: _data_usage_pb2.DateRange
    resolution: _duration_pb2.Duration
    aggregate: _containers.RepeatedScalarFieldContainer[AggregateBy]
    def __init__(self, date_range: _Optional[_Union[_data_usage_pb2.DateRange, _Mapping]] = ..., resolution: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., aggregate: _Optional[_Iterable[_Union[AggregateBy, str]]] = ...) -> None: ...

class GetTeamDetailedDataUsageResponse(_message.Message):
    __slots__ = ("timestamp", "size_gb", "units", "dimensions")
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    SIZE_GB_FIELD_NUMBER: _ClassVar[int]
    UNITS_FIELD_NUMBER: _ClassVar[int]
    DIMENSIONS_FIELD_NUMBER: _ClassVar[int]
    timestamp: _timestamp_pb2.Timestamp
    size_gb: _wrappers_pb2.FloatValue
    units: _wrappers_pb2.FloatValue
    dimensions: _containers.RepeatedCompositeFieldContainer[_data_usage_pb2.Dimension]
    def __init__(self, timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., size_gb: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., units: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., dimensions: _Optional[_Iterable[_Union[_data_usage_pb2.Dimension, _Mapping]]] = ...) -> None: ...

class GetSpansCountRequest(_message.Message):
    __slots__ = ("date_range", "resolution", "filters")
    DATE_RANGE_FIELD_NUMBER: _ClassVar[int]
    RESOLUTION_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    date_range: _data_usage_pb2.DateRange
    resolution: _duration_pb2.Duration
    filters: ScopesFilter
    def __init__(self, date_range: _Optional[_Union[_data_usage_pb2.DateRange, _Mapping]] = ..., resolution: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., filters: _Optional[_Union[ScopesFilter, _Mapping]] = ...) -> None: ...

class ScopesFilter(_message.Message):
    __slots__ = ("application", "subsystem")
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
    application: _containers.RepeatedScalarFieldContainer[str]
    subsystem: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, application: _Optional[_Iterable[str]] = ..., subsystem: _Optional[_Iterable[str]] = ...) -> None: ...

class GetSpansCountResponse(_message.Message):
    __slots__ = ("timestamp", "success_span_count", "error_span_count")
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    SUCCESS_SPAN_COUNT_FIELD_NUMBER: _ClassVar[int]
    ERROR_SPAN_COUNT_FIELD_NUMBER: _ClassVar[int]
    timestamp: _timestamp_pb2.Timestamp
    success_span_count: _wrappers_pb2.Int64Value
    error_span_count: _wrappers_pb2.Int64Value
    def __init__(self, timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., success_span_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., error_span_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class GetDataUsageMetricsExportStatusRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetDataUsageMetricsExportStatusResponse(_message.Message):
    __slots__ = ("enabled",)
    ENABLED_FIELD_NUMBER: _ClassVar[int]
    enabled: bool
    def __init__(self, enabled: bool = ...) -> None: ...

class UpdateDataUsageMetricsExportStatusRequest(_message.Message):
    __slots__ = ("enabled",)
    ENABLED_FIELD_NUMBER: _ClassVar[int]
    enabled: bool
    def __init__(self, enabled: bool = ...) -> None: ...

class UpdateDataUsageMetricsExportStatusResponse(_message.Message):
    __slots__ = ("enabled",)
    ENABLED_FIELD_NUMBER: _ClassVar[int]
    enabled: bool
    def __init__(self, enabled: bool = ...) -> None: ...
