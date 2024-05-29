from com.coralogix.scopes.v1 import named_filters_pb2 as _named_filters_pb2
from com.coralogixapis.scopes.v1 import entity_type_pb2 as _entity_type_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetCrossTeamScopeExpressionsByUserIdRequest(_message.Message):
    __slots__ = ("entity_type", "user_id", "team_ids")
    ENTITY_TYPE_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    TEAM_IDS_FIELD_NUMBER: _ClassVar[int]
    entity_type: _entity_type_pb2.EntityType
    user_id: str
    team_ids: _containers.RepeatedScalarFieldContainer[int]
    def __init__(self, entity_type: _Optional[_Union[_entity_type_pb2.EntityType, str]] = ..., user_id: _Optional[str] = ..., team_ids: _Optional[_Iterable[int]] = ...) -> None: ...

class GetCrossTeamScopeExpressionsByUserIdResponse(_message.Message):
    __slots__ = ("user_expressions",)
    USER_EXPRESSIONS_FIELD_NUMBER: _ClassVar[int]
    user_expressions: _containers.RepeatedCompositeFieldContainer[UserExpression]
    def __init__(self, user_expressions: _Optional[_Iterable[_Union[UserExpression, _Mapping]]] = ...) -> None: ...

class UserExpression(_message.Message):
    __slots__ = ("user_id", "team_id", "formatted_dpxl_expression")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    FORMATTED_DPXL_EXPRESSION_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    team_id: int
    formatted_dpxl_expression: str
    def __init__(self, user_id: _Optional[str] = ..., team_id: _Optional[int] = ..., formatted_dpxl_expression: _Optional[str] = ...) -> None: ...
