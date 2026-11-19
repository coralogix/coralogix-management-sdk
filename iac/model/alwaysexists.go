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

// impliedCandidates reports whether the operations that match each step make an always-exists
// singleton: one Get and one Update on one path with no path parameters, and no Create and no
// Delete. A policy that changes how the resource is read or deleted turns the rule off.
func impliedCandidates(candidates map[verb][]foundOp, p Policy) bool {
	if p.GetFromList || p.Overwrite || p.GetByIDs != "" || p.DeleteOperation != "" || p.DeleteIDs != "" || p.UpdateIDInBody {
		return false
	}
	if len(candidates[opCreate]) != 0 || len(candidates[opDelete]) != 0 {
		return false
	}
	if len(candidates[opGet]) != 1 || len(candidates[opUpdate]) != 1 {
		return false
	}
	get, update := candidates[opGet][0], candidates[opUpdate][0]
	return len(pathParams(get)) == 0 && len(pathParams(update)) == 0 && get.path == update.path
}

// impliedLifecycle adds the Create and the Delete of an always-exists singleton to found. The API
// cannot create or remove such a resource, so Create is the Update call, and Delete is the Get
// call, which no tool makes. found must hold Get and Update and no Create and no Delete.
func impliedLifecycle(found map[verb]foundOp, p Policy) {
	candidates := map[verb][]foundOp{}
	for v, f := range found {
		candidates[v] = []foundOp{f}
	}
	if !impliedCandidates(candidates, p) {
		return
	}
	create, del := found[opUpdate], found[opGet]
	create.implied, del.implied = true, true
	found[opCreate], found[opDelete] = create, del
}
