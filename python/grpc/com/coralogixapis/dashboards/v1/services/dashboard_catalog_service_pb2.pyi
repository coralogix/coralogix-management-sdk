from com.coralogixapis.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogixapis.dashboards.v1.common import folder_pb2 as _folder_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetDashboardCatalogRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetDashboardCatalogResponse(_message.Message):
    __slots__ = ("items",)
    ITEMS_FIELD_NUMBER: _ClassVar[int]
    items: _containers.RepeatedCompositeFieldContainer[DashboardCatalogItem]
    def __init__(self, items: _Optional[_Iterable[_Union[DashboardCatalogItem, _Mapping]]] = ...) -> None: ...

class DashboardCatalogItem(_message.Message):
    __slots__ = ("id", "name", "description", "is_default", "is_pinned", "create_time", "update_time", "folder")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    IS_DEFAULT_FIELD_NUMBER: _ClassVar[int]
    IS_PINNED_FIELD_NUMBER: _ClassVar[int]
    CREATE_TIME_FIELD_NUMBER: _ClassVar[int]
    UPDATE_TIME_FIELD_NUMBER: _ClassVar[int]
    FOLDER_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    is_default: _wrappers_pb2.BoolValue
    is_pinned: _wrappers_pb2.BoolValue
    create_time: _timestamp_pb2.Timestamp
    update_time: _timestamp_pb2.Timestamp
    folder: _folder_pb2.DashboardFolder
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., is_default: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., is_pinned: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., create_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., update_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., folder: _Optional[_Union[_folder_pb2.DashboardFolder, _Mapping]] = ...) -> None: ...
