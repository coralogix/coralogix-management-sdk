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

	"github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/orderedmap"
)

// Parent is the object field of the API that holds a top-level field of the resource, when the
// policy hoists that object (Policy.Hoist): the tool shows its fields as top-level fields of the
// resource, and the API sends and returns them inside the object.
type Parent struct {
	Field  string // API property of the object, for example "spec"
	Schema string // component of the object
	// Arm: the object is an arm of a oneOf group of the resource response, and the file skips the
	// other arms. When Get returns no such object, the resource does not exist.
	Arm bool
}

// hoisted is a request body or the resource response with each hoisted object replaced by its
// fields.
type hoisted struct {
	schema  *base.Schema
	parents map[string]Parent // top-level field → the object that holds it in the API
}

// hoist returns s with the objects that the policy hoists replaced by their fields, in place
// and in order. A hoisted object must be sent in every request that has it, so the request has
// its required fields: it must be required there. In the response it is required, or an arm of a
// oneOf group whose other arms the file skips: an absent arm then means that the resource does not
// exist. resource is the component of the resource, whose lines name the skipped arms.
func (p Policy) hoist(s *base.Schema, resource string, response bool) (hoisted, error) {
	out := hoisted{schema: s}
	if s == nil || len(p.Hoist) == 0 {
		return out, nil
	}
	names := propertyNames(s)
	if !slices.ContainsFunc(names, p.hoists) {
		return out, nil
	}
	if len(s.AllOf) != 0 || len(s.AnyOf) != 0 {
		return out, fmt.Errorf("hoist: an object with allOf or anyOf is not supported")
	}
	arms, err := hoistedArms(s, p, resource, response)
	if err != nil {
		return out, err
	}
	flat := &base.Schema{Type: s.Type, Title: s.Title, Description: s.Description, Properties: orderedmap.New[string, *base.SchemaProxy]()}
	out.schema, out.parents = flat, map[string]Parent{}
	add := func(name string, proxy *base.SchemaProxy, required bool) error {
		if flat.Properties.GetOrZero(name) != nil {
			return fmt.Errorf("hoist: two fields have the name %q", name)
		}
		flat.Properties.Set(name, proxy)
		if required {
			flat.Required = append(flat.Required, name)
		}
		return nil
	}
	for _, name := range names {
		proxy := propertyOf(s, name)
		if !p.hoists(name) {
			if err := add(name, proxy, slices.Contains(s.Required, name)); err != nil {
				return out, err
			}
			continue
		}
		parent, object, err := hoistedObject(name, proxy)
		if err != nil {
			return out, err
		}
		parent.Arm = arms[name]
		if !parent.Arm && !slices.Contains(s.Required, name) {
			return out, fmt.Errorf("hoist: %s must be required, so that its fields are always in the API value", name)
		}
		for _, child := range propertyNames(object) {
			if p.hoists(child) {
				return out, fmt.Errorf("hoist: %s.%s is hoisted too, but only a field of the resource can be", name, child)
			}
			if err := add(child, propertyOf(object, child), slices.Contains(object.Required, child)); err != nil {
				return out, err
			}
			out.parents[child] = parent
		}
	}
	if len(arms) == 0 {
		flat.OneOf = s.OneOf // the group has no hoisted arm, so it stays a group of fields
	}
	return out, nil
}

// hoistedArms returns the hoisted fields that are an arm of a oneOf group of the response s.
// The other arms of that group must be skipped, so the group means that the hoisted arm is set,
// or that the resource does not exist. A request cannot have a hoisted arm: the request must
// always send the object.
func hoistedArms(s *base.Schema, p Policy, resource string, response bool) (map[string]bool, error) {
	if len(s.OneOf) == 0 {
		return nil, nil
	}
	group, err := oneOfArms(s)
	if err != nil {
		return nil, err
	}
	var hoistedArm string
	for _, arm := range group.Arms {
		if p.hoists(arm) {
			hoistedArm = arm
		}
	}
	if hoistedArm == "" {
		return nil, nil
	}
	if !response {
		return nil, fmt.Errorf("hoist: %s is an arm of a oneOf group of a request, so the request does not always send it", hoistedArm)
	}
	for _, arm := range group.Arms {
		if arm != hoistedArm && !p.skips(resource, arm) {
			return nil, fmt.Errorf("hoist: %s is an arm of a oneOf group; skip the other arm %s, or the resource could be that arm", hoistedArm, arm)
		}
	}
	return map[string]bool{hoistedArm: true}, nil
}

// hoistedObject returns the parent and the object schema of a hoisted field: a property that is
// a $ref to an object component, or an allOf with that one $ref.
func hoistedObject(name string, proxy *base.SchemaProxy) (Parent, *base.Schema, error) {
	ref := unwrapRef(proxy)
	if ref == "" {
		return Parent{}, nil, fmt.Errorf("hoist: %s must be a $ref to an object component", name)
	}
	component, err := componentName(ref)
	if err != nil {
		return Parent{}, nil, fmt.Errorf("hoist: %s: %w", name, err)
	}
	object, err := schemaOf(proxy)
	if err != nil {
		return Parent{}, nil, fmt.Errorf("hoist: %s: %w", name, err)
	}
	if !proxy.IsReference() {
		if object, err = schemaOf(object.AllOf[0]); err != nil {
			return Parent{}, nil, fmt.Errorf("hoist: %s: %w", name, err)
		}
	}
	if object.Properties == nil || object.Properties.Len() == 0 || len(object.OneOf) != 0 || len(object.AllOf) != 0 || len(object.AnyOf) != 0 {
		return Parent{}, nil, fmt.Errorf("hoist: %s must be an object with fields and no oneOf, allOf, or anyOf", name)
	}
	return Parent{Field: name, Schema: component}, object, nil
}

// hoistAll returns the Create body, the Update body, and the resource response with the
// hoisted objects replaced by their fields.
func (p Policy) hoistAll(resource string, create, update, get *base.Schema) (c, u, g hoisted, err error) {
	if c, err = p.hoist(create, resource, false); err != nil {
		return c, u, g, fmt.Errorf("create body: %w", err)
	}
	if u, err = p.hoist(update, resource, false); err != nil {
		return c, u, g, fmt.Errorf("update body: %w", err)
	}
	if g, err = p.hoist(get, resource, true); err != nil {
		return c, u, g, fmt.Errorf("%s: %w", resource, err)
	}
	return c, u, g, nil
}
