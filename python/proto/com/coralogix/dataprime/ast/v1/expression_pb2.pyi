from com.coralogix.dataprime.ast.v1 import types_pb2 as _types_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Expression(_message.Message):
    __slots__ = ("null_e", "keypath", "boolean", "integer", "str", "regex", "infix_op", "function_call", "cast", "match", "long", "double", "time_unit", "interval", "timestamp", "date_unit", "array", "array_slice", "enum_value", "array_element", "datatype")
    class TimeUnit(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        TIME_UNIT_UNSPECIFIED: _ClassVar[Expression.TimeUnit]
        TIME_UNIT_NANO: _ClassVar[Expression.TimeUnit]
        TIME_UNIT_MICRO: _ClassVar[Expression.TimeUnit]
        TIME_UNIT_MILLI: _ClassVar[Expression.TimeUnit]
        TIME_UNIT_SECOND: _ClassVar[Expression.TimeUnit]
        TIME_UNIT_MINUTE: _ClassVar[Expression.TimeUnit]
        TIME_UNIT_HOUR: _ClassVar[Expression.TimeUnit]
        TIME_UNIT_DAY: _ClassVar[Expression.TimeUnit]
    TIME_UNIT_UNSPECIFIED: Expression.TimeUnit
    TIME_UNIT_NANO: Expression.TimeUnit
    TIME_UNIT_MICRO: Expression.TimeUnit
    TIME_UNIT_MILLI: Expression.TimeUnit
    TIME_UNIT_SECOND: Expression.TimeUnit
    TIME_UNIT_MINUTE: Expression.TimeUnit
    TIME_UNIT_HOUR: Expression.TimeUnit
    TIME_UNIT_DAY: Expression.TimeUnit
    class DateUnit(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        DATE_UNIT_UNSPECIFIED: _ClassVar[Expression.DateUnit]
        DATE_UNIT_YEAR: _ClassVar[Expression.DateUnit]
        DATE_UNIT_MONTH: _ClassVar[Expression.DateUnit]
        DATE_UNIT_WEEK: _ClassVar[Expression.DateUnit]
        DATE_UNIT_DAY_OF_YEAR: _ClassVar[Expression.DateUnit]
        DATE_UNIT_DAY_OF_WEEK: _ClassVar[Expression.DateUnit]
    DATE_UNIT_UNSPECIFIED: Expression.DateUnit
    DATE_UNIT_YEAR: Expression.DateUnit
    DATE_UNIT_MONTH: Expression.DateUnit
    DATE_UNIT_WEEK: Expression.DateUnit
    DATE_UNIT_DAY_OF_YEAR: Expression.DateUnit
    DATE_UNIT_DAY_OF_WEEK: Expression.DateUnit
    class Match(_message.Message):
        __slots__ = ("str", "regex")
        STR_FIELD_NUMBER: _ClassVar[int]
        REGEX_FIELD_NUMBER: _ClassVar[int]
        str: _wrappers_pb2.StringValue
        regex: Expression.Regex
        def __init__(self, str: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., regex: _Optional[_Union[Expression.Regex, _Mapping]] = ...) -> None: ...
    class Null(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class Keypath(_message.Message):
        __slots__ = ("root", "path_elements", "datatype")
        class Root(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            ROOT_UNSPECIFIED: _ClassVar[Expression.Keypath.Root]
            ROOT_EVENT_LABELS: _ClassVar[Expression.Keypath.Root]
            ROOT_EVENT_METADATA: _ClassVar[Expression.Keypath.Root]
            ROOT_USER_DATA: _ClassVar[Expression.Keypath.Root]
        ROOT_UNSPECIFIED: Expression.Keypath.Root
        ROOT_EVENT_LABELS: Expression.Keypath.Root
        ROOT_EVENT_METADATA: Expression.Keypath.Root
        ROOT_USER_DATA: Expression.Keypath.Root
        ROOT_FIELD_NUMBER: _ClassVar[int]
        PATH_ELEMENTS_FIELD_NUMBER: _ClassVar[int]
        DATATYPE_FIELD_NUMBER: _ClassVar[int]
        root: Expression.Keypath.Root
        path_elements: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        datatype: _types_pb2.Datatype
        def __init__(self, root: _Optional[_Union[Expression.Keypath.Root, str]] = ..., path_elements: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., datatype: _Optional[_Union[_types_pb2.Datatype, _Mapping]] = ...) -> None: ...
    class Regex(_message.Message):
        __slots__ = ("regex", "flags")
        REGEX_FIELD_NUMBER: _ClassVar[int]
        FLAGS_FIELD_NUMBER: _ClassVar[int]
        regex: _wrappers_pb2.StringValue
        flags: _wrappers_pb2.StringValue
        def __init__(self, regex: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., flags: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class InfixOp(_message.Message):
        __slots__ = ("left", "op", "right")
        class Op(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            OP_UNSPECIFIED: _ClassVar[Expression.InfixOp.Op]
            OP_ADD: _ClassVar[Expression.InfixOp.Op]
            OP_SUB: _ClassVar[Expression.InfixOp.Op]
            OP_MUL: _ClassVar[Expression.InfixOp.Op]
            OP_DIV: _ClassVar[Expression.InfixOp.Op]
            OP_AND: _ClassVar[Expression.InfixOp.Op]
            OP_OR: _ClassVar[Expression.InfixOp.Op]
            OP_EQ: _ClassVar[Expression.InfixOp.Op]
            OP_NEQ: _ClassVar[Expression.InfixOp.Op]
            OP_LT: _ClassVar[Expression.InfixOp.Op]
            OP_LTE: _ClassVar[Expression.InfixOp.Op]
            OP_GT: _ClassVar[Expression.InfixOp.Op]
            OP_GTE: _ClassVar[Expression.InfixOp.Op]
            OP_LIKE: _ClassVar[Expression.InfixOp.Op]
            OP_MOD: _ClassVar[Expression.InfixOp.Op]
            OP_ILIKE: _ClassVar[Expression.InfixOp.Op]
        OP_UNSPECIFIED: Expression.InfixOp.Op
        OP_ADD: Expression.InfixOp.Op
        OP_SUB: Expression.InfixOp.Op
        OP_MUL: Expression.InfixOp.Op
        OP_DIV: Expression.InfixOp.Op
        OP_AND: Expression.InfixOp.Op
        OP_OR: Expression.InfixOp.Op
        OP_EQ: Expression.InfixOp.Op
        OP_NEQ: Expression.InfixOp.Op
        OP_LT: Expression.InfixOp.Op
        OP_LTE: Expression.InfixOp.Op
        OP_GT: Expression.InfixOp.Op
        OP_GTE: Expression.InfixOp.Op
        OP_LIKE: Expression.InfixOp.Op
        OP_MOD: Expression.InfixOp.Op
        OP_ILIKE: Expression.InfixOp.Op
        LEFT_FIELD_NUMBER: _ClassVar[int]
        OP_FIELD_NUMBER: _ClassVar[int]
        RIGHT_FIELD_NUMBER: _ClassVar[int]
        left: Expression
        op: Expression.InfixOp.Op
        right: Expression
        def __init__(self, left: _Optional[_Union[Expression, _Mapping]] = ..., op: _Optional[_Union[Expression.InfixOp.Op, str]] = ..., right: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
    class Not(_message.Message):
        __slots__ = ("expression",)
        EXPRESSION_FIELD_NUMBER: _ClassVar[int]
        expression: Expression
        def __init__(self, expression: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
    class FunctionCall(_message.Message):
        __slots__ = ("v2",)
        V2_FIELD_NUMBER: _ClassVar[int]
        v2: Expression.FunctionCallV2
        def __init__(self, v2: _Optional[_Union[Expression.FunctionCallV2, _Mapping]] = ...) -> None: ...
    class Interval(_message.Message):
        __slots__ = ("nanos",)
        NANOS_FIELD_NUMBER: _ClassVar[int]
        nanos: int
        def __init__(self, nanos: _Optional[int] = ...) -> None: ...
    class Timestamp(_message.Message):
        __slots__ = ("epoch_nanos",)
        EPOCH_NANOS_FIELD_NUMBER: _ClassVar[int]
        epoch_nanos: int
        def __init__(self, epoch_nanos: _Optional[int] = ...) -> None: ...
    class Array(_message.Message):
        __slots__ = ("elements",)
        ELEMENTS_FIELD_NUMBER: _ClassVar[int]
        elements: _containers.RepeatedCompositeFieldContainer[Expression]
        def __init__(self, elements: _Optional[_Iterable[_Union[Expression, _Mapping]]] = ...) -> None: ...
    class ArraySlice(_message.Message):
        __slots__ = ("array", "start", "end")
        ARRAY_FIELD_NUMBER: _ClassVar[int]
        START_FIELD_NUMBER: _ClassVar[int]
        END_FIELD_NUMBER: _ClassVar[int]
        array: Expression
        start: Expression
        end: Expression
        def __init__(self, array: _Optional[_Union[Expression, _Mapping]] = ..., start: _Optional[_Union[Expression, _Mapping]] = ..., end: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
    class ArrayElement(_message.Message):
        __slots__ = ("array", "index")
        ARRAY_FIELD_NUMBER: _ClassVar[int]
        INDEX_FIELD_NUMBER: _ClassVar[int]
        array: Expression
        index: Expression
        def __init__(self, array: _Optional[_Union[Expression, _Mapping]] = ..., index: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
    class FunctionCallV2(_message.Message):
        __slots__ = ("first_non_null", "decode_base_64", "encode_base_64", "length", "substr", "concat", "trim", "rtrim", "ltrim", "to_lower_case", "to_upper_case", "round_time", "contains", "index_of", "starts_with", "ends_with", "round", "abs", "min", "max", "floor", "ceil", "power", "log", "log2", "ln", "to_iso_8601_date_time", "from_unix_time", "pad", "pad_left", "pad_right", "chr", "codepoint", "split_parts", "from_base", "to_base", "mod", "sqrt", "e", "pi", "random", "random_int", "url_encode", "url_decode", "matches", "regexp_split_parts", "ip_prefix", "ip_in_subnet", "random_uuid", "is_uuid", "to_unix_time", "now", "record_location", "parse_interval", "format_interval", "to_interval", "parse_timestamp", "format_timestamp", "round_interval", "extract_time", "array_length", "array_is_empty", "cardinality", "array_contains", "array_split", "array_join", "set_union", "set_intersection", "set_diff", "set_diff_symmetric", "is_subset", "is_superset", "set_equals_to", "array_concat", "array_append", "array_insert_at", "array_replace_at", "array_replace_all", "array_remove_at", "array_remove")
        class FirstNonNull(_message.Message):
            __slots__ = ("value", "values")
            VALUE_FIELD_NUMBER: _ClassVar[int]
            VALUES_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            values: _containers.RepeatedCompositeFieldContainer[Expression]
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ..., values: _Optional[_Iterable[_Union[Expression, _Mapping]]] = ...) -> None: ...
        class In(_message.Message):
            __slots__ = ("comparand", "value", "values")
            COMPARAND_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            VALUES_FIELD_NUMBER: _ClassVar[int]
            comparand: Expression
            value: Expression
            values: _containers.RepeatedCompositeFieldContainer[Expression]
            def __init__(self, comparand: _Optional[_Union[Expression, _Mapping]] = ..., value: _Optional[_Union[Expression, _Mapping]] = ..., values: _Optional[_Iterable[_Union[Expression, _Mapping]]] = ...) -> None: ...
        class If(_message.Message):
            __slots__ = ("condition", "then")
            CONDITION_FIELD_NUMBER: _ClassVar[int]
            THEN_FIELD_NUMBER: _ClassVar[int]
            ELSE_FIELD_NUMBER: _ClassVar[int]
            condition: Expression
            then: Expression
            def __init__(self, condition: _Optional[_Union[Expression, _Mapping]] = ..., then: _Optional[_Union[Expression, _Mapping]] = ..., **kwargs) -> None: ...
        class DecodeBase64(_message.Message):
            __slots__ = ("value",)
            VALUE_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class EncodeBase64(_message.Message):
            __slots__ = ("value",)
            VALUE_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Length(_message.Message):
            __slots__ = ("value",)
            VALUE_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Substr(_message.Message):
            __slots__ = ("value", "length")
            VALUE_FIELD_NUMBER: _ClassVar[int]
            FROM_FIELD_NUMBER: _ClassVar[int]
            LENGTH_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            length: Expression
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ..., length: _Optional[_Union[Expression, _Mapping]] = ..., **kwargs) -> None: ...
        class Concat(_message.Message):
            __slots__ = ("value", "values")
            VALUE_FIELD_NUMBER: _ClassVar[int]
            VALUES_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            values: _containers.RepeatedCompositeFieldContainer[Expression]
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ..., values: _Optional[_Iterable[_Union[Expression, _Mapping]]] = ...) -> None: ...
        class Trim(_message.Message):
            __slots__ = ("value",)
            VALUE_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Rtrim(_message.Message):
            __slots__ = ("value",)
            VALUE_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Ltrim(_message.Message):
            __slots__ = ("value",)
            VALUE_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Pad(_message.Message):
            __slots__ = ("value", "char_count", "fill_with")
            VALUE_FIELD_NUMBER: _ClassVar[int]
            CHAR_COUNT_FIELD_NUMBER: _ClassVar[int]
            FILL_WITH_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            char_count: Expression
            fill_with: Expression
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ..., char_count: _Optional[_Union[Expression, _Mapping]] = ..., fill_with: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class PadLeft(_message.Message):
            __slots__ = ("value", "char_count", "fill_with")
            VALUE_FIELD_NUMBER: _ClassVar[int]
            CHAR_COUNT_FIELD_NUMBER: _ClassVar[int]
            FILL_WITH_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            char_count: Expression
            fill_with: Expression
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ..., char_count: _Optional[_Union[Expression, _Mapping]] = ..., fill_with: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class PadRight(_message.Message):
            __slots__ = ("value", "char_count", "fill_with")
            VALUE_FIELD_NUMBER: _ClassVar[int]
            CHAR_COUNT_FIELD_NUMBER: _ClassVar[int]
            FILL_WITH_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            char_count: Expression
            fill_with: Expression
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ..., char_count: _Optional[_Union[Expression, _Mapping]] = ..., fill_with: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Chr(_message.Message):
            __slots__ = ("number",)
            NUMBER_FIELD_NUMBER: _ClassVar[int]
            number: Expression
            def __init__(self, number: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Codepoint(_message.Message):
            __slots__ = ("string",)
            STRING_FIELD_NUMBER: _ClassVar[int]
            string: Expression
            def __init__(self, string: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class SplitParts(_message.Message):
            __slots__ = ("string", "delimiter", "index")
            STRING_FIELD_NUMBER: _ClassVar[int]
            DELIMITER_FIELD_NUMBER: _ClassVar[int]
            INDEX_FIELD_NUMBER: _ClassVar[int]
            string: Expression
            delimiter: Expression
            index: Expression
            def __init__(self, string: _Optional[_Union[Expression, _Mapping]] = ..., delimiter: _Optional[_Union[Expression, _Mapping]] = ..., index: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ToLowerCase(_message.Message):
            __slots__ = ("value",)
            VALUE_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ToUpperCase(_message.Message):
            __slots__ = ("value",)
            VALUE_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Contains(_message.Message):
            __slots__ = ("string", "substring")
            STRING_FIELD_NUMBER: _ClassVar[int]
            SUBSTRING_FIELD_NUMBER: _ClassVar[int]
            string: Expression
            substring: Expression
            def __init__(self, string: _Optional[_Union[Expression, _Mapping]] = ..., substring: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class IndexOf(_message.Message):
            __slots__ = ("string", "substring")
            STRING_FIELD_NUMBER: _ClassVar[int]
            SUBSTRING_FIELD_NUMBER: _ClassVar[int]
            string: Expression
            substring: Expression
            def __init__(self, string: _Optional[_Union[Expression, _Mapping]] = ..., substring: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class StartsWith(_message.Message):
            __slots__ = ("string", "prefix")
            STRING_FIELD_NUMBER: _ClassVar[int]
            PREFIX_FIELD_NUMBER: _ClassVar[int]
            string: Expression
            prefix: Expression
            def __init__(self, string: _Optional[_Union[Expression, _Mapping]] = ..., prefix: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class EndsWith(_message.Message):
            __slots__ = ("string", "suffix")
            STRING_FIELD_NUMBER: _ClassVar[int]
            SUFFIX_FIELD_NUMBER: _ClassVar[int]
            string: Expression
            suffix: Expression
            def __init__(self, string: _Optional[_Union[Expression, _Mapping]] = ..., suffix: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Round(_message.Message):
            __slots__ = ("number", "digits")
            NUMBER_FIELD_NUMBER: _ClassVar[int]
            DIGITS_FIELD_NUMBER: _ClassVar[int]
            number: Expression
            digits: Expression
            def __init__(self, number: _Optional[_Union[Expression, _Mapping]] = ..., digits: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class FromBase(_message.Message):
            __slots__ = ("string", "radix")
            STRING_FIELD_NUMBER: _ClassVar[int]
            RADIX_FIELD_NUMBER: _ClassVar[int]
            string: Expression
            radix: Expression
            def __init__(self, string: _Optional[_Union[Expression, _Mapping]] = ..., radix: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ToBase(_message.Message):
            __slots__ = ("number", "radix")
            NUMBER_FIELD_NUMBER: _ClassVar[int]
            RADIX_FIELD_NUMBER: _ClassVar[int]
            number: Expression
            radix: Expression
            def __init__(self, number: _Optional[_Union[Expression, _Mapping]] = ..., radix: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Mod(_message.Message):
            __slots__ = ("number", "divisor")
            NUMBER_FIELD_NUMBER: _ClassVar[int]
            DIVISOR_FIELD_NUMBER: _ClassVar[int]
            number: Expression
            divisor: Expression
            def __init__(self, number: _Optional[_Union[Expression, _Mapping]] = ..., divisor: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Sqrt(_message.Message):
            __slots__ = ("number",)
            NUMBER_FIELD_NUMBER: _ClassVar[int]
            number: Expression
            def __init__(self, number: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Abs(_message.Message):
            __slots__ = ("number",)
            NUMBER_FIELD_NUMBER: _ClassVar[int]
            number: Expression
            def __init__(self, number: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class E(_message.Message):
            __slots__ = ()
            def __init__(self) -> None: ...
        class Pi(_message.Message):
            __slots__ = ()
            def __init__(self) -> None: ...
        class Random(_message.Message):
            __slots__ = ()
            def __init__(self) -> None: ...
        class RandomInt(_message.Message):
            __slots__ = ("upper_bound",)
            UPPER_BOUND_FIELD_NUMBER: _ClassVar[int]
            upper_bound: Expression
            def __init__(self, upper_bound: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Min(_message.Message):
            __slots__ = ("value", "values")
            VALUE_FIELD_NUMBER: _ClassVar[int]
            VALUES_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            values: _containers.RepeatedCompositeFieldContainer[Expression]
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ..., values: _Optional[_Iterable[_Union[Expression, _Mapping]]] = ...) -> None: ...
        class Max(_message.Message):
            __slots__ = ("value", "values")
            VALUE_FIELD_NUMBER: _ClassVar[int]
            VALUES_FIELD_NUMBER: _ClassVar[int]
            value: Expression
            values: _containers.RepeatedCompositeFieldContainer[Expression]
            def __init__(self, value: _Optional[_Union[Expression, _Mapping]] = ..., values: _Optional[_Iterable[_Union[Expression, _Mapping]]] = ...) -> None: ...
        class Floor(_message.Message):
            __slots__ = ("number",)
            NUMBER_FIELD_NUMBER: _ClassVar[int]
            number: Expression
            def __init__(self, number: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Ceil(_message.Message):
            __slots__ = ("number",)
            NUMBER_FIELD_NUMBER: _ClassVar[int]
            number: Expression
            def __init__(self, number: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Power(_message.Message):
            __slots__ = ("number", "exponent")
            NUMBER_FIELD_NUMBER: _ClassVar[int]
            EXPONENT_FIELD_NUMBER: _ClassVar[int]
            number: Expression
            exponent: Expression
            def __init__(self, number: _Optional[_Union[Expression, _Mapping]] = ..., exponent: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Log2(_message.Message):
            __slots__ = ("number",)
            NUMBER_FIELD_NUMBER: _ClassVar[int]
            number: Expression
            def __init__(self, number: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Log(_message.Message):
            __slots__ = ("base", "number")
            BASE_FIELD_NUMBER: _ClassVar[int]
            NUMBER_FIELD_NUMBER: _ClassVar[int]
            base: Expression
            number: Expression
            def __init__(self, base: _Optional[_Union[Expression, _Mapping]] = ..., number: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Ln(_message.Message):
            __slots__ = ("number",)
            NUMBER_FIELD_NUMBER: _ClassVar[int]
            number: Expression
            def __init__(self, number: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ToIso8601DateTime(_message.Message):
            __slots__ = ("timestamp",)
            TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
            timestamp: Expression
            def __init__(self, timestamp: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class RoundTime(_message.Message):
            __slots__ = ("interval", "date")
            INTERVAL_FIELD_NUMBER: _ClassVar[int]
            DATE_FIELD_NUMBER: _ClassVar[int]
            interval: Expression
            date: Expression
            def __init__(self, interval: _Optional[_Union[Expression, _Mapping]] = ..., date: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class FromUnixTime(_message.Message):
            __slots__ = ("unix_time", "time_unit")
            UNIX_TIME_FIELD_NUMBER: _ClassVar[int]
            TIME_UNIT_FIELD_NUMBER: _ClassVar[int]
            unix_time: Expression
            time_unit: Expression
            def __init__(self, unix_time: _Optional[_Union[Expression, _Mapping]] = ..., time_unit: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ToUnixTime(_message.Message):
            __slots__ = ("timestamp", "time_unit")
            TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
            TIME_UNIT_FIELD_NUMBER: _ClassVar[int]
            timestamp: Expression
            time_unit: Expression
            def __init__(self, timestamp: _Optional[_Union[Expression, _Mapping]] = ..., time_unit: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Now(_message.Message):
            __slots__ = ()
            def __init__(self) -> None: ...
        class ParseInterval(_message.Message):
            __slots__ = ("string",)
            STRING_FIELD_NUMBER: _ClassVar[int]
            string: Expression
            def __init__(self, string: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class FormatInterval(_message.Message):
            __slots__ = ("interval", "scale")
            INTERVAL_FIELD_NUMBER: _ClassVar[int]
            SCALE_FIELD_NUMBER: _ClassVar[int]
            interval: Expression
            scale: Expression
            def __init__(self, interval: _Optional[_Union[Expression, _Mapping]] = ..., scale: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ToInterval(_message.Message):
            __slots__ = ("number", "time_unit")
            NUMBER_FIELD_NUMBER: _ClassVar[int]
            TIME_UNIT_FIELD_NUMBER: _ClassVar[int]
            number: Expression
            time_unit: Expression
            def __init__(self, number: _Optional[_Union[Expression, _Mapping]] = ..., time_unit: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ParseTimestamp(_message.Message):
            __slots__ = ("string", "format", "tz")
            STRING_FIELD_NUMBER: _ClassVar[int]
            FORMAT_FIELD_NUMBER: _ClassVar[int]
            TZ_FIELD_NUMBER: _ClassVar[int]
            string: Expression
            format: Expression
            tz: Expression
            def __init__(self, string: _Optional[_Union[Expression, _Mapping]] = ..., format: _Optional[_Union[Expression, _Mapping]] = ..., tz: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class FormatTimestamp(_message.Message):
            __slots__ = ("timestamp", "format", "tz")
            TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
            FORMAT_FIELD_NUMBER: _ClassVar[int]
            TZ_FIELD_NUMBER: _ClassVar[int]
            timestamp: Expression
            format: Expression
            tz: Expression
            def __init__(self, timestamp: _Optional[_Union[Expression, _Mapping]] = ..., format: _Optional[_Union[Expression, _Mapping]] = ..., tz: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class RoundInterval(_message.Message):
            __slots__ = ("interval", "scale")
            INTERVAL_FIELD_NUMBER: _ClassVar[int]
            SCALE_FIELD_NUMBER: _ClassVar[int]
            interval: Expression
            scale: Expression
            def __init__(self, interval: _Optional[_Union[Expression, _Mapping]] = ..., scale: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ExtractTime(_message.Message):
            __slots__ = ("timestamp", "unit", "tz")
            TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
            UNIT_FIELD_NUMBER: _ClassVar[int]
            TZ_FIELD_NUMBER: _ClassVar[int]
            timestamp: Expression
            unit: Expression
            tz: Expression
            def __init__(self, timestamp: _Optional[_Union[Expression, _Mapping]] = ..., unit: _Optional[_Union[Expression, _Mapping]] = ..., tz: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class UrlEncode(_message.Message):
            __slots__ = ("string",)
            STRING_FIELD_NUMBER: _ClassVar[int]
            string: Expression
            def __init__(self, string: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class UrlDecode(_message.Message):
            __slots__ = ("string",)
            STRING_FIELD_NUMBER: _ClassVar[int]
            string: Expression
            def __init__(self, string: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Matches(_message.Message):
            __slots__ = ("string", "regexp")
            STRING_FIELD_NUMBER: _ClassVar[int]
            REGEXP_FIELD_NUMBER: _ClassVar[int]
            string: Expression
            regexp: Expression
            def __init__(self, string: _Optional[_Union[Expression, _Mapping]] = ..., regexp: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class RegexpSplitParts(_message.Message):
            __slots__ = ("string", "delimiter", "index")
            STRING_FIELD_NUMBER: _ClassVar[int]
            DELIMITER_FIELD_NUMBER: _ClassVar[int]
            INDEX_FIELD_NUMBER: _ClassVar[int]
            string: Expression
            delimiter: Expression
            index: Expression
            def __init__(self, string: _Optional[_Union[Expression, _Mapping]] = ..., delimiter: _Optional[_Union[Expression, _Mapping]] = ..., index: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class IpPrefix(_message.Message):
            __slots__ = ("ip", "subnet_size")
            IP_FIELD_NUMBER: _ClassVar[int]
            SUBNET_SIZE_FIELD_NUMBER: _ClassVar[int]
            ip: Expression
            subnet_size: Expression
            def __init__(self, ip: _Optional[_Union[Expression, _Mapping]] = ..., subnet_size: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class IpInSubnet(_message.Message):
            __slots__ = ("ip", "ip_prefix")
            IP_FIELD_NUMBER: _ClassVar[int]
            IP_PREFIX_FIELD_NUMBER: _ClassVar[int]
            ip: Expression
            ip_prefix: Expression
            def __init__(self, ip: _Optional[_Union[Expression, _Mapping]] = ..., ip_prefix: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class RandomUuid(_message.Message):
            __slots__ = ()
            def __init__(self) -> None: ...
        class IsUuid(_message.Message):
            __slots__ = ("uuid",)
            UUID_FIELD_NUMBER: _ClassVar[int]
            uuid: Expression
            def __init__(self, uuid: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class RecordLocation(_message.Message):
            __slots__ = ()
            def __init__(self) -> None: ...
        class ArrayLength(_message.Message):
            __slots__ = ("array",)
            ARRAY_FIELD_NUMBER: _ClassVar[int]
            array: Expression
            def __init__(self, array: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class IsEmpty(_message.Message):
            __slots__ = ("array",)
            ARRAY_FIELD_NUMBER: _ClassVar[int]
            array: Expression
            def __init__(self, array: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class Cardinality(_message.Message):
            __slots__ = ("array",)
            ARRAY_FIELD_NUMBER: _ClassVar[int]
            array: Expression
            def __init__(self, array: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ArrayContains(_message.Message):
            __slots__ = ("array", "element")
            ARRAY_FIELD_NUMBER: _ClassVar[int]
            ELEMENT_FIELD_NUMBER: _ClassVar[int]
            array: Expression
            element: Expression
            def __init__(self, array: _Optional[_Union[Expression, _Mapping]] = ..., element: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ArraySplit(_message.Message):
            __slots__ = ("string", "delimiter")
            STRING_FIELD_NUMBER: _ClassVar[int]
            DELIMITER_FIELD_NUMBER: _ClassVar[int]
            string: Expression
            delimiter: Expression
            def __init__(self, string: _Optional[_Union[Expression, _Mapping]] = ..., delimiter: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ArrayJoin(_message.Message):
            __slots__ = ("array", "delimiter")
            ARRAY_FIELD_NUMBER: _ClassVar[int]
            DELIMITER_FIELD_NUMBER: _ClassVar[int]
            array: Expression
            delimiter: Expression
            def __init__(self, array: _Optional[_Union[Expression, _Mapping]] = ..., delimiter: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ArrayConcat(_message.Message):
            __slots__ = ("array1", "array2")
            ARRAY1_FIELD_NUMBER: _ClassVar[int]
            ARRAY2_FIELD_NUMBER: _ClassVar[int]
            array1: Expression
            array2: Expression
            def __init__(self, array1: _Optional[_Union[Expression, _Mapping]] = ..., array2: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ArrayAppend(_message.Message):
            __slots__ = ("array", "element")
            ARRAY_FIELD_NUMBER: _ClassVar[int]
            ELEMENT_FIELD_NUMBER: _ClassVar[int]
            array: Expression
            element: Expression
            def __init__(self, array: _Optional[_Union[Expression, _Mapping]] = ..., element: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ArrayInsertAt(_message.Message):
            __slots__ = ("array", "position", "value")
            ARRAY_FIELD_NUMBER: _ClassVar[int]
            POSITION_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            array: Expression
            position: Expression
            value: Expression
            def __init__(self, array: _Optional[_Union[Expression, _Mapping]] = ..., position: _Optional[_Union[Expression, _Mapping]] = ..., value: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ArrayReplaceAt(_message.Message):
            __slots__ = ("array", "position", "value")
            ARRAY_FIELD_NUMBER: _ClassVar[int]
            POSITION_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            array: Expression
            position: Expression
            value: Expression
            def __init__(self, array: _Optional[_Union[Expression, _Mapping]] = ..., position: _Optional[_Union[Expression, _Mapping]] = ..., value: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ArrayReplaceAll(_message.Message):
            __slots__ = ("array", "value", "new_value")
            ARRAY_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            NEW_VALUE_FIELD_NUMBER: _ClassVar[int]
            array: Expression
            value: Expression
            new_value: Expression
            def __init__(self, array: _Optional[_Union[Expression, _Mapping]] = ..., value: _Optional[_Union[Expression, _Mapping]] = ..., new_value: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ArrayRemoveAt(_message.Message):
            __slots__ = ("array", "position")
            ARRAY_FIELD_NUMBER: _ClassVar[int]
            POSITION_FIELD_NUMBER: _ClassVar[int]
            array: Expression
            position: Expression
            def __init__(self, array: _Optional[_Union[Expression, _Mapping]] = ..., position: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class ArrayRemove(_message.Message):
            __slots__ = ("array", "element")
            ARRAY_FIELD_NUMBER: _ClassVar[int]
            ELEMENT_FIELD_NUMBER: _ClassVar[int]
            array: Expression
            element: Expression
            def __init__(self, array: _Optional[_Union[Expression, _Mapping]] = ..., element: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class SetUnion(_message.Message):
            __slots__ = ("array1", "array2")
            ARRAY1_FIELD_NUMBER: _ClassVar[int]
            ARRAY2_FIELD_NUMBER: _ClassVar[int]
            array1: Expression
            array2: Expression
            def __init__(self, array1: _Optional[_Union[Expression, _Mapping]] = ..., array2: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class SetIntersection(_message.Message):
            __slots__ = ("array1", "array2")
            ARRAY1_FIELD_NUMBER: _ClassVar[int]
            ARRAY2_FIELD_NUMBER: _ClassVar[int]
            array1: Expression
            array2: Expression
            def __init__(self, array1: _Optional[_Union[Expression, _Mapping]] = ..., array2: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class SetDiff(_message.Message):
            __slots__ = ("array1", "array2")
            ARRAY1_FIELD_NUMBER: _ClassVar[int]
            ARRAY2_FIELD_NUMBER: _ClassVar[int]
            array1: Expression
            array2: Expression
            def __init__(self, array1: _Optional[_Union[Expression, _Mapping]] = ..., array2: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class SetDiffSymmetric(_message.Message):
            __slots__ = ("array1", "array2")
            ARRAY1_FIELD_NUMBER: _ClassVar[int]
            ARRAY2_FIELD_NUMBER: _ClassVar[int]
            array1: Expression
            array2: Expression
            def __init__(self, array1: _Optional[_Union[Expression, _Mapping]] = ..., array2: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class IsSubset(_message.Message):
            __slots__ = ("array1", "array2")
            ARRAY1_FIELD_NUMBER: _ClassVar[int]
            ARRAY2_FIELD_NUMBER: _ClassVar[int]
            array1: Expression
            array2: Expression
            def __init__(self, array1: _Optional[_Union[Expression, _Mapping]] = ..., array2: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class IsSuperset(_message.Message):
            __slots__ = ("array1", "array2")
            ARRAY1_FIELD_NUMBER: _ClassVar[int]
            ARRAY2_FIELD_NUMBER: _ClassVar[int]
            array1: Expression
            array2: Expression
            def __init__(self, array1: _Optional[_Union[Expression, _Mapping]] = ..., array2: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        class SetEqualsTo(_message.Message):
            __slots__ = ("array1", "array2")
            ARRAY1_FIELD_NUMBER: _ClassVar[int]
            ARRAY2_FIELD_NUMBER: _ClassVar[int]
            array1: Expression
            array2: Expression
            def __init__(self, array1: _Optional[_Union[Expression, _Mapping]] = ..., array2: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
        FIRST_NON_NULL_FIELD_NUMBER: _ClassVar[int]
        IF_FIELD_NUMBER: _ClassVar[int]
        DECODE_BASE_64_FIELD_NUMBER: _ClassVar[int]
        ENCODE_BASE_64_FIELD_NUMBER: _ClassVar[int]
        LENGTH_FIELD_NUMBER: _ClassVar[int]
        SUBSTR_FIELD_NUMBER: _ClassVar[int]
        CONCAT_FIELD_NUMBER: _ClassVar[int]
        TRIM_FIELD_NUMBER: _ClassVar[int]
        RTRIM_FIELD_NUMBER: _ClassVar[int]
        LTRIM_FIELD_NUMBER: _ClassVar[int]
        TO_LOWER_CASE_FIELD_NUMBER: _ClassVar[int]
        TO_UPPER_CASE_FIELD_NUMBER: _ClassVar[int]
        ROUND_TIME_FIELD_NUMBER: _ClassVar[int]
        CONTAINS_FIELD_NUMBER: _ClassVar[int]
        INDEX_OF_FIELD_NUMBER: _ClassVar[int]
        STARTS_WITH_FIELD_NUMBER: _ClassVar[int]
        ENDS_WITH_FIELD_NUMBER: _ClassVar[int]
        ROUND_FIELD_NUMBER: _ClassVar[int]
        ABS_FIELD_NUMBER: _ClassVar[int]
        MIN_FIELD_NUMBER: _ClassVar[int]
        MAX_FIELD_NUMBER: _ClassVar[int]
        FLOOR_FIELD_NUMBER: _ClassVar[int]
        CEIL_FIELD_NUMBER: _ClassVar[int]
        POWER_FIELD_NUMBER: _ClassVar[int]
        LOG_FIELD_NUMBER: _ClassVar[int]
        LOG2_FIELD_NUMBER: _ClassVar[int]
        LN_FIELD_NUMBER: _ClassVar[int]
        TO_ISO_8601_DATE_TIME_FIELD_NUMBER: _ClassVar[int]
        FROM_UNIX_TIME_FIELD_NUMBER: _ClassVar[int]
        IN_FIELD_NUMBER: _ClassVar[int]
        PAD_FIELD_NUMBER: _ClassVar[int]
        PAD_LEFT_FIELD_NUMBER: _ClassVar[int]
        PAD_RIGHT_FIELD_NUMBER: _ClassVar[int]
        CHR_FIELD_NUMBER: _ClassVar[int]
        CODEPOINT_FIELD_NUMBER: _ClassVar[int]
        SPLIT_PARTS_FIELD_NUMBER: _ClassVar[int]
        FROM_BASE_FIELD_NUMBER: _ClassVar[int]
        TO_BASE_FIELD_NUMBER: _ClassVar[int]
        MOD_FIELD_NUMBER: _ClassVar[int]
        SQRT_FIELD_NUMBER: _ClassVar[int]
        E_FIELD_NUMBER: _ClassVar[int]
        PI_FIELD_NUMBER: _ClassVar[int]
        RANDOM_FIELD_NUMBER: _ClassVar[int]
        RANDOM_INT_FIELD_NUMBER: _ClassVar[int]
        URL_ENCODE_FIELD_NUMBER: _ClassVar[int]
        URL_DECODE_FIELD_NUMBER: _ClassVar[int]
        MATCHES_FIELD_NUMBER: _ClassVar[int]
        REGEXP_SPLIT_PARTS_FIELD_NUMBER: _ClassVar[int]
        IP_PREFIX_FIELD_NUMBER: _ClassVar[int]
        IP_IN_SUBNET_FIELD_NUMBER: _ClassVar[int]
        RANDOM_UUID_FIELD_NUMBER: _ClassVar[int]
        IS_UUID_FIELD_NUMBER: _ClassVar[int]
        TO_UNIX_TIME_FIELD_NUMBER: _ClassVar[int]
        NOW_FIELD_NUMBER: _ClassVar[int]
        RECORD_LOCATION_FIELD_NUMBER: _ClassVar[int]
        PARSE_INTERVAL_FIELD_NUMBER: _ClassVar[int]
        FORMAT_INTERVAL_FIELD_NUMBER: _ClassVar[int]
        TO_INTERVAL_FIELD_NUMBER: _ClassVar[int]
        PARSE_TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
        FORMAT_TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
        ROUND_INTERVAL_FIELD_NUMBER: _ClassVar[int]
        EXTRACT_TIME_FIELD_NUMBER: _ClassVar[int]
        ARRAY_LENGTH_FIELD_NUMBER: _ClassVar[int]
        ARRAY_IS_EMPTY_FIELD_NUMBER: _ClassVar[int]
        CARDINALITY_FIELD_NUMBER: _ClassVar[int]
        ARRAY_CONTAINS_FIELD_NUMBER: _ClassVar[int]
        ARRAY_SPLIT_FIELD_NUMBER: _ClassVar[int]
        ARRAY_JOIN_FIELD_NUMBER: _ClassVar[int]
        SET_UNION_FIELD_NUMBER: _ClassVar[int]
        SET_INTERSECTION_FIELD_NUMBER: _ClassVar[int]
        SET_DIFF_FIELD_NUMBER: _ClassVar[int]
        SET_DIFF_SYMMETRIC_FIELD_NUMBER: _ClassVar[int]
        IS_SUBSET_FIELD_NUMBER: _ClassVar[int]
        IS_SUPERSET_FIELD_NUMBER: _ClassVar[int]
        SET_EQUALS_TO_FIELD_NUMBER: _ClassVar[int]
        ARRAY_CONCAT_FIELD_NUMBER: _ClassVar[int]
        ARRAY_APPEND_FIELD_NUMBER: _ClassVar[int]
        ARRAY_INSERT_AT_FIELD_NUMBER: _ClassVar[int]
        ARRAY_REPLACE_AT_FIELD_NUMBER: _ClassVar[int]
        ARRAY_REPLACE_ALL_FIELD_NUMBER: _ClassVar[int]
        ARRAY_REMOVE_AT_FIELD_NUMBER: _ClassVar[int]
        ARRAY_REMOVE_FIELD_NUMBER: _ClassVar[int]
        first_non_null: Expression.FunctionCallV2.FirstNonNull
        decode_base_64: Expression.FunctionCallV2.DecodeBase64
        encode_base_64: Expression.FunctionCallV2.EncodeBase64
        length: Expression.FunctionCallV2.Length
        substr: Expression.FunctionCallV2.Substr
        concat: Expression.FunctionCallV2.Concat
        trim: Expression.FunctionCallV2.Trim
        rtrim: Expression.FunctionCallV2.Rtrim
        ltrim: Expression.FunctionCallV2.Ltrim
        to_lower_case: Expression.FunctionCallV2.ToLowerCase
        to_upper_case: Expression.FunctionCallV2.ToUpperCase
        round_time: Expression.FunctionCallV2.RoundTime
        contains: Expression.FunctionCallV2.Contains
        index_of: Expression.FunctionCallV2.IndexOf
        starts_with: Expression.FunctionCallV2.StartsWith
        ends_with: Expression.FunctionCallV2.EndsWith
        round: Expression.FunctionCallV2.Round
        abs: Expression.FunctionCallV2.Abs
        min: Expression.FunctionCallV2.Min
        max: Expression.FunctionCallV2.Max
        floor: Expression.FunctionCallV2.Floor
        ceil: Expression.FunctionCallV2.Ceil
        power: Expression.FunctionCallV2.Power
        log: Expression.FunctionCallV2.Log
        log2: Expression.FunctionCallV2.Log2
        ln: Expression.FunctionCallV2.Ln
        to_iso_8601_date_time: Expression.FunctionCallV2.ToIso8601DateTime
        from_unix_time: Expression.FunctionCallV2.FromUnixTime
        pad: Expression.FunctionCallV2.Pad
        pad_left: Expression.FunctionCallV2.PadLeft
        pad_right: Expression.FunctionCallV2.PadRight
        chr: Expression.FunctionCallV2.Chr
        codepoint: Expression.FunctionCallV2.Codepoint
        split_parts: Expression.FunctionCallV2.SplitParts
        from_base: Expression.FunctionCallV2.FromBase
        to_base: Expression.FunctionCallV2.ToBase
        mod: Expression.FunctionCallV2.Mod
        sqrt: Expression.FunctionCallV2.Sqrt
        e: Expression.FunctionCallV2.E
        pi: Expression.FunctionCallV2.Pi
        random: Expression.FunctionCallV2.Random
        random_int: Expression.FunctionCallV2.RandomInt
        url_encode: Expression.FunctionCallV2.UrlEncode
        url_decode: Expression.FunctionCallV2.UrlDecode
        matches: Expression.FunctionCallV2.Matches
        regexp_split_parts: Expression.FunctionCallV2.RegexpSplitParts
        ip_prefix: Expression.FunctionCallV2.IpPrefix
        ip_in_subnet: Expression.FunctionCallV2.IpInSubnet
        random_uuid: Expression.FunctionCallV2.RandomUuid
        is_uuid: Expression.FunctionCallV2.IsUuid
        to_unix_time: Expression.FunctionCallV2.ToUnixTime
        now: Expression.FunctionCallV2.Now
        record_location: Expression.FunctionCallV2.RecordLocation
        parse_interval: Expression.FunctionCallV2.ParseInterval
        format_interval: Expression.FunctionCallV2.FormatInterval
        to_interval: Expression.FunctionCallV2.ToInterval
        parse_timestamp: Expression.FunctionCallV2.ParseTimestamp
        format_timestamp: Expression.FunctionCallV2.FormatTimestamp
        round_interval: Expression.FunctionCallV2.RoundInterval
        extract_time: Expression.FunctionCallV2.ExtractTime
        array_length: Expression.FunctionCallV2.ArrayLength
        array_is_empty: Expression.FunctionCallV2.IsEmpty
        cardinality: Expression.FunctionCallV2.Cardinality
        array_contains: Expression.FunctionCallV2.ArrayContains
        array_split: Expression.FunctionCallV2.ArraySplit
        array_join: Expression.FunctionCallV2.ArrayJoin
        set_union: Expression.FunctionCallV2.SetUnion
        set_intersection: Expression.FunctionCallV2.SetIntersection
        set_diff: Expression.FunctionCallV2.SetDiff
        set_diff_symmetric: Expression.FunctionCallV2.SetDiffSymmetric
        is_subset: Expression.FunctionCallV2.IsSubset
        is_superset: Expression.FunctionCallV2.IsSuperset
        set_equals_to: Expression.FunctionCallV2.SetEqualsTo
        array_concat: Expression.FunctionCallV2.ArrayConcat
        array_append: Expression.FunctionCallV2.ArrayAppend
        array_insert_at: Expression.FunctionCallV2.ArrayInsertAt
        array_replace_at: Expression.FunctionCallV2.ArrayReplaceAt
        array_replace_all: Expression.FunctionCallV2.ArrayReplaceAll
        array_remove_at: Expression.FunctionCallV2.ArrayRemoveAt
        array_remove: Expression.FunctionCallV2.ArrayRemove
        def __init__(self, first_non_null: _Optional[_Union[Expression.FunctionCallV2.FirstNonNull, _Mapping]] = ..., decode_base_64: _Optional[_Union[Expression.FunctionCallV2.DecodeBase64, _Mapping]] = ..., encode_base_64: _Optional[_Union[Expression.FunctionCallV2.EncodeBase64, _Mapping]] = ..., length: _Optional[_Union[Expression.FunctionCallV2.Length, _Mapping]] = ..., substr: _Optional[_Union[Expression.FunctionCallV2.Substr, _Mapping]] = ..., concat: _Optional[_Union[Expression.FunctionCallV2.Concat, _Mapping]] = ..., trim: _Optional[_Union[Expression.FunctionCallV2.Trim, _Mapping]] = ..., rtrim: _Optional[_Union[Expression.FunctionCallV2.Rtrim, _Mapping]] = ..., ltrim: _Optional[_Union[Expression.FunctionCallV2.Ltrim, _Mapping]] = ..., to_lower_case: _Optional[_Union[Expression.FunctionCallV2.ToLowerCase, _Mapping]] = ..., to_upper_case: _Optional[_Union[Expression.FunctionCallV2.ToUpperCase, _Mapping]] = ..., round_time: _Optional[_Union[Expression.FunctionCallV2.RoundTime, _Mapping]] = ..., contains: _Optional[_Union[Expression.FunctionCallV2.Contains, _Mapping]] = ..., index_of: _Optional[_Union[Expression.FunctionCallV2.IndexOf, _Mapping]] = ..., starts_with: _Optional[_Union[Expression.FunctionCallV2.StartsWith, _Mapping]] = ..., ends_with: _Optional[_Union[Expression.FunctionCallV2.EndsWith, _Mapping]] = ..., round: _Optional[_Union[Expression.FunctionCallV2.Round, _Mapping]] = ..., abs: _Optional[_Union[Expression.FunctionCallV2.Abs, _Mapping]] = ..., min: _Optional[_Union[Expression.FunctionCallV2.Min, _Mapping]] = ..., max: _Optional[_Union[Expression.FunctionCallV2.Max, _Mapping]] = ..., floor: _Optional[_Union[Expression.FunctionCallV2.Floor, _Mapping]] = ..., ceil: _Optional[_Union[Expression.FunctionCallV2.Ceil, _Mapping]] = ..., power: _Optional[_Union[Expression.FunctionCallV2.Power, _Mapping]] = ..., log: _Optional[_Union[Expression.FunctionCallV2.Log, _Mapping]] = ..., log2: _Optional[_Union[Expression.FunctionCallV2.Log2, _Mapping]] = ..., ln: _Optional[_Union[Expression.FunctionCallV2.Ln, _Mapping]] = ..., to_iso_8601_date_time: _Optional[_Union[Expression.FunctionCallV2.ToIso8601DateTime, _Mapping]] = ..., from_unix_time: _Optional[_Union[Expression.FunctionCallV2.FromUnixTime, _Mapping]] = ..., pad: _Optional[_Union[Expression.FunctionCallV2.Pad, _Mapping]] = ..., pad_left: _Optional[_Union[Expression.FunctionCallV2.PadLeft, _Mapping]] = ..., pad_right: _Optional[_Union[Expression.FunctionCallV2.PadRight, _Mapping]] = ..., chr: _Optional[_Union[Expression.FunctionCallV2.Chr, _Mapping]] = ..., codepoint: _Optional[_Union[Expression.FunctionCallV2.Codepoint, _Mapping]] = ..., split_parts: _Optional[_Union[Expression.FunctionCallV2.SplitParts, _Mapping]] = ..., from_base: _Optional[_Union[Expression.FunctionCallV2.FromBase, _Mapping]] = ..., to_base: _Optional[_Union[Expression.FunctionCallV2.ToBase, _Mapping]] = ..., mod: _Optional[_Union[Expression.FunctionCallV2.Mod, _Mapping]] = ..., sqrt: _Optional[_Union[Expression.FunctionCallV2.Sqrt, _Mapping]] = ..., e: _Optional[_Union[Expression.FunctionCallV2.E, _Mapping]] = ..., pi: _Optional[_Union[Expression.FunctionCallV2.Pi, _Mapping]] = ..., random: _Optional[_Union[Expression.FunctionCallV2.Random, _Mapping]] = ..., random_int: _Optional[_Union[Expression.FunctionCallV2.RandomInt, _Mapping]] = ..., url_encode: _Optional[_Union[Expression.FunctionCallV2.UrlEncode, _Mapping]] = ..., url_decode: _Optional[_Union[Expression.FunctionCallV2.UrlDecode, _Mapping]] = ..., matches: _Optional[_Union[Expression.FunctionCallV2.Matches, _Mapping]] = ..., regexp_split_parts: _Optional[_Union[Expression.FunctionCallV2.RegexpSplitParts, _Mapping]] = ..., ip_prefix: _Optional[_Union[Expression.FunctionCallV2.IpPrefix, _Mapping]] = ..., ip_in_subnet: _Optional[_Union[Expression.FunctionCallV2.IpInSubnet, _Mapping]] = ..., random_uuid: _Optional[_Union[Expression.FunctionCallV2.RandomUuid, _Mapping]] = ..., is_uuid: _Optional[_Union[Expression.FunctionCallV2.IsUuid, _Mapping]] = ..., to_unix_time: _Optional[_Union[Expression.FunctionCallV2.ToUnixTime, _Mapping]] = ..., now: _Optional[_Union[Expression.FunctionCallV2.Now, _Mapping]] = ..., record_location: _Optional[_Union[Expression.FunctionCallV2.RecordLocation, _Mapping]] = ..., parse_interval: _Optional[_Union[Expression.FunctionCallV2.ParseInterval, _Mapping]] = ..., format_interval: _Optional[_Union[Expression.FunctionCallV2.FormatInterval, _Mapping]] = ..., to_interval: _Optional[_Union[Expression.FunctionCallV2.ToInterval, _Mapping]] = ..., parse_timestamp: _Optional[_Union[Expression.FunctionCallV2.ParseTimestamp, _Mapping]] = ..., format_timestamp: _Optional[_Union[Expression.FunctionCallV2.FormatTimestamp, _Mapping]] = ..., round_interval: _Optional[_Union[Expression.FunctionCallV2.RoundInterval, _Mapping]] = ..., extract_time: _Optional[_Union[Expression.FunctionCallV2.ExtractTime, _Mapping]] = ..., array_length: _Optional[_Union[Expression.FunctionCallV2.ArrayLength, _Mapping]] = ..., array_is_empty: _Optional[_Union[Expression.FunctionCallV2.IsEmpty, _Mapping]] = ..., cardinality: _Optional[_Union[Expression.FunctionCallV2.Cardinality, _Mapping]] = ..., array_contains: _Optional[_Union[Expression.FunctionCallV2.ArrayContains, _Mapping]] = ..., array_split: _Optional[_Union[Expression.FunctionCallV2.ArraySplit, _Mapping]] = ..., array_join: _Optional[_Union[Expression.FunctionCallV2.ArrayJoin, _Mapping]] = ..., set_union: _Optional[_Union[Expression.FunctionCallV2.SetUnion, _Mapping]] = ..., set_intersection: _Optional[_Union[Expression.FunctionCallV2.SetIntersection, _Mapping]] = ..., set_diff: _Optional[_Union[Expression.FunctionCallV2.SetDiff, _Mapping]] = ..., set_diff_symmetric: _Optional[_Union[Expression.FunctionCallV2.SetDiffSymmetric, _Mapping]] = ..., is_subset: _Optional[_Union[Expression.FunctionCallV2.IsSubset, _Mapping]] = ..., is_superset: _Optional[_Union[Expression.FunctionCallV2.IsSuperset, _Mapping]] = ..., set_equals_to: _Optional[_Union[Expression.FunctionCallV2.SetEqualsTo, _Mapping]] = ..., array_concat: _Optional[_Union[Expression.FunctionCallV2.ArrayConcat, _Mapping]] = ..., array_append: _Optional[_Union[Expression.FunctionCallV2.ArrayAppend, _Mapping]] = ..., array_insert_at: _Optional[_Union[Expression.FunctionCallV2.ArrayInsertAt, _Mapping]] = ..., array_replace_at: _Optional[_Union[Expression.FunctionCallV2.ArrayReplaceAt, _Mapping]] = ..., array_replace_all: _Optional[_Union[Expression.FunctionCallV2.ArrayReplaceAll, _Mapping]] = ..., array_remove_at: _Optional[_Union[Expression.FunctionCallV2.ArrayRemoveAt, _Mapping]] = ..., array_remove: _Optional[_Union[Expression.FunctionCallV2.ArrayRemove, _Mapping]] = ..., **kwargs) -> None: ...
    class Cast(_message.Message):
        __slots__ = ("expression", "datatype")
        EXPRESSION_FIELD_NUMBER: _ClassVar[int]
        DATATYPE_FIELD_NUMBER: _ClassVar[int]
        expression: Expression
        datatype: _types_pb2.Datatype
        def __init__(self, expression: _Optional[_Union[Expression, _Mapping]] = ..., datatype: _Optional[_Union[_types_pb2.Datatype, _Mapping]] = ...) -> None: ...
    class MatchSubtree(_message.Message):
        __slots__ = ("path", "term")
        PATH_FIELD_NUMBER: _ClassVar[int]
        TERM_FIELD_NUMBER: _ClassVar[int]
        path: Expression.Keypath
        term: Expression
        def __init__(self, path: _Optional[_Union[Expression.Keypath, _Mapping]] = ..., term: _Optional[_Union[Expression, _Mapping]] = ...) -> None: ...
    class EnumValue(_message.Message):
        __slots__ = ("name", "value")
        NAME_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        name: str
        value: str
        def __init__(self, name: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    NULL_E_FIELD_NUMBER: _ClassVar[int]
    KEYPATH_FIELD_NUMBER: _ClassVar[int]
    BOOLEAN_FIELD_NUMBER: _ClassVar[int]
    INTEGER_FIELD_NUMBER: _ClassVar[int]
    STR_FIELD_NUMBER: _ClassVar[int]
    REGEX_FIELD_NUMBER: _ClassVar[int]
    INFIX_OP_FIELD_NUMBER: _ClassVar[int]
    NOT_FIELD_NUMBER: _ClassVar[int]
    FUNCTION_CALL_FIELD_NUMBER: _ClassVar[int]
    CAST_FIELD_NUMBER: _ClassVar[int]
    MATCH_FIELD_NUMBER: _ClassVar[int]
    LONG_FIELD_NUMBER: _ClassVar[int]
    DOUBLE_FIELD_NUMBER: _ClassVar[int]
    TIME_UNIT_FIELD_NUMBER: _ClassVar[int]
    INTERVAL_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    DATE_UNIT_FIELD_NUMBER: _ClassVar[int]
    ARRAY_FIELD_NUMBER: _ClassVar[int]
    ARRAY_SLICE_FIELD_NUMBER: _ClassVar[int]
    ENUM_VALUE_FIELD_NUMBER: _ClassVar[int]
    ARRAY_ELEMENT_FIELD_NUMBER: _ClassVar[int]
    DATATYPE_FIELD_NUMBER: _ClassVar[int]
    null_e: Expression.Null
    keypath: Expression.Keypath
    boolean: _wrappers_pb2.BoolValue
    integer: _wrappers_pb2.Int32Value
    str: _wrappers_pb2.StringValue
    regex: Expression.Regex
    infix_op: Expression.InfixOp
    function_call: Expression.FunctionCall
    cast: Expression.Cast
    match: Expression.MatchSubtree
    long: _wrappers_pb2.Int64Value
    double: _wrappers_pb2.DoubleValue
    time_unit: Expression.TimeUnit
    interval: Expression.Interval
    timestamp: Expression.Timestamp
    date_unit: Expression.DateUnit
    array: Expression.Array
    array_slice: Expression.ArraySlice
    enum_value: Expression.EnumValue
    array_element: Expression.ArrayElement
    datatype: _types_pb2.Datatype
    def __init__(self, null_e: _Optional[_Union[Expression.Null, _Mapping]] = ..., keypath: _Optional[_Union[Expression.Keypath, _Mapping]] = ..., boolean: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., integer: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., str: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., regex: _Optional[_Union[Expression.Regex, _Mapping]] = ..., infix_op: _Optional[_Union[Expression.InfixOp, _Mapping]] = ..., function_call: _Optional[_Union[Expression.FunctionCall, _Mapping]] = ..., cast: _Optional[_Union[Expression.Cast, _Mapping]] = ..., match: _Optional[_Union[Expression.MatchSubtree, _Mapping]] = ..., long: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., double: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ..., time_unit: _Optional[_Union[Expression.TimeUnit, str]] = ..., interval: _Optional[_Union[Expression.Interval, _Mapping]] = ..., timestamp: _Optional[_Union[Expression.Timestamp, _Mapping]] = ..., date_unit: _Optional[_Union[Expression.DateUnit, str]] = ..., array: _Optional[_Union[Expression.Array, _Mapping]] = ..., array_slice: _Optional[_Union[Expression.ArraySlice, _Mapping]] = ..., enum_value: _Optional[_Union[Expression.EnumValue, _Mapping]] = ..., array_element: _Optional[_Union[Expression.ArrayElement, _Mapping]] = ..., datatype: _Optional[_Union[_types_pb2.Datatype, _Mapping]] = ..., **kwargs) -> None: ...
