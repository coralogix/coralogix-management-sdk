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

class GetDeployedExtensionsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetDeployedExtensionsResponse(_message.Message):
    __slots__ = ("deployed_extensions",)
    class DeployedExtensionSummary(_message.Message):
        __slots__ = ("deployed_item_counts",)
        DEPLOYED_ITEM_COUNTS_FIELD_NUMBER: _ClassVar[int]
        deployed_item_counts: _extension_pb2.ItemCounts
        def __init__(self, deployed_item_counts: _Optional[_Union[_extension_pb2.ItemCounts, _Mapping]] = ...) -> None: ...
    class DeployedExtension(_message.Message):
        __slots__ = ("id", "version", "applications", "subsystems", "item_ids", "summary")
        ID_FIELD_NUMBER: _ClassVar[int]
        VERSION_FIELD_NUMBER: _ClassVar[int]
        APPLICATIONS_FIELD_NUMBER: _ClassVar[int]
        SUBSYSTEMS_FIELD_NUMBER: _ClassVar[int]
        ITEM_IDS_FIELD_NUMBER: _ClassVar[int]
        SUMMARY_FIELD_NUMBER: _ClassVar[int]
        id: _wrappers_pb2.StringValue
        version: _wrappers_pb2.StringValue
        applications: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        subsystems: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        item_ids: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        summary: GetDeployedExtensionsResponse.DeployedExtensionSummary
        def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., version: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., applications: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystems: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., item_ids: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., summary: _Optional[_Union[GetDeployedExtensionsResponse.DeployedExtensionSummary, _Mapping]] = ...) -> None: ...
    DEPLOYED_EXTENSIONS_FIELD_NUMBER: _ClassVar[int]
    deployed_extensions: _containers.RepeatedCompositeFieldContainer[GetDeployedExtensionsResponse.DeployedExtension]
    def __init__(self, deployed_extensions: _Optional[_Iterable[_Union[GetDeployedExtensionsResponse.DeployedExtension, _Mapping]]] = ...) -> None: ...

class DeployExtensionRequest(_message.Message):
    __slots__ = ("id", "version", "item_ids", "applications", "subsystems")
    ID_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    ITEM_IDS_FIELD_NUMBER: _ClassVar[int]
    APPLICATIONS_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEMS_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    version: _wrappers_pb2.StringValue
    item_ids: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    applications: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    subsystems: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., version: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., item_ids: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., applications: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystems: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class DeployExtensionResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class UpdateExtensionRequest(_message.Message):
    __slots__ = ("id", "version", "item_ids", "applications", "subsystems")
    ID_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    ITEM_IDS_FIELD_NUMBER: _ClassVar[int]
    APPLICATIONS_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEMS_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    version: _wrappers_pb2.StringValue
    item_ids: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    applications: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    subsystems: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., version: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., item_ids: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., applications: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystems: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class UpdateExtensionResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class UndeployExtensionRequest(_message.Message):
    __slots__ = ("id", "kept_extension_items")
    ID_FIELD_NUMBER: _ClassVar[int]
    KEPT_EXTENSION_ITEMS_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    kept_extension_items: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., kept_extension_items: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class UndeployExtensionResponse(_message.Message):
    __slots__ = ("failed_items",)
    class FailedItem(_message.Message):
        __slots__ = ("item_id", "remote_id", "reason")
        ITEM_ID_FIELD_NUMBER: _ClassVar[int]
        REMOTE_ID_FIELD_NUMBER: _ClassVar[int]
        REASON_FIELD_NUMBER: _ClassVar[int]
        item_id: _wrappers_pb2.StringValue
        remote_id: _wrappers_pb2.StringValue
        reason: _wrappers_pb2.StringValue
        def __init__(self, item_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., remote_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., reason: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    FAILED_ITEMS_FIELD_NUMBER: _ClassVar[int]
    failed_items: _containers.RepeatedCompositeFieldContainer[UndeployExtensionResponse.FailedItem]
    def __init__(self, failed_items: _Optional[_Iterable[_Union[UndeployExtensionResponse.FailedItem, _Mapping]]] = ...) -> None: ...
