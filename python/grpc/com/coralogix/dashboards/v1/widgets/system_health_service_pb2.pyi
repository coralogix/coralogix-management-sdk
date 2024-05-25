from com.coralogix.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.dashboards.v1.common import search_filter_pb2 as _search_filter_pb2
from com.coralogix.dashboards.v1.common import time_frame_pb2 as _time_frame_pb2
from com.coralogix.dashboards.v1.common import time_series_pb2 as _time_series_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SystemHealthStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    SYSTEM_HEALTH_STATUS_UNSPECIFIED: _ClassVar[SystemHealthStatus]
    SYSTEM_HEALTH_STATUS_OK: _ClassVar[SystemHealthStatus]
    SYSTEM_HEALTH_STATUS_NO_DATA: _ClassVar[SystemHealthStatus]
    SYSTEM_HEALTH_STATUS_MISSING: _ClassVar[SystemHealthStatus]
SYSTEM_HEALTH_STATUS_UNSPECIFIED: SystemHealthStatus
SYSTEM_HEALTH_STATUS_OK: SystemHealthStatus
SYSTEM_HEALTH_STATUS_NO_DATA: SystemHealthStatus
SYSTEM_HEALTH_STATUS_MISSING: SystemHealthStatus

class GetSystemHealthRequest(_message.Message):
    __slots__ = ("filters", "time_frame", "graph_size")
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    GRAPH_SIZE_FIELD_NUMBER: _ClassVar[int]
    filters: _containers.RepeatedCompositeFieldContainer[_search_filter_pb2.SearchFilter]
    time_frame: _time_frame_pb2.TimeFrame
    graph_size: _wrappers_pb2.Int32Value
    def __init__(self, filters: _Optional[_Iterable[_Union[_search_filter_pb2.SearchFilter, _Mapping]]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., graph_size: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class GetSystemHealthResponse(_message.Message):
    __slots__ = ("status", "time_series")
    STATUS_FIELD_NUMBER: _ClassVar[int]
    TIME_SERIES_FIELD_NUMBER: _ClassVar[int]
    status: SystemHealthStatus
    time_series: _time_series_pb2.TimeSeries
    def __init__(self, status: _Optional[_Union[SystemHealthStatus, str]] = ..., time_series: _Optional[_Union[_time_series_pb2.TimeSeries, _Mapping]] = ...) -> None: ...
