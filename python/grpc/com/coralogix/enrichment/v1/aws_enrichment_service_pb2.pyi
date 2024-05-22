from com.coralogix.enrichment.v1 import audit_log_pb2 as _audit_log_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetSupportedAwsResourceTypesRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetSupportedAwsResourceTypesResponse(_message.Message):
    __slots__ = ("resource_types",)
    RESOURCE_TYPES_FIELD_NUMBER: _ClassVar[int]
    resource_types: _containers.RepeatedCompositeFieldContainer[SupportedAwsResourceType]
    def __init__(self, resource_types: _Optional[_Iterable[_Union[SupportedAwsResourceType, _Mapping]]] = ...) -> None: ...

class SupportedAwsResourceType(_message.Message):
    __slots__ = ("type", "display_name")
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DISPLAY_NAME_FIELD_NUMBER: _ClassVar[int]
    type: str
    display_name: str
    def __init__(self, type: _Optional[str] = ..., display_name: _Optional[str] = ...) -> None: ...
