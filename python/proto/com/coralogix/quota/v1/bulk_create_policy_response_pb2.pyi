from com.coralogix.quota.v1 import create_policy_response_pb2 as _create_policy_response_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AtomicBatchCreatePolicyResponse(_message.Message):
    __slots__ = ("create_responses",)
    CREATE_RESPONSES_FIELD_NUMBER: _ClassVar[int]
    create_responses: _containers.RepeatedCompositeFieldContainer[_create_policy_response_pb2.CreatePolicyResponse]
    def __init__(self, create_responses: _Optional[_Iterable[_Union[_create_policy_response_pb2.CreatePolicyResponse, _Mapping]]] = ...) -> None: ...
