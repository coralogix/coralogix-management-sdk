from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Query(_message.Message):
    __slots__ = ("vector",)
    VECTOR_FIELD_NUMBER: _ClassVar[int]
    vector: VectorOperation
    def __init__(self, vector: _Optional[_Union[VectorOperation, _Mapping]] = ...) -> None: ...

class VectorOperation(_message.Message):
    __slots__ = ("single", "unary", "binary", "subquery")
    class Unary(_message.Message):
        __slots__ = ("op", "operand")
        class Op(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            OP_UNSPECIFIED: _ClassVar[VectorOperation.Unary.Op]
            OP_PLUS: _ClassVar[VectorOperation.Unary.Op]
            OP_MINUS: _ClassVar[VectorOperation.Unary.Op]
        OP_UNSPECIFIED: VectorOperation.Unary.Op
        OP_PLUS: VectorOperation.Unary.Op
        OP_MINUS: VectorOperation.Unary.Op
        OP_FIELD_NUMBER: _ClassVar[int]
        OPERAND_FIELD_NUMBER: _ClassVar[int]
        op: VectorOperation.Unary.Op
        operand: VectorOperation
        def __init__(self, op: _Optional[_Union[VectorOperation.Unary.Op, str]] = ..., operand: _Optional[_Union[VectorOperation, _Mapping]] = ...) -> None: ...
    class Binary(_message.Message):
        __slots__ = ("left", "op", "grouping", "right")
        class Op(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            OP_UNSPECIFIED: _ClassVar[VectorOperation.Binary.Op]
            OP_ADD: _ClassVar[VectorOperation.Binary.Op]
            OP_SUB: _ClassVar[VectorOperation.Binary.Op]
            OP_MUL: _ClassVar[VectorOperation.Binary.Op]
            OP_DIV: _ClassVar[VectorOperation.Binary.Op]
            OP_MOD: _ClassVar[VectorOperation.Binary.Op]
            OP_ATAN2: _ClassVar[VectorOperation.Binary.Op]
            OP_POW: _ClassVar[VectorOperation.Binary.Op]
            OP_AND: _ClassVar[VectorOperation.Binary.Op]
            OP_OR: _ClassVar[VectorOperation.Binary.Op]
            OP_UNLESS: _ClassVar[VectorOperation.Binary.Op]
            OP_EQ: _ClassVar[VectorOperation.Binary.Op]
            OP_NEQ: _ClassVar[VectorOperation.Binary.Op]
            OP_LT: _ClassVar[VectorOperation.Binary.Op]
            OP_LTE: _ClassVar[VectorOperation.Binary.Op]
            OP_GT: _ClassVar[VectorOperation.Binary.Op]
            OP_GTE: _ClassVar[VectorOperation.Binary.Op]
            OP_EQ_BOOL: _ClassVar[VectorOperation.Binary.Op]
            OP_NEQ_BOOL: _ClassVar[VectorOperation.Binary.Op]
            OP_LT_BOOL: _ClassVar[VectorOperation.Binary.Op]
            OP_LTE_BOOL: _ClassVar[VectorOperation.Binary.Op]
            OP_GT_BOOL: _ClassVar[VectorOperation.Binary.Op]
            OP_GTE_BOOL: _ClassVar[VectorOperation.Binary.Op]
        OP_UNSPECIFIED: VectorOperation.Binary.Op
        OP_ADD: VectorOperation.Binary.Op
        OP_SUB: VectorOperation.Binary.Op
        OP_MUL: VectorOperation.Binary.Op
        OP_DIV: VectorOperation.Binary.Op
        OP_MOD: VectorOperation.Binary.Op
        OP_ATAN2: VectorOperation.Binary.Op
        OP_POW: VectorOperation.Binary.Op
        OP_AND: VectorOperation.Binary.Op
        OP_OR: VectorOperation.Binary.Op
        OP_UNLESS: VectorOperation.Binary.Op
        OP_EQ: VectorOperation.Binary.Op
        OP_NEQ: VectorOperation.Binary.Op
        OP_LT: VectorOperation.Binary.Op
        OP_LTE: VectorOperation.Binary.Op
        OP_GT: VectorOperation.Binary.Op
        OP_GTE: VectorOperation.Binary.Op
        OP_EQ_BOOL: VectorOperation.Binary.Op
        OP_NEQ_BOOL: VectorOperation.Binary.Op
        OP_LT_BOOL: VectorOperation.Binary.Op
        OP_LTE_BOOL: VectorOperation.Binary.Op
        OP_GT_BOOL: VectorOperation.Binary.Op
        OP_GTE_BOOL: VectorOperation.Binary.Op
        class Grouping(_message.Message):
            __slots__ = ("modifier", "bias")
            class Modifier(_message.Message):
                __slots__ = ("kind", "labels")
                class Kind(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
                    __slots__ = ()
                    KIND_UNSPECIFIED: _ClassVar[VectorOperation.Binary.Grouping.Modifier.Kind]
                    KIND_ON: _ClassVar[VectorOperation.Binary.Grouping.Modifier.Kind]
                    KIND_IGNORING: _ClassVar[VectorOperation.Binary.Grouping.Modifier.Kind]
                KIND_UNSPECIFIED: VectorOperation.Binary.Grouping.Modifier.Kind
                KIND_ON: VectorOperation.Binary.Grouping.Modifier.Kind
                KIND_IGNORING: VectorOperation.Binary.Grouping.Modifier.Kind
                KIND_FIELD_NUMBER: _ClassVar[int]
                LABELS_FIELD_NUMBER: _ClassVar[int]
                kind: VectorOperation.Binary.Grouping.Modifier.Kind
                labels: _containers.RepeatedScalarFieldContainer[str]
                def __init__(self, kind: _Optional[_Union[VectorOperation.Binary.Grouping.Modifier.Kind, str]] = ..., labels: _Optional[_Iterable[str]] = ...) -> None: ...
            class Bias(_message.Message):
                __slots__ = ("kind", "labels")
                class Kind(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
                    __slots__ = ()
                    KIND_UNSPECIFIED: _ClassVar[VectorOperation.Binary.Grouping.Bias.Kind]
                    KIND_LEFT: _ClassVar[VectorOperation.Binary.Grouping.Bias.Kind]
                    KIND_RIGHT: _ClassVar[VectorOperation.Binary.Grouping.Bias.Kind]
                KIND_UNSPECIFIED: VectorOperation.Binary.Grouping.Bias.Kind
                KIND_LEFT: VectorOperation.Binary.Grouping.Bias.Kind
                KIND_RIGHT: VectorOperation.Binary.Grouping.Bias.Kind
                KIND_FIELD_NUMBER: _ClassVar[int]
                LABELS_FIELD_NUMBER: _ClassVar[int]
                kind: VectorOperation.Binary.Grouping.Bias.Kind
                labels: _containers.RepeatedScalarFieldContainer[str]
                def __init__(self, kind: _Optional[_Union[VectorOperation.Binary.Grouping.Bias.Kind, str]] = ..., labels: _Optional[_Iterable[str]] = ...) -> None: ...
            MODIFIER_FIELD_NUMBER: _ClassVar[int]
            BIAS_FIELD_NUMBER: _ClassVar[int]
            modifier: VectorOperation.Binary.Grouping.Modifier
            bias: VectorOperation.Binary.Grouping.Bias
            def __init__(self, modifier: _Optional[_Union[VectorOperation.Binary.Grouping.Modifier, _Mapping]] = ..., bias: _Optional[_Union[VectorOperation.Binary.Grouping.Bias, _Mapping]] = ...) -> None: ...
        LEFT_FIELD_NUMBER: _ClassVar[int]
        OP_FIELD_NUMBER: _ClassVar[int]
        GROUPING_FIELD_NUMBER: _ClassVar[int]
        RIGHT_FIELD_NUMBER: _ClassVar[int]
        left: VectorOperation
        op: VectorOperation.Binary.Op
        grouping: VectorOperation.Binary.Grouping
        right: VectorOperation
        def __init__(self, left: _Optional[_Union[VectorOperation, _Mapping]] = ..., op: _Optional[_Union[VectorOperation.Binary.Op, str]] = ..., grouping: _Optional[_Union[VectorOperation.Binary.Grouping, _Mapping]] = ..., right: _Optional[_Union[VectorOperation, _Mapping]] = ...) -> None: ...
    class Subquery(_message.Message):
        __slots__ = ("operation", "range", "offset")
        class Range(_message.Message):
            __slots__ = ("to",)
            FROM_FIELD_NUMBER: _ClassVar[int]
            TO_FIELD_NUMBER: _ClassVar[int]
            to: Duration
            def __init__(self, to: _Optional[_Union[Duration, _Mapping]] = ..., **kwargs) -> None: ...
        OPERATION_FIELD_NUMBER: _ClassVar[int]
        RANGE_FIELD_NUMBER: _ClassVar[int]
        OFFSET_FIELD_NUMBER: _ClassVar[int]
        operation: VectorOperation
        range: VectorOperation.Subquery.Range
        offset: Duration
        def __init__(self, operation: _Optional[_Union[VectorOperation, _Mapping]] = ..., range: _Optional[_Union[VectorOperation.Subquery.Range, _Mapping]] = ..., offset: _Optional[_Union[Duration, _Mapping]] = ...) -> None: ...
    SINGLE_FIELD_NUMBER: _ClassVar[int]
    UNARY_FIELD_NUMBER: _ClassVar[int]
    BINARY_FIELD_NUMBER: _ClassVar[int]
    SUBQUERY_FIELD_NUMBER: _ClassVar[int]
    single: VectorExpression
    unary: VectorOperation.Unary
    binary: VectorOperation.Binary
    subquery: VectorOperation.Subquery
    def __init__(self, single: _Optional[_Union[VectorExpression, _Mapping]] = ..., unary: _Optional[_Union[VectorOperation.Unary, _Mapping]] = ..., binary: _Optional[_Union[VectorOperation.Binary, _Mapping]] = ..., subquery: _Optional[_Union[VectorOperation.Subquery, _Mapping]] = ...) -> None: ...

class VectorExpression(_message.Message):
    __slots__ = ("number", "string", "function_call", "aggregation", "selector", "grouped")
    class FunctionCall(_message.Message):
        __slots__ = ("name", "parameters")
        NAME_FIELD_NUMBER: _ClassVar[int]
        PARAMETERS_FIELD_NUMBER: _ClassVar[int]
        name: str
        parameters: _containers.RepeatedCompositeFieldContainer[VectorOperation]
        def __init__(self, name: _Optional[str] = ..., parameters: _Optional[_Iterable[_Union[VectorOperation, _Mapping]]] = ...) -> None: ...
    class Aggregation(_message.Message):
        __slots__ = ("name", "front_modifier", "parameters", "back_modifier")
        class Modifier(_message.Message):
            __slots__ = ("kind", "labels")
            class Kind(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
                __slots__ = ()
                KIND_UNSPECIFIED: _ClassVar[VectorExpression.Aggregation.Modifier.Kind]
                KIND_BY: _ClassVar[VectorExpression.Aggregation.Modifier.Kind]
                KIND_WITHOUT: _ClassVar[VectorExpression.Aggregation.Modifier.Kind]
            KIND_UNSPECIFIED: VectorExpression.Aggregation.Modifier.Kind
            KIND_BY: VectorExpression.Aggregation.Modifier.Kind
            KIND_WITHOUT: VectorExpression.Aggregation.Modifier.Kind
            KIND_FIELD_NUMBER: _ClassVar[int]
            LABELS_FIELD_NUMBER: _ClassVar[int]
            kind: VectorExpression.Aggregation.Modifier.Kind
            labels: _containers.RepeatedScalarFieldContainer[str]
            def __init__(self, kind: _Optional[_Union[VectorExpression.Aggregation.Modifier.Kind, str]] = ..., labels: _Optional[_Iterable[str]] = ...) -> None: ...
        NAME_FIELD_NUMBER: _ClassVar[int]
        FRONT_MODIFIER_FIELD_NUMBER: _ClassVar[int]
        PARAMETERS_FIELD_NUMBER: _ClassVar[int]
        BACK_MODIFIER_FIELD_NUMBER: _ClassVar[int]
        name: str
        front_modifier: VectorExpression.Aggregation.Modifier
        parameters: _containers.RepeatedCompositeFieldContainer[VectorOperation]
        back_modifier: VectorExpression.Aggregation.Modifier
        def __init__(self, name: _Optional[str] = ..., front_modifier: _Optional[_Union[VectorExpression.Aggregation.Modifier, _Mapping]] = ..., parameters: _Optional[_Iterable[_Union[VectorOperation, _Mapping]]] = ..., back_modifier: _Optional[_Union[VectorExpression.Aggregation.Modifier, _Mapping]] = ...) -> None: ...
    class Selector(_message.Message):
        __slots__ = ("metric", "label_matchers", "time_range", "offset", "at_modifier")
        class LabelMatcher(_message.Message):
            __slots__ = ("name", "op", "value")
            class Op(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
                __slots__ = ()
                OP_UNSPECIFIED: _ClassVar[VectorExpression.Selector.LabelMatcher.Op]
                OP_EQ: _ClassVar[VectorExpression.Selector.LabelMatcher.Op]
                OP_NEQ: _ClassVar[VectorExpression.Selector.LabelMatcher.Op]
                OP_MATCHES: _ClassVar[VectorExpression.Selector.LabelMatcher.Op]
                OP_NOT_MATCHES: _ClassVar[VectorExpression.Selector.LabelMatcher.Op]
            OP_UNSPECIFIED: VectorExpression.Selector.LabelMatcher.Op
            OP_EQ: VectorExpression.Selector.LabelMatcher.Op
            OP_NEQ: VectorExpression.Selector.LabelMatcher.Op
            OP_MATCHES: VectorExpression.Selector.LabelMatcher.Op
            OP_NOT_MATCHES: VectorExpression.Selector.LabelMatcher.Op
            NAME_FIELD_NUMBER: _ClassVar[int]
            OP_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            name: str
            op: VectorExpression.Selector.LabelMatcher.Op
            value: str
            def __init__(self, name: _Optional[str] = ..., op: _Optional[_Union[VectorExpression.Selector.LabelMatcher.Op, str]] = ..., value: _Optional[str] = ...) -> None: ...
        class Modifier(_message.Message):
            __slots__ = ("timestamp", "boundary")
            class Boundary(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
                __slots__ = ()
                BOUNDARY_UNSPECIFIED: _ClassVar[VectorExpression.Selector.Modifier.Boundary]
                BOUNDARY_START: _ClassVar[VectorExpression.Selector.Modifier.Boundary]
                BOUNDARY_END: _ClassVar[VectorExpression.Selector.Modifier.Boundary]
            BOUNDARY_UNSPECIFIED: VectorExpression.Selector.Modifier.Boundary
            BOUNDARY_START: VectorExpression.Selector.Modifier.Boundary
            BOUNDARY_END: VectorExpression.Selector.Modifier.Boundary
            TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
            BOUNDARY_FIELD_NUMBER: _ClassVar[int]
            timestamp: _timestamp_pb2.Timestamp
            boundary: VectorExpression.Selector.Modifier.Boundary
            def __init__(self, timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., boundary: _Optional[_Union[VectorExpression.Selector.Modifier.Boundary, str]] = ...) -> None: ...
        METRIC_FIELD_NUMBER: _ClassVar[int]
        LABEL_MATCHERS_FIELD_NUMBER: _ClassVar[int]
        TIME_RANGE_FIELD_NUMBER: _ClassVar[int]
        OFFSET_FIELD_NUMBER: _ClassVar[int]
        AT_MODIFIER_FIELD_NUMBER: _ClassVar[int]
        metric: _wrappers_pb2.StringValue
        label_matchers: _containers.RepeatedCompositeFieldContainer[VectorExpression.Selector.LabelMatcher]
        time_range: Duration
        offset: Duration
        at_modifier: VectorExpression.Selector.Modifier
        def __init__(self, metric: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., label_matchers: _Optional[_Iterable[_Union[VectorExpression.Selector.LabelMatcher, _Mapping]]] = ..., time_range: _Optional[_Union[Duration, _Mapping]] = ..., offset: _Optional[_Union[Duration, _Mapping]] = ..., at_modifier: _Optional[_Union[VectorExpression.Selector.Modifier, _Mapping]] = ...) -> None: ...
    NUMBER_FIELD_NUMBER: _ClassVar[int]
    STRING_FIELD_NUMBER: _ClassVar[int]
    FUNCTION_CALL_FIELD_NUMBER: _ClassVar[int]
    AGGREGATION_FIELD_NUMBER: _ClassVar[int]
    SELECTOR_FIELD_NUMBER: _ClassVar[int]
    GROUPED_FIELD_NUMBER: _ClassVar[int]
    number: Num
    string: _wrappers_pb2.StringValue
    function_call: VectorExpression.FunctionCall
    aggregation: VectorExpression.Aggregation
    selector: VectorExpression.Selector
    grouped: VectorOperation
    def __init__(self, number: _Optional[_Union[Num, _Mapping]] = ..., string: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., function_call: _Optional[_Union[VectorExpression.FunctionCall, _Mapping]] = ..., aggregation: _Optional[_Union[VectorExpression.Aggregation, _Mapping]] = ..., selector: _Optional[_Union[VectorExpression.Selector, _Mapping]] = ..., grouped: _Optional[_Union[VectorOperation, _Mapping]] = ...) -> None: ...

class Num(_message.Message):
    __slots__ = ("integer", "floating", "infinity")
    class Infinity(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    INTEGER_FIELD_NUMBER: _ClassVar[int]
    FLOATING_FIELD_NUMBER: _ClassVar[int]
    INFINITY_FIELD_NUMBER: _ClassVar[int]
    integer: _wrappers_pb2.Int64Value
    floating: _wrappers_pb2.DoubleValue
    infinity: Num.Infinity
    def __init__(self, integer: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., floating: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ..., infinity: _Optional[_Union[Num.Infinity, _Mapping]] = ...) -> None: ...

class Duration(_message.Message):
    __slots__ = ("amount", "unit")
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    UNIT_FIELD_NUMBER: _ClassVar[int]
    amount: Num
    unit: str
    def __init__(self, amount: _Optional[_Union[Num, _Mapping]] = ..., unit: _Optional[str] = ...) -> None: ...
