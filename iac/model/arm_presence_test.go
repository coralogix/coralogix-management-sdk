package model

import (
	"strings"
	"testing"

	"go.yaml.in/yaml/v4"
)

// A scalar arm of a oneOf is marked with x-coralogix-presence by the OpenAPI fork. The mark states
// what the arm is already, so it is no attribute of the arm. Any other attribute still is.
func TestOneOfArmMayCarryPresenceOnly(t *testing.T) {
	doc := loadComponent(t, `
    Choice:
      type: object
      properties:
        a: {type: boolean}
        b: {type: string}
      oneOf:
        - required: [a]
        - required: [b]
`)
	schema, err := schemaOf(doc.Components.Schemas.GetOrZero("Choice"))
	if err != nil {
		t.Fatal(err)
	}
	fields := func(a, b Attrs) []*Field {
		return []*Field{{Name: "a", Attrs: a}, {Name: "b", Attrs: b}}
	}
	if _, err := oneOfGroups(schema, "choice", fields(Attrs{Presence: true}, Attrs{Presence: true})); err != nil {
		t.Fatalf("presence on an arm must be accepted: %v", err)
	}
	for name, attrs := range map[string]Attrs{
		"required": {Presence: true, Required: true},
		"computed": {Computed: true},
		"readOnly": {ReadOnly: true},
	} {
		_, err := oneOfGroups(schema, "choice", fields(attrs, Attrs{}))
		if err == nil || !strings.Contains(err.Error(), "has attributes") {
			t.Fatalf("%s on an arm must still be refused, got %v", name, err)
		}
	}
	def := "true"
	if _, err := oneOfGroups(schema, "choice", fields(Attrs{Default: &def}, Attrs{})); err == nil {
		t.Fatal("a default on an arm must still be refused")
	}
}

// An enum of an existing resource may have one value, which the proto names _UNSPECIFIED because it
// is the first. A new resource still needs a business value next to the zero value.
func TestEnumWithOneValue(t *testing.T) {
	values := func(names ...string) []*yaml.Node {
		var out []*yaml.Node
		for _, n := range names {
			out = append(out, &yaml.Node{Kind: yaml.ScalarNode, Value: n})
		}
		return out
	}
	one := values("THING_KIND_MORE_THAN_OR_UNSPECIFIED")
	if _, err := enumZero(one, false); err == nil || !strings.Contains(err.Error(), "at least one business value") {
		t.Fatalf("a new resource needs a business value, got %v", err)
	}
	zero, err := enumZero(one, true)
	if err != nil || zero != "THING_KIND_MORE_THAN_OR_UNSPECIFIED" {
		t.Fatalf("an existing resource has the one value as its zero value: zero %q, err %v", zero, err)
	}
	// A one-value enum whose value is no _UNSPECIFIED stays a business value with no zero.
	if zero, err := enumZero(values("THING_KIND_AUTO"), true); err != nil || zero != "" {
		t.Fatalf("zero %q, err %v", zero, err)
	}
	// Two values still follow the exact rule, with or without the existing mode.
	if zero, err := enumZero(values("THING_KIND_UNSPECIFIED", "THING_KIND_A"), true); err != nil || zero != "THING_KIND_UNSPECIFIED" {
		t.Fatalf("zero %q, err %v", zero, err)
	}
}
