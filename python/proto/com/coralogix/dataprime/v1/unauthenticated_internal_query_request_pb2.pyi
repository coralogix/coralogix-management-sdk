from com.coralogix.dataprime.v1 import internal_metadata_pb2 as _internal_metadata_pb2
from com.coralogix.dataprime.v1 import internal_query_request_pb2 as _internal_query_request_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UnauthenticatedInternalDataprimeQueryServiceQueryRequest(_message.Message):
    __slots__ = ("ast_query", "text_query", "internal_client_id", "metadata", "team_id", "user_id", "query_id")
    AST_QUERY_FIELD_NUMBER: _ClassVar[int]
    TEXT_QUERY_FIELD_NUMBER: _ClassVar[int]
    INTERNAL_CLIENT_ID_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    ast_query: _internal_query_request_pb2.InternalAstQueryRequest
    text_query: _internal_query_request_pb2.InternalTextQueryRequest
    internal_client_id: str
    metadata: _internal_metadata_pb2.QueryMetadata
    team_id: _wrappers_pb2.StringValue
    user_id: _wrappers_pb2.StringValue
    query_id: str
    def __init__(self, ast_query: _Optional[_Union[_internal_query_request_pb2.InternalAstQueryRequest, _Mapping]] = ..., text_query: _Optional[_Union[_internal_query_request_pb2.InternalTextQueryRequest, _Mapping]] = ..., internal_client_id: _Optional[str] = ..., metadata: _Optional[_Union[_internal_metadata_pb2.QueryMetadata, _Mapping]] = ..., team_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., user_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., query_id: _Optional[str] = ...) -> None: ...
