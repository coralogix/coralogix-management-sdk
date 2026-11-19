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
	"testing"
)

// The released users contract carries limits the requests lost: an owner removed the request
// limits the backend never enforced (it checks the membership of the login modes, not the
// count), while the response still caps what the server stores. A field whose line says
// releasedLimits keeps its released behavior: the model compares the request types with the
// response type without the limits.

// memberLimitsLostEdits drop the request limit that the response keeps: the loginModes of the
// template lose the maximum of two that the response still declares.
var memberLimitsLostEdits = []string{
	"    MemberTemplate:\n      type: object\n      properties:\n        displayName: {type: string, minLength: 0}\n        login: {type: string, minLength: 0}\n        loginModes:\n          type: array\n          minItems: 0\n          maxItems: 2\n          items: {$ref: '#/components/schemas/LoginMode'}\n",
	"    MemberTemplate:\n      type: object\n      properties:\n        displayName: {type: string, minLength: 0}\n        login: {type: string, minLength: 0}\n        loginModes:\n          type: array\n          minItems: 0\n          items: {$ref: '#/components/schemas/LoginMode'}\n",
}

// The members contract without the releasedLimits lines refuses the fields whose limits the
// requests and the response no longer share.
func TestMemberLimitsDriftIsRefused(t *testing.T) {
	source := memberSource(t, memberLimitsLostEdits)
	memberIssue(t, source, memberPolicy(), "FIELD_TYPE_INCONSISTENT", "different limits")
}

// With the releasedLimits lines, the model compares the types without the limits, and the
// resource builds: the released resource sends what the configuration holds and stores what
// the response echoes, so it enforces no limit itself.
func TestMemberReleasedLimitsHold(t *testing.T) {
	source := memberSource(t, memberLimitsLostEdits)
	p := memberPolicy()
	p.ReleasedLimits = []string{"Member.loginModes"}
	buildMember(t, source, p)
}

// The line covers only the fields it names: the loginModes drift builds with the line, and the
// same drift on login — a field the line does not name — still refuses. A second releasedLimits
// line covers it, and the resource builds.
func TestMemberReleasedLimitsCoverTheItems(t *testing.T) {
	drifted := memberSource(t, memberLimitsLostEdits,
		[]string{
			"        login: {type: string, minLength: 0}\n        loginModes:",
			"        login: {type: string, minLength: 0, maxLength: 100}\n        loginModes:",
		})
	p := memberPolicy()
	p.ReleasedLimits = []string{"Member.loginModes"}
	// login is not in the line: its response-only limit keeps the strict comparison.
	memberIssue(t, drifted, p, "FIELD_TYPE_INCONSISTENT", "different limits")
	// The line names it too: the resource builds.
	p.ReleasedLimits = append(p.ReleasedLimits, "Member.login")
	buildMember(t, drifted, p)
}
