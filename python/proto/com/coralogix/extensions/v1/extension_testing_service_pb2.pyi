from com.coralogix.extensions.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.extensions.v1 import extension_pb2 as _extension_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class InitializeTestingRevisionRequest(_message.Message):
    __slots__ = ("extension_data",)
    EXTENSION_DATA_FIELD_NUMBER: _ClassVar[int]
    extension_data: _extension_pb2.ExtensionData
    def __init__(self, extension_data: _Optional[_Union[_extension_pb2.ExtensionData, _Mapping]] = ...) -> None: ...

class InitializeTestingRevisionResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class CleanupTestingRevisionRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class CleanupTestingRevisionResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class TestExtensionRevisionRequest(_message.Message):
    __slots__ = ("extension_data", "cleanup_after_test")
    EXTENSION_DATA_FIELD_NUMBER: _ClassVar[int]
    CLEANUP_AFTER_TEST_FIELD_NUMBER: _ClassVar[int]
    extension_data: _extension_pb2.ExtensionData
    cleanup_after_test: _wrappers_pb2.BoolValue
    def __init__(self, extension_data: _Optional[_Union[_extension_pb2.ExtensionData, _Mapping]] = ..., cleanup_after_test: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class TestExtensionRevisionResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
