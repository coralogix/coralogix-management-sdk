from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class FeatureControl(_message.Message):
    __slots__ = ("feature", "on_query_duration_gt", "control")
    class Feature(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        FEATURE_UNSPECIFIED: _ClassVar[FeatureControl.Feature]
        FEATURE_UNKNOWN_FIELDS: _ClassVar[FeatureControl.Feature]
        FEATURE_OPTIONAL_SOURCE: _ClassVar[FeatureControl.Feature]
        FEATURE_INVALID_REGEX: _ClassVar[FeatureControl.Feature]
        FEATURE_WILDFIND: _ClassVar[FeatureControl.Feature]
    FEATURE_UNSPECIFIED: FeatureControl.Feature
    FEATURE_UNKNOWN_FIELDS: FeatureControl.Feature
    FEATURE_OPTIONAL_SOURCE: FeatureControl.Feature
    FEATURE_INVALID_REGEX: FeatureControl.Feature
    FEATURE_WILDFIND: FeatureControl.Feature
    class Control(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        CONTROL_UNSPECIFIED: _ClassVar[FeatureControl.Control]
        CONTROL_WARNING: _ClassVar[FeatureControl.Control]
        CONTROL_ERROR: _ClassVar[FeatureControl.Control]
    CONTROL_UNSPECIFIED: FeatureControl.Control
    CONTROL_WARNING: FeatureControl.Control
    CONTROL_ERROR: FeatureControl.Control
    FEATURE_FIELD_NUMBER: _ClassVar[int]
    ON_QUERY_DURATION_GT_FIELD_NUMBER: _ClassVar[int]
    CONTROL_FIELD_NUMBER: _ClassVar[int]
    feature: FeatureControl.Feature
    on_query_duration_gt: _wrappers_pb2.StringValue
    control: FeatureControl.Control
    def __init__(self, feature: _Optional[_Union[FeatureControl.Feature, str]] = ..., on_query_duration_gt: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., control: _Optional[_Union[FeatureControl.Control, str]] = ...) -> None: ...

class ListFeatureControlRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ListFeatureControlResponse(_message.Message):
    __slots__ = ("feature_controls",)
    FEATURE_CONTROLS_FIELD_NUMBER: _ClassVar[int]
    feature_controls: _containers.RepeatedCompositeFieldContainer[FeatureControl]
    def __init__(self, feature_controls: _Optional[_Iterable[_Union[FeatureControl, _Mapping]]] = ...) -> None: ...
