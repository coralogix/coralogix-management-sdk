from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class OriginatorType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ORIGINATOR_TYPE_UNSPECIFIED: _ClassVar[OriginatorType]
    ORIGINATOR_TYPE_OPERATIONAL: _ClassVar[OriginatorType]
    ORIGINATOR_TYPE_ADMINISTRATIVE: _ClassVar[OriginatorType]
ORIGINATOR_TYPE_UNSPECIFIED: OriginatorType
ORIGINATOR_TYPE_OPERATIONAL: OriginatorType
ORIGINATOR_TYPE_ADMINISTRATIVE: OriginatorType
