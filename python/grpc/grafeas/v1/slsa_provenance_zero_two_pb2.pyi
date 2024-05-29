from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SlsaProvenanceZeroTwo(_message.Message):
    __slots__ = ("builder", "build_type", "invocation", "build_config", "metadata", "materials")
    class SlsaBuilder(_message.Message):
        __slots__ = ("id",)
        ID_FIELD_NUMBER: _ClassVar[int]
        id: str
        def __init__(self, id: _Optional[str] = ...) -> None: ...
    class SlsaMaterial(_message.Message):
        __slots__ = ("uri", "digest")
        class DigestEntry(_message.Message):
            __slots__ = ("key", "value")
            KEY_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            key: str
            value: str
            def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
        URI_FIELD_NUMBER: _ClassVar[int]
        DIGEST_FIELD_NUMBER: _ClassVar[int]
        uri: str
        digest: _containers.ScalarMap[str, str]
        def __init__(self, uri: _Optional[str] = ..., digest: _Optional[_Mapping[str, str]] = ...) -> None: ...
    class SlsaInvocation(_message.Message):
        __slots__ = ("config_source", "parameters", "environment")
        CONFIG_SOURCE_FIELD_NUMBER: _ClassVar[int]
        PARAMETERS_FIELD_NUMBER: _ClassVar[int]
        ENVIRONMENT_FIELD_NUMBER: _ClassVar[int]
        config_source: SlsaProvenanceZeroTwo.SlsaConfigSource
        parameters: _struct_pb2.Struct
        environment: _struct_pb2.Struct
        def __init__(self, config_source: _Optional[_Union[SlsaProvenanceZeroTwo.SlsaConfigSource, _Mapping]] = ..., parameters: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., environment: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...
    class SlsaConfigSource(_message.Message):
        __slots__ = ("uri", "digest", "entry_point")
        class DigestEntry(_message.Message):
            __slots__ = ("key", "value")
            KEY_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            key: str
            value: str
            def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
        URI_FIELD_NUMBER: _ClassVar[int]
        DIGEST_FIELD_NUMBER: _ClassVar[int]
        ENTRY_POINT_FIELD_NUMBER: _ClassVar[int]
        uri: str
        digest: _containers.ScalarMap[str, str]
        entry_point: str
        def __init__(self, uri: _Optional[str] = ..., digest: _Optional[_Mapping[str, str]] = ..., entry_point: _Optional[str] = ...) -> None: ...
    class SlsaMetadata(_message.Message):
        __slots__ = ("build_invocation_id", "build_started_on", "build_finished_on", "completeness", "reproducible")
        BUILD_INVOCATION_ID_FIELD_NUMBER: _ClassVar[int]
        BUILD_STARTED_ON_FIELD_NUMBER: _ClassVar[int]
        BUILD_FINISHED_ON_FIELD_NUMBER: _ClassVar[int]
        COMPLETENESS_FIELD_NUMBER: _ClassVar[int]
        REPRODUCIBLE_FIELD_NUMBER: _ClassVar[int]
        build_invocation_id: str
        build_started_on: _timestamp_pb2.Timestamp
        build_finished_on: _timestamp_pb2.Timestamp
        completeness: SlsaProvenanceZeroTwo.SlsaCompleteness
        reproducible: bool
        def __init__(self, build_invocation_id: _Optional[str] = ..., build_started_on: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., build_finished_on: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., completeness: _Optional[_Union[SlsaProvenanceZeroTwo.SlsaCompleteness, _Mapping]] = ..., reproducible: bool = ...) -> None: ...
    class SlsaCompleteness(_message.Message):
        __slots__ = ("parameters", "environment", "materials")
        PARAMETERS_FIELD_NUMBER: _ClassVar[int]
        ENVIRONMENT_FIELD_NUMBER: _ClassVar[int]
        MATERIALS_FIELD_NUMBER: _ClassVar[int]
        parameters: bool
        environment: bool
        materials: bool
        def __init__(self, parameters: bool = ..., environment: bool = ..., materials: bool = ...) -> None: ...
    BUILDER_FIELD_NUMBER: _ClassVar[int]
    BUILD_TYPE_FIELD_NUMBER: _ClassVar[int]
    INVOCATION_FIELD_NUMBER: _ClassVar[int]
    BUILD_CONFIG_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    MATERIALS_FIELD_NUMBER: _ClassVar[int]
    builder: SlsaProvenanceZeroTwo.SlsaBuilder
    build_type: str
    invocation: SlsaProvenanceZeroTwo.SlsaInvocation
    build_config: _struct_pb2.Struct
    metadata: SlsaProvenanceZeroTwo.SlsaMetadata
    materials: _containers.RepeatedCompositeFieldContainer[SlsaProvenanceZeroTwo.SlsaMaterial]
    def __init__(self, builder: _Optional[_Union[SlsaProvenanceZeroTwo.SlsaBuilder, _Mapping]] = ..., build_type: _Optional[str] = ..., invocation: _Optional[_Union[SlsaProvenanceZeroTwo.SlsaInvocation, _Mapping]] = ..., build_config: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., metadata: _Optional[_Union[SlsaProvenanceZeroTwo.SlsaMetadata, _Mapping]] = ..., materials: _Optional[_Iterable[_Union[SlsaProvenanceZeroTwo.SlsaMaterial, _Mapping]]] = ...) -> None: ...
