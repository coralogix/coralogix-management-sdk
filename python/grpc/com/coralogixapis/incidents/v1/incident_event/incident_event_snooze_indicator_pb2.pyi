from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class IncidentEventSnoozeIndicator(_message.Message):
    __slots__ = ("start_time", "duration_minutes", "user_id")
    START_TIME_FIELD_NUMBER: _ClassVar[int]
    DURATION_MINUTES_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    start_time: _timestamp_pb2.Timestamp
    duration_minutes: _wrappers_pb2.Int32Value
    user_id: _wrappers_pb2.StringValue
    def __init__(self, start_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., duration_minutes: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., user_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
