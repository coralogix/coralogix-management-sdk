from com.coralogixapis.views.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogixapis.views.v1 import view_folder_pb2 as _view_folder_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateViewFolderRequest(_message.Message):
    __slots__ = ("name",)
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class CreateViewFolderResponse(_message.Message):
    __slots__ = ("folder",)
    FOLDER_FIELD_NUMBER: _ClassVar[int]
    folder: _view_folder_pb2.ViewFolder
    def __init__(self, folder: _Optional[_Union[_view_folder_pb2.ViewFolder, _Mapping]] = ...) -> None: ...

class GetViewFolderRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetViewFolderResponse(_message.Message):
    __slots__ = ("folder",)
    FOLDER_FIELD_NUMBER: _ClassVar[int]
    folder: _view_folder_pb2.ViewFolder
    def __init__(self, folder: _Optional[_Union[_view_folder_pb2.ViewFolder, _Mapping]] = ...) -> None: ...

class ListViewFoldersRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListViewFoldersResponse(_message.Message):
    __slots__ = ("folders",)
    FOLDERS_FIELD_NUMBER: _ClassVar[int]
    folders: _containers.RepeatedCompositeFieldContainer[_view_folder_pb2.ViewFolder]
    def __init__(self, folders: _Optional[_Iterable[_Union[_view_folder_pb2.ViewFolder, _Mapping]]] = ...) -> None: ...

class DeleteViewFolderRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class DeleteViewFolderResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ReplaceViewFolderRequest(_message.Message):
    __slots__ = ("folder",)
    FOLDER_FIELD_NUMBER: _ClassVar[int]
    folder: _view_folder_pb2.ViewFolder
    def __init__(self, folder: _Optional[_Union[_view_folder_pb2.ViewFolder, _Mapping]] = ...) -> None: ...

class ReplaceViewFolderResponse(_message.Message):
    __slots__ = ("folder",)
    FOLDER_FIELD_NUMBER: _ClassVar[int]
    folder: _view_folder_pb2.ViewFolder
    def __init__(self, folder: _Optional[_Union[_view_folder_pb2.ViewFolder, _Mapping]] = ...) -> None: ...
