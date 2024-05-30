from com.coralogix.integrations.v1 import integration_pb2 as _integration_pb2
from com.coralogix.integrations.v1 import audit_log_pb2 as _audit_log_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DeployIntegrationExtensionsRequest(_message.Message):
    __slots__ = ("integration_key",)
    INTEGRATION_KEY_FIELD_NUMBER: _ClassVar[int]
    integration_key: _wrappers_pb2.StringValue
    def __init__(self, integration_key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class DeployIntegrationExtensionsResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
