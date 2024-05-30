from com.coralogix.quota.v1 import policy_order_pb2 as _policy_order_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ReorderPoliciesResponse(_message.Message):
    __slots__ = ("orders",)
    ORDERS_FIELD_NUMBER: _ClassVar[int]
    orders: _containers.RepeatedCompositeFieldContainer[_policy_order_pb2.PolicyOrder]
    def __init__(self, orders: _Optional[_Iterable[_Union[_policy_order_pb2.PolicyOrder, _Mapping]]] = ...) -> None: ...
