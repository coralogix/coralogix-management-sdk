from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SpanField(_message.Message):
    __slots__ = ("metadata_field", "tag_field", "process_tag_field")
    class MetadataField(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        METADATA_FIELD_UNSPECIFIED: _ClassVar[SpanField.MetadataField]
        METADATA_FIELD_APPLICATION_NAME: _ClassVar[SpanField.MetadataField]
        METADATA_FIELD_SUBSYSTEM_NAME: _ClassVar[SpanField.MetadataField]
        METADATA_FIELD_SERVICE_NAME: _ClassVar[SpanField.MetadataField]
        METADATA_FIELD_OPERATION_NAME: _ClassVar[SpanField.MetadataField]
    METADATA_FIELD_UNSPECIFIED: SpanField.MetadataField
    METADATA_FIELD_APPLICATION_NAME: SpanField.MetadataField
    METADATA_FIELD_SUBSYSTEM_NAME: SpanField.MetadataField
    METADATA_FIELD_SERVICE_NAME: SpanField.MetadataField
    METADATA_FIELD_OPERATION_NAME: SpanField.MetadataField
    METADATA_FIELD_FIELD_NUMBER: _ClassVar[int]
    TAG_FIELD_FIELD_NUMBER: _ClassVar[int]
    PROCESS_TAG_FIELD_FIELD_NUMBER: _ClassVar[int]
    metadata_field: SpanField.MetadataField
    tag_field: _wrappers_pb2.StringValue
    process_tag_field: _wrappers_pb2.StringValue
    def __init__(self, metadata_field: _Optional[_Union[SpanField.MetadataField, str]] = ..., tag_field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., process_tag_field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
