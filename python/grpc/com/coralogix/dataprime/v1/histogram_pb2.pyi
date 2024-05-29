from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.dataprime.v1 import common_pb2 as _common_pb2
from com.coralogix.dataprime.v1 import aggregation_operators_pb2 as _aggregation_operators_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Aggregation(_message.Message):
    __slots__ = ("field", "type")
    FIELD_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    field: _common_pb2.UntypedKeypath
    type: _aggregation_operators_pb2.Operator
    def __init__(self, field: _Optional[_Union[_common_pb2.UntypedKeypath, _Mapping]] = ..., type: _Optional[_Union[_aggregation_operators_pb2.Operator, str]] = ...) -> None: ...

class Data(_message.Message):
    __slots__ = ("key", "doc_count", "agg_val", "compared")
    KEY_FIELD_NUMBER: _ClassVar[int]
    DOC_COUNT_FIELD_NUMBER: _ClassVar[int]
    AGG_VAL_FIELD_NUMBER: _ClassVar[int]
    COMPARED_FIELD_NUMBER: _ClassVar[int]
    key: _wrappers_pb2.StringValue
    doc_count: _wrappers_pb2.Int64Value
    agg_val: _wrappers_pb2.DoubleValue
    compared: ComparedData
    def __init__(self, key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., doc_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., agg_val: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ..., compared: _Optional[_Union[ComparedData, _Mapping]] = ...) -> None: ...

class ComparedData(_message.Message):
    __slots__ = ("historical_doc_count", "historical_agg_val")
    HISTORICAL_DOC_COUNT_FIELD_NUMBER: _ClassVar[int]
    HISTORICAL_AGG_VAL_FIELD_NUMBER: _ClassVar[int]
    historical_doc_count: _wrappers_pb2.Int64Value
    historical_agg_val: _wrappers_pb2.DoubleValue
    def __init__(self, historical_doc_count: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., historical_agg_val: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ...) -> None: ...

class HistogramSlice(_message.Message):
    __slots__ = ("value", "data")
    VALUE_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    value: _wrappers_pb2.Int64Value
    data: _containers.RepeatedCompositeFieldContainer[Data]
    def __init__(self, value: _Optional[_Union[_wrappers_pb2.Int64Value, _Mapping]] = ..., data: _Optional[_Iterable[_Union[Data, _Mapping]]] = ...) -> None: ...
