from com.coralogix.integrations.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.integrations.v1 import webhook_pb2 as _webhook_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateWebhookIntegrationRequest(_message.Message):
    __slots__ = ("name", "api_key_value", "application", "subsystem", "is_private", "json", "text")
    NAME_FIELD_NUMBER: _ClassVar[int]
    API_KEY_VALUE_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
    IS_PRIVATE_FIELD_NUMBER: _ClassVar[int]
    JSON_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    api_key_value: _wrappers_pb2.StringValue
    application: _wrappers_pb2.StringValue
    subsystem: _wrappers_pb2.StringValue
    is_private: _wrappers_pb2.BoolValue
    json: _webhook_pb2.JsonContentType
    text: _webhook_pb2.TextContentType
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., api_key_value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., application: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., is_private: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., json: _Optional[_Union[_webhook_pb2.JsonContentType, _Mapping]] = ..., text: _Optional[_Union[_webhook_pb2.TextContentType, _Mapping]] = ...) -> None: ...

class CreateWebhookIntegrationResponse(_message.Message):
    __slots__ = ("id", "webhook_url")
    ID_FIELD_NUMBER: _ClassVar[int]
    WEBHOOK_URL_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    webhook_url: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., webhook_url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class ListWebhookIntegrationsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListWebhookIntegrationsResponse(_message.Message):
    __slots__ = ("webhooks",)
    class WebhookIntegration(_message.Message):
        __slots__ = ("webhook_id", "company_id", "name", "application", "subsystem", "is_private", "url", "is_active", "created_at", "updated_at")
        WEBHOOK_ID_FIELD_NUMBER: _ClassVar[int]
        COMPANY_ID_FIELD_NUMBER: _ClassVar[int]
        NAME_FIELD_NUMBER: _ClassVar[int]
        APPLICATION_FIELD_NUMBER: _ClassVar[int]
        SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
        IS_PRIVATE_FIELD_NUMBER: _ClassVar[int]
        URL_FIELD_NUMBER: _ClassVar[int]
        IS_ACTIVE_FIELD_NUMBER: _ClassVar[int]
        CREATED_AT_FIELD_NUMBER: _ClassVar[int]
        UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
        webhook_id: _wrappers_pb2.StringValue
        company_id: _wrappers_pb2.StringValue
        name: _wrappers_pb2.StringValue
        application: _wrappers_pb2.StringValue
        subsystem: _wrappers_pb2.StringValue
        is_private: _wrappers_pb2.BoolValue
        url: _wrappers_pb2.StringValue
        is_active: _wrappers_pb2.BoolValue
        created_at: _timestamp_pb2.Timestamp
        updated_at: _timestamp_pb2.Timestamp
        def __init__(self, webhook_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., company_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., application: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., is_private: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., is_active: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
    WEBHOOKS_FIELD_NUMBER: _ClassVar[int]
    webhooks: _containers.RepeatedCompositeFieldContainer[ListWebhookIntegrationsResponse.WebhookIntegration]
    def __init__(self, webhooks: _Optional[_Iterable[_Union[ListWebhookIntegrationsResponse.WebhookIntegration, _Mapping]]] = ...) -> None: ...

class DeleteWebhookIntegrationRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class DeleteWebhookIntegrationResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ToggleWebhookIntegrationActivationRequest(_message.Message):
    __slots__ = ("id", "is_active")
    ID_FIELD_NUMBER: _ClassVar[int]
    IS_ACTIVE_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    is_active: _wrappers_pb2.BoolValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., is_active: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class ToggleWebhookIntegrationActivationResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class CountWebhookIntegrationsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class CountWebhookIntegrationsResponse(_message.Message):
    __slots__ = ("count",)
    COUNT_FIELD_NUMBER: _ClassVar[int]
    count: _wrappers_pb2.UInt32Value
    def __init__(self, count: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...
