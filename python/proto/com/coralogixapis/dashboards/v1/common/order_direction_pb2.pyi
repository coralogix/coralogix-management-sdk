from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class OrderDirection(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ORDER_DIRECTION_UNSPECIFIED: _ClassVar[OrderDirection]
    ORDER_DIRECTION_ASC: _ClassVar[OrderDirection]
    ORDER_DIRECTION_DESC: _ClassVar[OrderDirection]
ORDER_DIRECTION_UNSPECIFIED: OrderDirection
ORDER_DIRECTION_ASC: OrderDirection
ORDER_DIRECTION_DESC: OrderDirection
