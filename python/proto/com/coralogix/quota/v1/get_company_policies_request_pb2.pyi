from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.quota.v1 import enums_pb2 as _enums_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetCompanyPoliciesRequest(_message.Message):
    __slots__ = ("enabled_only", "source_type")
    ENABLED_ONLY_FIELD_NUMBER: _ClassVar[int]
    SOURCE_TYPE_FIELD_NUMBER: _ClassVar[int]
    enabled_only: _wrappers_pb2.BoolValue
    source_type: _enums_pb2.SourceType
    def __init__(self, enabled_only: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., source_type: _Optional[_Union[_enums_pb2.SourceType, str]] = ...) -> None: ...
