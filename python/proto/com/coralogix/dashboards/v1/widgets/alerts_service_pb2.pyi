from com.coralogix.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.dashboards.v1.common import search_filter_pb2 as _search_filter_pb2
from com.coralogix.dashboards.v1.common import time_frame_pb2 as _time_frame_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AlertGroupSeverity(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ALERT_GROUP_SEVERITY_UNSPECIFIED: _ClassVar[AlertGroupSeverity]
    ALERT_GROUP_SEVERITY_INFO: _ClassVar[AlertGroupSeverity]
    ALERT_GROUP_SEVERITY_WARNING: _ClassVar[AlertGroupSeverity]
    ALERT_GROUP_SEVERITY_CRITICAL: _ClassVar[AlertGroupSeverity]
    ALERT_GROUP_SEVERITY_ERROR: _ClassVar[AlertGroupSeverity]
ALERT_GROUP_SEVERITY_UNSPECIFIED: AlertGroupSeverity
ALERT_GROUP_SEVERITY_INFO: AlertGroupSeverity
ALERT_GROUP_SEVERITY_WARNING: AlertGroupSeverity
ALERT_GROUP_SEVERITY_CRITICAL: AlertGroupSeverity
ALERT_GROUP_SEVERITY_ERROR: AlertGroupSeverity

class SearchAlertsGroupedBySeverityRequest(_message.Message):
    __slots__ = ("filters", "time_frame")
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    filters: _containers.RepeatedCompositeFieldContainer[_search_filter_pb2.SearchFilter]
    time_frame: _time_frame_pb2.TimeFrame
    def __init__(self, filters: _Optional[_Iterable[_Union[_search_filter_pb2.SearchFilter, _Mapping]]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ...) -> None: ...

class SearchAlertsGroupedBySeverityResponse(_message.Message):
    __slots__ = ("alerts",)
    ALERTS_FIELD_NUMBER: _ClassVar[int]
    alerts: _containers.RepeatedCompositeFieldContainer[AlertGroup]
    def __init__(self, alerts: _Optional[_Iterable[_Union[AlertGroup, _Mapping]]] = ...) -> None: ...

class AlertGroup(_message.Message):
    __slots__ = ("severity", "amount")
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    severity: AlertGroupSeverity
    amount: _wrappers_pb2.Int32Value
    def __init__(self, severity: _Optional[_Union[AlertGroupSeverity, str]] = ..., amount: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
