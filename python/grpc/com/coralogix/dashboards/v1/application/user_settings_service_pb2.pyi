from com.coralogix.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UpdateUserDataSourceVisibilityRequest(_message.Message):
    __slots__ = ("data_source_visibility",)
    DATA_SOURCE_VISIBILITY_FIELD_NUMBER: _ClassVar[int]
    data_source_visibility: DataSourceVisibility
    def __init__(self, data_source_visibility: _Optional[_Union[DataSourceVisibility, _Mapping]] = ...) -> None: ...

class UpdateUserDataSourceVisibilityResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class SearchUserSettingsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class SearchUserSettingsResponse(_message.Message):
    __slots__ = ("data_source_visibility",)
    DATA_SOURCE_VISIBILITY_FIELD_NUMBER: _ClassVar[int]
    data_source_visibility: DataSourceVisibility
    def __init__(self, data_source_visibility: _Optional[_Union[DataSourceVisibility, _Mapping]] = ...) -> None: ...

class DataSourceVisibility(_message.Message):
    __slots__ = ("logs", "metrics", "traces", "security")
    LOGS_FIELD_NUMBER: _ClassVar[int]
    METRICS_FIELD_NUMBER: _ClassVar[int]
    TRACES_FIELD_NUMBER: _ClassVar[int]
    SECURITY_FIELD_NUMBER: _ClassVar[int]
    logs: _wrappers_pb2.BoolValue
    metrics: _wrappers_pb2.BoolValue
    traces: _wrappers_pb2.BoolValue
    security: _wrappers_pb2.BoolValue
    def __init__(self, logs: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., metrics: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., traces: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., security: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
