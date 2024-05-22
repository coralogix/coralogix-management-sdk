from com.coralogix.integrations.v1 import integration_pb2 as _integration_pb2
from com.coralogix.integrations.v1 import audit_log_pb2 as _audit_log_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetContextualDataIntegrationsRequest(_message.Message):
    __slots__ = ("include_testing_integrations",)
    INCLUDE_TESTING_INTEGRATIONS_FIELD_NUMBER: _ClassVar[int]
    include_testing_integrations: _wrappers_pb2.BoolValue
    def __init__(self, include_testing_integrations: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class GetContextualDataIntegrationsResponse(_message.Message):
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
    integrations: _containers.RepeatedCompositeFieldContainer[GetContextualDataIntegrationsResponse.IntegrationWithCounts]
    def __init__(self, integrations: _Optional[_Iterable[_Union[GetContextualDataIntegrationsResponse.IntegrationWithCounts, _Mapping]]] = ...) -> None: ...

class GetContextualDataIntegrationDefinitionRequest(_message.Message):
    __slots__ = ("id", "include_testing_integrations")
    ID_FIELD_NUMBER: _ClassVar[int]
    INCLUDE_TESTING_INTEGRATIONS_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    include_testing_integrations: _wrappers_pb2.BoolValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., include_testing_integrations: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class GetContextualDataIntegrationDefinitionResponse(_message.Message):
    __slots__ = ("integration_definition",)
    INTEGRATION_DEFINITION_FIELD_NUMBER: _ClassVar[int]
    integration_definition: _integration_pb2.IntegrationDefinition
    def __init__(self, integration_definition: _Optional[_Union[_integration_pb2.IntegrationDefinition, _Mapping]] = ...) -> None: ...

class GetContextualDataIntegrationDetailsRequest(_message.Message):
    __slots__ = ("id", "include_testing_revisions")
    ID_FIELD_NUMBER: _ClassVar[int]
    INCLUDE_TESTING_REVISIONS_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    include_testing_revisions: _wrappers_pb2.BoolValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., include_testing_revisions: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class GetContextualDataIntegrationDetailsResponse(_message.Message):
    __slots__ = ("integration_detail",)
    INTEGRATION_DETAIL_FIELD_NUMBER: _ClassVar[int]
    integration_detail: _integration_pb2.IntegrationDetails
    def __init__(self, integration_detail: _Optional[_Union[_integration_pb2.IntegrationDetails, _Mapping]] = ...) -> None: ...

class UpdateContextualDataIntegrationRequest(_message.Message):
    __slots__ = ("integration_id", "metadata")
    INTEGRATION_ID_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    integration_id: _wrappers_pb2.StringValue
    metadata: _integration_pb2.IntegrationMetadata
    def __init__(self, integration_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., metadata: _Optional[_Union[_integration_pb2.IntegrationMetadata, _Mapping]] = ...) -> None: ...

class UpdateContextualDataIntegrationResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class TestContextualDataIntegrationRequest(_message.Message):
    __slots__ = ("integration_data",)
    INTEGRATION_DATA_FIELD_NUMBER: _ClassVar[int]
    integration_data: _integration_pb2.IntegrationMetadata
    def __init__(self, integration_data: _Optional[_Union[_integration_pb2.IntegrationMetadata, _Mapping]] = ...) -> None: ...

class TestContextualDataIntegrationResponse(_message.Message):
    __slots__ = ("result",)
    RESULT_FIELD_NUMBER: _ClassVar[int]
    result: _integration_pb2.TestIntegrationResult
    def __init__(self, result: _Optional[_Union[_integration_pb2.TestIntegrationResult, _Mapping]] = ...) -> None: ...

class SaveContextualDataIntegrationRequest(_message.Message):
    __slots__ = ("metadata",)
    METADATA_FIELD_NUMBER: _ClassVar[int]
    metadata: _integration_pb2.IntegrationMetadata
    def __init__(self, metadata: _Optional[_Union[_integration_pb2.IntegrationMetadata, _Mapping]] = ...) -> None: ...

class SaveContextualDataIntegrationResponse(_message.Message):
    __slots__ = ("integration_id",)
    INTEGRATION_ID_FIELD_NUMBER: _ClassVar[int]
    integration_id: _wrappers_pb2.StringValue
    def __init__(self, integration_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class DeleteContextualDataIntegrationRequest(_message.Message):
    __slots__ = ("integration_id",)
    INTEGRATION_ID_FIELD_NUMBER: _ClassVar[int]
    integration_id: _wrappers_pb2.StringValue
    def __init__(self, integration_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class DeleteContextualDataIntegrationResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
