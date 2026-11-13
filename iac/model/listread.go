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
	"errors"
	"fmt"
	"maps"
	"slices"

	"github.com/pb33f/libopenapi/datamodel/high/base"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

// getFromListLocation is the line of the behavior-overrides file that says Get reads the List.
const getFromListLocation = "behavior-overrides.yaml:api.getFromList"

// readListOperations reads a resource that no operation reads by its id: Get is the List, and a tool
// selects the item whose id is the resource id. Create and Get have no path parameters, Update sends
// the id in its body on the Create path, and Delete takes the id in its path, or in the id array of
// its body (Policy.DeleteIDs).
func (r *Resource) readListOperations(ops map[verb]foundOp) error {
	switch {
	case !r.Policy.Existing:
		return errors.New("api.getFromList is only for existing resources")
	case r.Policy.GetByIDs != "":
		return errors.New("api.getFromList and api.getByIDs both say how Get reads the resource; keep one")
	case r.Policy.IDProperty == "":
		return &contractError{code: "GET_FROM_LIST_ID_PROPERTY_REQUIRED", location: getFromListLocation,
			message:     "No operation has the id in a path parameter of Get, so the overrides must name the id property.",
			remediation: "Set api.idProperty to the property of the resource that holds its id."}
	case !r.Policy.UpdateIDInBody:
		return errors.New("api.getFromList needs api.updateIDInBody: Get has no id path parameter for Update to share")
	}
	if len(pathParams(ops[opCreate])) != 0 {
		return errors.New("create must have no path parameters")
	}
	r.IDField = r.Policy.IDProperty
	get := ops[opGet]
	list, problem := r.readListOperation(get, get)
	if problem != nil {
		return &contractError{code: "GET_FROM_LIST_NOT_LIST", location: problem.location,
			message:     fmt.Sprintf("Get %s is not a List that the model can describe: %s", get.op.OperationId, problem.message),
			remediation: problem.remediation}
	}
	if !list.Full {
		return &contractError{code: "GET_FROM_LIST_NOT_FULL", location: "components.schemas." + list.ItemComponent,
			message:     fmt.Sprintf("The List item %s is not %s, so the List cannot read every field.", list.ItemComponent, r.Name),
			remediation: "Name a List whose items are the resource component, or delete api.getFromList."}
	}
	if _, err := r.itemOperations(ops); err != nil {
		return err
	}
	if err := r.readDeleteOperation(ops); err != nil {
		return err
	}
	if err := r.readTargets(ops); err != nil {
		return err
	}
	r.GetFromList, r.List = true, list
	if err := r.readPaging(get); err != nil {
		return err
	}
	return r.readListQuery(get)
}

// readListQuery reads the query parameters that the List read of Get always sends (Policy.GetQuery),
// for example the status of the resources that a remove has not deactivated. A tool pages the List
// and narrows it by its filters itself, so a fixed value is neither a paging parameter nor a filter.
func (r *Resource) readListQuery(get foundOp) error {
	for _, name := range slices.Sorted(maps.Keys(r.Policy.GetQuery)) {
		if pagingName(name) {
			return &contractError{code: "GET_QUERY_UNKNOWN", location: getFromListLocation + ".query." + name,
				message:     fmt.Sprintf("The query parameter %q pages the List, which a tool does itself.", name),
				remediation: "Name a query parameter of the List that is not a paging parameter, or delete the line."}
		}
		if _, filter := r.Policy.ListFilters[name]; filter {
			return fmt.Errorf("api.getFromList.query and api.getFromList.filters both name %q; keep one", name)
		}
	}
	return r.readGetQuery(get, getFromListLocation+".query")
}

// readPaging records the query parameters that page Get, and the property of its response that holds
// the next page token.
func (r *Resource) readPaging(get foundOp) error {
	overridden := map[string]bool{}
	for _, p := range get.op.Parameters {
		if p.In == "query" {
			overridden[p.Name] = true
		}
	}
	parameters := slices.DeleteFunc(slices.Clone(get.item.Parameters), func(p *v3.Parameter) bool {
		return p.In == "query" && overridden[p.Name]
	})
	parameters = append(parameters, get.op.Parameters...)
	for _, p := range parameters {
		if p.In != "query" || !pagingName(p.Name) {
			continue
		}
		if t, err := typeOf(p.Schema, listLocation+"."+p.Name, walk{policy: &r.Policy}); err != nil {
			return &contractError{code: "GET_FROM_LIST_PAGING_TYPE", location: listLocation + "." + p.Name,
				message:     fmt.Sprintf("The paging parameter %q has a type the model cannot read: %v", p.Name, err),
				remediation: "Page the List with a scalar page token and page size, or a pagination object the model can read."}
		} else {
			r.PageParameters = append(r.PageParameters, Parameter{Name: p.Name, Type: t})
		}
	}
	response := mustSchema(responseRef(mustProxy(okSchema(get.op))))
	for _, name := range propertyNames(response) {
		if normalName(name) != "nextpagetoken" {
			continue
		}
		if t, err := typeOf(propertyOf(response, name), listLocation+"."+name, walk{policy: &r.Policy}); err != nil {
			return &contractError{code: "GET_FROM_LIST_PAGING_TYPE", location: listLocation + "." + name,
				message:     fmt.Sprintf("The next page property %q has a type the model cannot read: %v", name, err),
				remediation: "Page the List with a scalar page token and page size, or a pagination object the model can read."}
		} else {
			r.NextPage = &Parameter{Name: name, Type: t}
		}
	}
	return nil
}

// mustProxy returns proxy, or nil when err is set.
func mustProxy(proxy *base.SchemaProxy, err error) *base.SchemaProxy {
	if err != nil {
		return nil
	}
	return proxy
}

// readDeleteOperation reads the Delete of a GetFromList resource: it takes the id in its one path
// parameter on the Create path, or in the id array of its body (Policy.DeleteIDs).
func (r *Resource) readDeleteOperation(ops map[verb]foundOp) error {
	if r.Policy.DeleteIDs != "" {
		return r.readDeleteIDs(ops[opDelete])
	}
	id, err := idParam(ops[opDelete])
	if err != nil {
		return fmt.Errorf("api.getFromList delete: %w", err)
	}
	r.IDParam = id
	if want := ops[opCreate].path + "/{" + id + "}"; !r.Policy.CustomMethods && ops[opDelete].path != want {
		return fmt.Errorf("api.getFromList delete: path %s, want %s", ops[opDelete].path, want)
	}
	return nil
}

// readDeleteIDs reads a Delete that removes resources through its body (Policy.DeleteIDs): it has no
// path parameter, the body holds the id in a string array that accepts one item, and the other
// properties that it requires are DeleteValues. The response is ignored.
func (r *Resource) readDeleteIDs(del foundOp) error {
	if n := len(pathParams(del)); n != 0 {
		return fmt.Errorf("api.delete.ids: %s has %d path parameters, want none", del.op.OperationId, n)
	}
	location := "paths.delete." + del.op.OperationId + ".requestBody"
	fail := func(code, at, message, remediation string) error {
		return &contractError{code: code, location: at, message: message, remediation: remediation}
	}
	body, _, required := validationObject(bodySchemaOrNil(bodyProxy(del.op)))
	ids := propertyOf(body, r.Policy.DeleteIDs)
	if ids == nil {
		return fail("DELETE_IDS_NOT_FOUND", location, fmt.Sprintf("The body of %s has no property %q for the ids.", del.op.OperationId, r.Policy.DeleteIDs),
			"Name the body property that holds the ids to remove.")
	}
	if item := mustSchema(arrayItems(ids)); item == nil || !slices.Equal(item.Type, []string{"string"}) {
		return fail("DELETE_IDS_NOT_STRING_ARRAY", location+"."+r.Policy.DeleteIDs, fmt.Sprintf("The property %q is not an array of string ids.", r.Policy.DeleteIDs),
			"Name the body property that holds the string ids of the resources.")
	}
	if s := mustSchema(ids); s.MinItems != nil && *s.MinItems > 1 || s.MaxItems != nil && *s.MaxItems < 1 {
		return fail("DELETE_IDS_ONE_ITEM_REFUSED", location+"."+r.Policy.DeleteIDs, fmt.Sprintf("The property %q does not accept one id. A tool removes one resource at a time.", r.Policy.DeleteIDs),
			"Accept one id in the array.")
	}
	for _, name := range slices.Sorted(maps.Keys(r.Policy.DeleteValues)) {
		at := "behavior-overrides.yaml:api.delete.values." + name
		proxy := propertyOf(body, name)
		if proxy == nil || name == r.Policy.DeleteIDs {
			return fail("DELETE_VALUE_UNKNOWN", at, fmt.Sprintf("The body of %s has no property %q beside the ids.", del.op.OperationId, name),
				"Name a property of the Delete body, or delete the line.")
		}
		value, problem := r.fixedValue(proxy, name, r.Policy.DeleteValues[name], at)
		if problem != "" {
			return fail("DELETE_VALUE_INVALID", at, problem, "Use a value that the property accepts.")
		}
		r.DeleteValues = append(r.DeleteValues, value)
	}
	for _, name := range required {
		if _, sent := r.Policy.DeleteValues[name]; name != r.Policy.DeleteIDs && !sent {
			return fail("DELETE_BODY_REQUIRED_PROPERTY", location+"."+name, fmt.Sprintf("The body of %s requires %q, which a remove does not send.", del.op.OperationId, name),
				"Give the property a value in api.delete.values, or make it optional in the contract.")
		}
	}
	return nil
}

// readListFilters reads the query parameters of the List that narrow it by a resource field
// (Policy.ListFilters). A filter is a hint: it can match more than the resource, so a tool still
// selects the item by its id.
func (r *Resource) readListFilters(get foundOp) error {
	if len(r.Policy.ListFilters) == 0 {
		return nil
	}
	if !r.GetFromList {
		return errors.New("api.getFromList.filters needs api.getFromList")
	}
	params := map[string]*v3.Parameter{}
	for _, p := range slices.Concat(get.item.Parameters, get.op.Parameters) {
		if p.In == "query" && !pagingName(p.Name) {
			params[p.Name] = p
		}
	}
	for _, param := range slices.Sorted(maps.Keys(r.Policy.ListFilters)) {
		name := r.Policy.ListFilters[param]
		location := getFromListLocation + ".filters." + param
		fail := func(code, message, remediation string) error {
			return &contractError{code: code, location: location, message: message, remediation: remediation}
		}
		p := params[param]
		if p == nil {
			return fail("LIST_FILTER_UNKNOWN", fmt.Sprintf("Get %s has no query parameter %q that filters the List.", get.op.OperationId, param),
				"Name a query parameter of the List that is not a paging parameter, or delete the line.")
		}
		i := slices.IndexFunc(r.Fields, func(f *ResourceField) bool { return f.Name == name })
		if i < 0 {
			return fail("LIST_FILTER_FIELD_UNKNOWN", fmt.Sprintf("%s has no field %q for the filter %q.", r.Name, name, param),
				"Name the resource field that the filter matches.")
		}
		field := r.Fields[i].Type
		t, err := typeOf(p.Schema, location, walk{policy: &r.Policy})
		if err != nil || len(field.Wrappers) != 0 || !slices.Contains([]Kind{String, Bool, Integer, Number, Enum}, t.Kind) ||
			t.Kind != field.Kind || t.Format != field.Format || t.Schema != field.Schema {
			return fail("LIST_FILTER_TYPE_INCOMPATIBLE", fmt.Sprintf("The filter %q does not have the scalar type of the field %q.", param, name),
				"Name a filter of the type of the field, or delete the line.")
		}
		r.ListFilters = append(r.ListFilters, ListFilter{Param: param, Field: name, Type: t})
	}
	return nil
}
