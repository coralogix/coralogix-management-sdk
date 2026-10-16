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

package overrides

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"go.yaml.in/yaml/v4"

	"github.com/coralogix/coralogix-management-sdk/iac/model"
)

func TestPolicyStatesTheAPIKeys(t *testing.T) {
	on, off := true, false
	e := Exceptions{
		API: API{
			RequestWrapper: "router", UpdateIDInBody: true, ClientSetID: true, CustomMethods: true, UpdateMaskInBody: true,
			Operations: &Operations{Create: "C", Get: "G", Update: "U", Delete: "D"},
		},
		Unwrap: []string{"LuceneQuery"},
		Fields: map[string]map[string]Field{
			"Thing":  {"id": {ReadOnly: true}, "query": {Unwrap: &off}},
			"Target": {"id": {Skip: true}, "label": {Unwrap: &on}},
		},
	}
	want := model.Policy{
		Existing: true, RequestWrapper: "router", UpdateIDInBody: true, ClientSetID: true, CustomMethods: true, UpdateMaskInBody: true,
		Operations:   model.OperationIDs{Create: "C", Get: "G", Update: "U", Delete: "D"},
		Unwrap:       []string{"LuceneQuery"},
		UnwrapFields: map[string]bool{"Target.label": true, "Thing.query": false},
		Skip:         []string{"Target.id"},
		ReadOnly:     []string{"Thing.id"},
		Released:     []string{"Target.id", "Thing.id"},
	}
	if got := e.Policy(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Policy() =\n%+v\nwant\n%+v", got, want)
	}
	if got := (Exceptions{API: API{Delete: &Delete{Operation: "Archive"}}}).Policy(); got.DeleteOperation != "Archive" {
		t.Fatalf("DeleteOperation = %q", got.DeleteOperation)
	}
}

func TestPolicyStatesTheFieldShapeKeys(t *testing.T) {
	e := Exceptions{
		API: API{IDProperty: "uniqueIdentifier"},
		Fields: map[string]map[string]Field{
			"Thing": {
				"name":   {Optional: true},
				"tags":   {Collection: "set"},
				"labels": {KeyedBy: &KeyedBy{Key: "key", Value: "value"}},
				"limits": {CreateAs: map[string]string{"limit": "permutationsLimit"}},
			},
		},
	}
	want := model.Policy{
		Existing: true, IDProperty: "uniqueIdentifier",
		Optional: []string{"Thing.name"},
		Sets:     []string{"Thing.tags"},
		KeyedBy:  map[string]model.KeyedLine{"Thing.labels": {Key: "key", Value: "value"}},
		CreateAs: map[string]map[string]string{"Thing.limits": {"limit": "permutationsLimit"}},
	}
	if got := e.Policy(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Policy() =\n%+v\nwant\n%+v", got, want)
	}
}

func TestFieldCheckRejectsBadShapeKeys(t *testing.T) {
	for name, l := range map[string]Field{
		"keyedBy without a key":        {KeyedBy: &KeyedBy{Value: "value"}},
		"keyedBy value is the key":     {KeyedBy: &KeyedBy{Key: "key", Value: "key"}},
		"keyedBy object with value":    {KeyedBy: &KeyedBy{Key: "key", Value: "value", Object: true}},
		"keyedBy arms without object":  {KeyedBy: &KeyedBy{Key: "key", Arms: map[string]string{"A": "a"}}},
		"keyedBy and collection":       {KeyedBy: &KeyedBy{Key: "key"}, Collection: "set"},
		"createAs with an empty name":  {CreateAs: map[string]string{"limit": ""}},
		"createAs names a field twice": {CreateAs: map[string]string{"a": "x", "b": "x"}},
		"collection other than set":    {Collection: "list"},
	} {
		if l.Check() == nil {
			t.Errorf("%s: Check() = nil", name)
		}
	}
	ok := Field{Collection: "set", CreateAs: map[string]string{"limit": "permutationsLimit"}}
	if err := ok.Check(); err != nil {
		t.Fatal(err)
	}
}

func TestCheckRejectsConflictingAPIKeys(t *testing.T) {
	for name, a := range map[string]API{
		"both id owners":           {ClientSetID: true, ServerSetID: true},
		"singleton id and id path": {SingletonID: true, ServerSetID: true},
		"delete with no name":      {Delete: &Delete{}},
		"incomplete operations":    {Operations: &Operations{Create: "C"}},
		"two delete operations":    {Operations: &Operations{Create: "C", Get: "G", Update: "U", Delete: "D"}, Delete: &Delete{Operation: "Archive"}},
	} {
		if a.Check() == nil {
			t.Errorf("%s: Check() = nil", name)
		}
	}
	if err := (API{Delete: &Delete{Operation: "Archive"}}).Check(); err != nil {
		t.Fatal(err)
	}
	if CheckUnwrap([]string{"A", "A"}) == nil || CheckUnwrap([]string{""}) == nil || CheckUnwrap([]string{"A", "B"}) != nil {
		t.Fatal("CheckUnwrap does not reject an empty or a repeated name")
	}
}

// A tool decodes its whole file once. Its field line embeds Field inline, so the API keys and
// its own keys share one line, and a key that neither part knows is still an error.
func TestToolEmbedsFieldInItsLine(t *testing.T) {
	type toolField struct {
		Field       `yaml:",inline"`
		Description *string `yaml:"description"`
	}
	type toolFile struct {
		API   API                             `yaml:"api"`
		Types map[string]map[string]toolField `yaml:"types"`
	}
	decode := func(text string) (toolFile, error) {
		dec := yaml.NewDecoder(bytes.NewReader([]byte(text)))
		dec.KnownFields(true)
		var f toolFile
		return f, dec.Decode(&f)
	}
	f, err := decode("api: {requestWrapper: router}\ntypes:\n  Target:\n    id: {skip: true, description: The id.}\n")
	if err != nil {
		t.Fatal(err)
	}
	if line := f.Types["Target"]["id"]; !line.Skip || line.Description == nil || *line.Description != "The id." || f.API.RequestWrapper != "router" {
		t.Fatalf("decoded %+v", f)
	}
	if _, err := decode("types:\n  Target:\n    id: {skipp: true}\n"); err == nil || !strings.Contains(err.Error(), "skipp") {
		t.Fatalf("unknown key error = %v", err)
	}
}

func TestContractReadOnly(t *testing.T) {
	doc, err := model.Load([]byte(`openapi: 3.1.0
info: {title: t, version: "1"}
paths: {}
components:
  schemas:
    Thing:
      type: object
      properties:
        id: {type: string, readOnly: true}
        name: {type: string}
`))
	if err != nil {
		t.Fatal(err)
	}
	schema, err := doc.Components.Schemas.GetOrZero("Thing").BuildSchema()
	if err != nil {
		t.Fatal(err)
	}
	if !ContractReadOnly(schema, "id") || ContractReadOnly(schema, "name") {
		t.Fatal("ContractReadOnly does not match the readOnly keyword")
	}
}
