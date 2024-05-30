from com.coralogix.extensions.v1 import audit_log_pb2 as _audit_log_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetQuotasRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetQuotasResponse(_message.Message):
    __slots__ = ("company_id", "alert", "enrichment", "parsing_rule", "parsing_rule_group", "parsing_theme", "dynamic_alert", "events_2_metrics")
    class Usage(_message.Message):
        __slots__ = ("used", "limit")
        USED_FIELD_NUMBER: _ClassVar[int]
        LIMIT_FIELD_NUMBER: _ClassVar[int]
        used: _wrappers_pb2.Int32Value
        limit: _wrappers_pb2.Int32Value
        def __init__(self, used: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    class Events2MetricsUsage(_message.Message):
        __slots__ = ("labels_limit", "permutations", "metrics")
        LABELS_LIMIT_FIELD_NUMBER: _ClassVar[int]
        PERMUTATIONS_FIELD_NUMBER: _ClassVar[int]
        METRICS_FIELD_NUMBER: _ClassVar[int]
        labels_limit: _wrappers_pb2.Int32Value
        permutations: GetQuotasResponse.Usage
        metrics: GetQuotasResponse.Usage
        def __init__(self, labels_limit: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., permutations: _Optional[_Union[GetQuotasResponse.Usage, _Mapping]] = ..., metrics: _Optional[_Union[GetQuotasResponse.Usage, _Mapping]] = ...) -> None: ...
    COMPANY_ID_FIELD_NUMBER: _ClassVar[int]
    ALERT_FIELD_NUMBER: _ClassVar[int]
    ENRICHMENT_FIELD_NUMBER: _ClassVar[int]
    PARSING_RULE_FIELD_NUMBER: _ClassVar[int]
    PARSING_RULE_GROUP_FIELD_NUMBER: _ClassVar[int]
    PARSING_THEME_FIELD_NUMBER: _ClassVar[int]
    DYNAMIC_ALERT_FIELD_NUMBER: _ClassVar[int]
    EVENTS_2_METRICS_FIELD_NUMBER: _ClassVar[int]
    company_id: _wrappers_pb2.StringValue
    alert: GetQuotasResponse.Usage
    enrichment: GetQuotasResponse.Usage
    parsing_rule: GetQuotasResponse.Usage
    parsing_rule_group: GetQuotasResponse.Usage
    parsing_theme: GetQuotasResponse.Usage
    dynamic_alert: GetQuotasResponse.Usage
    events_2_metrics: GetQuotasResponse.Events2MetricsUsage
    def __init__(self, company_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., alert: _Optional[_Union[GetQuotasResponse.Usage, _Mapping]] = ..., enrichment: _Optional[_Union[GetQuotasResponse.Usage, _Mapping]] = ..., parsing_rule: _Optional[_Union[GetQuotasResponse.Usage, _Mapping]] = ..., parsing_rule_group: _Optional[_Union[GetQuotasResponse.Usage, _Mapping]] = ..., parsing_theme: _Optional[_Union[GetQuotasResponse.Usage, _Mapping]] = ..., dynamic_alert: _Optional[_Union[GetQuotasResponse.Usage, _Mapping]] = ..., events_2_metrics: _Optional[_Union[GetQuotasResponse.Events2MetricsUsage, _Mapping]] = ...) -> None: ...
