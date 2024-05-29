from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.enrichment.v1 import enrichment_type_pb2 as _enrichment_type_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Enrichment(_message.Message):
    __slots__ = ("id", "field_name", "enrichment_type", "enriched_field_name")
    ID_FIELD_NUMBER: _ClassVar[int]
    FIELD_NAME_FIELD_NUMBER: _ClassVar[int]
    ENRICHMENT_TYPE_FIELD_NUMBER: _ClassVar[int]
    ENRICHED_FIELD_NAME_FIELD_NUMBER: _ClassVar[int]
    id: int
    field_name: str
    enrichment_type: _enrichment_type_pb2.EnrichmentType
    enriched_field_name: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[int] = ..., field_name: _Optional[str] = ..., enrichment_type: _Optional[_Union[_enrichment_type_pb2.EnrichmentType, _Mapping]] = ..., enriched_field_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
