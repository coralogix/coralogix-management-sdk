from com.coralogixapis.dashboards.v1.common import observation_field_pb2 as _observation_field_pb2
from com.coralogixapis.dashboards.v1.common import span_field_pb2 as _span_field_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Filter(_message.Message):
    __slots__ = ("source", "enabled", "collapsed")
    class Source(_message.Message):
        __slots__ = ("logs", "spans", "metrics")
        LOGS_FIELD_NUMBER: _ClassVar[int]
        SPANS_FIELD_NUMBER: _ClassVar[int]
        METRICS_FIELD_NUMBER: _ClassVar[int]
        logs: Filter.LogsFilter
        spans: Filter.SpansFilter
        metrics: Filter.MetricsFilter
        def __init__(self, logs: _Optional[_Union[Filter.LogsFilter, _Mapping]] = ..., spans: _Optional[_Union[Filter.SpansFilter, _Mapping]] = ..., metrics: _Optional[_Union[Filter.MetricsFilter, _Mapping]] = ...) -> None: ...
    class LogsFilter(_message.Message):
        __slots__ = ("field", "operator", "observation_field")
        FIELD_FIELD_NUMBER: _ClassVar[int]
        OPERATOR_FIELD_NUMBER: _ClassVar[int]
        OBSERVATION_FIELD_FIELD_NUMBER: _ClassVar[int]
        field: _wrappers_pb2.StringValue
        operator: Filter.Operator
        observation_field: _observation_field_pb2.ObservationField
        def __init__(self, field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., operator: _Optional[_Union[Filter.Operator, _Mapping]] = ..., observation_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
    class SpansFilter(_message.Message):
        __slots__ = ("field", "operator")
        FIELD_FIELD_NUMBER: _ClassVar[int]
        OPERATOR_FIELD_NUMBER: _ClassVar[int]
        field: _span_field_pb2.SpanField
        operator: Filter.Operator
        def __init__(self, field: _Optional[_Union[_span_field_pb2.SpanField, _Mapping]] = ..., operator: _Optional[_Union[Filter.Operator, _Mapping]] = ...) -> None: ...
    class MetricsFilter(_message.Message):
        __slots__ = ("metric", "label", "operator")
        METRIC_FIELD_NUMBER: _ClassVar[int]
        LABEL_FIELD_NUMBER: _ClassVar[int]
        OPERATOR_FIELD_NUMBER: _ClassVar[int]
        metric: _wrappers_pb2.StringValue
        label: _wrappers_pb2.StringValue
        operator: Filter.Operator
        def __init__(self, metric: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., label: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., operator: _Optional[_Union[Filter.Operator, _Mapping]] = ...) -> None: ...
    class Operator(_message.Message):
        __slots__ = ("equals", "not_equals")
        EQUALS_FIELD_NUMBER: _ClassVar[int]
        NOT_EQUALS_FIELD_NUMBER: _ClassVar[int]
        equals: Filter.Equals
        not_equals: Filter.NotEquals
        def __init__(self, equals: _Optional[_Union[Filter.Equals, _Mapping]] = ..., not_equals: _Optional[_Union[Filter.NotEquals, _Mapping]] = ...) -> None: ...
    class Equals(_message.Message):
        __slots__ = ("selection",)
        class Selection(_message.Message):
            __slots__ = ("all", "list")
            class AllSelection(_message.Message):
                __slots__ = ()
                def __init__(self) -> None: ...
            class ListSelection(_message.Message):
                __slots__ = ("values",)
                VALUES_FIELD_NUMBER: _ClassVar[int]
                values: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
                def __init__(self, values: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...
            ALL_FIELD_NUMBER: _ClassVar[int]
            LIST_FIELD_NUMBER: _ClassVar[int]
            all: Filter.Equals.Selection.AllSelection
            list: Filter.Equals.Selection.ListSelection
            def __init__(self, all: _Optional[_Union[Filter.Equals.Selection.AllSelection, _Mapping]] = ..., list: _Optional[_Union[Filter.Equals.Selection.ListSelection, _Mapping]] = ...) -> None: ...
        SELECTION_FIELD_NUMBER: _ClassVar[int]
        selection: Filter.Equals.Selection
        def __init__(self, selection: _Optional[_Union[Filter.Equals.Selection, _Mapping]] = ...) -> None: ...
    class NotEquals(_message.Message):
        __slots__ = ("selection",)
        class Selection(_message.Message):
            __slots__ = ("list",)
            class ListSelection(_message.Message):
                __slots__ = ("values",)
                VALUES_FIELD_NUMBER: _ClassVar[int]
                values: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
                def __init__(self, values: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...
            LIST_FIELD_NUMBER: _ClassVar[int]
            list: Filter.NotEquals.Selection.ListSelection
            def __init__(self, list: _Optional[_Union[Filter.NotEquals.Selection.ListSelection, _Mapping]] = ...) -> None: ...
        SELECTION_FIELD_NUMBER: _ClassVar[int]
        selection: Filter.NotEquals.Selection
        def __init__(self, selection: _Optional[_Union[Filter.NotEquals.Selection, _Mapping]] = ...) -> None: ...
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    ENABLED_FIELD_NUMBER: _ClassVar[int]
    COLLAPSED_FIELD_NUMBER: _ClassVar[int]
    source: Filter.Source
    enabled: _wrappers_pb2.BoolValue
    collapsed: _wrappers_pb2.BoolValue
    def __init__(self, source: _Optional[_Union[Filter.Source, _Mapping]] = ..., enabled: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., collapsed: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
