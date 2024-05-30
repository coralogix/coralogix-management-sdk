from com.coralogix.outgoing_webhooks.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogix.outgoing_webhooks.v1 import outgoing_webhook_pb2 as _outgoing_webhook_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ListOutgoingWebhookTypesRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListOutgoingWebhookTypesResponse(_message.Message):
    __slots__ = ("webhooks",)
    class OutgoingWebhookType(_message.Message):
        __slots__ = ("type", "label", "count")
        TYPE_FIELD_NUMBER: _ClassVar[int]
        LABEL_FIELD_NUMBER: _ClassVar[int]
        COUNT_FIELD_NUMBER: _ClassVar[int]
        type: _outgoing_webhook_pb2.WebhookType
        label: _wrappers_pb2.StringValue
        count: _wrappers_pb2.UInt32Value
        def __init__(self, type: _Optional[_Union[_outgoing_webhook_pb2.WebhookType, str]] = ..., label: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., count: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...
    WEBHOOKS_FIELD_NUMBER: _ClassVar[int]
    webhooks: _containers.RepeatedCompositeFieldContainer[ListOutgoingWebhookTypesResponse.OutgoingWebhookType]
    def __init__(self, webhooks: _Optional[_Iterable[_Union[ListOutgoingWebhookTypesResponse.OutgoingWebhookType, _Mapping]]] = ...) -> None: ...

class GetOutgoingWebhookTypeDetailsRequest(_message.Message):
    __slots__ = ("type",)
    TYPE_FIELD_NUMBER: _ClassVar[int]
    type: _outgoing_webhook_pb2.WebhookType
    def __init__(self, type: _Optional[_Union[_outgoing_webhook_pb2.WebhookType, str]] = ...) -> None: ...

class GetOutgoingWebhookTypeDetailsResponse(_message.Message):
    __slots__ = ("details",)
    DETAILS_FIELD_NUMBER: _ClassVar[int]
    details: _outgoing_webhook_pb2.OutgoingWebhookDetails
    def __init__(self, details: _Optional[_Union[_outgoing_webhook_pb2.OutgoingWebhookDetails, _Mapping]] = ...) -> None: ...

class ListOutgoingWebhooksRequest(_message.Message):
    __slots__ = ("type",)
    TYPE_FIELD_NUMBER: _ClassVar[int]
    type: _outgoing_webhook_pb2.WebhookType
    def __init__(self, type: _Optional[_Union[_outgoing_webhook_pb2.WebhookType, str]] = ...) -> None: ...

class ListOutgoingWebhooksResponse(_message.Message):
    __slots__ = ("deployed",)
    DEPLOYED_FIELD_NUMBER: _ClassVar[int]
    deployed: _containers.RepeatedCompositeFieldContainer[_outgoing_webhook_pb2.OutgoingWebhookSummary]
    def __init__(self, deployed: _Optional[_Iterable[_Union[_outgoing_webhook_pb2.OutgoingWebhookSummary, _Mapping]]] = ...) -> None: ...

class ListAllOutgoingWebhooksRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListAllOutgoingWebhooksResponse(_message.Message):
    __slots__ = ("deployed",)
    DEPLOYED_FIELD_NUMBER: _ClassVar[int]
    deployed: _containers.RepeatedCompositeFieldContainer[_outgoing_webhook_pb2.OutgoingWebhookExtendedSummary]
    def __init__(self, deployed: _Optional[_Iterable[_Union[_outgoing_webhook_pb2.OutgoingWebhookExtendedSummary, _Mapping]]] = ...) -> None: ...

class GetOutgoingWebhookRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class GetOutgoingWebhookResponse(_message.Message):
    __slots__ = ("webhook",)
    WEBHOOK_FIELD_NUMBER: _ClassVar[int]
    webhook: _outgoing_webhook_pb2.OutgoingWebhook
    def __init__(self, webhook: _Optional[_Union[_outgoing_webhook_pb2.OutgoingWebhook, _Mapping]] = ...) -> None: ...

class CreateOutgoingWebhookRequest(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _outgoing_webhook_pb2.OutgoingWebhookInputData
    def __init__(self, data: _Optional[_Union[_outgoing_webhook_pb2.OutgoingWebhookInputData, _Mapping]] = ...) -> None: ...

class CreateOutgoingWebhookResponse(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class UpdateOutgoingWebhookRequest(_message.Message):
    __slots__ = ("id", "data")
    ID_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    id: str
    data: _outgoing_webhook_pb2.OutgoingWebhookInputData
    def __init__(self, id: _Optional[str] = ..., data: _Optional[_Union[_outgoing_webhook_pb2.OutgoingWebhookInputData, _Mapping]] = ...) -> None: ...

class UpdateOutgoingWebhookResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DeleteOutgoingWebhookRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class DeleteOutgoingWebhookResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class TestOutgoingWebhookRequest(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _outgoing_webhook_pb2.OutgoingWebhookInputData
    def __init__(self, data: _Optional[_Union[_outgoing_webhook_pb2.OutgoingWebhookInputData, _Mapping]] = ...) -> None: ...

class TestExistingOutgoingWebhookRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class TestOutgoingWebhookResponse(_message.Message):
    __slots__ = ("success", "failure")
    class Success(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class Failure(_message.Message):
        __slots__ = ("error_message", "display_message", "status_code")
        ERROR_MESSAGE_FIELD_NUMBER: _ClassVar[int]
        DISPLAY_MESSAGE_FIELD_NUMBER: _ClassVar[int]
        STATUS_CODE_FIELD_NUMBER: _ClassVar[int]
        error_message: _wrappers_pb2.StringValue
        display_message: _wrappers_pb2.StringValue
        status_code: _wrappers_pb2.UInt32Value
        def __init__(self, error_message: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., display_message: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., status_code: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    FAILURE_FIELD_NUMBER: _ClassVar[int]
    success: TestOutgoingWebhookResponse.Success
    failure: TestOutgoingWebhookResponse.Failure
    def __init__(self, success: _Optional[_Union[TestOutgoingWebhookResponse.Success, _Mapping]] = ..., failure: _Optional[_Union[TestOutgoingWebhookResponse.Failure, _Mapping]] = ...) -> None: ...

class ListIbmEventNotificationsInstancesRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListIbmEventNotificationsInstancesResponse(_message.Message):
    __slots__ = ("instances",)
    class EventNotificationsInstance(_message.Message):
        __slots__ = ("instance_id", "region_id", "name", "crn", "is_used")
        INSTANCE_ID_FIELD_NUMBER: _ClassVar[int]
        REGION_ID_FIELD_NUMBER: _ClassVar[int]
        NAME_FIELD_NUMBER: _ClassVar[int]
        CRN_FIELD_NUMBER: _ClassVar[int]
        IS_USED_FIELD_NUMBER: _ClassVar[int]
        instance_id: str
        region_id: str
        name: str
        crn: str
        is_used: bool
        def __init__(self, instance_id: _Optional[str] = ..., region_id: _Optional[str] = ..., name: _Optional[str] = ..., crn: _Optional[str] = ..., is_used: bool = ...) -> None: ...
    INSTANCES_FIELD_NUMBER: _ClassVar[int]
    instances: _containers.RepeatedCompositeFieldContainer[ListIbmEventNotificationsInstancesResponse.EventNotificationsInstance]
    def __init__(self, instances: _Optional[_Iterable[_Union[ListIbmEventNotificationsInstancesResponse.EventNotificationsInstance, _Mapping]]] = ...) -> None: ...
