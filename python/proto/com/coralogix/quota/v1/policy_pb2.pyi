from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.quota.v1 import archive_retention_pb2 as _archive_retention_pb2
from com.coralogix.quota.v1 import enums_pb2 as _enums_pb2
from com.coralogix.quota.v1 import rule_pb2 as _rule_pb2
from com.coralogix.quota.v1 import log_rules_pb2 as _log_rules_pb2
from com.coralogix.quota.v1 import span_rules_pb2 as _span_rules_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Policy(_message.Message):
    __slots__ = ("id", "company_id", "name", "description", "priority", "deleted", "enabled", "order", "application_rule", "subsystem_rule", "log_rules", "span_rules", "created_at", "updated_at", "archive_retention")
    ID_FIELD_NUMBER: _ClassVar[int]
    COMPANY_ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    PRIORITY_FIELD_NUMBER: _ClassVar[int]
    DELETED_FIELD_NUMBER: _ClassVar[int]
    ENABLED_FIELD_NUMBER: _ClassVar[int]
    ORDER_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_RULE_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_RULE_FIELD_NUMBER: _ClassVar[int]
    LOG_RULES_FIELD_NUMBER: _ClassVar[int]
    SPAN_RULES_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    ARCHIVE_RETENTION_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    company_id: _wrappers_pb2.Int32Value
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    priority: _enums_pb2.Priority
    deleted: _wrappers_pb2.BoolValue
    enabled: _wrappers_pb2.BoolValue
    order: _wrappers_pb2.Int32Value
    application_rule: _rule_pb2.Rule
    subsystem_rule: _rule_pb2.Rule
    log_rules: _log_rules_pb2.LogRules
    span_rules: _span_rules_pb2.SpanRules
    created_at: _wrappers_pb2.StringValue
    updated_at: _wrappers_pb2.StringValue
    archive_retention: _archive_retention_pb2.ArchiveRetention
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., company_id: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., priority: _Optional[_Union[_enums_pb2.Priority, str]] = ..., deleted: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., enabled: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., order: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., application_rule: _Optional[_Union[_rule_pb2.Rule, _Mapping]] = ..., subsystem_rule: _Optional[_Union[_rule_pb2.Rule, _Mapping]] = ..., log_rules: _Optional[_Union[_log_rules_pb2.LogRules, _Mapping]] = ..., span_rules: _Optional[_Union[_span_rules_pb2.SpanRules, _Mapping]] = ..., created_at: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., updated_at: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., archive_retention: _Optional[_Union[_archive_retention_pb2.ArchiveRetention, _Mapping]] = ...) -> None: ...
