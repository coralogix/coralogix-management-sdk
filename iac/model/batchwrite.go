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
	"strings"

	"github.com/pb33f/libopenapi/datamodel/high/base"
)

// batchWriteLocation is the line of the behavior-overrides file that says the writes take arrays.
const batchWriteLocation = "behavior-overrides.yaml:api.batchWrite"

// readBatchBodies reads the Create and Update bodies of a BatchWrite resource. Each is an array of
// one request component that accepts one item, and the Operation names the component. The rest of
// the model reads the item where it reads a body (Policy.bodyOf).
func (r *Resource) readBatchBodies(ops map[verb]foundOp) error {
	if !r.Policy.BatchWrite {
		return nil
	}
	if !r.Policy.Existing {
		return errors.New("api.batchWrite is only for existing resources")
	}
	if arrayItems(bodyProxy(ops[opCreate].op)) == nil && arrayItems(bodyProxy(ops[opUpdate].op)) == nil {
		return &contractError{code: "OVERRIDE_UNUSED", location: batchWriteLocation,
			message:     "Neither the Create nor the Update body is an array.",
			remediation: "Delete api.batchWrite."}
	}
	writes := map[verb]*Operation{opCreate: &r.Create, opUpdate: &r.Update}
	for _, v := range []verb{opCreate, opUpdate} {
		o := writes[v]
		location := "paths." + strings.ToLower(string(v)) + "." + o.OperationID + ".requestBody"
		notArray := func(message string) error {
			return &contractError{code: "BATCH_WRITE_BODY_NOT_ARRAY", location: location, message: message,
				remediation: "Describe a resource whose Create and Update both take an array of one request component, or delete api.batchWrite."}
		}
		body := bodyProxy(ops[v].op)
		items := arrayItems(body)
		if items == nil {
			return notArray(fmt.Sprintf("The %s body is not an array. api.batchWrite needs both writes to take an array.", v))
		}
		component, err := componentName(unwrapRef(items))
		if err != nil {
			return notArray(fmt.Sprintf("The items of the %s body are not a component.", v))
		}
		if s := mustSchema(body); s.MinItems != nil && *s.MinItems > 1 || s.MaxItems != nil && *s.MaxItems < 1 {
			return &contractError{code: "BATCH_WRITE_ONE_ITEM_REFUSED", location: location,
				message:     fmt.Sprintf("The %s body does not accept an array of one item. A tool writes one resource at a time.", v),
				remediation: "Accept one item in the array body."}
		}
		o.Body, o.BodyItems = component, true
	}
	return nil
}

// resultResponse reads the Create response of a BatchWrite resource that answers with one result per
// request item (Policy.CreateResult). The response has one array of result components. A result holds
// the resource in a property, and the status of the item in an enum property.
func (r *Resource) resultResponse(out Response, proxy *base.SchemaProxy) (Response, error) {
	line := r.Policy.CreateResult
	if !r.Policy.BatchWrite {
		return Response{}, errors.New("api.batchWrite.createResult needs api.batchWrite")
	}
	location := batchWriteLocation + ".createResult"
	fail := func(code, message, remediation string) (Response, error) {
		return Response{}, &contractError{code: code, location: location, message: message, remediation: remediation}
	}
	s, err := schemaOf(proxy)
	if err != nil {
		return Response{}, fmt.Errorf("response %s: %w", out.Schema, err)
	}
	arrays := itemArrays(s)
	switch {
	case len(arrays) == 0:
		return fail("CREATE_RESULT_ITEMS_NOT_FOUND", fmt.Sprintf("The Create response %s has no array of result components.", out.Schema),
			"Name a Create that answers with one result per request item, or delete createResult.")
	case len(arrays) > 1:
		return fail("CREATE_RESULT_ITEMS_AMBIGUOUS", fmt.Sprintf("The Create response %s has more than one array of components.", out.Schema),
			"Return the results in the only array of components of the Create response.")
	}
	results := arrays[0]
	if resource, err := componentName(unwrapRef(propertyOf(results.schema, line.Resource))); err != nil || resource != r.Name {
		return fail("CREATE_RESULT_RESOURCE_NOT_FOUND", fmt.Sprintf("The result %s has no property %q that holds %s.", results.component, line.Resource, r.Name),
			"Name the result property that holds the created resource.")
	}
	status := propertyOf(results.schema, line.Status)
	statusType, err := typeOf(status, fieldLocation(results.component, line.Status), walk{policy: &r.Policy})
	if status == nil || err != nil || statusType.Kind != Enum {
		return fail("CREATE_RESULT_STATUS_NOT_ENUM", fmt.Sprintf("The result %s has no enum property %q.", results.component, line.Status),
			"Name the enum property that holds the status of the item.")
	}
	if len(line.Created) == 0 {
		return fail("CREATE_RESULT_STATUS_VALUE_INVALID", "createResult names no status value that means created.",
			"List the status values that mean that the item was created.")
	}
	for _, value := range line.Created {
		problem := queryValueProblem(status, value)
		if value == statusType.EnumZero {
			problem = "is the zero value"
		}
		if problem != "" {
			return fail("CREATE_RESULT_STATUS_VALUE_INVALID", fmt.Sprintf("The status value %q of %s.%s %s.", value, results.component, line.Status, problem),
				"List business values of the status enum.")
		}
	}
	if line.Message != "" {
		if message, err := schemaOf(propertyOf(results.schema, line.Message)); err != nil || !slices.Equal(message.Type, []string{"string"}) {
			return fail("CREATE_RESULT_MESSAGE_NOT_STRING", fmt.Sprintf("The result %s has no string property %q.", results.component, line.Message),
				"Name the string property that says why an item was not created, or delete message.")
		}
	}
	out.Result = &WriteResult{Items: results.field, Component: results.component, Resource: line.Resource, Status: line.Status,
		StatusType: statusType, Created: slices.Clone(line.Created), Message: line.Message}
	return out, nil
}

// readCreateValues reads the properties beside the request wrapper that Create always sends, with
// their values (Policy.CreateValues). Each is a scalar or an enum, and the value is one that it accepts
// and not its zero value, which sending nothing already means.
func (r *Resource) readCreateValues(create foundOp) error {
	if len(r.Policy.CreateValues) == 0 {
		return nil
	}
	if !r.Policy.Existing {
		return errors.New("api.createValues is only for existing resources")
	}
	wrapper := r.Policy.WrapperOf(true)
	body, _, _ := validationObject(bodySchemaOrNil(r.Policy.bodyOf(create.op)))
	inside := mustSchema(r.Policy.requestBody(create.op))
	for _, name := range slices.Sorted(maps.Keys(r.Policy.CreateValues)) {
		value := r.Policy.CreateValues[name]
		location := "behavior-overrides.yaml:api.createValues." + name
		fail := func(code, message, remediation string) error {
			return &contractError{code: code, location: location, message: message, remediation: remediation}
		}
		proxy := propertyOf(body, name)
		switch {
		case wrapper == "":
			return fail("CREATE_VALUE_UNKNOWN", fmt.Sprintf("Create %s has no request wrapper, so no property %q sits beside it.", create.op.OperationId, name),
				"Set requestWrapper, or delete the line.")
		case name == wrapper || propertyOf(inside, name) != nil:
			return fail("CREATE_VALUE_INSIDE_WRAPPER", fmt.Sprintf("%q is the request wrapper %q or a property inside it: a field of the resource, not a fixed value.", name, wrapper),
				"Name a property beside the request wrapper, or delete the line.")
		case proxy == nil:
			return fail("CREATE_VALUE_UNKNOWN", fmt.Sprintf("The Create body of %s has no property %q beside %q.", create.op.OperationId, name, wrapper),
				"Name a property beside the request wrapper, or delete the line.")
		}
		fixed, problem := r.fixedValue(proxy, name, value, location)
		if problem != "" {
			return fail("CREATE_VALUE_INVALID", problem, "Use a value that the property accepts.")
		}
		r.CreateValues = append(r.CreateValues, fixed)
	}
	return nil
}

// fixedValue reads a body property that a call always sends with value. problem says why the value
// does not fit: the property is a scalar or an enum, the value is one that it accepts, and it is not
// the zero value of an enum, which an omitted property already sends.
func (r *Resource) fixedValue(proxy *base.SchemaProxy, name, value, location string) (QueryValue, string) {
	t, err := typeOf(proxy, location, walk{policy: &r.Policy})
	problem := queryValueProblem(proxy, value)
	switch {
	case err != nil || !slices.Contains([]Kind{String, Bool, Integer, Number, Enum}, t.Kind):
		problem = "is not a value of a scalar or an enum property"
	case problem == "" && t.Kind == Enum && value == t.EnumZero:
		problem = "is the zero value, which an omitted property already sends"
	}
	if problem != "" {
		return QueryValue{}, fmt.Sprintf("The value %q of the property %q %s.", value, name, problem)
	}
	return QueryValue{Name: name, Value: value, Type: t}, ""
}
