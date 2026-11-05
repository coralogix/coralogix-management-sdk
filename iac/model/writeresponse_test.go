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
	"testing"
)

// Every object component of the real contract that the model reads holds itself, so a write that
// returns it can give the state. The one exception is a component with an inline object that has
// fields: the SDK names that object after the component, so another component cannot hold it.
func TestWriteResponseAcceptsEveryRealComponent(t *testing.T) {
	raw, err := os.ReadFile("../../openapi.yaml")
	if err != nil {
		t.Fatal(err)
	}
	doc, err := Load(raw)
	if err != nil {
		t.Fatal(err)
	}
	policy := Policy{Existing: true}
	checked := 0
	for name := range doc.Components.Schemas.KeysFromOldest() {
		resource, err := typeOf(doc.Components.Schemas.GetOrZero(name), name, walk{policy: &policy})
		if err != nil || resource.Kind != Object && resource.Kind != OneOf {
			continue // not a resource the model can read
		}
		checked++
		if problem := writeResponseDifference(doc, &policy, name, name); problem != "" && !hasInlineObjectWithFields(resource) {
			t.Errorf("%s does not hold itself: %s", name, problem)
		}
	}
	if checked < 100 {
		t.Fatalf("checked %d components; the contract should have hundreds", checked)
	}
}

// hasInlineObjectWithFields reports whether a field of t, at any depth outside other components,
// is an inline object with fields.
func hasInlineObjectWithFields(t *Type) bool {
	for _, f := range t.Fields {
		for ft := f.Type; ft != nil; ft = ft.Elem {
			if ft.Kind == Object && ft.Schema == "" && len(ft.Fields) != 0 {
				return true
			}
		}
	}
	return false
}
