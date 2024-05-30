from com.coralogix.quota.v1 import policy_pb2 as _policy_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UpdatePolicyResponse(_message.Message):
    __slots__ = ("policy",)
    POLICY_FIELD_NUMBER: _ClassVar[int]
    policy: _policy_pb2.Policy
    def __init__(self, policy: _Optional[_Union[_policy_pb2.Policy, _Mapping]] = ...) -> None: ...
