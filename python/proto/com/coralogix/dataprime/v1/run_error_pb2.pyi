from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.dataprime.v1 import run_request_pb2 as _run_request_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Error(_message.Message):
    __slots__ = ("message", "affected_components")
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    AFFECTED_COMPONENTS_FIELD_NUMBER: _ClassVar[int]
    message: _wrappers_pb2.StringValue
    affected_components: _containers.RepeatedCompositeFieldContainer[_run_request_pb2.RequestComponent]
    def __init__(self, message: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., affected_components: _Optional[_Iterable[_Union[_run_request_pb2.RequestComponent, _Mapping]]] = ...) -> None: ...
