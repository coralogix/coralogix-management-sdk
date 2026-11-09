// Copyright 2026 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sdkcheck

import (
	"os"
	"testing"

	"github.com/coralogix/coralogix-management-sdk/iac/model"
)

func TestBatchGetSDKNames(t *testing.T) {
	data, err := os.ReadFile("../model/testdata/batchthing.yaml")
	if err != nil {
		t.Fatal(err)
	}
	doc, err := model.Load(data)
	if err != nil {
		t.Fatal(err)
	}
	r, err := model.BuildWithPolicy(doc, "BatchThing", model.OperationIDs{}, model.Policy{Existing: true, GetByIDs: "ids", UpdateIDInBody: true, ReadOnly: []string{"BatchThing.id"}})
	if err != nil {
		t.Fatal(err)
	}
	refs, err := Names(r, "Batch Things Service", "example/sdk", nil)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]Ref{
		"get": {
			Kind: KindMethod, Owner: "BatchThingsServiceAPIService", Name: "BatchThingsServiceGetBatchThing",
			Want: "func(ctx context.Context) ApiBatchThingsServiceGetBatchThingRequest",
		},
		"get.ids": {
			Kind: KindMethod, Owner: "ApiBatchThingsServiceGetBatchThingRequest", Name: "Ids",
			Want: "func(ids []string) ApiBatchThingsServiceGetBatchThingRequest",
		},
		"get.response.things": {
			Kind: KindField, Owner: "GetBatchThingResponse", Name: "Things", Want: "[]BatchThing",
		},
	}
	for _, ref := range refs {
		expected, ok := want[ref.Path]
		if !ok || ref.Kind == KindType {
			continue
		}
		if ref.Kind != expected.Kind || ref.Owner != expected.Owner || ref.Name != expected.Name || ref.Want != expected.Want {
			t.Fatalf("%s: %+v, want %+v", ref.Path, ref, expected)
		}
		delete(want, ref.Path)
	}
	if len(want) != 0 {
		t.Fatalf("missing refs: %v", want)
	}
}
