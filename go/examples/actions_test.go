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

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestActions(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator, err := cxsdk.NewSDKCallPropertiesCreator(region, authContext)
	assertNilAndPrintError(t, err)
	c := cxsdk.NewActionsClient(creator)
	defer creator.CloseConnection()

	action, e := c.Create(context.Background(), &cxsdk.CreateActionRequest{
		Name:             wrapperspb.String("google search action " + strconv.FormatInt(time.Now().UnixMilli(), 10)),
		Url:              wrapperspb.String("https://www.google.com/search?q={{$p.selected_value}}"),
		IsPrivate:        wrapperspb.Bool(false),
		SourceType:       cxsdk.SourceTypeLog,
		ApplicationNames: []*wrapperspb.StringValue{},
		SubsystemNames:   []*wrapperspb.StringValue{},
	})
	assertNilAndPrintError(t, e)

	_, e = c.Replace(context.Background(), &cxsdk.ReplaceActionRequest{
		Action: &cxsdk.Action{
			Id:               action.Action.Id,
			Name:             wrapperspb.String("updated action " + strconv.FormatInt(time.Now().UnixMilli(), 10)),
			Url:              wrapperspb.String("https://www.bing.com/search?q={{$p.selected_value}}"),
			IsPrivate:        wrapperspb.Bool(false),
			IsHidden:         wrapperspb.Bool(false),
			SourceType:       cxsdk.SourceTypeLog,
			ApplicationNames: []*wrapperspb.StringValue{},
			SubsystemNames:   []*wrapperspb.StringValue{},
		},
	})

	assertNilAndPrintError(t, e)

	updated, _ := c.Get(context.Background(), &cxsdk.GetActionRequest{
		Id: action.Action.Id,
	})

	assert.Equal(t, updated.Action.Url.Value, "https://www.bing.com/search?q={{$p.selected_value}}")

	c.Delete(context.Background(), &cxsdk.DeleteActionRequest{
		Id: action.Action.Id,
	})
}
