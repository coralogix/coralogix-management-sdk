from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UpsertIncidentStatePayload(_message.Message):
    __slots__ = ("cx_event_key",)
    CX_EVENT_KEY_FIELD_NUMBER: _ClassVar[int]
    cx_event_key: _wrappers_pb2.StringValue
    def __init__(self, cx_event_key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
