// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cxsdk

import (
	"context"

	rulesgroups "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/rules/v1"
)

// CreateRuleGroupRequest represents a request to create a rule group.
type CreateRuleGroupRequest = rulesgroups.CreateRuleGroupRequest

// CreateRuleGroupResponse represents the response after creating a rule group.
type CreateRuleGroupResponse = rulesgroups.CreateRuleGroupResponse

// GetRuleGroupRequest represents a request to get a rule group.
type GetRuleGroupRequest = rulesgroups.GetRuleGroupRequest

// GetRuleGroupResponse represents the response after getting a rule group.
type GetRuleGroupResponse = rulesgroups.GetRuleGroupResponse

// UpdateRuleGroupRequest represents a request to update a rule group.
type UpdateRuleGroupRequest = rulesgroups.UpdateRuleGroupRequest

// UpdateRuleGroupResponse represents the response after updating a rule group.
type UpdateRuleGroupResponse = rulesgroups.UpdateRuleGroupResponse

// DeleteRuleGroupRequest represents a request to delete a rule group.
type DeleteRuleGroupRequest = rulesgroups.DeleteRuleGroupRequest

// DeleteRuleGroupResponse represents the response after deleting a rule group.
type DeleteRuleGroupResponse = rulesgroups.DeleteRuleGroupResponse

// CreateRuleGroupRequestCreateRuleSubgroup represents a subgroup in the create rule group request.
type CreateRuleGroupRequestCreateRuleSubgroup = rulesgroups.CreateRuleGroupRequest_CreateRuleSubgroup

// CreateRuleGroupRequestCreateRuleSubgroupCreateRule represents a rule in the create rule group subgroup request.
type CreateRuleGroupRequestCreateRuleSubgroupCreateRule = rulesgroups.CreateRuleGroupRequest_CreateRuleSubgroup_CreateRule

// RuleParameters represents rule parameters.
type RuleParameters = rulesgroups.RuleParameters

// RuleParametersParseParameters represents parse parameters for a rule.
type RuleParametersParseParameters = rulesgroups.RuleParameters_ParseParameters

// ParseParameters represents parse parameters.
type ParseParameters = rulesgroups.ParseParameters

// RuleParametersJSONParseParameters represents JSON parse parameters for a rule.
type RuleParametersJSONParseParameters = rulesgroups.RuleParameters_JsonParseParameters

// JSONParseParameters represents JSON parse parameters.
type JSONParseParameters = rulesgroups.JsonParseParameters

// RuleParametersJSONStringifyParameters represents JSON stringify parameters for a rule.
type RuleParametersJSONStringifyParameters = rulesgroups.RuleParameters_JsonStringifyParameters

// JSONStringifyParameters represents JSON stringify parameters.
type JSONStringifyParameters = rulesgroups.JsonStringifyParameters

// RuleParametersJSONExtractParameters represents JSON extract parameters for a rule.
type RuleParametersJSONExtractParameters = rulesgroups.RuleParameters_JsonExtractParameters

// JSONExtractParameters represents JSON extract parameters.
type JSONExtractParameters = rulesgroups.JsonExtractParameters

// RuleParametersRemoveFieldsParameters represents parameters to remove fields in a rule.
type RuleParametersRemoveFieldsParameters = rulesgroups.RuleParameters_RemoveFieldsParameters

// RemoveFieldsParameters represents parameters to remove fields.
type RemoveFieldsParameters = rulesgroups.RemoveFieldsParameters

// RuleParametersExtractTimestampParameters represents parameters to extract timestamp in a rule.
type RuleParametersExtractTimestampParameters = rulesgroups.RuleParameters_ExtractTimestampParameters

// ExtractTimestampParameters represents parameters to extract timestamp.
type ExtractTimestampParameters = rulesgroups.ExtractTimestampParameters

// RuleParametersBlockParameters represents parameters to block in a rule.
type RuleParametersBlockParameters = rulesgroups.RuleParameters_BlockParameters

// BlockParameters represents block parameters.
type BlockParameters = rulesgroups.BlockParameters

// RuleParametersAllowParameters represents allow parameters in a rule.
type RuleParametersAllowParameters = rulesgroups.RuleParameters_AllowParameters

// AllowParameters represents allow parameters.
type AllowParameters = rulesgroups.AllowParameters

// RuleParametersReplaceParameters represents replace parameters in a rule.
type RuleParametersReplaceParameters = rulesgroups.RuleParameters_ReplaceParameters

// ReplaceParameters represents replace parameters.
type ReplaceParameters = rulesgroups.ReplaceParameters

// RuleParametersExtractParameters represents extract parameters in a rule.
type RuleParametersExtractParameters = rulesgroups.RuleParameters_ExtractParameters

// ExtractParameters represents extract parameters.
type ExtractParameters = rulesgroups.ExtractParameters

// RuleMatcher represents a rule matcher.
type RuleMatcher = rulesgroups.RuleMatcher

// ExtractTimestampParametersFormatStandard represents the standard timestamp format.
type ExtractTimestampParametersFormatStandard = rulesgroups.ExtractTimestampParameters_FormatStandard

// ApplicationNameConstraint represents an application name constraint.
type ApplicationNameConstraint = rulesgroups.ApplicationNameConstraint

// RuleMatcherApplicationName represents an application name matcher.
type RuleMatcherApplicationName = rulesgroups.RuleMatcher_ApplicationName

// SubsystemNameConstraint represents a subsystem name constraint.
type SubsystemNameConstraint = rulesgroups.SubsystemNameConstraint

// RuleMatcherSubsystemName represents a subsystem name matcher.
type RuleMatcherSubsystemName = rulesgroups.RuleMatcher_SubsystemName

// SeverityConstraint represents a severity constraint.
type SeverityConstraint = rulesgroups.SeverityConstraint

// SeverityConstraintValue represents a severity constraint value.
type SeverityConstraintValue = rulesgroups.SeverityConstraint_Value

// JSONExtractParametersDestinationField represents a JSON extract destination field.
type JSONExtractParametersDestinationField = rulesgroups.JsonExtractParameters_DestinationField

// RuleMatcherSeverity represents a severity matcher.
type RuleMatcherSeverity = rulesgroups.RuleMatcher_Severity

// RuleGroup represents a rule group.
type RuleGroup = rulesgroups.RuleGroup

// RuleSubgroup represents a subgroup of a rule.
type RuleSubgroup = rulesgroups.RuleSubgroup

// Rule represents a rule.
type Rule = rulesgroups.Rule

// SeverityConstraint values.
const (
	SeverityConstraintValueDebugOrUnspecified = rulesgroups.SeverityConstraint_VALUE_DEBUG_OR_UNSPECIFIED
	SeverityConstraintValueVerbose            = rulesgroups.SeverityConstraint_VALUE_VERBOSE
	SeverityConstraintValueInfo               = rulesgroups.SeverityConstraint_VALUE_INFO
	SeverityConstraintValueWarning            = rulesgroups.SeverityConstraint_VALUE_WARNING
	SeverityConstraintValueError              = rulesgroups.SeverityConstraint_VALUE_ERROR
	SeverityConstraintValueCritical           = rulesgroups.SeverityConstraint_VALUE_CRITICAL
)

// JSONExtractParameters values.
const (
	JSONExtractParametersDestinationFieldCategoryOrUnspecified = rulesgroups.JsonExtractParameters_DESTINATION_FIELD_CATEGORY_OR_UNSPECIFIED
	JSONExtractParametersDestinationFieldClassName             = rulesgroups.JsonExtractParameters_DESTINATION_FIELD_CLASSNAME
	JSONExtractParametersDestinationFieldMethodName            = rulesgroups.JsonExtractParameters_DESTINATION_FIELD_METHODNAME
	JSONExtractParametersDestinationFieldThreadID              = rulesgroups.JsonExtractParameters_DESTINATION_FIELD_THREADID
	JSONExtractParametersDestinationFieldSeverity              = rulesgroups.JsonExtractParameters_DESTINATION_FIELD_SEVERITY
)

// ExtractTimestampParameters values.
const (
	ExtractTimestampParametersFormatStandardStrftimeOrUnspecified = rulesgroups.ExtractTimestampParameters_FORMAT_STANDARD_STRFTIME_OR_UNSPECIFIED
	ExtractTimestampParametersFormatStandardJavasdf               = rulesgroups.ExtractTimestampParameters_FORMAT_STANDARD_JAVASDF
	ExtractTimestampParametersFormatStandardGolang                = rulesgroups.ExtractTimestampParameters_FORMAT_STANDARD_GOLANG
	ExtractTimestampParametersFormatStandardSecondsTS             = rulesgroups.ExtractTimestampParameters_FORMAT_STANDARD_SECONDSTS
	ExtractTimestampParametersFormatStandardMilliTS               = rulesgroups.ExtractTimestampParameters_FORMAT_STANDARD_MILLITS
	ExtractTimestampParametersFormatStandardMicroTS               = rulesgroups.ExtractTimestampParameters_FORMAT_STANDARD_MICROTS
	ExtractTimestampParametersFormatStandardNanoTS                = rulesgroups.ExtractTimestampParameters_FORMAT_STANDARD_NANOTS
)

// RuleGroupsClient is a client for the Coralogix Rules Groups API.
type RuleGroupsClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new rule group.
func (r RuleGroupsClient) Create(ctx context.Context, req *rulesgroups.CreateRuleGroupRequest) (*rulesgroups.CreateRuleGroupResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := rulesgroups.NewRuleGroupsServiceClient(conn)

	return client.CreateRuleGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Get gets a rule group.
func (r RuleGroupsClient) Get(ctx context.Context, req *rulesgroups.GetRuleGroupRequest) (*rulesgroups.GetRuleGroupResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := rulesgroups.NewRuleGroupsServiceClient(conn)

	return client.GetRuleGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Update updates a rule group.
func (r RuleGroupsClient) Update(ctx context.Context, req *rulesgroups.UpdateRuleGroupRequest) (*rulesgroups.UpdateRuleGroupResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()

	client := rulesgroups.NewRuleGroupsServiceClient(conn)

	return client.UpdateRuleGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes a rule group.
func (r RuleGroupsClient) Delete(ctx context.Context, req *rulesgroups.DeleteRuleGroupRequest) (*rulesgroups.DeleteRuleGroupResponse, error) {
	callProperties, err := r.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()

	client := rulesgroups.NewRuleGroupsServiceClient(conn)

	return client.DeleteRuleGroup(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewRuleGroupsClient creates a new rule groups client.
func NewRuleGroupsClient(c *CallPropertiesCreator) *RuleGroupsClient {
	return &RuleGroupsClient{callPropertiesCreator: c}
}
