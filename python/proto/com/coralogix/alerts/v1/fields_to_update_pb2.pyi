from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class FieldsToUpdate(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    FIELDS_TO_UPDATE_UNSUPPORTED_OR_UNSPECIFIED: _ClassVar[FieldsToUpdate]
    FIELDS_TO_UPDATE_GROUP_BY: _ClassVar[FieldsToUpdate]
    FIELDS_TO_UPDATE_ACTIVE_WHEN: _ClassVar[FieldsToUpdate]
FIELDS_TO_UPDATE_UNSUPPORTED_OR_UNSPECIFIED: FieldsToUpdate
FIELDS_TO_UPDATE_GROUP_BY: FieldsToUpdate
FIELDS_TO_UPDATE_ACTIVE_WHEN: FieldsToUpdate
