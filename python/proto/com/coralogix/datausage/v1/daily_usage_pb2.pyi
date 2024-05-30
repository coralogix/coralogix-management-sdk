from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DetailedDailyUsage(_message.Message):
    __slots__ = ("stats_date", "block", "block_count_quota", "low", "low_count_quota", "medium", "medium_count_quota", "high", "high_count_quota", "total", "total_count_quota", "high_metrics_count_quota", "high_metrics_quota", "high_tracing_count_quota", "high_tracing_quota", "medium_tracing_count_quota", "medium_tracing_quota", "low_tracing_count_quota", "low_tracing_quota", "low_session_recording_count_quota", "low_session_recording_quota")
    STATS_DATE_FIELD_NUMBER: _ClassVar[int]
    BLOCK_FIELD_NUMBER: _ClassVar[int]
    BLOCK_COUNT_QUOTA_FIELD_NUMBER: _ClassVar[int]
    LOW_FIELD_NUMBER: _ClassVar[int]
    LOW_COUNT_QUOTA_FIELD_NUMBER: _ClassVar[int]
    MEDIUM_FIELD_NUMBER: _ClassVar[int]
    MEDIUM_COUNT_QUOTA_FIELD_NUMBER: _ClassVar[int]
    HIGH_FIELD_NUMBER: _ClassVar[int]
    HIGH_COUNT_QUOTA_FIELD_NUMBER: _ClassVar[int]
    TOTAL_FIELD_NUMBER: _ClassVar[int]
    TOTAL_COUNT_QUOTA_FIELD_NUMBER: _ClassVar[int]
    HIGH_METRICS_COUNT_QUOTA_FIELD_NUMBER: _ClassVar[int]
    HIGH_METRICS_QUOTA_FIELD_NUMBER: _ClassVar[int]
    HIGH_TRACING_COUNT_QUOTA_FIELD_NUMBER: _ClassVar[int]
    HIGH_TRACING_QUOTA_FIELD_NUMBER: _ClassVar[int]
    MEDIUM_TRACING_COUNT_QUOTA_FIELD_NUMBER: _ClassVar[int]
    MEDIUM_TRACING_QUOTA_FIELD_NUMBER: _ClassVar[int]
    LOW_TRACING_COUNT_QUOTA_FIELD_NUMBER: _ClassVar[int]
    LOW_TRACING_QUOTA_FIELD_NUMBER: _ClassVar[int]
    LOW_SESSION_RECORDING_COUNT_QUOTA_FIELD_NUMBER: _ClassVar[int]
    LOW_SESSION_RECORDING_QUOTA_FIELD_NUMBER: _ClassVar[int]
    stats_date: _timestamp_pb2.Timestamp
    block: _wrappers_pb2.FloatValue
    block_count_quota: _wrappers_pb2.FloatValue
    low: _wrappers_pb2.FloatValue
    low_count_quota: _wrappers_pb2.FloatValue
    medium: _wrappers_pb2.FloatValue
    medium_count_quota: _wrappers_pb2.FloatValue
    high: _wrappers_pb2.FloatValue
    high_count_quota: _wrappers_pb2.FloatValue
    total: _wrappers_pb2.FloatValue
    total_count_quota: _wrappers_pb2.FloatValue
    high_metrics_count_quota: _wrappers_pb2.FloatValue
    high_metrics_quota: _wrappers_pb2.FloatValue
    high_tracing_count_quota: _wrappers_pb2.FloatValue
    high_tracing_quota: _wrappers_pb2.FloatValue
    medium_tracing_count_quota: _wrappers_pb2.FloatValue
    medium_tracing_quota: _wrappers_pb2.FloatValue
    low_tracing_count_quota: _wrappers_pb2.FloatValue
    low_tracing_quota: _wrappers_pb2.FloatValue
    low_session_recording_count_quota: _wrappers_pb2.FloatValue
    low_session_recording_quota: _wrappers_pb2.FloatValue
    def __init__(self, stats_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., block: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., block_count_quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., low: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., low_count_quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., medium: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., medium_count_quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., high: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., high_count_quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., total: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., total_count_quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., high_metrics_count_quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., high_metrics_quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., high_tracing_count_quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., high_tracing_quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., medium_tracing_count_quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., medium_tracing_quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., low_tracing_count_quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., low_tracing_quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., low_session_recording_count_quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., low_session_recording_quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ...) -> None: ...

class OverageDailyReport(_message.Message):
    __slots__ = ("date", "overage_unit")
    DATE_FIELD_NUMBER: _ClassVar[int]
    OVERAGE_UNIT_FIELD_NUMBER: _ClassVar[int]
    date: _timestamp_pb2.Timestamp
    overage_unit: _wrappers_pb2.UInt64Value
    def __init__(self, date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., overage_unit: _Optional[_Union[_wrappers_pb2.UInt64Value, _Mapping]] = ...) -> None: ...
