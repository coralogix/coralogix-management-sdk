from com.coralogixapis.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogixapis.dashboards.v1.common import folder_pb2 as _folder_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateDashboardFolderRequest(_message.Message):
    __slots__ = ("request_id", "folder")
    REQUEST_ID_FIELD_NUMBER: _ClassVar[int]
    FOLDER_FIELD_NUMBER: _ClassVar[int]
    request_id: _wrappers_pb2.StringValue
    folder: _folder_pb2.DashboardFolder
    def __init__(self, request_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., folder: _Optional[_Union[_folder_pb2.DashboardFolder, _Mapping]] = ...) -> None: ...

class CreateDashboardFolderResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ReplaceDashboardFolderRequest(_message.Message):
    __slots__ = ("request_id", "folder")
    REQUEST_ID_FIELD_NUMBER: _ClassVar[int]
    FOLDER_FIELD_NUMBER: _ClassVar[int]
    request_id: _wrappers_pb2.StringValue
    folder: _folder_pb2.DashboardFolder
    def __init__(self, request_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., folder: _Optional[_Union[_folder_pb2.DashboardFolder, _Mapping]] = ...) -> None: ...

class ReplaceDashboardFolderResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DeleteDashboardFolderRequest(_message.Message):
    __slots__ = ("request_id", "folder_id")
    REQUEST_ID_FIELD_NUMBER: _ClassVar[int]
    FOLDER_ID_FIELD_NUMBER: _ClassVar[int]
    request_id: _wrappers_pb2.StringValue
    folder_id: _wrappers_pb2.StringValue
    def __init__(self, request_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., folder_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class DeleteDashboardFolderResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListDashboardFoldersRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListDashboardFoldersResponse(_message.Message):
    __slots__ = ("folder",)
    FOLDER_FIELD_NUMBER: _ClassVar[int]
    folder: _containers.RepeatedCompositeFieldContainer[_folder_pb2.DashboardFolder]
    def __init__(self, folder: _Optional[_Iterable[_Union[_folder_pb2.DashboardFolder, _Mapping]]] = ...) -> None: ...
