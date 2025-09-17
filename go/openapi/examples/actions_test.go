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

package examples

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go/openapi"
)

func TestActions(t *testing.T) {
	region := os.Getenv(cxsdk.EnvCoralogixRegion)
	if region == "" {
		t.Fatalf("env %s not set", cxsdk.EnvCoralogixRegion)
	}

	apiKey := os.Getenv(cxsdk.EnvCoralogixAPIKey)
	if apiKey == "" {
		t.Fatalf("env %s not set", cxsdk.EnvCoralogixAPIKey)
	}

	client := cxsdk.NewClientBuilder().
		WithRegion(region).
		WithAPIKey(apiKey).
		Build()

	name := "google search action " + strconv.FormatInt(time.Now().UnixMilli(), 10)
	url := "https://www.google.com/search?q={{$p.selected_value}}"
	sourceType := "SOURCE_TYPE_LOG"
	isPrivate := false

	createReq := cxsdk.ActionsServiceCreateActionRequest{
		Name:             &name,
		Url:              &url,
		IsPrivate:        &isPrivate,
		SourceType:       &sourceType,
		ApplicationNames: []string{},
		SubsystemNames:   []string{},
	}

	created, _, err := client.ActionsServiceAPI.
		ActionsServiceCreateAction(context.Background()).
		ActionsServiceCreateActionRequest(createReq).
		Execute()
	assert.NoError(t, err)
	assert.NotNil(t, created.Action.Id)

	newName := "updated action " + strconv.FormatInt(time.Now().UnixMilli(), 10)
	newURL := "https://www.bing.com/search?q={{$p.selected_value}}"
	isHidden := false

	replaceReq := cxsdk.ActionsServiceReplaceActionRequest{
		Action: &cxsdk.ComCoralogixapisActionsV2Action{
			Id:               created.Action.Id,
			Name:             &newName,
			Url:              &newURL,
			IsPrivate:        &isPrivate,
			IsHidden:         &isHidden,
			SourceType:       &sourceType,
			ApplicationNames: []string{},
			SubsystemNames:   []string{},
		},
	}

	_, _, err = client.ActionsServiceAPI.
		ActionsServiceReplaceAction(context.Background()).
		ActionsServiceReplaceActionRequest(replaceReq).
		Execute()
	assert.NoError(t, err)

	updated, _, err := client.ActionsServiceAPI.
		ActionsServiceGetAction(context.Background(), *created.Action.Id).
		Execute()
	assert.NoError(t, err)
	assert.Equal(t, newURL, *updated.Action.Url)

	_, _, err = client.ActionsServiceAPI.
		ActionsServiceDeleteAction(context.Background(), *created.Action.Id).
		Execute()
	assert.NoError(t, err)
}
