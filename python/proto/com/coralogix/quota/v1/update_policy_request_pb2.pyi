from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.quota.v1 import enums_pb2 as _enums_pb2
from com.coralogix.quota.v1 import rule_pb2 as _rule_pb2
from com.coralogix.quota.v1 import archive_retention_pb2 as _archive_retention_pb2
from com.coralogix.quota.v1 import log_rules_pb2 as _log_rules_pb2
from com.coralogix.quota.v1 import span_rules_pb2 as _span_rules_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UpdatePolicyRequest(_message.Message):
    __slots__ = ("id", "name", "description", "priority", "application_rule", "subsystem_rule", "archive_retention", "log_rules", "span_rules", "enabled")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    PRIORITY_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_RULE_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_RULE_FIELD_NUMBER: _ClassVar[int]
    ARCHIVE_RETENTION_FIELD_NUMBER: _ClassVar[int]
    LOG_RULES_FIELD_NUMBER: _ClassVar[int]
    SPAN_RULES_FIELD_NUMBER: _ClassVar[int]
    ENABLED_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    priority: _enums_pb2.Priority
    application_rule: _rule_pb2.Rule
    subsystem_rule: _rule_pb2.Rule
    archive_retention: _archive_retention_pb2.ArchiveRetention
    log_rules: _log_rules_pb2.LogRules
    span_rules: _span_rules_pb2.SpanRules
    enabled: _wrappers_pb2.BoolValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., priority: _Optional[_Union[_enums_pb2.Priority, str]] = ..., application_rule: _Optional[_Union[_rule_pb2.Rule, _Mapping]] = ..., subsystem_rule: _Optional[_Union[_rule_pb2.Rule, _Mapping]] = ..., archive_retention: _Optional[_Union[_archive_retention_pb2.ArchiveRetention, _Mapping]] = ..., log_rules: _Optional[_Union[_log_rules_pb2.LogRules, _Mapping]] = ..., span_rules: _Optional[_Union[_span_rules_pb2.SpanRules, _Mapping]] = ..., enabled: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
