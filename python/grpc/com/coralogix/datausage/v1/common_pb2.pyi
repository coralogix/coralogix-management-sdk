from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class TcoTier(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    TCO_TIER_UNSPECIFIED: _ClassVar[TcoTier]
    TCO_TIER_LOW: _ClassVar[TcoTier]
    TCO_TIER_MEDIUM: _ClassVar[TcoTier]
    TCO_TIER_HIGH: _ClassVar[TcoTier]
    TCO_TIER_BLOCKED: _ClassVar[TcoTier]

class Priority(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    PRIORITY_UNSPECIFIED: _ClassVar[Priority]
    PRIORITY_LOW: _ClassVar[Priority]
    PRIORITY_MEDIUM: _ClassVar[Priority]
    PRIORITY_HIGH: _ClassVar[Priority]
    PRIORITY_BLOCKED: _ClassVar[Priority]

class DataType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    DATA_TYPE_UNSPECIFIED: _ClassVar[DataType]
    DATA_TYPE_LOGS: _ClassVar[DataType]
    DATA_TYPE_METRICS: _ClassVar[DataType]
    DATA_TYPE_TRACING: _ClassVar[DataType]

class Severity(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    SEVERITY_UNSPECIFIED: _ClassVar[Severity]
    SEVERITY_DEBUG: _ClassVar[Severity]
    SEVERITY_VERBOSE: _ClassVar[Severity]
    SEVERITY_INFO: _ClassVar[Severity]
    SEVERITY_WARNING: _ClassVar[Severity]
    SEVERITY_ERROR: _ClassVar[Severity]
    SEVERITY_CRITICAL: _ClassVar[Severity]

class Range(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    RANGE_UNSPECIFIED: _ClassVar[Range]
    RANGE_CURRENT_MONTH: _ClassVar[Range]
    RANGE_LAST_30_DAYS: _ClassVar[Range]
    RANGE_LAST_90_DAYS: _ClassVar[Range]
    RANGE_LAST_WEEK: _ClassVar[Range]

class Resolution(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    RESOLUTION_UNSPECIFIED: _ClassVar[Resolution]
    RESOLUTION_ONE_MINUTE: _ClassVar[Resolution]
    RESOLUTION_FIVE_MINUTES: _ClassVar[Resolution]
    RESOLUTION_FIFTEEN_MINUTES: _ClassVar[Resolution]
    RESOLUTION_ONE_HOUR: _ClassVar[Resolution]
    RESOLUTION_SIX_HOURS: _ClassVar[Resolution]
    RESOLUTION_ONE_DAY: _ClassVar[Resolution]
    RESOLUTION_ONE_WEEK: _ClassVar[Resolution]
TCO_TIER_UNSPECIFIED: TcoTier
TCO_TIER_LOW: TcoTier
TCO_TIER_MEDIUM: TcoTier
TCO_TIER_HIGH: TcoTier
TCO_TIER_BLOCKED: TcoTier
PRIORITY_UNSPECIFIED: Priority
PRIORITY_LOW: Priority
PRIORITY_MEDIUM: Priority
PRIORITY_HIGH: Priority
PRIORITY_BLOCKED: Priority
DATA_TYPE_UNSPECIFIED: DataType
DATA_TYPE_LOGS: DataType
DATA_TYPE_METRICS: DataType
DATA_TYPE_TRACING: DataType
SEVERITY_UNSPECIFIED: Severity
SEVERITY_DEBUG: Severity
SEVERITY_VERBOSE: Severity
SEVERITY_INFO: Severity
SEVERITY_WARNING: Severity
SEVERITY_ERROR: Severity
SEVERITY_CRITICAL: Severity
RANGE_UNSPECIFIED: Range
RANGE_CURRENT_MONTH: Range
RANGE_LAST_30_DAYS: Range
RANGE_LAST_90_DAYS: Range
RANGE_LAST_WEEK: Range
RESOLUTION_UNSPECIFIED: Resolution
RESOLUTION_ONE_MINUTE: Resolution
RESOLUTION_FIVE_MINUTES: Resolution
RESOLUTION_FIFTEEN_MINUTES: Resolution
RESOLUTION_ONE_HOUR: Resolution
RESOLUTION_SIX_HOURS: Resolution
RESOLUTION_ONE_DAY: Resolution
RESOLUTION_ONE_WEEK: Resolution

class Team(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.UInt64Value
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.UInt64Value, _Mapping]] = ...) -> None: ...

class Unit(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: _wrappers_pb2.FloatValue
    def __init__(self, value: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ...) -> None: ...

class GB(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: _wrappers_pb2.FloatValue
    def __init__(self, value: _Optional[_Union[_wrappers_pb2.FloatValue, _Mapping]] = ...) -> None: ...

class Retention(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: _wrappers_pb2.UInt64Value
    def __init__(self, value: _Optional[_Union[_wrappers_pb2.UInt64Value, _Mapping]] = ...) -> None: ...

class DateRange(_message.Message):
    __slots__ = ("from_date", "to_date")
    FROM_DATE_FIELD_NUMBER: _ClassVar[int]
    TO_DATE_FIELD_NUMBER: _ClassVar[int]
    from_date: _timestamp_pb2.Timestamp
    to_date: _timestamp_pb2.Timestamp
    def __init__(self, from_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., to_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
