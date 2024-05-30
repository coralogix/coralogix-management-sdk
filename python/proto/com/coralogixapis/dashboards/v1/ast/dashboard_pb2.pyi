from com.coralogixapis.dashboards.v1.ast import annotation_pb2 as _annotation_pb2
from com.coralogixapis.dashboards.v1.ast import filter_pb2 as _filter_pb2
from com.coralogixapis.dashboards.v1.ast import folder_path_pb2 as _folder_path_pb2
from com.coralogixapis.dashboards.v1.ast import layout_pb2 as _layout_pb2
from com.coralogixapis.dashboards.v1.ast import variable_pb2 as _variable_pb2
from com.coralogixapis.dashboards.v1.common import time_frame_pb2 as _time_frame_pb2
from com.coralogixapis.dashboards.v1 import types_pb2 as _types_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Dashboard(_message.Message):
    __slots__ = ("id", "name", "description", "layout", "variables", "filters", "absolute_time_frame", "relative_time_frame", "folder_id", "folder_path", "annotations", "off", "two_minutes", "five_minutes")
    class AutoRefreshOff(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class AutoRefreshTwoMinutes(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class AutoRefreshFiveMinutes(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    LAYOUT_FIELD_NUMBER: _ClassVar[int]
    VARIABLES_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    ABSOLUTE_TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    RELATIVE_TIME_FRAME_FIELD_NUMBER: _ClassVar[int]
    FOLDER_ID_FIELD_NUMBER: _ClassVar[int]
    FOLDER_PATH_FIELD_NUMBER: _ClassVar[int]
    ANNOTATIONS_FIELD_NUMBER: _ClassVar[int]
    OFF_FIELD_NUMBER: _ClassVar[int]
    TWO_MINUTES_FIELD_NUMBER: _ClassVar[int]
    FIVE_MINUTES_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    layout: _layout_pb2.Layout
    variables: _containers.RepeatedCompositeFieldContainer[_variable_pb2.Variable]
    filters: _containers.RepeatedCompositeFieldContainer[_filter_pb2.Filter]
    absolute_time_frame: _time_frame_pb2.TimeFrame
    relative_time_frame: _duration_pb2.Duration
    folder_id: _types_pb2.UUID
    folder_path: _folder_path_pb2.FolderPath
    annotations: _containers.RepeatedCompositeFieldContainer[_annotation_pb2.Annotation]
    off: Dashboard.AutoRefreshOff
    two_minutes: Dashboard.AutoRefreshTwoMinutes
    five_minutes: Dashboard.AutoRefreshFiveMinutes
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., layout: _Optional[_Union[_layout_pb2.Layout, _Mapping]] = ..., variables: _Optional[_Iterable[_Union[_variable_pb2.Variable, _Mapping]]] = ..., filters: _Optional[_Iterable[_Union[_filter_pb2.Filter, _Mapping]]] = ..., absolute_time_frame: _Optional[_Union[_time_frame_pb2.TimeFrame, _Mapping]] = ..., relative_time_frame: _Optional[_Union[_duration_pb2.Duration, _Mapping]] = ..., folder_id: _Optional[_Union[_types_pb2.UUID, _Mapping]] = ..., folder_path: _Optional[_Union[_folder_path_pb2.FolderPath, _Mapping]] = ..., annotations: _Optional[_Iterable[_Union[_annotation_pb2.Annotation, _Mapping]]] = ..., off: _Optional[_Union[Dashboard.AutoRefreshOff, _Mapping]] = ..., two_minutes: _Optional[_Union[Dashboard.AutoRefreshTwoMinutes, _Mapping]] = ..., five_minutes: _Optional[_Union[Dashboard.AutoRefreshFiveMinutes, _Mapping]] = ...) -> None: ...
