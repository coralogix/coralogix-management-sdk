from com.coralogix.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.dashboards.v1.common import search_filter_pb2 as _search_filter_pb2
from com.coralogix.dashboards.v1.common import time_frame_pb2 as _time_frame_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AnomalyGroupType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ANOMALY_GROUP_TYPE_UNSPECIFIED: _ClassVar[AnomalyGroupType]
    ANOMALY_GROUP_TYPE_VOLUME: _ClassVar[AnomalyGroupType]
    ANOMALY_GROUP_TYPE_FLOW: _ClassVar[AnomalyGroupType]
ANOMALY_GROUP_TYPE_UNSPECIFIED: AnomalyGroupType
ANOMALY_GROUP_TYPE_VOLUME: AnomalyGroupType
ANOMALY_GROUP_TYPE_FLOW: AnomalyGroupType

class SearchAnomaliesGroupedByTypeRequest(_message.Message):
    __slots__ = ("filters", "time_frame")
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    filters: _containers.RepeatedCompositeFieldContainer[_search_filter_pb2.SearchFilter]
    time_frame: _time_frame_pb2.TimeFrame
    def __init__(self, filters: _Optional[_Iterable[_Union[_search_filter_pb2.SearchFilter, _Mapping]]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ...) -> None: ...

class SearchAnomaliesGroupedByTypeResponse(_message.Message):
    __slots__ = ("anomalies",)
    ANOMALIES_FIELD_NUMBER: _ClassVar[int]
    anomalies: _containers.RepeatedCompositeFieldContainer[AnomalyGroupByType]
    def __init__(self, anomalies: _Optional[_Iterable[_Union[AnomalyGroupByType, _Mapping]]] = ...) -> None: ...

class AnomalyGroupByType(_message.Message):
    __slots__ = ("type", "amount")
    TYPE_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    type: AnomalyGroupType
    amount: _wrappers_pb2.Int32Value
    def __init__(self, type: _Optional[_Union[AnomalyGroupType, str]] = ..., amount: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
