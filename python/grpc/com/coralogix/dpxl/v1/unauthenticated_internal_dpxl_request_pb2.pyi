from com.coralogix.dpxl.v1 import dpxl_pb2 as _dpxl_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UnauthenticatedInternalDpxlServiceCompileRequest(_message.Message):
    __slots__ = ("expression", "entity_types", "entity_type")
    EXPRESSION_FIELD_NUMBER: _ClassVar[int]
    ENTITY_TYPES_FIELD_NUMBER: _ClassVar[int]
    ENTITY_TYPE_FIELD_NUMBER: _ClassVar[int]
    expression: _dpxl_pb2.Expression
    entity_types: _containers.RepeatedCompositeFieldContainer[_dpxl_pb2.EntityType]
    entity_type: _dpxl_pb2.EntityType
    def __init__(self, expression: _Optional[_Union[_dpxl_pb2.Expression, _Mapping]] = ..., entity_types: _Optional[_Iterable[_Union[_dpxl_pb2.EntityType, _Mapping]]] = ..., entity_type: _Optional[_Union[_dpxl_pb2.EntityType, _Mapping]] = ...) -> None: ...

class UnauthenticatedInternalDpxlServiceToDpxlTextRequest(_message.Message):
    __slots__ = ("serialized",)
    SERIALIZED_FIELD_NUMBER: _ClassVar[int]
    serialized: _dpxl_pb2.SerializedExpression
    def __init__(self, serialized: _Optional[_Union[_dpxl_pb2.SerializedExpression, _Mapping]] = ...) -> None: ...

class UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstRequest(_message.Message):
    __slots__ = ("serialized",)
    SERIALIZED_FIELD_NUMBER: _ClassVar[int]
    serialized: _dpxl_pb2.SerializedExpression
    def __init__(self, serialized: _Optional[_Union[_dpxl_pb2.SerializedExpression, _Mapping]] = ...) -> None: ...

class UnauthenticatedInternalDpxlServiceToOpenSearchClauseRequest(_message.Message):
    __slots__ = ("serialized", "source")
    SERIALIZED_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    serialized: _dpxl_pb2.SerializedExpression
    source: _dpxl_pb2.Source
    def __init__(self, serialized: _Optional[_Union[_dpxl_pb2.SerializedExpression, _Mapping]] = ..., source: _Optional[_Union[_dpxl_pb2.Source, _Mapping]] = ...) -> None: ...

class UnauthenticatedInternalDpxlServiceCombineRequest(_message.Message):
    __slots__ = ("serialized1", "serialized2", "operation", "entity_types", "entity_type")
    class Operation(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        OPERATION_UNSPECIFIED: _ClassVar[UnauthenticatedInternalDpxlServiceCombineRequest.Operation]
        OPERATION_AND: _ClassVar[UnauthenticatedInternalDpxlServiceCombineRequest.Operation]
        OPERATION_OR: _ClassVar[UnauthenticatedInternalDpxlServiceCombineRequest.Operation]
    OPERATION_UNSPECIFIED: UnauthenticatedInternalDpxlServiceCombineRequest.Operation
    OPERATION_AND: UnauthenticatedInternalDpxlServiceCombineRequest.Operation
    OPERATION_OR: UnauthenticatedInternalDpxlServiceCombineRequest.Operation
    SERIALIZED1_FIELD_NUMBER: _ClassVar[int]
    SERIALIZED2_FIELD_NUMBER: _ClassVar[int]
    OPERATION_FIELD_NUMBER: _ClassVar[int]
    ENTITY_TYPES_FIELD_NUMBER: _ClassVar[int]
    ENTITY_TYPE_FIELD_NUMBER: _ClassVar[int]
    serialized1: _dpxl_pb2.SerializedExpression
    serialized2: _dpxl_pb2.SerializedExpression
    operation: UnauthenticatedInternalDpxlServiceCombineRequest.Operation
    entity_types: _containers.RepeatedCompositeFieldContainer[_dpxl_pb2.EntityType]
    entity_type: _dpxl_pb2.EntityType
    def __init__(self, serialized1: _Optional[_Union[_dpxl_pb2.SerializedExpression, _Mapping]] = ..., serialized2: _Optional[_Union[_dpxl_pb2.SerializedExpression, _Mapping]] = ..., operation: _Optional[_Union[UnauthenticatedInternalDpxlServiceCombineRequest.Operation, str]] = ..., entity_types: _Optional[_Iterable[_Union[_dpxl_pb2.EntityType, _Mapping]]] = ..., entity_type: _Optional[_Union[_dpxl_pb2.EntityType, _Mapping]] = ...) -> None: ...
