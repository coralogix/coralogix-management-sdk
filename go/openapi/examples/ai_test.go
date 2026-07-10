// Copyright 2026 Coralogix Ltd.
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

package examples

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	aiapplications "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/ai_applications_service"
	aievaluations "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/ai_evaluations_service"
)

func TestAIApplications(t *testing.T) {
	cfg := newTestConfig()
	clientSet := cxsdk.NewClientSet(cfg)
	client := clientSet.AIApplications()
	require.NotNil(t, client)

	result, httpResp, err := client.
		AiApplicationsServiceListAiApplications(context.Background()).
		PageSize(10).
		PageOffset(0).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, result)

	for _, app := range result.GetAiApplications() {
		if app.Id == nil {
			continue
		}

		getResp, httpResp, err := client.
			AiApplicationsServiceGetAiApplicationById(context.Background(), *app.Id).
			Execute()
		require.NoError(t, cxsdk.NewAPIError(httpResp, err))
		gotApp := getResp.GetAiApplication()
		require.Equal(t, *app.Id, gotApp.GetId())
		return
	}
}

func TestAIEvaluations(t *testing.T) {
	cfg := newTestConfig()
	clientSet := cxsdk.NewClientSet(cfg)
	client := clientSet.AIEvaluations()
	require.NotNil(t, client)

	app, ok := firstAIApplication(t, clientSet.AIApplications())
	if !ok {
		t.Fatal("no existing AI application found")
	}

	createdThreshold := 0.8
	config := aievaluations.EvaluationConfig{
		Pii: &aievaluations.PiiConfig{
			Categories: []aievaluations.PiiCategory{
				aievaluations.PIICATEGORY_EMAIL_ADDRESS,
				aievaluations.PIICATEGORY_CREDIT_CARD,
			},
		},
	}

	createReq := aievaluations.AiEvaluationsServiceCreateAiEvaluationRequest{
		Application: aievaluations.PtrString(app.application),
		Target:      aievaluations.EVALUATIONTARGET_RESPONSE.Ptr(),
		Threshold:   aievaluations.PtrFloat64(createdThreshold),
		IsEnabled:   aievaluations.PtrBool(true),
		Config:      &config,
	}
	if app.subsystem != "" {
		createReq.Subsystem = aievaluations.PtrString(app.subsystem)
	}

	createResp, httpResp, err := client.
		AiEvaluationsServiceCreateAiEvaluation(context.Background()).
		AiEvaluationsServiceCreateAiEvaluationRequest(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	createdEvaluation := createResp.GetAiEvaluation()
	require.NotEmpty(t, createdEvaluation.GetId())

	evaluationID := createdEvaluation.GetId()
	deletedEvaluation := false
	t.Cleanup(func() {
		if !deletedEvaluation {
			deleteAIEvaluation(t, client, evaluationID)
		}
	})
	require.Equal(t, app.application, createdEvaluation.GetApplication())
	if app.subsystem != "" {
		require.Equal(t, app.subsystem, createdEvaluation.GetSubsystem())
	}
	require.Equal(t, createdThreshold, createdEvaluation.GetThreshold())
	requirePIICategories(t, createdEvaluation, []aievaluations.PiiCategory{
		aievaluations.PIICATEGORY_EMAIL_ADDRESS,
		aievaluations.PIICATEGORY_CREDIT_CARD,
	})

	getResp, httpResp, err := client.
		AiEvaluationsServiceGetAiEvaluation(context.Background(), evaluationID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	gotEvaluation := getResp.GetAiEvaluation()
	require.Equal(t, evaluationID, gotEvaluation.GetId())

	updatedConfig := aievaluations.EvaluationConfig{
		Pii: &aievaluations.PiiConfig{
			Categories: []aievaluations.PiiCategory{
				aievaluations.PIICATEGORY_PHONE_NUMBER,
				aievaluations.PIICATEGORY_US_SSN,
			},
		},
	}
	updateReq := aievaluations.AiEvaluationsServiceUpdateAiEvaluationRequest{
		Config:     &updatedConfig,
		UpdateMask: aievaluations.PtrString("config"),
	}
	updateResp, httpResp, err := client.
		AiEvaluationsServiceUpdateAiEvaluation(context.Background(), evaluationID).
		AiEvaluationsServiceUpdateAiEvaluationRequest(updateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	updatedEvaluation := updateResp.GetAiEvaluation()
	requirePIICategories(t, updatedEvaluation, []aievaluations.PiiCategory{
		aievaluations.PIICATEGORY_PHONE_NUMBER,
		aievaluations.PIICATEGORY_US_SSN,
	})

	listReq := client.
		AiEvaluationsServiceListAiEvaluations(context.Background()).
		Application(app.application).
		EvaluationType(aievaluations.EVALUATIONTYPE_PII).
		PageSize(10).
		PageOffset(0)
	if app.subsystem != "" {
		listReq = listReq.Subsystem(app.subsystem)
	}
	evaluations, httpResp, err := listReq.Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, evaluations)
	require.True(t, aiEvaluationExists(evaluations.GetAiEvaluations(), evaluationID))

	counts, httpResp, err := client.
		AiEvaluationsServiceCountAiAppsPerEvalType(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, counts)

	deleteAIEvaluation(t, client, evaluationID)
	deletedEvaluation = true

	_, httpResp, err = client.
		AiEvaluationsServiceGetAiEvaluation(context.Background(), evaluationID).
		Execute()
	apiErr := cxsdk.NewAPIError(httpResp, err)
	require.Equalf(t, http.StatusNotFound, cxsdk.Code(apiErr), "expected deleted AI evaluation to be gone, got %v", apiErr)
}

func TestAICustomEvaluations(t *testing.T) {
	cfg := newTestConfig()
	clientSet := cxsdk.NewClientSet(cfg)
	client := clientSet.AIEvaluations()
	require.NotNil(t, client)

	_, httpResp, err := client.
		AiEvaluationsServiceGetCustomEvaluations(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	name := "sdk-openapi-custom-ai-" + uuid.NewString()

	createReq := aievaluations.AiEvaluationsServiceCreateCustomEvaluationRequest{
		Name:                      aievaluations.PtrString(name),
		Description:               aievaluations.PtrString("OpenAPI SDK custom evaluation example"),
		Instructions:              aievaluations.PtrString("Score 1 when the response mentions a competitor, otherwise score 0."),
		PolicyType:                aievaluations.PtrString("competition"),
		Safe:                      aievaluations.PtrString("does not mention competitors"),
		Violates:                  aievaluations.PtrString("mentions a competitor product"),
		ShouldIncludeSystemPrompt: aievaluations.PtrBool(false),
		Examples: []aievaluations.CustomEvaluationExample{
			{
				Conversation: aievaluations.PtrString("User: which tool is best? Assistant: CompetitorX is great."),
				Score:        aievaluations.PtrString("1"),
			},
		},
	}

	createResp, httpResp, err := client.
		AiEvaluationsServiceCreateCustomEvaluation(context.Background()).
		AiEvaluationsServiceCreateCustomEvaluationRequest(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	createdCustomEvaluation := createResp.GetItem()
	require.NotEmpty(t, createdCustomEvaluation.GetId())

	customEvaluationID := createdCustomEvaluation.GetId()
	deletedCustomEvaluation := false
	t.Cleanup(func() {
		if !deletedCustomEvaluation {
			deleteCustomEvaluation(t, client, customEvaluationID)
		}
	})
	require.Equal(t, name, createdCustomEvaluation.GetName())

	updatedName := name + " updated"
	updateReq := aievaluations.AiEvaluationsServiceUpdateCustomEvaluationRequest{
		Name:       aievaluations.PtrString(updatedName),
		UpdateMask: aievaluations.PtrString("name"),
	}
	updateResp, httpResp, err := client.
		AiEvaluationsServiceUpdateCustomEvaluation(context.Background(), customEvaluationID).
		AiEvaluationsServiceUpdateCustomEvaluationRequest(updateReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	updatedCustomEvaluation := updateResp.GetItem()
	require.Equal(t, updatedName, updatedCustomEvaluation.GetName())

	customEvaluations, httpResp, err := client.
		AiEvaluationsServiceGetCustomEvaluations(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.True(t, customEvaluationExists(customEvaluations.GetItems(), customEvaluationID))

	app, ok := firstAIApplication(t, clientSet.AIApplications())
	if !ok {
		t.Fatal("no existing AI application found")
	}

	linkedToApp := false
	t.Cleanup(func() {
		if linkedToApp {
			unlinkCustomEvaluationFromApp(t, client, customEvaluationID, app.id)
		}
	})

	_, httpResp, err = client.
		AiEvaluationsServiceLinkCustomEvaluation(context.Background(), customEvaluationID, app.id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	linkedToApp = true

	linkedEvaluations, httpResp, err := client.
		AiEvaluationsServiceListCustomEvaluationsForApplication(context.Background(), app.id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.True(t, customEvaluationExists(linkedEvaluations.GetItems(), customEvaluationID))

	unlinkCustomEvaluationFromApp(t, client, customEvaluationID, app.id)
	linkedToApp = false

	linkedEvaluations, httpResp, err = client.
		AiEvaluationsServiceListCustomEvaluationsForApplication(context.Background(), app.id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.False(t, customEvaluationExists(linkedEvaluations.GetItems(), customEvaluationID))

	deleteCustomEvaluation(t, client, customEvaluationID)
	deletedCustomEvaluation = true

	customEvaluations, httpResp, err = client.
		AiEvaluationsServiceGetCustomEvaluations(context.Background()).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.False(t, customEvaluationExists(customEvaluations.GetItems(), customEvaluationID))
}

func TestAIEvaluationConfigDecodesPromptInjection(t *testing.T) {
	payload := []byte(`{
		"promptInjection": {
			"additionalContext": "Only inspect the user prompt."
		}
	}`)

	var config aievaluations.EvaluationConfig
	require.NoError(t, json.Unmarshal(payload, &config))

	promptInjectionConfig, ok := config.GetPromptInjectionOk()
	require.True(t, ok)
	require.Equal(t, "Only inspect the user prompt.", promptInjectionConfig.GetAdditionalContext())
}

func TestAIEvaluationConfigDecodesCustomEvaluation(t *testing.T) {
	payload := []byte(`{
		"customEvaluation": {
			"instructions": "Score 1 when the response mentions a competitor.\nScore 0 otherwise.",
			"policyType": "competition",
			"safe": "does not mention competitors",
			"violates": "mentions a competitor product",
			"shouldIncludeSystemPrompt": false,
			"examples": [
				{
					"conversation": "User: which tool is best?\nAssistant: CompetitorX is great.",
					"score": "1"
				}
			]
		}
	}`)

	var config aievaluations.EvaluationConfig
	require.NoError(t, json.Unmarshal(payload, &config))

	customEvaluationConfig, ok := config.GetCustomEvaluationOk()
	require.True(t, ok)
	require.Equal(t, "competition", customEvaluationConfig.GetPolicyType())
	require.Len(t, customEvaluationConfig.GetExamples(), 1)
}

func deleteAIEvaluation(t *testing.T, client *aievaluations.AIEvaluationsServiceAPIService, id string) {
	t.Helper()

	_, httpResp, err := client.
		AiEvaluationsServiceDeleteAiEvaluation(context.Background(), id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func deleteCustomEvaluation(t *testing.T, client *aievaluations.AIEvaluationsServiceAPIService, id string) {
	t.Helper()

	_, httpResp, err := client.
		AiEvaluationsServiceDeleteCustomEvaluation(context.Background(), id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func unlinkCustomEvaluationFromApp(t *testing.T, client *aievaluations.AIEvaluationsServiceAPIService, id, applicationID string) {
	t.Helper()

	_, httpResp, err := client.
		AiEvaluationsServiceUnlinkCustomEvaluationFromApp(context.Background(), id, applicationID).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}

func requirePIICategories(t *testing.T, evaluation aievaluations.AiEvaluation, expected []aievaluations.PiiCategory) {
	t.Helper()

	config := evaluation.GetConfig()
	pii, ok := config.GetPiiOk()
	require.True(t, ok)
	require.ElementsMatch(t, expected, pii.GetCategories())
}

func aiEvaluationExists(evaluations []aievaluations.AiEvaluation, id string) bool {
	for _, evaluation := range evaluations {
		if evaluation.GetId() == id {
			return true
		}
	}
	return false
}

func customEvaluationExists(evaluations []aievaluations.CustomEvaluation, id string) bool {
	for _, evaluation := range evaluations {
		if evaluation.GetId() == id {
			return true
		}
	}
	return false
}

type aiApplicationRef struct {
	id          string
	application string
	subsystem   string
}

func firstAIApplication(t *testing.T, client *aiapplications.AIApplicationsServiceAPIService) (aiApplicationRef, bool) {
	t.Helper()

	result, httpResp, err := client.
		AiApplicationsServiceListAiApplications(context.Background()).
		PageSize(10).
		PageOffset(0).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	for _, app := range result.GetAiApplications() {
		if app.GetId() != "" && app.GetApplication() != "" {
			return aiApplicationRef{
				id:          app.GetId(),
				application: app.GetApplication(),
				subsystem:   app.GetSubsystem(),
			}, true
		}
	}
	return aiApplicationRef{}, false
}
