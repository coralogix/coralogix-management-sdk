from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.alerts.v1 import alert_severity_pb2 as _alert_severity_pb2
from com.coralogix.alerts.v1 import alert_meta_label_pb2 as _alert_meta_label_pb2
from com.coralogix.alerts.v1 import alert_condition_pb2 as _alert_condition_pb2
from com.coralogix.alerts.v1 import alert_notifications_pb2 as _alert_notifications_pb2
from com.coralogix.alerts.v1 import alert_filters_pb2 as _alert_filters_pb2
from com.coralogix.alerts.v1 import alert_active_when_pb2 as _alert_active_when_pb2
from com.coralogix.alerts.v1 import date_time_pb2 as _date_time_pb2
from com.coralogix.alerts.v1 import tracing_alert_pb2 as _tracing_alert_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Alert(_message.Message):
    __slots__ = ("id", "name", "description", "is_active", "severity", "expiration", "condition", "notifications", "filters", "notify_every", "active_when", "notification_payload_filters", "meta_labels", "meta_labels_strings", "tracing_alert", "unique_identifier")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    IS_ACTIVE_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    EXPIRATION_FIELD_NUMBER: _ClassVar[int]
    CONDITION_FIELD_NUMBER: _ClassVar[int]
    NOTIFICATIONS_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    NOTIFY_EVERY_FIELD_NUMBER: _ClassVar[int]
    ACTIVE_WHEN_FIELD_NUMBER: _ClassVar[int]
    NOTIFICATION_PAYLOAD_FILTERS_FIELD_NUMBER: _ClassVar[int]
    META_LABELS_FIELD_NUMBER: _ClassVar[int]
    META_LABELS_STRINGS_FIELD_NUMBER: _ClassVar[int]
    TRACING_ALERT_FIELD_NUMBER: _ClassVar[int]
    UNIQUE_IDENTIFIER_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    is_active: _wrappers_pb2.BoolValue
    severity: _alert_severity_pb2.AlertSeverity
    expiration: _date_time_pb2.Date
    condition: _alert_condition_pb2.AlertCondition
    notifications: _alert_notifications_pb2.AlertNotifications
    filters: _alert_filters_pb2.AlertFilters
    notify_every: _wrappers_pb2.DoubleValue
    active_when: _alert_active_when_pb2.AlertActiveWhen
    notification_payload_filters: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    meta_labels: _containers.RepeatedCompositeFieldContainer[_alert_meta_label_pb2.MetaLabel]
    meta_labels_strings: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    tracing_alert: _tracing_alert_pb2.TracingAlert
    unique_identifier: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., is_active: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., severity: _Optional[_Union[_alert_severity_pb2.AlertSeverity, str]] = ..., expiration: _Optional[_Union[_date_time_pb2.Date, _Mapping]] = ..., condition: _Optional[_Union[_alert_condition_pb2.AlertCondition, _Mapping]] = ..., notifications: _Optional[_Union[_alert_notifications_pb2.AlertNotifications, _Mapping]] = ..., filters: _Optional[_Union[_alert_filters_pb2.AlertFilters, _Mapping]] = ..., notify_every: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ..., active_when: _Optional[_Union[_alert_active_when_pb2.AlertActiveWhen, _Mapping]] = ..., notification_payload_filters: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., meta_labels: _Optional[_Iterable[_Union[_alert_meta_label_pb2.MetaLabel, _Mapping]]] = ..., meta_labels_strings: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., tracing_alert: _Optional[_Union[_tracing_alert_pb2.TracingAlert, _Mapping]] = ..., unique_identifier: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
