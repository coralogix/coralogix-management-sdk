from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class TargetDomain(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ACTION: _ClassVar[TargetDomain]
    ALERT: _ClassVar[TargetDomain]
    ENRICHMENT: _ClassVar[TargetDomain]
    GRAFANA_DASHBOARD: _ClassVar[TargetDomain]
    KIBANA_DASHBOARD: _ClassVar[TargetDomain]
    PARSING_RULE: _ClassVar[TargetDomain]
    SAVED_VIEW: _ClassVar[TargetDomain]
    CX_CUSTOM_DASHBOARD: _ClassVar[TargetDomain]
    METRICS_RULE_GROUP: _ClassVar[TargetDomain]
    EVENTS_TO_METRICS: _ClassVar[TargetDomain]
ACTION: TargetDomain
ALERT: TargetDomain
ENRICHMENT: TargetDomain
GRAFANA_DASHBOARD: TargetDomain
KIBANA_DASHBOARD: TargetDomain
PARSING_RULE: TargetDomain
SAVED_VIEW: TargetDomain
CX_CUSTOM_DASHBOARD: TargetDomain
METRICS_RULE_GROUP: TargetDomain
EVENTS_TO_METRICS: TargetDomain

class ExtensionItemBinary(_message.Message):
    __slots__ = ("type", "data", "file_name")
    class BinaryType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        PREVIEW_IMAGE: _ClassVar[ExtensionItemBinary.BinaryType]
        KIBANA_DASHBOARD_DEFINITION: _ClassVar[ExtensionItemBinary.BinaryType]
        GRAFANA_DASHBOARD_DEFINITION: _ClassVar[ExtensionItemBinary.BinaryType]
        ENRICHMENT_CSV: _ClassVar[ExtensionItemBinary.BinaryType]
        CX_CUSTOM_DASHBOARD_DEFINITION: _ClassVar[ExtensionItemBinary.BinaryType]
    PREVIEW_IMAGE: ExtensionItemBinary.BinaryType
    KIBANA_DASHBOARD_DEFINITION: ExtensionItemBinary.BinaryType
    GRAFANA_DASHBOARD_DEFINITION: ExtensionItemBinary.BinaryType
    ENRICHMENT_CSV: ExtensionItemBinary.BinaryType
    CX_CUSTOM_DASHBOARD_DEFINITION: ExtensionItemBinary.BinaryType
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    FILE_NAME_FIELD_NUMBER: _ClassVar[int]
    type: ExtensionItemBinary.BinaryType
    data: _wrappers_pb2.StringValue
    file_name: _wrappers_pb2.StringValue
    def __init__(self, type: _Optional[_Union[ExtensionItemBinary.BinaryType, str]] = ..., data: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., file_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class ExtensionItem(_message.Message):
    __slots__ = ("id", "name", "description", "target_domain", "data", "binaries", "is_mandatory", "permission_resource", "extended_internal_id", "unique_id")
    class PermissionResource(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        UNKNOWN: _ClassVar[ExtensionItem.PermissionResource]
        ACTION: _ClassVar[ExtensionItem.PermissionResource]
        ALERT: _ClassVar[ExtensionItem.PermissionResource]
        CUSTOM_ENRICHMENT: _ClassVar[ExtensionItem.PermissionResource]
        GEO_ENRICHMENT: _ClassVar[ExtensionItem.PermissionResource]
        SECURITY_ENRICHMENT: _ClassVar[ExtensionItem.PermissionResource]
        RESOURCE_CLOUD_METADATA_ENRICHMENT: _ClassVar[ExtensionItem.PermissionResource]
        GRAFANA_DASHBOARD: _ClassVar[ExtensionItem.PermissionResource]
        KIBANA_DASHBOARD: _ClassVar[ExtensionItem.PermissionResource]
        PARSING_RULE: _ClassVar[ExtensionItem.PermissionResource]
        SAVED_VIEW: _ClassVar[ExtensionItem.PermissionResource]
        CX_CUSTOM_DASHBOARD: _ClassVar[ExtensionItem.PermissionResource]
        METRICS_RULE_GROUP: _ClassVar[ExtensionItem.PermissionResource]
        SPAN_EVENTS_TO_METRICS: _ClassVar[ExtensionItem.PermissionResource]
        LOGS_EVENTS_TO_METRICS: _ClassVar[ExtensionItem.PermissionResource]
    UNKNOWN: ExtensionItem.PermissionResource
    ACTION: ExtensionItem.PermissionResource
    ALERT: ExtensionItem.PermissionResource
    CUSTOM_ENRICHMENT: ExtensionItem.PermissionResource
    GEO_ENRICHMENT: ExtensionItem.PermissionResource
    SECURITY_ENRICHMENT: ExtensionItem.PermissionResource
    RESOURCE_CLOUD_METADATA_ENRICHMENT: ExtensionItem.PermissionResource
    GRAFANA_DASHBOARD: ExtensionItem.PermissionResource
    KIBANA_DASHBOARD: ExtensionItem.PermissionResource
    PARSING_RULE: ExtensionItem.PermissionResource
    SAVED_VIEW: ExtensionItem.PermissionResource
    CX_CUSTOM_DASHBOARD: ExtensionItem.PermissionResource
    METRICS_RULE_GROUP: ExtensionItem.PermissionResource
    SPAN_EVENTS_TO_METRICS: ExtensionItem.PermissionResource
    LOGS_EVENTS_TO_METRICS: ExtensionItem.PermissionResource
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    TARGET_DOMAIN_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    BINARIES_FIELD_NUMBER: _ClassVar[int]
    IS_MANDATORY_FIELD_NUMBER: _ClassVar[int]
    PERMISSION_RESOURCE_FIELD_NUMBER: _ClassVar[int]
    EXTENDED_INTERNAL_ID_FIELD_NUMBER: _ClassVar[int]
    UNIQUE_ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    target_domain: TargetDomain
    data: _struct_pb2.Struct
    binaries: _containers.RepeatedCompositeFieldContainer[ExtensionItemBinary]
    is_mandatory: _wrappers_pb2.BoolValue
    permission_resource: ExtensionItem.PermissionResource
    extended_internal_id: _wrappers_pb2.StringValue
    unique_id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., target_domain: _Optional[_Union[TargetDomain, str]] = ..., data: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., binaries: _Optional[_Iterable[_Union[ExtensionItemBinary, _Mapping]]] = ..., is_mandatory: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., permission_resource: _Optional[_Union[ExtensionItem.PermissionResource, str]] = ..., extended_internal_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., unique_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class ItemCounts(_message.Message):
    __slots__ = ("actions", "alerts", "custom_dashboards", "enrichments", "events_to_metrics", "grafana_dashboards", "kibana_dashboards", "metrics_rule_group", "parsing_rules", "saved_views")
    ACTIONS_FIELD_NUMBER: _ClassVar[int]
    ALERTS_FIELD_NUMBER: _ClassVar[int]
    CUSTOM_DASHBOARDS_FIELD_NUMBER: _ClassVar[int]
    ENRICHMENTS_FIELD_NUMBER: _ClassVar[int]
    EVENTS_TO_METRICS_FIELD_NUMBER: _ClassVar[int]
    GRAFANA_DASHBOARDS_FIELD_NUMBER: _ClassVar[int]
    KIBANA_DASHBOARDS_FIELD_NUMBER: _ClassVar[int]
    METRICS_RULE_GROUP_FIELD_NUMBER: _ClassVar[int]
    PARSING_RULES_FIELD_NUMBER: _ClassVar[int]
    SAVED_VIEWS_FIELD_NUMBER: _ClassVar[int]
    actions: _wrappers_pb2.UInt32Value
    alerts: _wrappers_pb2.UInt32Value
    custom_dashboards: _wrappers_pb2.UInt32Value
    enrichments: _wrappers_pb2.UInt32Value
    events_to_metrics: _wrappers_pb2.UInt32Value
    grafana_dashboards: _wrappers_pb2.UInt32Value
    kibana_dashboards: _wrappers_pb2.UInt32Value
    metrics_rule_group: _wrappers_pb2.UInt32Value
    parsing_rules: _wrappers_pb2.UInt32Value
    saved_views: _wrappers_pb2.UInt32Value
    def __init__(self, actions: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., alerts: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., custom_dashboards: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., enrichments: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., events_to_metrics: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., grafana_dashboards: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., kibana_dashboards: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., metrics_rule_group: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., parsing_rules: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., saved_views: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...

class IntegrationDetail(_message.Message):
    __slots__ = ("name", "link")
    NAME_FIELD_NUMBER: _ClassVar[int]
    LINK_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    link: _wrappers_pb2.StringValue
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., link: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class ExtensionBinary(_message.Message):
    __slots__ = ("type", "data")
    class BinaryType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        KIBANA_INDEX_PATTERN: _ClassVar[ExtensionBinary.BinaryType]
    KIBANA_INDEX_PATTERN: ExtensionBinary.BinaryType
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    type: ExtensionBinary.BinaryType
    data: _wrappers_pb2.StringValue
    def __init__(self, type: _Optional[_Union[ExtensionBinary.BinaryType, str]] = ..., data: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class ChangelogEntry(_message.Message):
    __slots__ = ("version", "description_md")
    VERSION_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_MD_FIELD_NUMBER: _ClassVar[int]
    version: _wrappers_pb2.StringValue
    description_md: _wrappers_pb2.StringValue
    def __init__(self, version: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description_md: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class Deprecation(_message.Message):
    __slots__ = ("reason", "replacement_extensions")
    REASON_FIELD_NUMBER: _ClassVar[int]
    REPLACEMENT_EXTENSIONS_FIELD_NUMBER: _ClassVar[int]
    reason: _wrappers_pb2.StringValue
    replacement_extensions: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, reason: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., replacement_extensions: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class ExtensionRevision(_message.Message):
    __slots__ = ("version", "description", "excerpt", "labels", "integration_details", "items", "binaries", "permission_denied_items", "is_testing")
    VERSION_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    EXCERPT_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    INTEGRATION_DETAILS_FIELD_NUMBER: _ClassVar[int]
    ITEMS_FIELD_NUMBER: _ClassVar[int]
    BINARIES_FIELD_NUMBER: _ClassVar[int]
    PERMISSION_DENIED_ITEMS_FIELD_NUMBER: _ClassVar[int]
    IS_TESTING_FIELD_NUMBER: _ClassVar[int]
    version: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    excerpt: _wrappers_pb2.StringValue
    labels: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    integration_details: _containers.RepeatedCompositeFieldContainer[IntegrationDetail]
    items: _containers.RepeatedCompositeFieldContainer[ExtensionItem]
    binaries: _containers.RepeatedCompositeFieldContainer[ExtensionBinary]
    permission_denied_items: _containers.RepeatedCompositeFieldContainer[ExtensionItem]
    is_testing: _wrappers_pb2.BoolValue
    def __init__(self, version: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., excerpt: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., labels: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., integration_details: _Optional[_Iterable[_Union[IntegrationDetail, _Mapping]]] = ..., items: _Optional[_Iterable[_Union[ExtensionItem, _Mapping]]] = ..., binaries: _Optional[_Iterable[_Union[ExtensionBinary, _Mapping]]] = ..., permission_denied_items: _Optional[_Iterable[_Union[ExtensionItem, _Mapping]]] = ..., is_testing: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class Extension(_message.Message):
    __slots__ = ("id", "name", "image", "dark_mode_image", "revisions", "is_hidden", "integrations", "keywords", "permission_denied_revisions", "changelog", "deprecation")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    DARK_MODE_IMAGE_FIELD_NUMBER: _ClassVar[int]
    REVISIONS_FIELD_NUMBER: _ClassVar[int]
    IS_HIDDEN_FIELD_NUMBER: _ClassVar[int]
    INTEGRATIONS_FIELD_NUMBER: _ClassVar[int]
    KEYWORDS_FIELD_NUMBER: _ClassVar[int]
    PERMISSION_DENIED_REVISIONS_FIELD_NUMBER: _ClassVar[int]
    CHANGELOG_FIELD_NUMBER: _ClassVar[int]
    DEPRECATION_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    image: _wrappers_pb2.StringValue
    dark_mode_image: _wrappers_pb2.StringValue
    revisions: _containers.RepeatedCompositeFieldContainer[ExtensionRevision]
    is_hidden: _wrappers_pb2.BoolValue
    integrations: _containers.RepeatedScalarFieldContainer[str]
    keywords: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    permission_denied_revisions: _containers.RepeatedCompositeFieldContainer[ExtensionRevision]
    changelog: _containers.RepeatedCompositeFieldContainer[ChangelogEntry]
    deprecation: Deprecation
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., image: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., dark_mode_image: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., revisions: _Optional[_Iterable[_Union[ExtensionRevision, _Mapping]]] = ..., is_hidden: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., integrations: _Optional[_Iterable[str]] = ..., keywords: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., permission_denied_revisions: _Optional[_Iterable[_Union[ExtensionRevision, _Mapping]]] = ..., changelog: _Optional[_Iterable[_Union[ChangelogEntry, _Mapping]]] = ..., deprecation: _Optional[_Union[Deprecation, _Mapping]] = ...) -> None: ...

class ExtensionData(_message.Message):
    __slots__ = ("id", "name", "description", "excerpt", "image", "dark_mode_image", "labels", "version", "items", "integration_details", "is_hidden", "binaries", "integrations", "keywords", "changelog", "deprecation")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    EXCERPT_FIELD_NUMBER: _ClassVar[int]
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    DARK_MODE_IMAGE_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    ITEMS_FIELD_NUMBER: _ClassVar[int]
    INTEGRATION_DETAILS_FIELD_NUMBER: _ClassVar[int]
    IS_HIDDEN_FIELD_NUMBER: _ClassVar[int]
    BINARIES_FIELD_NUMBER: _ClassVar[int]
    INTEGRATIONS_FIELD_NUMBER: _ClassVar[int]
    KEYWORDS_FIELD_NUMBER: _ClassVar[int]
    CHANGELOG_FIELD_NUMBER: _ClassVar[int]
    DEPRECATION_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    excerpt: _wrappers_pb2.StringValue
    image: _wrappers_pb2.StringValue
    dark_mode_image: _wrappers_pb2.StringValue
    labels: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    version: _wrappers_pb2.StringValue
    items: _containers.RepeatedCompositeFieldContainer[ExtensionItemData]
    integration_details: _containers.RepeatedCompositeFieldContainer[IntegrationDetail]
    is_hidden: _wrappers_pb2.BoolValue
    binaries: _containers.RepeatedCompositeFieldContainer[ExtensionBinary]
    integrations: _containers.RepeatedScalarFieldContainer[str]
    keywords: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    changelog: _containers.RepeatedCompositeFieldContainer[ChangelogEntry]
    deprecation: Deprecation
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., excerpt: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., image: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., dark_mode_image: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., labels: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., version: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., items: _Optional[_Iterable[_Union[ExtensionItemData, _Mapping]]] = ..., integration_details: _Optional[_Iterable[_Union[IntegrationDetail, _Mapping]]] = ..., is_hidden: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., binaries: _Optional[_Iterable[_Union[ExtensionBinary, _Mapping]]] = ..., integrations: _Optional[_Iterable[str]] = ..., keywords: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., changelog: _Optional[_Iterable[_Union[ChangelogEntry, _Mapping]]] = ..., deprecation: _Optional[_Union[Deprecation, _Mapping]] = ...) -> None: ...

class ExtensionItemData(_message.Message):
    __slots__ = ("name", "description", "target_domain", "data", "binaries", "is_mandatory", "internal_id", "unique_id", "permission_resource")
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    TARGET_DOMAIN_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    BINARIES_FIELD_NUMBER: _ClassVar[int]
    IS_MANDATORY_FIELD_NUMBER: _ClassVar[int]
    INTERNAL_ID_FIELD_NUMBER: _ClassVar[int]
    UNIQUE_ID_FIELD_NUMBER: _ClassVar[int]
    PERMISSION_RESOURCE_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    target_domain: TargetDomain
    data: _struct_pb2.Struct
    binaries: _containers.RepeatedCompositeFieldContainer[ExtensionItemBinary]
    is_mandatory: _wrappers_pb2.BoolValue
    internal_id: _wrappers_pb2.Int32Value
    unique_id: _wrappers_pb2.StringValue
    permission_resource: ExtensionItem.PermissionResource
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., target_domain: _Optional[_Union[TargetDomain, str]] = ..., data: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ..., binaries: _Optional[_Iterable[_Union[ExtensionItemBinary, _Mapping]]] = ..., is_mandatory: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., internal_id: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., unique_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., permission_resource: _Optional[_Union[ExtensionItem.PermissionResource, str]] = ...) -> None: ...
