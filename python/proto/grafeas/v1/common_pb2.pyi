from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class NoteKind(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    NOTE_KIND_UNSPECIFIED: _ClassVar[NoteKind]
    VULNERABILITY: _ClassVar[NoteKind]
    BUILD: _ClassVar[NoteKind]
    IMAGE: _ClassVar[NoteKind]
    PACKAGE: _ClassVar[NoteKind]
    DEPLOYMENT: _ClassVar[NoteKind]
    DISCOVERY: _ClassVar[NoteKind]
    ATTESTATION: _ClassVar[NoteKind]
    UPGRADE: _ClassVar[NoteKind]
    COMPLIANCE: _ClassVar[NoteKind]
    DSSE_ATTESTATION: _ClassVar[NoteKind]
    VULNERABILITY_ASSESSMENT: _ClassVar[NoteKind]
    SBOM_REFERENCE: _ClassVar[NoteKind]
NOTE_KIND_UNSPECIFIED: NoteKind
VULNERABILITY: NoteKind
BUILD: NoteKind
IMAGE: NoteKind
PACKAGE: NoteKind
DEPLOYMENT: NoteKind
DISCOVERY: NoteKind
ATTESTATION: NoteKind
UPGRADE: NoteKind
COMPLIANCE: NoteKind
DSSE_ATTESTATION: NoteKind
VULNERABILITY_ASSESSMENT: NoteKind
SBOM_REFERENCE: NoteKind

class RelatedUrl(_message.Message):
    __slots__ = ("url", "label")
    URL_FIELD_NUMBER: _ClassVar[int]
    LABEL_FIELD_NUMBER: _ClassVar[int]
    url: str
    label: str
    def __init__(self, url: _Optional[str] = ..., label: _Optional[str] = ...) -> None: ...

class Signature(_message.Message):
    __slots__ = ("signature", "public_key_id")
    SIGNATURE_FIELD_NUMBER: _ClassVar[int]
    PUBLIC_KEY_ID_FIELD_NUMBER: _ClassVar[int]
    signature: bytes
    public_key_id: str
    def __init__(self, signature: _Optional[bytes] = ..., public_key_id: _Optional[str] = ...) -> None: ...

class Envelope(_message.Message):
    __slots__ = ("payload", "payload_type", "signatures")
    PAYLOAD_FIELD_NUMBER: _ClassVar[int]
    PAYLOAD_TYPE_FIELD_NUMBER: _ClassVar[int]
    SIGNATURES_FIELD_NUMBER: _ClassVar[int]
    payload: bytes
    payload_type: str
    signatures: _containers.RepeatedCompositeFieldContainer[EnvelopeSignature]
    def __init__(self, payload: _Optional[bytes] = ..., payload_type: _Optional[str] = ..., signatures: _Optional[_Iterable[_Union[EnvelopeSignature, _Mapping]]] = ...) -> None: ...

class EnvelopeSignature(_message.Message):
    __slots__ = ("sig", "keyid")
    SIG_FIELD_NUMBER: _ClassVar[int]
    KEYID_FIELD_NUMBER: _ClassVar[int]
    sig: bytes
    keyid: str
    def __init__(self, sig: _Optional[bytes] = ..., keyid: _Optional[str] = ...) -> None: ...

class FileLocation(_message.Message):
    __slots__ = ("file_path",)
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    file_path: str
    def __init__(self, file_path: _Optional[str] = ...) -> None: ...

class License(_message.Message):
    __slots__ = ("expression", "comments")
    EXPRESSION_FIELD_NUMBER: _ClassVar[int]
    COMMENTS_FIELD_NUMBER: _ClassVar[int]
    expression: str
    comments: str
    def __init__(self, expression: _Optional[str] = ..., comments: _Optional[str] = ...) -> None: ...

class Digest(_message.Message):
    __slots__ = ("algo", "digest_bytes")
    ALGO_FIELD_NUMBER: _ClassVar[int]
    DIGEST_BYTES_FIELD_NUMBER: _ClassVar[int]
    algo: str
    digest_bytes: bytes
    def __init__(self, algo: _Optional[str] = ..., digest_bytes: _Optional[bytes] = ...) -> None: ...
