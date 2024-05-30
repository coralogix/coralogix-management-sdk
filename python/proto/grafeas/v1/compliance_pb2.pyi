from grafeas.v1 import severity_pb2 as _severity_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ComplianceNote(_message.Message):
    __slots__ = ("title", "description", "version", "rationale", "remediation", "cis_benchmark", "scan_instructions", "impact")
    class CisBenchmark(_message.Message):
        __slots__ = ("profile_level", "severity")
        PROFILE_LEVEL_FIELD_NUMBER: _ClassVar[int]
        SEVERITY_FIELD_NUMBER: _ClassVar[int]
        profile_level: int
        severity: _severity_pb2.Severity
        def __init__(self, profile_level: _Optional[int] = ..., severity: _Optional[_Union[_severity_pb2.Severity, str]] = ...) -> None: ...
    TITLE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    RATIONALE_FIELD_NUMBER: _ClassVar[int]
    REMEDIATION_FIELD_NUMBER: _ClassVar[int]
    CIS_BENCHMARK_FIELD_NUMBER: _ClassVar[int]
    SCAN_INSTRUCTIONS_FIELD_NUMBER: _ClassVar[int]
    IMPACT_FIELD_NUMBER: _ClassVar[int]
    title: str
    description: str
    version: _containers.RepeatedCompositeFieldContainer[ComplianceVersion]
    rationale: str
    remediation: str
    cis_benchmark: ComplianceNote.CisBenchmark
    scan_instructions: bytes
    impact: str
    def __init__(self, title: _Optional[str] = ..., description: _Optional[str] = ..., version: _Optional[_Iterable[_Union[ComplianceVersion, _Mapping]]] = ..., rationale: _Optional[str] = ..., remediation: _Optional[str] = ..., cis_benchmark: _Optional[_Union[ComplianceNote.CisBenchmark, _Mapping]] = ..., scan_instructions: _Optional[bytes] = ..., impact: _Optional[str] = ...) -> None: ...

class ComplianceVersion(_message.Message):
    __slots__ = ("cpe_uri", "benchmark_document", "version")
    CPE_URI_FIELD_NUMBER: _ClassVar[int]
    BENCHMARK_DOCUMENT_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    cpe_uri: str
    benchmark_document: str
    version: str
    def __init__(self, cpe_uri: _Optional[str] = ..., benchmark_document: _Optional[str] = ..., version: _Optional[str] = ...) -> None: ...

class ComplianceOccurrence(_message.Message):
    __slots__ = ("non_compliant_files", "non_compliance_reason")
    NON_COMPLIANT_FILES_FIELD_NUMBER: _ClassVar[int]
    NON_COMPLIANCE_REASON_FIELD_NUMBER: _ClassVar[int]
    non_compliant_files: _containers.RepeatedCompositeFieldContainer[NonCompliantFile]
    non_compliance_reason: str
    def __init__(self, non_compliant_files: _Optional[_Iterable[_Union[NonCompliantFile, _Mapping]]] = ..., non_compliance_reason: _Optional[str] = ...) -> None: ...

class NonCompliantFile(_message.Message):
    __slots__ = ("path", "display_command", "reason")
    PATH_FIELD_NUMBER: _ClassVar[int]
    DISPLAY_COMMAND_FIELD_NUMBER: _ClassVar[int]
    REASON_FIELD_NUMBER: _ClassVar[int]
    path: str
    display_command: str
    reason: str
    def __init__(self, path: _Optional[str] = ..., display_command: _Optional[str] = ..., reason: _Optional[str] = ...) -> None: ...
