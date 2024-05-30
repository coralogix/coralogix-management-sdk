from com.coralogixapis.dataprime.v1 import response_pb2 as _response_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UnauthenticatedInternalDataprimeQueryServiceQueryResponse(_message.Message):
    __slots__ = ("error", "result", "warning", "query_id")
    ERROR_FIELD_NUMBER: _ClassVar[int]
    RESULT_FIELD_NUMBER: _ClassVar[int]
    WARNING_FIELD_NUMBER: _ClassVar[int]
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    error: _response_pb2.DataprimeError
    result: _response_pb2.DataprimeResult
    warning: _response_pb2.DataprimeWarning
    query_id: _response_pb2.QueryId
    def __init__(self, error: _Optional[_Union[_response_pb2.DataprimeError, _Mapping]] = ..., result: _Optional[_Union[_response_pb2.DataprimeResult, _Mapping]] = ..., warning: _Optional[_Union[_response_pb2.DataprimeWarning, _Mapping]] = ..., query_id: _Optional[_Union[_response_pb2.QueryId, _Mapping]] = ...) -> None: ...
