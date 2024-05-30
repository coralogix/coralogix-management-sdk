from com.coralogixapis.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetTeamSettingsRequest(_message.Message):
    __slots__ = ("request_id",)
    REQUEST_ID_FIELD_NUMBER: _ClassVar[int]
    request_id: _wrappers_pb2.StringValue
    def __init__(self, request_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetTeamSettingsResponse(_message.Message):
    __slots__ = ("max_group_by_logs", "max_group_by_metrics")
    MAX_GROUP_BY_LOGS_FIELD_NUMBER: _ClassVar[int]
    MAX_GROUP_BY_METRICS_FIELD_NUMBER: _ClassVar[int]
    max_group_by_logs: _wrappers_pb2.Int32Value
    max_group_by_metrics: _wrappers_pb2.Int32Value
    def __init__(self, max_group_by_logs: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., max_group_by_metrics: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
