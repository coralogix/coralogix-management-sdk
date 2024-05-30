from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from grafeas.v1 import intoto_provenance_pb2 as _intoto_provenance_pb2
from grafeas.v1 import slsa_provenance_pb2 as _slsa_provenance_pb2
from grafeas.v1 import slsa_provenance_zero_two_pb2 as _slsa_provenance_zero_two_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class InTotoStatement(_message.Message):
    __slots__ = ("type", "subject", "predicate_type", "provenance", "slsa_provenance", "slsa_provenance_zero_two")
    TYPE_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_FIELD_NUMBER: _ClassVar[int]
    PREDICATE_TYPE_FIELD_NUMBER: _ClassVar[int]
    PROVENANCE_FIELD_NUMBER: _ClassVar[int]
    SLSA_PROVENANCE_FIELD_NUMBER: _ClassVar[int]
    SLSA_PROVENANCE_ZERO_TWO_FIELD_NUMBER: _ClassVar[int]
    type: str
    subject: _containers.RepeatedCompositeFieldContainer[Subject]
    predicate_type: str
    provenance: _intoto_provenance_pb2.InTotoProvenance
    slsa_provenance: _slsa_provenance_pb2.SlsaProvenance
    slsa_provenance_zero_two: _slsa_provenance_zero_two_pb2.SlsaProvenanceZeroTwo
    def __init__(self, type: _Optional[str] = ..., subject: _Optional[_Iterable[_Union[Subject, _Mapping]]] = ..., predicate_type: _Optional[str] = ..., provenance: _Optional[_Union[_intoto_provenance_pb2.InTotoProvenance, _Mapping]] = ..., slsa_provenance: _Optional[_Union[_slsa_provenance_pb2.SlsaProvenance, _Mapping]] = ..., slsa_provenance_zero_two: _Optional[_Union[_slsa_provenance_zero_two_pb2.SlsaProvenanceZeroTwo, _Mapping]] = ...) -> None: ...

class Subject(_message.Message):
    __slots__ = ("name", "digest")
    class DigestEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    NAME_FIELD_NUMBER: _ClassVar[int]
    DIGEST_FIELD_NUMBER: _ClassVar[int]
    name: str
    digest: _containers.ScalarMap[str, str]
    def __init__(self, name: _Optional[str] = ..., digest: _Optional[_Mapping[str, str]] = ...) -> None: ...

class InTotoSlsaProvenanceV1(_message.Message):
    __slots__ = ("type", "subject", "predicate_type", "predicate")
    class SlsaProvenanceV1(_message.Message):
        __slots__ = ("build_definition", "run_details")
        BUILD_DEFINITION_FIELD_NUMBER: _ClassVar[int]
        RUN_DETAILS_FIELD_NUMBER: _ClassVar[int]
        build_definition: InTotoSlsaProvenanceV1.BuildDefinition
        run_details: InTotoSlsaProvenanceV1.RunDetails
        def __init__(self, build_definition: _Optional[_Union[InTotoSlsaProvenanceV1.BuildDefinition, _Mapping]] = ..., run_details: _Optional[_Union[InTotoSlsaProvenanceV1.RunDetails, _Mapping]] = ...) -> None: ...
    class BuildDefinition(_message.Message):
        __slots__ = ("build_type", "external_parameters", "internal_parameters", "resolved_dependencies")
        BUILD_TYPE_FIELD_NUMBER: _ClassVar[int]
        EXTERNAL_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
        INTERNAL_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
        RESOLVED_DEPENDENCIES_FIELD_NUMBER: _ClassVar[int]
        build_type: str
        external_parameters: _struct_pb2.Struct
        internal_parameters: _struct_pb2.Struct
        resolved_dependencies: _containers.RepeatedCompositeFieldContainer[InTotoSlsaProvenanceV1.ResourceDescriptor]
        def __init__(self, build_type: _Optional[str] = ..., external_parameters: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., internal_parameters: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., resolved_dependencies: _Optional[_Iterable[_Union[InTotoSlsaProvenanceV1.ResourceDescriptor, _Mapping]]] = ...) -> None: ...
    class ResourceDescriptor(_message.Message):
        __slots__ = ("name", "uri", "digest", "content", "download_location", "media_type", "annotations")
        class DigestEntry(_message.Message):
            __slots__ = ("key", "value")
            KEY_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            key: str
            value: str
            def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
        class AnnotationsEntry(_message.Message):
            __slots__ = ("key", "value")
            KEY_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            key: str
            value: _struct_pb2.Value
            def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[_struct_pb2.Value, _Mapping]] = ...) -> None: ...
        NAME_FIELD_NUMBER: _ClassVar[int]
        URI_FIELD_NUMBER: _ClassVar[int]
        DIGEST_FIELD_NUMBER: _ClassVar[int]
        CONTENT_FIELD_NUMBER: _ClassVar[int]
        DOWNLOAD_LOCATION_FIELD_NUMBER: _ClassVar[int]
        MEDIA_TYPE_FIELD_NUMBER: _ClassVar[int]
        ANNOTATIONS_FIELD_NUMBER: _ClassVar[int]
        name: str
        uri: str
        digest: _containers.ScalarMap[str, str]
        content: bytes
        download_location: str
        media_type: str
        annotations: _containers.MessageMap[str, _struct_pb2.Value]
        def __init__(self, name: _Optional[str] = ..., uri: _Optional[str] = ..., digest: _Optional[_Mapping[str, str]] = ..., content: _Optional[bytes] = ..., download_location: _Optional[str] = ..., media_type: _Optional[str] = ..., annotations: _Optional[_Mapping[str, _struct_pb2.Value]] = ...) -> None: ...
    class RunDetails(_message.Message):
        __slots__ = ("builder", "metadata", "byproducts")
        BUILDER_FIELD_NUMBER: _ClassVar[int]
        METADATA_FIELD_NUMBER: _ClassVar[int]
        BYPRODUCTS_FIELD_NUMBER: _ClassVar[int]
        builder: InTotoSlsaProvenanceV1.ProvenanceBuilder
        metadata: InTotoSlsaProvenanceV1.BuildMetadata
        byproducts: _containers.RepeatedCompositeFieldContainer[InTotoSlsaProvenanceV1.ResourceDescriptor]
        def __init__(self, builder: _Optional[_Union[InTotoSlsaProvenanceV1.ProvenanceBuilder, _Mapping]] = ..., metadata: _Optional[_Union[InTotoSlsaProvenanceV1.BuildMetadata, _Mapping]] = ..., byproducts: _Optional[_Iterable[_Union[InTotoSlsaProvenanceV1.ResourceDescriptor, _Mapping]]] = ...) -> None: ...
    class ProvenanceBuilder(_message.Message):
        __slots__ = ("id", "version", "builder_dependencies")
        class VersionEntry(_message.Message):
            __slots__ = ("key", "value")
            KEY_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            key: str
            value: str
            def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
        ID_FIELD_NUMBER: _ClassVar[int]
        VERSION_FIELD_NUMBER: _ClassVar[int]
        BUILDER_DEPENDENCIES_FIELD_NUMBER: _ClassVar[int]
        id: str
        version: _containers.ScalarMap[str, str]
        builder_dependencies: _containers.RepeatedCompositeFieldContainer[InTotoSlsaProvenanceV1.ResourceDescriptor]
        def __init__(self, id: _Optional[str] = ..., version: _Optional[_Mapping[str, str]] = ..., builder_dependencies: _Optional[_Iterable[_Union[InTotoSlsaProvenanceV1.ResourceDescriptor, _Mapping]]] = ...) -> None: ...
    class BuildMetadata(_message.Message):
        __slots__ = ("invocation_id", "started_on", "finished_on")
        INVOCATION_ID_FIELD_NUMBER: _ClassVar[int]
        STARTED_ON_FIELD_NUMBER: _ClassVar[int]
        FINISHED_ON_FIELD_NUMBER: _ClassVar[int]
        invocation_id: str
        started_on: _timestamp_pb2.Timestamp
        finished_on: _timestamp_pb2.Timestamp
        def __init__(self, invocation_id: _Optional[str] = ..., started_on: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., finished_on: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
    TYPE_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_FIELD_NUMBER: _ClassVar[int]
    PREDICATE_TYPE_FIELD_NUMBER: _ClassVar[int]
    PREDICATE_FIELD_NUMBER: _ClassVar[int]
    type: str
    subject: _containers.RepeatedCompositeFieldContainer[Subject]
    predicate_type: str
    predicate: InTotoSlsaProvenanceV1.SlsaProvenanceV1
    def __init__(self, type: _Optional[str] = ..., subject: _Optional[_Iterable[_Union[Subject, _Mapping]]] = ..., predicate_type: _Optional[str] = ..., predicate: _Optional[_Union[InTotoSlsaProvenanceV1.SlsaProvenanceV1, _Mapping]] = ...) -> None: ...
