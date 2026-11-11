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
	"slices"
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

func TestCreateResultPolicy(t *testing.T) {
	e := exceptionsOf(t, "api:\n  batchWrite:\n    createResult: {resource: member, status: status, created: [CREATED], message: message}\n")
	if err := e.API.Check(); err != nil {
		t.Fatal(err)
	}
	r := e.Policy().CreateResult
	if r == nil || r.Resource != "member" || r.Status != "status" || r.Message != "message" || !slices.Equal(r.Created, []string{"CREATED"}) {
		t.Fatalf("create result %+v", r)
	}
	for _, line := range []string{"{status: status, created: [CREATED]}", "{resource: member, created: [CREATED]}", "{resource: member, status: status}"} {
		if err := exceptionsOf(t, "api: {batchWrite: {createResult: "+line+"}}\n").API.Check(); err == nil {
			t.Errorf("createResult %s passes the check", line)
		}
	}
}

func TestCreateValuesPolicy(t *testing.T) {
	e := exceptionsOf(t, "api: {requestWrapper: template, createValues: {onboardingMode: ONBOARDING_MODE_NO_INVITE}}\n")
	if err := e.API.Check(); err != nil {
		t.Fatal(err)
	}
	if got := e.Policy().CreateValues; len(got) != 1 || got["onboardingMode"] != "ONBOARDING_MODE_NO_INVITE" {
		t.Fatalf("create values %v", got)
	}
	if err := exceptionsOf(t, "api: {createValues: {onboardingMode: ''}}\n").API.Check(); err == nil {
		t.Fatal("an empty create value passes the check")
	}
}
