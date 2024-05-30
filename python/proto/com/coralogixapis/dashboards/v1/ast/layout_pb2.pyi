from com.coralogixapis.dashboards.v1.ast import widget_pb2 as _widget_pb2
from com.coralogixapis.dashboards.v1 import types_pb2 as _types_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Layout(_message.Message):
    __slots__ = ("sections",)
    SECTIONS_FIELD_NUMBER: _ClassVar[int]
    sections: _containers.RepeatedCompositeFieldContainer[Section]
    def __init__(self, sections: _Optional[_Iterable[_Union[Section, _Mapping]]] = ...) -> None: ...

class Section(_message.Message):
    __slots__ = ("id", "rows", "options")
    ID_FIELD_NUMBER: _ClassVar[int]
    ROWS_FIELD_NUMBER: _ClassVar[int]
    OPTIONS_FIELD_NUMBER: _ClassVar[int]
    id: _types_pb2.UUID
    rows: _containers.RepeatedCompositeFieldContainer[Row]
    options: SectionOptions
    def __init__(self, id: _Optional[_Union[_types_pb2.UUID, _Mapping]] = ..., rows: _Optional[_Iterable[_Union[Row, _Mapping]]] = ..., options: _Optional[_Union[SectionOptions, _Mapping]] = ...) -> None: ...

class Row(_message.Message):
    __slots__ = ("id", "appearance", "widgets")
    class Appearance(_message.Message):
        __slots__ = ("height",)
        HEIGHT_FIELD_NUMBER: _ClassVar[int]
        height: _wrappers_pb2.Int32Value
        def __init__(self, height: _Optional[_Union[_wrappers_pb2.Int32Value, _Mapping]] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    APPEARANCE_FIELD_NUMBER: _ClassVar[int]
    WIDGETS_FIELD_NUMBER: _ClassVar[int]
    id: _types_pb2.UUID
    appearance: Row.Appearance
    widgets: _containers.RepeatedCompositeFieldContainer[_widget_pb2.Widget]
    def __init__(self, id: _Optional[_Union[_types_pb2.UUID, _Mapping]] = ..., appearance: _Optional[_Union[Row.Appearance, _Mapping]] = ..., widgets: _Optional[_Iterable[_Union[_widget_pb2.Widget, _Mapping]]] = ...) -> None: ...

class SectionOptions(_message.Message):
    __slots__ = ("internal", "custom")
    INTERNAL_FIELD_NUMBER: _ClassVar[int]
    CUSTOM_FIELD_NUMBER: _ClassVar[int]
    internal: InternalSectionOptions
    custom: CustomSectionOptions
    def __init__(self, internal: _Optional[_Union[InternalSectionOptions, _Mapping]] = ..., custom: _Optional[_Union[CustomSectionOptions, _Mapping]] = ...) -> None: ...

class InternalSectionOptions(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class CustomSectionOptions(_message.Message):
    __slots__ = ("name", "description", "collapsed")
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    COLLAPSED_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    collapsed: _wrappers_pb2.BoolValue
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., collapsed: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
