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
	"fmt"
	"slices"
	"strings"

	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

// discoverOperations finds the operations of each lifecycle step. It searches for the operationId
// suffix _<Role><Name> first. When that search does not resolve the resource, it adds a service
// fallback for each step that has no match and no explicit operationId (see serviceOperations).
//
// The suffix search resolves the resource when every step has exactly one match, or when the
// matches make an always-exists singleton. The fallback never runs for such a resource, so it
// cannot change a result that the suffix search already gives.
//
// hints holds, for a step that stays unresolved, a sentence that tells what the fallback found.
func discoverOperations(doc *v3.Document, name string, explicit map[verb]string, p Policy) (candidates map[verb][]foundOp, hints map[verb]string) {
	candidates = map[verb][]foundOp{}
	for _, role := range verbs {
		candidates[role] = matchingOperations(doc, name, role, explicit[role])
	}
	if suffixResolved(candidates, explicit, p) {
		return candidates, nil
	}
	hints = map[verb]string{}
	for _, role := range verbs {
		if len(candidates[role]) != 0 || explicit[role] != "" {
			continue
		}
		matches := serviceOperations(doc, name, role, p)
		if len(matches) == 1 {
			candidates[role] = matches
			continue
		}
		hints[role] = fallbackHint(name, role, matches)
	}
	return candidates, hints
}

func suffixResolved(candidates map[verb][]foundOp, explicit map[verb]string, p Policy) bool {
	if impliedCandidates(candidates, explicit, p) {
		return true
	}
	for _, role := range verbs {
		if len(candidates[role]) != 1 {
			return false
		}
	}
	return true
}

// serviceOperations returns the operations of a service whose operationId is
// <Service>_<Role><anything> and whose success response is the component name. The service is the
// operationId prefix before the first "_". Update also takes the prefix Replace, as the suffix
// search does.
func serviceOperations(doc *v3.Document, name string, role verb, p Policy) []foundOp {
	prefixes := []string{string(role)}
	if role == opUpdate {
		prefixes = append(prefixes, string(opReplace))
	}
	var matches []foundOp
	for path, item := range doc.Paths.PathItems.FromOldest() {
		for method, operation := range item.GetOperations().FromOldest() {
			_, action, ok := strings.Cut(operation.OperationId, "_")
			if !ok || !slices.ContainsFunc(prefixes, func(prefix string) bool { return strings.HasPrefix(action, prefix) }) {
				continue
			}
			if p.responseResourceProxy(operation, name) == nil {
				continue
			}
			matches = append(matches, foundOp{path: path, method: strings.ToUpper(method), item: item, op: operation})
		}
	}
	return matches
}

func fallbackHint(name string, role verb, matches []foundOp) string {
	base := fmt.Sprintf("Discovery also looked for <Service>_%s<anything> operations that return %s", role, name)
	if len(matches) == 0 {
		return base + ", and found none."
	}
	var ids []string
	for _, m := range matches {
		ids = append(ids, m.op.OperationId)
	}
	slices.Sort(ids)
	return fmt.Sprintf("%s, and found %d: %s. Select one exact operation ID.", base, len(ids), strings.Join(ids, ", "))
}

// hintSentence is the hint as a sentence to add after an error message, or "".
func hintSentence(hint string) string {
	if hint == "" {
		return ""
	}
	return ". " + hint
}
