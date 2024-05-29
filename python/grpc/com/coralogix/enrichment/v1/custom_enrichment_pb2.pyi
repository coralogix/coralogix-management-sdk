from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CustomEnrichment(_message.Message):
    __slots__ = ("id", "name", "description", "version")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    id: int
    name: str
    description: str
    version: int
    def __init__(self, id: _Optional[int] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., version: _Optional[int] = ...) -> None: ...

class CustomEnrichmentData(_message.Message):
    __slots__ = ("definition", "textual", "binary")
    DEFINITION_FIELD_NUMBER: _ClassVar[int]
    TEXTUAL_FIELD_NUMBER: _ClassVar[int]
    BINARY_FIELD_NUMBER: _ClassVar[int]
    definition: CustomEnrichment
    textual: _wrappers_pb2.StringValue
    binary: _wrappers_pb2.BytesValue
    def __init__(self, definition: _Optional[_Union[CustomEnrichment, _Mapping]] = ..., textual: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., binary: _Optional[_Union[_wrappers_pb2.BytesValue, _Mapping]] = ...) -> None: ...
