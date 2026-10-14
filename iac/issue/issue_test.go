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

func TestReportTextNamesTheTool(t *testing.T) {
	report := Report{{Code: "A", Location: "a", Message: "first", Remediation: "fix"}}
	want := "Resource is not eligible for X generation: 1 problems\n\nA a\n  first\n  fix"
	if got := report.Text("Resource is not eligible for X generation"); got != want {
		t.Fatalf("Text() =\n%s\nwant:\n%s", got, want)
	}
	if got := report.Error(); !strings.HasPrefix(got, "Resource is not eligible for generation: 1 problems\n") {
		t.Fatalf("Error() =\n%s", got)
	}
	if got := (Report{}).Text("heading"); got != "" {
		t.Fatalf("empty report Text() = %q, want empty", got)
	}
}
