from com.coralogix.integrations.v1 import integration_pb2 as _integration_pb2
from com.coralogix.integrations.v1 import audit_log_pb2 as _audit_log_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UpdateIntegrationRequest(_message.Message):
    __slots__ = ("id", "metadata")
    ID_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    metadata: _integration_pb2.IntegrationMetadata
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., metadata: _Optional[_Union[_integration_pb2.IntegrationMetadata, _Mapping]] = ...) -> None: ...

class UpdateIntegrationResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetIntegrationDefinitionRequest(_message.Message):
    __slots__ = ("id", "include_testing_revision")
    ID_FIELD_NUMBER: _ClassVar[int]
    INCLUDE_TESTING_REVISION_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    include_testing_revision: _wrappers_pb2.BoolValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., include_testing_revision: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class GetIntegrationDefinitionResponse(_message.Message):
    __slots__ = ("integration_definition",)
    INTEGRATION_DEFINITION_FIELD_NUMBER: _ClassVar[int]
    integration_definition: _integration_pb2.IntegrationDefinition
    def __init__(self, integration_definition: _Optional[_Union[_integration_pb2.IntegrationDefinition, _Mapping]] = ...) -> None: ...

class GetIntegrationsRequest(_message.Message):
    __slots__ = ("include_testing_revision",)
    INCLUDE_TESTING_REVISION_FIELD_NUMBER: _ClassVar[int]
    include_testing_revision: _wrappers_pb2.BoolValue
    def __init__(self, include_testing_revision: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class GetIntegrationsResponse(_message.Message):
    __slots__ = ("integrations",)
    class IntegrationWithCounts(_message.Message):
        __slots__ = ("integration", "amount_integrations", "errors", "upgrade_available", "is_new")
        INTEGRATION_FIELD_NUMBER: _ClassVar[int]
        AMOUNT_INTEGRATIONS_FIELD_NUMBER: _ClassVar[int]
        ERRORS_FIELD_NUMBER: _ClassVar[int]
        UPGRADE_AVAILABLE_FIELD_NUMBER: _ClassVar[int]
        IS_NEW_FIELD_NUMBER: _ClassVar[int]
        integration: _integration_pb2.Integration
        amount_integrations: _wrappers_pb2.UInt32Value
        errors: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        upgrade_available: _wrappers_pb2.BoolValue
        is_new: _wrappers_pb2.BoolValue
        def __init__(self, integration: _Optional[_Union[_integration_pb2.Integration, _Mapping]] = ..., amount_integrations: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., errors: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., upgrade_available: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., is_new: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
    INTEGRATIONS_FIELD_NUMBER: _ClassVar[int]
    integrations: _containers.RepeatedCompositeFieldContainer[GetIntegrationsResponse.IntegrationWithCounts]
    def __init__(self, integrations: _Optional[_Iterable[_Union[GetIntegrationsResponse.IntegrationWithCounts, _Mapping]]] = ...) -> None: ...

class GetIntegrationDetailsRequest(_message.Message):
    __slots__ = ("id", "include_testing_revision")
    ID_FIELD_NUMBER: _ClassVar[int]
    INCLUDE_TESTING_REVISION_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    include_testing_revision: _wrappers_pb2.BoolValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., include_testing_revision: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class GetIntegrationDetailsResponse(_message.Message):
    __slots__ = ("integration_detail",)
    INTEGRATION_DETAIL_FIELD_NUMBER: _ClassVar[int]
    integration_detail: _integration_pb2.IntegrationDetails
    def __init__(self, integration_detail: _Optional[_Union[_integration_pb2.IntegrationDetails, _Mapping]] = ...) -> None: ...

class GetManagedIntegrationStatusRequest(_message.Message):
    __slots__ = ("integration_id",)
    INTEGRATION_ID_FIELD_NUMBER: _ClassVar[int]
    integration_id: str
    def __init__(self, integration_id: _Optional[str] = ...) -> None: ...

class GetManagedIntegrationStatusResponse(_message.Message):
    __slots__ = ("integration_id", "status")
    INTEGRATION_ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    integration_id: str
    status: _integration_pb2.IntegrationStatus
    def __init__(self, integration_id: _Optional[str] = ..., status: _Optional[_Union[_integration_pb2.IntegrationStatus, _Mapping]] = ...) -> None: ...

class SaveIntegrationRequest(_message.Message):
    __slots__ = ("metadata",)
    METADATA_FIELD_NUMBER: _ClassVar[int]
    metadata: _integration_pb2.IntegrationMetadata
    def __init__(self, metadata: _Optional[_Union[_integration_pb2.IntegrationMetadata, _Mapping]] = ...) -> None: ...

class SaveIntegrationResponse(_message.Message):
    __slots__ = ("integration_id",)
    INTEGRATION_ID_FIELD_NUMBER: _ClassVar[int]
    integration_id: _wrappers_pb2.StringValue
    def __init__(self, integration_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class DeleteIntegrationRequest(_message.Message):
    __slots__ = ("integration_id",)
    INTEGRATION_ID_FIELD_NUMBER: _ClassVar[int]
    integration_id: _wrappers_pb2.StringValue
    def __init__(self, integration_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class DeleteIntegrationResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetTemplateRequest(_message.Message):
    __slots__ = ("integration_id", "common_arm_params", "empty")
    class CommonARMParams(_message.Message):
        __slots__ = ("logs_url", "api_key", "cgx_domain")
        LOGS_URL_FIELD_NUMBER: _ClassVar[int]
        API_KEY_FIELD_NUMBER: _ClassVar[int]
        CGX_DOMAIN_FIELD_NUMBER: _ClassVar[int]
        logs_url: _wrappers_pb2.StringValue
        api_key: _wrappers_pb2.StringValue
        cgx_domain: _wrappers_pb2.StringValue
        def __init__(self, logs_url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., api_key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., cgx_domain: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class Empty(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    INTEGRATION_ID_FIELD_NUMBER: _ClassVar[int]
    COMMON_ARM_PARAMS_FIELD_NUMBER: _ClassVar[int]
    EMPTY_FIELD_NUMBER: _ClassVar[int]
    integration_id: _wrappers_pb2.StringValue
    common_arm_params: GetTemplateRequest.CommonARMParams
    empty: GetTemplateRequest.Empty
    def __init__(self, integration_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., common_arm_params: _Optional[_Union[GetTemplateRequest.CommonARMParams, _Mapping]] = ..., empty: _Optional[_Union[GetTemplateRequest.Empty, _Mapping]] = ...) -> None: ...

class GetTemplateResponse(_message.Message):
    __slots__ = ("template_url",)
    TEMPLATE_URL_FIELD_NUMBER: _ClassVar[int]
    template_url: _wrappers_pb2.StringValue
    def __init__(self, template_url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetRumApplicationVersionDataRequest(_message.Message):
    __slots__ = ("application_name",)
    APPLICATION_NAME_FIELD_NUMBER: _ClassVar[int]
    application_name: _wrappers_pb2.StringValue
    def __init__(self, application_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetRumApplicationVersionDataResponse(_message.Message):
    __slots__ = ("version_data",)
    VERSION_DATA_FIELD_NUMBER: _ClassVar[int]
    version_data: _integration_pb2.RumVersionData
    def __init__(self, version_data: _Optional[_Union[_integration_pb2.RumVersionData, _Mapping]] = ...) -> None: ...

class SyncRumDataRequest(_message.Message):
    __slots__ = ("force",)
    FORCE_FIELD_NUMBER: _ClassVar[int]
    force: _wrappers_pb2.BoolValue
    def __init__(self, force: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class SyncRumDataResponse(_message.Message):
    __slots__ = ("sync_executed", "synced_at")
    SYNC_EXECUTED_FIELD_NUMBER: _ClassVar[int]
    SYNCED_AT_FIELD_NUMBER: _ClassVar[int]
    sync_executed: _wrappers_pb2.BoolValue
    synced_at: _timestamp_pb2.Timestamp
    def __init__(self, sync_executed: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., synced_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class TestIntegrationRequest(_message.Message):
    __slots__ = ("integration_data",)
    INTEGRATION_DATA_FIELD_NUMBER: _ClassVar[int]
    integration_data: _integration_pb2.IntegrationMetadata
    def __init__(self, integration_data: _Optional[_Union[_integration_pb2.IntegrationMetadata, _Mapping]] = ...) -> None: ...

class TestIntegrationResponse(_message.Message):
    __slots__ = ("result",)
    RESULT_FIELD_NUMBER: _ClassVar[int]
    result: _integration_pb2.TestIntegrationResult
    def __init__(self, result: _Optional[_Union[_integration_pb2.TestIntegrationResult, _Mapping]] = ...) -> None: ...
