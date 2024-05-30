from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.quota.v1 import enums_pb2 as _enums_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class TagRule(_message.Message):
    __slots__ = ("rule_type_id", "tag_name", "tag_value")
    RULE_TYPE_ID_FIELD_NUMBER: _ClassVar[int]
    TAG_NAME_FIELD_NUMBER: _ClassVar[int]
    TAG_VALUE_FIELD_NUMBER: _ClassVar[int]
    rule_type_id: _enums_pb2.RuleTypeId
    tag_name: _wrappers_pb2.StringValue
    tag_value: _wrappers_pb2.StringValue
    def __init__(self, rule_type_id: _Optional[_Union[_enums_pb2.RuleTypeId, str]] = ..., tag_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., tag_value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
