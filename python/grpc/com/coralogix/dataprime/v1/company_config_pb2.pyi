from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetCompanyConfigRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetCompanyConfigResponse(_message.Message):
    __slots__ = ("config",)
    CONFIG_FIELD_NUMBER: _ClassVar[int]
    config: CompanyConfig
    def __init__(self, config: _Optional[_Union[CompanyConfig, _Mapping]] = ...) -> None: ...

class CompanyConfig(_message.Message):
    __slots__ = ("migrated_companies", "max_teams_per_query")
    MIGRATED_COMPANIES_FIELD_NUMBER: _ClassVar[int]
    MAX_TEAMS_PER_QUERY_FIELD_NUMBER: _ClassVar[int]
    migrated_companies: _containers.RepeatedCompositeFieldContainer[MigratedCompany]
    max_teams_per_query: int
    def __init__(self, migrated_companies: _Optional[_Iterable[_Union[MigratedCompany, _Mapping]]] = ..., max_teams_per_query: _Optional[int] = ...) -> None: ...

class MigratedCompany(_message.Message):
    __slots__ = ("company_id", "cutoff")
    COMPANY_ID_FIELD_NUMBER: _ClassVar[int]
    CUTOFF_FIELD_NUMBER: _ClassVar[int]
    company_id: int
    cutoff: _timestamp_pb2.Timestamp
    def __init__(self, company_id: _Optional[int] = ..., cutoff: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
