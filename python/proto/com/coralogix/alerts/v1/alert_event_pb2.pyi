from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.alerts.v1 import alert_severity_pb2 as _alert_severity_pb2
from com.coralogix.alerts.v1 import alert_snoozed_pb2 as _alert_snoozed_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AlertEvent(_message.Message):
    __slots__ = ("id", "type", "alert_id", "company_id", "severity", "subsystem", "application", "occurred_on", "name", "snoozed", "sub_type")
    class AlertEventType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        ALERT_EVENT_TYPE_UNSPECIFIED: _ClassVar[AlertEvent.AlertEventType]
        ALERT_EVENT_TYPE_USER_ALERT: _ClassVar[AlertEvent.AlertEventType]
        ALERT_EVENT_TYPE_ANOMALY: _ClassVar[AlertEvent.AlertEventType]
        ALERT_EVENT_TYPE_RATIO: _ClassVar[AlertEvent.AlertEventType]
        ALERT_EVENT_TYPE_NEW_VALUE: _ClassVar[AlertEvent.AlertEventType]
        ALERT_EVENT_TYPE_UNIQUE_COUNT: _ClassVar[AlertEvent.AlertEventType]
        ALERT_EVENT_TYPE_TIME_RELATIVE: _ClassVar[AlertEvent.AlertEventType]
        ALERT_EVENT_TYPE_METRIC: _ClassVar[AlertEvent.AlertEventType]
        ALERT_EVENT_TYPE_FLOW: _ClassVar[AlertEvent.AlertEventType]
    ALERT_EVENT_TYPE_UNSPECIFIED: AlertEvent.AlertEventType
    ALERT_EVENT_TYPE_USER_ALERT: AlertEvent.AlertEventType
    ALERT_EVENT_TYPE_ANOMALY: AlertEvent.AlertEventType
    ALERT_EVENT_TYPE_RATIO: AlertEvent.AlertEventType
    ALERT_EVENT_TYPE_NEW_VALUE: AlertEvent.AlertEventType
    ALERT_EVENT_TYPE_UNIQUE_COUNT: AlertEvent.AlertEventType
    ALERT_EVENT_TYPE_TIME_RELATIVE: AlertEvent.AlertEventType
    ALERT_EVENT_TYPE_METRIC: AlertEvent.AlertEventType
    ALERT_EVENT_TYPE_FLOW: AlertEvent.AlertEventType
    class AlertEventSubType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        ALERT_EVENT_SUB_TYPE_UNSPECIFIED: _ClassVar[AlertEvent.AlertEventSubType]
        ALERT_EVENT_SUB_TYPE_FLOW: _ClassVar[AlertEvent.AlertEventSubType]
        ALERT_EVENT_SUB_TYPE_VOLUME: _ClassVar[AlertEvent.AlertEventSubType]
    ALERT_EVENT_SUB_TYPE_UNSPECIFIED: AlertEvent.AlertEventSubType
    ALERT_EVENT_SUB_TYPE_FLOW: AlertEvent.AlertEventSubType
    ALERT_EVENT_SUB_TYPE_VOLUME: AlertEvent.AlertEventSubType
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    ALERT_ID_FIELD_NUMBER: _ClassVar[int]
    COMPANY_ID_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    OCCURRED_ON_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    SNOOZED_FIELD_NUMBER: _ClassVar[int]
    SUB_TYPE_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    type: AlertEvent.AlertEventType
    alert_id: _wrappers_pb2.StringValue
    company_id: _wrappers_pb2.StringValue
    severity: _alert_severity_pb2.AlertSeverity
    subsystem: _wrappers_pb2.StringValue
    application: _wrappers_pb2.StringValue
    occurred_on: _timestamp_pb2.Timestamp
    name: _wrappers_pb2.StringValue
    snoozed: _alert_snoozed_pb2.AlertSnoozed
    sub_type: AlertEvent.AlertEventSubType
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., type: _Optional[_Union[AlertEvent.AlertEventType, str]] = ..., alert_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., company_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., severity: _Optional[_Union[_alert_severity_pb2.AlertSeverity, str]] = ..., subsystem: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., application: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., occurred_on: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., snoozed: _Optional[_Union[_alert_snoozed_pb2.AlertSnoozed, _Mapping]] = ..., sub_type: _Optional[_Union[AlertEvent.AlertEventSubType, str]] = ...) -> None: ...
