from com.coralogix.dataprime.ast.restricted.v1 import expression_pb2 as _expression_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Ast(_message.Message):
    __slots__ = ("expression",)
    EXPRESSION_FIELD_NUMBER: _ClassVar[int]
    expression: _expression_pb2.Expression
    def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ...) -> None: ...
