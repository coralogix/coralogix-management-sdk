from com.coralogix.datausage.v2 import data_usage_pb2 as _data_usage_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ExportTeamsDetailedUsageRequest(_message.Message):
    __slots__ = ("date_range", "resolution")
    DATE_RANGE_FIELD_NUMBER: _ClassVar[int]
    RESOLUTION_FIELD_NUMBER: _ClassVar[int]
    date_range: _data_usage_pb2.DateRange
    resolution: _duration_pb2.Duration
    def __init__(self, date_range: _Optional[_Union[_data_usage_pb2.DateRange, _Mapping]] = ..., resolution: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ...) -> None: ...

class DataUsage(_message.Message):
    __slots__ = ("timestamp", "size_gb", "dimensions")
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    SIZE_GB_FIELD_NUMBER: _ClassVar[int]
    DIMENSIONS_FIELD_NUMBER: _ClassVar[int]
    timestamp: _timestamp_pb2.Timestamp
    size_gb: _wrappers_pb2.FloatValue
    dimensions: _containers.RepeatedCompositeFieldContainer[_data_usage_pb2.Dimension]
    def __init__(self, timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., size_gb: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., dimensions: _Optional[_Iterable[_Union[_data_usage_pb2.Dimension, _Mapping]]] = ...) -> None: ...

class TeamDetailedUsage(_message.Message):
    __slots__ = ("team_id", "external_id", "retention", "data_usage")
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    EXTERNAL_ID_FIELD_NUMBER: _ClassVar[int]
    RETENTION_FIELD_NUMBER: _ClassVar[int]
    DATA_USAGE_FIELD_NUMBER: _ClassVar[int]
    team_id: int
    external_id: str
    retention: int
    data_usage: _containers.RepeatedCompositeFieldContainer[DataUsage]
    def __init__(self, team_id: _Optional[int] = ..., external_id: _Optional[str] = ..., retention: _Optional[int] = ..., data_usage: _Optional[_Iterable[_Union[DataUsage, _Mapping]]] = ...) -> None: ...

class ExportTeamsDetailedUsageResponse(_message.Message):
    __slots__ = ("team_usage",)
    TEAM_USAGE_FIELD_NUMBER: _ClassVar[int]
    team_usage: TeamDetailedUsage
    def __init__(self, team_usage: _Optional[_Union[TeamDetailedUsage, _Mapping]] = ...) -> None: ...
