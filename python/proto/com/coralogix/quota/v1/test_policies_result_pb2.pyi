from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.quota.v1 import policy_pb2 as _policy_pb2
from com.coralogix.quota.v1 import log_meta_field_values_pb2 as _log_meta_field_values_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class TestPoliciesResult(_message.Message):
    __slots__ = ("meta_fields_values", "matched", "policy")
    META_FIELDS_VALUES_FIELD_NUMBER: _ClassVar[int]
    MATCHED_FIELD_NUMBER: _ClassVar[int]
    POLICY_FIELD_NUMBER: _ClassVar[int]
    meta_fields_values: _log_meta_field_values_pb2.LogMetaFieldsValues
    matched: _wrappers_pb2.BoolValue
    policy: _policy_pb2.Policy
    def __init__(self, meta_fields_values: _Optional[_Union[_log_meta_field_values_pb2.LogMetaFieldsValues, _Mapping]] = ..., matched: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., policy: _Optional[_Union[_policy_pb2.Policy, _Mapping]] = ...) -> None: ...
