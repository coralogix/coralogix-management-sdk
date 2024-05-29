from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class QueryRequest(_message.Message):
    __slots__ = ("query", "metadata")
    QUERY_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELD_NUMBER: _ClassVar[int]
    query: _wrappers_pb2.StringValue
    metadata: Metadata
    def __init__(self, query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., metadata: _Optional[_Union[Metadata, _Mapping]] = ...) -> None: ...

class Metadata(_message.Message):
    __slots__ = ("start_date", "end_date", "default_source", "tier", "syntax", "limit", "strict_fields_validation")
    class Tier(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        TIER_UNSPECIFIED: _ClassVar[Metadata.Tier]
        TIER_ARCHIVE: _ClassVar[Metadata.Tier]
        TIER_FREQUENT_SEARCH: _ClassVar[Metadata.Tier]
    TIER_UNSPECIFIED: Metadata.Tier
    TIER_ARCHIVE: Metadata.Tier
    TIER_FREQUENT_SEARCH: Metadata.Tier
    class QuerySyntax(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        QUERY_SYNTAX_UNSPECIFIED: _ClassVar[Metadata.QuerySyntax]
        QUERY_SYNTAX_LUCENE: _ClassVar[Metadata.QuerySyntax]
        QUERY_SYNTAX_DATAPRIME: _ClassVar[Metadata.QuerySyntax]
    QUERY_SYNTAX_UNSPECIFIED: Metadata.QuerySyntax
    QUERY_SYNTAX_LUCENE: Metadata.QuerySyntax
    QUERY_SYNTAX_DATAPRIME: Metadata.QuerySyntax
    START_DATE_FIELD_NUMBER: _ClassVar[int]
    END_DATE_FIELD_NUMBER: _ClassVar[int]
    DEFAULT_SOURCE_FIELD_NUMBER: _ClassVar[int]
    TIER_FIELD_NUMBER: _ClassVar[int]
    SYNTAX_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    STRICT_FIELDS_VALIDATION_FIELD_NUMBER: _ClassVar[int]
    start_date: _timestamp_pb2.Timestamp
    end_date: _timestamp_pb2.Timestamp
    default_source: _wrappers_pb2.StringValue
    tier: Metadata.Tier
    syntax: Metadata.QuerySyntax
    limit: int
    strict_fields_validation: bool
    def __init__(self, start_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., end_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., default_source: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., tier: _Optional[_Union[Metadata.Tier, str]] = ..., syntax: _Optional[_Union[Metadata.QuerySyntax, str]] = ..., limit: _Optional[int] = ..., strict_fields_validation: bool = ...) -> None: ...
