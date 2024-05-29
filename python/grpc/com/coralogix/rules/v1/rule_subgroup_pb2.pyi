from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.rules.v1 import rule_pb2 as _rule_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class RuleSubgroup(_message.Message):
    __slots__ = ("id", "rules", "enabled", "order")
    ID_FIELD_NUMBER: _ClassVar[int]
    RULES_FIELD_NUMBER: _ClassVar[int]
    ENABLED_FIELD_NUMBER: _ClassVar[int]
    ORDER_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    rules: _containers.RepeatedCompositeFieldContainer[_rule_pb2.Rule]
    enabled: _wrappers_pb2.BoolValue
    order: _wrappers_pb2.UInt32Value
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., rules: _Optional[_Iterable[_Union[_rule_pb2.Rule, _Mapping]]] = ..., enabled: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., order: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...
