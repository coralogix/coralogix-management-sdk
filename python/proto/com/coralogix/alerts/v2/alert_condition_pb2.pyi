from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.alerts.v1 import flow_alert_pb2 as _flow_alert_pb2
from com.coralogix.alerts.v1 import alert_condition_pb2 as _alert_condition_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

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
    evaluation_window: _alert_condition_pb2.EvaluationWindow
    def __init__(self, parameters: _Optional[_Union[ConditionParameters, _Mapping]] = ..., evaluation_window: _Optional[_Union[_alert_condition_pb2.EvaluationWindow, str]] = ...) -> None: ...

class MoreThanOrEqualCondition(_message.Message):
    __slots__ = ("parameters", "evaluation_window")
    PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    EVALUATION_WINDOW_FIELD_NUMBER: _ClassVar[int]
    parameters: ConditionParameters
    evaluation_window: _alert_condition_pb2.EvaluationWindow
    def __init__(self, parameters: _Optional[_Union[ConditionParameters, _Mapping]] = ..., evaluation_window: _Optional[_Union[_alert_condition_pb2.EvaluationWindow, str]] = ...) -> None: ...

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
    __slots__ = ("threshold", "timeframe", "group_by", "metric_alert_parameters", "metric_alert_promql_parameters", "ignore_infinity", "relative_timeframe", "notify_group_by_only_alerts", "cardinality_fields", "max_unique_count_values_for_group_by_key", "related_extended_data")
    THRESHOLD_FIELD_NUMBER: _ClassVar[int]
    TIMEFRAME_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_FIELD_NUMBER: _ClassVar[int]
    METRIC_ALERT_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    METRIC_ALERT_PROMQL_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    IGNORE_INFINITY_FIELD_NUMBER: _ClassVar[int]
    RELATIVE_TIMEFRAME_FIELD_NUMBER: _ClassVar[int]
    NOTIFY_GROUP_BY_ONLY_ALERTS_FIELD_NUMBER: _ClassVar[int]
    CARDINALITY_FIELDS_FIELD_NUMBER: _ClassVar[int]
    MAX_UNIQUE_COUNT_VALUES_FOR_GROUP_BY_KEY_FIELD_NUMBER: _ClassVar[int]
    RELATED_EXTENDED_DATA_FIELD_NUMBER: _ClassVar[int]
    threshold: _wrappers_pb2.DoubleValue
    timeframe: _alert_condition_pb2.Timeframe
    group_by: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    metric_alert_parameters: _alert_condition_pb2.MetricAlertConditionParameters
    metric_alert_promql_parameters: _alert_condition_pb2.MetricAlertPromqlConditionParameters
    ignore_infinity: _wrappers_pb2.BoolValue
    relative_timeframe: _alert_condition_pb2.RelativeTimeframe
    notify_group_by_only_alerts: _wrappers_pb2.BoolValue
    cardinality_fields: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    max_unique_count_values_for_group_by_key: _wrappers_pb2.UInt32Value
    related_extended_data: _alert_condition_pb2.RelatedExtendedData
    def __init__(self, threshold: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ..., timeframe: _Optional[_Union[_alert_condition_pb2.Timeframe, str]] = ..., group_by: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., metric_alert_parameters: _Optional[_Union[_alert_condition_pb2.MetricAlertConditionParameters, _Mapping]] = ..., metric_alert_promql_parameters: _Optional[_Union[_alert_condition_pb2.MetricAlertPromqlConditionParameters, _Mapping]] = ..., ignore_infinity: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., relative_timeframe: _Optional[_Union[_alert_condition_pb2.RelativeTimeframe, str]] = ..., notify_group_by_only_alerts: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., cardinality_fields: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., max_unique_count_values_for_group_by_key: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ..., related_extended_data: _Optional[_Union[_alert_condition_pb2.RelatedExtendedData, _Mapping]] = ...) -> None: ...
