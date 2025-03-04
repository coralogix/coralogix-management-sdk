// Copyright 2025 Coralogix Ltd.
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
	"testing"

	cxsdk "github.com/coralogix/coralogix-management-sdk/go"
)

func TestViews(t *testing.T) {
	region, err := cxsdk.CoralogixRegionFromEnv()
	assertNilAndPrintError(t, err)
	authContext, err := cxsdk.AuthContextFromEnv()
	assertNilAndPrintError(t, err)
	creator := cxsdk.NewCallPropertiesCreator(region, authContext)
	c := cxsdk.NewViewsClient(creator)

	action, e := c.Create(context.Background(), &cxsdk.CreateViewRequest{
		Name: wrapperspb.StringValue("MyView"), 
		SearchQuery: &cxsdk.SearchQuery{
			Query: wrapperspb.StringValue("source logs
| filter $l.applicationname == 'default'"), 
		}, 
		TimeSelection: &cxsdk.TimeSelection{
			SelectionType: cxsdk.ViewTimeSelectionQuick{
				QuickSelection: cxsdk.QuickTimeSelection {
					Seconds: uint32(86400), // 24h
				},
			},
		}, 
		Filters: []*cxsdk.ViewFilter{},
	})
	assertNilAndPrintError(t, e)

	_, e = c.Replace(context.Background(), &cxsdk.ReplaceViewRequest{})

	assertNilAndPrintError(t, e)

	updated, _ := c.Get(context.Background(), &cxsdk.GetViewRequest{
		Id: 
	})

	c.Delete(context.Background(), &cxsdk.DeleteViewRequest{
		Id: 
	})
}
