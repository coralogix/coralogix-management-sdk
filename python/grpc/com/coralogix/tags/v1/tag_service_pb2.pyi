from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.tags.v1 import tag_pb2 as _tag_pb2
from com.coralogix.tags.v1 import audit_log_pb2 as _audit_log_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateTagRequest(_message.Message):
    __slots__ = ("key", "name", "company_id", "icon_url", "timestamp", "application", "subsystem")
    KEY_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    COMPANY_ID_FIELD_NUMBER: _ClassVar[int]
    ICON_URL_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
    key: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    company_id: _wrappers_pb2.UInt32Value
    icon_url: _wrappers_pb2.StringValue
    timestamp: _timestamp_pb2.Timestamp
    application: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    subsystem: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., company_id: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., icon_url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., application: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystem: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class GetTagsRequest(_message.Message):
    __slots__ = ("query_def",)
    class QueryDef(_message.Message):
        __slots__ = ("start_date", "end_date", "page_index", "page_size", "query_params")
        class QueryParams(_message.Message):
            __slots__ = ("metadata", "sort_model")
            class Metadata(_message.Message):
                __slots__ = ("application_name", "subsystem_name")
                APPLICATION_NAME_FIELD_NUMBER: _ClassVar[int]
                SUBSYSTEM_NAME_FIELD_NUMBER: _ClassVar[int]
                application_name: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
                subsystem_name: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
                def __init__(self, application_name: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystem_name: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...
            class SortModel(_message.Message):
                __slots__ = ("field", "ordering", "missing")
                FIELD_FIELD_NUMBER: _ClassVar[int]
                ORDERING_FIELD_NUMBER: _ClassVar[int]
                MISSING_FIELD_NUMBER: _ClassVar[int]
                field: _wrappers_pb2.StringValue
                ordering: _wrappers_pb2.StringValue
                missing: _wrappers_pb2.StringValue
                def __init__(self, field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., ordering: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., missing: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
            METADATA_FIELD_NUMBER: _ClassVar[int]
            SORT_MODEL_FIELD_NUMBER: _ClassVar[int]
            metadata: GetTagsRequest.QueryDef.QueryParams.Metadata
            sort_model: _containers.RepeatedCompositeFieldContainer[GetTagsRequest.QueryDef.QueryParams.SortModel]
            def __init__(self, metadata: _Optional[_Union[GetTagsRequest.QueryDef.QueryParams.Metadata, _Mapping]] = ..., sort_model: _Optional[_Iterable[_Union[GetTagsRequest.QueryDef.QueryParams.SortModel, _Mapping]]] = ...) -> None: ...
        START_DATE_FIELD_NUMBER: _ClassVar[int]
        END_DATE_FIELD_NUMBER: _ClassVar[int]
        PAGE_INDEX_FIELD_NUMBER: _ClassVar[int]
        PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
        QUERY_PARAMS_FIELD_NUMBER: _ClassVar[int]
        start_date: _timestamp_pb2.Timestamp
        end_date: _timestamp_pb2.Timestamp
        page_index: _wrappers_pb2.UInt32Value
        page_size: _wrappers_pb2.UInt32Value
        query_params: GetTagsRequest.QueryDef.QueryParams
        def __init__(self, start_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., end_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., page_index: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., page_size: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., query_params: _Optional[_Union[GetTagsRequest.QueryDef.QueryParams, _Mapping]] = ...) -> None: ...
    QUERY_DEF_FIELD_NUMBER: _ClassVar[int]
    query_def: GetTagsRequest.QueryDef
    def __init__(self, query_def: _Optional[_Union[GetTagsRequest.QueryDef, _Mapping]] = ...) -> None: ...

class UpdateTagRequest(_message.Message):
    __slots__ = ("tag",)
    TAG_FIELD_NUMBER: _ClassVar[int]
    tag: _tag_pb2.Tag
    def __init__(self, tag: _Optional[_Union[_tag_pb2.Tag, _Mapping]] = ...) -> None: ...

class Query(_message.Message):
    __slots__ = ("application_name", "end_time", "start_time", "subsystem_name", "step")
    APPLICATION_NAME_FIELD_NUMBER: _ClassVar[int]
    END_TIME_FIELD_NUMBER: _ClassVar[int]
    START_TIME_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_NAME_FIELD_NUMBER: _ClassVar[int]
    STEP_FIELD_NUMBER: _ClassVar[int]
    application_name: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    end_time: _timestamp_pb2.Timestamp
    start_time: _timestamp_pb2.Timestamp
    subsystem_name: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    step: _wrappers_pb2.Int64Value
    def __init__(self, application_name: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., end_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., start_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., subsystem_name: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., step: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class DeleteTagRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.UInt64Value
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.UInt64Value, _Mapping]] = ...) -> None: ...

class GetTagSummaryRequest(_message.Message):
    __slots__ = ("compare_tag", "tag")
    COMPARE_TAG_FIELD_NUMBER: _ClassVar[int]
    TAG_FIELD_NUMBER: _ClassVar[int]
    compare_tag: Query
    tag: Query
    def __init__(self, compare_tag: _Optional[_Union[Query, _Mapping]] = ..., tag: _Optional[_Union[Query, _Mapping]] = ...) -> None: ...

class GetTagAlertsRequest(_message.Message):
    __slots__ = ("compare_tag", "tag")
    COMPARE_TAG_FIELD_NUMBER: _ClassVar[int]
    TAG_FIELD_NUMBER: _ClassVar[int]
    compare_tag: Query
    tag: Query
    def __init__(self, compare_tag: _Optional[_Union[Query, _Mapping]] = ..., tag: _Optional[_Union[Query, _Mapping]] = ...) -> None: ...

class GetTagErrorVolumeRequest(_message.Message):
    __slots__ = ("compare_tag", "tag")
    COMPARE_TAG_FIELD_NUMBER: _ClassVar[int]
    TAG_FIELD_NUMBER: _ClassVar[int]
    compare_tag: Query
    tag: Query
    def __init__(self, compare_tag: _Optional[_Union[Query, _Mapping]] = ..., tag: _Optional[_Union[Query, _Mapping]] = ...) -> None: ...

class CreateTagResponse(_message.Message):
    __slots__ = ("tag",)
    TAG_FIELD_NUMBER: _ClassVar[int]
    tag: _tag_pb2.Tag
    def __init__(self, tag: _Optional[_Union[_tag_pb2.Tag, _Mapping]] = ...) -> None: ...

class GetTagsResponse(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[_tag_pb2.Tag]
    def __init__(self, data: _Optional[_Iterable[_Union[_tag_pb2.Tag, _Mapping]]] = ...) -> None: ...

class UpdateTagResponse(_message.Message):
    __slots__ = ("tag",)
    TAG_FIELD_NUMBER: _ClassVar[int]
    tag: _tag_pb2.Tag
    def __init__(self, tag: _Optional[_Union[_tag_pb2.Tag, _Mapping]] = ...) -> None: ...

class DeleteTagResponse(_message.Message):
    __slots__ = ("response",)
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    response: _wrappers_pb2.StringValue
    def __init__(self, response: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetTagSummaryResponse(_message.Message):
    __slots__ = ("data",)
    class Summary(_message.Message):
        __slots__ = ("tag_result", "compare_result", "link_cache_id", "type")
        TAG_RESULT_FIELD_NUMBER: _ClassVar[int]
        COMPARE_RESULT_FIELD_NUMBER: _ClassVar[int]
        LINK_CACHE_ID_FIELD_NUMBER: _ClassVar[int]
        TYPE_FIELD_NUMBER: _ClassVar[int]
        tag_result: _wrappers_pb2.DoubleValue
        compare_result: _wrappers_pb2.DoubleValue
        link_cache_id: _wrappers_pb2.StringValue
        type: _wrappers_pb2.StringValue
        def __init__(self, tag_result: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ..., compare_result: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ..., link_cache_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., type: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[GetTagSummaryResponse.Summary]
    def __init__(self, data: _Optional[_Iterable[_Union[GetTagSummaryResponse.Summary, _Mapping]]] = ...) -> None: ...

class GraphPoint(_message.Message):
    __slots__ = ("timestamp", "value")
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    timestamp: _timestamp_pb2.Timestamp
    value: _wrappers_pb2.Int32Value
    def __init__(self, timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., value: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class GetTagAlertsResponse(_message.Message):
    __slots__ = ("alerts_list_info", "severity_info", "volume_compare_graph")
    class AlertInfo(_message.Message):
        __slots__ = ("name", "severity", "size")
        NAME_FIELD_NUMBER: _ClassVar[int]
        SEVERITY_FIELD_NUMBER: _ClassVar[int]
        SIZE_FIELD_NUMBER: _ClassVar[int]
        name: _wrappers_pb2.StringValue
        severity: _wrappers_pb2.Int32Value
        size: _wrappers_pb2.Int32Value
        def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., severity: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., size: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    class SeverityInfo(_message.Message):
        __slots__ = ("severity", "size")
        SEVERITY_FIELD_NUMBER: _ClassVar[int]
        SIZE_FIELD_NUMBER: _ClassVar[int]
        severity: _wrappers_pb2.Int32Value
        size: _wrappers_pb2.Int32Value
        def __init__(self, severity: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., size: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    class VolumeCompareGraph(_message.Message):
        __slots__ = ("compare_tag", "tag")
        COMPARE_TAG_FIELD_NUMBER: _ClassVar[int]
        TAG_FIELD_NUMBER: _ClassVar[int]
        compare_tag: _containers.RepeatedCompositeFieldContainer[GraphPoint]
        tag: _containers.RepeatedCompositeFieldContainer[GraphPoint]
        def __init__(self, compare_tag: _Optional[_Iterable[_Union[GraphPoint, _Mapping]]] = ..., tag: _Optional[_Iterable[_Union[GraphPoint, _Mapping]]] = ...) -> None: ...
    ALERTS_LIST_INFO_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_INFO_FIELD_NUMBER: _ClassVar[int]
    VOLUME_COMPARE_GRAPH_FIELD_NUMBER: _ClassVar[int]
    alerts_list_info: _containers.RepeatedCompositeFieldContainer[GetTagAlertsResponse.AlertInfo]
    severity_info: _containers.RepeatedCompositeFieldContainer[GetTagAlertsResponse.SeverityInfo]
    volume_compare_graph: GetTagAlertsResponse.VolumeCompareGraph
    def __init__(self, alerts_list_info: _Optional[_Iterable[_Union[GetTagAlertsResponse.AlertInfo, _Mapping]]] = ..., severity_info: _Optional[_Iterable[_Union[GetTagAlertsResponse.SeverityInfo, _Mapping]]] = ..., volume_compare_graph: _Optional[_Union[GetTagAlertsResponse.VolumeCompareGraph, _Mapping]] = ...) -> None: ...

class GetTagErrorVolumeResponse(_message.Message):
    __slots__ = ("compare_tag", "tag")
    COMPARE_TAG_FIELD_NUMBER: _ClassVar[int]
    TAG_FIELD_NUMBER: _ClassVar[int]
    compare_tag: _containers.RepeatedCompositeFieldContainer[GraphPoint]
    tag: _containers.RepeatedCompositeFieldContainer[GraphPoint]
    def __init__(self, compare_tag: _Optional[_Iterable[_Union[GraphPoint, _Mapping]]] = ..., tag: _Optional[_Iterable[_Union[GraphPoint, _Mapping]]] = ...) -> None: ...
