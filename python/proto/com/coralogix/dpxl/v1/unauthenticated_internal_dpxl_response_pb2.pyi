from com.coralogix.dpxl.v1 import dpxl_pb2 as _dpxl_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UnauthenticatedInternalDpxlServiceCompileResponse(_message.Message):
    __slots__ = ("success", "failure")
    class Success(_message.Message):
        __slots__ = ("serialized", "warnings")
        SERIALIZED_FIELD_NUMBER: _ClassVar[int]
        WARNINGS_FIELD_NUMBER: _ClassVar[int]
        serialized: _dpxl_pb2.SerializedExpression
        warnings: _containers.RepeatedCompositeFieldContainer[_dpxl_pb2.Diagnostic.Warning]
        def __init__(self, serialized: _Optional[_Union[_dpxl_pb2.SerializedExpression, _Mapping]] = ..., warnings: _Optional[_Iterable[_Union[_dpxl_pb2.Diagnostic.Warning, _Mapping]]] = ...) -> None: ...
    class Failure(_message.Message):
        __slots__ = ("compile_error",)
        COMPILE_ERROR_FIELD_NUMBER: _ClassVar[int]
        compile_error: CompileError
        def __init__(self, compile_error: _Optional[_Union[CompileError, _Mapping]] = ...) -> None: ...
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    FAILURE_FIELD_NUMBER: _ClassVar[int]
    success: UnauthenticatedInternalDpxlServiceCompileResponse.Success
    failure: UnauthenticatedInternalDpxlServiceCompileResponse.Failure
    def __init__(self, success: _Optional[_Union[UnauthenticatedInternalDpxlServiceCompileResponse.Success, _Mapping]] = ..., failure: _Optional[_Union[UnauthenticatedInternalDpxlServiceCompileResponse.Failure, _Mapping]] = ...) -> None: ...

class UnauthenticatedInternalDpxlServiceToDpxlTextResponse(_message.Message):
    __slots__ = ("success",)
    class Success(_message.Message):
        __slots__ = ("expression",)
        EXPRESSION_FIELD_NUMBER: _ClassVar[int]
        expression: _dpxl_pb2.Expression
        def __init__(self, expression: _Optional[_Union[_dpxl_pb2.Expression, _Mapping]] = ...) -> None: ...
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: UnauthenticatedInternalDpxlServiceToDpxlTextResponse.Success
    def __init__(self, success: _Optional[_Union[UnauthenticatedInternalDpxlServiceToDpxlTextResponse.Success, _Mapping]] = ...) -> None: ...

class UnauthenticatedInternalDpxlServiceToOpenSearchClauseResponse(_message.Message):
    __slots__ = ("success", "failure")
    class Success(_message.Message):
        __slots__ = ("clause",)
        CLAUSE_FIELD_NUMBER: _ClassVar[int]
        clause: str
        def __init__(self, clause: _Optional[str] = ...) -> None: ...
    class Failure(_message.Message):
        __slots__ = ("compile_error", "translate_error")
        class TranslateError(_message.Message):
            __slots__ = ("errors",)
            ERRORS_FIELD_NUMBER: _ClassVar[int]
            errors: _containers.RepeatedCompositeFieldContainer[_dpxl_pb2.Diagnostic.Error]
            def __init__(self, errors: _Optional[_Iterable[_Union[_dpxl_pb2.Diagnostic.Error, _Mapping]]] = ...) -> None: ...
        COMPILE_ERROR_FIELD_NUMBER: _ClassVar[int]
        TRANSLATE_ERROR_FIELD_NUMBER: _ClassVar[int]
        compile_error: CompileError
        translate_error: UnauthenticatedInternalDpxlServiceToOpenSearchClauseResponse.Failure.TranslateError
        def __init__(self, compile_error: _Optional[_Union[CompileError, _Mapping]] = ..., translate_error: _Optional[_Union[UnauthenticatedInternalDpxlServiceToOpenSearchClauseResponse.Failure.TranslateError, _Mapping]] = ...) -> None: ...
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    FAILURE_FIELD_NUMBER: _ClassVar[int]
    success: UnauthenticatedInternalDpxlServiceToOpenSearchClauseResponse.Success
    failure: UnauthenticatedInternalDpxlServiceToOpenSearchClauseResponse.Failure
    def __init__(self, success: _Optional[_Union[UnauthenticatedInternalDpxlServiceToOpenSearchClauseResponse.Success, _Mapping]] = ..., failure: _Optional[_Union[UnauthenticatedInternalDpxlServiceToOpenSearchClauseResponse.Failure, _Mapping]] = ...) -> None: ...

class CompileError(_message.Message):
    __slots__ = ("compile_expression_error", "entity_type_validation_error")
    COMPILE_EXPRESSION_ERROR_FIELD_NUMBER: _ClassVar[int]
    ENTITY_TYPE_VALIDATION_ERROR_FIELD_NUMBER: _ClassVar[int]
    compile_expression_error: CompileExpressionError
    entity_type_validation_error: EntityTypeValidationError
    def __init__(self, compile_expression_error: _Optional[_Union[CompileExpressionError, _Mapping]] = ..., entity_type_validation_error: _Optional[_Union[EntityTypeValidationError, _Mapping]] = ...) -> None: ...

class CombineError(_message.Message):
    __slots__ = ("entity_type_validation_error", "incompatible_keypaths_error")
    ENTITY_TYPE_VALIDATION_ERROR_FIELD_NUMBER: _ClassVar[int]
    INCOMPATIBLE_KEYPATHS_ERROR_FIELD_NUMBER: _ClassVar[int]
    entity_type_validation_error: EntityTypeValidationError
    incompatible_keypaths_error: IncompatibleKeypathsError
    def __init__(self, entity_type_validation_error: _Optional[_Union[EntityTypeValidationError, _Mapping]] = ..., incompatible_keypaths_error: _Optional[_Union[IncompatibleKeypathsError, _Mapping]] = ...) -> None: ...

class CompileExpressionError(_message.Message):
    __slots__ = ("errors",)
    ERRORS_FIELD_NUMBER: _ClassVar[int]
    errors: _containers.RepeatedCompositeFieldContainer[_dpxl_pb2.Diagnostic.Error]
    def __init__(self, errors: _Optional[_Iterable[_Union[_dpxl_pb2.Diagnostic.Error, _Mapping]]] = ...) -> None: ...

class IncompatibleKeypathsError(_message.Message):
    __slots__ = ("keypaths",)
    KEYPATHS_FIELD_NUMBER: _ClassVar[int]
    keypaths: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, keypaths: _Optional[_Iterable[str]] = ...) -> None: ...

class EntityTypeValidationError(_message.Message):
    __slots__ = ("errors",)
    class Error(_message.Message):
        __slots__ = ("entity_type", "message", "location")
        ENTITY_TYPE_FIELD_NUMBER: _ClassVar[int]
        MESSAGE_FIELD_NUMBER: _ClassVar[int]
        LOCATION_FIELD_NUMBER: _ClassVar[int]
        entity_type: _dpxl_pb2.EntityType
        message: str
        location: _dpxl_pb2.Diagnostic.Range
        def __init__(self, entity_type: _Optional[_Union[_dpxl_pb2.EntityType, _Mapping]] = ..., message: _Optional[str] = ..., location: _Optional[_Union[_dpxl_pb2.Diagnostic.Range, _Mapping]] = ...) -> None: ...
    ERRORS_FIELD_NUMBER: _ClassVar[int]
    errors: _containers.RepeatedCompositeFieldContainer[EntityTypeValidationError.Error]
    def __init__(self, errors: _Optional[_Iterable[_Union[EntityTypeValidationError.Error, _Mapping]]] = ...) -> None: ...

class UnauthenticatedInternalDpxlServiceCombineResponse(_message.Message):
    __slots__ = ("success", "failure")
    class Success(_message.Message):
        __slots__ = ("serialized",)
        SERIALIZED_FIELD_NUMBER: _ClassVar[int]
        serialized: _dpxl_pb2.SerializedExpression
        def __init__(self, serialized: _Optional[_Union[_dpxl_pb2.SerializedExpression, _Mapping]] = ...) -> None: ...
    class Failure(_message.Message):
        __slots__ = ("compile_error",)
        COMPILE_ERROR_FIELD_NUMBER: _ClassVar[int]
        compile_error: CombineError
        def __init__(self, compile_error: _Optional[_Union[CombineError, _Mapping]] = ...) -> None: ...
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    FAILURE_FIELD_NUMBER: _ClassVar[int]
    success: UnauthenticatedInternalDpxlServiceCombineResponse.Success
    failure: UnauthenticatedInternalDpxlServiceCombineResponse.Failure
    def __init__(self, success: _Optional[_Union[UnauthenticatedInternalDpxlServiceCombineResponse.Success, _Mapping]] = ..., failure: _Optional[_Union[UnauthenticatedInternalDpxlServiceCombineResponse.Failure, _Mapping]] = ...) -> None: ...

class UnauthenticatedInternalDpxlServiceToSerializedDataprimeCoreAstResponse(_message.Message):
    __slots__ = ("serialized_dataprime_core_ast",)
    SERIALIZED_DATAPRIME_CORE_AST_FIELD_NUMBER: _ClassVar[int]
    serialized_dataprime_core_ast: bytes
    def __init__(self, serialized_dataprime_core_ast: _Optional[bytes] = ...) -> None: ...
