from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.enrichment.v1 import custom_enrichment_pb2 as _custom_enrichment_pb2
from com.coralogix.enrichment.v1 import audit_log_pb2 as _audit_log_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class File(_message.Message):
    __slots__ = ("name", "extension", "textual", "binary")
    NAME_FIELD_NUMBER: _ClassVar[int]
    EXTENSION_FIELD_NUMBER: _ClassVar[int]
    TEXTUAL_FIELD_NUMBER: _ClassVar[int]
    BINARY_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    extension: _wrappers_pb2.StringValue
    textual: _wrappers_pb2.StringValue
    binary: _wrappers_pb2.BytesValue
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., extension: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., textual: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., binary: _Optional[_Union[_wrappers_pb2.BytesValue, _Mapping]] = ...) -> None: ...

class GetCustomEnrichmentRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.UInt32Value
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...

class GetCustomEnrichmentResponse(_message.Message):
    __slots__ = ("custom_enrichment",)
    CUSTOM_ENRICHMENT_FIELD_NUMBER: _ClassVar[int]
    custom_enrichment: _custom_enrichment_pb2.CustomEnrichment
    def __init__(self, custom_enrichment: _Optional[_Union[_custom_enrichment_pb2.CustomEnrichment, _Mapping]] = ...) -> None: ...

class GetCustomEnrichmentsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetCustomEnrichmentsResponse(_message.Message):
    __slots__ = ("custom_enrichments",)
    CUSTOM_ENRICHMENTS_FIELD_NUMBER: _ClassVar[int]
    custom_enrichments: _containers.RepeatedCompositeFieldContainer[_custom_enrichment_pb2.CustomEnrichment]
    def __init__(self, custom_enrichments: _Optional[_Iterable[_Union[_custom_enrichment_pb2.CustomEnrichment, _Mapping]]] = ...) -> None: ...

class CreateCustomEnrichmentRequest(_message.Message):
    __slots__ = ("name", "description", "file")
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    FILE_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    file: File
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., file: _Optional[_Union[File, _Mapping]] = ...) -> None: ...

class CreateCustomEnrichmentResponse(_message.Message):
    __slots__ = ("message", "custom_enrichment")
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    CUSTOM_ENRICHMENT_FIELD_NUMBER: _ClassVar[int]
    message: str
    custom_enrichment: _custom_enrichment_pb2.CustomEnrichment
    def __init__(self, message: _Optional[str] = ..., custom_enrichment: _Optional[_Union[_custom_enrichment_pb2.CustomEnrichment, _Mapping]] = ...) -> None: ...

class UpdateCustomEnrichmentRequest(_message.Message):
    __slots__ = ("custom_enrichment_id", "name", "description", "file")
    CUSTOM_ENRICHMENT_ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    FILE_FIELD_NUMBER: _ClassVar[int]
    custom_enrichment_id: _wrappers_pb2.UInt32Value
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    file: File
    def __init__(self, custom_enrichment_id: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., file: _Optional[_Union[File, _Mapping]] = ...) -> None: ...

class UpdateCustomEnrichmentResponse(_message.Message):
    __slots__ = ("message", "custom_enrichment")
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    CUSTOM_ENRICHMENT_FIELD_NUMBER: _ClassVar[int]
    message: str
    custom_enrichment: _custom_enrichment_pb2.CustomEnrichment
    def __init__(self, message: _Optional[str] = ..., custom_enrichment: _Optional[_Union[_custom_enrichment_pb2.CustomEnrichment, _Mapping]] = ...) -> None: ...

class DeleteCustomEnrichmentRequest(_message.Message):
    __slots__ = ("custom_enrichment_id",)
    CUSTOM_ENRICHMENT_ID_FIELD_NUMBER: _ClassVar[int]
    custom_enrichment_id: _wrappers_pb2.UInt32Value
    def __init__(self, custom_enrichment_id: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...

class DeleteCustomEnrichmentResponse(_message.Message):
    __slots__ = ("message", "custom_enrichment_id")
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    CUSTOM_ENRICHMENT_ID_FIELD_NUMBER: _ClassVar[int]
    message: str
    custom_enrichment_id: int
    def __init__(self, message: _Optional[str] = ..., custom_enrichment_id: _Optional[int] = ...) -> None: ...

class SearchCustomEnrichmentDataRequest(_message.Message):
    __slots__ = ("search_clauses",)
    class SearchClause(_message.Message):
        __slots__ = ("id", "name")
        ID_FIELD_NUMBER: _ClassVar[int]
        NAME_FIELD_NUMBER: _ClassVar[int]
        id: _wrappers_pb2.UInt32Value
        name: _wrappers_pb2.StringValue
        def __init__(self, id: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    SEARCH_CLAUSES_FIELD_NUMBER: _ClassVar[int]
    search_clauses: _containers.RepeatedCompositeFieldContainer[SearchCustomEnrichmentDataRequest.SearchClause]
    def __init__(self, search_clauses: _Optional[_Iterable[_Union[SearchCustomEnrichmentDataRequest.SearchClause, _Mapping]]] = ...) -> None: ...

class SearchCustomEnrichmentDataResponse(_message.Message):
    __slots__ = ("custom_enrichments_data",)
    CUSTOM_ENRICHMENTS_DATA_FIELD_NUMBER: _ClassVar[int]
    custom_enrichments_data: _containers.RepeatedCompositeFieldContainer[_custom_enrichment_pb2.CustomEnrichmentData]
    def __init__(self, custom_enrichments_data: _Optional[_Iterable[_Union[_custom_enrichment_pb2.CustomEnrichmentData, _Mapping]]] = ...) -> None: ...
