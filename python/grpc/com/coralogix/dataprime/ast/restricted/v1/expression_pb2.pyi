from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.dataprime.ast.restricted.v1 import types_pb2 as _types_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Expression(_message.Message):
    __slots__ = ("datatype", "null_e", "keypath", "boolean", "integer", "str", "infix_op", "function_call", "long", "double", "timestamp")
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
    class InfixOp(_message.Message):
        __slots__ = ("left", "op", "right")
        class Op(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            OP_UNSPECIFIED: _ClassVar[Expression.InfixOp.Op]
            OP_AND: _ClassVar[Expression.InfixOp.Op]
            OP_OR: _ClassVar[Expression.InfixOp.Op]
            OP_EQ: _ClassVar[Expression.InfixOp.Op]
            OP_NEQ: _ClassVar[Expression.InfixOp.Op]
            OP_LT: _ClassVar[Expression.InfixOp.Op]
            OP_LTE: _ClassVar[Expression.InfixOp.Op]
            OP_GT: _ClassVar[Expression.InfixOp.Op]
            OP_GTE: _ClassVar[Expression.InfixOp.Op]
        OP_UNSPECIFIED: Expression.InfixOp.Op
        OP_AND: Expression.InfixOp.Op
        OP_OR: Expression.InfixOp.Op
        OP_EQ: Expression.InfixOp.Op
        OP_NEQ: Expression.InfixOp.Op
        OP_LT: Expression.InfixOp.Op
        OP_LTE: Expression.InfixOp.Op
        OP_GT: Expression.InfixOp.Op
        OP_GTE: Expression.InfixOp.Op
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
        __slots__ = ("contains", "starts_with", "ends_with")
        class In(_message.Message):
            __slots__ = ("comparand", "value", "values")
            COMPARAND_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            VALUES_FIELD_NUMBER: _ClassVar[int]
            comparand: Expression
            value: Expression
            values: _containers.RepeatedCompositeFieldContainer[Expression]
            def __init__(self, comparand: _Optional[_Union[Expression, _Mapping]] = ..., value: _Optional[_Union[Expression, _Mapping]] = ..., values: _Optional[_Iterable[_Union[Expression, _Mapping]]] = ...) -> None: ...
        class Contains(_message.Message):
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
        CONTAINS_FIELD_NUMBER: _ClassVar[int]
        STARTS_WITH_FIELD_NUMBER: _ClassVar[int]
        ENDS_WITH_FIELD_NUMBER: _ClassVar[int]
        IN_FIELD_NUMBER: _ClassVar[int]
        contains: Expression.FunctionCall.Contains
        starts_with: Expression.FunctionCall.StartsWith
        ends_with: Expression.FunctionCall.EndsWith
        def __init__(self, contains: _Optional[_Union[Expression.FunctionCall.Contains, _Mapping]] = ..., starts_with: _Optional[_Union[Expression.FunctionCall.StartsWith, _Mapping]] = ..., ends_with: _Optional[_Union[Expression.FunctionCall.EndsWith, _Mapping]] = ..., **kwargs) -> None: ...
    class Timestamp(_message.Message):
        __slots__ = ("epoch_nanos",)
        EPOCH_NANOS_FIELD_NUMBER: _ClassVar[int]
        epoch_nanos: int
        def __init__(self, epoch_nanos: _Optional[int] = ...) -> None: ...
    DATATYPE_FIELD_NUMBER: _ClassVar[int]
    NULL_E_FIELD_NUMBER: _ClassVar[int]
    KEYPATH_FIELD_NUMBER: _ClassVar[int]
    BOOLEAN_FIELD_NUMBER: _ClassVar[int]
    INTEGER_FIELD_NUMBER: _ClassVar[int]
    STR_FIELD_NUMBER: _ClassVar[int]
    INFIX_OP_FIELD_NUMBER: _ClassVar[int]
    NOT_FIELD_NUMBER: _ClassVar[int]
    FUNCTION_CALL_FIELD_NUMBER: _ClassVar[int]
    LONG_FIELD_NUMBER: _ClassVar[int]
    DOUBLE_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    datatype: _types_pb2.Datatype
    null_e: Expression.Null
    keypath: Expression.Keypath
    boolean: _wrappers_pb2.BoolValue
    integer: _wrappers_pb2.Int32Value
    str: _wrappers_pb2.StringValue
    infix_op: Expression.InfixOp
    function_call: Expression.FunctionCall
    long: _wrappers_pb2.Int64Value
    double: _wrappers_pb2.DoubleValue
    timestamp: Expression.Timestamp
    def __init__(self, datatype: _Optional[_Union[_types_pb2.Datatype, _Mapping]] = ..., null_e: _Optional[_Union[Expression.Null, _Mapping]] = ..., keypath: _Optional[_Union[Expression.Keypath, _Mapping]] = ..., boolean: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., integer: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., str: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., infix_op: _Optional[_Union[Expression.InfixOp, _Mapping]] = ..., function_call: _Optional[_Union[Expression.FunctionCall, _Mapping]] = ..., long: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., double: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ..., timestamp: _Optional[_Union[Expression.Timestamp, _Mapping]] = ..., **kwargs) -> None: ...
