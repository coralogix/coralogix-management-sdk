from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SpansAggregation(_message.Message):
    __slots__ = ("metric_aggregation", "dimension_aggregation")
    class MetricAggregation(_message.Message):
        __slots__ = ("metric_field", "aggregation_type")
        class MetricField(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            METRIC_FIELD_UNSPECIFIED: _ClassVar[SpansAggregation.MetricAggregation.MetricField]
            METRIC_FIELD_DURATION: _ClassVar[SpansAggregation.MetricAggregation.MetricField]
        METRIC_FIELD_UNSPECIFIED: SpansAggregation.MetricAggregation.MetricField
        METRIC_FIELD_DURATION: SpansAggregation.MetricAggregation.MetricField
        class MetricAggregationType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            METRIC_AGGREGATION_TYPE_UNSPECIFIED: _ClassVar[SpansAggregation.MetricAggregation.MetricAggregationType]
            METRIC_AGGREGATION_TYPE_MIN: _ClassVar[SpansAggregation.MetricAggregation.MetricAggregationType]
            METRIC_AGGREGATION_TYPE_MAX: _ClassVar[SpansAggregation.MetricAggregation.MetricAggregationType]
            METRIC_AGGREGATION_TYPE_AVERAGE: _ClassVar[SpansAggregation.MetricAggregation.MetricAggregationType]
            METRIC_AGGREGATION_TYPE_SUM: _ClassVar[SpansAggregation.MetricAggregation.MetricAggregationType]
            METRIC_AGGREGATION_TYPE_PERCENTILE_99: _ClassVar[SpansAggregation.MetricAggregation.MetricAggregationType]
            METRIC_AGGREGATION_TYPE_PERCENTILE_95: _ClassVar[SpansAggregation.MetricAggregation.MetricAggregationType]
            METRIC_AGGREGATION_TYPE_PERCENTILE_50: _ClassVar[SpansAggregation.MetricAggregation.MetricAggregationType]
        METRIC_AGGREGATION_TYPE_UNSPECIFIED: SpansAggregation.MetricAggregation.MetricAggregationType
        METRIC_AGGREGATION_TYPE_MIN: SpansAggregation.MetricAggregation.MetricAggregationType
        METRIC_AGGREGATION_TYPE_MAX: SpansAggregation.MetricAggregation.MetricAggregationType
        METRIC_AGGREGATION_TYPE_AVERAGE: SpansAggregation.MetricAggregation.MetricAggregationType
        METRIC_AGGREGATION_TYPE_SUM: SpansAggregation.MetricAggregation.MetricAggregationType
        METRIC_AGGREGATION_TYPE_PERCENTILE_99: SpansAggregation.MetricAggregation.MetricAggregationType
        METRIC_AGGREGATION_TYPE_PERCENTILE_95: SpansAggregation.MetricAggregation.MetricAggregationType
        METRIC_AGGREGATION_TYPE_PERCENTILE_50: SpansAggregation.MetricAggregation.MetricAggregationType
        METRIC_FIELD_FIELD_NUMBER: _ClassVar[int]
        AGGREGATION_TYPE_FIELD_NUMBER: _ClassVar[int]
        metric_field: SpansAggregation.MetricAggregation.MetricField
        aggregation_type: SpansAggregation.MetricAggregation.MetricAggregationType
        def __init__(self, metric_field: _Optional[_Union[SpansAggregation.MetricAggregation.MetricField, str]] = ..., aggregation_type: _Optional[_Union[SpansAggregation.MetricAggregation.MetricAggregationType, str]] = ...) -> None: ...
    class DimensionAggregation(_message.Message):
        __slots__ = ("dimension_field", "aggregation_type")
        class DimensionField(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            DIMENSION_FIELD_UNSPECIFIED: _ClassVar[SpansAggregation.DimensionAggregation.DimensionField]
            DIMENSION_FIELD_TRACE_ID: _ClassVar[SpansAggregation.DimensionAggregation.DimensionField]
        DIMENSION_FIELD_UNSPECIFIED: SpansAggregation.DimensionAggregation.DimensionField
        DIMENSION_FIELD_TRACE_ID: SpansAggregation.DimensionAggregation.DimensionField
        class DimensionAggregationType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            DIMENSION_AGGREGATION_TYPE_UNSPECIFIED: _ClassVar[SpansAggregation.DimensionAggregation.DimensionAggregationType]
            DIMENSION_AGGREGATION_TYPE_UNIQUE_COUNT: _ClassVar[SpansAggregation.DimensionAggregation.DimensionAggregationType]
            DIMENSION_AGGREGATION_TYPE_ERROR_COUNT: _ClassVar[SpansAggregation.DimensionAggregation.DimensionAggregationType]
        DIMENSION_AGGREGATION_TYPE_UNSPECIFIED: SpansAggregation.DimensionAggregation.DimensionAggregationType
        DIMENSION_AGGREGATION_TYPE_UNIQUE_COUNT: SpansAggregation.DimensionAggregation.DimensionAggregationType
        DIMENSION_AGGREGATION_TYPE_ERROR_COUNT: SpansAggregation.DimensionAggregation.DimensionAggregationType
        DIMENSION_FIELD_FIELD_NUMBER: _ClassVar[int]
        AGGREGATION_TYPE_FIELD_NUMBER: _ClassVar[int]
        dimension_field: SpansAggregation.DimensionAggregation.DimensionField
        aggregation_type: SpansAggregation.DimensionAggregation.DimensionAggregationType
        def __init__(self, dimension_field: _Optional[_Union[SpansAggregation.DimensionAggregation.DimensionField, str]] = ..., aggregation_type: _Optional[_Union[SpansAggregation.DimensionAggregation.DimensionAggregationType, str]] = ...) -> None: ...
    METRIC_AGGREGATION_FIELD_NUMBER: _ClassVar[int]
    DIMENSION_AGGREGATION_FIELD_NUMBER: _ClassVar[int]
    metric_aggregation: SpansAggregation.MetricAggregation
    dimension_aggregation: SpansAggregation.DimensionAggregation
    def __init__(self, metric_aggregation: _Optional[_Union[SpansAggregation.MetricAggregation, _Mapping]] = ..., dimension_aggregation: _Optional[_Union[SpansAggregation.DimensionAggregation, _Mapping]] = ...) -> None: ...
