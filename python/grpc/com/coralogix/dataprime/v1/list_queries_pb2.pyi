from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ListQueryRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class RunningQuery(_message.Message):
    __slots__ = ("id", "durations_ms", "total_task_duration_ms", "total_tasks", "completed_tasks")
    ID_FIELD_NUMBER: _ClassVar[int]
    DURATIONS_MS_FIELD_NUMBER: _ClassVar[int]
    TOTAL_TASK_DURATION_MS_FIELD_NUMBER: _ClassVar[int]
    TOTAL_TASKS_FIELD_NUMBER: _ClassVar[int]
    COMPLETED_TASKS_FIELD_NUMBER: _ClassVar[int]
    id: str
    durations_ms: int
    total_task_duration_ms: int
    total_tasks: int
    completed_tasks: int
    def __init__(self, id: _Optional[str] = ..., durations_ms: _Optional[int] = ..., total_task_duration_ms: _Optional[int] = ..., total_tasks: _Optional[int] = ..., completed_tasks: _Optional[int] = ...) -> None: ...

class ListQueryResponse(_message.Message):
    __slots__ = ("running_queries",)
    RUNNING_QUERIES_FIELD_NUMBER: _ClassVar[int]
    running_queries: _containers.RepeatedCompositeFieldContainer[RunningQuery]
    def __init__(self, running_queries: _Optional[_Iterable[_Union[RunningQuery, _Mapping]]] = ...) -> None: ...
