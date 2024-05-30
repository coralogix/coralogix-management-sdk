from com.coralogix.quota.v1 import create_policy_request_pb2 as _create_policy_request_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AtomicBatchCreatePolicyRequest(_message.Message):
    __slots__ = ("policy_requests",)
    POLICY_REQUESTS_FIELD_NUMBER: _ClassVar[int]
    policy_requests: _containers.RepeatedCompositeFieldContainer[_create_policy_request_pb2.CreatePolicyRequest]
    def __init__(self, policy_requests: _Optional[_Iterable[_Union[_create_policy_request_pb2.CreatePolicyRequest, _Mapping]]] = ...) -> None: ...
