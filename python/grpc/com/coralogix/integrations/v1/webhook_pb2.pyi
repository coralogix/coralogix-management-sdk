from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class TimestampSourceGenerate(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class TimestampFormatAutomatic(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class TimestampFormatCustom(_message.Message):
    __slots__ = ("format",)
    FORMAT_FIELD_NUMBER: _ClassVar[int]
    format: _wrappers_pb2.StringValue
    def __init__(self, format: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class TimestampSourceJsonPath(_message.Message):
    __slots__ = ("json_path", "automatic", "custom")
    JSON_PATH_FIELD_NUMBER: _ClassVar[int]
    AUTOMATIC_FIELD_NUMBER: _ClassVar[int]
    CUSTOM_FIELD_NUMBER: _ClassVar[int]
    json_path: _wrappers_pb2.StringValue
    automatic: TimestampFormatAutomatic
    custom: TimestampFormatCustom
    def __init__(self, json_path: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., automatic: _Optional[_Union[TimestampFormatAutomatic, _Mapping]] = ..., custom: _Optional[_Union[TimestampFormatCustom, _Mapping]] = ...) -> None: ...

class TimestampSourceRegex(_message.Message):
    __slots__ = ("regex", "automatic", "custom")
    REGEX_FIELD_NUMBER: _ClassVar[int]
    AUTOMATIC_FIELD_NUMBER: _ClassVar[int]
    CUSTOM_FIELD_NUMBER: _ClassVar[int]
    regex: _wrappers_pb2.StringValue
    automatic: TimestampFormatAutomatic
    custom: TimestampFormatCustom
    def __init__(self, regex: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., automatic: _Optional[_Union[TimestampFormatAutomatic, _Mapping]] = ..., custom: _Optional[_Union[TimestampFormatCustom, _Mapping]] = ...) -> None: ...

class JsonContentType(_message.Message):
    __slots__ = ("generate", "json_path")
    GENERATE_FIELD_NUMBER: _ClassVar[int]
    JSON_PATH_FIELD_NUMBER: _ClassVar[int]
    generate: TimestampSourceGenerate
    json_path: TimestampSourceJsonPath
    def __init__(self, generate: _Optional[_Union[TimestampSourceGenerate, _Mapping]] = ..., json_path: _Optional[_Union[TimestampSourceJsonPath, _Mapping]] = ...) -> None: ...

class TextContentType(_message.Message):
    __slots__ = ("generate", "regex")
    GENERATE_FIELD_NUMBER: _ClassVar[int]
    REGEX_FIELD_NUMBER: _ClassVar[int]
    generate: TimestampSourceGenerate
    regex: TimestampSourceRegex
    def __init__(self, generate: _Optional[_Union[TimestampSourceGenerate, _Mapping]] = ..., regex: _Optional[_Union[TimestampSourceRegex, _Mapping]] = ...) -> None: ...
