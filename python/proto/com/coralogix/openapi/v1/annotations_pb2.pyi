from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Cloud(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    CLOUD_UNSPECIFIED: _ClassVar[Cloud]
    CLOUD_IBM: _ClassVar[Cloud]
    CLOUD_AWS: _ClassVar[Cloud]
    CLOUD_AZURE: _ClassVar[Cloud]
    CLOUD_GCP: _ClassVar[Cloud]
CLOUD_UNSPECIFIED: Cloud
CLOUD_IBM: Cloud
CLOUD_AWS: Cloud
CLOUD_AZURE: Cloud
CLOUD_GCP: Cloud
SERVICE_FIELD_NUMBER: _ClassVar[int]
service: _descriptor.FieldDescriptor
METHOD_FIELD_NUMBER: _ClassVar[int]
method: _descriptor.FieldDescriptor
MESSAGE_FIELD_NUMBER: _ClassVar[int]
message: _descriptor.FieldDescriptor
ENUM_FIELD_NUMBER: _ClassVar[int]
enum: _descriptor.FieldDescriptor
ENUM_VALUE_FIELD_NUMBER: _ClassVar[int]
enum_value: _descriptor.FieldDescriptor
FIELD_FIELD_NUMBER: _ClassVar[int]
field: _descriptor.FieldDescriptor

class ServiceOpenAPI(_message.Message):
    __slots__ = ("description", "version", "error_model", "resource_name")
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    ERROR_MODEL_FIELD_NUMBER: _ClassVar[int]
    RESOURCE_NAME_FIELD_NUMBER: _ClassVar[int]
    description: str
    version: str
    error_model: str
    resource_name: str
    def __init__(self, description: _Optional[str] = ..., version: _Optional[str] = ..., error_model: _Optional[str] = ..., resource_name: _Optional[str] = ...) -> None: ...

class MethodOpenAPI(_message.Message):
    __slots__ = ("resource", "description", "request_schema", "response_schema", "error_responses", "http_rule_index", "path_parameter_name", "extensions", "datasource", "request_body_example", "response_body_example")
    class ErrorResponsesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: int
        value: str
        def __init__(self, key: _Optional[int] = ..., value: _Optional[str] = ...) -> None: ...
    class ExtensionsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    RESOURCE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    REQUEST_SCHEMA_FIELD_NUMBER: _ClassVar[int]
    RESPONSE_SCHEMA_FIELD_NUMBER: _ClassVar[int]
    ERROR_RESPONSES_FIELD_NUMBER: _ClassVar[int]
    HTTP_RULE_INDEX_FIELD_NUMBER: _ClassVar[int]
    PATH_PARAMETER_NAME_FIELD_NUMBER: _ClassVar[int]
    EXTENSIONS_FIELD_NUMBER: _ClassVar[int]
    DATASOURCE_FIELD_NUMBER: _ClassVar[int]
    REQUEST_BODY_EXAMPLE_FIELD_NUMBER: _ClassVar[int]
    RESPONSE_BODY_EXAMPLE_FIELD_NUMBER: _ClassVar[int]
    resource: bool
    description: str
    request_schema: str
    response_schema: str
    error_responses: _containers.ScalarMap[int, str]
    http_rule_index: int
    path_parameter_name: str
    extensions: _containers.ScalarMap[str, str]
    datasource: bool
    request_body_example: str
    response_body_example: str
    def __init__(self, resource: bool = ..., description: _Optional[str] = ..., request_schema: _Optional[str] = ..., response_schema: _Optional[str] = ..., error_responses: _Optional[_Mapping[int, str]] = ..., http_rule_index: _Optional[int] = ..., path_parameter_name: _Optional[str] = ..., extensions: _Optional[_Mapping[str, str]] = ..., datasource: bool = ..., request_body_example: _Optional[str] = ..., response_body_example: _Optional[str] = ...) -> None: ...

class MessageOpenAPI(_message.Message):
    __slots__ = ("name", "description", "resource", "extensions")
    class ExtensionsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    RESOURCE_FIELD_NUMBER: _ClassVar[int]
    EXTENSIONS_FIELD_NUMBER: _ClassVar[int]
    name: str
    description: str
    resource: bool
    extensions: _containers.ScalarMap[str, str]
    def __init__(self, name: _Optional[str] = ..., description: _Optional[str] = ..., resource: bool = ..., extensions: _Optional[_Mapping[str, str]] = ...) -> None: ...

class EnumOpenAPI(_message.Message):
    __slots__ = ("name", "description", "extensions")
    class ExtensionsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    EXTENSIONS_FIELD_NUMBER: _ClassVar[int]
    name: str
    description: str
    extensions: _containers.ScalarMap[str, str]
    def __init__(self, name: _Optional[str] = ..., description: _Optional[str] = ..., extensions: _Optional[_Mapping[str, str]] = ...) -> None: ...

class EnumValueOpenAPI(_message.Message):
    __slots__ = ("skip_in",)
    SKIP_IN_FIELD_NUMBER: _ClassVar[int]
    skip_in: _containers.RepeatedScalarFieldContainer[Cloud]
    def __init__(self, skip_in: _Optional[_Iterable[_Union[Cloud, str]]] = ...) -> None: ...

class FieldOpenAPI(_message.Message):
    __slots__ = ("name", "description", "example", "pattern", "min_length", "max_length", "min_items", "max_items", "required", "identifier", "format", "skip_in")
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    EXAMPLE_FIELD_NUMBER: _ClassVar[int]
    PATTERN_FIELD_NUMBER: _ClassVar[int]
    MIN_LENGTH_FIELD_NUMBER: _ClassVar[int]
    MAX_LENGTH_FIELD_NUMBER: _ClassVar[int]
    MIN_ITEMS_FIELD_NUMBER: _ClassVar[int]
    MAX_ITEMS_FIELD_NUMBER: _ClassVar[int]
    REQUIRED_FIELD_NUMBER: _ClassVar[int]
    IDENTIFIER_FIELD_NUMBER: _ClassVar[int]
    FORMAT_FIELD_NUMBER: _ClassVar[int]
    SKIP_IN_FIELD_NUMBER: _ClassVar[int]
    name: str
    description: str
    example: str
    pattern: str
    min_length: int
    max_length: int
    min_items: int
    max_items: int
    required: bool
    identifier: bool
    format: str
    skip_in: _containers.RepeatedScalarFieldContainer[Cloud]
    def __init__(self, name: _Optional[str] = ..., description: _Optional[str] = ..., example: _Optional[str] = ..., pattern: _Optional[str] = ..., min_length: _Optional[int] = ..., max_length: _Optional[int] = ..., min_items: _Optional[int] = ..., max_items: _Optional[int] = ..., required: bool = ..., identifier: bool = ..., format: _Optional[str] = ..., skip_in: _Optional[_Iterable[_Union[Cloud, str]]] = ...) -> None: ...

class ApiError(_message.Message):
    __slots__ = ("errors", "trace", "status_code")
    class Error(_message.Message):
        __slots__ = ("code", "message", "more_info")
        class ErrorCode(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            ERROR_CODE_BAD_REQUEST_OR_UNSPECIFIED: _ClassVar[ApiError.Error.ErrorCode]
            ERROR_CODE_UNAUTHORIZED: _ClassVar[ApiError.Error.ErrorCode]
            ERROR_CODE_FORBIDDEN: _ClassVar[ApiError.Error.ErrorCode]
            ERROR_CODE_NOT_FOUND: _ClassVar[ApiError.Error.ErrorCode]
            ERROR_CODE_METHOD_INTERNAL_ERROR: _ClassVar[ApiError.Error.ErrorCode]
            ERROR_CODE_CONFLICT: _ClassVar[ApiError.Error.ErrorCode]
            ERROR_CODE_UNAUTHENTICATED: _ClassVar[ApiError.Error.ErrorCode]
            ERROR_CODE_RESOURCE_EXHAUSTED: _ClassVar[ApiError.Error.ErrorCode]
            ERROR_CODE_DEADLINE_EXCEEDED: _ClassVar[ApiError.Error.ErrorCode]
        ERROR_CODE_BAD_REQUEST_OR_UNSPECIFIED: ApiError.Error.ErrorCode
        ERROR_CODE_UNAUTHORIZED: ApiError.Error.ErrorCode
        ERROR_CODE_FORBIDDEN: ApiError.Error.ErrorCode
        ERROR_CODE_NOT_FOUND: ApiError.Error.ErrorCode
        ERROR_CODE_METHOD_INTERNAL_ERROR: ApiError.Error.ErrorCode
        ERROR_CODE_CONFLICT: ApiError.Error.ErrorCode
        ERROR_CODE_UNAUTHENTICATED: ApiError.Error.ErrorCode
        ERROR_CODE_RESOURCE_EXHAUSTED: ApiError.Error.ErrorCode
        ERROR_CODE_DEADLINE_EXCEEDED: ApiError.Error.ErrorCode
        CODE_FIELD_NUMBER: _ClassVar[int]
        MESSAGE_FIELD_NUMBER: _ClassVar[int]
        MORE_INFO_FIELD_NUMBER: _ClassVar[int]
        code: ApiError.Error.ErrorCode
        message: _wrappers_pb2.StringValue
        more_info: _wrappers_pb2.StringValue
        def __init__(self, code: _Optional[_Union[ApiError.Error.ErrorCode, str]] = ..., message: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., more_info: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    ERRORS_FIELD_NUMBER: _ClassVar[int]
    TRACE_FIELD_NUMBER: _ClassVar[int]
    STATUS_CODE_FIELD_NUMBER: _ClassVar[int]
    errors: _containers.RepeatedCompositeFieldContainer[ApiError.Error]
    trace: _wrappers_pb2.StringValue
    status_code: _wrappers_pb2.Int32Value
    def __init__(self, errors: _Optional[_Iterable[_Union[ApiError.Error, _Mapping]]] = ..., trace: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., status_code: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
