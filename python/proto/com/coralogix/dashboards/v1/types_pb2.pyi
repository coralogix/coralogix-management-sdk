from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Data(_message.Message):
    __slots__ = ("primitive", "composite")
    class Primitive(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        PRIMITIVE_UNSPECIFIED: _ClassVar[Data.Primitive]
        PRIMITIVE_BOOL: _ClassVar[Data.Primitive]
        PRIMITIVE_NUMBER: _ClassVar[Data.Primitive]
        PRIMITIVE_STRING: _ClassVar[Data.Primitive]
    PRIMITIVE_UNSPECIFIED: Data.Primitive
    PRIMITIVE_BOOL: Data.Primitive
    PRIMITIVE_NUMBER: Data.Primitive
    PRIMITIVE_STRING: Data.Primitive
    class Composite(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        COMPOSITE_UNSPECIFIED: _ClassVar[Data.Composite]
        COMPOSITE_ARRAY: _ClassVar[Data.Composite]
        COMPOSITE_OBJECT: _ClassVar[Data.Composite]
    COMPOSITE_UNSPECIFIED: Data.Composite
    COMPOSITE_ARRAY: Data.Composite
    COMPOSITE_OBJECT: Data.Composite
    PRIMITIVE_FIELD_NUMBER: _ClassVar[int]
    COMPOSITE_FIELD_NUMBER: _ClassVar[int]
    primitive: Data.Primitive
    composite: Data.Composite
    def __init__(self, primitive: _Optional[_Union[Data.Primitive, str]] = ..., composite: _Optional[_Union[Data.Composite, str]] = ...) -> None: ...
