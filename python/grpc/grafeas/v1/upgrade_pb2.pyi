from google.protobuf import timestamp_pb2 as _timestamp_pb2
from grafeas.v1 import package_pb2 as _package_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UpgradeNote(_message.Message):
    __slots__ = ("package", "version", "distributions", "windows_update")
    PACKAGE_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    DISTRIBUTIONS_FIELD_NUMBER: _ClassVar[int]
    WINDOWS_UPDATE_FIELD_NUMBER: _ClassVar[int]
    package: str
    version: _package_pb2.Version
    distributions: _containers.RepeatedCompositeFieldContainer[UpgradeDistribution]
    windows_update: WindowsUpdate
    def __init__(self, package: _Optional[str] = ..., version: _Optional[_Union[_package_pb2.Version, _Mapping]] = ..., distributions: _Optional[_Iterable[_Union[UpgradeDistribution, _Mapping]]] = ..., windows_update: _Optional[_Union[WindowsUpdate, _Mapping]] = ...) -> None: ...

class UpgradeDistribution(_message.Message):
    __slots__ = ("cpe_uri", "classification", "severity", "cve")
    CPE_URI_FIELD_NUMBER: _ClassVar[int]
    CLASSIFICATION_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_FIELD_NUMBER: _ClassVar[int]
    CVE_FIELD_NUMBER: _ClassVar[int]
    cpe_uri: str
    classification: str
    severity: str
    cve: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, cpe_uri: _Optional[str] = ..., classification: _Optional[str] = ..., severity: _Optional[str] = ..., cve: _Optional[_Iterable[str]] = ...) -> None: ...

class WindowsUpdate(_message.Message):
    __slots__ = ("identity", "title", "description", "categories", "kb_article_ids", "support_url", "last_published_timestamp")
    class Identity(_message.Message):
        __slots__ = ("update_id", "revision")
        UPDATE_ID_FIELD_NUMBER: _ClassVar[int]
        REVISION_FIELD_NUMBER: _ClassVar[int]
        update_id: str
        revision: int
        def __init__(self, update_id: _Optional[str] = ..., revision: _Optional[int] = ...) -> None: ...
    class Category(_message.Message):
        __slots__ = ("category_id", "name")
        CATEGORY_ID_FIELD_NUMBER: _ClassVar[int]
        NAME_FIELD_NUMBER: _ClassVar[int]
        category_id: str
        name: str
        def __init__(self, category_id: _Optional[str] = ..., name: _Optional[str] = ...) -> None: ...
    IDENTITY_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    CATEGORIES_FIELD_NUMBER: _ClassVar[int]
    KB_ARTICLE_IDS_FIELD_NUMBER: _ClassVar[int]
    SUPPORT_URL_FIELD_NUMBER: _ClassVar[int]
    LAST_PUBLISHED_TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    identity: WindowsUpdate.Identity
    title: str
    description: str
    categories: _containers.RepeatedCompositeFieldContainer[WindowsUpdate.Category]
    kb_article_ids: _containers.RepeatedScalarFieldContainer[str]
    support_url: str
    last_published_timestamp: _timestamp_pb2.Timestamp
    def __init__(self, identity: _Optional[_Union[WindowsUpdate.Identity, _Mapping]] = ..., title: _Optional[str] = ..., description: _Optional[str] = ..., categories: _Optional[_Iterable[_Union[WindowsUpdate.Category, _Mapping]]] = ..., kb_article_ids: _Optional[_Iterable[str]] = ..., support_url: _Optional[str] = ..., last_published_timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class UpgradeOccurrence(_message.Message):
    __slots__ = ("package", "parsed_version", "distribution", "windows_update")
    PACKAGE_FIELD_NUMBER: _ClassVar[int]
    PARSED_VERSION_FIELD_NUMBER: _ClassVar[int]
    DISTRIBUTION_FIELD_NUMBER: _ClassVar[int]
    WINDOWS_UPDATE_FIELD_NUMBER: _ClassVar[int]
    package: str
    parsed_version: _package_pb2.Version
    distribution: UpgradeDistribution
    windows_update: WindowsUpdate
    def __init__(self, package: _Optional[str] = ..., parsed_version: _Optional[_Union[_package_pb2.Version, _Mapping]] = ..., distribution: _Optional[_Union[UpgradeDistribution, _Mapping]] = ..., windows_update: _Optional[_Union[WindowsUpdate, _Mapping]] = ...) -> None: ...
