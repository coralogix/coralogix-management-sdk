from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class PushBasedPlatform(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    UNDEFINED: _ClassVar[PushBasedPlatform]
    PLATFORM_BITBUCKET: _ClassVar[PushBasedPlatform]
    PLATFORM_GITHUB: _ClassVar[PushBasedPlatform]
    PLATFORM_GITLAB: _ClassVar[PushBasedPlatform]
    PLATFORM_AWS_SNS: _ClassVar[PushBasedPlatform]
    PLATFORM_OPSGENIE: _ClassVar[PushBasedPlatform]
    PLATFORM_PAGERDUTY: _ClassVar[PushBasedPlatform]
    PLATFORM_PROMETHEUS: _ClassVar[PushBasedPlatform]
    PLATFORM_SLACK: _ClassVar[PushBasedPlatform]
    PLATFORM_INTERCOM: _ClassVar[PushBasedPlatform]
UNDEFINED: PushBasedPlatform
PLATFORM_BITBUCKET: PushBasedPlatform
PLATFORM_GITHUB: PushBasedPlatform
PLATFORM_GITLAB: PushBasedPlatform
PLATFORM_AWS_SNS: PushBasedPlatform
PLATFORM_OPSGENIE: PushBasedPlatform
PLATFORM_PAGERDUTY: PushBasedPlatform
PLATFORM_PROMETHEUS: PushBasedPlatform
PLATFORM_SLACK: PushBasedPlatform
PLATFORM_INTERCOM: PushBasedPlatform
