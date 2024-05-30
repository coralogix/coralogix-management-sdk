from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class ScaleType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    SCALE_TYPE_UNSPECIFIED: _ClassVar[ScaleType]
    SCALE_TYPE_LINEAR: _ClassVar[ScaleType]
    SCALE_TYPE_LOGARITHMIC: _ClassVar[ScaleType]
SCALE_TYPE_UNSPECIFIED: ScaleType
SCALE_TYPE_LINEAR: ScaleType
SCALE_TYPE_LOGARITHMIC: ScaleType
