from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Date(_message.Message):
    __slots__ = ("year", "month", "day")
    YEAR_FIELD_NUMBER: _ClassVar[int]
    MONTH_FIELD_NUMBER: _ClassVar[int]
    DAY_FIELD_NUMBER: _ClassVar[int]
    year: int
    month: int
    day: int
    def __init__(self, year: _Optional[int] = ..., month: _Optional[int] = ..., day: _Optional[int] = ...) -> None: ...

class Time(_message.Message):
    __slots__ = ("hours", "minutes", "seconds")
    HOURS_FIELD_NUMBER: _ClassVar[int]
    MINUTES_FIELD_NUMBER: _ClassVar[int]
    SECONDS_FIELD_NUMBER: _ClassVar[int]
    hours: int
    minutes: int
    seconds: int
    def __init__(self, hours: _Optional[int] = ..., minutes: _Optional[int] = ..., seconds: _Optional[int] = ...) -> None: ...
