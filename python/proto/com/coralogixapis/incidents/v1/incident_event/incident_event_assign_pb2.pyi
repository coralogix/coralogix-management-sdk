from com.coralogixapis.incidents.v1 import assignee_pb2 as _assignee_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class IncidentEventAssign(_message.Message):
    __slots__ = ("assignment",)
    ASSIGNMENT_FIELD_NUMBER: _ClassVar[int]
    assignment: _assignee_pb2.Assignment
    def __init__(self, assignment: _Optional[_Union[_assignee_pb2.Assignment, _Mapping]] = ...) -> None: ...
