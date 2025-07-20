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

	"google.golang.org/protobuf/types/known/emptypb"

	aicenter "github.com/coralogix/coralogix-management-sdk/go/internal/coralogixapis/ai/v3"
)

// CreateAiEvaluationRequest is a request to create an AI evaluation.
type CreateAiEvaluationRequest = aicenter.CreateAiEvaluationRequest

// GetAiEvaluationRequest is a request to get an AI evaluation.
type GetAiEvaluationRequest = aicenter.GetAiEvaluationRequest

// UpdateAiEvaluationRequest is a request to update an AI evaluation.
type UpdateAiEvaluationRequest = aicenter.UpdateAiEvaluationRequest

// DeleteAiEvaluationRequest is a request to delete an AI evaluation.
type DeleteAiEvaluationRequest = aicenter.DeleteAiEvaluationRequest

// ListAiEvaluationsRequest is a request to list AI evaluations.
type ListAiEvaluationsRequest = aicenter.ListAiEvaluationsRequest

// ListAiEvaluationsResponse is a response containing a list of AI evaluations.
type ListAiEvaluationsResponse = aicenter.ListAiEvaluationsResponse

// CountAiAppsPerEvalTypeRequest is a request to count AI apps per evaluation type.
type CountAiAppsPerEvalTypeRequest = aicenter.CountAiAppsPerEvalTypeRequest

// CountAiAppsPerEvalTypeResponse is a response containing the count of AI applications per evaluation type.
type CountAiAppsPerEvalTypeResponse = aicenter.CountAiAppsPerEvalTypeResponse

// ListAiApplicationsRequest is a request to list AI applications.
type ListAiApplicationsRequest = aicenter.ListAiApplicationsRequest

// ListAiApplicationsResponse is a response containing a list of AI applications.
type ListAiApplicationsResponse = aicenter.ListAiApplicationsResponse

// AiEvaluation is an AI evaluation.
type AiEvaluation = aicenter.AiEvaluation

// EvaluationTarget is the target of an AI evaluation.
type EvaluationTarget = aicenter.EvaluationTarget

// AiApplication is an AI application.
type AiApplication = aicenter.AiApplication

// EvaluationType is the type of evaluation for an AI application.
type EvaluationType = aicenter.EvaluationType

// EvaluationConfig is the configuration for an AI evaluation.
type EvaluationConfig = aicenter.EvaluationConfig

// EvaluationConfigPromptInjection is the configuration for prompt injection evaluation.
type EvaluationConfigPromptInjection = aicenter.EvaluationConfig_PromptInjection

// add the others
// PromptInjectionConfig is the configuration for prompt injection evaluation.
type PromptInjectionConfig = aicenter.PromptInjectionConfig

// EvaluationTarget values
const (
	EvaluationTargetPrompt   = aicenter.EvaluationTarget_prompt
	EvaluationTargetResponse = aicenter.EvaluationTarget_response
)

// EvaluationType values
const (
	allowedTopics                 = aicenter.EvaluationType_allowed_topics
	competition                   = aicenter.EvaluationType_competition
	hallucinationCompleteness     = aicenter.EvaluationType_hallucination_completeness
	hallucinationContextAdherence = aicenter.EvaluationType_hallucination_context_adherence
	hallucinationContextRelevance = aicenter.EvaluationType_hallucination_context_relevance
	hallucinationCorrectness      = aicenter.EvaluationType_hallucination_correctness
	pii                           = aicenter.EvaluationType_pii
	promptInjection               = aicenter.EvaluationType_prompt_injection
	restrictedTopics              = aicenter.EvaluationType_restricted_topics
	sexism                        = aicenter.EvaluationType_sexism
	sqlAllowedTables              = aicenter.EvaluationType_sql_allowed_tables
	sqlHallucination              = aicenter.EvaluationType_sql_hallucination
	sqlReadOnly                   = aicenter.EvaluationType_sql_read_only
	sqlRestrictedTables           = aicenter.EvaluationType_sql_restricted_tables
	toxicity                      = aicenter.EvaluationType_toxicity
	hallucinationTaskAdherence    = aicenter.EvaluationType_hallucination_task_adherence
	sqlLoad                       = aicenter.EvaluationType_sql_load
	languageMismatch              = aicenter.EvaluationType_language_mismatch
	customEvaluation              = aicenter.EvaluationType_custom_evaluation
)

// apiKeysFeatureGroupID is the feature group ID for API keys.
const aiCenterFeatureGroupID = "ai"

// RPC names
const (
	CreateAiEvaluationRPC     = aicenter.AiEvaluationsService_CreateAiEvaluation_FullMethodName
	GetAiEvaluationRPC        = aicenter.AiEvaluationsService_GetAiEvaluation_FullMethodName
	UpdateAiEvaluationRPC     = aicenter.AiEvaluationsService_UpdateAiEvaluation_FullMethodName
	DeleteAiEvaluationRPC     = aicenter.AiEvaluationsService_DeleteAiEvaluation_FullMethodName
	ListAiEvaluationsRPC      = aicenter.AiEvaluationsService_ListAiEvaluations_FullMethodName
	CountAiAppsPerEvalTypeRPC = aicenter.AiEvaluationsService_CountAiAppsPerEvalType_FullMethodName
	ListAiApplicationsRPC     = aicenter.AiApplicationsService_ListAiApplications_FullMethodName
)

// AICenterClient is a client for the Coralogix AI Center API.
type AICenterClient struct {
	callPropertiesCreator *CallPropertiesCreator
}

// Create creates a new AI evaluation.
func (c AICenterClient) Create(ctx context.Context, req *CreateAiEvaluationRequest) (*AiEvaluation, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := aicenter.NewAiEvaluationsServiceClient(conn)

	resp, err := client.CreateAiEvaluation(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, CreateAiEvaluationRPC, aiCenterFeatureGroupID)
	}
	return resp, nil
}

// Get retrieves an existing AI evaluation.
func (c AICenterClient) Get(ctx context.Context, req *GetAiEvaluationRequest) (*AiEvaluation, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := aicenter.NewAiEvaluationsServiceClient(conn)

	resp, err := client.GetAiEvaluation(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, GetAiEvaluationRPC, aiCenterFeatureGroupID)
	}
	return resp, nil
}

// Update modifies an existing AI evaluation.
func (c AICenterClient) Update(ctx context.Context, req *UpdateAiEvaluationRequest) (*AiEvaluation, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := aicenter.NewAiEvaluationsServiceClient(conn)

	resp, err := client.UpdateAiEvaluation(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, UpdateAiEvaluationRPC, aiCenterFeatureGroupID)
	}
	return resp, nil
}

// Delete removes an existing AI evaluation.
func (c AICenterClient) Delete(ctx context.Context, req *DeleteAiEvaluationRequest) (*emptypb.Empty, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := aicenter.NewAiEvaluationsServiceClient(conn)

	resp, err := client.DeleteAiEvaluation(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, DeleteAiEvaluationRPC, aiCenterFeatureGroupID)
	}
	return resp, nil
}

// List lists AI evaluations.
func (c AICenterClient) List(ctx context.Context, req *ListAiEvaluationsRequest) (*ListAiEvaluationsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := aicenter.NewAiEvaluationsServiceClient(conn)

	resp, err := client.ListAiEvaluations(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ListAiEvaluationsRPC, aiCenterFeatureGroupID)
	}
	return resp, nil
}

// CountAppsPerEvalType counts AI applications per evaluation type.
func (c AICenterClient) CountAppsPerEvalType(ctx context.Context, req *CountAiAppsPerEvalTypeRequest) (*CountAiAppsPerEvalTypeResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := aicenter.NewAiEvaluationsServiceClient(conn)

	resp, err := client.CountAiAppsPerEvalType(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, CountAiAppsPerEvalTypeRPC, aiCenterFeatureGroupID)
	}
	return resp, nil
}

// ListApplications lists AI applications.
func (c AICenterClient) ListApplications(ctx context.Context, req *ListAiApplicationsRequest) (*ListAiApplicationsResponse, error) {
	callProperties, err := c.callPropertiesCreator.GetTeamsLevelCallProperties(ctx)
	if err != nil {
		return nil, err
	}

	conn := callProperties.Connection
	defer conn.Close()
	client := aicenter.NewAiApplicationsServiceClient(conn)

	resp, err := client.ListAiApplications(callProperties.Ctx, req, callProperties.CallOptions...)
	if err != nil {
		return nil, NewSdkAPIError(err, ListAiApplicationsRPC, aiCenterFeatureGroupID)
	}
	return resp, nil
}

// NewAICenterClient creates a new AI Center client.
func NewAICenterClient(c *CallPropertiesCreator) *AICenterClient {
	return &AICenterClient{callPropertiesCreator: c}
}
