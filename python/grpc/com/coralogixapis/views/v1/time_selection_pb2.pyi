from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class TimeSelection(_message.Message):
    __slots__ = ("quick_selection", "custom_selection")
    QUICK_SELECTION_FIELD_NUMBER: _ClassVar[int]
    CUSTOM_SELECTION_FIELD_NUMBER: _ClassVar[int]
    quick_selection: QuickTimeSelection
    custom_selection: CustomTimeSelection
    def __init__(self, quick_selection: _Optional[_Union[QuickTimeSelection, _Mapping]] = ..., custom_selection: _Optional[_Union[CustomTimeSelection, _Mapping]] = ...) -> None: ...

class QuickTimeSelection(_message.Message):
    __slots__ = ("caption", "seconds")
    CAPTION_FIELD_NUMBER: _ClassVar[int]
    SECONDS_FIELD_NUMBER: _ClassVar[int]
    caption: _wrappers_pb2.StringValue
    seconds: int
    def __init__(self, caption: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., seconds: _Optional[int] = ...) -> None: ...

class CustomTimeSelection(_message.Message):
    __slots__ = ("from_time", "to_time")
    FROM_TIME_FIELD_NUMBER: _ClassVar[int]
    TO_TIME_FIELD_NUMBER: _ClassVar[int]
    from_time: _timestamp_pb2.Timestamp
    to_time: _timestamp_pb2.Timestamp
    def __init__(self, from_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., to_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
