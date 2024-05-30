from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ColorsBy(_message.Message):
    __slots__ = ("stack", "group_by", "aggregation")
    class ColorsByStack(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class ColorsByGroupBy(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class ColorsByAggregation(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    STACK_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_FIELD_NUMBER: _ClassVar[int]
    stack: ColorsBy.ColorsByStack
    group_by: ColorsBy.ColorsByGroupBy
    aggregation: ColorsBy.ColorsByAggregation
    def __init__(self, stack: _Optional[_Union[ColorsBy.ColorsByStack, _Mapping]] = ..., group_by: _Optional[_Union[ColorsBy.ColorsByGroupBy, _Mapping]] = ..., aggregation: _Optional[_Union[ColorsBy.ColorsByAggregation, _Mapping]] = ...) -> None: ...
