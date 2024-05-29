from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AlertSnoozed(_message.Message):
    __slots__ = ("user_id", "until")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    UNTIL_FIELD_NUMBER: _ClassVar[int]
    user_id: _wrappers_pb2.StringValue
    until: _timestamp_pb2.Timestamp
    def __init__(self, user_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., until: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
