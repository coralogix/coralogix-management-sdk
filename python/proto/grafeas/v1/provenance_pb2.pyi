from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class BuildProvenance(_message.Message):
    __slots__ = ("id", "project_id", "commands", "built_artifacts", "create_time", "start_time", "end_time", "creator", "logs_uri", "source_provenance", "trigger_id", "build_options", "builder_version")
    class BuildOptionsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_ID_FIELD_NUMBER: _ClassVar[int]
    COMMANDS_FIELD_NUMBER: _ClassVar[int]
    BUILT_ARTIFACTS_FIELD_NUMBER: _ClassVar[int]
    CREATE_TIME_FIELD_NUMBER: _ClassVar[int]
    START_TIME_FIELD_NUMBER: _ClassVar[int]
    END_TIME_FIELD_NUMBER: _ClassVar[int]
    CREATOR_FIELD_NUMBER: _ClassVar[int]
    LOGS_URI_FIELD_NUMBER: _ClassVar[int]
    SOURCE_PROVENANCE_FIELD_NUMBER: _ClassVar[int]
    TRIGGER_ID_FIELD_NUMBER: _ClassVar[int]
    BUILD_OPTIONS_FIELD_NUMBER: _ClassVar[int]
    BUILDER_VERSION_FIELD_NUMBER: _ClassVar[int]
    id: str
    project_id: str
    commands: _containers.RepeatedCompositeFieldContainer[Command]
    built_artifacts: _containers.RepeatedCompositeFieldContainer[Artifact]
    create_time: _timestamp_pb2.Timestamp
    start_time: _timestamp_pb2.Timestamp
    end_time: _timestamp_pb2.Timestamp
    creator: str
    logs_uri: str
    source_provenance: Source
    trigger_id: str
    build_options: _containers.ScalarMap[str, str]
    builder_version: str
    def __init__(self, id: _Optional[str] = ..., project_id: _Optional[str] = ..., commands: _Optional[_Iterable[_Union[Command, _Mapping]]] = ..., built_artifacts: _Optional[_Iterable[_Union[Artifact, _Mapping]]] = ..., create_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., start_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., end_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., creator: _Optional[str] = ..., logs_uri: _Optional[str] = ..., source_provenance: _Optional[_Union[Source, _Mapping]] = ..., trigger_id: _Optional[str] = ..., build_options: _Optional[_Mapping[str, str]] = ..., builder_version: _Optional[str] = ...) -> None: ...

class Source(_message.Message):
    __slots__ = ("artifact_storage_source_uri", "file_hashes", "context", "additional_contexts")
    class FileHashesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: FileHashes
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[FileHashes, _Mapping]] = ...) -> None: ...
    ARTIFACT_STORAGE_SOURCE_URI_FIELD_NUMBER: _ClassVar[int]
    FILE_HASHES_FIELD_NUMBER: _ClassVar[int]
    CONTEXT_FIELD_NUMBER: _ClassVar[int]
    ADDITIONAL_CONTEXTS_FIELD_NUMBER: _ClassVar[int]
    artifact_storage_source_uri: str
    file_hashes: _containers.MessageMap[str, FileHashes]
    context: SourceContext
    additional_contexts: _containers.RepeatedCompositeFieldContainer[SourceContext]
    def __init__(self, artifact_storage_source_uri: _Optional[str] = ..., file_hashes: _Optional[_Mapping[str, FileHashes]] = ..., context: _Optional[_Union[SourceContext, _Mapping]] = ..., additional_contexts: _Optional[_Iterable[_Union[SourceContext, _Mapping]]] = ...) -> None: ...

class FileHashes(_message.Message):
    __slots__ = ("file_hash",)
    FILE_HASH_FIELD_NUMBER: _ClassVar[int]
    file_hash: _containers.RepeatedCompositeFieldContainer[Hash]
    def __init__(self, file_hash: _Optional[_Iterable[_Union[Hash, _Mapping]]] = ...) -> None: ...

class Hash(_message.Message):
    __slots__ = ("type", "value")
    TYPE_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    type: str
    value: bytes
    def __init__(self, type: _Optional[str] = ..., value: _Optional[bytes] = ...) -> None: ...

class Command(_message.Message):
    __slots__ = ("name", "env", "args", "dir", "id", "wait_for")
    NAME_FIELD_NUMBER: _ClassVar[int]
    ENV_FIELD_NUMBER: _ClassVar[int]
    ARGS_FIELD_NUMBER: _ClassVar[int]
    DIR_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    WAIT_FOR_FIELD_NUMBER: _ClassVar[int]
    name: str
    env: _containers.RepeatedScalarFieldContainer[str]
    args: _containers.RepeatedScalarFieldContainer[str]
    dir: str
    id: str
    wait_for: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, name: _Optional[str] = ..., env: _Optional[_Iterable[str]] = ..., args: _Optional[_Iterable[str]] = ..., dir: _Optional[str] = ..., id: _Optional[str] = ..., wait_for: _Optional[_Iterable[str]] = ...) -> None: ...

class Artifact(_message.Message):
    __slots__ = ("checksum", "id", "names")
    CHECKSUM_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    NAMES_FIELD_NUMBER: _ClassVar[int]
    checksum: str
    id: str
    names: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, checksum: _Optional[str] = ..., id: _Optional[str] = ..., names: _Optional[_Iterable[str]] = ...) -> None: ...

class SourceContext(_message.Message):
    __slots__ = ("cloud_repo", "gerrit", "git", "labels")
    class LabelsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    CLOUD_REPO_FIELD_NUMBER: _ClassVar[int]
    GERRIT_FIELD_NUMBER: _ClassVar[int]
    GIT_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    cloud_repo: CloudRepoSourceContext
    gerrit: GerritSourceContext
    git: GitSourceContext
    labels: _containers.ScalarMap[str, str]
    def __init__(self, cloud_repo: _Optional[_Union[CloudRepoSourceContext, _Mapping]] = ..., gerrit: _Optional[_Union[GerritSourceContext, _Mapping]] = ..., git: _Optional[_Union[GitSourceContext, _Mapping]] = ..., labels: _Optional[_Mapping[str, str]] = ...) -> None: ...

class AliasContext(_message.Message):
    __slots__ = ("kind", "name")
    class Kind(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        KIND_UNSPECIFIED: _ClassVar[AliasContext.Kind]
        FIXED: _ClassVar[AliasContext.Kind]
        MOVABLE: _ClassVar[AliasContext.Kind]
        OTHER: _ClassVar[AliasContext.Kind]
    KIND_UNSPECIFIED: AliasContext.Kind
    FIXED: AliasContext.Kind
    MOVABLE: AliasContext.Kind
    OTHER: AliasContext.Kind
    KIND_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    kind: AliasContext.Kind
    name: str
    def __init__(self, kind: _Optional[_Union[AliasContext.Kind, str]] = ..., name: _Optional[str] = ...) -> None: ...

class CloudRepoSourceContext(_message.Message):
    __slots__ = ("repo_id", "revision_id", "alias_context")
    REPO_ID_FIELD_NUMBER: _ClassVar[int]
    REVISION_ID_FIELD_NUMBER: _ClassVar[int]
    ALIAS_CONTEXT_FIELD_NUMBER: _ClassVar[int]
    repo_id: RepoId
    revision_id: str
    alias_context: AliasContext
    def __init__(self, repo_id: _Optional[_Union[RepoId, _Mapping]] = ..., revision_id: _Optional[str] = ..., alias_context: _Optional[_Union[AliasContext, _Mapping]] = ...) -> None: ...

class GerritSourceContext(_message.Message):
    __slots__ = ("host_uri", "gerrit_project", "revision_id", "alias_context")
    HOST_URI_FIELD_NUMBER: _ClassVar[int]
    GERRIT_PROJECT_FIELD_NUMBER: _ClassVar[int]
    REVISION_ID_FIELD_NUMBER: _ClassVar[int]
    ALIAS_CONTEXT_FIELD_NUMBER: _ClassVar[int]
    host_uri: str
    gerrit_project: str
    revision_id: str
    alias_context: AliasContext
    def __init__(self, host_uri: _Optional[str] = ..., gerrit_project: _Optional[str] = ..., revision_id: _Optional[str] = ..., alias_context: _Optional[_Union[AliasContext, _Mapping]] = ...) -> None: ...

class GitSourceContext(_message.Message):
    __slots__ = ("url", "revision_id")
    URL_FIELD_NUMBER: _ClassVar[int]
    REVISION_ID_FIELD_NUMBER: _ClassVar[int]
    url: str
    revision_id: str
    def __init__(self, url: _Optional[str] = ..., revision_id: _Optional[str] = ...) -> None: ...

class RepoId(_message.Message):
    __slots__ = ("project_repo_id", "uid")
    PROJECT_REPO_ID_FIELD_NUMBER: _ClassVar[int]
    UID_FIELD_NUMBER: _ClassVar[int]
    project_repo_id: ProjectRepoId
    uid: str
    def __init__(self, project_repo_id: _Optional[_Union[ProjectRepoId, _Mapping]] = ..., uid: _Optional[str] = ...) -> None: ...

class ProjectRepoId(_message.Message):
    __slots__ = ("project_id", "repo_name")
    PROJECT_ID_FIELD_NUMBER: _ClassVar[int]
    REPO_NAME_FIELD_NUMBER: _ClassVar[int]
    project_id: str
    repo_name: str
    def __init__(self, project_id: _Optional[str] = ..., repo_name: _Optional[str] = ...) -> None: ...
