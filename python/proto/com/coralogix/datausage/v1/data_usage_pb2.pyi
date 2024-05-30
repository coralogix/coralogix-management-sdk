from com.coralogix.datausage.v1 import common_pb2 as _common_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Pillar(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    PILLAR_UNSPECIFIED: _ClassVar[Pillar]
    PILLAR_METRICS: _ClassVar[Pillar]
    PILLAR_LOGS: _ClassVar[Pillar]
    PILLAR_SPANS: _ClassVar[Pillar]
    PILLAR_BINARY: _ClassVar[Pillar]
PILLAR_UNSPECIFIED: Pillar
PILLAR_METRICS: Pillar
PILLAR_LOGS: Pillar
PILLAR_SPANS: Pillar
PILLAR_BINARY: Pillar

class TeamsAndTimeRange(_message.Message):
    __slots__ = ("teams", "range")
    TEAMS_FIELD_NUMBER: _ClassVar[int]
    RANGE_FIELD_NUMBER: _ClassVar[int]
    teams: _containers.RepeatedCompositeFieldContainer[_common_pb2.Team]
    range: _common_pb2.Range
    def __init__(self, teams: _Optional[_Iterable[_Union[_common_pb2.Team, _Mapping]]] = ..., range: _Optional[_Union[_common_pb2.Range, str]] = ...) -> None: ...

class TeamsAndTime(_message.Message):
    __slots__ = ("teams", "time")
    TEAMS_FIELD_NUMBER: _ClassVar[int]
    TIME_FIELD_NUMBER: _ClassVar[int]
    teams: _containers.RepeatedCompositeFieldContainer[_common_pb2.Team]
    time: _timestamp_pb2.Timestamp
    def __init__(self, teams: _Optional[_Iterable[_Union[_common_pb2.Team, _Mapping]]] = ..., time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class UnitUsage(_message.Message):
    __slots__ = ("low", "medium", "high", "blocked")
    LOW_FIELD_NUMBER: _ClassVar[int]
    MEDIUM_FIELD_NUMBER: _ClassVar[int]
    HIGH_FIELD_NUMBER: _ClassVar[int]
    BLOCKED_FIELD_NUMBER: _ClassVar[int]
    low: _common_pb2.Unit
    medium: _common_pb2.Unit
    high: _common_pb2.Unit
    blocked: _common_pb2.Unit
    def __init__(self, low: _Optional[_Union[_common_pb2.Unit, _Mapping]] = ..., medium: _Optional[_Union[_common_pb2.Unit, _Mapping]] = ..., high: _Optional[_Union[_common_pb2.Unit, _Mapping]] = ..., blocked: _Optional[_Union[_common_pb2.Unit, _Mapping]] = ...) -> None: ...

class GBUsage(_message.Message):
    __slots__ = ("low", "medium", "high", "blocked")
    LOW_FIELD_NUMBER: _ClassVar[int]
    MEDIUM_FIELD_NUMBER: _ClassVar[int]
    HIGH_FIELD_NUMBER: _ClassVar[int]
    BLOCKED_FIELD_NUMBER: _ClassVar[int]
    low: _common_pb2.GB
    medium: _common_pb2.GB
    high: _common_pb2.GB
    blocked: _common_pb2.GB
    def __init__(self, low: _Optional[_Union[_common_pb2.GB, _Mapping]] = ..., medium: _Optional[_Union[_common_pb2.GB, _Mapping]] = ..., high: _Optional[_Union[_common_pb2.GB, _Mapping]] = ..., blocked: _Optional[_Union[_common_pb2.GB, _Mapping]] = ...) -> None: ...

class DailyUsageMetrics(_message.Message):
    __slots__ = ("daily_avg", "daily_usage", "date", "pay_as_you_go_usage", "daily_avg_data_sent", "daily_data_sent", "logs_sent", "metrics_sent", "traces_sent", "logs_quota", "metrics_quota", "traces_quota", "session_recording_sent", "session_recording_quota")
    DAILY_AVG_FIELD_NUMBER: _ClassVar[int]
    DAILY_USAGE_FIELD_NUMBER: _ClassVar[int]
    DATE_FIELD_NUMBER: _ClassVar[int]
    PAY_AS_YOU_GO_USAGE_FIELD_NUMBER: _ClassVar[int]
    DAILY_AVG_DATA_SENT_FIELD_NUMBER: _ClassVar[int]
    DAILY_DATA_SENT_FIELD_NUMBER: _ClassVar[int]
    LOGS_SENT_FIELD_NUMBER: _ClassVar[int]
    METRICS_SENT_FIELD_NUMBER: _ClassVar[int]
    TRACES_SENT_FIELD_NUMBER: _ClassVar[int]
    LOGS_QUOTA_FIELD_NUMBER: _ClassVar[int]
    METRICS_QUOTA_FIELD_NUMBER: _ClassVar[int]
    TRACES_QUOTA_FIELD_NUMBER: _ClassVar[int]
    SESSION_RECORDING_SENT_FIELD_NUMBER: _ClassVar[int]
    SESSION_RECORDING_QUOTA_FIELD_NUMBER: _ClassVar[int]
    daily_avg: _common_pb2.Unit
    daily_usage: _common_pb2.Unit
    date: _timestamp_pb2.Timestamp
    pay_as_you_go_usage: _common_pb2.Unit
    daily_avg_data_sent: _common_pb2.GB
    daily_data_sent: _common_pb2.GB
    logs_sent: GBUsage
    metrics_sent: GBUsage
    traces_sent: GBUsage
    logs_quota: UnitUsage
    metrics_quota: UnitUsage
    traces_quota: UnitUsage
    session_recording_sent: GBUsage
    session_recording_quota: UnitUsage
    def __init__(self, daily_avg: _Optional[_Union[_common_pb2.Unit, _Mapping]] = ..., daily_usage: _Optional[_Union[_common_pb2.Unit, _Mapping]] = ..., date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., pay_as_you_go_usage: _Optional[_Union[_common_pb2.Unit, _Mapping]] = ..., daily_avg_data_sent: _Optional[_Union[_common_pb2.GB, _Mapping]] = ..., daily_data_sent: _Optional[_Union[_common_pb2.GB, _Mapping]] = ..., logs_sent: _Optional[_Union[GBUsage, _Mapping]] = ..., metrics_sent: _Optional[_Union[GBUsage, _Mapping]] = ..., traces_sent: _Optional[_Union[GBUsage, _Mapping]] = ..., logs_quota: _Optional[_Union[UnitUsage, _Mapping]] = ..., metrics_quota: _Optional[_Union[UnitUsage, _Mapping]] = ..., traces_quota: _Optional[_Union[UnitUsage, _Mapping]] = ..., session_recording_sent: _Optional[_Union[GBUsage, _Mapping]] = ..., session_recording_quota: _Optional[_Union[UnitUsage, _Mapping]] = ...) -> None: ...

class TeamsDailyDataUsage(_message.Message):
    __slots__ = ("team", "metrics")
    TEAM_FIELD_NUMBER: _ClassVar[int]
    METRICS_FIELD_NUMBER: _ClassVar[int]
    team: _common_pb2.Team
    metrics: _containers.RepeatedCompositeFieldContainer[DailyUsageMetrics]
    def __init__(self, team: _Optional[_Union[_common_pb2.Team, _Mapping]] = ..., metrics: _Optional[_Iterable[_Union[DailyUsageMetrics, _Mapping]]] = ...) -> None: ...

class TeamBlockEvents(_message.Message):
    __slots__ = ("team", "block_event")
    TEAM_FIELD_NUMBER: _ClassVar[int]
    BLOCK_EVENT_FIELD_NUMBER: _ClassVar[int]
    team: _common_pb2.Team
    block_event: _containers.RepeatedCompositeFieldContainer[BlockEvent]
    def __init__(self, team: _Optional[_Union[_common_pb2.Team, _Mapping]] = ..., block_event: _Optional[_Iterable[_Union[BlockEvent, _Mapping]]] = ...) -> None: ...

class BlockEvent(_message.Message):
    __slots__ = ("data_units_ingested", "quota", "blocked_from")
    DATA_UNITS_INGESTED_FIELD_NUMBER: _ClassVar[int]
    QUOTA_FIELD_NUMBER: _ClassVar[int]
    BLOCKED_FROM_FIELD_NUMBER: _ClassVar[int]
    data_units_ingested: _common_pb2.Unit
    quota: _common_pb2.Unit
    blocked_from: _timestamp_pb2.Timestamp
    def __init__(self, data_units_ingested: _Optional[_Union[_common_pb2.Unit, _Mapping]] = ..., quota: _Optional[_Union[_common_pb2.Unit, _Mapping]] = ..., blocked_from: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class TeamQuotaHistory(_message.Message):
    __slots__ = ("team", "quota_entry")
    TEAM_FIELD_NUMBER: _ClassVar[int]
    QUOTA_ENTRY_FIELD_NUMBER: _ClassVar[int]
    team: _common_pb2.Team
    quota_entry: _containers.RepeatedCompositeFieldContainer[QuotaHistoryEntry]
    def __init__(self, team: _Optional[_Union[_common_pb2.Team, _Mapping]] = ..., quota_entry: _Optional[_Iterable[_Union[QuotaHistoryEntry, _Mapping]]] = ...) -> None: ...

class QuotaHistoryEntry(_message.Message):
    __slots__ = ("quota", "timestamp", "retention")
    QUOTA_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    RETENTION_FIELD_NUMBER: _ClassVar[int]
    quota: _common_pb2.Unit
    timestamp: _timestamp_pb2.Timestamp
    retention: _common_pb2.Retention
    def __init__(self, quota: _Optional[_Union[_common_pb2.Unit, _Mapping]] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., retention: _Optional[_Union[_common_pb2.Retention, _Mapping]] = ...) -> None: ...

class TeamQuota(_message.Message):
    __slots__ = ("team", "quota", "retention")
    TEAM_FIELD_NUMBER: _ClassVar[int]
    QUOTA_FIELD_NUMBER: _ClassVar[int]
    RETENTION_FIELD_NUMBER: _ClassVar[int]
    team: _common_pb2.Team
    quota: _common_pb2.Unit
    retention: _common_pb2.Retention
    def __init__(self, team: _Optional[_Union[_common_pb2.Team, _Mapping]] = ..., quota: _Optional[_Union[_common_pb2.Unit, _Mapping]] = ..., retention: _Optional[_Union[_common_pb2.Retention, _Mapping]] = ...) -> None: ...

class DataUsageReport(_message.Message):
    __slots__ = ("team", "type", "tier", "date", "application", "subsystem", "severity", "usage_in_gb", "units")
    TEAM_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    TIER_FIELD_NUMBER: _ClassVar[int]
    DATE_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    USAGE_IN_GB_FIELD_NUMBER: _ClassVar[int]
    UNITS_FIELD_NUMBER: _ClassVar[int]
    team: _common_pb2.Team
    type: _common_pb2.DataType
    tier: _common_pb2.TcoTier
    date: _timestamp_pb2.Timestamp
    application: _wrappers_pb2.StringValue
    subsystem: _wrappers_pb2.StringValue
    severity: _common_pb2.Severity
    usage_in_gb: _wrappers_pb2.DoubleValue
    units: _common_pb2.Unit
    def __init__(self, team: _Optional[_Union[_common_pb2.Team, _Mapping]] = ..., type: _Optional[_Union[_common_pb2.DataType, str]] = ..., tier: _Optional[_Union[_common_pb2.TcoTier, str]] = ..., date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., application: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., severity: _Optional[_Union[_common_pb2.Severity, str]] = ..., usage_in_gb: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ..., units: _Optional[_Union[_common_pb2.Unit, _Mapping]] = ...) -> None: ...

class GenericDimension(_message.Message):
    __slots__ = ("key", "value")
    KEY_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    key: str
    value: str
    def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...

class Dimension(_message.Message):
    __slots__ = ("pillar", "generic_dimension", "tier", "severity", "priority")
    PILLAR_FIELD_NUMBER: _ClassVar[int]
    GENERIC_DIMENSION_FIELD_NUMBER: _ClassVar[int]
    TIER_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    PRIORITY_FIELD_NUMBER: _ClassVar[int]
    pillar: Pillar
    generic_dimension: GenericDimension
    tier: _common_pb2.TcoTier
    severity: _common_pb2.Severity
    priority: _common_pb2.Priority
    def __init__(self, pillar: _Optional[_Union[Pillar, str]] = ..., generic_dimension: _Optional[_Union[GenericDimension, _Mapping]] = ..., tier: _Optional[_Union[_common_pb2.TcoTier, str]] = ..., severity: _Optional[_Union[_common_pb2.Severity, str]] = ..., priority: _Optional[_Union[_common_pb2.Priority, str]] = ...) -> None: ...

class DataUsage(_message.Message):
    __slots__ = ("usage",)
    class Usage(_message.Message):
        __slots__ = ("timestamp", "size")
        TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
        SIZE_FIELD_NUMBER: _ClassVar[int]
        timestamp: _timestamp_pb2.Timestamp
        size: _wrappers_pb2.FloatValue
        def __init__(self, timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., size: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ...) -> None: ...
    USAGE_FIELD_NUMBER: _ClassVar[int]
    usage: _containers.RepeatedCompositeFieldContainer[DataUsage.Usage]
    def __init__(self, usage: _Optional[_Iterable[_Union[DataUsage.Usage, _Mapping]]] = ...) -> None: ...
