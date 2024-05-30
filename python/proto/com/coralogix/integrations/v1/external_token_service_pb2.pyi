from com.coralogix.integrations.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.integrations.v1 import push_based_platform_pb2 as _push_based_platform_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GenerateNewTokenRequest(_message.Message):
    __slots__ = ("platform",)
    PLATFORM_FIELD_NUMBER: _ClassVar[int]
    platform: _push_based_platform_pb2.PushBasedPlatform
    def __init__(self, platform: _Optional[_Union[_push_based_platform_pb2.PushBasedPlatform, str]] = ...) -> None: ...

class GenerateNewTokenResponse(_message.Message):
    __slots__ = ("token", "id")
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    token: _wrappers_pb2.StringValue
    id: _wrappers_pb2.StringValue
    def __init__(self, token: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class UpdateTokenRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class UpdateTokenResponse(_message.Message):
    __slots__ = ("token",)
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    token: _wrappers_pb2.StringValue
    def __init__(self, token: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
