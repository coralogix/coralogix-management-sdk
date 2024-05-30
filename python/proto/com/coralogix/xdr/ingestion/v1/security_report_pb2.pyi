from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SecurityReport(_message.Message):
    __slots__ = ("context", "test_results")
    class Context(_message.Message):
        __slots__ = ("application_name", "subsystem_name", "computer_name", "provider", "service", "execution_id")
        APPLICATION_NAME_FIELD_NUMBER: _ClassVar[int]
        SUBSYSTEM_NAME_FIELD_NUMBER: _ClassVar[int]
        COMPUTER_NAME_FIELD_NUMBER: _ClassVar[int]
        PROVIDER_FIELD_NUMBER: _ClassVar[int]
        SERVICE_FIELD_NUMBER: _ClassVar[int]
        EXECUTION_ID_FIELD_NUMBER: _ClassVar[int]
        application_name: _wrappers_pb2.StringValue
        subsystem_name: _wrappers_pb2.StringValue
        computer_name: _wrappers_pb2.StringValue
        provider: _wrappers_pb2.StringValue
        service: _wrappers_pb2.StringValue
        execution_id: _wrappers_pb2.StringValue
        def __init__(self, application_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., subsystem_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., computer_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., provider: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., service: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., execution_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class TestResult(_message.Message):
        __slots__ = ("name", "start_time", "end_time", "item", "item_type", "result", "additional_data")
        class Result(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            TEST_PASSED: _ClassVar[SecurityReport.TestResult.Result]
            TEST_FAILED: _ClassVar[SecurityReport.TestResult.Result]
        TEST_PASSED: SecurityReport.TestResult.Result
        TEST_FAILED: SecurityReport.TestResult.Result
        NAME_FIELD_NUMBER: _ClassVar[int]
        START_TIME_FIELD_NUMBER: _ClassVar[int]
        END_TIME_FIELD_NUMBER: _ClassVar[int]
        ITEM_FIELD_NUMBER: _ClassVar[int]
        ITEM_TYPE_FIELD_NUMBER: _ClassVar[int]
        RESULT_FIELD_NUMBER: _ClassVar[int]
        ADDITIONAL_DATA_FIELD_NUMBER: _ClassVar[int]
        name: _wrappers_pb2.StringValue
        start_time: _timestamp_pb2.Timestamp
        end_time: _timestamp_pb2.Timestamp
        item: _wrappers_pb2.StringValue
        item_type: _wrappers_pb2.StringValue
        result: SecurityReport.TestResult.Result
        additional_data: _struct_pb2.Struct
        def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., start_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., end_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., item: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., item_type: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., result: _Optional[_Union[SecurityReport.TestResult.Result, str]] = ..., additional_data: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...
    CONTEXT_FIELD_NUMBER: _ClassVar[int]
    TEST_RESULTS_FIELD_NUMBER: _ClassVar[int]
    context: SecurityReport.Context
    test_results: _containers.RepeatedCompositeFieldContainer[SecurityReport.TestResult]
    def __init__(self, context: _Optional[_Union[SecurityReport.Context, _Mapping]] = ..., test_results: _Optional[_Iterable[_Union[SecurityReport.TestResult, _Mapping]]] = ...) -> None: ...
