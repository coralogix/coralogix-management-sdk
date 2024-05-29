from com.coralogix.extensions.v1 import extension_pb2 as _extension_pb2
from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ConnectionStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    PENDING: _ClassVar[ConnectionStatus]
    ACTIVE: _ClassVar[ConnectionStatus]
    FAILING: _ClassVar[ConnectionStatus]
    STATUS_UNKNOWN: _ClassVar[ConnectionStatus]
PENDING: ConnectionStatus
ACTIVE: ConnectionStatus
FAILING: ConnectionStatus
STATUS_UNKNOWN: ConnectionStatus

class CloudFormationStack(_message.Message):
    __slots__ = ("arn", "region")
    ARN_FIELD_NUMBER: _ClassVar[int]
    REGION_FIELD_NUMBER: _ClassVar[int]
    arn: _wrappers_pb2.StringValue
    region: _wrappers_pb2.StringValue
    def __init__(self, arn: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., region: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class ARMStack(_message.Message):
    __slots__ = ("subscription_id", "resource_group_name")
    SUBSCRIPTION_ID_FIELD_NUMBER: _ClassVar[int]
    RESOURCE_GROUP_NAME_FIELD_NUMBER: _ClassVar[int]
    subscription_id: _wrappers_pb2.StringValue
    resource_group_name: _wrappers_pb2.StringValue
    def __init__(self, subscription_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., resource_group_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class NoDeployment(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class IntegrationStatus(_message.Message):
    __slots__ = ("connection_status", "message", "details")
    class DetailsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    CONNECTION_STATUS_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    DETAILS_FIELD_NUMBER: _ClassVar[int]
    connection_status: ConnectionStatus
    message: _wrappers_pb2.StringValue
    details: _containers.ScalarMap[str, str]
    def __init__(self, connection_status: _Optional[_Union[ConnectionStatus, str]] = ..., message: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., details: _Optional[_Mapping[str, str]] = ...) -> None: ...

class IntegrationDoc(_message.Message):
    __slots__ = ("name", "link")
    NAME_FIELD_NUMBER: _ClassVar[int]
    LINK_FIELD_NUMBER: _ClassVar[int]
    name: _wrappers_pb2.StringValue
    link: _wrappers_pb2.StringValue
    def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., link: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...

class Integration(_message.Message):
    __slots__ = ("id", "name", "description", "icon", "dark_icon", "tags", "versions")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    ICON_FIELD_NUMBER: _ClassVar[int]
    DARK_ICON_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    VERSIONS_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    name: _wrappers_pb2.StringValue
    description: _wrappers_pb2.StringValue
    icon: _wrappers_pb2.StringValue
    dark_icon: _wrappers_pb2.StringValue
    tags: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    versions: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., icon: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., dark_icon: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., tags: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., versions: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...

class RevisionRef(_message.Message):
    __slots__ = ("version", "description_md")
    VERSION_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_MD_FIELD_NUMBER: _ClassVar[int]
    version: str
    description_md: str
    def __init__(self, version: _Optional[str] = ..., description_md: _Optional[str] = ...) -> None: ...

class LocalChangelog(_message.Message):
    __slots__ = ("changes",)
    CHANGES_FIELD_NUMBER: _ClassVar[int]
    changes: _containers.RepeatedCompositeFieldContainer[RevisionRef]
    def __init__(self, changes: _Optional[_Iterable[_Union[RevisionRef, _Mapping]]] = ...) -> None: ...

class ExternalUrl(_message.Message):
    __slots__ = ("url",)
    URL_FIELD_NUMBER: _ClassVar[int]
    url: str
    def __init__(self, url: _Optional[str] = ...) -> None: ...

class IntegrationDetails(_message.Message):
    __slots__ = ("integration", "extensions", "docs", "default", "local", "external")
    class DefaultIntegrationDetails(_message.Message):
        __slots__ = ("registered",)
        class RegisteredInstance(_message.Message):
            __slots__ = ("id", "definition_version", "last_updated", "parameters", "integration_status", "empty", "cloudformation", "arm", "is_testing")
            ID_FIELD_NUMBER: _ClassVar[int]
            DEFINITION_VERSION_FIELD_NUMBER: _ClassVar[int]
            LAST_UPDATED_FIELD_NUMBER: _ClassVar[int]
            PARAMETERS_FIELD_NUMBER: _ClassVar[int]
            INTEGRATION_STATUS_FIELD_NUMBER: _ClassVar[int]
            EMPTY_FIELD_NUMBER: _ClassVar[int]
            CLOUDFORMATION_FIELD_NUMBER: _ClassVar[int]
            ARM_FIELD_NUMBER: _ClassVar[int]
            IS_TESTING_FIELD_NUMBER: _ClassVar[int]
            id: _wrappers_pb2.StringValue
            definition_version: _wrappers_pb2.StringValue
            last_updated: _timestamp_pb2.Timestamp
            parameters: _containers.RepeatedCompositeFieldContainer[Parameter]
            integration_status: IntegrationStatus
            empty: NoDeployment
            cloudformation: CloudFormationStack
            arm: ARMStack
            is_testing: _wrappers_pb2.BoolValue
            def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., definition_version: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., last_updated: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., parameters: _Optional[_Iterable[_Union[Parameter, _Mapping]]] = ..., integration_status: _Optional[_Union[IntegrationStatus, _Mapping]] = ..., empty: _Optional[_Union[NoDeployment, _Mapping]] = ..., cloudformation: _Optional[_Union[CloudFormationStack, _Mapping]] = ..., arm: _Optional[_Union[ARMStack, _Mapping]] = ..., is_testing: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
        REGISTERED_FIELD_NUMBER: _ClassVar[int]
        registered: _containers.RepeatedCompositeFieldContainer[IntegrationDetails.DefaultIntegrationDetails.RegisteredInstance]
        def __init__(self, registered: _Optional[_Iterable[_Union[IntegrationDetails.DefaultIntegrationDetails.RegisteredInstance, _Mapping]]] = ...) -> None: ...
    INTEGRATION_FIELD_NUMBER: _ClassVar[int]
    EXTENSIONS_FIELD_NUMBER: _ClassVar[int]
    DOCS_FIELD_NUMBER: _ClassVar[int]
    DEFAULT_FIELD_NUMBER: _ClassVar[int]
    LOCAL_FIELD_NUMBER: _ClassVar[int]
    EXTERNAL_FIELD_NUMBER: _ClassVar[int]
    integration: Integration
    extensions: _containers.RepeatedCompositeFieldContainer[_extension_pb2.Extension]
    docs: _containers.RepeatedCompositeFieldContainer[IntegrationDoc]
    default: IntegrationDetails.DefaultIntegrationDetails
    local: LocalChangelog
    external: ExternalUrl
    def __init__(self, integration: _Optional[_Union[Integration, _Mapping]] = ..., extensions: _Optional[_Iterable[_Union[_extension_pb2.Extension, _Mapping]]] = ..., docs: _Optional[_Iterable[_Union[IntegrationDoc, _Mapping]]] = ..., default: _Optional[_Union[IntegrationDetails.DefaultIntegrationDetails, _Mapping]] = ..., local: _Optional[_Union[LocalChangelog, _Mapping]] = ..., external: _Optional[_Union[ExternalUrl, _Mapping]] = ...) -> None: ...

class IntegrationDefinition(_message.Message):
    __slots__ = ("key", "revisions")
    KEY_FIELD_NUMBER: _ClassVar[int]
    REVISIONS_FIELD_NUMBER: _ClassVar[int]
    key: _wrappers_pb2.StringValue
    revisions: _containers.RepeatedCompositeFieldContainer[IntegrationRevision]
    def __init__(self, key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., revisions: _Optional[_Iterable[_Union[IntegrationRevision, _Mapping]]] = ...) -> None: ...

class IntegrationRevision(_message.Message):
    __slots__ = ("id", "fields", "groups", "cloud_formation", "managed_service", "helm_chart", "azure_arm_template", "rum", "terraform", "upgrade_instructions_md", "revision_deployment_supported")
    class InputType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        API_KEY: _ClassVar[IntegrationRevision.InputType]
        TEXT: _ClassVar[IntegrationRevision.InputType]
        LIST_TEXT: _ClassVar[IntegrationRevision.InputType]
        MULTIPLE_SELECTION: _ClassVar[IntegrationRevision.InputType]
        BOOLEAN: _ClassVar[IntegrationRevision.InputType]
        SELECT: _ClassVar[IntegrationRevision.InputType]
        JSON: _ClassVar[IntegrationRevision.InputType]
        NUMBER: _ClassVar[IntegrationRevision.InputType]
    API_KEY: IntegrationRevision.InputType
    TEXT: IntegrationRevision.InputType
    LIST_TEXT: IntegrationRevision.InputType
    MULTIPLE_SELECTION: IntegrationRevision.InputType
    BOOLEAN: IntegrationRevision.InputType
    SELECT: IntegrationRevision.InputType
    JSON: IntegrationRevision.InputType
    NUMBER: IntegrationRevision.InputType
    class CloudFormationTemplate(_message.Message):
        __slots__ = ("template_url", "parameters", "post_installation_steps")
        class ParametersEntry(_message.Message):
            __slots__ = ("key", "value")
            KEY_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            key: str
            value: str
            def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
        class PostInstallationStepsEntry(_message.Message):
            __slots__ = ("key", "value")
            KEY_FIELD_NUMBER: _ClassVar[int]
            VALUE_FIELD_NUMBER: _ClassVar[int]
            key: str
            value: str
            def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
        TEMPLATE_URL_FIELD_NUMBER: _ClassVar[int]
        PARAMETERS_FIELD_NUMBER: _ClassVar[int]
        POST_INSTALLATION_STEPS_FIELD_NUMBER: _ClassVar[int]
        template_url: _wrappers_pb2.StringValue
        parameters: _containers.ScalarMap[str, str]
        post_installation_steps: _containers.ScalarMap[str, str]
        def __init__(self, template_url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., parameters: _Optional[_Mapping[str, str]] = ..., post_installation_steps: _Optional[_Mapping[str, str]] = ...) -> None: ...
    class ManagedService(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class HelmChart(_message.Message):
        __slots__ = ("template", "commands", "guide")
        TEMPLATE_FIELD_NUMBER: _ClassVar[int]
        COMMANDS_FIELD_NUMBER: _ClassVar[int]
        GUIDE_FIELD_NUMBER: _ClassVar[int]
        template: _wrappers_pb2.StringValue
        commands: _containers.RepeatedCompositeFieldContainer[IntegrationRevision.CommandInformation]
        guide: IntegrationRevision.IntegrationGuide
        def __init__(self, template: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., commands: _Optional[_Iterable[_Union[IntegrationRevision.CommandInformation, _Mapping]]] = ..., guide: _Optional[_Union[IntegrationRevision.IntegrationGuide, _Mapping]] = ...) -> None: ...
    class Terraform(_message.Message):
        __slots__ = ("configuration_blocks",)
        CONFIGURATION_BLOCKS_FIELD_NUMBER: _ClassVar[int]
        configuration_blocks: _containers.RepeatedCompositeFieldContainer[IntegrationRevision.ConfigurationBlock]
        def __init__(self, configuration_blocks: _Optional[_Iterable[_Union[IntegrationRevision.ConfigurationBlock, _Mapping]]] = ...) -> None: ...
    class Rum(_message.Message):
        __slots__ = ("browser_sdk_commands", "source_map_commands")
        BROWSER_SDK_COMMANDS_FIELD_NUMBER: _ClassVar[int]
        SOURCE_MAP_COMMANDS_FIELD_NUMBER: _ClassVar[int]
        browser_sdk_commands: _containers.RepeatedCompositeFieldContainer[IntegrationRevision.CommandInformation]
        source_map_commands: _containers.RepeatedCompositeFieldContainer[IntegrationRevision.CommandInformation]
        def __init__(self, browser_sdk_commands: _Optional[_Iterable[_Union[IntegrationRevision.CommandInformation, _Mapping]]] = ..., source_map_commands: _Optional[_Iterable[_Union[IntegrationRevision.CommandInformation, _Mapping]]] = ...) -> None: ...
    class AzureArmTemplate(_message.Message):
        __slots__ = ("template_url",)
        TEMPLATE_URL_FIELD_NUMBER: _ClassVar[int]
        template_url: _wrappers_pb2.StringValue
        def __init__(self, template_url: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class ListTextValue(_message.Message):
        __slots__ = ("options", "default_values")
        OPTIONS_FIELD_NUMBER: _ClassVar[int]
        DEFAULT_VALUES_FIELD_NUMBER: _ClassVar[int]
        options: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        default_values: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        def __init__(self, options: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., default_values: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...
    class SingleValue(_message.Message):
        __slots__ = ("default_value",)
        DEFAULT_VALUE_FIELD_NUMBER: _ClassVar[int]
        default_value: _wrappers_pb2.StringValue
        def __init__(self, default_value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class SingleBooleanValue(_message.Message):
        __slots__ = ("default_value",)
        DEFAULT_VALUE_FIELD_NUMBER: _ClassVar[int]
        default_value: _wrappers_pb2.BoolValue
        def __init__(self, default_value: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
    class SingleNumericValue(_message.Message):
        __slots__ = ("default_value",)
        DEFAULT_VALUE_FIELD_NUMBER: _ClassVar[int]
        default_value: _wrappers_pb2.DoubleValue
        def __init__(self, default_value: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ...) -> None: ...
    class MultipleSelectionValue(_message.Message):
        __slots__ = ("options",)
        OPTIONS_FIELD_NUMBER: _ClassVar[int]
        options: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        def __init__(self, options: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...
    class SelectionValue(_message.Message):
        __slots__ = ("options", "default_value")
        OPTIONS_FIELD_NUMBER: _ClassVar[int]
        DEFAULT_VALUE_FIELD_NUMBER: _ClassVar[int]
        options: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        default_value: _wrappers_pb2.StringValue
        def __init__(self, options: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ..., default_value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class FieldCondition(_message.Message):
        __slots__ = ("type", "values")
        class ConditionType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            UNKNOWN: _ClassVar[IntegrationRevision.FieldCondition.ConditionType]
            OR: _ClassVar[IntegrationRevision.FieldCondition.ConditionType]
            AND: _ClassVar[IntegrationRevision.FieldCondition.ConditionType]
        UNKNOWN: IntegrationRevision.FieldCondition.ConditionType
        OR: IntegrationRevision.FieldCondition.ConditionType
        AND: IntegrationRevision.FieldCondition.ConditionType
        class FieldValue(_message.Message):
            __slots__ = ("field_name", "value_pattern")
            FIELD_NAME_FIELD_NUMBER: _ClassVar[int]
            VALUE_PATTERN_FIELD_NUMBER: _ClassVar[int]
            field_name: str
            value_pattern: str
            def __init__(self, field_name: _Optional[str] = ..., value_pattern: _Optional[str] = ...) -> None: ...
        TYPE_FIELD_NUMBER: _ClassVar[int]
        VALUES_FIELD_NUMBER: _ClassVar[int]
        type: IntegrationRevision.FieldCondition.ConditionType
        values: _containers.RepeatedCompositeFieldContainer[IntegrationRevision.FieldCondition.FieldValue]
        def __init__(self, type: _Optional[_Union[IntegrationRevision.FieldCondition.ConditionType, str]] = ..., values: _Optional[_Iterable[_Union[IntegrationRevision.FieldCondition.FieldValue, _Mapping]]] = ...) -> None: ...
    class Group(_message.Message):
        __slots__ = ("id", "name")
        ID_FIELD_NUMBER: _ClassVar[int]
        NAME_FIELD_NUMBER: _ClassVar[int]
        id: _wrappers_pb2.StringValue
        name: _wrappers_pb2.StringValue
        def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class FieldInformation(_message.Message):
        __slots__ = ("single", "multi_text", "multiple_selection", "single_boolean", "selection", "single_number", "type", "name", "tooltip", "template_param_name", "placeholder", "required", "predefined", "visible", "readonly", "applicable_if", "group_id", "upgrade_notice", "allowed_pattern")
        SINGLE_FIELD_NUMBER: _ClassVar[int]
        MULTI_TEXT_FIELD_NUMBER: _ClassVar[int]
        MULTIPLE_SELECTION_FIELD_NUMBER: _ClassVar[int]
        SINGLE_BOOLEAN_FIELD_NUMBER: _ClassVar[int]
        SELECTION_FIELD_NUMBER: _ClassVar[int]
        SINGLE_NUMBER_FIELD_NUMBER: _ClassVar[int]
        TYPE_FIELD_NUMBER: _ClassVar[int]
        NAME_FIELD_NUMBER: _ClassVar[int]
        TOOLTIP_FIELD_NUMBER: _ClassVar[int]
        TEMPLATE_PARAM_NAME_FIELD_NUMBER: _ClassVar[int]
        PLACEHOLDER_FIELD_NUMBER: _ClassVar[int]
        REQUIRED_FIELD_NUMBER: _ClassVar[int]
        PREDEFINED_FIELD_NUMBER: _ClassVar[int]
        VISIBLE_FIELD_NUMBER: _ClassVar[int]
        READONLY_FIELD_NUMBER: _ClassVar[int]
        APPLICABLE_IF_FIELD_NUMBER: _ClassVar[int]
        GROUP_ID_FIELD_NUMBER: _ClassVar[int]
        UPGRADE_NOTICE_FIELD_NUMBER: _ClassVar[int]
        ALLOWED_PATTERN_FIELD_NUMBER: _ClassVar[int]
        single: IntegrationRevision.SingleValue
        multi_text: IntegrationRevision.ListTextValue
        multiple_selection: IntegrationRevision.MultipleSelectionValue
        single_boolean: IntegrationRevision.SingleBooleanValue
        selection: IntegrationRevision.SelectionValue
        single_number: IntegrationRevision.SingleNumericValue
        type: IntegrationRevision.InputType
        name: _wrappers_pb2.StringValue
        tooltip: _wrappers_pb2.StringValue
        template_param_name: _wrappers_pb2.StringValue
        placeholder: _wrappers_pb2.StringValue
        required: _wrappers_pb2.BoolValue
        predefined: _wrappers_pb2.BoolValue
        visible: _wrappers_pb2.BoolValue
        readonly: _wrappers_pb2.BoolValue
        applicable_if: IntegrationRevision.FieldCondition
        group_id: _wrappers_pb2.StringValue
        upgrade_notice: _wrappers_pb2.StringValue
        allowed_pattern: _wrappers_pb2.StringValue
        def __init__(self, single: _Optional[_Union[IntegrationRevision.SingleValue, _Mapping]] = ..., multi_text: _Optional[_Union[IntegrationRevision.ListTextValue, _Mapping]] = ..., multiple_selection: _Optional[_Union[IntegrationRevision.MultipleSelectionValue, _Mapping]] = ..., single_boolean: _Optional[_Union[IntegrationRevision.SingleBooleanValue, _Mapping]] = ..., selection: _Optional[_Union[IntegrationRevision.SelectionValue, _Mapping]] = ..., single_number: _Optional[_Union[IntegrationRevision.SingleNumericValue, _Mapping]] = ..., type: _Optional[_Union[IntegrationRevision.InputType, str]] = ..., name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., tooltip: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., template_param_name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., placeholder: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., required: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., predefined: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., visible: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., readonly: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., applicable_if: _Optional[_Union[IntegrationRevision.FieldCondition, _Mapping]] = ..., group_id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., upgrade_notice: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., allowed_pattern: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class CommandInformation(_message.Message):
        __slots__ = ("name", "command", "description", "tooltip_text", "language", "links")
        class Language(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
            __slots__ = ()
            UNKNOWN: _ClassVar[IntegrationRevision.CommandInformation.Language]
            BASH: _ClassVar[IntegrationRevision.CommandInformation.Language]
            JAVASCRIPT: _ClassVar[IntegrationRevision.CommandInformation.Language]
        UNKNOWN: IntegrationRevision.CommandInformation.Language
        BASH: IntegrationRevision.CommandInformation.Language
        JAVASCRIPT: IntegrationRevision.CommandInformation.Language
        class Link(_message.Message):
            __slots__ = ("key", "text", "url")
            KEY_FIELD_NUMBER: _ClassVar[int]
            TEXT_FIELD_NUMBER: _ClassVar[int]
            URL_FIELD_NUMBER: _ClassVar[int]
            key: str
            text: str
            url: str
            def __init__(self, key: _Optional[str] = ..., text: _Optional[str] = ..., url: _Optional[str] = ...) -> None: ...
        NAME_FIELD_NUMBER: _ClassVar[int]
        COMMAND_FIELD_NUMBER: _ClassVar[int]
        DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
        TOOLTIP_TEXT_FIELD_NUMBER: _ClassVar[int]
        LANGUAGE_FIELD_NUMBER: _ClassVar[int]
        LINKS_FIELD_NUMBER: _ClassVar[int]
        name: _wrappers_pb2.StringValue
        command: _wrappers_pb2.StringValue
        description: _wrappers_pb2.StringValue
        tooltip_text: _wrappers_pb2.StringValue
        language: IntegrationRevision.CommandInformation.Language
        links: _containers.RepeatedCompositeFieldContainer[IntegrationRevision.CommandInformation.Link]
        def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., command: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., tooltip_text: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., language: _Optional[_Union[IntegrationRevision.CommandInformation.Language, str]] = ..., links: _Optional[_Iterable[_Union[IntegrationRevision.CommandInformation.Link, _Mapping]]] = ...) -> None: ...
    class ConfigurationBlock(_message.Message):
        __slots__ = ("name", "value", "description")
        NAME_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
        name: _wrappers_pb2.StringValue
        value: _wrappers_pb2.StringValue
        description: _wrappers_pb2.StringValue
        def __init__(self, name: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., description: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    class IntegrationGuide(_message.Message):
        __slots__ = ("introduction", "installation_requirements")
        INTRODUCTION_FIELD_NUMBER: _ClassVar[int]
        INSTALLATION_REQUIREMENTS_FIELD_NUMBER: _ClassVar[int]
        introduction: _wrappers_pb2.StringValue
        installation_requirements: _wrappers_pb2.StringValue
        def __init__(self, introduction: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., installation_requirements: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    FIELDS_FIELD_NUMBER: _ClassVar[int]
    GROUPS_FIELD_NUMBER: _ClassVar[int]
    CLOUD_FORMATION_FIELD_NUMBER: _ClassVar[int]
    MANAGED_SERVICE_FIELD_NUMBER: _ClassVar[int]
    HELM_CHART_FIELD_NUMBER: _ClassVar[int]
    AZURE_ARM_TEMPLATE_FIELD_NUMBER: _ClassVar[int]
    RUM_FIELD_NUMBER: _ClassVar[int]
    TERRAFORM_FIELD_NUMBER: _ClassVar[int]
    UPGRADE_INSTRUCTIONS_MD_FIELD_NUMBER: _ClassVar[int]
    REVISION_DEPLOYMENT_SUPPORTED_FIELD_NUMBER: _ClassVar[int]
    id: _wrappers_pb2.StringValue
    fields: _containers.RepeatedCompositeFieldContainer[IntegrationRevision.FieldInformation]
    groups: _containers.RepeatedCompositeFieldContainer[IntegrationRevision.Group]
    cloud_formation: IntegrationRevision.CloudFormationTemplate
    managed_service: IntegrationRevision.ManagedService
    helm_chart: IntegrationRevision.HelmChart
    azure_arm_template: IntegrationRevision.AzureArmTemplate
    rum: IntegrationRevision.Rum
    terraform: IntegrationRevision.Terraform
    upgrade_instructions_md: str
    revision_deployment_supported: bool
    def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., fields: _Optional[_Iterable[_Union[IntegrationRevision.FieldInformation, _Mapping]]] = ..., groups: _Optional[_Iterable[_Union[IntegrationRevision.Group, _Mapping]]] = ..., cloud_formation: _Optional[_Union[IntegrationRevision.CloudFormationTemplate, _Mapping]] = ..., managed_service: _Optional[_Union[IntegrationRevision.ManagedService, _Mapping]] = ..., helm_chart: _Optional[_Union[IntegrationRevision.HelmChart, _Mapping]] = ..., azure_arm_template: _Optional[_Union[IntegrationRevision.AzureArmTemplate, _Mapping]] = ..., rum: _Optional[_Union[IntegrationRevision.Rum, _Mapping]] = ..., terraform: _Optional[_Union[IntegrationRevision.Terraform, _Mapping]] = ..., upgrade_instructions_md: _Optional[str] = ..., revision_deployment_supported: bool = ...) -> None: ...

class Parameter(_message.Message):
    __slots__ = ("key", "string_value", "boolean_value", "string_list", "api_key", "numeric_value")
    class StringList(_message.Message):
        __slots__ = ("values",)
        VALUES_FIELD_NUMBER: _ClassVar[int]
        values: _containers.RepeatedCompositeFieldContainer[_wrappers_pb2.StringValue]
        def __init__(self, values: _Optional[_Iterable[_Union[_wrappers_pb2.StringValue, _Mapping]]] = ...) -> None: ...
    class ApiKeyData(_message.Message):
        __slots__ = ("id", "value")
        ID_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        id: _wrappers_pb2.StringValue
        value: _wrappers_pb2.StringValue
        def __init__(self, id: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    KEY_FIELD_NUMBER: _ClassVar[int]
    STRING_VALUE_FIELD_NUMBER: _ClassVar[int]
    BOOLEAN_VALUE_FIELD_NUMBER: _ClassVar[int]
    STRING_LIST_FIELD_NUMBER: _ClassVar[int]
    API_KEY_FIELD_NUMBER: _ClassVar[int]
    NUMERIC_VALUE_FIELD_NUMBER: _ClassVar[int]
    key: str
    string_value: _wrappers_pb2.StringValue
    boolean_value: _wrappers_pb2.BoolValue
    string_list: Parameter.StringList
    api_key: Parameter.ApiKeyData
    numeric_value: _wrappers_pb2.DoubleValue
    def __init__(self, key: _Optional[str] = ..., string_value: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., boolean_value: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ..., string_list: _Optional[_Union[Parameter.StringList, _Mapping]] = ..., api_key: _Optional[_Union[Parameter.ApiKeyData, _Mapping]] = ..., numeric_value: _Optional[_Union[_wrappers_pb2.DoubleValue, _Mapping]] = ...) -> None: ...

class GenericIntegrationParameters(_message.Message):
    __slots__ = ("parameters",)
    PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    parameters: _containers.RepeatedCompositeFieldContainer[Parameter]
    def __init__(self, parameters: _Optional[_Iterable[_Union[Parameter, _Mapping]]] = ...) -> None: ...

class IntegrationMetadata(_message.Message):
    __slots__ = ("integration_key", "version", "integration_parameters")
    INTEGRATION_KEY_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    INTEGRATION_PARAMETERS_FIELD_NUMBER: _ClassVar[int]
    integration_key: _wrappers_pb2.StringValue
    version: _wrappers_pb2.StringValue
    integration_parameters: GenericIntegrationParameters
    def __init__(self, integration_key: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., version: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., integration_parameters: _Optional[_Union[GenericIntegrationParameters, _Mapping]] = ...) -> None: ...

class TestIntegrationResult(_message.Message):
    __slots__ = ("success", "failure")
    class Success(_message.Message):
        __slots__ = ()
        def __init__(self) -> None: ...
    class Failure(_message.Message):
        __slots__ = ("error_message",)
        ERROR_MESSAGE_FIELD_NUMBER: _ClassVar[int]
        error_message: _wrappers_pb2.StringValue
        def __init__(self, error_message: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ...) -> None: ...
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    FAILURE_FIELD_NUMBER: _ClassVar[int]
    success: TestIntegrationResult.Success
    failure: TestIntegrationResult.Failure
    def __init__(self, success: _Optional[_Union[TestIntegrationResult.Success, _Mapping]] = ..., failure: _Optional[_Union[TestIntegrationResult.Failure, _Mapping]] = ...) -> None: ...

class RumVersionData(_message.Message):
    __slots__ = ("versions", "synced_at")
    class SourceMapMetadata(_message.Message):
        __slots__ = ("created_at", "is_uploaded_successful")
        CREATED_AT_FIELD_NUMBER: _ClassVar[int]
        IS_UPLOADED_SUCCESSFUL_FIELD_NUMBER: _ClassVar[int]
        created_at: _timestamp_pb2.Timestamp
        is_uploaded_successful: _wrappers_pb2.BoolValue
        def __init__(self, created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., is_uploaded_successful: _Optional[_Union[_wrappers_pb2.BoolValue, _Mapping]] = ...) -> None: ...
    class LogMetadata(_message.Message):
        __slots__ = ("first_occurrence", "last_occurrence")
        FIRST_OCCURRENCE_FIELD_NUMBER: _ClassVar[int]
        LAST_OCCURRENCE_FIELD_NUMBER: _ClassVar[int]
        first_occurrence: _timestamp_pb2.Timestamp
        last_occurrence: _timestamp_pb2.Timestamp
        def __init__(self, first_occurrence: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., last_occurrence: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
    class Version(_message.Message):
        __slots__ = ("version", "log_metadata", "source_map_metadata")
        VERSION_FIELD_NUMBER: _ClassVar[int]
        LOG_METADATA_FIELD_NUMBER: _ClassVar[int]
        SOURCE_MAP_METADATA_FIELD_NUMBER: _ClassVar[int]
        version: _wrappers_pb2.StringValue
        log_metadata: RumVersionData.LogMetadata
        source_map_metadata: RumVersionData.SourceMapMetadata
        def __init__(self, version: _Optional[_Union[_wrappers_pb2.StringValue, _Mapping]] = ..., log_metadata: _Optional[_Union[RumVersionData.LogMetadata, _Mapping]] = ..., source_map_metadata: _Optional[_Union[RumVersionData.SourceMapMetadata, _Mapping]] = ...) -> None: ...
    VERSIONS_FIELD_NUMBER: _ClassVar[int]
    SYNCED_AT_FIELD_NUMBER: _ClassVar[int]
    versions: _containers.RepeatedCompositeFieldContainer[RumVersionData.Version]
    synced_at: _timestamp_pb2.Timestamp
    def __init__(self, versions: _Optional[_Iterable[_Union[RumVersionData.Version, _Mapping]]] = ..., synced_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
