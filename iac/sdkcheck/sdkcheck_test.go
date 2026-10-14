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

package sdkcheck

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/coralogix/coralogix-management-sdk/iac/issue"
	"github.com/coralogix/coralogix-management-sdk/iac/model"
)

func thingNames(t *testing.T, zeroConstants func(string) bool) []Ref {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", "model", "testdata", "openapi.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := model.Load(data)
	if err != nil {
		t.Fatal(err)
	}
	r, err := model.Build(doc, "Thing")
	if err != nil {
		t.Fatal(err)
	}
	tag, err := ResourceTag(doc, r)
	if err != nil {
		t.Fatal(err)
	}
	if tag != "Things Service" || PackageName(tag) != "things_service" || ClientType(tag) != "ThingsServiceAPIService" {
		t.Fatalf("tag %q, package %q, client %q", tag, PackageName(tag), ClientType(tag))
	}
	refs, err := Names(r, tag, "example.com/sdk", zeroConstants)
	if err != nil {
		t.Fatal(err)
	}
	return refs
}

// Names lists only names of the SDK module: the generated package of the tag and the cxsdk
// package. A tool appends the names of its own code.
func TestNamesListsOnlySDKNames(t *testing.T) {
	refs := thingNames(t, nil)
	gen, cxsdk := "example.com/sdk/go/openapi/gen/things_service", "example.com/sdk/go/openapi/cxsdk"
	for _, ref := range refs {
		if ref.Pkg != gen && ref.Pkg != cxsdk {
			t.Errorf("ref %s %s is in package %s", ref.Kind, ref.SDKName(), ref.Pkg)
		}
	}
	for _, want := range []string{"ThingsServiceAPIService.ThingsServiceGetThing", "Thing.Name", "IsNotFound"} {
		if !slices.ContainsFunc(refs, func(ref Ref) bool { return ref.SDKName() == want }) {
			t.Errorf("refs do not contain %s", want)
		}
	}
}

// The zero constant of an enum is in the list only when the caller uses it.
func TestNamesListsZeroConstantsOnRequest(t *testing.T) {
	zero := func(refs []Ref) bool {
		return slices.ContainsFunc(refs, func(ref Ref) bool { return ref.Name == "THINGKIND_THING_KIND_UNSPECIFIED" })
	}
	if zero(thingNames(t, nil)) {
		t.Fatal("nil zeroConstants lists the zero constant")
	}
	if !zero(thingNames(t, func(component string) bool { return component == "ThingKind" })) {
		t.Fatal("zeroConstants for ThingKind does not list its zero constant")
	}
}

func TestIssuesChecksNamesAndSources(t *testing.T) {
	const pkg = "example.com/fakesdk/widgets"
	refs := []Ref{
		{Path: "resource", Kind: KindPackage, Pkg: pkg, Name: "widgets"},
		{Path: "fields", Kind: KindType, Pkg: pkg, Name: "Widget"},
		{Path: "fields.name", Kind: KindField, Pkg: pkg, Owner: "Widget", Name: "Name", Want: "*string"},
		{Path: "fields.count", Kind: KindField, Pkg: pkg, Owner: "Widget", Name: "Count", Want: "*int32", WantValue: "int32"},
		{Path: "get", Kind: KindMethod, Pkg: pkg, Owner: "WidgetsAPIService", Name: "Widget", Want: "func() string", ByType: true},
		{Path: "fields.size", Kind: KindField, Pkg: pkg, Owner: "Widget", Name: "Size", Want: "*int64"},
	}
	loaded, err := Load(refs, filepath.Join("testdata", "fakesdk"))
	if err != nil {
		t.Fatal(err)
	}
	report := Issues(refs, loaded, func(Ref) Source { return Source{Module: "example.com/fakesdk"} })
	if codes := codesOf(report); !slices.Equal(codes, []string{"SDK_SYMBOL_MISSING"}) || report[0].Location != "fields.size" {
		t.Fatalf("report = %v", report)
	}
	// The SDK has the value type and another method name: the refs now name what the SDK has.
	if refs[3].Want != "int32" || refs[4].Name != "GetWidget" {
		t.Fatalf("refs = %+v, %+v", refs[3], refs[4])
	}

	report = Issues(refs[:1], loaded, func(Ref) Source { return Source{Module: "example.com/other", Remediation: "Load the other module."} })
	if codes := codesOf(report); !slices.Equal(codes, []string{"SDK_SOURCE_MISMATCH"}) || report[0].Remediation != "Load the other module." {
		t.Fatalf("report = %v", report)
	}
	if !strings.Contains(report[0].Message, "expected example.com/other") {
		t.Fatalf("message = %s", report[0].Message)
	}
}

func codesOf(report issue.Report) []string {
	var codes []string
	for _, item := range report {
		codes = append(codes, item.Code)
	}
	return codes
}
