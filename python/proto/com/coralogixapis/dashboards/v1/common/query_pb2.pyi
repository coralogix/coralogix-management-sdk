from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DataprimeQuery(_message.Message):
    __slots__ = ("text",)
    TEXT_FIELD_NUMBER: _ClassVar[int]
    text: str
    def __init__(self, text: _Optional[str] = ...) -> None: ...

class SerializedDataprimeQuery(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: bytes
    def __init__(self, data: _Optional[bytes] = ...) -> None: ...

class FullDataprimeQuery(_message.Message):
    __slots__ = ("serialized", "raw")
    SERIALIZED_FIELD_NUMBER: _ClassVar[int]
    RAW_FIELD_NUMBER: _ClassVar[int]
    serialized: SerializedDataprimeQuery
    raw: DataprimeQuery
    def __init__(self, serialized: _Optional[_Union[SerializedDataprimeQuery, _Mapping]] = ..., raw: _Optional[_Union[DataprimeQuery, _Mapping]] = ...) -> None: ...

class PromQlQuery(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: _wrappers_pb2.StringValue
    def __init__(self, value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class LuceneQuery(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: _wrappers_pb2.StringValue
    def __init__(self, value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
