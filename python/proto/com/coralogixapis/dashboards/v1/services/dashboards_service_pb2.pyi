from com.coralogixapis.dashboards.v1.ast import dashboard_pb2 as _dashboard_pb2
from com.coralogixapis.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateDashboardRequest(_message.Message):
    __slots__ = ("request_id", "dashboard")
    REQUEST_ID_FIELD_NUMBER: _ClassVar[int]
    DASHBOARD_FIELD_NUMBER: _ClassVar[int]
    request_id: _wrappers_pb2.StringValue
    dashboard: _dashboard_pb2.Dashboard
    def __init__(self, request_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., dashboard: _Optional[_Union[_dashboard_pb2.Dashboard, _Mapping]] = ...) -> None: ...

class CreateDashboardResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ReplaceDashboardRequest(_message.Message):
    __slots__ = ("request_id", "dashboard")
    REQUEST_ID_FIELD_NUMBER: _ClassVar[int]
    DASHBOARD_FIELD_NUMBER: _ClassVar[int]
    request_id: _wrappers_pb2.StringValue
    dashboard: _dashboard_pb2.Dashboard
    def __init__(self, request_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., dashboard: _Optional[_Union[_dashboard_pb2.Dashboard, _Mapping]] = ...) -> None: ...

class ReplaceDashboardResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DeleteDashboardRequest(_message.Message):
    __slots__ = ("request_id", "dashboard_id")
    REQUEST_ID_FIELD_NUMBER: _ClassVar[int]
    DASHBOARD_ID_FIELD_NUMBER: _ClassVar[int]
    request_id: _wrappers_pb2.StringValue
    dashboard_id: _wrappers_pb2.StringValue
    def __init__(self, request_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., dashboard_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class DeleteDashboardResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetDashboardRequest(_message.Message):
    __slots__ = ("dashboard_id",)
    DASHBOARD_ID_FIELD_NUMBER: _ClassVar[int]
    dashboard_id: _wrappers_pb2.StringValue
    def __init__(self, dashboard_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetDashboardResponse(_message.Message):
    __slots__ = ("dashboard",)
    DASHBOARD_FIELD_NUMBER: _ClassVar[int]
    dashboard: _dashboard_pb2.Dashboard
    def __init__(self, dashboard: _Optional[_Union[_dashboard_pb2.Dashboard, _Mapping]] = ...) -> None: ...

class PinDashboardRequest(_message.Message):
    __slots__ = ("request_id", "dashboard_id")
    REQUEST_ID_FIELD_NUMBER: _ClassVar[int]
    DASHBOARD_ID_FIELD_NUMBER: _ClassVar[int]
    request_id: _wrappers_pb2.StringValue
    dashboard_id: _wrappers_pb2.StringValue
    def __init__(self, request_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., dashboard_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class PinDashboardResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class UnpinDashboardRequest(_message.Message):
    __slots__ = ("request_id", "dashboard_id")
    REQUEST_ID_FIELD_NUMBER: _ClassVar[int]
    DASHBOARD_ID_FIELD_NUMBER: _ClassVar[int]
    request_id: _wrappers_pb2.StringValue
    dashboard_id: _wrappers_pb2.StringValue
    def __init__(self, request_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., dashboard_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class UnpinDashboardResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ReplaceDefaultDashboardRequest(_message.Message):
    __slots__ = ("request_id", "dashboard_id")
    REQUEST_ID_FIELD_NUMBER: _ClassVar[int]
    DASHBOARD_ID_FIELD_NUMBER: _ClassVar[int]
    request_id: _wrappers_pb2.StringValue
    dashboard_id: _wrappers_pb2.StringValue
    def __init__(self, request_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., dashboard_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class ReplaceDefaultDashboardResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class AssignDashboardFolderRequest(_message.Message):
    __slots__ = ("request_id", "dashboard_id", "folder_id")
    REQUEST_ID_FIELD_NUMBER: _ClassVar[int]
    DASHBOARD_ID_FIELD_NUMBER: _ClassVar[int]
    FOLDER_ID_FIELD_NUMBER: _ClassVar[int]
    request_id: _wrappers_pb2.StringValue
    dashboard_id: _wrappers_pb2.StringValue
    folder_id: _wrappers_pb2.StringValue
    def __init__(self, request_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., dashboard_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., folder_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class AssignDashboardFolderResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
