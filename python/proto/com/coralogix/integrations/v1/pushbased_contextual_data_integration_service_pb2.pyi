from com.coralogix.integrations.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.integrations.v1 import push_based_platform_pb2 as _push_based_platform_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreatePushBasedIntegrationRequest(_message.Message):
    __slots__ = ("bitbucket", "amazon_sns", "pager_duty", "github", "gitlab", "opsgenie", "prometheus", "intercom", "slack", "name", "application", "subsystem")
    class Bitbucket(_message.Message):
        __slots__ = ("internal_token_id",)
        INTERNAL_TOKEN_ID_FIELD_NUMBER: _ClassVar[int]
        internal_token_id: _wrappers_pb2.StringValue
        def __init__(self, internal_token_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class AmazonSns(_message.Message):
        __slots__ = ("internal_token_id",)
        INTERNAL_TOKEN_ID_FIELD_NUMBER: _ClassVar[int]
        internal_token_id: _wrappers_pb2.StringValue
        def __init__(self, internal_token_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class PagerDuty(_message.Message):
        __slots__ = ("internal_token_id",)
        INTERNAL_TOKEN_ID_FIELD_NUMBER: _ClassVar[int]
        internal_token_id: _wrappers_pb2.StringValue
        def __init__(self, internal_token_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class Github(_message.Message):
        __slots__ = ("internal_token_id",)
        INTERNAL_TOKEN_ID_FIELD_NUMBER: _ClassVar[int]
        internal_token_id: _wrappers_pb2.StringValue
        def __init__(self, internal_token_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class Gitlab(_message.Message):
        __slots__ = ("internal_token_id",)
        INTERNAL_TOKEN_ID_FIELD_NUMBER: _ClassVar[int]
        internal_token_id: _wrappers_pb2.StringValue
        def __init__(self, internal_token_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class Opsgenie(_message.Message):
        __slots__ = ("internal_token_id",)
        INTERNAL_TOKEN_ID_FIELD_NUMBER: _ClassVar[int]
        internal_token_id: _wrappers_pb2.StringValue
        def __init__(self, internal_token_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class Prometheus(_message.Message):
        __slots__ = ("internal_token_id",)
        INTERNAL_TOKEN_ID_FIELD_NUMBER: _ClassVar[int]
        internal_token_id: _wrappers_pb2.StringValue
        def __init__(self, internal_token_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class Intercom(_message.Message):
        __slots__ = ("token",)
        TOKEN_FIELD_NUMBER: _ClassVar[int]
        token: _wrappers_pb2.StringValue
        def __init__(self, token: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class Slack(_message.Message):
        __slots__ = ("internal_token_id",)
        INTERNAL_TOKEN_ID_FIELD_NUMBER: _ClassVar[int]
        internal_token_id: _wrappers_pb2.StringValue
        def __init__(self, internal_token_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    BITBUCKET_FIELD_NUMBER: _ClassVar[int]
    AMAZON_SNS_FIELD_NUMBER: _ClassVar[int]
    PAGER_DUTY_FIELD_NUMBER: _ClassVar[int]
    GITHUB_FIELD_NUMBER: _ClassVar[int]
    GITLAB_FIELD_NUMBER: _ClassVar[int]
    OPSGENIE_FIELD_NUMBER: _ClassVar[int]
    PROMETHEUS_FIELD_NUMBER: _ClassVar[int]
    INTERCOM_FIELD_NUMBER: _ClassVar[int]
    SLACK_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
    bitbucket: CreatePushBasedIntegrationRequest.Bitbucket
    amazon_sns: CreatePushBasedIntegrationRequest.AmazonSns
    pager_duty: CreatePushBasedIntegrationRequest.PagerDuty
    github: CreatePushBasedIntegrationRequest.Github
    gitlab: CreatePushBasedIntegrationRequest.Gitlab
    opsgenie: CreatePushBasedIntegrationRequest.Opsgenie
    prometheus: CreatePushBasedIntegrationRequest.Prometheus
    intercom: CreatePushBasedIntegrationRequest.Intercom
    slack: CreatePushBasedIntegrationRequest.Slack
    name: _wrappers_pb2.StringValue
    application: _wrappers_pb2.StringValue
    subsystem: _wrappers_pb2.StringValue
    def __init__(self, bitbucket: _Optional[_Union[CreatePushBasedIntegrationRequest.Bitbucket, _Mapping]] = ..., amazon_sns: _Optional[_Union[CreatePushBasedIntegrationRequest.AmazonSns, _Mapping]] = ..., pager_duty: _Optional[_Union[CreatePushBasedIntegrationRequest.PagerDuty, _Mapping]] = ..., github: _Optional[_Union[CreatePushBasedIntegrationRequest.Github, _Mapping]] = ..., gitlab: _Optional[_Union[CreatePushBasedIntegrationRequest.Gitlab, _Mapping]] = ..., opsgenie: _Optional[_Union[CreatePushBasedIntegrationRequest.Opsgenie, _Mapping]] = ..., prometheus: _Optional[_Union[CreatePushBasedIntegrationRequest.Prometheus, _Mapping]] = ..., intercom: _Optional[_Union[CreatePushBasedIntegrationRequest.Intercom, _Mapping]] = ..., slack: _Optional[_Union[CreatePushBasedIntegrationRequest.Slack, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., application: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class CreatePushBasedIntegrationResponse(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class ListPushBasedContextualDataIntegrationsRequest(_message.Message):
    __slots__ = ("platform",)
    PLATFORM_FIELD_NUMBER: _ClassVar[int]
    platform: _push_based_platform_pb2.PushBasedPlatform
    def __init__(self, platform: _Optional[_Union[_push_based_platform_pb2.PushBasedPlatform, str]] = ...) -> None: ...

class ListPushBasedContextualDataIntegrationsResponse(_message.Message):
    __slots__ = ("integrations",)
    class PushBasedContextualDataIntegration(_message.Message):
        __slots__ = ("id", "name", "application", "subsystem", "platform", "token", "created_at", "updated_at")
        ID_FIELD_NUMBER: _ClassVar[int]
        NAME_FIELD_NUMBER: _ClassVar[int]
        APPLICATION_FIELD_NUMBER: _ClassVar[int]
        SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
        PLATFORM_FIELD_NUMBER: _ClassVar[int]
        TOKEN_FIELD_NUMBER: _ClassVar[int]
        CREATED_AT_FIELD_NUMBER: _ClassVar[int]
        UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
        id: _wrappers_pb2.StringValue
        name: _wrappers_pb2.StringValue
        application: _wrappers_pb2.StringValue
        subsystem: _wrappers_pb2.StringValue
        platform: _push_based_platform_pb2.PushBasedPlatform
        token: _wrappers_pb2.StringValue
        created_at: _timestamp_pb2.Timestamp
        updated_at: _timestamp_pb2.Timestamp
        def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., application: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., platform: _Optional[_Union[_push_based_platform_pb2.PushBasedPlatform, str]] = ..., token: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
    INTEGRATIONS_FIELD_NUMBER: _ClassVar[int]
    integrations: _containers.RepeatedCompositeFieldContainer[ListPushBasedContextualDataIntegrationsResponse.PushBasedContextualDataIntegration]
    def __init__(self, integrations: _Optional[_Iterable[_Union[ListPushBasedContextualDataIntegrationsResponse.PushBasedContextualDataIntegration, _Mapping]]] = ...) -> None: ...

class DeletePushBasedContextualDataIntegrationRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class DeletePushBasedContextualDataIntegrationResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class CountPushBasedContextualDataIntegrationsRequest(_message.Message):
    __slots__ = ("platform",)
    PLATFORM_FIELD_NUMBER: _ClassVar[int]
    platform: _push_based_platform_pb2.PushBasedPlatform
    def __init__(self, platform: _Optional[_Union[_push_based_platform_pb2.PushBasedPlatform, str]] = ...) -> None: ...

class CountPushBasedContextualDataIntegrationsResponse(_message.Message):
    __slots__ = ("counts",)
    class IntegrationCount(_message.Message):
        __slots__ = ("platform", "count")
        PLATFORM_FIELD_NUMBER: _ClassVar[int]
        COUNT_FIELD_NUMBER: _ClassVar[int]
        platform: _push_based_platform_pb2.PushBasedPlatform
        count: _wrappers_pb2.UInt32Value
        def __init__(self, platform: _Optional[_Union[_push_based_platform_pb2.PushBasedPlatform, str]] = ..., count: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...
    COUNTS_FIELD_NUMBER: _ClassVar[int]
    counts: _containers.RepeatedCompositeFieldContainer[CountPushBasedContextualDataIntegrationsResponse.IntegrationCount]
    def __init__(self, counts: _Optional[_Iterable[_Union[CountPushBasedContextualDataIntegrationsResponse.IntegrationCount, _Mapping]]] = ...) -> None: ...

class UpdatePushBasedContextualDataIntegrationsRequest(_message.Message):
    __slots__ = ("id", "name", "application", "subsystem")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    application: _wrappers_pb2.StringValue
    subsystem: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., application: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class UpdatePushBasedContextualDataIntegrationsResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
