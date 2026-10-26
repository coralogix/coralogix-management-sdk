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

	"github.com/pb33f/libopenapi/datamodel/high/base"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

// listLocation is the location of the issues of a List.
const listLocation = "paths.list"

// pagingParameters are the query parameters that page a List, and pagingProperties are the
// response properties that return the next page. A name is compared in lower case without "_"
// and "-", so page_token and pageToken are both "pagetoken". A parameter in a pagination
// object, for example pagination.page_size, also pages.
var (
	pagingParameters = map[string]bool{
		"pagetoken": true, "nextpagetoken": true, "pagesize": true, "pageoffset": true,
		"offset": true, "limit": true, "pagination": true,
	}
	pagingProperties = map[string]bool{"nextpagetoken": true, "pagination": true}
)

// readList sets the List of the resource. explicit is the operationId that names the List. A
// List that it names and that the model cannot describe is an error. With no explicit name, the
// List is the GET on the Create path. When the model cannot describe it, List stays nil and
// ListProblem says why, so a resource without a usable List is still eligible.
func (r *Resource) readList(doc *v3.Document, get foundOp, explicit string) error {
	list, err := r.findList(doc, get, explicit)
	switch {
	case err == nil:
		r.List = list
		return nil
	case explicit != "":
		return err
	}
	r.ListProblem = err.message
	return nil
}

// findList finds the List and reads it. It returns nil and no error when no name is explicit
// and the Create path has no GET.
func (r *Resource) findList(doc *v3.Document, get foundOp, explicit string) (*ListOperation, *contractError) {
	var candidates []foundOp
	switch {
	case explicit != "":
		candidates = matchingOperations(doc, r.Name, opList, explicit)
	case !r.Singleton:
		// The Get of a singleton is the GET on the Create path, so a singleton has no List.
		candidates = collectionGets(doc, r.Create.Path)
	}
	switch {
	case len(candidates) == 0 && explicit == "":
		return nil, nil
	case len(candidates) == 0:
		return nil, &contractError{code: "LIST_OPERATION_NOT_FOUND", location: listLocation,
			message:     fmt.Sprintf("No operation has the operationId %q.", explicit),
			remediation: "Name a GET operation of the API contract that lists the resources."}
	case len(candidates) > 1:
		var names []string
		for _, c := range candidates {
			names = append(names, c.op.OperationId)
		}
		slices.Sort(names)
		return nil, &contractError{code: "LIST_OPERATION_AMBIGUOUS", location: listLocation,
			message:     fmt.Sprintf("More than one GET operation lists the resources: %s.", strings.Join(names, ", ")),
			remediation: "Name one List operation, or keep one GET on the collection path in the source API contract."}
	}
	return r.readListOperation(get, candidates[0])
}

// collectionGets returns the GET operations on the Create path. A trailing "/" does not
// change the path.
func collectionGets(doc *v3.Document, createPath string) []foundOp {
	var out []foundOp
	for path, item := range doc.Paths.PathItems.FromOldest() {
		if strings.TrimSuffix(path, "/") != strings.TrimSuffix(createPath, "/") || item.Get == nil {
			continue
		}
		out = append(out, foundOp{path: path, method: "GET", item: item, op: item.Get})
	}
	return out
}

// readListOperation checks the List operation f and reads its items.
func (r *Resource) readListOperation(get, f foundOp) (*ListOperation, *contractError) {
	location := listLocation + "." + f.op.OperationId
	fail := func(code, at, message, remediation string) (*ListOperation, *contractError) {
		return nil, &contractError{code: code, location: at, message: message, remediation: remediation}
	}
	if r.Singleton {
		return fail("LIST_SINGLETON_UNSUPPORTED", location, "A singleton has one resource, so it has no List.", "Delete the List operation id.")
	}
	if f.method != "GET" {
		return fail("LIST_OPERATION_METHOD_INCOMPATIBLE", location, fmt.Sprintf("The method is %s. A List needs GET.", f.method), "Name a GET operation.")
	}
	if params := pathParams(f); len(params) != 0 {
		return fail("LIST_PATH_PARAMETER_UNSUPPORTED", location+".parameters.path."+params[0].Name,
			fmt.Sprintf("The List has the path parameter %q. A List needs none.", params[0].Name), "Name a GET operation with no path parameters.")
	}
	for _, p := range slices.Concat(f.item.Parameters, f.op.Parameters) {
		if p.In != "path" && p.Required != nil && *p.Required && !pagingName(p.Name) {
			return fail("LIST_REQUIRED_PARAMETER_UNSUPPORTED", fmt.Sprintf("%s.parameters.%s.%s", location, p.In, p.Name),
				fmt.Sprintf("The List requires the %s parameter %q, which does not page the result.", p.In, p.Name), "Make the parameter optional in the source API contract.")
		}
	}
	if tag, getTag := oneTag(f.op), oneTag(get.op); tag == "" || tag != getTag {
		return fail("LIST_TAG_INCOMPATIBLE", location, fmt.Sprintf("The List has the tags %v. A List needs the one tag of Get, %q.", f.op.Tags, getTag),
			"Give the List the tag of the other operations, so the SDK has it in the same package.")
	}
	schema, err := okSchema(f.op)
	if err != nil {
		return fail("LIST_RESPONSE_UNSUPPORTED", location, fmt.Sprintf("The List has %s.", err), "Return a response component from the List.")
	}
	proxy := responseRef(schema)
	if proxy == nil {
		return fail("LIST_RESPONSE_UNSUPPORTED", location, "The 200 response schema of the List is inline. A List needs a $ref.", "Return a response component from the List.")
	}
	response, cerr := componentName(proxy.GetReference())
	if cerr != nil {
		return fail("LIST_RESPONSE_UNSUPPORTED", location, cerr.Error()+".", "Return a response component from the List.")
	}
	s, err := schemaOf(proxy)
	if err != nil {
		return fail("LIST_RESPONSE_UNSUPPORTED", location, fmt.Sprintf("The response %s: %s.", response, err), "Return a response component from the List.")
	}
	arrays := itemArrays(s)
	switch len(arrays) {
	case 0:
		return fail("LIST_ITEMS_NOT_FOUND", "components.schemas."+response, fmt.Sprintf("The response %s has no array of component objects.", response),
			"Return the resources in one array of a component in the List response.")
	case 1:
	default:
		var names []string
		for _, a := range arrays {
			names = append(names, a.field)
		}
		return fail("LIST_ITEMS_AMBIGUOUS", "components.schemas."+response, fmt.Sprintf("The response %s has more than one array of component objects: %s.", response, strings.Join(names, ", ")),
			"Return the resources in the only array of objects of the List response.")
	}
	items := arrays[0]
	list := &ListOperation{
		Operation:     Operation{Method: f.method, Path: f.path, OperationID: f.op.OperationId, Response: Response{Schema: response}},
		ItemsField:    items.field,
		ItemComponent: items.component,
		Full:          items.component == r.Name,
		IDField:       r.IDField,
		Paged:         pagedList(f, s),
	}
	if err := r.listItemID(get, items); err != nil {
		return nil, err
	}
	if name, err := schemaOf(propertyOf(items.schema, "name")); err == nil && slices.Equal(name.Type, []string{"string"}) {
		list.NameField = "name"
	}
	return list, nil
}

// itemArray is a property of a List response that is an array of component objects.
type itemArray struct {
	field, component string
	schema           *base.Schema // the item component
}

// itemArrays returns the properties of the response s that are arrays of component objects:
// the items are "$ref: X" or "allOf: [$ref: X]", and X is an object.
func itemArrays(s *base.Schema) []itemArray {
	var out []itemArray
	for _, field := range propertyNames(s) {
		property, err := schemaOf(propertyOf(s, field))
		if err != nil || !slices.Equal(property.Type, []string{"array"}) {
			continue
		}
		proxy := itemsOf(property)
		component, err := componentName(unwrapRef(proxy))
		if err != nil {
			continue
		}
		if !proxy.IsReference() {
			proxy = proxy.Schema().AllOf[0] // unwrapRef found the $ref in a single allOf
		}
		item, err := schemaOf(proxy)
		if err != nil || !objectSchema(item) {
			continue
		}
		out = append(out, itemArray{field: field, component: component, schema: item})
	}
	return out
}

// objectSchema reports whether s is an object with properties, or has the type object.
func objectSchema(s *base.Schema) bool {
	return slices.Equal(s.Type, []string{"object"}) || len(s.Type) == 0 && s.Properties != nil && s.Properties.Len() != 0
}

// listItemID checks that the item has the id property of the resource, with the type of the
// id in the Get response. A tool reads the id of an item to call Get.
func (r *Resource) listItemID(get foundOp, items itemArray) *contractError {
	location := "components.schemas." + items.component + "." + r.IDField
	id := propertyOf(items.schema, r.IDField)
	if id == nil {
		return &contractError{code: "LIST_ITEM_ID_MISSING", location: "components.schemas." + items.component,
			message:     fmt.Sprintf("The List item %s has no id property %q.", items.component, r.IDField),
			remediation: "Return the resource id in each List item."}
	}
	want, err := schemaOf(propertyOf(mustSchema(responseResourceProxy(get.op, r.Name)), r.IDField))
	got, gotErr := schemaOf(id)
	if err != nil || gotErr != nil || !slices.Equal(got.Type, want.Type) || got.Format != want.Format {
		return &contractError{code: "LIST_ITEM_ID_TYPE_INCOMPATIBLE", location: location,
			message:     fmt.Sprintf("The id property %q of the List item %s does not have the type of the id in the Get response.", r.IDField, items.component),
			remediation: "Give the id in each List item the type of the resource id."}
	}
	return nil
}

// mustSchema returns the schema of proxy, or nil.
func mustSchema(proxy *base.SchemaProxy) *base.Schema {
	s, err := schemaOf(proxy)
	if err != nil {
		return nil
	}
	return s
}

// pagedList reports whether the List f pages its result: a query parameter pages it, or the
// response s has a property that returns the next page.
func pagedList(f foundOp, s *base.Schema) bool {
	for _, p := range slices.Concat(f.item.Parameters, f.op.Parameters) {
		if p.In == "query" && pagingName(p.Name) {
			return true
		}
	}
	return slices.ContainsFunc(propertyNames(s), func(name string) bool { return pagingProperties[normalName(name)] })
}

// pagingName reports whether the query parameter name pages a List.
func pagingName(name string) bool {
	head, _, _ := strings.Cut(normalName(name), ".")
	return pagingParameters[normalName(name)] || head == "pagination"
}

// normalName is name in lower case without "_" and "-".
func normalName(name string) string {
	return strings.ToLower(strings.NewReplacer("_", "", "-", "").Replace(name))
}

// oneTag returns the tag of op, or "" when op does not have exactly one.
func oneTag(op *v3.Operation) string {
	if len(op.Tags) != 1 {
		return ""
	}
	return op.Tags[0]
}
