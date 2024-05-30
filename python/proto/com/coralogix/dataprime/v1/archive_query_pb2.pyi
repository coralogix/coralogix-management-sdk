from com.coralogix.dataprime.v1 import compile_pb2 as _compile_pb2
from com.coralogix.dataprime.v1 import query_pb2 as _query_pb2
from com.coralogix.dataprime.v1 import warnings_pb2 as _warnings_pb2
from com.coralogix.dataprime.v1 import archive_log_pb2 as _archive_log_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class RunQueryRequest(_message.Message):
    __slots__ = ("name", "description", "syntax", "query", "raw_query", "from_date", "to_date", "application_filters", "severity_filters", "subsystem_filters", "widget_id", "request_params_hash")
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    SYNTAX_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    RAW_QUERY_FIELD_NUMBER: _ClassVar[int]
    FROM_DATE_FIELD_NUMBER: _ClassVar[int]
    TO_DATE_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_FILTERS_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FILTERS_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FILTERS_FIELD_NUMBER: _ClassVar[int]
    WIDGET_ID_FIELD_NUMBER: _ClassVar[int]
    REQUEST_PARAMS_HASH_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    syntax: _compile_pb2.QuerySyntax
    query: _query_pb2.QueryPayload
    raw_query: _wrappers_pb2.StringValue
    from_date: _timestamp_pb2.Timestamp
    to_date: _timestamp_pb2.Timestamp
    application_filters: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    severity_filters: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    subsystem_filters: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    widget_id: _wrappers_pb2.StringValue
    request_params_hash: _wrappers_pb2.StringValue
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., syntax: _Optional[_Union[_compile_pb2.QuerySyntax, str]] = ..., query: _Optional[_Union[_query_pb2.QueryPayload, _Mapping]] = ..., raw_query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., from_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., to_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., application_filters: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., severity_filters: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystem_filters: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., widget_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., request_params_hash: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class RunQueryResponse(_message.Message):
    __slots__ = ("archive_query",)
    ARCHIVE_QUERY_FIELD_NUMBER: _ClassVar[int]
    archive_query: ArchiveQuery
    def __init__(self, archive_query: _Optional[_Union[ArchiveQuery, _Mapping]] = ...) -> None: ...

class ListQueriesRequest(_message.Message):
    __slots__ = ("page", "size", "filter", "request_params_hash", "ordering", "only_dashboard_queries")
    class Ordering(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        ORDERING_DESC_UNSPECIFIED: _ClassVar[ListQueriesRequest.Ordering]
        ORDERING_ASC: _ClassVar[ListQueriesRequest.Ordering]
    ORDERING_DESC_UNSPECIFIED: ListQueriesRequest.Ordering
    ORDERING_ASC: ListQueriesRequest.Ordering
    PAGE_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    FILTER_FIELD_NUMBER: _ClassVar[int]
    REQUEST_PARAMS_HASH_FIELD_NUMBER: _ClassVar[int]
    ORDERING_FIELD_NUMBER: _ClassVar[int]
    ONLY_DASHBOARD_QUERIES_FIELD_NUMBER: _ClassVar[int]
    page: int
    size: int
    filter: _wrappers_pb2.StringValue
    request_params_hash: _wrappers_pb2.StringValue
    ordering: ListQueriesRequest.Ordering
    only_dashboard_queries: _wrappers_pb2.BoolValue
    def __init__(self, page: _Optional[int] = ..., size: _Optional[int] = ..., filter: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., request_params_hash: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., ordering: _Optional[_Union[ListQueriesRequest.Ordering, str]] = ..., only_dashboard_queries: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class ListQueriesResponse(_message.Message):
    __slots__ = ("archive_queries", "total_count", "request_params_hash")
    ARCHIVE_QUERIES_FIELD_NUMBER: _ClassVar[int]
    TOTAL_COUNT_FIELD_NUMBER: _ClassVar[int]
    REQUEST_PARAMS_HASH_FIELD_NUMBER: _ClassVar[int]
    archive_queries: _containers.RepeatedCompositeFieldContainer[ArchiveQuery]
    total_count: int
    request_params_hash: _wrappers_pb2.StringValue
    def __init__(self, archive_queries: _Optional[_Iterable[_Union[ArchiveQuery, _Mapping]]] = ..., total_count: _Optional[int] = ..., request_params_hash: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class ListQueryNamesRequest(_message.Message):
    __slots__ = ("limit", "contains_filter")
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    CONTAINS_FILTER_FIELD_NUMBER: _ClassVar[int]
    limit: int
    contains_filter: _wrappers_pb2.StringValue
    def __init__(self, limit: _Optional[int] = ..., contains_filter: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class ListQueryNamesResponse(_message.Message):
    __slots__ = ("query_names",)
    QUERY_NAMES_FIELD_NUMBER: _ClassVar[int]
    query_names: _containers.RepeatedCompositeFieldContainer[QueryName]
    def __init__(self, query_names: _Optional[_Iterable[_Union[QueryName, _Mapping]]] = ...) -> None: ...

class QueryName(_message.Message):
    __slots__ = ("query_name",)
    QUERY_NAME_FIELD_NUMBER: _ClassVar[int]
    query_name: _wrappers_pb2.StringValue
    def __init__(self, query_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class QueryResultRequest(_message.Message):
    __slots__ = ("archive_query_id", "query", "page", "size", "format")
    class Format(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        FORMAT_JSON_UNSPECIFIED: _ClassVar[QueryResultRequest.Format]
        FORMAT_CSV: _ClassVar[QueryResultRequest.Format]
    FORMAT_JSON_UNSPECIFIED: QueryResultRequest.Format
    FORMAT_CSV: QueryResultRequest.Format
    ARCHIVE_QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    PAGE_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    FORMAT_FIELD_NUMBER: _ClassVar[int]
    archive_query_id: _wrappers_pb2.StringValue
    query: _query_pb2.QueryPayload
    page: int
    size: int
    format: QueryResultRequest.Format
    def __init__(self, archive_query_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., query: _Optional[_Union[_query_pb2.QueryPayload, _Mapping]] = ..., page: _Optional[int] = ..., size: _Optional[int] = ..., format: _Optional[_Union[QueryResultRequest.Format, str]] = ...) -> None: ...

class QueryResultResponse(_message.Message):
    __slots__ = ("log",)
    LOG_FIELD_NUMBER: _ClassVar[int]
    log: _archive_log_pb2.ArchiveLog
    def __init__(self, log: _Optional[_Union[_archive_log_pb2.ArchiveLog, _Mapping]] = ...) -> None: ...

class ArchiveQuery(_message.Message):
    __slots__ = ("id", "name", "description", "syntax", "query", "from_date", "to_date", "application_filters", "severity_filters", "subsystem_filters", "status", "created_by", "started_at", "completed_at", "expires_at", "status_failed_description", "rows_total_count", "widget_id", "scan_limit_reached", "blocks_limit_reached", "archive_warning")
    class Status(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        STATUS_PENDING_UNSPECIFIED: _ClassVar[ArchiveQuery.Status]
        STATUS_IN_PROGRESS: _ClassVar[ArchiveQuery.Status]
        STATUS_COMPLETED: _ClassVar[ArchiveQuery.Status]
        STATUS_FAILED: _ClassVar[ArchiveQuery.Status]
    STATUS_PENDING_UNSPECIFIED: ArchiveQuery.Status
    STATUS_IN_PROGRESS: ArchiveQuery.Status
    STATUS_COMPLETED: ArchiveQuery.Status
    STATUS_FAILED: ArchiveQuery.Status
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    SYNTAX_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    FROM_DATE_FIELD_NUMBER: _ClassVar[int]
    TO_DATE_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_FILTERS_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FILTERS_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FILTERS_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    CREATED_BY_FIELD_NUMBER: _ClassVar[int]
    STARTED_AT_FIELD_NUMBER: _ClassVar[int]
    COMPLETED_AT_FIELD_NUMBER: _ClassVar[int]
    EXPIRES_AT_FIELD_NUMBER: _ClassVar[int]
    STATUS_FAILED_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    ROWS_TOTAL_COUNT_FIELD_NUMBER: _ClassVar[int]
    WIDGET_ID_FIELD_NUMBER: _ClassVar[int]
    SCAN_LIMIT_REACHED_FIELD_NUMBER: _ClassVar[int]
    BLOCKS_LIMIT_REACHED_FIELD_NUMBER: _ClassVar[int]
    ARCHIVE_WARNING_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    syntax: _compile_pb2.QuerySyntax
    query: _wrappers_pb2.StringValue
    from_date: _timestamp_pb2.Timestamp
    to_date: _timestamp_pb2.Timestamp
    application_filters: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    severity_filters: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    subsystem_filters: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    status: ArchiveQuery.Status
    created_by: _wrappers_pb2.StringValue
    started_at: _timestamp_pb2.Timestamp
    completed_at: _timestamp_pb2.Timestamp
    expires_at: _timestamp_pb2.Timestamp
    status_failed_description: _wrappers_pb2.StringValue
    rows_total_count: int
    widget_id: _wrappers_pb2.StringValue
    scan_limit_reached: bool
    blocks_limit_reached: bool
    archive_warning: _warnings_pb2.ArchiveWarning
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., syntax: _Optional[_Union[_compile_pb2.QuerySyntax, str]] = ..., query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., from_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., to_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., application_filters: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., severity_filters: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystem_filters: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., status: _Optional[_Union[ArchiveQuery.Status, str]] = ..., created_by: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., started_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., completed_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., expires_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., status_failed_description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., rows_total_count: _Optional[int] = ..., widget_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., scan_limit_reached: bool = ..., blocks_limit_reached: bool = ..., archive_warning: _Optional[_Union[_warnings_pb2.ArchiveWarning, _Mapping]] = ...) -> None: ...
