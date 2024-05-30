from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class DataModeType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    DATA_MODE_TYPE_HIGH_UNSPECIFIED: _ClassVar[DataModeType]
    DATA_MODE_TYPE_ARCHIVE: _ClassVar[DataModeType]
DATA_MODE_TYPE_HIGH_UNSPECIFIED: DataModeType
DATA_MODE_TYPE_ARCHIVE: DataModeType
