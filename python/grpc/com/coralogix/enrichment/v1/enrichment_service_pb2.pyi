from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.enrichment.v1 import enrichment_pb2 as _enrichment_pb2
from com.coralogix.enrichment.v1 import enrichment_request_model_pb2 as _enrichment_request_model_pb2
from com.coralogix.enrichment.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.enrichment.v1 import enrichment_type_pb2 as _enrichment_type_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetEnrichmentLimitRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetEnrichmentLimitResponse(_message.Message):
    __slots__ = ("limit", "used")
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    USED_FIELD_NUMBER: _ClassVar[int]
    limit: int
    used: int
    def __init__(self, limit: _Optional[int] = ..., used: _Optional[int] = ...) -> None: ...

class GetEnrichmentsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetEnrichmentsResponse(_message.Message):
    __slots__ = ("enrichments",)
    ENRICHMENTS_FIELD_NUMBER: _ClassVar[int]
    enrichments: _containers.RepeatedCompositeFieldContainer[_enrichment_pb2.Enrichment]
    def __init__(self, enrichments: _Optional[_Iterable[_Union[_enrichment_pb2.Enrichment, _Mapping]]] = ...) -> None: ...

class AddEnrichmentsRequest(_message.Message):
    __slots__ = ("request_enrichments",)
    REQUEST_ENRICHMENTS_FIELD_NUMBER: _ClassVar[int]
    request_enrichments: _containers.RepeatedCompositeFieldContainer[_enrichment_request_model_pb2.EnrichmentRequestModel]
    def __init__(self, request_enrichments: _Optional[_Iterable[_Union[_enrichment_request_model_pb2.EnrichmentRequestModel, _Mapping]]] = ...) -> None: ...

class AddEnrichmentsResponse(_message.Message):
    __slots__ = ("enrichments",)
    ENRICHMENTS_FIELD_NUMBER: _ClassVar[int]
    enrichments: _containers.RepeatedCompositeFieldContainer[_enrichment_pb2.Enrichment]
    def __init__(self, enrichments: _Optional[_Iterable[_Union[_enrichment_pb2.Enrichment, _Mapping]]] = ...) -> None: ...

class RemoveEnrichmentsRequest(_message.Message):
    __slots__ = ("enrichment_ids",)
    ENRICHMENT_IDS_FIELD_NUMBER: _ClassVar[int]
    enrichment_ids: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.UInt32Value]
    def __init__(self, enrichment_ids: _Optional[_Iterable[_Union[_wrappers_pb2.UInt32Value, _Mapping]]] = ...) -> None: ...

class RemoveEnrichmentsResponse(_message.Message):
    __slots__ = ("remaining_enrichments",)
    REMAINING_ENRICHMENTS_FIELD_NUMBER: _ClassVar[int]
    remaining_enrichments: _containers.RepeatedCompositeFieldContainer[_enrichment_pb2.Enrichment]
    def __init__(self, remaining_enrichments: _Optional[_Iterable[_Union[_enrichment_pb2.Enrichment, _Mapping]]] = ...) -> None: ...

class AtomicOverwriteEnrichmentsRequest(_message.Message):
    __slots__ = ("enrichment_type", "enrichment_fields")
    ENRICHMENT_TYPE_FIELD_NUMBER: _ClassVar[int]
    ENRICHMENT_FIELDS_FIELD_NUMBER: _ClassVar[int]
    enrichment_type: _enrichment_type_pb2.EnrichmentType
    enrichment_fields: _containers.RepeatedCompositeFieldContainer[_enrichment_request_model_pb2.EnrichmentFieldDefinition]
    def __init__(self, enrichment_type: _Optional[_Union[_enrichment_type_pb2.EnrichmentType, _Mapping]] = ..., enrichment_fields: _Optional[_Iterable[_Union[_enrichment_request_model_pb2.EnrichmentFieldDefinition, _Mapping]]] = ...) -> None: ...

class AtomicOverwriteEnrichmentsResponse(_message.Message):
    __slots__ = ("enrichments",)
    ENRICHMENTS_FIELD_NUMBER: _ClassVar[int]
    enrichments: _containers.RepeatedCompositeFieldContainer[_enrichment_pb2.Enrichment]
    def __init__(self, enrichments: _Optional[_Iterable[_Union[_enrichment_pb2.Enrichment, _Mapping]]] = ...) -> None: ...
