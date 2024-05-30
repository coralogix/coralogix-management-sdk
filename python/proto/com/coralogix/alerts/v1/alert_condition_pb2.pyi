from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.alerts.v1 import flow_alert_pb2 as _flow_alert_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Timeframe(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    TIMEFRAME_5_MIN_OR_UNSPECIFIED: _ClassVar[Timeframe]
    TIMEFRAME_1_MIN: _ClassVar[Timeframe]
    TIMEFRAME_2_MIN: _ClassVar[Timeframe]
    TIMEFRAME_10_MIN: _ClassVar[Timeframe]
    TIMEFRAME_15_MIN: _ClassVar[Timeframe]
    TIMEFRAME_20_MIN: _ClassVar[Timeframe]
    TIMEFRAME_30_MIN: _ClassVar[Timeframe]
    TIMEFRAME_1_H: _ClassVar[Timeframe]
    TIMEFRAME_2_H: _ClassVar[Timeframe]
    TIMEFRAME_3_H: _ClassVar[Timeframe]
    TIMEFRAME_4_H: _ClassVar[Timeframe]
    TIMEFRAME_6_H: _ClassVar[Timeframe]
    TIMEFRAME_12_H: _ClassVar[Timeframe]
    TIMEFRAME_24_H: _ClassVar[Timeframe]
    TIMEFRAME_36_H: _ClassVar[Timeframe]
    TIMEFRAME_48_H: _ClassVar[Timeframe]
    TIMEFRAME_72_H: _ClassVar[Timeframe]
    TIMEFRAME_1_W: _ClassVar[Timeframe]
    TIMEFRAME_1_M: _ClassVar[Timeframe]
    TIMEFRAME_2_M: _ClassVar[Timeframe]
    TIMEFRAME_3_M: _ClassVar[Timeframe]

class RelativeTimeframe(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    RELATIVE_TIMEFRAME_HOUR_OR_UNSPECIFIED: _ClassVar[RelativeTimeframe]
    RELATIVE_TIMEFRAME_DAY: _ClassVar[RelativeTimeframe]
    RELATIVE_TIMEFRAME_WEEK: _ClassVar[RelativeTimeframe]
    RELATIVE_TIMEFRAME_MONTH: _ClassVar[RelativeTimeframe]

class CleanupDeadmanDuration(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    CLEANUP_DEADMAN_DURATION_NEVER_OR_UNSPECIFIED: _ClassVar[CleanupDeadmanDuration]
    CLEANUP_DEADMAN_DURATION_5MIN: _ClassVar[CleanupDeadmanDuration]
    CLEANUP_DEADMAN_DURATION_10MIN: _ClassVar[CleanupDeadmanDuration]
    CLEANUP_DEADMAN_DURATION_1H: _ClassVar[CleanupDeadmanDuration]
    CLEANUP_DEADMAN_DURATION_2H: _ClassVar[CleanupDeadmanDuration]
    CLEANUP_DEADMAN_DURATION_6H: _ClassVar[CleanupDeadmanDuration]
    CLEANUP_DEADMAN_DURATION_12H: _ClassVar[CleanupDeadmanDuration]
    CLEANUP_DEADMAN_DURATION_24H: _ClassVar[CleanupDeadmanDuration]

class EvaluationWindow(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    EVALUATION_WINDOW_ROLLING_OR_UNSPECIFIED: _ClassVar[EvaluationWindow]
    EVALUATION_WINDOW_DYNAMIC: _ClassVar[EvaluationWindow]
TIMEFRAME_5_MIN_OR_UNSPECIFIED: Timeframe
TIMEFRAME_1_MIN: Timeframe
TIMEFRAME_2_MIN: Timeframe
TIMEFRAME_10_MIN: Timeframe
TIMEFRAME_15_MIN: Timeframe
TIMEFRAME_20_MIN: Timeframe
TIMEFRAME_30_MIN: Timeframe
TIMEFRAME_1_H: Timeframe
TIMEFRAME_2_H: Timeframe
TIMEFRAME_3_H: Timeframe
TIMEFRAME_4_H: Timeframe
TIMEFRAME_6_H: Timeframe
TIMEFRAME_12_H: Timeframe
TIMEFRAME_24_H: Timeframe
TIMEFRAME_36_H: Timeframe
TIMEFRAME_48_H: Timeframe
TIMEFRAME_72_H: Timeframe
TIMEFRAME_1_W: Timeframe
TIMEFRAME_1_M: Timeframe
TIMEFRAME_2_M: Timeframe
TIMEFRAME_3_M: Timeframe
RELATIVE_TIMEFRAME_HOUR_OR_UNSPECIFIED: RelativeTimeframe
RELATIVE_TIMEFRAME_DAY: RelativeTimeframe
RELATIVE_TIMEFRAME_WEEK: RelativeTimeframe
RELATIVE_TIMEFRAME_MONTH: RelativeTimeframe
CLEANUP_DEADMAN_DURATION_NEVER_OR_UNSPECIFIED: CleanupDeadmanDuration
CLEANUP_DEADMAN_DURATION_5MIN: CleanupDeadmanDuration
CLEANUP_DEADMAN_DURATION_10MIN: CleanupDeadmanDuration
CLEANUP_DEADMAN_DURATION_1H: CleanupDeadmanDuration
CLEANUP_DEADMAN_DURATION_2H: CleanupDeadmanDuration
CLEANUP_DEADMAN_DURATION_6H: CleanupDeadmanDuration
CLEANUP_DEADMAN_DURATION_12H: CleanupDeadmanDuration
CLEANUP_DEADMAN_DURATION_24H: CleanupDeadmanDuration
EVALUATION_WINDOW_ROLLING_OR_UNSPECIFIED: EvaluationWindow
EVALUATION_WINDOW_DYNAMIC: EvaluationWindow

class AlertCondition(_message.Message):
    __slots__ = ("immediate", "less_than", "more_than", "more_than_usual", "new_value", "flow", "unique_count", "less_than_usual", "more_than_or_equal", "less_than_or_equal")
    IMMEDIATE_FIELD_NUMBER: _ClassVar[int]
    LESS_THAN_FIELD_NUMBER: _ClassVar[int]
    MORE_THAN_FIELD_NUMBER: _ClassVar[int]
    MORE_THAN_USUAL_FIELD_NUMBER: _ClassVar[int]
    NEW_VALUE_FIELD_NUMBER: _ClassVar[int]
    FLOW_FIELD_NUMBER: _ClassVar[int]
    UNIQUE_COUNT_FIELD_NUMBER: _ClassVar[int]
    LESS_THAN_USUAL_FIELD_NUMBER: _ClassVar[int]
    MORE_THAN_OR_EQUAL_FIELD_NUMBER: _ClassVar[int]
    LESS_THAN_OR_EQUAL_FIELD_NUMBER: _ClassVar[int]
    immediate: ImmediateCondition
    less_than: LessThanCondition
    more_than: MoreThanCondition
    more_than_usual: MoreThanUsualCondition
    new_value: NewValueCondition
    flow: FlowCondition
    unique_count: UniqueCountCondition
    less_than_usual: LessThanUsualCondition
    more_than_or_equal: MoreThanOrEqualCondition
    less_than_or_equal: LessThanOrEqualCondition
    def __init__(self, immediate: _Optional[_Union[ImmediateCondition, _Mapping]] = ..., less_than: _Optional[_Union[LessThanCondition, _Mapping]] = ..., more_than: _Optional[_Union[MoreThanCondition, _Mapping]] = ..., more_than_usual: _Optional[_Union[MoreThanUsualCondition, _Mapping]] = ..., new_value: _Optional[_Union[NewValueCondition, _Mapping]] = ..., flow: _Optional[_Union[FlowCondition, _Mapping]] = ..., unique_count: _Optional[_Union[UniqueCountCondition, _Mapping]] = ..., less_than_usual: _Optional[_Union[LessThanUsualCondition, _Mapping]] = ..., more_than_or_equal: _Optional[_Union[MoreThanOrEqualCondition, _Mapping]] = ..., less_than_or_equal: _Optional[_Union[LessThanOrEqualCondition, _Mapping]] = ...) -> None: ...

class ImmediateCondition(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class LessThanCondition(_message.Message):
    __slots__ = ("parameters",)
    PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    parameters: ConditionParameters
    def __init__(self, parameters: _Optional[_Union[ConditionParameters, _Mapping]] = ...) -> None: ...

class LessThanOrEqualCondition(_message.Message):
    __slots__ = ("parameters",)
    PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    parameters: ConditionParameters
    def __init__(self, parameters: _Optional[_Union[ConditionParameters, _Mapping]] = ...) -> None: ...

class MoreThanCondition(_message.Message):
    __slots__ = ("parameters", "evaluation_window")
    PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    EVALUATION_WINDOW_FIELD_NUMBER: _ClassVar[int]
    parameters: ConditionParameters
    evaluation_window: EvaluationWindow
    def __init__(self, parameters: _Optional[_Union[ConditionParameters, _Mapping]] = ..., evaluation_window: _Optional[_Union[EvaluationWindow, str]] = ...) -> None: ...

class MoreThanOrEqualCondition(_message.Message):
    __slots__ = ("parameters", "evaluation_window")
    PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    EVALUATION_WINDOW_FIELD_NUMBER: _ClassVar[int]
    parameters: ConditionParameters
    evaluation_window: EvaluationWindow
    def __init__(self, parameters: _Optional[_Union[ConditionParameters, _Mapping]] = ..., evaluation_window: _Optional[_Union[EvaluationWindow, str]] = ...) -> None: ...

class MoreThanUsualCondition(_message.Message):
    __slots__ = ("parameters",)
    PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    parameters: ConditionParameters
    def __init__(self, parameters: _Optional[_Union[ConditionParameters, _Mapping]] = ...) -> None: ...

class LessThanUsualCondition(_message.Message):
    __slots__ = ("parameters",)
    PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    parameters: ConditionParameters
    def __init__(self, parameters: _Optional[_Union[ConditionParameters, _Mapping]] = ...) -> None: ...

class NewValueCondition(_message.Message):
    __slots__ = ("parameters",)
    PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    parameters: ConditionParameters
    def __init__(self, parameters: _Optional[_Union[ConditionParameters, _Mapping]] = ...) -> None: ...

class UniqueCountCondition(_message.Message):
    __slots__ = ("parameters",)
    PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    parameters: ConditionParameters
    def __init__(self, parameters: _Optional[_Union[ConditionParameters, _Mapping]] = ...) -> None: ...

class FlowCondition(_message.Message):
    __slots__ = ("stages", "parameters", "enforce_suppression")
    STAGES_FIELD_NUMBER: _ClassVar[int]
    PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    ENFORCE_SUPPRESSION_FIELD_NUMBER: _ClassVar[int]
    stages: _containers.RepeatedCompositeFieldContainer[_flow_alert_pb2.FlowStage]
    parameters: ConditionParameters
    enforce_suppression: _wrappers_pb2.BoolValue
    def __init__(self, stages: _Optional[_Iterable[_Union[_flow_alert_pb2.FlowStage, _Mapping]]] = ..., parameters: _Optional[_Union[ConditionParameters, _Mapping]] = ..., enforce_suppression: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class ConditionParameters(_message.Message):
    __slots__ = ("threshold", "timeframe", "group_by", "metric_alert_parameters", "metric_alert_promql_parameters", "notify_on_resolved", "ignore_infinity", "relative_timeframe", "notify_group_by_only_alerts", "notify_per_group_by_value", "cardinality_fields", "max_unique_count_values_for_group_by_key", "related_extended_data")
    THRESHOLD_FIELD_NUMBER: _ClassVar[int]
    TIMEFRAME_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_FIELD_NUMBER: _ClassVar[int]
    METRIC_ALERT_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    METRIC_ALERT_PROMQL_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    NOTIFY_ON_RESOLVED_FIELD_NUMBER: _ClassVar[int]
    IGNORE_INFINITY_FIELD_NUMBER: _ClassVar[int]
    RELATIVE_TIMEFRAME_FIELD_NUMBER: _ClassVar[int]
    NOTIFY_GROUP_BY_ONLY_ALERTS_FIELD_NUMBER: _ClassVar[int]
    NOTIFY_PER_GROUP_BY_VALUE_FIELD_NUMBER: _ClassVar[int]
    CARDINALITY_FIELDS_FIELD_NUMBER: _ClassVar[int]
    MAX_UNIQUE_COUNT_VALUES_FOR_GROUP_BY_KEY_FIELD_NUMBER: _ClassVar[int]
    RELATED_EXTENDED_DATA_FIELD_NUMBER: _ClassVar[int]
    threshold: _wrappers_pb2.DoubleValue
    timeframe: Timeframe
    group_by: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    metric_alert_parameters: MetricAlertConditionParameters
    metric_alert_promql_parameters: MetricAlertPromqlConditionParameters
    notify_on_resolved: _wrappers_pb2.BoolValue
    ignore_infinity: _wrappers_pb2.BoolValue
    relative_timeframe: RelativeTimeframe
    notify_group_by_only_alerts: _wrappers_pb2.BoolValue
    notify_per_group_by_value: _wrappers_pb2.BoolValue
    cardinality_fields: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    max_unique_count_values_for_group_by_key: _wrappers_pb2.UInt32Value
    related_extended_data: RelatedExtendedData
    def __init__(self, threshold: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ..., timeframe: _Optional[_Union[Timeframe, str]] = ..., group_by: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., metric_alert_parameters: _Optional[_Union[MetricAlertConditionParameters, _Mapping]] = ..., metric_alert_promql_parameters: _Optional[_Union[MetricAlertPromqlConditionParameters, _Mapping]] = ..., notify_on_resolved: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., ignore_infinity: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., relative_timeframe: _Optional[_Union[RelativeTimeframe, str]] = ..., notify_group_by_only_alerts: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., notify_per_group_by_value: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., cardinality_fields: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., max_unique_count_values_for_group_by_key: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., related_extended_data: _Optional[_Union[RelatedExtendedData, _Mapping]] = ...) -> None: ...

class MetricAlertConditionParameters(_message.Message):
    __slots__ = ("metric_field", "metric_source", "arithmetic_operator", "arithmetic_operator_modifier", "sample_threshold_percentage", "non_null_percentage", "swap_null_values")
    class MetricSource(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        METRIC_SOURCE_LOGS2METRICS_OR_UNSPECIFIED: _ClassVar[MetricAlertConditionParameters.MetricSource]
        METRIC_SOURCE_PROMETHEUS: _ClassVar[MetricAlertConditionParameters.MetricSource]
    METRIC_SOURCE_LOGS2METRICS_OR_UNSPECIFIED: MetricAlertConditionParameters.MetricSource
    METRIC_SOURCE_PROMETHEUS: MetricAlertConditionParameters.MetricSource
    class ArithmeticOperator(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        ARITHMETIC_OPERATOR_AVG_OR_UNSPECIFIED: _ClassVar[MetricAlertConditionParameters.ArithmeticOperator]
        ARITHMETIC_OPERATOR_MIN: _ClassVar[MetricAlertConditionParameters.ArithmeticOperator]
        ARITHMETIC_OPERATOR_MAX: _ClassVar[MetricAlertConditionParameters.ArithmeticOperator]
        ARITHMETIC_OPERATOR_SUM: _ClassVar[MetricAlertConditionParameters.ArithmeticOperator]
        ARITHMETIC_OPERATOR_COUNT: _ClassVar[MetricAlertConditionParameters.ArithmeticOperator]
        ARITHMETIC_OPERATOR_PERCENTILE: _ClassVar[MetricAlertConditionParameters.ArithmeticOperator]
    ARITHMETIC_OPERATOR_AVG_OR_UNSPECIFIED: MetricAlertConditionParameters.ArithmeticOperator
    ARITHMETIC_OPERATOR_MIN: MetricAlertConditionParameters.ArithmeticOperator
    ARITHMETIC_OPERATOR_MAX: MetricAlertConditionParameters.ArithmeticOperator
    ARITHMETIC_OPERATOR_SUM: MetricAlertConditionParameters.ArithmeticOperator
    ARITHMETIC_OPERATOR_COUNT: MetricAlertConditionParameters.ArithmeticOperator
    ARITHMETIC_OPERATOR_PERCENTILE: MetricAlertConditionParameters.ArithmeticOperator
    METRIC_FIELD_FIELD_NUMBER: _ClassVar[int]
    METRIC_SOURCE_FIELD_NUMBER: _ClassVar[int]
    ARITHMETIC_OPERATOR_FIELD_NUMBER: _ClassVar[int]
    ARITHMETIC_OPERATOR_MODIFIER_FIELD_NUMBER: _ClassVar[int]
    SAMPLE_THRESHOLD_PERCENTAGE_FIELD_NUMBER: _ClassVar[int]
    NON_NULL_PERCENTAGE_FIELD_NUMBER: _ClassVar[int]
    SWAP_NULL_VALUES_FIELD_NUMBER: _ClassVar[int]
    metric_field: _wrappers_pb2.StringValue
    metric_source: MetricAlertConditionParameters.MetricSource
    arithmetic_operator: MetricAlertConditionParameters.ArithmeticOperator
    arithmetic_operator_modifier: _wrappers_pb2.UInt32Value
    sample_threshold_percentage: _wrappers_pb2.UInt32Value
    non_null_percentage: _wrappers_pb2.UInt32Value
    swap_null_values: _wrappers_pb2.BoolValue
    def __init__(self, metric_field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., metric_source: _Optional[_Union[MetricAlertConditionParameters.MetricSource, str]] = ..., arithmetic_operator: _Optional[_Union[MetricAlertConditionParameters.ArithmeticOperator, str]] = ..., arithmetic_operator_modifier: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., sample_threshold_percentage: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., non_null_percentage: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., swap_null_values: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class MetricAlertPromqlConditionParameters(_message.Message):
    __slots__ = ("promql_text", "arithmetic_operator_modifier", "sample_threshold_percentage", "non_null_percentage", "swap_null_values")
    PROMQL_TEXT_FIELD_NUMBER: _ClassVar[int]
    ARITHMETIC_OPERATOR_MODIFIER_FIELD_NUMBER: _ClassVar[int]
    SAMPLE_THRESHOLD_PERCENTAGE_FIELD_NUMBER: _ClassVar[int]
    NON_NULL_PERCENTAGE_FIELD_NUMBER: _ClassVar[int]
    SWAP_NULL_VALUES_FIELD_NUMBER: _ClassVar[int]
    promql_text: _wrappers_pb2.StringValue
    arithmetic_operator_modifier: _wrappers_pb2.UInt32Value
    sample_threshold_percentage: _wrappers_pb2.UInt32Value
    non_null_percentage: _wrappers_pb2.UInt32Value
    swap_null_values: _wrappers_pb2.BoolValue
    def __init__(self, promql_text: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., arithmetic_operator_modifier: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., sample_threshold_percentage: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., non_null_percentage: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., swap_null_values: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class RelatedExtendedData(_message.Message):
    __slots__ = ("cleanup_deadman_duration", "should_trigger_deadman")
    CLEANUP_DEADMAN_DURATION_FIELD_NUMBER: _ClassVar[int]
    SHOULD_TRIGGER_DEADMAN_FIELD_NUMBER: _ClassVar[int]
    cleanup_deadman_duration: CleanupDeadmanDuration
    should_trigger_deadman: _wrappers_pb2.BoolValue
    def __init__(self, cleanup_deadman_duration: _Optional[_Union[CleanupDeadmanDuration, str]] = ..., should_trigger_deadman: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
