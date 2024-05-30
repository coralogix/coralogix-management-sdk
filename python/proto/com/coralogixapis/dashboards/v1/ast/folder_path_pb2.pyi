from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class FolderPath(_message.Message):
    __slots__ = ("segments",)
    SEGMENTS_FIELD_NUMBER: _ClassVar[int]
    segments: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, segments: _Optional[_Iterable[str]] = ...) -> None: ...
