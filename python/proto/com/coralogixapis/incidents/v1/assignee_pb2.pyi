from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Assignment(_message.Message):
    __slots__ = ("assigned_to", "assigned_by")
    ASSIGNED_TO_FIELD_NUMBER: _ClassVar[int]
    ASSIGNED_BY_FIELD_NUMBER: _ClassVar[int]
    assigned_to: UserDetails
    assigned_by: UserDetails
    def __init__(self, assigned_to: _Optional[_Union[UserDetails, _Mapping]] = ..., assigned_by: _Optional[_Union[UserDetails, _Mapping]] = ...) -> None: ...

class UserDetails(_message.Message):
    __slots__ = ("user_id",)
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: _wrappers_pb2.StringValue
    def __init__(self, user_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
