from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Pillar(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    PILLAR_UNSPECIFIED: _ClassVar[Pillar]
    PILLAR_METRICS: _ClassVar[Pillar]
    PILLAR_LOGS: _ClassVar[Pillar]
    PILLAR_SPANS: _ClassVar[Pillar]
    PILLAR_BINARY: _ClassVar[Pillar]

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

class Severity(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    SEVERITY_UNSPECIFIED: _ClassVar[Severity]
    SEVERITY_DEBUG: _ClassVar[Severity]
    SEVERITY_VERBOSE: _ClassVar[Severity]
    SEVERITY_INFO: _ClassVar[Severity]
    SEVERITY_WARNING: _ClassVar[Severity]
    SEVERITY_ERROR: _ClassVar[Severity]
    SEVERITY_CRITICAL: _ClassVar[Severity]
PILLAR_UNSPECIFIED: Pillar
PILLAR_METRICS: Pillar
PILLAR_LOGS: Pillar
PILLAR_SPANS: Pillar
PILLAR_BINARY: Pillar
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
SEVERITY_UNSPECIFIED: Severity
SEVERITY_DEBUG: Severity
SEVERITY_VERBOSE: Severity
SEVERITY_INFO: Severity
SEVERITY_WARNING: Severity
SEVERITY_ERROR: Severity
SEVERITY_CRITICAL: Severity

class Team(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.UInt64Value
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.UInt64Value, _Mapping]] = ...) -> None: ...

class DateRange(_message.Message):
    __slots__ = ("from_date", "to_date")
    FROM_DATE_FIELD_NUMBER: _ClassVar[int]
    TO_DATE_FIELD_NUMBER: _ClassVar[int]
    from_date: _timestamp_pb2.Timestamp
    to_date: _timestamp_pb2.Timestamp
    def __init__(self, from_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., to_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class GenericDimension(_message.Message):
    __slots__ = ("key", "value")
    KEY_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    key: str
    value: str
    def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...

class Dimension(_message.Message):
    __slots__ = ("pillar", "generic_dimension", "tier", "severity", "priority")
    PILLAR_FIELD_NUMBER: _ClassVar[int]
    GENERIC_DIMENSION_FIELD_NUMBER: _ClassVar[int]
    TIER_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    PRIORITY_FIELD_NUMBER: _ClassVar[int]
    pillar: Pillar
    generic_dimension: GenericDimension
    tier: TcoTier
    severity: Severity
    priority: Priority
    def __init__(self, pillar: _Optional[_Union[Pillar, str]] = ..., generic_dimension: _Optional[_Union[GenericDimension, _Mapping]] = ..., tier: _Optional[_Union[TcoTier, str]] = ..., severity: _Optional[_Union[Severity, str]] = ..., priority: _Optional[_Union[Priority, str]] = ...) -> None: ...
