from com.coralogix.extensions.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.extensions.v1 import extension_pb2 as _extension_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetAllExtensionsRequest(_message.Message):
    __slots__ = ("include_hidden_extensions", "filter")
    class Filter(_message.Message):
        __slots__ = ("integrations",)
        INTEGRATIONS_FIELD_NUMBER: _ClassVar[int]
        integrations: _containers.RepeatedScalarFieldContainer[str]
        def __init__(self, integrations: _Optional[_Iterable[str]] = ...) -> None: ...
    INCLUDE_HIDDEN_EXTENSIONS_FIELD_NUMBER: _ClassVar[int]
    FILTER_FIELD_NUMBER: _ClassVar[int]
    include_hidden_extensions: _wrappers_pb2.BoolValue
    filter: GetAllExtensionsRequest.Filter
    def __init__(self, include_hidden_extensions: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., filter: _Optional[_Union[GetAllExtensionsRequest.Filter, _Mapping]] = ...) -> None: ...

class GetAllExtensionsResponse(_message.Message):
    __slots__ = ("extensions",)
    class RevisionSummary(_message.Message):
        __slots__ = ("item_counts", "is_new")
        ITEM_COUNTS_FIELD_NUMBER: _ClassVar[int]
        IS_NEW_FIELD_NUMBER: _ClassVar[int]
        item_counts: _extension_pb2.ItemCounts
        is_new: _wrappers_pb2.BoolValue
        def __init__(self, item_counts: _Optional[_Union[_extension_pb2.ItemCounts, _Mapping]] = ..., is_new: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
    class Revision(_message.Message):
        __slots__ = ("version", "description", "excerpt", "labels", "integration_details", "summary")
        VERSION_FIELD_NUMBER: _ClassVar[int]
        DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
        EXCERPT_FIELD_NUMBER: _ClassVar[int]
        LABELS_FIELD_NUMBER: _ClassVar[int]
        INTEGRATION_DETAILS_FIELD_NUMBER: _ClassVar[int]
        SUMMARY_FIELD_NUMBER: _ClassVar[int]
        version: _wrappers_pb2.StringValue
        description: _wrappers_pb2.StringValue
        excerpt: _wrappers_pb2.StringValue
        labels: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        integration_details: _containers.RepeatedCompositeFieldContainer[_extension_pb2.IntegrationDetail]
        summary: GetAllExtensionsResponse.RevisionSummary
        def __init__(self, version: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., excerpt: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., labels: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., integration_details: _Optional[_Iterable[_Union[_extension_pb2.IntegrationDetail, _Mapping]]] = ..., summary: _Optional[_Union[GetAllExtensionsResponse.RevisionSummary, _Mapping]] = ...) -> None: ...
    class Extension(_message.Message):
        __slots__ = ("id", "name", "image", "dark_mode_image", "revisions", "is_hidden", "integrations", "keywords", "deprecation")
        ID_FIELD_NUMBER: _ClassVar[int]
        NAME_FIELD_NUMBER: _ClassVar[int]
        IMAGE_FIELD_NUMBER: _ClassVar[int]
        DARK_MODE_IMAGE_FIELD_NUMBER: _ClassVar[int]
        REVISIONS_FIELD_NUMBER: _ClassVar[int]
        IS_HIDDEN_FIELD_NUMBER: _ClassVar[int]
        INTEGRATIONS_FIELD_NUMBER: _ClassVar[int]
        KEYWORDS_FIELD_NUMBER: _ClassVar[int]
        DEPRECATION_FIELD_NUMBER: _ClassVar[int]
        id: _wrappers_pb2.StringValue
        name: _wrappers_pb2.StringValue
        image: _wrappers_pb2.StringValue
        dark_mode_image: _wrappers_pb2.StringValue
        revisions: _containers.RepeatedCompositeFieldContainer[GetAllExtensionsResponse.Revision]
        is_hidden: _wrappers_pb2.BoolValue
        integrations: _containers.RepeatedScalarFieldContainer[str]
        keywords: _containers.RepeatedScalarFieldContainer[str]
        deprecation: _extension_pb2.Deprecation
        def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., image: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., dark_mode_image: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., revisions: _Optional[_Iterable[_Union[GetAllExtensionsResponse.Revision, _Mapping]]] = ..., is_hidden: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., integrations: _Optional[_Iterable[str]] = ..., keywords: _Optional[_Iterable[str]] = ..., deprecation: _Optional[_Union[_extension_pb2.Deprecation, _Mapping]] = ...) -> None: ...
    EXTENSIONS_FIELD_NUMBER: _ClassVar[int]
    extensions: _containers.RepeatedCompositeFieldContainer[GetAllExtensionsResponse.Extension]
    def __init__(self, extensions: _Optional[_Iterable[_Union[GetAllExtensionsResponse.Extension, _Mapping]]] = ...) -> None: ...

class GetExtensionRequest(_message.Message):
    __slots__ = ("id", "include_dashboard_binaries", "include_testing_revision")
    ID_FIELD_NUMBER: _ClassVar[int]
    INCLUDE_DASHBOARD_BINARIES_FIELD_NUMBER: _ClassVar[int]
    INCLUDE_TESTING_REVISION_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    include_dashboard_binaries: _wrappers_pb2.BoolValue
    include_testing_revision: _wrappers_pb2.BoolValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., include_dashboard_binaries: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., include_testing_revision: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class GetExtensionResponse(_message.Message):
    __slots__ = ("extension",)
    EXTENSION_FIELD_NUMBER: _ClassVar[int]
    extension: _extension_pb2.Extension
    def __init__(self, extension: _Optional[_Union[_extension_pb2.Extension, _Mapping]] = ...) -> None: ...
