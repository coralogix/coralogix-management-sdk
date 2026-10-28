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
	"github.com/pb33f/libopenapi/datamodel/high/base"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/orderedmap"
)

// updateIDOutsideWrapper reports whether the Update body carries the id beside the request wrapper.
func (p Policy) updateIDOutsideWrapper() bool {
	return p.UpdateIDOutsideWrapper && p.RequestWrapper != "" && p.UpdateIDInBody
}

// shapeRequests returns the Create and Update bodies as the resource sees them, when the policy
// says that a request is shaped differently from the resource component: requestAs moves a
// property of a request to the top level under the name that the resource has, and
// updateIDOutsideWrapper adds the id property that sits beside the wrapper to the Update body.
// Validation and Build both read the bodies this way, so the field has the normal lifecycle and
// its type is checked as usual. Without these keys it returns the bodies unchanged.
func (p Policy) shapeRequests(name, id string, ops map[verb]foundOp, create, update, get *base.Schema) (c, u *base.Schema, report issue.Report) {
	if len(p.RequestNames) == 0 && !p.UpdateIDOutsideWrapper {
		return create, update, nil
	}
	c, u = create, update
	if p.UpdateIDOutsideWrapper {
		u, report = p.withOutsideID(ops, id, update)
	}
	if len(p.RequestNames) != 0 {
		var names issue.Report
		c, u, names = p.renameRequestFields(name, c, u, get)
		report = append(report, names...)
	}
	return c, u, report
}

// withOutsideID adds the id property that sits beside the wrapper of the Update body to update, the
// schema inside the wrapper.
func (p Policy) withOutsideID(ops map[verb]foundOp, id string, update *base.Schema) (*base.Schema, issue.Report) {
	location := "paths.update." + ops[opUpdate].op.OperationId + ".requestBody"
	if !p.UpdateIDInBody || p.RequestWrapper == "" {
		return update, issue.Report{{
			Code:        "UPDATE_ID_OUTSIDE_WRAPPER_UNSUPPORTED",
			Location:    location,
			Message:     "The override updateIDOutsideWrapper needs requestWrapper and updateIDInBody.",
			Remediation: "Set requestWrapper and updateIDInBody, or delete updateIDOutsideWrapper.",
		}}
	}
	outer, _, required := validationObject(bodySchemaOrNil(bodyProxy(ops[opUpdate].op)))
	idProperty := propertyOf(outer, id)
	switch {
	case id == "":
		return update, nil // Build reports the unsupported identity shape.
	case idProperty == nil:
		return update, issue.Report{{
			Code:        "UPDATE_ID_NOT_IN_BODY",
			Location:    location,
			Message:     fmt.Sprintf("The override updateIDOutsideWrapper says that the Update body has the id %q beside %q, but it has no such property.", id, p.RequestWrapper),
			Remediation: "Delete updateIDOutsideWrapper from the overrides, or add the id property to the Update body.",
		}}
	case propertyOf(update, id) != nil:
		return update, issue.Report{{
			Code:        "UPDATE_ID_INSIDE_WRAPPER",
			Location:    location,
			Message:     fmt.Sprintf("The override updateIDOutsideWrapper says that the id %q is beside %q, but %q has a property of that name too.", id, p.RequestWrapper, p.RequestWrapper),
			Remediation: "Delete updateIDOutsideWrapper from the overrides if the id is inside the wrapper.",
		}}
	}
	out := copySchema(update, "", update.Required)
	out.Properties.Set(id, idProperty)
	if slices.Contains(required, id) {
		out.Required = append(out.Required, id)
	}
	return out, nil
}

// bodySchemaOrNil returns the schema of proxy, or nil when it has none: callers read a missing body as
// a body without properties, and the operation checks report it.
func bodySchemaOrNil(proxy *base.SchemaProxy) *base.Schema {
	s, _ := schemaOf(referencedSchema(proxy))
	return s
}

// renameRequestFields moves each property that a requestAs line names to the top level of the Create and
// Update bodies under the name of the resource field.
func (p Policy) renameRequestFields(name string, create, update, get *base.Schema) (c, u *base.Schema, report issue.Report) {
	names := p.requestNames(name)
	fields := make([]string, 0, len(names))
	for field := range names {
		fields = append(fields, field)
	}
	slices.Sort(fields)
	for _, key := range slices.Sorted(maps.Keys(p.RequestNames)) {
		if component, _, _ := strings.Cut(key, "."); component != name {
			report = append(report, issue.Issue{Code: "REQUEST_NAME_UNKNOWN_FIELD", Location: "behavior-overrides.yaml:types." + key, Message: fmt.Sprintf("The field line %q has requestAs, but the resource component is %s.", key, name), Remediation: "Put requestAs on a field of the resource component, or delete it."})
		}
	}
	var valid []string // the fields whose line can be applied
	byPath := map[string]string{}
	for _, field := range fields {
		path := names[field]
		location := fieldLocation(name, field)
		if propertyOf(get, field) == nil {
			report = append(report, issue.Issue{Code: "REQUEST_NAME_UNKNOWN_FIELD", Location: location, Message: fmt.Sprintf("The overrides map the request path %q to %q, which is not a property of %s.", path, field, name), Remediation: "Name a property of the resource component, or delete the line."})
			continue
		}
		if path == field {
			report = append(report, issue.Issue{Code: "OVERRIDE_UNUSED", Location: "behavior-overrides.yaml:types." + name + "." + field, Message: fmt.Sprintf("the requests already name the field %q like the resource", field), Remediation: "Delete the line."})
			continue
		}
		if other, clash := overlappingPath(byPath, path); clash {
			report = append(report, issue.Issue{Code: "REQUEST_NAME_DUPLICATE", Location: location, Message: fmt.Sprintf("The request path %q is also used by %q (the paths are the same, or one holds the other).", path, other), Remediation: "Map each request property to one resource field."})
			continue
		}
		byPath[path] = field
		valid = append(valid, field)
	}
	c, u = create, update
	var found [2]map[string]bool
	for i, body := range []*base.Schema{create, update} {
		var shaped *base.Schema
		shaped, found[i], report = p.renameIn(name, body, valid, report)
		if i == 0 {
			c = shaped
		} else {
			u = shaped
		}
	}
	for _, field := range valid {
		if !found[0][field] && !found[1][field] {
			report = append(report, issue.Issue{Code: "REQUEST_NAME_NOT_FOUND", Location: fieldLocation(name, field), Message: fmt.Sprintf("Neither the Create nor the Update body has the property %q that the overrides map to %q.", names[field], field), Remediation: "Fix the path in requestAs, or delete the line."})
		}
	}
	return c, u, report
}

// overlappingPath returns the field of a path in byPath that is path, or a prefix of it, or
// starts with it.
func overlappingPath(byPath map[string]string, path string) (string, bool) {
	for other, field := range byPath {
		if other == path || strings.HasPrefix(other, path+".") || strings.HasPrefix(path, other+".") {
			return field, true
		}
	}
	return "", false
}

// renameIn moves the properties of the fields that it finds in body to the top level. found says
// which fields the body has.
func (p Policy) renameIn(name string, body *base.Schema, fields []string, report issue.Report) (*base.Schema, map[string]bool, issue.Report) {
	found := map[string]bool{}
	if body == nil {
		return body, found, report
	}
	names := p.requestNames(name)
	type move struct {
		field    string
		leaf     *base.SchemaProxy
		required bool
	}
	var moves []move
	for _, field := range fields {
		leaf, required, problem := resolveRequestPath(body, strings.Split(names[field], "."))
		switch problem {
		case "":
			found[field] = true
			moves = append(moves, move{field, leaf, required})
		case problemNotObject:
			report = append(report, issue.Issue{Code: "REQUEST_NAME_NOT_OBJECT", Location: fieldLocation(name, field), Message: fmt.Sprintf("A segment of the request path %q is not an object, so the path cannot go through it.", names[field]), Remediation: "Dot-separate only object properties in requestAs."})
		}
	}
	if len(moves) == 0 {
		return body, found, report
	}
	obj, _, required := validationObject(body)
	shaped := obj
	for _, m := range moves {
		shaped = removePath(shaped, required, strings.Split(names[m.field], "."))
		required = shaped.Required
	}
	shaped = copySchema(shaped, "", required)
	for _, m := range moves {
		if propertyOf(shaped, m.field) != nil {
			report = append(report, issue.Issue{Code: "REQUEST_NAME_CONFLICT", Location: fieldLocation(name, m.field), Message: fmt.Sprintf("A request body has its own property %q, and the overrides also map %q to it.", m.field, names[m.field]), Remediation: "Map the request property to a field that no request property has, or delete the line."})
			delete(found, m.field)
			continue
		}
		shaped.Properties.Set(m.field, m.leaf)
		if m.required {
			shaped.Required = append(shaped.Required, m.field)
		}
	}
	return shaped, found, report
}

const (
	problemNotFound  = "not found"
	problemNotObject = "not object"
)

// resolveRequestPath returns the property at path in body, and whether every property on the way is
// required. problem says why there is none: the path is not there, or it goes through a value that
// is not an object.
func resolveRequestPath(body *base.Schema, path []string) (leaf *base.SchemaProxy, required bool, problem string) {
	cur, required := body, true
	for i, segment := range path {
		obj, _, req := validationObject(cur)
		proxy := propertyOf(obj, segment)
		if proxy == nil {
			return nil, false, problemNotFound
		}
		required = required && slices.Contains(req, segment)
		if i == len(path)-1 {
			return proxy, required, ""
		}
		s, err := schemaOf(proxy)
		if err != nil {
			return nil, false, problemNotFound
		}
		if inner, _, _ := validationObject(s); inner == nil || inner.Properties == nil {
			return nil, false, problemNotObject
		}
		cur = s
	}
	return nil, false, problemNotFound
}

// removePath returns a copy of obj, a resolved object whose required list is required, without the
// property at path. An object that loses its last property goes too.
func removePath(obj *base.Schema, required, path []string) *base.Schema {
	if len(path) > 1 {
		if child, err := schemaOf(propertyOf(obj, path[0])); err == nil {
			inner, _, innerRequired := validationObject(child)
			rest := removePath(inner, innerRequired, path[1:])
			if rest.Properties.Len() != 0 {
				out := copySchema(obj, "", required)
				out.Properties.Set(path[0], base.CreateSchemaProxy(rest))
				return out
			}
		}
	}
	return copySchema(obj, path[0], slices.DeleteFunc(slices.Clone(required), func(n string) bool { return n == path[0] }))
}

// copySchema returns a copy of obj with its own properties, without the property skip, and with
// the required list given.
func copySchema(obj *base.Schema, skip string, required []string) *base.Schema {
	out := *obj
	out.Properties = copyProperties(obj, skip)
	out.Required = slices.Clone(required)
	return &out
}

// copyProperties returns the properties of obj in order, without skip.
func copyProperties(obj *base.Schema, skip string) *orderedmap.Map[string, *base.SchemaProxy] {
	out := orderedmap.New[string, *base.SchemaProxy]()
	if obj != nil && obj.Properties != nil {
		for name, proxy := range obj.Properties.FromOldest() {
			if name != skip {
				out.Set(name, proxy)
			}
		}
	}
	return out
}

// readRequestComponents records the components of the request wrapper property of the Create and
// Update bodies, and the objects that each request path goes through, so a tool can name the types
// that the SDK body holds. Without a request wrapper it records nothing.
func (r *Resource) readRequestComponents(ops map[verb]foundOp) {
	if r.Policy.RequestWrapper == "" {
		return
	}
	for v, o := range map[verb]*Operation{opCreate: &r.Create, opUpdate: &r.Update} {
		if op := ops[v].op; op != nil {
			o.RequestComponent = r.Policy.wrapperComponent(op)
		}
	}
	for _, f := range r.Fields {
		if f.RequestPath == "" || strings.Count(f.RequestPath, ".") == 0 {
			continue
		}
		path := strings.Split(f.RequestPath, ".")
		for _, v := range []verb{opCreate, opUpdate} {
			if op := ops[v].op; op != nil {
				if objects := r.Policy.requestObjects(op, path); objects != nil {
					f.RequestObjects = objects
					break
				}
			}
		}
	}
}

// wrapperComponent returns the component of the request wrapper property of op, or "".
func (p Policy) wrapperComponent(op *v3.Operation) string {
	body := bodyProxy(op)
	if body == nil {
		return ""
	}
	s, err := schemaOf(body)
	if err != nil {
		return ""
	}
	name, err := componentName(unwrapRef(propertyOf(s, p.RequestWrapper)))
	if err != nil {
		return ""
	}
	return name
}

// requestObjects returns the component of each object that path goes through in the request body
// of op, one for each segment before the last. It returns nil when a segment has no component.
func (p Policy) requestObjects(op *v3.Operation, path []string) []string {
	cur, err := schemaOf(p.requestBody(op))
	if err != nil {
		return nil
	}
	var out []string
	for _, segment := range path[:len(path)-1] {
		obj, _, _ := validationObject(cur)
		proxy := propertyOf(obj, segment)
		name, err := componentName(unwrapRef(proxy))
		if err != nil {
			return nil
		}
		out = append(out, name)
		if cur, err = schemaOf(proxy); err != nil {
			return nil
		}
	}
	return out
}

// requestNames returns the request path of each field of the resource component that a field
// line gives a requestAs. RequestNames is keyed by component and field, like the other field
// lines.
func (p Policy) requestNames(resource string) map[string]string {
	out := map[string]string{}
	for key, path := range p.RequestNames {
		if component, field, ok := strings.Cut(key, "."); ok && component == resource {
			out[field] = path
		}
	}
	return out
}
