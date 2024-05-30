from com.coralogixapis.views.v1 import audit_log_pb2 as _audit_log_pb2
from com.coralogixapis.views.v1 import filters_pb2 as _filters_pb2
from com.coralogixapis.views.v1 import search_query_pb2 as _search_query_pb2
from com.coralogixapis.views.v1 import time_selection_pb2 as _time_selection_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateViewRequest(_message.Message):
    __slots__ = ("name", "search_query", "time_selection", "filters", "folder_id")
    NAME_FIELD_NUMBER: _ClassVar[int]
    SEARCH_QUERY_FIELD_NUMBER: _ClassVar[int]
    TIME_SELECTION_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    FOLDER_ID_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    search_query: _search_query_pb2.SearchQuery
    time_selection: _time_selection_pb2.TimeSelection
    filters: _filters_pb2.SelectedFilters
    folder_id: _wrappers_pb2.StringValue
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., search_query: _Optional[_Union[_search_query_pb2.SearchQuery, _Mapping]] = ..., time_selection: _Optional[_Union[_time_selection_pb2.TimeSelection, _Mapping]] = ..., filters: _Optional[_Union[_filters_pb2.SelectedFilters, _Mapping]] = ..., folder_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class CreateViewResponse(_message.Message):
    __slots__ = ("view",)
    VIEW_FIELD_NUMBER: _ClassVar[int]
    view: View
    def __init__(self, view: _Optional[_Union[View, _Mapping]] = ...) -> None: ...

class ReplaceViewRequest(_message.Message):
    __slots__ = ("view",)
    VIEW_FIELD_NUMBER: _ClassVar[int]
    view: View
    def __init__(self, view: _Optional[_Union[View, _Mapping]] = ...) -> None: ...

class ReplaceViewResponse(_message.Message):
    __slots__ = ("view",)
    VIEW_FIELD_NUMBER: _ClassVar[int]
    view: View
    def __init__(self, view: _Optional[_Union[View, _Mapping]] = ...) -> None: ...

class GetViewRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.Int32Value
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class GetViewResponse(_message.Message):
    __slots__ = ("view",)
    VIEW_FIELD_NUMBER: _ClassVar[int]
    view: View
    def __init__(self, view: _Optional[_Union[View, _Mapping]] = ...) -> None: ...

class DeleteViewRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.Int32Value
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...

class DeleteViewResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListViewsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListViewsResponse(_message.Message):
    __slots__ = ("views",)
    VIEWS_FIELD_NUMBER: _ClassVar[int]
    views: _containers.RepeatedCompositeFieldContainer[View]
    def __init__(self, views: _Optional[_Iterable[_Union[View, _Mapping]]] = ...) -> None: ...

class View(_message.Message):
    __slots__ = ("id", "name", "search_query", "time_selection", "filters", "folder_id")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    SEARCH_QUERY_FIELD_NUMBER: _ClassVar[int]
    TIME_SELECTION_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    FOLDER_ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.Int32Value
    name: _wrappers_pb2.StringValue
    search_query: _search_query_pb2.SearchQuery
    time_selection: _time_selection_pb2.TimeSelection
    filters: _filters_pb2.SelectedFilters
    folder_id: _wrappers_pb2.StringValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., search_query: _Optional[_Union[_search_query_pb2.SearchQuery, _Mapping]] = ..., time_selection: _Optional[_Union[_time_selection_pb2.TimeSelection, _Mapping]] = ..., filters: _Optional[_Union[_filters_pb2.SelectedFilters, _Mapping]] = ..., folder_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
