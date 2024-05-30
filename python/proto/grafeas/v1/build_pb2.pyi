from grafeas.v1 import intoto_provenance_pb2 as _intoto_provenance_pb2
from grafeas.v1 import intoto_statement_pb2 as _intoto_statement_pb2
from grafeas.v1 import provenance_pb2 as _provenance_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class BuildNote(_message.Message):
    __slots__ = ("builder_version",)
    BUILDER_VERSION_FIELD_NUMBER: _ClassVar[int]
    builder_version: str
    def __init__(self, builder_version: _Optional[str] = ...) -> None: ...

class BuildOccurrence(_message.Message):
    __slots__ = ("provenance", "provenance_bytes", "intoto_provenance", "intoto_statement", "in_toto_slsa_provenance_v1")
    PROVENANCE_FIELD_NUMBER: _ClassVar[int]
    PROVENANCE_BYTES_FIELD_NUMBER: _ClassVar[int]
    INTOTO_PROVENANCE_FIELD_NUMBER: _ClassVar[int]
    INTOTO_STATEMENT_FIELD_NUMBER: _ClassVar[int]
    IN_TOTO_SLSA_PROVENANCE_V1_FIELD_NUMBER: _ClassVar[int]
    provenance: _provenance_pb2.BuildProvenance
    provenance_bytes: str
    intoto_provenance: _intoto_provenance_pb2.InTotoProvenance
    intoto_statement: _intoto_statement_pb2.InTotoStatement
    in_toto_slsa_provenance_v1: _intoto_statement_pb2.InTotoSlsaProvenanceV1
    def __init__(self, provenance: _Optional[_Union[_provenance_pb2.BuildProvenance, _Mapping]] = ..., provenance_bytes: _Optional[str] = ..., intoto_provenance: _Optional[_Union[_intoto_provenance_pb2.InTotoProvenance, _Mapping]] = ..., intoto_statement: _Optional[_Union[_intoto_statement_pb2.InTotoStatement, _Mapping]] = ..., in_toto_slsa_provenance_v1: _Optional[_Union[_intoto_statement_pb2.InTotoSlsaProvenanceV1, _Mapping]] = ...) -> None: ...
