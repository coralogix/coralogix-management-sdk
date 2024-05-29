from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class Operator(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    OPERATOR_COUNT_UNSPECIFIED: _ClassVar[Operator]
    OPERATOR_AVG: _ClassVar[Operator]
    OPERATOR_SUM: _ClassVar[Operator]
    OPERATOR_MIN: _ClassVar[Operator]
    OPERATOR_MAX: _ClassVar[Operator]
OPERATOR_COUNT_UNSPECIFIED: Operator
OPERATOR_AVG: Operator
OPERATOR_SUM: Operator
OPERATOR_MIN: Operator
OPERATOR_MAX: Operator
