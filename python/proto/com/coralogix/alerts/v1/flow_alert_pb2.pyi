from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class FlowOperator(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    AND: _ClassVar[FlowOperator]
    OR: _ClassVar[FlowOperator]
AND: FlowOperator
OR: FlowOperator

class FlowStage(_message.Message):
    __slots__ = ("groups", "timeframe")
    GROUPS_FIELD_NUMBER: _ClassVar[int]
    TIMEFRAME_FIELD_NUMBER: _ClassVar[int]
    groups: _containers.RepeatedCompositeFieldContainer[FlowGroup]
    timeframe: FlowTimeframe
    def __init__(self, groups: _Optional[_Iterable[_Union[FlowGroup, _Mapping]]] = ..., timeframe: _Optional[_Union[FlowTimeframe, _Mapping]] = ...) -> None: ...

class FlowAlert(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    NOT_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., **kwargs) -> None: ...

class FlowAlerts(_message.Message):
    __slots__ = ("op", "values")
    OP_FIELD_NUMBER: _ClassVar[int]
    VALUES_FIELD_NUMBER: _ClassVar[int]
    op: FlowOperator
    values: _containers.RepeatedCompositeFieldContainer[FlowAlert]
    def __init__(self, op: _Optional[_Union[FlowOperator, str]] = ..., values: _Optional[_Iterable[_Union[FlowAlert, _Mapping]]] = ...) -> None: ...

class FlowGroup(_message.Message):
    __slots__ = ("alerts", "nextOp")
    ALERTS_FIELD_NUMBER: _ClassVar[int]
    NEXTOP_FIELD_NUMBER: _ClassVar[int]
    alerts: FlowAlerts
    nextOp: FlowOperator
    def __init__(self, alerts: _Optional[_Union[FlowAlerts, _Mapping]] = ..., nextOp: _Optional[_Union[FlowOperator, str]] = ...) -> None: ...

class FlowTimeframe(_message.Message):
    __slots__ = ("ms",)
    MS_FIELD_NUMBER: _ClassVar[int]
    ms: _wrappers_pb2.UInt32Value
    def __init__(self, ms: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...
