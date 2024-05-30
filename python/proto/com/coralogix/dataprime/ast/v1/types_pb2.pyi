from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Datatype(_message.Message):
    __slots__ = ("primitive_type", "composite_type", "union_type", "array_type", "enum_type")
    class PrimitiveType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        PRIMITIVE_TYPE_UNSPECIFIED: _ClassVar[Datatype.PrimitiveType]
        PRIMITIVE_TYPE_BOOL: _ClassVar[Datatype.PrimitiveType]
        PRIMITIVE_TYPE_NUM: _ClassVar[Datatype.PrimitiveType]
        PRIMITIVE_TYPE_STRING: _ClassVar[Datatype.PrimitiveType]
        PRIMITIVE_TYPE_TIMESTAMP: _ClassVar[Datatype.PrimitiveType]
        PRIMITIVE_TYPE_REGEX: _ClassVar[Datatype.PrimitiveType]
        PRIMITIVE_TYPE_INTERVAL: _ClassVar[Datatype.PrimitiveType]
        PRIMITIVE_TYPE_DATATYPE: _ClassVar[Datatype.PrimitiveType]
        PRIMITIVE_TYPE_TIME_UNIT: _ClassVar[Datatype.PrimitiveType]
        PRIMITIVE_TYPE_DATE_UNIT: _ClassVar[Datatype.PrimitiveType]
        PRIMITIVE_TYPE_NULL: _ClassVar[Datatype.PrimitiveType]
    PRIMITIVE_TYPE_UNSPECIFIED: Datatype.PrimitiveType
    PRIMITIVE_TYPE_BOOL: Datatype.PrimitiveType
    PRIMITIVE_TYPE_NUM: Datatype.PrimitiveType
    PRIMITIVE_TYPE_STRING: Datatype.PrimitiveType
    PRIMITIVE_TYPE_TIMESTAMP: Datatype.PrimitiveType
    PRIMITIVE_TYPE_REGEX: Datatype.PrimitiveType
    PRIMITIVE_TYPE_INTERVAL: Datatype.PrimitiveType
    PRIMITIVE_TYPE_DATATYPE: Datatype.PrimitiveType
    PRIMITIVE_TYPE_TIME_UNIT: Datatype.PrimitiveType
    PRIMITIVE_TYPE_DATE_UNIT: Datatype.PrimitiveType
    PRIMITIVE_TYPE_NULL: Datatype.PrimitiveType
    class CompositeType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        COMPOSITE_TYPE_UNSPECIFIED: _ClassVar[Datatype.CompositeType]
        COMPOSITE_TYPE_ARRAY: _ClassVar[Datatype.CompositeType]
        COMPOSITE_TYPE_OBJECT: _ClassVar[Datatype.CompositeType]
    COMPOSITE_TYPE_UNSPECIFIED: Datatype.CompositeType
    COMPOSITE_TYPE_ARRAY: Datatype.CompositeType
    COMPOSITE_TYPE_OBJECT: Datatype.CompositeType
    class UnionType(_message.Message):
        __slots__ = ("types",)
        TYPES_FIELD_NUMBER: _ClassVar[int]
        types: _containers.RepeatedCompositeFieldContainer[Datatype]
        def __init__(self, types: _Optional[_Iterable[_Union[Datatype, _Mapping]]] = ...) -> None: ...
    class ArrayType(_message.Message):
        __slots__ = ("element",)
        ELEMENT_FIELD_NUMBER: _ClassVar[int]
        element: Datatype
        def __init__(self, element: _Optional[_Union[Datatype, _Mapping]] = ...) -> None: ...
    class EnumType(_message.Message):
        __slots__ = ("name",)
        NAME_FIELD_NUMBER: _ClassVar[int]
        name: str
        def __init__(self, name: _Optional[str] = ...) -> None: ...
    PRIMITIVE_TYPE_FIELD_NUMBER: _ClassVar[int]
    COMPOSITE_TYPE_FIELD_NUMBER: _ClassVar[int]
    UNION_TYPE_FIELD_NUMBER: _ClassVar[int]
    ARRAY_TYPE_FIELD_NUMBER: _ClassVar[int]
    ENUM_TYPE_FIELD_NUMBER: _ClassVar[int]
    primitive_type: Datatype.PrimitiveType
    composite_type: Datatype.CompositeType
    union_type: Datatype.UnionType
    array_type: Datatype.ArrayType
    enum_type: Datatype.EnumType
    def __init__(self, primitive_type: _Optional[_Union[Datatype.PrimitiveType, str]] = ..., composite_type: _Optional[_Union[Datatype.CompositeType, str]] = ..., union_type: _Optional[_Union[Datatype.UnionType, _Mapping]] = ..., array_type: _Optional[_Union[Datatype.ArrayType, _Mapping]] = ..., enum_type: _Optional[_Union[Datatype.EnumType, _Mapping]] = ...) -> None: ...
