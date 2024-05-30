from com.coralogix.datausage.v1 import common_pb2 as _common_pb2
from com.coralogix.datausage.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.datausage.v1 import data_usage_pb2 as _data_usage_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetDataUsageRequest(_message.Message):
    __slots__ = ("team_ids", "date_range", "resolution", "dimensions")
    TEAM_IDS_FIELD_NUMBER: _ClassVar[int]
    DATE_RANGE_FIELD_NUMBER: _ClassVar[int]
    RESOLUTION_FIELD_NUMBER: _ClassVar[int]
    DIMENSIONS_FIELD_NUMBER: _ClassVar[int]
    team_ids: _containers.RepeatedCompositeFieldContainer[_common_pb2.Team]
    date_range: _common_pb2.DateRange
    resolution: _duration_pb2.Duration
    dimensions: _containers.RepeatedCompositeFieldContainer[_data_usage_pb2.Dimension]
    def __init__(self, team_ids: _Optional[_Iterable[_Union[_common_pb2.Team, _Mapping]]] = ..., date_range: _Optional[_Union[_common_pb2.DateRange, _Mapping]] = ..., resolution: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., dimensions: _Optional[_Iterable[_Union[_data_usage_pb2.Dimension, _Mapping]]] = ...) -> None: ...

class GetDataUsageResponse(_message.Message):
    __slots__ = ("data_usage",)
    class DataUsageEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: int
        value: _data_usage_pb2.DataUsage
        def __init__(self, key: _Optional[int] = ..., value: _Optional[_Union[_data_usage_pb2.DataUsage, _Mapping]] = ...) -> None: ...
    DATA_USAGE_FIELD_NUMBER: _ClassVar[int]
    data_usage: _containers.MessageMap[int, _data_usage_pb2.DataUsage]
    def __init__(self, data_usage: _Optional[_Mapping[int, _data_usage_pb2.DataUsage]] = ...) -> None: ...

class GetTeamsBlocksRequest(_message.Message):
    __slots__ = ("param",)
    PARAM_FIELD_NUMBER: _ClassVar[int]
    param: _data_usage_pb2.TeamsAndTimeRange
    def __init__(self, param: _Optional[_Union[_data_usage_pb2.TeamsAndTimeRange, _Mapping]] = ...) -> None: ...

class GetTeamsBlocksResponse(_message.Message):
    __slots__ = ("events",)
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    events: _containers.RepeatedCompositeFieldContainer[_data_usage_pb2.TeamBlockEvents]
    def __init__(self, events: _Optional[_Iterable[_Union[_data_usage_pb2.TeamBlockEvents, _Mapping]]] = ...) -> None: ...

class GetTeamsQuotaHistoryRequest(_message.Message):
    __slots__ = ("param",)
    PARAM_FIELD_NUMBER: _ClassVar[int]
    param: _data_usage_pb2.TeamsAndTimeRange
    def __init__(self, param: _Optional[_Union[_data_usage_pb2.TeamsAndTimeRange, _Mapping]] = ...) -> None: ...

class GetTeamsQuotaHistoryResponse(_message.Message):
    __slots__ = ("history_entries",)
    HISTORY_ENTRIES_FIELD_NUMBER: _ClassVar[int]
    history_entries: _containers.RepeatedCompositeFieldContainer[_data_usage_pb2.TeamQuotaHistory]
    def __init__(self, history_entries: _Optional[_Iterable[_Union[_data_usage_pb2.TeamQuotaHistory, _Mapping]]] = ...) -> None: ...

class GetTeamsDailyUsageRequest(_message.Message):
    __slots__ = ("param",)
    PARAM_FIELD_NUMBER: _ClassVar[int]
    param: _data_usage_pb2.TeamsAndTimeRange
    def __init__(self, param: _Optional[_Union[_data_usage_pb2.TeamsAndTimeRange, _Mapping]] = ...) -> None: ...

class GetTeamsDailyUsageResponse(_message.Message):
    __slots__ = ("teams_usage",)
    TEAMS_USAGE_FIELD_NUMBER: _ClassVar[int]
    teams_usage: _containers.RepeatedCompositeFieldContainer[_data_usage_pb2.TeamsDailyDataUsage]
    def __init__(self, teams_usage: _Optional[_Iterable[_Union[_data_usage_pb2.TeamsDailyDataUsage, _Mapping]]] = ...) -> None: ...

class GetTeamsQuotaRequest(_message.Message):
    __slots__ = ("param",)
    PARAM_FIELD_NUMBER: _ClassVar[int]
    param: _data_usage_pb2.TeamsAndTime
    def __init__(self, param: _Optional[_Union[_data_usage_pb2.TeamsAndTime, _Mapping]] = ...) -> None: ...

class GetTeamsQuotaResponse(_message.Message):
    __slots__ = ("teams_quota",)
    TEAMS_QUOTA_FIELD_NUMBER: _ClassVar[int]
    teams_quota: _containers.RepeatedCompositeFieldContainer[_data_usage_pb2.TeamQuota]
    def __init__(self, teams_quota: _Optional[_Iterable[_Union[_data_usage_pb2.TeamQuota, _Mapping]]] = ...) -> None: ...

class GetTeamsDetailedUsageRequest(_message.Message):
    __slots__ = ("param",)
    PARAM_FIELD_NUMBER: _ClassVar[int]
    param: _data_usage_pb2.TeamsAndTimeRange
    def __init__(self, param: _Optional[_Union[_data_usage_pb2.TeamsAndTimeRange, _Mapping]] = ...) -> None: ...

class GetTeamsDetailedUsageResponse(_message.Message):
    __slots__ = ("reports",)
    REPORTS_FIELD_NUMBER: _ClassVar[int]
    reports: _containers.RepeatedCompositeFieldContainer[_data_usage_pb2.DataUsageReport]
    def __init__(self, reports: _Optional[_Iterable[_Union[_data_usage_pb2.DataUsageReport, _Mapping]]] = ...) -> None: ...

class GetDetailedDataUsageRequest(_message.Message):
    __slots__ = ("team_id", "range", "date_range", "resolution", "dimensions")
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    RANGE_FIELD_NUMBER: _ClassVar[int]
    DATE_RANGE_FIELD_NUMBER: _ClassVar[int]
    RESOLUTION_FIELD_NUMBER: _ClassVar[int]
    DIMENSIONS_FIELD_NUMBER: _ClassVar[int]
    team_id: _common_pb2.Team
    range: _common_pb2.Range
    date_range: _common_pb2.DateRange
    resolution: _common_pb2.Resolution
    dimensions: _containers.RepeatedCompositeFieldContainer[_data_usage_pb2.Dimension]
    def __init__(self, team_id: _Optional[_Union[_common_pb2.Team, _Mapping]] = ..., range: _Optional[_Union[_common_pb2.Range, str]] = ..., date_range: _Optional[_Union[_common_pb2.DateRange, _Mapping]] = ..., resolution: _Optional[_Union[_common_pb2.Resolution, str]] = ..., dimensions: _Optional[_Iterable[_Union[_data_usage_pb2.Dimension, _Mapping]]] = ...) -> None: ...

class GetDetailedDataUsageResponse(_message.Message):
    __slots__ = ("timestamp", "size", "quota", "dimensions")
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    QUOTA_FIELD_NUMBER: _ClassVar[int]
    DIMENSIONS_FIELD_NUMBER: _ClassVar[int]
    timestamp: _timestamp_pb2.Timestamp
    size: _wrappers_pb2.FloatValue
    quota: _wrappers_pb2.FloatValue
    dimensions: _containers.RepeatedCompositeFieldContainer[_data_usage_pb2.Dimension]
    def __init__(self, timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., size: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., quota: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ..., dimensions: _Optional[_Iterable[_Union[_data_usage_pb2.Dimension, _Mapping]]] = ...) -> None: ...
