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

package overrides

import (
	"testing"

	"go.yaml.in/yaml/v4"
)

func exceptionsOf(t *testing.T, source string) Exceptions {
	t.Helper()
	var e Exceptions
	if err := yaml.Unmarshal([]byte(source), &e); err != nil {
		t.Fatal(err)
	}
	return e
}

func TestBatchWritePolicy(t *testing.T) {
	e := exceptionsOf(t, "api: {batchWrite: {}, requestWrapper: template}\n")
	if err := e.API.Check(); err != nil {
		t.Fatal(err)
	}
	if p := e.Policy(); !p.BatchWrite || !p.Existing {
		t.Fatalf("policy %+v, want an existing batch write", p)
	}
	if p := exceptionsOf(t, "api: {requestWrapper: template}\n").Policy(); p.BatchWrite {
		t.Fatal("batchWrite is set without its key")
	}
	overwrite := exceptionsOf(t, "api: {batchWrite: {}, overwrite: {}, operations: {create: W, get: G, update: W, delete: W}}\n")
	if err := overwrite.API.Check(); err == nil {
		t.Fatal("api.overwrite goes with api.batchWrite")
	}
}
