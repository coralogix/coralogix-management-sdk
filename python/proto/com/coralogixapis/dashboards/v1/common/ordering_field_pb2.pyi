from com.coralogixapis.dashboards.v1.common import order_direction_pb2 as _order_direction_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class OrderingField(_message.Message):
    __slots__ = ("field", "order_direction")
    FIELD_FIELD_NUMBER: _ClassVar[int]
    ORDER_DIRECTION_FIELD_NUMBER: _ClassVar[int]
    field: _wrappers_pb2.StringValue
    order_direction: _order_direction_pb2.OrderDirection
    def __init__(self, field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., order_direction: _Optional[_Union[_order_direction_pb2.OrderDirection, str]] = ...) -> None: ...
