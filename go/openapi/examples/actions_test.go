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
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/coralogix/coralogix-management-sdk/go/openapi/cxsdk"
	actions "github.com/coralogix/coralogix-management-sdk/go/openapi/gen/actions_service"
)

func TestActions(t *testing.T) {
	cfg := cxsdk.NewConfigBuilder().WithAPIKeyEnv().WithRegionEnv().WithHTTPLogging().Build()
	client := cxsdk.NewActionsClient(cfg)

	name := "google search action " + strconv.FormatInt(time.Now().UnixMilli(), 10)
	url := "https://www.google.com/search?q={{$p.selected_value}}"
	isPrivate := false

	createReq := actions.ActionsServiceCreateActionRequest{
		Name:             &name,
		Url:              &url,
		IsPrivate:        &isPrivate,
		SourceType:       actions.V2SOURCETYPE_SOURCE_TYPE_LOG.Ptr(),
		ApplicationNames: []string{},
		SubsystemNames:   []string{},
	}

	created, httpResp, err := client.
		ActionsServiceCreateAction(context.Background()).
		ActionsServiceCreateActionRequest(createReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.NotNil(t, created.Action.Id)

	newName := "updated action " + strconv.FormatInt(time.Now().UnixMilli(), 10)
	newURL := "https://www.bing.com/search?q={{$p.selected_value}}"
	isHidden := false

	replaceReq := actions.ActionsServiceReplaceActionRequest{
		Action: &actions.V2Action{
			Id:               created.Action.Id,
			Name:             &newName,
			Url:              &newURL,
			IsPrivate:        &isPrivate,
			IsHidden:         &isHidden,
			SourceType:       actions.V2SOURCETYPE_SOURCE_TYPE_LOG.Ptr(),
			ApplicationNames: []string{},
			SubsystemNames:   []string{},
		},
	}

	_, httpResp, err = client.
		ActionsServiceReplaceAction(context.Background()).
		ActionsServiceReplaceActionRequest(replaceReq).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))

	updated, httpResp, err := client.
		ActionsServiceGetAction(context.Background(), *created.Action.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
	require.Equal(t, newURL, *updated.Action.Url)

	_, httpResp, err = client.
		ActionsServiceDeleteAction(context.Background(), *created.Action.Id).
		Execute()
	require.NoError(t, cxsdk.NewAPIError(httpResp, err))
}
