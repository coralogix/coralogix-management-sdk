from google.protobuf import any_pb2 as _any_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Recipe(_message.Message):
    __slots__ = ("type", "defined_in_material", "entry_point", "arguments", "environment")
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DEFINED_IN_MATERIAL_FIELD_NUMBER: _ClassVar[int]
    ENTRY_POINT_FIELD_NUMBER: _ClassVar[int]
    ARGUMENTS_FIELD_NUMBER: _ClassVar[int]
    ENVIRONMENT_FIELD_NUMBER: _ClassVar[int]
    type: str
    defined_in_material: int
    entry_point: str
    arguments: _containers.RepeatedCompositeFieldContainer[_any_pb2.Any]
    environment: _containers.RepeatedCompositeFieldContainer[_any_pb2.Any]
    def __init__(self, type: _Optional[str] = ..., defined_in_material: _Optional[int] = ..., entry_point: _Optional[str] = ..., arguments: _Optional[_Iterable[_Union[_any_pb2.Any, _Mapping]]] = ..., environment: _Optional[_Iterable[_Union[_any_pb2.Any, _Mapping]]] = ...) -> None: ...

class Completeness(_message.Message):
    __slots__ = ("arguments", "environment", "materials")
    ARGUMENTS_FIELD_NUMBER: _ClassVar[int]
    ENVIRONMENT_FIELD_NUMBER: _ClassVar[int]
    MATERIALS_FIELD_NUMBER: _ClassVar[int]
    arguments: bool
    environment: bool
    materials: bool
    def __init__(self, arguments: bool = ..., environment: bool = ..., materials: bool = ...) -> None: ...

class Metadata(_message.Message):
    __slots__ = ("build_invocation_id", "build_started_on", "build_finished_on", "completeness", "reproducible")
    BUILD_INVOCATION_ID_FIELD_NUMBER: _ClassVar[int]
    BUILD_STARTED_ON_FIELD_NUMBER: _ClassVar[int]
    BUILD_FINISHED_ON_FIELD_NUMBER: _ClassVar[int]
    COMPLETENESS_FIELD_NUMBER: _ClassVar[int]
    REPRODUCIBLE_FIELD_NUMBER: _ClassVar[int]
    build_invocation_id: str
    build_started_on: _timestamp_pb2.Timestamp
    build_finished_on: _timestamp_pb2.Timestamp
    completeness: Completeness
    reproducible: bool
    def __init__(self, build_invocation_id: _Optional[str] = ..., build_started_on: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., build_finished_on: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., completeness: _Optional[_Union[Completeness, _Mapping]] = ..., reproducible: bool = ...) -> None: ...

class BuilderConfig(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class InTotoProvenance(_message.Message):
    __slots__ = ("builder_config", "recipe", "metadata", "materials")
    BUILDER_CONFIG_FIELD_NUMBER: _ClassVar[int]
    RECIPE_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    MATERIALS_FIELD_NUMBER: _ClassVar[int]
    builder_config: BuilderConfig
    recipe: Recipe
    metadata: Metadata
    materials: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, builder_config: _Optional[_Union[BuilderConfig, _Mapping]] = ..., recipe: _Optional[_Union[Recipe, _Mapping]] = ..., metadata: _Optional[_Union[Metadata, _Mapping]] = ..., materials: _Optional[_Iterable[str]] = ...) -> None: ...
