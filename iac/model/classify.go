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

import "fmt"

// classifyTopLevel returns the behavior of a top-level field. With updateIDInBody,
// the Update call has no id in its path, so a non-singleton's Get path id is a
// field of the Update body. A singleton can also have a server-assigned id in
// its Update body. When Create does not send the id, it is computed and Update
// sends the value from state.
//
// A field that Create does not have, but that Update can set and that declares
// a server default in Update (updateDefault), is managed: Create stores the
// default, and the generated Create follows with an Update when the
// configuration sets another value.
// id is the id property: the id path parameter of Get, or the property that api.idProperty names.
func classifyTopLevel(p Policy, id, field string, inCreate, inUpdate, inGet, updateDefault bool) (Behavior, error) {
	// A singleton can return a server-assigned id and require it in the body of
	// a later Replace, even though it has no path id parameter.
	isID := id != "" && field == id || id == "" && field == "id"
	if p.UpdateIDInBody && isID && !inCreate && inUpdate && inGet {
		return Computed, nil
	}
	if !inCreate && inUpdate && inGet && updateDefault {
		return Normal, nil
	}
	return Classify(inCreate, inUpdate, inGet)
}

// Classify returns the behavior of a top-level field from where it appears
// (README.md, "Field location"). Any other combination is an error.
func Classify(inCreate, inUpdate, inGet bool) (Behavior, error) {
	switch {
	case inCreate && inUpdate && inGet:
		return Normal, nil
	case inCreate && !inUpdate && inGet:
		return Immutable, nil
	case !inCreate && !inUpdate && inGet:
		return Computed, nil
	}
	return "", fmt.Errorf("unsupported field location: create %t, update %t, get %t", inCreate, inUpdate, inGet)
}
