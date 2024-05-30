from com.coralogix.alerts.v1 import date_time_pb2 as _date_time_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DayOfWeek(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    DAY_OF_WEEK_MONDAY_OR_UNSPECIFIED: _ClassVar[DayOfWeek]
    DAY_OF_WEEK_TUESDAY: _ClassVar[DayOfWeek]
    DAY_OF_WEEK_WEDNESDAY: _ClassVar[DayOfWeek]
    DAY_OF_WEEK_THURSDAY: _ClassVar[DayOfWeek]
    DAY_OF_WEEK_FRIDAY: _ClassVar[DayOfWeek]
    DAY_OF_WEEK_SATURDAY: _ClassVar[DayOfWeek]
    DAY_OF_WEEK_SUNDAY: _ClassVar[DayOfWeek]
DAY_OF_WEEK_MONDAY_OR_UNSPECIFIED: DayOfWeek
DAY_OF_WEEK_TUESDAY: DayOfWeek
DAY_OF_WEEK_WEDNESDAY: DayOfWeek
DAY_OF_WEEK_THURSDAY: DayOfWeek
DAY_OF_WEEK_FRIDAY: DayOfWeek
DAY_OF_WEEK_SATURDAY: DayOfWeek
DAY_OF_WEEK_SUNDAY: DayOfWeek

class AlertActiveWhen(_message.Message):
    __slots__ = ("timeframes",)
    TIMEFRAMES_FIELD_NUMBER: _ClassVar[int]
    timeframes: _containers.RepeatedCompositeFieldContainer[AlertActiveTimeframe]
    def __init__(self, timeframes: _Optional[_Iterable[_Union[AlertActiveTimeframe, _Mapping]]] = ...) -> None: ...

class AlertActiveTimeframe(_message.Message):
    __slots__ = ("days_of_week", "range")
    DAYS_OF_WEEK_FIELD_NUMBER: _ClassVar[int]
    RANGE_FIELD_NUMBER: _ClassVar[int]
    days_of_week: _containers.RepeatedScalarFieldContainer[DayOfWeek]
    range: TimeRange
    def __init__(self, days_of_week: _Optional[_Iterable[_Union[DayOfWeek, str]]] = ..., range: _Optional[_Union[TimeRange, _Mapping]] = ...) -> None: ...

class TimeRange(_message.Message):
    __slots__ = ("start", "end")
    START_FIELD_NUMBER: _ClassVar[int]
    END_FIELD_NUMBER: _ClassVar[int]
    start: _date_time_pb2.Time
    end: _date_time_pb2.Time
    def __init__(self, start: _Optional[_Union[_date_time_pb2.Time, _Mapping]] = ..., end: _Optional[_Union[_date_time_pb2.Time, _Mapping]] = ...) -> None: ...
