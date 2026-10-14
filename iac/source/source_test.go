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

package source

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveUsesPinnedSDK(t *testing.T) {
	root := t.TempDir()
	version := "v1.2.3"
	toolMod := "module example.com/tool\n\ngo 1.26.0\n\nrequire " + SDKModule + " " + version + "\n"
	writeTestFile(t, filepath.Join(root, "go.mod"), toolMod)
	start := filepath.Join(root, "tools", "codegen")
	writeTestFile(t, filepath.Join(start, "go.mod"), "module example.com/tool/tools/codegen\n\ngo 1.26.0\n")
	cache := filepath.Join(root, "module-cache")
	openAPI := []byte("openapi: 3.0.3\n")
	writeTestFile(t, filepath.Join(cache, SDKModule+"@"+version, "openapi.yaml"), string(openAPI))
	t.Setenv("GOMODCACHE", cache)

	input, err := Resolve(start)
	if err != nil {
		t.Fatal(err)
	}
	if input.ModuleRoot != root || input.Module != "example.com/tool" || input.SDKVersion != version || string(input.OpenAPI) != string(openAPI) {
		t.Fatalf("resolved %+v", input)
	}
}

// A tool module that requires only the iac module of the SDK does not pin the SDK, also
// though its go.mod contains the SDK module path. Resolve goes on to the module that pins it.
func TestResolveSkipsModuleThatRequiresOnlyIAC(t *testing.T) {
	root := t.TempDir()
	version := "v1.2.3"
	writeTestFile(t, filepath.Join(root, "go.mod"), "module example.com/tool\n\ngo 1.26.0\n\nrequire "+SDKModule+" "+version+"\n")
	start := filepath.Join(root, "tools", "codegen")
	writeTestFile(t, filepath.Join(start, "go.mod"), "module example.com/tool/tools/codegen\n\ngo 1.26.0\n\nrequire "+SDKModule+"/iac v0.0.0-20261020101500-1a2b3c4d5e6f\n")
	cache := filepath.Join(root, "module-cache")
	writeTestFile(t, filepath.Join(cache, SDKModule+"@"+version, "openapi.yaml"), "openapi: 3.0.3\n")
	t.Setenv("GOMODCACHE", cache)

	input, err := Resolve(start)
	if err != nil {
		t.Fatal(err)
	}
	if input.ModuleRoot != root || input.SDKVersion != version {
		t.Fatalf("resolved %+v", input)
	}
}

func TestResolveDoesNotFetchMissingSDK(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, filepath.Join(root, "go.mod"), "module example.com/tool\n\ngo 1.26.0\n\nrequire "+SDKModule+" v1.2.3\n")
	t.Setenv("GOMODCACHE", filepath.Join(root, "empty-cache"))
	_, err := Resolve(root)
	if err == nil {
		t.Fatal("expected missing bundled OpenAPI error")
	}
}

func writeTestFile(t *testing.T, path, data string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}
}
