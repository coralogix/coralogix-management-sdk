from com.coralogix.quota.v1 import test_policies_result_pb2 as _test_policies_result_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class BulkTestLogPoliciesResponse(_message.Message):
    __slots__ = ("test_policies_bulk_result",)
    TEST_POLICIES_BULK_RESULT_FIELD_NUMBER: _ClassVar[int]
    test_policies_bulk_result: _containers.RepeatedCompositeFieldContainer[_test_policies_result_pb2.TestPoliciesResult]
    def __init__(self, test_policies_bulk_result: _Optional[_Iterable[_Union[_test_policies_result_pb2.TestPoliciesResult, _Mapping]]] = ...) -> None: ...
