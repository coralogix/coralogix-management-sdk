from com.coralogix.datausage.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.datausage.v1 import daily_usage_pb2 as _daily_usage_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.api import annotations_pb2 as _annotations_pb2
from com.coralogix.datausage.v1 import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DailyUsageRequest(_message.Message):
    __slots__ = ("range", "date_range")
    RANGE_FIELD_NUMBER: _ClassVar[int]
    DATE_RANGE_FIELD_NUMBER: _ClassVar[int]
    range: _common_pb2.Range
    date_range: _common_pb2.DateRange
    def __init__(self, range: _Optional[_Union[_common_pb2.Range, str]] = ..., date_range: _Optional[_Union[_common_pb2.DateRange, _Mapping]] = ...) -> None: ...

class DailyUsageResponse(_message.Message):
    __slots__ = ("response",)
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    response: _containers.RepeatedCompositeFieldContainer[_daily_usage_pb2.DetailedDailyUsage]
    def __init__(self, response: _Optional[_Iterable[_Union[_daily_usage_pb2.DetailedDailyUsage, _Mapping]]] = ...) -> None: ...

class OverageDailyUsageRequest(_message.Message):
    __slots__ = ("range", "date_range")
    RANGE_FIELD_NUMBER: _ClassVar[int]
    DATE_RANGE_FIELD_NUMBER: _ClassVar[int]
    range: _common_pb2.Range
    date_range: _common_pb2.DateRange
    def __init__(self, range: _Optional[_Union[_common_pb2.Range, str]] = ..., date_range: _Optional[_Union[_common_pb2.DateRange, _Mapping]] = ...) -> None: ...

class OverageDailyUsageResponse(_message.Message):
    __slots__ = ("overage_detail",)
    OVERAGE_DETAIL_FIELD_NUMBER: _ClassVar[int]
    overage_detail: _containers.RepeatedCompositeFieldContainer[_daily_usage_pb2.OverageDailyReport]
    def __init__(self, overage_detail: _Optional[_Iterable[_Union[_daily_usage_pb2.OverageDailyReport, _Mapping]]] = ...) -> None: ...
