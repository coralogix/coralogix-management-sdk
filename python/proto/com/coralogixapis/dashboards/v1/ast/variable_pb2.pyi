from com.coralogixapis.dashboards.v1.common import observation_field_pb2 as _observation_field_pb2
from com.coralogixapis.dashboards.v1.common import order_direction_pb2 as _order_direction_pb2
from com.coralogixapis.dashboards.v1.common import span_field_pb2 as _span_field_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class VariableDisplayType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    VARIABLE_DISPLAY_TYPE_UNSPECIFIED: _ClassVar[VariableDisplayType]
    VARIABLE_DISPLAY_TYPE_LABEL_VALUE: _ClassVar[VariableDisplayType]
    VARIABLE_DISPLAY_TYPE_VALUE: _ClassVar[VariableDisplayType]
    VARIABLE_DISPLAY_TYPE_NOTHING: _ClassVar[VariableDisplayType]
VARIABLE_DISPLAY_TYPE_UNSPECIFIED: VariableDisplayType
VARIABLE_DISPLAY_TYPE_LABEL_VALUE: VariableDisplayType
VARIABLE_DISPLAY_TYPE_VALUE: VariableDisplayType
VARIABLE_DISPLAY_TYPE_NOTHING: VariableDisplayType

class Variable(_message.Message):
    __slots__ = ("name", "definition", "display_name", "description", "display_type")
    class Definition(_message.Message):
        __slots__ = ("constant", "multi_select")
        CONSTANT_FIELD_NUMBER: _ClassVar[int]
        MULTI_SELECT_FIELD_NUMBER: _ClassVar[int]
        constant: Constant
        multi_select: MultiSelect
        def __init__(self, constant: _Optional[_Union[Constant, _Mapping]] = ..., multi_select: _Optional[_Union[MultiSelect, _Mapping]] = ...) -> None: ...
    NAME_FIELD_NUMBER: _ClassVar[int]
    DEFINITION_FIELD_NUMBER: _ClassVar[int]
    DISPLAY_NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    DISPLAY_TYPE_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    definition: Variable.Definition
    display_name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    display_type: VariableDisplayType
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., definition: _Optional[_Union[Variable.Definition, _Mapping]] = ..., display_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., display_type: _Optional[_Union[VariableDisplayType, str]] = ...) -> None: ...

class Constant(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: _wrappers_pb2.StringValue
    def __init__(self, value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class MultiSelect(_message.Message):
    __slots__ = ("selected", "source", "selection", "values_order_direction", "selection_options")
    class RefreshStrategy(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        REFRESH_STRATEGY_UNSPECIFIED: _ClassVar[MultiSelect.RefreshStrategy]
        REFRESH_STRATEGY_ON_DASHBOARD_LOAD: _ClassVar[MultiSelect.RefreshStrategy]
        REFRESH_STRATEGY_ON_TIME_FRAME_CHANGE: _ClassVar[MultiSelect.RefreshStrategy]
    REFRESH_STRATEGY_UNSPECIFIED: MultiSelect.RefreshStrategy
    REFRESH_STRATEGY_ON_DASHBOARD_LOAD: MultiSelect.RefreshStrategy
    REFRESH_STRATEGY_ON_TIME_FRAME_CHANGE: MultiSelect.RefreshStrategy
    class Source(_message.Message):
        __slots__ = ("logs_path", "metric_label", "constant_list", "span_field", "query")
        LOGS_PATH_FIELD_NUMBER: _ClassVar[int]
        METRIC_LABEL_FIELD_NUMBER: _ClassVar[int]
        CONSTANT_LIST_FIELD_NUMBER: _ClassVar[int]
        SPAN_FIELD_FIELD_NUMBER: _ClassVar[int]
        QUERY_FIELD_NUMBER: _ClassVar[int]
        logs_path: MultiSelect.LogsPathSource
        metric_label: MultiSelect.MetricLabelSource
        constant_list: MultiSelect.ConstantListSource
        span_field: MultiSelect.SpanFieldSource
        query: MultiSelect.QuerySource
        def __init__(self, logs_path: _Optional[_Union[MultiSelect.LogsPathSource, _Mapping]] = ..., metric_label: _Optional[_Union[MultiSelect.MetricLabelSource, _Mapping]] = ..., constant_list: _Optional[_Union[MultiSelect.ConstantListSource, _Mapping]] = ..., span_field: _Optional[_Union[MultiSelect.SpanFieldSource, _Mapping]] = ..., query: _Optional[_Union[MultiSelect.QuerySource, _Mapping]] = ...) -> None: ...
    class LogsPathSource(_message.Message):
        __slots__ = ("value", "observation_field")
        VALUE_FIELD_NUMBER: _ClassVar[int]
        OBSERVATION_FIELD_FIELD_NUMBER: _ClassVar[int]
        value: _wrappers_pb2.StringValue
        observation_field: _observation_field_pb2.ObservationField
        def __init__(self, value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., observation_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
    class MetricLabelSource(_message.Message):
        __slots__ = ("metric_name", "label")
        METRIC_NAME_FIELD_NUMBER: _ClassVar[int]
        LABEL_FIELD_NUMBER: _ClassVar[int]
        metric_name: _wrappers_pb2.StringValue
        label: _wrappers_pb2.StringValue
        def __init__(self, metric_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., label: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class SpanFieldSource(_message.Message):
        __slots__ = ("value",)
        VALUE_FIELD_NUMBER: _ClassVar[int]
        value: _span_field_pb2.SpanField
        def __init__(self, value: _Optional[_Union[_span_field_pb2.SpanField, _Mapping]] = ...) -> None: ...
    class ConstantListSource(_message.Message):
        __slots__ = ("values",)
        VALUES_FIELD_NUMBER: _ClassVar[int]
        values: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        def __init__(self, values: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...
    class QuerySource(_message.Message):
        __slots__ = ("query", "refresh_strategy", "value_display_options")
        QUERY_FIELD_NUMBER: _ClassVar[int]
        REFRESH_STRATEGY_FIELD_NUMBER: _ClassVar[int]
        VALUE_DISPLAY_OPTIONS_FIELD_NUMBER: _ClassVar[int]
        query: MultiSelect.Query
        refresh_strategy: MultiSelect.RefreshStrategy
        value_display_options: MultiSelect.ValueDisplayOptions
        def __init__(self, query: _Optional[_Union[MultiSelect.Query, _Mapping]] = ..., refresh_strategy: _Optional[_Union[MultiSelect.RefreshStrategy, str]] = ..., value_display_options: _Optional[_Union[MultiSelect.ValueDisplayOptions, _Mapping]] = ...) -> None: ...
    class Query(_message.Message):
        __slots__ = ("logs_query", "metrics_query", "spans_query")
        class LogsQuery(_message.Message):
            __slots__ = ("type",)
            class Type(_message.Message):
                __slots__ = ("field_name", "field_value")
                class FieldName(_message.Message):
                    __slots__ = ("log_regex",)
                    LOG_REGEX_FIELD_NUMBER: _ClassVar[int]
                    log_regex: _wrappers_pb2.StringValue
                    def __init__(self, log_regex: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
                class FieldValue(_message.Message):
                    __slots__ = ("observation_field",)
                    OBSERVATION_FIELD_FIELD_NUMBER: _ClassVar[int]
                    observation_field: _observation_field_pb2.ObservationField
                    def __init__(self, observation_field: _Optional[_Union[_observation_field_pb2.ObservationField, _Mapping]] = ...) -> None: ...
                FIELD_NAME_FIELD_NUMBER: _ClassVar[int]
                FIELD_VALUE_FIELD_NUMBER: _ClassVar[int]
                field_name: MultiSelect.Query.LogsQuery.Type.FieldName
                field_value: MultiSelect.Query.LogsQuery.Type.FieldValue
                def __init__(self, field_name: _Optional[_Union[MultiSelect.Query.LogsQuery.Type.FieldName, _Mapping]] = ..., field_value: _Optional[_Union[MultiSelect.Query.LogsQuery.Type.FieldValue, _Mapping]] = ...) -> None: ...
            TYPE_FIELD_NUMBER: _ClassVar[int]
            type: MultiSelect.Query.LogsQuery.Type
            def __init__(self, type: _Optional[_Union[MultiSelect.Query.LogsQuery.Type, _Mapping]] = ...) -> None: ...
        class SpansQuery(_message.Message):
            __slots__ = ("type",)
            class Type(_message.Message):
                __slots__ = ("field_name", "field_value")
                class FieldName(_message.Message):
                    __slots__ = ("span_regex",)
                    SPAN_REGEX_FIELD_NUMBER: _ClassVar[int]
                    span_regex: _wrappers_pb2.StringValue
                    def __init__(self, span_regex: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
                class FieldValue(_message.Message):
                    __slots__ = ("value",)
                    VALUE_FIELD_NUMBER: _ClassVar[int]
                    value: _span_field_pb2.SpanField
                    def __init__(self, value: _Optional[_Union[_span_field_pb2.SpanField, _Mapping]] = ...) -> None: ...
                FIELD_NAME_FIELD_NUMBER: _ClassVar[int]
                FIELD_VALUE_FIELD_NUMBER: _ClassVar[int]
                field_name: MultiSelect.Query.SpansQuery.Type.FieldName
                field_value: MultiSelect.Query.SpansQuery.Type.FieldValue
                def __init__(self, field_name: _Optional[_Union[MultiSelect.Query.SpansQuery.Type.FieldName, _Mapping]] = ..., field_value: _Optional[_Union[MultiSelect.Query.SpansQuery.Type.FieldValue, _Mapping]] = ...) -> None: ...
            TYPE_FIELD_NUMBER: _ClassVar[int]
            type: MultiSelect.Query.SpansQuery.Type
            def __init__(self, type: _Optional[_Union[MultiSelect.Query.SpansQuery.Type, _Mapping]] = ...) -> None: ...
        class MetricsQuery(_message.Message):
            __slots__ = ("type",)
            class Type(_message.Message):
                __slots__ = ("metric_name", "label_name", "label_value")
                class MetricName(_message.Message):
                    __slots__ = ("metric_regex",)
                    METRIC_REGEX_FIELD_NUMBER: _ClassVar[int]
                    metric_regex: _wrappers_pb2.StringValue
                    def __init__(self, metric_regex: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
                class LabelName(_message.Message):
                    __slots__ = ("metric_regex",)
                    METRIC_REGEX_FIELD_NUMBER: _ClassVar[int]
                    metric_regex: _wrappers_pb2.StringValue
                    def __init__(self, metric_regex: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
                class LabelValue(_message.Message):
                    __slots__ = ("metric_name", "label_name", "label_filters")
                    METRIC_NAME_FIELD_NUMBER: _ClassVar[int]
                    LABEL_NAME_FIELD_NUMBER: _ClassVar[int]
                    LABEL_FILTERS_FIELD_NUMBER: _ClassVar[int]
                    metric_name: MultiSelect.Query.MetricsQuery.StringOrVariable
                    label_name: MultiSelect.Query.MetricsQuery.StringOrVariable
                    label_filters: _containers.RepeatedCompositeFieldContainer[MultiSelect.Query.MetricsQuery.MetricsLabelFilter]
                    def __init__(self, metric_name: _Optional[_Union[MultiSelect.Query.MetricsQuery.StringOrVariable, _Mapping]] = ..., label_name: _Optional[_Union[MultiSelect.Query.MetricsQuery.StringOrVariable, _Mapping]] = ..., label_filters: _Optional[_Iterable[_Union[MultiSelect.Query.MetricsQuery.MetricsLabelFilter, _Mapping]]] = ...) -> None: ...
                METRIC_NAME_FIELD_NUMBER: _ClassVar[int]
                LABEL_NAME_FIELD_NUMBER: _ClassVar[int]
                LABEL_VALUE_FIELD_NUMBER: _ClassVar[int]
                metric_name: MultiSelect.Query.MetricsQuery.Type.MetricName
                label_name: MultiSelect.Query.MetricsQuery.Type.LabelName
                label_value: MultiSelect.Query.MetricsQuery.Type.LabelValue
                def __init__(self, metric_name: _Optional[_Union[MultiSelect.Query.MetricsQuery.Type.MetricName, _Mapping]] = ..., label_name: _Optional[_Union[MultiSelect.Query.MetricsQuery.Type.LabelName, _Mapping]] = ..., label_value: _Optional[_Union[MultiSelect.Query.MetricsQuery.Type.LabelValue, _Mapping]] = ...) -> None: ...
            class StringOrVariable(_message.Message):
                __slots__ = ("string_value", "variable_name")
                STRING_VALUE_FIELD_NUMBER: _ClassVar[int]
                VARIABLE_NAME_FIELD_NUMBER: _ClassVar[int]
                string_value: _wrappers_pb2.StringValue
                variable_name: _wrappers_pb2.StringValue
                def __init__(self, string_value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., variable_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
            class MetricsLabelFilter(_message.Message):
                __slots__ = ("metric", "label", "operator")
                METRIC_FIELD_NUMBER: _ClassVar[int]
                LABEL_FIELD_NUMBER: _ClassVar[int]
                OPERATOR_FIELD_NUMBER: _ClassVar[int]
                metric: MultiSelect.Query.MetricsQuery.StringOrVariable
                label: MultiSelect.Query.MetricsQuery.StringOrVariable
                operator: MultiSelect.Query.MetricsQuery.Operator
                def __init__(self, metric: _Optional[_Union[MultiSelect.Query.MetricsQuery.StringOrVariable, _Mapping]] = ..., label: _Optional[_Union[MultiSelect.Query.MetricsQuery.StringOrVariable, _Mapping]] = ..., operator: _Optional[_Union[MultiSelect.Query.MetricsQuery.Operator, _Mapping]] = ...) -> None: ...
            class Operator(_message.Message):
                __slots__ = ("equals", "not_equals")
                EQUALS_FIELD_NUMBER: _ClassVar[int]
                NOT_EQUALS_FIELD_NUMBER: _ClassVar[int]
                equals: MultiSelect.Query.MetricsQuery.Equals
                not_equals: MultiSelect.Query.MetricsQuery.NotEquals
                def __init__(self, equals: _Optional[_Union[MultiSelect.Query.MetricsQuery.Equals, _Mapping]] = ..., not_equals: _Optional[_Union[MultiSelect.Query.MetricsQuery.NotEquals, _Mapping]] = ...) -> None: ...
            class Equals(_message.Message):
                __slots__ = ("selection",)
                SELECTION_FIELD_NUMBER: _ClassVar[int]
                selection: MultiSelect.Query.MetricsQuery.Selection
                def __init__(self, selection: _Optional[_Union[MultiSelect.Query.MetricsQuery.Selection, _Mapping]] = ...) -> None: ...
            class NotEquals(_message.Message):
                __slots__ = ("selection",)
                SELECTION_FIELD_NUMBER: _ClassVar[int]
                selection: MultiSelect.Query.MetricsQuery.Selection
                def __init__(self, selection: _Optional[_Union[MultiSelect.Query.MetricsQuery.Selection, _Mapping]] = ...) -> None: ...
            class Selection(_message.Message):
                __slots__ = ("list",)
                class ListSelection(_message.Message):
                    __slots__ = ("values",)
                    VALUES_FIELD_NUMBER: _ClassVar[int]
                    values: _containers.RepeatedCompositeFieldContainer[MultiSelect.Query.MetricsQuery.StringOrVariable]
                    def __init__(self, values: _Optional[_Iterable[_Union[MultiSelect.Query.MetricsQuery.StringOrVariable, _Mapping]]] = ...) -> None: ...
                LIST_FIELD_NUMBER: _ClassVar[int]
                list: MultiSelect.Query.MetricsQuery.Selection.ListSelection
                def __init__(self, list: _Optional[_Union[MultiSelect.Query.MetricsQuery.Selection.ListSelection, _Mapping]] = ...) -> None: ...
            TYPE_FIELD_NUMBER: _ClassVar[int]
            type: MultiSelect.Query.MetricsQuery.Type
            def __init__(self, type: _Optional[_Union[MultiSelect.Query.MetricsQuery.Type, _Mapping]] = ...) -> None: ...
        LOGS_QUERY_FIELD_NUMBER: _ClassVar[int]
        METRICS_QUERY_FIELD_NUMBER: _ClassVar[int]
        SPANS_QUERY_FIELD_NUMBER: _ClassVar[int]
        logs_query: MultiSelect.Query.LogsQuery
        metrics_query: MultiSelect.Query.MetricsQuery
        spans_query: MultiSelect.Query.SpansQuery
        def __init__(self, logs_query: _Optional[_Union[MultiSelect.Query.LogsQuery, _Mapping]] = ..., metrics_query: _Optional[_Union[MultiSelect.Query.MetricsQuery, _Mapping]] = ..., spans_query: _Optional[_Union[MultiSelect.Query.SpansQuery, _Mapping]] = ...) -> None: ...
    class ValueDisplayOptions(_message.Message):
        __slots__ = ("value_regex", "label_regex")
        VALUE_REGEX_FIELD_NUMBER: _ClassVar[int]
        LABEL_REGEX_FIELD_NUMBER: _ClassVar[int]
        value_regex: _wrappers_pb2.StringValue
        label_regex: _wrappers_pb2.StringValue
        def __init__(self, value_regex: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., label_regex: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class Selection(_message.Message):
        __slots__ = ("all", "list")
        class AllSelection(_message.Message):
            __slots__ = ()
            def __init__(self) -> None: ...
        class ListSelection(_message.Message):
            __slots__ = ("values", "labels")
            VALUES_FIELD_NUMBER: _ClassVar[int]
            LABELS_FIELD_NUMBER: _ClassVar[int]
            values: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
            labels: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
            def __init__(self, values: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., labels: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...
        ALL_FIELD_NUMBER: _ClassVar[int]
        LIST_FIELD_NUMBER: _ClassVar[int]
        all: MultiSelect.Selection.AllSelection
        list: MultiSelect.Selection.ListSelection
        def __init__(self, all: _Optional[_Union[MultiSelect.Selection.AllSelection, _Mapping]] = ..., list: _Optional[_Union[MultiSelect.Selection.ListSelection, _Mapping]] = ...) -> None: ...
    class VariableSelectionOptions(_message.Message):
        __slots__ = ("selection_type",)
        class SelectionType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            SELECTION_TYPE_UNSPECIFIED: _ClassVar[MultiSelect.VariableSelectionOptions.SelectionType]
            SELECTION_TYPE_MULTI_ALL: _ClassVar[MultiSelect.VariableSelectionOptions.SelectionType]
            SELECTION_TYPE_MULTI: _ClassVar[MultiSelect.VariableSelectionOptions.SelectionType]
            SELECTION_TYPE_SINGLE: _ClassVar[MultiSelect.VariableSelectionOptions.SelectionType]
        SELECTION_TYPE_UNSPECIFIED: MultiSelect.VariableSelectionOptions.SelectionType
        SELECTION_TYPE_MULTI_ALL: MultiSelect.VariableSelectionOptions.SelectionType
        SELECTION_TYPE_MULTI: MultiSelect.VariableSelectionOptions.SelectionType
        SELECTION_TYPE_SINGLE: MultiSelect.VariableSelectionOptions.SelectionType
        SELECTION_TYPE_FIELD_NUMBER: _ClassVar[int]
        selection_type: MultiSelect.VariableSelectionOptions.SelectionType
        def __init__(self, selection_type: _Optional[_Union[MultiSelect.VariableSelectionOptions.SelectionType, str]] = ...) -> None: ...
    SELECTED_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    SELECTION_FIELD_NUMBER: _ClassVar[int]
    VALUES_ORDER_DIRECTION_FIELD_NUMBER: _ClassVar[int]
    SELECTION_OPTIONS_FIELD_NUMBER: _ClassVar[int]
    selected: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    source: MultiSelect.Source
    selection: MultiSelect.Selection
    values_order_direction: _order_direction_pb2.OrderDirection
    selection_options: MultiSelect.VariableSelectionOptions
    def __init__(self, selected: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., source: _Optional[_Union[MultiSelect.Source, _Mapping]] = ..., selection: _Optional[_Union[MultiSelect.Selection, _Mapping]] = ..., values_order_direction: _Optional[_Union[_order_direction_pb2.OrderDirection, str]] = ..., selection_options: _Optional[_Union[MultiSelect.VariableSelectionOptions, _Mapping]] = ...) -> None: ...
