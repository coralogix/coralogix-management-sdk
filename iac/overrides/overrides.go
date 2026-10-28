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

// Package overrides has the keys of a behavior-overrides file that describe the API of a
// resource that users already have: where its API differs from the contract of a new
// resource. Every tool reads them the same way. A tool owns the file: it adds its own keys
// next to these, decodes the whole file, and adds the exceptions that its own keys cover to
// the Policy.
package overrides

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"sort"
	"strings"

	"github.com/pb33f/libopenapi/datamodel/high/base"

	"github.com/coralogix/coralogix-management-sdk/iac/model"
)

// API is the HTTP contract of a released resource.
type API struct {
	// RequestWrapper is the property of the Create and Update bodies that holds the resource.
	RequestWrapper string `yaml:"requestWrapper"`
	// UpdateIDInBody: Update has no id in its path. The id is a field of the resource.
	UpdateIDInBody bool `yaml:"updateIDInBody"`
	// UpdateIDOutsideWrapper: with requestWrapper and updateIDInBody, the Update body carries the
	// id property beside the wrapper property, not inside it.
	UpdateIDOutsideWrapper bool `yaml:"updateIDOutsideWrapper"`
	// RequestNames maps a property of the resource component to the path of the property that
	// holds it in the Create and Update bodies, after the request wrapper. The path is
	// dot-separated through object properties, for example {definitionKey: integrationKey,
	// parameters: integrationParameters.parameters}. A request property that the key covers is
	// not a separate field.
	RequestNames map[string]string `yaml:"requestNames"`
	// ClientSetID: the client can send the id on Create.
	ClientSetID bool `yaml:"clientSetID"`
	// ServerSetID: the server sets the id and returns it in every response, but the Get
	// response cannot require it, because the request shares the resource schema.
	ServerSetID bool `yaml:"serverSetID"`
	// IDProperty is the property of the resource that holds the id when it does not have the name
	// of the path parameter, for example uniqueIdentifier for /rules/{rule_id}. "" means the
	// property named like the path parameter.
	IDProperty string `yaml:"idProperty"`
	// CreateResponseOnlyIDProperty is the property of the Create response that holds only the id of
	// the new resource, when the response does not return the resource. A tool then reads the
	// resource with Get.
	CreateResponseOnlyIDProperty string `yaml:"createResponseOnlyIDProperty"`
	// CustomMethods: Get, Update, and Delete are not the Create path with an id.
	// Get and Delete each have the same {id} path parameter on their own paths.
	// Update has no path parameter, and the id stays in the body. New resources
	// keep the stricter path rules.
	CustomMethods bool `yaml:"customMethods"`
	// Operations names the lifecycle operationIds when they do not end in
	// _Create, _Get, _Update, or _Delete plus the resource name. A generator then
	// needs no operation flags. It can also name the List.
	Operations *Operations `yaml:"operations"`
	// Delete names the operation that removes the resource when the API has no DELETE.
	Delete *Delete `yaml:"delete"`
	// UpdateMaskInBody: PATCH updateMask is an optional JSON body property.
	// The generated update sets it to the same changed-field mask a query
	// parameter would send.
	UpdateMaskInBody bool `yaml:"updateMaskInBody"`
}

// Operations names the four lifecycle operationIds, and the List. The four lifecycle ids are
// set together, or not at all.
type Operations struct {
	Create string `yaml:"create"`
	Get    string `yaml:"get"`
	Update string `yaml:"update"`
	Delete string `yaml:"delete"`
	// List is the operationId of the GET that lists the resources, when it is not the GET on
	// the Create path. A List that it names must be one that the model can describe.
	List string `yaml:"list"`
}

// Delete is how a released resource is removed when the API has no DELETE, for example a
// POST /things/{id}/archive. It is an object, so that a later key can add a step before the call.
type Delete struct {
	// Operation is the operationId of a POST whose path is the Get path plus one segment. It
	// takes the id path parameter of Get and no request body. Its response is ignored.
	Operation string `yaml:"operation"`
}

// Field is the part of a field line that describes the API. A tool embeds it in its own
// field line with the yaml inline flag, so both parts are one line of the file.
type Field struct {
	// Skip: the resource does not manage the field. It is not sent and not stored.
	Skip bool `yaml:"skip"`
	// ReadOnly: true says that the server sets the field, although the contract does not mark
	// it readOnly yet. The field is read and stored, and never sent. Delete this key when the
	// contract marks the field readOnly.
	ReadOnly bool `yaml:"readOnly"`
	// Optional: true makes the field optional, although a request of the contract requires it.
	// The released resource lets the user leave it out, and the API accepts that, but the
	// contract cannot say so: the request shares the component with a response that always has
	// the field.
	Optional bool `yaml:"optional"`
	// CreateAs folds top-level Create fields into this object field: Create sends each
	// property of the object as the top-level Create field it names, for example
	// {limit: permutationsLimit}. Update and Get have the object itself.
	CreateAs map[string]string `yaml:"createAs"`
	// Collection: set says that the order of a list of the contract has no meaning. The API
	// keeps the order.
	Collection string `yaml:"collection"`
	// KeyedBy says that a list of objects is a map keyed by a string property of the object.
	// With Value, the map value is that one property, and the object must have only the two.
	KeyedBy *KeyedBy `yaml:"keyedBy"`
	// Unwrap: true shows the value inside the wrapper of the field, or of its items, in this
	// place. false keeps the object of a component that the unwrap list names.
	Unwrap *bool `yaml:"unwrap"`
}

// KeyedBy names the key property, and the value property, of a list of objects that is a map.
type KeyedBy struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
	// Object shows the list as an object with one attribute per accepted value of the key, an
	// enum with an enums line. The attribute has the name that the tool gives the enum value.
	Object bool `yaml:"object"`
	// Arms gives an enum value of the key the oneOf arm of the item that it holds. The other
	// arms are not in its attribute, and the arm is required in it.
	Arms map[string]string `yaml:"arms"`
	// KeyAs is a string property of the item that expand sets to the attribute name.
	KeyAs string `yaml:"keyAs"`
}

// Exceptions are the API keys of one behavior-overrides file.
type Exceptions struct {
	API API
	// Unwrap lists the components with one property that the released resource shows as the
	// value of that property, for example LuceneQuery in {"luceneQuery": {"value": "error"}}. It
	// applies in every place that has no unwrap key on its field line.
	Unwrap []string
	// Fields has the API part of each field line: component, then field.
	Fields map[string]map[string]Field
}

// Check checks that a Delete override names its operation, and that only one line names it.
func (a API) Check() error {
	if a.ClientSetID && a.ServerSetID {
		return errors.New("api.clientSetID and api.serverSetID both describe the id: keep the one that states who sets it")
	}
	if a.UpdateIDOutsideWrapper && (a.RequestWrapper == "" || !a.UpdateIDInBody) {
		return errors.New("api.updateIDOutsideWrapper needs api.requestWrapper and api.updateIDInBody")
	}
	if err := checkRequestNames(a.RequestNames); err != nil {
		return err
	}
	if a.Delete != nil && a.Delete.Operation == "" {
		return errors.New("api.delete.operation is required")
	}
	if a.Operations == nil {
		return nil
	}
	o := a.Operations
	if o.List != "" && o.Create == "" && o.Get == "" && o.Update == "" && o.Delete == "" {
		return nil // only the List
	}
	if o.Create == "" || o.Get == "" || o.Update == "" || o.Delete == "" {
		return errors.New("api.operations needs create, get, update, and delete")
	}
	if a.Delete != nil {
		return errors.New("api.delete and api.operations both name the Delete operation: keep only api.operations.delete")
	}
	return nil
}

// Check checks the createAs, collection, and keyedBy keys of a field line.
func (l Field) Check() error {
	if l.KeyedBy != nil && (l.KeyedBy.Key == "" || l.KeyedBy.Key == l.KeyedBy.Value) {
		return errors.New("keyedBy needs a key, and a value other than the key")
	}
	if k := l.KeyedBy; k != nil && k.Object && k.Value != "" {
		return errors.New("keyedBy: object holds the whole item, so value cannot be set")
	}
	if k := l.KeyedBy; k != nil && !k.Object && (len(k.Arms) != 0 || k.KeyAs != "") {
		return errors.New("keyedBy: arms and keyAs need object: true")
	}
	if l.KeyedBy != nil && l.Collection != "" {
		return errors.New("keyedBy makes a map, so collection cannot be set")
	}
	if err := checkCreateAs(l.CreateAs); err != nil {
		return err
	}
	if l.Collection != "" && l.Collection != "set" {
		return fmt.Errorf("collection is %q, want set", l.Collection)
	}
	return nil
}

// checkCreateAs checks a createAs map: each property names one Create field, and no two
// properties name the same one.
func checkCreateAs(into map[string]string) error {
	seen := map[string]bool{}
	for nested, flat := range into {
		if nested == "" || flat == "" {
			return errors.New("createAs maps a property to a Create field; neither can be empty")
		}
		if seen[flat] {
			return fmt.Errorf("createAs names the Create field %q twice", flat)
		}
		seen[flat] = true
	}
	return nil
}

// checkRequestNames checks a requestNames map: each resource field and each request path is a
// non-empty name, and no two fields name the same path.
func checkRequestNames(names map[string]string) error {
	seen := map[string]string{}
	for _, field := range sortedKeys(names) {
		path := names[field]
		if field == "" || path == "" || slices.Contains(strings.Split(path, "."), "") {
			return fmt.Errorf("api.requestNames maps a resource field to a request path of dot-separated names; %q: %q is not one", field, path)
		}
		if other, ok := seen[path]; ok {
			return fmt.Errorf("api.requestNames maps both %q and %q to the request path %q", other, field, path)
		}
		seen[path] = field
	}
	return nil
}

// CheckUnwrap checks that the unwrap list names each component once.
func CheckUnwrap(unwrap []string) error {
	for i, name := range unwrap {
		if name == "" {
			return fmt.Errorf("unwrap[%d] is empty", i)
		}
		if slices.Contains(unwrap[:i], name) {
			return fmt.Errorf("unwrap names %s twice", name)
		}
	}
	return nil
}

// Policy returns the rule set of a released resource that the API keys state. A field that
// a line skips or marks readOnly is released: the line says how an omitted value behaves.
// The tool adds the exceptions that its own keys cover.
func (e Exceptions) Policy() model.Policy {
	p := model.Policy{
		Existing:                     true,
		RequestWrapper:               e.API.RequestWrapper,
		UpdateIDInBody:               e.API.UpdateIDInBody,
		UpdateIDOutsideWrapper:       e.API.UpdateIDOutsideWrapper,
		RequestNames:                 maps.Clone(e.API.RequestNames),
		ClientSetID:                  e.API.ClientSetID,
		ServerSetID:                  e.API.ServerSetID,
		IDProperty:                   e.API.IDProperty,
		CreateResponseOnlyIDProperty: e.API.CreateResponseOnlyIDProperty,
		CustomMethods:                e.API.CustomMethods,
		UpdateMaskInBody:             e.API.UpdateMaskInBody,
	}
	if e.API.Operations != nil {
		p.Operations = model.OperationIDs{
			Create: e.API.Operations.Create,
			Get:    e.API.Operations.Get,
			Update: e.API.Operations.Update,
			Delete: e.API.Operations.Delete,
			List:   e.API.Operations.List,
		}
	}
	if e.API.Delete != nil {
		p.DeleteOperation = e.API.Delete.Operation
	}
	p.Unwrap = slices.Clone(e.Unwrap)
	for _, name := range sortedKeys(e.Fields) {
		fields := e.Fields[name]
		for _, field := range sortedKeys(fields) {
			line := fields[field]
			if line.Unwrap != nil {
				if p.UnwrapFields == nil {
					p.UnwrapFields = map[string]bool{}
				}
				p.UnwrapFields[name+"."+field] = *line.Unwrap
			}
			if line.Skip {
				p.Skip = append(p.Skip, name+"."+field)
			}
			if line.ReadOnly {
				p.ReadOnly = append(p.ReadOnly, name+"."+field)
			}
			if line.Optional {
				p.Optional = append(p.Optional, name+"."+field)
			}
			if line.Collection == "set" {
				p.Sets = append(p.Sets, name+"."+field)
			}
			if k := line.KeyedBy; k != nil {
				if p.KeyedBy == nil {
					p.KeyedBy = map[string]model.KeyedLine{}
				}
				p.KeyedBy[name+"."+field] = model.KeyedLine{Key: k.Key, Value: k.Value, Object: k.Object, Arms: maps.Clone(k.Arms), KeyAs: k.KeyAs}
			}
			if into := line.CreateAs; len(into) != 0 {
				if p.CreateAs == nil {
					p.CreateAs = map[string]map[string]string{}
				}
				p.CreateAs[name+"."+field] = maps.Clone(into)
			}
			if line.Skip || line.ReadOnly {
				p.Released = append(p.Released, name+"."+field)
			}
		}
	}
	return p
}

// ContractReadOnly reports whether the contract marks field of schema readOnly. A readOnly
// line on that field is then stale.
func ContractReadOnly(schema *base.Schema, field string) bool {
	built, err := schema.Properties.GetOrZero(field).BuildSchema()
	return err == nil && built.ReadOnly != nil && *built.ReadOnly
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return slices.Clip(keys)
}
