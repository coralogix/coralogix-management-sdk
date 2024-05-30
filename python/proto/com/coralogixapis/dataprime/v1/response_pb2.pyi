from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class QueryResponse(_message.Message):
    __slots__ = ("error", "result", "warning", "query_id")
    ERROR_FIELD_NUMBER: _ClassVar[int]
    RESULT_FIELD_NUMBER: _ClassVar[int]
    WARNING_FIELD_NUMBER: _ClassVar[int]
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    error: DataprimeError
    result: DataprimeResult
    warning: DataprimeWarning
    query_id: QueryId
    def __init__(self, error: _Optional[_Union[DataprimeError, _Mapping]] = ..., result: _Optional[_Union[DataprimeResult, _Mapping]] = ..., warning: _Optional[_Union[DataprimeWarning, _Mapping]] = ..., query_id: _Optional[_Union[QueryId, _Mapping]] = ...) -> None: ...

class DataprimeWarning(_message.Message):
    __slots__ = ("compile_warning", "time_range_warning", "number_of_results_limit_warning", "bytes_scanned_limit_warning", "deprecation_warning", "blocks_limit_warning", "aggregation_buckets_limit_warning", "archive_warning")
    COMPILE_WARNING_FIELD_NUMBER: _ClassVar[int]
    TIME_RANGE_WARNING_FIELD_NUMBER: _ClassVar[int]
    NUMBER_OF_RESULTS_LIMIT_WARNING_FIELD_NUMBER: _ClassVar[int]
    BYTES_SCANNED_LIMIT_WARNING_FIELD_NUMBER: _ClassVar[int]
    DEPRECATION_WARNING_FIELD_NUMBER: _ClassVar[int]
    BLOCKS_LIMIT_WARNING_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_BUCKETS_LIMIT_WARNING_FIELD_NUMBER: _ClassVar[int]
    ARCHIVE_WARNING_FIELD_NUMBER: _ClassVar[int]
    compile_warning: CompileWarning
    time_range_warning: TimeRangeWarning
    number_of_results_limit_warning: NumberOfResultsLimitWarning
    bytes_scanned_limit_warning: BytesScannedLimitWarning
    deprecation_warning: DeprecationWarning
    blocks_limit_warning: BlocksLimitWarning
    aggregation_buckets_limit_warning: AggregationBucketsLimitWarning
    archive_warning: ArchiveWarning
    def __init__(self, compile_warning: _Optional[_Union[CompileWarning, _Mapping]] = ..., time_range_warning: _Optional[_Union[TimeRangeWarning, _Mapping]] = ..., number_of_results_limit_warning: _Optional[_Union[NumberOfResultsLimitWarning, _Mapping]] = ..., bytes_scanned_limit_warning: _Optional[_Union[BytesScannedLimitWarning, _Mapping]] = ..., deprecation_warning: _Optional[_Union[DeprecationWarning, _Mapping]] = ..., blocks_limit_warning: _Optional[_Union[BlocksLimitWarning, _Mapping]] = ..., aggregation_buckets_limit_warning: _Optional[_Union[AggregationBucketsLimitWarning, _Mapping]] = ..., archive_warning: _Optional[_Union[ArchiveWarning, _Mapping]] = ...) -> None: ...

class CompileWarning(_message.Message):
    __slots__ = ("warning_message",)
    WARNING_MESSAGE_FIELD_NUMBER: _ClassVar[int]
    warning_message: str
    def __init__(self, warning_message: _Optional[str] = ...) -> None: ...

class TimeRangeWarning(_message.Message):
    __slots__ = ("warning_message", "start_date", "end_date")
    WARNING_MESSAGE_FIELD_NUMBER: _ClassVar[int]
    START_DATE_FIELD_NUMBER: _ClassVar[int]
    END_DATE_FIELD_NUMBER: _ClassVar[int]
    warning_message: str
    start_date: _timestamp_pb2.Timestamp
    end_date: _timestamp_pb2.Timestamp
    def __init__(self, warning_message: _Optional[str] = ..., start_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., end_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class NumberOfResultsLimitWarning(_message.Message):
    __slots__ = ("number_of_results_limit",)
    NUMBER_OF_RESULTS_LIMIT_FIELD_NUMBER: _ClassVar[int]
    number_of_results_limit: int
    def __init__(self, number_of_results_limit: _Optional[int] = ...) -> None: ...

class BytesScannedLimitWarning(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class AggregationBucketsLimitWarning(_message.Message):
    __slots__ = ("aggregation_buckets_limit",)
    AGGREGATION_BUCKETS_LIMIT_FIELD_NUMBER: _ClassVar[int]
    aggregation_buckets_limit: int
    def __init__(self, aggregation_buckets_limit: _Optional[int] = ...) -> None: ...

class DeprecationWarning(_message.Message):
    __slots__ = ("warning_message",)
    WARNING_MESSAGE_FIELD_NUMBER: _ClassVar[int]
    warning_message: str
    def __init__(self, warning_message: _Optional[str] = ...) -> None: ...

class BlocksLimitWarning(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DataprimeResult(_message.Message):
    __slots__ = ("results",)
    RESULTS_FIELD_NUMBER: _ClassVar[int]
    results: _containers.RepeatedCompositeFieldContainer[DataprimeResults]
    def __init__(self, results: _Optional[_Iterable[_Union[DataprimeResults, _Mapping]]] = ...) -> None: ...

class DataprimeResults(_message.Message):
    __slots__ = ("metadata", "labels", "user_data")
    class KeyValue(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    METADATA_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    USER_DATA_FIELD_NUMBER: _ClassVar[int]
    metadata: _containers.RepeatedCompositeFieldContainer[DataprimeResults.KeyValue]
    labels: _containers.RepeatedCompositeFieldContainer[DataprimeResults.KeyValue]
    user_data: str
    def __init__(self, metadata: _Optional[_Iterable[_Union[DataprimeResults.KeyValue, _Mapping]]] = ..., labels: _Optional[_Iterable[_Union[DataprimeResults.KeyValue, _Mapping]]] = ..., user_data: _Optional[str] = ...) -> None: ...

class DataprimeError(_message.Message):
    __slots__ = ("message", "code")
    class RateLimitReached(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class Code(_message.Message):
        __slots__ = ("rate_limit_reached",)
        RATE_LIMIT_REACHED_FIELD_NUMBER: _ClassVar[int]
        rate_limit_reached: DataprimeError.RateLimitReached
        def __init__(self, rate_limit_reached: _Optional[_Union[DataprimeError.RateLimitReached, _Mapping]] = ...) -> None: ...
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    message: _wrappers_pb2.StringValue
    code: DataprimeError.Code
    def __init__(self, message: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., code: _Optional[_Union[DataprimeError.Code, _Mapping]] = ...) -> None: ...

class SerializedDataprime(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: bytes
    def __init__(self, data: _Optional[bytes] = ...) -> None: ...

class QueryId(_message.Message):
    __slots__ = ("query_id",)
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    query_id: str
    def __init__(self, query_id: _Optional[str] = ...) -> None: ...

class ArchiveWarning(_message.Message):
    __slots__ = ("no_metastore_data", "bucket_access_denied", "bucket_read_failed", "missing_data")
    class NoMetastoreData(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class BucketAccessDenied(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class BucketReadFailed(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class MissingData(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    NO_METASTORE_DATA_FIELD_NUMBER: _ClassVar[int]
    BUCKET_ACCESS_DENIED_FIELD_NUMBER: _ClassVar[int]
    BUCKET_READ_FAILED_FIELD_NUMBER: _ClassVar[int]
    MISSING_DATA_FIELD_NUMBER: _ClassVar[int]
    no_metastore_data: ArchiveWarning.NoMetastoreData
    bucket_access_denied: ArchiveWarning.BucketAccessDenied
    bucket_read_failed: ArchiveWarning.BucketReadFailed
    missing_data: ArchiveWarning.MissingData
    def __init__(self, no_metastore_data: _Optional[_Union[ArchiveWarning.NoMetastoreData, _Mapping]] = ..., bucket_access_denied: _Optional[_Union[ArchiveWarning.BucketAccessDenied, _Mapping]] = ..., bucket_read_failed: _Optional[_Union[ArchiveWarning.BucketReadFailed, _Mapping]] = ..., missing_data: _Optional[_Union[ArchiveWarning.MissingData, _Mapping]] = ...) -> None: ...
