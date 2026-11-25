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
	"math"
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/coralogix/coralogix-management-sdk/iac/issue"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	"go.yaml.in/yaml/v4"
)

// extIac is the extension that the OpenAPI fork writes from openapiv3_field.iac.
const extIac = "x-coralogix-iac"

// IaCHints are the facts that the API owner states about one field, for the tests that a tool
// generates. A hint never changes the schema of the resource. The fork writes them from the knobs
// of the proto field (openapiv3_field.iac), in x-coralogix-iac.
//
// Names are the JSON property names of the sibling fields, the same as Field.Name.
type IaCHints struct {
	// Examples are values of the field, as the API writes them (JSON). Examples[0] is the value
	// for a create, and Examples[1] is the value for an update. A map or an object example holds
	// the whole value. Each example is a string, a bool, a number, a []any, or a map[string]any.
	Examples []any
	// RequiredByAPI: the API refuses a create that does not set the field.
	RequiredByAPI bool
	// ConflictsWith are the sibling fields that the API refuses together with this field.
	ConflictsWith []string
	// Requires are the sibling fields that the API needs when this field is set.
	Requires []IaCRequires
	// ArmFixed is true on a oneOf arm: an update cannot switch to or from this arm.
	ArmFixed bool

	// problems are the mistakes that the reader found. Validate reports them.
	problems []iacProblem
}

// IaCRequires says that a field needs the sibling field Field, and its value Equals when
// HasEquals is true.
type IaCRequires struct {
	Field     string
	Equals    any
	HasEquals bool
}

type iacProblem struct{ message, fix string }

// IaC returns the hints of the field, or nil. The hints sit on the schema of the field in the
// resource response, in the Create request, or in the Update request. The first one in that
// order is the result. Validate reports hints that differ between these places.
func (f *ResourceField) IaC() *IaCHints { return firstIaC(f.Get, f.Create, f.Update) }

// IaC returns the hints of the field, or nil. See ResourceField.IaC.
func (f *Field) IaC() *IaCHints { return firstIaC(&f.Attrs, f.Create, f.Update) }

func firstIaC(attrs ...*Attrs) *IaCHints {
	for _, a := range attrs {
		if a != nil && a.IaC != nil {
			return a.IaC
		}
	}
	return nil
}

// iacOf reads x-coralogix-iac of a property schema. It returns nil when the extension is absent.
// A mistake does not stop the build: Validate reports it, because the hints change no schema.
// Keys that this reader does not know are ignored, so an older reader works with a newer fork.
func iacOf(s *base.Schema) *IaCHints {
	node := s.Extensions.GetOrZero(extIac)
	if node == nil {
		return nil
	}
	h := &IaCHints{}
	bad := func(message, fix string) { h.problems = append(h.problems, iacProblem{message, fix}) }
	if node.Kind != yaml.MappingNode {
		bad(extIac+" is not an object.", "Write the knobs in openapiv3_field.iac, for example iac: { required_by_api: true }.")
		return h
	}
	for i := 0; i+1 < len(node.Content); i += 2 {
		key, value := node.Content[i].Value, node.Content[i+1]
		switch key {
		case "examples":
			iacList(value, key, bad, func(n *yaml.Node, at string) {
				var v any
				if err := n.Decode(&v); err != nil {
					bad(at+" is not a JSON value.", "Write the example as JSON, for example '\"text\"' or '{\"key\": 1}'.")
					return
				}
				h.Examples = append(h.Examples, v)
			})
		case "requiredByApi":
			h.RequiredByAPI = iacBool(value, key, bad)
		case "armFixed":
			h.ArmFixed = iacBool(value, key, bad)
		case "conflictsWith":
			iacList(value, key, bad, func(n *yaml.Node, at string) {
				h.ConflictsWith = append(h.ConflictsWith, n.Value)
			})
		case "requires":
			iacList(value, key, bad, func(n *yaml.Node, at string) {
				h.Requires = append(h.Requires, iacRequire(n, at, bad))
			})
		}
	}
	return h
}

// iacList calls add for each item of a sequence node.
func iacList(n *yaml.Node, key string, bad func(message, fix string), add func(*yaml.Node, string)) {
	if n.Kind != yaml.SequenceNode {
		bad(fmt.Sprintf("%s is not a list.", key), "Write a list.")
		return
	}
	for i, item := range n.Content {
		add(item, fmt.Sprintf("%s[%d]", key, i))
	}
}

func iacBool(n *yaml.Node, key string, bad func(message, fix string)) bool {
	switch n.Value {
	case "true":
		return true
	case "false":
		return false
	}
	bad(fmt.Sprintf("%s is %q, want true or false.", key, n.Value), "Write true or false.")
	return false
}

func iacRequire(n *yaml.Node, at string, bad func(message, fix string)) IaCRequires {
	var r IaCRequires
	if n.Kind != yaml.MappingNode {
		bad(at+" is not an object.", "Write requires: { field: \"other_field\" }.")
		return r
	}
	for i := 0; i+1 < len(n.Content); i += 2 {
		switch n.Content[i].Value {
		case "field":
			r.Field = n.Content[i+1].Value
		case "equals":
			if err := n.Content[i+1].Decode(&r.Equals); err != nil {
				bad(at+".equals is not a JSON value.", "Write equals as JSON, for example '\"apm\"'.")
				continue
			}
			r.HasEquals = true
		}
	}
	return r
}

// iacSubject is one field with the facts that the hint checks need.
type iacSubject struct {
	name  string
	typ   *Type
	attrs []*Attrs // response, Create, Update
}

// iacHintIssues reports the hints of the resource that the API contract cannot hold: a hint that
// the reader could not parse, an example that does not fit the type of the field, a name that is
// not a sibling field, and an arm_fixed on a field that is not a oneOf arm.
func iacHintIssues(r *Resource) issue.Report {
	subjects := make([]iacSubject, 0, len(r.Fields))
	for _, f := range r.Fields {
		subjects = append(subjects, iacSubject{f.Name, f.Type, []*Attrs{f.Get, f.Create, f.Update}})
	}
	return iacFieldIssues("components.schemas."+r.Name, subjects, armNames(r.Groups, false))
}

func iacFieldIssues(location string, subjects []iacSubject, arms []string) issue.Report {
	var report issue.Report
	byName := make(map[string]*Type, len(subjects))
	for _, s := range subjects {
		byName[s.name] = s.typ
	}
	for _, s := range subjects {
		at := location + "." + s.name
		report = append(report, iacOwnIssues(at, s, byName, slices.Contains(arms, s.name))...)
		report = append(report, iacNestedIssues(at, s.typ)...)
	}
	return report
}

func iacNestedIssues(location string, t *Type) issue.Report {
	if t == nil {
		return nil
	}
	switch t.Kind {
	case List, Set, Map:
		return iacNestedIssues(location+"[]", t.Elem)
	case Object, OneOf:
		subjects := make([]iacSubject, 0, len(t.Fields))
		for _, f := range t.Fields {
			subjects = append(subjects, iacSubject{f.Name, f.Type, []*Attrs{&f.Attrs, f.Create, f.Update}})
		}
		return iacFieldIssues(location, subjects, armNames(t.Groups, t.Kind == OneOf, t.Fields...))
	}
	return nil
}

// armNames returns the names of the oneOf arms: every field of a OneOf, and the arms of the groups.
func armNames(groups []OneOfGroup, all bool, fields ...*Field) []string {
	var out []string
	if all {
		for _, f := range fields {
			out = append(out, f.Name)
		}
	}
	for _, g := range groups {
		out = append(out, g.Arms...)
	}
	return out
}

func iacOwnIssues(location string, s iacSubject, siblings map[string]*Type, isArm bool) issue.Report {
	var found []*IaCHints
	for _, a := range s.attrs {
		if a != nil && a.IaC != nil {
			found = append(found, a.IaC)
		}
	}
	if len(found) == 0 {
		return nil
	}
	h := found[0]
	var problems []iacProblem
	for _, other := range found[1:] {
		if !reflect.DeepEqual(h, other) {
			problems = append(problems, iacProblem{
				"The field has different iac knobs in the response and in a request.",
				"Put the knobs on one proto field, or write the same knobs on each.",
			})
			break
		}
	}
	for _, found := range found {
		problems = append(problems, found.problems...)
	}
	problems = append(problems, iacValueProblems(h, s, siblings, isArm)...)
	var report issue.Report
	for _, p := range problems {
		report = append(report, issue.Issue{Code: "IAC_HINT_INVALID", Location: location, Message: p.message, Remediation: p.fix})
	}
	return report
}

func iacValueProblems(h *IaCHints, s iacSubject, siblings map[string]*Type, isArm bool) []iacProblem {
	var out []iacProblem
	for i, example := range h.Examples {
		for _, message := range iacValueProblemsOf(s.typ, example, fmt.Sprintf("examples[%d]", i)) {
			out = append(out, iacProblem{message, "Change the example to a value that the API accepts for this field, written as JSON."})
		}
	}
	sibling := func(key, name string) (*Type, bool) {
		t, ok := siblings[name]
		switch {
		case name == s.name:
			out = append(out, iacProblem{fmt.Sprintf("%s names the field itself.", key), "Name another field of the same message."})
		case !ok:
			out = append(out, iacProblem{fmt.Sprintf("%s: %q is not a sibling field.", key, name),
				"Use the proto name (snake_case) of a field in the same message. The JSON names of the siblings are: " + strings.Join(sortedNames(siblings), ", ") + "."})
		}
		return t, ok && name != s.name
	}
	for i, name := range h.ConflictsWith {
		sibling(fmt.Sprintf("conflicts_with[%d]", i), name)
	}
	for i, req := range h.Requires {
		key := fmt.Sprintf("requires[%d]", i)
		t, ok := sibling(key+".field", req.Field)
		if !ok || !req.HasEquals {
			continue
		}
		for _, message := range iacValueProblemsOf(t, req.Equals, key+".equals") {
			out = append(out, iacProblem{message, "Change equals to a value that field " + req.Field + " accepts, written as JSON."})
		}
	}
	if h.ArmFixed && !isArm {
		out = append(out, iacProblem{"arm_fixed is on a field that is not a oneOf arm.", "Remove arm_fixed, or put it on a field of a oneof."})
	}
	return out
}

func sortedNames(siblings map[string]*Type) []string {
	names := make([]string, 0, len(siblings))
	for name := range siblings {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// iacValueProblemsOf returns why v is not a value of t. at names the value in the messages. A
// type that the API holds in another shape than the model (a wrapper, a projection, a keyed map)
// is not checked: its JSON shape is not the one of the model.
func iacValueProblemsOf(t *Type, v any, at string) []string {
	if t == nil || len(t.Wrappers) != 0 || t.Projected != nil || t.Keyed != nil || t.RequestWrapper != nil {
		return nil
	}
	if v == nil {
		return []string{at + " is null. A null is not a value."}
	}
	switch t.Kind {
	case String:
		return iacStringProblems(t, v, at)
	case Enum:
		return iacEnumProblems(t, v, at)
	case Bool:
		if _, ok := v.(bool); !ok {
			return []string{fmt.Sprintf("%s is %s, want a bool.", at, jsonKind(v))}
		}
	case Number, Integer:
		return iacNumberProblems(t, v, at)
	case List, Set:
		return iacListProblems(t, v, at)
	case Map:
		return iacMapProblems(t, v, at)
	case Object, OneOf:
		return iacObjectProblems(t, v, at)
	}
	return nil
}

func iacStringProblems(t *Type, v any, at string) []string {
	s, ok := v.(string)
	if !ok {
		return []string{fmt.Sprintf("%s is %s, want a string.", at, jsonKind(v))}
	}
	var out []string
	n := int64(utf8.RuneCountInString(s))
	if t.MinLength != nil && n < *t.MinLength {
		out = append(out, fmt.Sprintf("%s has %d characters, the minimum is %d.", at, n, *t.MinLength))
	}
	if t.MaxLength != nil && n > *t.MaxLength {
		out = append(out, fmt.Sprintf("%s has %d characters, the maximum is %d.", at, n, *t.MaxLength))
	}
	if t.Pattern != "" {
		if re, err := regexp.Compile(t.Pattern); err == nil && !re.MatchString(s) {
			out = append(out, fmt.Sprintf("%s does not match the pattern %s.", at, t.Pattern))
		}
	}
	if t.Format == "date-time" {
		if _, err := time.Parse(time.RFC3339Nano, s); err != nil {
			out = append(out, fmt.Sprintf("%s is not an RFC 3339 time, for example \"2026-01-02T03:04:05Z\".", at))
		}
	}
	return out
}

func iacEnumProblems(t *Type, v any, at string) []string {
	s, ok := v.(string)
	switch {
	case !ok:
		return []string{fmt.Sprintf("%s is %s, want an enum value.", at, jsonKind(v))}
	case s == t.EnumZero:
		return []string{fmt.Sprintf("%s is the zero value %s, which means unset. Use one of: %s.", at, s, strings.Join(t.Values, ", "))}
	case !slices.Contains(t.Values, s):
		return []string{fmt.Sprintf("%s is %q, which is not a value of the enum. Use one of: %s.", at, s, strings.Join(t.Values, ", "))}
	}
	return nil
}

func iacNumberProblems(t *Type, v any, at string) []string {
	f, ok := numberOf(v)
	if !ok && t.WireString {
		// The API writes a 64-bit number as a string, and accepts digits.
		if s, isString := v.(string); isString {
			if n, err := strconv.ParseFloat(s, 64); err == nil && !strings.ContainsAny(s, ".eE") {
				f, ok = n, true
			}
		}
	}
	if !ok {
		return []string{fmt.Sprintf("%s is %s, want a number.", at, jsonKind(v))}
	}
	var out []string
	if t.Kind == Integer && f != math.Trunc(f) {
		out = append(out, fmt.Sprintf("%s is %v, want an integer.", at, f))
	}
	if t.Minimum != nil && f < *t.Minimum {
		out = append(out, fmt.Sprintf("%s is %v, the minimum is %v.", at, f, *t.Minimum))
	}
	if t.Maximum != nil && f > *t.Maximum {
		out = append(out, fmt.Sprintf("%s is %v, the maximum is %v.", at, f, *t.Maximum))
	}
	return out
}

func iacListProblems(t *Type, v any, at string) []string {
	items, ok := v.([]any)
	if !ok {
		return []string{fmt.Sprintf("%s is %s, want a list.", at, jsonKind(v))}
	}
	var out []string
	if t.MinItems != nil && int64(len(items)) < *t.MinItems {
		out = append(out, fmt.Sprintf("%s has %d items, the minimum is %d.", at, len(items), *t.MinItems))
	}
	if t.MaxItems != nil && int64(len(items)) > *t.MaxItems {
		out = append(out, fmt.Sprintf("%s has %d items, the maximum is %d.", at, len(items), *t.MaxItems))
	}
	for i, item := range items {
		out = append(out, iacValueProblemsOf(t.Elem, item, fmt.Sprintf("%s[%d]", at, i))...)
	}
	return out
}

func iacMapProblems(t *Type, v any, at string) []string {
	entries, ok := v.(map[string]any)
	if !ok {
		return []string{fmt.Sprintf("%s is %s, want an object that maps keys to values.", at, jsonKind(v))}
	}
	var out []string
	if t.MinItems != nil && int64(len(entries)) < *t.MinItems {
		out = append(out, fmt.Sprintf("%s has %d entries, the minimum is %d.", at, len(entries), *t.MinItems))
	}
	if t.MaxItems != nil && int64(len(entries)) > *t.MaxItems {
		out = append(out, fmt.Sprintf("%s has %d entries, the maximum is %d.", at, len(entries), *t.MaxItems))
	}
	for _, key := range sortedKeys(entries) {
		out = append(out, iacValueProblemsOf(t.Elem, entries[key], at+"."+key)...)
	}
	return out
}

func iacObjectProblems(t *Type, v any, at string) []string {
	object, ok := v.(map[string]any)
	if !ok {
		return []string{fmt.Sprintf("%s is %s, want an object.", at, jsonKind(v))}
	}
	var out []string
	for _, key := range sortedKeys(object) {
		f := fieldNamed(t.Fields, key)
		if f == nil {
			names := make([]string, 0, len(t.Fields))
			for _, f := range t.Fields {
				names = append(names, f.Name)
			}
			out = append(out, fmt.Sprintf("%s has the key %q, which is not a field. The fields are: %s.", at, key, strings.Join(names, ", ")))
			continue
		}
		if f.Attrs.ReadOnly {
			out = append(out, fmt.Sprintf("%s.%s is read only: the server sets it.", at, key))
			continue
		}
		out = append(out, iacValueProblemsOf(f.Type, object[key], at+"."+key)...)
	}
	for _, f := range t.Fields {
		if _, set := object[f.Name]; !set && f.Create != nil && f.Create.Required && t.Kind == Object && !slices.Contains(armNames(t.Groups, false), f.Name) {
			out = append(out, fmt.Sprintf("%s does not set the required field %q.", at, f.Name))
		}
	}
	return append(out, iacArmProblems(t, object, at)...)
}

// iacArmProblems checks that an example sets one arm of each oneOf group, not several.
func iacArmProblems(t *Type, object map[string]any, at string) []string {
	groups := t.Groups
	if t.Kind == OneOf {
		g := OneOfGroup{AllowNone: t.AllowNone}
		for _, f := range t.Fields {
			g.Arms = append(g.Arms, f.Name)
		}
		groups = []OneOfGroup{g}
	}
	var out []string
	for _, g := range groups {
		var set []string
		for _, arm := range g.Arms {
			if _, ok := object[arm]; ok {
				set = append(set, arm)
			}
		}
		switch {
		case len(set) > 1:
			out = append(out, fmt.Sprintf("%s sets %s, but at most one of %s can be set.", at, strings.Join(set, " and "), strings.Join(g.Arms, ", ")))
		case len(set) == 0 && !g.AllowNone:
			out = append(out, fmt.Sprintf("%s sets none of %s, but one must be set.", at, strings.Join(g.Arms, ", ")))
		}
	}
	return out
}

func numberOf(v any) (float64, bool) {
	switch n := v.(type) {
	case int:
		return float64(n), true
	case int64:
		return float64(n), true
	case uint64:
		return float64(n), true
	case float64:
		return n, true
	}
	return 0, false
}

func jsonKind(v any) string {
	switch v.(type) {
	case string:
		return "a string"
	case bool:
		return "a bool"
	case []any:
		return "a list"
	case map[string]any:
		return "an object"
	case nil:
		return "null"
	}
	return "a number"
}

func sortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
