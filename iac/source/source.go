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

// Package source resolves the OpenAPI and Go SDK that a tool's Go module pins.
package source

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
)

// SDKModule is the Go module of the management SDK. Its module root has the openapi.yaml
// that the SDK was generated from.
const SDKModule = "github.com/coralogix/coralogix-management-sdk"

// Input is one matched OpenAPI and SDK source.
type Input struct {
	// ModuleRoot is the directory of the go.mod that pins the SDK, and Module is its module path.
	ModuleRoot string
	Module     string
	SDKDir     string
	SDKModule  string
	SDKVersion string
	OpenAPI    []byte
}

// Resolve walks from start to the go.mod that pins the SDK and resolves its exact SDK version.
func Resolve(start string) (Input, error) {
	root, data, err := findModule(start)
	if err != nil {
		return Input{}, err
	}
	parsed, err := modfile.Parse(filepath.Join(root, "go.mod"), data, nil)
	if err != nil {
		return Input{}, fmt.Errorf("parse %s: %w", filepath.Join(root, "go.mod"), err)
	}
	if parsed.Module == nil || parsed.Module.Mod.Path == "" {
		return Input{}, fmt.Errorf("%s must declare a module path", filepath.Join(root, "go.mod"))
	}
	var versions []string
	for _, req := range parsed.Require {
		if req.Mod.Path == SDKModule {
			versions = append(versions, req.Mod.Version)
		}
	}
	if len(versions) != 1 || versions[0] == "" {
		return Input{}, fmt.Errorf("%s must pin exactly one %s version; found %v", filepath.Join(root, "go.mod"), SDKModule, versions)
	}
	cache, err := moduleCache()
	if err != nil {
		return Input{}, err
	}
	escapedPath, err := module.EscapePath(SDKModule)
	if err != nil {
		return Input{}, fmt.Errorf("escape SDK module path: %w", err)
	}
	escapedVersion, err := module.EscapeVersion(versions[0])
	if err != nil {
		return Input{}, fmt.Errorf("escape SDK module version: %w", err)
	}
	sdkDir := filepath.Join(cache, escapedPath+"@"+escapedVersion)
	openAPIPath := filepath.Join(sdkDir, "openapi.yaml")
	openAPI, err := os.ReadFile(openAPIPath)
	if err != nil {
		return Input{}, fmt.Errorf("read OpenAPI for pinned SDK %s from %s: %w; download the pinned module before generation", versions[0], openAPIPath, err)
	}
	return Input{ModuleRoot: root, Module: parsed.Module.Mod.Path, SDKDir: sdkDir, SDKModule: SDKModule, SDKVersion: versions[0], OpenAPI: openAPI}, nil
}

func findModule(start string) (string, []byte, error) {
	dir, err := filepath.Abs(start)
	if err != nil {
		return "", nil, fmt.Errorf("resolve start directory: %w", err)
	}
	for {
		path := filepath.Join(dir, "go.mod")
		data, readErr := os.ReadFile(path)
		if readErr == nil && strings.Contains(string(data), SDKModule) {
			return dir, data, nil
		}
		if readErr != nil && !errors.Is(readErr, os.ErrNotExist) {
			return "", nil, fmt.Errorf("read %s: %w", path, readErr)
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", nil, fmt.Errorf("find the module root from %s: no ancestor go.mod requires %s", start, SDKModule)
		}
		dir = parent
	}
}

func moduleCache() (string, error) {
	if value := os.Getenv("GOMODCACHE"); value != "" {
		return value, nil
	}
	out, err := exec.Command("go", "env", "GOMODCACHE").Output()
	if err != nil {
		return "", fmt.Errorf("find Go module cache: %w", err)
	}
	value := strings.TrimSpace(string(out))
	if value == "" {
		return "", errors.New("find Go module cache: go env GOMODCACHE returned an empty path")
	}
	return value, nil
}
