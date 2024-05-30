from grafeas.v1 import common_pb2 as _common_pb2
from grafeas.v1 import intoto_statement_pb2 as _intoto_statement_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SBOMReferenceNote(_message.Message):
    __slots__ = ("format", "version")
    FORMAT_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    format: str
    version: str
    def __init__(self, format: _Optional[str] = ..., version: _Optional[str] = ...) -> None: ...

class SBOMReferenceOccurrence(_message.Message):
    __slots__ = ("payload", "payload_type", "signatures")
    PAYLOAD_FIELD_NUMBER: _ClassVar[int]
    PAYLOAD_TYPE_FIELD_NUMBER: _ClassVar[int]
    SIGNATURES_FIELD_NUMBER: _ClassVar[int]
    payload: SbomReferenceIntotoPayload
    payload_type: str
    signatures: _containers.RepeatedCompositeFieldContainer[_common_pb2.EnvelopeSignature]
    def __init__(self, payload: _Optional[_Union[SbomReferenceIntotoPayload, _Mapping]] = ..., payload_type: _Optional[str] = ..., signatures: _Optional[_Iterable[_Union[_common_pb2.EnvelopeSignature, _Mapping]]] = ...) -> None: ...

class SbomReferenceIntotoPayload(_message.Message):
    __slots__ = ("type", "predicate_type", "subject", "predicate")
    TYPE_FIELD_NUMBER: _ClassVar[int]
    PREDICATE_TYPE_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_FIELD_NUMBER: _ClassVar[int]
    PREDICATE_FIELD_NUMBER: _ClassVar[int]
    type: str
    predicate_type: str
    subject: _containers.RepeatedCompositeFieldContainer[_intoto_statement_pb2.Subject]
    predicate: SbomReferenceIntotoPredicate
    def __init__(self, type: _Optional[str] = ..., predicate_type: _Optional[str] = ..., subject: _Optional[_Iterable[_Union[_intoto_statement_pb2.Subject, _Mapping]]] = ..., predicate: _Optional[_Union[SbomReferenceIntotoPredicate, _Mapping]] = ...) -> None: ...

class SbomReferenceIntotoPredicate(_message.Message):
    __slots__ = ("referrer_id", "location", "mime_type", "digest")
    class DigestEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    REFERRER_ID_FIELD_NUMBER: _ClassVar[int]
    LOCATION_FIELD_NUMBER: _ClassVar[int]
    MIME_TYPE_FIELD_NUMBER: _ClassVar[int]
    DIGEST_FIELD_NUMBER: _ClassVar[int]
    referrer_id: str
    location: str
    mime_type: str
    digest: _containers.ScalarMap[str, str]
    def __init__(self, referrer_id: _Optional[str] = ..., location: _Optional[str] = ..., mime_type: _Optional[str] = ..., digest: _Optional[_Mapping[str, str]] = ...) -> None: ...
