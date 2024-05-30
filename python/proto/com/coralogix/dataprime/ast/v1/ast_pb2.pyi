from google.protobuf import timestamp_pb2 as _timestamp_pb2
from com.coralogix.dataprime.ast.v1 import query_pb2 as _query_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Ast(_message.Message):
    __slots__ = ("query", "time_range")
    QUERY_FIELD_NUMBER: _ClassVar[int]
    TIME_RANGE_FIELD_NUMBER: _ClassVar[int]
    query: _query_pb2.Query
    time_range: TimeRange
    def __init__(self, query: _Optional[_Union[_query_pb2.Query, _Mapping]] = ..., time_range: _Optional[_Union[TimeRange, _Mapping]] = ...) -> None: ...

class DdlAst(_message.Message):
    __slots__ = ("ast", "dataset_name")
    AST_FIELD_NUMBER: _ClassVar[int]
    DATASET_NAME_FIELD_NUMBER: _ClassVar[int]
    ast: Ast
    dataset_name: str
    def __init__(self, ast: _Optional[_Union[Ast, _Mapping]] = ..., dataset_name: _Optional[str] = ...) -> None: ...

class TimeRange(_message.Message):
    __slots__ = ("to",)
    FROM_FIELD_NUMBER: _ClassVar[int]
    TO_FIELD_NUMBER: _ClassVar[int]
    to: _timestamp_pb2.Timestamp
    def __init__(self, to: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., **kwargs) -> None: ...
