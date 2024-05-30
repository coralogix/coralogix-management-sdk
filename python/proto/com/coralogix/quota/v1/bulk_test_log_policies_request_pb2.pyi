from com.coralogix.quota.v1 import log_meta_field_values_pb2 as _log_meta_field_values_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class BulkTestLogPoliciesRequest(_message.Message):
    __slots__ = ("meta_fields_values_list",)
    META_FIELDS_VALUES_LIST_FIELD_NUMBER: _ClassVar[int]
    meta_fields_values_list: _containers.RepeatedCompositeFieldContainer[_log_meta_field_values_pb2.LogMetaFieldsValues]
    def __init__(self, meta_fields_values_list: _Optional[_Iterable[_Union[_log_meta_field_values_pb2.LogMetaFieldsValues, _Mapping]]] = ...) -> None: ...
