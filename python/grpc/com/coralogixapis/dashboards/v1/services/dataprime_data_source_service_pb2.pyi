from com.coralogixapis.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogixapis.dashboards.v1.common import dataprime_result_pb2 as _dataprime_result_pb2
from com.coralogixapis.dashboards.v1.common import query_pb2 as _query_pb2
from com.coralogixapis.dashboards.v1.common import time_frame_pb2 as _time_frame_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SearchDataprimeRequest(_message.Message):
    __slots__ = ("dataprime_query", "dataprime_query_raw", "time_frame", "limit")
    DATAPRIME_QUERY_FIELD_NUMBER: _ClassVar[int]
    DATAPRIME_QUERY_RAW_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    dataprime_query: _query_pb2.SerializedDataprimeQuery
    dataprime_query_raw: _query_pb2.DataprimeQuery
    time_frame: _time_frame_pb2.TimeFrame
    limit: _wrappers_pb2.Int32Value
    def __init__(self, dataprime_query: _Optional[_Union[_query_pb2.SerializedDataprimeQuery, _Mapping]] = ..., dataprime_query_raw: _Optional[_Union[_query_pb2.DataprimeQuery, _Mapping]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class SearchDataprimeResponse(_message.Message):
    __slots__ = ("results",)
    RESULTS_FIELD_NUMBER: _ClassVar[int]
    results: _containers.RepeatedCompositeFieldContainer[_dataprime_result_pb2.DataprimeResult]
    def __init__(self, results: _Optional[_Iterable[_Union[_dataprime_result_pb2.DataprimeResult, _Mapping]]] = ...) -> None: ...

class SearchDataprimeArchiveRequest(_message.Message):
    __slots__ = ("dataprime_query", "dataprime_query_raw", "time_frame", "limit", "widget_id")
    DATAPRIME_QUERY_FIELD_NUMBER: _ClassVar[int]
    DATAPRIME_QUERY_RAW_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    WIDGET_ID_FIELD_NUMBER: _ClassVar[int]
    dataprime_query: _query_pb2.SerializedDataprimeQuery
    dataprime_query_raw: _query_pb2.DataprimeQuery
    time_frame: _time_frame_pb2.TimeFrame
    limit: _wrappers_pb2.Int32Value
    widget_id: _wrappers_pb2.StringValue
    def __init__(self, dataprime_query: _Optional[_Union[_query_pb2.SerializedDataprimeQuery, _Mapping]] = ..., dataprime_query_raw: _Optional[_Union[_query_pb2.DataprimeQuery, _Mapping]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., widget_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class SearchDataprimeArchiveResponse(_message.Message):
    __slots__ = ("results",)
    RESULTS_FIELD_NUMBER: _ClassVar[int]
    results: _containers.RepeatedCompositeFieldContainer[_dataprime_result_pb2.DataprimeResult]
    def __init__(self, results: _Optional[_Iterable[_Union[_dataprime_result_pb2.DataprimeResult, _Mapping]]] = ...) -> None: ...
