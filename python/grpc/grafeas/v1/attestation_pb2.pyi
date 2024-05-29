from grafeas.v1 import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AttestationNote(_message.Message):
    __slots__ = ("hint",)
    class Hint(_message.Message):
        __slots__ = ("human_readable_name",)
        HUMAN_READABLE_NAME_FIELD_NUMBER: _ClassVar[int]
        human_readable_name: str
        def __init__(self, human_readable_name: _Optional[str] = ...) -> None: ...
    HINT_FIELD_NUMBER: _ClassVar[int]
    hint: AttestationNote.Hint
    def __init__(self, hint: _Optional[_Union[AttestationNote.Hint, _Mapping]] = ...) -> None: ...

class Jwt(_message.Message):
    __slots__ = ("compact_jwt",)
    COMPACT_JWT_FIELD_NUMBER: _ClassVar[int]
    compact_jwt: str
    def __init__(self, compact_jwt: _Optional[str] = ...) -> None: ...

class AttestationOccurrence(_message.Message):
    __slots__ = ("serialized_payload", "signatures", "jwts")
    SERIALIZED_PAYLOAD_FIELD_NUMBER: _ClassVar[int]
    SIGNATURES_FIELD_NUMBER: _ClassVar[int]
    JWTS_FIELD_NUMBER: _ClassVar[int]
    serialized_payload: bytes
    signatures: _containers.RepeatedCompositeFieldContainer[_common_pb2.Signature]
    jwts: _containers.RepeatedCompositeFieldContainer[Jwt]
    def __init__(self, serialized_payload: _Optional[bytes] = ..., signatures: _Optional[_Iterable[_Union[_common_pb2.Signature, _Mapping]]] = ..., jwts: _Optional[_Iterable[_Union[Jwt, _Mapping]]] = ...) -> None: ...
