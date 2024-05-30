from com.coralogixapis.dashboards.v1.common import observation_field_pb2 as _observation_field_pb2
from com.coralogixapis.dashboards.v1.common import query_pb2 as _query_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Annotation(_message.Message):
    __slots__ = ("id", "name", "enabled", "source")
    class Source(_message.Message):
        __slots__ = ("metrics", "logs", "spans")
        METRICS_FIELD_NUMBER: _ClassVar[int]
        LOGS_FIELD_NUMBER: _ClassVar[int]
        SPANS_FIELD_NUMBER: _ClassVar[int]
        metrics: Annotation.MetricsSource
        logs: Annotation.LogsSource
        spans: Annotation.SpansSource
        def __init__(self, metrics: _Optional[_Union[Annotation.MetricsSource, _Mapping]] = ..., logs: _Optional[_Union[Annotation.LogsSource, _Mapping]] = ..., spans: _Optional[_Union[Annotation.SpansSource, _Mapping]] = ...) -> None: ...
    class MetricsSource(_message.Message):
        __slots__ = ("promql_query", "strategy", "message_template", "labels")
        class Strategy(_message.Message):
            __slots__ = ("start_time_metric",)
            START_TIME_METRIC_FIELD_NUMBER: _ClassVar[int]
            start_time_metric: Annotation.MetricsSource.StartTimeMetric
            def __init__(self, start_time_metric: _Optional[_Union[Annotation.MetricsSource.StartTimeMetric, _Mapping]] = ...) -> None: ...
        class StartTimeMetric(_message.Message):
            __slots__ = ()
            def __init__(self) -> None: ...
        PROMQL_QUERY_FIELD_NUMBER: _ClassVar[int]
        STRATEGY_FIELD_NUMBER: _ClassVar[int]
        MESSAGE_TEMPLATE_FIELD_NUMBER: _ClassVar[int]
        LABELS_FIELD_NUMBER: _ClassVar[int]
        promql_query: _query_pb2.PromQlQuery
        strategy: Annotation.MetricsSource.Strategy
        message_template: _wrappers_pb2.StringValue
        labels: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        def __init__(self, promql_query: _Optional[_Union[_query_pb2.PromQlQuery, _Mapping]] = ..., strategy: _Optional[_Union[Annotation.MetricsSource.Strategy, _Mapping]] = ..., message_template: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., labels: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...
    class LogsSource(_message.Message):
        __slots__ = ("lucene_query", "strategy", "message_template", "label_fields")
        class Strategy(_message.Message):
            __slots__ = ("instant", "range", "duration")
            class Instant(_message.Message):
                __slots__ = ("timestamp_field",)
                TIMESTAMP_FIELD_FIELD_NUMBER: _ClassVar[int]
                timestamp_field: _observation_field_pb2.ObservationField
                def __init__(self, timestamp_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
            class Range(_message.Message):
                __slots__ = ("start_timestamp_field", "end_timestamp_field")
                START_TIMESTAMP_FIELD_FIELD_NUMBER: _ClassVar[int]
                END_TIMESTAMP_FIELD_FIELD_NUMBER: _ClassVar[int]
                start_timestamp_field: _observation_field_pb2.ObservationField
                end_timestamp_field: _observation_field_pb2.ObservationField
                def __init__(self, start_timestamp_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ..., end_timestamp_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
            class Duration(_message.Message):
                __slots__ = ("start_timestamp_field", "duration_field")
                START_TIMESTAMP_FIELD_FIELD_NUMBER: _ClassVar[int]
                DURATION_FIELD_FIELD_NUMBER: _ClassVar[int]
                start_timestamp_field: _observation_field_pb2.ObservationField
                duration_field: _observation_field_pb2.ObservationField
                def __init__(self, start_timestamp_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ..., duration_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
            INSTANT_FIELD_NUMBER: _ClassVar[int]
            RANGE_FIELD_NUMBER: _ClassVar[int]
            DURATION_FIELD_NUMBER: _ClassVar[int]
            instant: Annotation.LogsSource.Strategy.Instant
            range: Annotation.LogsSource.Strategy.Range
            duration: Annotation.LogsSource.Strategy.Duration
            def __init__(self, instant: _Optional[_Union[Annotation.LogsSource.Strategy.Instant, _Mapping]] = ..., range: _Optional[_Union[Annotation.LogsSource.Strategy.Range, _Mapping]] = ..., duration: _Optional[_Union[Annotation.LogsSource.Strategy.Duration, _Mapping]] = ...) -> None: ...
        LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
        STRATEGY_FIELD_NUMBER: _ClassVar[int]
        MESSAGE_TEMPLATE_FIELD_NUMBER: _ClassVar[int]
        LABEL_FIELDS_FIELD_NUMBER: _ClassVar[int]
        lucene_query: _query_pb2.LuceneQuery
        strategy: Annotation.LogsSource.Strategy
        message_template: _wrappers_pb2.StringValue
        label_fields: _containers.RepeatedCompositeFieldContainer[_observation_field_pb2.ObservationField]
        def __init__(self, lucene_query: _Optional[_Union[_query_pb2.LuceneQuery, _Mapping]] = ..., strategy: _Optional[_Union[Annotation.LogsSource.Strategy, _Mapping]] = ..., message_template: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., label_fields: _Optional[_Iterable[_Union[_observation_field_pb2.ObservationField, _Mapping]]] = ...) -> None: ...
    class SpansSource(_message.Message):
        __slots__ = ("lucene_query", "strategy", "message_template", "label_fields")
        class Strategy(_message.Message):
            __slots__ = ("instant", "range", "duration")
            class Instant(_message.Message):
                __slots__ = ("timestamp_field",)
                TIMESTAMP_FIELD_FIELD_NUMBER: _ClassVar[int]
                timestamp_field: _observation_field_pb2.ObservationField
                def __init__(self, timestamp_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
            class Range(_message.Message):
                __slots__ = ("start_timestamp_field", "end_timestamp_field")
                START_TIMESTAMP_FIELD_FIELD_NUMBER: _ClassVar[int]
                END_TIMESTAMP_FIELD_FIELD_NUMBER: _ClassVar[int]
                start_timestamp_field: _observation_field_pb2.ObservationField
                end_timestamp_field: _observation_field_pb2.ObservationField
                def __init__(self, start_timestamp_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ..., end_timestamp_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
            class Duration(_message.Message):
                __slots__ = ("start_timestamp_field", "duration_field")
                START_TIMESTAMP_FIELD_FIELD_NUMBER: _ClassVar[int]
                DURATION_FIELD_FIELD_NUMBER: _ClassVar[int]
                start_timestamp_field: _observation_field_pb2.ObservationField
                duration_field: _observation_field_pb2.ObservationField
                def __init__(self, start_timestamp_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ..., duration_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
            INSTANT_FIELD_NUMBER: _ClassVar[int]
            RANGE_FIELD_NUMBER: _ClassVar[int]
            DURATION_FIELD_NUMBER: _ClassVar[int]
            instant: Annotation.SpansSource.Strategy.Instant
            range: Annotation.SpansSource.Strategy.Range
            duration: Annotation.SpansSource.Strategy.Duration
            def __init__(self, instant: _Optional[_Union[Annotation.SpansSource.Strategy.Instant, _Mapping]] = ..., range: _Optional[_Union[Annotation.SpansSource.Strategy.Range, _Mapping]] = ..., duration: _Optional[_Union[Annotation.SpansSource.Strategy.Duration, _Mapping]] = ...) -> None: ...
        LUCENE_QUERY_FIELD_NUMBER: _ClassVar[int]
        STRATEGY_FIELD_NUMBER: _ClassVar[int]
        MESSAGE_TEMPLATE_FIELD_NUMBER: _ClassVar[int]
        LABEL_FIELDS_FIELD_NUMBER: _ClassVar[int]
        lucene_query: _query_pb2.LuceneQuery
        strategy: Annotation.SpansSource.Strategy
        message_template: _wrappers_pb2.StringValue
        label_fields: _containers.RepeatedCompositeFieldContainer[_observation_field_pb2.ObservationField]
        def __init__(self, lucene_query: _Optional[_Union[_query_pb2.LuceneQuery, _Mapping]] = ..., strategy: _Optional[_Union[Annotation.SpansSource.Strategy, _Mapping]] = ..., message_template: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., label_fields: _Optional[_Iterable[_Union[_observation_field_pb2.ObservationField, _Mapping]]] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    ENABLED_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    enabled: _wrappers_pb2.BoolValue
    source: Annotation.Source
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., enabled: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., source: _Optional[_Union[Annotation.Source, _Mapping]] = ...) -> None: ...
