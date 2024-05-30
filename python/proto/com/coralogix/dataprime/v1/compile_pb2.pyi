from com.coralogix.dataprime.ast.v1 import ast_pb2 as _ast_pb2
from com.coralogix.dataprime.ast.v1 import query_pb2 as _query_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.dataprime.v1 import schema_pb2 as _schema_pb2
from com.coralogix.dataprime.v1 import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class QuerySyntax(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    QUERY_SYNTAX_DATAPRIME_UNSPECIFIED: _ClassVar[QuerySyntax]
    QUERY_SYNTAX_LUCENE: _ClassVar[QuerySyntax]
QUERY_SYNTAX_DATAPRIME_UNSPECIFIED: QuerySyntax
QUERY_SYNTAX_LUCENE: QuerySyntax

class CompileRequest(_message.Message):
    __slots__ = ("source_code", "team_id", "time_range", "default_source", "syntax", "available_sources", "override_schema", "backend")
    SOURCE_CODE_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    TIME_RANGE_FIELD_NUMBER: _ClassVar[int]
    DEFAULT_SOURCE_FIELD_NUMBER: _ClassVar[int]
    SYNTAX_FIELD_NUMBER: _ClassVar[int]
    AVAILABLE_SOURCES_FIELD_NUMBER: _ClassVar[int]
    OVERRIDE_SCHEMA_FIELD_NUMBER: _ClassVar[int]
    BACKEND_FIELD_NUMBER: _ClassVar[int]
    source_code: _wrappers_pb2.StringValue
    team_id: int
    time_range: _ast_pb2.TimeRange
    default_source: _wrappers_pb2.StringValue
    syntax: QuerySyntax
    available_sources: _containers.RepeatedCompositeFieldContainer[_query_pb2.Source]
    override_schema: _schema_pb2.SchemaLookup
    backend: _common_pb2.Backend
    def __init__(self, source_code: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., team_id: _Optional[int] = ..., time_range: _Optional[_Union[_ast_pb2.TimeRange, _Mapping]] = ..., default_source: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., syntax: _Optional[_Union[QuerySyntax, str]] = ..., available_sources: _Optional[_Iterable[_Union[_query_pb2.Source, _Mapping]]] = ..., override_schema: _Optional[_Union[_schema_pb2.SchemaLookup, _Mapping]] = ..., backend: _Optional[_Union[_common_pb2.Backend, str]] = ...) -> None: ...

class CompileResponse(_message.Message):
    __slots__ = ("response",)
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    response: CompilationResponse
    def __init__(self, response: _Optional[_Union[CompilationResponse, _Mapping]] = ...) -> None: ...

class CheckRequest(_message.Message):
    __slots__ = ("ast",)
    AST_FIELD_NUMBER: _ClassVar[int]
    ast: _ast_pb2.Ast
    def __init__(self, ast: _Optional[_Union[_ast_pb2.Ast, _Mapping]] = ...) -> None: ...

class CheckResponse(_message.Message):
    __slots__ = ("response",)
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    response: CompilationResponse
    def __init__(self, response: _Optional[_Union[CompilationResponse, _Mapping]] = ...) -> None: ...

class CompilationResponse(_message.Message):
    __slots__ = ("success", "failure")
    class CompilationSuccessful(_message.Message):
        __slots__ = ("ast", "warnings")
        AST_FIELD_NUMBER: _ClassVar[int]
        WARNINGS_FIELD_NUMBER: _ClassVar[int]
        ast: _ast_pb2.Ast
        warnings: _containers.RepeatedCompositeFieldContainer[CompilationResponse.Warning]
        def __init__(self, ast: _Optional[_Union[_ast_pb2.Ast, _Mapping]] = ..., warnings: _Optional[_Iterable[_Union[CompilationResponse.Warning, _Mapping]]] = ...) -> None: ...
    class CompilationFailure(_message.Message):
        __slots__ = ("problems",)
        PROBLEMS_FIELD_NUMBER: _ClassVar[int]
        problems: _containers.RepeatedCompositeFieldContainer[CompilationResponse.Problem]
        def __init__(self, problems: _Optional[_Iterable[_Union[CompilationResponse.Problem, _Mapping]]] = ...) -> None: ...
    class Problem(_message.Message):
        __slots__ = ("error", "warning")
        ERROR_FIELD_NUMBER: _ClassVar[int]
        WARNING_FIELD_NUMBER: _ClassVar[int]
        error: CompilationResponse.Error
        warning: CompilationResponse.Warning
        def __init__(self, error: _Optional[_Union[CompilationResponse.Error, _Mapping]] = ..., warning: _Optional[_Union[CompilationResponse.Warning, _Mapping]] = ...) -> None: ...
    class Error(_message.Message):
        __slots__ = ("msg", "position")
        MSG_FIELD_NUMBER: _ClassVar[int]
        POSITION_FIELD_NUMBER: _ClassVar[int]
        msg: _wrappers_pb2.StringValue
        position: CompilationResponse.Position
        def __init__(self, msg: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., position: _Optional[_Union[CompilationResponse.Position, _Mapping]] = ...) -> None: ...
    class Warning(_message.Message):
        __slots__ = ("msg", "position")
        MSG_FIELD_NUMBER: _ClassVar[int]
        POSITION_FIELD_NUMBER: _ClassVar[int]
        msg: _wrappers_pb2.StringValue
        position: CompilationResponse.Position
        def __init__(self, msg: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., position: _Optional[_Union[CompilationResponse.Position, _Mapping]] = ...) -> None: ...
    class Position(_message.Message):
        __slots__ = ("line", "col")
        LINE_FIELD_NUMBER: _ClassVar[int]
        COL_FIELD_NUMBER: _ClassVar[int]
        line: _wrappers_pb2.UInt32Value
        col: _wrappers_pb2.UInt32Value
        def __init__(self, line: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., col: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    FAILURE_FIELD_NUMBER: _ClassVar[int]
    success: CompilationResponse.CompilationSuccessful
    failure: CompilationResponse.CompilationFailure
    def __init__(self, success: _Optional[_Union[CompilationResponse.CompilationSuccessful, _Mapping]] = ..., failure: _Optional[_Union[CompilationResponse.CompilationFailure, _Mapping]] = ...) -> None: ...
