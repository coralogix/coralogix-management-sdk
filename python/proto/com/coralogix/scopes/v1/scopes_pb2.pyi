from com.coralogix.scopes.v1 import named_filters_pb2 as _named_filters_pb2
from com.coralogixapis.scopes.v1 import entity_type_pb2 as _entity_type_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Scope(_message.Message):
    __slots__ = ("id", "display_name", "description", "team_id", "named_filters")
    ID_FIELD_NUMBER: _ClassVar[int]
    DISPLAY_NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    NAMED_FILTERS_FIELD_NUMBER: _ClassVar[int]
    id: str
    display_name: str
    description: str
    team_id: int
    named_filters: _containers.RepeatedCompositeFieldContainer[_named_filters_pb2.NamedFilter]
    def __init__(self, id: _Optional[str] = ..., display_name: _Optional[str] = ..., description: _Optional[str] = ..., team_id: _Optional[int] = ..., named_filters: _Optional[_Iterable[_Union[_named_filters_pb2.NamedFilter, _Mapping]]] = ...) -> None: ...

class GetScopeExpressionsRequest(_message.Message):
    __slots__ = ("entity_type",)
    ENTITY_TYPE_FIELD_NUMBER: _ClassVar[int]
    entity_type: _entity_type_pb2.EntityType
    def __init__(self, entity_type: _Optional[_Union[_entity_type_pb2.EntityType, str]] = ...) -> None: ...

class GetScopeExpressionResponse(_message.Message):
    __slots__ = ("formatted_dpxl_expression",)
    FORMATTED_DPXL_EXPRESSION_FIELD_NUMBER: _ClassVar[int]
    formatted_dpxl_expression: str
    def __init__(self, formatted_dpxl_expression: _Optional[str] = ...) -> None: ...

class GetScopesByIdsRequest(_message.Message):
    __slots__ = ("ids",)
    IDS_FIELD_NUMBER: _ClassVar[int]
    ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, ids: _Optional[_Iterable[str]] = ...) -> None: ...

class GetScopesByTeamIdRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetScopesResponse(_message.Message):
    __slots__ = ("scopes",)
    SCOPES_FIELD_NUMBER: _ClassVar[int]
    scopes: _containers.RepeatedCompositeFieldContainer[Scope]
    def __init__(self, scopes: _Optional[_Iterable[_Union[Scope, _Mapping]]] = ...) -> None: ...

class CreateScopeRequest(_message.Message):
    __slots__ = ("id", "display_name", "description", "new_named_filter_definitions", "existing_named_filter_ids")
    ID_FIELD_NUMBER: _ClassVar[int]
    DISPLAY_NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    NEW_NAMED_FILTER_DEFINITIONS_FIELD_NUMBER: _ClassVar[int]
    EXISTING_NAMED_FILTER_IDS_FIELD_NUMBER: _ClassVar[int]
    id: str
    display_name: str
    description: str
    new_named_filter_definitions: _containers.RepeatedCompositeFieldContainer[_named_filters_pb2.NamedFilterCreateDefinition]
    existing_named_filter_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, id: _Optional[str] = ..., display_name: _Optional[str] = ..., description: _Optional[str] = ..., new_named_filter_definitions: _Optional[_Iterable[_Union[_named_filters_pb2.NamedFilterCreateDefinition, _Mapping]]] = ..., existing_named_filter_ids: _Optional[_Iterable[str]] = ...) -> None: ...

class CreateScopeResponse(_message.Message):
    __slots__ = ("scope",)
    SCOPE_FIELD_NUMBER: _ClassVar[int]
    scope: Scope
    def __init__(self, scope: _Optional[_Union[Scope, _Mapping]] = ...) -> None: ...

class UpdateScopeRequest(_message.Message):
    __slots__ = ("id", "display_name", "description", "new_named_filter_definitions", "existing_named_filter_ids")
    ID_FIELD_NUMBER: _ClassVar[int]
    DISPLAY_NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    NEW_NAMED_FILTER_DEFINITIONS_FIELD_NUMBER: _ClassVar[int]
    EXISTING_NAMED_FILTER_IDS_FIELD_NUMBER: _ClassVar[int]
    id: str
    display_name: str
    description: str
    new_named_filter_definitions: _containers.RepeatedCompositeFieldContainer[_named_filters_pb2.NamedFilterCreateDefinition]
    existing_named_filter_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, id: _Optional[str] = ..., display_name: _Optional[str] = ..., description: _Optional[str] = ..., new_named_filter_definitions: _Optional[_Iterable[_Union[_named_filters_pb2.NamedFilterCreateDefinition, _Mapping]]] = ..., existing_named_filter_ids: _Optional[_Iterable[str]] = ...) -> None: ...

class UpdateScopeResponse(_message.Message):
    __slots__ = ("scope",)
    SCOPE_FIELD_NUMBER: _ClassVar[int]
    scope: Scope
    def __init__(self, scope: _Optional[_Union[Scope, _Mapping]] = ...) -> None: ...

class DeleteScopeRequest(_message.Message):
    __slots__ = ("scope_id",)
    SCOPE_ID_FIELD_NUMBER: _ClassVar[int]
    scope_id: str
    def __init__(self, scope_id: _Optional[str] = ...) -> None: ...

class DeleteScopeResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
