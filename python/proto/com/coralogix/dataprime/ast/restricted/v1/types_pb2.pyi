from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Datatype(_message.Message):
    __slots__ = ("primitive_type", "composite_type")
    class PrimitiveType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        PRIMITIVE_TYPE_UNSPECIFIED: _ClassVar[Datatype.PrimitiveType]
        PRIMITIVE_TYPE_BOOL: _ClassVar[Datatype.PrimitiveType]
        PRIMITIVE_TYPE_NUM: _ClassVar[Datatype.PrimitiveType]
        PRIMITIVE_TYPE_STRING: _ClassVar[Datatype.PrimitiveType]
        PRIMITIVE_TYPE_TIMESTAMP: _ClassVar[Datatype.PrimitiveType]
    PRIMITIVE_TYPE_UNSPECIFIED: Datatype.PrimitiveType
    PRIMITIVE_TYPE_BOOL: Datatype.PrimitiveType
    PRIMITIVE_TYPE_NUM: Datatype.PrimitiveType
    PRIMITIVE_TYPE_STRING: Datatype.PrimitiveType
    PRIMITIVE_TYPE_TIMESTAMP: Datatype.PrimitiveType
    class CompositeType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        COMPOSITE_TYPE_UNSPECIFIED: _ClassVar[Datatype.CompositeType]
        COMPOSITE_TYPE_ARRAY: _ClassVar[Datatype.CompositeType]
        COMPOSITE_TYPE_OBJECT: _ClassVar[Datatype.CompositeType]
    COMPOSITE_TYPE_UNSPECIFIED: Datatype.CompositeType
    COMPOSITE_TYPE_ARRAY: Datatype.CompositeType
    COMPOSITE_TYPE_OBJECT: Datatype.CompositeType
    PRIMITIVE_TYPE_FIELD_NUMBER: _ClassVar[int]
    COMPOSITE_TYPE_FIELD_NUMBER: _ClassVar[int]
    primitive_type: Datatype.PrimitiveType
    composite_type: Datatype.CompositeType
    def __init__(self, primitive_type: _Optional[_Union[Datatype.PrimitiveType, str]] = ..., composite_type: _Optional[_Union[Datatype.CompositeType, str]] = ...) -> None: ...
