from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Legend(_message.Message):
    __slots__ = ("is_visible", "columns", "group_by_query")
    class LegendColumn(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        LEGEND_COLUMN_UNSPECIFIED: _ClassVar[Legend.LegendColumn]
        LEGEND_COLUMN_MIN: _ClassVar[Legend.LegendColumn]
        LEGEND_COLUMN_MAX: _ClassVar[Legend.LegendColumn]
        LEGEND_COLUMN_SUM: _ClassVar[Legend.LegendColumn]
        LEGEND_COLUMN_AVG: _ClassVar[Legend.LegendColumn]
        LEGEND_COLUMN_LAST: _ClassVar[Legend.LegendColumn]
        LEGEND_COLUMN_NAME: _ClassVar[Legend.LegendColumn]
    LEGEND_COLUMN_UNSPECIFIED: Legend.LegendColumn
    LEGEND_COLUMN_MIN: Legend.LegendColumn
    LEGEND_COLUMN_MAX: Legend.LegendColumn
    LEGEND_COLUMN_SUM: Legend.LegendColumn
    LEGEND_COLUMN_AVG: Legend.LegendColumn
    LEGEND_COLUMN_LAST: Legend.LegendColumn
    LEGEND_COLUMN_NAME: Legend.LegendColumn
    IS_VISIBLE_FIELD_NUMBER: _ClassVar[int]
    COLUMNS_FIELD_NUMBER: _ClassVar[int]
    GROUP_BY_QUERY_FIELD_NUMBER: _ClassVar[int]
    is_visible: _wrappers_pb2.BoolValue
    columns: _containers.RepeatedScalarFieldContainer[Legend.LegendColumn]
    group_by_query: _wrappers_pb2.BoolValue
    def __init__(self, is_visible: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., columns: _Optional[_Iterable[_Union[Legend.LegendColumn, str]]] = ..., group_by_query: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
