from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Rule(_message.Message):
    __slots__ = ("id", "name", "description", "source_field", "parameters", "enabled", "order")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_FIELD_NUMBER: _ClassVar[int]
    PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    ENABLED_FIELD_NUMBER: _ClassVar[int]
    ORDER_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    source_field: _wrappers_pb2.StringValue
    parameters: RuleParameters
    enabled: _wrappers_pb2.BoolValue
    order: _wrappers_pb2.UInt32Value
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., source_field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., parameters: _Optional[_Union[RuleParameters, _Mapping]] = ..., enabled: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., order: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...

class RuleParameters(_message.Message):
    __slots__ = ("extract_parameters", "json_extract_parameters", "replace_parameters", "parse_parameters", "allow_parameters", "block_parameters", "extract_timestamp_parameters", "remove_fields_parameters", "json_stringify_parameters", "json_parse_parameters")
    EXTRACT_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    JSON_EXTRACT_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    REPLACE_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    PARSE_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    ALLOW_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    BLOCK_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    EXTRACT_TIMESTAMP_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    REMOVE_FIELDS_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    JSON_STRINGIFY_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    JSON_PARSE_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    extract_parameters: ExtractParameters
    json_extract_parameters: JsonExtractParameters
    replace_parameters: ReplaceParameters
    parse_parameters: ParseParameters
    allow_parameters: AllowParameters
    block_parameters: BlockParameters
    extract_timestamp_parameters: ExtractTimestampParameters
    remove_fields_parameters: RemoveFieldsParameters
    json_stringify_parameters: JsonStringifyParameters
    json_parse_parameters: JsonParseParameters
    def __init__(self, extract_parameters: _Optional[_Union[ExtractParameters, _Mapping]] = ..., json_extract_parameters: _Optional[_Union[JsonExtractParameters, _Mapping]] = ..., replace_parameters: _Optional[_Union[ReplaceParameters, _Mapping]] = ..., parse_parameters: _Optional[_Union[ParseParameters, _Mapping]] = ..., allow_parameters: _Optional[_Union[AllowParameters, _Mapping]] = ..., block_parameters: _Optional[_Union[BlockParameters, _Mapping]] = ..., extract_timestamp_parameters: _Optional[_Union[ExtractTimestampParameters, _Mapping]] = ..., remove_fields_parameters: _Optional[_Union[RemoveFieldsParameters, _Mapping]] = ..., json_stringify_parameters: _Optional[_Union[JsonStringifyParameters, _Mapping]] = ..., json_parse_parameters: _Optional[_Union[JsonParseParameters, _Mapping]] = ...) -> None: ...

class ExtractParameters(_message.Message):
    __slots__ = ("rule",)
    RULE_FIELD_NUMBER: _ClassVar[int]
    rule: _wrappers_pb2.StringValue
    def __init__(self, rule: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class JsonExtractParameters(_message.Message):
    __slots__ = ("destination_field", "rule")
    class DestinationField(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        DESTINATION_FIELD_CATEGORY_OR_UNSPECIFIED: _ClassVar[JsonExtractParameters.DestinationField]
        DESTINATION_FIELD_CLASSNAME: _ClassVar[JsonExtractParameters.DestinationField]
        DESTINATION_FIELD_METHODNAME: _ClassVar[JsonExtractParameters.DestinationField]
        DESTINATION_FIELD_THREADID: _ClassVar[JsonExtractParameters.DestinationField]
        DESTINATION_FIELD_SEVERITY: _ClassVar[JsonExtractParameters.DestinationField]
    DESTINATION_FIELD_CATEGORY_OR_UNSPECIFIED: JsonExtractParameters.DestinationField
    DESTINATION_FIELD_CLASSNAME: JsonExtractParameters.DestinationField
    DESTINATION_FIELD_METHODNAME: JsonExtractParameters.DestinationField
    DESTINATION_FIELD_THREADID: JsonExtractParameters.DestinationField
    DESTINATION_FIELD_SEVERITY: JsonExtractParameters.DestinationField
    DESTINATION_FIELD_FIELD_NUMBER: _ClassVar[int]
    RULE_FIELD_NUMBER: _ClassVar[int]
    destination_field: JsonExtractParameters.DestinationField
    rule: _wrappers_pb2.StringValue
    def __init__(self, destination_field: _Optional[_Union[JsonExtractParameters.DestinationField, str]] = ..., rule: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class ReplaceParameters(_message.Message):
    __slots__ = ("destination_field", "replace_new_val", "rule")
    DESTINATION_FIELD_FIELD_NUMBER: _ClassVar[int]
    REPLACE_NEW_VAL_FIELD_NUMBER: _ClassVar[int]
    RULE_FIELD_NUMBER: _ClassVar[int]
    destination_field: _wrappers_pb2.StringValue
    replace_new_val: _wrappers_pb2.StringValue
    rule: _wrappers_pb2.StringValue
    def __init__(self, destination_field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., replace_new_val: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., rule: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class ParseParameters(_message.Message):
    __slots__ = ("destination_field", "rule")
    DESTINATION_FIELD_FIELD_NUMBER: _ClassVar[int]
    RULE_FIELD_NUMBER: _ClassVar[int]
    destination_field: _wrappers_pb2.StringValue
    rule: _wrappers_pb2.StringValue
    def __init__(self, destination_field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., rule: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class AllowParameters(_message.Message):
    __slots__ = ("keep_blocked_logs", "rule")
    KEEP_BLOCKED_LOGS_FIELD_NUMBER: _ClassVar[int]
    RULE_FIELD_NUMBER: _ClassVar[int]
    keep_blocked_logs: _wrappers_pb2.BoolValue
    rule: _wrappers_pb2.StringValue
    def __init__(self, keep_blocked_logs: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., rule: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class BlockParameters(_message.Message):
    __slots__ = ("keep_blocked_logs", "rule")
    KEEP_BLOCKED_LOGS_FIELD_NUMBER: _ClassVar[int]
    RULE_FIELD_NUMBER: _ClassVar[int]
    keep_blocked_logs: _wrappers_pb2.BoolValue
    rule: _wrappers_pb2.StringValue
    def __init__(self, keep_blocked_logs: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., rule: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class ExtractTimestampParameters(_message.Message):
    __slots__ = ("standard", "format")
    class FormatStandard(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        FORMAT_STANDARD_STRFTIME_OR_UNSPECIFIED: _ClassVar[ExtractTimestampParameters.FormatStandard]
        FORMAT_STANDARD_JAVASDF: _ClassVar[ExtractTimestampParameters.FormatStandard]
        FORMAT_STANDARD_GOLANG: _ClassVar[ExtractTimestampParameters.FormatStandard]
        FORMAT_STANDARD_SECONDSTS: _ClassVar[ExtractTimestampParameters.FormatStandard]
        FORMAT_STANDARD_MILLITS: _ClassVar[ExtractTimestampParameters.FormatStandard]
        FORMAT_STANDARD_MICROTS: _ClassVar[ExtractTimestampParameters.FormatStandard]
        FORMAT_STANDARD_NANOTS: _ClassVar[ExtractTimestampParameters.FormatStandard]
    FORMAT_STANDARD_STRFTIME_OR_UNSPECIFIED: ExtractTimestampParameters.FormatStandard
    FORMAT_STANDARD_JAVASDF: ExtractTimestampParameters.FormatStandard
    FORMAT_STANDARD_GOLANG: ExtractTimestampParameters.FormatStandard
    FORMAT_STANDARD_SECONDSTS: ExtractTimestampParameters.FormatStandard
    FORMAT_STANDARD_MILLITS: ExtractTimestampParameters.FormatStandard
    FORMAT_STANDARD_MICROTS: ExtractTimestampParameters.FormatStandard
    FORMAT_STANDARD_NANOTS: ExtractTimestampParameters.FormatStandard
    STANDARD_FIELD_NUMBER: _ClassVar[int]
    FORMAT_FIELD_NUMBER: _ClassVar[int]
    standard: ExtractTimestampParameters.FormatStandard
    format: _wrappers_pb2.StringValue
    def __init__(self, standard: _Optional[_Union[ExtractTimestampParameters.FormatStandard, str]] = ..., format: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class RemoveFieldsParameters(_message.Message):
    __slots__ = ("fields",)
    FIELDS_FIELD_NUMBER: _ClassVar[int]
    fields: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, fields: _Optional[_Iterable[str]] = ...) -> None: ...

class JsonStringifyParameters(_message.Message):
    __slots__ = ("destination_field", "delete_source")
    DESTINATION_FIELD_FIELD_NUMBER: _ClassVar[int]
    DELETE_SOURCE_FIELD_NUMBER: _ClassVar[int]
    destination_field: _wrappers_pb2.StringValue
    delete_source: _wrappers_pb2.BoolValue
    def __init__(self, destination_field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., delete_source: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...

class JsonParseParameters(_message.Message):
    __slots__ = ("destination_field", "delete_source", "escaped_value", "override_dest")
    DESTINATION_FIELD_FIELD_NUMBER: _ClassVar[int]
    DELETE_SOURCE_FIELD_NUMBER: _ClassVar[int]
    ESCAPED_VALUE_FIELD_NUMBER: _ClassVar[int]
    OVERRIDE_DEST_FIELD_NUMBER: _ClassVar[int]
    destination_field: _wrappers_pb2.StringValue
    delete_source: _wrappers_pb2.BoolValue
    escaped_value: _wrappers_pb2.BoolValue
    override_dest: _wrappers_pb2.BoolValue
    def __init__(self, destination_field: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., delete_source: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., escaped_value: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., override_dest: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
