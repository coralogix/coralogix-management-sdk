from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Expression(_message.Message):
    __slots__ = ("predicate",)
    PREDICATE_FIELD_NUMBER: _ClassVar[int]
    predicate: str
    def __init__(self, predicate: _Optional[str] = ...) -> None: ...

class SerializedExpression(_message.Message):
    __slots__ = ("content",)
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    content: str
    def __init__(self, content: _Optional[str] = ...) -> None: ...

class Diagnostic(_message.Message):
    __slots__ = ()
    class Error(_message.Message):
        __slots__ = ("message", "range")
        MESSAGE_FIELD_NUMBER: _ClassVar[int]
        RANGE_FIELD_NUMBER: _ClassVar[int]
        message: str
        range: Diagnostic.Range
        def __init__(self, message: _Optional[str] = ..., range: _Optional[_Union[Diagnostic.Range, _Mapping]] = ...) -> None: ...
    class Warning(_message.Message):
        __slots__ = ("message", "range")
        MESSAGE_FIELD_NUMBER: _ClassVar[int]
        RANGE_FIELD_NUMBER: _ClassVar[int]
        message: str
        range: Diagnostic.Range
        def __init__(self, message: _Optional[str] = ..., range: _Optional[_Union[Diagnostic.Range, _Mapping]] = ...) -> None: ...
    class Range(_message.Message):
        __slots__ = ("to",)
        FROM_FIELD_NUMBER: _ClassVar[int]
        TO_FIELD_NUMBER: _ClassVar[int]
        to: Diagnostic.Position
        def __init__(self, to: _Optional[_Union[Diagnostic.Position, _Mapping]] = ..., **kwargs) -> None: ...
    class Position(_message.Message):
        __slots__ = ("line", "col")
        LINE_FIELD_NUMBER: _ClassVar[int]
        COL_FIELD_NUMBER: _ClassVar[int]
        line: int
        col: int
        def __init__(self, line: _Optional[int] = ..., col: _Optional[int] = ...) -> None: ...
    def __init__(self) -> None: ...

class Source(_message.Message):
    __slots__ = ("logs", "spans")
    class Logs(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class Spans(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    LOGS_FIELD_NUMBER: _ClassVar[int]
    SPANS_FIELD_NUMBER: _ClassVar[int]
    logs: Source.Logs
    spans: Source.Spans
    def __init__(self, logs: _Optional[_Union[Source.Logs, _Mapping]] = ..., spans: _Optional[_Union[Source.Spans, _Mapping]] = ...) -> None: ...

class EntityType(_message.Message):
    __slots__ = ("logs", "spans")
    class Logs(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class Spans(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    LOGS_FIELD_NUMBER: _ClassVar[int]
    SPANS_FIELD_NUMBER: _ClassVar[int]
    logs: EntityType.Logs
    spans: EntityType.Spans
    def __init__(self, logs: _Optional[_Union[EntityType.Logs, _Mapping]] = ..., spans: _Optional[_Union[EntityType.Spans, _Mapping]] = ...) -> None: ...
