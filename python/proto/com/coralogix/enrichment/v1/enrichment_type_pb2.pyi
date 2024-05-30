from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class EnrichmentType(_message.Message):
    __slots__ = ("geo_ip", "suspicious_ip", "aws", "custom_enrichment")
    GEO_IP_FIELD_NUMBER: _ClassVar[int]
    SUSPICIOUS_IP_FIELD_NUMBER: _ClassVar[int]
    AWS_FIELD_NUMBER: _ClassVar[int]
    CUSTOM_ENRICHMENT_FIELD_NUMBER: _ClassVar[int]
    geo_ip: GeoIpType
    suspicious_ip: SuspiciousIpType
    aws: AwsType
    custom_enrichment: CustomEnrichmentType
    def __init__(self, geo_ip: _Optional[_Union[GeoIpType, _Mapping]] = ..., suspicious_ip: _Optional[_Union[SuspiciousIpType, _Mapping]] = ..., aws: _Optional[_Union[AwsType, _Mapping]] = ..., custom_enrichment: _Optional[_Union[CustomEnrichmentType, _Mapping]] = ...) -> None: ...

class GeoIpType(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class SuspiciousIpType(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class AwsType(_message.Message):
    __slots__ = ("resource_type",)
    RESOURCE_TYPE_FIELD_NUMBER: _ClassVar[int]
    resource_type: _wrappers_pb2.StringValue
    def __init__(self, resource_type: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class CustomEnrichmentType(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.UInt32Value
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.UInt32Value, _Mapping]] = ...) -> None: ...
