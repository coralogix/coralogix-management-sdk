from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class LogMetaFieldsValues(_message.Message):
    __slots__ = ("application_name_values", "severity_values", "subsystem_name_values")
    APPLICATION_NAME_VALUES_FIELD_NUMBER: _ClassVar[int]
    SEVERITY_VALUES_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_NAME_VALUES_FIELD_NUMBER: _ClassVar[int]
    application_name_values: _wrappers_pb2.StringValue
    severity_values: _wrappers_pb2.StringValue
    subsystem_name_values: _wrappers_pb2.StringValue
    def __init__(self, application_name_values: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., severity_values: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem_name_values: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
