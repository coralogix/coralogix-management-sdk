from com.coralogix.dashboards.v1 import audit_log_pb2 as _audit_log_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SearchAnnouncementsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class SearchAnnouncementsResponse(_message.Message):
    __slots__ = ("announcements",)
    ANNOUNCEMENTS_FIELD_NUMBER: _ClassVar[int]
    announcements: _containers.RepeatedCompositeFieldContainer[Announcement]
    def __init__(self, announcements: _Optional[_Iterable[_Union[Announcement, _Mapping]]] = ...) -> None: ...

class SearchDismissedAnnouncementsIdsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class SearchDismissedAnnouncementsIdsResponse(_message.Message):
    __slots__ = ("dismissed_announcements_ids",)
    DISMISSED_ANNOUNCEMENTS_IDS_FIELD_NUMBER: _ClassVar[int]
    dismissed_announcements_ids: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, dismissed_announcements_ids: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class UpdateDismissedAnnouncementsIdsRequest(_message.Message):
    __slots__ = ("dismissed_announcements_ids",)
    DISMISSED_ANNOUNCEMENTS_IDS_FIELD_NUMBER: _ClassVar[int]
    dismissed_announcements_ids: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, dismissed_announcements_ids: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class UpdateDismissedAnnouncementsIdsResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class Announcement(_message.Message):
    __slots__ = ("id", "message", "link", "is_beta")
    class Message(_message.Message):
        __slots__ = ("headline", "description")
        HEADLINE_FIELD_NUMBER: _ClassVar[int]
        DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
        headline: _wrappers_pb2.StringValue
        description: _wrappers_pb2.StringValue
        def __init__(self, headline: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class Link(_message.Message):
        __slots__ = ("text", "url")
        TEXT_FIELD_NUMBER: _ClassVar[int]
        URL_FIELD_NUMBER: _ClassVar[int]
        text: _wrappers_pb2.StringValue
        url: _wrappers_pb2.StringValue
        def __init__(self, text: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    LINK_FIELD_NUMBER: _ClassVar[int]
    IS_BETA_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    message: Announcement.Message
    link: Announcement.Link
    is_beta: _wrappers_pb2.BoolValue
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., message: _Optional[_Union[Announcement.Message, _Mapping]] = ..., link: _Optional[_Union[Announcement.Link, _Mapping]] = ..., is_beta: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
