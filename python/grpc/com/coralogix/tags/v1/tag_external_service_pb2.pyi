from google.protobuf import wrappers_pb2 as _wrappers_pb2
from com.coralogix.tags.v1 import tag_pb2 as _tag_pb2
from com.coralogix.tags.v1 import tag_type_pb2 as _tag_type_pb2
from com.coralogix.tags.v1 import audit_log_pb2 as _audit_log_pb2
from google.protobuf import descriptor_pb2 as _descriptor_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateExternalTagRequest(_message.Message):
    __slots__ = ("name", "key", "application", "subsystem", "icon_url", "timestamp", "type")
    NAME_FIELD_NUMBER: _ClassVar[int]
    KEY_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
    ICON_URL_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    key: _wrappers_pb2.StringValue
    application: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    subsystem: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    icon_url: _wrappers_pb2.StringValue
    timestamp: _timestamp_pb2.Timestamp
    type: _tag_type_pb2.Type
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., application: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystem: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., icon_url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., type: _Optional[_Union[_tag_type_pb2.Type, str]] = ...) -> None: ...

class AddExternalTagRequest(_message.Message):
    __slots__ = ("name", "key", "application", "subsystem", "icon_url", "timestamp", "type")
    NAME_FIELD_NUMBER: _ClassVar[int]
    KEY_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
    ICON_URL_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    key: _wrappers_pb2.StringValue
    application: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    subsystem: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    icon_url: _wrappers_pb2.StringValue
    timestamp: _wrappers_pb2.StringValue
    type: _tag_type_pb2.Type
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., application: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystem: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., icon_url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., timestamp: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., type: _Optional[_Union[_tag_type_pb2.Type, str]] = ...) -> None: ...

class CreateBitbucketTagRequest(_message.Message):
    __slots__ = ("name", "application", "subsystem", "body")
    class BitbucketBody(_message.Message):
        __slots__ = ("commit_status", "repository")
        class CommitStatus(_message.Message):
            __slots__ = ("state", "updated_on", "repository", "url")
            class Repository(_message.Message):
                __slots__ = ("full_name",)
                FULL_NAME_FIELD_NUMBER: _ClassVar[int]
                full_name: _wrappers_pb2.StringValue
                def __init__(self, full_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
            STATE_FIELD_NUMBER: _ClassVar[int]
            UPDATED_ON_FIELD_NUMBER: _ClassVar[int]
            REPOSITORY_FIELD_NUMBER: _ClassVar[int]
            URL_FIELD_NUMBER: _ClassVar[int]
            state: _wrappers_pb2.StringValue
            updated_on: _timestamp_pb2.Timestamp
            repository: CreateBitbucketTagRequest.BitbucketBody.CommitStatus.Repository
            url: _wrappers_pb2.StringValue
            def __init__(self, state: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., updated_on: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., repository: _Optional[_Union[CreateBitbucketTagRequest.BitbucketBody.CommitStatus.Repository, _Mapping]] = ..., url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
        class Repository(_message.Message):
            __slots__ = ("links",)
            class Links(_message.Message):
                __slots__ = ("avatar",)
                class Avatar(_message.Message):
                    __slots__ = ("href",)
                    HREF_FIELD_NUMBER: _ClassVar[int]
                    href: _wrappers_pb2.StringValue
                    def __init__(self, href: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
                AVATAR_FIELD_NUMBER: _ClassVar[int]
                avatar: CreateBitbucketTagRequest.BitbucketBody.Repository.Links.Avatar
                def __init__(self, avatar: _Optional[_Union[CreateBitbucketTagRequest.BitbucketBody.Repository.Links.Avatar, _Mapping]] = ...) -> None: ...
            LINKS_FIELD_NUMBER: _ClassVar[int]
            links: CreateBitbucketTagRequest.BitbucketBody.Repository.Links
            def __init__(self, links: _Optional[_Union[CreateBitbucketTagRequest.BitbucketBody.Repository.Links, _Mapping]] = ...) -> None: ...
        COMMIT_STATUS_FIELD_NUMBER: _ClassVar[int]
        REPOSITORY_FIELD_NUMBER: _ClassVar[int]
        commit_status: CreateBitbucketTagRequest.BitbucketBody.CommitStatus
        repository: CreateBitbucketTagRequest.BitbucketBody.Repository
        def __init__(self, commit_status: _Optional[_Union[CreateBitbucketTagRequest.BitbucketBody.CommitStatus, _Mapping]] = ..., repository: _Optional[_Union[CreateBitbucketTagRequest.BitbucketBody.Repository, _Mapping]] = ...) -> None: ...
    NAME_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
    BODY_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    application: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    subsystem: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    body: _struct_pb2.Struct
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., application: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystem: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., body: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...

class CreateTfsTagRequest(_message.Message):
    __slots__ = ("name", "application", "subsystem", "body")
    class TfsBody(_message.Message):
        __slots__ = ("resource", "created_date")
        class Resource(_message.Message):
            __slots__ = ("status", "environment", "finish_time", "source_branch", "definition", "deployment", "project")
            class Project(_message.Message):
                __slots__ = ("name",)
                NAME_FIELD_NUMBER: _ClassVar[int]
                name: _wrappers_pb2.StringValue
                def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
            class Environment(_message.Message):
                __slots__ = ("status", "source_branch", "project")
                STATUS_FIELD_NUMBER: _ClassVar[int]
                SOURCE_BRANCH_FIELD_NUMBER: _ClassVar[int]
                PROJECT_FIELD_NUMBER: _ClassVar[int]
                status: _wrappers_pb2.StringValue
                source_branch: _wrappers_pb2.StringValue
                project: CreateTfsTagRequest.TfsBody.Resource.Project
                def __init__(self, status: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., source_branch: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., project: _Optional[_Union[CreateTfsTagRequest.TfsBody.Resource.Project, _Mapping]] = ...) -> None: ...
            class Definition(_message.Message):
                __slots__ = ("url",)
                URL_FIELD_NUMBER: _ClassVar[int]
                url: _wrappers_pb2.StringValue
                def __init__(self, url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
            class Deployment(_message.Message):
                __slots__ = ("links",)
                class Links(_message.Message):
                    __slots__ = ("web",)
                    class Web(_message.Message):
                        __slots__ = ("href",)
                        HREF_FIELD_NUMBER: _ClassVar[int]
                        href: _wrappers_pb2.StringValue
                        def __init__(self, href: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
                    WEB_FIELD_NUMBER: _ClassVar[int]
                    web: CreateTfsTagRequest.TfsBody.Resource.Deployment.Links.Web
                    def __init__(self, web: _Optional[_Union[CreateTfsTagRequest.TfsBody.Resource.Deployment.Links.Web, _Mapping]] = ...) -> None: ...
                LINKS_FIELD_NUMBER: _ClassVar[int]
                links: CreateTfsTagRequest.TfsBody.Resource.Deployment.Links
                def __init__(self, links: _Optional[_Union[CreateTfsTagRequest.TfsBody.Resource.Deployment.Links, _Mapping]] = ...) -> None: ...
            STATUS_FIELD_NUMBER: _ClassVar[int]
            ENVIRONMENT_FIELD_NUMBER: _ClassVar[int]
            FINISH_TIME_FIELD_NUMBER: _ClassVar[int]
            SOURCE_BRANCH_FIELD_NUMBER: _ClassVar[int]
            DEFINITION_FIELD_NUMBER: _ClassVar[int]
            DEPLOYMENT_FIELD_NUMBER: _ClassVar[int]
            PROJECT_FIELD_NUMBER: _ClassVar[int]
            status: _wrappers_pb2.StringValue
            environment: CreateTfsTagRequest.TfsBody.Resource.Environment
            finish_time: _timestamp_pb2.Timestamp
            source_branch: _wrappers_pb2.StringValue
            definition: CreateTfsTagRequest.TfsBody.Resource.Definition
            deployment: CreateTfsTagRequest.TfsBody.Resource.Deployment
            project: CreateTfsTagRequest.TfsBody.Resource.Project
            def __init__(self, status: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., environment: _Optional[_Union[CreateTfsTagRequest.TfsBody.Resource.Environment, _Mapping]] = ..., finish_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., source_branch: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., definition: _Optional[_Union[CreateTfsTagRequest.TfsBody.Resource.Definition, _Mapping]] = ..., deployment: _Optional[_Union[CreateTfsTagRequest.TfsBody.Resource.Deployment, _Mapping]] = ..., project: _Optional[_Union[CreateTfsTagRequest.TfsBody.Resource.Project, _Mapping]] = ...) -> None: ...
        RESOURCE_FIELD_NUMBER: _ClassVar[int]
        CREATED_DATE_FIELD_NUMBER: _ClassVar[int]
        resource: CreateTfsTagRequest.TfsBody.Resource
        created_date: _timestamp_pb2.Timestamp
        def __init__(self, resource: _Optional[_Union[CreateTfsTagRequest.TfsBody.Resource, _Mapping]] = ..., created_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
    NAME_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
    BODY_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    application: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    subsystem: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    body: _struct_pb2.Struct
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., application: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystem: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., body: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...

class CreateGitlabTagRequest(_message.Message):
    __slots__ = ("name", "application", "subsystem", "body")
    class GitlabBody(_message.Message):
        __slots__ = ("object_attributes", "project")
        class ObjectAttributes(_message.Message):
            __slots__ = ("status", "finished_at", "ref")
            STATUS_FIELD_NUMBER: _ClassVar[int]
            FINISHED_AT_FIELD_NUMBER: _ClassVar[int]
            REF_FIELD_NUMBER: _ClassVar[int]
            status: _wrappers_pb2.StringValue
            finished_at: _timestamp_pb2.Timestamp
            ref: _wrappers_pb2.StringValue
            def __init__(self, status: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., finished_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., ref: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
        class Project(_message.Message):
            __slots__ = ("name", "git_http_url")
            NAME_FIELD_NUMBER: _ClassVar[int]
            GIT_HTTP_URL_FIELD_NUMBER: _ClassVar[int]
            name: _wrappers_pb2.StringValue
            git_http_url: _wrappers_pb2.StringValue
            def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., git_http_url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
        OBJECT_ATTRIBUTES_FIELD_NUMBER: _ClassVar[int]
        PROJECT_FIELD_NUMBER: _ClassVar[int]
        object_attributes: CreateGitlabTagRequest.GitlabBody.ObjectAttributes
        project: CreateGitlabTagRequest.GitlabBody.Project
        def __init__(self, object_attributes: _Optional[_Union[CreateGitlabTagRequest.GitlabBody.ObjectAttributes, _Mapping]] = ..., project: _Optional[_Union[CreateGitlabTagRequest.GitlabBody.Project, _Mapping]] = ...) -> None: ...
    NAME_FIELD_NUMBER: _ClassVar[int]
    APPLICATION_FIELD_NUMBER: _ClassVar[int]
    SUBSYSTEM_FIELD_NUMBER: _ClassVar[int]
    BODY_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    application: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    subsystem: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    body: _struct_pb2.Struct
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., application: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., subsystem: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., body: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...

class CreateExternalTagResponse(_message.Message):
    __slots__ = ("tag",)
    TAG_FIELD_NUMBER: _ClassVar[int]
    tag: _tag_pb2.Tag
    def __init__(self, tag: _Optional[_Union[_tag_pb2.Tag, _Mapping]] = ...) -> None: ...

class AddExternalTagResponse(_message.Message):
    __slots__ = ("tag",)
    TAG_FIELD_NUMBER: _ClassVar[int]
    tag: _tag_pb2.Tag
    def __init__(self, tag: _Optional[_Union[_tag_pb2.Tag, _Mapping]] = ...) -> None: ...

class CreateGitlabTagResponse(_message.Message):
    __slots__ = ("response",)
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    response: _wrappers_pb2.StringValue
    def __init__(self, response: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class CreateTfsTagResponse(_message.Message):
    __slots__ = ("response",)
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    response: _wrappers_pb2.StringValue
    def __init__(self, response: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class CreateBitbucketTagResponse(_message.Message):
    __slots__ = ("response",)
    RESPONSE_FIELD_NUMBER: _ClassVar[int]
    response: _wrappers_pb2.StringValue
    def __init__(self, response: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
