from com.coralogix.xdr.ingestion.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.xdr.ingestion.v1 import security_report_pb2 as _security_report_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class PostSecurityReportRequest(_message.Message):
    __slots__ = ("security_report",)
    SECURITY_REPORT_FIELD_NUMBER: _ClassVar[int]
    security_report: _security_report_pb2.SecurityReport
    def __init__(self, security_report: _Optional[_Union[_security_report_pb2.SecurityReport, _Mapping]] = ...) -> None: ...

class PostSecurityReportResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetExecutionStatusRequest(_message.Message):
    __slots__ = ("provider", "account")
    PROVIDER_FIELD_NUMBER: _ClassVar[int]
    ACCOUNT_FIELD_NUMBER: _ClassVar[int]
    provider: _wrappers_pb2.StringValue
    account: _wrappers_pb2.StringValue
    def __init__(self, provider: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., account: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetExecutionStatusResponse(_message.Message):
    __slots__ = ("status", "execution_id")
    class Status(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        IDLE: _ClassVar[GetExecutionStatusResponse.Status]
        READY: _ClassVar[GetExecutionStatusResponse.Status]
    IDLE: GetExecutionStatusResponse.Status
    READY: GetExecutionStatusResponse.Status
    STATUS_FIELD_NUMBER: _ClassVar[int]
    EXECUTION_ID_FIELD_NUMBER: _ClassVar[int]
    status: GetExecutionStatusResponse.Status
    execution_id: _wrappers_pb2.StringValue
    def __init__(self, status: _Optional[_Union[GetExecutionStatusResponse.Status, str]] = ..., execution_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetExecutionStatusWithDeactivationsRequest(_message.Message):
    __slots__ = ("provider", "account")
    PROVIDER_FIELD_NUMBER: _ClassVar[int]
    ACCOUNT_FIELD_NUMBER: _ClassVar[int]
    provider: _wrappers_pb2.StringValue
    account: _wrappers_pb2.StringValue
    def __init__(self, provider: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., account: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetExecutionStatusWithDeactivationsResponse(_message.Message):
    __slots__ = ("idle", "ready")
    class Deactivation(_message.Message):
        __slots__ = ("details", "deactivated_at")
        class RuleDeactivation(_message.Message):
            __slots__ = ("test_identity",)
            TEST_IDENTITY_FIELD_NUMBER: _ClassVar[int]
            test_identity: _wrappers_pb2.StringValue
            def __init__(self, test_identity: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
        class ResourceDeactivation(_message.Message):
            __slots__ = ("resource_id",)
            RESOURCE_ID_FIELD_NUMBER: _ClassVar[int]
            resource_id: _wrappers_pb2.StringValue
            def __init__(self, resource_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
        class RuleAndResourceDeactivation(_message.Message):
            __slots__ = ("test_identity", "resource_id")
            TEST_IDENTITY_FIELD_NUMBER: _ClassVar[int]
            RESOURCE_ID_FIELD_NUMBER: _ClassVar[int]
            test_identity: _wrappers_pb2.StringValue
            resource_id: _wrappers_pb2.StringValue
            def __init__(self, test_identity: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., resource_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
        class DeactivationDetails(_message.Message):
            __slots__ = ("rule_deactivation", "resource_deactivation", "rule_and_resource_deactivation")
            RULE_DEACTIVATION_FIELD_NUMBER: _ClassVar[int]
            RESOURCE_DEACTIVATION_FIELD_NUMBER: _ClassVar[int]
            RULE_AND_RESOURCE_DEACTIVATION_FIELD_NUMBER: _ClassVar[int]
            rule_deactivation: GetExecutionStatusWithDeactivationsResponse.Deactivation.RuleDeactivation
            resource_deactivation: GetExecutionStatusWithDeactivationsResponse.Deactivation.ResourceDeactivation
            rule_and_resource_deactivation: GetExecutionStatusWithDeactivationsResponse.Deactivation.RuleAndResourceDeactivation
            def __init__(self, rule_deactivation: _Optional[_Union[GetExecutionStatusWithDeactivationsResponse.Deactivation.RuleDeactivation, _Mapping]] = ..., resource_deactivation: _Optional[_Union[GetExecutionStatusWithDeactivationsResponse.Deactivation.ResourceDeactivation, _Mapping]] = ..., rule_and_resource_deactivation: _Optional[_Union[GetExecutionStatusWithDeactivationsResponse.Deactivation.RuleAndResourceDeactivation, _Mapping]] = ...) -> None: ...
        DETAILS_FIELD_NUMBER: _ClassVar[int]
        DEACTIVATED_AT_FIELD_NUMBER: _ClassVar[int]
        details: GetExecutionStatusWithDeactivationsResponse.Deactivation.DeactivationDetails
        deactivated_at: _timestamp_pb2.Timestamp
        def __init__(self, details: _Optional[_Union[GetExecutionStatusWithDeactivationsResponse.Deactivation.DeactivationDetails, _Mapping]] = ..., deactivated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
    class Idle(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class Ready(_message.Message):
        __slots__ = ("execution_id", "deactivations")
        EXECUTION_ID_FIELD_NUMBER: _ClassVar[int]
        DEACTIVATIONS_FIELD_NUMBER: _ClassVar[int]
        execution_id: _wrappers_pb2.StringValue
        deactivations: _containers.RepeatedCompositeFieldContainer[GetExecutionStatusWithDeactivationsResponse.Deactivation]
        def __init__(self, execution_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., deactivations: _Optional[_Iterable[_Union[GetExecutionStatusWithDeactivationsResponse.Deactivation, _Mapping]]] = ...) -> None: ...
    IDLE_FIELD_NUMBER: _ClassVar[int]
    READY_FIELD_NUMBER: _ClassVar[int]
    idle: GetExecutionStatusWithDeactivationsResponse.Idle
    ready: GetExecutionStatusWithDeactivationsResponse.Ready
    def __init__(self, idle: _Optional[_Union[GetExecutionStatusWithDeactivationsResponse.Idle, _Mapping]] = ..., ready: _Optional[_Union[GetExecutionStatusWithDeactivationsResponse.Ready, _Mapping]] = ...) -> None: ...
