from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetQueryResultsRequest(_message.Message):
    __slots__ = ("query_id",)
    QUERY_ID_FIELD_NUMBER: _ClassVar[int]
    query_id: str
    def __init__(self, query_id: _Optional[str] = ...) -> None: ...

class GetQueryResultsResponse(_message.Message):
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
    metadata: _containers.RepeatedCompositeFieldContainer[GetQueryResultsResponse.KeyValue]
    labels: _containers.RepeatedCompositeFieldContainer[GetQueryResultsResponse.KeyValue]
    user_data: str
    def __init__(self, metadata: _Optional[_Iterable[_Union[GetQueryResultsResponse.KeyValue, _Mapping]]] = ..., labels: _Optional[_Iterable[_Union[GetQueryResultsResponse.KeyValue, _Mapping]]] = ..., user_data: _Optional[str] = ...) -> None: ...
