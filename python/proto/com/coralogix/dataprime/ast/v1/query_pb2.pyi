from com.coralogix.dataprime.ast.v1 import expression_pb2 as _expression_pb2
from com.coralogix.dataprime.ast.v1 import misc_pb2 as _misc_pb2
from com.coralogix.dataprime.ast.v1 import types_pb2 as _types_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Query(_message.Message):
    __slots__ = ("source", "operators")
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    OPERATORS_FIELD_NUMBER: _ClassVar[int]
    source: Source
    operators: _containers.RepeatedCompositeFieldContainer[Operator]
    def __init__(self, source: _Optional[_Union[Source, _Mapping]] = ..., operators: _Optional[_Iterable[_Union[Operator, _Mapping]]] = ...) -> None: ...

class Source(_message.Message):
    __slots__ = ("logs", "spans", "custom_dataset", "custom_enrichment", "schema")
    LOGS_FIELD_NUMBER: _ClassVar[int]
    SPANS_FIELD_NUMBER: _ClassVar[int]
    CUSTOM_DATASET_FIELD_NUMBER: _ClassVar[int]
    CUSTOM_ENRICHMENT_FIELD_NUMBER: _ClassVar[int]
    SCHEMA_FIELD_NUMBER: _ClassVar[int]
    logs: Logs
    spans: Spans
    custom_dataset: CustomDataset
    custom_enrichment: CustomEnrichment
    schema: Schema
    def __init__(self, logs: _Optional[_Union[Logs, _Mapping]] = ..., spans: _Optional[_Union[Spans, _Mapping]] = ..., custom_dataset: _Optional[_Union[CustomDataset, _Mapping]] = ..., custom_enrichment: _Optional[_Union[CustomEnrichment, _Mapping]] = ..., schema: _Optional[_Union[Schema, _Mapping]] = ...) -> None: ...

class Logs(_message.Message):
    __slots__ = ("team_id",)
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    team_id: str
    def __init__(self, team_id: _Optional[str] = ...) -> None: ...

class Spans(_message.Message):
    __slots__ = ("team_id",)
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    team_id: str
    def __init__(self, team_id: _Optional[str] = ...) -> None: ...

class CustomDataset(_message.Message):
    __slots__ = ("name", "team_id")
    NAME_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    team_id: str
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., team_id: _Optional[str] = ...) -> None: ...

class CustomEnrichment(_message.Message):
    __slots__ = ("name", "team_id")
    NAME_FIELD_NUMBER: _ClassVar[int]
    TEAM_ID_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    team_id: str
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., team_id: _Optional[str] = ...) -> None: ...

class Schema(_message.Message):
    __slots__ = ("name",)
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class Operator(_message.Message):
    __slots__ = ("filter", "create", "move", "remove", "replace", "redact", "choose", "convert", "extract", "limit", "order_by", "group_by", "union", "enrich")
    class Filter(_message.Message):
        __slots__ = ("condition",)
        CONDITION_FIELD_NUMBER: _ClassVar[int]
        condition: _expression_pb2.Expression
        def __init__(self, condition: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ...) -> None: ...
    class Create(_message.Message):
        __slots__ = ("path", "expression")
        PATH_FIELD_NUMBER: _ClassVar[int]
        EXPRESSION_FIELD_NUMBER: _ClassVar[int]
        path: _expression_pb2.Expression.Keypath
        expression: _expression_pb2.Expression
        def __init__(self, path: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ..., expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ...) -> None: ...
    class Move(_message.Message):
        __slots__ = ("to",)
        FROM_FIELD_NUMBER: _ClassVar[int]
        TO_FIELD_NUMBER: _ClassVar[int]
        to: _expression_pb2.Expression.Keypath
        def __init__(self, to: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ..., **kwargs) -> None: ...
    class Remove(_message.Message):
        __slots__ = ("targets",)
        TARGETS_FIELD_NUMBER: _ClassVar[int]
        targets: _containers.RepeatedCompositeFieldContainer[_expression_pb2.Expression.Keypath]
        def __init__(self, targets: _Optional[_Iterable[_Union[_expression_pb2.Expression.Keypath, _Mapping]]] = ...) -> None: ...
    class Replace(_message.Message):
        __slots__ = ("path", "expression")
        PATH_FIELD_NUMBER: _ClassVar[int]
        EXPRESSION_FIELD_NUMBER: _ClassVar[int]
        path: _expression_pb2.Expression.Keypath
        expression: _expression_pb2.Expression
        def __init__(self, path: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ..., expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ...) -> None: ...
    class Redact(_message.Message):
        __slots__ = ("path", "match", "replacement")
        PATH_FIELD_NUMBER: _ClassVar[int]
        MATCH_FIELD_NUMBER: _ClassVar[int]
        REPLACEMENT_FIELD_NUMBER: _ClassVar[int]
        path: _expression_pb2.Expression.Keypath
        match: _expression_pb2.Expression.Match
        replacement: _expression_pb2.Expression
        def __init__(self, path: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ..., match: _Optional[_Union[_expression_pb2.Expression.Match, _Mapping]] = ..., replacement: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ...) -> None: ...
    class Choose(_message.Message):
        __slots__ = ("fields",)
        class Chosen(_message.Message):
            __slots__ = ("expression",)
            EXPRESSION_FIELD_NUMBER: _ClassVar[int]
            AS_FIELD_NUMBER: _ClassVar[int]
            expression: _expression_pb2.Expression
            def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., **kwargs) -> None: ...
        FIELDS_FIELD_NUMBER: _ClassVar[int]
        fields: _containers.RepeatedCompositeFieldContainer[Operator.Choose.Chosen]
        def __init__(self, fields: _Optional[_Iterable[_Union[Operator.Choose.Chosen, _Mapping]]] = ...) -> None: ...
    class Convert(_message.Message):
        __slots__ = ("conversions",)
        CONVERSIONS_FIELD_NUMBER: _ClassVar[int]
        conversions: _containers.RepeatedCompositeFieldContainer[_misc_pb2.KeypathType]
        def __init__(self, conversions: _Optional[_Iterable[_Union[_misc_pb2.KeypathType, _Mapping]]] = ...) -> None: ...
    class Extract(_message.Message):
        __slots__ = ("expression", "into", "datatypes", "regex", "key_value", "delimited", "json", "multi_regex")
        class ExtractRegex(_message.Message):
            __slots__ = ("regex",)
            REGEX_FIELD_NUMBER: _ClassVar[int]
            regex: _expression_pb2.Expression.Regex
            def __init__(self, regex: _Optional[_Union[_expression_pb2.Expression.Regex, _Mapping]] = ...) -> None: ...
        class ExtractMultiRegex(_message.Message):
            __slots__ = ("regex",)
            REGEX_FIELD_NUMBER: _ClassVar[int]
            regex: _expression_pb2.Expression.Regex
            def __init__(self, regex: _Optional[_Union[_expression_pb2.Expression.Regex, _Mapping]] = ...) -> None: ...
        class ExtractKeyValue(_message.Message):
            __slots__ = ("pair_delimiter", "key_delimiter", "quote_char_value")
            PAIR_DELIMITER_FIELD_NUMBER: _ClassVar[int]
            KEY_DELIMITER_FIELD_NUMBER: _ClassVar[int]
            QUOTE_CHAR_VALUE_FIELD_NUMBER: _ClassVar[int]
            pair_delimiter: _wrappers_pb2.StringValue
            key_delimiter: _wrappers_pb2.StringValue
            quote_char_value: _wrappers_pb2.StringValue
            def __init__(self, pair_delimiter: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., key_delimiter: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., quote_char_value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
        class ExtractDelimited(_message.Message):
            __slots__ = ("delimiter", "type")
            DELIMITER_FIELD_NUMBER: _ClassVar[int]
            TYPE_FIELD_NUMBER: _ClassVar[int]
            delimiter: _wrappers_pb2.StringValue
            type: _types_pb2.Datatype.PrimitiveType
            def __init__(self, delimiter: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., type: _Optional[_Union[_types_pb2.Datatype.PrimitiveType, str]] = ...) -> None: ...
        class ExtractJson(_message.Message):
            __slots__ = ("max_unescape_count",)
            MAX_UNESCAPE_COUNT_FIELD_NUMBER: _ClassVar[int]
            max_unescape_count: _wrappers_pb2.Int32Value
            def __init__(self, max_unescape_count: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
        EXPRESSION_FIELD_NUMBER: _ClassVar[int]
        INTO_FIELD_NUMBER: _ClassVar[int]
        DATATYPES_FIELD_NUMBER: _ClassVar[int]
        REGEX_FIELD_NUMBER: _ClassVar[int]
        KEY_VALUE_FIELD_NUMBER: _ClassVar[int]
        DELIMITED_FIELD_NUMBER: _ClassVar[int]
        JSON_FIELD_NUMBER: _ClassVar[int]
        MULTI_REGEX_FIELD_NUMBER: _ClassVar[int]
        expression: _expression_pb2.Expression
        into: _expression_pb2.Expression.Keypath
        datatypes: _containers.RepeatedCompositeFieldContainer[_misc_pb2.KeypathType]
        regex: Operator.Extract.ExtractRegex
        key_value: Operator.Extract.ExtractKeyValue
        delimited: Operator.Extract.ExtractDelimited
        json: Operator.Extract.ExtractJson
        multi_regex: Operator.Extract.ExtractMultiRegex
        def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., into: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ..., datatypes: _Optional[_Iterable[_Union[_misc_pb2.KeypathType, _Mapping]]] = ..., regex: _Optional[_Union[Operator.Extract.ExtractRegex, _Mapping]] = ..., key_value: _Optional[_Union[Operator.Extract.ExtractKeyValue, _Mapping]] = ..., delimited: _Optional[_Union[Operator.Extract.ExtractDelimited, _Mapping]] = ..., json: _Optional[_Union[Operator.Extract.ExtractJson, _Mapping]] = ..., multi_regex: _Optional[_Union[Operator.Extract.ExtractMultiRegex, _Mapping]] = ...) -> None: ...
    class Limit(_message.Message):
        __slots__ = ("event_count",)
        EVENT_COUNT_FIELD_NUMBER: _ClassVar[int]
        event_count: _wrappers_pb2.Int32Value
        def __init__(self, event_count: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    class OrderBy(_message.Message):
        __slots__ = ("expressions",)
        class OrderByExpression(_message.Message):
            __slots__ = ("expression", "order")
            class Order(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
                __slots__ = ()
                ORDER_UNSPECIFIED: _ClassVar[Operator.OrderBy.OrderByExpression.Order]
                ORDER_ASCENDING: _ClassVar[Operator.OrderBy.OrderByExpression.Order]
                ORDER_DESCENDING: _ClassVar[Operator.OrderBy.OrderByExpression.Order]
            ORDER_UNSPECIFIED: Operator.OrderBy.OrderByExpression.Order
            ORDER_ASCENDING: Operator.OrderBy.OrderByExpression.Order
            ORDER_DESCENDING: Operator.OrderBy.OrderByExpression.Order
            EXPRESSION_FIELD_NUMBER: _ClassVar[int]
            ORDER_FIELD_NUMBER: _ClassVar[int]
            expression: _expression_pb2.Expression
            order: Operator.OrderBy.OrderByExpression.Order
            def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., order: _Optional[_Union[Operator.OrderBy.OrderByExpression.Order, str]] = ...) -> None: ...
        EXPRESSIONS_FIELD_NUMBER: _ClassVar[int]
        expressions: _containers.RepeatedCompositeFieldContainer[Operator.OrderBy.OrderByExpression]
        def __init__(self, expressions: _Optional[_Iterable[_Union[Operator.OrderBy.OrderByExpression, _Mapping]]] = ...) -> None: ...
    class GroupBy(_message.Message):
        __slots__ = ("functions", "groupings")
        class GroupByFunction(_message.Message):
            __slots__ = ("count", "distinct_count", "sum", "min", "max", "avg", "stddev", "sample_stddev", "variance", "sample_variance", "percentile", "any_value", "approx_count_distinct", "collect")
            class Count(_message.Message):
                __slots__ = ("expression", "alias")
                EXPRESSION_FIELD_NUMBER: _ClassVar[int]
                ALIAS_FIELD_NUMBER: _ClassVar[int]
                expression: _expression_pb2.Expression
                alias: _expression_pb2.Expression.Keypath
                def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., alias: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ...) -> None: ...
            class DistinctCount(_message.Message):
                __slots__ = ("expression", "alias")
                EXPRESSION_FIELD_NUMBER: _ClassVar[int]
                ALIAS_FIELD_NUMBER: _ClassVar[int]
                expression: _expression_pb2.Expression
                alias: _expression_pb2.Expression.Keypath
                def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., alias: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ...) -> None: ...
            class Sum(_message.Message):
                __slots__ = ("expression", "alias")
                EXPRESSION_FIELD_NUMBER: _ClassVar[int]
                ALIAS_FIELD_NUMBER: _ClassVar[int]
                expression: _expression_pb2.Expression
                alias: _expression_pb2.Expression.Keypath
                def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., alias: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ...) -> None: ...
            class Min(_message.Message):
                __slots__ = ("expression", "alias")
                EXPRESSION_FIELD_NUMBER: _ClassVar[int]
                ALIAS_FIELD_NUMBER: _ClassVar[int]
                expression: _expression_pb2.Expression
                alias: _expression_pb2.Expression.Keypath
                def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., alias: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ...) -> None: ...
            class Max(_message.Message):
                __slots__ = ("expression", "alias")
                EXPRESSION_FIELD_NUMBER: _ClassVar[int]
                ALIAS_FIELD_NUMBER: _ClassVar[int]
                expression: _expression_pb2.Expression
                alias: _expression_pb2.Expression.Keypath
                def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., alias: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ...) -> None: ...
            class Avg(_message.Message):
                __slots__ = ("expression", "alias")
                EXPRESSION_FIELD_NUMBER: _ClassVar[int]
                ALIAS_FIELD_NUMBER: _ClassVar[int]
                expression: _expression_pb2.Expression
                alias: _expression_pb2.Expression.Keypath
                def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., alias: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ...) -> None: ...
            class StdDev(_message.Message):
                __slots__ = ("expression", "alias")
                EXPRESSION_FIELD_NUMBER: _ClassVar[int]
                ALIAS_FIELD_NUMBER: _ClassVar[int]
                expression: _expression_pb2.Expression
                alias: _expression_pb2.Expression.Keypath
                def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., alias: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ...) -> None: ...
            class SampleStdDev(_message.Message):
                __slots__ = ("expression", "alias")
                EXPRESSION_FIELD_NUMBER: _ClassVar[int]
                ALIAS_FIELD_NUMBER: _ClassVar[int]
                expression: _expression_pb2.Expression
                alias: _expression_pb2.Expression.Keypath
                def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., alias: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ...) -> None: ...
            class Variance(_message.Message):
                __slots__ = ("expression", "alias")
                EXPRESSION_FIELD_NUMBER: _ClassVar[int]
                ALIAS_FIELD_NUMBER: _ClassVar[int]
                expression: _expression_pb2.Expression
                alias: _expression_pb2.Expression.Keypath
                def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., alias: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ...) -> None: ...
            class SampleVariance(_message.Message):
                __slots__ = ("expression", "alias")
                EXPRESSION_FIELD_NUMBER: _ClassVar[int]
                ALIAS_FIELD_NUMBER: _ClassVar[int]
                expression: _expression_pb2.Expression
                alias: _expression_pb2.Expression.Keypath
                def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., alias: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ...) -> None: ...
            class Percentile(_message.Message):
                __slots__ = ("percentile", "expression", "error_threshold", "alias")
                PERCENTILE_FIELD_NUMBER: _ClassVar[int]
                EXPRESSION_FIELD_NUMBER: _ClassVar[int]
                ERROR_THRESHOLD_FIELD_NUMBER: _ClassVar[int]
                ALIAS_FIELD_NUMBER: _ClassVar[int]
                percentile: _expression_pb2.Expression
                expression: _expression_pb2.Expression
                error_threshold: _expression_pb2.Expression
                alias: _expression_pb2.Expression.Keypath
                def __init__(self, percentile: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., error_threshold: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., alias: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ...) -> None: ...
            class AnyValue(_message.Message):
                __slots__ = ("expression", "alias")
                EXPRESSION_FIELD_NUMBER: _ClassVar[int]
                ALIAS_FIELD_NUMBER: _ClassVar[int]
                expression: _expression_pb2.Expression
                alias: _expression_pb2.Expression.Keypath
                def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., alias: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ...) -> None: ...
            class ApproxCountDistinct(_message.Message):
                __slots__ = ("value", "values", "alias")
                VALUE_FIELD_NUMBER: _ClassVar[int]
                VALUES_FIELD_NUMBER: _ClassVar[int]
                ALIAS_FIELD_NUMBER: _ClassVar[int]
                value: _expression_pb2.Expression
                values: _containers.RepeatedCompositeFieldContainer[_expression_pb2.Expression]
                alias: _expression_pb2.Expression.Keypath
                def __init__(self, value: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., values: _Optional[_Iterable[_Union[_expression_pb2.Expression, _Mapping]]] = ..., alias: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ...) -> None: ...
            class Collect(_message.Message):
                __slots__ = ("expression", "order_by", "alias", "distinct")
                EXPRESSION_FIELD_NUMBER: _ClassVar[int]
                ORDER_BY_FIELD_NUMBER: _ClassVar[int]
                ALIAS_FIELD_NUMBER: _ClassVar[int]
                DISTINCT_FIELD_NUMBER: _ClassVar[int]
                expression: _expression_pb2.Expression
                order_by: _containers.RepeatedCompositeFieldContainer[_expression_pb2.Expression]
                alias: _expression_pb2.Expression.Keypath
                distinct: bool
                def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., order_by: _Optional[_Iterable[_Union[_expression_pb2.Expression, _Mapping]]] = ..., alias: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ..., distinct: bool = ...) -> None: ...
            COUNT_FIELD_NUMBER: _ClassVar[int]
            DISTINCT_COUNT_FIELD_NUMBER: _ClassVar[int]
            SUM_FIELD_NUMBER: _ClassVar[int]
            MIN_FIELD_NUMBER: _ClassVar[int]
            MAX_FIELD_NUMBER: _ClassVar[int]
            AVG_FIELD_NUMBER: _ClassVar[int]
            STDDEV_FIELD_NUMBER: _ClassVar[int]
            SAMPLE_STDDEV_FIELD_NUMBER: _ClassVar[int]
            VARIANCE_FIELD_NUMBER: _ClassVar[int]
            SAMPLE_VARIANCE_FIELD_NUMBER: _ClassVar[int]
            PERCENTILE_FIELD_NUMBER: _ClassVar[int]
            ANY_VALUE_FIELD_NUMBER: _ClassVar[int]
            APPROX_COUNT_DISTINCT_FIELD_NUMBER: _ClassVar[int]
            COLLECT_FIELD_NUMBER: _ClassVar[int]
            count: Operator.GroupBy.GroupByFunction.Count
            distinct_count: Operator.GroupBy.GroupByFunction.DistinctCount
            sum: Operator.GroupBy.GroupByFunction.Sum
            min: Operator.GroupBy.GroupByFunction.Min
            max: Operator.GroupBy.GroupByFunction.Max
            avg: Operator.GroupBy.GroupByFunction.Avg
            stddev: Operator.GroupBy.GroupByFunction.StdDev
            sample_stddev: Operator.GroupBy.GroupByFunction.SampleStdDev
            variance: Operator.GroupBy.GroupByFunction.Variance
            sample_variance: Operator.GroupBy.GroupByFunction.SampleVariance
            percentile: Operator.GroupBy.GroupByFunction.Percentile
            any_value: Operator.GroupBy.GroupByFunction.AnyValue
            approx_count_distinct: Operator.GroupBy.GroupByFunction.ApproxCountDistinct
            collect: Operator.GroupBy.GroupByFunction.Collect
            def __init__(self, count: _Optional[_Union[Operator.GroupBy.GroupByFunction.Count, _Mapping]] = ..., distinct_count: _Optional[_Union[Operator.GroupBy.GroupByFunction.DistinctCount, _Mapping]] = ..., sum: _Optional[_Union[Operator.GroupBy.GroupByFunction.Sum, _Mapping]] = ..., min: _Optional[_Union[Operator.GroupBy.GroupByFunction.Min, _Mapping]] = ..., max: _Optional[_Union[Operator.GroupBy.GroupByFunction.Max, _Mapping]] = ..., avg: _Optional[_Union[Operator.GroupBy.GroupByFunction.Avg, _Mapping]] = ..., stddev: _Optional[_Union[Operator.GroupBy.GroupByFunction.StdDev, _Mapping]] = ..., sample_stddev: _Optional[_Union[Operator.GroupBy.GroupByFunction.SampleStdDev, _Mapping]] = ..., variance: _Optional[_Union[Operator.GroupBy.GroupByFunction.Variance, _Mapping]] = ..., sample_variance: _Optional[_Union[Operator.GroupBy.GroupByFunction.SampleVariance, _Mapping]] = ..., percentile: _Optional[_Union[Operator.GroupBy.GroupByFunction.Percentile, _Mapping]] = ..., any_value: _Optional[_Union[Operator.GroupBy.GroupByFunction.AnyValue, _Mapping]] = ..., approx_count_distinct: _Optional[_Union[Operator.GroupBy.GroupByFunction.ApproxCountDistinct, _Mapping]] = ..., collect: _Optional[_Union[Operator.GroupBy.GroupByFunction.Collect, _Mapping]] = ...) -> None: ...
        class GroupingSet(_message.Message):
            __slots__ = ("fields",)
            FIELDS_FIELD_NUMBER: _ClassVar[int]
            fields: _containers.RepeatedCompositeFieldContainer[Operator.GroupBy.GroupingSetField]
            def __init__(self, fields: _Optional[_Iterable[_Union[Operator.GroupBy.GroupingSetField, _Mapping]]] = ...) -> None: ...
        class GroupingSetField(_message.Message):
            __slots__ = ("expression", "alias")
            EXPRESSION_FIELD_NUMBER: _ClassVar[int]
            ALIAS_FIELD_NUMBER: _ClassVar[int]
            expression: _expression_pb2.Expression
            alias: _expression_pb2.Expression.Keypath
            def __init__(self, expression: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., alias: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ...) -> None: ...
        FUNCTIONS_FIELD_NUMBER: _ClassVar[int]
        GROUPINGS_FIELD_NUMBER: _ClassVar[int]
        functions: _containers.RepeatedCompositeFieldContainer[Operator.GroupBy.GroupByFunction]
        groupings: _containers.RepeatedCompositeFieldContainer[Operator.GroupBy.GroupingSet]
        def __init__(self, functions: _Optional[_Iterable[_Union[Operator.GroupBy.GroupByFunction, _Mapping]]] = ..., groupings: _Optional[_Iterable[_Union[Operator.GroupBy.GroupingSet, _Mapping]]] = ...) -> None: ...
    class Union(_message.Message):
        __slots__ = ("query",)
        QUERY_FIELD_NUMBER: _ClassVar[int]
        query: Query
        def __init__(self, query: _Optional[_Union[Query, _Mapping]] = ...) -> None: ...
    class EnrichmentSource(_message.Message):
        __slots__ = ("custom_enrichment",)
        CUSTOM_ENRICHMENT_FIELD_NUMBER: _ClassVar[int]
        custom_enrichment: CustomEnrichment
        def __init__(self, custom_enrichment: _Optional[_Union[CustomEnrichment, _Mapping]] = ...) -> None: ...
    class Enrich(_message.Message):
        __slots__ = ("source", "left", "into")
        SOURCE_FIELD_NUMBER: _ClassVar[int]
        LEFT_FIELD_NUMBER: _ClassVar[int]
        INTO_FIELD_NUMBER: _ClassVar[int]
        source: Operator.EnrichmentSource
        left: _expression_pb2.Expression
        into: _expression_pb2.Expression.Keypath
        def __init__(self, source: _Optional[_Union[Operator.EnrichmentSource, _Mapping]] = ..., left: _Optional[_Union[_expression_pb2.Expression, _Mapping]] = ..., into: _Optional[_Union[_expression_pb2.Expression.Keypath, _Mapping]] = ...) -> None: ...
    FILTER_FIELD_NUMBER: _ClassVar[int]
    CREATE_FIELD_NUMBER: _ClassVar[int]
    MOVE_FIELD_NUMBER: _ClassVar[int]
    REMOVE_FIELD_NUMBER: _ClassVar[int]
    REPLACE_FIELD_NUMBER: _ClassVar[int]
    REDACT_FIELD_NUMBER: _ClassVar[int]
    CHOOSE_FIELD_NUMBER: _ClassVar[int]
    CONVERT_FIELD_NUMBER: _ClassVar[int]
    EXTRACT_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_FIELD_NUMBER: _ClassVar[int]
    UNION_FIELD_NUMBER: _ClassVar[int]
    ENRICH_FIELD_NUMBER: _ClassVar[int]
    filter: Operator.Filter
    create: Operator.Create
    move: Operator.Move
    remove: Operator.Remove
    replace: Operator.Replace
    redact: Operator.Redact
    choose: Operator.Choose
    convert: Operator.Convert
    extract: Operator.Extract
    limit: Operator.Limit
    order_by: Operator.OrderBy
    group_by: Operator.GroupBy
    union: Operator.Union
    enrich: Operator.Enrich
    def __init__(self, filter: _Optional[_Union[Operator.Filter, _Mapping]] = ..., create: _Optional[_Union[Operator.Create, _Mapping]] = ..., move: _Optional[_Union[Operator.Move, _Mapping]] = ..., remove: _Optional[_Union[Operator.Remove, _Mapping]] = ..., replace: _Optional[_Union[Operator.Replace, _Mapping]] = ..., redact: _Optional[_Union[Operator.Redact, _Mapping]] = ..., choose: _Optional[_Union[Operator.Choose, _Mapping]] = ..., convert: _Optional[_Union[Operator.Convert, _Mapping]] = ..., extract: _Optional[_Union[Operator.Extract, _Mapping]] = ..., limit: _Optional[_Union[Operator.Limit, _Mapping]] = ..., order_by: _Optional[_Union[Operator.OrderBy, _Mapping]] = ..., group_by: _Optional[_Union[Operator.GroupBy, _Mapping]] = ..., union: _Optional[_Union[Operator.Union, _Mapping]] = ..., enrich: _Optional[_Union[Operator.Enrich, _Mapping]] = ...) -> None: ...
