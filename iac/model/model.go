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

// Package model is the tool-neutral model of one API resource. It has no
// types of a generated tool and no SDK types. Build reads it from an OpenAPI document.
package model

import "strings"

// OperationIDs selects lifecycle operations when suffix discovery is ambiguous.
// Empty fields use deterministic suffix discovery. An empty List uses the GET on the Create
// path, and a List that the model cannot describe is then no error.
type OperationIDs struct {
	Create string
	Get    string
	Update string
	Delete string
	List   string
}

// Resource is one API resource with Create, Get, Update, and Delete operations.
type Resource struct {
	Name   string // component name of the resource schema
	Create Operation
	Get    Operation
	Update Operation
	Delete Operation
	// IDParam is the path parameter of Get, Update, and Delete. The Get
	// response has a field with the same name. It is "" for a singleton.
	IDParam string
	// IDField is the resource field that holds the id: the field named like IDParam, or the
	// IDProperty of the policy. It is "" for a singleton.
	IDField string
	// IDType is the type shared by the path parameter and response ID field.
	// It is nil for a singleton.
	IDType *Type
	// Singleton is true when Get has no path parameter: there is one
	// resource per company (D18). Create, Get, Update, and Delete use one
	// path, except for an Overwrite resource.
	Singleton bool
	// Overwrite is true when the resource is the whole list that one write operation replaces
	// (Policy.Overwrite). Create, Update, and Delete are that write; Delete sends every field
	// empty. The write response is not read: a tool reads the resource with Get after the write,
	// unless WriteReturnsResource is set. An Overwrite resource is a singleton.
	Overwrite bool
	// WriteReturnsResource is true when the write of an Overwrite resource returns the resource
	// (Policy.WriteReturnsResource): the response of Create and Update has every field of the
	// resource under its name, with the same shape. A tool takes the state from it, with no Get.
	WriteReturnsResource bool
	// GetFromList is true when Get is the List (Policy.GetFromList): no operation reads one resource by
	// its id. A tool pages List with GetQuery, optionally narrowed by ListFilters, and selects the
	// item whose IDField is the id; no such item means that the resource is missing. Get.Response holds the
	// items (Items) in List.ItemsField.
	GetFromList bool
	// ListFilters are the query parameters that narrow the List by a resource field, sorted by
	// parameter (Policy.ListFilters). A filter can match more than the resource.
	ListFilters []ListFilter
	// PageParameters are the query parameters that page the Get of a GetFromList resource, in the
	// order of the contract, and NextPage is the response property that holds the token of the next
	// page, nil when the response has none. A parameter or property whose type the model cannot read,
	// such as a pagination object, is left out.
	PageParameters []Parameter
	NextPage       *Parameter
	// DeleteValues are the properties of the Delete body that a remove always sends beside the id
	// array Policy.DeleteIDs, with their values, sorted by name (Policy.DeleteValues).
	DeleteValues []QueryValue
	// GetQuery are the query parameters that Get always sends, with their values, sorted by name
	// (Policy.GetQuery). It is nil when Get sends none.
	GetQuery []QueryValue
	// CreateValues are the properties beside the request wrapper that Create always sends, with their
	// values, sorted by name (Policy.CreateValues). They are not resource fields.
	CreateValues []QueryValue
	// Replace is true when Update is a full replace (PUT, E11): the body has
	// every Update field, and the server clears a field that the body does
	// not have. It has no update mask. False: PATCH with an update mask.
	Replace bool
	// UpdateMask is the update-mask name: the query parameter, or the JSON body
	// property when MaskInBody is set. It is "" for a full replace. It is not
	// a resource field.
	UpdateMask string
	// UpdateMaskPattern is the "pattern" of the update mask string, "" when
	// the spec has none. It shows which mask paths the API accepts.
	UpdateMaskPattern string
	// MaskInBody: PATCH updateMask is an optional JSON body property. Update
	// sets it to the same changed-field mask a query parameter would send.
	MaskInBody bool
	// UpdateIDOutsideWrapper is true when the Update body carries the id property beside the request
	// wrapper, not inside it (Policy.UpdateIDOutsideWrapper). The id is then a field of the
	// resource that Update sends, and a tool places it next to the wrapper property.
	UpdateIDOutsideWrapper bool
	// Policy is the rule set that Build used. It is the zero value for a new resource.
	Policy Policy
	// Groups are the oneOf groups among the top-level fields.
	Groups []OneOfGroup
	// Fields are the top-level resource fields: the Get fields in spec order,
	// then Create-only and Update-only fields (Build rejects those).
	Fields []*ResourceField
	// UnwrapUsed are the unwrap lines of the policy that decided a place: "unwrap.<Component>"
	// for a component of the list, and "<Component>.<field>" for a field line.
	UnwrapUsed map[string]bool
	// HiddenWrappers are the wrapper components that collapse in every place, so the generated
	// code has no attribute for their property.
	HiddenWrappers map[string]bool
	// List is the operation that lists the resources. It is nil when the API has none, or has
	// one that the model cannot describe.
	List *ListOperation
	// shapeUsed collects the nested request lines that Build applied (Policy.shapeNested).
	shapeUsed map[string]bool
	// requestUnwrapUsed collects the unwrap lines that collapsed a wrapper that only a request holds.
	requestUnwrapUsed map[string]bool
	// ListProblem says in one sentence why List is nil although the Create path has a GET. It
	// is "" when List is set, or the Create path has no GET. A List that OperationIDs names
	// has no problem: Build fails instead.
	ListProblem string
}

// ID returns the resource field that holds the id: IDField, or IDParam for a resource that Build
// did not make. It is "" for a singleton.
func (r *Resource) ID() string {
	if r.IDField != "" {
		return r.IDField
	}
	return r.IDParam
}

// Operation is one HTTP operation of the resource.
type Operation struct {
	Method      string // upper case, for example "POST"
	Path        string
	OperationID string
	Body        string // "" (no body), "inline", or a component name
	// BodyItems: the body is an array of the component Body (Policy.BatchWrite), and a tool sends one
	// item, for example []CreateThingRequest.
	BodyItems bool
	// BodyTitle is the "title" of an inline body, "" when it has none.
	// BodyTitleIsComponent: a component schema has the same name.
	BodyTitle            string
	BodyTitleIsComponent bool
	// BodyAsCreate: the Update body is an inline schema without a title that is the same as the
	// inline Create body. openapi-generator makes one type of the two, with the name of the Create
	// body, so the Update builder takes that type.
	BodyAsCreate bool
	// RequestComponent is the component schema of the request wrapper property (Policy.RequestWrapper)
	// of a Create or Update body, for example "IntegrationMetadata". It is "" without a wrapper, and
	// when the wrapper property has an inline schema. It is the type that the SDK body holds: the
	// resource component only when the request has the shape of the resource.
	RequestComponent string
	Response         Response
}

// QueryValue is a query parameter or a body property that an operation always sends, and its value.
type QueryValue struct {
	Name  string // the parameter or property name, for example "source_type"
	Value string // the value as the query string holds it, for example "SOURCE_TYPE_LOGS"
	Type  *Type  // the type of the parameter: an Enum names its component
}

// ListFilter is a query parameter of the List that narrows it by a resource field.
type ListFilter struct {
	Param string // the query parameter, for example "username"
	Field string // the resource field that it matches
	Type  *Type  // the type of the parameter, the scalar type of the field
}

// ListOperation is a GET that returns the resources in one array. The model records the facts
// about it. A tool decides what to generate from them.
type ListOperation struct {
	Operation
	// ItemsField is the response property that holds the items, for example "connectors".
	ItemsField string
	// ItemComponent is the component of an item, for example "Connector" or "PresetSummary".
	ItemComponent string
	// Full is true when the item component is the resource component, so an item has every
	// field of a Get. False: an item is a summary, and a tool calls Get for the other fields.
	Full bool
	// IDField is the item property that holds the resource id. It is the IDField of the
	// resource, and it has the type of the resource id.
	IDField string
	// NameField is "name" when the item has a string property name, and "" otherwise.
	NameField string
	// Paged is true when the List has a query parameter that pages the result (a page token,
	// size, or offset, a limit, or a pagination object), or the response has a next page token.
	Paged bool
}

// Parameter is a query parameter or a response property, and its type.
type Parameter struct {
	Name string
	Type *Type
}

// Response is the 200 response body.
type Response struct {
	// Items: Field holds an array of resources. GetByIDs reads one id through this batch API, and the
	// Update of a BatchWrite resource returns the item that it wrote.
	Items  bool
	Schema string // component name, "" for an inline empty object
	// Field is the property that wraps the resource, for example "aiEvaluation".
	// It is empty when the response does not return the resource, or is the
	// resource itself (Direct).
	Field string
	// Direct is true when the response is the resource itself, with no
	// wrapper (proto google.api.http response_body).
	Direct bool
	// ID is the property that holds only the id of the resource, for a Create response that does
	// not return the resource (an existing resource). It is a string, or the same wrapper of the id
	// as the id field of the resource has, and Create reads the resource with Get.
	ID string
	// idType is the type of ID after the wrapper that the policy unwraps collapses, or nil when ID
	// holds a string id itself.
	idType *Type
	// Empty is true when the response schema has no fields, for example many
	// Delete responses. The SDK then returns map[string]interface{}.
	Empty bool
	// Result is the answer of a batch Create that holds one result per request item
	// (Policy.CreateResult), nil otherwise. Field is then "".
	Result *WriteResult
}

// WriteResult is the answer of a batch Create: one result per request item, in request order. A tool
// sends one item and reads the first result. The item is created only when Status holds one of
// Created, and Resource then holds it. Otherwise Message, when the result has it, says why.
type WriteResult struct {
	Items     string // the response property that holds the results, for example "results"
	Component string // the component of a result, for example "CreateThingResult"
	Resource  string // the result property that holds the resource
	Status    string // the enum property with the status of the item
	// StatusType is the enum of Status. Created are its values that mean created, in the order of
	// the line; none is the zero value.
	StatusType *Type
	Created    []string
	Message    string // the string property with a message, "" when the line names none
}

// Behavior is how a top-level field is managed. See Classify.
type Behavior string

const (
	Normal    Behavior = "normal"    // Create, Update, Get
	Immutable Behavior = "immutable" // Create, Get: a change needs a new resource
	Computed  Behavior = "computed"  // Get only: read and store, never send
)

// ResourceField is a top-level field of the resource.
type ResourceField struct {
	Name        string
	Description string // from the Get schema
	Type        *Type
	Behavior    Behavior
	Create      *Attrs // nil: not in the Create body
	Update      *Attrs // nil: not in the Update body
	Get         *Attrs // attributes in the resource response
	InGet       bool
	// RequestPath is the dot-separated path of the property that holds the field in the Create and
	// Update bodies, after the request wrapper, when Policy.RequestNames gives the field another
	// name there. "" means that the requests use the name of the field.
	RequestPath string
	// RequestObjects are the component schemas of the objects that RequestPath goes through, one
	// for each segment before the last: "integrationParameters.parameters" has the component of
	// integrationParameters. It is nil for a path of one segment, and when a segment has an inline
	// schema.
	RequestObjects []string
	// UpdateRequestPath and UpdateRequestObjects are the same for the Update body, when it holds the
	// field at another path than the Create body (an updateAs line). "" means that the Update body
	// uses RequestPath.
	UpdateRequestPath    string
	UpdateRequestObjects []string
}

// RequestPathOf returns the dot-separated path of the property that holds the field in the Update
// body (update) or in the Create body, after the request wrapper. "" means the field name.
func (f *ResourceField) RequestPathOf(update bool) string {
	if update && f.UpdateRequestPath != "" {
		return f.UpdateRequestPath
	}
	return f.RequestPath
}

// RequestObjectsOf returns the components of the objects that RequestPathOf goes through.
func (f *ResourceField) RequestObjectsOf(update bool) []string {
	if update && f.UpdateRequestPath != "" {
		return f.UpdateRequestObjects
	}
	return f.RequestObjects
}

// Field is a field of a nested object, or an arm of a oneOf.
type Field struct {
	Name        string
	Description string
	Attrs       Attrs // attributes in the resource response
	Type        *Type
	// Behavior is how the field is managed inside its parent object, as for a
	// top-level field. It is "" inside a computed value: the whole value is
	// computed.
	Behavior Behavior
	// Create and Update are the attributes of the field in the request type of
	// its parent object. nil: that request does not have the field.
	Create, Update *Attrs
	// RequestPath is the path of the property that holds the field in the request object of its
	// parent, when it is not the field name: for example ["policy", "name"] when the request item
	// holds the field in a wrapper (Policy.RequestWrappers), or ["policy", "disabled"] after a
	// requestAs line. nil means the field name.
	RequestPath []string
	// RequestObjects are the components of the request objects that RequestPath goes through,
	// outer first, for example ["CreateGenericPolicyRequest"]. A tool needs them to build the request.
	RequestObjects []string
	// Negated: the request sends the opposite of the boolean field (Policy.RequestNegated), for
	// example disabled for enabled.
	Negated bool
}

// Attrs are the facts about a field in one place (a request body or an object).
type Attrs struct {
	// Computed: a computed: true line keeps the value of the field in state. With a Get that
	// requires the field, state always has a value that Update can send.
	Computed bool
	// KeepState: a line keeps the state value for an unknown plan value (useStateForUnknown or
	// useNonNullStateForUnknown).
	KeepState bool
	Required  bool    // in the parent "required" list
	Presence  bool    // x-coralogix-presence: true
	Default   *string // YAML text of "default", nil when there is none
	// ReadOnly: readOnly: true. In a request, the server sets the field, so the request does not send it.
	ReadOnly bool
}

// Kind is the kind of value a Type holds.
type Kind string

const (
	String  Kind = "string"
	Bool    Kind = "bool"
	Number  Kind = "number"
	Integer Kind = "integer"
	Enum    Kind = "enum"
	Object  Kind = "object"
	OneOf   Kind = "oneOf" // an object with exactly one arm set
	List    Kind = "list"  // ordered
	Set     Kind = "set"   // unordered, unique items
	Map     Kind = "map"   // string keys; Elem is the value type
)

// OneOfGroup is a set of fields of an object of which at most one is set.
type OneOfGroup struct {
	Arms []string // field names
	// AllowNone: the value can have no arm set. When Create sends the value,
	// it is the Create rule, which the client configuration follows.
	AllowNone bool
}

// Type is the type of a field.
type Type struct {
	Kind Kind
	// WireKind is the collection kind used by this operation's API schema. Kind is
	// the consumer-facing kind and can differ when an explicit collection: set
	// override normalizes a list or a map of objects.
	WireKind Kind
	// CreateWireKind and UpdateWireKind preserve the corresponding request shape
	// on the resource type so requestType can expose it to generators.
	CreateWireKind, UpdateWireKind Kind
	Schema                         string // component name, "" for an inline schema
	// CreateSchema and UpdateSchema are the components of an Object or a
	// OneOf in the Create and Update requests. They can differ from Schema,
	// the component in the resource response. "" when the request does not
	// send the value.
	CreateSchema, UpdateSchema string
	// Model is the component of the resource response that holds the
	// generated model of a request type. It is set only on the types that
	// CreateType and UpdateType return.
	Model  string
	Format string // for example "double", "uint64", "date-time"
	// WireString is true for an integer that JSON sends as a string
	// (protobuf 64-bit numbers).
	WireString bool
	// DigitString is "uint64" or "int64" for a string without a format whose pattern is the one
	// that protoc-gen-openapiv3 writes for a 64-bit number (^[0-9]+$ or ^-?[0-9]+$), and "" for
	// any other string. The pattern cannot tell such a number from a string that a proto author
	// validates as digits (a StringValue company id has the same schema), so the type stays a
	// String. A tool can read it as a number where its own policy names the field.
	DigitString string
	Values      []string // Enum: the business values, without EnumZero
	EnumZero    string   // Enum: the exact protobuf zero value
	Elem        *Type    // List, Set, Map
	Fields      []*Field // Object: the properties; OneOf: the arms
	AllowNone   bool     // OneOf: the value can have no arm set, as in OneOfGroup
	// Groups are the oneOf groups of an Object that also has normal fields,
	// or has more than one group. Each arm is one of Fields.
	Groups []OneOfGroup
	// Discriminator is the string field that names the set arm (an OpenAPI
	// discriminator with no mapping). It is also a normal field (F36).
	Discriminator string

	// Zero is the protobuf zero value of a scalar, as the text that a default holds: "false",
	// "0", or "". A 64-bit number that the API writes as a string has the zero "0".
	Zero string
	// ZeroValid: Zero meets the limits and the pattern of the contract. It is false for an
	// enum, a time, and a value that is not a scalar.
	ZeroValid bool
	// ReadZeroAsNull: a request sends the scalar without presence or a default, and its zero value
	// is not valid, so a configuration never holds it. A read of the zero value is an omitted one.
	ReadZeroAsNull bool
	// ImplicitDefault: the contract declares no default. The default is the zero value that a
	// request scalar without presence reads back as. The service does not fill an omitted object
	// that holds it.
	ImplicitDefault bool

	MinLength, MaxLength *int64
	// Pattern is the regular expression that a String must match, "" for
	// none. The free-text pattern is none. Existing resources have none.
	Pattern            string
	Minimum, Maximum   *float64
	MinItems, MaxItems *int64 // List, Set: items. Map: entries (minProperties, maxProperties).
	// requestMinItems: a request's item minimum is merged into MinItems. The first request may
	// require more items than the response declares; every later request must require as many.
	requestMinItems bool
	// Wrappers are the one-property objects that hold the value in the API, outer first. The
	// rest of the type is the value inside them, which the generated code shows. They are set only by
	// the unwrap lines of a behavior-overrides file.
	Wrappers []Wrapper
	// RequestWrapper is a wrapper that only a request holds, set on a top-level field by an unwrap
	// entry: Create and Get hold the value, and Update holds it in a one-property object. Only its
	// CreateSchema or UpdateSchema is set. The resource response view does not have it; the request
	// view of that request has it as its outermost wrapper (CreateType and UpdateType).
	RequestWrapper *Wrapper
	// Projected is set on a List or Set of strings that the API holds as a list of objects in the
	// resource response, and that the requests send as the strings: an itemName line. Elem is the
	// type of the property that the resource shows. It is nil in the Create and Update request types,
	// whose items are the strings.
	Projected *Projection
	// Keyed is set on a Map that the API sends as a list of objects, keyed by a property of the
	// object (a keyedBy line). Elem is the object without the key, or the value property. On an
	// Object (keyedBy object: true), each field is the item of one value of the enum key.
	Keyed *Keyed
	// Variant tells apart the items that one component makes for a keyedBy object: the Go name
	// of the attribute that holds the item of each value.
	Variant string
}

// Projection is the object that holds each item of a list in the resource response, when the
// resource shows one string property of it.
type Projection struct {
	// Schema is the component of the item in the resource response, and Property the string
	// property that the resource shows.
	Schema, Property string
	// Required says whether the response requires Property in the item.
	Required bool
}

// Keyed is the list of objects that the API holds for a map of the generated resource.
type Keyed struct {
	// Key is the string property of Item that is the map key.
	Key string
	// Value is the property of Item that is the map value, "" when the value is the rest of Item.
	Value string
	// Item is the object of the API list, with every property.
	Item *Type
	// Object: the generated value is an object with an attribute per value of the enum key.
	// Values are those values. KeyAs is the string property that expand sets to the attribute.
	Object bool
	Values []KeyedValue
	KeyAs  string
}

// KeyedValue is an API value of the key of a keyedBy object and the field that holds its item.
type KeyedValue struct {
	API, Field string
}

// Wrapper is one object with one property that holds a value in the API, for example
// LuceneQuery in {"luceneQuery": {"value": "error"}}. The generated code shows only the value.
type Wrapper struct {
	// Schema is the component in the resource response. CreateSchema and UpdateSchema are the
	// components in the requests, "" when that request does not send the value.
	Schema, CreateSchema, UpdateSchema string
	// Field is the one property that holds the value.
	Field string
	// Attrs are the attributes of Field in the resource response. Create and Update are its
	// attributes in the requests, nil when that request does not send the value.
	Attrs          Attrs
	Create, Update *Attrs
}

// WrapperPath returns the API path from the outer wrapper to the value, for example
// "filter.value", or "" when the value has no wrapper.
func (t *Type) WrapperPath() string {
	names := make([]string, 0, len(t.Wrappers))
	for _, w := range t.Wrappers {
		names = append(names, w.Field)
	}
	return strings.Join(names, ".")
}

// Unwrapped returns the type of the value inside the wrappers. It is t when t has none.
func (t *Type) Unwrapped() *Type {
	if len(t.Wrappers) == 0 {
		return t
	}
	out := *t
	out.Wrappers = nil
	return &out
}
