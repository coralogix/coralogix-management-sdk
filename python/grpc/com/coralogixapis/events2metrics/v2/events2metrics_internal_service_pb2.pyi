from com.coralogixapis.events2metrics.v2 import events2metrics_definition_pb2 as _events2metrics_definition_pb2
from com.coralogixapis.logs2metrics.v2 import audit_log_pb2 as _audit_log_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ListE2MRequestInternal(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListE2MResponseInternal(_message.Message):
    __slots__ = ("e2m",)
    E2M_FIELD_NUMBER: _ClassVar[int]
    e2m: _containers.RepeatedCompositeFieldContainer[_events2metrics_definition_pb2.E2M]
    def __init__(self, e2m: _Optional[_Iterable[_Union[_events2metrics_definition_pb2.E2M, _Mapping]]] = ...) -> None: ...

class CreateE2MRequestInternal(_message.Message):
    __slots__ = ("e2m", "is_internal")
    E2M_FIELD_NUMBER: _ClassVar[int]
    IS_INTERNAL_FIELD_NUMBER: _ClassVar[int]
    e2m: _events2metrics_definition_pb2.E2MCreateParams
    is_internal: _wrappers_pb2.BoolValue
    def __init__(self, e2m: _Optional[_Union[_events2metrics_definition_pb2.E2MCreateParams, _Mapping]] = ..., is_internal: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class CreateE2MResponseInternal(_message.Message):
    __slots__ = ("e2m",)
    E2M_FIELD_NUMBER: _ClassVar[int]
    e2m: _events2metrics_definition_pb2.E2M
    def __init__(self, e2m: _Optional[_Union[_events2metrics_definition_pb2.E2M, _Mapping]] = ...) -> None: ...

class ReplaceE2MRequestInternal(_message.Message):
    __slots__ = ("e2m",)
    E2M_FIELD_NUMBER: _ClassVar[int]
    e2m: _events2metrics_definition_pb2.E2M
    def __init__(self, e2m: _Optional[_Union[_events2metrics_definition_pb2.E2M, _Mapping]] = ...) -> None: ...

class ReplaceE2MResponseInternal(_message.Message):
    __slots__ = ("e2m",)
    E2M_FIELD_NUMBER: _ClassVar[int]
    e2m: _events2metrics_definition_pb2.E2M
    def __init__(self, e2m: _Optional[_Union[_events2metrics_definition_pb2.E2M, _Mapping]] = ...) -> None: ...

class DeleteE2MRequestInternal(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class DeleteE2MResponseInternal(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
