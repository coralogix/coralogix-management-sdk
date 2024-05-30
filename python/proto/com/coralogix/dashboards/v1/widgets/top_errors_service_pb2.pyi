from com.coralogix.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.dashboards.v1.common import search_filter_pb2 as _search_filter_pb2
from com.coralogix.dashboards.v1.common import time_frame_pb2 as _time_frame_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class FetchTopErrorsRequest(_message.Message):
    __slots__ = ("filters", "time_frame", "aggregate_by")
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    AGGREGATE_BY_FIELD_NUMBER: _ClassVar[int]
    filters: _containers.RepeatedCompositeFieldContainer[_search_filter_pb2.SearchFilter]
    time_frame: _time_frame_pb2.TimeFrame
    aggregate_by: _wrappers_pb2.StringValue
    def __init__(self, filters: _Optional[_Iterable[_Union[_search_filter_pb2.SearchFilter, _Mapping]]] = ..., time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., aggregate_by: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class FetchTopErrorsResponse(_message.Message):
    __slots__ = ("total_error_count", "errors_by_group")
    class GroupErrorCount(_message.Message):
        __slots__ = ("application", "subsystem", "count", "errors_by_template", "no_template_count")
        APPLICATION_FIELD_NUMBER: _ClassVar[int]
        SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
        COUNT_FIELD_NUMBER: _ClassVar[int]
        ERRORS_BY_TEMPLATE_FIELD_NUMBER: _ClassVar[int]
        NO_TEMPLATE_COUNT_FIELD_NUMBER: _ClassVar[int]
        application: _wrappers_pb2.StringValue
        subsystem: _wrappers_pb2.StringValue
        count: _wrappers_pb2.Int64Value
        errors_by_template: _containers.RepeatedCompositeFieldContainer[FetchTopErrorsResponse.TemplateErrorCount]
        no_template_count: _wrappers_pb2.Int64Value
        def __init__(self, application: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., errors_by_template: _Optional[_Iterable[_Union[FetchTopErrorsResponse.TemplateErrorCount, _Mapping]]] = ..., no_template_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...
    class TemplateErrorCount(_message.Message):
        __slots__ = ("template_id", "count", "log_example", "template_created_at")
        TEMPLATE_ID_FIELD_NUMBER: _ClassVar[int]
        COUNT_FIELD_NUMBER: _ClassVar[int]
        LOG_EXAMPLE_FIELD_NUMBER: _ClassVar[int]
        TEMPLATE_CREATED_AT_FIELD_NUMBER: _ClassVar[int]
        template_id: _wrappers_pb2.StringValue
        count: _wrappers_pb2.Int64Value
        log_example: _wrappers_pb2.StringValue
        template_created_at: _timestamp_pb2.Timestamp
        def __init__(self, template_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., log_example: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., template_created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
    TOTAL_ERROR_COUNT_FIELD_NUMBER: _ClassVar[int]
    ERRORS_BY_GROUP_FIELD_NUMBER: _ClassVar[int]
    total_error_count: _wrappers_pb2.Int64Value
    errors_by_group: _containers.RepeatedCompositeFieldContainer[FetchTopErrorsResponse.GroupErrorCount]
    def __init__(self, total_error_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., errors_by_group: _Optional[_Iterable[_Union[FetchTopErrorsResponse.GroupErrorCount, _Mapping]]] = ...) -> None: ...
