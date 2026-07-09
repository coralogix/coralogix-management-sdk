//go:build !jsonhcl

// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEnsureTerraformResourceName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		display  string
		id       string
		expected string
	}{
		{
			name:     "uses sanitized display name",
			display:  "My Dashboard",
			id:       "abc-123",
			expected: "my_dashboard",
		},
		{
			name:     "falls back to id for empty display name",
			display:  "",
			id:       "abc-123-def",
			expected: "abc-123-def",
		},
		{
			name:     "falls back to id when display name sanitizes to empty",
			display:  "---",
			id:       "dashboard-id",
			expected: "dashboard-id",
		},
		{
			name:     "uses unnamed when both are empty",
			display:  "",
			id:       "",
			expected: "unnamed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := ensureTerraformResourceName(convertToTerraformResourceName(tt.display), tt.id); got != tt.expected {
				t.Fatalf("ensureTerraformResourceName() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestGenerateImportsFromIdsEmptyDashboardName(t *testing.T) {
	t.Parallel()

	outputPath := filepath.Join(t.TempDir(), "imports.tf")
	err := generateImportsFromIds("dashboard", outputPath, []IdAndName{
		{Id: "dashboard-id-1", Name: ""},
		{Id: "dashboard-id-2", Name: "valid_name"},
	})
	if err != nil {
		t.Fatalf("generateImportsFromIds() error = %v", err)
	}

	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("failed to read imports.tf: %v", err)
	}

	imports := string(content)
	if strings.Contains(imports, "coralogix_dashboard.\n") || strings.Contains(imports, "coralogix_dashboard. ") {
		t.Fatalf("generated invalid import address for empty dashboard name:\n%s", imports)
	}
	if !strings.Contains(imports, `to = coralogix_dashboard.dashboard-id-1`) {
		t.Fatalf("expected fallback resource name from id, got:\n%s", imports)
	}
	if !strings.Contains(imports, `to = coralogix_dashboard.valid_name`) {
		t.Fatalf("expected named dashboard import, got:\n%s", imports)
	}
}
