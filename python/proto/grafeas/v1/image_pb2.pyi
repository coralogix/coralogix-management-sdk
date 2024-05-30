from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Layer(_message.Message):
    __slots__ = ("directive", "arguments")
    DIRECTIVE_FIELD_NUMBER: _ClassVar[int]
    ARGUMENTS_FIELD_NUMBER: _ClassVar[int]
    directive: str
    arguments: str
    def __init__(self, directive: _Optional[str] = ..., arguments: _Optional[str] = ...) -> None: ...

class Fingerprint(_message.Message):
    __slots__ = ("v1_name", "v2_blob", "v2_name")
    V1_NAME_FIELD_NUMBER: _ClassVar[int]
    V2_BLOB_FIELD_NUMBER: _ClassVar[int]
    V2_NAME_FIELD_NUMBER: _ClassVar[int]
    v1_name: str
    v2_blob: _containers.RepeatedScalarFieldContainer[str]
    v2_name: str
    def __init__(self, v1_name: _Optional[str] = ..., v2_blob: _Optional[_Iterable[str]] = ..., v2_name: _Optional[str] = ...) -> None: ...

class ImageNote(_message.Message):
    __slots__ = ("resource_url", "fingerprint")
    RESOURCE_URL_FIELD_NUMBER: _ClassVar[int]
    FINGERPRINT_FIELD_NUMBER: _ClassVar[int]
    resource_url: str
    fingerprint: Fingerprint
    def __init__(self, resource_url: _Optional[str] = ..., fingerprint: _Optional[_Union[Fingerprint, _Mapping]] = ...) -> None: ...

class ImageOccurrence(_message.Message):
    __slots__ = ("fingerprint", "distance", "layer_info", "base_resource_url")
    FINGERPRINT_FIELD_NUMBER: _ClassVar[int]
    DISTANCE_FIELD_NUMBER: _ClassVar[int]
    LAYER_INFO_FIELD_NUMBER: _ClassVar[int]
    BASE_RESOURCE_URL_FIELD_NUMBER: _ClassVar[int]
    fingerprint: Fingerprint
    distance: int
    layer_info: _containers.RepeatedCompositeFieldContainer[Layer]
    base_resource_url: str
    def __init__(self, fingerprint: _Optional[_Union[Fingerprint, _Mapping]] = ..., distance: _Optional[int] = ..., layer_info: _Optional[_Iterable[_Union[Layer, _Mapping]]] = ..., base_resource_url: _Optional[str] = ...) -> None: ...
