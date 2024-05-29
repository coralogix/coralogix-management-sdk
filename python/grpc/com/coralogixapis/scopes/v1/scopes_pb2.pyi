from com.coralogixapis.scopes.v1 import entity_type_pb2 as _entity_type_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Scope(_message.Message):
    __slots__ = ("id", "display_name", "description", "team_id", "filters", "default_expression")
    ID_FIELD_NUMBER: _ClassVar[int]
    DISPLAY_NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    DEFAULT_EXPRESSION_FIELD_NUMBER: _ClassVar[int]
    id: str
    display_name: str
    description: str
    team_id: int
    filters: _containers.RepeatedCompositeFieldContainer[Filter]
    default_expression: str
    def __init__(self, id: _Optional[str] = ..., display_name: _Optional[str] = ..., description: _Optional[str] = ..., team_id: _Optional[int] = ..., filters: _Optional[_Iterable[_Union[Filter, _Mapping]]] = ..., default_expression: _Optional[str] = ...) -> None: ...

class Filter(_message.Message):
    __slots__ = ("entity_type", "expression")
    ENTITY_TYPE_FIELD_NUMBER: _ClassVar[int]
    EXPRESSION_FIELD_NUMBER: _ClassVar[int]
    entity_type: _entity_type_pb2.EntityType
    expression: str
    def __init__(self, entity_type: _Optional[_Union[_entity_type_pb2.EntityType, str]] = ..., expression: _Optional[str] = ...) -> None: ...

class GetTeamScopesRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetTeamScopesByIdsRequest(_message.Message):
    __slots__ = ("ids",)
    IDS_FIELD_NUMBER: _ClassVar[int]
    ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, ids: _Optional[_Iterable[str]] = ...) -> None: ...

class GetScopesResponse(_message.Message):
    __slots__ = ("scopes",)
    SCOPES_FIELD_NUMBER: _ClassVar[int]
    scopes: _containers.RepeatedCompositeFieldContainer[Scope]
    def __init__(self, scopes: _Optional[_Iterable[_Union[Scope, _Mapping]]] = ...) -> None: ...

class CreateScopeRequest(_message.Message):
    __slots__ = ("display_name", "description", "filters", "default_expression")
    DISPLAY_NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    DEFAULT_EXPRESSION_FIELD_NUMBER: _ClassVar[int]
    display_name: str
    description: str
    filters: _containers.RepeatedCompositeFieldContainer[Filter]
    default_expression: str
    def __init__(self, display_name: _Optional[str] = ..., description: _Optional[str] = ..., filters: _Optional[_Iterable[_Union[Filter, _Mapping]]] = ..., default_expression: _Optional[str] = ...) -> None: ...

class CreateScopeResponse(_message.Message):
    __slots__ = ("scope",)
    SCOPE_FIELD_NUMBER: _ClassVar[int]
    scope: Scope
    def __init__(self, scope: _Optional[_Union[Scope, _Mapping]] = ...) -> None: ...

class UpdateScopeRequest(_message.Message):
    __slots__ = ("id", "display_name", "description", "filters", "default_expression")
    ID_FIELD_NUMBER: _ClassVar[int]
    DISPLAY_NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    DEFAULT_EXPRESSION_FIELD_NUMBER: _ClassVar[int]
    id: str
    display_name: str
    description: str
    filters: _containers.RepeatedCompositeFieldContainer[Filter]
    default_expression: str
    def __init__(self, id: _Optional[str] = ..., display_name: _Optional[str] = ..., description: _Optional[str] = ..., filters: _Optional[_Iterable[_Union[Filter, _Mapping]]] = ..., default_expression: _Optional[str] = ...) -> None: ...

class UpdateScopeResponse(_message.Message):
    __slots__ = ("scope",)
    SCOPE_FIELD_NUMBER: _ClassVar[int]
    scope: Scope
    def __init__(self, scope: _Optional[_Union[Scope, _Mapping]] = ...) -> None: ...

class DeleteScopeRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class DeleteScopeResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
