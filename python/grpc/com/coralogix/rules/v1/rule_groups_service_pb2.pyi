from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.rules.v1 import rule_pb2 as _rule_pb2
from com.coralogix.rules.v1 import rule_group_pb2 as _rule_group_pb2
from com.coralogix.rules.v1 import rule_matcher_pb2 as _rule_matcher_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor
AUDIT_LOG_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
audit_log_description: _descriptor.FieldDescriptor

class AuditLogDescription(_message.Message):
    __slots__ = ("description",)
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    description: str
    def __init__(self, description: _Optional[str] = ...) -> None: ...

class GetRuleGroupRequest(_message.Message):
    __slots__ = ("group_id",)
    GROUP_ID_FIELD_NUMBER: _ClassVar[int]
    group_id: str
    def __init__(self, group_id: _Optional[str] = ...) -> None: ...

class ListRuleGroupsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListRuleGroupsResponse(_message.Message):
    __slots__ = ("rule_groups",)
    RULE_GROUPS_FIELD_NUMBER: _ClassVar[int]
    rule_groups: _containers.RepeatedCompositeFieldContainer[_rule_group_pb2.RuleGroup]
    def __init__(self, rule_groups: _Optional[_Iterable[_Union[_rule_group_pb2.RuleGroup, _Mapping]]] = ...) -> None: ...

class GetRuleGroupResponse(_message.Message):
    __slots__ = ("rule_group",)
    RULE_GROUP_FIELD_NUMBER: _ClassVar[int]
    rule_group: _rule_group_pb2.RuleGroup
    def __init__(self, rule_group: _Optional[_Union[_rule_group_pb2.RuleGroup, _Mapping]] = ...) -> None: ...

class CreateRuleGroupRequest(_message.Message):
    __slots__ = ("name", "description", "enabled", "hidden", "creator", "rule_matchers", "rule_subgroups", "order", "team_id")
    class CreateRuleSubgroup(_message.Message):
        __slots__ = ("rules", "enabled", "order")
        class CreateRule(_message.Message):
            __slots__ = ("name", "description", "source_field", "parameters", "enabled", "order")
            NAME_FIELD_NUMBER: _ClassVar[int]
            DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
            SOURCE_FIELD_FIELD_NUMBER: _ClassVar[int]
            PARAMETERS_FIELD_NUMBER: _ClassVar[int]
            ENABLED_FIELD_NUMBER: _ClassVar[int]
            ORDER_FIELD_NUMBER: _ClassVar[int]
            name: _wrappers_pb2.StringValue
            description: _wrappers_pb2.StringValue
            source_field: _wrappers_pb2.StringValue
            parameters: _rule_pb2.RuleParameters
            enabled: _wrappers_pb2.BoolValue
            order: _wrappers_pb2.UInt32Value
            def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., source_field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., parameters: _Optional[_Union[_rule_pb2.RuleParameters, _Mapping]] = ..., enabled: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., order: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...
        RULES_FIELD_NUMBER: _ClassVar[int]
        ENABLED_FIELD_NUMBER: _ClassVar[int]
        ORDER_FIELD_NUMBER: _ClassVar[int]
        rules: _containers.RepeatedCompositeFieldContainer[CreateRuleGroupRequest.CreateRuleSubgroup.CreateRule]
        enabled: _wrappers_pb2.BoolValue
        order: _wrappers_pb2.UInt32Value
        def __init__(self, rules: _Optional[_Iterable[_Union[CreateRuleGroupRequest.CreateRuleSubgroup.CreateRule, _Mapping]]] = ..., enabled: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., order: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    ENABLED_FIELD_NUMBER: _ClassVar[int]
    HIDDEN_FIELD_NUMBER: _ClassVar[int]
    CREATOR_FIELD_NUMBER: _ClassVar[int]
    RULE_MATCHERS_FIELD_NUMBER: _ClassVar[int]
    RULE_SUBGROUPS_FIELD_NUMBER: _ClassVar[int]
    ORDER_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    enabled: _wrappers_pb2.BoolValue
    hidden: _wrappers_pb2.BoolValue
    creator: _wrappers_pb2.StringValue
    rule_matchers: _containers.RepeatedCompositeFieldContainer[_rule_matcher_pb2.RuleMatcher]
    rule_subgroups: _containers.RepeatedCompositeFieldContainer[CreateRuleGroupRequest.CreateRuleSubgroup]
    order: _wrappers_pb2.UInt32Value
    team_id: _rule_group_pb2.TeamId
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., enabled: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., hidden: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., creator: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., rule_matchers: _Optional[_Iterable[_Union[_rule_matcher_pb2.RuleMatcher, _Mapping]]] = ..., rule_subgroups: _Optional[_Iterable[_Union[CreateRuleGroupRequest.CreateRuleSubgroup, _Mapping]]] = ..., order: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., team_id: _Optional[_Union[_rule_group_pb2.TeamId, _Mapping]] = ...) -> None: ...

class CreateRuleGroupResponse(_message.Message):
    __slots__ = ("rule_group",)
    RULE_GROUP_FIELD_NUMBER: _ClassVar[int]
    rule_group: _rule_group_pb2.RuleGroup
    def __init__(self, rule_group: _Optional[_Union[_rule_group_pb2.RuleGroup, _Mapping]] = ...) -> None: ...

class UpdateRuleGroupRequest(_message.Message):
    __slots__ = ("group_id", "rule_group")
    GROUP_ID_FIELD_NUMBER: _ClassVar[int]
    RULE_GROUP_FIELD_NUMBER: _ClassVar[int]
    group_id: _wrappers_pb2.StringValue
    rule_group: CreateRuleGroupRequest
    def __init__(self, group_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., rule_group: _Optional[_Union[CreateRuleGroupRequest, _Mapping]] = ...) -> None: ...

class UpdateRuleGroupResponse(_message.Message):
    __slots__ = ("rule_group",)
    RULE_GROUP_FIELD_NUMBER: _ClassVar[int]
    rule_group: _rule_group_pb2.RuleGroup
    def __init__(self, rule_group: _Optional[_Union[_rule_group_pb2.RuleGroup, _Mapping]] = ...) -> None: ...

class DeleteRuleGroupRequest(_message.Message):
    __slots__ = ("group_id",)
    GROUP_ID_FIELD_NUMBER: _ClassVar[int]
    group_id: str
    def __init__(self, group_id: _Optional[str] = ...) -> None: ...

class BulkDeleteRuleGroupRequest(_message.Message):
    __slots__ = ("group_ids",)
    GROUP_IDS_FIELD_NUMBER: _ClassVar[int]
    group_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, group_ids: _Optional[_Iterable[str]] = ...) -> None: ...

class DeleteRuleGroupResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class BulkDeleteRuleGroupResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetRuleGroupModelMappingRequest(_message.Message):
    __slots__ = ("name", "description", "enabled", "hidden", "creator", "rule_matchers", "rule_subgroups", "order")
    class CreateRuleSubgroup(_message.Message):
        __slots__ = ("rules", "enabled", "order")
        class CreateRule(_message.Message):
            __slots__ = ("name", "description", "source_field", "parameters", "enabled", "order")
            NAME_FIELD_NUMBER: _ClassVar[int]
            DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
            SOURCE_FIELD_FIELD_NUMBER: _ClassVar[int]
            PARAMETERS_FIELD_NUMBER: _ClassVar[int]
            ENABLED_FIELD_NUMBER: _ClassVar[int]
            ORDER_FIELD_NUMBER: _ClassVar[int]
            name: _wrappers_pb2.StringValue
            description: _wrappers_pb2.StringValue
            source_field: _wrappers_pb2.StringValue
            parameters: _rule_pb2.RuleParameters
            enabled: _wrappers_pb2.BoolValue
            order: _wrappers_pb2.UInt32Value
            def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., source_field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., parameters: _Optional[_Union[_rule_pb2.RuleParameters, _Mapping]] = ..., enabled: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., order: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...
        RULES_FIELD_NUMBER: _ClassVar[int]
        ENABLED_FIELD_NUMBER: _ClassVar[int]
        ORDER_FIELD_NUMBER: _ClassVar[int]
        rules: _containers.RepeatedCompositeFieldContainer[GetRuleGroupModelMappingRequest.CreateRuleSubgroup.CreateRule]
        enabled: _wrappers_pb2.BoolValue
        order: _wrappers_pb2.UInt32Value
        def __init__(self, rules: _Optional[_Iterable[_Union[GetRuleGroupModelMappingRequest.CreateRuleSubgroup.CreateRule, _Mapping]]] = ..., enabled: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., order: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    ENABLED_FIELD_NUMBER: _ClassVar[int]
    HIDDEN_FIELD_NUMBER: _ClassVar[int]
    CREATOR_FIELD_NUMBER: _ClassVar[int]
    RULE_MATCHERS_FIELD_NUMBER: _ClassVar[int]
    RULE_SUBGROUPS_FIELD_NUMBER: _ClassVar[int]
    ORDER_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    enabled: _wrappers_pb2.BoolValue
    hidden: _wrappers_pb2.BoolValue
    creator: _wrappers_pb2.StringValue
    rule_matchers: _containers.RepeatedCompositeFieldContainer[_rule_matcher_pb2.RuleMatcher]
    rule_subgroups: _containers.RepeatedCompositeFieldContainer[GetRuleGroupModelMappingRequest.CreateRuleSubgroup]
    order: _wrappers_pb2.UInt32Value
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., enabled: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., hidden: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., creator: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., rule_matchers: _Optional[_Iterable[_Union[_rule_matcher_pb2.RuleMatcher, _Mapping]]] = ..., rule_subgroups: _Optional[_Iterable[_Union[GetRuleGroupModelMappingRequest.CreateRuleSubgroup, _Mapping]]] = ..., order: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...

class GetRuleGroupModelMappingResponse(_message.Message):
    __slots__ = ("rule_definition",)
    RULE_DEFINITION_FIELD_NUMBER: _ClassVar[int]
    rule_definition: _struct_pb2.Struct
    def __init__(self, rule_definition: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...

class GetCompanyUsageLimitsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetCompanyUsageLimitsResponse(_message.Message):
    __slots__ = ("company_id", "limits", "usage")
    class Counts(_message.Message):
        __slots__ = ("parsing_themes", "groups", "rules")
        PARSING_THEMES_FIELD_NUMBER: _ClassVar[int]
        GROUPS_FIELD_NUMBER: _ClassVar[int]
        RULES_FIELD_NUMBER: _ClassVar[int]
        parsing_themes: _wrappers_pb2.Int32Value
        groups: _wrappers_pb2.Int32Value
        rules: _wrappers_pb2.Int32Value
        def __init__(self, parsing_themes: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., groups: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., rules: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    COMPANY_ID_FIELD_NUMBER: _ClassVar[int]
    LIMITS_FIELD_NUMBER: _ClassVar[int]
    USAGE_FIELD_NUMBER: _ClassVar[int]
    company_id: _wrappers_pb2.StringValue
    limits: GetCompanyUsageLimitsResponse.Counts
    usage: GetCompanyUsageLimitsResponse.Counts
    def __init__(self, company_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., limits: _Optional[_Union[GetCompanyUsageLimitsResponse.Counts, _Mapping]] = ..., usage: _Optional[_Union[GetCompanyUsageLimitsResponse.Counts, _Mapping]] = ...) -> None: ...
