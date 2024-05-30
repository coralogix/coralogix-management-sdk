from grafeas.v1 import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class VulnerabilityAssessmentNote(_message.Message):
    __slots__ = ("title", "short_description", "long_description", "language_code", "publisher", "product", "assessment")
    class Publisher(_message.Message):
        __slots__ = ("name", "issuing_authority", "publisher_namespace")
        NAME_FIELD_NUMBER: _ClassVar[int]
        ISSUING_AUTHORITY_FIELD_NUMBER: _ClassVar[int]
        PUBLISHER_NAMESPACE_FIELD_NUMBER: _ClassVar[int]
        name: str
        issuing_authority: str
        publisher_namespace: str
        def __init__(self, name: _Optional[str] = ..., issuing_authority: _Optional[str] = ..., publisher_namespace: _Optional[str] = ...) -> None: ...
    class Product(_message.Message):
        __slots__ = ("name", "id", "generic_uri")
        NAME_FIELD_NUMBER: _ClassVar[int]
        ID_FIELD_NUMBER: _ClassVar[int]
        GENERIC_URI_FIELD_NUMBER: _ClassVar[int]
        name: str
        id: str
        generic_uri: str
        def __init__(self, name: _Optional[str] = ..., id: _Optional[str] = ..., generic_uri: _Optional[str] = ...) -> None: ...
    class Assessment(_message.Message):
        __slots__ = ("cve", "vulnerability_id", "short_description", "long_description", "related_uris", "state", "impacts", "justification", "remediations")
        class State(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            STATE_UNSPECIFIED: _ClassVar[VulnerabilityAssessmentNote.Assessment.State]
            AFFECTED: _ClassVar[VulnerabilityAssessmentNote.Assessment.State]
            NOT_AFFECTED: _ClassVar[VulnerabilityAssessmentNote.Assessment.State]
            FIXED: _ClassVar[VulnerabilityAssessmentNote.Assessment.State]
            UNDER_INVESTIGATION: _ClassVar[VulnerabilityAssessmentNote.Assessment.State]
        STATE_UNSPECIFIED: VulnerabilityAssessmentNote.Assessment.State
        AFFECTED: VulnerabilityAssessmentNote.Assessment.State
        NOT_AFFECTED: VulnerabilityAssessmentNote.Assessment.State
        FIXED: VulnerabilityAssessmentNote.Assessment.State
        UNDER_INVESTIGATION: VulnerabilityAssessmentNote.Assessment.State
        class Justification(_message.Message):
            __slots__ = ("justification_type", "details")
            class JustificationType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
                __slots__ = ()
                JUSTIFICATION_TYPE_UNSPECIFIED: _ClassVar[VulnerabilityAssessmentNote.Assessment.Justification.JustificationType]
                COMPONENT_NOT_PRESENT: _ClassVar[VulnerabilityAssessmentNote.Assessment.Justification.JustificationType]
                VULNERABLE_CODE_NOT_PRESENT: _ClassVar[VulnerabilityAssessmentNote.Assessment.Justification.JustificationType]
                VULNERABLE_CODE_NOT_IN_EXECUTE_PATH: _ClassVar[VulnerabilityAssessmentNote.Assessment.Justification.JustificationType]
                VULNERABLE_CODE_CANNOT_BE_CONTROLLED_BY_ADVERSARY: _ClassVar[VulnerabilityAssessmentNote.Assessment.Justification.JustificationType]
                INLINE_MITIGATIONS_ALREADY_EXIST: _ClassVar[VulnerabilityAssessmentNote.Assessment.Justification.JustificationType]
            JUSTIFICATION_TYPE_UNSPECIFIED: VulnerabilityAssessmentNote.Assessment.Justification.JustificationType
            COMPONENT_NOT_PRESENT: VulnerabilityAssessmentNote.Assessment.Justification.JustificationType
            VULNERABLE_CODE_NOT_PRESENT: VulnerabilityAssessmentNote.Assessment.Justification.JustificationType
            VULNERABLE_CODE_NOT_IN_EXECUTE_PATH: VulnerabilityAssessmentNote.Assessment.Justification.JustificationType
            VULNERABLE_CODE_CANNOT_BE_CONTROLLED_BY_ADVERSARY: VulnerabilityAssessmentNote.Assessment.Justification.JustificationType
            INLINE_MITIGATIONS_ALREADY_EXIST: VulnerabilityAssessmentNote.Assessment.Justification.JustificationType
            JUSTIFICATION_TYPE_FIELD_NUMBER: _ClassVar[int]
            DETAILS_FIELD_NUMBER: _ClassVar[int]
            justification_type: VulnerabilityAssessmentNote.Assessment.Justification.JustificationType
            details: str
            def __init__(self, justification_type: _Optional[_Union[VulnerabilityAssessmentNote.Assessment.Justification.JustificationType, str]] = ..., details: _Optional[str] = ...) -> None: ...
        class Remediation(_message.Message):
            __slots__ = ("remediation_type", "details", "remediation_uri")
            class RemediationType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
                __slots__ = ()
                REMEDIATION_TYPE_UNSPECIFIED: _ClassVar[VulnerabilityAssessmentNote.Assessment.Remediation.RemediationType]
                MITIGATION: _ClassVar[VulnerabilityAssessmentNote.Assessment.Remediation.RemediationType]
                NO_FIX_PLANNED: _ClassVar[VulnerabilityAssessmentNote.Assessment.Remediation.RemediationType]
                NONE_AVAILABLE: _ClassVar[VulnerabilityAssessmentNote.Assessment.Remediation.RemediationType]
                VENDOR_FIX: _ClassVar[VulnerabilityAssessmentNote.Assessment.Remediation.RemediationType]
                WORKAROUND: _ClassVar[VulnerabilityAssessmentNote.Assessment.Remediation.RemediationType]
            REMEDIATION_TYPE_UNSPECIFIED: VulnerabilityAssessmentNote.Assessment.Remediation.RemediationType
            MITIGATION: VulnerabilityAssessmentNote.Assessment.Remediation.RemediationType
            NO_FIX_PLANNED: VulnerabilityAssessmentNote.Assessment.Remediation.RemediationType
            NONE_AVAILABLE: VulnerabilityAssessmentNote.Assessment.Remediation.RemediationType
            VENDOR_FIX: VulnerabilityAssessmentNote.Assessment.Remediation.RemediationType
            WORKAROUND: VulnerabilityAssessmentNote.Assessment.Remediation.RemediationType
            REMEDIATION_TYPE_FIELD_NUMBER: _ClassVar[int]
            DETAILS_FIELD_NUMBER: _ClassVar[int]
            REMEDIATION_URI_FIELD_NUMBER: _ClassVar[int]
            remediation_type: VulnerabilityAssessmentNote.Assessment.Remediation.RemediationType
            details: str
            remediation_uri: _common_pb2.RelatedUrl
            def __init__(self, remediation_type: _Optional[_Union[VulnerabilityAssessmentNote.Assessment.Remediation.RemediationType, str]] = ..., details: _Optional[str] = ..., remediation_uri: _Optional[_Union[_common_pb2.RelatedUrl, _Mapping]] = ...) -> None: ...
        CVE_FIELD_NUMBER: _ClassVar[int]
        VULNERABILITY_ID_FIELD_NUMBER: _ClassVar[int]
        SHORT_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
        LONG_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
        RELATED_URIS_FIELD_NUMBER: _ClassVar[int]
        STATE_FIELD_NUMBER: _ClassVar[int]
        IMPACTS_FIELD_NUMBER: _ClassVar[int]
        JUSTIFICATION_FIELD_NUMBER: _ClassVar[int]
        REMEDIATIONS_FIELD_NUMBER: _ClassVar[int]
        cve: str
        vulnerability_id: str
        short_description: str
        long_description: str
        related_uris: _containers.RepeatedCompositeFieldContainer[_common_pb2.RelatedUrl]
        state: VulnerabilityAssessmentNote.Assessment.State
        impacts: _containers.RepeatedScalarFieldContainer[str]
        justification: VulnerabilityAssessmentNote.Assessment.Justification
        remediations: _containers.RepeatedCompositeFieldContainer[VulnerabilityAssessmentNote.Assessment.Remediation]
        def __init__(self, cve: _Optional[str] = ..., vulnerability_id: _Optional[str] = ..., short_description: _Optional[str] = ..., long_description: _Optional[str] = ..., related_uris: _Optional[_Iterable[_Union[_common_pb2.RelatedUrl, _Mapping]]] = ..., state: _Optional[_Union[VulnerabilityAssessmentNote.Assessment.State, str]] = ..., impacts: _Optional[_Iterable[str]] = ..., justification: _Optional[_Union[VulnerabilityAssessmentNote.Assessment.Justification, _Mapping]] = ..., remediations: _Optional[_Iterable[_Union[VulnerabilityAssessmentNote.Assessment.Remediation, _Mapping]]] = ...) -> None: ...
    TITLE_FIELD_NUMBER: _ClassVar[int]
    SHORT_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    LONG_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    LANGUAGE_CODE_FIELD_NUMBER: _ClassVar[int]
    PUBLISHER_FIELD_NUMBER: _ClassVar[int]
    PRODUCT_FIELD_NUMBER: _ClassVar[int]
    ASSESSMENT_FIELD_NUMBER: _ClassVar[int]
    title: str
    short_description: str
    long_description: str
    language_code: str
    publisher: VulnerabilityAssessmentNote.Publisher
    product: VulnerabilityAssessmentNote.Product
    assessment: VulnerabilityAssessmentNote.Assessment
    def __init__(self, title: _Optional[str] = ..., short_description: _Optional[str] = ..., long_description: _Optional[str] = ..., language_code: _Optional[str] = ..., publisher: _Optional[_Union[VulnerabilityAssessmentNote.Publisher, _Mapping]] = ..., product: _Optional[_Union[VulnerabilityAssessmentNote.Product, _Mapping]] = ..., assessment: _Optional[_Union[VulnerabilityAssessmentNote.Assessment, _Mapping]] = ...) -> None: ...
