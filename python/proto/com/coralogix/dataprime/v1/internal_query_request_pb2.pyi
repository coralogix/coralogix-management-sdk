from com.coralogix.dataprime.v1 import internal_metadata_pb2 as _internal_metadata_pb2
from com.coralogix.dataprime.v1 import query_pb2 as _query_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DpxlScopesPlacementType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    DPXL_SCOPES_PLACEMENT_TYPE_UNSPECIFIED: _ClassVar[DpxlScopesPlacementType]
    DPXL_SCOPES_PLACEMENT_TYPE_FIRST_FILTER: _ClassVar[DpxlScopesPlacementType]
    DPXL_SCOPES_PLACEMENT_TYPE_PLACEHOLDER: _ClassVar[DpxlScopesPlacementType]
    DPXL_SCOPES_PLACEMENT_TYPE_SKIP: _ClassVar[DpxlScopesPlacementType]
DPXL_SCOPES_PLACEMENT_TYPE_UNSPECIFIED: DpxlScopesPlacementType
DPXL_SCOPES_PLACEMENT_TYPE_FIRST_FILTER: DpxlScopesPlacementType
DPXL_SCOPES_PLACEMENT_TYPE_PLACEHOLDER: DpxlScopesPlacementType
DPXL_SCOPES_PLACEMENT_TYPE_SKIP: DpxlScopesPlacementType

class QueryRequest(_message.Message):
    __slots__ = ("ast_query", "text_query", "internal_client_id", "metadata", "query_id", "dpxl_scopes_placement")
    AST_QUERY_FIELD_NUMBER: _ClassVar[int]
    TEXT_QUERY_FIELD_NUMBER: _ClassVar[int]
    INTERNAL_CLIENT_ID_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    DPXL_SCOPES_PLACEMENT_FIELD_NUMBER: _ClassVar[int]
    ast_query: InternalAstQueryRequest
    text_query: InternalTextQueryRequest
    internal_client_id: str
    metadata: _internal_metadata_pb2.QueryMetadata
    query_id: str
    dpxl_scopes_placement: DpxlScopesPlacement
    def __init__(self, ast_query: _Optional[_Union[InternalAstQueryRequest, _Mapping]] = ..., text_query: _Optional[_Union[InternalTextQueryRequest, _Mapping]] = ..., internal_client_id: _Optional[str] = ..., metadata: _Optional[_Union[_internal_metadata_pb2.QueryMetadata, _Mapping]] = ..., query_id: _Optional[str] = ..., dpxl_scopes_placement: _Optional[_Union[DpxlScopesPlacement, _Mapping]] = ...) -> None: ...

class DpxlScopesPlacement(_message.Message):
    __slots__ = ("type", "placeholder")
    TYPE_FIELD_NUMBER: _ClassVar[int]
    PLACEHOLDER_FIELD_NUMBER: _ClassVar[int]
    type: DpxlScopesPlacementType
    placeholder: str
    def __init__(self, type: _Optional[_Union[DpxlScopesPlacementType, str]] = ..., placeholder: _Optional[str] = ...) -> None: ...

class InternalAstQueryRequest(_message.Message):
    __slots__ = ("ast",)
    AST_FIELD_NUMBER: _ClassVar[int]
    ast: _query_pb2.SerializedDataprime
    def __init__(self, ast: _Optional[_Union[_query_pb2.SerializedDataprime, _Mapping]] = ...) -> None: ...

class InternalTextQueryRequest(_message.Message):
    __slots__ = ("text", "metadata")
    TEXT_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    text: _wrappers_pb2.StringValue
    metadata: _internal_metadata_pb2.TextQueryMetadata
    def __init__(self, text: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., metadata: _Optional[_Union[_internal_metadata_pb2.TextQueryMetadata, _Mapping]] = ...) -> None: ...
