from com.coralogix.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.dashboards.v1.common import search_filter_pb2 as _search_filter_pb2
from com.coralogix.dashboards.v1.common import time_frame_pb2 as _time_frame_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class InsightSeverity(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    INSIGHT_SEVERITY_UNSPECIFIED: _ClassVar[InsightSeverity]
    INSIGHT_SEVERITY_INFO: _ClassVar[InsightSeverity]
    INSIGHT_SEVERITY_WARNING: _ClassVar[InsightSeverity]
    INSIGHT_SEVERITY_CRITICAL: _ClassVar[InsightSeverity]
    INSIGHT_SEVERITY_ERROR: _ClassVar[InsightSeverity]

class InsightType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    INSIGHT_TYPE_UNSPECIFIED: _ClassVar[InsightType]
    INSIGHT_TYPE_USER_ALERT: _ClassVar[InsightType]
    INSIGHT_TYPE_ANOMALY: _ClassVar[InsightType]
    INSIGHT_TYPE_RATIO: _ClassVar[InsightType]
    INSIGHT_TYPE_NEW_VALUE: _ClassVar[InsightType]
    INSIGHT_TYPE_UNIQUE_COUNT: _ClassVar[InsightType]
    INSIGHT_TYPE_TIME_RELATIVE: _ClassVar[InsightType]
    INSIGHT_TYPE_METRIC: _ClassVar[InsightType]
    INSIGHT_TYPE_FLOW: _ClassVar[InsightType]
INSIGHT_SEVERITY_UNSPECIFIED: InsightSeverity
INSIGHT_SEVERITY_INFO: InsightSeverity
INSIGHT_SEVERITY_WARNING: InsightSeverity
INSIGHT_SEVERITY_CRITICAL: InsightSeverity
INSIGHT_SEVERITY_ERROR: InsightSeverity
INSIGHT_TYPE_UNSPECIFIED: InsightType
INSIGHT_TYPE_USER_ALERT: InsightType
INSIGHT_TYPE_ANOMALY: InsightType
INSIGHT_TYPE_RATIO: InsightType
INSIGHT_TYPE_NEW_VALUE: InsightType
INSIGHT_TYPE_UNIQUE_COUNT: InsightType
INSIGHT_TYPE_TIME_RELATIVE: InsightType
INSIGHT_TYPE_METRIC: InsightType
INSIGHT_TYPE_FLOW: InsightType

class SearchInsightsRequest(_message.Message):
    __slots__ = ("filters", "time_frame")
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    filters: _containers.RepeatedCompositeFieldContainer[_search_filter_pb2.SearchFilter]
    time_frame: _time_frame_pb2.TimeFrame
    def __init__(self, filters: _Optional[_Iterable[_Union[_search_filter_pb2.SearchFilter, _Mapping]]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ...) -> None: ...

class SearchInsightsResponse(_message.Message):
    __slots__ = ("insights",)
    INSIGHTS_FIELD_NUMBER: _ClassVar[int]
    insights: _containers.RepeatedCompositeFieldContainer[Insight]
    def __init__(self, insights: _Optional[_Iterable[_Union[Insight, _Mapping]]] = ...) -> None: ...

class Insight(_message.Message):
    __slots__ = ("id", "alert_id", "name", "severity", "application", "subsystem", "repetitions", "type", "last_occurrence", "alert_snoozed")
    ID_FIELD_NUMBER: _ClassVar[int]
    ALERT_ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
    REPETITIONS_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    LAST_OCCURRENCE_FIELD_NUMBER: _ClassVar[int]
    ALERT_SNOOZED_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    alert_id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    severity: InsightSeverity
    application: _wrappers_pb2.StringValue
    subsystem: _wrappers_pb2.StringValue
    repetitions: _wrappers_pb2.Int32Value
    type: InsightType
    last_occurrence: _timestamp_pb2.Timestamp
    alert_snoozed: AlertSnoozed
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., alert_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., severity: _Optional[_Union[InsightSeverity, str]] = ..., application: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., repetitions: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., type: _Optional[_Union[InsightType, str]] = ..., last_occurrence: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., alert_snoozed: _Optional[_Union[AlertSnoozed, _Mapping]] = ...) -> None: ...

class AlertSnoozed(_message.Message):
    __slots__ = ("user_id", "until")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    UNTIL_FIELD_NUMBER: _ClassVar[int]
    user_id: _wrappers_pb2.StringValue
    until: _timestamp_pb2.Timestamp
    def __init__(self, user_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., until: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
