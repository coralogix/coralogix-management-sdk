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

	tcopolicies "github.com/coralogix/coralogix-management-sdk/go/internal/coralogix/quota/v1"
)

// TCOPolicy is a policy type.
type TCOPolicy = tcopolicies.Policy

// TCOPolicyLogRules is a rule type.
type TCOPolicyLogRules = tcopolicies.Policy_LogRules

// TCOPolicySpanRules is a spans type.
type TCOPolicySpanRules = tcopolicies.Policy_SpanRules

// TCOPolicyTagRule is a rule type.
type TCOSpanRules = tcopolicies.SpanRules

// TCOPolicyTagRule is a rule type.
type TCOPolicyTagRule = tcopolicies.TagRule

// TCOPolicyRule is a rule type.
type TCOPolicyRule = tcopolicies.Rule

// TCOLogRules is a rules type.
type TCOLogRules = tcopolicies.LogRules

// GetPolicyRequest is a request type.
type GetPolicyRequest = tcopolicies.GetPolicyRequest

// UpdatePolicyRequest is a request type.
type UpdatePolicyRequest = tcopolicies.UpdatePolicyRequest

// DeletePolicyRequest is a request type.
type DeletePolicyRequest = tcopolicies.DeletePolicyRequest

// ReorderPoliciesRequest is a request type.
type ReorderPoliciesRequest = tcopolicies.ReorderPoliciesRequest

// GetCompanyPoliciesRequest is a request type.
type GetCompanyPoliciesRequest = tcopolicies.GetCompanyPoliciesRequest

// GetCompanyPoliciesResponse is a request type.
type GetCompanyPoliciesResponse = tcopolicies.GetCompanyPoliciesResponse

// AtomicOverwriteLogPoliciesRequest is a request type.
type AtomicOverwriteLogPoliciesRequest = tcopolicies.AtomicOverwriteLogPoliciesRequest

// AtomicOverwriteLogPoliciesResponse is a request type.
type AtomicOverwriteLogPoliciesResponse = tcopolicies.AtomicOverwriteLogPoliciesResponse

// AtomicOverwriteSpanPoliciesRequest is a request type.
type AtomicOverwriteSpanPoliciesRequest = tcopolicies.AtomicOverwriteSpanPoliciesRequest

// CreatePolicyRequest is a request type.
type CreatePolicyRequest = tcopolicies.CreatePolicyRequest

// CreateLogPolicyRequest is a request type.
type CreateLogPolicyRequest = tcopolicies.CreateLogPolicyRequest

// CreateGenericPolicyRequest is a request type.
type CreateGenericPolicyRequest = tcopolicies.CreateGenericPolicyRequest

// CreateSpanPolicyRequest is  a request type.
type CreateSpanPolicyRequest = tcopolicies.CreateSpanPolicyRequest

// AtomicOverwriteSpanPoliciesResponse is  a response type.
type AtomicOverwriteSpanPoliciesResponse = tcopolicies.AtomicOverwriteSpanPoliciesResponse

// ArchiveRetention is a configuration type.
type ArchiveRetention = tcopolicies.ArchiveRetention

// TCOPolicySource is a type for TCO policy source setting.
type TCOPolicySource = tcopolicies.SourceType

// TCOPolicySourceType values.
const (
	TCOPolicySourceTypeUnspecified = tcopolicies.SourceType_SOURCE_TYPE_UNSPECIFIED
	TCOPolicySourceTypeLogs        = tcopolicies.SourceType_SOURCE_TYPE_LOGS
	TCOPolicySourceTypeSpans       = tcopolicies.SourceType_SOURCE_TYPE_SPANS
)

// TCOPolicySeverity is a type for TCO policy severity setting.
type TCOPolicySeverity = tcopolicies.Severity

// TCOPolicySeverity values.
const (
	TCOPolicySeverityDebug    = tcopolicies.Severity_SEVERITY_DEBUG
	TCOPolicySeverityVerbose  = tcopolicies.Severity_SEVERITY_VERBOSE
	TCOPolicySeverityInfo     = tcopolicies.Severity_SEVERITY_INFO
	TCOPolicySeverityWarning  = tcopolicies.Severity_SEVERITY_WARNING
	TCOPolicySeverityError    = tcopolicies.Severity_SEVERITY_ERROR
	TCOPolicySeverityCritical = tcopolicies.Severity_SEVERITY_CRITICAL
)

// TCOPolicyRuleTypeId is a type for TCO policy rule type ID.
type TCOPolicyRuleTypeId = tcopolicies.RuleTypeId

// TCOPolicyRuleTypeId values.
const (
	TCOPolicyRuleTypeIdUnspecified = tcopolicies.RuleTypeId_RULE_TYPE_ID_UNSPECIFIED
	TCOPolicyRuleTypeIdIs          = tcopolicies.RuleTypeId_RULE_TYPE_ID_IS
	TCOPolicyRuleTypeIdIsNot       = tcopolicies.RuleTypeId_RULE_TYPE_ID_IS_NOT
	TCOPolicyRuleTypeIdStartWith   = tcopolicies.RuleTypeId_RULE_TYPE_ID_START_WITH
	TCOPolicyRuleTypeIdIncludes    = tcopolicies.RuleTypeId_RULE_TYPE_ID_INCLUDES
)

// TCOPolicyPriority is a type for TCO policy priority.
type TCOPolicyPriority = tcopolicies.Priority

// TCOPolicyPriority values.
const (
	TCOPolicyPriorityBlock  = tcopolicies.Priority_PRIORITY_TYPE_BLOCK
	TCOPolicyPriorityHigh   = tcopolicies.Priority_PRIORITY_TYPE_HIGH
	TCOPolicyPriorityLow    = tcopolicies.Priority_PRIORITY_TYPE_LOW
	TCOPolicyPriorityMedium = tcopolicies.Priority_PRIORITY_TYPE_MEDIUM
)

// RPC names.
const (
	TCOPoliciesGetPolicyRPC                   = tcopolicies.PoliciesService_GetPolicy_FullMethodName
	TCOPoliciesCreatePolicyRPC                = tcopolicies.PoliciesService_CreatePolicy_FullMethodName
	TCOPoliciesUpdatePolicyRPC                = tcopolicies.PoliciesService_UpdatePolicy_FullMethodName
	TCOPoliciesGetCompanyPoliciesRPC          = tcopolicies.PoliciesService_GetCompanyPolicies_FullMethodName
	TCOPoliciesDeletePolicyRPC                = tcopolicies.PoliciesService_DeletePolicy_FullMethodName
	TCOPoliciesReorderPoliciesRPC             = tcopolicies.PoliciesService_ReorderPolicies_FullMethodName
	TCOPoliciesBulkTestLogPoliciesRPC         = tcopolicies.PoliciesService_BulkTestLogPolicies_FullMethodName
	TCOPoliciesTogglePolicyRPC                = tcopolicies.PoliciesService_TogglePolicy_FullMethodName
	TCOPoliciesAtomicBatchCreatePolicyRPC     = tcopolicies.PoliciesService_AtomicBatchCreatePolicy_FullMethodName
	TCOPoliciesAtomicOverwriteLogPoliciesRPC  = tcopolicies.PoliciesService_AtomicOverwriteLogPolicies_FullMethodName
	TCOPoliciesAtomicOverwriteSpanPoliciesRPC = tcopolicies.PoliciesService_AtomicOverwriteSpanPolicies_FullMethodName
)

// TCOPoliciesClient is a client for the Coralogix TCO Policies API.
type TCOPoliciesClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// CreateTCOPolicy creates a new TCO policy.
func (t TCOPoliciesClient) Create(ctx context.Context, req *CreatePolicyRequest) (*tcopolicies.CreatePolicyResponse, error) {
	callProperties, err := t.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := tcopolicies.NewPoliciesServiceClient(conn)

	return client.CreatePolicy(callProperties.Ctx, req, callProperties.CallOptions...)
}

// GetTCOPolicy gets the specified TCO policy.
func (t TCOPoliciesClient) Get(ctx context.Context, req *GetPolicyRequest) (*tcopolicies.GetPolicyResponse, error) {
	callProperties, err := t.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := tcopolicies.NewPoliciesServiceClient(conn)

	return client.GetPolicy(callProperties.Ctx, req, callProperties.CallOptions...)
}

// OverwriteTCOTracesPolicies overwrites the specified TCO traces policies.
func (t TCOPoliciesClient) OverwriteTCOTracesPolicies(ctx context.Context, req *AtomicOverwriteSpanPoliciesRequest) (*tcopolicies.AtomicOverwriteSpanPoliciesResponse, error) {
	callProperties, err := t.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := tcopolicies.NewPoliciesServiceClient(conn)

	return client.AtomicOverwriteSpanPolicies(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Update updates the specified TCO policy.
func (t TCOPoliciesClient) Update(ctx context.Context, req *UpdatePolicyRequest) (*tcopolicies.UpdatePolicyResponse, error) {
	callProperties, err := t.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := tcopolicies.NewPoliciesServiceClient(conn)

	return client.UpdatePolicy(callProperties.Ctx, req, callProperties.CallOptions...)
}

// OverwriteTCOLogsPolicies overwrites the specified TCO logs policies.
func (t TCOPoliciesClient) OverwriteTCOLogsPolicies(ctx context.Context, req *tcopolicies.AtomicOverwriteLogPoliciesRequest) (*tcopolicies.AtomicOverwriteLogPoliciesResponse, error) {
	callProperties, err := t.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := tcopolicies.NewPoliciesServiceClient(conn)

	return client.AtomicOverwriteLogPolicies(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Delete deletes the specified TCO policy.
func (t TCOPoliciesClient) Delete(ctx context.Context, req *DeletePolicyRequest) (*tcopolicies.DeletePolicyResponse, error) {
	callProperties, err := t.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := tcopolicies.NewPoliciesServiceClient(conn)

	return client.DeletePolicy(callProperties.Ctx, req, callProperties.CallOptions...)
}

// Reorder reorders the specified TCO policies.
func (t TCOPoliciesClient) Reorder(ctx context.Context, req *ReorderPoliciesRequest) (*tcopolicies.ReorderPoliciesResponse, error) {
	callProperties, err := t.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := tcopolicies.NewPoliciesServiceClient(conn)

	return client.ReorderPolicies(callProperties.Ctx, req, callProperties.CallOptions...)
}

// NewTCOPoliciesClient Creates a new TCO policies client.
func NewTCOPoliciesClient(c *CallPropertiesCreator) *TCOPoliciesClient {
	return &TCOPoliciesClient{callPropertiesCreator: c}
}

// List gets the specified TCO logs policies.
func (t TCOPoliciesClient) List(ctx context.Context, req *GetCompanyPoliciesRequest) (*tcopolicies.GetCompanyPoliciesResponse, error) {
	callProperties, err := t.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := tcopolicies.NewPoliciesServiceClient(conn)

	return client.GetCompanyPolicies(callProperties.Ctx, req, callProperties.CallOptions...)
}
