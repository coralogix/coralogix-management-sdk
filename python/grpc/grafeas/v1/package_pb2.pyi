from google.api import field_behavior_pb2 as _field_behavior_pb2
from grafeas.v1 import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Architecture(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ARCHITECTURE_UNSPECIFIED: _ClassVar[Architecture]
    X86: _ClassVar[Architecture]
    X64: _ClassVar[Architecture]
ARCHITECTURE_UNSPECIFIED: Architecture
X86: Architecture
X64: Architecture

class Distribution(_message.Message):
    __slots__ = ("cpe_uri", "architecture", "latest_version", "maintainer", "url", "description")
    CPE_URI_FIELD_NUMBER: _ClassVar[int]
    ARCHITECTURE_FIELD_NUMBER: _ClassVar[int]
    LATEST_VERSION_FIELD_NUMBER: _ClassVar[int]
    MAINTAINER_FIELD_NUMBER: _ClassVar[int]
    URL_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    cpe_uri: str
    architecture: Architecture
    latest_version: Version
    maintainer: str
    url: str
    description: str
    def __init__(self, cpe_uri: _Optional[str] = ..., architecture: _Optional[_Union[Architecture, str]] = ..., latest_version: _Optional[_Union[Version, _Mapping]] = ..., maintainer: _Optional[str] = ..., url: _Optional[str] = ..., description: _Optional[str] = ...) -> None: ...

class Location(_message.Message):
    __slots__ = ("cpe_uri", "version", "path")
    CPE_URI_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    PATH_FIELD_NUMBER: _ClassVar[int]
    cpe_uri: str
    version: Version
    path: str
    def __init__(self, cpe_uri: _Optional[str] = ..., version: _Optional[_Union[Version, _Mapping]] = ..., path: _Optional[str] = ...) -> None: ...

class PackageNote(_message.Message):
    __slots__ = ("name", "distribution", "package_type", "cpe_uri", "architecture", "version", "maintainer", "url", "description", "license", "digest")
    NAME_FIELD_NUMBER: _ClassVar[int]
    DISTRIBUTION_FIELD_NUMBER: _ClassVar[int]
    PACKAGE_TYPE_FIELD_NUMBER: _ClassVar[int]
    CPE_URI_FIELD_NUMBER: _ClassVar[int]
    ARCHITECTURE_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    MAINTAINER_FIELD_NUMBER: _ClassVar[int]
    URL_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    LICENSE_FIELD_NUMBER: _ClassVar[int]
    DIGEST_FIELD_NUMBER: _ClassVar[int]
    name: str
    distribution: _containers.RepeatedCompositeFieldContainer[Distribution]
    package_type: str
    cpe_uri: str
    architecture: Architecture
    version: Version
    maintainer: str
    url: str
    description: str
    license: _common_pb2.License
    digest: _containers.RepeatedCompositeFieldContainer[_common_pb2.Digest]
    def __init__(self, name: _Optional[str] = ..., distribution: _Optional[_Iterable[_Union[Distribution, _Mapping]]] = ..., package_type: _Optional[str] = ..., cpe_uri: _Optional[str] = ..., architecture: _Optional[_Union[Architecture, str]] = ..., version: _Optional[_Union[Version, _Mapping]] = ..., maintainer: _Optional[str] = ..., url: _Optional[str] = ..., description: _Optional[str] = ..., license: _Optional[_Union[_common_pb2.License, _Mapping]] = ..., digest: _Optional[_Iterable[_Union[_common_pb2.Digest, _Mapping]]] = ...) -> None: ...

class PackageOccurrence(_message.Message):
    __slots__ = ("name", "location", "package_type", "cpe_uri", "architecture", "license", "version")
    NAME_FIELD_NUMBER: _ClassVar[int]
    LOCATION_FIELD_NUMBER: _ClassVar[int]
    PACKAGE_TYPE_FIELD_NUMBER: _ClassVar[int]
    CPE_URI_FIELD_NUMBER: _ClassVar[int]
    ARCHITECTURE_FIELD_NUMBER: _ClassVar[int]
    LICENSE_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    name: str
    location: _containers.RepeatedCompositeFieldContainer[Location]
    package_type: str
    cpe_uri: str
    architecture: Architecture
    license: _common_pb2.License
    version: Version
    def __init__(self, name: _Optional[str] = ..., location: _Optional[_Iterable[_Union[Location, _Mapping]]] = ..., package_type: _Optional[str] = ..., cpe_uri: _Optional[str] = ..., architecture: _Optional[_Union[Architecture, str]] = ..., license: _Optional[_Union[_common_pb2.License, _Mapping]] = ..., version: _Optional[_Union[Version, _Mapping]] = ...) -> None: ...

class Version(_message.Message):
    __slots__ = ("epoch", "name", "revision", "inclusive", "kind", "full_name")
    class VersionKind(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        VERSION_KIND_UNSPECIFIED: _ClassVar[Version.VersionKind]
        NORMAL: _ClassVar[Version.VersionKind]
        MINIMUM: _ClassVar[Version.VersionKind]
        MAXIMUM: _ClassVar[Version.VersionKind]
    VERSION_KIND_UNSPECIFIED: Version.VersionKind
    NORMAL: Version.VersionKind
    MINIMUM: Version.VersionKind
    MAXIMUM: Version.VersionKind
    EPOCH_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    REVISION_FIELD_NUMBER: _ClassVar[int]
    INCLUSIVE_FIELD_NUMBER: _ClassVar[int]
    KIND_FIELD_NUMBER: _ClassVar[int]
    FULL_NAME_FIELD_NUMBER: _ClassVar[int]
    epoch: int
    name: str
    revision: str
    inclusive: bool
    kind: Version.VersionKind
    full_name: str
    def __init__(self, epoch: _Optional[int] = ..., name: _Optional[str] = ..., revision: _Optional[str] = ..., inclusive: bool = ..., kind: _Optional[_Union[Version.VersionKind, str]] = ..., full_name: _Optional[str] = ...) -> None: ...
