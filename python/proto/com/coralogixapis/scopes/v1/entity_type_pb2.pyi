from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class EntityType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    UNSPECIFIED: _ClassVar[EntityType]
    LOGS: _ClassVar[EntityType]
    SPANS: _ClassVar[EntityType]
UNSPECIFIED: EntityType
LOGS: EntityType
SPANS: EntityType
