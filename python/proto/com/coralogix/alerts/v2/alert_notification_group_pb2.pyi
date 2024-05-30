from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class NotifyOn(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    TRIGGERED_ONLY: _ClassVar[NotifyOn]
    TRIGGERED_AND_RESOLVED: _ClassVar[NotifyOn]
TRIGGERED_ONLY: NotifyOn
TRIGGERED_AND_RESOLVED: NotifyOn

class ShowInInsight(_message.Message):
    __slots__ = ("retriggering_period_seconds", "notify_on")
    RETRIGGERING_PERIOD_SECONDS_FIELD_NUMBER: _ClassVar[int]
    NOTIFY_ON_FIELD_NUMBER: _ClassVar[int]
    retriggering_period_seconds: _wrappers_pb2.UInt32Value
    notify_on: NotifyOn
    def __init__(self, retriggering_period_seconds: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., notify_on: _Optional[_Union[NotifyOn, str]] = ...) -> None: ...

class AlertNotificationGroups(_message.Message):
    __slots__ = ("group_by_fields", "notifications")
    GROUP_BY_FIELDS_FIELD_NUMBER: _ClassVar[int]
    NOTIFICATIONS_FIELD_NUMBER: _ClassVar[int]
    group_by_fields: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    notifications: _containers.RepeatedCompositeFieldContainer[AlertNotification]
    def __init__(self, group_by_fields: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., notifications: _Optional[_Iterable[_Union[AlertNotification, _Mapping]]] = ...) -> None: ...

class AlertNotification(_message.Message):
    __slots__ = ("retriggering_period_seconds", "notify_on", "integration_id", "recipients")
    RETRIGGERING_PERIOD_SECONDS_FIELD_NUMBER: _ClassVar[int]
    NOTIFY_ON_FIELD_NUMBER: _ClassVar[int]
    INTEGRATION_ID_FIELD_NUMBER: _ClassVar[int]
    RECIPIENTS_FIELD_NUMBER: _ClassVar[int]
    retriggering_period_seconds: _wrappers_pb2.UInt32Value
    notify_on: NotifyOn
    integration_id: _wrappers_pb2.UInt32Value
    recipients: Recipients
    def __init__(self, retriggering_period_seconds: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., notify_on: _Optional[_Union[NotifyOn, str]] = ..., integration_id: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., recipients: _Optional[_Union[Recipients, _Mapping]] = ...) -> None: ...

class Recipients(_message.Message):
    __slots__ = ("emails",)
    EMAILS_FIELD_NUMBER: _ClassVar[int]
    emails: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, emails: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...
