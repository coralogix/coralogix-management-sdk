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

package model

import (
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"testing"

	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

// listsPolicy is the rule set of a resource of testdata/lists.yaml. Preset has a custom Create.
func listsPolicy(name string) Policy {
	return Policy{CustomMethods: name == "Preset"}
}

// listsDoc returns testdata/lists.yaml with each pair of old and new text replaced.
func listsDoc(t *testing.T, replacements ...string) *v3.Document {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", "lists.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	spec := string(data)
	for i := 0; i < len(replacements); i += 2 {
		if !strings.Contains(spec, replacements[i]) {
			t.Fatalf("the spec has no %q", replacements[i])
		}
		spec = strings.Replace(spec, replacements[i], replacements[i+1], 1)
	}
	doc, err := Load([]byte(spec))
	if err != nil {
		t.Fatal(err)
	}
	return doc
}

func TestListDiscovery(t *testing.T) {
	list := func(id, path, response, items, item string, full bool, name string, paged bool) *ListOperation {
		return &ListOperation{
			Operation:  Operation{Method: "GET", Path: path, OperationID: id, Response: Response{Schema: response}},
			ItemsField: items, ItemComponent: item, Full: full, IDField: "id", NameField: name, Paged: paged,
		}
	}
	tests := map[string]struct {
		list    string // OperationIDs.List
		want    *ListOperation
		problem string
	}{
		"Connector":    {want: list("ConnectorsService_ListConnectors", "/connectors", "ListConnectorsResponse", "connectors", "Connector", true, "name", false)},
		"GlobalRouter": {want: list("GlobalRoutersService_ListGlobalRouters", "/routers", "ListGlobalRoutersResponse", "routers", "GlobalRouter", true, "name", false)},
		// The List of a custom Create is not on the Create path, so only its name finds it.
		"Preset":     {},
		"ViewFolder": {want: list("FoldersService_ListViewFolders", "/folders", "ListViewFoldersResponse", "folders", "ViewFolder", true, "name", false)},
		"Evaluation": {want: list("EvaluationsService_ListEvaluations", "/evaluations", "ListEvaluationsResponse", "evaluations", "Evaluation", true, "name", true)},
		"Label":      {want: list("LabelsService_ListLabels", "/labels", "ListLabelsResponse", "labels", "LabelSummary", false, "", false)},
		// A List that the model cannot describe is no issue: the resource stays eligible.
		"Pair": {problem: "The response ListPairsResponse has more than one array of component objects: pairs, archived."},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			doc := listsDoc(t)
			ids := OperationIDs{List: test.list}
			if report := ValidateWithPolicy(doc, name, ids, listsPolicy(name)); len(report) != 0 {
				t.Fatalf("issues:\n%s", report.Error())
			}
			r, err := BuildWithPolicy(doc, name, ids, listsPolicy(name))
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(r.List, test.want) || r.ListProblem != test.problem {
				t.Fatalf("List = %+v, problem %q\nwant %+v, problem %q", r.List, r.ListProblem, test.want, test.problem)
			}
		})
	}
}

// The overrides file or a flag names a List that is not on the Create path.
func TestListExplicitSelection(t *testing.T) {
	want := &ListOperation{
		Operation:  Operation{Method: "GET", Path: "/presets:summariesList", OperationID: "PresetsService_ListPresetSummaries", Response: Response{Schema: "ListPresetSummariesResponse"}},
		ItemsField: "presetSummaries", ItemComponent: "PresetSummary", IDField: "id", NameField: "name",
	}
	fromFile := listsPolicy("Preset")
	fromFile.Operations.List = want.OperationID
	for name, test := range map[string]struct {
		ids    OperationIDs
		policy Policy
	}{
		"flag":      {OperationIDs{List: want.OperationID}, listsPolicy("Preset")},
		"overrides": {OperationIDs{}, fromFile},
	} {
		t.Run(name, func(t *testing.T) {
			r, err := BuildWithPolicy(listsDoc(t), "Preset", test.ids, test.policy)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(r.List, want) || r.ListProblem != "" {
				t.Fatalf("List = %+v, problem %q, want %+v", r.List, r.ListProblem, want)
			}
		})
	}
	conflict := fromFile
	conflict.Operations.List = "PresetsService_Other"
	if _, err := BuildWithPolicy(listsDoc(t), "Preset", OperationIDs{List: want.OperationID}, conflict); err == nil || !strings.Contains(err.Error(), "List: the overrides name") {
		t.Fatalf("err = %v, want the List conflict", err)
	}
}

// A List that a name selects and that the model cannot describe is an issue, because the user
// asked for it. Without the name, the same List is only a ListProblem.
func TestListExplicitIssues(t *testing.T) {
	tests := map[string]struct {
		resource, list string
		replacements   []string
		code           string
	}{
		"not found":          {"ViewFolder", "FoldersService_Missing", nil, "LIST_OPERATION_NOT_FOUND"},
		"not a GET":          {"ViewFolder", "FoldersService_CreateViewFolder", nil, "LIST_OPERATION_METHOD_INCOMPATIBLE"},
		"path parameter":     {"ViewFolder", "FoldersService_GetViewFolder", nil, "LIST_PATH_PARAMETER_UNSUPPORTED"},
		"two arrays":         {"Pair", "PairsService_ListPairs", nil, "LIST_ITEMS_AMBIGUOUS"},
		"no array":           {"Evaluation", "EvaluationsService_ListEvaluations", []string{"        evaluations:\n          type: array\n          items: {$ref: '#/components/schemas/Evaluation'}\n", ""}, "LIST_ITEMS_NOT_FOUND"},
		"item without an id": {"Label", "LabelsService_ListLabels", []string{"    LabelSummary:\n      type: object\n      properties:\n        id: {type: string}\n", "    LabelSummary:\n      type: object\n      properties:\n"}, "LIST_ITEM_ID_MISSING"},
		"item id type":       {"Label", "LabelsService_ListLabels", []string{"    LabelSummary:\n      type: object\n      properties:\n        id: {type: string}\n", "    LabelSummary:\n      type: object\n      properties:\n        id: {type: integer, format: int32}\n"}, "LIST_ITEM_ID_TYPE_INCOMPATIBLE"},
		"other tag":          {"ViewFolder", "FoldersService_ListViewFolders", []string{"      tags: [Folders Service]\n      operationId: FoldersService_ListViewFolders", "      tags: [Views Service]\n      operationId: FoldersService_ListViewFolders"}, "LIST_TAG_INCOMPATIBLE"},
		"required filter":    {"Connector", "ConnectorsService_ListConnectors", []string{"          name: supported_by_entity_type\n          required: false", "          name: supported_by_entity_type\n          required: true"}, "LIST_REQUIRED_PARAMETER_UNSUPPORTED"},
		"inline response":    {"ViewFolder", "FoldersService_ListViewFolders", []string{"schema: {$ref: '#/components/schemas/ListViewFoldersResponse'}", "schema: {type: object}"}, "LIST_RESPONSE_UNSUPPORTED"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			doc := listsDoc(t, test.replacements...)
			policy := listsPolicy(test.resource)
			codes := reportCodes(ValidateWithPolicy(doc, test.resource, OperationIDs{List: test.list}, policy))
			if !slices.Equal(codes, []string{test.code}) {
				t.Fatalf("codes = %v, want [%s]", codes, test.code)
			}
			if _, err := BuildWithPolicy(doc, test.resource, OperationIDs{List: test.list}, policy); err == nil {
				t.Fatal("Build passed, want the List error")
			}
			// Discovery finds only the GET on the Create path, so an explicit name of another
			// operation has nothing to compare with.
			if !strings.Contains(test.list, "List") {
				return
			}
			if report := ValidateWithPolicy(doc, test.resource, OperationIDs{}, policy); len(report) != 0 {
				t.Fatalf("discovery issues:\n%s", report.Error())
			}
			r, err := BuildWithPolicy(doc, test.resource, OperationIDs{}, policy)
			if err != nil {
				t.Fatal(err)
			}
			if r.List != nil || r.ListProblem == "" {
				t.Fatalf("discovered List = %+v, problem %q, want a problem", r.List, r.ListProblem)
			}
		})
	}
}

// Two GET operations on the collection path, with and without a trailing "/", are ambiguous.
func TestListDiscoveryAmbiguous(t *testing.T) {
	doc := listsDoc(t, "  /folders/{id}:\n", `  /folders/:
    get:
      tags: [Folders Service]
      operationId: FoldersService_ListViewFoldersAgain
      responses:
        '200': {description: ok, content: {application/json: {schema: {$ref: '#/components/schemas/ListViewFoldersResponse'}}}}
  /folders/{id}:
`)
	r, err := Build(doc, "ViewFolder")
	if err != nil {
		t.Fatal(err)
	}
	want := "More than one GET operation lists the resources: FoldersService_ListViewFolders, FoldersService_ListViewFoldersAgain."
	if r.List != nil || r.ListProblem != want {
		t.Fatalf("List = %+v, problem %q, want %q", r.List, r.ListProblem, want)
	}
}

// A resource whose Create path has no GET has no List and no problem.
func TestNoListOnTheCreatePath(t *testing.T) {
	doc, err := Load(gadgetSpec(t))
	if err != nil {
		t.Fatal(err)
	}
	r, err := Build(doc, "Gadget")
	if err != nil {
		t.Fatal(err)
	}
	if r.List != nil || r.ListProblem != "" {
		t.Fatalf("List = %+v, problem %q, want none", r.List, r.ListProblem)
	}
}

func TestPagingNames(t *testing.T) {
	for name, want := range map[string]bool{
		"page_token": true, "pageToken": true, "pageSize": true, "page_size": true, "pageOffset": true,
		"page_offset": true, "next_page_token": true, "limit": true, "offset": true, "pagination": true,
		"pagination.page_size": true, "pagination.pageToken": true,
		"connector_type": false, "entity_type": false, "order_by": false, "page": false,
	} {
		if got := pagingName(name); got != want {
			t.Errorf("pagingName(%q) = %t, want %t", name, got, want)
		}
	}
}

// A response with a next page token pages the List, also when no query parameter does.
func TestListPagedByResponse(t *testing.T) {
	doc := listsDoc(t, "    ListViewFoldersResponse:\n      type: object\n      properties:\n", "    ListViewFoldersResponse:\n      type: object\n      properties:\n        next_page_token: {type: string}\n")
	r, err := Build(doc, "ViewFolder")
	if err != nil {
		t.Fatal(err)
	}
	if r.List == nil || !r.List.Paged {
		t.Fatalf("List = %+v, want a paged List", r.List)
	}
}
