from com.coralogix.extensions.v1 import audit_log_pb2 as _audit_log_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class InternalDeployExtensionRequest(_message.Message):
    __slots__ = ("company_id", "id")
    COMPANY_ID_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    company_id: _wrappers_pb2.StringValue
    id: _wrappers_pb2.StringValue
    def __init__(self, company_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class InternalDeployExtensionResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
