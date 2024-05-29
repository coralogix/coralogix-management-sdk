from google.protobuf import timestamp_pb2 as _timestamp_pb2
from com.coralogixapis.dashboards.v1.ast.widgets import bar_chart_pb2 as _bar_chart_pb2
from com.coralogixapis.dashboards.v1.ast.widgets import data_table_pb2 as _data_table_pb2
from com.coralogixapis.dashboards.v1.ast.widgets import gauge_pb2 as _gauge_pb2
from com.coralogixapis.dashboards.v1.ast.widgets import horizontal_bar_chart_pb2 as _horizontal_bar_chart_pb2
from com.coralogixapis.dashboards.v1.ast.widgets import line_chart_pb2 as _line_chart_pb2
from com.coralogixapis.dashboards.v1.ast.widgets import markdown_pb2 as _markdown_pb2
from com.coralogixapis.dashboards.v1.ast.widgets import pie_chart_pb2 as _pie_chart_pb2
from com.coralogixapis.dashboards.v1 import types_pb2 as _types_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Widget(_message.Message):
    __slots__ = ("id", "title", "description", "definition", "appearance", "created_at", "updated_at")
    class Definition(_message.Message):
        __slots__ = ("line_chart", "data_table", "gauge", "pie_chart", "bar_chart", "horizontal_bar_chart", "markdown")
        LINE_CHART_FIELD_NUMBER: _ClassVar[int]
        DATA_TABLE_FIELD_NUMBER: _ClassVar[int]
        GAUGE_FIELD_NUMBER: _ClassVar[int]
        PIE_CHART_FIELD_NUMBER: _ClassVar[int]
        BAR_CHART_FIELD_NUMBER: _ClassVar[int]
        HORIZONTAL_BAR_CHART_FIELD_NUMBER: _ClassVar[int]
        MARKDOWN_FIELD_NUMBER: _ClassVar[int]
        line_chart: _line_chart_pb2.LineChart
        data_table: _data_table_pb2.DataTable
        gauge: _gauge_pb2.Gauge
        pie_chart: _pie_chart_pb2.PieChart
        bar_chart: _bar_chart_pb2.BarChart
        horizontal_bar_chart: _horizontal_bar_chart_pb2.HorizontalBarChart
        markdown: _markdown_pb2.Markdown
        def __init__(self, line_chart: _Optional[_Union[_line_chart_pb2.LineChart, _Mapping]] = ..., data_table: _Optional[_Union[_data_table_pb2.DataTable, _Mapping]] = ..., gauge: _Optional[_Union[_gauge_pb2.Gauge, _Mapping]] = ..., pie_chart: _Optional[_Union[_pie_chart_pb2.PieChart, _Mapping]] = ..., bar_chart: _Optional[_Union[_bar_chart_pb2.BarChart, _Mapping]] = ..., horizontal_bar_chart: _Optional[_Union[_horizontal_bar_chart_pb2.HorizontalBarChart, _Mapping]] = ..., markdown: _Optional[_Union[_markdown_pb2.Markdown, _Mapping]] = ...) -> None: ...
    class Appearance(_message.Message):
        __slots__ = ("width",)
        WIDTH_FIELD_NUMBER: _ClassVar[int]
        width: _wrappers_pb2.Int32Value
        def __init__(self, width: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    DEFINITION_FIELD_NUMBER: _ClassVar[int]
    APPEARANCE_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: _types_pb2.UUID
    title: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    definition: Widget.Definition
    appearance: Widget.Appearance
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[_Union[_types_pb2.UUID, _Mapping]] = ..., title: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., definition: _Optional[_Union[Widget.Definition, _Mapping]] = ..., appearance: _Optional[_Union[Widget.Appearance, _Mapping]] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
