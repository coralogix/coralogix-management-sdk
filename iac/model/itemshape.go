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
	"maps"
	"slices"
	"strings"

	"github.com/coralogix/coralogix-management-sdk/iac/issue"
)

// shapeNested returns req, the type of a value in a request, with each nested object that the
// policy shapes rewritten into the shape of its type in the resource response, get: the fields of
// a RequestWrappers property move up beside the other fields, a nested requestAs line renames a
// field, and RequestNegated marks a field. Each moved field records its RequestPath and
// RequestObjects, so the merge sees the shape of the response and a tool can still build the
// request. used collects the "Component.field" keys and the wrapped components that took effect.
// Without such lines it returns req unchanged.
func (p Policy) shapeNested(resource, location string, get, req *Type, used map[string]bool) (*Type, error) {
	if get == nil || req == nil || !p.shapesNested(resource) {
		return req, nil
	}
	switch get.Kind {
	case List, Set, Map:
		elem, err := p.shapeNested(resource, location+"[]", get.Elem, req.Elem, used)
		if err != nil || elem == req.Elem {
			return req, err
		}
		out := *req
		out.Elem = elem
		return &out, nil
	case Object, OneOf:
	default:
		return req, nil
	}
	out := req
	if p.shapesComponent(resource, get.Schema) {
		var err error
		if out, err = p.shapeObject(location, get, req, used); err != nil {
			return nil, err
		}
	}
	var fields []*Field
	for i, rf := range out.Fields {
		gf := fieldNamed(get.Fields, rf.Name)
		if gf == nil {
			continue
		}
		shaped, err := p.shapeNested(resource, location+"."+rf.Name, gf.Type, rf.Type, used)
		if err != nil {
			return nil, err
		}
		if shaped == rf.Type {
			continue
		}
		if fields == nil {
			fields = slices.Clone(out.Fields)
		}
		copied := *rf
		copied.Type = shaped
		fields[i] = &copied
	}
	if fields == nil {
		return out, nil
	}
	copied := *out
	copied.Fields = fields
	return &copied, nil
}

// shapesNested reports whether the policy shapes any nested request object of the resource.
func (p Policy) shapesNested(resource string) bool {
	return len(p.RequestWrappers) != 0 || len(p.RequestNegated) != 0 || len(p.nestedRequestNames(resource)) != 0
}

// shapesComponent reports whether the policy shapes the request objects of the response component.
func (p Policy) shapesComponent(resource, component string) bool {
	if component == "" || component == resource {
		return false
	}
	if _, ok := p.RequestWrappers[component]; ok {
		return true
	}
	prefix := component + "."
	for key := range p.nestedRequestNames(resource) {
		if strings.HasPrefix(key, prefix) {
			return true
		}
	}
	return slices.ContainsFunc(p.RequestNegated, func(key string) bool { return strings.HasPrefix(key, prefix) })
}

// nestedRequestNames are the requestAs lines of fields of components other than the top level of
// the resource. The top-level lines are renamed on the request bodies (shapeRequests); a nested
// line is renamed here, on the type of the request object.
func (p Policy) nestedRequestNames(resource string) map[string]string {
	out := map[string]string{}
	for key, path := range p.RequestNames {
		if component, _, _ := strings.Cut(key, "."); component != resource {
			out[key] = path
		}
	}
	return out
}

// shapeObject rewrites req, a request object of the response component get.Schema, into the shape
// of get.
func (p Policy) shapeObject(location string, get, req *Type, used map[string]bool) (*Type, error) {
	component := get.Schema
	wrapper := p.RequestWrappers[component]
	fields := make([]*Field, 0, len(req.Fields))
	var inner *Field
	for _, rf := range req.Fields {
		if wrapper != "" && rf.Name == wrapper {
			inner = rf
			continue
		}
		copied := *rf
		fields = append(fields, &copied)
	}
	if wrapper != "" {
		if inner == nil || inner.Type == nil || inner.Type.Kind != Object {
			return nil, &contractError{
				code:        "REQUEST_WRAPPER_NOT_FOUND",
				location:    location,
				message:     fmt.Sprintf("The request object %s has no object property %q, which the overrides name as the wrapper of %s.", req.Schema, wrapper, component),
				remediation: "Name the property of the request object that holds the fields of " + component + ", or delete the line.",
			}
		}
		used["requestWrapper."+component] = true
		for _, wf := range inner.Type.Fields {
			if fieldNamed(fields, wf.Name) != nil {
				return nil, &contractError{
					code:        "REQUEST_NAME_CONFLICT",
					location:    location + "." + wf.Name,
					message:     fmt.Sprintf("The request object %s has %q both beside and inside its wrapper %q.", req.Schema, wf.Name, wrapper),
					remediation: "Give the fields of the wrapper and of the request object different names.",
				}
			}
			copied := *wf
			copied.RequestPath = []string{wrapper, wf.Name}
			copied.RequestObjects = []string{inner.Type.Schema}
			if inner.Attrs.Required {
				// The wrapper is required, so a field that it requires is required in the request.
				copied.Attrs.Required = wf.Attrs.Required
			} else {
				copied.Attrs.Required = false
			}
			fields = append(fields, &copied)
		}
	}
	prefix := component + "."
	for _, key := range slices.Sorted(maps.Keys(p.RequestNames)) {
		if !strings.HasPrefix(key, prefix) {
			continue
		}
		field := strings.TrimPrefix(key, prefix)
		path := strings.Split(p.RequestNames[key], ".")
		if wrapper != "" {
			path = append([]string{wrapper}, path...)
		}
		index := slices.IndexFunc(fields, func(f *Field) bool {
			return slices.Equal(requestPathOf(f), path)
		})
		if index < 0 {
			return nil, &contractError{
				code:        "REQUEST_NAME_NOT_FOUND",
				location:    location + "." + field,
				message:     fmt.Sprintf("The request object %s has no property at %q for %s.", req.Schema, strings.Join(path, "."), key),
				remediation: "Name the path of the property after the wrapper, or delete the line.",
			}
		}
		if other := fieldNamed(fields, field); other != nil && other != fields[index] {
			return nil, &contractError{
				code:        "REQUEST_NAME_CONFLICT",
				location:    location + "." + field,
				message:     fmt.Sprintf("The request object %s already has %q, which the requestAs line of %s also maps.", req.Schema, field, key),
				remediation: "Map each request property to one field.",
			}
		}
		used[key] = true
		fields[index].RequestPath = path
		fields[index].Name = field
	}
	for _, key := range p.RequestNegated {
		field, ok := strings.CutPrefix(key, prefix)
		if !ok {
			continue
		}
		f := fieldNamed(fields, field)
		if f == nil || f.Type == nil || f.Type.Kind != Bool {
			return nil, &contractError{
				code:        "REQUEST_NEGATED_UNSUPPORTED",
				location:    location + "." + field,
				message:     fmt.Sprintf("The overrides negate %s, but the request object %s has no boolean field of that name.", key, req.Schema),
				remediation: "Negate a boolean field, with a requestAs line when the request names it differently.",
			}
		}
		used["negate."+key] = true
		f.Negated = true
	}
	out := *req
	out.Fields = fields
	// A write that sends one arm of a oneOf of the response as a plain field always selects that
	// arm. The other arms are not in the request, so they are computed.
	out.Groups = nil
	for _, g := range get.Groups {
		present := slices.ContainsFunc(g.Arms, func(arm string) bool { return fieldNamed(fields, arm) != nil })
		out.Groups = append(out.Groups, OneOfGroup{Arms: slices.Clone(g.Arms), AllowNone: g.AllowNone && !present})
	}
	return &out, nil
}

// requestPathOf is the request path of f: RequestPath, or its name.
func requestPathOf(f *Field) []string {
	if f.RequestPath != nil {
		return f.RequestPath
	}
	return []string{f.Name}
}

// unusedNestedShapes reports the nested lines that shaped no request object of the resource: a
// requestAs line on a component other than the resource, a requestWrapper line, or a negation.
func (p Policy) unusedNestedShapes(resource string, used map[string]bool) error {
	var keys []string
	for key := range p.nestedRequestNames(resource) {
		if !used[key] {
			keys = append(keys, "types."+key+".requestAs")
		}
	}
	for component := range p.RequestWrappers {
		if !used["requestWrapper."+component] {
			keys = append(keys, "types."+component+".requestWrapper")
		}
	}
	for _, key := range p.RequestNegated {
		if !used["negate."+key] {
			keys = append(keys, "types."+key+".negate")
		}
	}
	if len(keys) == 0 {
		return nil
	}
	slices.Sort(keys)
	return &contractError{
		code:        "REQUEST_NAME_UNKNOWN_FIELD",
		location:    "behavior-overrides.yaml:" + keys[0],
		message:     fmt.Sprintf("These lines shape no request object of %s: %s.", resource, strings.Join(keys, ", ")),
		remediation: "Put the line on a component that a request of the resource holds, or delete it.",
	}
}

// nestedShapeIssues reports a nested request line on a component that the resource does not hold.
// Build reports a line on a component that it holds but whose requests the line does not match.
func nestedShapeIssues(p Policy, resource string, root *Type) issue.Report {
	held := map[string]bool{}
	var walkType func(t *Type)
	walkType = func(t *Type) {
		if t == nil || (t.Schema != "" && held[t.Schema]) {
			return
		}
		if t.Schema != "" {
			held[t.Schema] = true
		}
		walkType(t.Elem)
		for _, f := range t.Fields {
			walkType(f.Type)
		}
	}
	walkType(root)
	components := map[string]string{}
	for key := range p.nestedRequestNames(resource) {
		component, _, _ := strings.Cut(key, ".")
		components[component] = "types." + key + ".requestAs"
	}
	for component := range p.RequestWrappers {
		components[component] = "types." + component + ".requestWrapper"
	}
	for _, key := range p.RequestNegated {
		component, _, _ := strings.Cut(key, ".")
		if component != resource {
			components[component] = "types." + key + ".negate"
		}
	}
	var report issue.Report
	for _, component := range slices.Sorted(maps.Keys(components)) {
		if held[component] {
			continue
		}
		report = append(report, issue.Issue{
			Code:        "REQUEST_NAME_UNKNOWN_FIELD",
			Location:    "behavior-overrides.yaml:" + components[component],
			Message:     fmt.Sprintf("The line shapes the requests of %s, which %s does not hold.", component, resource),
			Remediation: "Put the line on a component of the resource, or delete it.",
		})
	}
	return report
}
