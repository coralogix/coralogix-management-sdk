from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CVSSVersion(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    CVSS_VERSION_UNSPECIFIED: _ClassVar[CVSSVersion]
    CVSS_VERSION_2: _ClassVar[CVSSVersion]
    CVSS_VERSION_3: _ClassVar[CVSSVersion]
CVSS_VERSION_UNSPECIFIED: CVSSVersion
CVSS_VERSION_2: CVSSVersion
CVSS_VERSION_3: CVSSVersion

class CVSSv3(_message.Message):
    __slots__ = ("base_score", "exploitability_score", "impact_score", "attack_vector", "attack_complexity", "privileges_required", "user_interaction", "scope", "confidentiality_impact", "integrity_impact", "availability_impact")
    class AttackVector(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        ATTACK_VECTOR_UNSPECIFIED: _ClassVar[CVSSv3.AttackVector]
        ATTACK_VECTOR_NETWORK: _ClassVar[CVSSv3.AttackVector]
        ATTACK_VECTOR_ADJACENT: _ClassVar[CVSSv3.AttackVector]
        ATTACK_VECTOR_LOCAL: _ClassVar[CVSSv3.AttackVector]
        ATTACK_VECTOR_PHYSICAL: _ClassVar[CVSSv3.AttackVector]
    ATTACK_VECTOR_UNSPECIFIED: CVSSv3.AttackVector
    ATTACK_VECTOR_NETWORK: CVSSv3.AttackVector
    ATTACK_VECTOR_ADJACENT: CVSSv3.AttackVector
    ATTACK_VECTOR_LOCAL: CVSSv3.AttackVector
    ATTACK_VECTOR_PHYSICAL: CVSSv3.AttackVector
    class AttackComplexity(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        ATTACK_COMPLEXITY_UNSPECIFIED: _ClassVar[CVSSv3.AttackComplexity]
        ATTACK_COMPLEXITY_LOW: _ClassVar[CVSSv3.AttackComplexity]
        ATTACK_COMPLEXITY_HIGH: _ClassVar[CVSSv3.AttackComplexity]
    ATTACK_COMPLEXITY_UNSPECIFIED: CVSSv3.AttackComplexity
    ATTACK_COMPLEXITY_LOW: CVSSv3.AttackComplexity
    ATTACK_COMPLEXITY_HIGH: CVSSv3.AttackComplexity
    class PrivilegesRequired(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        PRIVILEGES_REQUIRED_UNSPECIFIED: _ClassVar[CVSSv3.PrivilegesRequired]
        PRIVILEGES_REQUIRED_NONE: _ClassVar[CVSSv3.PrivilegesRequired]
        PRIVILEGES_REQUIRED_LOW: _ClassVar[CVSSv3.PrivilegesRequired]
        PRIVILEGES_REQUIRED_HIGH: _ClassVar[CVSSv3.PrivilegesRequired]
    PRIVILEGES_REQUIRED_UNSPECIFIED: CVSSv3.PrivilegesRequired
    PRIVILEGES_REQUIRED_NONE: CVSSv3.PrivilegesRequired
    PRIVILEGES_REQUIRED_LOW: CVSSv3.PrivilegesRequired
    PRIVILEGES_REQUIRED_HIGH: CVSSv3.PrivilegesRequired
    class UserInteraction(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        USER_INTERACTION_UNSPECIFIED: _ClassVar[CVSSv3.UserInteraction]
        USER_INTERACTION_NONE: _ClassVar[CVSSv3.UserInteraction]
        USER_INTERACTION_REQUIRED: _ClassVar[CVSSv3.UserInteraction]
    USER_INTERACTION_UNSPECIFIED: CVSSv3.UserInteraction
    USER_INTERACTION_NONE: CVSSv3.UserInteraction
    USER_INTERACTION_REQUIRED: CVSSv3.UserInteraction
    class Scope(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        SCOPE_UNSPECIFIED: _ClassVar[CVSSv3.Scope]
        SCOPE_UNCHANGED: _ClassVar[CVSSv3.Scope]
        SCOPE_CHANGED: _ClassVar[CVSSv3.Scope]
    SCOPE_UNSPECIFIED: CVSSv3.Scope
    SCOPE_UNCHANGED: CVSSv3.Scope
    SCOPE_CHANGED: CVSSv3.Scope
    class Impact(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        IMPACT_UNSPECIFIED: _ClassVar[CVSSv3.Impact]
        IMPACT_HIGH: _ClassVar[CVSSv3.Impact]
        IMPACT_LOW: _ClassVar[CVSSv3.Impact]
        IMPACT_NONE: _ClassVar[CVSSv3.Impact]
    IMPACT_UNSPECIFIED: CVSSv3.Impact
    IMPACT_HIGH: CVSSv3.Impact
    IMPACT_LOW: CVSSv3.Impact
    IMPACT_NONE: CVSSv3.Impact
    BASE_SCORE_FIELD_NUMBER: _ClassVar[int]
    EXPLOITABILITY_SCORE_FIELD_NUMBER: _ClassVar[int]
    IMPACT_SCORE_FIELD_NUMBER: _ClassVar[int]
    ATTACK_VECTOR_FIELD_NUMBER: _ClassVar[int]
    ATTACK_COMPLEXITY_FIELD_NUMBER: _ClassVar[int]
    PRIVILEGES_REQUIRED_FIELD_NUMBER: _ClassVar[int]
    USER_INTERACTION_FIELD_NUMBER: _ClassVar[int]
    SCOPE_FIELD_NUMBER: _ClassVar[int]
    CONFIDENTIALITY_IMPACT_FIELD_NUMBER: _ClassVar[int]
    INTEGRITY_IMPACT_FIELD_NUMBER: _ClassVar[int]
    AVAILABILITY_IMPACT_FIELD_NUMBER: _ClassVar[int]
    base_score: float
    exploitability_score: float
    impact_score: float
    attack_vector: CVSSv3.AttackVector
    attack_complexity: CVSSv3.AttackComplexity
    privileges_required: CVSSv3.PrivilegesRequired
    user_interaction: CVSSv3.UserInteraction
    scope: CVSSv3.Scope
    confidentiality_impact: CVSSv3.Impact
    integrity_impact: CVSSv3.Impact
    availability_impact: CVSSv3.Impact
    def __init__(self, base_score: _Optional[float] = ..., exploitability_score: _Optional[float] = ..., impact_score: _Optional[float] = ..., attack_vector: _Optional[_Union[CVSSv3.AttackVector, str]] = ..., attack_complexity: _Optional[_Union[CVSSv3.AttackComplexity, str]] = ..., privileges_required: _Optional[_Union[CVSSv3.PrivilegesRequired, str]] = ..., user_interaction: _Optional[_Union[CVSSv3.UserInteraction, str]] = ..., scope: _Optional[_Union[CVSSv3.Scope, str]] = ..., confidentiality_impact: _Optional[_Union[CVSSv3.Impact, str]] = ..., integrity_impact: _Optional[_Union[CVSSv3.Impact, str]] = ..., availability_impact: _Optional[_Union[CVSSv3.Impact, str]] = ...) -> None: ...

class CVSS(_message.Message):
    __slots__ = ("base_score", "exploitability_score", "impact_score", "attack_vector", "attack_complexity", "authentication", "privileges_required", "user_interaction", "scope", "confidentiality_impact", "integrity_impact", "availability_impact")
    class AttackVector(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        ATTACK_VECTOR_UNSPECIFIED: _ClassVar[CVSS.AttackVector]
        ATTACK_VECTOR_NETWORK: _ClassVar[CVSS.AttackVector]
        ATTACK_VECTOR_ADJACENT: _ClassVar[CVSS.AttackVector]
        ATTACK_VECTOR_LOCAL: _ClassVar[CVSS.AttackVector]
        ATTACK_VECTOR_PHYSICAL: _ClassVar[CVSS.AttackVector]
    ATTACK_VECTOR_UNSPECIFIED: CVSS.AttackVector
    ATTACK_VECTOR_NETWORK: CVSS.AttackVector
    ATTACK_VECTOR_ADJACENT: CVSS.AttackVector
    ATTACK_VECTOR_LOCAL: CVSS.AttackVector
    ATTACK_VECTOR_PHYSICAL: CVSS.AttackVector
    class AttackComplexity(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        ATTACK_COMPLEXITY_UNSPECIFIED: _ClassVar[CVSS.AttackComplexity]
        ATTACK_COMPLEXITY_LOW: _ClassVar[CVSS.AttackComplexity]
        ATTACK_COMPLEXITY_HIGH: _ClassVar[CVSS.AttackComplexity]
        ATTACK_COMPLEXITY_MEDIUM: _ClassVar[CVSS.AttackComplexity]
    ATTACK_COMPLEXITY_UNSPECIFIED: CVSS.AttackComplexity
    ATTACK_COMPLEXITY_LOW: CVSS.AttackComplexity
    ATTACK_COMPLEXITY_HIGH: CVSS.AttackComplexity
    ATTACK_COMPLEXITY_MEDIUM: CVSS.AttackComplexity
    class Authentication(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        AUTHENTICATION_UNSPECIFIED: _ClassVar[CVSS.Authentication]
        AUTHENTICATION_MULTIPLE: _ClassVar[CVSS.Authentication]
        AUTHENTICATION_SINGLE: _ClassVar[CVSS.Authentication]
        AUTHENTICATION_NONE: _ClassVar[CVSS.Authentication]
    AUTHENTICATION_UNSPECIFIED: CVSS.Authentication
    AUTHENTICATION_MULTIPLE: CVSS.Authentication
    AUTHENTICATION_SINGLE: CVSS.Authentication
    AUTHENTICATION_NONE: CVSS.Authentication
    class PrivilegesRequired(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        PRIVILEGES_REQUIRED_UNSPECIFIED: _ClassVar[CVSS.PrivilegesRequired]
        PRIVILEGES_REQUIRED_NONE: _ClassVar[CVSS.PrivilegesRequired]
        PRIVILEGES_REQUIRED_LOW: _ClassVar[CVSS.PrivilegesRequired]
        PRIVILEGES_REQUIRED_HIGH: _ClassVar[CVSS.PrivilegesRequired]
    PRIVILEGES_REQUIRED_UNSPECIFIED: CVSS.PrivilegesRequired
    PRIVILEGES_REQUIRED_NONE: CVSS.PrivilegesRequired
    PRIVILEGES_REQUIRED_LOW: CVSS.PrivilegesRequired
    PRIVILEGES_REQUIRED_HIGH: CVSS.PrivilegesRequired
    class UserInteraction(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        USER_INTERACTION_UNSPECIFIED: _ClassVar[CVSS.UserInteraction]
        USER_INTERACTION_NONE: _ClassVar[CVSS.UserInteraction]
        USER_INTERACTION_REQUIRED: _ClassVar[CVSS.UserInteraction]
    USER_INTERACTION_UNSPECIFIED: CVSS.UserInteraction
    USER_INTERACTION_NONE: CVSS.UserInteraction
    USER_INTERACTION_REQUIRED: CVSS.UserInteraction
    class Scope(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        SCOPE_UNSPECIFIED: _ClassVar[CVSS.Scope]
        SCOPE_UNCHANGED: _ClassVar[CVSS.Scope]
        SCOPE_CHANGED: _ClassVar[CVSS.Scope]
    SCOPE_UNSPECIFIED: CVSS.Scope
    SCOPE_UNCHANGED: CVSS.Scope
    SCOPE_CHANGED: CVSS.Scope
    class Impact(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        IMPACT_UNSPECIFIED: _ClassVar[CVSS.Impact]
        IMPACT_HIGH: _ClassVar[CVSS.Impact]
        IMPACT_LOW: _ClassVar[CVSS.Impact]
        IMPACT_NONE: _ClassVar[CVSS.Impact]
        IMPACT_PARTIAL: _ClassVar[CVSS.Impact]
        IMPACT_COMPLETE: _ClassVar[CVSS.Impact]
    IMPACT_UNSPECIFIED: CVSS.Impact
    IMPACT_HIGH: CVSS.Impact
    IMPACT_LOW: CVSS.Impact
    IMPACT_NONE: CVSS.Impact
    IMPACT_PARTIAL: CVSS.Impact
    IMPACT_COMPLETE: CVSS.Impact
    BASE_SCORE_FIELD_NUMBER: _ClassVar[int]
    EXPLOITABILITY_SCORE_FIELD_NUMBER: _ClassVar[int]
    IMPACT_SCORE_FIELD_NUMBER: _ClassVar[int]
    ATTACK_VECTOR_FIELD_NUMBER: _ClassVar[int]
    ATTACK_COMPLEXITY_FIELD_NUMBER: _ClassVar[int]
    AUTHENTICATION_FIELD_NUMBER: _ClassVar[int]
    PRIVILEGES_REQUIRED_FIELD_NUMBER: _ClassVar[int]
    USER_INTERACTION_FIELD_NUMBER: _ClassVar[int]
    SCOPE_FIELD_NUMBER: _ClassVar[int]
    CONFIDENTIALITY_IMPACT_FIELD_NUMBER: _ClassVar[int]
    INTEGRITY_IMPACT_FIELD_NUMBER: _ClassVar[int]
    AVAILABILITY_IMPACT_FIELD_NUMBER: _ClassVar[int]
    base_score: float
    exploitability_score: float
    impact_score: float
    attack_vector: CVSS.AttackVector
    attack_complexity: CVSS.AttackComplexity
    authentication: CVSS.Authentication
    privileges_required: CVSS.PrivilegesRequired
    user_interaction: CVSS.UserInteraction
    scope: CVSS.Scope
    confidentiality_impact: CVSS.Impact
    integrity_impact: CVSS.Impact
    availability_impact: CVSS.Impact
    def __init__(self, base_score: _Optional[float] = ..., exploitability_score: _Optional[float] = ..., impact_score: _Optional[float] = ..., attack_vector: _Optional[_Union[CVSS.AttackVector, str]] = ..., attack_complexity: _Optional[_Union[CVSS.AttackComplexity, str]] = ..., authentication: _Optional[_Union[CVSS.Authentication, str]] = ..., privileges_required: _Optional[_Union[CVSS.PrivilegesRequired, str]] = ..., user_interaction: _Optional[_Union[CVSS.UserInteraction, str]] = ..., scope: _Optional[_Union[CVSS.Scope, str]] = ..., confidentiality_impact: _Optional[_Union[CVSS.Impact, str]] = ..., integrity_impact: _Optional[_Union[CVSS.Impact, str]] = ..., availability_impact: _Optional[_Union[CVSS.Impact, str]] = ...) -> None: ...
