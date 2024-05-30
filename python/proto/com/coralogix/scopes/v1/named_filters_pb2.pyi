from com.coralogixapis.scopes.v1 import entity_type_pb2 as _entity_type_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class NamedFilterCreateDefinition(_message.Message):
    __slots__ = ("id", "entity_type", "serialized_expression", "display_name")
    ID_FIELD_NUMBER: _ClassVar[int]
    ENTITY_TYPE_FIELD_NUMBER: _ClassVar[int]
    SERIALIZED_EXPRESSION_FIELD_NUMBER: _ClassVar[int]
    DISPLAY_NAME_FIELD_NUMBER: _ClassVar[int]
    id: str
    entity_type: _entity_type_pb2.EntityType
    serialized_expression: str
    display_name: str
    def __init__(self, id: _Optional[str] = ..., entity_type: _Optional[_Union[_entity_type_pb2.EntityType, str]] = ..., serialized_expression: _Optional[str] = ..., display_name: _Optional[str] = ...) -> None: ...

class NamedFilterUpdateDefinition(_message.Message):
    __slots__ = ("id", "entity_type", "serialized_expression", "display_name")
    ID_FIELD_NUMBER: _ClassVar[int]
    ENTITY_TYPE_FIELD_NUMBER: _ClassVar[int]
    SERIALIZED_EXPRESSION_FIELD_NUMBER: _ClassVar[int]
    DISPLAY_NAME_FIELD_NUMBER: _ClassVar[int]
    id: str
    entity_type: _entity_type_pb2.EntityType
    serialized_expression: str
    display_name: str
    def __init__(self, id: _Optional[str] = ..., entity_type: _Optional[_Union[_entity_type_pb2.EntityType, str]] = ..., serialized_expression: _Optional[str] = ..., display_name: _Optional[str] = ...) -> None: ...

class NamedFilter(_message.Message):
    __slots__ = ("id", "team_id", "entity_type", "serialized_expression", "display_name")
    ID_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    ENTITY_TYPE_FIELD_NUMBER: _ClassVar[int]
    SERIALIZED_EXPRESSION_FIELD_NUMBER: _ClassVar[int]
    DISPLAY_NAME_FIELD_NUMBER: _ClassVar[int]
    id: str
    team_id: int
    entity_type: _entity_type_pb2.EntityType
    serialized_expression: str
    display_name: str
    def __init__(self, id: _Optional[str] = ..., team_id: _Optional[int] = ..., entity_type: _Optional[_Union[_entity_type_pb2.EntityType, str]] = ..., serialized_expression: _Optional[str] = ..., display_name: _Optional[str] = ...) -> None: ...

class GetNamedFiltersByTeamIdRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetNamedFiltersByIdsRequest(_message.Message):
    __slots__ = ("named_filter_ids",)
    NAMED_FILTER_IDS_FIELD_NUMBER: _ClassVar[int]
    named_filter_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, named_filter_ids: _Optional[_Iterable[str]] = ...) -> None: ...

class GetNamedFiltersResponse(_message.Message):
    __slots__ = ("named_filters",)
    NAMED_FILTERS_FIELD_NUMBER: _ClassVar[int]
    named_filters: _containers.RepeatedCompositeFieldContainer[NamedFilter]
    def __init__(self, named_filters: _Optional[_Iterable[_Union[NamedFilter, _Mapping]]] = ...) -> None: ...

class CreateNamedFiltersRequest(_message.Message):
    __slots__ = ("named_filter_definitions",)
    NAMED_FILTER_DEFINITIONS_FIELD_NUMBER: _ClassVar[int]
    named_filter_definitions: _containers.RepeatedCompositeFieldContainer[NamedFilterCreateDefinition]
    def __init__(self, named_filter_definitions: _Optional[_Iterable[_Union[NamedFilterCreateDefinition, _Mapping]]] = ...) -> None: ...

class CreateNamedFiltersResponse(_message.Message):
    __slots__ = ("named_filters",)
    NAMED_FILTERS_FIELD_NUMBER: _ClassVar[int]
    named_filters: _containers.RepeatedCompositeFieldContainer[NamedFilter]
    def __init__(self, named_filters: _Optional[_Iterable[_Union[NamedFilter, _Mapping]]] = ...) -> None: ...

class UpdateNamedFiltersRequest(_message.Message):
    __slots__ = ("named_filter_definitions",)
    NAMED_FILTER_DEFINITIONS_FIELD_NUMBER: _ClassVar[int]
    named_filter_definitions: _containers.RepeatedCompositeFieldContainer[NamedFilterUpdateDefinition]
    def __init__(self, named_filter_definitions: _Optional[_Iterable[_Union[NamedFilterUpdateDefinition, _Mapping]]] = ...) -> None: ...

class UpdateNamedFiltersResponse(_message.Message):
    __slots__ = ("named_filters",)
    NAMED_FILTERS_FIELD_NUMBER: _ClassVar[int]
    named_filters: _containers.RepeatedCompositeFieldContainer[NamedFilter]
    def __init__(self, named_filters: _Optional[_Iterable[_Union[NamedFilter, _Mapping]]] = ...) -> None: ...

class DeleteNamedFiltersRequest(_message.Message):
    __slots__ = ("named_filter_ids",)
    NAMED_FILTER_IDS_FIELD_NUMBER: _ClassVar[int]
    named_filter_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, named_filter_ids: _Optional[_Iterable[str]] = ...) -> None: ...

class DeleteNamedFiltersResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
