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
	"slices"
	"sort"

	"github.com/pb33f/libopenapi/datamodel/high/base"

	"github.com/coralogix/coralogix-management-sdk/iac/model"
)

// API is the HTTP contract of a released resource.
type API struct {
	// RequestWrapper is the property of the Create and Update bodies that holds the resource.
	RequestWrapper string `yaml:"requestWrapper"`
	// UpdateIDInBody: Update has no id in its path. The id is a field of the resource.
	UpdateIDInBody bool `yaml:"updateIDInBody"`
	// ClientSetID: the client can send the id on Create.
	ClientSetID bool `yaml:"clientSetID"`
	// ServerSetID: the server sets the id and returns it in every response, but the Get
	// response cannot require it, because the request shares the resource schema.
	ServerSetID bool `yaml:"serverSetID"`
	// CustomMethods: Get, Update, and Delete are not the Create path with an id.
	// Get and Delete each have the same {id} path parameter on their own paths.
	// Update has no path parameter, and the id stays in the body. New resources
	// keep the stricter path rules.
	CustomMethods bool `yaml:"customMethods"`
	// Operations names the lifecycle operationIds when they do not end in
	// _Create, _Get, _Update, or _Delete plus the resource name. A generator then
	// needs no operation flags.
	Operations *Operations `yaml:"operations"`
	// Delete names the operation that removes the resource when the API has no DELETE.
	Delete *Delete `yaml:"delete"`
	// UpdateMaskInBody: PATCH updateMask is an optional JSON body property.
	// The generated update sets it to the same changed-field mask a query
	// parameter would send.
	UpdateMaskInBody bool `yaml:"updateMaskInBody"`
}

// Operations names the four lifecycle operationIds. All four are required.
type Operations struct {
	Create string `yaml:"create"`
	Get    string `yaml:"get"`
	Update string `yaml:"update"`
	Delete string `yaml:"delete"`
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
	// Unwrap: true shows the value inside the wrapper of the field, or of its items, in this
	// place. false keeps the object of a component that the unwrap list names.
	Unwrap *bool `yaml:"unwrap"`
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
	if a.Delete != nil && a.Delete.Operation == "" {
		return errors.New("api.delete.operation is required")
	}
	if a.Operations == nil {
		return nil
	}
	if a.Operations.Create == "" || a.Operations.Get == "" || a.Operations.Update == "" || a.Operations.Delete == "" {
		return errors.New("api.operations needs create, get, update, and delete")
	}
	if a.Delete != nil {
		return errors.New("api.delete and api.operations both name the Delete operation: keep only api.operations.delete")
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
		Existing:         true,
		RequestWrapper:   e.API.RequestWrapper,
		UpdateIDInBody:   e.API.UpdateIDInBody,
		ClientSetID:      e.API.ClientSetID,
		ServerSetID:      e.API.ServerSetID,
		CustomMethods:    e.API.CustomMethods,
		UpdateMaskInBody: e.API.UpdateMaskInBody,
	}
	if e.API.Operations != nil {
		p.Operations = model.OperationIDs{
			Create: e.API.Operations.Create,
			Get:    e.API.Operations.Get,
			Update: e.API.Operations.Update,
			Delete: e.API.Operations.Delete,
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
