from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.alerts.v2 import alert_notification_group_pb2 as _alert_notification_group_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AlertIncidentSettings(_message.Message):
    __slots__ = ("retriggering_period_seconds", "notify_on", "use_as_notification_settings")
    RETRIGGERING_PERIOD_SECONDS_FIELD_NUMBER: _ClassVar[int]
    NOTIFY_ON_FIELD_NUMBER: _ClassVar[int]
    USE_AS_NOTIFICATION_SETTINGS_FIELD_NUMBER: _ClassVar[int]
    retriggering_period_seconds: _wrappers_pb2.UInt32Value
    notify_on: _alert_notification_group_pb2.NotifyOn
    use_as_notification_settings: _wrappers_pb2.BoolValue
    def __init__(self, retriggering_period_seconds: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., notify_on: _Optional[_Union[_alert_notification_group_pb2.NotifyOn, str]] = ..., use_as_notification_settings: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
