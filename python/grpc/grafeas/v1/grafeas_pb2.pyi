from google.api import annotations_pb2 as _annotations_pb2
from google.api import client_pb2 as _client_pb2
from google.api import field_behavior_pb2 as _field_behavior_pb2
from google.api import resource_pb2 as _resource_pb2
from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf import field_mask_pb2 as _field_mask_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from grafeas.v1 import attestation_pb2 as _attestation_pb2
from grafeas.v1 import build_pb2 as _build_pb2
from grafeas.v1 import common_pb2 as _common_pb2
from grafeas.v1 import compliance_pb2 as _compliance_pb2
from grafeas.v1 import deployment_pb2 as _deployment_pb2
from grafeas.v1 import discovery_pb2 as _discovery_pb2
from grafeas.v1 import dsse_attestation_pb2 as _dsse_attestation_pb2
from grafeas.v1 import image_pb2 as _image_pb2
from grafeas.v1 import package_pb2 as _package_pb2
from grafeas.v1 import sbom_pb2 as _sbom_pb2
from grafeas.v1 import upgrade_pb2 as _upgrade_pb2
from grafeas.v1 import vex_pb2 as _vex_pb2
from grafeas.v1 import vulnerability_pb2 as _vulnerability_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Occurrence(_message.Message):
    __slots__ = ("name", "resource_uri", "note_name", "kind", "remediation", "create_time", "update_time", "vulnerability", "build", "image", "package", "deployment", "discovery", "attestation", "upgrade", "compliance", "dsse_attestation", "sbom_reference", "envelope")
    NAME_FIELD_NUMBER: _ClassVar[int]
    RESOURCE_URI_FIELD_NUMBER: _ClassVar[int]
    NOTE_NAME_FIELD_NUMBER: _ClassVar[int]
    KIND_FIELD_NUMBER: _ClassVar[int]
    REMEDIATION_FIELD_NUMBER: _ClassVar[int]
    CREATE_TIME_FIELD_NUMBER: _ClassVar[int]
    UPDATE_TIME_FIELD_NUMBER: _ClassVar[int]
    VULNERABILITY_FIELD_NUMBER: _ClassVar[int]
    BUILD_FIELD_NUMBER: _ClassVar[int]
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    PACKAGE_FIELD_NUMBER: _ClassVar[int]
    DEPLOYMENT_FIELD_NUMBER: _ClassVar[int]
    DISCOVERY_FIELD_NUMBER: _ClassVar[int]
    ATTESTATION_FIELD_NUMBER: _ClassVar[int]
    UPGRADE_FIELD_NUMBER: _ClassVar[int]
    COMPLIANCE_FIELD_NUMBER: _ClassVar[int]
    DSSE_ATTESTATION_FIELD_NUMBER: _ClassVar[int]
    SBOM_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    ENVELOPE_FIELD_NUMBER: _ClassVar[int]
    name: str
    resource_uri: str
    note_name: str
    kind: _common_pb2.NoteKind
    remediation: str
    create_time: _timestamp_pb2.Timestamp
    update_time: _timestamp_pb2.Timestamp
    vulnerability: _vulnerability_pb2.VulnerabilityOccurrence
    build: _build_pb2.BuildOccurrence
    image: _image_pb2.ImageOccurrence
    package: _package_pb2.PackageOccurrence
    deployment: _deployment_pb2.DeploymentOccurrence
    discovery: _discovery_pb2.DiscoveryOccurrence
    attestation: _attestation_pb2.AttestationOccurrence
    upgrade: _upgrade_pb2.UpgradeOccurrence
    compliance: _compliance_pb2.ComplianceOccurrence
    dsse_attestation: _dsse_attestation_pb2.DSSEAttestationOccurrence
    sbom_reference: _sbom_pb2.SBOMReferenceOccurrence
    envelope: _common_pb2.Envelope
    def __init__(self, name: _Optional[str] = ..., resource_uri: _Optional[str] = ..., note_name: _Optional[str] = ..., kind: _Optional[_Union[_common_pb2.NoteKind, str]] = ..., remediation: _Optional[str] = ..., create_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., update_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., vulnerability: _Optional[_Union[_vulnerability_pb2.VulnerabilityOccurrence, _Mapping]] = ..., build: _Optional[_Union[_build_pb2.BuildOccurrence, _Mapping]] = ..., image: _Optional[_Union[_image_pb2.ImageOccurrence, _Mapping]] = ..., package: _Optional[_Union[_package_pb2.PackageOccurrence, _Mapping]] = ..., deployment: _Optional[_Union[_deployment_pb2.DeploymentOccurrence, _Mapping]] = ..., discovery: _Optional[_Union[_discovery_pb2.DiscoveryOccurrence, _Mapping]] = ..., attestation: _Optional[_Union[_attestation_pb2.AttestationOccurrence, _Mapping]] = ..., upgrade: _Optional[_Union[_upgrade_pb2.UpgradeOccurrence, _Mapping]] = ..., compliance: _Optional[_Union[_compliance_pb2.ComplianceOccurrence, _Mapping]] = ..., dsse_attestation: _Optional[_Union[_dsse_attestation_pb2.DSSEAttestationOccurrence, _Mapping]] = ..., sbom_reference: _Optional[_Union[_sbom_pb2.SBOMReferenceOccurrence, _Mapping]] = ..., envelope: _Optional[_Union[_common_pb2.Envelope, _Mapping]] = ...) -> None: ...

class Note(_message.Message):
    __slots__ = ("name", "short_description", "long_description", "kind", "related_url", "expiration_time", "create_time", "update_time", "related_note_names", "vulnerability", "build", "image", "package", "deployment", "discovery", "attestation", "upgrade", "compliance", "dsse_attestation", "vulnerability_assessment", "sbom_reference")
    NAME_FIELD_NUMBER: _ClassVar[int]
    SHORT_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    LONG_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    KIND_FIELD_NUMBER: _ClassVar[int]
    RELATED_URL_FIELD_NUMBER: _ClassVar[int]
    EXPIRATION_TIME_FIELD_NUMBER: _ClassVar[int]
    CREATE_TIME_FIELD_NUMBER: _ClassVar[int]
    UPDATE_TIME_FIELD_NUMBER: _ClassVar[int]
    RELATED_NOTE_NAMES_FIELD_NUMBER: _ClassVar[int]
    VULNERABILITY_FIELD_NUMBER: _ClassVar[int]
    BUILD_FIELD_NUMBER: _ClassVar[int]
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    PACKAGE_FIELD_NUMBER: _ClassVar[int]
    DEPLOYMENT_FIELD_NUMBER: _ClassVar[int]
    DISCOVERY_FIELD_NUMBER: _ClassVar[int]
    ATTESTATION_FIELD_NUMBER: _ClassVar[int]
    UPGRADE_FIELD_NUMBER: _ClassVar[int]
    COMPLIANCE_FIELD_NUMBER: _ClassVar[int]
    DSSE_ATTESTATION_FIELD_NUMBER: _ClassVar[int]
    VULNERABILITY_ASSESSMENT_FIELD_NUMBER: _ClassVar[int]
    SBOM_REFERENCE_FIELD_NUMBER: _ClassVar[int]
    name: str
    short_description: str
    long_description: str
    kind: _common_pb2.NoteKind
    related_url: _containers.RepeatedCompositeFieldContainer[_common_pb2.RelatedUrl]
    expiration_time: _timestamp_pb2.Timestamp
    create_time: _timestamp_pb2.Timestamp
    update_time: _timestamp_pb2.Timestamp
    related_note_names: _containers.RepeatedScalarFieldContainer[str]
    vulnerability: _vulnerability_pb2.VulnerabilityNote
    build: _build_pb2.BuildNote
    image: _image_pb2.ImageNote
    package: _package_pb2.PackageNote
    deployment: _deployment_pb2.DeploymentNote
    discovery: _discovery_pb2.DiscoveryNote
    attestation: _attestation_pb2.AttestationNote
    upgrade: _upgrade_pb2.UpgradeNote
    compliance: _compliance_pb2.ComplianceNote
    dsse_attestation: _dsse_attestation_pb2.DSSEAttestationNote
    vulnerability_assessment: _vex_pb2.VulnerabilityAssessmentNote
    sbom_reference: _sbom_pb2.SBOMReferenceNote
    def __init__(self, name: _Optional[str] = ..., short_description: _Optional[str] = ..., long_description: _Optional[str] = ..., kind: _Optional[_Union[_common_pb2.NoteKind, str]] = ..., related_url: _Optional[_Iterable[_Union[_common_pb2.RelatedUrl, _Mapping]]] = ..., expiration_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., create_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., update_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., related_note_names: _Optional[_Iterable[str]] = ..., vulnerability: _Optional[_Union[_vulnerability_pb2.VulnerabilityNote, _Mapping]] = ..., build: _Optional[_Union[_build_pb2.BuildNote, _Mapping]] = ..., image: _Optional[_Union[_image_pb2.ImageNote, _Mapping]] = ..., package: _Optional[_Union[_package_pb2.PackageNote, _Mapping]] = ..., deployment: _Optional[_Union[_deployment_pb2.DeploymentNote, _Mapping]] = ..., discovery: _Optional[_Union[_discovery_pb2.DiscoveryNote, _Mapping]] = ..., attestation: _Optional[_Union[_attestation_pb2.AttestationNote, _Mapping]] = ..., upgrade: _Optional[_Union[_upgrade_pb2.UpgradeNote, _Mapping]] = ..., compliance: _Optional[_Union[_compliance_pb2.ComplianceNote, _Mapping]] = ..., dsse_attestation: _Optional[_Union[_dsse_attestation_pb2.DSSEAttestationNote, _Mapping]] = ..., vulnerability_assessment: _Optional[_Union[_vex_pb2.VulnerabilityAssessmentNote, _Mapping]] = ..., sbom_reference: _Optional[_Union[_sbom_pb2.SBOMReferenceNote, _Mapping]] = ...) -> None: ...

class GetOccurrenceRequest(_message.Message):
    __slots__ = ("name",)
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class ListOccurrencesRequest(_message.Message):
    __slots__ = ("parent", "filter", "page_size", "page_token")
    PARENT_FIELD_NUMBER: _ClassVar[int]
    FILTER_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    PAGE_TOKEN_FIELD_NUMBER: _ClassVar[int]
    parent: str
    filter: str
    page_size: int
    page_token: str
    def __init__(self, parent: _Optional[str] = ..., filter: _Optional[str] = ..., page_size: _Optional[int] = ..., page_token: _Optional[str] = ...) -> None: ...

class ListOccurrencesResponse(_message.Message):
    __slots__ = ("occurrences", "next_page_token")
    OCCURRENCES_FIELD_NUMBER: _ClassVar[int]
    NEXT_PAGE_TOKEN_FIELD_NUMBER: _ClassVar[int]
    occurrences: _containers.RepeatedCompositeFieldContainer[Occurrence]
    next_page_token: str
    def __init__(self, occurrences: _Optional[_Iterable[_Union[Occurrence, _Mapping]]] = ..., next_page_token: _Optional[str] = ...) -> None: ...

class DeleteOccurrenceRequest(_message.Message):
    __slots__ = ("name",)
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class CreateOccurrenceRequest(_message.Message):
    __slots__ = ("parent", "occurrence")
    PARENT_FIELD_NUMBER: _ClassVar[int]
    OCCURRENCE_FIELD_NUMBER: _ClassVar[int]
    parent: str
    occurrence: Occurrence
    def __init__(self, parent: _Optional[str] = ..., occurrence: _Optional[_Union[Occurrence, _Mapping]] = ...) -> None: ...

class UpdateOccurrenceRequest(_message.Message):
    __slots__ = ("name", "occurrence", "update_mask")
    NAME_FIELD_NUMBER: _ClassVar[int]
    OCCURRENCE_FIELD_NUMBER: _ClassVar[int]
    UPDATE_MASK_FIELD_NUMBER: _ClassVar[int]
    name: str
    occurrence: Occurrence
    update_mask: _field_mask_pb2.FieldMask
    def __init__(self, name: _Optional[str] = ..., occurrence: _Optional[_Union[Occurrence, _Mapping]] = ..., update_mask: _Optional[_Union[_field_mask_pb2.FieldMask, _Mapping]] = ...) -> None: ...

class GetNoteRequest(_message.Message):
    __slots__ = ("name",)
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class GetOccurrenceNoteRequest(_message.Message):
    __slots__ = ("name",)
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class ListNotesRequest(_message.Message):
    __slots__ = ("parent", "filter", "page_size", "page_token")
    PARENT_FIELD_NUMBER: _ClassVar[int]
    FILTER_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    PAGE_TOKEN_FIELD_NUMBER: _ClassVar[int]
    parent: str
    filter: str
    page_size: int
    page_token: str
    def __init__(self, parent: _Optional[str] = ..., filter: _Optional[str] = ..., page_size: _Optional[int] = ..., page_token: _Optional[str] = ...) -> None: ...

class ListNotesResponse(_message.Message):
    __slots__ = ("notes", "next_page_token")
    NOTES_FIELD_NUMBER: _ClassVar[int]
    NEXT_PAGE_TOKEN_FIELD_NUMBER: _ClassVar[int]
    notes: _containers.RepeatedCompositeFieldContainer[Note]
    next_page_token: str
    def __init__(self, notes: _Optional[_Iterable[_Union[Note, _Mapping]]] = ..., next_page_token: _Optional[str] = ...) -> None: ...

class DeleteNoteRequest(_message.Message):
    __slots__ = ("name",)
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class CreateNoteRequest(_message.Message):
    __slots__ = ("parent", "note_id", "note")
    PARENT_FIELD_NUMBER: _ClassVar[int]
    NOTE_ID_FIELD_NUMBER: _ClassVar[int]
    NOTE_FIELD_NUMBER: _ClassVar[int]
    parent: str
    note_id: str
    note: Note
    def __init__(self, parent: _Optional[str] = ..., note_id: _Optional[str] = ..., note: _Optional[_Union[Note, _Mapping]] = ...) -> None: ...

class UpdateNoteRequest(_message.Message):
    __slots__ = ("name", "note", "update_mask")
    NAME_FIELD_NUMBER: _ClassVar[int]
    NOTE_FIELD_NUMBER: _ClassVar[int]
    UPDATE_MASK_FIELD_NUMBER: _ClassVar[int]
    name: str
    note: Note
    update_mask: _field_mask_pb2.FieldMask
    def __init__(self, name: _Optional[str] = ..., note: _Optional[_Union[Note, _Mapping]] = ..., update_mask: _Optional[_Union[_field_mask_pb2.FieldMask, _Mapping]] = ...) -> None: ...

class ListNoteOccurrencesRequest(_message.Message):
    __slots__ = ("name", "filter", "page_size", "page_token")
    NAME_FIELD_NUMBER: _ClassVar[int]
    FILTER_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    PAGE_TOKEN_FIELD_NUMBER: _ClassVar[int]
    name: str
    filter: str
    page_size: int
    page_token: str
    def __init__(self, name: _Optional[str] = ..., filter: _Optional[str] = ..., page_size: _Optional[int] = ..., page_token: _Optional[str] = ...) -> None: ...

class ListNoteOccurrencesResponse(_message.Message):
    __slots__ = ("occurrences", "next_page_token")
    OCCURRENCES_FIELD_NUMBER: _ClassVar[int]
    NEXT_PAGE_TOKEN_FIELD_NUMBER: _ClassVar[int]
    occurrences: _containers.RepeatedCompositeFieldContainer[Occurrence]
    next_page_token: str
    def __init__(self, occurrences: _Optional[_Iterable[_Union[Occurrence, _Mapping]]] = ..., next_page_token: _Optional[str] = ...) -> None: ...

class BatchCreateNotesRequest(_message.Message):
    __slots__ = ("parent", "notes")
    class NotesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: Note
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[Note, _Mapping]] = ...) -> None: ...
    PARENT_FIELD_NUMBER: _ClassVar[int]
    NOTES_FIELD_NUMBER: _ClassVar[int]
    parent: str
    notes: _containers.MessageMap[str, Note]
    def __init__(self, parent: _Optional[str] = ..., notes: _Optional[_Mapping[str, Note]] = ...) -> None: ...

class BatchCreateNotesResponse(_message.Message):
    __slots__ = ("notes",)
    NOTES_FIELD_NUMBER: _ClassVar[int]
    notes: _containers.RepeatedCompositeFieldContainer[Note]
    def __init__(self, notes: _Optional[_Iterable[_Union[Note, _Mapping]]] = ...) -> None: ...

class BatchCreateOccurrencesRequest(_message.Message):
    __slots__ = ("parent", "occurrences")
    PARENT_FIELD_NUMBER: _ClassVar[int]
    OCCURRENCES_FIELD_NUMBER: _ClassVar[int]
    parent: str
    occurrences: _containers.RepeatedCompositeFieldContainer[Occurrence]
    def __init__(self, parent: _Optional[str] = ..., occurrences: _Optional[_Iterable[_Union[Occurrence, _Mapping]]] = ...) -> None: ...

class BatchCreateOccurrencesResponse(_message.Message):
    __slots__ = ("occurrences",)
    OCCURRENCES_FIELD_NUMBER: _ClassVar[int]
    occurrences: _containers.RepeatedCompositeFieldContainer[Occurrence]
    def __init__(self, occurrences: _Optional[_Iterable[_Union[Occurrence, _Mapping]]] = ...) -> None: ...
