from com.coralogix.dataprime.ast.v1 import types_pb2 as _types_pb2
from com.coralogix.dataprime.ast.v1 import query_pb2 as _query_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SchemaLookup(_message.Message):
    __slots__ = ("schemas",)
    SCHEMAS_FIELD_NUMBER: _ClassVar[int]
    schemas: _containers.RepeatedCompositeFieldContainer[NamedSchema]
    def __init__(self, schemas: _Optional[_Iterable[_Union[NamedSchema, _Mapping]]] = ...) -> None: ...

class NamedSchema(_message.Message):
    __slots__ = ("source", "schema")
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    SCHEMA_FIELD_NUMBER: _ClassVar[int]
    source: _query_pb2.Source
    schema: Schema
    def __init__(self, source: _Optional[_Union[_query_pb2.Source, _Mapping]] = ..., schema: _Optional[_Union[Schema, _Mapping]] = ...) -> None: ...

class Schema(_message.Message):
    __slots__ = ("userdata_fields", "metadata_fields", "labels_fields")
    USERDATA_FIELDS_FIELD_NUMBER: _ClassVar[int]
    METADATA_FIELDS_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELDS_FIELD_NUMBER: _ClassVar[int]
    userdata_fields: _containers.RepeatedCompositeFieldContainer[SchemaField]
    metadata_fields: _containers.RepeatedCompositeFieldContainer[SchemaField]
    labels_fields: _containers.RepeatedCompositeFieldContainer[SchemaField]
    def __init__(self, userdata_fields: _Optional[_Iterable[_Union[SchemaField, _Mapping]]] = ..., metadata_fields: _Optional[_Iterable[_Union[SchemaField, _Mapping]]] = ..., labels_fields: _Optional[_Iterable[_Union[SchemaField, _Mapping]]] = ...) -> None: ...

class SchemaField(_message.Message):
    __slots__ = ("path", "lucene_path", "datatype")
    PATH_FIELD_NUMBER: _ClassVar[int]
    LUCENE_PATH_FIELD_NUMBER: _ClassVar[int]
    DATATYPE_FIELD_NUMBER: _ClassVar[int]
    path: _containers.RepeatedScalarFieldContainer[str]
    lucene_path: LucenePath
    datatype: _types_pb2.Datatype
    def __init__(self, path: _Optional[_Iterable[str]] = ..., lucene_path: _Optional[_Union[LucenePath, _Mapping]] = ..., datatype: _Optional[_Union[_types_pb2.Datatype, _Mapping]] = ...) -> None: ...

class LucenePath(_message.Message):
    __slots__ = ("path",)
    PATH_FIELD_NUMBER: _ClassVar[int]
    path: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, path: _Optional[_Iterable[str]] = ...) -> None: ...
