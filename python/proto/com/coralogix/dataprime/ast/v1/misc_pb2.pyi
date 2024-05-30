from com.coralogix.dataprime.ast.v1 import expression_pb2 as _expression_pb2
from com.coralogix.dataprime.ast.v1 import types_pb2 as _types_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class KeypathType(_message.Message):
    __slots__ = ("path", "datatype")
    PATH_FIELD_NUMBER: _ClassVar[int]
    DATATYPE_FIELD_NUMBER: _ClassVar[int]
    path: _expression_pb2.Expression.Keypath
    datatype: _types_pb2.Datatype
    def __init__(self, path: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ..., datatype: _Optional[_Union[_types_pb2.Datatype, _Mapping]] = ...) -> None: ...
