from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from com.coralogix.tags.v1 import tag_type_pb2 as _tag_type_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Status(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    STATUS_UNSPECIFIED: _ClassVar[Status]
    STATUS_SUCCESSFUL: _ClassVar[Status]
STATUS_UNSPECIFIED: Status
STATUS_SUCCESSFUL: Status

class Tag(_message.Message):
    __slots__ = ("id", "key", "name", "company_id", "status", "icon_url", "timestamp", "application", "subsystem", "updated_at", "created_at", "type")
    ID_FIELD_NUMBER: _ClassVar[int]
    KEY_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    COMPANY_ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    ICON_URL_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.UInt64Value
    key: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    company_id: _wrappers_pb2.UInt32Value
    status: Status
    icon_url: _wrappers_pb2.StringValue
    timestamp: _timestamp_pb2.Timestamp
    application: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    subsystem: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    updated_at: _timestamp_pb2.Timestamp
    created_at: _timestamp_pb2.Timestamp
    type: _tag_type_pb2.TagType
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.UInt64Value, _Mapping]] = ..., key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., company_id: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., status: _Optional[_Union[Status, str]] = ..., icon_url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., application: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystem: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., type: _Optional[_Union[_tag_type_pb2.TagType, _Mapping]] = ...) -> None: ...
