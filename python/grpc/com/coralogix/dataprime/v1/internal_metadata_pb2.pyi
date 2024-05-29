from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class MetadataEnums(_message.Message):
    __slots__ = ()
    class Tier(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        TIER_UNSPECIFIED: _ClassVar[MetadataEnums.Tier]
        TIER_ARCHIVE: _ClassVar[MetadataEnums.Tier]
        TIER_FREQUENT_SEARCH: _ClassVar[MetadataEnums.Tier]
    TIER_UNSPECIFIED: MetadataEnums.Tier
    TIER_ARCHIVE: MetadataEnums.Tier
    TIER_FREQUENT_SEARCH: MetadataEnums.Tier
    class QuerySyntax(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        QUERY_SYNTAX_UNSPECIFIED: _ClassVar[MetadataEnums.QuerySyntax]
        QUERY_SYNTAX_LUCENE: _ClassVar[MetadataEnums.QuerySyntax]
        QUERY_SYNTAX_DATAPRIME: _ClassVar[MetadataEnums.QuerySyntax]
    QUERY_SYNTAX_UNSPECIFIED: MetadataEnums.QuerySyntax
    QUERY_SYNTAX_LUCENE: MetadataEnums.QuerySyntax
    QUERY_SYNTAX_DATAPRIME: MetadataEnums.QuerySyntax
    def __init__(self) -> None: ...

class QueryMetadata(_message.Message):
    __slots__ = ("start_date", "end_date", "tier", "limit", "strict_fields_validation", "skip_rbac_filters")
    START_DATE_FIELD_NUMBER: _ClassVar[int]
    END_DATE_FIELD_NUMBER: _ClassVar[int]
    TIER_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    STRICT_FIELDS_VALIDATION_FIELD_NUMBER: _ClassVar[int]
    SKIP_RBAC_FILTERS_FIELD_NUMBER: _ClassVar[int]
    start_date: _timestamp_pb2.Timestamp
    end_date: _timestamp_pb2.Timestamp
    tier: MetadataEnums.Tier
    limit: int
    strict_fields_validation: bool
    skip_rbac_filters: bool
    def __init__(self, start_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., end_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., tier: _Optional[_Union[MetadataEnums.Tier, str]] = ..., limit: _Optional[int] = ..., strict_fields_validation: bool = ..., skip_rbac_filters: bool = ...) -> None: ...

class TextQueryMetadata(_message.Message):
    __slots__ = ("syntax", "default_source")
    SYNTAX_FIELD_NUMBER: _ClassVar[int]
    DEFAULT_SOURCE_FIELD_NUMBER: _ClassVar[int]
    syntax: MetadataEnums.QuerySyntax
    default_source: _wrappers_pb2.StringValue
    def __init__(self, syntax: _Optional[_Union[MetadataEnums.QuerySyntax, str]] = ..., default_source: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
