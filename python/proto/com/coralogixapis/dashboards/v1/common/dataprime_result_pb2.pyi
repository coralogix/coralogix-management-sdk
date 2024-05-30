from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DataprimeResult(_message.Message):
    __slots__ = ("metadata", "labels", "user_data")
    class KeyValue(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    METADATA_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    USER_DATA_FIELD_NUMBER: _ClassVar[int]
    metadata: _containers.RepeatedCompositeFieldContainer[DataprimeResult.KeyValue]
    labels: _containers.RepeatedCompositeFieldContainer[DataprimeResult.KeyValue]
    user_data: str
    def __init__(self, metadata: _Optional[_Iterable[_Union[DataprimeResult.KeyValue, _Mapping]]] = ..., labels: _Optional[_Iterable[_Union[DataprimeResult.KeyValue, _Mapping]]] = ..., user_data: _Optional[str] = ...) -> None: ...
