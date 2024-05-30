from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class UpsertIncidentStateType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    UPSERT_INCIDENT_STATE_TYPE_UNSPECIFIED: _ClassVar[UpsertIncidentStateType]
    UPSERT_INCIDENT_STATE_TYPE_TRIGGERED: _ClassVar[UpsertIncidentStateType]
    UPSERT_INCIDENT_STATE_TYPE_RESOLVED: _ClassVar[UpsertIncidentStateType]
UPSERT_INCIDENT_STATE_TYPE_UNSPECIFIED: UpsertIncidentStateType
UPSERT_INCIDENT_STATE_TYPE_TRIGGERED: UpsertIncidentStateType
UPSERT_INCIDENT_STATE_TYPE_RESOLVED: UpsertIncidentStateType
