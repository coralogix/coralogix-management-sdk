from com.coralogix.dataprime.ast.v1 import ast_pb2 as _ast_pb2
from com.coralogix.dataprime.ast.v1 import query_pb2 as _query_pb2
from com.coralogix.dataprime.v1 import object_store_warning_pb2 as _object_store_warning_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Metastore(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    METASTORE_UNSPECIFIED: _ClassVar[Metastore]
    METASTORE_SCYLLA: _ClassVar[Metastore]
    METASTORE_POSTGRES: _ClassVar[Metastore]
METASTORE_UNSPECIFIED: Metastore
METASTORE_SCYLLA: Metastore
METASTORE_POSTGRES: Metastore

class SubmitQueryRequest(_message.Message):
    __slots__ = ("query_ast", "execution_config", "query_id", "named_query_sources")
    QUERY_AST_FIELD_NUMBER: _ClassVar[int]
    EXECUTION_CONFIG_FIELD_NUMBER: _ClassVar[int]
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    NAMED_QUERY_SOURCES_FIELD_NUMBER: _ClassVar[int]
    query_ast: _ast_pb2.Ast
    execution_config: ExecutionConfig
    query_id: _wrappers_pb2.StringValue
    named_query_sources: _containers.RepeatedCompositeFieldContainer[NamedQuerySource]
    def __init__(self, query_ast: _Optional[_Union[_ast_pb2.Ast, _Mapping]] = ..., execution_config: _Optional[_Union[ExecutionConfig, _Mapping]] = ..., query_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., named_query_sources: _Optional[_Iterable[_Union[NamedQuerySource, _Mapping]]] = ...) -> None: ...

class SubmitDdlQueryRequest(_message.Message):
    __slots__ = ("ddl_ast", "execution_config", "query_id", "dataset_location", "named_query_sources")
    DDL_AST_FIELD_NUMBER: _ClassVar[int]
    EXECUTION_CONFIG_FIELD_NUMBER: _ClassVar[int]
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    DATASET_LOCATION_FIELD_NUMBER: _ClassVar[int]
    NAMED_QUERY_SOURCES_FIELD_NUMBER: _ClassVar[int]
    ddl_ast: _ast_pb2.DdlAst
    execution_config: ExecutionConfig
    query_id: _wrappers_pb2.StringValue
    dataset_location: DatasetLocation
    named_query_sources: _containers.RepeatedCompositeFieldContainer[NamedQuerySource]
    def __init__(self, ddl_ast: _Optional[_Union[_ast_pb2.DdlAst, _Mapping]] = ..., execution_config: _Optional[_Union[ExecutionConfig, _Mapping]] = ..., query_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., dataset_location: _Optional[_Union[DatasetLocation, _Mapping]] = ..., named_query_sources: _Optional[_Iterable[_Union[NamedQuerySource, _Mapping]]] = ...) -> None: ...

class DatasetLocation(_message.Message):
    __slots__ = ("object_store_url", "path", "region")
    OBJECT_STORE_URL_FIELD_NUMBER: _ClassVar[int]
    PATH_FIELD_NUMBER: _ClassVar[int]
    REGION_FIELD_NUMBER: _ClassVar[int]
    object_store_url: str
    path: str
    region: str
    def __init__(self, object_store_url: _Optional[str] = ..., path: _Optional[str] = ..., region: _Optional[str] = ...) -> None: ...

class ExplainQueryRequest(_message.Message):
    __slots__ = ("query_ast", "execution_config", "query_id", "named_query_sources")
    QUERY_AST_FIELD_NUMBER: _ClassVar[int]
    EXECUTION_CONFIG_FIELD_NUMBER: _ClassVar[int]
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    NAMED_QUERY_SOURCES_FIELD_NUMBER: _ClassVar[int]
    query_ast: _ast_pb2.Ast
    execution_config: ExecutionConfig
    query_id: _wrappers_pb2.StringValue
    named_query_sources: _containers.RepeatedCompositeFieldContainer[NamedQuerySource]
    def __init__(self, query_ast: _Optional[_Union[_ast_pb2.Ast, _Mapping]] = ..., execution_config: _Optional[_Union[ExecutionConfig, _Mapping]] = ..., query_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., named_query_sources: _Optional[_Iterable[_Union[NamedQuerySource, _Mapping]]] = ...) -> None: ...

class ExplainStage(_message.Message):
    __slots__ = ("stage", "partitions", "plan")
    STAGE_FIELD_NUMBER: _ClassVar[int]
    PARTITIONS_FIELD_NUMBER: _ClassVar[int]
    PLAN_FIELD_NUMBER: _ClassVar[int]
    stage: int
    partitions: int
    plan: str
    def __init__(self, stage: _Optional[int] = ..., partitions: _Optional[int] = ..., plan: _Optional[str] = ...) -> None: ...

class ExplainQueryResponse(_message.Message):
    __slots__ = ("initial_logical_plan", "optimized_logical_plan", "physical_plan", "stages", "full_explanation")
    INITIAL_LOGICAL_PLAN_FIELD_NUMBER: _ClassVar[int]
    OPTIMIZED_LOGICAL_PLAN_FIELD_NUMBER: _ClassVar[int]
    PHYSICAL_PLAN_FIELD_NUMBER: _ClassVar[int]
    STAGES_FIELD_NUMBER: _ClassVar[int]
    FULL_EXPLANATION_FIELD_NUMBER: _ClassVar[int]
    initial_logical_plan: str
    optimized_logical_plan: str
    physical_plan: str
    stages: _containers.RepeatedCompositeFieldContainer[ExplainStage]
    full_explanation: str
    def __init__(self, initial_logical_plan: _Optional[str] = ..., optimized_logical_plan: _Optional[str] = ..., physical_plan: _Optional[str] = ..., stages: _Optional[_Iterable[_Union[ExplainStage, _Mapping]]] = ..., full_explanation: _Optional[str] = ...) -> None: ...

class QuerySource(_message.Message):
    __slots__ = ("archive_source", "open_search_source", "custom_enrichment_csv", "async_query_source", "schema_store_source")
    class ArchiveSource(_message.Message):
        __slots__ = ("bucket_name", "region_name", "max_archive_version", "bucket_v2")
        BUCKET_NAME_FIELD_NUMBER: _ClassVar[int]
        REGION_NAME_FIELD_NUMBER: _ClassVar[int]
        MAX_ARCHIVE_VERSION_FIELD_NUMBER: _ClassVar[int]
        BUCKET_V2_FIELD_NUMBER: _ClassVar[int]
        bucket_name: str
        region_name: str
        max_archive_version: int
        bucket_v2: _wrappers_pb2.StringValue
        def __init__(self, bucket_name: _Optional[str] = ..., region_name: _Optional[str] = ..., max_archive_version: _Optional[int] = ..., bucket_v2: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class OpenSearchSource(_message.Message):
        __slots__ = ("index",)
        INDEX_FIELD_NUMBER: _ClassVar[int]
        index: str
        def __init__(self, index: _Optional[str] = ...) -> None: ...
    class CustomEnrichmentCsv(_message.Message):
        __slots__ = ("bucket_name", "region_name")
        BUCKET_NAME_FIELD_NUMBER: _ClassVar[int]
        REGION_NAME_FIELD_NUMBER: _ClassVar[int]
        bucket_name: _wrappers_pb2.StringValue
        region_name: _wrappers_pb2.StringValue
        def __init__(self, bucket_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., region_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class AsyncQuerySource(_message.Message):
        __slots__ = ("archive_source",)
        ARCHIVE_SOURCE_FIELD_NUMBER: _ClassVar[int]
        archive_source: QuerySource.ArchiveSource
        def __init__(self, archive_source: _Optional[_Union[QuerySource.ArchiveSource, _Mapping]] = ...) -> None: ...
    class SchemaStoreSource(_message.Message):
        __slots__ = ("team_ids",)
        TEAM_IDS_FIELD_NUMBER: _ClassVar[int]
        team_ids: _containers.RepeatedScalarFieldContainer[str]
        def __init__(self, team_ids: _Optional[_Iterable[str]] = ...) -> None: ...
    ARCHIVE_SOURCE_FIELD_NUMBER: _ClassVar[int]
    OPEN_SEARCH_SOURCE_FIELD_NUMBER: _ClassVar[int]
    CUSTOM_ENRICHMENT_CSV_FIELD_NUMBER: _ClassVar[int]
    ASYNC_QUERY_SOURCE_FIELD_NUMBER: _ClassVar[int]
    SCHEMA_STORE_SOURCE_FIELD_NUMBER: _ClassVar[int]
    archive_source: QuerySource.ArchiveSource
    open_search_source: QuerySource.OpenSearchSource
    custom_enrichment_csv: QuerySource.CustomEnrichmentCsv
    async_query_source: QuerySource.AsyncQuerySource
    schema_store_source: QuerySource.SchemaStoreSource
    def __init__(self, archive_source: _Optional[_Union[QuerySource.ArchiveSource, _Mapping]] = ..., open_search_source: _Optional[_Union[QuerySource.OpenSearchSource, _Mapping]] = ..., custom_enrichment_csv: _Optional[_Union[QuerySource.CustomEnrichmentCsv, _Mapping]] = ..., async_query_source: _Optional[_Union[QuerySource.AsyncQuerySource, _Mapping]] = ..., schema_store_source: _Optional[_Union[QuerySource.SchemaStoreSource, _Mapping]] = ...) -> None: ...

class NamedQuerySource(_message.Message):
    __slots__ = ("name", "source")
    NAME_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    name: _query_pb2.Source
    source: QuerySource
    def __init__(self, name: _Optional[_Union[_query_pb2.Source, _Mapping]] = ..., source: _Optional[_Union[QuerySource, _Mapping]] = ...) -> None: ...

class SubmitQueryResponse(_message.Message):
    __slots__ = ("query_id",)
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    query_id: str
    def __init__(self, query_id: _Optional[str] = ...) -> None: ...

class SubmitDdlQueryResponse(_message.Message):
    __slots__ = ("query_id",)
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    query_id: str
    def __init__(self, query_id: _Optional[str] = ...) -> None: ...

class GetDatasetRequest(_message.Message):
    __slots__ = ("query_ast", "execution_config", "query_id", "named_query_source", "skip", "limit", "json", "csv")
    class Json(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class Csv(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    QUERY_AST_FIELD_NUMBER: _ClassVar[int]
    EXECUTION_CONFIG_FIELD_NUMBER: _ClassVar[int]
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    NAMED_QUERY_SOURCE_FIELD_NUMBER: _ClassVar[int]
    SKIP_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    JSON_FIELD_NUMBER: _ClassVar[int]
    CSV_FIELD_NUMBER: _ClassVar[int]
    query_ast: _ast_pb2.Ast
    execution_config: ExecutionConfig
    query_id: _wrappers_pb2.StringValue
    named_query_source: NamedQuerySource
    skip: int
    limit: int
    json: GetDatasetRequest.Json
    csv: GetDatasetRequest.Csv
    def __init__(self, query_ast: _Optional[_Union[_ast_pb2.Ast, _Mapping]] = ..., execution_config: _Optional[_Union[ExecutionConfig, _Mapping]] = ..., query_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., named_query_source: _Optional[_Union[NamedQuerySource, _Mapping]] = ..., skip: _Optional[int] = ..., limit: _Optional[int] = ..., json: _Optional[_Union[GetDatasetRequest.Json, _Mapping]] = ..., csv: _Optional[_Union[GetDatasetRequest.Csv, _Mapping]] = ...) -> None: ...

class GetDatasetResponse(_message.Message):
    __slots__ = ("row",)
    ROW_FIELD_NUMBER: _ClassVar[int]
    row: str
    def __init__(self, row: _Optional[str] = ...) -> None: ...

class DropDatasetRequest(_message.Message):
    __slots__ = ("dataset_location", "dataset_name")
    DATASET_LOCATION_FIELD_NUMBER: _ClassVar[int]
    DATASET_NAME_FIELD_NUMBER: _ClassVar[int]
    dataset_location: DatasetLocation
    dataset_name: str
    def __init__(self, dataset_location: _Optional[_Union[DatasetLocation, _Mapping]] = ..., dataset_name: _Optional[str] = ...) -> None: ...

class DropDatasetResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ExecutionConfig(_message.Message):
    __slots__ = ("target_partitions", "max_bytes_per_partition", "metastore", "use_page_bloom_filter", "page_index_path", "timeout_sec", "use_column_statistics", "use_text_index", "max_bytes_read_from_s3", "blocks_limit", "use_correct_sorting", "open_search_agg_buckets_size", "open_search_agg_max_pages", "optimize_jsona_parsing")
    TARGET_PARTITIONS_FIELD_NUMBER: _ClassVar[int]
    MAX_BYTES_PER_PARTITION_FIELD_NUMBER: _ClassVar[int]
    METASTORE_FIELD_NUMBER: _ClassVar[int]
    USE_PAGE_BLOOM_FILTER_FIELD_NUMBER: _ClassVar[int]
    PAGE_INDEX_PATH_FIELD_NUMBER: _ClassVar[int]
    TIMEOUT_SEC_FIELD_NUMBER: _ClassVar[int]
    USE_COLUMN_STATISTICS_FIELD_NUMBER: _ClassVar[int]
    USE_TEXT_INDEX_FIELD_NUMBER: _ClassVar[int]
    MAX_BYTES_READ_FROM_S3_FIELD_NUMBER: _ClassVar[int]
    BLOCKS_LIMIT_FIELD_NUMBER: _ClassVar[int]
    USE_CORRECT_SORTING_FIELD_NUMBER: _ClassVar[int]
    OPEN_SEARCH_AGG_BUCKETS_SIZE_FIELD_NUMBER: _ClassVar[int]
    OPEN_SEARCH_AGG_MAX_PAGES_FIELD_NUMBER: _ClassVar[int]
    OPTIMIZE_JSONA_PARSING_FIELD_NUMBER: _ClassVar[int]
    target_partitions: int
    max_bytes_per_partition: int
    metastore: Metastore
    use_page_bloom_filter: bool
    page_index_path: str
    timeout_sec: int
    use_column_statistics: bool
    use_text_index: bool
    max_bytes_read_from_s3: int
    blocks_limit: int
    use_correct_sorting: bool
    open_search_agg_buckets_size: int
    open_search_agg_max_pages: int
    optimize_jsona_parsing: bool
    def __init__(self, target_partitions: _Optional[int] = ..., max_bytes_per_partition: _Optional[int] = ..., metastore: _Optional[_Union[Metastore, str]] = ..., use_page_bloom_filter: bool = ..., page_index_path: _Optional[str] = ..., timeout_sec: _Optional[int] = ..., use_column_statistics: bool = ..., use_text_index: bool = ..., max_bytes_read_from_s3: _Optional[int] = ..., blocks_limit: _Optional[int] = ..., use_correct_sorting: bool = ..., open_search_agg_buckets_size: _Optional[int] = ..., open_search_agg_max_pages: _Optional[int] = ..., optimize_jsona_parsing: bool = ...) -> None: ...

class AwaitReadyRequest(_message.Message):
    __slots__ = ("query_id",)
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    query_id: str
    def __init__(self, query_id: _Optional[str] = ...) -> None: ...

class Failed(_message.Message):
    __slots__ = ("internal", "external", "cancelled")
    class Internal(_message.Message):
        __slots__ = ("message",)
        MESSAGE_FIELD_NUMBER: _ClassVar[int]
        message: str
        def __init__(self, message: _Optional[str] = ...) -> None: ...
    class External(_message.Message):
        __slots__ = ("message",)
        MESSAGE_FIELD_NUMBER: _ClassVar[int]
        message: str
        def __init__(self, message: _Optional[str] = ...) -> None: ...
    class Cancelled(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    INTERNAL_FIELD_NUMBER: _ClassVar[int]
    EXTERNAL_FIELD_NUMBER: _ClassVar[int]
    CANCELLED_FIELD_NUMBER: _ClassVar[int]
    internal: Failed.Internal
    external: Failed.External
    cancelled: Failed.Cancelled
    def __init__(self, internal: _Optional[_Union[Failed.Internal, _Mapping]] = ..., external: _Optional[_Union[Failed.External, _Mapping]] = ..., cancelled: _Optional[_Union[Failed.Cancelled, _Mapping]] = ...) -> None: ...

class AwaitReadyResponse(_message.Message):
    __slots__ = ("success", "failed")
    class Success(_message.Message):
        __slots__ = ("scan_limit_reached", "blocks_limit_reached", "object_store_warning")
        SCAN_LIMIT_REACHED_FIELD_NUMBER: _ClassVar[int]
        BLOCKS_LIMIT_REACHED_FIELD_NUMBER: _ClassVar[int]
        OBJECT_STORE_WARNING_FIELD_NUMBER: _ClassVar[int]
        scan_limit_reached: bool
        blocks_limit_reached: bool
        object_store_warning: _object_store_warning_pb2.ObjectStoreWarning
        def __init__(self, scan_limit_reached: bool = ..., blocks_limit_reached: bool = ..., object_store_warning: _Optional[_Union[_object_store_warning_pb2.ObjectStoreWarning, _Mapping]] = ...) -> None: ...
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    FAILED_FIELD_NUMBER: _ClassVar[int]
    success: AwaitReadyResponse.Success
    failed: Failed
    def __init__(self, success: _Optional[_Union[AwaitReadyResponse.Success, _Mapping]] = ..., failed: _Optional[_Union[Failed, _Mapping]] = ...) -> None: ...

class GetQueryStatusRequest(_message.Message):
    __slots__ = ("query_id",)
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    query_id: str
    def __init__(self, query_id: _Optional[str] = ...) -> None: ...

class QueryStatus(_message.Message):
    __slots__ = ("queued", "running", "failed", "completed", "planning", "not_found", "cancelled", "query_id")
    class Queued(_message.Message):
        __slots__ = ("queued_at",)
        QUEUED_AT_FIELD_NUMBER: _ClassVar[int]
        queued_at: int
        def __init__(self, queued_at: _Optional[int] = ...) -> None: ...
    class Running(_message.Message):
        __slots__ = ("queued_at",)
        QUEUED_AT_FIELD_NUMBER: _ClassVar[int]
        queued_at: int
        def __init__(self, queued_at: _Optional[int] = ...) -> None: ...
    class Completed(_message.Message):
        __slots__ = ("queued_at", "scan_limit_reached", "ended_at", "num_rows", "blocks_limit_reached", "object_store_warning")
        QUEUED_AT_FIELD_NUMBER: _ClassVar[int]
        SCAN_LIMIT_REACHED_FIELD_NUMBER: _ClassVar[int]
        ENDED_AT_FIELD_NUMBER: _ClassVar[int]
        NUM_ROWS_FIELD_NUMBER: _ClassVar[int]
        BLOCKS_LIMIT_REACHED_FIELD_NUMBER: _ClassVar[int]
        OBJECT_STORE_WARNING_FIELD_NUMBER: _ClassVar[int]
        queued_at: int
        scan_limit_reached: bool
        ended_at: int
        num_rows: int
        blocks_limit_reached: bool
        object_store_warning: _object_store_warning_pb2.ObjectStoreWarning
        def __init__(self, queued_at: _Optional[int] = ..., scan_limit_reached: bool = ..., ended_at: _Optional[int] = ..., num_rows: _Optional[int] = ..., blocks_limit_reached: bool = ..., object_store_warning: _Optional[_Union[_object_store_warning_pb2.ObjectStoreWarning, _Mapping]] = ...) -> None: ...
    class Planning(_message.Message):
        __slots__ = ("queued_at",)
        QUEUED_AT_FIELD_NUMBER: _ClassVar[int]
        queued_at: int
        def __init__(self, queued_at: _Optional[int] = ...) -> None: ...
    class NotFound(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class Cancelled(_message.Message):
        __slots__ = ("queued_at", "ended_at")
        QUEUED_AT_FIELD_NUMBER: _ClassVar[int]
        ENDED_AT_FIELD_NUMBER: _ClassVar[int]
        queued_at: int
        ended_at: int
        def __init__(self, queued_at: _Optional[int] = ..., ended_at: _Optional[int] = ...) -> None: ...
    QUEUED_FIELD_NUMBER: _ClassVar[int]
    RUNNING_FIELD_NUMBER: _ClassVar[int]
    FAILED_FIELD_NUMBER: _ClassVar[int]
    COMPLETED_FIELD_NUMBER: _ClassVar[int]
    PLANNING_FIELD_NUMBER: _ClassVar[int]
    NOT_FOUND_FIELD_NUMBER: _ClassVar[int]
    CANCELLED_FIELD_NUMBER: _ClassVar[int]
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    queued: QueryStatus.Queued
    running: QueryStatus.Running
    failed: Failed
    completed: QueryStatus.Completed
    planning: QueryStatus.Planning
    not_found: QueryStatus.NotFound
    cancelled: QueryStatus.Cancelled
    query_id: str
    def __init__(self, queued: _Optional[_Union[QueryStatus.Queued, _Mapping]] = ..., running: _Optional[_Union[QueryStatus.Running, _Mapping]] = ..., failed: _Optional[_Union[Failed, _Mapping]] = ..., completed: _Optional[_Union[QueryStatus.Completed, _Mapping]] = ..., planning: _Optional[_Union[QueryStatus.Planning, _Mapping]] = ..., not_found: _Optional[_Union[QueryStatus.NotFound, _Mapping]] = ..., cancelled: _Optional[_Union[QueryStatus.Cancelled, _Mapping]] = ..., query_id: _Optional[str] = ...) -> None: ...

class GetQueryStatusResponse(_message.Message):
    __slots__ = ("status",)
    STATUS_FIELD_NUMBER: _ClassVar[int]
    status: QueryStatus
    def __init__(self, status: _Optional[_Union[QueryStatus, _Mapping]] = ...) -> None: ...

class CancelQueryRequest(_message.Message):
    __slots__ = ("query_id",)
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    query_id: str
    def __init__(self, query_id: _Optional[str] = ...) -> None: ...

class CancelQueryResponse(_message.Message):
    __slots__ = ("cancelled",)
    CANCELLED_FIELD_NUMBER: _ClassVar[int]
    cancelled: bool
    def __init__(self, cancelled: bool = ...) -> None: ...

class QueryPayload(_message.Message):
    __slots__ = ("dataprime",)
    DATAPRIME_FIELD_NUMBER: _ClassVar[int]
    dataprime: SerializedDataprime
    def __init__(self, dataprime: _Optional[_Union[SerializedDataprime, _Mapping]] = ...) -> None: ...

class SerializedDataprime(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: bytes
    def __init__(self, data: _Optional[bytes] = ...) -> None: ...
