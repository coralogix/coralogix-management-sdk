from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class WebhookType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    UNKNOWN: _ClassVar[WebhookType]
    GENERIC: _ClassVar[WebhookType]
    SLACK: _ClassVar[WebhookType]
    PAGERDUTY: _ClassVar[WebhookType]
    SEND_LOG: _ClassVar[WebhookType]
    EMAIL_GROUP: _ClassVar[WebhookType]
    MICROSOFT_TEAMS: _ClassVar[WebhookType]
    JIRA: _ClassVar[WebhookType]
    OPSGENIE: _ClassVar[WebhookType]
    DEMISTO: _ClassVar[WebhookType]
    AWS_EVENT_BRIDGE: _ClassVar[WebhookType]
    IBM_EVENT_NOTIFICATIONS: _ClassVar[WebhookType]
UNKNOWN: WebhookType
GENERIC: WebhookType
SLACK: WebhookType
PAGERDUTY: WebhookType
SEND_LOG: WebhookType
EMAIL_GROUP: WebhookType
MICROSOFT_TEAMS: WebhookType
JIRA: WebhookType
OPSGENIE: WebhookType
DEMISTO: WebhookType
AWS_EVENT_BRIDGE: WebhookType
IBM_EVENT_NOTIFICATIONS: WebhookType

class GenericWebhookConfig(_message.Message):
    __slots__ = ("uuid", "method", "headers", "payload")
    class MethodType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        UNKNOWN: _ClassVar[GenericWebhookConfig.MethodType]
        GET: _ClassVar[GenericWebhookConfig.MethodType]
        POST: _ClassVar[GenericWebhookConfig.MethodType]
        PUT: _ClassVar[GenericWebhookConfig.MethodType]
    UNKNOWN: GenericWebhookConfig.MethodType
    GET: GenericWebhookConfig.MethodType
    POST: GenericWebhookConfig.MethodType
    PUT: GenericWebhookConfig.MethodType
    class HeadersEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    UUID_FIELD_NUMBER: _ClassVar[int]
    METHOD_FIELD_NUMBER: _ClassVar[int]
    HEADERS_FIELD_NUMBER: _ClassVar[int]
    PAYLOAD_FIELD_NUMBER: _ClassVar[int]
    uuid: _wrappers_pb2.StringValue
    method: GenericWebhookConfig.MethodType
    headers: _containers.ScalarMap[str, str]
    payload: _wrappers_pb2.StringValue
    def __init__(self, uuid: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., method: _Optional[_Union[GenericWebhookConfig.MethodType, str]] = ..., headers: _Optional[_Mapping[str, str]] = ..., payload: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class SlackConfig(_message.Message):
    __slots__ = ("digests", "attachments")
    class DigestType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        UNKNOWN: _ClassVar[SlackConfig.DigestType]
        ERROR_AND_CRITICAL_LOGS: _ClassVar[SlackConfig.DigestType]
        FLOW_ANOMALIES: _ClassVar[SlackConfig.DigestType]
        SPIKE_ANOMALIES: _ClassVar[SlackConfig.DigestType]
        DATA_USAGE: _ClassVar[SlackConfig.DigestType]
    UNKNOWN: SlackConfig.DigestType
    ERROR_AND_CRITICAL_LOGS: SlackConfig.DigestType
    FLOW_ANOMALIES: SlackConfig.DigestType
    SPIKE_ANOMALIES: SlackConfig.DigestType
    DATA_USAGE: SlackConfig.DigestType
    class AttachmentType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        EMPTY: _ClassVar[SlackConfig.AttachmentType]
        METRIC_SNAPSHOT: _ClassVar[SlackConfig.AttachmentType]
        LOGS: _ClassVar[SlackConfig.AttachmentType]
    EMPTY: SlackConfig.AttachmentType
    METRIC_SNAPSHOT: SlackConfig.AttachmentType
    LOGS: SlackConfig.AttachmentType
    class Digest(_message.Message):
        __slots__ = ("type", "is_active")
        TYPE_FIELD_NUMBER: _ClassVar[int]
        IS_ACTIVE_FIELD_NUMBER: _ClassVar[int]
        type: SlackConfig.DigestType
        is_active: _wrappers_pb2.BoolValue
        def __init__(self, type: _Optional[_Union[SlackConfig.DigestType, str]] = ..., is_active: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
    class Attachment(_message.Message):
        __slots__ = ("type", "is_active")
        TYPE_FIELD_NUMBER: _ClassVar[int]
        IS_ACTIVE_FIELD_NUMBER: _ClassVar[int]
        type: SlackConfig.AttachmentType
        is_active: _wrappers_pb2.BoolValue
        def __init__(self, type: _Optional[_Union[SlackConfig.AttachmentType, str]] = ..., is_active: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
    DIGESTS_FIELD_NUMBER: _ClassVar[int]
    ATTACHMENTS_FIELD_NUMBER: _ClassVar[int]
    digests: _containers.RepeatedCompositeFieldContainer[SlackConfig.Digest]
    attachments: _containers.RepeatedCompositeFieldContainer[SlackConfig.Attachment]
    def __init__(self, digests: _Optional[_Iterable[_Union[SlackConfig.Digest, _Mapping]]] = ..., attachments: _Optional[_Iterable[_Union[SlackConfig.Attachment, _Mapping]]] = ...) -> None: ...

class PagerDutyConfig(_message.Message):
    __slots__ = ("service_key",)
    SERVICE_KEY_FIELD_NUMBER: _ClassVar[int]
    service_key: _wrappers_pb2.StringValue
    def __init__(self, service_key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class SendLogConfig(_message.Message):
    __slots__ = ("uuid", "payload")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PAYLOAD_FIELD_NUMBER: _ClassVar[int]
    uuid: _wrappers_pb2.StringValue
    payload: _wrappers_pb2.StringValue
    def __init__(self, uuid: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., payload: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class EmailGroupConfig(_message.Message):
    __slots__ = ("email_addresses",)
    EMAIL_ADDRESSES_FIELD_NUMBER: _ClassVar[int]
    email_addresses: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, email_addresses: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class MicrosoftTeamsConfig(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class JiraConfig(_message.Message):
    __slots__ = ("api_token", "email", "project_key")
    API_TOKEN_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    PROJECT_KEY_FIELD_NUMBER: _ClassVar[int]
    api_token: _wrappers_pb2.StringValue
    email: _wrappers_pb2.StringValue
    project_key: _wrappers_pb2.StringValue
    def __init__(self, api_token: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., email: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., project_key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class OpsgenieConfig(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DemistoConfig(_message.Message):
    __slots__ = ("uuid", "payload")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PAYLOAD_FIELD_NUMBER: _ClassVar[int]
    uuid: _wrappers_pb2.StringValue
    payload: _wrappers_pb2.StringValue
    def __init__(self, uuid: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., payload: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class AwsEventBridgeConfig(_message.Message):
    __slots__ = ("event_bus_arn", "detail", "detail_type", "source", "role_name")
    EVENT_BUS_ARN_FIELD_NUMBER: _ClassVar[int]
    DETAIL_FIELD_NUMBER: _ClassVar[int]
    DETAIL_TYPE_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    ROLE_NAME_FIELD_NUMBER: _ClassVar[int]
    event_bus_arn: _wrappers_pb2.StringValue
    detail: _wrappers_pb2.StringValue
    detail_type: _wrappers_pb2.StringValue
    source: _wrappers_pb2.StringValue
    role_name: _wrappers_pb2.StringValue
    def __init__(self, event_bus_arn: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., detail: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., detail_type: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., source: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., role_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class IbmEventNotificationsConfig(_message.Message):
    __slots__ = ("event_notifications_instance_id", "region_id")
    EVENT_NOTIFICATIONS_INSTANCE_ID_FIELD_NUMBER: _ClassVar[int]
    REGION_ID_FIELD_NUMBER: _ClassVar[int]
    event_notifications_instance_id: _wrappers_pb2.StringValue
    region_id: _wrappers_pb2.StringValue
    def __init__(self, event_notifications_instance_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., region_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class OutgoingWebhook(_message.Message):
    __slots__ = ("id", "type", "name", "url", "created_at", "updated_at", "external_id", "generic_webhook", "slack", "pager_duty", "send_log", "email_group", "microsoft_teams", "jira", "opsgenie", "demisto", "aws_event_bridge", "ibm_event_notifications")
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    URL_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    EXTERNAL_ID_FIELD_NUMBER: _ClassVar[int]
    GENERIC_WEBHOOK_FIELD_NUMBER: _ClassVar[int]
    SLACK_FIELD_NUMBER: _ClassVar[int]
    PAGER_DUTY_FIELD_NUMBER: _ClassVar[int]
    SEND_LOG_FIELD_NUMBER: _ClassVar[int]
    EMAIL_GROUP_FIELD_NUMBER: _ClassVar[int]
    MICROSOFT_TEAMS_FIELD_NUMBER: _ClassVar[int]
    JIRA_FIELD_NUMBER: _ClassVar[int]
    OPSGENIE_FIELD_NUMBER: _ClassVar[int]
    DEMISTO_FIELD_NUMBER: _ClassVar[int]
    AWS_EVENT_BRIDGE_FIELD_NUMBER: _ClassVar[int]
    IBM_EVENT_NOTIFICATIONS_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    type: WebhookType
    name: _wrappers_pb2.StringValue
    url: _wrappers_pb2.StringValue
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    external_id: _wrappers_pb2.UInt32Value
    generic_webhook: GenericWebhookConfig
    slack: SlackConfig
    pager_duty: PagerDutyConfig
    send_log: SendLogConfig
    email_group: EmailGroupConfig
    microsoft_teams: MicrosoftTeamsConfig
    jira: JiraConfig
    opsgenie: OpsgenieConfig
    demisto: DemistoConfig
    aws_event_bridge: AwsEventBridgeConfig
    ibm_event_notifications: IbmEventNotificationsConfig
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., type: _Optional[_Union[WebhookType, str]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., external_id: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., generic_webhook: _Optional[_Union[GenericWebhookConfig, _Mapping]] = ..., slack: _Optional[_Union[SlackConfig, _Mapping]] = ..., pager_duty: _Optional[_Union[PagerDutyConfig, _Mapping]] = ..., send_log: _Optional[_Union[SendLogConfig, _Mapping]] = ..., email_group: _Optional[_Union[EmailGroupConfig, _Mapping]] = ..., microsoft_teams: _Optional[_Union[MicrosoftTeamsConfig, _Mapping]] = ..., jira: _Optional[_Union[JiraConfig, _Mapping]] = ..., opsgenie: _Optional[_Union[OpsgenieConfig, _Mapping]] = ..., demisto: _Optional[_Union[DemistoConfig, _Mapping]] = ..., aws_event_bridge: _Optional[_Union[AwsEventBridgeConfig, _Mapping]] = ..., ibm_event_notifications: _Optional[_Union[IbmEventNotificationsConfig, _Mapping]] = ...) -> None: ...

class OutgoingWebhookInputData(_message.Message):
    __slots__ = ("type", "name", "url", "generic_webhook", "slack", "pager_duty", "send_log", "email_group", "microsoft_teams", "jira", "opsgenie", "demisto", "aws_event_bridge", "ibm_event_notifications")
    TYPE_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    URL_FIELD_NUMBER: _ClassVar[int]
    GENERIC_WEBHOOK_FIELD_NUMBER: _ClassVar[int]
    SLACK_FIELD_NUMBER: _ClassVar[int]
    PAGER_DUTY_FIELD_NUMBER: _ClassVar[int]
    SEND_LOG_FIELD_NUMBER: _ClassVar[int]
    EMAIL_GROUP_FIELD_NUMBER: _ClassVar[int]
    MICROSOFT_TEAMS_FIELD_NUMBER: _ClassVar[int]
    JIRA_FIELD_NUMBER: _ClassVar[int]
    OPSGENIE_FIELD_NUMBER: _ClassVar[int]
    DEMISTO_FIELD_NUMBER: _ClassVar[int]
    AWS_EVENT_BRIDGE_FIELD_NUMBER: _ClassVar[int]
    IBM_EVENT_NOTIFICATIONS_FIELD_NUMBER: _ClassVar[int]
    type: WebhookType
    name: _wrappers_pb2.StringValue
    url: _wrappers_pb2.StringValue
    generic_webhook: GenericWebhookConfig
    slack: SlackConfig
    pager_duty: PagerDutyConfig
    send_log: SendLogConfig
    email_group: EmailGroupConfig
    microsoft_teams: MicrosoftTeamsConfig
    jira: JiraConfig
    opsgenie: OpsgenieConfig
    demisto: DemistoConfig
    aws_event_bridge: AwsEventBridgeConfig
    ibm_event_notifications: IbmEventNotificationsConfig
    def __init__(self, type: _Optional[_Union[WebhookType, str]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., generic_webhook: _Optional[_Union[GenericWebhookConfig, _Mapping]] = ..., slack: _Optional[_Union[SlackConfig, _Mapping]] = ..., pager_duty: _Optional[_Union[PagerDutyConfig, _Mapping]] = ..., send_log: _Optional[_Union[SendLogConfig, _Mapping]] = ..., email_group: _Optional[_Union[EmailGroupConfig, _Mapping]] = ..., microsoft_teams: _Optional[_Union[MicrosoftTeamsConfig, _Mapping]] = ..., jira: _Optional[_Union[JiraConfig, _Mapping]] = ..., opsgenie: _Optional[_Union[OpsgenieConfig, _Mapping]] = ..., demisto: _Optional[_Union[DemistoConfig, _Mapping]] = ..., aws_event_bridge: _Optional[_Union[AwsEventBridgeConfig, _Mapping]] = ..., ibm_event_notifications: _Optional[_Union[IbmEventNotificationsConfig, _Mapping]] = ...) -> None: ...

class OutgoingWebhookSummary(_message.Message):
    __slots__ = ("id", "name", "url", "created_at", "updated_at", "external_id")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    URL_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    EXTERNAL_ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    url: _wrappers_pb2.StringValue
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    external_id: _wrappers_pb2.UInt32Value
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., external_id: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...

class OutgoingWebhookExtendedSummary(_message.Message):
    __slots__ = ("id", "type", "name", "url", "created_at", "updated_at", "external_id")
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    URL_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    EXTERNAL_ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    type: WebhookType
    name: _wrappers_pb2.StringValue
    url: _wrappers_pb2.StringValue
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    external_id: _wrappers_pb2.UInt32Value
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., type: _Optional[_Union[WebhookType, str]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., external_id: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...

class OutgoingWebhookDetails(_message.Message):
    __slots__ = ("type", "label")
    TYPE_FIELD_NUMBER: _ClassVar[int]
    LABEL_FIELD_NUMBER: _ClassVar[int]
    type: WebhookType
    label: _wrappers_pb2.StringValue
    def __init__(self, type: _Optional[_Union[WebhookType, str]] = ..., label: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
