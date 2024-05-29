from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from com.coralogix.dataprime.v1 import common_pb2 as _common_pb2
from com.coralogix.dataprime.v1 import histogram_pb2 as _histogram_pb2
from com.coralogix.dataprime.v1 import run_error_pb2 as _run_error_pb2
from com.coralogix.dataprime.v1 import warnings_pb2 as _warnings_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class RunDataprimeQueryResponse(_message.Message):
    __slots__ = ("error", "log_result", "severity_histogram", "count", "key_distribution", "histogram", "compared_histogram", "sidebar_filters", "request_statistics", "team_id_counts")
    ERROR_FIELD_NUMBER: _ClassVar[int]
    LOG_RESULT_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_HISTOGRAM_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    KEY_DISTRIBUTION_FIELD_NUMBER: _ClassVar[int]
    HISTOGRAM_FIELD_NUMBER: _ClassVar[int]
    COMPARED_HISTOGRAM_FIELD_NUMBER: _ClassVar[int]
    SIDEBAR_FILTERS_FIELD_NUMBER: _ClassVar[int]
    REQUEST_STATISTICS_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_COUNTS_FIELD_NUMBER: _ClassVar[int]
    error: _run_error_pb2.Error
    log_result: LogResult
    severity_histogram: SeverityHistogram
    count: Count
    key_distribution: KeyDistribution
    histogram: Histogram
    compared_histogram: Histogram
    sidebar_filters: SideBarFilters
    request_statistics: RequestStatistics
    team_id_counts: TeamIdCounts
    def __init__(self, error: _Optional[_Union[_run_error_pb2.Error, _Mapping]] = ..., log_result: _Optional[_Union[LogResult, _Mapping]] = ..., severity_histogram: _Optional[_Union[SeverityHistogram, _Mapping]] = ..., count: _Optional[_Union[Count, _Mapping]] = ..., key_distribution: _Optional[_Union[KeyDistribution, _Mapping]] = ..., histogram: _Optional[_Union[Histogram, _Mapping]] = ..., compared_histogram: _Optional[_Union[Histogram, _Mapping]] = ..., sidebar_filters: _Optional[_Union[SideBarFilters, _Mapping]] = ..., request_statistics: _Optional[_Union[RequestStatistics, _Mapping]] = ..., team_id_counts: _Optional[_Union[TeamIdCounts, _Mapping]] = ...) -> None: ...

class LogResult(_message.Message):
    __slots__ = ("results", "warnings", "skipped")
    RESULTS_FIELD_NUMBER: _ClassVar[int]
    WARNINGS_FIELD_NUMBER: _ClassVar[int]
    SKIPPED_FIELD_NUMBER: _ClassVar[int]
    results: _containers.RepeatedCompositeFieldContainer[LogEntry]
    warnings: _containers.RepeatedCompositeFieldContainer[_warnings_pb2.DataprimeWarning]
    skipped: _wrappers_pb2.Int64Value
    def __init__(self, results: _Optional[_Iterable[_Union[LogEntry, _Mapping]]] = ..., warnings: _Optional[_Iterable[_Union[_warnings_pb2.DataprimeWarning, _Mapping]]] = ..., skipped: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class SeverityCounts(_message.Message):
    __slots__ = ("debug_count", "info_count", "warning_count", "error_count", "critical_count", "verbose_count", "warnings")
    DEBUG_COUNT_FIELD_NUMBER: _ClassVar[int]
    INFO_COUNT_FIELD_NUMBER: _ClassVar[int]
    WARNING_COUNT_FIELD_NUMBER: _ClassVar[int]
    ERROR_COUNT_FIELD_NUMBER: _ClassVar[int]
    CRITICAL_COUNT_FIELD_NUMBER: _ClassVar[int]
    VERBOSE_COUNT_FIELD_NUMBER: _ClassVar[int]
    WARNINGS_FIELD_NUMBER: _ClassVar[int]
    debug_count: _wrappers_pb2.Int64Value
    info_count: _wrappers_pb2.Int64Value
    warning_count: _wrappers_pb2.Int64Value
    error_count: _wrappers_pb2.Int64Value
    critical_count: _wrappers_pb2.Int64Value
    verbose_count: _wrappers_pb2.Int64Value
    warnings: _containers.RepeatedCompositeFieldContainer[_warnings_pb2.DataprimeWarning]
    def __init__(self, debug_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., info_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., warning_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., error_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., critical_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., verbose_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., warnings: _Optional[_Iterable[_Union[_warnings_pb2.DataprimeWarning, _Mapping]]] = ...) -> None: ...

class SeverityHistogram(_message.Message):
    __slots__ = ("buckets", "warnings")
    class Bucket(_message.Message):
        __slots__ = ("bucket", "severity_counts")
        BUCKET_FIELD_NUMBER: _ClassVar[int]
        SEVERITY_COUNTS_FIELD_NUMBER: _ClassVar[int]
        bucket: _timestamp_pb2.Timestamp
        severity_counts: SeverityCounts
        def __init__(self, bucket: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., severity_counts: _Optional[_Union[SeverityCounts, _Mapping]] = ...) -> None: ...
    BUCKETS_FIELD_NUMBER: _ClassVar[int]
    WARNINGS_FIELD_NUMBER: _ClassVar[int]
    buckets: _containers.RepeatedCompositeFieldContainer[SeverityHistogram.Bucket]
    warnings: _containers.RepeatedCompositeFieldContainer[_warnings_pb2.DataprimeWarning]
    def __init__(self, buckets: _Optional[_Iterable[_Union[SeverityHistogram.Bucket, _Mapping]]] = ..., warnings: _Optional[_Iterable[_Union[_warnings_pb2.DataprimeWarning, _Mapping]]] = ...) -> None: ...

class DefaultTermCounts(_message.Message):
    __slots__ = ("term_counts", "warnings")
    class TermCount(_message.Message):
        __slots__ = ("term", "count")
        TERM_FIELD_NUMBER: _ClassVar[int]
        COUNT_FIELD_NUMBER: _ClassVar[int]
        term: _wrappers_pb2.StringValue
        count: _wrappers_pb2.Int64Value
        def __init__(self, term: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...
    TERM_COUNTS_FIELD_NUMBER: _ClassVar[int]
    WARNINGS_FIELD_NUMBER: _ClassVar[int]
    term_counts: _containers.RepeatedCompositeFieldContainer[DefaultTermCounts.TermCount]
    warnings: _containers.RepeatedCompositeFieldContainer[_warnings_pb2.DataprimeWarning]
    def __init__(self, term_counts: _Optional[_Iterable[_Union[DefaultTermCounts.TermCount, _Mapping]]] = ..., warnings: _Optional[_Iterable[_Union[_warnings_pb2.DataprimeWarning, _Mapping]]] = ...) -> None: ...

class SideBarFilters(_message.Message):
    __slots__ = ("sidebar_filter", "warnings")
    SIDEBAR_FILTER_FIELD_NUMBER: _ClassVar[int]
    WARNINGS_FIELD_NUMBER: _ClassVar[int]
    sidebar_filter: _containers.RepeatedCompositeFieldContainer[SideBarFilter]
    warnings: _containers.RepeatedCompositeFieldContainer[_warnings_pb2.DataprimeWarning]
    def __init__(self, sidebar_filter: _Optional[_Iterable[_Union[SideBarFilter, _Mapping]]] = ..., warnings: _Optional[_Iterable[_Union[_warnings_pb2.DataprimeWarning, _Mapping]]] = ...) -> None: ...

class SideBarFilter(_message.Message):
    __slots__ = ("filter_name", "term")
    FILTER_NAME_FIELD_NUMBER: _ClassVar[int]
    TERM_FIELD_NUMBER: _ClassVar[int]
    filter_name: _wrappers_pb2.StringValue
    term: DefaultTermCounts
    def __init__(self, filter_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., term: _Optional[_Union[DefaultTermCounts, _Mapping]] = ...) -> None: ...

class Count(_message.Message):
    __slots__ = ("total_count", "result_count", "warnings")
    TOTAL_COUNT_FIELD_NUMBER: _ClassVar[int]
    RESULT_COUNT_FIELD_NUMBER: _ClassVar[int]
    WARNINGS_FIELD_NUMBER: _ClassVar[int]
    total_count: _wrappers_pb2.Int64Value
    result_count: _wrappers_pb2.Int64Value
    warnings: _containers.RepeatedCompositeFieldContainer[_warnings_pb2.DataprimeWarning]
    def __init__(self, total_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., result_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., warnings: _Optional[_Iterable[_Union[_warnings_pb2.DataprimeWarning, _Mapping]]] = ...) -> None: ...

class KeyDistribution(_message.Message):
    __slots__ = ("key", "distribution", "warnings")
    class DistributionEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: int
        def __init__(self, key: _Optional[str] = ..., value: _Optional[int] = ...) -> None: ...
    KEY_FIELD_NUMBER: _ClassVar[int]
    DISTRIBUTION_FIELD_NUMBER: _ClassVar[int]
    WARNINGS_FIELD_NUMBER: _ClassVar[int]
    key: str
    distribution: _containers.ScalarMap[str, int]
    warnings: _containers.RepeatedCompositeFieldContainer[_warnings_pb2.DataprimeWarning]
    def __init__(self, key: _Optional[str] = ..., distribution: _Optional[_Mapping[str, int]] = ..., warnings: _Optional[_Iterable[_Union[_warnings_pb2.DataprimeWarning, _Mapping]]] = ...) -> None: ...

class Histogram(_message.Message):
    __slots__ = ("values", "timestamps", "warnings")
    VALUES_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMPS_FIELD_NUMBER: _ClassVar[int]
    WARNINGS_FIELD_NUMBER: _ClassVar[int]
    values: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    timestamps: _containers.RepeatedCompositeFieldContainer[_histogram_pb2.HistogramSlice]
    warnings: _containers.RepeatedCompositeFieldContainer[_warnings_pb2.DataprimeWarning]
    def __init__(self, values: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., timestamps: _Optional[_Iterable[_Union[_histogram_pb2.HistogramSlice, _Mapping]]] = ..., warnings: _Optional[_Iterable[_Union[_warnings_pb2.DataprimeWarning, _Mapping]]] = ...) -> None: ...

class RequestStatistics(_message.Message):
    __slots__ = ("bytes_scanned",)
    BYTES_SCANNED_FIELD_NUMBER: _ClassVar[int]
    bytes_scanned: _wrappers_pb2.Int64Value
    def __init__(self, bytes_scanned: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class TeamIdCounts(_message.Message):
    __slots__ = ("team_id_counts",)
    TEAM_ID_COUNTS_FIELD_NUMBER: _ClassVar[int]
    team_id_counts: _containers.RepeatedCompositeFieldContainer[TeamIdCount]
    def __init__(self, team_id_counts: _Optional[_Iterable[_Union[TeamIdCount, _Mapping]]] = ...) -> None: ...

class TeamIdCount(_message.Message):
    __slots__ = ("term", "count")
    TERM_FIELD_NUMBER: _ClassVar[int]
    COUNT_FIELD_NUMBER: _ClassVar[int]
    term: _wrappers_pb2.StringValue
    count: _wrappers_pb2.Int64Value
    def __init__(self, term: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ...) -> None: ...

class LogEntry(_message.Message):
    __slots__ = ("index", "labels", "metadata", "data")
    INDEX_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    index: _wrappers_pb2.Int64Value
    labels: _containers.RepeatedCompositeFieldContainer[_common_pb2.KeyValue]
    metadata: _containers.RepeatedCompositeFieldContainer[_common_pb2.KeyValue]
    data: str
    def __init__(self, index: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., labels: _Optional[_Iterable[_Union[_common_pb2.KeyValue, _Mapping]]] = ..., metadata: _Optional[_Iterable[_Union[_common_pb2.KeyValue, _Mapping]]] = ..., data: _Optional[str] = ...) -> None: ...

class Result(_message.Message):
    __slots__ = ("index", "labels", "metadata", "data")
    INDEX_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    index: _wrappers_pb2.Int64Value
    labels: _containers.RepeatedCompositeFieldContainer[_common_pb2.KeyValue]
    metadata: _containers.RepeatedCompositeFieldContainer[_common_pb2.KeyValue]
    data: str
    def __init__(self, index: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., labels: _Optional[_Iterable[_Union[_common_pb2.KeyValue, _Mapping]]] = ..., metadata: _Optional[_Iterable[_Union[_common_pb2.KeyValue, _Mapping]]] = ..., data: _Optional[str] = ...) -> None: ...
