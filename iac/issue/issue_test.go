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

package issue

import (
	"strings"
	"testing"
)

func TestReportIsDeduplicatedAndSorted(t *testing.T) {
	a := Issue{Code: "B", Location: "z", Message: "second", Remediation: "fix"}
	b := Issue{Code: "A", Location: "a", Message: "first", Remediation: "fix"}
	report := Report{a, b, a}.Normalize()
	if len(report) != 2 || report[0] != b || report[1] != a {
		t.Fatalf("normalized report = %#v", report)
	}
	text := report.Error()
	if !strings.Contains(text, "2 problems") || strings.Index(text, "A a") > strings.Index(text, "B z") {
		t.Fatalf("unexpected report:\n%s", text)
	}
}
