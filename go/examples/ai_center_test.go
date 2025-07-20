package examples

import (
	"context"
	"testing"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
	"github.com/stretchr/testify/assert"
)

func TestAiEvaluation(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assert.Nil(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assert.Nil(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	client := cxsdk.NewAICenterClient(creator)

	app := "test-app"
	subsystem := "test-subsystem"
	createReq := &cxsdk.CreateAiEvaluationRequest{
		Application: app,
		Subsystem:   subsystem,
		Target:      cxsdk.EvaluationTargetPrompt,
		IsEnabled:   true,
		Config: &cxsdk.EvaluationConfig{
			Evaluation: &cxsdk.EvaluationConfigPromptInjection{
				PromptInjection: &cxsdk.PromptInjectionConfig{},
			},
		},
	}

	createdRes, err := client.Create(context.Background(), createReq)
	if err != nil {
		t.Fatal(err)
	}

	aiEvalID := createdRes.Id
	assert.NotEmpty(t, aiEvalID, "Expected AI Evaluation ID to be non-empty")

	getRes, err := client.Get(context.Background(), &cxsdk.GetAiEvaluationRequest{
		Id: aiEvalID,
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, getRes.Application, app, "Expected AI Evaluation application to match")
	assert.Equal(t, getRes.Subsystem, subsystem, "Expected AI Evaluation subsystem to match")
}
