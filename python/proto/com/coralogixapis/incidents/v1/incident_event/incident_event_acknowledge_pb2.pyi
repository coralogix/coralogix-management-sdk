from com.coralogixapis.incidents.v1 import assignee_pb2 as _assignee_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class IncidentEventAcknowledge(_message.Message):
    __slots__ = ("acknowledged_by",)
    ACKNOWLEDGED_BY_FIELD_NUMBER: _ClassVar[int]
    acknowledged_by: _assignee_pb2.UserDetails
    def __init__(self, acknowledged_by: _Optional[_Union[_assignee_pb2.UserDetails, _Mapping]] = ...) -> None: ...
