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
	"testing"
)

// actionPolicy is the policy of a released resource like the actions: the Update body wraps the
// resource and has its id, and the Create body is the resource itself.
func actionPolicy() Policy {
	return Policy{
		Existing:       true,
		RequestWrapper: "thing",
		UpdateIDInBody: true,
		ServerSetID:    true,
		ReadOnly:       []string{"ActionThing.id"},
	}
}

func loadActionThing(t *testing.T) *Resource {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", "actionthing.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := Load(data)
	if err != nil {
		t.Fatal(err)
	}
	r, err := BuildWithPolicy(doc, "ActionThing", OperationIDs{}, actionPolicy())
	if err != nil {
		t.Fatal(err)
	}
	return r
}

// A Create body without the wrapper property, next to an Update body that has it, is read as the
// resource itself. No file states it: the contract does.
func TestCreateWithoutTheWrapperIsUnwrapped(t *testing.T) {
	r := loadActionThing(t)
	if !r.Policy.CreateUnwrapped || r.Policy.WrapperOf(true) != "" || r.Policy.WrapperOf(false) != "thing" {
		t.Fatalf("policy = %+v, want only the Update body wrapped", r.Policy)
	}
	fields := map[string]*ResourceField{}
	for _, f := range r.Fields {
		fields[f.Name] = f
	}
	if f := fields["name"]; f == nil || f.Create == nil || f.Update == nil {
		t.Fatalf("name = %+v, want a field in Create (flat body) and Update (wrapped)", f)
	}
	// The contract declares the default of hidden in Update, so Create stores it.
	if f := fields["hidden"]; f == nil || f.Create != nil || f.Update == nil || f.Behavior != Normal {
		t.Fatalf("hidden = %+v, want a managed field that only Update has", f)
	}
}

// A Create body that has the wrapper too is wrapped.
func TestCreateWithTheWrapperIsNotUnwrapped(t *testing.T) {
	r, err := buildFolderThing(t, folderPolicy())
	if err != nil {
		t.Fatal(err)
	}
	if r.Policy.CreateUnwrapped || r.Policy.WrapperOf(true) != "thing" {
		t.Fatalf("policy = %+v, want both bodies wrapped", r.Policy)
	}
}

// With no wrapper in the policy, nothing is unwrapped.
func TestNoRequestWrapperMeansNoUnwrappedCreate(t *testing.T) {
	p := Policy{Existing: true}
	if p.WrapperOf(true) != "" || p.WrapperOf(false) != "" {
		t.Fatalf("WrapperOf = %q, %q, want none", p.WrapperOf(true), p.WrapperOf(false))
	}
}
