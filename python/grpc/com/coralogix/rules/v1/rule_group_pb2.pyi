from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.rules.v1 import rule_matcher_pb2 as _rule_matcher_pb2
from com.coralogix.rules.v1 import rule_subgroup_pb2 as _rule_subgroup_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class RuleGroup(_message.Message):
    __slots__ = ("id", "name", "description", "creator", "enabled", "hidden", "rule_matchers", "rule_subgroups", "order")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    CREATOR_FIELD_NUMBER: _ClassVar[int]
    ENABLED_FIELD_NUMBER: _ClassVar[int]
    HIDDEN_FIELD_NUMBER: _ClassVar[int]
    RULE_MATCHERS_FIELD_NUMBER: _ClassVar[int]
    RULE_SUBGROUPS_FIELD_NUMBER: _ClassVar[int]
    ORDER_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    creator: _wrappers_pb2.StringValue
    enabled: _wrappers_pb2.BoolValue
    hidden: _wrappers_pb2.BoolValue
    rule_matchers: _containers.RepeatedCompositeFieldContainer[_rule_matcher_pb2.RuleMatcher]
    rule_subgroups: _containers.RepeatedCompositeFieldContainer[_rule_subgroup_pb2.RuleSubgroup]
    order: _wrappers_pb2.UInt32Value
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., creator: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., enabled: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., hidden: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., rule_matchers: _Optional[_Iterable[_Union[_rule_matcher_pb2.RuleMatcher, _Mapping]]] = ..., rule_subgroups: _Optional[_Iterable[_Union[_rule_subgroup_pb2.RuleSubgroup, _Mapping]]] = ..., order: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...

class TeamId(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...
