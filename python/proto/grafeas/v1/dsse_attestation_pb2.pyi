from grafeas.v1 import common_pb2 as _common_pb2
from grafeas.v1 import intoto_statement_pb2 as _intoto_statement_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DSSEAttestationNote(_message.Message):
    __slots__ = ("hint",)
    class DSSEHint(_message.Message):
        __slots__ = ("human_readable_name",)
        HUMAN_READABLE_NAME_FIELD_NUMBER: _ClassVar[int]
        human_readable_name: str
        def __init__(self, human_readable_name: _Optional[str] = ...) -> None: ...
    HINT_FIELD_NUMBER: _ClassVar[int]
    hint: DSSEAttestationNote.DSSEHint
    def __init__(self, hint: _Optional[_Union[DSSEAttestationNote.DSSEHint, _Mapping]] = ...) -> None: ...

class DSSEAttestationOccurrence(_message.Message):
    __slots__ = ("envelope", "statement")
    ENVELOPE_FIELD_NUMBER: _ClassVar[int]
    STATEMENT_FIELD_NUMBER: _ClassVar[int]
    envelope: _common_pb2.Envelope
    statement: _intoto_statement_pb2.InTotoStatement
    def __init__(self, envelope: _Optional[_Union[_common_pb2.Envelope, _Mapping]] = ..., statement: _Optional[_Union[_intoto_statement_pb2.InTotoStatement, _Mapping]] = ...) -> None: ...
