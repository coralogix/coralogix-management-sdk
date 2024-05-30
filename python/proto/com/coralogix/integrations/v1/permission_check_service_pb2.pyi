from com.coralogix.integrations.v1 import audit_log_pb2 as _audit_log_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ResourceId(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.UInt32Value
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...

class ActionId(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.UInt32Value
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...

class ApiKeyValue(_message.Message):
    __slots__ = ("key",)
    KEY_FIELD_NUMBER: _ClassVar[int]
    key: _wrappers_pb2.StringValue
    def __init__(self, key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class ApiKeyId(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class Permission(_message.Message):
    __slots__ = ("resource", "action")
    RESOURCE_FIELD_NUMBER: _ClassVar[int]
    ACTION_FIELD_NUMBER: _ClassVar[int]
    resource: ResourceId
    action: ActionId
    def __init__(self, resource: _Optional[_Union[ResourceId, _Mapping]] = ..., action: _Optional[_Union[ActionId, _Mapping]] = ...) -> None: ...

class CheckApiKeyPermissionRequest(_message.Message):
    __slots__ = ("permission", "api_key")
    PERMISSION_FIELD_NUMBER: _ClassVar[int]
    API_KEY_FIELD_NUMBER: _ClassVar[int]
    permission: Permission
    api_key: ApiKeyValue
    def __init__(self, permission: _Optional[_Union[Permission, _Mapping]] = ..., api_key: _Optional[_Union[ApiKeyValue, _Mapping]] = ...) -> None: ...

class CheckApiKeyPermissionResponse(_message.Message):
    __slots__ = ("unauthorized", "authorization_result")
    class KeyAuthorizationResult(_message.Message):
        __slots__ = ("api_key_id",)
        API_KEY_ID_FIELD_NUMBER: _ClassVar[int]
        api_key_id: ApiKeyId
        def __init__(self, api_key_id: _Optional[_Union[ApiKeyId, _Mapping]] = ...) -> None: ...
    class Unauthorized(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    UNAUTHORIZED_FIELD_NUMBER: _ClassVar[int]
    AUTHORIZATION_RESULT_FIELD_NUMBER: _ClassVar[int]
    unauthorized: CheckApiKeyPermissionResponse.Unauthorized
    authorization_result: CheckApiKeyPermissionResponse.KeyAuthorizationResult
    def __init__(self, unauthorized: _Optional[_Union[CheckApiKeyPermissionResponse.Unauthorized, _Mapping]] = ..., authorization_result: _Optional[_Union[CheckApiKeyPermissionResponse.KeyAuthorizationResult, _Mapping]] = ...) -> None: ...
