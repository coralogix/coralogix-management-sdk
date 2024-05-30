from com.coralogixapis.incidents.v1 import incident_pb2 as _incident_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class OrderByFields(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ORDER_BY_FIELDS_UNSPECIFIED: _ClassVar[OrderByFields]
    ORDER_BY_FIELDS_ID: _ClassVar[OrderByFields]
    ORDER_BY_FIELDS_SEVERITY: _ClassVar[OrderByFields]
    ORDER_BY_FIELDS_NAME: _ClassVar[OrderByFields]
    ORDER_BY_FIELDS_CREATED_TIME: _ClassVar[OrderByFields]
    ORDER_BY_FIELDS_CLOSED_TIME: _ClassVar[OrderByFields]

class OrderByDirection(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ORDER_BY_DIRECTION_UNSPECIFIED: _ClassVar[OrderByDirection]
    ORDER_BY_DIRECTION_ASC: _ClassVar[OrderByDirection]
    ORDER_BY_DIRECTION_DESC: _ClassVar[OrderByDirection]
ORDER_BY_FIELDS_UNSPECIFIED: OrderByFields
ORDER_BY_FIELDS_ID: OrderByFields
ORDER_BY_FIELDS_SEVERITY: OrderByFields
ORDER_BY_FIELDS_NAME: OrderByFields
ORDER_BY_FIELDS_CREATED_TIME: OrderByFields
ORDER_BY_FIELDS_CLOSED_TIME: OrderByFields
ORDER_BY_DIRECTION_UNSPECIFIED: OrderByDirection
ORDER_BY_DIRECTION_ASC: OrderByDirection
ORDER_BY_DIRECTION_DESC: OrderByDirection

class OrderBy(_message.Message):
    __slots__ = ("incident_field", "contextual_label", "direction")
    INCIDENT_FIELD_FIELD_NUMBER: _ClassVar[int]
    CONTEXTUAL_LABEL_FIELD_NUMBER: _ClassVar[int]
    DIRECTION_FIELD_NUMBER: _ClassVar[int]
    incident_field: _incident_pb2.IncidentFields
    contextual_label: _wrappers_pb2.StringValue
    direction: OrderByDirection
    def __init__(self, incident_field: _Optional[_Union[_incident_pb2.IncidentFields, str]] = ..., contextual_label: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., direction: _Optional[_Union[OrderByDirection, str]] = ...) -> None: ...

class GroupBy(_message.Message):
    __slots__ = ("incident_field", "contextual_label", "order_by_direction")
    INCIDENT_FIELD_FIELD_NUMBER: _ClassVar[int]
    CONTEXTUAL_LABEL_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_DIRECTION_FIELD_NUMBER: _ClassVar[int]
    incident_field: _incident_pb2.IncidentFields
    contextual_label: _wrappers_pb2.StringValue
    order_by_direction: OrderByDirection
    def __init__(self, incident_field: _Optional[_Union[_incident_pb2.IncidentFields, str]] = ..., contextual_label: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., order_by_direction: _Optional[_Union[OrderByDirection, str]] = ...) -> None: ...

class IncidentSearchQuery(_message.Message):
    __slots__ = ("query", "incident_field", "contextual_label")
    QUERY_FIELD_NUMBER: _ClassVar[int]
    INCIDENT_FIELD_FIELD_NUMBER: _ClassVar[int]
    CONTEXTUAL_LABEL_FIELD_NUMBER: _ClassVar[int]
    query: _wrappers_pb2.StringValue
    incident_field: _incident_pb2.IncidentFields
    contextual_label: _wrappers_pb2.StringValue
    def __init__(self, query: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., incident_field: _Optional[_Union[_incident_pb2.IncidentFields, str]] = ..., contextual_label: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
