from com.coralogix.quota.v1 import rule_pb2 as _rule_pb2
from com.coralogix.quota.v1 import tag_rule_pb2 as _tag_rule_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SpanRules(_message.Message):
    __slots__ = ("service_rule", "action_rule", "tag_rules")
    SERVICE_RULE_FIELD_NUMBER: _ClassVar[int]
    ACTION_RULE_FIELD_NUMBER: _ClassVar[int]
    TAG_RULES_FIELD_NUMBER: _ClassVar[int]
    service_rule: _rule_pb2.Rule
    action_rule: _rule_pb2.Rule
    tag_rules: _containers.RepeatedCompositeFieldContainer[_tag_rule_pb2.TagRule]
    def __init__(self, service_rule: _Optional[_Union[_rule_pb2.Rule, _Mapping]] = ..., action_rule: _Optional[_Union[_rule_pb2.Rule, _Mapping]] = ..., tag_rules: _Optional[_Iterable[_Union[_tag_rule_pb2.TagRule, _Mapping]]] = ...) -> None: ...
