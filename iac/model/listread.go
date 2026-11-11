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

	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

// getFromListLocation is the line of the behavior-overrides file that says Get reads the List.
const getFromListLocation = "behavior-overrides.yaml:api.getFromList"

// readListOperations reads a resource that no operation reads by its id: Get is the List, and a tool
// selects the item whose id is the resource id. Create and Get have no path parameters, Update sends
// the id in its body on the Create path, and Delete takes the id in its path.
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
	id, err := idParam(ops[opDelete])
	if err != nil {
		return fmt.Errorf("api.getFromList delete: %w", err)
	}
	r.IDParam = id
	if want := ops[opCreate].path + "/{" + id + "}"; !r.Policy.CustomMethods && ops[opDelete].path != want {
		return fmt.Errorf("api.getFromList delete: path %s, want %s", ops[opDelete].path, want)
	}
	if err := r.readTargets(ops); err != nil {
		return err
	}
	r.GetFromList, r.List = true, list
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
