from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ObjectStoreWarning(_message.Message):
    __slots__ = ("metastore_warning", "block_warning")
    class MetastoreWarning(_message.Message):
        __slots__ = ("no_files_found_warning",)
        class NoFilesFound(_message.Message):
            __slots__ = ()
            def __init__(self) -> None: ...
        NO_FILES_FOUND_WARNING_FIELD_NUMBER: _ClassVar[int]
        no_files_found_warning: ObjectStoreWarning.MetastoreWarning.NoFilesFound
        def __init__(self, no_files_found_warning: _Optional[_Union[ObjectStoreWarning.MetastoreWarning.NoFilesFound, _Mapping]] = ...) -> None: ...
    class BlockWarning(_message.Message):
        __slots__ = ("file_not_found_warning", "access_denied_warning", "read_failed_warning")
        class FileNotFound(_message.Message):
            __slots__ = ()
            def __init__(self) -> None: ...
        class AccessDenied(_message.Message):
            __slots__ = ()
            def __init__(self) -> None: ...
        class ReadFailed(_message.Message):
            __slots__ = ()
            def __init__(self) -> None: ...
        FILE_NOT_FOUND_WARNING_FIELD_NUMBER: _ClassVar[int]
        ACCESS_DENIED_WARNING_FIELD_NUMBER: _ClassVar[int]
        READ_FAILED_WARNING_FIELD_NUMBER: _ClassVar[int]
        file_not_found_warning: ObjectStoreWarning.BlockWarning.FileNotFound
        access_denied_warning: ObjectStoreWarning.BlockWarning.AccessDenied
        read_failed_warning: ObjectStoreWarning.BlockWarning.ReadFailed
        def __init__(self, file_not_found_warning: _Optional[_Union[ObjectStoreWarning.BlockWarning.FileNotFound, _Mapping]] = ..., access_denied_warning: _Optional[_Union[ObjectStoreWarning.BlockWarning.AccessDenied, _Mapping]] = ..., read_failed_warning: _Optional[_Union[ObjectStoreWarning.BlockWarning.ReadFailed, _Mapping]] = ...) -> None: ...
    METASTORE_WARNING_FIELD_NUMBER: _ClassVar[int]
    BLOCK_WARNING_FIELD_NUMBER: _ClassVar[int]
    metastore_warning: ObjectStoreWarning.MetastoreWarning
    block_warning: ObjectStoreWarning.BlockWarning
    def __init__(self, metastore_warning: _Optional[_Union[ObjectStoreWarning.MetastoreWarning, _Mapping]] = ..., block_warning: _Optional[_Union[ObjectStoreWarning.BlockWarning, _Mapping]] = ...) -> None: ...
