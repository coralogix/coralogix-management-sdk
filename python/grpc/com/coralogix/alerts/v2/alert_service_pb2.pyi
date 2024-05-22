from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.alerts.v2 import alert_pb2 as _alert_pb2
from com.coralogix.alerts.v1 import alert_event_pb2 as _alert_event_pb2
from com.coralogix.alerts.v1 import alert_meta_label_pb2 as _alert_meta_label_pb2
from com.coralogix.alerts.v1 import alert_severity_pb2 as _alert_severity_pb2
from com.coralogix.alerts.v2 import alert_condition_pb2 as _alert_condition_pb2
from com.coralogix.alerts.v1 import alert_filters_pb2 as _alert_filters_pb2
from com.coralogix.alerts.v1 import alert_service_pb2 as _alert_service_pb2
from com.coralogix.alerts.v1 import alert_active_when_pb2 as _alert_active_when_pb2
from com.coralogix.alerts.v1 import date_time_pb2 as _date_time_pb2
from com.coralogix.alerts.v1 import fields_to_update_pb2 as _fields_to_update_pb2
from com.coralogix.alerts.v2 import alert_notification_group_pb2 as _alert_notification_group_pb2
from com.coralogix.alerts.v2 import alert_incident_settings_pb2 as _alert_incident_settings_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from com.coralogix.alerts.v1 import tracing_alert_pb2 as _tracing_alert_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetAlertRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetAlertResponse(_message.Message):
    __slots__ = ("alert",)
    ALERT_FIELD_NUMBER: _ClassVar[int]
    alert: _alert_pb2.Alert
    def __init__(self, alert: _Optional[_Union[_alert_pb2.Alert, _Mapping]] = ...) -> None: ...

class GetAlertByUniqueIdRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetAlertByUniqueIdResponse(_message.Message):
    __slots__ = ("alert",)
    ALERT_FIELD_NUMBER: _ClassVar[int]
    alert: _alert_pb2.Alert
    def __init__(self, alert: _Optional[_Union[_alert_pb2.Alert, _Mapping]] = ...) -> None: ...

class CreateAlertRequest(_message.Message):
    __slots__ = ("name", "description", "is_active", "severity", "expiration", "condition", "show_in_insight", "notification_groups", "incident_settings", "filters", "active_when", "notification_payload_filters", "meta_labels", "meta_labels_strings", "tracing_alert")
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    IS_ACTIVE_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    EXPIRATION_FIELD_NUMBER: _ClassVar[int]
    CONDITION_FIELD_NUMBER: _ClassVar[int]
    SHOW_IN_INSIGHT_FIELD_NUMBER: _ClassVar[int]
    NOTIFICATION_GROUPS_FIELD_NUMBER: _ClassVar[int]
    INCIDENT_SETTINGS_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    ACTIVE_WHEN_FIELD_NUMBER: _ClassVar[int]
    NOTIFICATION_PAYLOAD_FILTERS_FIELD_NUMBER: _ClassVar[int]
    META_LABELS_FIELD_NUMBER: _ClassVar[int]
    META_LABELS_STRINGS_FIELD_NUMBER: _ClassVar[int]
    TRACING_ALERT_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    is_active: _wrappers_pb2.BoolValue
    severity: _alert_severity_pb2.AlertSeverity
    expiration: _date_time_pb2.Date
    condition: _alert_condition_pb2.AlertCondition
    show_in_insight: _alert_notification_group_pb2.ShowInInsight
    notification_groups: _containers.RepeatedCompositeFieldContainer[_alert_notification_group_pb2.AlertNotificationGroups]
    incident_settings: _alert_incident_settings_pb2.AlertIncidentSettings
    filters: _alert_filters_pb2.AlertFilters
    active_when: _alert_active_when_pb2.AlertActiveWhen
    notification_payload_filters: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    meta_labels: _containers.RepeatedCompositeFieldContainer[_alert_meta_label_pb2.MetaLabel]
    meta_labels_strings: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    tracing_alert: _tracing_alert_pb2.TracingAlert
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., is_active: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., severity: _Optional[_Union[_alert_severity_pb2.AlertSeverity, str]] = ..., expiration: _Optional[_Union[_date_time_pb2.Date, _Mapping]] = ..., condition: _Optional[_Union[_alert_condition_pb2.AlertCondition, _Mapping]] = ..., show_in_insight: _Optional[_Union[_alert_notification_group_pb2.ShowInInsight, _Mapping]] = ..., notification_groups: _Optional[_Iterable[_Union[_alert_notification_group_pb2.AlertNotificationGroups, _Mapping]]] = ..., incident_settings: _Optional[_Union[_alert_incident_settings_pb2.AlertIncidentSettings, _Mapping]] = ..., filters: _Optional[_Union[_alert_filters_pb2.AlertFilters, _Mapping]] = ..., active_when: _Optional[_Union[_alert_active_when_pb2.AlertActiveWhen, _Mapping]] = ..., notification_payload_filters: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., meta_labels: _Optional[_Iterable[_Union[_alert_meta_label_pb2.MetaLabel, _Mapping]]] = ..., meta_labels_strings: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., tracing_alert: _Optional[_Union[_tracing_alert_pb2.TracingAlert, _Mapping]] = ...) -> None: ...

class CreateAlertResponse(_message.Message):
    __slots__ = ("alert",)
    ALERT_FIELD_NUMBER: _ClassVar[int]
    alert: _alert_pb2.Alert
    def __init__(self, alert: _Optional[_Union[_alert_pb2.Alert, _Mapping]] = ...) -> None: ...

class GetAlertModelMappingRequest(_message.Message):
    __slots__ = ("name", "description", "is_active", "severity", "expiration", "condition", "show_in_insight", "notification_groups", "incident_settings", "filters", "active_when", "notification_payload_filters", "meta_labels", "meta_labels_strings", "tracing_alert")
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    IS_ACTIVE_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    EXPIRATION_FIELD_NUMBER: _ClassVar[int]
    CONDITION_FIELD_NUMBER: _ClassVar[int]
    SHOW_IN_INSIGHT_FIELD_NUMBER: _ClassVar[int]
    NOTIFICATION_GROUPS_FIELD_NUMBER: _ClassVar[int]
    INCIDENT_SETTINGS_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    ACTIVE_WHEN_FIELD_NUMBER: _ClassVar[int]
    NOTIFICATION_PAYLOAD_FILTERS_FIELD_NUMBER: _ClassVar[int]
    META_LABELS_FIELD_NUMBER: _ClassVar[int]
    META_LABELS_STRINGS_FIELD_NUMBER: _ClassVar[int]
    TRACING_ALERT_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    is_active: _wrappers_pb2.BoolValue
    severity: _alert_severity_pb2.AlertSeverity
    expiration: _date_time_pb2.Date
    condition: _alert_condition_pb2.AlertCondition
    show_in_insight: _alert_notification_group_pb2.ShowInInsight
    notification_groups: _containers.RepeatedCompositeFieldContainer[_alert_notification_group_pb2.AlertNotificationGroups]
    incident_settings: _alert_incident_settings_pb2.AlertIncidentSettings
    filters: _alert_filters_pb2.AlertFilters
    active_when: _alert_active_when_pb2.AlertActiveWhen
    notification_payload_filters: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    meta_labels: _containers.RepeatedCompositeFieldContainer[_alert_meta_label_pb2.MetaLabel]
    meta_labels_strings: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    tracing_alert: _tracing_alert_pb2.TracingAlert
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., is_active: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., severity: _Optional[_Union[_alert_severity_pb2.AlertSeverity, str]] = ..., expiration: _Optional[_Union[_date_time_pb2.Date, _Mapping]] = ..., condition: _Optional[_Union[_alert_condition_pb2.AlertCondition, _Mapping]] = ..., show_in_insight: _Optional[_Union[_alert_notification_group_pb2.ShowInInsight, _Mapping]] = ..., notification_groups: _Optional[_Iterable[_Union[_alert_notification_group_pb2.AlertNotificationGroups, _Mapping]]] = ..., incident_settings: _Optional[_Union[_alert_incident_settings_pb2.AlertIncidentSettings, _Mapping]] = ..., filters: _Optional[_Union[_alert_filters_pb2.AlertFilters, _Mapping]] = ..., active_when: _Optional[_Union[_alert_active_when_pb2.AlertActiveWhen, _Mapping]] = ..., notification_payload_filters: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., meta_labels: _Optional[_Iterable[_Union[_alert_meta_label_pb2.MetaLabel, _Mapping]]] = ..., meta_labels_strings: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., tracing_alert: _Optional[_Union[_tracing_alert_pb2.TracingAlert, _Mapping]] = ...) -> None: ...

class GetAlertModelMappingResponse(_message.Message):
    __slots__ = ("alert_definition",)
    ALERT_DEFINITION_FIELD_NUMBER: _ClassVar[int]
    alert_definition: _struct_pb2.Struct
    def __init__(self, alert_definition: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...

class UpdateAlertRequest(_message.Message):
    __slots__ = ("alert", "fields_to_update")
    ALERT_FIELD_NUMBER: _ClassVar[int]
    FIELDS_TO_UPDATE_FIELD_NUMBER: _ClassVar[int]
    alert: _alert_pb2.Alert
    fields_to_update: _containers.RepeatedScalarFieldContainer[_fields_to_update_pb2.FieldsToUpdate]
    def __init__(self, alert: _Optional[_Union[_alert_pb2.Alert, _Mapping]] = ..., fields_to_update: _Optional[_Iterable[_Union[_fields_to_update_pb2.FieldsToUpdate, str]]] = ...) -> None: ...

class UpdateAlertResponse(_message.Message):
    __slots__ = ("alert",)
    ALERT_FIELD_NUMBER: _ClassVar[int]
    alert: _alert_pb2.Alert
    def __init__(self, alert: _Optional[_Union[_alert_pb2.Alert, _Mapping]] = ...) -> None: ...

class UpdateAlertByUniqueIdRequest(_message.Message):
    __slots__ = ("alert", "fields_to_update")
    ALERT_FIELD_NUMBER: _ClassVar[int]
    FIELDS_TO_UPDATE_FIELD_NUMBER: _ClassVar[int]
    alert: _alert_pb2.Alert
    fields_to_update: _containers.RepeatedScalarFieldContainer[_fields_to_update_pb2.FieldsToUpdate]
    def __init__(self, alert: _Optional[_Union[_alert_pb2.Alert, _Mapping]] = ..., fields_to_update: _Optional[_Iterable[_Union[_fields_to_update_pb2.FieldsToUpdate, str]]] = ...) -> None: ...

class UpdateAlertByUniqueIdResponse(_message.Message):
    __slots__ = ("alert",)
    ALERT_FIELD_NUMBER: _ClassVar[int]
    alert: _alert_pb2.Alert
    def __init__(self, alert: _Optional[_Union[_alert_pb2.Alert, _Mapping]] = ...) -> None: ...

class GetAlertsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetAlertsResponse(_message.Message):
    __slots__ = ("alerts",)
    ALERTS_FIELD_NUMBER: _ClassVar[int]
    alerts: _containers.RepeatedCompositeFieldContainer[_alert_pb2.Alert]
    def __init__(self, alerts: _Optional[_Iterable[_Union[_alert_pb2.Alert, _Mapping]]] = ...) -> None: ...

class ValidateAlertRequest(_message.Message):
    __slots__ = ("alert",)
    ALERT_FIELD_NUMBER: _ClassVar[int]
    alert: _alert_pb2.Alert
    def __init__(self, alert: _Optional[_Union[_alert_pb2.Alert, _Mapping]] = ...) -> None: ...

class ValidateAlertResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DeleteAlertRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class DeleteAlertResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DeleteAlertByUniqueIdRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class DeleteAlertByUniqueIdResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetAlertEventsRequest(_message.Message):
    __slots__ = ("to",)
    FROM_FIELD_NUMBER: _ClassVar[int]
    TO_FIELD_NUMBER: _ClassVar[int]
    to: _timestamp_pb2.Timestamp
    def __init__(self, to: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., **kwargs) -> None: ...

class GetAlertEventsResponse(_message.Message):
    __slots__ = ("events",)
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    events: _containers.RepeatedCompositeFieldContainer[_alert_event_pb2.AlertEvent]
    def __init__(self, events: _Optional[_Iterable[_Union[_alert_event_pb2.AlertEvent, _Mapping]]] = ...) -> None: ...

class GetAlertEventsCountBySeverityRequest(_message.Message):
    __slots__ = ("to", "applications", "subsystems")
    FROM_FIELD_NUMBER: _ClassVar[int]
    TO_FIELD_NUMBER: _ClassVar[int]
    APPLICATIONS_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEMS_FIELD_NUMBER: _ClassVar[int]
    to: _timestamp_pb2.Timestamp
    applications: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    subsystems: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, to: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., applications: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystems: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., **kwargs) -> None: ...

class GetAlertEventsCountBySeverityResponse(_message.Message):
    __slots__ = ("event_counts",)
    EVENT_COUNTS_FIELD_NUMBER: _ClassVar[int]
    event_counts: _containers.RepeatedCompositeFieldContainer[AlertEventsCountBySeverity]
    def __init__(self, event_counts: _Optional[_Iterable[_Union[AlertEventsCountBySeverity, _Mapping]]] = ...) -> None: ...

class AlertEventsCountBySeverity(_message.Message):
    __slots__ = ("severity", "type", "sub_type", "count")
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    SUB_TYPE_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    severity: _alert_severity_pb2.AlertSeverity
    type: _alert_service_pb2.AlertEventCountType
    sub_type: _alert_event_pb2.AlertEvent.AlertEventSubType
    count: _wrappers_pb2.Int32Value
    def __init__(self, severity: _Optional[_Union[_alert_severity_pb2.AlertSeverity, str]] = ..., type: _Optional[_Union[_alert_service_pb2.AlertEventCountType, str]] = ..., sub_type: _Optional[_Union[_alert_event_pb2.AlertEvent.AlertEventSubType, str]] = ..., count: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class GetCompanyAlertsLimitRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetCompanyAlertsLimitResponse(_message.Message):
    __slots__ = ("company_id", "limit", "used", "dynamic_alert_limit", "dynamic_alert_used")
    COMPANY_ID_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    USED_FIELD_NUMBER: _ClassVar[int]
    DYNAMIC_ALERT_LIMIT_FIELD_NUMBER: _ClassVar[int]
    DYNAMIC_ALERT_USED_FIELD_NUMBER: _ClassVar[int]
    company_id: _wrappers_pb2.StringValue
    limit: _wrappers_pb2.Int32Value
    used: _wrappers_pb2.Int32Value
    dynamic_alert_limit: _wrappers_pb2.Int32Value
    dynamic_alert_used: _wrappers_pb2.Int32Value
    def __init__(self, company_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., used: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., dynamic_alert_limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., dynamic_alert_used: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
