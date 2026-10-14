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
	"errors"
	"fmt"
	"go/types"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"

	"github.com/coralogix/coralogix-management-sdk/iac/issue"
)

// Source is the module that the package of a Ref must come from. An empty Version or Dir is
// not checked.
type Source struct {
	Module, Version, Dir string
	// Remediation is the fix that the SDK_SOURCE_MISMATCH issue states.
	Remediation string
}

// Issues checks that each loaded package comes from the module that source returns for the
// first ref of that package, and that every ref exists in its package. It changes refs in
// place, so the caller writes the names that the SDK has: Want of a field becomes the value
// type when the SDK uses a value, and Name of a ByType method becomes the method it found.
func Issues(refs []Ref, loaded map[string]*packages.Package, source func(Ref) Source) issue.Report {
	var report issue.Report
	seenPackages := map[string]bool{}
	for _, ref := range refs {
		if seenPackages[ref.Pkg] {
			continue
		}
		seenPackages[ref.Pkg] = true
		pkg := loaded[ref.Pkg]
		if pkg == nil || pkg.Module == nil {
			continue
		}
		want := source(ref)
		if pkg.Module.Path != want.Module || want.Version != "" && pkg.Module.Version != want.Version || want.Dir != "" && !samePath(pkg.Module.Dir, want.Dir) {
			report = append(report, issue.Issue{Code: "SDK_SOURCE_MISMATCH", Location: ref.Pkg, Message: fmt.Sprintf("The loaded package comes from module %s %s at %s; expected %s %s at %s.", pkg.Module.Path, pkg.Module.Version, pkg.Module.Dir, want.Module, want.Version, want.Dir), Remediation: want.Remediation})
		}
	}
	for index := range refs {
		ref := &refs[index]
		if err := checkRef(ref, loaded[ref.Pkg]); err != nil {
			report = append(report, issue.Issue{Code: "SDK_SYMBOL_MISSING", Location: ref.Path, Message: fmt.Sprintf("SDK %s %s: %s", ref.Kind, ref.SDKName(), err), Remediation: "Regenerate and publish the SDK from the same OpenAPI contract, then pin that SDK version."})
		}
	}
	return report
}

func samePath(a, b string) bool {
	a, aErr := filepath.EvalSymlinks(a)
	b, bErr := filepath.EvalSymlinks(b)
	return aErr == nil && bErr == nil && filepath.Clean(a) == filepath.Clean(b)
}

// Load loads the packages that the refs use from the selected module graph.
func Load(refs []Ref, dir string) (map[string]*packages.Package, error) {
	var paths []string
	seen := map[string]bool{}
	for _, ref := range refs {
		if !seen[ref.Pkg] {
			seen[ref.Pkg] = true
			paths = append(paths, ref.Pkg)
		}
	}
	cfg := &packages.Config{
		Dir:  dir,
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedModule,
		Env:  offlineGoEnv(),
	}
	pkgs, err := packages.Load(cfg, paths...)
	if err != nil {
		return nil, fmt.Errorf("load SDK: %w", err)
	}
	byPath := map[string]*packages.Package{}
	for _, p := range pkgs {
		if len(p.Errors) > 0 {
			return nil, fmt.Errorf("load SDK package %s: %v", p.PkgPath, p.Errors[0])
		}
		byPath[p.PkgPath] = p
	}
	return byPath, nil
}

func offlineGoEnv() []string {
	env := make([]string, 0, len(os.Environ())+2)
	for _, item := range os.Environ() {
		if strings.HasPrefix(item, "GOPROXY=") || strings.HasPrefix(item, "GOSUMDB=") {
			continue
		}
		env = append(env, item)
	}
	return append(env, "GOPROXY=off", "GOSUMDB=off")
}

func checkRef(ref *Ref, pkg *packages.Package) error {
	if pkg == nil {
		return fmt.Errorf("package %s not loaded", ref.Pkg)
	}
	qualifier := func(p *types.Package) string {
		if p == pkg.Types {
			return ""
		}
		return p.Name()
	}
	if ref.Kind == KindField || ref.Kind == KindMethod {
		return checkMember(ref, pkg, qualifier)
	}
	return checkTopLevel(ref, pkg, qualifier)
}

// checkTopLevel checks a package, or a type, const, or func of the package.
func checkTopLevel(ref *Ref, pkg *packages.Package, qualifier types.Qualifier) error {
	scope := pkg.Types.Scope()
	switch ref.Kind {
	case KindPackage:
		if pkg.Name != ref.Name {
			return fmt.Errorf("package name is %s", pkg.Name)
		}
		return nil
	case KindType:
		if _, ok := scope.Lookup(ref.Name).(*types.TypeName); !ok {
			return errors.New("not found")
		}
		return nil
	case KindConst:
		if _, ok := scope.Lookup(ref.Name).(*types.Const); !ok {
			return errors.New("not found")
		}
		return nil
	case KindFunc:
		fn, ok := scope.Lookup(ref.Name).(*types.Func)
		if !ok {
			return errors.New("not found")
		}
		if s := types.TypeString(fn.Type(), qualifier); s != ref.Want {
			return fmt.Errorf("type is %s, want %s", s, ref.Want)
		}
		return nil
	}
	return fmt.Errorf("unknown kind %q", ref.Kind)
}

// checkMember checks a field or a method of the type ref.Owner.
func checkMember(ref *Ref, pkg *packages.Package, qualifier types.Qualifier) error {
	owner, ok := pkg.Types.Scope().Lookup(ref.Owner).(*types.TypeName)
	if !ok {
		return fmt.Errorf("type %s not found", ref.Owner)
	}
	var got types.Type
	if ref.Kind == KindField {
		st, ok := owner.Type().Underlying().(*types.Struct)
		if !ok {
			return fmt.Errorf("%s is not a struct", ref.Owner)
		}
		for f := range st.Fields() {
			if f.Name() == ref.Name {
				got = f.Type()
			}
		}
	} else if sel := types.NewMethodSet(types.NewPointer(owner.Type())).Lookup(pkg.Types, ref.Name); sel != nil {
		got = sel.Type()
	}
	if got == nil && ref.ByType {
		return findMethodByType(ref, owner, qualifier)
	}
	if got == nil {
		return errors.New("not found")
	}
	switch s := types.TypeString(got, qualifier); {
	case s == ref.Want:
	case ref.WantValue != "" && s == ref.WantValue:
		ref.Want = s
	default:
		return fmt.Errorf("type is %s, want %s", s, ref.Want)
	}
	return nil
}

// findMethodByType sets ref.Name to the one exported method of owner whose
// type is ref.Want.
func findMethodByType(ref *Ref, owner *types.TypeName, qualifier types.Qualifier) error {
	var found []string
	for sel := range types.NewMethodSet(types.NewPointer(owner.Type())).Methods() {
		if sel.Obj().Exported() && types.TypeString(sel.Type(), qualifier) == ref.Want {
			found = append(found, sel.Obj().Name())
		}
	}
	if len(found) != 1 {
		return fmt.Errorf("not found, and %d methods of %s have type %s: %v", len(found), ref.Owner, ref.Want, found)
	}
	ref.Name = found[0]
	return nil
}
