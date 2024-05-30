from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AlertFilters(_message.Message):
    __slots__ = ("severities", "metadata", "alias", "text", "ratio_alerts", "filter_type")
    class LogSeverity(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        LOG_SEVERITY_DEBUG_OR_UNSPECIFIED: _ClassVar[AlertFilters.LogSeverity]
        LOG_SEVERITY_VERBOSE: _ClassVar[AlertFilters.LogSeverity]
        LOG_SEVERITY_INFO: _ClassVar[AlertFilters.LogSeverity]
        LOG_SEVERITY_WARNING: _ClassVar[AlertFilters.LogSeverity]
        LOG_SEVERITY_ERROR: _ClassVar[AlertFilters.LogSeverity]
        LOG_SEVERITY_CRITICAL: _ClassVar[AlertFilters.LogSeverity]
    LOG_SEVERITY_DEBUG_OR_UNSPECIFIED: AlertFilters.LogSeverity
    LOG_SEVERITY_VERBOSE: AlertFilters.LogSeverity
    LOG_SEVERITY_INFO: AlertFilters.LogSeverity
    LOG_SEVERITY_WARNING: AlertFilters.LogSeverity
    LOG_SEVERITY_ERROR: AlertFilters.LogSeverity
    LOG_SEVERITY_CRITICAL: AlertFilters.LogSeverity
    class FilterType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        FILTER_TYPE_TEXT_OR_UNSPECIFIED: _ClassVar[AlertFilters.FilterType]
        FILTER_TYPE_TEMPLATE: _ClassVar[AlertFilters.FilterType]
        FILTER_TYPE_RATIO: _ClassVar[AlertFilters.FilterType]
        FILTER_TYPE_UNIQUE_COUNT: _ClassVar[AlertFilters.FilterType]
        FILTER_TYPE_TIME_RELATIVE: _ClassVar[AlertFilters.FilterType]
        FILTER_TYPE_METRIC: _ClassVar[AlertFilters.FilterType]
        FILTER_TYPE_TRACING: _ClassVar[AlertFilters.FilterType]
        FILTER_TYPE_FLOW: _ClassVar[AlertFilters.FilterType]
    FILTER_TYPE_TEXT_OR_UNSPECIFIED: AlertFilters.FilterType
    FILTER_TYPE_TEMPLATE: AlertFilters.FilterType
    FILTER_TYPE_RATIO: AlertFilters.FilterType
    FILTER_TYPE_UNIQUE_COUNT: AlertFilters.FilterType
    FILTER_TYPE_TIME_RELATIVE: AlertFilters.FilterType
    FILTER_TYPE_METRIC: AlertFilters.FilterType
    FILTER_TYPE_TRACING: AlertFilters.FilterType
    FILTER_TYPE_FLOW: AlertFilters.FilterType
    class MetadataFilters(_message.Message):
        __slots__ = ("categories", "applications", "subsystems", "computers", "classes", "methods", "ip_addresses")
        CATEGORIES_FIELD_NUMBER: _ClassVar[int]
        APPLICATIONS_FIELD_NUMBER: _ClassVar[int]
        SUBSYSTEMS_FIELD_NUMBER: _ClassVar[int]
        COMPUTERS_FIELD_NUMBER: _ClassVar[int]
        CLASSES_FIELD_NUMBER: _ClassVar[int]
        METHODS_FIELD_NUMBER: _ClassVar[int]
        IP_ADDRESSES_FIELD_NUMBER: _ClassVar[int]
        categories: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        applications: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        subsystems: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        computers: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        classes: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        methods: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        ip_addresses: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        def __init__(self, categories: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., applications: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystems: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., computers: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., classes: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., methods: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., ip_addresses: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...
    class RatioAlert(_message.Message):
        __slots__ = ("alias", "text", "severities", "applications", "subsystems", "group_by")
        ALIAS_FIELD_NUMBER: _ClassVar[int]
        TEXT_FIELD_NUMBER: _ClassVar[int]
        SEVERITIES_FIELD_NUMBER: _ClassVar[int]
        APPLICATIONS_FIELD_NUMBER: _ClassVar[int]
        SUBSYSTEMS_FIELD_NUMBER: _ClassVar[int]
        GROUP_BY_FIELD_NUMBER: _ClassVar[int]
        alias: _wrappers_pb2.StringValue
        text: _wrappers_pb2.StringValue
        severities: _containers.RepeatedScalarFieldContainer[AlertFilters.LogSeverity]
        applications: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        subsystems: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        group_by: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        def __init__(self, alias: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., text: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., severities: _Optional[_Iterable[_Union[AlertFilters.LogSeverity, str]]] = ..., applications: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystems: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., group_by: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...
    SEVERITIES_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    ALIAS_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    RATIO_ALERTS_FIELD_NUMBER: _ClassVar[int]
    FILTER_TYPE_FIELD_NUMBER: _ClassVar[int]
    severities: _containers.RepeatedScalarFieldContainer[AlertFilters.LogSeverity]
    metadata: AlertFilters.MetadataFilters
    alias: _wrappers_pb2.StringValue
    text: _wrappers_pb2.StringValue
    ratio_alerts: _containers.RepeatedCompositeFieldContainer[AlertFilters.RatioAlert]
    filter_type: AlertFilters.FilterType
    def __init__(self, severities: _Optional[_Iterable[_Union[AlertFilters.LogSeverity, str]]] = ..., metadata: _Optional[_Union[AlertFilters.MetadataFilters, _Mapping]] = ..., alias: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., text: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., ratio_alerts: _Optional[_Iterable[_Union[AlertFilters.RatioAlert, _Mapping]]] = ..., filter_type: _Optional[_Union[AlertFilters.FilterType, str]] = ...) -> None: ...
