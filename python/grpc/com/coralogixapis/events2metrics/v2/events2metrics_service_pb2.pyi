from com.coralogixapis.events2metrics.v2 import events2metrics_definition_pb2 as _events2metrics_definition_pb2
from com.coralogixapis.logs2metrics.v2 import audit_log_pb2 as _audit_log_pb2
from com.coralogixapis.logs2metrics.v2 import logs_query_pb2 as _logs_query_pb2
from com.coralogixapis.spans2metrics.v2 import spans_query_pb2 as _spans_query_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.openapi.v1 import annotations_pb2 as _annotations_pb2_1
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateE2MRequest(_message.Message):
    __slots__ = ("e2m",)
    E2M_FIELD_NUMBER: _ClassVar[int]
    e2m: _events2metrics_definition_pb2.E2MCreateParams
    def __init__(self, e2m: _Optional[_Union[_events2metrics_definition_pb2.E2MCreateParams, _Mapping]] = ...) -> None: ...

class CreateE2MResponse(_message.Message):
    __slots__ = ("e2m",)
    E2M_FIELD_NUMBER: _ClassVar[int]
    e2m: _events2metrics_definition_pb2.E2M
    def __init__(self, e2m: _Optional[_Union[_events2metrics_definition_pb2.E2M, _Mapping]] = ...) -> None: ...

class ListE2MRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListE2MResponse(_message.Message):
    __slots__ = ("e2m",)
    E2M_FIELD_NUMBER: _ClassVar[int]
    e2m: _containers.RepeatedCompositeFieldContainer[_events2metrics_definition_pb2.E2M]
    def __init__(self, e2m: _Optional[_Iterable[_Union[_events2metrics_definition_pb2.E2M, _Mapping]]] = ...) -> None: ...

class ReplaceE2MRequest(_message.Message):
    __slots__ = ("e2m",)
    E2M_FIELD_NUMBER: _ClassVar[int]
    e2m: _events2metrics_definition_pb2.E2M
    def __init__(self, e2m: _Optional[_Union[_events2metrics_definition_pb2.E2M, _Mapping]] = ...) -> None: ...

class ReplaceE2MResponse(_message.Message):
    __slots__ = ("e2m",)
    E2M_FIELD_NUMBER: _ClassVar[int]
    e2m: _events2metrics_definition_pb2.E2M
    def __init__(self, e2m: _Optional[_Union[_events2metrics_definition_pb2.E2M, _Mapping]] = ...) -> None: ...

class GetE2MRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetE2MResponse(_message.Message):
    __slots__ = ("e2m",)
    E2M_FIELD_NUMBER: _ClassVar[int]
    e2m: _events2metrics_definition_pb2.E2M
    def __init__(self, e2m: _Optional[_Union[_events2metrics_definition_pb2.E2M, _Mapping]] = ...) -> None: ...

class DeleteE2MRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class DeleteE2MResponse(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class E2MExecutionRequest(_message.Message):
    __slots__ = ("create", "replace", "delete")
    CREATE_FIELD_NUMBER: _ClassVar[int]
    REPLACE_FIELD_NUMBER: _ClassVar[int]
    DELETE_FIELD_NUMBER: _ClassVar[int]
    create: CreateE2MRequest
    replace: ReplaceE2MRequest
    delete: DeleteE2MRequest
    def __init__(self, create: _Optional[_Union[CreateE2MRequest, _Mapping]] = ..., replace: _Optional[_Union[ReplaceE2MRequest, _Mapping]] = ..., delete: _Optional[_Union[DeleteE2MRequest, _Mapping]] = ...) -> None: ...

class AtomicBatchExecuteE2MRequest(_message.Message):
    __slots__ = ("requests",)
    REQUESTS_FIELD_NUMBER: _ClassVar[int]
    requests: _containers.RepeatedCompositeFieldContainer[E2MExecutionRequest]
    def __init__(self, requests: _Optional[_Iterable[_Union[E2MExecutionRequest, _Mapping]]] = ...) -> None: ...

class E2MExecutionResponse(_message.Message):
    __slots__ = ("created", "replaced", "deleted")
    CREATED_FIELD_NUMBER: _ClassVar[int]
    REPLACED_FIELD_NUMBER: _ClassVar[int]
    DELETED_FIELD_NUMBER: _ClassVar[int]
    created: CreateE2MResponse
    replaced: ReplaceE2MResponse
    deleted: DeleteE2MResponse
    def __init__(self, created: _Optional[_Union[CreateE2MResponse, _Mapping]] = ..., replaced: _Optional[_Union[ReplaceE2MResponse, _Mapping]] = ..., deleted: _Optional[_Union[DeleteE2MResponse, _Mapping]] = ...) -> None: ...

class AtomicBatchExecuteE2MResponse(_message.Message):
    __slots__ = ("matching_responses",)
    MATCHING_RESPONSES_FIELD_NUMBER: _ClassVar[int]
    matching_responses: _containers.RepeatedCompositeFieldContainer[E2MExecutionResponse]
    def __init__(self, matching_responses: _Optional[_Iterable[_Union[E2MExecutionResponse, _Mapping]]] = ...) -> None: ...

class ListLabelsCardinalityRequest(_message.Message):
    __slots__ = ("spans_query", "logs_query", "metric_labels")
    SPANS_QUERY_FIELD_NUMBER: _ClassVar[int]
    LOGS_QUERY_FIELD_NUMBER: _ClassVar[int]
    METRIC_LABELS_FIELD_NUMBER: _ClassVar[int]
    spans_query: _spans_query_pb2.SpansQuery
    logs_query: _logs_query_pb2.LogsQuery
    metric_labels: _containers.RepeatedCompositeFieldContainer[_events2metrics_definition_pb2.MetricLabel]
    def __init__(self, spans_query: _Optional[_Union[_spans_query_pb2.SpansQuery, _Mapping]] = ..., logs_query: _Optional[_Union[_logs_query_pb2.LogsQuery, _Mapping]] = ..., metric_labels: _Optional[_Iterable[_Union[_events2metrics_definition_pb2.MetricLabel, _Mapping]]] = ...) -> None: ...

class LabelsPermutationsCardinalityDay(_message.Message):
    __slots__ = ("day", "permutations")
    DAY_FIELD_NUMBER: _ClassVar[int]
    PERMUTATIONS_FIELD_NUMBER: _ClassVar[int]
    day: str
    permutations: int
    def __init__(self, day: _Optional[str] = ..., permutations: _Optional[int] = ...) -> None: ...

class ListLabelsCardinalityResponse(_message.Message):
    __slots__ = ("permutations",)
    PERMUTATIONS_FIELD_NUMBER: _ClassVar[int]
    permutations: _containers.RepeatedCompositeFieldContainer[LabelsPermutationsCardinalityDay]
    def __init__(self, permutations: _Optional[_Iterable[_Union[LabelsPermutationsCardinalityDay, _Mapping]]] = ...) -> None: ...

class GetLimitsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetLimitsResponse(_message.Message):
    __slots__ = ("company_id", "labels_limit", "permutations_limit", "metrics_limit")
    class LimitUsage(_message.Message):
        __slots__ = ("limit", "used")
        LIMIT_FIELD_NUMBER: _ClassVar[int]
        USED_FIELD_NUMBER: _ClassVar[int]
        limit: _wrappers_pb2.Int32Value
        used: _wrappers_pb2.Int32Value
        def __init__(self, limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., used: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    COMPANY_ID_FIELD_NUMBER: _ClassVar[int]
    LABELS_LIMIT_FIELD_NUMBER: _ClassVar[int]
    PERMUTATIONS_LIMIT_FIELD_NUMBER: _ClassVar[int]
    METRICS_LIMIT_FIELD_NUMBER: _ClassVar[int]
    company_id: _wrappers_pb2.StringValue
    labels_limit: int
    permutations_limit: GetLimitsResponse.LimitUsage
    metrics_limit: GetLimitsResponse.LimitUsage
    def __init__(self, company_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., labels_limit: _Optional[int] = ..., permutations_limit: _Optional[_Union[GetLimitsResponse.LimitUsage, _Mapping]] = ..., metrics_limit: _Optional[_Union[GetLimitsResponse.LimitUsage, _Mapping]] = ...) -> None: ...
