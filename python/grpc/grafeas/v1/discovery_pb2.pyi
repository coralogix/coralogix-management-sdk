from google.api import field_behavior_pb2 as _field_behavior_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.rpc import status_pb2 as _status_pb2
from grafeas.v1 import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DiscoveryNote(_message.Message):
    __slots__ = ("analysis_kind",)
    ANALYSIS_KIND_FIELD_NUMBER: _ClassVar[int]
    analysis_kind: _common_pb2.NoteKind
    def __init__(self, analysis_kind: _Optional[_Union[_common_pb2.NoteKind, str]] = ...) -> None: ...

class DiscoveryOccurrence(_message.Message):
    __slots__ = ("continuous_analysis", "analysis_status", "analysis_completed", "analysis_error", "analysis_status_error", "cpe", "last_scan_time", "archive_time", "sbom_status")
    class ContinuousAnalysis(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        CONTINUOUS_ANALYSIS_UNSPECIFIED: _ClassVar[DiscoveryOccurrence.ContinuousAnalysis]
        ACTIVE: _ClassVar[DiscoveryOccurrence.ContinuousAnalysis]
        INACTIVE: _ClassVar[DiscoveryOccurrence.ContinuousAnalysis]
    CONTINUOUS_ANALYSIS_UNSPECIFIED: DiscoveryOccurrence.ContinuousAnalysis
    ACTIVE: DiscoveryOccurrence.ContinuousAnalysis
    INACTIVE: DiscoveryOccurrence.ContinuousAnalysis
    class AnalysisStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        ANALYSIS_STATUS_UNSPECIFIED: _ClassVar[DiscoveryOccurrence.AnalysisStatus]
        PENDING: _ClassVar[DiscoveryOccurrence.AnalysisStatus]
        SCANNING: _ClassVar[DiscoveryOccurrence.AnalysisStatus]
        FINISHED_SUCCESS: _ClassVar[DiscoveryOccurrence.AnalysisStatus]
        COMPLETE: _ClassVar[DiscoveryOccurrence.AnalysisStatus]
        FINISHED_FAILED: _ClassVar[DiscoveryOccurrence.AnalysisStatus]
        FINISHED_UNSUPPORTED: _ClassVar[DiscoveryOccurrence.AnalysisStatus]
    ANALYSIS_STATUS_UNSPECIFIED: DiscoveryOccurrence.AnalysisStatus
    PENDING: DiscoveryOccurrence.AnalysisStatus
    SCANNING: DiscoveryOccurrence.AnalysisStatus
    FINISHED_SUCCESS: DiscoveryOccurrence.AnalysisStatus
    COMPLETE: DiscoveryOccurrence.AnalysisStatus
    FINISHED_FAILED: DiscoveryOccurrence.AnalysisStatus
    FINISHED_UNSUPPORTED: DiscoveryOccurrence.AnalysisStatus
    class AnalysisCompleted(_message.Message):
        __slots__ = ("analysis_type",)
        ANALYSIS_TYPE_FIELD_NUMBER: _ClassVar[int]
        analysis_type: _containers.RepeatedScalarFieldContainer[str]
        def __init__(self, analysis_type: _Optional[_Iterable[str]] = ...) -> None: ...
    class SBOMStatus(_message.Message):
        __slots__ = ("sbom_state", "error")
        class SBOMState(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            SBOM_STATE_UNSPECIFIED: _ClassVar[DiscoveryOccurrence.SBOMStatus.SBOMState]
            PENDING: _ClassVar[DiscoveryOccurrence.SBOMStatus.SBOMState]
            COMPLETE: _ClassVar[DiscoveryOccurrence.SBOMStatus.SBOMState]
        SBOM_STATE_UNSPECIFIED: DiscoveryOccurrence.SBOMStatus.SBOMState
        PENDING: DiscoveryOccurrence.SBOMStatus.SBOMState
        COMPLETE: DiscoveryOccurrence.SBOMStatus.SBOMState
        SBOM_STATE_FIELD_NUMBER: _ClassVar[int]
        ERROR_FIELD_NUMBER: _ClassVar[int]
        sbom_state: DiscoveryOccurrence.SBOMStatus.SBOMState
        error: str
        def __init__(self, sbom_state: _Optional[_Union[DiscoveryOccurrence.SBOMStatus.SBOMState, str]] = ..., error: _Optional[str] = ...) -> None: ...
    CONTINUOUS_ANALYSIS_FIELD_NUMBER: _ClassVar[int]
    ANALYSIS_STATUS_FIELD_NUMBER: _ClassVar[int]
    ANALYSIS_COMPLETED_FIELD_NUMBER: _ClassVar[int]
    ANALYSIS_ERROR_FIELD_NUMBER: _ClassVar[int]
    ANALYSIS_STATUS_ERROR_FIELD_NUMBER: _ClassVar[int]
    CPE_FIELD_NUMBER: _ClassVar[int]
    LAST_SCAN_TIME_FIELD_NUMBER: _ClassVar[int]
    ARCHIVE_TIME_FIELD_NUMBER: _ClassVar[int]
    SBOM_STATUS_FIELD_NUMBER: _ClassVar[int]
    continuous_analysis: DiscoveryOccurrence.ContinuousAnalysis
    analysis_status: DiscoveryOccurrence.AnalysisStatus
    analysis_completed: DiscoveryOccurrence.AnalysisCompleted
    analysis_error: _containers.RepeatedCompositeFieldContainer[_status_pb2.Status]
    analysis_status_error: _status_pb2.Status
    cpe: str
    last_scan_time: _timestamp_pb2.Timestamp
    archive_time: _timestamp_pb2.Timestamp
    sbom_status: DiscoveryOccurrence.SBOMStatus
    def __init__(self, continuous_analysis: _Optional[_Union[DiscoveryOccurrence.ContinuousAnalysis, str]] = ..., analysis_status: _Optional[_Union[DiscoveryOccurrence.AnalysisStatus, str]] = ..., analysis_completed: _Optional[_Union[DiscoveryOccurrence.AnalysisCompleted, _Mapping]] = ..., analysis_error: _Optional[_Iterable[_Union[_status_pb2.Status, _Mapping]]] = ..., analysis_status_error: _Optional[_Union[_status_pb2.Status, _Mapping]] = ..., cpe: _Optional[str] = ..., last_scan_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., archive_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., sbom_status: _Optional[_Union[DiscoveryOccurrence.SBOMStatus, _Mapping]] = ...) -> None: ...
