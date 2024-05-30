from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Markdown(_message.Message):
    __slots__ = ("markdown_text", "tooltip_text")
    MARKDOWN_TEXT_FIELD_NUMBER: _ClassVar[int]
    TOOLTIP_TEXT_FIELD_NUMBER: _ClassVar[int]
    markdown_text: _wrappers_pb2.StringValue
    tooltip_text: _wrappers_pb2.StringValue
    def __init__(self, markdown_text: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., tooltip_text: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
