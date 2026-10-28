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

// Package sdkcheck lists the Go names of the management SDK that generated code uses, and
// checks that the pinned SDK has each of them.
package sdkcheck

import (
	"fmt"
	"slices"
	"strings"
	"unicode"

	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"

	"github.com/coralogix/coralogix-management-sdk/iac/model"
)

// sdkGenRoot is the parent of the generated SDK packages of module.
func sdkGenRoot(module string) string { return module + "/go/openapi/gen" }

// sdkClientSet is the handwritten cxsdk package of module.
func sdkClientSet(module string) string { return module + "/go/openapi/cxsdk" }

// Kind is the kind of Go object a Ref names.
type Kind string

const (
	KindPackage Kind = "package"
	KindType    Kind = "type"
	KindField   Kind = "field"
	KindMethod  Kind = "method"
	KindFunc    Kind = "func"
	KindConst   Kind = "const"
)

// Rule is how an SDK name is derived from the spec. The "special" rules copy
// a choice that the SDK makes and the spec does not show
// (findings F15–F18).
type Rule string

const (
	RuleTag         Rule = "tag"                   // the operation tag
	RuleOperationID Rule = "operationId"           // the operationId
	RuleParameter   Rule = "parameter"             // an operation parameter
	RuleComponent   Rule = "component"             // the component schema name
	RuleProperty    Rule = "property"              // the property name
	RuleEnumValue   Rule = "enum value"            // the component name and the enum value
	RuleInlineBody  Rule = "special: inline body"  // F15
	RuleEmptyObject Rule = "special: empty object" // F16
	RuleClientSet   Rule = "special: cxsdk"        // F17
)

// Ref is one SDK name that the generated code uses.
type Ref struct {
	Path  string // model path, for example "fields.config.sqlLoad.joinLimit"
	Kind  Kind
	Pkg   string // import path of the package that has the name
	Owner string // field, method: the type that has it
	Name  string // package: the package name
	// Want is the Go type of a field or the signature of a method, written
	// relative to Pkg. It is "" for the other kinds.
	Want string
	// WantValue is the value type that a field may have instead of the
	// pointer type Want. Issues sets Want to it when the SDK has it.
	WantValue string
	// ByType: a method that the SDK may name another way. When Name is not
	// found, the check uses the one method of Owner with type Want (F17).
	ByType bool
	Rule   Rule
	// Schema is the component name of a type ref that comes from the spec.
	// The generator uses it to find the SDK type of a model type.
	Schema string
}

// sdkName is the name as the Go code writes it, for example "SqlLoadConfig.JoinLimit".
func (s Ref) SDKName() string {
	if s.Owner == "" {
		return s.Name
	}
	return s.Owner + "." + s.Name
}

// ResourceTag returns the one tag that all four operations share. The SDK
// makes one package per tag.
func ResourceTag(doc *v3.Document, r *model.Resource) (string, error) {
	tag := ""
	for _, op := range []model.Operation{r.Create, r.Get, r.Update, r.Delete} {
		item, ok := doc.Paths.PathItems.Get(op.Path)
		if !ok {
			return "", fmt.Errorf("%s: path %s not in the spec", op.OperationID, op.Path)
		}
		o, ok := item.GetOperations().Get(strings.ToLower(op.Method))
		if !ok {
			return "", fmt.Errorf("%s: method %s not in the spec", op.OperationID, op.Method)
		}
		if len(o.Tags) != 1 {
			return "", fmt.Errorf("%s: %d tags, want 1", op.OperationID, len(o.Tags))
		}
		if tag != "" && o.Tags[0] != tag {
			return "", fmt.Errorf("%s: tag %q, want %q (as in the other operations)", op.OperationID, o.Tags[0], tag)
		}
		tag = o.Tags[0]
	}
	return tag, nil
}

// PackageName is the name of the generated SDK package of an operation tag, for example
// "global_routers_service" for "Global Routers Service".
func PackageName(tag string) string { return strings.ToLower(strings.ReplaceAll(tag, " ", "_")) }

// ClientType is the API client type of the SDK package of tag, for example
// "GlobalRoutersServiceAPIService".
func ClientType(tag string) string { return camelize(tag) + "APIService" }

// Names lists the SDK names that the generated code uses for the resource. It only applies
// the naming rules. Issues confirms that the names exist. zeroConstants reports the enum
// components whose zero value the generated code uses, so the SDK must have its constant. It
// can be nil.
//
// A tool appends the names of its own code that the generated code uses, for example its
// client set, before it loads and checks the list.
func Names(r *model.Resource, tag, module string, zeroConstants func(component string) bool) ([]Ref, error) {
	pkgName := PackageName(tag)
	s := &resolver{pkg: sdkGenRoot(module) + "/" + pkgName, seen: map[string]bool{}, bodies: map[string]string{}, policy: r.Policy, zeroConstants: zeroConstants}
	client := ClientType(tag)
	s.add(Ref{Path: "resource", Kind: KindPackage, Name: pkgName, Rule: RuleTag})
	s.add(Ref{Path: "resource", Kind: KindType, Name: client, Rule: RuleTag})

	if err := s.operations(r, client); err != nil {
		return nil, err
	}

	resource := goTypeName(r.Name)
	s.add(Ref{Path: "fields", Kind: KindType, Name: resource, Rule: RuleComponent, Schema: r.Name})
	for _, f := range r.Fields {
		path := "fields." + f.Name
		if f.InGet {
			if err := s.field(path, resource, f.Name, f.Type, f.Get != nil && f.Get.Required); err != nil {
				return nil, err
			}
		}
		if err := s.nested(path, f.Type); err != nil {
			return nil, err
		}
	}
	if err := s.bodyFields(r, resource); err != nil {
		return nil, err
	}
	// The SDK cxsdk package owns the API error helpers.
	for _, ref := range []Ref{
		{Path: "cxsdk", Kind: KindPackage, Name: "cxsdk"},
		{Path: "cxsdk.errors", Kind: KindFunc, Name: "NewAPIError", Want: "func(resp *http.Response, err error) error"},
		{Path: "cxsdk.errors", Kind: KindFunc, Name: "IsNotFound", Want: "func(err error) bool"},
	} {
		ref.Pkg, ref.Rule = sdkClientSet(module), RuleClientSet
		s.refs = append(s.refs, ref)
	}
	return s.refs, nil
}

// bodyFields adds the SDK fields of the Create and Update bodies. A body that wraps the resource has
// one field that holds it. The resource fields are then the fields of the resource type, which the
// "fields" names cover.
func (s *resolver) bodyFields(r *model.Resource, resource string) error {
	if r.MaskInBody {
		// The body property carries the same mask a query parameter would. It is
		// *string, or a string when the SDK uses a value.
		s.add(Ref{Path: "update.body." + r.UpdateMask, Kind: KindField, Owner: s.bodies["update"], Name: goFieldName(r.UpdateMask),
			Want: "*string", WantValue: "string", Rule: RuleProperty})
	}
	if id := ServerSetBodyID(r); id != nil && r.Policy.RequestWrapper == "" {
		// No Update field sends a server-set id, so Update sets this field of the body itself.
		if err := s.field("update.body."+id.Name, s.bodies["update"], id.Name, id.Type, true); err != nil {
			return err
		}
	}
	if wrapper := r.Policy.RequestWrapper; wrapper != "" {
		return s.wrappedBodyFields(r, wrapper, resource)
	}
	for _, f := range r.Fields {
		if folds := r.CreateFolds(f); folds != nil {
			// Create sends the properties of the object as top-level fields.
			for _, fd := range folds {
				if err := s.requestField("create.body."+fd.Field, s.bodies["create"], fd.Field, fd.Type, fd.Required); err != nil {
					return err
				}
			}
			continue
		}
		if f.Create != nil {
			if err := s.requestField("create.body."+f.Name, s.bodies["create"], f.Name, f.Type.CreateType(), f.Create.Required); err != nil {
				return err
			}
		}
	}
	for _, f := range r.Fields {
		if f.Update != nil {
			if err := s.requestField("update.body."+f.Name, s.bodies["update"], f.Name, f.Type.UpdateType(), f.Update.Required); err != nil {
				return err
			}
		}
	}
	return nil
}

// wrappedBodyFields adds the SDK names of Create and Update bodies that hold the request in one
// wrapper property. The wrapper has the component that the spec gives it (RequestComponent), which
// is the resource only when the request has the shape of the resource. A field that a request holds
// under another name or deeper (RequestPath) is a field of that component, or of an object inside
// it. An Update id beside the wrapper is a field of the body itself.
func (s *resolver) wrappedBodyFields(r *model.Resource, wrapper, resource string) error {
	for _, name := range []string{"create", "update"} {
		op := r.Create
		if name == "update" {
			op = r.Update
		}
		component := resource
		if op.RequestComponent != "" {
			component = goTypeName(op.RequestComponent)
		}
		s.add(Ref{Path: name + ".body." + wrapper, Kind: KindField, Owner: s.bodies[name], Name: goFieldName(wrapper),
			Want: "*" + component, WantValue: component, Rule: RuleProperty})
		if component == resource {
			continue // the "fields" names cover the resource
		}
		if err := s.requestComponent(r, name, wrapper, op.RequestComponent); err != nil {
			return err
		}
	}
	if r.UpdateIDOutsideWrapper {
		i := slices.IndexFunc(r.Fields, func(f *model.ResourceField) bool { return f.Name == r.ID() })
		if i < 0 {
			return fmt.Errorf("update.body: no field %q for the id beside the wrapper", r.ID())
		}
		f := r.Fields[i]
		// A tool that makes the id a server-set body id drops its Update attributes, so the check
		// accepts the value type too.
		return s.field("update.body."+f.Name, s.bodies["update"], f.Name, f.Type, true)
	}
	return nil
}

// requestComponent adds the type of the request component that the wrapper of the Create or Update
// body holds, and the Go field of each resource field that the request sends in it.
func (s *resolver) requestComponent(r *model.Resource, name, wrapper, component string) error {
	owner := goTypeName(component)
	s.addOnce(Ref{Path: name + ".body." + wrapper, Kind: KindType, Name: owner, Rule: RuleComponent, Schema: component})
	for _, f := range r.Fields {
		attrs, t := f.Create, f.Type.CreateType()
		if name == "update" {
			attrs, t = f.Update, f.Type.UpdateType()
			if r.UpdateIDOutsideWrapper && f.Name == r.ID() {
				continue // a field of the body, not of the wrapper
			}
		}
		if attrs == nil {
			continue
		}
		segments := strings.Split(f.Name, ".")
		if f.RequestPath != "" {
			segments = strings.Split(f.RequestPath, ".")
		}
		if len(f.RequestObjects) != len(segments)-1 {
			return fmt.Errorf("%s.body.%s: the request path %q goes through %d objects with a component, want %d", name, f.Name, f.RequestPath, len(f.RequestObjects), len(segments)-1)
		}
		path := name + ".body." + f.Name
		holder := owner
		for i, segment := range segments[:len(segments)-1] {
			object := goTypeName(f.RequestObjects[i])
			via := path + ".via." + segment // the object that the path goes through, not the field itself
			s.add(Ref{Path: via, Kind: KindField, Owner: holder, Name: goFieldName(segment), Want: "*" + object, WantValue: object, Rule: RuleProperty})
			s.addOnce(Ref{Path: via, Kind: KindType, Name: object, Rule: RuleComponent, Schema: f.RequestObjects[i]})
			holder = object
		}
		if err := s.requestField(path, holder, segments[len(segments)-1], t, attrs.Required); err != nil {
			return err
		}
	}
	return nil
}

// requestField adds the Go field of a request body, and the request
// components inside it. A component that the response also uses is already
// walked.
func (s *resolver) requestField(path, owner, name string, t *model.Type, allowValue bool) error {
	if err := s.field(path, owner, name, t, allowValue); err != nil {
		return err
	}
	return s.nested(path, t)
}

// operations adds the SDK names of the four operations, and of the List when the resource has one.
func (s *resolver) operations(r *model.Resource, client string) error {
	ops := []struct {
		name string
		op   model.Operation
		id   bool // the method takes the id path parameter
	}{
		{"create", r.Create, false},
		{"get", r.Get, true},
		{"update", r.Update, true},
		{"delete", r.Delete, true},
	}
	for _, o := range ops {
		// A singleton has no id in the path (D18).
		withID := o.id && !r.Singleton && (o.name != "update" || !r.Policy.UpdateIDInBody)
		if err := s.operation(r, o.name, o.op, client, withID); err != nil {
			return err
		}
	}
	if r.List != nil {
		return s.list(r, client)
	}
	return nil
}

// list adds the SDK names of the List: the call, the items field of the response, and the
// item fields that a tool reads. A resource without a List needs none of them.
func (s *resolver) list(r *model.Resource, client string) error {
	l := r.List
	if err := s.operation(r, "list", l.Operation, client, false); err != nil {
		return err
	}
	item := goTypeName(l.ItemComponent)
	s.add(Ref{Path: "list.response." + l.ItemsField, Kind: KindField, Owner: goTypeName(l.Response.Schema),
		Name: goFieldName(l.ItemsField), Want: "[]" + item, Rule: RuleProperty})
	if l.Full {
		return nil // the "fields" names cover the resource
	}
	s.add(Ref{Path: "list.item", Kind: KindType, Name: item, Rule: RuleComponent, Schema: l.ItemComponent})
	if err := s.field("list.item."+l.IDField, item, l.IDField, r.IDType, true); err != nil {
		return err
	}
	if l.NameField == "" {
		return nil
	}
	return s.field("list.item."+l.NameField, item, l.NameField, &model.Type{Kind: model.String}, true)
}

type resolver struct {
	pkg    string
	refs   []Ref
	seen   map[string]bool   // component schemas already walked
	bodies map[string]string // operation name → request body type
	policy model.Policy
	// zeroConstants reports the enum components whose zero constant the generated code uses.
	zeroConstants func(component string) bool
}

// addOnce adds a type ref that another walk of the same component may have added.
func (s *resolver) addOnce(ref Ref) {
	if !slices.ContainsFunc(s.refs, func(r Ref) bool { return r.Kind == ref.Kind && r.Name == ref.Name && r.Path == ref.Path }) {
		s.add(ref)
	}
}

func (s *resolver) add(ref Ref) {
	ref.Pkg = s.pkg
	s.refs = append(s.refs, ref)
}

// operation adds the method, the request builder, the body, and the response
// of one operation. The generated code calls, for example:
//
//	client.WidgetsServiceGetWidget(ctx, id).Execute()
func (s *resolver) operation(r *model.Resource, name string, op model.Operation, client string, withID bool) error {
	path := name
	method := camelize(op.OperationID)
	builder := "Api" + method + "Request"
	params := "ctx context.Context"
	if withID {
		idType, err := ValueType(r.IDType)
		if err != nil {
			return fmt.Errorf("%s.id: %w", path, err)
		}
		// The Go SDK names the parameter in lower camel case: alert_scheduler_rule_id becomes
		// alertSchedulerRuleId, and id stays id.
		params += ", " + lowerFirst(camelize(r.IDParam)) + " " + idType
	}
	s.add(Ref{Path: path, Kind: KindMethod, Owner: client, Name: method,
		Want: "func(" + params + ") " + builder, Rule: RuleOperationID})
	s.add(Ref{Path: path, Kind: KindType, Name: builder, Rule: RuleOperationID})
	if name == "update" && !r.Replace && !r.MaskInBody {
		s.add(Ref{Path: path + ".mask", Kind: KindMethod, Owner: builder, Name: goFieldName(r.UpdateMask),
			Want: "func(" + lowerFirst(goFieldName(r.UpdateMask)) + " string) " + builder, Rule: RuleParameter})
	}

	switch op.Body {
	case "":
	case "inline":
		body := inlineBodyName(method, op)
		s.bodies[name] = body
		s.add(Ref{Path: path + ".body", Kind: KindType, Name: body, Rule: RuleInlineBody})
		s.add(Ref{Path: path + ".body", Kind: KindMethod, Owner: builder, Name: body,
			Want: "func(" + lowerFirst(body) + " " + body + ") " + builder, Rule: RuleInlineBody})
	default:
		return fmt.Errorf("%s.body: component body %s is not supported", path, op.Body)
	}

	resp := goTypeName(op.Response.Schema)
	if op.Response.Empty {
		// An object with no fields has no SDK type (F16).
		s.add(Ref{Path: path, Kind: KindMethod, Owner: builder, Name: "Execute",
			Want: "func() (map[string]interface{}, *http.Response, error)", Rule: RuleEmptyObject})
		return nil
	}
	s.add(Ref{Path: path + ".response", Kind: KindType, Name: resp, Rule: RuleComponent})
	s.add(Ref{Path: path, Kind: KindMethod, Owner: builder, Name: "Execute",
		Want: "func() (*" + resp + ", *http.Response, error)", Rule: RuleOperationID})
	if op.Response.ID != "" {
		s.add(Ref{Path: path + ".response." + op.Response.ID, Kind: KindField, Owner: resp,
			Name: goFieldName(op.Response.ID), Want: "*string", WantValue: "string", Rule: RuleProperty})
	}
	if op.Response.Field != "" {
		s.add(Ref{Path: path + ".response." + op.Response.Field, Kind: KindField, Owner: resp,
			Name: goFieldName(op.Response.Field), Want: "*" + goTypeName(r.Name), WantValue: goTypeName(r.Name), Rule: RuleProperty})
	}
	return nil
}

// inlineBodyName is the SDK type of an inline request body. openapi-generator
// names it after its title, or after the operationId when it has none. When
// a component has the title as its name, it adds "1" (F15, F45).
func inlineBodyName(method string, op model.Operation) string {
	switch {
	case op.BodyTitle == "":
		return method + "Request"
	case op.BodyTitleIsComponent:
		return goTypeName(op.BodyTitle) + "1"
	}
	return goTypeName(op.BodyTitle)
}

// field adds the Go field for the property name of the owner type.
func (s *resolver) field(path, owner, name string, t *model.Type, allowValue bool) error {
	goType, r, err := fieldType(t)
	if err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	ref := Ref{Path: path, Kind: KindField, Owner: owner, Name: goFieldName(name), Want: goType, Rule: r}
	// openapi-generator uses a value, not a pointer, for a required field.
	// Only the SDK knows which spec it was generated from (F18), so the
	// check accepts both and keeps the one that the SDK has.
	if allowValue && strings.HasPrefix(goType, "*") {
		ref.WantValue = goType[1:]
	}
	s.add(ref)
	return nil
}

// nested adds the types, fields, and enum constants inside t. It walks each
// component schema once, so a path shows the first place that uses it.
func (s *resolver) nested(path string, t *model.Type) error {
	if len(t.Wrappers) != 0 {
		return s.wrappers(path, t)
	}
	if t.Keyed != nil {
		// The SDK holds the list of whole items, the key and the value too.
		return s.nested(path, t.Keyed.Item)
	}
	switch t.Kind {
	case model.List, model.Set, model.Map:
		return s.nested(path, t.Elem)
	case model.Enum:
		if s.seen[t.Schema] {
			return nil
		}
		s.seen[t.Schema] = true
		name := goTypeName(t.Schema)
		s.add(Ref{Path: path, Kind: KindType, Name: name, Rule: RuleComponent, Schema: t.Schema})
		for _, v := range t.Values {
			s.add(Ref{Path: path + "." + v, Kind: KindConst, Name: enumConstName(name, v), Rule: RuleEnumValue})
		}
		if s.zeroConstants != nil && s.zeroConstants(t.Schema) && t.EnumZero != "" {
			// The generated code uses the zero value, so the SDK must have it.
			s.add(Ref{Path: path + "." + t.EnumZero, Kind: KindConst, Name: enumConstName(name, t.EnumZero), Rule: RuleEnumValue})
		}
	case model.Object, model.OneOf:
		// An object with no fields has no SDK type (F16).
		if len(t.Fields) == 0 || s.seen[t.Schema] {
			return nil
		}
		s.seen[t.Schema] = true
		name := goTypeName(t.Schema)
		s.add(Ref{Path: path, Kind: KindType, Name: name, Rule: RuleComponent, Schema: t.Schema})
		for _, f := range t.Fields {
			child := path + "." + f.Name
			if err := s.field(child, name, f.Name, f.Type, f.Attrs.Required); err != nil {
				return err
			}
			if err := s.nested(child, f.Type); err != nil {
				return err
			}
		}
	}
	return nil
}

// wrappers adds the type of each wrapper of t and the field that holds the value, then the names
// inside the value. A wrapper that another place walked as an object has the same field type.
func (s *resolver) wrappers(path string, t *model.Type) error {
	for i, w := range t.Wrappers {
		if !s.seen[w.Schema] {
			s.seen[w.Schema] = true
			name := goTypeName(w.Schema)
			s.add(Ref{Path: path, Kind: KindType, Name: name, Rule: RuleComponent, Schema: w.Schema})
			rest := *t
			rest.Wrappers = t.Wrappers[i+1:]
			if err := s.field(path+"."+w.Field, name, w.Field, &rest, w.Attrs.Required); err != nil {
				return err
			}
		}
		path += "." + w.Field
	}
	return s.nested(path, t.Unwrapped())
}

// fieldType is the Go type of an optional SDK struct field. A required field
// can also be the value type (see field).
func fieldType(t *model.Type) (string, Rule, error) {
	if len(t.Wrappers) != 0 {
		return "*" + goTypeName(t.Wrappers[0].Schema), RuleProperty, nil
	}
	if (t.Kind == model.Object || t.Kind == model.OneOf) && len(t.Fields) == 0 {
		return "map[string]interface{}", RuleEmptyObject, nil
	}
	if t.Keyed != nil {
		return "[]" + goTypeName(t.Keyed.Item.Schema), RuleProperty, nil
	}
	if t.Kind == model.List || t.Kind == model.Set || t.Kind == model.Map {
		elem, _, err := fieldType(t.Elem)
		if err != nil {
			return "", "", err
		}
		if !strings.HasPrefix(elem, "*") {
			return "", "", fmt.Errorf("%s of %s is not supported", t.Kind, elem)
		}
		if t.Kind == model.Map {
			return "map[string]" + elem[1:], RuleProperty, nil
		}
		return "[]" + elem[1:], RuleProperty, nil
	}
	v, err := ValueType(t)
	if err != nil {
		return "", "", err
	}
	return "*" + v, RuleProperty, nil
}

func ValueType(t *model.Type) (string, error) {
	switch t.Kind {
	case model.String:
		if t.Format == "date-time" {
			return "time.Time", nil
		}
		return "string", nil
	case model.Bool:
		return "bool", nil
	case model.Number:
		switch t.Format {
		case "double":
			return "float64", nil
		case "float":
			return "float32", nil
		}
	case model.Integer:
		if t.WireString {
			return "string", nil // D7: the SDK has the JSON string
		}
		switch t.Format {
		case "int32":
			return "int32", nil
		case "int64":
			return "int64", nil
		}
	case model.Enum, model.Object, model.OneOf:
		if t.Schema == "" {
			return "", fmt.Errorf("inline %s schema is not supported", t.Kind)
		}
		return goTypeName(t.Schema), nil
	}
	return "", fmt.Errorf("%s with format %q is not supported", t.Kind, t.Format)
}

// enumConstName is the SDK constant of an enum value: the upper-case type name, "_", and the
// value with "_" between a digit and the letter after it, as openapi-generator writes it
// (E2MType, E2M_TYPE_LOGS2METRICS → E2MTYPE_E2_M_TYPE_LOGS2_METRICS).
func enumConstName(typeName, value string) string {
	var b strings.Builder
	for i, r := range value {
		if i > 0 && unicode.IsLetter(r) && unicode.IsDigit(rune(value[i-1])) {
			b.WriteByte('_')
		}
		b.WriteRune(r)
	}
	return strings.ToUpper(typeName) + "_" + b.String()
}

// goTypeName is the SDK type for a component name: "v3.FilterOperator" → "V3FilterOperator".
func goTypeName(component string) string { return camelize(component) }

// goFieldName is the SDK struct field for a property name: "joinLimit" → "JoinLimit".
func goFieldName(property string) string { return camelize(property) }

// camelize splits s at each character that is not a letter or a digit, and
// joins the parts with the first letter in upper case. It keeps the case of
// the other letters: "AI Evaluations Service" → "AIEvaluationsService".
func camelize(s string) string {
	return model.GoName(s)
}

func lowerFirst(s string) string { return strings.ToLower(s[:1]) + s[1:] }

// ServerSetBodyID returns the id field when Update carries the id in its body and no Update field
// sends it, because only the server sets it. It returns nil otherwise.
func ServerSetBodyID(r *model.Resource) *model.ResourceField {
	if r.Singleton || !r.Policy.UpdateIDInBody {
		return nil
	}
	i := slices.IndexFunc(r.Fields, func(f *model.ResourceField) bool { return f.Name == r.ID() })
	if i < 0 || r.Fields[i].Update != nil {
		return nil
	}
	return r.Fields[i]
}
