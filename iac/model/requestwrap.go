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

import "fmt"

// collapseRequestWrappers removes the wrapper that a request of a top-level field holds and the
// resource response does not: Create and Get hold a list, and Update holds the same list in a
// one-property object. A wrapper that the unwrap lines name collapses before the request type is
// compared with the type in the response, as it does when the response holds it too. It replaces
// the request types in requests with the values inside the wrappers, and returns the wrapper to
// record on the response type once the types are merged (nil when no request has one), and the
// policy lines that decided.
//
// A request whose wrapper the response holds too is left alone: unwrap collapses it with the
// response, in every request.
func (p Policy) collapseRequestWrappers(resource, field string, get *Type, requests *[2]*Type) (*Wrapper, map[string]bool, error) {
	u := &unwrapper{policy: p, used: map[string]bool{}, collapsed: map[string]bool{}, kept: map[string]bool{}}
	var w *Wrapper
	for i, req := range requests {
		if req == nil || req.Kind != Object || req.Schema == "" || wrapperProblem(req) != "" {
			continue
		}
		if get != nil && get.Kind == Object && get.Schema == req.Schema {
			continue // the response holds the wrapper: unwrap handles it for every request
		}
		if !u.decide(resource, field, req.Schema) {
			continue
		}
		inner := req.Fields[0]
		attrs := inner.Attrs
		if w == nil {
			w = &Wrapper{Field: inner.Name}
		} else if w.Field != inner.Name {
			return nil, nil, unwrapError(fieldLocation(resource, field), "UNWRAP_COMPONENT_UNSUPPORTED",
				fmt.Sprintf("Create and Update hold the value in wrappers with the properties %s and %s.", w.Field, inner.Name),
				"Name the wrapper of one request only, or give both wrappers the same property.")
		}
		if i == 0 {
			w.CreateSchema, w.Create = req.Schema, &attrs
		} else {
			w.UpdateSchema, w.Update = req.Schema, &attrs
		}
		value := *inner.Type
		// The request types were shaped as sets while they still were the wrapper.
		p.asSet(resource, field, &value)
		requests[i] = &value
	}
	return w, u.used, nil
}
